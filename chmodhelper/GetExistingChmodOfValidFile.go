package chmodhelper

import "os"

// GetExistingChmodOfValidFile
//
//	on file invalid returns 0
//
// Warning:
//   - error swallowed.
func GetExistingChmodOfValidFile(filePath string) (chmod os.FileMode, isInvalid bool) {
	fileInfo, err := os.Stat(filePath)

	if err != nil || fileInfo == nil {
		return 0, true
	}

	return fileInfo.Mode(), false
}
