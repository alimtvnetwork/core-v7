package stringutil

import (
	"math"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func ToUint16Default(
	s string,
) uint16 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	if toInt >= constants.Zero && toInt <= math.MaxUint16 {
		return uint16(toInt)
	}

	return constants.Zero
}
