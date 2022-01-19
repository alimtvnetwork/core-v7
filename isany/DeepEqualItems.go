package isany

func DeepEqualItems(
	items ...interface{},
) (isAllEqual bool) {
	length := len(items)

	if length == 0 || length == 1 {
		return true
	}

	if length == 2 {
		return DeepEqual(items[0], items[1])
	}

	for i := 1; i < length; i++ {
		first := items[i-1]
		second := items[i]

		if NotDeepEqual(first, second) {
			return false
		}
	}

	return true
}
