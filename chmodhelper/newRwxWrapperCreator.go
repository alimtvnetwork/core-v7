package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/messages"
)

type newRwxWrapperCreator struct{}

// CreatePtr
//
//	mode length needs to 3, not more not less
//	mode chars should be digits only (0-7)
//	example "777", "755", "655"
func (it newRwxWrapperCreator) CreatePtr(mode string) (*RwxWrapper, error) {
	rwx, err := it.Create(mode)

	if err != nil {
		return nil, err
	}

	return &rwx, err
}

// Create
//
//	mode length needs to 3, not more not less
//	mode chars should be digits only (0-7)
//	example "777", "755", "655"
func (it newRwxWrapperCreator) Create(mode string) (RwxWrapper, error) {
	length := len(mode)

	if length != SingleRwxLength {
		panic(errcore.OutOfRangeLengthType.Combine(
			"mode length should be "+SingleRwxLengthString,
			length))
	}

	allBytes := []byte(mode)

	for i, allByte := range allBytes {
		n := allByte - constants.ZeroChar

		if n > 7 || n < 0 {
			err := errcore.
				InvalidCharType.
				Error(
					messages.ModeCharShouldBeAllNumbersAndWithin0To7,
					n+constants.ZeroChar)

			return RwxWrapper{}, err
		}

		allBytes[i] = n
	}

	return it.UsingSpecificByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex]), nil
}

// UsingBytes
//
//	each byte should not be more than 7
func (it newRwxWrapperCreator) UsingBytes(allBytes [3]byte) RwxWrapper {
	return it.UsingSpecificByte(
		allBytes[OwnerIndex],
		allBytes[GroupIndex],
		allBytes[OtherIndex])
}

func (it newRwxWrapperCreator) Invalid() RwxWrapper {
	return RwxWrapper{}
}

func (it newRwxWrapperCreator) InvalidPtr() *RwxWrapper {
	return &RwxWrapper{}
}

func (it newRwxWrapperCreator) UsingChmod(
	fileMode os.FileMode,
) *RwxWrapper {
	return it.UsingChmod(fileMode)
}

// UsingFileModePtr
//
//	Hint. os.FileMode.String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Reference:
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) UsingFileModePtr(
	fileMode os.FileMode,
) *RwxWrapper {
	if fileMode == 0 {
		return it.Empty()
	}

	str := fileMode.String()
	// Reference : https://play.golang.org/p/Qq_rKl_pAqe
	owner := str[1:4]
	group := str[4:7]
	other := str[7:10]

	return &RwxWrapper{
		Owner: New.Attribute.UsingRwxString(owner),
		Group: New.Attribute.UsingRwxString(group),
		Other: New.Attribute.UsingRwxString(other),
	}
}

// UsingFileMode
//
//	Hint. os.FileMode String() returns "-rwxrwxrwx" full rwx
//	https://go.dev/play/p/Qq_rKl_pAqe
//
// Reference:
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) UsingFileMode(
	fileMode os.FileMode,
) RwxWrapper {
	if fileMode == 0 {
		return it.Empty().ToNonPtr()
	}

	str := fileMode.String()
	// Reference : https://play.golang.org/p/Qq_rKl_pAqe
	owner := str[1:4]
	group := str[4:7]
	other := str[7:10]

	return RwxWrapper{
		Owner: New.Attribute.UsingRwxString(owner),
		Group: New.Attribute.UsingRwxString(group),
		Other: New.Attribute.UsingRwxString(other),
	}
}

func (it newRwxWrapperCreator) UsingRwxOwnerGroupOther(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (RwxWrapper, error) {
	return it.RwxFullStringWtHyphen(
		rwxOwnerGroupOther.ToString(
			false))
}

// UsingSpecificByte
//
//	each byte should not be more than 7
func (it newRwxWrapperCreator) UsingSpecificByte(
	owner, group, other byte,
) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: New.Attribute.UsingByteMust(owner),
		Group: New.Attribute.UsingByteMust(group),
		Other: New.Attribute.UsingByteMust(other),
	}

	return wrapper
}

func (it newRwxWrapperCreator) UsingAttrVariants(
	owner, group, other AttrVariant,
) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: New.Attribute.UsingVariantMust(owner),
		Group: New.Attribute.UsingVariantMust(group),
		Other: New.Attribute.UsingVariantMust(other),
	}

	return wrapper
}

func (it newRwxWrapperCreator) UsingAttrs(
	owner, group, other Attribute,
) RwxWrapper {
	wrapper := RwxWrapper{
		Owner: owner,
		Group: group,
		Other: other,
	}

	return wrapper
}

