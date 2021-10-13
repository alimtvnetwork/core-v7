package errcore

func MeaningfulMessageError(
	rawErrType RawErrorType,
	funcName string,
	err error,
	message string,
) error {
	if err == nil {
		return nil
	}

	return rawErrType.Error(
		funcName,
		err.Error()+message)
}
