package simplewrap

import "gitlab.com/auk-go/core/constants"

func wrapDoubleQuoteByExistenceCheck(
	inputSlice []string,
	newSlice []string,
) []string {
	for i, item := range inputSlice {
		itemLength := len(item)
		if itemLength < 2 {
			newSlice[i] = WithDoubleQuote(item)

			continue
		}

		// more than 2 char
		if item[0] == constants.DoubleQuoteChar && item[itemLength-1] == constants.DoubleQuoteChar {
			continue
		}

		// quote not there or one is there.
		newSlice[i] = WithDoubleQuote(item)
	}

	return newSlice
}
