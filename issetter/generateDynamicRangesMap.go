package issetter

func generateDynamicRangesMap() map[string]interface{} {
	newMap := make(map[string]interface{}, len(valuesNames))

	for i, name := range valuesNames {
		newMap[name] = i
	}

	return newMap
}
