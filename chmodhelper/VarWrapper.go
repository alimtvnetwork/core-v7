package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodinstruction"
	"gitlab.com/evatix-go/core/constants"
)

type VarWrapper struct {
	rawInput            string
	isFixedType         bool
	Owner, Group, Other VarAttribute
}

func ParseRwxInstructionToStringRwx(
	rwxInstruction *chmodinstruction.RwxInstruction,
	isIncludeHyphen bool,
) string {
	if rwxInstruction == nil {
		return constants.EmptyString
	}

	var hyphen string

	if isIncludeHyphen {
		hyphen = constants.Hyphen
	}

	compiled := hyphen +
		rwxInstruction.Owner +
		rwxInstruction.Group +
		rwxInstruction.Other

	return compiled
}

func (varWrapper *VarWrapper) IsFixedType() bool {
	return varWrapper.isFixedType
}

// ToCompileWrapper if Fixed type then fixed input can be nil.
func (varWrapper *VarWrapper) ToCompileWrapper(fixed *Wrapper) Wrapper {
	return *varWrapper.ToCompileWrapperPtr(fixed)
}

// ToCompileWrapperPtr if Fixed type then fixed input can be nil.
func (varWrapper *VarWrapper) ToCompileWrapperPtr(fixed *Wrapper) *Wrapper {
	if varWrapper.IsFixedType() {
		return &Wrapper{
			Owner: *varWrapper.Owner.ToCompileFixAttr(),
			Group: *varWrapper.Group.ToCompileFixAttr(),
			Other: *varWrapper.Other.ToCompileFixAttr(),
		}
	}

	return &Wrapper{
		Owner: varWrapper.Owner.ToCompileAttr(&fixed.Owner),
		Group: varWrapper.Group.ToCompileAttr(&fixed.Group),
		Other: varWrapper.Other.ToCompileAttr(&fixed.Other),
	}
}
