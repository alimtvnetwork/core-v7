package enumimpl

func toHashset(names ...string) map[string]bool {
	if len(names) == 0 {
		return map[string]bool{}
	}

	newMap := make(map[string]bool, len(names)+1)

	for _, name := range names {
		newMap[name] = true
	}

	return newMap
}
