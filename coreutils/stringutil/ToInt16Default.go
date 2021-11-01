package stringutil

import (
	"math"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

func ToInt16Default(
	s string,
) int16 {
	toInt, err := strconv.Atoi(s)

	if err != nil {
		return constants.Zero
	}

	if toInt >= math.MinInt16 && toInt <= math.MaxInt16 {
		return int16(toInt)
	}

	return constants.Zero
}
