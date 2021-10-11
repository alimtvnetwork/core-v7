package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func GetSearchLineNumberExpectationMessage(
	counter int,
	lineNumberExpect int,
	lineNumberActualContent int,
	content interface{},
	searchTerm interface{},
	additionalInfo interface{},
) string {
	return fmt.Sprintf(
		msgformats.PrintSearchLineNumberDidntMatchFormat,
		counter,
		lineNumberExpect,
		lineNumberActualContent,
		content,
		searchTerm,
		lineNumberExpect,
		lineNumberActualContent,
		additionalInfo,
	)
}
