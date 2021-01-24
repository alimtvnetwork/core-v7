package converters

// PointerStringsToStrings will give empty or converted results array (not nil)
func PointerStringsToStrings(pointerStringOfArray *[]*string) *[]string {
	if pointerStringOfArray == nil || *pointerStringOfArray == nil {
		var emptyResult []string

		return &emptyResult
	}

	newArray := make([]string, len(*pointerStringOfArray))

	for i, value := range *pointerStringOfArray {
		newArray[i] = *value
	}

	return &newArray
}
