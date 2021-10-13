package errcore

func MeaningfulErrorHandle(
	rawErrType RawErrorType,
	funcName string,
	err error,
) {
	if err == nil {
		return
	}

	err2 := MeaningfulError(rawErrType, funcName, err)

	panic(err2)
}
