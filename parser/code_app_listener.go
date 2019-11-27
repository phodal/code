package parser

import (
	. "../languages/code"
	. "../model"
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
	allParameters := ctx.AllParameter()

	var parameters []CodeParameter
	for _, parameter := range allParameters {
		stringCodeType := &CodeType{
			Type: "string",
		}
		var paramValue = &CodeParameterValue{Value: parameter.GetText()}
		parameter := &CodeParameter{*stringCodeType, *paramValue}
		parameters = append(parameters, *parameter)
	}

	functionCall := CreateFunctionCall(ctx.IDENTIFIER().GetText(), parameters)
	fmt.Println(functionCall)
}
