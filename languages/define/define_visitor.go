// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DefineParser.
type DefineVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DefineParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by DefineParser#symbolDeclaration.
	VisitSymbolDeclaration(ctx *SymbolDeclarationContext) interface{}

	// Visit a parse tree produced by DefineParser#normalDeclarations.
	VisitNormalDeclarations(ctx *NormalDeclarationsContext) interface{}

	// Visit a parse tree produced by DefineParser#defineDeclaration.
	VisitDefineDeclaration(ctx *DefineDeclarationContext) interface{}

	// Visit a parse tree produced by DefineParser#defineExpress.
	VisitDefineExpress(ctx *DefineExpressContext) interface{}

	// Visit a parse tree produced by DefineParser#defineAttribute.
	VisitDefineAttribute(ctx *DefineAttributeContext) interface{}

	// Visit a parse tree produced by DefineParser#defineTemplate.
	VisitDefineTemplate(ctx *DefineTemplateContext) interface{}

	// Visit a parse tree produced by DefineParser#symbolKey.
	VisitSymbolKey(ctx *SymbolKeyContext) interface{}

	// Visit a parse tree produced by DefineParser#symbolValue.
	VisitSymbolValue(ctx *SymbolValueContext) interface{}

	// Visit a parse tree produced by DefineParser#defineBody.
	VisitDefineBody(ctx *DefineBodyContext) interface{}

	// Visit a parse tree produced by DefineParser#templateData.
	VisitTemplateData(ctx *TemplateDataContext) interface{}

	// Visit a parse tree produced by DefineParser#systemDeclaration.
	VisitSystemDeclaration(ctx *SystemDeclarationContext) interface{}

	// Visit a parse tree produced by DefineParser#defineKey.
	VisitDefineKey(ctx *DefineKeyContext) interface{}

	// Visit a parse tree produced by DefineParser#defineValue.
	VisitDefineValue(ctx *DefineValueContext) interface{}

	// Visit a parse tree produced by DefineParser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{}

	// Visit a parse tree produced by DefineParser#moduleDefine.
	VisitModuleDefine(ctx *ModuleDefineContext) interface{}

	// Visit a parse tree produced by DefineParser#moduleAttribute.
	VisitModuleAttribute(ctx *ModuleAttributeContext) interface{}
}
