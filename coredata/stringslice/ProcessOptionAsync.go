package stringslice

func ProcessOptionAsync(
	isSkipOnNil bool,
	processor func(index int, item interface{}) string,
	items ...interface{},
) []string {
	if len(items) == 0 {
		return []string{}
	}

	list := ProcessAsync(processor, items...)

	if !isSkipOnNil {
		return list
	}

	newSlice := make([]string, 0, len(list))

	for _, item := range list {
		if item == "" {
			continue
		}

		newSlice = append(newSlice, item)
	}

	return newSlice
}
