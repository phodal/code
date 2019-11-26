// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CodeParser.
type CodeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CodeParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by CodeParser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#expressDeclaration.
	VisitExpressDeclaration(ctx *ExpressDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#methodCallDeclaration.
	VisitMethodCallDeclaration(ctx *MethodCallDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by CodeParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by CodeParser#dataStructDeclaration.
	VisitDataStructDeclaration(ctx *DataStructDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by CodeParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}
}
