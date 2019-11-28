package code

import (
	. "../../languages/define"
	"fmt"
	"reflect"
)

type Symbol struct {
	Key   string
	Value string
}

var symbols []Symbol
var symbolMaps map[string]string

func NewDefineAppListener() *DefineAppListener {
	symbolMaps = make(map[string]string)
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
	symbolMaps[key] = value
}

func (s *DefineAppListener) EnterDefineDeclaration(ctx *DefineDeclarationContext) {
	switch ctx.DefineKey().GetText() {
	case "entry_point":
		buildEntryPoint(ctx)
	}
}

func buildEntryPoint(ctx *DefineDeclarationContext) {
	for _, express := range ctx.AllDefineExpress() {
		fmt.Println(express.(*DefineExpressContext).GetText())
		firstChild := express.(*DefineExpressContext).GetChild(0)
		typeOfExpress := reflect.TypeOf(firstChild).String()
		switch typeOfExpress {
		case "*parser.DefineAttributeContext":
			attributeCtx := firstChild.(*DefineAttributeContext)
			fmt.Println(attributeCtx.GetText())
		case "*parser.DefineTemplateContext":
			defineCtx := firstChild.(*DefineTemplateContext)
			templateKey := defineCtx.IDENTIFIER().GetText()
			defineBody := defineCtx.DefineBody().(*DefineBodyContext)
			funcName := defineBody.STRING_LITERAL().GetText()
			templateData := defineBody.TemplateData().GetText()

			firstSymbol := defineBody.AllSymbolKey()[:1][0].(*SymbolKeyContext).GetText()
			fmt.Println(firstSymbol)
			//fmt.Println(symbolMaps[firstSymbol])
			for _, symbolKey := range defineBody.AllSymbolKey() {
				symbolCtx := symbolKey.(*SymbolKeyContext)
				symbolCtx.GetText()
				//symbolText := symbolCtx.GetText()
			}

			fmt.Println(templateKey, funcName, templateData)
		}
	}
}

func (s *DefineAppListener) getDefineInformation() {
	fmt.Println(symbols)
}
