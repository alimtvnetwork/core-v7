package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

type RwxVariableWrapper struct {
	rawInput            string
	isFixedType         bool
	Owner, Group, Other VarAttribute
}

// NewRwxVariableWrapper
//
//	partialRwx can be any length in
//	between 1-10 (rest will be fixed by wildcard)
//
//	Hyphened prefix MUST.
//
// rwxPartial:
//   - "-rwx"    will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
//
// Restrictions:
//   - cannot have first char other than hyphen(-) or else things will not work
func NewRwxVariableWrapper(partialRwx string) (*RwxVariableWrapper, error) {
	fullRwxWithWildcard := chmodins.FixRwxFullStringWithWildcards(partialRwx)

	owner := fullRwxWithWildcard[1:4]
	group := fullRwxWithWildcard[4:7]
	other := fullRwxWithWildcard[7:10]
	ownerAttr, err := ParseRwxToVarAttribute(owner)
	groupAttr, err2 := ParseRwxToVarAttribute(group)
	otherAttr, err3 := ParseRwxToVarAttribute(other)

	mergedErr := errcore.MergeErrors(
		err,
		err2,
		err3)

	if mergedErr != nil {
		return nil, mergedErr
	}

	isAllFixedType := ownerAttr.IsFixedType() &&
		groupAttr.IsFixedType() &&
		otherAttr.IsFixedType()

	return &RwxVariableWrapper{
		rawInput:    fullRwxWithWildcard,
		isFixedType: isAllFixedType,
		Owner:       *ownerAttr,
		Group:       *groupAttr,
		Other:       *otherAttr,
	}, nil
}

func (varWrapper *RwxVariableWrapper) IsFixedType() bool {
	return varWrapper.isFixedType
}

func (varWrapper *RwxVariableWrapper) HasWildcard() bool {
	return !varWrapper.isFixedType
}

func (varWrapper *RwxVariableWrapper) ToCompileFixedPtr() *RwxWrapper {
	if varWrapper.IsFixedType() {
		return varWrapper.ToCompileWrapperPtr(nil)
	}

	return nil
}

// ToCompileWrapper if Fixed type then fixed input can be nil.
func (varWrapper *RwxVariableWrapper) ToCompileWrapper(fixed *RwxWrapper) RwxWrapper {
	return *varWrapper.ToCompileWrapperPtr(fixed)
}

func (varWrapper *RwxVariableWrapper) ToCompileWrapperUsingLocationPtr(location string) (*RwxWrapper, error) {
	if varWrapper.IsFixedType() {
		return varWrapper.ToCompileFixedPtr(), nil
	}

	existingRwxWrapper, err := GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		return nil, err
	}

	return varWrapper.ToCompileWrapperPtr(existingRwxWrapper), nil
}

// ToCompileWrapperPtr if Fixed type then fixed input can be nil.
func (varWrapper *RwxVariableWrapper) ToCompileWrapperPtr(fixed *RwxWrapper) *RwxWrapper {
	if varWrapper.IsFixedType() {
		return &RwxWrapper{
			Owner: *varWrapper.Owner.ToCompileFixAttr(),
			Group: *varWrapper.Group.ToCompileFixAttr(),
			Other: *varWrapper.Other.ToCompileFixAttr(),
		}
	}

	return &RwxWrapper{
		Owner: varWrapper.Owner.ToCompileAttr(&fixed.Owner),
		Group: varWrapper.Group.ToCompileAttr(&fixed.Group),
		Other: varWrapper.Other.ToCompileAttr(&fixed.Other),
	}
}

func (varWrapper *RwxVariableWrapper) Clone() *RwxVariableWrapper {
	if varWrapper == nil {
		return nil
	}

	return &RwxVariableWrapper{
		rawInput:    varWrapper.rawInput,
		isFixedType: varWrapper.IsFixedType(),
		Owner:       *varWrapper.Owner.Clone(),
		Group:       *varWrapper.Group.Clone(),
		Other:       *varWrapper.Other.Clone(),
	}
}

