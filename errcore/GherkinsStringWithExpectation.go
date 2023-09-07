package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GherkinsStringWithExpectation(
	testCaseIndex int,
	feature,
	given,
	when,
	then,
	actual,
	expectation interface{},
) string {
	return fmt.Sprintf(
		msgformats.SimpleGherkinsWithExpectationFormat,
		testCaseIndex,
		feature,
		testCaseIndex,
		given,
		when,
		then,
		actual,
		expectation,
	)
}
