package chmodhelper

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strconv"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/internal/fsinternal"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

type RwxWrapper struct {
	Owner, Group, Other Attribute
}

func (rwxWrapper *RwxWrapper) Verify(location string) error {
	return VerifyChmod(location, rwxWrapper.ToFullRwxValueString())
}

func (rwxWrapper *RwxWrapper) VerifyPaths(location *[]string, isContinueOnError bool) error {
	return VerifyChmodPaths(
		location,
		rwxWrapper.ToFullRwxValueString(),
		isContinueOnError)
}

func (rwxWrapper *RwxWrapper) HasChmod(location string) bool {
	return IsChmod(location, rwxWrapper.ToFullRwxValueString())
}

// Bytes return rwx, (Owner)(Group)(Other) byte values under 1-7
func (rwxWrapper *RwxWrapper) Bytes() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := rwxWrapper.Owner.ToSum()
	group := rwxWrapper.Group.ToSum()
	other := rwxWrapper.Other.ToSum()

	allBytes := [3]byte{owner, group, other}

	return allBytes
}

func (rwxWrapper *RwxWrapper) ToUint32Octal() uint32 {
	// # https://play.golang.org/p/dX-wsvJmFie
	str := rwxWrapper.ToFileModeString()

	// # https://bit.ly/35aBepk
	octal, err := strconv.ParseUint(str, bitsize.Of8, bitsize.Of32)

	if err != nil {
		msgtype.
			MeaningFulErrorHandle(
				msgtype.PathChmodConvertFailedMessage,
				"ToUint32Octal",
				err)
	}

	return uint32(octal)
}

// ToCompiledOctalBytes4Digits return 0rwx, '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 0777, 0555, 0755 NOT 0rwx
func (rwxWrapper *RwxWrapper) ToCompiledOctalBytes4Digits() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := rwxWrapper.Owner.ToStringByte()
	group := rwxWrapper.Group.ToStringByte()
	other := rwxWrapper.Other.ToStringByte()

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
func (rwxWrapper *RwxWrapper) ToCompiledOctalBytes3Digits() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := rwxWrapper.Owner.ToStringByte()
	group := rwxWrapper.Group.ToStringByte()
	other := rwxWrapper.Other.ToStringByte()

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
func (rwxWrapper *RwxWrapper) ToCompiledSplitValues() (owner, group, other byte) {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner = rwxWrapper.Owner.ToStringByte()
	group = rwxWrapper.Group.ToStringByte()
	other = rwxWrapper.Other.ToStringByte()

	return owner, group, other
}

// ToFileModeString 4 digit string 0rwx, example 0777
func (rwxWrapper *RwxWrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := rwxWrapper.ToCompiledOctalBytes4Digits()

	return string(allBytes[:])
}

// ToRwxCompiledStr 3 digit string, example 777
func (rwxWrapper *RwxWrapper) ToRwxCompiledStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := rwxWrapper.ToCompiledOctalBytes4Digits()

	return string(allBytes[1:])
}

// ToFullRwxValueString returns "-rwxrwxrwx"
func (rwxWrapper *RwxWrapper) ToFullRwxValueString() string {
	owner := rwxWrapper.Owner.ToRwxString()
	group := rwxWrapper.Group.ToRwxString()
	other := rwxWrapper.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return constants.Hyphen + owner + group + other
}

// ToFullRwxValuesChars "-rwxrwxrwx" Bytes values
func (rwxWrapper *RwxWrapper) ToFullRwxValuesChars() []byte {
	str := rwxWrapper.ToFullRwxValueString()
	chars := []byte(str)

	return chars
}

func (rwxWrapper *RwxWrapper) String() string {
	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return rwxWrapper.ToFullRwxValueString()
}

func (rwxWrapper *RwxWrapper) ToFileMode() os.FileMode {
	// # https://play.golang.org/p/dX-wsvJmFie
	octalUint32 := rwxWrapper.ToUint32Octal()

	return os.FileMode(octalUint32)
}

func (rwxWrapper *RwxWrapper) ApplyChmod(
	isSkipOnNonExist bool,
	fileOrDirectoryPath string,
) error {
	isFileExist := fsinternal.IsPathExists(fileOrDirectoryPath)

	if isSkipOnNonExist && !isFileExist {
		return nil
	}

	if !isSkipOnNonExist && !isFileExist {
		return msgtype.
			PathInvalidErrorMessage.
			Error(
				messages.PathNotExist, fileOrDirectoryPath)
	}

	err := os.Chmod(fileOrDirectoryPath, rwxWrapper.ToFileMode())

	if err != nil {
		return msgtype.
			PathChmodApplyMessage.
			Error(err.Error(), fileOrDirectoryPath)
	}

	return nil
}

