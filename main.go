package main

import (
	. "./parser/code"
	. "./parser/define"
	"fmt"
)

func main()  {
	app := NewCodeApp()
	app.Start("examples/helloworld.code")

	startTemplateSymbol := "{{"
	endTemplateSymbol := "}}"
	defineApp := NewDefineApp(startTemplateSymbol, endTemplateSymbol)
	info := defineApp.Start("examples/mu.define")

	fmt.Println(info)
	//template.New(info, startTemplateSymbol, endTemplateSymbol)
}
