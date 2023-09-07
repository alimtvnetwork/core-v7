package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func ToStringValues(any interface{}) string {
	if any == nil {
		return constants.NilAngelBracket
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}
