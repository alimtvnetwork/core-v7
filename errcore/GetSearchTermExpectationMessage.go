package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GetSearchTermExpectationMessage(
	counter int,
	header string,
	expectationMessage string,
	lineProcessingIndex int,
	actual interface{},
	expected interface{},
	additionalInfo interface{}, // can be nil
) string {
	if additionalInfo == nil {
		return fmt.Sprintf(
			msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedWithoutAdditionalFormat,
			counter,
			header,
			expectationMessage,
			lineProcessingIndex,
			actual,
			expected,
		)
	}

	return fmt.Sprintf(
		msgformats.PrintHeaderForSearchWithActualAndExpectedProcessedFormat,
		counter,
		header,
		expectationMessage,
		lineProcessingIndex,
		actual,
		expected,
		additionalInfo,
	)
}
