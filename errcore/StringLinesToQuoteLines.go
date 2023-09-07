package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

// StringLinesToQuoteLines
//
// Each line will be wrapped with "\"%s\", quotation and comma
func StringLinesToQuoteLines(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	slice := make(
		[]string,
		len(lines))

	for i, line := range lines {
		slice[i] = fmt.Sprintf(
			msgformats.LinePrinterFormat,
			line)
	}

	return slice
}
