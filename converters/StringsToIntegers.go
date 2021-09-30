package converters

import (
	"strconv"
)

// StringsToIntegers panic if not a number
func StringsToIntegers(isPanic bool, strArray []string) []int {
	results := make([]int, len(strArray))

	for i, v := range strArray {
		vInt, err := strconv.Atoi(v)

		if isPanic && err != nil {
			panic(err)
		} else if err != nil {
			continue
		}

		results[i] = vInt
	}

	return results
}
