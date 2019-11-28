package main

import (
	. "./parser/code"
	. "./parser/define"
	codetemplate "./parser/template"
	"fmt"
)

func main() {
	app := NewCodeApp()
	app.Start("examples/helloworld.code")

	startTemplateSymbol := "{{"
	endTemplateSymbol := "}}"
	defineApp := NewDefineApp(startTemplateSymbol, endTemplateSymbol)
	info := defineApp.Start("examples/mu.define")

	templates := info.DefineTemplates
	fmt.Println(info)
	template := codetemplate.New(templates["code"], startTemplateSymbol, endTemplateSymbol)
	result := template.ExecuteString(map[string]interface{}{
		"code": "aaaa",
	})
	fmt.Printf("%s", result)
}
