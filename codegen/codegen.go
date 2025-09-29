package codegen

import (
	"fmt"
	"log"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/sef-comp/truby-compiler/parser"
	gen "github.com/sef-comp/truby-compiler/parser" // сгенерированный парсер и лексер
)

type VarScope map[string]*ir.InstAlloca

func NewVisitor(debug bool) *MyRubyKVisitor {
	return &MyRubyKVisitor{Debug: debug}
}

type MyRubyKVisitor struct {
	*antlr.BaseParseTreeVisitor

	Debug bool

	Module      *ir.Module
	CurrentType types.Type
	Function    *ir.Func

	Main           *ir.Block
	Blocks         []*ir.Block
	VariableScopes []VarScope
}

func (v *MyRubyKVisitor) block() *ir.Block {
	return v.Blocks[len(v.Blocks)-1]
}

func (v *MyRubyKVisitor) pushBlock(block *ir.Block) {
	if len(v.Blocks) == 0 {
		v.Main = block
	} else {
		endWithBr(block, v.block())
	}
	v.Blocks = append(v.Blocks, block)
}

func (v *MyRubyKVisitor) popBlock(block *ir.Block) {
	v.Blocks = v.Blocks[:len(v.Blocks)-1]
	if len(v.Blocks) == 0 {
		v.Main = nil
	}
}

// returns replaced block
func (v *MyRubyKVisitor) replaceBlock(block *ir.Block) *ir.Block {
	curBlock := v.block()
	if brTerm, ok := curBlock.Term.(*ir.TermBr); ok {
		block.NewBr(brTerm.Succs()[0])
	}
	v.Blocks[len(v.Blocks)-1] = block
	return curBlock
}

func endWithBr(block, nextBlock *ir.Block) {
	if block.Term == nil {
		block.NewBr(nextBlock)
	}
}

func (v *MyRubyKVisitor) pushScope() {
	v.VariableScopes = append(v.VariableScopes, VarScope{})
}
func (v *MyRubyKVisitor) popScope() {
	v.VariableScopes = v.VariableScopes[:len(v.VariableScopes)-1]
}

func (v *MyRubyKVisitor) newVariable(t types.Type, name string) *ir.InstAlloca {
	_, exists := v.variable(name)
	if exists {
		panic("already defined")
	}

	ptr := v.Main.NewAlloca(t)

	// assing new variable to current variable scope
	v.VariableScopes[len(v.VariableScopes)-1][name] = ptr

	return ptr
}

func (v *MyRubyKVisitor) identifier(name string) value.Named {
	if id, _ := v.variable(name); id != nil {
		return id
	}
	if id := v.argumentId(name); id != nil {
		return id
	}
	if id := v.funcId(name); id != nil {
		return id
	}
	return nil
}

func (v *MyRubyKVisitor) funcId(name string) *ir.Func {
	for _, function := range v.Module.Funcs {
		if function.Name() == name {
			return function
		}
	}
	return nil
}

func (v *MyRubyKVisitor) argumentId(name string) *ir.Param {
	for _, param := range v.Function.Params {
		if param.Name() == name {
			return param
		}
	}
	return nil
}

func (v *MyRubyKVisitor) variable(name string) (*ir.InstAlloca, bool) {
	for i := len(v.VariableScopes) - 1; i >= 0; i = i - 1 {
		scope := v.VariableScopes[i]
		if variable, ok := scope[name]; ok {
			return variable, true
		}
	}
	return nil, false
}

func (v *MyRubyKVisitor) dereference(expr value.Value) value.Value {
	if !types.IsPointer(expr.Type()) {
		return expr
	}

	if ptr, ok := expr.(*ir.InstAlloca); ok {
		return v.block().NewLoad(ptr.ElemType, ptr)
	}
	if ptr, ok := expr.(*ir.InstGetElementPtr); ok {
		arrType := baseArrType(ptr)

		arrElem := v.block().NewLoad(arrType.ElemType, ptr)
		return v.dereference(arrElem)
	}
	return expr
}

func baseArrType(ptr *ir.InstGetElementPtr) *types.ArrayType {
	arrType := ptr.ElemType.(*types.ArrayType)

	for {
		if elem, ok := arrType.ElemType.(*types.ArrayType); ok {
			arrType = elem
		} else {
			break
		}
	}
	return arrType
}

func (v *MyRubyKVisitor) sameT(expr, nextExpr value.Value) (value.Value, value.Value) {
	// TODO: check size compatibility
	return v.dereference(expr), v.dereference(nextExpr)
}

func (v *MyRubyKVisitor) castCond(cond value.Value) value.Value {
	cond = v.dereference(cond)

	if cond.Type().Equal(types.I1) {
		return cond
	}

	if cond.Type().Equal(types.I32) {
		return v.block().NewICmp(enum.IPredNE, cond, constant.NewInt(types.I32, 0))
	}

	panic("Bad Condition")
}

// func (v *MyRubyKVisitor) Visit(ctx *gen.Pro)

func (v *MyRubyKVisitor) VisitProgram(ctx *gen.ProgramContext) interface{} {
	if v.Debug {
		log.Println("Program visited", ctx.GetText())
	}
	v.Module = ir.NewModule()

	v.Function = v.Module.NewFunc("main", types.I32, []*ir.Param{}...) // mainFunc := v.Module.NewFunc("main", types.I32, nil)
	block := v.Function.NewBlock("entry")
	v.Main = block
	v.pushBlock(block)
	v.pushScope()

	expr_list := ctx.Expression_list()
	v.VisitExpression_list(expr_list.(*gen.Expression_listContext))
	return nil
	// return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitExpression_list(ctx *gen.Expression_listContext) interface{} {
	// expression_list
	//     : expression terminator
	//     | expression_list expression terminator
	//     | terminator
	//     ;
	if v.Debug {
		log.Println("Expression_list visited", ctx.GetText())
	}

	if ctx_exprlist := ctx.Expression_list(); ctx_exprlist != nil {
		v.VisitExpression_list(ctx_exprlist.(*gen.Expression_listContext))
	}
	if ctx_expr := ctx.Expression(); ctx_expr != nil {
		v.VisitExpression(ctx_expr.(*gen.ExpressionContext))
	}
	if ctx_term := ctx.Terminator(); ctx_term != nil {
		v.VisitTerminator(ctx_term.(*gen.TerminatorContext))
		return nil
	} else {
		panic("expression_list error")
	}
}

