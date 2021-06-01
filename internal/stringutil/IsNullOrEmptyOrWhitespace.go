package stringutil

func IsNullOrEmptyOrWhitespace(stringPtr *string) bool {
	return stringPtr == nil || IsEmptyOrWhitespace(*stringPtr)
}
