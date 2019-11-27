// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDefineListener is a complete listener for a parse tree produced by DefineParser.
type BaseDefineListener struct{}

var _ DefineListener = &BaseDefineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseDefineListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseDefineListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterSymbolDelcaration is called when production symbolDelcaration is entered.
func (s *BaseDefineListener) EnterSymbolDelcaration(ctx *SymbolDelcarationContext) {}

// ExitSymbolDelcaration is called when production symbolDelcaration is exited.
func (s *BaseDefineListener) ExitSymbolDelcaration(ctx *SymbolDelcarationContext) {}

// EnterNormalDeclarations is called when production normalDeclarations is entered.
func (s *BaseDefineListener) EnterNormalDeclarations(ctx *NormalDeclarationsContext) {}

// ExitNormalDeclarations is called when production normalDeclarations is exited.
func (s *BaseDefineListener) ExitNormalDeclarations(ctx *NormalDeclarationsContext) {}

// EnterDefaultDeclaration is called when production defaultDeclaration is entered.
func (s *BaseDefineListener) EnterDefaultDeclaration(ctx *DefaultDeclarationContext) {}

// ExitDefaultDeclaration is called when production defaultDeclaration is exited.
func (s *BaseDefineListener) ExitDefaultDeclaration(ctx *DefaultDeclarationContext) {}

// EnterSystemDeclaration is called when production systemDeclaration is entered.
func (s *BaseDefineListener) EnterSystemDeclaration(ctx *SystemDeclarationContext) {}

// ExitSystemDeclaration is called when production systemDeclaration is exited.
func (s *BaseDefineListener) ExitSystemDeclaration(ctx *SystemDeclarationContext) {}

// EnterDefineKey is called when production defineKey is entered.
func (s *BaseDefineListener) EnterDefineKey(ctx *DefineKeyContext) {}

// ExitDefineKey is called when production defineKey is exited.
func (s *BaseDefineListener) ExitDefineKey(ctx *DefineKeyContext) {}

// EnterDefineValue is called when production defineValue is entered.
func (s *BaseDefineListener) EnterDefineValue(ctx *DefineValueContext) {}

// ExitDefineValue is called when production defineValue is exited.
func (s *BaseDefineListener) ExitDefineValue(ctx *DefineValueContext) {}

// EnterModuleDeclaration is called when production moduleDeclaration is entered.
func (s *BaseDefineListener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {}

// ExitModuleDeclaration is called when production moduleDeclaration is exited.
func (s *BaseDefineListener) ExitModuleDeclaration(ctx *ModuleDeclarationContext) {}

// EnterModuleDefine is called when production moduleDefine is entered.
func (s *BaseDefineListener) EnterModuleDefine(ctx *ModuleDefineContext) {}

// ExitModuleDefine is called when production moduleDefine is exited.
func (s *BaseDefineListener) ExitModuleDefine(ctx *ModuleDefineContext) {}

// EnterModuleAttributes is called when production moduleAttributes is entered.
func (s *BaseDefineListener) EnterModuleAttributes(ctx *ModuleAttributesContext) {}

// ExitModuleAttributes is called when production moduleAttributes is exited.
func (s *BaseDefineListener) ExitModuleAttributes(ctx *ModuleAttributesContext) {}
