package chmodhelper

import "gitlab.com/auk-go/core/errcore"

type newAttributeCreator struct{}

func (it newAttributeCreator) Create(
	isRead, isWrite, isExecute bool,
) Attribute {
	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

func (it newAttributeCreator) Default(
	isRead, isWrite, isExecute bool,
) Attribute {
	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

// UsingRwxString
//
// Length must be 3
// "rwx" should be put for attributes.
//
// Examples:
//   - read enable all disable    : "r--"
//   - write enable all disable   : "-w-"
//   - execute enable all disable : "--x"
//   - all enabled                : "rwx"
func (it newAttributeCreator) UsingRwxString(
	rwx string,
) Attribute {
	length := len(rwx)

	if length != SingleRwxLength {
		panic(GetRwxLengthError(rwx))
	}

	r := rwx[0]
	w := rwx[1]
	e := rwx[2]

	return Attribute{
		IsRead:    r == ReadChar,
		IsWrite:   w == WriteChar,
		IsExecute: e == ExecuteChar,
	}
}

// UsingByteMust
//
//	Byte can be at most 0 to 7
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingByteMust(v7 byte) Attribute {
	attr, err := it.UsingByte(v7)

	if err != nil {
		panic(attr)
	}

	return attr
}

// UsingByte
//
//	Byte can be at most 0 to 7
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingByte(v7 byte) (Attribute, error) {
	if ReadWriteExecute.IsGreaterThan(v7) {
		return Attribute{}, errcore.
			ShouldBeLessThanEqualType.
			Error(
				"v7 byte should not be more than "+ReadWriteExecute.String(),
				v7)
	}

	// TODO optimize logic in future.
	isRead := v7 >= ReadValue
	isWrite := (isRead && v7 >= ReadWriteValue) || (!isRead && v7 >= WriteValue)
	isExecute := (isWrite && isRead && v7 >= ReadWriteExecuteValue) ||
		(isRead && !isWrite && v7 >= ReadExecuteValue) ||
		(isWrite && !isRead && v7 >= WriteExecuteValue) ||
		(!isRead && !isWrite && v7 >= ExecuteValue)

	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}, nil
}

// UsingVariantMust
//
// safe because converting AttrVariant should never exceed 7
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingVariantMust(v AttrVariant) Attribute {
	return it.UsingByteMust(v.Value())
}

func (it newAttributeCreator) UsingVariant(v AttrVariant) (Attribute, error) {
	return it.UsingByte(v.Value())
}
