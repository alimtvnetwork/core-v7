package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func GherkinsString(
	testCaseIndex int,
	feature,
	given,
	when,
	then interface{},
) string {
	return fmt.Sprintf(
		msgformats.SimpleGherkinsFormat,
		testCaseIndex,
		feature,
		testCaseIndex,
		given,
		when,
		then,
	)
}
