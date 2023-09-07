package chmodhelper

import (
	"errors"
	"strings"

	"gitlab.com/auk-go/core/chmodhelper/chmodclasstype"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

type SingleRwx struct {
	// Rwx Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Rwx       string
	ClassType chmodclasstype.Variant
}

func NewSingleRwx(
	rwx string,
	classType chmodclasstype.Variant,
) (*SingleRwx, error) {
	err := GetRwxLengthError(rwx)

	if err != nil {
		return nil, err
	}

	return &SingleRwx{
		Rwx:       rwx,
		ClassType: classType,
	}, nil
}

func (it *SingleRwx) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	switch it.ClassType {
	case chmodclasstype.All:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: it.Rwx,
			Other: it.Rwx,
		}
	case chmodclasstype.Owner:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: AllWildcards,
			Other: AllWildcards,
		}
	case chmodclasstype.Group:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: it.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.Other:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: AllWildcards,
			Other: it.Rwx,
		}

	case chmodclasstype.OwnerGroup:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: it.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.GroupOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: it.Rwx,
			Other: it.Rwx,
		}

	case chmodclasstype.OwnerOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: AllWildcards,
			Other: it.Rwx,
		}

	default:
		panic(chmodclasstype.BasicEnumImpl.RangesInvalidErr())
	}
}

func (it *SingleRwx) ToRwxInstruction(
	conditionalIns *chmodins.Condition,
) *chmodins.RwxInstruction {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()

	return &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition:          *conditionalIns,
	}
}

func (it *SingleRwx) ToVarRwxWrapper() (*RwxVariableWrapper, error) {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()

	return ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwxOwnerGroupOther)
}

func (it *SingleRwx) ToDisabledRwxWrapper() (*RwxWrapper, error) {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()
	rwxFullString := rwxOwnerGroupOther.String()
	rwxFullString = strings.ReplaceAll(
		rwxFullString,
		constants.WildcardSymbol,
		constants.Hyphen)

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxFullString)

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (it *SingleRwx) ToRwxWrapper() (*RwxWrapper, error) {
	if !it.ClassType.IsAll() {
		return nil, errcore.MeaningfulError(
			errcore.CannotConvertToRwxWhereVarRwxPossibleType,
			"ToRwxWrapper",
			errors.New("use ToVarRwx"))
	}

	rwxWrapper, err := New.RwxWrapper.UsingRwxOwnerGroupOther(
		it.ToRwxOwnerGroupOther())

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (it *SingleRwx) ApplyOnMany(
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	toRwxInstruction := it.ToRwxInstruction(condition)
	executor, err := ParseRwxInstructionToExecutor(toRwxInstruction)

	if err != nil {
		return err
	}

	return executor.ApplyOnPathsPtr(&locations)
}
