package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

// New mode length needs to 3, not more not less
// mode chars should be digits only (0-7)
// example "777", "755", "655"
func New(mode string) (RwxWrapper, error) {
	length := len(mode)

	if length != SingleRwxLength {
		panic(msgtype.OutOfRangeLength.Combine(
			"mode length should be "+SingleRwxLengthString,
			length))
	}

	allBytes := []byte(mode)

	for i, allByte := range allBytes {
		n := allByte - constants.ZeroChar

		if n > 7 || n < 0 {
			err := msgtype.
				InvalidCharErrorMessage.
				Error(
					messages.ModeCharShouldBeAllNumbersAndWithin0To7,
					n+constants.ZeroChar)

			return RwxWrapper{}, err
		}

		allBytes[i] = n
	}

	return NewUsingByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex]), nil
}

// NewUsingBytes each byte should not be more than 7
func NewUsingBytes(allBytes [3]byte) RwxWrapper {
	return NewUsingByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex])
}

func NewUsingFileModePtr(fileMode os.FileMode) *RwxWrapper {
	str := fileMode.String()
	// Reference : https://play.golang.org/p/Qq_rKl_pAqe
	owner := str[1:4]
	group := str[4:7]
	other := str[7:10]

	return &RwxWrapper{
		Owner: NewAttributeUsingRwx(owner),
		Group: NewAttributeUsingRwx(group),
		Other: NewAttributeUsingRwx(other),
	}
}

func NewUsingFileMode(fileMode os.FileMode) RwxWrapper {
	str := fileMode.String()
	// Reference : https://play.golang.org/p/Qq_rKl_pAqe
	owner := str[1:4]
	group := str[4:7]
	other := str[7:10]

	return RwxWrapper{
		Owner: NewAttributeUsingRwx(owner),
		Group: NewAttributeUsingRwx(group),
		Other: NewAttributeUsingRwx(other),
	}
}

func NewUsingRwxOwnerGroupOther(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (RwxWrapper, error) {
	return NewUsingRwxFullString(
		rwxOwnerGroupOther.ToString(
			false))
}

// NewUsingByte each byte should not be more than 7
func NewUsingByte(owner, group, other byte) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: NewAttributeUsingByte(owner),
		Group: NewAttributeUsingByte(group),
		Other: NewAttributeUsingByte(other),
	}

	return wrapper
}

func NewUsingAttrVariants(owner, group, other AttrVariant) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: NewAttributeUsingVariant(owner),
		Group: NewAttributeUsingVariant(group),
		Other: NewAttributeUsingVariant(other),
	}

	return wrapper
}

func NewUsingAttrs(owner, group, other Attribute) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: owner,
		Group: group,
		Other: other,
	}

	return wrapper
}

// NewUsingHyphenedRwxFullString Format "-rwxrwxrwx"
//
// eg. owener all enabled only "-rwx------"
//
// eg. group all enabled only  "----rwx---"
//
// length must be 10 always.
//
// Reference: https://ss64.com/bash/chmod.html
func NewUsingHyphenedRwxFullString(hyphenedRwxRwxRwx string) (RwxWrapper, error) {
	length := len(hyphenedRwxRwxRwx)

	if length != HyphenedRwxLength {
		return RwxWrapper{}, hyphenedRwxLengthErr
	}

	return NewUsingRwxFullString(hyphenedRwxRwxRwx[constants.One:])
}

// NewUsingRwxFullString Format "rwxrwxrwx"
//
// eg. owener all enabled only "rwx------"
//
// eg. group all enabled only  "---rwx---"
//
// length must be 9 always.
func NewUsingRwxFullString(rwxFullStringWithoutHyphen string) (RwxWrapper, error) {
	length := len(rwxFullStringWithoutHyphen)

	if length != FullRwxLengthWithoutHyphen {
		return RwxWrapper{}, fullRwxLengthWithoutHyphenErr
	}

	owner := rwxFullStringWithoutHyphen[0:3]
	group := rwxFullStringWithoutHyphen[3:6]
	other := rwxFullStringWithoutHyphen[6:9]

	wrapper := RwxWrapper{
		Owner: NewAttributeUsingRwx(owner),
		Group: NewAttributeUsingRwx(group),
		Other: NewAttributeUsingRwx(other),
	}

	return wrapper, nil
}

func NewUsingVariant(variant Variant) (RwxWrapper, error) {
	return New(variant.String())
}
