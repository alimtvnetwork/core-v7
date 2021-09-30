package converters

func AnyToStringsUsingProcessor(
	isSkipOnNil bool,
	processor func(index int, in interface{}) (out string, isTake, isBreak bool),
	any interface{},
) []string {
	if any == nil {
		return []string{}
	}

	anyItems := AnyToAnyItems(isSkipOnNil, any)
	slice := make([]string, 0, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out, isTake, isBreak := processor(i, item)

		if isTake {
			slice = append(slice, out)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}
