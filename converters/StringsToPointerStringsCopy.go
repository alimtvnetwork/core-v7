package converters

// StringsToPointerStringsCopy will give empty or converted results array (not nil)
//
// Copy each item to the new array
func StringsToPointerStringsCopy(ptrStrArray *[]string) *[]*string {
	if ptrStrArray == nil || *ptrStrArray == nil {
		var emptyResult []*string

		return &emptyResult
	}

	newArray := make([]*string, len(*ptrStrArray))

	for i, value := range *ptrStrArray {
		// here copy is important
		valueCopy := value
		newArray[i] = &valueCopy
	}

	return &newArray
}
