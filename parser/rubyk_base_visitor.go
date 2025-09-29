// Code generated from RubyK.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // RubyK

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRubyKVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRubyKVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitGlobal_get(ctx *Global_getContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitGlobal_set(ctx *Global_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitGlobal_result(ctx *Global_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_inline_call(ctx *Function_inline_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition(ctx *Function_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition_body(ctx *Function_definition_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition_header(ctx *Function_definition_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition_params(ctx *Function_definition_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition_params_list(ctx *Function_definition_params_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_definition_param_id(ctx *Function_definition_param_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_call_param_list(ctx *Function_call_param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_call_params(ctx *Function_call_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_param(ctx *Function_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_unnamed_param(ctx *Function_unnamed_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_named_param(ctx *Function_named_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFunction_call_assignment(ctx *Function_call_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitAll_result(ctx *All_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitElsif_statement(ctx *Elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitIf_elsif_statement(ctx *If_elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitUnless_statement(ctx *Unless_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitInit_expression(ctx *Init_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitAll_assignment(ctx *All_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFor_init_list(ctx *For_init_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitCond_expression(ctx *Cond_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitLoop_expression(ctx *Loop_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFor_loop_list(ctx *For_loop_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitStatement_body(ctx *Statement_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitStatement_expression_list(ctx *Statement_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitDynamic_assignment(ctx *Dynamic_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitInt_assignment(ctx *Int_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFloat_assignment(ctx *Float_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitString_assignment(ctx *String_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitInitial_array_assignment(ctx *Initial_array_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitArray_assignment(ctx *Array_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitArray_definition(ctx *Array_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitArray_definition_elements(ctx *Array_definition_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitArray_selector(ctx *Array_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitDynamic_result(ctx *Dynamic_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitDynamic_(ctx *Dynamic_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitInt_result(ctx *Int_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFloat_result(ctx *Float_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitString_result(ctx *String_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitComparison_list(ctx *Comparison_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitComp_var(ctx *Comp_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitLvalue(ctx *LvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitRvalue(ctx *RvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitBreak_expression(ctx *Break_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitLiteral_t(ctx *Literal_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitFloat_t(ctx *Float_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitInt_t(ctx *Int_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitBool_t(ctx *Bool_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitNil_t(ctx *Nil_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitId_(ctx *Id_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitId_global(ctx *Id_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitId_function(ctx *Id_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitTerminator(ctx *TerminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitElse_token(ctx *Else_tokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubyKVisitor) VisitCrlf(ctx *CrlfContext) interface{} {
	return v.VisitChildren(ctx)
}
