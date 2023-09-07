package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

func PrintValue(header string, any interface{}) {
	toString := ToStringValues(any)

	fmt.Printf(msgformats.PrintValuesFormat, header, any, toString)
}
