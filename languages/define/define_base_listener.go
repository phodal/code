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

// EnterSymbolDeclaration is called when production symbolDeclaration is entered.
func (s *BaseDefineListener) EnterSymbolDeclaration(ctx *SymbolDeclarationContext) {}

// ExitSymbolDeclaration is called when production symbolDeclaration is exited.
func (s *BaseDefineListener) ExitSymbolDeclaration(ctx *SymbolDeclarationContext) {}

// EnterNormalDeclarations is called when production normalDeclarations is entered.
func (s *BaseDefineListener) EnterNormalDeclarations(ctx *NormalDeclarationsContext) {}

// ExitNormalDeclarations is called when production normalDeclarations is exited.
func (s *BaseDefineListener) ExitNormalDeclarations(ctx *NormalDeclarationsContext) {}

// EnterDefineDeclaration is called when production defineDeclaration is entered.
func (s *BaseDefineListener) EnterDefineDeclaration(ctx *DefineDeclarationContext) {}

// ExitDefineDeclaration is called when production defineDeclaration is exited.
func (s *BaseDefineListener) ExitDefineDeclaration(ctx *DefineDeclarationContext) {}

// EnterDefineExpress is called when production defineExpress is entered.
func (s *BaseDefineListener) EnterDefineExpress(ctx *DefineExpressContext) {}

// ExitDefineExpress is called when production defineExpress is exited.
func (s *BaseDefineListener) ExitDefineExpress(ctx *DefineExpressContext) {}

// EnterDefineAttribute is called when production defineAttribute is entered.
func (s *BaseDefineListener) EnterDefineAttribute(ctx *DefineAttributeContext) {}

// ExitDefineAttribute is called when production defineAttribute is exited.
func (s *BaseDefineListener) ExitDefineAttribute(ctx *DefineAttributeContext) {}

// EnterDefineTemplate is called when production defineTemplate is entered.
func (s *BaseDefineListener) EnterDefineTemplate(ctx *DefineTemplateContext) {}

// ExitDefineTemplate is called when production defineTemplate is exited.
func (s *BaseDefineListener) ExitDefineTemplate(ctx *DefineTemplateContext) {}

// EnterSymbolKey is called when production symbolKey is entered.
func (s *BaseDefineListener) EnterSymbolKey(ctx *SymbolKeyContext) {}

// ExitSymbolKey is called when production symbolKey is exited.
func (s *BaseDefineListener) ExitSymbolKey(ctx *SymbolKeyContext) {}

// EnterSymbolValue is called when production symbolValue is entered.
func (s *BaseDefineListener) EnterSymbolValue(ctx *SymbolValueContext) {}

// ExitSymbolValue is called when production symbolValue is exited.
func (s *BaseDefineListener) ExitSymbolValue(ctx *SymbolValueContext) {}

// EnterDefineBody is called when production defineBody is entered.
func (s *BaseDefineListener) EnterDefineBody(ctx *DefineBodyContext) {}

// ExitDefineBody is called when production defineBody is exited.
func (s *BaseDefineListener) ExitDefineBody(ctx *DefineBodyContext) {}

// EnterTemplateData is called when production templateData is entered.
func (s *BaseDefineListener) EnterTemplateData(ctx *TemplateDataContext) {}

// ExitTemplateData is called when production templateData is exited.
func (s *BaseDefineListener) ExitTemplateData(ctx *TemplateDataContext) {}

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