func (v *MyRubyKVisitor) VisitExpression(ctx *gen.ExpressionContext) interface{} {
	// expression
	//     : function_definition
	//     | function_inline_call
	//     | if_statement
	//     | rvalue
	//     | return_statement
	//     | while_statement
	//     | for_statement
	//     ;
	if v.Debug {
		log.Println("Expression visited", ctx.GetText())
	}
	// В зависимости от типа выражения, вызовем соответствующий визит
	if ctx.Function_definition() != nil {
		return v.VisitFunction_definition(ctx.Function_definition().(*gen.Function_definitionContext))
	}
	if ctx.Function_inline_call() != nil {
		return v.VisitFunction_inline_call(ctx.Function_inline_call().(*gen.Function_inline_callContext))
	}
	if ctx.If_statement() != nil {
		return v.VisitIf_statement(ctx.If_statement().(*gen.If_statementContext))
	}
	if ctx.Rvalue() != nil {
		return v.VisitRvalue(ctx.Rvalue().(*gen.RvalueContext))
	}
	if ctx.Return_statement() != nil {
		return v.VisitReturn_statement(ctx.Return_statement().(*gen.Return_statementContext))
	}
	if ctx.While_statement() != nil {
		return v.VisitWhile_statement(ctx.While_statement().(*gen.While_statementContext))
	}
	if ctx.For_statement() != nil {
		return v.VisitFor_statement(ctx.For_statement().(*gen.For_statementContext))
	}
	return nil
}

