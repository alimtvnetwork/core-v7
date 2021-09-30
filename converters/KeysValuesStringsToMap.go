package converters

// KeysValuesStringsToMap keys nil will return empty map[string]string
func KeysValuesStringsToMap(keys, values []string) map[string]string {
	if keys == nil || values == nil {
		return map[string]string{}
	}

	newArray := make(map[string]string, len(keys))

	for i, key := range keys {
		newArray[key] = values[i]
	}

	return newArray
}
