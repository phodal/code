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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 185,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 7, 2, 51, 10, 2, 12, 2, 14, 2, 54, 11, 2,
	3, 2, 7, 2, 57, 10, 2, 12, 2, 14, 2, 60, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 74, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 80, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 7, 7, 89, 10, 7, 12, 7, 14, 7, 92, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 104, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 114, 10, 9, 3, 9, 7, 9, 117, 10, 9,
	12, 9, 14, 9, 120, 11, 9, 3, 10, 3, 10, 7, 10, 124, 10, 10, 12, 10, 14,
	10, 127, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 137, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 147, 10, 14, 12, 14, 14, 14, 150, 11, 14, 3, 15, 3, 15, 3, 15,
	5, 15, 155, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 7, 19, 166, 10, 19, 12, 19, 14, 19, 169, 11, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 5, 20, 175, 10, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 2, 3, 16, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 5, 3, 2, 9, 20,
	4, 2, 29, 29, 33, 34, 3, 2, 33, 34, 2, 183, 2, 47, 3, 2, 2, 2, 4, 63, 3,
	2, 2, 2, 6, 66, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12,
	86, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2, 16, 103, 3, 2, 2, 2, 18, 121, 3, 2,
	2, 2, 20, 136, 3, 2, 2, 2, 22, 138, 3, 2, 2, 2, 24, 140, 3, 2, 2, 2, 26,
	142, 3, 2, 2, 2, 28, 151, 3, 2, 2, 2, 30, 156, 3, 2, 2, 2, 32, 158, 3,
	2, 2, 2, 34, 160, 3, 2, 2, 2, 36, 162, 3, 2, 2, 2, 38, 174, 3, 2, 2, 2,
	40, 176, 3, 2, 2, 2, 42, 178, 3, 2, 2, 2, 44, 181, 3, 2, 2, 2, 46, 48,
	5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 52, 3, 2, 2, 2,
	49, 51, 5, 6, 4, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3,
	2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 58, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55,
	57, 5, 8, 5, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2,
	2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62,
	7, 2, 2, 3, 62, 3, 3, 2, 2, 2, 63, 64, 7, 24, 2, 2, 64, 65, 7, 29, 2, 2,
	65, 5, 3, 2, 2, 2, 66, 67, 7, 25, 2, 2, 67, 68, 7, 29, 2, 2, 68, 7, 3,
	2, 2, 2, 69, 74, 5, 42, 22, 2, 70, 74, 5, 44, 23, 2, 71, 74, 5, 10, 6,
	2, 72, 74, 5, 16, 9, 2, 73, 69, 3, 2, 2, 2, 73, 70, 3, 2, 2, 2, 73, 71,
	3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 9, 3, 2, 2, 2, 75, 76, 7, 28, 2, 2,
	76, 77, 7, 29, 2, 2, 77, 79, 7, 3, 2, 2, 78, 80, 5, 38, 20, 2, 79, 78,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 7, 4, 2, 2,
	82, 83, 7, 5, 2, 2, 83, 84, 5, 12, 7, 2, 84, 85, 7, 6, 2, 2, 85, 11, 3,
	2, 2, 2, 86, 90, 5, 16, 9, 2, 87, 89, 5, 16, 9, 2, 88, 87, 3, 2, 2, 2,
	89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 13, 3,
	2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 29, 2, 2, 94, 15, 3, 2, 2, 2, 95,
	96, 8, 9, 1, 2, 96, 104, 5, 36, 19, 2, 97, 104, 5, 20, 11, 2, 98, 104,
	5, 14, 8, 2, 99, 100, 5, 34, 18, 2, 100, 101, 9, 2, 2, 2, 101, 102, 5,
	34, 18, 2, 102, 104, 3, 2, 2, 2, 103, 95, 3, 2, 2, 2, 103, 97, 3, 2, 2,
	2, 103, 98, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 104, 118, 3, 2, 2, 2, 105,
	113, 12, 4, 2, 2, 106, 107, 7, 7, 2, 2, 107, 114, 7, 7, 2, 2, 108, 109,
	7, 8, 2, 2, 109, 110, 7, 8, 2, 2, 110, 114, 7, 8, 2, 2, 111, 112, 7, 8,
	2, 2, 112, 114, 7, 8, 2, 2, 113, 106, 3, 2, 2, 2, 113, 108, 3, 2, 2, 2,
	113, 111, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 5, 16, 9, 5, 116,
	105, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 17, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 125, 7, 5,
	2, 2, 122, 124, 5, 20, 11, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2,
	2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128, 3, 2, 2, 2, 127,
	125, 3, 2, 2, 2, 128, 129, 7, 6, 2, 2, 129, 19, 3, 2, 2, 2, 130, 137, 5,
	18, 10, 2, 131, 137, 5, 24, 13, 2, 132, 133, 7, 22, 2, 2, 133, 134, 5,
	22, 12, 2, 134, 135, 5, 20, 11, 2, 135, 137, 3, 2, 2, 2, 136, 130, 3, 2,
	2, 2, 136, 131, 3, 2, 2, 2, 136, 132, 3, 2, 2, 2, 137, 21, 3, 2, 2, 2,
	138, 139, 5, 16, 9, 2, 139, 23, 3, 2, 2, 2, 140, 141, 5, 26, 14, 2, 141,
	25, 3, 2, 2, 2, 142, 143, 7, 23, 2, 2, 143, 148, 5, 28, 15, 2, 144, 145,
	7, 21, 2, 2, 145, 147, 5, 28, 15, 2, 146, 144, 3, 2, 2, 2, 147, 150, 3,
	2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 27, 3, 2, 2,
	2, 150, 148, 3, 2, 2, 2, 151, 154, 5, 30, 16, 2, 152, 153, 7, 9, 2, 2,
	153, 155, 5, 32, 17, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155,
	29, 3, 2, 2, 2, 156, 157, 7, 29, 2, 2, 157, 31, 3, 2, 2, 2, 158, 159, 5,
	34, 18, 2, 159, 33, 3, 2, 2, 2, 160, 161, 9, 3, 2, 2, 161, 35, 3, 2, 2,
	2, 162, 163, 7, 29, 2, 2, 163, 167, 7, 3, 2, 2, 164, 166, 5, 38, 20, 2,
	165, 164, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167,
	168, 3, 2, 2, 2, 168, 170, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 171,
	7, 4, 2, 2, 171, 37, 3, 2, 2, 2, 172, 175, 5, 40, 21, 2, 173, 175, 7, 29,
	2, 2, 174, 172, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 39, 3, 2, 2, 2,
	176, 177, 9, 4, 2, 2, 177, 41, 3, 2, 2, 2, 178, 179, 7, 26, 2, 2, 179,
	180, 7, 29, 2, 2, 180, 43, 3, 2, 2, 2, 181, 182, 7, 27, 2, 2, 182, 183,
	7, 29, 2, 2, 183, 45, 3, 2, 2, 2, 17, 47, 52, 58, 73, 79, 90, 103, 113,
	118, 125, 136, 148, 154, 167, 174,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'<'", "'>'", "'='", "'+='", "'-='", "'*='",
	"'/='", "'&='", "'|='", "'^='", "'>>='", "'>>>='", "'<<='", "'%='", "','",
	"'for'", "'var'", "'package'", "'import'", "'struct'", "'member'", "'func'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "FOR", "VAR", "PACKAGE", "IMPORT", "DATA_STRUCT", "MEMBER", "FUNCTION",
	"IDENTIFIER", "WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL", "DECIMAL_LITERAL",
}

