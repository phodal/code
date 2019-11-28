package transform

import (
	. "../model"
	"strings"
)

type Transform struct {
}

func (transform Transform) ToCode(call CodeFunctionCall, modules []DefineModule) string {
	var parameters []string
	for _, parameter := range call.Parameters {
		parameters = append(parameters, parameter.Value.Value)
	}

	paramList := strings.Join(parameters, ",")

	return call.MemberId + "(" + paramList + ")"
}

func (transform Transform) BuildImport(call CodeFunctionCall, modules []DefineModule) []string {
	var imports []string

	for _, module := range modules {
		for _, function := range module.ModuleFunctions {
			if function.FunctionName == call.MemberId {
				imports = append(imports, "import " + function.EqualName + "\n")
			}
		}
	}
	return imports
}

func (transform Transform) BuildPackage(s string) string {
	return "package " + s + "\n"
}
