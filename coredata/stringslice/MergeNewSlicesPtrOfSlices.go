package stringslice

// MergeNewSlicesPtrOfSlices Don't include nil or length 0 slices
func MergeNewSlicesPtrOfSlices(slices ...*[]string) *[]string {
	sliceLength := len(slices)

	if sliceLength == 0 {
		return &[]string{}
	}

	return MergeNewSlicesPtrOfSlicesPtr(&slices)
}
