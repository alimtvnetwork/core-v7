package msgtype

func MeaningFulErrorHandle(
	msgType Variation,
	funcName string,
	err error,
) {
	if err == nil {
		return
	}

	err2 := MeaningFulError(msgType, funcName, err)

	panic(err2)
}
