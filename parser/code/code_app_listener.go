package code

import (
	. "../../languages/code"
	. "../../model"
	"fmt"
	"reflect"
)

type CodeModel struct {
	FunctionCalls []CodeFunctionCall
	Functions     []CodeFunction
}

var currentCodeModel CodeModel
var currentFunction = CreateFunction("")

func NewCodeAppListener() *CodeAppListener {
	currentCodeModel = *&CodeModel{nil, nil}
	return &CodeAppListener{}
}

type CodeAppListener struct {
	BaseCodeListener
}

func (s *CodeAppListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {
	if currentFunction.MemberId == "" {
		allParameters := ctx.AllParameter()
		functionName := ctx.IDENTIFIER().GetText()

		functionCall := BuildFunctionCall(allParameters, functionName)
		currentCodeModel.FunctionCalls = append(currentCodeModel.FunctionCalls, functionCall)
	}
}

func BuildFunctionCall(allParameters []IParameterContext, functionName string) CodeFunctionCall {
	var parameters []CodeParameter
	for _, parameter := range allParameters {
		stringCodeType := &CodeType{
			Type: "string",
		}
		var paramValue = &CodeParameterValue{Value: parameter.GetText()}
		parameter := &CodeParameter{*stringCodeType, *paramValue}
		parameters = append(parameters, *parameter)
	}
	functionCall := CreateFunctionCall(functionName, parameters)
	return functionCall
}

func (s *CodeAppListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {
	function := CreateFunction(ctx.IDENTIFIER().GetText())
	child := ctx.FunctionBody().GetChild(0)
	if child != nil {
		expressCtx := child.(*ExpressDeclarationContext)
		firstChildCtx := expressCtx.GetChild(0)
		switch reflect.TypeOf(firstChildCtx).String() {
		case "*parser.MethodCallDeclarationContext":
			context := firstChildCtx.(*MethodCallDeclarationContext)
			functionCall := BuildFunctionCall(context.AllParameter(), context.IDENTIFIER().GetText())
			function.CodeFunctionCalls = append(function.CodeFunctionCalls, functionCall)
		}
	}

	currentCodeModel.Functions = append(currentCodeModel.Functions, function)
	currentFunction = CreateFunction("")
}

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	fmt.Println(ctx.GetText())
}

func (s *CodeAppListener) getCode() CodeModel {
	return currentCodeModel
}
