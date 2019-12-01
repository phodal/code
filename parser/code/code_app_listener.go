package code

import (
	. "../../languages/code"
	. "../../model"
	"reflect"
)

var currentCodeModel CodeModel
var currentFunction = CreateFunction("")

func NewCodeAppListener() *CodeAppListener {
	currentCodeModel = *&CodeModel{nil, nil, nil}
	return &CodeAppListener{}
}

type CodeAppListener struct {
	BaseCodeListener
}

func (s *CodeAppListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {
	allParameters := ctx.AllParameter()
	functionName := ctx.IDENTIFIER().GetText()

	parentParentType := reflect.TypeOf(ctx.GetParent().GetParent()).String()
	switch parentParentType {
	case "*parser.TypeDeclarationContext":
		functionCall := BuildFunctionCall(allParameters, functionName)
		currentCodeModel.FunctionCalls = append(currentCodeModel.FunctionCalls, functionCall)
	case "*parser.FunctionBodyContext":
		functionCall := BuildFunctionCall(allParameters, functionName)
		currentFunction.CodeFunctionCalls = append(currentFunction.CodeFunctionCalls, functionCall)
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
	currentFunction = function

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
}

var varMaps = make(map[string]string)

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	for _, varDeclarator := range ctx.AllVariableDeclarator() {
		varCtx := varDeclarator.(*VariableDeclaratorContext)
		ident := varCtx.VariableDeclaratorId().GetText()
		value := ""

		if varCtx.VariableInitializer() != nil {
			value = varCtx.VariableInitializer().GetText()
		}
		varMaps[ident] = value
	}
}

func (s *CodeAppListener) getCode() CodeModel {
	currentCodeModel.Variables = varMaps
	return currentCodeModel
}
