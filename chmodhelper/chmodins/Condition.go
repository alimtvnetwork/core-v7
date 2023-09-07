package chmodins

type Condition struct {
	IsSkipOnInvalid   bool `json:"IsSkipOnInvalid"`
	IsContinueOnError bool `json:"IsContinueOnError"`
	IsRecursive       bool `json:"IsRecursive"`
}

func DefaultAllTrueCondition() *Condition {
	return &Condition{
		IsSkipOnInvalid:   true,
		IsContinueOnError: true,
		IsRecursive:       true,
	}
}

func DefaultAllFalseCondition() *Condition {
	return &Condition{
		IsSkipOnInvalid:   false,
		IsContinueOnError: false,
		IsRecursive:       false,
	}
}

// DefaultAllFalseExceptRecurse only IsRecursive will be true
func DefaultAllFalseExceptRecurse() *Condition {
	return &Condition{
		IsSkipOnInvalid:   false,
		IsContinueOnError: false,
		IsRecursive:       true,
	}
}

// IsExitOnInvalid
//
//  returns true Condition is null or invert of IsSkipOnInvalid
func (it *Condition) IsExitOnInvalid() bool {
	return it == nil || !it.IsSkipOnInvalid
}

// IsCollectErrorOnInvalid
//
//  returns true Condition is null or invert of IsSkipOnInvalid
func (it *Condition) IsCollectErrorOnInvalid() bool {
	return it == nil || !it.IsSkipOnInvalid
}

func (it *Condition) Clone() *Condition {
	if it == nil {
		return nil
	}

	return &Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}
}

func (it Condition) CloneNonPtr() Condition {
	return Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}
}
