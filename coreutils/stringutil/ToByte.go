package stringutil

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ToByte(
	s string,
	defVal byte,
) byte {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return defVal
	}

	if toInt >= constants.Zero && toInt <= constants.MaxUnit8AsInt {
		return byte(toInt)
	}

	return defVal
}
