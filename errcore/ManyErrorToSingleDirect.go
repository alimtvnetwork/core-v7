package errcore

func ManyErrorToSingleDirect(errorItems ...error) error {
	if errorItems == nil || len(errorItems) == 0 {
		return nil
	}

	return ManyErrorToSingle(errorItems)
}
