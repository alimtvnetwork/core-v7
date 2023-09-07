package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GetActualAndExpectProcessedMessage(
	counter int,
	actual interface{},
	expected interface{},
	actualProcessed interface{},
	expectedProcessed interface{},
) string {
	return fmt.Sprintf(
		msgformats.PrintActualAndExpectedProcessedFormat,
		counter,
		actual,
		expected,
		actualProcessed,
		expectedProcessed,
	)
}
