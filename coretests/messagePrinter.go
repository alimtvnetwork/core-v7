package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

type printMessage struct{}

func (it printMessage) FailedExpected(
	isFailed bool,
	when,
	actual,
	expected interface{},
	counter int,
) {
	if isFailed {
		message := GetAssert.Quick(when, actual, expected, counter)

		fmt.Println(message)
	}
}

// PrintNameValue
//
// Print using msgformats.PrintValuesFormat
func (it printMessage) NameValue(header string, any interface{}) {
	toString := ToStringNameValues(any)

	fmt.Printf(
		msgformats.PrintValuesFormat,
		header,
		any,
		toString,
	)
}

// PrintValue
//
// Print values using msgformats.PrintValuesFormat
func (it printMessage) Value(header string, any interface{}) {
	toString := ToStringValues(any)

	fmt.Printf(
		msgformats.PrintValuesFormat,
		header,
		any,
		toString,
	)
}
