package enumimpl

func stringsToHashSet(slice []string) map[string]bool {
	length := len(slice)
	hashset := make(map[string]bool, length)

	for _, s := range slice {
		hashset[s] = true
	}

	return hashset
}
