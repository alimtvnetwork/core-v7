package msgtype

func MeaningFulError(
	msgType Variation,
	funcName string,
	err error,
) error {
	if err == nil {
		return nil
	}

	return msgType.Error(funcName, err.Error())
}
