package converters

import (
	"strconv"

	"gitlab.com/auk-go/core/constants/bitsize"
)

func StringToFloat64Conditional(
	input string, defaultFloat64 float64,
) (value float64, isSuccess bool) {
	value, err2 := strconv.ParseFloat(input, bitsize.Of64)

	if err2 != nil {
		return defaultFloat64, false
	}

	return value, true
}
