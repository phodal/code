package main

import (
	. "./model"
	. "./parser/code"
	. "./parser/define"
	codetemplate "./parser/template"
	. "./transform"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

var transform Transform

func main() {
	transform = *&Transform{}

	app := NewCodeApp()
	codeModel := app.Start("examples/helloworld.code")

	bytes, _ := json.MarshalIndent(codeModel, "", "\t")
	_ = ioutil.WriteFile("code.json", bytes, 0644)

	startTemplateSymbol := "{{"
	endTemplateSymbol := "}}"
	defineApp := NewDefineApp(startTemplateSymbol, endTemplateSymbol)
	info := defineApp.Start("examples/mu.define")

	code := ""

	code = code + transformMainCode(codeModel, info, startTemplateSymbol, endTemplateSymbol)
	code = code + transformNormalCode(codeModel, info)

	runCode(code)
}

func transformNormalCode(model CodeModel, information DefineInformation) string {
	funcStr := "\n"
	for _, function := range model.Functions {
		funcStr = funcStr + transform.BuildFunction(function, information)
	}

	return funcStr
}

func runCode(codeWithImport string) {
	_ = ioutil.WriteFile("test/main/main.go", []byte(codeWithImport), 0644)
	cmd := exec.Command("go", "run", "test/main/main.go")
	stdout, err := cmd.StdoutPipe()
	_ = cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}

func transformMainCode(codeModel CodeModel, info DefineInformation, startTemplateSymbol string, endTemplateSymbol string) string {
	var packageInfo string
	var imports []string
	var code = ""
	var result = ""
	for _, call := range codeModel.FunctionCalls {
		code = code + "\n" + transform.BuildFunctionCall(call, info.DefineModules)
		imports = transform.BuildImport(call, info.DefineModules)
	}
	templates := info.DefineTemplates
	template := codetemplate.New(templates["code"], startTemplateSymbol, endTemplateSymbol)
	result = template.ExecuteString(map[string]interface{}{
		"code": code,
	})
	packageInfo = transform.BuildPackage("main")
	codeWithImport := packageInfo + strings.Join(imports, "") + result

	return codeWithImport
}
