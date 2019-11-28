package code

import (
	. "../../languages/code"
	. "../../model"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewCodeApp() *CodeApp {
	return &CodeApp{}
}

type CodeApp struct {
}

func (j *CodeApp) Start(path string) CodeFunctionCall {
	context := (*CodeApp)(nil).ProcessFile(path).CompilationUnit()
	listener := NewCodeAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	code := listener.getCode()
	return code
}

func (j *CodeApp) ProcessFile(path string) *CodeParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewCodeLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewCodeParser(stream)
	return parser
}
