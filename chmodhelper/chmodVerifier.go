package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/fsinternal"
)

type chmodVerifier struct{}

// IsEqualRwxFull
//
//	expectedHyphenedRwx must be 10 chars in "-rwxrwxrwx"
//
//	Hint. os.FileMode. String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Format (length must be 10)
//
//	"-rwxrwxrwx"
//
// Example:
//   - owner all enabled only "-rwx------"
//   - group all enabled only "----rwx---"
//
// Must have or restrictions:
//   - string length must be 10.
//
// Reference:
//   - https://ss64.com/bash/chmod.html
func (it chmodVerifier) IsEqualRwxFull(
	location string,
	expectedHyphenedRwx string,
) bool {
	return IsChmod(
		location,
		expectedHyphenedRwx)
}

// IsEqualRwxFullSkipInvalid
//
//	On invalid path it is assumed to be equal.
//	expectedHyphenedRwx must be 10 chars in "-rwxrwxrwx"
//
//	Hint. os.FileMode. String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Format (length must be 10)
//
//	"-rwxrwxrwx"
//
// Example:
//   - owner all enabled only "-rwx------"
//   - group all enabled only "----rwx---"
//
// Must have or restrictions:
//   - string length must be 10.
//
// Reference:
//   - https://ss64.com/bash/chmod.html
func (it chmodVerifier) IsEqualRwxFullSkipInvalid(
	location string,
	expectedHyphenedRwx string,
) bool {
	if fsinternal.IsPathInvalid(location) {
		return true
	}

	return IsChmod(
		location,
		expectedHyphenedRwx)
}

func (it chmodVerifier) IsEqual(
	location string,
	expectedFileMode os.FileMode,
) bool {
	return IsChmod(
		location,
		expectedFileMode.String())
}

// IsEqualSkipInvalid
//
//	On invalid path it is assumed to be equal.
func (it chmodVerifier) IsEqualSkipInvalid(
	location string,
	expectedFileMode os.FileMode,
) bool {
	if fsinternal.IsPathInvalid(location) {
		return true
	}

	return IsChmod(
		location,
		expectedFileMode.String())
}

func (it chmodVerifier) IsMismatch(
	location string,
	expectedFileMode os.FileMode,
) bool {
	return !IsChmod(
		location,
		expectedFileMode.String())
}

func (it chmodVerifier) MismatchError(
	location string,
	expectedFileMode os.FileMode,
) error {
	return it.RwxFull(
		location,
		expectedFileMode.String())
}

func (it chmodVerifier) MismatchErrorUsingRwxFull(
	location string,
	rwxFull string,
) error {
	return it.RwxFull(
		location,
		rwxFull)
}

// GetRwxFull
//
//	Hint. os.FileMode. String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
func (it chmodVerifier) GetRwxFull(fileMode os.FileMode) string {
	return fileMode.String()
}

// GetRwx9
//
//	return "rwxrwxrwx"
//
//	Hint. os.FileMode. String() returns "-rwxrwxrwx" full rwx
//	then substring "-rwxrwxrwx"[1:] to return "rwxrwxrwx"
//
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Format (length must be 9)
//
//	"rwxrwxrwx"
//
// Understanding Examples:
//   - owner all enabled only "rwx------"
//   - group all enabled only "---rwx---"
//
// Reference:
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it chmodVerifier) GetRwx9(fileMode os.FileMode) string {
	rwxFull := fileMode.String()

	if len(rwxFull) > 9 {
		return rwxFull[1:]
	}

	return ""
}

func (it chmodVerifier) GetExisting(
	filePath string,
) (os.FileMode, error) {
	return GetExistingChmod(filePath)
}

func (it chmodVerifier) GetExistingRwxWrapper(
	location string,
) (RwxWrapper, error) {
	return GetExistingChmodRwxWrapper(location)
}

func (it chmodVerifier) GetExistingRwxWrapperMust(
	location string,
) RwxWrapper {
	rwx, err := GetExistingChmodRwxWrapper(location)

	if err != nil {
		panic(err)
	}

	return rwx
}

func (it chmodVerifier) GetExistingChmodRwxWrappers(
	isContinueOnError bool,
	locations ...string,
) (filePathToRwxWrapper map[string]*RwxWrapper, err error) {
	return GetExistingChmodRwxWrappers(
		isContinueOnError,
		locations...)
}

func (it chmodVerifier) GetExistsFilteredPathFileInfoMap(
	isSkipOnInvalid bool,
	locations ...string,
) *FilteredPathFileInfoMap {
	return GetExistsFilteredPathFileInfoMap(
		isSkipOnInvalid,
		locations...)
}

func (it chmodVerifier) PathIf(
	isVerify bool,
	location string,
	expectedFileMode os.FileMode,
) error {
	if !isVerify {
		return nil
	}

	return it.RwxFull(
		location,
		expectedFileMode.String())
}

func (it chmodVerifier) Path(
	location string,
	expectedFileMode os.FileMode,
) error {
	return it.RwxFull(
		location,
		expectedFileMode.String())
}