var ruleNames = []string{
	"compilationUnit", "packageDeclaration", "importDeclaration", "typeDeclaration",
	"functionDeclaration", "functionBody", "primary", "expressDeclaration",
	"block", "blockStatement", "forControl", "localVariableDeclaration", "variableDeclarators",
	"variableDeclarator", "variableDeclaratorId", "variableInitializer", "expression",
	"methodCallDeclaration", "parameter", "literal", "dataStructDeclaration",
	"memberDeclaration",
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
	CodeParserT__4            = 5
	CodeParserT__5            = 6
	CodeParserT__6            = 7
	CodeParserT__7            = 8
	CodeParserT__8            = 9
	CodeParserT__9            = 10
	CodeParserT__10           = 11
	CodeParserT__11           = 12
	CodeParserT__12           = 13
	CodeParserT__13           = 14
	CodeParserT__14           = 15
	CodeParserT__15           = 16
	CodeParserT__16           = 17
	CodeParserT__17           = 18
	CodeParserT__18           = 19
	CodeParserFOR             = 20
	CodeParserVAR             = 21
	CodeParserPACKAGE         = 22
	CodeParserIMPORT          = 23
	CodeParserDATA_STRUCT     = 24
	CodeParserMEMBER          = 25
	CodeParserFUNCTION        = 26
	CodeParserIDENTIFIER      = 27
	CodeParserWS              = 28
	CodeParserCOMMENT         = 29
	CodeParserLINE_COMMENT    = 30
	CodeParserSTRING_LITERAL  = 31
	CodeParserDECIMAL_LITERAL = 32
)

