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
