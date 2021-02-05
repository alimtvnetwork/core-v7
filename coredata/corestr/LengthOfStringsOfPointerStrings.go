package corestr

func LengthOfStringsOfPointerStrings(stringItems *[]*[]string) int {
	if stringItems == nil || *stringItems == nil {
		return 0
	}

	return len(*stringItems)
}
