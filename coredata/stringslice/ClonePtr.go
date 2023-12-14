package stringslice

// ClonePtr on nil or empty makes new  &[]string{}
// else makes a copy of itself
func ClonePtr(slice *[]string) (slicePtr *[]string) {
	if slice == nil || len(*slice) == 0 {
		return &[]string{}
	}

	newSlice := make([]string, len(*slice))

	for i, s := range *slice {
		newSlice[i] = s
	}

	return &newSlice
}
