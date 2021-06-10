package stringslice

import "gitlab.com/evatix-go/core/constants"

func MergeNewSimple(simpleSlices ...[]string) []string {
	sliceLength := len(simpleSlices)

	if sliceLength == 0 {
		return []string{}
	}

	slicesPtr := make([]*[]string, constants.Zero, sliceLength)

	for _, slice := range simpleSlices {
		if len(slice) == 0 {
			continue
		}

		slicesPtr = append(slicesPtr, &slice)
	}

	return *MergeNewSlicesPtrOfSlicesPtr(&slicesPtr)
}
