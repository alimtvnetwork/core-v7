package chmodhelper

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper/chmodclasstype"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
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

func (receiver *SingleRwx) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	switch receiver.ClassType {
	case chmodclasstype.All:
		return &chmodins.RwxOwnerGroupOther{
			Owner: receiver.Rwx,
			Group: receiver.Rwx,
			Other: receiver.Rwx,
		}
	case chmodclasstype.Owner:
		return &chmodins.RwxOwnerGroupOther{
			Owner: receiver.Rwx,
			Group: AllWildcards,
			Other: AllWildcards,
		}
	case chmodclasstype.Group:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: receiver.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.Other:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: AllWildcards,
			Other: receiver.Rwx,
		}

	case chmodclasstype.OwnerGroup:
		return &chmodins.RwxOwnerGroupOther{
			Owner: receiver.Rwx,
			Group: receiver.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.GroupOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: receiver.Rwx,
			Other: receiver.Rwx,
		}

	case chmodclasstype.OwnerOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: receiver.Rwx,
			Group: AllWildcards,
			Other: receiver.Rwx,
		}

	default:
		panic(chmodclasstype.BasicEnumImpl.RangesInvalidErr())
	}
}

func (receiver *SingleRwx) ToRwxInstruction(
	conditionalIns *chmodins.Condition,
) *chmodins.RwxInstruction {
	rwxOwnerGroupOther := receiver.ToRwxOwnerGroupOther()

	return &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition:          *conditionalIns,
	}
}

func (receiver *SingleRwx) ToVarRwxWrapper() (*RwxVariableWrapper, error) {
	rwxOwnerGroupOther := receiver.ToRwxOwnerGroupOther()

	return ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwxOwnerGroupOther)
}

func (receiver *SingleRwx) ToDisabledRwxWrapper() (*RwxWrapper, error) {
	rwxOwnerGroupOther := receiver.ToRwxOwnerGroupOther()
	rwxFullString := rwxOwnerGroupOther.String()
	rwxFullString = strings.ReplaceAll(
		rwxFullString,
		constants.WildcardSymbol,
		constants.Hyphen)

	rwxWrapper, err := NewUsingHyphenedRwxFullString(rwxFullString)

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (receiver *SingleRwx) ToRwxWrapper() (*RwxWrapper, error) {
	if !receiver.ClassType.IsAll() {
		return nil, msgtype.MeaningFulError(msgtype.CannotConvertToRwxWhereVarRwxPossible,
			"ToRwxWrapper", errors.New("use ToVarRwx"))
	}

	rwxWrapper, err := NewUsingRwxOwnerGroupOther(
		receiver.ToRwxOwnerGroupOther())

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (receiver *SingleRwx) ApplyOnMany(
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	toRwxInstruction := receiver.ToRwxInstruction(condition)
	executor, err := ParseRwxInstructionToExecutor(toRwxInstruction)

	if err != nil {
		return err
	}

	return executor.ApplyOnPathsPtr(&locations)
}
