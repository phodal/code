// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CodeListener is a complete listener for a parse tree produced by CodeParser.
type CodeListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPackageDeclaration is called when entering the packageDeclaration production.
	EnterPackageDeclaration(c *PackageDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterExpressDeclaration is called when entering the expressDeclaration production.
	EnterExpressDeclaration(c *ExpressDeclarationContext)

	// EnterMethodCallDeclaration is called when entering the methodCallDeclaration production.
	EnterMethodCallDeclaration(c *MethodCallDeclarationContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterDataStructDeclaration is called when entering the dataStructDeclaration production.
	EnterDataStructDeclaration(c *DataStructDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitExpressDeclaration is called when exiting the expressDeclaration production.
	ExitExpressDeclaration(c *ExpressDeclarationContext)

	// ExitMethodCallDeclaration is called when exiting the methodCallDeclaration production.
	ExitMethodCallDeclaration(c *MethodCallDeclarationContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitDataStructDeclaration is called when exiting the dataStructDeclaration production.
	ExitDataStructDeclaration(c *DataStructDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)
}
