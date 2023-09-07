package converters

import "strconv"

func IntegersToStrings(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for index, value := range intSlice {
		stringSlice[index] = strconv.Itoa(value)
	}

	return stringSlice
}
