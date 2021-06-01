package chmodhelper

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/msgtype"
)

type Wrapper struct {
	Owner, Group, Other Attribute
}

// Bytes return rwx, (Owner)(Group)(Other) byte values under 1-7
func (wrapper Wrapper) Bytes() [3]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToSum()
	group := wrapper.Group.ToSum()
	other := wrapper.Other.ToSum()

	allBytes := [3]byte{owner, group, other}

	return allBytes
}

func (wrapper Wrapper) ToUint32Octal() uint32 {
	// # https://play.golang.org/p/dX-wsvJmFie
	str := wrapper.ToFileModeString()

	// # https://bit.ly/35aBepk
	octal, err := strconv.ParseUint(str, bitsize.Of8, bitsize.Of32)

	if err != nil {
		msgtype.
			MeaningFulErrorHandle(
				msgtype.FileChmodConvertFailedMessage,
				"ToUint32Octal",
				err)
	}

	return uint32(octal)
}

// Chars return 0rwx, '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 0777, 0555, 0755 NOT 0rwx
func (wrapper Wrapper) Chars() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToChar()
	group := wrapper.Group.ToChar()
	other := wrapper.Other.ToChar()

	allBytes := [4]byte{constants.ZeroChar, owner, group, other}

	return allBytes
}

// ToFileModeString 4 digit string 0rwx, example 0777
func (wrapper Wrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[:])
}

// ToModeStr 3 digit string, example 777
func (wrapper Wrapper) ToModeStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[1:])
}

// ToRwxes returns "-rwxrwxrwx"
func (wrapper Wrapper) ToRwxes() string {
	owner := wrapper.Owner.ToRwxString()
	group := wrapper.Group.ToRwxString()
	other := wrapper.Other.ToRwxString()

	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return constants.Hyphen + owner + group + other
}

func (wrapper Wrapper) ToRwxesChars() []byte {
	str := wrapper.ToRwxes()
	chars := []byte(str)

	return chars
}

func (wrapper Wrapper) String() string {
	// # https://ss64.com/bash/chmod.html, needs to be 10 always
	return wrapper.ToRwxes()
}

func (wrapper Wrapper) ToFileMode() os.FileMode {
	// # https://play.golang.org/p/dX-wsvJmFie
	octalUint32 := wrapper.ToUint32Octal()

	return os.FileMode(octalUint32)
}

func (wrapper Wrapper) ApplyChmod(
	fileOrDirectoryPath string,
	isSkipOnNonExist bool,
) error {
	if isSkipOnNonExist && !isPathExist(fileOrDirectoryPath) {
		return nil
	}

	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		return msgtype.
			FileChmodApplyMessage.
			Error(err.Error(), fileOrDirectoryPath)
	}

	return nil
}

// UnixApplyRecursive skip if it is a non dir path
func (wrapper Wrapper) UnixApplyRecursive(
	dirPath string,
	isSkipOnNonExist bool,
) error {
	if isSkipOnNonExist && !isPathExist(dirPath) {
		return nil
	}

	fileMode := wrapper.ToFileMode()

	if fileMode.IsDir() {
		return wrapper.applyRecursiveChmodUsingCmd(
			dirPath)
	}

	return nil
}

func (wrapper Wrapper) applyRecursiveChmodUsingCmd(dirPath string) error {
	cmd := wrapper.getRecursiveCmdForChmod()

	if cmd == nil {
		return msgtype.
			FailedToCreateCmd.Error(
			constants.ChmodCommand,
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

func (wrapper Wrapper) getRecursiveCmdForChmod() *exec.Cmd {
	return exec.Command(
		constants.ChmodCommand,
		constants.RecursiveCommandFlag,
		wrapper.ToModeStr())
}

func (wrapper Wrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		panic(err)
	}
}
