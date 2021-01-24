package converters

func StringsToMap(inputArray *[]string) *map[string]bool {
	length := len(*inputArray)
	hashset := make(map[string]bool, length)

	for i := 0; i < length; i++ {
		hashset[(*inputArray)[i]] = true
	}

	return &hashset
}
