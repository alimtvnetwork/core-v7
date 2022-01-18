package corecmp

func IsStringsEqualPtr(leftLines, rightLines *[]string) bool {
	if leftLines == nil && rightLines == nil {
		return true
	}

	if leftLines == nil || rightLines == nil {
		return false
	}

	length := len(*leftLines)

	if length != len(*rightLines) {
		return false
	}

	return IsStringsEqual(*leftLines, *rightLines)
}
