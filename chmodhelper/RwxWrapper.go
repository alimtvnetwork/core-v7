package chmodhelper

import (
	"bytes"
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

func (wrapper RwxWrapper) Verify(location string) error {
	return VerifyChmod(location, wrapper.ToFullRwxValueString())
}

func (wrapper RwxWrapper) VerifyPaths(location *[]string, isContinueOnError bool) error {
	return VerifyChmodPaths(
		location,
		wrapper.ToFullRwxValueString(),
		isContinueOnError)
}

func (wrapper RwxWrapper) HasChmod(location string) bool {
	return IsChmod(location, wrapper.ToFullRwxValueString())
}

// Bytes return rwx, (Owner)(Group)(Other) byte values under 1-7
func (wrapper RwxWrapper) Bytes() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToSum()
	group := wrapper.Group.ToSum()
	other := wrapper.Other.ToSum()

	allBytes := [3]byte{owner, group, other}

	return allBytes
}

func (wrapper RwxWrapper) ToUint32Octal() uint32 {
	// # https://play.golang.org/p/dX-wsvJmFie
	str := wrapper.ToFileModeString()

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
func (wrapper RwxWrapper) ToCompiledOctalBytes4Digits() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToStringByte()
	group := wrapper.Group.ToStringByte()
	other := wrapper.Other.ToStringByte()

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
func (wrapper RwxWrapper) ToCompiledOctalBytes3Digits() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToStringByte()
	group := wrapper.Group.ToStringByte()
	other := wrapper.Other.ToStringByte()

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
func (wrapper RwxWrapper) ToCompiledSplitValues() (owner, group, other byte) {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner = wrapper.Owner.ToStringByte()
	group = wrapper.Group.ToStringByte()
	other = wrapper.Other.ToStringByte()

	return owner, group, other
}

// ToFileModeString 4 digit string 0rwx, example 0777
func (wrapper RwxWrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.ToCompiledOctalBytes4Digits()

	return string(allBytes[:])
}

// ToRwxCompiledStr 3 digit string, example 777
func (wrapper RwxWrapper) ToRwxCompiledStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.ToCompiledOctalBytes4Digits()

	return string(allBytes[1:])
}

// ToFullRwxValueString returns "-rwxrwxrwx"
func (wrapper RwxWrapper) ToFullRwxValueString() string {
	owner := wrapper.Owner.ToRwxString()
	group := wrapper.Group.ToRwxString()
	other := wrapper.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return constants.Hyphen + owner + group + other
}

// ToFullRwxValuesChars "-rwxrwxrwx" Bytes values
func (wrapper RwxWrapper) ToFullRwxValuesChars() []byte {
	str := wrapper.ToFullRwxValueString()
	chars := []byte(str)

	return chars
}

func (wrapper RwxWrapper) String() string {
	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return wrapper.ToFullRwxValueString()
}

func (wrapper RwxWrapper) ToFileMode() os.FileMode {
	// # https://play.golang.org/p/dX-wsvJmFie
	octalUint32 := wrapper.ToUint32Octal()

	return os.FileMode(octalUint32)
}

func (wrapper RwxWrapper) ApplyChmod(
	fileOrDirectoryPath string,
	isSkipOnNonExist bool,
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

	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		return msgtype.
			PathChmodApplyMessage.
			Error(err.Error(), fileOrDirectoryPath)
	}

	return nil
}

// UnixApplyRecursive skip if it is a non dir path
func (wrapper RwxWrapper) UnixApplyRecursive(
	dirPath string,
	isSkipOnNonExist bool,
) error {
	isFileExist := fsinternal.IsPathExists(dirPath)

	if isSkipOnNonExist && !isFileExist {
		return nil
	}

	if !isSkipOnNonExist && !isFileExist {
		return msgtype.
			PathInvalidErrorMessage.
			Error(
				"Path doesn't exist", dirPath)
	}

	return wrapper.applyUnixRecursiveChmodUsingCmd(
		dirPath)
}

func (wrapper RwxWrapper) applyUnixRecursiveChmodUsingCmd(dirPath string) error {
	cmd := wrapper.getUnixRecursiveCmdForChmod(dirPath)

	if cmd == nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.BashCommandline,
			dirPath)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.ChmodCommand,
			err.Error()+constants.NewLineUnix+stderr.String())
	}

	return nil
}

func (wrapper RwxWrapper) getUnixRecursiveCmdForChmod(dirPath string) *exec.Cmd {
	instructionLine := constants.ChmodCommand +
		constants.Space +
		constants.RecursiveCommandFlag +
		constants.Space +
		wrapper.ToRwxCompiledStr() +
		constants.Space +
		dirPath

	return exec.Command(
		constants.BinShellCmd,
		constants.NonInteractiveFlag,
		instructionLine)
}

func (wrapper RwxWrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		panic(err)
	}
}

func (wrapper RwxWrapper) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	return &chmodins.RwxOwnerGroupOther{
		Owner: wrapper.Owner.ToRwxString(),
		Group: wrapper.Group.ToRwxString(),
		Other: wrapper.Other.ToRwxString(),
	}
}
