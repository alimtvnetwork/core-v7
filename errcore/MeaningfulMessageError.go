package errcore

func MeaningfulMessageError(
	msgType Variation,
	funcName string,
	err error,
	message string,
) error {
	if err == nil {
		return nil
	}

	return msgType.Error(
		funcName,
		err.Error()+message)
}
