package chmodinstruction

type Instructions struct {
	PathModifiers *[]*PathModifier `json:"PathModifiers"`
}

type ChmodCommand struct {
	IsRecursive bool   `json:"IsRecursive"`
	Command     string `json:"Command"`
}

type UserNamePlusGroupName struct {
	UserName  string `json:"UserName"`
	GroupName string `json:"GroupName"`
}

type Chown struct {
	IsRecursive bool `json:"IsRecursive"`
	UserNamePlusGroupName
}

type ChangeGroup struct {
	IsRecursive bool   `json:"IsRecursive"`
	Group       string `json:"Group"`
}

type RwxInstruction struct {
	IsSkipOnNonExist  bool   `json:"IsSkipOnNonExist"`
	IsContinueOnError bool   `json:"IsContinueOnError"`
	IsRecursive       bool   `json:"IsRecursive"`
	Owner             string `json:"Owner"`
	Group             string `json:"Group"`
	Other             string `json:"Other"`
}

type PathModifier struct {
	RwxInstructions *[]*RwxInstruction `json:"RwxInstruction"`
}
