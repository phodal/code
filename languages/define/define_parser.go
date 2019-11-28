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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 142,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 2, 5, 2, 44, 10,
	2, 3, 2, 7, 2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 60, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7,
	5, 66, 10, 5, 12, 5, 14, 5, 69, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 75,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 97, 10,
	11, 12, 11, 14, 11, 100, 11, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 7, 16, 119, 10, 16, 12, 16, 14, 16, 122, 11, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 7, 17, 129, 10, 17, 12, 17, 14, 17, 132, 11, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 140, 10, 18, 3, 18, 2, 2, 19,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 2, 2,
	134, 2, 39, 3, 2, 2, 2, 4, 53, 3, 2, 2, 2, 6, 59, 3, 2, 2, 2, 8, 61, 3,
	2, 2, 2, 10, 74, 3, 2, 2, 2, 12, 76, 3, 2, 2, 2, 14, 81, 3, 2, 2, 2, 16,
	89, 3, 2, 2, 2, 18, 91, 3, 2, 2, 2, 20, 93, 3, 2, 2, 2, 22, 104, 3, 2,
	2, 2, 24, 106, 3, 2, 2, 2, 26, 110, 3, 2, 2, 2, 28, 112, 3, 2, 2, 2, 30,
	114, 3, 2, 2, 2, 32, 125, 3, 2, 2, 2, 34, 139, 3, 2, 2, 2, 36, 38, 5, 4,
	3, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 44, 5, 8, 5, 2,
	43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 48, 3, 2, 2, 2, 45, 47, 5,
	6, 4, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48,
	49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 7, 2, 2,
	3, 52, 3, 3, 2, 2, 2, 53, 54, 7, 3, 2, 2, 54, 55, 7, 10, 2, 2, 55, 56,
	7, 24, 2, 2, 56, 5, 3, 2, 2, 2, 57, 60, 5, 24, 13, 2, 58, 60, 5, 30, 16,
	2, 59, 57, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 7, 3, 2, 2, 2, 61, 62, 7,
	4, 2, 2, 62, 63, 5, 26, 14, 2, 63, 67, 7, 13, 2, 2, 64, 66, 5, 10, 6, 2,
	65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 71, 7, 14, 2, 2, 71,
	9, 3, 2, 2, 2, 72, 75, 5, 12, 7, 2, 73, 75, 5, 14, 8, 2, 74, 72, 3, 2,
	2, 2, 74, 73, 3, 2, 2, 2, 75, 11, 3, 2, 2, 2, 76, 77, 7, 5, 2, 2, 77, 78,
	7, 20, 2, 2, 78, 79, 5, 16, 9, 2, 79, 80, 5, 18, 10, 2, 80, 13, 3, 2, 2,
	2, 81, 82, 7, 6, 2, 2, 82, 83, 7, 11, 2, 2, 83, 84, 7, 10, 2, 2, 84, 85,
	7, 12, 2, 2, 85, 86, 7, 13, 2, 2, 86, 87, 5, 20, 11, 2, 87, 88, 7, 14,
	2, 2, 88, 15, 3, 2, 2, 2, 89, 90, 7, 10, 2, 2, 90, 17, 3, 2, 2, 2, 91,
	92, 7, 10, 2, 2, 92, 19, 3, 2, 2, 2, 93, 94, 5, 16, 9, 2, 94, 98, 7, 24,
	2, 2, 95, 97, 5, 16, 9, 2, 96, 95, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 101, 102, 5, 22, 12, 2, 102, 103, 5, 16, 9, 2, 103, 21, 3, 2, 2,
	2, 104, 105, 7, 10, 2, 2, 105, 23, 3, 2, 2, 2, 106, 107, 5, 26, 14, 2,
	107, 108, 7, 20, 2, 2, 108, 109, 5, 28, 15, 2, 109, 25, 3, 2, 2, 2, 110,
	111, 7, 10, 2, 2, 111, 27, 3, 2, 2, 2, 112, 113, 7, 10, 2, 2, 113, 29,
	3, 2, 2, 2, 114, 115, 7, 7, 2, 2, 115, 116, 7, 10, 2, 2, 116, 120, 7, 13,
	2, 2, 117, 119, 5, 32, 17, 2, 118, 117, 3, 2, 2, 2, 119, 122, 3, 2, 2,
	2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 123, 124, 7, 14, 2, 2, 124, 31, 3, 2, 2, 2, 125, 126,
	7, 10, 2, 2, 126, 130, 7, 13, 2, 2, 127, 129, 5, 34, 18, 2, 128, 127, 3,
	2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2,
	2, 131, 133, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 134, 7, 14, 2, 2, 134,
	33, 3, 2, 2, 2, 135, 136, 7, 8, 2, 2, 136, 140, 7, 24, 2, 2, 137, 138,
	7, 9, 2, 2, 138, 140, 7, 24, 2, 2, 139, 135, 3, 2, 2, 2, 139, 137, 3, 2,
	2, 2, 140, 35, 3, 2, 2, 2, 12, 39, 43, 48, 59, 67, 74, 98, 120, 130, 139,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'symbol'", "'define'", "'defaultSymbol'", "'defaultTemplate'", "'module'",
	"'import'", "'equal'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'",
	"','", "'.'", "':'",
}
var symbolicNames = []string{
	"", "SYMBOL_TEXT", "DEFINE", "DEFAULT_SYMBOL", "DEFAULT_TEMPLATE", "MODULE",
	"IMPORT", "EQUAL", "IDENTIFIER", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
	"LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "COLON", "WS", "COMMENT", "LINE_COMMENT",
	"STRING_LITERAL",
}

