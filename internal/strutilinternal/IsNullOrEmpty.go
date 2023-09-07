package strutilinternal

func IsNullOrEmpty(stringPtr *string) bool {
	return stringPtr == nil || *stringPtr == ""
}
