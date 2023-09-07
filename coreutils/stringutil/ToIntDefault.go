package stringutil

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ToIntDefault(
	s string,
) int {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	return toInt
}