var ruleNames = []string{
	"compilationUnit", "symbolDeclaration", "normalDeclarations", "defineDeclaration",
	"defineExpress", "defineAttribute", "defineTemplate", "symbolKey", "symbolValue",
	"defineBody", "templateData", "systemDeclaration", "defineKey", "defineValue",
	"moduleDeclaration", "moduleDefines", "moduleAttribute",
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
	DefineParserEOF              = antlr.TokenEOF
	DefineParserSYMBOL_TEXT      = 1
	DefineParserDEFINE           = 2
	DefineParserDEFAULT_SYMBOL   = 3
	DefineParserDEFAULT_TEMPLATE = 4
	DefineParserMODULE           = 5
	DefineParserIMPORT           = 6
	DefineParserEQUAL            = 7
	DefineParserIDENTIFIER       = 8
	DefineParserLPAREN           = 9
	DefineParserRPAREN           = 10
	DefineParserLBRACE           = 11
	DefineParserRBRACE           = 12
	DefineParserLBRACK           = 13
	DefineParserRBRACK           = 14
	DefineParserSEMI             = 15
	DefineParserCOMMA            = 16
	DefineParserDOT              = 17
	DefineParserCOLON            = 18
	DefineParserWS               = 19
	DefineParserCOMMENT          = 20
	DefineParserLINE_COMMENT     = 21
	DefineParserSTRING_LITERAL   = 22
)

// DefineParser rules.
const (
	DefineParserRULE_compilationUnit    = 0
	DefineParserRULE_symbolDeclaration  = 1
	DefineParserRULE_normalDeclarations = 2
	DefineParserRULE_defineDeclaration  = 3
	DefineParserRULE_defineExpress      = 4
	DefineParserRULE_defineAttribute    = 5
	DefineParserRULE_defineTemplate     = 6
	DefineParserRULE_symbolKey          = 7
	DefineParserRULE_symbolValue        = 8
	DefineParserRULE_defineBody         = 9
	DefineParserRULE_templateData       = 10
	DefineParserRULE_systemDeclaration  = 11
	DefineParserRULE_defineKey          = 12
	DefineParserRULE_defineValue        = 13
	DefineParserRULE_moduleDeclaration  = 14
	DefineParserRULE_moduleDefines      = 15
	DefineParserRULE_moduleAttribute    = 16
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

func (s *CompilationUnitContext) AllSymbolDeclaration() []ISymbolDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolDeclarationContext)(nil)).Elem())
	var tst = make([]ISymbolDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) SymbolDeclaration(i int) ISymbolDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolDeclarationContext)
}

