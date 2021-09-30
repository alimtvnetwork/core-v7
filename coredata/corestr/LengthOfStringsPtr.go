package corestr

func LengthOfStringsPtr(stringItems *[]string) int {
	if stringItems == nil {
		return 0
	}

	return len(*stringItems)
}
