package errcore

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/internal/msgformats"
)

// LinesToDoubleQuoteLinesWithTabs
//
// Each line will be wrapped with "\"%s\", quotation and comma
func LinesToDoubleQuoteLinesWithTabs(
	tabCount int,
	lines []string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	slice := make(
		[]string,
		len(lines),
	)

	space := strings.Repeat(" ", tabCount)

	for i, line := range lines {
		slice[i] = fmt.Sprintf(
			space+msgformats.LinePrinterFormat,
			line,
		)
	}

	return slice
}
