package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func StringToIntegerDefault(
	input string,
) int {
	value, err2 := strconv.Atoi(input)

	if err2 != nil {
		return constants.Zero
	}

	return value
}
