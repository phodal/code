package code

import (
	. "../../languages/code"
	. "../../model"
	"fmt"
)

type CodeModel struct {
	FunctionCalls []CodeFunctionCall
}

var currentCodeModel CodeModel

func NewCodeAppListener() *CodeAppListener {
	//functionCall = *&CodeFunctionCall{"", nil, nil}
	currentCodeModel = *&CodeModel{nil}
	return &CodeAppListener{}
}

type CodeAppListener struct {
	BaseCodeListener
}

func (s *CodeAppListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {
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
	currentCodeModel.FunctionCalls = append(currentCodeModel.FunctionCalls, functionCall)
}

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	fmt.Println(ctx.GetText())
}

func (s *CodeAppListener) getCode() CodeModel {
	return currentCodeModel
}

