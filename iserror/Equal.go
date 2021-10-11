package iserror

func Equal(left, right error) bool {
	if left == right {
		return true
	}

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Error() == right.Error()
}
