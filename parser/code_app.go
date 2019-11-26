package parser

import (
	. "../languages/code"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewCodeApp() *CodeApp {
	return &CodeApp{}
}

type CodeApp struct {
}

func (j *CodeApp) Start(path string) {
	context := (*CodeApp)(nil).ProcessFile(path).CompilationUnit()
	listener := NewCodeAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func (j *CodeApp) ProcessFile(path string) *CodeParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewCodeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewCodeParser(stream)
	return parser
}
