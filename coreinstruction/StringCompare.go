package coreinstruction

import (
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/regexnew"
)

type StringCompare struct {
	StringSearch
	Content string
}

func NewStringCompare(
	method stringcompareas.Variant,
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: method,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareEqual(
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        search,
		},
		Content: content,
	}
}

func NewStringCompareRegex(
	regex,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Regex,
			Search:        regex,
		},
		Content: content,
	}
}

func NewStringCompareStartsWith(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareEndsWith(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.EndsWith,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareContains(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Contains,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func (it *StringCompare) IsInvalid() bool {
	return it == nil
}

func (it *StringCompare) IsDefined() bool {
	return it != nil
}

func (it *StringCompare) IsMatch() bool {
	if it == nil {
		return true
	}

	return it.CompareMethod.IsCompareSuccess(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}

func (it *StringCompare) IsMatchFailed() bool {
	return !it.IsMatch()
}

func (it *StringCompare) VerifyError() error {
	if it == nil {
		return nil
	}

	if it.CompareMethod.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			it.Content)
	}

	return it.CompareMethod.VerifyError(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}
