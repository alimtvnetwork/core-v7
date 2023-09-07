package csvinternal

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func StringersToStringDefault(
	stringerFunctions ...fmt.Stringer,
) string {
	return StringersToString(
		constants.CommaSpace,
		true,
		false,
		stringerFunctions...)
}
