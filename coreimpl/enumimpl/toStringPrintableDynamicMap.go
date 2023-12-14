package enumimpl

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func toStringPrintableDynamicMap(diffMap DynamicMap) string {
	if diffMap.IsEmpty() {
		return ""
	}

	slice := toStringsSliceOfDiffMap(diffMap)
	compiledString := strings.Join(
		slice,
		constants.CommaUnixNewLine,
	)

	return fmt.Sprintf(
		curlyWrapFormat,
		compiledString,
	)
}

func toStringPrintableDynamicMapLines(diffMap DynamicMap) []string {
	if diffMap.IsEmpty() {
		return []string{}
	}

	toString := toStringPrintableDynamicMap(diffMap)

	return strings.Split(toString, constants.NewLineUnix)
}
