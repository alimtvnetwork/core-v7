package errcore

func SliceErrorsToStrings(
	errorItems ...error,
) []string {
	if errorItems == nil {
		return []string{}
	}

	slice := make([]string, 0, len(errorItems))

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		slice = append(slice, err.Error())
	}

	return slice
}
