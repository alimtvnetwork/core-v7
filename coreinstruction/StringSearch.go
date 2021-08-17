package coreinstruction

import (
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/regexnew"
)

type StringSearch struct {
	IsIgnoreCase bool
	Method       stringcompareas.Variant
	Search       string
}

func (it *StringSearch) IsEmpty() bool {
	return it == nil
}

func (it *StringSearch) IsExist() bool {
	return it != nil
}

func (it *StringSearch) Has() bool {
	return it != nil
}

func (it *StringSearch) IsCaseSensitive() bool {
	return !it.IsIgnoreCase
}

func (it *StringSearch) IsMatch(content string) bool {
	if it == nil {
		return true
	}

	return it.Method.IsCompareSuccess(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}

func (it *StringSearch) VerifyError(content string) error {
	if it == nil {
		return nil
	}

	if it.Method.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			content)
	}

	return it.Method.VerifyError(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}