func (v *MyRubyKVisitor) VisitGlobal_get(ctx *gen.Global_getContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitGlobal_set(ctx *gen.Global_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitGlobal_result(ctx *gen.Global_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFunction_inline_call(ctx *gen.Function_inline_callContext) interface{} {
	// function_inline_call
	//     : function_call
	//     ;

	return v.VisitFunction_call(ctx.Function_call().(*gen.Function_callContext))
}

func (v *MyRubyKVisitor) VisitFunction_definition(ctx *gen.Function_definitionContext) interface{} {
	// function_definition
	//     : function_definition_header function_definition_body END
	//     ;

	if v.Debug {
		log.Println("Function_definition visited", ctx.GetText())
	}

	block := v.VisitFunction_definition_header(ctx.Function_definition_header().(*gen.Function_definition_headerContext))
	v.VisitFunction_definition_body(ctx.Function_definition_body().(*gen.Function_definition_bodyContext))

	v.popBlock(block)
	v.Function = nil

	return nil
}

func (v *MyRubyKVisitor) VisitFunction_definition_body(ctx *gen.Function_definition_bodyContext) interface{} {
	// function_definition_body
	//     : expression_list
	//     ;

	return v.VisitExpression_list(ctx.Expression_list().(*gen.Expression_listContext))
}

func (v *MyRubyKVisitor) VisitFunction_definition_header(ctx *gen.Function_definition_headerContext) *ir.Block {
	// function_definition_header
	//     : DEF function_name crlf
	//     | DEF function_name function_definition_params crlf
	//     ;

	funcName := ctx.Function_name().GetText()

	// Определение типа функции (например, возвращаемое значение - int, параметры - пусто)
	// funcType := types.NewFunc(types.I32)

	var params = []*ir.Param{}
	if ctx.Function_definition_params() != nil {
		params = v.VisitFunction_definition_params(ctx.Function_definition_params().(*gen.Function_definition_paramsContext))
	}

	v.Function = v.Module.NewFunc(funcName, types.I32, params...)

	blockNum := len(v.block().Parent.Blocks)
	block := v.Function.NewBlock(fmt.Sprintf("fblock-%d", blockNum))

	v.pushBlock(block)

	for i, param := range params {
		alloca := v.block().NewAlloca(types.I32)
		v.block().NewStore(v.Function.Params[i], alloca)
		v.newVariable(param.Type(), param.Name())
	}

	return block
}

func (v *MyRubyKVisitor) VisitFunction_name(ctx *gen.Function_nameContext) interface{} {
	// function_name
	//     : id_function
	//     | id_
	//     ;

	if ctx.Id_function() != nil {
		return v.VisitId_function(ctx.Id_function().(*gen.Id_functionContext))
	}
	if ctx.Id_() != nil {
		return v.VisitId_(ctx.Id_().(*gen.Id_Context))
	}

	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFunction_definition_params(ctx *gen.Function_definition_paramsContext) []*ir.Param {
	// function_definition_params
	//     : LEFT_RBRACKET RIGHT_RBRACKET
	//     | LEFT_RBRACKET function_definition_params_list RIGHT_RBRACKET
	//     | function_definition_params_list

	if ctx.Function_definition_params_list() == nil {
		return []*ir.Param{}
	}

	return v.VisitFunction_definition_params_list(ctx.Function_definition_params_list().(*gen.Function_definition_params_listContext))
}

func (v *MyRubyKVisitor) VisitFunction_definition_params_list(ctx *gen.Function_definition_params_listContext) []*ir.Param {
	// function_definition_params_list
	//     : function_definition_param_id
	//     | function_definition_params_list COMMA function_definition_param_id
	//     ;

	var params = []*ir.Param{}

	if ctx.Function_definition_params_list() != nil {
		params = v.VisitFunction_definition_params_list(ctx.Function_definition_params_list().(*gen.Function_definition_params_listContext))
	}

	params = append(params, v.VisitFunction_definition_param_id(ctx.Function_definition_param_id().(*gen.Function_definition_param_idContext)))
	return params
}

func (v *MyRubyKVisitor) VisitFunction_definition_param_id(ctx *gen.Function_definition_param_idContext) *ir.Param {
	// function_definition_param_id
	//     : id_
	//     ;

	param_name := v.VisitId_(ctx.Id_().(*gen.Id_Context))
	return ir.NewParam(param_name, types.I32)
}

func (v *MyRubyKVisitor) VisitReturn_statement(ctx *gen.Return_statementContext) interface{} {
	if v.Debug {
		log.Println("Return_statement visited", ctx.GetText())
	}

	ctx_node_all_res := ctx.All_result()
	retValue := v.VisitAll_result(ctx_node_all_res.(*gen.All_resultContext))
	v.block().NewRet(v.dereference(retValue))

	return nil
}

func (v *MyRubyKVisitor) VisitFunction_call(ctx *gen.Function_callContext) value.Value {
	// function_call
	//     : name = function_name LEFT_RBRACKET params = function_call_param_list RIGHT_RBRACKET
	//     | name = function_name params = function_call_param_list
	//     | name = function_name LEFT_RBRACKET RIGHT_RBRACKET
	//     ;

	fname := v.VisitFunction_name(ctx.Function_name().(*gen.Function_nameContext))

	var args = []value.Value{}

	if ctx.Function_call_param_list() != nil {
		args = v.VisitFunction_call_param_list(ctx.Function_call_param_list().(*gen.Function_call_param_listContext))
	}

	fn := v.funcId(fname.(string))

	result := v.block().NewCall(fn, args...)

	if v.Debug {
		log.Println(result.Type())
	}

	if result == nil {
		panic("empty res")
	}
	return result
}

func (v *MyRubyKVisitor) VisitFunction_call_param_list(ctx *gen.Function_call_param_listContext) []value.Value {
	// function_call_param_list
	//     : function_call_params
	//     ;

	return v.VisitFunction_call_params(ctx.Function_call_params().(*gen.Function_call_paramsContext))
}

func (v *MyRubyKVisitor) VisitFunction_call_params(ctx *gen.Function_call_paramsContext) []value.Value {
	// function_call_params
	//     : function_param
	//     | function_call_params COMMA function_param
	//     ;

	var params = []value.Value{}
	if ctx.Function_call_params() != nil {
		params = v.VisitFunction_call_params(ctx.Function_call_params().(*gen.Function_call_paramsContext))
	}
	params = append(params, v.VisitFunction_param(ctx.Function_param().(*gen.Function_paramContext)))

	return params

}

func (v *MyRubyKVisitor) VisitFunction_param(ctx *gen.Function_paramContext) value.Value {
	// function_param
	//     : (function_unnamed_param | function_named_param)
	//     ;

	if ctx.Function_unnamed_param() != nil {
		ret := v.VisitFunction_unnamed_param(ctx.Function_unnamed_param().(*gen.Function_unnamed_paramContext))
		return ret
	}
	if ctx.Function_named_param() != nil {
		_, ret := v.VisitFunction_named_param(ctx.Function_named_param().(*gen.Function_named_paramContext))
		return ret
	}
	panic("argument exception")
}

func (v *MyRubyKVisitor) VisitFunction_unnamed_param(ctx *gen.Function_unnamed_paramContext) value.Value {
	// function_unnamed_param
	//     : (int_result | float_result | string_result | dynamic_result)
	//     ;

	switch true {
	case ctx.Int_result() != nil:
		return v.VisitInt_result(ctx.Int_result().(*gen.Int_resultContext))
	case ctx.Dynamic_result() != nil:
		return v.VisitDynamic_result(ctx.Dynamic_result().(*gen.Dynamic_resultContext))
	default:
		panic("not implemented")
	}
}

func (v *MyRubyKVisitor) VisitFunction_named_param(ctx *gen.Function_named_paramContext) (*ir.Param, value.Value) {
	// function_named_param
	//     : id_ op = ASSIGN (int_result | float_result | string_result | dynamic_result)
	//     ;

	paramName := v.VisitId_(ctx.Id_().(*gen.Id_Context))

	switch true {
	case ctx.Int_result() != nil:
		return ir.NewParam(paramName, types.I32), v.VisitInt_result(ctx.Int_result().(*gen.Int_resultContext))
	case ctx.Float_result() != nil:
		// return ir.NewParam(paramName, types.FP128), v.VisitFloat_result(ctx.Float_result().(*gen.Float_resultContext))
	// case ctx.String_result() != nil:
	// return ir.NewParam(paramName, type)
	case ctx.Dynamic_result() != nil:
		res := v.VisitDynamic_result(ctx.Dynamic_result().(*gen.Dynamic_resultContext))
		return ir.NewParam(paramName, res.Type()), res
	}

	panic("wrong arguments call")
}

func (v *MyRubyKVisitor) VisitFunction_call_assignment(ctx *gen.Function_call_assignmentContext) value.Value {
	// function_call_assignment
	//     : function_call
	//     ;

	return v.VisitFunction_call(ctx.Function_call().(*gen.Function_callContext))
}

func (v *MyRubyKVisitor) VisitAll_result(ctx *gen.All_resultContext) value.Value {
	if v.Debug {
		log.Println("All_result visited", ctx.GetText())
	}
	var res value.Value

	if ctx_node := ctx.Dynamic_result(); ctx_node != nil {
		res = v.VisitDynamic_result(ctx_node.(*gen.Dynamic_resultContext))
		return res
	}
	if ctx_node := ctx.Int_result(); ctx_node != nil {
		res = v.VisitInt_result(ctx_node.(*gen.Int_resultContext))
		return res
	}

	panic("not inplemented")
	// return res
}

func (v *MyRubyKVisitor) VisitElsif_statement(ctx *gen.Elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitIf_elsif_statement(ctx *gen.If_elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitIf_statement(ctx *gen.If_statementContext) interface{} {
	// if_statement
	// : IF cond_expression crlf statement_body END
	// | IF cond_expression crlf statement_body else_token crlf statement_body END
	// | IF cond_expression crlf statement_body elsif_statement END
	// ;

	condCtx := ctx.Cond_expression()
	cond := v.VisitCond_expression(condCtx.(*parser.Cond_expressionContext))

	blockNum := len(v.block().Parent.Blocks)
	ifBlock := v.block().Parent.NewBlock(fmt.Sprintf("if-%d", blockNum))
	elseBlock := v.block().Parent.NewBlock(fmt.Sprintf("else-%d", blockNum))
	nextBlock := v.block().Parent.NewBlock(fmt.Sprintf("main-%d", blockNum))

	preBlock := v.replaceBlock(nextBlock)
	preBlock.NewCondBr(v.castCond(cond), ifBlock, elseBlock)
	ifBlock.NewBr(nextBlock)
	elseBlock.NewBr(nextBlock)

	{
		// if
		v.pushBlock(ifBlock)
		stCtx := ctx.AllStatement_body()[0]
		v.VisitStatement_body(stCtx.(*gen.Statement_bodyContext))
		v.popBlock(ifBlock)
	}

	if len(ctx.AllStatement_body()) > 1 {
		// else (optional)
		v.pushBlock(elseBlock)
		stCtx := ctx.AllStatement_body()[1]
		v.VisitStatement_body(stCtx.(*parser.Statement_bodyContext))
		v.popBlock(elseBlock)
	}

	return nil
}

func (v *MyRubyKVisitor) VisitUnless_statement(ctx *gen.Unless_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitWhile_statement(ctx *gen.While_statementContext) interface{} {
	// while_statement
	//     : WHILE cond_expression crlf statement_body END
	//     ;
	if v.Debug {
		log.Println("While_statement visited", ctx.GetText())
	}
	blockNum := len(v.block().Parent.Blocks)

	condBlock := v.block().Parent.NewBlock(fmt.Sprintf("while.cond-%d", blockNum))
	bodyBlock := v.block().Parent.NewBlock(fmt.Sprintf("while.body-%d", blockNum))
	nextBlock := v.block().Parent.NewBlock(fmt.Sprintf("main-%d", blockNum))

	preBlock := v.replaceBlock(nextBlock)
	preBlock.NewBr(condBlock)
	bodyBlock.NewBr(condBlock)

	// {
	// condition
	v.pushBlock(condBlock)
	condCtx := ctx.Cond_expression()
	cond := v.VisitCond_expression(condCtx.(*gen.Cond_expressionContext))
	condBlock.NewCondBr(v.castCond(cond), bodyBlock, nextBlock)
	v.popBlock(condBlock)
	// }

	// {
	// statement
	v.pushBlock(bodyBlock)
	stmtCtx := ctx.Statement_body()
	v.VisitStatement_body(stmtCtx.(*parser.Statement_bodyContext))
	v.popBlock(bodyBlock)
	// }
	return nil
}

func (v *MyRubyKVisitor) VisitFor_statement(ctx *gen.For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitInit_expression(ctx *gen.Init_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitAll_assignment(ctx *gen.All_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFor_init_list(ctx *gen.For_init_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitCond_expression(ctx *gen.Cond_expressionContext) value.Value {
	// cond_expression
	//
	//	: comparison_list
	//	;
	compctx := ctx.Comparison_list()
	val := v.VisitComparison_list(compctx.(*gen.Comparison_listContext))
	return val
}

func (v *MyRubyKVisitor) VisitLoop_expression(ctx *gen.Loop_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFor_loop_list(ctx *gen.For_loop_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitStatement_body(ctx *gen.Statement_bodyContext) interface{} {
	// statement_body
	//     : statement_expression_list
	//     ;

	v.VisitStatement_expression_list(ctx.Statement_expression_list().(*gen.Statement_expression_listContext))
	return nil
}

func (v *MyRubyKVisitor) VisitStatement_expression_list(ctx *gen.Statement_expression_listContext) interface{} {
	// statement_expression_list
	//     : expression terminator
	//     | RETRY terminator
	//     | break_expression terminator
	//     | statement_expression_list expression terminator
	//     | statement_expression_list RETRY terminator
	//     | statement_expression_list break_expression terminator
	//     ;

	if ctx.Statement_expression_list() != nil {
		v.VisitStatement_expression_list(ctx.Statement_expression_list().(*gen.Statement_expression_listContext))
	}

	if ctx.Expression() != nil {
		return v.VisitExpression(ctx.Expression().(*gen.ExpressionContext))
	}

	if ctx.Break_expression() != nil {
		return v.VisitBreak_expression(ctx.Break_expression().(*gen.Break_expressionContext))
	}

	panic("not implemented")
}

func (v *MyRubyKVisitor) VisitAssignment(ctx *gen.AssignmentContext) interface{} {
	if v.Debug {
		log.Println("VisitAssignment ", ctx.GetText())
	}

	if ctx_lval := ctx.Lvalue(); ctx_lval != nil {
		v.VisitLvalue(ctx_lval.(*gen.LvalueContext))
	}
	if ctx_rval := ctx.Rvalue(); ctx_rval != nil {
		v.VisitRvalue(ctx_rval.(*gen.RvalueContext))
	}
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitDynamic_assignment(ctx *gen.Dynamic_assignmentContext) interface{} {
	if v.Debug {
		log.Println("dynamic_assignment visited", ctx.GetText())
	}

	lval := ctx.Lvalue()
	idname := v.VisitLvalue(lval.(*gen.LvalueContext)).(string)

	vr, exists := v.variable(idname)
	if !exists {
		panic(fmt.Sprintf("variable %s not defined", idname))
	}
	// log.Println("[DEBUG] vr = ", vr)

	ctx_res := ctx.Dynamic_result()
	val := v.VisitDynamic_result(ctx_res.(*gen.Dynamic_resultContext))

	if v.Debug {
		log.Println(val.Type())
	}

	v.block().NewStore(val, vr)

	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitInt_assignment(ctx *gen.Int_assignmentContext) interface{} {
	// int_assignment
	//     : var_id = lvalue op = ASSIGN int_result
	//     | var_id = lvalue op = (
	//         PLUS_ASSIGN
	//         | MINUS_ASSIGN
	//         | MUL_ASSIGN
	//         | DIV_ASSIGN
	//         | MOD_ASSIGN
	//         | EXP_ASSIGN
	//     ) int_result
	//     ;
	if v.Debug {
		log.Println("int_assignment visited", ctx.GetText())
	}
	// ctx_varid := ctx.Lvalue()
	// name := v.VisitLvalue(ctx_varid.(*gen.LvalueContext)).(string)
	// var ptr *ir.InstAlloca
	// if ptr = v.variable(name); ptr == nil {
	// 	ptr = v.newVariable(types.I32, name)
	// }
	//
	// // log.Println(ptr)
	//
	// ctx_intres := ctx.Int_result()
	// val := v.VisitInt_result(ctx_intres.(*gen.Int_resultContext))
	//
	// log.Println(val)
	// v.block().NewStore(val, ptr)

	// log.Println("lval = ", ctx.Lvalue().GetText())
	// log.Println("intRes = ", ctx.Int_result().GetText()

	ctx_varid := ctx.Lvalue()
	name := v.VisitLvalue(ctx_varid.(*gen.LvalueContext)).(string)

	// Получаем значение
	ctx_intres := ctx.Int_result()
	val := v.VisitInt_result(ctx_intres.(*gen.Int_resultContext))

	// Проверяем, существует ли переменная в текущей области видимости
	var alloc *ir.InstAlloca
	var exists bool
	// if alloc, exists = v.VariableScopes[len(v.VariableScopes)-1][name]; !exists {
	if alloc, exists = v.variable(name); exists == false {
		// Если переменная не существует, выделяем память (alloca) и сохраняем в текущей области видимости
		alloc = v.block().NewAlloca(types.I32)
		v.VariableScopes[len(v.VariableScopes)-1][name] = alloc
	}

	// Присваиваем значение переменной (store)
	v.Main.NewStore(val, alloc)

	return nil

	// return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFloat_assignment(ctx *gen.Float_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitString_assignment(ctx *gen.String_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitInitial_array_assignment(ctx *gen.Initial_array_assignmentContext) interface{} {
	// initial_array_assignment
	//     : var_id = lvalue op = ASSIGN LEFT_SBRACKET RIGHT_SBRACKET
	//     ;
	if v.Debug {
		log.Println("initial_array_assignment visited", ctx.GetText())
	}

	varName := ctx.Lvalue().GetText()
	elementType := types.I32

	// TODO Переделать на кучу
	// mallocFunc := v.Module.NewFunc("malloc", types.NewPointer(elementType), types.I32)
	// arrayAlloc := v.Main.NewCall(mallocFunc, arraySize)

	// arrayAlloc := v.block().NewAlloca(types.NewArray(32, elementType))
	// array := v.dereference(arrayAlloc)

	if v.Debug {
		// log.Println("array type: ", array.Type())
	}

	// v.VariableScopes[len(v.VariableScopes)-1][varName] = arrayAlloc
	_ = v.newVariable(types.NewArray(32, elementType), varName)

	// ptr, _ := v.variable(varName);

	//  log.Println(ptr.Type())
	//  arr := v.dereference(ptr)
	//  log.Println(arr.Type())
	//  // log.Println(arrayAlloc.Type())
	//  arrayPtr := v.block().NewGetElementPtr(
	// 	ptr.ElemType, // Тип элементов
	// 	ptr,          // Указатель на массив
	// 	constant.NewInt(types.I32, 0), // Базовый указатель
	// 	constant.NewInt(types.I32, 0), // Индекс первого элемента
	// )
	//
	//  log.Println(arrayPtr)
	//
	// // Пример инициализации первого элемента
	// v.block().NewStore(constant.NewInt(types.I32, 42), arrayPtr)
	//
	//  log.Println("ok")

	return nil
}

func (v *MyRubyKVisitor) VisitArray_assignment(ctx *gen.Array_assignmentContext) interface{} {
	// array_assignment
	//     : arr_def = array_selector op = ASSIGN arr_val = all_result
	//     ;

	if v.Debug {
		log.Println("array_assignment ", ctx.GetText())
	}

	ctx_sel := ctx.Array_selector()
	elem := v.VisitArray_selector(ctx_sel.(*gen.Array_selectorContext))

	if v.Debug {
		log.Println("ElemType: ", elem.Type())
	}

	ctx_res := ctx.All_result()
	res := v.VisitAll_result(ctx_res.(*gen.All_resultContext))
	val := v.dereference(res)

	v.block().NewStore(val, elem)
	return nil
}

func (v *MyRubyKVisitor) VisitArray_definition(ctx *gen.Array_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitArray_definition_elements(ctx *gen.Array_definition_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitArray_selector(ctx *gen.Array_selectorContext) value.Value {
	// id_ LEFT_SBRACKET (int_result | dynamic_result) RIGHT_SBRACKET

	if v.Debug {
		log.Println("Array_selector visited", ctx.GetText())
	}

	name := ctx.Id_().GetText()

	var index_val value.Value
	if ctx_int := ctx.Int_result(); ctx_int != nil {
		index_val = v.VisitInt_result(ctx_int.(*gen.Int_resultContext))
	} else if ctx_res := ctx.Dynamic_result(); ctx_res != nil {
		index_val = v.VisitDynamic_result(ctx_res.(*gen.Dynamic_resultContext))
	}

	index := v.dereference(index_val)
	if v.Debug {
		log.Println(index.Type())
	}
	if index.Type() != types.I32 {
		panic(fmt.Sprintf("cannot index with %s type", index.Type()))
	}

	arrayPtr, exists := v.variable(name)
	if !exists {
		panic(fmt.Sprintf("Array %s not found", name))
	}

	elementPtr := v.block().NewGetElementPtr(
		arrayPtr.ElemType,             // Тип элемента массива
		arrayPtr,                      // Указатель на массив
		constant.NewInt(types.I32, 0), // Индекс 0 для доступа к самому массиву
		index,                         // Индекс элемента в массиве
	)

	if v.Debug {
		log.Println("elemPtr type: ", elementPtr.Type())
	}
	// Загружаем значение элемента массива
	// elementValue := v.block().NewLoad(arrayPtr.ElemType, elementPtr)
	return elementPtr
}

func (v *MyRubyKVisitor) VisitDynamic_result(ctx *gen.Dynamic_resultContext) value.Value {
	// dynamic_result
	//     : dynamic_result op = (MUL | DIV | MOD) int_result
	//     | int_result op = ( MUL | DIV | MOD) dynamic_result
	//     | dynamic_result op = ( MUL | DIV | MOD) float_result
	//     | float_result op = ( MUL | DIV | MOD) dynamic_result
	//     | dynamic_result op = ( MUL | DIV | MOD) dynamic_result
	//     | dynamic_result op = MUL string_result
	//     | string_result op = MUL dynamic_result
	//     | dynamic_result op = ( PLUS | MINUS) int_result
	//     | int_result op = ( PLUS | MINUS) dynamic_result
	//     | dynamic_result op = ( PLUS | MINUS) float_result
	//     | float_result op = ( PLUS | MINUS) dynamic_result
	//     | dynamic_result op = ( PLUS | MINUS) dynamic_result
	//     | LEFT_RBRACKET dynamic_result RIGHT_RBRACKET
	//     | dynamic_
	//     ;
	if v.Debug {
		log.Println("Dynamic_result visited", ctx.GetText())
	}
	if ctx.Dynamic_() != nil {
		return v.VisitDynamic_(ctx.Dynamic_().(*gen.Dynamic_Context))
	}

	var left, right value.Value
	var op antlr.Token

	// Определяем левый операнд
	if ctx.Dynamic_result(0) != nil {
		left = v.VisitDynamic_result(ctx.Dynamic_result(0).(*gen.Dynamic_resultContext))
	} else if ctx.Int_result() != nil {
		left = v.VisitInt_result(ctx.Int_result().(*gen.Int_resultContext))
	} else if ctx.Float_result() != nil {
		// left = v.VisitFloat_result(ctx.Float_result().(*gen.Float_resultContext))
	} else if ctx.String_result() != nil {
		// left = v.VisitString_result(ctx.String_result().(*gen.String_resultContext))
	}

	// Проверяем, есть ли оператор и правый операнд
	if ctx.GetOp() != nil {
		op = ctx.GetOp()

		// Определяем правый операнд
		if ctx.Dynamic_result(1) != nil {
			right = v.VisitDynamic_result(ctx.Dynamic_result(1).(*gen.Dynamic_resultContext))
		} else if ctx.Int_result() != nil {
			right = v.VisitInt_result(ctx.Int_result().(*gen.Int_resultContext))
		} else if ctx.Float_result() != nil {
			// right = v.VisitFloat_result(ctx.Float_result().(*gen.Float_resultContext))
		}
		// log.Println("left", left)
		// log.Println("right", right)

		// В зависимости от оператора, выполняем нужную операцию
		leftval, rightval := v.dereference(left), v.dereference(right)
		switch op.GetTokenType() {
		case gen.RubyKParserPLUS:
			// log.Println(leftval.Type(), rightval.Type())
			if leftval.Type().Equal(types.I32) && rightval.Type().Equal(types.I32) {
				// log.Print("ok")
				return v.block().NewAdd(leftval, rightval)
			} else if leftval.Type().Equal(types.Double) && rightval.Type().Equal(types.Double) {
				// return v.block().NewFAdd(left, right)
			}
		case gen.RubyKParserMINUS:
			if leftval.Type().Equal(types.I32) && rightval.Type().Equal(types.I32) {
				return v.block().NewSub(leftval, rightval)
			} else if leftval.Type().Equal(types.Double) && rightval.Type().Equal(types.Double) {
				// return v.Function.NewFSub(left, right)
			}
		case gen.RubyKParserMUL:
			if leftval.Type().Equal(types.I32) && rightval.Type().Equal(types.I32) {
				return v.block().NewMul(leftval, rightval)
			} else if leftval.Type().Equal(types.Double) && rightval.Type().Equal(types.Double) {
				// return v.Function.NewFMul(left, right)
			}
		case gen.RubyKParserDIV:
			if leftval.Type().Equal(types.I32) && rightval.Type().Equal(types.I32) {
				return v.block().NewSDiv(leftval, rightval)
			} else if left.Type().Equal(types.Double) && right.Type().Equal(types.Double) {
				// return v.Function.NewFDiv(left, right)
			}
		case gen.RubyKParserMOD:
			if leftval.Type().Equal(types.I32) && rightval.Type().Equal(types.I32) {
				return v.block().NewSRem(leftval, rightval)
			}
		}
	}

	// Если нет оператора, возвращаем левый операнд
	return left
}

func (v *MyRubyKVisitor) VisitDynamic_(ctx *gen.Dynamic_Context) value.Value {
	// dynamic_
	//     : id_
	//     | function_call_assignment
	//     | array_selector
	//     ;

	if v.Debug {
		log.Println("Dynamic_ visited", ctx.GetText())
	}

	if ctx_id := ctx.Id_(); ctx_id != nil {
		name := v.VisitId_(ctx_id.(*gen.Id_Context))

		vr, exists := v.variable(name)
		if !exists {
			panic(fmt.Sprintf("variable %s not defined", name))
		}
		return v.dereference(vr) //////////////////////////////////////////////////////////////////////
	}

	if ctx_arr := ctx.Array_selector(); ctx_arr != nil {
		elem_ptr := v.VisitArray_selector(ctx_arr.(*gen.Array_selectorContext))
		elem := v.dereference(elem_ptr)
		return elem
	}

	if ctx_f := ctx.Function_call_assignment(); ctx_f != nil {
		return v.VisitFunction_call_assignment(ctx.Function_call_assignment().(*gen.Function_call_assignmentContext))
	}

	panic("not implemented")
}

func (v *MyRubyKVisitor) VisitInt_result(ctx *gen.Int_resultContext) value.Value {
	// int_result
	//     : int_result op = (MUL | DIV | MOD) int_result
	//     | int_result op = ( PLUS | MINUS) int_result
	//     | LEFT_RBRACKET int_result RIGHT_RBRACKET
	//     | int_t
	if v.Debug {
		log.Println("Int_result visited", ctx.GetText())
	}

	if ctx.GetToken(gen.RubyKLexerLEFT_RBRACKET, 0) != nil && ctx.GetToken(gen.RubyKLexerRIGHT_RBRACKET, 0) != nil {
		return v.VisitInt_result(ctx.Int_result(0).(*gen.Int_resultContext))
	}

	if ctx_int := ctx.Int_t(); ctx_int != nil {
		return v.VisitInt_t(ctx_int.(*gen.Int_tContext))
	}

	if ctx.GetOp() != nil {
		op := ctx.GetOp()

		leftval := v.VisitInt_result(ctx.Int_result(0).(*gen.Int_resultContext))
		rightval := v.VisitInt_result(ctx.Int_result(1).(*gen.Int_resultContext))

		left, right := v.dereference(leftval), v.dereference(rightval)
		switch op.GetTokenType() {
		case gen.RubyKParserPLUS:
			return v.block().NewAdd(left, right)
		case gen.RubyKParserMINUS:
			return v.block().NewSub(left, right)
		case gen.RubyKParserMUL:
			return v.block().NewMul(left, right)
		case gen.RubyKParserDIV:
			return v.block().NewSDiv(left, right)
		case gen.RubyKParserMOD:
			return v.block().NewSRem(left, right)
		}
	}
	panic("not implemented")
}

func (v *MyRubyKVisitor) VisitFloat_result(ctx *gen.Float_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitString_result(ctx *gen.String_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitComparison_list(ctx *gen.Comparison_listContext) value.Value {
	// comparison_list
	//     : left = comparison op = BIT_AND right = comparison_list
	//     | left = comparison op = AND right = comparison_list
	//     | left = comparison op = BIT_OR right = comparison_list
	//     | left = comparison op = OR right = comparison_list
	//     | LEFT_RBRACKET comparison_list RIGHT_RBRACKET
	//     | comparison
	//     ;

	if ctx.GetOp() != nil {

		op := ctx.GetOp()

		rightval := v.VisitComparison_list(ctx.Comparison_list().(*gen.Comparison_listContext))
		leftval := v.VisitComparison(ctx.Comparison().(*gen.ComparisonContext))

		left, right := v.dereference(leftval), v.dereference(rightval)

		switch op.GetTokenType() {
		case gen.RubyKLexerBIT_AND:
			return v.block().NewAnd(left, right)
		case gen.RubyKLexerAND:
			return v.block().NewAnd(left, right)
		case gen.RubyKLexerBIT_OR:
			return v.block().NewOr(left, right)
		case gen.RubyKLexerOR:
			return v.block().NewOr(left, right)
		}

	} else {
		return v.VisitComparison(ctx.Comparison().(*gen.ComparisonContext))
	}

	panic("not implemented")
}

func (v *MyRubyKVisitor) VisitComparison(ctx *gen.ComparisonContext) value.Value {
	// comparison
	//     : left = comp_var op = (LESS | GREATER | LESS_EQUAL | GREATER_EQUAL) right = comp_var
	//     | left = comp_var op = ( EQUAL | NOT_EQUAL) right = comp_var
	//     ;

	// Получаем левую и правую части сравнения
	leftval := v.VisitComp_var(ctx.GetLeft().(*gen.Comp_varContext))
	rightval := v.VisitComp_var(ctx.GetRight().(*gen.Comp_varContext))

	// Получаем оператор сравнения
	op := ctx.GetOp()

	left, right := v.dereference(leftval), v.dereference(rightval)

	// Определяем тип сравнения (целые или плавающие)
	if left.Type().Equal(types.I32) {
		// Для целых чисел используем icmp
		switch op.GetTokenType() {
		case gen.RubyKParserLESS:
			return v.block().NewICmp(enum.IPredSLT, left, right)
		case gen.RubyKParserGREATER:
			return v.block().NewICmp(enum.IPredSGT, left, right)
		case gen.RubyKParserLESS_EQUAL:
			return v.block().NewICmp(enum.IPredSLE, left, right)
		case gen.RubyKParserGREATER_EQUAL:
			return v.block().NewICmp(enum.IPredSGE, left, right)
		case gen.RubyKParserEQUAL:
			return v.block().NewICmp(enum.IPredEQ, left, right)
		case gen.RubyKParserNOT_EQUAL:
			return v.block().NewICmp(enum.IPredNE, left, right)
		default:
			panic("Неизвестный оператор сравнения для целых чисел")
		}
	} else if left.Type().Equal(types.Double) {
		// Для чисел с плавающей точкой используем fcmp
		switch op.GetTokenType() {
		case gen.RubyKParserLESS:
			return v.Main.NewFCmp(enum.FPredOLT, left, right)
		case gen.RubyKParserGREATER:
			return v.Main.NewFCmp(enum.FPredOGT, left, right)
		case gen.RubyKParserLESS_EQUAL:
			return v.Main.NewFCmp(enum.FPredOLE, left, right)
		case gen.RubyKParserGREATER_EQUAL:
			return v.Main.NewFCmp(enum.FPredOGE, left, right)
		case gen.RubyKParserEQUAL:
			return v.Main.NewFCmp(enum.FPredOEQ, left, right)
		case gen.RubyKParserNOT_EQUAL:
			return v.Main.NewFCmp(enum.FPredONE, left, right)
		default:
			panic("Неизвестный оператор сравнения для чисел с плавающей точкой")
		}
	} else {
		panic("Типы данных не поддерживаются для сравнения")
	}
}

func (v *MyRubyKVisitor) VisitComp_var(ctx *gen.Comp_varContext) value.Value {
	// comp_var
	//     : all_result
	//     | array_selector
	//     | id_
	//     ;

	switch true {
	case ctx.All_result() != nil:
		return v.VisitAll_result(ctx.All_result().(*gen.All_resultContext))
	case ctx.Array_selector() != nil:
		return nil
		// return v.VisitArray_selector(ctx.Array_selector().(*gen.Array_selectorContext))
	case ctx.Id_() != nil:
		name := v.VisitId_(ctx.Id_().(*gen.Id_Context))
		vr, exists := v.variable(name)
		if !exists {
			panic(fmt.Sprintf("variable %s not defined", name))
		}
		return vr
	}

	return nil
}

func (v *MyRubyKVisitor) VisitLvalue(ctx *gen.LvalueContext) interface{} {
	if ctx_id := ctx.Id_(); ctx != nil {
		return v.VisitId_(ctx_id.(*gen.Id_Context))
	}
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitRvalue(ctx *gen.RvalueContext) interface{} {
	// rvalue
	//     : lvalue
	//     | initial_array_assignment
	//     | int_result
	//     | array_assignment
	//     | float_result
	//     | string_result
	//     | string_assignment
	//     | float_assignment
	//     | int_assignment
	//     | assignment
	//     | function_call
	//     | literal_t
	//     | bool_t
	//     | float_t
	//     | int_t
	//     | nil_t
	//     | rvalue EXP rvalue
	//     | ( NOT | BIT_NOT) rvalue
	//     | rvalue ( MUL | DIV | MOD) rvalue
	//     | rvalue ( PLUS | MINUS) rvalue
	//     | rvalue ( BIT_SHL | BIT_SHR) rvalue
	//     | rvalue BIT_AND rvalue
	//     | rvalue ( BIT_OR | BIT_XOR) rvalue
	//     | rvalue ( LESS | GREATER | LESS_EQUAL | GREATER_EQUAL) rvalue
	//     | rvalue ( EQUAL | NOT_EQUAL) rvalue
	//     | rvalue ( OR | AND) rvalue
	//     | LEFT_RBRACKET rvalue RIGHT_RBRACKET
	//     | dynamic_assignment
	//     ;
	if v.Debug {
		log.Println("Rvalue visited", ctx.GetText())
	}

	if ctx_intass := ctx.Int_assignment(); ctx_intass != nil {
		v.VisitInt_assignment(ctx_intass.(*gen.Int_assignmentContext))
		return nil
	}

	if ctx_arrass := ctx.Array_assignment(); ctx_arrass != nil {
		v.VisitArray_assignment(ctx_arrass.(*gen.Array_assignmentContext))
		return nil
	}

	if ctx_arrass := ctx.Initial_array_assignment(); ctx_arrass != nil {
		v.VisitInitial_array_assignment(ctx_arrass.(*gen.Initial_array_assignmentContext))
		return nil
	}

	if ctx_ass := ctx.Assignment(); ctx_ass != nil {
		v.VisitAssignment(ctx_ass.(*gen.AssignmentContext))
		return nil
	}

	if ctx_f := ctx.Function_call(); ctx_f != nil {
		return v.VisitFunction_call(ctx_f.(*gen.Function_callContext))
	}

	if ctx_ass := ctx.Dynamic_assignment(); ctx_ass != nil {
		v.VisitDynamic_assignment(ctx_ass.(*gen.Dynamic_assignmentContext))
		return nil
	}

	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitBreak_expression(ctx *gen.Break_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitLiteral_t(ctx *gen.Literal_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitFloat_t(ctx *gen.Float_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitInt_t(ctx *gen.Int_tContext) value.Value {
	if v.Debug {
		log.Println("int_t visited", ctx.GetText())
	}

	val, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}

	return constant.NewInt(types.I32, int64(val))
}

func (v *MyRubyKVisitor) VisitBool_t(ctx *gen.Bool_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitNil_t(ctx *gen.Nil_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitId_(ctx *gen.Id_Context) string {
	if v.Debug {
		log.Println("id_ visited: ", ctx.GetText())
	}
	return ctx.GetText()
}

func (v *MyRubyKVisitor) VisitId_global(ctx *gen.Id_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitId_function(ctx *gen.Id_functionContext) string {
	// id_function
	// : ID_FUNCTION
	// ;
	return ctx.ID_FUNCTION().GetText()
}

func (v *MyRubyKVisitor) VisitTerminator(ctx *gen.TerminatorContext) interface{} {
	// terminator
	//     : terminator SEMICOLON
	//     | terminator crlf
	//     | SEMICOLON
	//     | crlf
	//     ;
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitElse_token(ctx *gen.Else_tokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyRubyKVisitor) VisitCrlf(ctx *gen.CrlfContext) interface{} {
	return v.VisitChildren(ctx)
}
