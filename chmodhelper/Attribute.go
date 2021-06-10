package chmodhelper

import (
	"gitlab.com/evatix-go/core/conditional"
	"gitlab.com/evatix-go/core/constants"
)

type Attribute struct {
	IsRead    bool
	IsWrite   bool
	IsExecute bool
}

func (attribute Attribute) ToAttributeValue() AttributeValue {
	read, write, exe, sum := attribute.ToSpecificBytes()

	return AttributeValue{
		Read:    read,
		Write:   write,
		Execute: exe,
		Sum:     sum,
	}
}

func (attribute Attribute) ToSpecificBytes() (read, write, exe, sum byte) {
	read = conditional.Byte(attribute.IsRead, ReadValue, constants.Zero)
	write = conditional.Byte(attribute.IsWrite, WriteValue, constants.Zero)
	exe = conditional.Byte(attribute.IsExecute, ExecuteValue, constants.Zero)

	return read, write, exe, read + write + exe
}

// ToByte refers to the compiled byte value in between 0-7
func (attribute Attribute) ToByte() byte {
	r := conditional.Byte(attribute.IsRead, ReadValue, constants.Zero)
	w := conditional.Byte(attribute.IsWrite, WriteValue, constants.Zero)
	e := conditional.Byte(attribute.IsExecute, ExecuteValue, constants.Zero)

	return r + w + e
}

// ToSum refers to the compiled byte value in between 0-7
func (attribute Attribute) ToSum() byte {
	return attribute.ToByte()
}

func (attribute Attribute) ToRwx() [3]byte {
	return [3]byte{
		conditional.Byte(attribute.IsRead, ReadChar, constants.HyphenChar),
		conditional.Byte(attribute.IsWrite, WriteChar, constants.HyphenChar),
		conditional.Byte(attribute.IsExecute, ExecuteChar, constants.HyphenChar),
	}
}

// ToRwxString returns "rwx"
func (attribute Attribute) ToRwxString() string {
	rwxBytes := attribute.ToRwx()

	return string(rwxBytes[:])
}

func (attribute Attribute) ToVariant() AttrVariant {
	b := attribute.ToByte()

	return AttrVariant(b)
}

// ToStringByte returns the compiled byte value as Char byte value
//
// It is not restricted between 0-7 but 0-7 + char '0', which makes it string 0-7
func (attribute Attribute) ToStringByte() byte {
	return attribute.ToByte() + constants.ZeroChar
}
