package chmodins

type BaseRwxInstructions struct {
	RwxInstructions []RwxInstruction `json:"RwxInstructions,omitempty"`
}

func (it *BaseRwxInstructions) Length() int {
	if it == nil || it.RwxInstructions == nil {
		return 0
	}

	return len(it.RwxInstructions)
}

func (it *BaseRwxInstructions) IsEmpty() bool {
	return it.Length() == 0
}

func (it *BaseRwxInstructions) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *BaseRwxInstructions) Clone() *BaseRwxInstructions {
	if it == nil {
		return nil
	}

	cloned := it.CloneNonPtr()

	return &cloned
}

func (it BaseRwxInstructions) CloneNonPtr() BaseRwxInstructions {
	instructions := make(
		[]RwxInstruction,
		it.Length())

	for i, instruction := range it.RwxInstructions {
		instructions[i] = *instruction.Clone()
	}

	return BaseRwxInstructions{
		RwxInstructions: instructions,
	}
}
