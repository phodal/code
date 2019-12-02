package transform

import (
	. "../model"
	codetemplate "../parser/template"
	"fmt"
	"strings"
)

var imports = make(map[string]string)

type Transform struct {
}

func (transform Transform) BuildFunctionCall(call CodeFunctionCall, info DefineInformation) string {
	var parameters []string
	for _, parameter := range call.Parameters {
		switch parameter.Key.Type {
		case "string":
			parameters = append(parameters, parameter.Value.Value)
		case "integer":
			parameters = append(parameters, parameter.Value.Value)
		}
	}

	modules := info.DefineModules

	callName := call.MemberId
	for _, module := range modules {
		for _, function := range module.ModuleFunctions {
			if function.FunctionName == call.MemberId {
				callName = function.EqualName[1 : len(function.EqualName)-1]
			}
		}
	}

	paramList := strings.Join(parameters, ",")

	return callName + "(" + paramList + ")"
}

func (transform Transform) BuildImport(call CodeFunctionCall, modules []DefineModule) {
	for _, module := range modules {
		for _, function := range module.ModuleFunctions {
			if function.FunctionName == call.MemberId {
				imports[function.ImportName] = function.ImportName
			}
		}
	}
}

func (transform Transform) GetImports() string {
	var str = ""
	for _, imp := range imports {
		str += "import " + imp + "\n"
	}

	return str + "\n"
}

func (transform Transform) BuildPackage(s string) string {
	return "package " + s + "\n"
}

func (transform Transform) BuildFunction(function CodeFunction, information DefineInformation) string {
	symbolMap := information.SymbolsMap
	funcBody := ""
	funcName := function.MemberId
	params := ""
	callCode := ""

	for _, param := range function.Parameters {
		fmt.Println(param.Type)
	}

	for _, call := range function.CodeFunctionCalls {
		callCode = transform.BuildFunctionCall(call, information)
	}

	funcBody = "\n" + callCode

	return symbolMap["FUNCTION"] + " " + funcName + symbolMap["PARAMETER_START"] + params + symbolMap["PARAMETER_END"] + symbolMap["METHOD_START"] + funcBody + symbolMap["METHOD_END"]
}

func  (transform Transform) TransformMainCode(codeModel CodeModel, info DefineInformation, startTemplateSymbol string, endTemplateSymbol string) string {
	var packageInfo string
	var importStr string
	var code = ""
	var result = ""
	for _, call := range codeModel.FunctionCalls {
		code = code + "\n" + transform.BuildFunctionCall(call, info)
		transform.BuildImport(call, info.DefineModules)
	}

	importStr = transform.GetImports()

	templates := info.DefineTemplates
	template := codetemplate.New(templates["code"], startTemplateSymbol, endTemplateSymbol)
	result = template.ExecuteString(map[string]interface{}{
		"code": code,
	})
	packageInfo = transform.BuildPackage("main")
	codeWithImport := packageInfo + importStr + result

	return codeWithImport
}

func (transform Transform) TransformNormalCode(model CodeModel, information DefineInformation) string {
	funcStr := "\n"
	for _, function := range model.Functions {
		funcStr = funcStr + transform.BuildFunction(function, information)
	}

	return funcStr
}

