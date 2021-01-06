package filemode

import (
	"os"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

type Wrapper struct {
	Owner, Group, Other Attribute
}

// return rwx, (Owner)(Group)(Other) byte values under 1-7
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
	octal, err := strconv.ParseUint(str, 8, 32)

	if err != nil {
		panic(err)
	}

	return uint32(octal)
}

// return 0rwx, '0'(Owner + '0')(Group + '0')(Other + '0')
// eg. 0777, 0555, 0755 NOT 0rwx
func (wrapper Wrapper) Chars() [4]byte {
	// # https://play.golang.org/p/dX-wsvJmFie
	owner := wrapper.Owner.ToChar()
	group := wrapper.Group.ToChar()
	other := wrapper.Other.ToChar()

	allBytes := [4]byte{constants.ZeroChar, owner, group, other}

	return allBytes
}

// 4 digit string 0rwx, example 0777
func (wrapper Wrapper) ToFileModeString() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[:])
}

// 3 digit string, example 777
func (wrapper Wrapper) ToModeStr() string {
	// # https://play.golang.org/p/dX-wsvJmFie
	allBytes := wrapper.Chars()

	return string(allBytes[1:])
}

// returns "-rwxrwxrwx"
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

func (wrapper Wrapper) ApplyChmod(fileOrDirectoryPath string) error {
	return os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())
}

func (wrapper Wrapper) MustApplyChmod(fileOrDirectoryPath string) {
	err := os.Chmod(fileOrDirectoryPath, wrapper.ToFileMode())

	if err != nil {
		panic(err)
	}
}
