// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DefineListener is a complete listener for a parse tree produced by DefineParser.
type DefineListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterSymbolDelcaration is called when entering the symbolDelcaration production.
	EnterSymbolDelcaration(c *SymbolDelcarationContext)

	// EnterNormalDeclarations is called when entering the normalDeclarations production.
	EnterNormalDeclarations(c *NormalDeclarationsContext)

	// EnterDefaultDeclaration is called when entering the defaultDeclaration production.
	EnterDefaultDeclaration(c *DefaultDeclarationContext)

	// EnterSystemDeclaration is called when entering the systemDeclaration production.
	EnterSystemDeclaration(c *SystemDeclarationContext)

	// EnterDefineKey is called when entering the defineKey production.
	EnterDefineKey(c *DefineKeyContext)

	// EnterDefineValue is called when entering the defineValue production.
	EnterDefineValue(c *DefineValueContext)

	// EnterModuleDeclaration is called when entering the moduleDeclaration production.
	EnterModuleDeclaration(c *ModuleDeclarationContext)

	// EnterModuleDefine is called when entering the moduleDefine production.
	EnterModuleDefine(c *ModuleDefineContext)

	// EnterModuleAttributes is called when entering the moduleAttributes production.
	EnterModuleAttributes(c *ModuleAttributesContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitSymbolDelcaration is called when exiting the symbolDelcaration production.
	ExitSymbolDelcaration(c *SymbolDelcarationContext)

	// ExitNormalDeclarations is called when exiting the normalDeclarations production.
	ExitNormalDeclarations(c *NormalDeclarationsContext)

	// ExitDefaultDeclaration is called when exiting the defaultDeclaration production.
	ExitDefaultDeclaration(c *DefaultDeclarationContext)

	// ExitSystemDeclaration is called when exiting the systemDeclaration production.
	ExitSystemDeclaration(c *SystemDeclarationContext)

	// ExitDefineKey is called when exiting the defineKey production.
	ExitDefineKey(c *DefineKeyContext)

	// ExitDefineValue is called when exiting the defineValue production.
	ExitDefineValue(c *DefineValueContext)

	// ExitModuleDeclaration is called when exiting the moduleDeclaration production.
	ExitModuleDeclaration(c *ModuleDeclarationContext)

	// ExitModuleDefine is called when exiting the moduleDefine production.
	ExitModuleDefine(c *ModuleDefineContext)

	// ExitModuleAttributes is called when exiting the moduleAttributes production.
	ExitModuleAttributes(c *ModuleAttributesContext)
}
