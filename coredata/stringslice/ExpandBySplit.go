package stringslice

// ExpandBySplit Take each slice item, split and add to the new slice array and returns it.
func ExpandBySplit(slice []string, splitter string) []string {
	length := len(slice)
	if length == 0 {
		return []string{}
	}

	return ExpandBySplits(slice, splitter)
}
