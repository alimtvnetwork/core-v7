package stringslice

func MergeNewPointers(slices ...*[]string) *[]string {
	sliceLength := len(slices)

	if sliceLength == 0 {
		return &[]string{}
	}

	return MergeNewSlicesPtrOfSlices(slices...)
}
