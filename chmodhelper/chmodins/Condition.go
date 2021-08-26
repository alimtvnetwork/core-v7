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

func (receiver *Condition) Clone() *Condition {
	if receiver == nil {
		return nil
	}

	return &Condition{
		IsSkipOnInvalid:   receiver.IsSkipOnInvalid,
		IsContinueOnError: receiver.IsContinueOnError,
		IsRecursive:       receiver.IsRecursive,
	}
}

func (receiver Condition) CloneNonPtr() Condition {
	return Condition{
		IsSkipOnInvalid:   receiver.IsSkipOnInvalid,
		IsContinueOnError: receiver.IsContinueOnError,
		IsRecursive:       receiver.IsRecursive,
	}
}
