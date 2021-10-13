package errcore

func MeaningfulError(
	rawErrType RawErrorType,
	funcName string,
	err error,
) error {
	if err == nil {
		return nil
	}

	return rawErrType.Error(
		funcName, err.Error())
}
