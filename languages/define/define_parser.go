// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 77, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 2, 5, 2, 30, 10, 2, 3, 2, 7, 2, 33, 10, 2,
	12, 2, 14, 2, 36, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	5, 4, 46, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 75, 10, 11, 3, 11,
	2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 71, 2, 25, 3, 2,
	2, 2, 4, 39, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 51,
	3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 57, 3, 2, 2, 2, 16, 59, 3, 2, 2, 2,
	18, 65, 3, 2, 2, 2, 20, 74, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2, 23, 22, 3,
	2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26,
	29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 30, 5, 8, 5, 2, 29, 28, 3, 2, 2,
	2, 29, 30, 3, 2, 2, 2, 30, 34, 3, 2, 2, 2, 31, 33, 5, 6, 4, 2, 32, 31,
	3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 37, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 7, 2, 2, 3, 38, 3, 3, 2,
	2, 2, 39, 40, 7, 4, 2, 2, 40, 41, 7, 9, 2, 2, 41, 42, 7, 22, 2, 2, 42,
	5, 3, 2, 2, 2, 43, 46, 5, 10, 6, 2, 44, 46, 5, 16, 9, 2, 45, 43, 3, 2,
	2, 2, 45, 44, 3, 2, 2, 2, 46, 7, 3, 2, 2, 2, 47, 48, 5, 12, 7, 2, 48, 49,
	7, 3, 2, 2, 49, 50, 5, 14, 8, 2, 50, 9, 3, 2, 2, 2, 51, 52, 5, 12, 7, 2,
	52, 53, 7, 3, 2, 2, 53, 54, 5, 14, 8, 2, 54, 11, 3, 2, 2, 2, 55, 56, 7,
	9, 2, 2, 56, 13, 3, 2, 2, 2, 57, 58, 7, 9, 2, 2, 58, 15, 3, 2, 2, 2, 59,
	60, 7, 6, 2, 2, 60, 61, 7, 9, 2, 2, 61, 62, 7, 12, 2, 2, 62, 63, 5, 18,
	10, 2, 63, 64, 7, 13, 2, 2, 64, 17, 3, 2, 2, 2, 65, 66, 7, 9, 2, 2, 66,
	67, 7, 12, 2, 2, 67, 68, 5, 20, 11, 2, 68, 69, 7, 13, 2, 2, 69, 19, 3,
	2, 2, 2, 70, 71, 7, 7, 2, 2, 71, 75, 7, 22, 2, 2, 72, 73, 7, 8, 2, 2, 73,
	75, 7, 22, 2, 2, 74, 70, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 21, 3, 2,
	2, 2, 7, 25, 29, 34, 45, 74,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'symbol'", "", "'module'", "'import'", "'equal'", "", "'('",
	"')'", "'{'", "'}'", "'['", "']'", "';'", "','", "'.'",
}
var symbolicNames = []string{
	"", "", "SYMBOL_TEXT", "SPECIAL_SYMBOL", "MODULE", "IMPORT", "EQUAL", "IDENTIFIER",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA",
	"DOT", "WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL",
}

