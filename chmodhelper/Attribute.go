package chmodhelper

import (
	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/constants"
)

type Attribute struct {
	IsRead    bool
	IsWrite   bool
	IsExecute bool
}

func (it *Attribute) IsNull() bool {
	return it == nil
}

func (it *Attribute) IsAnyNull() bool {
	return it == nil
}

func (it *Attribute) IsEmpty() bool {
	return it == nil ||
		!it.IsRead &&
			!it.IsWrite &&
			!it.IsExecute
}

func (it *Attribute) IsZero() bool {
	return it.IsEmpty()
}

func (it *Attribute) IsInvalid() bool {
	return it.IsEmpty()
}

func (it *Attribute) IsDefined() bool {
	return !it.IsEmpty()
}

func (it *Attribute) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *Attribute) ToAttributeValue() AttributeValue {
	read, write, exe, sum := it.ToSpecificBytes()

	return AttributeValue{
		Read:    read,
		Write:   write,
		Execute: exe,
		Sum:     sum,
	}
}

func (it *Attribute) ToSpecificBytes() (read, write, exe, sum byte) {
	read = conditional.Byte(it.IsRead, ReadValue, constants.Zero)
	write = conditional.Byte(it.IsWrite, WriteValue, constants.Zero)
	exe = conditional.Byte(it.IsExecute, ExecuteValue, constants.Zero)

	return read, write, exe, read + write + exe
}

// ToByte refers to the compiled byte value in between 0-7
func (it *Attribute) ToByte() byte {
	r := conditional.Byte(it.IsRead, ReadValue, constants.Zero)
	w := conditional.Byte(it.IsWrite, WriteValue, constants.Zero)
	e := conditional.Byte(it.IsExecute, ExecuteValue, constants.Zero)

	return r + w + e
}

// ToSum refers to the compiled byte value in between 0-7
func (it *Attribute) ToSum() byte {
	return it.ToByte()
}

func (it *Attribute) ToRwx() [3]byte {
	return [3]byte{
		conditional.Byte(it.IsRead, ReadChar, constants.HyphenChar),
		conditional.Byte(it.IsWrite, WriteChar, constants.HyphenChar),
		conditional.Byte(it.IsExecute, ExecuteChar, constants.HyphenChar),
	}
}

// ToRwxString returns "rwx"
func (it *Attribute) ToRwxString() string {
	rwxBytes := it.ToRwx()

	return string(rwxBytes[:])
}

func (it *Attribute) ToVariant() AttrVariant {
	b := it.ToByte()

	return AttrVariant(b)
}

// ToStringByte returns the compiled byte value as Char byte value
//
// It is not restricted between 0-7 but 0-7 + char '0', which makes it string 0-7
func (it *Attribute) ToStringByte() byte {
	return it.ToByte() + constants.ZeroChar
}

func (it *Attribute) Clone() *Attribute {
	if it == nil {
		return nil
	}

	return &Attribute{
		IsRead:    it.IsRead,
		IsWrite:   it.IsWrite,
		IsExecute: it.IsExecute,
	}
}

func (it *Attribute) IsEqualPtr(next *Attribute) bool {
	if it == nil && next == nil {
		return true
	}

	if it == nil || next == nil {
		return false
	}

	isRead := it.IsRead == next.IsRead
	isWrite := it.IsWrite == next.IsWrite
	isExecute := it.IsExecute == next.IsExecute

	return isRead &&
		isWrite &&
		isExecute
}

func (it Attribute) IsEqual(next Attribute) bool {
	isRead := it.IsRead == next.IsRead
	isWrite := it.IsWrite == next.IsWrite
	isExecute := it.IsExecute == next.IsExecute

	return isRead &&
		isWrite &&
		isExecute
}
