package filemode

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

// mode length needs to 3, not more not less
// mode chars should be digits only (0-7)
// example "777", "755", "655"
func New(mode string) (Wrapper, error) {
	length := len(mode)

	if length != SupportedLength {
		panic(msgtype.OutOfRangeLength.Combine(
			"mode length should be "+SupportedLengthString,
			length))
	}

	allBytes := []byte(mode)

	for i, allByte := range allBytes {
		n := allByte - constants.ZeroChar

		if n > 7 || n < 0 {
			err := msgtype.
				InvalidCharErrorMessage.
				Error(
					"mode char should be all digits and under 0 to 7",
					n+constants.ZeroChar)

			return Wrapper{}, err
		}

		allBytes[i] = n
	}

	return NewUsingByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex]), nil
}

// each byte should not be more than 7
func NewUsingBytes(allBytes [3]byte) Wrapper {
	return NewUsingByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex])
}

func NewUsingFileMode(fileMode os.FileMode) Wrapper {
	str := fileMode.String()

	// Reference : https://play.golang.org/p/Qq_rKl_pAqe
	owner := str[1:4]
	group := str[4:7]
	other := str[7:10]

	wrapper := Wrapper{
		Owner: NewAttributeUsingRwx(owner),
		Group: NewAttributeUsingRwx(group),
		Other: NewAttributeUsingRwx(other),
	}

	return wrapper
}

// each byte should not be more than 7
func NewUsingByte(owner, group, other byte) Wrapper {
	wrapper := Wrapper{
		Owner: NewAttributeUsingByte(owner),
		Group: NewAttributeUsingByte(group),
		Other: NewAttributeUsingByte(other),
	}

	return wrapper
}

func NewUsingAttrVariants(owner, group, other AttrVariant) Wrapper {
	wrapper := Wrapper{
		Owner: NewAttributeUsingVariant(owner),
		Group: NewAttributeUsingVariant(group),
		Other: NewAttributeUsingVariant(other),
	}

	return wrapper
}

func NewUsingAttrs(owner, group, other Attribute) Wrapper {
	wrapper := Wrapper{
		Owner: owner,
		Group: group,
		Other: other,
	}

	return wrapper
}

// Format "rwxrwxrwx"
// eg. owener all enabled only "rwx------"
// length must be 9 always.
func NewUsingRwxes(rwxrwxrwx string) (Wrapper, error) {
	length := len(rwxrwxrwx)

	if length != 9 {
		err := msgtype.
			LengthShouldBeEqualToMessage.
			Error("rwxrwxrwx length must be 9", length)

		return Wrapper{}, err
	}

	owner := rwxrwxrwx[0:3]
	group := rwxrwxrwx[3:6]
	other := rwxrwxrwx[6:9]

	wrapper := Wrapper{
		Owner: NewAttributeUsingRwx(owner),
		Group: NewAttributeUsingRwx(group),
		Other: NewAttributeUsingRwx(other),
	}

	return wrapper, nil
}

func NewUsingVariant(variant Variant) (Wrapper, error) {
	return New(variant.String())
}
