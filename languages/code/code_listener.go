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

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterExpressDeclaration is called when entering the expressDeclaration production.
	EnterExpressDeclaration(c *ExpressDeclarationContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMethodCallDeclaration is called when entering the methodCallDeclaration production.
	EnterMethodCallDeclaration(c *MethodCallDeclarationContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterDataStructDeclaration is called when entering the dataStructDeclaration production.
	EnterDataStructDeclaration(c *DataStructDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitExpressDeclaration is called when exiting the expressDeclaration production.
	ExitExpressDeclaration(c *ExpressDeclarationContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMethodCallDeclaration is called when exiting the methodCallDeclaration production.
	ExitMethodCallDeclaration(c *MethodCallDeclarationContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitDataStructDeclaration is called when exiting the dataStructDeclaration production.
	ExitDataStructDeclaration(c *DataStructDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)
}
