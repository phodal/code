package code

import (
	. "../../languages/define"
	. "../../model"
	"reflect"
	"strings"
)

type Symbol struct {
	Key   string
	Value string
}

var symbols []Symbol

var startSymbol string
var endSymbol string

var defineInformation DefineInformation

func NewDefineAppListener(start string, end string) *DefineAppListener {
	startSymbol = start
	endSymbol = end
	defineInformation = *&DefineInformation{
		DefineTemplates: make(map[string]string),
	}
	defineInformation.SymbolsMap = make(map[string]string)
	return &DefineAppListener{}
}

type DefineAppListener struct {
	BaseDefineListener
}

func (s *DefineAppListener) EnterSymbolDeclaration(ctx *SymbolDeclarationContext) {
	key := ctx.IDENTIFIER().GetText()
	value := ctx.STRING_LITERAL().GetText()
	symbol := &Symbol{key, value}
	symbols = append(symbols, *symbol)

	defineInformation.SymbolsMap[key] = value[1: len(value)-1]

}

func (s *DefineAppListener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {
	defineModule := &DefineModule{
		ModuleName:      ctx.IDENTIFIER().GetText(),
		ModuleFunctions: nil,
	}

	for _, moduleDefine := range ctx.AllModuleDefines() {
		moduleFunction := &ModuleFunction{}
		moduleDefineCtx := moduleDefine.(*ModuleDefinesContext)
		for _, attribute := range moduleDefineCtx.AllModuleAttribute() {
			moduleFunction.FunctionName = moduleDefineCtx.IDENTIFIER().GetText()
			attr := attribute.(*ModuleAttributeContext)
			if attr.IMPORT() != nil {
				moduleFunction.ImportName = attr.STRING_LITERAL().GetText()
			}
			if attr.EQUAL() != nil {
				moduleFunction.EqualName = attr.STRING_LITERAL().GetText()
			}
		}

		defineModule.ModuleFunctions = append(defineModule.ModuleFunctions, *moduleFunction)

	}

	defineInformation.ModuleDeclarations = append(defineInformation.ModuleDeclarations, *defineModule)
}

func (s *DefineAppListener) EnterDefineDeclaration(ctx *DefineDeclarationContext) {
	switch ctx.DefineKey().GetText() {
	case "entry_point":
		buildEntryPoint(ctx)
	}
}

func buildEntryPoint(ctx *DefineDeclarationContext) {
	for _, express := range ctx.AllDefineExpress() {
		firstChild := express.(*DefineExpressContext).GetChild(0)
		typeOfExpress := reflect.TypeOf(firstChild).String()
		switch typeOfExpress {
		case "*parser.DefineAttributeContext":
			attributeCtx := firstChild.(*DefineAttributeContext)
			attributeCtx.GetText()
		case "*parser.DefineTemplateContext":
			defineCtx := firstChild.(*DefineTemplateContext)
			templateKey := defineCtx.IDENTIFIER().GetText()
			defineBody := defineCtx.DefineBody().(*DefineBodyContext)
			funcName := defineBody.STRING_LITERAL().GetText()
			templateData := defineBody.TemplateData().GetText()

			allSymbol := defineBody.AllSymbolKey()

			var codeText []string
			for index, symbolKey := range allSymbol {
				symbolCtx := symbolKey.(*SymbolKeyContext)
				symbolText := symbolCtx.GetText()

				codeText = append(codeText, defineInformation.SymbolsMap[symbolText])
				if index == 0 {
					codeText = append(codeText, funcName[1:len(funcName)-1])
				}
				if index == len(allSymbol)-2 {
					codeText = append(codeText, buildTemplate(templateData))
				}
			}

			defineInformation.DefineTemplates[templateKey] = strings.Join(codeText, " ")
		}
	}
}

func buildTemplate(templateData string) string {
	return startSymbol + templateData + endSymbol
}

func (s *DefineAppListener) getDefineInformation() DefineInformation {
	return defineInformation
}
