package keymk

type Option struct {
	Joiner                          string
	IsSkipEmptyEntry, IsUseBrackets bool
	StartBracket, EndBracket        string
}

func (it *Option) IsAddEntryRegardlessOfEmptiness() bool {
	if it == nil {
		return false
	}

	return !it.IsSkipEmptyEntry
}

func (it *Option) ClonePtr() *Option {
	if it == nil {
		return nil
	}

	return &Option{
		Joiner:           it.Joiner,
		IsSkipEmptyEntry: it.IsSkipEmptyEntry,
		IsUseBrackets:    it.IsUseBrackets,
		StartBracket:     it.StartBracket,
		EndBracket:       it.EndBracket,
	}
}

func (it Option) Clone() Option {
	return Option{
		Joiner:           it.Joiner,
		IsSkipEmptyEntry: it.IsSkipEmptyEntry,
		IsUseBrackets:    it.IsUseBrackets,
		StartBracket:     it.StartBracket,
		EndBracket:       it.EndBracket,
	}
}
