package corevalidator

type ValidatorParamsBase struct {
	CaseIndex                         int
	IsIgnoreCompareOnActualInputEmpty bool
	IsAttachUserInputs                bool
	IsCaseSensitive                   bool
}

func (it ValidatorParamsBase) IsIgnoreCase() bool {
	return !it.IsCaseSensitive
}