// LinuxApplyRecursive skip if it is a non dir path
func (rwxWrapper *RwxWrapper) LinuxApplyRecursive(
	isSkipOnNonExist bool,
	location string,
) error {
	isFileExist := fsinternal.IsPathExists(location)

	if isSkipOnNonExist && !isFileExist {
		return nil
	}

	if !isSkipOnNonExist && !isFileExist {
		return msgtype.
			PathInvalidErrorMessage.
			Error(
				"Path doesn't exist. path : ", location)
	}

	return rwxWrapper.applyLinuxRecursiveChmodUsingCmd(
		location)
}

func (rwxWrapper *RwxWrapper) applyLinuxRecursiveChmodUsingCmd(location string) error {
	cmd := rwxWrapper.getLinuxRecursiveCmdForChmod(location)

	if cmd == nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.BashCommandline,
			location)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.ChmodCommand,
			err.Error()+constants.NewLineUnix+stderr.String()+"location:"+location)
	}

	return nil
}

func (rwxWrapper *RwxWrapper) getLinuxRecursiveCmdForChmod(dirPath string) *exec.Cmd {
	instructionLine := constants.ChmodCommand +
		constants.Space +
		constants.RecursiveCommandFlag +
		constants.Space +
		rwxWrapper.ToRwxCompiledStr() +
		constants.Space +
		dirPath

	return exec.Command(
		constants.BinShellCmd,
		constants.NonInteractiveFlag,
		instructionLine)
}

func (rwxWrapper *RwxWrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(
		fileOrDirectoryPath,
		rwxWrapper.ToFileMode())

	if err != nil {
		finalErr := errors.New(err.Error() + fileOrDirectoryPath)

		panic(msgtype.MeaningFulError(
			msgtype.PathChmodApplyMessage,
			"MustApplyChmod",
			finalErr))
	}
}

func (rwxWrapper *RwxWrapper) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	return &chmodins.RwxOwnerGroupOther{
		Owner: rwxWrapper.Owner.ToRwxString(),
		Group: rwxWrapper.Group.ToRwxString(),
		Other: rwxWrapper.Other.ToRwxString(),
	}
}

func (rwxWrapper *RwxWrapper) ToRwxInstruction(
	condition *chmodins.Condition,
) *chmodins.RwxInstruction {
	rwxOwnerGroupOther := rwxWrapper.ToRwxOwnerGroupOther()

	return &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition:          *condition,
	}
}

func (rwxWrapper *RwxWrapper) Clone() *RwxWrapper {
	if rwxWrapper == nil {
		return nil
	}

	return &RwxWrapper{
		Owner: *rwxWrapper.Owner.Clone(),
		Group: *rwxWrapper.Group.Clone(),
		Other: *rwxWrapper.Other.Clone(),
	}
}

func (rwxWrapper *RwxWrapper) applyLinuxChmodOnManyNonRecursive(
	condition *chmodins.Condition,
	locations []string,
) error {
	if condition.IsContinueOnError {
		// continue on error
		return rwxWrapper.applyLinuxChmodNonRecursiveManyContinueOnError(
			condition,
			locations)
	}

	for _, location := range locations {
		err := rwxWrapper.ApplyChmod(
			condition.IsSkipOnNonExist,
			location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (rwxWrapper *RwxWrapper) ApplyLinuxChmodOnMany(
	condition *chmodins.Condition,
	locations ...string,
) error {
	if condition.IsRecursive {
		return rwxWrapper.applyLinuxChmodOnManyRecursive(
			condition,
			locations)
	}

	return rwxWrapper.applyLinuxChmodOnManyNonRecursive(
		condition, locations)
}

func (rwxWrapper *RwxWrapper) applyLinuxChmodOnManyRecursive(
	condition *chmodins.Condition,
	locations []string,
) error {
	if condition.IsContinueOnError {
		// continue on error
		return rwxWrapper.applyLinuxChmodRecursiveManyContinueOnError(
			condition,
			locations)
	}

	for _, location := range locations {
		err := rwxWrapper.LinuxApplyRecursive(
			condition.IsSkipOnNonExist,
			location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (rwxWrapper *RwxWrapper) applyLinuxChmodRecursiveManyContinueOnError(
	condition *chmodins.Condition,
	locations []string,
) error {
	var errSlice []string

	for _, location := range locations {
		err := rwxWrapper.LinuxApplyRecursive(
			condition.IsSkipOnNonExist,
			location)

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	return msgtype.SliceToErrorPtr(&errSlice)
}

func (rwxWrapper *RwxWrapper) applyLinuxChmodNonRecursiveManyContinueOnError(
	condition *chmodins.Condition,
	locations []string,
) error {
	var errSlice []string

	for _, location := range locations {
		err := rwxWrapper.ApplyChmod(
			condition.IsSkipOnNonExist,
			location)

		if err != nil {
			errSlice = append(errSlice, err.Error())
		}
	}

	return msgtype.SliceToErrorPtr(&errSlice)
}
