package corevalidator

type Parameter struct {
	CaseIndex                  int    // which index
	Header                     string // case title
	IsSkipCompareOnActualEmpty bool   // don't activate comparison if actual is empty or nil.
	IsAttachUserInputs         bool   // attach the user inputs to the console output
	IsCaseSensitive            bool
}

func (it Parameter) IsIgnoreCase() bool {
	return !it.IsCaseSensitive
}
