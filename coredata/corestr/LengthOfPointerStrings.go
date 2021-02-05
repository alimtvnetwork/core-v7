package corestr

func LengthOfPointerStrings(stringItems *[]*string) int {
	if stringItems == nil || *stringItems == nil {
		return 0
	}

	return len(*stringItems)
}
