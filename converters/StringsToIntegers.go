package converters

import (
	"strconv"
)

// StringsToIntegers panic if not a number
func StringsToIntegers(strArray *[]string) *[]int {
	results := make([]int, len(*strArray))

	for i, v := range *strArray {
		vInt, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		results[i] = vInt
	}

	return &results
}
