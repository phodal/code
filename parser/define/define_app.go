package code

import (
	. "../../languages/define"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewDefineApp() *DefineApp {
	return &DefineApp{}
}

type DefineApp struct {
}

func (j *DefineApp) Start(path string) {
	context := (*DefineApp)(nil).ProcessFile(path).CompilationUnit()
	listener := NewDefineAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func (j *DefineApp) ProcessFile(path string) *DefineParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewDefineLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewDefineParser(stream)
	return parser
}
