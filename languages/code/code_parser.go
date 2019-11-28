// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 119,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 5, 2, 38, 10, 2, 3, 2, 7, 2, 41, 10, 2, 12, 2, 14, 2, 44, 11,
	2, 3, 2, 7, 2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 2, 7, 2, 53, 10,
	2, 12, 2, 14, 2, 56, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 5, 5, 69, 10, 5, 3, 6, 3, 6, 5, 6, 73, 10, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85,
	11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 90, 10, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 101, 10, 14, 12, 14, 14,
	14, 104, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 4, 4, 2, 12, 12, 17, 17,
	4, 2, 12, 12, 16, 17, 2, 111, 2, 37, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2, 6,
	62, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12, 74, 3, 2, 2,
	2, 14, 76, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 91,
	3, 2, 2, 2, 22, 93, 3, 2, 2, 2, 24, 95, 3, 2, 2, 2, 26, 97, 3, 2, 2, 2,
	28, 107, 3, 2, 2, 2, 30, 109, 3, 2, 2, 2, 32, 112, 3, 2, 2, 2, 34, 115,
	3, 2, 2, 2, 36, 38, 5, 4, 3, 2, 37, 36, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2,
	38, 42, 3, 2, 2, 2, 39, 41, 5, 6, 4, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3,
	2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 48, 3, 2, 2, 2, 44,
	42, 3, 2, 2, 2, 45, 47, 5, 8, 5, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2,
	2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 54, 3, 2, 2, 2, 50, 48,
	3, 2, 2, 2, 51, 53, 5, 10, 6, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2,
	54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54, 3,
	2, 2, 2, 57, 58, 7, 2, 2, 3, 58, 3, 3, 2, 2, 2, 59, 60, 7, 7, 2, 2, 60,
	61, 7, 12, 2, 2, 61, 5, 3, 2, 2, 2, 62, 63, 7, 8, 2, 2, 63, 64, 7, 12,
	2, 2, 64, 7, 3, 2, 2, 2, 65, 69, 5, 30, 16, 2, 66, 69, 5, 32, 17, 2, 67,
	69, 5, 34, 18, 2, 68, 65, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2,
	2, 2, 69, 9, 3, 2, 2, 2, 70, 73, 5, 26, 14, 2, 71, 73, 5, 12, 7, 2, 72,
	70, 3, 2, 2, 2, 72, 71, 3, 2, 2, 2, 73, 11, 3, 2, 2, 2, 74, 75, 5, 14,
	8, 2, 75, 13, 3, 2, 2, 2, 76, 77, 5, 16, 9, 2, 77, 15, 3, 2, 2, 2, 78,
	83, 5, 18, 10, 2, 79, 80, 7, 3, 2, 2, 80, 82, 5, 18, 10, 2, 81, 79, 3,
	2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84,
	17, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 89, 5, 20, 11, 2, 87, 88, 7, 4,
	2, 2, 88, 90, 5, 22, 12, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90,
	19, 3, 2, 2, 2, 91, 92, 7, 12, 2, 2, 92, 21, 3, 2, 2, 2, 93, 94, 5, 24,
	13, 2, 94, 23, 3, 2, 2, 2, 95, 96, 9, 2, 2, 2, 96, 25, 3, 2, 2, 2, 97,
	98, 7, 12, 2, 2, 98, 102, 7, 5, 2, 2, 99, 101, 5, 28, 15, 2, 100, 99, 3,
	2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2,
	2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 7, 6, 2, 2, 106,
	27, 3, 2, 2, 2, 107, 108, 9, 3, 2, 2, 108, 29, 3, 2, 2, 2, 109, 110, 7,
	9, 2, 2, 110, 111, 7, 12, 2, 2, 111, 31, 3, 2, 2, 2, 112, 113, 7, 10, 2,
	2, 113, 114, 7, 12, 2, 2, 114, 33, 3, 2, 2, 2, 115, 116, 7, 11, 2, 2, 116,
	117, 7, 12, 2, 2, 117, 35, 3, 2, 2, 2, 11, 37, 42, 48, 54, 68, 72, 83,
	89, 102,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'='", "'('", "')'", "'package'", "'import'", "'struct'", "'member'",
	"'function'",
}
var symbolicNames = []string{
	"", "", "", "", "", "PACKAGE", "IMPORT", "DATA_STRUCT", "MEMBER", "FUNCTION",
	"IDENTIFIER", "WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL", "DECIMAL_LITERAL",
}

