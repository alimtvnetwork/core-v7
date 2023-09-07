package stringutil

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ToByteDefault(
	s string,
) byte {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	if toInt >= constants.Zero && toInt <= constants.MaxUnit8AsInt {
		return byte(toInt)
	}

	return constants.Zero
}
