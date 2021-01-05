package filemode

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coremath"
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
	read = coremath.GetValue(attribute.IsRead, ReadValue)
	write = coremath.GetValue(attribute.IsWrite, WriteValue)
	exe = coremath.GetValue(attribute.IsExecute, ExecuteValue)

	return read, write, exe, read + write + exe
}

func (attribute Attribute) ToByte() byte {
	r := coremath.GetValue(attribute.IsRead, ReadValue)
	w := coremath.GetValue(attribute.IsWrite, WriteValue)
	e := coremath.GetValue(attribute.IsExecute, ExecuteValue)

	return r + w + e
}

func (attribute Attribute) ToSum() byte {
	return attribute.ToByte()
}

func (attribute Attribute) ToRwx() [3]byte {
	return [3]byte{
		coremath.GetValueWithDefault(attribute.IsRead, ReadChar, constants.HyphenChar),
		coremath.GetValueWithDefault(attribute.IsWrite, WriteChar, constants.HyphenChar),
		coremath.GetValueWithDefault(attribute.IsExecute, ExecuteChar, constants.HyphenChar),
	}
}

func (attribute Attribute) ToRwxString() string {
	rwxBytes := attribute.ToRwx()

	return string(rwxBytes[:])
}

func (attribute Attribute) ToVariant() AttrVariant {
	b := attribute.ToByte()

	return AttrVariant(b)
}

func (attribute Attribute) ToChar() byte {
	return attribute.ToByte() + constants.ZeroChar
}
