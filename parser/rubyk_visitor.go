// Code generated from RubyK.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // RubyK

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RubyKParser.
type RubyKVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RubyKParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RubyKParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RubyKParser#global_get.
	VisitGlobal_get(ctx *Global_getContext) interface{}

	// Visit a parse tree produced by RubyKParser#global_set.
	VisitGlobal_set(ctx *Global_setContext) interface{}

	// Visit a parse tree produced by RubyKParser#global_result.
	VisitGlobal_result(ctx *Global_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_inline_call.
	VisitFunction_inline_call(ctx *Function_inline_callContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition.
	VisitFunction_definition(ctx *Function_definitionContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition_body.
	VisitFunction_definition_body(ctx *Function_definition_bodyContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition_header.
	VisitFunction_definition_header(ctx *Function_definition_headerContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition_params.
	VisitFunction_definition_params(ctx *Function_definition_paramsContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition_params_list.
	VisitFunction_definition_params_list(ctx *Function_definition_params_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_definition_param_id.
	VisitFunction_definition_param_id(ctx *Function_definition_param_idContext) interface{}

	// Visit a parse tree produced by RubyKParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_call.
	VisitFunction_call(ctx *Function_callContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_call_param_list.
	VisitFunction_call_param_list(ctx *Function_call_param_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_call_params.
	VisitFunction_call_params(ctx *Function_call_paramsContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_param.
	VisitFunction_param(ctx *Function_paramContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_unnamed_param.
	VisitFunction_unnamed_param(ctx *Function_unnamed_paramContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_named_param.
	VisitFunction_named_param(ctx *Function_named_paramContext) interface{}

	// Visit a parse tree produced by RubyKParser#function_call_assignment.
	VisitFunction_call_assignment(ctx *Function_call_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#all_result.
	VisitAll_result(ctx *All_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#elsif_statement.
	VisitElsif_statement(ctx *Elsif_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#if_elsif_statement.
	VisitIf_elsif_statement(ctx *If_elsif_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#unless_statement.
	VisitUnless_statement(ctx *Unless_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by RubyKParser#init_expression.
	VisitInit_expression(ctx *Init_expressionContext) interface{}

	// Visit a parse tree produced by RubyKParser#all_assignment.
	VisitAll_assignment(ctx *All_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#for_init_list.
	VisitFor_init_list(ctx *For_init_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#cond_expression.
	VisitCond_expression(ctx *Cond_expressionContext) interface{}

	// Visit a parse tree produced by RubyKParser#loop_expression.
	VisitLoop_expression(ctx *Loop_expressionContext) interface{}

	// Visit a parse tree produced by RubyKParser#for_loop_list.
	VisitFor_loop_list(ctx *For_loop_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#statement_body.
	VisitStatement_body(ctx *Statement_bodyContext) interface{}

	// Visit a parse tree produced by RubyKParser#statement_expression_list.
	VisitStatement_expression_list(ctx *Statement_expression_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#dynamic_assignment.
	VisitDynamic_assignment(ctx *Dynamic_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#int_assignment.
	VisitInt_assignment(ctx *Int_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#float_assignment.
	VisitFloat_assignment(ctx *Float_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#string_assignment.
	VisitString_assignment(ctx *String_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#initial_array_assignment.
	VisitInitial_array_assignment(ctx *Initial_array_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#array_assignment.
	VisitArray_assignment(ctx *Array_assignmentContext) interface{}

	// Visit a parse tree produced by RubyKParser#array_definition.
	VisitArray_definition(ctx *Array_definitionContext) interface{}

	// Visit a parse tree produced by RubyKParser#array_definition_elements.
	VisitArray_definition_elements(ctx *Array_definition_elementsContext) interface{}

	// Visit a parse tree produced by RubyKParser#array_selector.
	VisitArray_selector(ctx *Array_selectorContext) interface{}

	// Visit a parse tree produced by RubyKParser#dynamic_result.
	VisitDynamic_result(ctx *Dynamic_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#dynamic_.
	VisitDynamic_(ctx *Dynamic_Context) interface{}

	// Visit a parse tree produced by RubyKParser#int_result.
	VisitInt_result(ctx *Int_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#float_result.
	VisitFloat_result(ctx *Float_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#string_result.
	VisitString_result(ctx *String_resultContext) interface{}

	// Visit a parse tree produced by RubyKParser#comparison_list.
	VisitComparison_list(ctx *Comparison_listContext) interface{}

	// Visit a parse tree produced by RubyKParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by RubyKParser#comp_var.
	VisitComp_var(ctx *Comp_varContext) interface{}

	// Visit a parse tree produced by RubyKParser#lvalue.
	VisitLvalue(ctx *LvalueContext) interface{}

	// Visit a parse tree produced by RubyKParser#rvalue.
	VisitRvalue(ctx *RvalueContext) interface{}

	// Visit a parse tree produced by RubyKParser#break_expression.
	VisitBreak_expression(ctx *Break_expressionContext) interface{}

	// Visit a parse tree produced by RubyKParser#literal_t.
	VisitLiteral_t(ctx *Literal_tContext) interface{}

	// Visit a parse tree produced by RubyKParser#float_t.
	VisitFloat_t(ctx *Float_tContext) interface{}

	// Visit a parse tree produced by RubyKParser#int_t.
	VisitInt_t(ctx *Int_tContext) interface{}

	// Visit a parse tree produced by RubyKParser#bool_t.
	VisitBool_t(ctx *Bool_tContext) interface{}

	// Visit a parse tree produced by RubyKParser#nil_t.
	VisitNil_t(ctx *Nil_tContext) interface{}

	// Visit a parse tree produced by RubyKParser#id_.
	VisitId_(ctx *Id_Context) interface{}

	// Visit a parse tree produced by RubyKParser#id_global.
	VisitId_global(ctx *Id_globalContext) interface{}

	// Visit a parse tree produced by RubyKParser#id_function.
	VisitId_function(ctx *Id_functionContext) interface{}

	// Visit a parse tree produced by RubyKParser#terminator.
	VisitTerminator(ctx *TerminatorContext) interface{}

	// Visit a parse tree produced by RubyKParser#else_token.
	VisitElse_token(ctx *Else_tokenContext) interface{}

	// Visit a parse tree produced by RubyKParser#crlf.
	VisitCrlf(ctx *CrlfContext) interface{}
}
