package iserror

func AnyEmpty(errorItems ...error) bool {
	if len(errorItems) == 0 {
		return true
	}

	for _, err := range errorItems {
		if err == nil {
			return true
		}
	}

	return false
}
