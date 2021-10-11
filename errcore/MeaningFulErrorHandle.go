package errcore

func MeaningfulErrorHandle(
	msgType Variation,
	funcName string,
	err error,
) {
	if err == nil {
		return
	}

	err2 := MeaningfulError(msgType, funcName, err)

	panic(err2)
}