func (s *CompilationUnitContext) DefineDeclaration() IDefineDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineDeclarationContext)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserSYMBOL_TEXT {
		{
			p.SetState(34)
			p.SymbolDeclaration()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DefineParserDEFINE {
		{
			p.SetState(40)
			p.DefineDeclaration()
		}

	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserMODULE || _la == DefineParserIDENTIFIER {
		{
			p.SetState(43)
			p.NormalDeclarations()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(DefineParserEOF)
	}

	return localctx
}

// ISymbolDeclarationContext is an interface to support dynamic dispatch.
type ISymbolDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolDeclarationContext differentiates from other interfaces.
	IsSymbolDeclarationContext()
}

type SymbolDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolDeclarationContext() *SymbolDeclarationContext {
	var p = new(SymbolDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolDeclaration
	return p
}

func (*SymbolDeclarationContext) IsSymbolDeclarationContext() {}

func NewSymbolDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolDeclarationContext {
	var p = new(SymbolDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolDeclaration

	return p
}

func (s *SymbolDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolDeclarationContext) SYMBOL_TEXT() antlr.TerminalNode {
	return s.GetToken(DefineParserSYMBOL_TEXT, 0)
}

func (s *SymbolDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *SymbolDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolDeclaration(s)
	}
}

func (s *SymbolDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolDeclaration(s)
	}
}

func (s *SymbolDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolDeclaration() (localctx ISymbolDeclarationContext) {
	localctx = NewSymbolDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefineParserRULE_symbolDeclaration)

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
		p.SetState(51)
		p.Match(DefineParserSYMBOL_TEXT)
	}
	{
		p.SetState(52)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(53)
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

	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.SystemDeclaration()
		}

	case DefineParserMODULE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.ModuleDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefineDeclarationContext is an interface to support dynamic dispatch.
type IDefineDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineDeclarationContext differentiates from other interfaces.
	IsDefineDeclarationContext()
}

type DefineDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineDeclarationContext() *DefineDeclarationContext {
	var p = new(DefineDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineDeclaration
	return p
}

func (*DefineDeclarationContext) IsDefineDeclarationContext() {}

func NewDefineDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineDeclarationContext {
	var p = new(DefineDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineDeclaration

	return p
}

func (s *DefineDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineDeclarationContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFINE, 0)
}

func (s *DefineDeclarationContext) DefineKey() IDefineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineKeyContext)
}

func (s *DefineDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *DefineDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *DefineDeclarationContext) AllDefineExpress() []IDefineExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefineExpressContext)(nil)).Elem())
	var tst = make([]IDefineExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefineExpressContext)
		}
	}

	return tst
}

func (s *DefineDeclarationContext) DefineExpress(i int) IDefineExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefineExpressContext)
}

func (s *DefineDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineDeclaration(s)
	}
}

func (s *DefineDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineDeclaration(s)
	}
}

func (s *DefineDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineDeclaration() (localctx IDefineDeclarationContext) {
	localctx = NewDefineDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DefineParserRULE_defineDeclaration)
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
		p.SetState(59)
		p.Match(DefineParserDEFINE)
	}
	{
		p.SetState(60)
		p.DefineKey()
	}
	{
		p.SetState(61)
		p.Match(DefineParserLBRACE)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserDEFAULT_SYMBOL || _la == DefineParserDEFAULT_TEMPLATE {
		{
			p.SetState(62)
			p.DefineExpress()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IDefineExpressContext is an interface to support dynamic dispatch.
type IDefineExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineExpressContext differentiates from other interfaces.
	IsDefineExpressContext()
}

type DefineExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineExpressContext() *DefineExpressContext {
	var p = new(DefineExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineExpress
	return p
}

func (*DefineExpressContext) IsDefineExpressContext() {}

func NewDefineExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineExpressContext {
	var p = new(DefineExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineExpress

	return p
}

func (s *DefineExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineExpressContext) DefineAttribute() IDefineAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineAttributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineAttributeContext)
}

func (s *DefineExpressContext) DefineTemplate() IDefineTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineTemplateContext)
}

