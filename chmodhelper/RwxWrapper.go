package chmodhelper

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/fsinternal"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/osconsts"
)

type RwxWrapper struct {
	Owner, Group, Other Attribute
}

func (it *RwxWrapper) IsEmpty() bool {
	return it == nil ||
		it.Owner.IsEmpty() &&
			it.Group.IsEmpty() &&
			it.Group.IsEmpty()
}

func (it *RwxWrapper) IsNull() bool {
	return it == nil
}

func (it *RwxWrapper) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *RwxWrapper) IsDefined() bool {
	return !it.IsEmpty()
}

func (it *RwxWrapper) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *RwxWrapper) Verify(location string) error {
	return ChmodVerify.RwxFull(location, it.ToFullRwxValueString())
}

func (it *RwxWrapper) VerifyPaths(isContinueOnError bool, locations ...string) error {
	return ChmodVerify.PathsUsingRwxFull(
		isContinueOnError,
		it.ToFullRwxValueString(),
		locations...,
	)
}

func (it *RwxWrapper) HasChmod(location string) bool {
	return IsChmod(location, it.ToFullRwxValueString())
}

// Bytes
//
// return rwx, (Owner)(Group)(Other) byte values under 1-7
func (it *RwxWrapper) Bytes() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := it.Owner.ToSum()
	group := it.Group.ToSum()
	other := it.Other.ToSum()

	allBytes := [3]byte{owner, group, other}

	return allBytes
}

func (it *RwxWrapper) ToUint32Octal() uint32 {
	// # https://play.golang.org/p/dX-wsvJmFie
	str := it.ToFileModeString()

	// # https://bit.ly/35aBepk
	octal, err := strconv.ParseUint(str, bitsize.Of8, bitsize.Of32)

	if err != nil {
		errcore.
			MeaningfulErrorHandle(
				errcore.PathChmodConvertFailedType,
				"ToUint32Octal",
				err)
	}

	return uint32(octal)
}

// ToCompiledOctalBytes4Digits return 0rwx, '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 0777, 0555, 0755 NOT 0rwx
func (it *RwxWrapper) ToCompiledOctalBytes4Digits() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := it.Owner.ToStringByte()
	group := it.Group.ToStringByte()
	other := it.Other.ToStringByte()

	allBytes := [4]byte{
		constants.ZeroChar,
		owner,
		group,
		other,
	}

	return allBytes
}

// ToCompiledOctalBytes3Digits return '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 777, 555, 755 NOT rwx
// return
//      owner -> (0 - 7 value)
//      group -> (0 - 7 value)
//      other -> (0 - 7 value)
func (it *RwxWrapper) ToCompiledOctalBytes3Digits() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := it.Owner.ToStringByte()
	group := it.Group.ToStringByte()
	other := it.Other.ToStringByte()

	allBytes := [3]byte{
		owner,
		group,
		other,
	}

	return allBytes
}

// ToCompiledSplitValues
// return
//      owner -> (0 - 7 value)
//      group -> (0 - 7 value)
//      other -> (0 - 7 value)
//      eg. 777, 755 etc
func (it *RwxWrapper) ToCompiledSplitValues() (owner, group, other byte) {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner = it.Owner.ToStringByte()
	group = it.Group.ToStringByte()
	other = it.Other.ToStringByte()

	return owner, group, other
}

// ToFileModeString 4 digit string 0rwx, example 0777
func (it *RwxWrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := it.ToCompiledOctalBytes4Digits()

	return string(allBytes[:])
}

// ToRwxCompiledStr 3 digit string, example 777
func (it *RwxWrapper) ToRwxCompiledStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := it.ToCompiledOctalBytes4Digits()

	return string(allBytes[1:])
}

// ToFullRwxValueString
//
//  returns "-rwxrwxrwx" / RwxFull (10)
func (it *RwxWrapper) ToFullRwxValueString() string {
	owner := it.Owner.ToRwxString()
	group := it.Group.ToRwxString()
	other := it.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return constants.Hyphen + owner + group + other
}

// ToFullRwxValueStringExceptHyphen returns "rwxrwxrwx", 9 chars
func (it *RwxWrapper) ToFullRwxValueStringExceptHyphen() string {
	owner := it.Owner.ToRwxString()
	group := it.Group.ToRwxString()
	other := it.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return owner + group + other
}

// ToFullRwxValuesChars "-rwxrwxrwx" Bytes values
func (it *RwxWrapper) ToFullRwxValuesChars() []byte {
	str := it.ToFullRwxValueString()
	chars := []byte(str)

	return chars
}

func (it *RwxWrapper) String() string {
	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return it.ToFullRwxValueString()
}

func (it *RwxWrapper) ToFileMode() os.FileMode {
	// # https://play.golang.org/p/dX-wsvJmFie
	octalUint32 := it.ToUint32Octal()

	return os.FileMode(octalUint32)
}

