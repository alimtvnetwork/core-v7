package args

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

type String string

func (it String) Concat(s ...string) String {
	toStr := stringslice.Joins("", s...)

	return it + String(toStr)
}

func (it String) Join(sep string, s ...string) String {
	toStr := stringslice.Joins(sep, s...)

	return it + String(sep) + String(toStr)
}

func (it String) Split(sep string) corestr.SimpleSlice {
	return strings.Split(string(it), sep)
}

func (it String) DoubleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintValueDoubleQuotationFormat,
			it,
		),
	)
}

func (it String) DoubleQuoteQ() String {
	return String(
		fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			it,
		),
	)
}

func (it String) SingleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintSingleQuoteFormat,
			it,
		),
	)
}

func (it String) ValueDoubleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			it,
		),
	)
}

func (it String) String() string {
	return string(it)
}

func (it String) Bytes() []byte {
	return []byte(it)
}

func (it String) Runes() []rune {
	return []rune(it)
}

func (it String) Length() int {
	return utf8.RuneCountInString(it.String())
}

func (it String) Count() int {
	return utf8.RuneCountInString(it.String())
}

func (it String) IsEmptyOrWhitespace() bool {
	return len(strings.TrimSpace(it.String())) == 0
}

func (it String) TrimSpace() String {
	return String(strings.TrimSpace(it.String()))
}

func (it String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(it.String(), old, new))
}

func (it String) TrimReplaceMap(replacerMap map[string]string) String {
	r := stringutil.ReplaceTemplate.DirectKeyUsingMapTrim(
		it.String(),
		replacerMap,
	)

	return String(r)
}

func (it String) Substring(start, length int) String {
	return it[start:length]
}

func (it String) IsEmpty() bool {
	return len(it) == 0
}

func (it String) HasCharacter() bool {
	return len(it) > 0
}

func (it String) IsDefined() bool {
	return len(it) > 0
}

func (it String) AscIILength() int {
	return len(it)
}
