package chmodins

type RwxInstruction struct {
	RwxOwnerGroupOther
	Condition
}

func (it *RwxInstruction) Clone() *RwxInstruction {
	if it == nil {
		return nil
	}

	return &RwxInstruction{
		RwxOwnerGroupOther: *it.RwxOwnerGroupOther.Clone(),
		Condition: Condition{
			IsSkipOnInvalid:   it.IsSkipOnInvalid,
			IsContinueOnError: it.IsContinueOnError,
			IsRecursive:       it.IsRecursive,
		},
	}
}
