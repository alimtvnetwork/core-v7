package namevalue

func AppendsIf(
	isAdd bool,
	nameValues []Instance,
	appendingItems ...Instance,
) []Instance {
	if !isAdd || len(appendingItems) == 0 {
		return nameValues
	}

	nameValues = append(
		nameValues,
		appendingItems...)

	return nameValues
}
