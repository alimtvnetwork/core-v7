package chmodhelper

import "gitlab.com/evatix-go/core/coredata/corestr"

// GetFilesChmodRwxFullMapDirect returns filePath -> "-rwxrwxrwx"
func GetFilesChmodRwxFullMapDirect(requestedPaths ...string) (*corestr.Hashmap, error) {
	return GetFilesChmodRwxFullMap(requestedPaths)
}