// CodeParser rules.
const (
	CodeParserRULE_compilationUnit          = 0
	CodeParserRULE_packageDeclaration       = 1
	CodeParserRULE_importDeclaration        = 2
	CodeParserRULE_typeDeclaration          = 3
	CodeParserRULE_functionDeclaration      = 4
	CodeParserRULE_functionBody             = 5
	CodeParserRULE_primary                  = 6
	CodeParserRULE_expressDeclaration       = 7
	CodeParserRULE_block                    = 8
	CodeParserRULE_blockStatement           = 9
	CodeParserRULE_forControl               = 10
	CodeParserRULE_localVariableDeclaration = 11
	CodeParserRULE_variableDeclarators      = 12
	CodeParserRULE_variableDeclarator       = 13
	CodeParserRULE_variableDeclaratorId     = 14
	CodeParserRULE_variableInitializer      = 15
	CodeParserRULE_expression               = 16
	CodeParserRULE_methodCallDeclaration    = 17
	CodeParserRULE_parameter                = 18
	CodeParserRULE_literal                  = 19
	CodeParserRULE_dataStructDeclaration    = 20
	CodeParserRULE_memberDeclaration        = 21
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CodeParserPACKAGE {
		{
			p.SetState(44)
			p.PackageDeclaration()
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIMPORT {
		{
			p.SetState(47)
			p.ImportDeclaration()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-3)&-(0x1f+1)) == 0 && ((1<<uint((_la-3)))&((1<<(CodeParserT__2-3))|(1<<(CodeParserFOR-3))|(1<<(CodeParserVAR-3))|(1<<(CodeParserDATA_STRUCT-3))|(1<<(CodeParserMEMBER-3))|(1<<(CodeParserFUNCTION-3))|(1<<(CodeParserIDENTIFIER-3))|(1<<(CodeParserSTRING_LITERAL-3))|(1<<(CodeParserDECIMAL_LITERAL-3)))) != 0 {
		{
			p.SetState(53)
			p.TypeDeclaration()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
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
		p.SetState(61)
		p.Match(CodeParserPACKAGE)
	}
	{
		p.SetState(62)
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
		p.SetState(64)
		p.Match(CodeParserIMPORT)
	}
	{
		p.SetState(65)
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

func (s *TypeDeclarationContext) ExpressDeclaration() IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CodeParserDATA_STRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.DataStructDeclaration()
		}

	case CodeParserMEMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.MemberDeclaration()
		}

	case CodeParserFUNCTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.FunctionDeclaration()
		}

	case CodeParserT__2, CodeParserFOR, CodeParserVAR, CodeParserIDENTIFIER, CodeParserSTRING_LITERAL, CodeParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.expressDeclaration(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) Parameter() IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
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
	p.EnterRule(localctx, 8, CodeParserRULE_functionDeclaration)
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
		p.SetState(73)
		p.Match(CodeParserFUNCTION)
	}
	{
		p.SetState(74)
		p.Match(CodeParserIDENTIFIER)
	}
	{
		p.SetState(75)
		p.Match(CodeParserT__0)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(CodeParserIDENTIFIER-27))|(1<<(CodeParserSTRING_LITERAL-27))|(1<<(CodeParserDECIMAL_LITERAL-27)))) != 0 {
		{
			p.SetState(76)
			p.Parameter()
		}

	}
	{
		p.SetState(79)
		p.Match(CodeParserT__1)
	}
	{
		p.SetState(80)
		p.Match(CodeParserT__2)
	}
	{
		p.SetState(81)
		p.FunctionBody()
	}
	{
		p.SetState(82)
		p.Match(CodeParserT__3)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) AllExpressDeclaration() []IExpressDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem())
	var tst = make([]IExpressDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressDeclarationContext)
		}
	}

	return tst
}

