package simplewrap

// DoubleQuoteWrapElements Returns new empty slice if nil or empty slice given.
//
// Reference : https://play.golang.org/p/s_uN2-ckk2F | https://stackoverflow.com/a/48832120
func DoubleQuoteWrapElements(
	isSkipQuoteOnlyOnExistence bool,
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

	if isSkipQuoteOnlyOnExistence {
		return wrapDoubleQuoteByExistenceCheck(inputElements, newSlice)
	}

	for i, item := range inputElements {
		newSlice[i] = WithDoubleQuote(item)
	}

	return newSlice
}
