package corejson

func JsonString(anyItem interface{}) (jsonString string, err error) {
	jsonResult := New(anyItem)

	jsonString = jsonResult.JsonString()
	err = jsonResult.MeaningfulError()
	jsonResult.Dispose()

	return jsonString, err
}
