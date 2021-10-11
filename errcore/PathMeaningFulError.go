package errcore

func PathMeaningFulError(
	msgType Variation,
	funcName string,
	err error,
	location string,
) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error() +
		", location: [" + location + "]"

	return msgType.Error(
		funcName,
		errMsg)
}
