package stringslice

// ExpandByFunc Don't include nil or empty slice from expand
func ExpandByFunc(
	slice *[]string,
	expandFunc func(line string) *[]string,
) *[]string {
	length := LengthOfPointer(slice)
	if length == 0 {
		return &[]string{}
	}

	sliceOfSlices := make(
		[]*[]string,
		0,
		length)

	for _, line := range *slice {
		expandedLines := expandFunc(line)
		if expandedLines == nil || len(*expandedLines) == 0 {
			continue
		}

		sliceOfSlices = append(sliceOfSlices, expandedLines)
	}

	return MergeNewSlicesPtrOfSlicesPtr(&sliceOfSlices)
}
