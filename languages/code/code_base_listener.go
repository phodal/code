// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCodeListener is a complete listener for a parse tree produced by CodeParser.
type BaseCodeListener struct{}

var _ CodeListener = &BaseCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseCodeListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseCodeListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseCodeListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseCodeListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseCodeListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseCodeListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseCodeListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseCodeListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseCodeListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseCodeListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseCodeListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseCodeListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterExpressDeclaration is called when production expressDeclaration is entered.
func (s *BaseCodeListener) EnterExpressDeclaration(ctx *ExpressDeclarationContext) {}

// ExitExpressDeclaration is called when production expressDeclaration is exited.
func (s *BaseCodeListener) ExitExpressDeclaration(ctx *ExpressDeclarationContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseCodeListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseCodeListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseCodeListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseCodeListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseCodeListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseCodeListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseCodeListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseCodeListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseCodeListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseCodeListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseCodeListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseCodeListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMethodCallDeclaration is called when production methodCallDeclaration is entered.
func (s *BaseCodeListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {}

// ExitMethodCallDeclaration is called when production methodCallDeclaration is exited.
func (s *BaseCodeListener) ExitMethodCallDeclaration(ctx *MethodCallDeclarationContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCodeListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCodeListener) ExitParameter(ctx *ParameterContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCodeListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCodeListener) ExitLiteral(ctx *LiteralContext) {}

// EnterDataStructDeclaration is called when production dataStructDeclaration is entered.
func (s *BaseCodeListener) EnterDataStructDeclaration(ctx *DataStructDeclarationContext) {}

// ExitDataStructDeclaration is called when production dataStructDeclaration is exited.
func (s *BaseCodeListener) ExitDataStructDeclaration(ctx *DataStructDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseCodeListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseCodeListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}