func (s *DefineExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineExpress(s)
	}
}

func (s *DefineExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineExpress(s)
	}
}

func (s *DefineExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineExpress() (localctx IDefineExpressContext) {
	localctx = NewDefineExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DefineParserRULE_defineExpress)

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
	case DefineParserDEFAULT_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.DefineAttribute()
		}

	case DefineParserDEFAULT_TEMPLATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.DefineTemplate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefineAttributeContext is an interface to support dynamic dispatch.
type IDefineAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineAttributeContext differentiates from other interfaces.
	IsDefineAttributeContext()
}

type DefineAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineAttributeContext() *DefineAttributeContext {
	var p = new(DefineAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineAttribute
	return p
}

func (*DefineAttributeContext) IsDefineAttributeContext() {}

func NewDefineAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineAttributeContext {
	var p = new(DefineAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineAttribute

	return p
}

func (s *DefineAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineAttributeContext) DEFAULT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFAULT_SYMBOL, 0)
}

func (s *DefineAttributeContext) COLON() antlr.TerminalNode {
	return s.GetToken(DefineParserCOLON, 0)
}

func (s *DefineAttributeContext) SymbolKey() ISymbolKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolKeyContext)
}

func (s *DefineAttributeContext) SymbolValue() ISymbolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolValueContext)
}

func (s *DefineAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineAttribute(s)
	}
}

func (s *DefineAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineAttribute(s)
	}
}

func (s *DefineAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineAttribute() (localctx IDefineAttributeContext) {
	localctx = NewDefineAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DefineParserRULE_defineAttribute)

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
		p.Match(DefineParserDEFAULT_SYMBOL)
	}
	{
		p.SetState(75)
		p.Match(DefineParserCOLON)
	}
	{
		p.SetState(76)
		p.SymbolKey()
	}
	{
		p.SetState(77)
		p.SymbolValue()
	}

	return localctx
}

// IDefineTemplateContext is an interface to support dynamic dispatch.
type IDefineTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineTemplateContext differentiates from other interfaces.
	IsDefineTemplateContext()
}

type DefineTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineTemplateContext() *DefineTemplateContext {
	var p = new(DefineTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineTemplate
	return p
}

func (*DefineTemplateContext) IsDefineTemplateContext() {}

func NewDefineTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineTemplateContext {
	var p = new(DefineTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineTemplate

	return p
}

func (s *DefineTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineTemplateContext) DEFAULT_TEMPLATE() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFAULT_TEMPLATE, 0)
}

func (s *DefineTemplateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DefineParserLPAREN, 0)
}

func (s *DefineTemplateContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineTemplateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DefineParserRPAREN, 0)
}

func (s *DefineTemplateContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *DefineTemplateContext) DefineBody() IDefineBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineBodyContext)
}

func (s *DefineTemplateContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *DefineTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineTemplate(s)
	}
}

func (s *DefineTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineTemplate(s)
	}
}

func (s *DefineTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineTemplate() (localctx IDefineTemplateContext) {
	localctx = NewDefineTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DefineParserRULE_defineTemplate)

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
		p.SetState(79)
		p.Match(DefineParserDEFAULT_TEMPLATE)
	}
	{
		p.SetState(80)
		p.Match(DefineParserLPAREN)
	}
	{
		p.SetState(81)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(82)
		p.Match(DefineParserRPAREN)
	}
	{
		p.SetState(83)
		p.Match(DefineParserLBRACE)
	}
	{
		p.SetState(84)
		p.DefineBody()
	}
	{
		p.SetState(85)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// ISymbolKeyContext is an interface to support dynamic dispatch.
type ISymbolKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolKeyContext differentiates from other interfaces.
	IsSymbolKeyContext()
}

type SymbolKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolKeyContext() *SymbolKeyContext {
	var p = new(SymbolKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolKey
	return p
}

func (*SymbolKeyContext) IsSymbolKeyContext() {}

func NewSymbolKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolKeyContext {
	var p = new(SymbolKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolKey

	return p
}

func (s *SymbolKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolKey(s)
	}
}

func (s *SymbolKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolKey(s)
	}
}

func (s *SymbolKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolKey() (localctx ISymbolKeyContext) {
	localctx = NewSymbolKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DefineParserRULE_symbolKey)

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
		p.SetState(87)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// ISymbolValueContext is an interface to support dynamic dispatch.
type ISymbolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolValueContext differentiates from other interfaces.
	IsSymbolValueContext()
}

type SymbolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolValueContext() *SymbolValueContext {
	var p = new(SymbolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolValue
	return p
}

func (*SymbolValueContext) IsSymbolValueContext() {}

func NewSymbolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolValueContext {
	var p = new(SymbolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolValue

	return p
}

func (s *SymbolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolValue(s)
	}
}

func (s *SymbolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolValue(s)
	}
}

func (s *SymbolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolValue() (localctx ISymbolValueContext) {
	localctx = NewSymbolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DefineParserRULE_symbolValue)

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
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IDefineBodyContext is an interface to support dynamic dispatch.
type IDefineBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineBodyContext differentiates from other interfaces.
	IsDefineBodyContext()
}

type DefineBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineBodyContext() *DefineBodyContext {
	var p = new(DefineBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineBody
	return p
}

func (*DefineBodyContext) IsDefineBodyContext() {}

func NewDefineBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineBodyContext {
	var p = new(DefineBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineBody

	return p
}

func (s *DefineBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineBodyContext) AllSymbolKey() []ISymbolKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem())
	var tst = make([]ISymbolKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolKeyContext)
		}
	}

	return tst
}

func (s *DefineBodyContext) SymbolKey(i int) ISymbolKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolKeyContext)
}

func (s *DefineBodyContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *DefineBodyContext) TemplateData() ITemplateDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateDataContext)
}

func (s *DefineBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineBody(s)
	}
}

func (s *DefineBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineBody(s)
	}
}

func (s *DefineBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineBody() (localctx IDefineBodyContext) {
	localctx = NewDefineBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DefineParserRULE_defineBody)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.SymbolKey()
	}
	{
		p.SetState(92)
		p.Match(DefineParserSTRING_LITERAL)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(93)
				p.SymbolKey()
			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	{
		p.SetState(99)
		p.TemplateData()
	}
	{
		p.SetState(100)
		p.SymbolKey()
	}

	return localctx
}

// ITemplateDataContext is an interface to support dynamic dispatch.
type ITemplateDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateDataContext differentiates from other interfaces.
	IsTemplateDataContext()
}

type TemplateDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateDataContext() *TemplateDataContext {
	var p = new(TemplateDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_templateData
	return p
}

func (*TemplateDataContext) IsTemplateDataContext() {}

func NewTemplateDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateDataContext {
	var p = new(TemplateDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_templateData

	return p
}

func (s *TemplateDataContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateDataContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *TemplateDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterTemplateData(s)
	}
}

func (s *TemplateDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitTemplateData(s)
	}
}

func (s *TemplateDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitTemplateData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) TemplateData() (localctx ITemplateDataContext) {
	localctx = NewTemplateDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DefineParserRULE_templateData)

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
		p.SetState(102)
		p.Match(DefineParserIDENTIFIER)
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

