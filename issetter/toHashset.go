package issetter

func toHashset(names ...string) map[string]bool {
	newMap := make(map[string]bool, len(names)+1)

	for _, name := range names {
		newMap[name] = true
	}

	return newMap
}
