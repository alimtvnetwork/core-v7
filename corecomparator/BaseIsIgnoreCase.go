package corecomparator

type BaseIsIgnoreCase struct {
	IsIgnoreCase bool `json:"IsCaseSensitive,omitempty"` // ignore case compare
}

func (it *BaseIsIgnoreCase) IsCaseSensitive() bool {
	return !it.IsIgnoreCase
}

func (it *BaseIsIgnoreCase) BaseIsCaseSensitive() BaseIsCaseSensitive {
	return BaseIsCaseSensitive{
		IsCaseSensitive: it.IsCaseSensitive(),
	}
}

func (it BaseIsIgnoreCase) Clone() BaseIsIgnoreCase {
	return BaseIsIgnoreCase{
		IsIgnoreCase: it.IsIgnoreCase,
	}
}

func (it *BaseIsIgnoreCase) ClonePtr() *BaseIsIgnoreCase {
	if it == nil {
		return nil
	}

	return &BaseIsIgnoreCase{
		IsIgnoreCase: it.IsIgnoreCase,
	}
}
