package coreinstruction

import (
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/regexnew"
)

type StringSearch struct {
	corecomparator.BaseIsIgnoreCase
	CompareMethod stringcompareas.Variant
	Search        string
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

func (it *StringSearch) IsMatch(content string) bool {
	if it == nil {
		return true
	}

	return it.CompareMethod.IsCompareSuccess(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}

func (it *StringSearch) IsAllMatch(contents ...string) bool {
	if len(contents) == 0 {
		return true
	}

	for _, content := range contents {
		if it.IsMatchFailed(content) {
			return false
		}
	}

	return true
}

func (it *StringSearch) IsAnyMatchFailed(contents ...string) bool {
	return !it.IsAllMatch(contents...)
}

func (it *StringSearch) IsMatchFailed(content string) bool {
	return !it.IsMatch(content)
}

func (it *StringSearch) VerifyError(content string) error {
	if it == nil {
		return nil
	}

	if it.CompareMethod.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			content)
	}

	return it.CompareMethod.VerifyError(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}
