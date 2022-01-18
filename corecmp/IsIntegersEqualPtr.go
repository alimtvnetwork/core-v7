package corecmp

func IsIntegersEqualPtr(leftSlicePtr, rightSlicePtr *[]int) bool {
	if leftSlicePtr == nil && rightSlicePtr == nil {
		return true
	}

	if leftSlicePtr == nil || rightSlicePtr == nil {
		return false
	}

	length := len(*leftSlicePtr)

	if length != len(*rightSlicePtr) {
		return false
	}

	return IsIntegersEqual(*leftSlicePtr, *rightSlicePtr)
}