func (it *RwxWrapper) ApplyChmod(
	isSkipOnInvalid bool,
	fileOrDirectoryPath string,
) error {
	isFileExist := fsinternal.IsPathExists(fileOrDirectoryPath)

	if isSkipOnInvalid && !isFileExist {
		return nil
	}

	if !isSkipOnInvalid && !isFileExist {
		return errcore.
			PathInvalidErrorType.
			Error(
				messages.PathNotExist, fileOrDirectoryPath)
	}

	err := os.Chmod(fileOrDirectoryPath, it.ToFileMode())

	if err != nil {
		return errcore.
			PathChmodApplyType.
			Error(err.Error(), fileOrDirectoryPath)
	}

	return nil
}

// LinuxApplyRecursive skip if it is a non dir path
func (it *RwxWrapper) LinuxApplyRecursive(
	isSkipOnInvalid bool,
	location string,
) error {
	isPathExists := fsinternal.IsPathExists(location)

	if isSkipOnInvalid && !isPathExists {
		return nil
	}

	if !isSkipOnInvalid && !isPathExists {
		return errcore.
			PathInvalidErrorType.
			Error(pathInvalidMessage,
				location)
	}

	return it.applyLinuxRecursiveChmodUsingCmd(
		location)
}

// ApplyRecursive skip if it is a non dir path
func (it *RwxWrapper) ApplyRecursive(
	isSkipOnInvalid bool,
	location string,
) error {
	stat := GetPathExistStat(location)

	if isSkipOnInvalid && !stat.IsExist {
		return nil
	}

	if !isSkipOnInvalid && !stat.IsExist {
		return errcore.
			PathInvalidErrorType.
			Error(pathInvalidMessage,
				location)
	}

	if osconsts.IsLinux {
		return it.LinuxApplyRecursive(
			false,
			location)
	}

	mode := it.ToFileMode()

	if stat.IsFile() {
		return os.Chmod(location, mode)
	}

	var sliceErr []string

	finalErr := RecursivePathsApply(
		location,
		func(currentPath string, info fs.FileInfo, err error) error {
			if err != nil {
				sliceErr = append(
					sliceErr,
					errcore.
						PathInvalidErrorType.Combine(
						err.Error()+pathInvalidMessage,
						currentPath))

				return err
			}

			if info == nil {
				sliceErr = append(
					sliceErr,
					errcore.
						PathInvalidErrorType.Combine(
						pathInvalidMessage,
						currentPath))

				return err
			}

			err2 := os.Chmod(currentPath, mode)

			if err2 != nil {
				sliceErr = append(
					sliceErr,
					errcore.
						PathInvalidErrorType.Combine(
						err2.Error()+pathInvalidMessage,
						currentPath))

				return err2
			}

			return nil
		})

	if finalErr != nil {
		sliceErr = append(
			sliceErr,
			errcore.
				PathInvalidErrorType.Combine(
				finalErr.Error()+pathInvalidMessage,
				location))
	}

	return errcore.SliceToError(sliceErr)
}

func (it *RwxWrapper) applyLinuxRecursiveChmodUsingCmd(location string) error {
	cmd := it.getLinuxRecursiveCmdForChmod(location)

	if cmd == nil {
		return errcore.
			FailedToCreateCmdType.Error(
			constants.BashCommandline,
			location)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return errcore.
			FailedToCreateCmdType.Error(
			constants.ChmodCommand,
			err.Error()+constants.NewLineUnix+stderr.String()+"location:"+location)
	}

	return nil
}

func (it *RwxWrapper) getLinuxRecursiveCmdForChmod(dirPath string) *exec.Cmd {
	instructionLine := constants.ChmodCommand +
		constants.Space +
		constants.RecursiveCommandFlag +
		constants.Space +
		it.ToRwxCompiledStr() +
		constants.Space +
		dirPath

	return exec.Command(
		constants.BinShellCmd,
		constants.NonInteractiveFlag,
		instructionLine)
}

func (it *RwxWrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(
		fileOrDirectoryPath,
		it.ToFileMode())

	if err != nil {
		finalErr := errors.New(err.Error() + fileOrDirectoryPath)

		panic(errcore.MeaningfulError(
			errcore.PathChmodApplyType,
			"MustApplyChmod",
			finalErr))
	}
}

func (it *RwxWrapper) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	return &chmodins.RwxOwnerGroupOther{
		Owner: it.Owner.ToRwxString(),
		Group: it.Group.ToRwxString(),
		Other: it.Other.ToRwxString(),
	}
}

func (it *RwxWrapper) ToRwxInstruction(
	condition *chmodins.Condition,
) *chmodins.RwxInstruction {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()

	return &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition:          *condition,
	}
}

func (it *RwxWrapper) Clone() *RwxWrapper {
	if it == nil {
		return nil
	}

	return &RwxWrapper{
		Owner: *it.Owner.Clone(),
		Group: *it.Group.Clone(),
		Other: *it.Other.Clone(),
	}
}

