package code

import (
	. "../../languages/code"
	. "../../model"
	"reflect"
)

var currentCodeModel CodeModel
var currentFunction = CreateFunction("")
var varMaps = make(map[string]string)

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
		childType := reflect.TypeOf(parameter.GetChild(0)).String()
		paraCodeType := &CodeType{
			Type: "string",
		}

		switch childType {
		case "*antlr.TerminalNodeImpl":
			paraCodeType.Type = "type"
		default:
		}

		var paramValue = &CodeParameterValue{Value: parameter.GetText()}
		parameter := &CodeParameter{*paraCodeType, *paramValue}
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

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	parentType := reflect.TypeOf(ctx.GetParent()).String()
	for _, varDeclarator := range ctx.AllVariableDeclarator() {
		varCtx := varDeclarator.(*VariableDeclaratorContext)
		ident := varCtx.VariableDeclaratorId().GetText()
		value := ""

		if varCtx.VariableInitializer() != nil {
			value = varCtx.VariableInitializer().GetText()
		}

		if parentType == "*parser.LocalVariableDeclarationContext" {
			currentFunction.Variables[ident] = value
		}

		varMaps[ident] = value
	}
}

func (s *CodeAppListener) getCode() CodeModel {
	currentCodeModel.Variables = varMaps
	return currentCodeModel
}
