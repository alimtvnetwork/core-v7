package corecomparator

type BaseIsCaseSensitive struct {
	IsCaseSensitive bool `json:"IsCaseSensitive,omitempty"` // strict case compare
}

func (it *BaseIsCaseSensitive) IsIgnoreCase() bool {
	return !it.IsCaseSensitive
}

func (it *BaseIsCaseSensitive) BaseIsIgnoreCase() BaseIsIgnoreCase {
	return BaseIsIgnoreCase{
		IsIgnoreCase: it.IsIgnoreCase(),
	}
}

func (it BaseIsCaseSensitive) Clone() BaseIsCaseSensitive {
	return BaseIsCaseSensitive{
		IsCaseSensitive: it.IsCaseSensitive,
	}
}

func (it *BaseIsCaseSensitive) ClonePtr() *BaseIsCaseSensitive {
	if it == nil {
		return nil
	}

	return &BaseIsCaseSensitive{
		IsCaseSensitive: it.IsCaseSensitive,
	}
}