func (s *SystemDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(DefineParserCOLON, 0)
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
	p.EnterRule(localctx, 22, DefineParserRULE_systemDeclaration)

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
		p.SetState(104)
		p.DefineKey()
	}
	{
		p.SetState(105)
		p.Match(DefineParserCOLON)
	}
	{
		p.SetState(106)
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
	p.EnterRule(localctx, 24, DefineParserRULE_defineKey)

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
		p.SetState(108)
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
	p.EnterRule(localctx, 26, DefineParserRULE_defineValue)

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

func (s *ModuleDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDeclarationContext) AllModuleDefines() []IModuleDefinesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModuleDefinesContext)(nil)).Elem())
	var tst = make([]IModuleDefinesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModuleDefinesContext)
		}
	}

	return tst
}

func (s *ModuleDeclarationContext) ModuleDefines(i int) IModuleDefinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDefinesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModuleDefinesContext)
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
	p.EnterRule(localctx, 28, DefineParserRULE_moduleDeclaration)
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
		p.SetState(112)
		p.Match(DefineParserMODULE)
	}
	{
		p.SetState(113)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(114)
		p.Match(DefineParserLBRACE)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserIDENTIFIER {
		{
			p.SetState(115)
			p.ModuleDefines()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleDefinesContext is an interface to support dynamic dispatch.
type IModuleDefinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDefinesContext differentiates from other interfaces.
	IsModuleDefinesContext()
}

type ModuleDefinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDefinesContext() *ModuleDefinesContext {
	var p = new(ModuleDefinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleDefines
	return p
}

func (*ModuleDefinesContext) IsModuleDefinesContext() {}

func NewModuleDefinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefinesContext {
	var p = new(ModuleDefinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleDefines

	return p
}

func (s *ModuleDefinesContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefinesContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *ModuleDefinesContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *ModuleDefinesContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDefinesContext) AllModuleAttribute() []IModuleAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModuleAttributeContext)(nil)).Elem())
	var tst = make([]IModuleAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModuleAttributeContext)
		}
	}

	return tst
}

func (s *ModuleDefinesContext) ModuleAttribute(i int) IModuleAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModuleAttributeContext)
}

func (s *ModuleDefinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleDefines(s)
	}
}

func (s *ModuleDefinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleDefines(s)
	}
}

func (s *ModuleDefinesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleDefines(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleDefines() (localctx IModuleDefinesContext) {
	localctx = NewModuleDefinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DefineParserRULE_moduleDefines)
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
		p.SetState(123)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(124)
		p.Match(DefineParserLBRACE)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserIMPORT || _la == DefineParserEQUAL {
		{
			p.SetState(125)
			p.ModuleAttribute()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleAttributeContext is an interface to support dynamic dispatch.
type IModuleAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleAttributeContext differentiates from other interfaces.
	IsModuleAttributeContext()
}

type ModuleAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleAttributeContext() *ModuleAttributeContext {
	var p = new(ModuleAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleAttribute
	return p
}

func (*ModuleAttributeContext) IsModuleAttributeContext() {}

func NewModuleAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleAttributeContext {
	var p = new(ModuleAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleAttribute

	return p
}

func (s *ModuleAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleAttributeContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(DefineParserIMPORT, 0)
}

func (s *ModuleAttributeContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *ModuleAttributeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(DefineParserEQUAL, 0)
}

func (s *ModuleAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleAttribute(s)
	}
}

func (s *ModuleAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleAttribute(s)
	}
}

func (s *ModuleAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleAttribute() (localctx IModuleAttributeContext) {
	localctx = NewModuleAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DefineParserRULE_moduleAttribute)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Match(DefineParserIMPORT)
		}
		{
			p.SetState(134)
			p.Match(DefineParserSTRING_LITERAL)
		}

	case DefineParserEQUAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(DefineParserEQUAL)
		}
		{
			p.SetState(136)
			p.Match(DefineParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
