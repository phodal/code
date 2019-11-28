// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DefineListener is a complete listener for a parse tree produced by DefineParser.
type DefineListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterSymbolDeclaration is called when entering the symbolDeclaration production.
	EnterSymbolDeclaration(c *SymbolDeclarationContext)

	// EnterNormalDeclarations is called when entering the normalDeclarations production.
	EnterNormalDeclarations(c *NormalDeclarationsContext)

	// EnterDefineDeclaration is called when entering the defineDeclaration production.
	EnterDefineDeclaration(c *DefineDeclarationContext)

	// EnterDefineExpress is called when entering the defineExpress production.
	EnterDefineExpress(c *DefineExpressContext)

	// EnterDefineAttribute is called when entering the defineAttribute production.
	EnterDefineAttribute(c *DefineAttributeContext)

	// EnterDefineTemplate is called when entering the defineTemplate production.
	EnterDefineTemplate(c *DefineTemplateContext)

	// EnterSymbolKey is called when entering the symbolKey production.
	EnterSymbolKey(c *SymbolKeyContext)

	// EnterSymbolValue is called when entering the symbolValue production.
	EnterSymbolValue(c *SymbolValueContext)

	// EnterDefineBody is called when entering the defineBody production.
	EnterDefineBody(c *DefineBodyContext)

	// EnterTemplateData is called when entering the templateData production.
	EnterTemplateData(c *TemplateDataContext)

	// EnterSystemDeclaration is called when entering the systemDeclaration production.
	EnterSystemDeclaration(c *SystemDeclarationContext)

	// EnterDefineKey is called when entering the defineKey production.
	EnterDefineKey(c *DefineKeyContext)

	// EnterDefineValue is called when entering the defineValue production.
	EnterDefineValue(c *DefineValueContext)

	// EnterModuleDeclaration is called when entering the moduleDeclaration production.
	EnterModuleDeclaration(c *ModuleDeclarationContext)

	// EnterModuleDefines is called when entering the moduleDefines production.
	EnterModuleDefines(c *ModuleDefinesContext)

	// EnterModuleAttribute is called when entering the moduleAttribute production.
	EnterModuleAttribute(c *ModuleAttributeContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitSymbolDeclaration is called when exiting the symbolDeclaration production.
	ExitSymbolDeclaration(c *SymbolDeclarationContext)

	// ExitNormalDeclarations is called when exiting the normalDeclarations production.
	ExitNormalDeclarations(c *NormalDeclarationsContext)

	// ExitDefineDeclaration is called when exiting the defineDeclaration production.
	ExitDefineDeclaration(c *DefineDeclarationContext)

	// ExitDefineExpress is called when exiting the defineExpress production.
	ExitDefineExpress(c *DefineExpressContext)

	// ExitDefineAttribute is called when exiting the defineAttribute production.
	ExitDefineAttribute(c *DefineAttributeContext)

	// ExitDefineTemplate is called when exiting the defineTemplate production.
	ExitDefineTemplate(c *DefineTemplateContext)

	// ExitSymbolKey is called when exiting the symbolKey production.
	ExitSymbolKey(c *SymbolKeyContext)

	// ExitSymbolValue is called when exiting the symbolValue production.
	ExitSymbolValue(c *SymbolValueContext)

	// ExitDefineBody is called when exiting the defineBody production.
	ExitDefineBody(c *DefineBodyContext)

	// ExitTemplateData is called when exiting the templateData production.
	ExitTemplateData(c *TemplateDataContext)

	// ExitSystemDeclaration is called when exiting the systemDeclaration production.
	ExitSystemDeclaration(c *SystemDeclarationContext)

	// ExitDefineKey is called when exiting the defineKey production.
	ExitDefineKey(c *DefineKeyContext)

	// ExitDefineValue is called when exiting the defineValue production.
	ExitDefineValue(c *DefineValueContext)

	// ExitModuleDeclaration is called when exiting the moduleDeclaration production.
	ExitModuleDeclaration(c *ModuleDeclarationContext)

	// ExitModuleDefines is called when exiting the moduleDefines production.
	ExitModuleDefines(c *ModuleDefinesContext)

	// ExitModuleAttribute is called when exiting the moduleAttribute production.
	ExitModuleAttribute(c *ModuleAttributeContext)
}
