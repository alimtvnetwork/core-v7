package chmodhelper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
)

// VerifyChmodUsingHashmap - expectedHyphenedRwx should be 10 chars example "-rwxrwxrwx"
//
// Multiple files verification error will be returned as once. nil will be returned if no error
func VerifyChmodUsingHashmap(
	filePathToRwxMap *corestr.Hashmap,
) error {
	var sliceError []string

	for filePath, expectedRwxFull := range filePathToRwxMap.Items() {
		err := VerifyChmod(filePath, expectedRwxFull)

		if err != nil {
			sliceError = append(sliceError, err.Error())
		}
	}

	return errcore.SliceToError(sliceError)
}
