package model


type DefineInformation struct {
	DefineTemplates    map[string]string
	ModuleDeclarations []DefineModule
}

type DefineModule struct {
	ModuleName      string
	ModuleFunctions []ModuleFunction
}

type ModuleFunction struct {
	FunctionName string
	EqualName    string
	ImportName   string
}
