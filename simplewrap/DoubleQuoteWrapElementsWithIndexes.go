package simplewrap

import (
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

// DoubleQuoteWrapElementsWithIndexes Returns new empty slice if nil or empty slice given.
func DoubleQuoteWrapElementsWithIndexes(
	inputSlice []string,
) (doubleQuoteWrappedItems []string) {
	if inputSlice == nil {
		return []string{}
	}

	length := len(inputSlice)
	newSlice := make([]string, length)

	if length == 0 {
		return newSlice
	}

	for i, item := range inputSlice {
		indexString := constants.SquareStart +
			strconv.Itoa(i) +
			constants.SquareEnd

		newSlice[i] = WithDoubleQuote(
			item + indexString)
	}

	return newSlice
}
