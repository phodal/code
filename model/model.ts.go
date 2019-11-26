package model

type CodeFile struct {
	FullName string
}

type CodePosition struct {
	StartLine       int
	StartLineColumn int
	StopLine        int
	StopLineColumn  int
}

type CodeMember struct {
	Name         string
	DataStructId string
	FileId       string
	PackageName  string // namespace
	Position     CodePosition
}

type CodeDataStruct struct {
	MemberId string
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

type CodeFunction struct {
	MemberId    string
	Parameters  []CodeType
	ReturnTypes []CodeType
	References  []MemberId
}