// Rwx10
//
//	alias for RwxFullString
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
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) Rwx10(
	hyphenedRwxRwxRwx string,
) (RwxWrapper, error) {
	return it.RwxFullString(hyphenedRwxRwxRwx)
}

// Rwx9
//
// alias of RwxFullStringWtHyphen
//
// Format (length must be 9)
//
//	"rwxrwxrwx"
//
// Example:
//   - owner all enabled only "rwx------"
//   - group all enabled only "---rwx---"
//
// Must have or restrictions:
//   - string length must be 9.
//
// Reference:
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) Rwx9(
	rwxRwxRwx string,
) (RwxWrapper, error) {
	return it.RwxFullStringWtHyphen(
		rwxRwxRwx)
}

// RwxFullString
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
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) RwxFullString(
	hyphenedRwxRwxRwx string,
) (RwxWrapper, error) {
	length := len(hyphenedRwxRwxRwx)

	if length != HyphenedRwxLength {
		return it.Invalid(), errHyphenedRwxLength
	}

	return it.RwxFullStringWtHyphen(
		hyphenedRwxRwxRwx[constants.One:])
}

// RwxFullStringWtHyphen
//
// Format (length must be 9)
//
//	"rwxrwxrwx"
//
// Example:
//   - owner all enabled only "rwx------"
//   - group all enabled only "---rwx---"
//
// Must have or restrictions:
//   - string length must be 9.
//
// Reference:
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) RwxFullStringWtHyphen(
	rwxFullStringWithoutHyphen string,
) (RwxWrapper, error) {
	length := len(rwxFullStringWithoutHyphen)

	if length != FullRwxLengthWithoutHyphen {
		return RwxWrapper{}, errFullRwxLengthWithoutHyphen
	}

	owner := rwxFullStringWithoutHyphen[0:3]
	group := rwxFullStringWithoutHyphen[3:6]
	other := rwxFullStringWithoutHyphen[6:9]

	wrapper := RwxWrapper{
		Owner: New.Attribute.UsingRwxString(owner),
		Group: New.Attribute.UsingRwxString(group),
		Other: New.Attribute.UsingRwxString(other),
	}

	return wrapper, nil
}

func (it newRwxWrapperCreator) UsingVariant(variant Variant) (RwxWrapper, error) {
	return it.Create(variant.String())
}

func (it newRwxWrapperCreator) UsingVariantPtr(
	variant Variant,
) (*RwxWrapper, error) {
	rwxWrapper, err := it.Create(variant.String())

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, nil
}

// Instruction
//
//	rwxFullString must be 10 chars in "-rwxrwxrwx"
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
//   - Chmod examples      : https://ss64.com/bash/chmod.html
//   - FileMode to RwxFull : https://go.dev/play/p/Qq_rKl_pAqe
func (it newRwxWrapperCreator) Instruction(
	rwxFullString string,
	condition chmodins.Condition,
) (*chmodins.RwxInstruction, error) {
	rwxWrapper, err := it.RwxFullString(rwxFullString)

	if err != nil {
		return nil, err
	}

	return rwxWrapper.ToRwxInstruction(&condition), nil
}

func (it newRwxWrapperCreator) UsingExistingFile(
	filePath string,
) (*RwxWrapper, error) {
	existingChmod, err := GetExistingChmodRwxWrapper(filePath)

	return existingChmod.ToPtr(), err
}

// UsingExistingFileSkipInvalidFile
//
// Warning:
//
//	swallows the error and invalid file
func (it newRwxWrapperCreator) UsingExistingFileSkipInvalidFile(
	filePath string,
) (rwxWrapper *RwxWrapper, isInvalidFile bool) {
	existingChmod, isInvalidFile := GetExistingChmodOfValidFile(filePath)

	if !isInvalidFile {
		// valid
		return it.UsingFileModePtr(existingChmod), isInvalidFile
	}

	return it.Empty(), isInvalidFile
}

// UsingExistingFileOption
//
// Warning:
//
//	swallows the error and invalid file
func (it newRwxWrapperCreator) UsingExistingFileOption(
	isSkipInvalidFile bool,
	filePath string,
) (rwxWrapper *RwxWrapper, err error, isInvalidFile bool) {
	if isSkipInvalidFile {
		rwxWrapper, isInvalidFile = it.UsingExistingFileSkipInvalidFile(
			filePath)

		return rwxWrapper, nil, isInvalidFile
	}

	existingChmod, err := GetExistingChmod(filePath)

	if err != nil || existingChmod == 0 {
		return it.Empty(), err, true
	}

	return it.UsingFileModePtr(existingChmod),
		err,
		err != nil && existingChmod != 0
}

func (it newRwxWrapperCreator) Empty() *RwxWrapper {
	return &RwxWrapper{}
}
