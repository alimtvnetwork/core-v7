package stringslice

func NonEmptySlicePtr(slice *[]string) *[]string {
	if slice == nil {
		return &[]string{}
	}

	length := len(*slice)

	if length == 0 {
		return &[]string{}
	}

	newSlice := MakeDefault(length)

	for _, s := range *slice {
		if s != "" {
			newSlice = append(newSlice, s)
		}
	}

	return &newSlice
}
