package chmodins

type Condition struct {
	IsSkipOnNonExist  bool `json:"IsSkipOnNonExist"`
	IsContinueOnError bool `json:"IsContinueOnError"`
	IsRecursive       bool `json:"IsRecursive"`
}

func DefaultAllTrueCondition() *Condition {
	return &Condition{
		IsSkipOnNonExist:  true,
		IsContinueOnError: true,
		IsRecursive:       true,
	}
}

func DefaultAllFalseCondition() *Condition {
	return &Condition{
		IsSkipOnNonExist:  false,
		IsContinueOnError: false,
		IsRecursive:       false,
	}
}

// DefaultAllFalseExceptRecurse only IsRecursive will be true
func DefaultAllFalseExceptRecurse() *Condition {
	return &Condition{
		IsSkipOnNonExist:  false,
		IsContinueOnError: false,
		IsRecursive:       true,
	}
}

func (receiver *Condition) Clone() *Condition {
	if receiver == nil {
		return nil
	}

	return &Condition{
		IsSkipOnNonExist:  receiver.IsSkipOnNonExist,
		IsContinueOnError: receiver.IsContinueOnError,
		IsRecursive:       receiver.IsRecursive,
	}
}
