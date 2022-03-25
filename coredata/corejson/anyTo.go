package corejson

type anyTo struct{}

func (it anyTo) SerializedJsonResult(
	fromAny interface{},
) *Result {
	switch castedTo := fromAny.(type) {
	case Result:
		return castedTo.Ptr()
	case *Result:
		return castedTo
	case []byte:
		return NewResult.UsingBytesTypePtr(
			"RawBytes",
			castedTo)
	case string:
		return NewResult.UsingBytesTypePtr(
			"RawString",
			[]byte(castedTo))
	case Jsoner:
		return castedTo.JsonPtr()
	case bytesSerializer:
		return NewResult.UsingSerializer(castedTo)
	}

	return Serialize.Apply(
		fromAny)
}

func (it anyTo) SerializedRaw(
	fromAny interface{},
) (allBytes []byte, err error) {
	return it.SerializedJsonResult(fromAny).Raw()
}

func (it anyTo) SerializedString(
	fromAny interface{},
) (serializedString string, err error) {
	jsonResult := it.SerializedJsonResult(fromAny)

	if jsonResult.HasError() {
		return "", jsonResult.MeaningfulError()
	}

	return jsonResult.JsonString(), nil
}

func (it anyTo) SerializedSafeString(
	fromAny interface{},
) (serializedString string) {
	jsonResult := it.SerializedJsonResult(fromAny)

	if jsonResult.HasError() {
		return ""
	}

	return jsonResult.JsonString()
}

func (it anyTo) SerializedStringMust(
	fromAny interface{},
) (serializedString string) {
	jsonResult := it.SerializedJsonResult(fromAny)
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

// SafeJsonString
//
//  warning : swallows error
func (it anyTo) SafeJsonString(
	anyItem interface{},
) string {
	jsonResult := New(anyItem)

	return jsonResult.JsonString()
}

// SafeJsonPrettyString
//
//  warning : swallows error
func (it anyTo) SafeJsonPrettyString(
	anyItem interface{},
) string {
	jsonResult := New(anyItem)

	return jsonResult.PrettyJsonString()
}

func (it anyTo) JsonStringMust(
	anyItem interface{},
) string {
	jsonResult := New(anyItem)
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it anyTo) PrettyStringMust(
	anyItem interface{},
) string {
	jsonResult := New(anyItem)
	jsonResult.MustBeSafe()

	return jsonResult.PrettyJsonString()
}

func (it anyTo) UsingSerializer(
	serializer bytesSerializer,
) *Result {
	return NewResult.UsingSerializer(
		serializer)
}
