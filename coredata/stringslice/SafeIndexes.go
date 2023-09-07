package stringslice

import "gitlab.com/auk-go/core/constants"

// SafeIndexes get safe indexes slice based on indexes given
//
// Reference : https://play.golang.org/p/uKLpK8go0lh
func SafeIndexes(slice []string, indexes ...int) []string {
	lastIndex := len(slice) - 1
	requestLength := len(indexes)
	responseSlice := MakeLen(requestLength)
	if lastIndex < constants.Zero {
		return responseSlice
	}

	for i := 0; i < requestLength; i++ {
		index := indexes[i]
		if index > lastIndex || index < 0 {
			continue
		}

		responseSlice[i] = slice[index]
	}

	return responseSlice
}
