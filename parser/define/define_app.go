package code

import (
	. "../../languages/define"
	. "../../model"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var startTemplateSymbol string
var endTemplateSymbol string

func NewDefineApp(startSymbol string, endSymbol string) *DefineApp {
	startTemplateSymbol = startSymbol
	endTemplateSymbol = endSymbol
	return &DefineApp{}
}

type DefineApp struct {
}

func (j *DefineApp) Start(path string) DefineInformation {
	context := (*DefineApp)(nil).ProcessFile(path).CompilationUnit()
	listener := NewDefineAppListener(startTemplateSymbol, endTemplateSymbol)

	antlr.NewParseTreeWalker().Walk(listener, context)

	information := listener.getDefineInformation()
	return information
}

func (j *DefineApp) ProcessFile(path string) *DefineParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewDefineLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewDefineParser(stream)
	return parser
}
