package stringutil

import (
	"math"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func ToUint32Default(
	s string,
) uint32 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	if toInt >= constants.Zero && toInt <= math.MaxUint32 {
		return uint32(toInt)
	}

	return constants.Zero
}
