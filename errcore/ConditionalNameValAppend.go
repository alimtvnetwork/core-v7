package errcore

func ConditionalNameValAppend(
	isAdd bool,
	nameValues []NameVal,
	appendingItems ...NameVal,
) []NameVal {
	if !isAdd || len(appendingItems) == 0 {
		return nameValues
	}

	nameValues = append(nameValues, appendingItems...)

	return nameValues
}
