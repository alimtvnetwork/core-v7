package namevalue

func PrependsIf(
	isAdd bool,
	nameValues []Instance,
	prependingItems ...Instance,
) []Instance {
	if !isAdd || len(prependingItems) == 0 {
		return nameValues
	}

	nameValues = append(prependingItems, nameValues...)

	return nameValues
}
