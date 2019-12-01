package main

import (
	. "./model"
	. "./parser/code"
	. "./parser/define"
	. "./transform"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
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

	code = code + transform.TransformMainCode(codeModel, info, startTemplateSymbol, endTemplateSymbol)
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