var ruleNames = []string{
	"compilationUnit", "symbolDelcaration", "normalDeclarations", "defaultDeclaration",
	"systemDeclaration", "defineKey", "defineValue", "moduleDeclaration", "moduleDefine",
	"moduleAttributes",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DefineParser struct {
	*antlr.BaseParser
}

func NewDefineParser(input antlr.TokenStream) *DefineParser {
	this := new(DefineParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Define.g4"

	return this
}

// DefineParser tokens.
const (
	DefineParserEOF            = antlr.TokenEOF
	DefineParserT__0           = 1
	DefineParserSYMBOL_TEXT    = 2
	DefineParserSPECIAL_SYMBOL = 3
	DefineParserMODULE         = 4
	DefineParserIMPORT         = 5
	DefineParserEQUAL          = 6
	DefineParserIDENTIFIER     = 7
	DefineParserLPAREN         = 8
	DefineParserRPAREN         = 9
	DefineParserLBRACE         = 10
	DefineParserRBRACE         = 11
	DefineParserLBRACK         = 12
	DefineParserRBRACK         = 13
	DefineParserSEMI           = 14
	DefineParserCOMMA          = 15
	DefineParserDOT            = 16
	DefineParserWS             = 17
	DefineParserCOMMENT        = 18
	DefineParserLINE_COMMENT   = 19
	DefineParserSTRING_LITERAL = 20
)

// DefineParser rules.
const (
	DefineParserRULE_compilationUnit    = 0
	DefineParserRULE_symbolDelcaration  = 1
	DefineParserRULE_normalDeclarations = 2
	DefineParserRULE_defaultDeclaration = 3
	DefineParserRULE_systemDeclaration  = 4
	DefineParserRULE_defineKey          = 5
	DefineParserRULE_defineValue        = 6
	DefineParserRULE_moduleDeclaration  = 7
	DefineParserRULE_moduleDefine       = 8
	DefineParserRULE_moduleAttributes   = 9
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
	p.RuleIndex = DefineParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(DefineParserEOF, 0)
}

func (s *CompilationUnitContext) AllSymbolDelcaration() []ISymbolDelcarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolDelcarationContext)(nil)).Elem())
	var tst = make([]ISymbolDelcarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolDelcarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) SymbolDelcaration(i int) ISymbolDelcarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolDelcarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolDelcarationContext)
}

func (s *CompilationUnitContext) DefaultDeclaration() IDefaultDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultDeclarationContext)
}

func (s *CompilationUnitContext) AllNormalDeclarations() []INormalDeclarationsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INormalDeclarationsContext)(nil)).Elem())
	var tst = make([]INormalDeclarationsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INormalDeclarationsContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) NormalDeclarations(i int) INormalDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalDeclarationsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INormalDeclarationsContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DefineParserRULE_compilationUnit)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserSYMBOL_TEXT {
		{
			p.SetState(20)
			p.SymbolDelcaration()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(26)
			p.DefaultDeclaration()
		}

	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserMODULE || _la == DefineParserIDENTIFIER {
		{
			p.SetState(29)
			p.NormalDeclarations()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(DefineParserEOF)
	}

	return localctx
}

// ISymbolDelcarationContext is an interface to support dynamic dispatch.
type ISymbolDelcarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolDelcarationContext differentiates from other interfaces.
	IsSymbolDelcarationContext()
}

type SymbolDelcarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolDelcarationContext() *SymbolDelcarationContext {
	var p = new(SymbolDelcarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolDelcaration
	return p
}

func (*SymbolDelcarationContext) IsSymbolDelcarationContext() {}

func NewSymbolDelcarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolDelcarationContext {
	var p = new(SymbolDelcarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolDelcaration

	return p
}

func (s *SymbolDelcarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolDelcarationContext) SYMBOL_TEXT() antlr.TerminalNode {
	return s.GetToken(DefineParserSYMBOL_TEXT, 0)
}

func (s *SymbolDelcarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolDelcarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *SymbolDelcarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolDelcarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolDelcarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolDelcaration(s)
	}
}

func (s *SymbolDelcarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolDelcaration(s)
	}
}

func (s *SymbolDelcarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolDelcaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolDelcaration() (localctx ISymbolDelcarationContext) {
	localctx = NewSymbolDelcarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefineParserRULE_symbolDelcaration)

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
		p.SetState(37)
		p.Match(DefineParserSYMBOL_TEXT)
	}
	{
		p.SetState(38)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(39)
		p.Match(DefineParserSTRING_LITERAL)
	}

	return localctx
}

// INormalDeclarationsContext is an interface to support dynamic dispatch.
type INormalDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormalDeclarationsContext differentiates from other interfaces.
	IsNormalDeclarationsContext()
}

type NormalDeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalDeclarationsContext() *NormalDeclarationsContext {
	var p = new(NormalDeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_normalDeclarations
	return p
}

func (*NormalDeclarationsContext) IsNormalDeclarationsContext() {}

func NewNormalDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalDeclarationsContext {
	var p = new(NormalDeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_normalDeclarations

	return p
}

func (s *NormalDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalDeclarationsContext) SystemDeclaration() ISystemDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISystemDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISystemDeclarationContext)
}

