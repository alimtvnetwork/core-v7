package iserror

func AllDefined(errorItems ...error) bool {
	if len(errorItems) == 0 {
		return false
	}

	for _, err := range errorItems {
		if err == nil {
			return false
		}
	}

	return true
}
