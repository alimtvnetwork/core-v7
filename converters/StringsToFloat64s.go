package converters

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants/bitsize"
)

// StringsToFloat64s panic if not a number
func StringsToFloat64s(strArray []string) []float64 {
	results := make([]float64, len(strArray))

	for i, v := range strArray {
		vFloat, err := strconv.ParseFloat(v, bitsize.Of64)

		if err != nil {
			panic(err)
		}

		results[i] = vFloat
	}

	return results
}