var ruleNames = []string{
	"compilationUnit", "packageDeclaration", "importDeclaration", "typeDeclaration",
	"expressDeclaration", "blockStatement", "localVariableDeclaration", "variableDeclarators",
	"variableDeclarator", "variableDeclaratorId", "variableInitializer", "expression",
	"methodCallDeclaration", "parameter", "dataStructDeclaration", "memberDeclaration",
	"functionDeclaration",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CodeParser struct {
	*antlr.BaseParser
}

func NewCodeParser(input antlr.TokenStream) *CodeParser {
	this := new(CodeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Code.g4"

	return this
}

// CodeParser tokens.
const (
	CodeParserEOF             = antlr.TokenEOF
	CodeParserT__0            = 1
	CodeParserT__1            = 2
	CodeParserT__2            = 3
	CodeParserT__3            = 4
	CodeParserPACKAGE         = 5
	CodeParserIMPORT          = 6
	CodeParserDATA_STRUCT     = 7
	CodeParserMEMBER          = 8
	CodeParserFUNCTION        = 9
	CodeParserIDENTIFIER      = 10
	CodeParserWS              = 11
	CodeParserCOMMENT         = 12
	CodeParserLINE_COMMENT    = 13
	CodeParserSTRING_LITERAL  = 14
	CodeParserDECIMAL_LITERAL = 15
)

// CodeParser rules.
const (
	CodeParserRULE_compilationUnit          = 0
	CodeParserRULE_packageDeclaration       = 1
	CodeParserRULE_importDeclaration        = 2
	CodeParserRULE_typeDeclaration          = 3
	CodeParserRULE_expressDeclaration       = 4
	CodeParserRULE_blockStatement           = 5
	CodeParserRULE_localVariableDeclaration = 6
	CodeParserRULE_variableDeclarators      = 7
	CodeParserRULE_variableDeclarator       = 8
	CodeParserRULE_variableDeclaratorId     = 9
	CodeParserRULE_variableInitializer      = 10
	CodeParserRULE_expression               = 11
	CodeParserRULE_methodCallDeclaration    = 12
	CodeParserRULE_parameter                = 13
	CodeParserRULE_dataStructDeclaration    = 14
	CodeParserRULE_memberDeclaration        = 15
	CodeParserRULE_functionDeclaration      = 16
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(CodeParserEOF, 0)
}

func (s *CompilationUnitContext) PackageDeclaration() IPackageDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageDeclarationContext)
}

func (s *CompilationUnitContext) AllImportDeclaration() []IImportDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem())
	var tst = make([]IImportDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) ImportDeclaration(i int) IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *CompilationUnitContext) AllTypeDeclaration() []ITypeDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem())
	var tst = make([]ITypeDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *CompilationUnitContext) AllExpressDeclaration() []IExpressDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem())
	var tst = make([]IExpressDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) ExpressDeclaration(i int) IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CodeParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CodeParserPACKAGE {
		{
			p.SetState(34)
			p.PackageDeclaration()
		}

	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIMPORT {
		{
			p.SetState(37)
			p.ImportDeclaration()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserDATA_STRUCT)|(1<<CodeParserMEMBER)|(1<<CodeParserFUNCTION))) != 0 {
		{
			p.SetState(43)
			p.TypeDeclaration()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIDENTIFIER {
		{
			p.SetState(49)
			p.ExpressDeclaration()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(CodeParserEOF)
	}

	return localctx
}

// IPackageDeclarationContext is an interface to support dynamic dispatch.
type IPackageDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageDeclarationContext differentiates from other interfaces.
	IsPackageDeclarationContext()
}

type PackageDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDeclarationContext() *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_packageDeclaration
	return p
}

func (*PackageDeclarationContext) IsPackageDeclarationContext() {}

func NewPackageDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_packageDeclaration

	return p
}

func (s *PackageDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDeclarationContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(CodeParserPACKAGE, 0)
}

func (s *PackageDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *PackageDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitPackageDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) PackageDeclaration() (localctx IPackageDeclarationContext) {
	localctx = NewPackageDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CodeParserRULE_packageDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(CodeParserPACKAGE)
	}
	{
		p.SetState(58)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(CodeParserIMPORT, 0)
}

func (s *ImportDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CodeParserRULE_importDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(CodeParserIMPORT)
	}
	{
		p.SetState(61)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) DataStructDeclaration() IDataStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataStructDeclarationContext)
}

func (s *TypeDeclarationContext) MemberDeclaration() IMemberDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberDeclarationContext)
}

func (s *TypeDeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CodeParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CodeParserDATA_STRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.DataStructDeclaration()
		}

	case CodeParserMEMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.MemberDeclaration()
		}

	case CodeParserFUNCTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.FunctionDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressDeclarationContext is an interface to support dynamic dispatch.
type IExpressDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressDeclarationContext differentiates from other interfaces.
	IsExpressDeclarationContext()
}

type ExpressDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressDeclarationContext() *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_expressDeclaration
	return p
}

func (*ExpressDeclarationContext) IsExpressDeclarationContext() {}

func NewExpressDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_expressDeclaration

	return p
}

func (s *ExpressDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressDeclarationContext) MethodCallDeclaration() IMethodCallDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallDeclarationContext)
}

func (s *ExpressDeclarationContext) BlockStatement() IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ExpressDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitExpressDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ExpressDeclaration() (localctx IExpressDeclarationContext) {
	localctx = NewExpressDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CodeParserRULE_expressDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.MethodCallDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.BlockStatement()
		}

	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) LocalVariableDeclaration() ILocalVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalVariableDeclarationContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CodeParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.LocalVariableDeclaration()
	}

	return localctx
}

// ILocalVariableDeclarationContext is an interface to support dynamic dispatch.
type ILocalVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalVariableDeclarationContext differentiates from other interfaces.
	IsLocalVariableDeclarationContext()
}

type LocalVariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalVariableDeclarationContext() *LocalVariableDeclarationContext {
	var p = new(LocalVariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_localVariableDeclaration
	return p
}

func (*LocalVariableDeclarationContext) IsLocalVariableDeclarationContext() {}

func NewLocalVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalVariableDeclarationContext {
	var p = new(LocalVariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_localVariableDeclaration

	return p
}

func (s *LocalVariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalVariableDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *LocalVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalVariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterLocalVariableDeclaration(s)
	}
}

func (s *LocalVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitLocalVariableDeclaration(s)
	}
}

func (s *LocalVariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitLocalVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) LocalVariableDeclaration() (localctx ILocalVariableDeclarationContext) {
	localctx = NewLocalVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CodeParserRULE_localVariableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.VariableDeclarators()
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CodeParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.VariableDeclarator()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserT__0 {
		{
			p.SetState(77)
			p.Match(CodeParserT__0)
		}
		{
			p.SetState(78)
			p.VariableDeclarator()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CodeParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.VariableDeclaratorId()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CodeParserT__1 {
		{
			p.SetState(85)
			p.Match(CodeParserT__1)
		}
		{
			p.SetState(86)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CodeParserRULE_variableDeclaratorId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CodeParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *ExpressionContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserDECIMAL_LITERAL, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CodeParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CodeParserIDENTIFIER || _la == CodeParserDECIMAL_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMethodCallDeclarationContext is an interface to support dynamic dispatch.
type IMethodCallDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallDeclarationContext differentiates from other interfaces.
	IsMethodCallDeclarationContext()
}

type MethodCallDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallDeclarationContext() *MethodCallDeclarationContext {
	var p = new(MethodCallDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_methodCallDeclaration
	return p
}

func (*MethodCallDeclarationContext) IsMethodCallDeclarationContext() {}

func NewMethodCallDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallDeclarationContext {
	var p = new(MethodCallDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_methodCallDeclaration

	return p
}

func (s *MethodCallDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *MethodCallDeclarationContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *MethodCallDeclarationContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *MethodCallDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterMethodCallDeclaration(s)
	}
}

func (s *MethodCallDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitMethodCallDeclaration(s)
	}
}

func (s *MethodCallDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitMethodCallDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) MethodCallDeclaration() (localctx IMethodCallDeclarationContext) {
	localctx = NewMethodCallDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CodeParserRULE_methodCallDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(CodeParserIDENTIFIER)
	}
	{
		p.SetState(96)
		p.Match(CodeParserT__2)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserIDENTIFIER)|(1<<CodeParserSTRING_LITERAL)|(1<<CodeParserDECIMAL_LITERAL))) != 0 {
		{
			p.SetState(97)
			p.Parameter()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(CodeParserT__3)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *ParameterContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserSTRING_LITERAL, 0)
}

func (s *ParameterContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserDECIMAL_LITERAL, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CodeParserRULE_parameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserIDENTIFIER)|(1<<CodeParserSTRING_LITERAL)|(1<<CodeParserDECIMAL_LITERAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDataStructDeclarationContext is an interface to support dynamic dispatch.
type IDataStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataStructDeclarationContext differentiates from other interfaces.
	IsDataStructDeclarationContext()
}

type DataStructDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataStructDeclarationContext() *DataStructDeclarationContext {
	var p = new(DataStructDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_dataStructDeclaration
	return p
}

func (*DataStructDeclarationContext) IsDataStructDeclarationContext() {}

func NewDataStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataStructDeclarationContext {
	var p = new(DataStructDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_dataStructDeclaration

	return p
}

func (s *DataStructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DataStructDeclarationContext) DATA_STRUCT() antlr.TerminalNode {
	return s.GetToken(CodeParserDATA_STRUCT, 0)
}

func (s *DataStructDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *DataStructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataStructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataStructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterDataStructDeclaration(s)
	}
}

func (s *DataStructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitDataStructDeclaration(s)
	}
}

func (s *DataStructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitDataStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) DataStructDeclaration() (localctx IDataStructDeclarationContext) {
	localctx = NewDataStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CodeParserRULE_dataStructDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(CodeParserDATA_STRUCT)
	}
	{
		p.SetState(108)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IMemberDeclarationContext is an interface to support dynamic dispatch.
type IMemberDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDeclarationContext differentiates from other interfaces.
	IsMemberDeclarationContext()
}

type MemberDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDeclarationContext() *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_memberDeclaration
	return p
}

func (*MemberDeclarationContext) IsMemberDeclarationContext() {}

func NewMemberDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_memberDeclaration

	return p
}

func (s *MemberDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDeclarationContext) MEMBER() antlr.TerminalNode {
	return s.GetToken(CodeParserMEMBER, 0)
}

func (s *MemberDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *MemberDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitMemberDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) MemberDeclaration() (localctx IMemberDeclarationContext) {
	localctx = NewMemberDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CodeParserRULE_memberDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(CodeParserMEMBER)
	}
	{
		p.SetState(111)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(CodeParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CodeParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(CodeParserFUNCTION)
	}
	{
		p.SetState(114)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}
