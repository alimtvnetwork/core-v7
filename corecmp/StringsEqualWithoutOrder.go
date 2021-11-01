package corecmp

import "sort"

// StringsEqualWithoutOrder
//
// Sort then compare
func StringsEqualWithoutOrder(leftItems, rightItems []string) bool {
	if leftItems == nil && rightItems == nil {
		return true
	}

	if leftItems == nil || rightItems == nil {
		return false
	}

	length := len(leftItems)
	if length != len(rightItems) {
		return false
	}

	sort.Strings(leftItems)
	sort.Strings(rightItems)

	for i := 0; i < length; i++ {
		if leftItems[i] != rightItems[i] {
			return false
		}
	}

	return true
}