func (s *FunctionBodyContext) ExpressDeclaration(i int) IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CodeParserRULE_functionBody)
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
		p.expressDeclaration(0)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-3)&-(0x1f+1)) == 0 && ((1<<uint((_la-3)))&((1<<(CodeParserT__2-3))|(1<<(CodeParserFOR-3))|(1<<(CodeParserVAR-3))|(1<<(CodeParserIDENTIFIER-3))|(1<<(CodeParserSTRING_LITERAL-3))|(1<<(CodeParserDECIMAL_LITERAL-3)))) != 0 {
		{
			p.SetState(85)
			p.expressDeclaration(0)
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CodeParserRULE_primary)

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
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IExpressDeclarationContext is an interface to support dynamic dispatch.
type IExpressDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// IsExpressDeclarationContext differentiates from other interfaces.
	IsExpressDeclarationContext()
}

type ExpressDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	bop    antlr.Token
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

func (s *ExpressDeclarationContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressDeclarationContext) SetBop(v antlr.Token) { s.bop = v }

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

func (s *ExpressDeclarationContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressDeclarationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressDeclarationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressDeclarationContext) AllExpressDeclaration() []IExpressDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem())
	var tst = make([]IExpressDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressDeclarationContext)
		}
	}

	return tst
}

func (s *ExpressDeclarationContext) ExpressDeclaration(i int) IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
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
	return p.expressDeclaration(0)
}

func (p *CodeParser) expressDeclaration(_p int) (localctx IExpressDeclarationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressDeclarationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressDeclarationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, CodeParserRULE_expressDeclaration, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(94)
			p.MethodCallDeclaration()
		}

	case 2:
		{
			p.SetState(95)
			p.BlockStatement()
		}

	case 3:
		{
			p.SetState(96)
			p.Primary()
		}

	case 4:
		{
			p.SetState(97)
			p.Expression()
		}
		{
			p.SetState(98)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressDeclarationContext).bop = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserT__6)|(1<<CodeParserT__7)|(1<<CodeParserT__8)|(1<<CodeParserT__9)|(1<<CodeParserT__10)|(1<<CodeParserT__11)|(1<<CodeParserT__12)|(1<<CodeParserT__13)|(1<<CodeParserT__14)|(1<<CodeParserT__15)|(1<<CodeParserT__16)|(1<<CodeParserT__17))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressDeclarationContext).bop = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(99)
			p.Expression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressDeclarationContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CodeParserRULE_expressDeclaration)
			p.SetState(103)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(104)
					p.Match(CodeParserT__4)
				}
				{
					p.SetState(105)
					p.Match(CodeParserT__4)
				}

			case 2:
				{
					p.SetState(106)
					p.Match(CodeParserT__5)
				}
				{
					p.SetState(107)
					p.Match(CodeParserT__5)
				}
				{
					p.SetState(108)
					p.Match(CodeParserT__5)
				}

			case 3:
				{
					p.SetState(109)
					p.Match(CodeParserT__5)
				}
				{
					p.SetState(110)
					p.Match(CodeParserT__5)
				}

			}
			{
				p.SetState(113)
				p.expressDeclaration(3)
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CodeParserRULE_block)
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
		p.SetState(119)
		p.Match(CodeParserT__2)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserT__2)|(1<<CodeParserFOR)|(1<<CodeParserVAR))) != 0 {
		{
			p.SetState(120)
			p.BlockStatement()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
		p.Match(CodeParserT__3)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBlockLabel returns the blockLabel rule contexts.
	GetBlockLabel() IBlockContext

	// SetBlockLabel sets the blockLabel rule contexts.
	SetBlockLabel(IBlockContext)

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	blockLabel IBlockContext
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

func (s *BlockStatementContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *BlockStatementContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *BlockStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStatementContext) LocalVariableDeclaration() ILocalVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalVariableDeclarationContext)
}

