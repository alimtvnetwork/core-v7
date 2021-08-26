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
