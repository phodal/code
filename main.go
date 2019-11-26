package main

import (
	. "./parser"
)

func main()  {
	app := NewCodeApp()
	app.Start("examples/helloworld.code")
}
