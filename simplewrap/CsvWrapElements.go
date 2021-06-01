package simplewrap

// DoubleQuoteWrapElements Returns new empty slice if nil or empty slice given.
//
// Reference : https://play.golang.org/p/s_uN2-ckk2F | https://stackoverflow.com/a/48832120
func DoubleQuoteWrapElements(
	inputSlice *[]string,
	isSkipQuoteOnlyOnExistence bool,
) (doubleQuoteWrappedItems *[]string) {
	if inputSlice == nil {
		return &[]string{}
	}

	length := len(*inputSlice)
	newSlice := make([]string, length)

	if length == 0 {
		return &newSlice
	}

	if isSkipQuoteOnlyOnExistence {
		return wrapDoubleQuoteByExistenceCheck(inputSlice, newSlice)
	}

	for i, item := range *inputSlice {
		newSlice[i] = WithDoubleQuote(item)
	}

	return &newSlice
}
