package coreinstruction

import (
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/regexnew"
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
			Method:       method,
			Search:       search,
			IsIgnoreCase: isIgnoreCase,
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
			Method: stringcompareas.Equal,
			Search: search,
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
			Method: stringcompareas.Regex,
			Search: regex,
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
			Method:       stringcompareas.StartsWith,
			Search:       search,
			IsIgnoreCase: isIgnoreCase,
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
			Method:       stringcompareas.EndsWith,
			Search:       search,
			IsIgnoreCase: isIgnoreCase,
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
			Method:       stringcompareas.Contains,
			Search:       search,
			IsIgnoreCase: isIgnoreCase,
		},
		Content: content,
	}
}

func (it *StringCompare) IsEmpty() bool {
	return it == nil
}

func (it *StringCompare) IsExist() bool {
	return it != nil
}

func (it *StringCompare) Has() bool {
	return it != nil
}

func (it *StringCompare) IsCaseSensitive() bool {
	return !it.IsIgnoreCase
}

func (it *StringCompare) IsMatch() bool {
	if it == nil {
		return true
	}

	return it.Method.IsCompareSuccess(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}

func (it *StringCompare) VerifyError() error {
	if it == nil {
		return nil
	}

	if it.Method.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			it.Content)
	}

	return it.Method.VerifyError(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}
