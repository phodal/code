package model

type DefineInformation struct {
	DefineTemplates map[string]string
	SymbolsMap      map[string]string
	DefineModules   []DefineModule
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

type CodeModel struct {
	FunctionCalls []CodeFunctionCall
	Functions     []CodeFunction
	Variables     map[string]string
}
