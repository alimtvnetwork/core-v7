package errcore

func PathMeaningFulError(
	rawErrType RawErrorType,
	funcName string,
	err error,
	location string,
) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error() +
		", location: [" + location + "]"

	return rawErrType.Error(
		funcName,
		errMsg)
}
