package stringslice

func TrimmedEachWordsIf(
	isNonWhitespaceTrim bool,
	slice []string,
) []string {
	if isNonWhitespaceTrim {
		return TrimmedEachWords(slice)
	}

	return NonNullStrings(slice)
}
