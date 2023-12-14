package simplewrap

import "gitlab.com/auk-go/core/constants"

func wrapDoubleQuoteByExistenceCheck(
	inputSlice []string,
	newSlice []string,
) []string {
	for i, item := range inputSlice {
		// quote not there or one is there.
		newSlice[i] = ConditionalWrapWith(
			constants.DoubleQuoteChar,
			item,
			constants.DoubleQuoteChar)
	}
	
	return newSlice
}