func (it *RwxWrapper) applyLinuxChmodOnManyNonRecursive(
	condition *chmodins.Condition,
	locations []string,
) error {
	if condition.IsContinueOnError {
		// continue on error
		return it.applyLinuxChmodNonRecursiveManyContinueOnError(
			condition,
			locations)
	}

	for _, location := range locations {
		err := it.ApplyChmod(
			condition.IsSkipOnInvalid,
			location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *RwxWrapper) ApplyLinuxChmodOnMany(
	condition *chmodins.Condition,
	locations ...string,
) error {
	if condition.IsRecursive {
		return it.applyLinuxChmodOnManyRecursive(
			condition,
			locations)
	}

	return it.applyLinuxChmodOnManyNonRecursive(
		condition, locations)
}

func (it *RwxWrapper) applyLinuxChmodOnManyRecursive(
	condition *chmodins.Condition,
	locations []string,
) error {
	if condition.IsContinueOnError {
		// continue on error
		return it.applyLinuxChmodRecursiveManyContinueOnError(
			condition,
			locations)
	}

	for _, location := range locations {
		err := it.LinuxApplyRecursive(
			condition.IsSkipOnInvalid,
			location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *RwxWrapper) applyLinuxChmodRecursiveManyContinueOnError(
	condition *chmodins.Condition,
	locations []string,
) error {
	var errSlice []string

	for _, location := range locations {
		err := it.LinuxApplyRecursive(
			condition.IsSkipOnInvalid,
			location)

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	return errcore.SliceToErrorPtr(&errSlice)
}

func (it *RwxWrapper) applyLinuxChmodNonRecursiveManyContinueOnError(
	condition *chmodins.Condition,
	locations []string,
) error {
	var errSlice []string

	for _, location := range locations {
		err := it.ApplyChmod(
			condition.IsSkipOnInvalid,
			location)

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	return errcore.SliceToErrorPtr(&errSlice)
}

// IsEqualVarWrapper if rwxVariableWrapper nil then returns false
func (it *RwxWrapper) IsEqualVarWrapper(
	rwxVariableWrapper *RwxVariableWrapper,
) bool {
	if rwxVariableWrapper == nil {
		return false
	}

	return rwxVariableWrapper.IsEqualRwxWrapperPtr(
		it)
}

// IsRwxEqualFileInfo if fileInfo nil then returns false
func (it *RwxWrapper) IsRwxEqualFileInfo(
	fileInfo os.FileInfo,
) bool {
	if fileInfo == nil {
		return false
	}

	return it.IsRwxFullEqual(
		fileInfo.Mode().String())
}

func (it *RwxWrapper) IsRwxEqualLocation(
	location string,
) bool {
	fileInfo, _ := os.Stat(location)

	if fileInfo == nil {
		return false
	}

	return it.IsRwxFullEqual(
		fileInfo.Mode().String())
}

func (it *RwxWrapper) IsRwxFullEqual(
	rwxFull string,
) bool {
	if len(rwxFull) < chmodins.RwxFullLength {
		return false
	}

	return it.ToFullRwxValueStringExceptHyphen() == rwxFull[1:]
}

func (it *RwxWrapper) IsEqualPtr(
	next *RwxWrapper,
) bool {
	if it == nil && next == nil {
		return true
	}

	if it == nil || next == nil {
		return false
	}

	return it.Owner.IsEqual(next.Owner) &&
		it.Group.IsEqual(next.Group) &&
		it.Other.IsEqual(next.Other)
}

func (it *RwxWrapper) IsEqualFileMode(
	mode os.FileMode,
) bool {
	toString := mode.String()[1:]
	wrapperString := it.ToFullRwxValueStringExceptHyphen()

	return toString == wrapperString
}

func (it *RwxWrapper) IsNotEqualFileMode(
	mode os.FileMode,
) bool {
	return !it.IsEqualFileMode(mode)
}

func (it RwxWrapper) ToPtr() *RwxWrapper {
	return &it
}

func (it *RwxWrapper) ToNonPtr() RwxWrapper {
	return *it
}

func (it RwxWrapper) MarshalJSON() ([]byte, error) {
	model := rwxWrapperModel{
		Chmod:   it.ToFileModeString(),
		RwxFull: it.ToFullRwxValueString(),
	}

	return corejson.Serialize.Raw(model)
}

func (it *RwxWrapper) UnmarshalJSON(jsonBytes []byte) error {
	var model rwxWrapperModel
	err := corejson.Deserialize.UsingBytes(
		jsonBytes, &model)

	if err == nil {
		// success
		*it, err = New.
			RwxWrapper.
			RwxFullString(model.RwxFull)
	}

	return err
}

// FriendlyDisplay
//
//  fileModeStringFriendlyDisplayFormat : "{chmod : \"%s (%s)\"}"
//  fileModeStringFriendlyDisplayFormat : "{chmod : \"0777 (-rw...)\"}"
func (it RwxWrapper) FriendlyDisplay() string {
	return fmt.Sprintf(
		fileModeStringFriendlyDisplayFormat,
		it.ToFileModeString(),
		it.ToFullRwxValueString())
}

func (it RwxWrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it RwxWrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *RwxWrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	return jsonResult.Deserialize(it)
}

func (it RwxWrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
