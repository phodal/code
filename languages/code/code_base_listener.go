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

// EnterExpressDeclaration is called when production expressDeclaration is entered.
func (s *BaseCodeListener) EnterExpressDeclaration(ctx *ExpressDeclarationContext) {}

// ExitExpressDeclaration is called when production expressDeclaration is exited.
func (s *BaseCodeListener) ExitExpressDeclaration(ctx *ExpressDeclarationContext) {}

// EnterMethodCallDeclaration is called when production methodCallDeclaration is entered.
func (s *BaseCodeListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {}

// ExitMethodCallDeclaration is called when production methodCallDeclaration is exited.
func (s *BaseCodeListener) ExitMethodCallDeclaration(ctx *MethodCallDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseCodeListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseCodeListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCodeListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCodeListener) ExitParameter(ctx *ParameterContext) {}

// EnterDataStructDeclaration is called when production dataStructDeclaration is entered.
func (s *BaseCodeListener) EnterDataStructDeclaration(ctx *DataStructDeclarationContext) {}

// ExitDataStructDeclaration is called when production dataStructDeclaration is exited.
func (s *BaseCodeListener) ExitDataStructDeclaration(ctx *DataStructDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseCodeListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseCodeListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseCodeListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseCodeListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}