func (varWrapper *RwxVariableWrapper) IsEqualPtr(next *RwxVariableWrapper) bool {
	if varWrapper == nil && next == nil {
		return true
	}

	if varWrapper == nil || next == nil {
		return false
	}

	isOwner := varWrapper.Owner.IsEqualPtr(&next.Owner)
	isGroup := varWrapper.Group.IsEqualPtr(&next.Group)
	isOther := varWrapper.Other.IsEqualPtr(&next.Other)

	return isOwner &&
		isGroup &&
		isOther
}

func (varWrapper *RwxVariableWrapper) IsOwnerPartialMatch(rwx string) bool {
	return IsPartialMatchVariableAttr(
		&varWrapper.Owner,
		rwx)
}

func (varWrapper *RwxVariableWrapper) IsGroupPartialMatch(rwx string) bool {
	return IsPartialMatchVariableAttr(
		&varWrapper.Group,
		rwx)
}

func (varWrapper *RwxVariableWrapper) IsOtherPartialMatch(rwx string) bool {
	return IsPartialMatchVariableAttr(
		&varWrapper.Other,
		rwx)
}

func (varWrapper *RwxVariableWrapper) ApplyRwxOnLocations(
	isContinueOnError,
	isSkipOnInvalid bool,
	locations ...string,
) error {
	existsFilteredPathFileInfoMap := GetExistsFilteredPathFileInfoMap(
		isSkipOnInvalid,
		locations...)
	if !isContinueOnError && existsFilteredPathFileInfoMap.Error != nil {
		return existsFilteredPathFileInfoMap.Error
	}

	locationsFileInfoRwx := existsFilteredPathFileInfoMap.
		LazyValidLocationFileInfoRwxWrappers()

	if isContinueOnError {
		var sliceErr []string
		for _, locationFileInfoRwx := range locationsFileInfoRwx {
			rwx := locationFileInfoRwx.
				RwxWrapper

			if rwx == nil {
				continue
			}

			err := rwx.
				ApplyChmod(
					isSkipOnInvalid,
					locationFileInfoRwx.Location)

			if err != nil {
				sliceErr = append(sliceErr, err.Error())
			}
		}

		return errcore.SliceToError(sliceErr)
	}

	for _, locationFileInfoRwx := range locationsFileInfoRwx {
		rwx := locationFileInfoRwx.
			RwxWrapper

		if rwx == nil {
			continue
		}

		err := rwx.
			ApplyChmod(
				isSkipOnInvalid,
				locationFileInfoRwx.Location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (varWrapper *RwxVariableWrapper) RwxMatchingStatus(
	isContinueOnError,
	isSkipOnInvalid bool,
	locations []string,
) *RwxMatchingStatus {
	existsFilteredPathFileInfoMap := GetExistsFilteredPathFileInfoMap(
		isSkipOnInvalid,
		locations...)
	if !isContinueOnError && existsFilteredPathFileInfoMap.Error != nil {
		return InvalidRwxMatchingStatus(existsFilteredPathFileInfoMap.Error)
	}

	rwxMismatchInfos := make(
		[]*RwxMismatchInfo,
		0,
		constants.Capacity1)

	for filePath, fileInfo := range existsFilteredPathFileInfoMap.FilesToInfoMap {
		fileRwx := fileInfo.Mode().String()

		if varWrapper.IsMismatchPartialFullRwx(fileRwx) {
			rwxMismatchInfos = append(rwxMismatchInfos,
				&RwxMismatchInfo{
					FilePath:  filePath,
					Expecting: varWrapper.ToString(false),
					Actual:    fileRwx[1:],
				})
		}
	}

	isAllMatching := len(rwxMismatchInfos) == 0 &&
		len(locations) == len(existsFilteredPathFileInfoMap.FilesToInfoMap)

	return &RwxMatchingStatus{
		RwxMismatchInfos:         rwxMismatchInfos,
		MissingOrPathsWithIssues: existsFilteredPathFileInfoMap.MissingOrOtherPathIssues,
		IsAllMatching:            isAllMatching,
		Error:                    existsFilteredPathFileInfoMap.Error,
	}
}

// IsEqualPartialRwxPartial
//
//	will make the partial to full rwx and then calls IsEqualPartialFullRwx
//	partialRwx can be any length in
//	between 1-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func (varWrapper *RwxVariableWrapper) IsEqualPartialRwxPartial(
	rwxPartial string,
) bool {
	fullRwxWithWildcard := chmodins.FixRwxFullStringWithWildcards(rwxPartial)

	return varWrapper.IsEqualPartialFullRwx(fullRwxWithWildcard)
}

// IsEqualUsingLocation returns by sending to IsEqualPartialFullRwx
//
// Returns false on non exist.
func (varWrapper *RwxVariableWrapper) IsEqualUsingLocation(
	location string,
) bool {
	fileInfo, _ := os.Stat(location)

	if fileInfo == nil {
		return false
	}

	return varWrapper.IsEqualPartialFullRwx(
		fileInfo.Mode().String())
}

// IsEqualUsingFileInfo returns by sending to IsEqualPartialFullRwx
//
// Returns false on nil.
func (varWrapper *RwxVariableWrapper) IsEqualUsingFileInfo(
	fileInfo os.FileInfo,
) bool {
	if fileInfo == nil {
		return false
	}

	return varWrapper.IsEqualPartialFullRwx(
		fileInfo.Mode().String())
}

// IsEqualUsingFileMode returns by sending to IsEqualPartialFullRwx
//
// Returns false on nil.
func (varWrapper *RwxVariableWrapper) IsEqualUsingFileMode(
	fileMode os.FileMode,
) bool {
	return varWrapper.IsEqualPartialFullRwx(
		fileMode.String())
}

// IsEqualRwxWrapperPtr returns by sending to IsEqualPartialFullRwx
//
// Returns false on nil.
func (varWrapper *RwxVariableWrapper) IsEqualRwxWrapperPtr(
	rwxWrapper *RwxWrapper,
) bool {
	if rwxWrapper == nil {
		return false
	}

	return varWrapper.IsEqualPartialFullRwx(
		rwxWrapper.ToFullRwxValueString())
}

// IsMismatchPartialFullRwx returns revert of IsEqualPartialFullRwx
//
// fullRwx (10 chars) where wildcard will be ignore during compare
func (varWrapper *RwxVariableWrapper) IsMismatchPartialFullRwx(
	fullRwx string,
) bool {
	return !varWrapper.IsEqualPartialFullRwx(fullRwx)
}

// IsEqualPartialFullRwx will compare with concrete FullRwx (10 chars) where wildcard will be ignore during compare
func (varWrapper *RwxVariableWrapper) IsEqualPartialFullRwx(
	fullRwx string,
) bool {
	if len(fullRwx) < HyphenedRwxLength {
		return false
	}

	owner := fullRwx[1:4]
	group := fullRwx[4:7]
	other := fullRwx[7:10]

	isOwner := varWrapper.IsOwnerPartialMatch(owner)
	isGroup := varWrapper.IsGroupPartialMatch(group)
	isOther := varWrapper.IsOtherPartialMatch(other)

	return isOwner &&
		isGroup &&
		isOther
}

// IsEqualPartialUsingFileMode will compare
// with concrete FullRwx (10 chars) where wildcard will be ignore during compare
func (varWrapper *RwxVariableWrapper) IsEqualPartialUsingFileMode(
	mode os.FileMode,
) bool {
	return varWrapper.IsEqualPartialFullRwx(
		mode.String())
}

func (varWrapper *RwxVariableWrapper) ToString(isIncludeHyphen bool) string {
	if isIncludeHyphen {
		return constants.Hyphen +
			varWrapper.Owner.String() +
			varWrapper.Group.String() +
			varWrapper.Other.String()
	}

	return varWrapper.Owner.String() +
		varWrapper.Group.String() +
		varWrapper.Other.String()
}

func (varWrapper *RwxVariableWrapper) String() string {
	return varWrapper.rawInput
}
