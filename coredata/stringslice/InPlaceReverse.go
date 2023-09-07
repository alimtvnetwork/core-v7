package stringslice

import "gitlab.com/auk-go/core/constants"

func InPlaceReverse(list *[]string) *[]string {
	if list == nil {
		return &[]string{}
	}

	nonPtrList := *list
	length := len(nonPtrList)

	if length <= 1 {
		return list
	}

	if length == constants.Capacity2 {
		nonPtrList[0], nonPtrList[1] =
			nonPtrList[1], nonPtrList[0]

		return &nonPtrList
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		nonPtrList[i], nonPtrList[lastIndex-i] =
			nonPtrList[lastIndex-i], nonPtrList[i]
	}

	return &nonPtrList
}
