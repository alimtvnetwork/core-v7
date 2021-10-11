package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

// VerifyChmod - expectedHyphenedRwx should be 10 chars example "-rwxrwxrwx"
func VerifyChmod(location string, expectedHyphenedRwx string) error {
	if len(expectedHyphenedRwx) != HyphenedRwxLength {
		return errcore.MeaningfulError(
			errcore.LengthShouldBeEqualToMessage,
			"VerifyChmod"+constants.HypenAngelRight+location,
			errHyphenedRwxLength)
	}

	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) || fileInfo == nil {
		return errcore.MeaningfulError(
			errcore.PathInvalidErrorMessage,
			"VerifyChmod"+constants.HypenAngelRight+location,
			err)
	}

	existingFileMode := fileInfo.Mode().String()[1:]
	if existingFileMode == expectedHyphenedRwx[1:] {
		return nil
	}

	expectationFailedMessage := errcore.ExpectingSimpleNoType(
		chmod,
		expectedHyphenedRwx,
		existingFileMode)

	return errcore.MeaningfulError(
		errcore.PathChmodMismatchErrorMessage,
		"VerifyChmod"+constants.HypenAngelRight+location,
		errors.New(expectationFailedMessage))
}
