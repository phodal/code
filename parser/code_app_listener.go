package parser

import (
	. "../languages/code"
	"fmt"
)

func NewCodeAppListener() *CodeAppListener {
	return &CodeAppListener{}
}


type CodeAppListener struct {
	BaseCodeListener
}

func (s *CodeAppListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {
	fmt.Println(ctx.IDENTIFIER().GetText())
	fmt.Println(ctx.ParameterList().GetText())
}