func (s *NormalDeclarationsContext) ModuleDeclaration() IModuleDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDeclarationContext)
}

func (s *NormalDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterNormalDeclarations(s)
	}
}

func (s *NormalDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitNormalDeclarations(s)
	}
}

func (s *NormalDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitNormalDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) NormalDeclarations() (localctx INormalDeclarationsContext) {
	localctx = NewNormalDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DefineParserRULE_normalDeclarations)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.SystemDeclaration()
		}

	case DefineParserMODULE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.ModuleDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefaultDeclarationContext is an interface to support dynamic dispatch.
type IDefaultDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultDeclarationContext differentiates from other interfaces.
	IsDefaultDeclarationContext()
}

type DefaultDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultDeclarationContext() *DefaultDeclarationContext {
	var p = new(DefaultDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defaultDeclaration
	return p
}

func (*DefaultDeclarationContext) IsDefaultDeclarationContext() {}

func NewDefaultDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultDeclarationContext {
	var p = new(DefaultDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defaultDeclaration

	return p
}

func (s *DefaultDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultDeclarationContext) DefineKey() IDefineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineKeyContext)
}

func (s *DefaultDeclarationContext) DefineValue() IDefineValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineValueContext)
}

func (s *DefaultDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefaultDeclaration(s)
	}
}

func (s *DefaultDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefaultDeclaration(s)
	}
}

func (s *DefaultDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefaultDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefaultDeclaration() (localctx IDefaultDeclarationContext) {
	localctx = NewDefaultDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DefineParserRULE_defaultDeclaration)

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
		p.SetState(45)
		p.DefineKey()
	}
	{
		p.SetState(46)
		p.Match(DefineParserT__0)
	}
	{
		p.SetState(47)
		p.DefineValue()
	}

	return localctx
}

// ISystemDeclarationContext is an interface to support dynamic dispatch.
type ISystemDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystemDeclarationContext differentiates from other interfaces.
	IsSystemDeclarationContext()
}

type SystemDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemDeclarationContext() *SystemDeclarationContext {
	var p = new(SystemDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_systemDeclaration
	return p
}

func (*SystemDeclarationContext) IsSystemDeclarationContext() {}

func NewSystemDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemDeclarationContext {
	var p = new(SystemDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_systemDeclaration

	return p
}

func (s *SystemDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemDeclarationContext) DefineKey() IDefineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineKeyContext)
}

func (s *SystemDeclarationContext) DefineValue() IDefineValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineValueContext)
}

func (s *SystemDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSystemDeclaration(s)
	}
}

func (s *SystemDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSystemDeclaration(s)
	}
}

func (s *SystemDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSystemDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SystemDeclaration() (localctx ISystemDeclarationContext) {
	localctx = NewSystemDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DefineParserRULE_systemDeclaration)

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
		p.SetState(49)
		p.DefineKey()
	}
	{
		p.SetState(50)
		p.Match(DefineParserT__0)
	}
	{
		p.SetState(51)
		p.DefineValue()
	}

	return localctx
}

// IDefineKeyContext is an interface to support dynamic dispatch.
type IDefineKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineKeyContext differentiates from other interfaces.
	IsDefineKeyContext()
}

type DefineKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineKeyContext() *DefineKeyContext {
	var p = new(DefineKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineKey
	return p
}

func (*DefineKeyContext) IsDefineKeyContext() {}

func NewDefineKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineKeyContext {
	var p = new(DefineKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineKey

	return p
}

func (s *DefineKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineKey(s)
	}
}

func (s *DefineKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineKey(s)
	}
}

func (s *DefineKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineKey() (localctx IDefineKeyContext) {
	localctx = NewDefineKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DefineParserRULE_defineKey)

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
		p.SetState(53)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IDefineValueContext is an interface to support dynamic dispatch.
type IDefineValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineValueContext differentiates from other interfaces.
	IsDefineValueContext()
}

type DefineValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineValueContext() *DefineValueContext {
	var p = new(DefineValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineValue
	return p
}

func (*DefineValueContext) IsDefineValueContext() {}

func NewDefineValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineValueContext {
	var p = new(DefineValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineValue

	return p
}

func (s *DefineValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineValue(s)
	}
}

func (s *DefineValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineValue(s)
	}
}

func (s *DefineValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineValue() (localctx IDefineValueContext) {
	localctx = NewDefineValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DefineParserRULE_defineValue)

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
		p.SetState(55)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IModuleDeclarationContext is an interface to support dynamic dispatch.
type IModuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDeclarationContext differentiates from other interfaces.
	IsModuleDeclarationContext()
}

type ModuleDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDeclarationContext() *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleDeclaration
	return p
}

func (*ModuleDeclarationContext) IsModuleDeclarationContext() {}

func NewModuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleDeclaration

	return p
}

func (s *ModuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclarationContext) MODULE() antlr.TerminalNode {
	return s.GetToken(DefineParserMODULE, 0)
}

func (s *ModuleDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *ModuleDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *ModuleDeclarationContext) ModuleDefine() IModuleDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDefineContext)
}

func (s *ModuleDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleDeclaration() (localctx IModuleDeclarationContext) {
	localctx = NewModuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DefineParserRULE_moduleDeclaration)

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
		p.Match(DefineParserMODULE)
	}
	{
		p.SetState(58)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(59)
		p.Match(DefineParserLBRACE)
	}
	{
		p.SetState(60)
		p.ModuleDefine()
	}
	{
		p.SetState(61)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleDefineContext is an interface to support dynamic dispatch.
type IModuleDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDefineContext differentiates from other interfaces.
	IsModuleDefineContext()
}

type ModuleDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDefineContext() *ModuleDefineContext {
	var p = new(ModuleDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleDefine
	return p
}

func (*ModuleDefineContext) IsModuleDefineContext() {}

func NewModuleDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefineContext {
	var p = new(ModuleDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleDefine

	return p
}

func (s *ModuleDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefineContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *ModuleDefineContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *ModuleDefineContext) ModuleAttributes() IModuleAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleAttributesContext)
}

func (s *ModuleDefineContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleDefine(s)
	}
}

func (s *ModuleDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleDefine(s)
	}
}

func (s *ModuleDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleDefine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleDefine() (localctx IModuleDefineContext) {
	localctx = NewModuleDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DefineParserRULE_moduleDefine)

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
		p.SetState(63)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(64)
		p.Match(DefineParserLBRACE)
	}
	{
		p.SetState(65)
		p.ModuleAttributes()
	}
	{
		p.SetState(66)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleAttributesContext is an interface to support dynamic dispatch.
type IModuleAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleAttributesContext differentiates from other interfaces.
	IsModuleAttributesContext()
}

type ModuleAttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleAttributesContext() *ModuleAttributesContext {
	var p = new(ModuleAttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleAttributes
	return p
}

func (*ModuleAttributesContext) IsModuleAttributesContext() {}

func NewModuleAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleAttributesContext {
	var p = new(ModuleAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleAttributes

	return p
}

func (s *ModuleAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleAttributesContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(DefineParserIMPORT, 0)
}

func (s *ModuleAttributesContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *ModuleAttributesContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(DefineParserEQUAL, 0)
}

func (s *ModuleAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleAttributes(s)
	}
}

func (s *ModuleAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleAttributes(s)
	}
}

func (s *ModuleAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleAttributes() (localctx IModuleAttributesContext) {
	localctx = NewModuleAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DefineParserRULE_moduleAttributes)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(DefineParserIMPORT)
		}
		{
			p.SetState(69)
			p.Match(DefineParserSTRING_LITERAL)
		}

	case DefineParserEQUAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(DefineParserEQUAL)
		}
		{
			p.SetState(71)
			p.Match(DefineParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
