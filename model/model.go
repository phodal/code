package model

type CodeFile struct {
	FileId   string
	FullName string
}

type CodePosition struct {
	StartLine       int
	StartLineColumn int
	StopLine        int
	StopLineColumn  int
}

type CodeMember struct {
	MemberId     string
	Name         string
	DataStructId string
	FileId       string
	PackageName  string // namespace
	Position     CodePosition
}

type CodeDataStruct struct {
	DataStructId string
	Name         string
	MemberIds    []string
}

type CodeCommit struct {
	FileId string
}

type CodeType struct {
	Type string
}

type MemberId struct {
	Id string
}

// define function
type CodeFunction struct {
	MemberId          string
	Parameters        []CodeType
	ReturnTypes       []CodeType
	References        []MemberId
	CodeFunctionCalls []CodeFunctionCall
}

type CodeFunctionCall struct {
	MemberId   string
	ReturnVars []CodeType
	Parameters []CodeParameter
}

type CodeParameter struct {
	Key   CodeType
	Value CodeParameterValue
}

type CodeParameterValue struct {
	Value string
}

func CreateParameter() {

}

func CreateFunctionCall(functionName string, parameters []CodeParameter) CodeFunctionCall {
	var call CodeFunctionCall
	switch functionName {
	case "print":
		call = *&CodeFunctionCall{
			MemberId:   functionName,
			ReturnVars: nil,
			Parameters: parameters,
		}
	}

	return call
}

func CreateMember() {

}

func CreateSystemMembers() []CodeMember {
	var systemMembers []CodeMember
	pos := &CodePosition{}
	printMember := &CodeMember{
		MemberId:     "sys" + "." + "print",
		Name:         "print",
		DataStructId: "",
		FileId:       "",
		PackageName:  "sys",
		Position:     *pos,
	}

	systemMembers = append(systemMembers, *printMember)
	return systemMembers
}

func CreateFunction(name string) CodeFunction {
	function := *&CodeFunction{
		MemberId:          "",
		Parameters:        nil,
		ReturnTypes:       nil,
		References:        nil,
		CodeFunctionCalls: nil,
	}

	return function
}

func CreateMainFunction() {
	//var defaultMember = "main"
	//ifunc := &CodeFunction{
	//	MemberId:          defaultMember,
	//	Parameters:        nil,
	//	ReturnTypes:       nil,
	//	References:        nil,
	//	CodeFunctionCalls: nil,
	//}

}
