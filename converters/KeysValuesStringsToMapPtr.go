package converters

// KeysValuesStringsToMapPtr keys nil will return empty map[string]string
func KeysValuesStringsToMapPtr(keys, values *[]string) *map[string]string {
	if keys == nil || *keys == nil {
		var emptyResult map[string]string

		return &emptyResult
	}

	newArray := make(map[string]string, len(*keys))

	for i, key := range *keys {
		newArray[key] = (*values)[i]
	}

	return &newArray
}
