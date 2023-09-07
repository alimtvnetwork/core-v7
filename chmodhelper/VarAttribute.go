package chmodhelper

import (
	"gitlab.com/auk-go/core/issetter"
)

type VarAttribute struct {
	rawInput    string
	isFixedType bool
	isRead      issetter.Value
	isWrite     issetter.Value
	isExecute   issetter.Value
}

func (it *VarAttribute) IsFixedType() bool {
	return it.isFixedType
}

func (it *VarAttribute) HasWildcard() bool {
	return !it.isFixedType
}

// ToCompileFixAttr
//
//  must check IsFixedType, before calling.
func (it *VarAttribute) ToCompileFixAttr() *Attribute {
	if it.isFixedType {
		return &Attribute{
			IsRead:    it.isRead.IsTrue(),
			IsWrite:   it.isWrite.IsTrue(),
			IsExecute: it.isExecute.IsTrue(),
		}
	}

	return nil
}

// ToCompileAttr
//
//  if fixed type then fixed param can be nil
func (it *VarAttribute) ToCompileAttr(fixed *Attribute) Attribute {
	if it.isFixedType {
		return Attribute{
			IsRead:    it.isRead.IsTrue(),
			IsWrite:   it.isWrite.IsTrue(),
			IsExecute: it.isExecute.IsTrue(),
		}
	}

	return Attribute{
		IsRead:    it.isRead.WildcardApply(fixed.IsRead),
		IsWrite:   it.isWrite.WildcardApply(fixed.IsWrite),
		IsExecute: it.isExecute.WildcardApply(fixed.IsExecute),
	}
}

func (it *VarAttribute) Clone() *VarAttribute {
	if it == nil {
		return nil
	}

	return &VarAttribute{
		rawInput:    it.rawInput,
		isFixedType: it.IsFixedType(),
		isRead:      it.isRead,
		isWrite:     it.isWrite,
		isExecute:   it.isExecute,
	}
}

func (it *VarAttribute) IsEqualPtr(next *VarAttribute) bool {
	if it == nil && next == nil {
		return true
	}

	if it == nil || next == nil {
		return false
	}

	isRead := next.isRead == it.isRead
	isWrite := next.isWrite == it.isWrite
	isExecute := next.isExecute == it.isExecute

	return isRead &&
		isWrite &&
		isExecute
}

func (it *VarAttribute) String() string {
	return it.rawInput
}
