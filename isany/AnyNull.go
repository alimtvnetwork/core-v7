package isany

func AnyNull(anyItems ...interface{}) bool {
	if len(anyItems) == 0 {
		return true
	}

	for _, anyItem := range anyItems {
		if Null(anyItem) {
			return true
		}
	}

	return false
}
