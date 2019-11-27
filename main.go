package main

import (
	. "./parser/code"
)

func main()  {
	app := NewCodeApp()
	app.Start("examples/helloworld.code")
}
