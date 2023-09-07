package fsinternal

func IsPathInvalid(location string) bool {
	return !IsPathExists(location)
}
