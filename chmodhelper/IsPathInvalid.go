package chmodhelper

func IsPathInvalid(location string) bool {
	return !IsPathExists(location)
}
