package mapdiffinternal

func isStringType(anyItem interface{}) bool {
	_, isSuccess := anyItem.(string)

	return isSuccess
}
