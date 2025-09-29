// Code generated from RubyK.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // RubyK

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRubyKListener is a complete listener for a parse tree produced by RubyKParser.
type BaseRubyKListener struct{}

var _ RubyKListener = &BaseRubyKListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRubyKListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRubyKListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRubyKListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRubyKListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRubyKListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRubyKListener) ExitProgram(ctx *ProgramContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseRubyKListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseRubyKListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRubyKListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRubyKListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGlobal_get is called when production global_get is entered.
func (s *BaseRubyKListener) EnterGlobal_get(ctx *Global_getContext) {}

// ExitGlobal_get is called when production global_get is exited.
func (s *BaseRubyKListener) ExitGlobal_get(ctx *Global_getContext) {}

// EnterGlobal_set is called when production global_set is entered.
func (s *BaseRubyKListener) EnterGlobal_set(ctx *Global_setContext) {}

// ExitGlobal_set is called when production global_set is exited.
func (s *BaseRubyKListener) ExitGlobal_set(ctx *Global_setContext) {}

// EnterGlobal_result is called when production global_result is entered.
func (s *BaseRubyKListener) EnterGlobal_result(ctx *Global_resultContext) {}

// ExitGlobal_result is called when production global_result is exited.
func (s *BaseRubyKListener) ExitGlobal_result(ctx *Global_resultContext) {}

// EnterFunction_inline_call is called when production function_inline_call is entered.
func (s *BaseRubyKListener) EnterFunction_inline_call(ctx *Function_inline_callContext) {}

// ExitFunction_inline_call is called when production function_inline_call is exited.
func (s *BaseRubyKListener) ExitFunction_inline_call(ctx *Function_inline_callContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseRubyKListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseRubyKListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterFunction_definition_body is called when production function_definition_body is entered.
func (s *BaseRubyKListener) EnterFunction_definition_body(ctx *Function_definition_bodyContext) {}

// ExitFunction_definition_body is called when production function_definition_body is exited.
func (s *BaseRubyKListener) ExitFunction_definition_body(ctx *Function_definition_bodyContext) {}

// EnterFunction_definition_header is called when production function_definition_header is entered.
func (s *BaseRubyKListener) EnterFunction_definition_header(ctx *Function_definition_headerContext) {}

// ExitFunction_definition_header is called when production function_definition_header is exited.
func (s *BaseRubyKListener) ExitFunction_definition_header(ctx *Function_definition_headerContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseRubyKListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseRubyKListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_definition_params is called when production function_definition_params is entered.
func (s *BaseRubyKListener) EnterFunction_definition_params(ctx *Function_definition_paramsContext) {}

// ExitFunction_definition_params is called when production function_definition_params is exited.
func (s *BaseRubyKListener) ExitFunction_definition_params(ctx *Function_definition_paramsContext) {}

// EnterFunction_definition_params_list is called when production function_definition_params_list is entered.
func (s *BaseRubyKListener) EnterFunction_definition_params_list(ctx *Function_definition_params_listContext) {
}

// ExitFunction_definition_params_list is called when production function_definition_params_list is exited.
func (s *BaseRubyKListener) ExitFunction_definition_params_list(ctx *Function_definition_params_listContext) {
}

// EnterFunction_definition_param_id is called when production function_definition_param_id is entered.
func (s *BaseRubyKListener) EnterFunction_definition_param_id(ctx *Function_definition_param_idContext) {
}

// ExitFunction_definition_param_id is called when production function_definition_param_id is exited.
func (s *BaseRubyKListener) ExitFunction_definition_param_id(ctx *Function_definition_param_idContext) {
}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseRubyKListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseRubyKListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseRubyKListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseRubyKListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFunction_call_param_list is called when production function_call_param_list is entered.
func (s *BaseRubyKListener) EnterFunction_call_param_list(ctx *Function_call_param_listContext) {}

// ExitFunction_call_param_list is called when production function_call_param_list is exited.
func (s *BaseRubyKListener) ExitFunction_call_param_list(ctx *Function_call_param_listContext) {}

// EnterFunction_call_params is called when production function_call_params is entered.
func (s *BaseRubyKListener) EnterFunction_call_params(ctx *Function_call_paramsContext) {}

// ExitFunction_call_params is called when production function_call_params is exited.
func (s *BaseRubyKListener) ExitFunction_call_params(ctx *Function_call_paramsContext) {}

// EnterFunction_param is called when production function_param is entered.
func (s *BaseRubyKListener) EnterFunction_param(ctx *Function_paramContext) {}

// ExitFunction_param is called when production function_param is exited.
func (s *BaseRubyKListener) ExitFunction_param(ctx *Function_paramContext) {}

// EnterFunction_unnamed_param is called when production function_unnamed_param is entered.
func (s *BaseRubyKListener) EnterFunction_unnamed_param(ctx *Function_unnamed_paramContext) {}

// ExitFunction_unnamed_param is called when production function_unnamed_param is exited.
func (s *BaseRubyKListener) ExitFunction_unnamed_param(ctx *Function_unnamed_paramContext) {}

// EnterFunction_named_param is called when production function_named_param is entered.
func (s *BaseRubyKListener) EnterFunction_named_param(ctx *Function_named_paramContext) {}

// ExitFunction_named_param is called when production function_named_param is exited.
func (s *BaseRubyKListener) ExitFunction_named_param(ctx *Function_named_paramContext) {}

// EnterFunction_call_assignment is called when production function_call_assignment is entered.
func (s *BaseRubyKListener) EnterFunction_call_assignment(ctx *Function_call_assignmentContext) {}

// ExitFunction_call_assignment is called when production function_call_assignment is exited.
func (s *BaseRubyKListener) ExitFunction_call_assignment(ctx *Function_call_assignmentContext) {}

// EnterAll_result is called when production all_result is entered.
func (s *BaseRubyKListener) EnterAll_result(ctx *All_resultContext) {}

// ExitAll_result is called when production all_result is exited.
func (s *BaseRubyKListener) ExitAll_result(ctx *All_resultContext) {}

// EnterElsif_statement is called when production elsif_statement is entered.
func (s *BaseRubyKListener) EnterElsif_statement(ctx *Elsif_statementContext) {}

// ExitElsif_statement is called when production elsif_statement is exited.
func (s *BaseRubyKListener) ExitElsif_statement(ctx *Elsif_statementContext) {}

// EnterIf_elsif_statement is called when production if_elsif_statement is entered.
func (s *BaseRubyKListener) EnterIf_elsif_statement(ctx *If_elsif_statementContext) {}

// ExitIf_elsif_statement is called when production if_elsif_statement is exited.
func (s *BaseRubyKListener) ExitIf_elsif_statement(ctx *If_elsif_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseRubyKListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseRubyKListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterUnless_statement is called when production unless_statement is entered.
func (s *BaseRubyKListener) EnterUnless_statement(ctx *Unless_statementContext) {}

// ExitUnless_statement is called when production unless_statement is exited.
func (s *BaseRubyKListener) ExitUnless_statement(ctx *Unless_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseRubyKListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseRubyKListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseRubyKListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseRubyKListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterInit_expression is called when production init_expression is entered.
func (s *BaseRubyKListener) EnterInit_expression(ctx *Init_expressionContext) {}

// ExitInit_expression is called when production init_expression is exited.
func (s *BaseRubyKListener) ExitInit_expression(ctx *Init_expressionContext) {}

// EnterAll_assignment is called when production all_assignment is entered.
func (s *BaseRubyKListener) EnterAll_assignment(ctx *All_assignmentContext) {}

// ExitAll_assignment is called when production all_assignment is exited.
func (s *BaseRubyKListener) ExitAll_assignment(ctx *All_assignmentContext) {}

// EnterFor_init_list is called when production for_init_list is entered.
func (s *BaseRubyKListener) EnterFor_init_list(ctx *For_init_listContext) {}

// ExitFor_init_list is called when production for_init_list is exited.
func (s *BaseRubyKListener) ExitFor_init_list(ctx *For_init_listContext) {}

// EnterCond_expression is called when production cond_expression is entered.
func (s *BaseRubyKListener) EnterCond_expression(ctx *Cond_expressionContext) {}

// ExitCond_expression is called when production cond_expression is exited.
func (s *BaseRubyKListener) ExitCond_expression(ctx *Cond_expressionContext) {}

// EnterLoop_expression is called when production loop_expression is entered.
func (s *BaseRubyKListener) EnterLoop_expression(ctx *Loop_expressionContext) {}

// ExitLoop_expression is called when production loop_expression is exited.
func (s *BaseRubyKListener) ExitLoop_expression(ctx *Loop_expressionContext) {}

// EnterFor_loop_list is called when production for_loop_list is entered.
func (s *BaseRubyKListener) EnterFor_loop_list(ctx *For_loop_listContext) {}

// ExitFor_loop_list is called when production for_loop_list is exited.
func (s *BaseRubyKListener) ExitFor_loop_list(ctx *For_loop_listContext) {}

// EnterStatement_body is called when production statement_body is entered.
func (s *BaseRubyKListener) EnterStatement_body(ctx *Statement_bodyContext) {}

// ExitStatement_body is called when production statement_body is exited.
func (s *BaseRubyKListener) ExitStatement_body(ctx *Statement_bodyContext) {}

// EnterStatement_expression_list is called when production statement_expression_list is entered.
func (s *BaseRubyKListener) EnterStatement_expression_list(ctx *Statement_expression_listContext) {}

// ExitStatement_expression_list is called when production statement_expression_list is exited.
func (s *BaseRubyKListener) ExitStatement_expression_list(ctx *Statement_expression_listContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseRubyKListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseRubyKListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterDynamic_assignment is called when production dynamic_assignment is entered.
func (s *BaseRubyKListener) EnterDynamic_assignment(ctx *Dynamic_assignmentContext) {}

// ExitDynamic_assignment is called when production dynamic_assignment is exited.
func (s *BaseRubyKListener) ExitDynamic_assignment(ctx *Dynamic_assignmentContext) {}

// EnterInt_assignment is called when production int_assignment is entered.
func (s *BaseRubyKListener) EnterInt_assignment(ctx *Int_assignmentContext) {}

// ExitInt_assignment is called when production int_assignment is exited.
func (s *BaseRubyKListener) ExitInt_assignment(ctx *Int_assignmentContext) {}

// EnterFloat_assignment is called when production float_assignment is entered.
func (s *BaseRubyKListener) EnterFloat_assignment(ctx *Float_assignmentContext) {}

// ExitFloat_assignment is called when production float_assignment is exited.
func (s *BaseRubyKListener) ExitFloat_assignment(ctx *Float_assignmentContext) {}

// EnterString_assignment is called when production string_assignment is entered.
func (s *BaseRubyKListener) EnterString_assignment(ctx *String_assignmentContext) {}

// ExitString_assignment is called when production string_assignment is exited.
func (s *BaseRubyKListener) ExitString_assignment(ctx *String_assignmentContext) {}

// EnterInitial_array_assignment is called when production initial_array_assignment is entered.
func (s *BaseRubyKListener) EnterInitial_array_assignment(ctx *Initial_array_assignmentContext) {}

// ExitInitial_array_assignment is called when production initial_array_assignment is exited.
func (s *BaseRubyKListener) ExitInitial_array_assignment(ctx *Initial_array_assignmentContext) {}

// EnterArray_assignment is called when production array_assignment is entered.
func (s *BaseRubyKListener) EnterArray_assignment(ctx *Array_assignmentContext) {}

// ExitArray_assignment is called when production array_assignment is exited.
func (s *BaseRubyKListener) ExitArray_assignment(ctx *Array_assignmentContext) {}

// EnterArray_definition is called when production array_definition is entered.
func (s *BaseRubyKListener) EnterArray_definition(ctx *Array_definitionContext) {}

// ExitArray_definition is called when production array_definition is exited.
func (s *BaseRubyKListener) ExitArray_definition(ctx *Array_definitionContext) {}

// EnterArray_definition_elements is called when production array_definition_elements is entered.
func (s *BaseRubyKListener) EnterArray_definition_elements(ctx *Array_definition_elementsContext) {}

// ExitArray_definition_elements is called when production array_definition_elements is exited.
func (s *BaseRubyKListener) ExitArray_definition_elements(ctx *Array_definition_elementsContext) {}

// EnterArray_selector is called when production array_selector is entered.
func (s *BaseRubyKListener) EnterArray_selector(ctx *Array_selectorContext) {}

// ExitArray_selector is called when production array_selector is exited.
func (s *BaseRubyKListener) ExitArray_selector(ctx *Array_selectorContext) {}

// EnterDynamic_result is called when production dynamic_result is entered.
func (s *BaseRubyKListener) EnterDynamic_result(ctx *Dynamic_resultContext) {}

// ExitDynamic_result is called when production dynamic_result is exited.
func (s *BaseRubyKListener) ExitDynamic_result(ctx *Dynamic_resultContext) {}

// EnterDynamic_ is called when production dynamic_ is entered.
func (s *BaseRubyKListener) EnterDynamic_(ctx *Dynamic_Context) {}

// ExitDynamic_ is called when production dynamic_ is exited.
func (s *BaseRubyKListener) ExitDynamic_(ctx *Dynamic_Context) {}

// EnterInt_result is called when production int_result is entered.
func (s *BaseRubyKListener) EnterInt_result(ctx *Int_resultContext) {}

// ExitInt_result is called when production int_result is exited.
func (s *BaseRubyKListener) ExitInt_result(ctx *Int_resultContext) {}

// EnterFloat_result is called when production float_result is entered.
func (s *BaseRubyKListener) EnterFloat_result(ctx *Float_resultContext) {}

// ExitFloat_result is called when production float_result is exited.
func (s *BaseRubyKListener) ExitFloat_result(ctx *Float_resultContext) {}

// EnterString_result is called when production string_result is entered.
func (s *BaseRubyKListener) EnterString_result(ctx *String_resultContext) {}

// ExitString_result is called when production string_result is exited.
func (s *BaseRubyKListener) ExitString_result(ctx *String_resultContext) {}

// EnterComparison_list is called when production comparison_list is entered.
func (s *BaseRubyKListener) EnterComparison_list(ctx *Comparison_listContext) {}

// ExitComparison_list is called when production comparison_list is exited.
func (s *BaseRubyKListener) ExitComparison_list(ctx *Comparison_listContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseRubyKListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseRubyKListener) ExitComparison(ctx *ComparisonContext) {}

// EnterComp_var is called when production comp_var is entered.
func (s *BaseRubyKListener) EnterComp_var(ctx *Comp_varContext) {}

// ExitComp_var is called when production comp_var is exited.
func (s *BaseRubyKListener) ExitComp_var(ctx *Comp_varContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseRubyKListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseRubyKListener) ExitLvalue(ctx *LvalueContext) {}

// EnterRvalue is called when production rvalue is entered.
func (s *BaseRubyKListener) EnterRvalue(ctx *RvalueContext) {}

// ExitRvalue is called when production rvalue is exited.
func (s *BaseRubyKListener) ExitRvalue(ctx *RvalueContext) {}

// EnterBreak_expression is called when production break_expression is entered.
func (s *BaseRubyKListener) EnterBreak_expression(ctx *Break_expressionContext) {}

// ExitBreak_expression is called when production break_expression is exited.
func (s *BaseRubyKListener) ExitBreak_expression(ctx *Break_expressionContext) {}

// EnterLiteral_t is called when production literal_t is entered.
func (s *BaseRubyKListener) EnterLiteral_t(ctx *Literal_tContext) {}

// ExitLiteral_t is called when production literal_t is exited.
func (s *BaseRubyKListener) ExitLiteral_t(ctx *Literal_tContext) {}

// EnterFloat_t is called when production float_t is entered.
func (s *BaseRubyKListener) EnterFloat_t(ctx *Float_tContext) {}

// ExitFloat_t is called when production float_t is exited.
func (s *BaseRubyKListener) ExitFloat_t(ctx *Float_tContext) {}

// EnterInt_t is called when production int_t is entered.
func (s *BaseRubyKListener) EnterInt_t(ctx *Int_tContext) {}

// ExitInt_t is called when production int_t is exited.
func (s *BaseRubyKListener) ExitInt_t(ctx *Int_tContext) {}

// EnterBool_t is called when production bool_t is entered.
func (s *BaseRubyKListener) EnterBool_t(ctx *Bool_tContext) {}

// ExitBool_t is called when production bool_t is exited.
func (s *BaseRubyKListener) ExitBool_t(ctx *Bool_tContext) {}

// EnterNil_t is called when production nil_t is entered.
func (s *BaseRubyKListener) EnterNil_t(ctx *Nil_tContext) {}

// ExitNil_t is called when production nil_t is exited.
func (s *BaseRubyKListener) ExitNil_t(ctx *Nil_tContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseRubyKListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseRubyKListener) ExitId_(ctx *Id_Context) {}

// EnterId_global is called when production id_global is entered.
func (s *BaseRubyKListener) EnterId_global(ctx *Id_globalContext) {}

// ExitId_global is called when production id_global is exited.
func (s *BaseRubyKListener) ExitId_global(ctx *Id_globalContext) {}

// EnterId_function is called when production id_function is entered.
func (s *BaseRubyKListener) EnterId_function(ctx *Id_functionContext) {}

// ExitId_function is called when production id_function is exited.
func (s *BaseRubyKListener) ExitId_function(ctx *Id_functionContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseRubyKListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseRubyKListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterElse_token is called when production else_token is entered.
func (s *BaseRubyKListener) EnterElse_token(ctx *Else_tokenContext) {}

// ExitElse_token is called when production else_token is exited.
func (s *BaseRubyKListener) ExitElse_token(ctx *Else_tokenContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BaseRubyKListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BaseRubyKListener) ExitCrlf(ctx *CrlfContext) {}
