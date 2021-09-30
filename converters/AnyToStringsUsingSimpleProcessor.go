package converters

func AnyToStringsUsingSimpleProcessor(
	isSkipOnNil bool,
	simpleProcessor func(index int, in interface{}) (out string),
	any interface{},
) []string {
	if any == nil {
		return []string{}
	}

	anyItems := AnyToAnyItems(isSkipOnNil, any)
	slice := make([]string, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out := simpleProcessor(i, item)

		slice[i] = out
	}

	return slice
}
