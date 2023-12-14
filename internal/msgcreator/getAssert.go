package msgcreator

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssert struct{}

// Quick
//
// Gives generic and consistent
// test message using msgformats.QuickIndexInputActualExpectedMessageFormat
func (it getAssert) Quick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexInputActualExpectedMessageFormat,
		counter,
		convertinteranl.AnyTo.SmartString(when),
		convertinteranl.AnyTo.SmartString(actual),
		convertinteranl.AnyTo.SmartString(expected),
	)
}

func (it getAssert) SortedMessage(
	isPrint bool,
	message,
	joiner string,
) string {
	whitespaceRemovedSplits := it.SortedArray(
		isPrint,
		true,
		message,
	)

	return strings.Join(whitespaceRemovedSplits, joiner)
}

func (it getAssert) SortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	if isPrint {
		fmt.Println(message)
	}

	return SplitByEachWordTrimmedNoSpace(
		message,
		isSort,
	)
}

// SortedArrayNoPrint
//
// isPrint: false, isSort: true
func (it getAssert) SortedArrayNoPrint(
	message string,
) []string {
	return it.SortedArray(
		false,
		true, message,
	)
}

// ToStrings
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//
// # Steps:
//  01. string to []string
//  02. []string to as is.
//  03. []interface{} to []string
//  04. map[string]interface{} (fmt - "%s : SmartJson(%s)") to []string
//  05. map[interface{}]interface{} (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//
// See also convertinteranl.AnyTo.Strings
func (it getAssert) ToStrings(
	any interface{},
) []string {
	return convertinteranl.AnyTo.Strings(any)
}

// ToStringsWithSpace
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//	Usages a space prefix for each line.
//
// # Steps:
//  01. string to []string
//  02. []string to as is.
//  03. []interface{} to []string
//  04. map[string]interface{} (fmt - "%s : SmartJson(%s)") to []string
//  05. map[interface{}]interface{} (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//
// See also convertinteranl.AnyTo.Strings
func (it getAssert) ToStringsWithSpace(
	spacePrefixCount int,
	any interface{},
) []string {
	lines := convertinteranl.AnyTo.Strings(any)

	return it.StringsToWithSpaceLines(
		spacePrefixCount,
		lines...,
	)
}

func (it getAssert) ToStringsWithSpaceDefault(
	any interface{},
) []string {
	return it.ToStringsWithSpace(2, any)
}

func (it getAssert) ToStringWithSpace(
	spacePrefixCount int,
	any interface{},
) string {
	lines := convertinteranl.AnyTo.Strings(any)

	withSpace := it.StringsToWithSpaceLines(
		spacePrefixCount,
		lines...,
	)

	return strings.Join(withSpace, constants.NewLineUnix)
}

func (it getAssert) StringsToWithSpaceLines(
	spaceCount int,
	lines ...string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = fmt.Sprintf(
			"%s%s",
			prefix,
			line,
		)
	}

	return newLines
}

func (it getAssert) StringsToSpaceStringUsingFunc(
	spaceCount int,
	toStringFunc func(i int, spacePrefix, line string) string,
	lines ...string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = toStringFunc(
			i,
			prefix,
			line,
		)
	}

	return newLines
}
