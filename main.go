package main

import (
	. "./model"
	. "./parser/code"
	. "./parser/define"
	codetemplate "./parser/template"
	. "./transform"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	app := NewCodeApp()
	codeModel := app.Start("examples/helloworld.code")

	startTemplateSymbol := "{{"
	endTemplateSymbol := "}}"
	defineApp := NewDefineApp(startTemplateSymbol, endTemplateSymbol)
	info := defineApp.Start("examples/mu.define")

	transformCode(codeModel, info, startTemplateSymbol, endTemplateSymbol)
}

func transformCode(codeModel CodeModel, info DefineInformation, startTemplateSymbol string, endTemplateSymbol string) {
	transform := &Transform{}
	var packageInfo string
	var imports []string
	var code = ""
	var result = ""
	for _, call := range codeModel.FunctionCalls {
		code = code + "\n" + transform.ToCode(call, info.ModuleDeclarations)
		imports = transform.BuildImport(call, info.ModuleDeclarations)
	}
	templates := info.DefineTemplates
	template := codetemplate.New(templates["code"], startTemplateSymbol, endTemplateSymbol)
	result = template.ExecuteString(map[string]interface{}{
		"code": code,
	})
	packageInfo = transform.BuildPackage("main")
	codeWithImport := packageInfo + strings.Join(imports, "") + result
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
