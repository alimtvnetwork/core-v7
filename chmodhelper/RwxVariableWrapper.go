package chmodhelper

type RwxVariableWrapper struct {
	rawInput            string
	isFixedType         bool
	Owner, Group, Other VarAttribute
}

func (varWrapper *RwxVariableWrapper) IsFixedType() bool {
	return varWrapper.isFixedType
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
