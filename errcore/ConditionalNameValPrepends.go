package errcore

func ConditionalNameValPrepends(
	isAdd bool,
	nameValues []NameVal,
	prependingItems ...NameVal,
) []NameVal {
	if !isAdd || len(prependingItems) == 0 {
		return nameValues
	}

	nameValues = append(prependingItems, nameValues...)

	return nameValues
}