func (s *BlockStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(CodeParserFOR, 0)
}

func (s *BlockStatementContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *BlockStatementContext) BlockStatement() IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
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
	p.EnterRule(localctx, 18, CodeParserRULE_blockStatement)

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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CodeParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)

			var _x = p.Block()

			localctx.(*BlockStatementContext).blockLabel = _x
		}

	case CodeParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.LocalVariableDeclaration()
		}

	case CodeParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(CodeParserFOR)
		}
		{
			p.SetState(131)
			p.ForControl()
		}
		{
			p.SetState(132)
			p.BlockStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForExpress returns the forExpress rule contexts.
	GetForExpress() IExpressDeclarationContext

	// SetForExpress sets the forExpress rule contexts.
	SetForExpress(IExpressDeclarationContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	forExpress IExpressDeclarationContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForExpress() IExpressDeclarationContext { return s.forExpress }

func (s *ForControlContext) SetForExpress(v IExpressDeclarationContext) { s.forExpress = v }

func (s *ForControlContext) ExpressDeclaration() IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CodeParserRULE_forControl)

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
		p.SetState(136)

		var _x = p.expressDeclaration(0)

		localctx.(*ForControlContext).forExpress = _x
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
	p.EnterRule(localctx, 22, CodeParserRULE_localVariableDeclaration)

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
		p.SetState(138)
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

func (s *VariableDeclaratorsContext) VAR() antlr.TerminalNode {
	return s.GetToken(CodeParserVAR, 0)
}

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
	p.EnterRule(localctx, 24, CodeParserRULE_variableDeclarators)

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
		p.SetState(140)
		p.Match(CodeParserVAR)
	}
	{
		p.SetState(141)
		p.VariableDeclarator()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(142)
				p.Match(CodeParserT__18)
			}
			{
				p.SetState(143)
				p.VariableDeclarator()
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, CodeParserRULE_variableDeclarator)

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
		p.SetState(149)
		p.VariableDeclaratorId()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(150)
			p.Match(CodeParserT__6)
		}
		{
			p.SetState(151)
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
	p.EnterRule(localctx, 28, CodeParserRULE_variableDeclaratorId)

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
		p.SetState(154)
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
	p.EnterRule(localctx, 30, CodeParserRULE_variableInitializer)

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
		p.SetState(156)
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

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserSTRING_LITERAL, 0)
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
	p.EnterRule(localctx, 32, CodeParserRULE_expression)
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
		p.SetState(158)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(CodeParserIDENTIFIER-27))|(1<<(CodeParserSTRING_LITERAL-27))|(1<<(CodeParserDECIMAL_LITERAL-27)))) != 0) {
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
	p.EnterRule(localctx, 34, CodeParserRULE_methodCallDeclaration)
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
		p.SetState(160)
		p.Match(CodeParserIDENTIFIER)
	}
	{
		p.SetState(161)
		p.Match(CodeParserT__0)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(CodeParserIDENTIFIER-27))|(1<<(CodeParserSTRING_LITERAL-27))|(1<<(CodeParserDECIMAL_LITERAL-27)))) != 0 {
		{
			p.SetState(162)
			p.Parameter()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(CodeParserT__1)
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

func (s *ParameterContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 36, CodeParserRULE_parameter)

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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CodeParserSTRING_LITERAL, CodeParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Literal()
		}

	case CodeParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(CodeParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserDECIMAL_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CodeParserRULE_literal)
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
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CodeParserSTRING_LITERAL || _la == CodeParserDECIMAL_LITERAL) {
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
	p.EnterRule(localctx, 40, CodeParserRULE_dataStructDeclaration)

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
		p.SetState(176)
		p.Match(CodeParserDATA_STRUCT)
	}
	{
		p.SetState(177)
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
	p.EnterRule(localctx, 42, CodeParserRULE_memberDeclaration)

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
		p.SetState(179)
		p.Match(CodeParserMEMBER)
	}
	{
		p.SetState(180)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

func (p *CodeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpressDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*ExpressDeclarationContext)
		}
		return p.ExpressDeclaration_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CodeParser) ExpressDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
