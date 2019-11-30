package transform

import (
	. "../model"
	"fmt"
	"strings"
)

type Transform struct {
}

func (transform Transform) ToCode(call CodeFunctionCall, modules []DefineModule) string {
	var parameters []string
	for _, parameter := range call.Parameters {
		parameters = append(parameters, parameter.Value.Value)
	}

	callName :=  call.MemberId
	for _, module := range modules {
		for _, function := range module.ModuleFunctions {
			if function.FunctionName == call.MemberId {
				callName = function.EqualName[1:len(function.EqualName)-1]
			}
		}
	}

	paramList := strings.Join(parameters, ",")

	return callName + "(" + paramList + ")"
}

func (transform Transform) BuildImport(call CodeFunctionCall, modules []DefineModule) []string {
	var imports []string

	for _, module := range modules {
		for _, function := range module.ModuleFunctions {
			if function.FunctionName == call.MemberId {
				imports = append(imports, "import " + function.ImportName + "\n")
			}
		}
	}
	return imports
}

func (transform Transform) BuildPackage(s string) string {
	return "package " + s + "\n"
}

func (transform Transform) BuildFunction(function CodeFunction, information DefineInformation) string {
	fmt.Println(information)
	return ""
}
