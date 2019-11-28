package code

import (
	. "../../languages/code"
	. "../../model"
	"fmt"
)

var functionCall CodeFunctionCall

func NewCodeAppListener() *CodeAppListener {
	functionCall = *&CodeFunctionCall{"", nil, nil}
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

	functionCall = CreateFunctionCall(ctx.IDENTIFIER().GetText(), parameters)
}

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	fmt.Println(ctx.GetText())
}

func (s *CodeAppListener) getCode() CodeFunctionCall {
	return functionCall
}