// RwxFull
//
//	expectedHyphenedRwx must be 10 chars in "-rwxrwxrwx"
//	Hint. os.FileMode.String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Format (length must be 10)
//
//	"-rwxrwxrwx"
//
// Example:
//   - owner all enabled only "-rwx------"
//   - group all enabled only "----rwx---"
//
// Must have or restrictions:
//   - string length must be 10.
//
// Reference:
//   - https://ss64.com/bash/chmod.html
func (it chmodVerifier) RwxFull(
	location,
	expectedHyphenedRwx string,
) error {
	if len(expectedHyphenedRwx) != HyphenedRwxLength {
		return errcore.MeaningfulError(
			errcore.LengthShouldBeEqualToType,
			"VerifyChmod"+constants.HyphenAngelRight+location,
			errHyphenedRwxLength)
	}

	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) || fileInfo == nil {
		return errcore.MeaningfulError(
			errcore.PathInvalidErrorType,
			"VerifyChmod"+constants.HyphenAngelRight+location,
			err)
	}

	existingFileMode := fileInfo.Mode().String()[1:]
	if existingFileMode == expectedHyphenedRwx[1:] {
		return nil
	}

	expectationFailedMessage := errcore.ExpectingSimpleNoType(
		chmodExpectationFailed,
		expectedHyphenedRwx,
		existingFileMode)

	return errcore.MeaningfulError(
		errcore.PathChmodMismatchErrorType,
		"VerifyChmod"+constants.HyphenAngelRight+location,
		errors.New(expectationFailedMessage))
}

// PathsUsingPartialRwxOptions
//
// partialRwx can be any length in
// between 1-10 (rest will be fixed by wildcard)
//
// partialRwx:
//   - "-rwx"    will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
//
// partialRwx Restrictions:
//   - cannot have first char other than hyphen(-) or else things will not work
//
// Options:
//   - isContinueOnError : on true don't stop until all locations are captured.
//   - isSkipOnInvalid   : on true invalid paths will not be considered into expectation.
func (it chmodVerifier) PathsUsingPartialRwxOptions(
	isContinueOnError,
	isSkipOnInvalid bool,
	partialRwx string,
	locations ...string,
) error {
	varWrapper, err := NewRwxVariableWrapper(partialRwx)

	if err != nil {
		return err
	}

	status := varWrapper.RwxMatchingStatus(
		isContinueOnError,
		isSkipOnInvalid,
		locations)

	return status.CreateErrFinalError()
}

// PathsUsingFileModeImmediateReturn
//
// on error quick return, don't wait for all.
func (it chmodVerifier) PathsUsingFileModeImmediateReturn(
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingRwxFull(
		false,
		fileMode.String(),
		locations...)
}

// PathsUsingFileModeContinueOnError
//
// continue on error and return collected error
func (it chmodVerifier) PathsUsingFileModeContinueOnError(
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingRwxFull(
		true,
		fileMode.String(),
		locations...)
}

// PathsUsingFileMode
//
// Options:
//   - isContinueOnError : on true don't stop until all locations are captured.
func (it chmodVerifier) PathsUsingFileMode(
	isContinueOnError bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingRwxFull(
		isContinueOnError,
		fileMode.String(),
		locations...)
}

// PathsUsingRwxFull
//
// expectedHyphenedRwx Format (length must be 10)
//
//	"-rwxrwxrwx"
//	Hint. os.FileMode. String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// expectedHyphenedRwx Example:
//   - owner all enabled only "-rwx------"
//   - group all enabled only "----rwx---"
//
// expectedHyphenedRwx Must have or restrictions:
//   - string length must be 10.
//
// Reference:
//   - https://ss64.com/bash/chmod.html
//
// Options:
//   - isContinueOnError : on true don't stop until all locations are captured.
func (it chmodVerifier) PathsUsingRwxFull(
	isContinueOnError bool,
	expectedHyphenedRwx string,
	locations ...string,
) error {
	if locations == nil || len(locations) == 0 {
		return errcore.CannotBeNilOrEmptyType.
			Error(constants.EmptyString, nil)
	}

	if !isContinueOnError {
		for _, location := range locations {
			err := it.RwxFull(location, expectedHyphenedRwx)

			if err != nil {
				return err
			}
		}
	}

	slice := corestr.New.Collection.Cap(constants.Zero)

	for _, location := range locations {
		err := it.RwxFull(location, expectedHyphenedRwx)

		//goland:noinspection ALL
		slice.AddIf(err != nil, err.Error())
	}

	return errcore.SliceErrorDefault(slice.ListPtr())
}

// UsingHashmap
//
//	Key - > Path, Value -> RwxFullString (10 chars, "-rwx------")
//
// map[key]RwxFullValue - RwxFullValue - Format (length must be 10)
//
//	"-rwxrwxrwx"
//
// RwxFullString Example:
//   - owner all enabled only "-rwx------"
//   - group all enabled only "----rwx---"
//
// RwxFullString Must have or restrictions:
//   - string length must be 10.
//
// RwxFullString Reference:
//   - https://ss64.com/bash/chmod.html
//
// Multiple files verification error will be returned as once.
// nil will be returned if no error
func (it chmodVerifier) UsingHashmap(
	filePathToRwxMap *corestr.Hashmap,
) error {
	var sliceError []string

	for filePath, expectedRwxFull := range filePathToRwxMap.Items() {
		err := it.RwxFull(filePath, expectedRwxFull)

		if err != nil {
			sliceError = append(sliceError, err.Error())
		}
	}

	return errcore.SliceToError(sliceError)
}

func (it chmodVerifier) UsingRwxOwnerGroupOther(
	rwx *chmodins.RwxOwnerGroupOther,
	location string,
) error {
	if rwx == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			Error("rwx is nil", location)
	}

	return it.RwxFull(
		location,
		rwx.String())
}
