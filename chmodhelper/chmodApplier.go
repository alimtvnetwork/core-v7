package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/errcore"
)

type chmodApplier struct{}

func (it chmodApplier) Default(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		false,
		location)
}

func (it chmodApplier) OnMismatchOption(
	isApply,
	isSkipOnInvalid bool,
	fileMode os.FileMode,
	location string,
) error {
	if !isApply {
		return nil
	}

	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		isApply,
		true,
		isSkipOnInvalid,
		location)
}

func (it chmodApplier) OnMismatch(
	isSkipOnInvalid bool,
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		true,
		true,
		isSkipOnInvalid,
		location)
}

func (it chmodApplier) OnMismatchSkipInvalid(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmodOptions(
		true,
		true,
		true,
		location)
}

func (it chmodApplier) SkipInvalidFile(
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		true,
		location)
}

func (it chmodApplier) ApplyIf(
	isApply bool,
	fileMode os.FileMode,
	location string,
) error {
	if !isApply {
		return nil
	}

	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	return rwx.ApplyChmod(
		false,
		location)
}

func (it chmodApplier) Options(
	isSkipInvalidPaths,
	isRecursive bool,
	fileMode os.FileMode,
	location string,
) error {
	rwx := New.RwxWrapper.UsingFileModePtr(
		fileMode)

	if isRecursive {
		return rwx.ApplyRecursive(
			isSkipInvalidPaths,
			location)
	}

	return rwx.ApplyChmod(
		isSkipInvalidPaths,
		location)
}

func (it chmodApplier) RecursivePath(
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	location string,
) error {
	return it.Options(
		isSkipInvalidPaths,
		true,
		fileMode,
		location)
}

func (it chmodApplier) RecursivePaths(
	isContinueOnError bool,
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipInvalidPaths,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) RecursivePathsContinueOnError(
	isSkipInvalidPaths bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipInvalidPaths,
			IsContinueOnError: true,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) RecursivePathsCaptureInvalids(
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: false,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeRecursive(
	isContinueOnError bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: isContinueOnError,
			IsRecursive:       true,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeContinueOnErr(
	isRecursive bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsContinueOnError: true,
			IsRecursive:       isRecursive,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeOptions(
	isSkipOnInvalid,
	isContinueOnError,
	isRecursive bool,
	fileMode os.FileMode,
	locations ...string,
) error {
	return it.PathsUsingFileModeConditions(
		fileMode,
		&chmodins.Condition{
			IsSkipOnInvalid:   isSkipOnInvalid,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       isRecursive,
		},
		locations...)
}

func (it chmodApplier) PathsUsingFileModeConditions(
	fileMode os.FileMode,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper := New.RwxWrapper.UsingFileMode(fileMode)

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}

// RwxPartial
//
// can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func (it chmodApplier) RwxPartial(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
		rwxPartial,
		condition)

	if err != nil {
		return err
	}

	return rwxInstructionExecutor.
		ApplyOnPathsPtr(&locations)
}

// RwxStringApplyChmod
//
//	rwxFullString 10 chars "-rwxrwxrwx"
func RwxStringApplyChmod(
	rwxFullString string, // rwxFullString 10 chars "-rwxrwxrwx"
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxFullStringErr := chmodins.GetRwxFullLengthError(
		rwxFullString)
	if rwxFullStringErr != nil {
		return rwxFullStringErr
	}

	if condition == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxFullString)

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}

// RwxOwnerGroupOtherApplyChmod rwxFullString 10 chars "-rwxrwxrwx"
func RwxOwnerGroupOtherApplyChmod(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if rwxOwnerGroupOther == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("rwxOwnerGroupOther")
	}

	if condition == nil {
		return errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxOwnerGroupOther.String())

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
