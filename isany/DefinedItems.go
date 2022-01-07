package isany

func DefinedItems(
	anyItems ...interface{},
) (isAllDefined bool, nonNullItems []interface{}) {
	if len(anyItems) == 0 {
		return false, nil
	}

	isAllDefined = true
	nonNullItems = make([]interface{}, 0, len(anyItems))

	for _, anyItem := range anyItems {
		if Null(anyItem) {
			isAllDefined = false
		}

		nonNullItems = append(nonNullItems, anyItem)
	}

	return isAllDefined, nonNullItems
}
