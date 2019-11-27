package main

import (
	. "./parser/code"
	. "./parser/define"
)

func main()  {
	app := NewCodeApp()
	app.Start("examples/helloworld.code")

	defineApp := NewDefineApp()
	defineApp.Start("examples/mu.define")
}
