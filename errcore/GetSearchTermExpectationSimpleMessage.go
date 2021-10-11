package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/internal/msgformats"
)

func GetSearchTermExpectationSimpleMessage(
	counter int,
	expectationErrorMessage string,
	processingIndex int,
	contentProcessed interface{},
	searchTermProcessed interface{},
) string {
	return fmt.Sprintf(
		msgformats.PrintHeaderForSearchActualAndExpectedProcessedSimpleFormat,
		counter,
		expectationErrorMessage,
		processingIndex,
		contentProcessed,
		searchTermProcessed,
	)
}
