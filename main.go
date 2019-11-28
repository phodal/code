package main

import (
	. "./parser/code"
	. "./parser/define"
	codetemplate "./parser/template"
	. "./transform"
	"io/ioutil"
	"strings"
)

func main() {
	app := NewCodeApp()
	codeModel := app.Start("examples/helloworld.code")

	startTemplateSymbol := "{{"
	endTemplateSymbol := "}}"
	defineApp := NewDefineApp(startTemplateSymbol, endTemplateSymbol)
	info := defineApp.Start("examples/mu.define")

	transform := &Transform{}
	code := transform.ToCode(codeModel, info.ModuleDeclarations)

	templates := info.DefineTemplates
	template := codetemplate.New(templates["code"], startTemplateSymbol, endTemplateSymbol)
	result := template.ExecuteString(map[string]interface{}{
		"code": code,
	})


	packageInfo := transform.BuildPackage("test")
	imports := transform.BuildImport(codeModel, info.ModuleDeclarations)

	codeWithImport := packageInfo + strings.Join(imports, "") + result

	_ = ioutil.WriteFile("test/test.go", []byte(codeWithImport), 0644)
}
