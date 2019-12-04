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
	Variables         map[string]string
	Position          CodePosition
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

func CreateFunctionCall(functionName string, parameters []CodeParameter) CodeFunctionCall {
	var call CodeFunctionCall
	switch functionName {
	case "print":
		call = *&CodeFunctionCall{
			MemberId:   functionName,
			ReturnVars: nil,
			Parameters: parameters,
		}
	default:
		call = *&CodeFunctionCall{
			MemberId:   functionName,
			ReturnVars: nil,
			Parameters: parameters,
		}
	}

	return call
}

func CreateFunction(name string) CodeFunction {
	function := *&CodeFunction{
		MemberId:          name,
		Parameters:        nil,
		ReturnTypes:       nil,
		References:        nil,
		CodeFunctionCalls: nil,
		Variables:         make(map[string]string),
	}

	return function
}
