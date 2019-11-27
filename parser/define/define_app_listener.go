package code

import (
	. "../../languages/define"
)

func NewDefineAppListener() *DefineAppListener {
	return &DefineAppListener{}
}

type DefineAppListener struct {
	BaseDefineListener
}
