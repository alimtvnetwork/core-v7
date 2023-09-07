package stringutil

import (
	"math"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

func ToUint32Default(
	s string,
) uint32 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	// fix https://t.ly/6aoW,
	// https://gitlab.com/auk-go/core/-/issues/81
	// use MaxInt32 instead of uint32Max
	if toInt >= 0 && toInt <= math.MaxInt32 {
		return uint32(toInt)
	}

	return constants.Zero
}
