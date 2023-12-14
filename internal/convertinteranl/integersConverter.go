package convertinteranl

import "strconv"

type integersConverter struct{}

func (it integersConverter) ToMapBool(
	items ...int,
) (hashMap map[int]bool) {
	if len(items) == 0 {
		return map[int]bool{}
	}

	hashMap = make(map[int]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}

func (it integersConverter) Int8ToMapBool(
	items ...int8,
) (hashMap map[int8]bool) {
	if len(items) == 0 {
		return map[int8]bool{}
	}

	hashMap = make(map[int8]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}

func (it integersConverter) FromIntegersToMap(inputArray ...int) map[int]bool {
	if len(inputArray) == 0 {
		return map[int]bool{}
	}

	return Map.FromIntegersToMap(inputArray...)
}

func (it integersConverter) IntegersToStrings(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for index, value := range intSlice {
		stringSlice[index] = strconv.Itoa(value)
	}

	return stringSlice
}
