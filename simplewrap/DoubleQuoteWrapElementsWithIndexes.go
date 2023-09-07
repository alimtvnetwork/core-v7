package simplewrap

import (
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

// DoubleQuoteWrapElementsWithIndexes Returns new empty slice if nil or empty slice given.
func DoubleQuoteWrapElementsWithIndexes(
	inputElements ...string,
) (doubleQuoteWrappedItems []string) {
	if inputElements == nil {
		return []string{}
	}

	length := len(inputElements)
	newSlice := make([]string, length)

	if length == 0 {
		return newSlice
	}

	for i, item := range inputElements {
		indexString := constants.SquareStart +
			strconv.Itoa(i) +
			constants.SquareEnd

		newSlice[i] = WithDoubleQuote(
			item + indexString)
	}

	return newSlice
}
