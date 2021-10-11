package chmodhelper

import "gitlab.com/evatix-go/core/errcore"

//goland:noinspection GoUnusedExportedFunction
func NewAttribute(isRead, isWrite, isExecute bool) Attribute {
	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

// NewAttributeUsingRwx Length must be 3
// "rwx" should be put for attributes.
// eg. read enable all disable : "r--"
// eg. write enable all disable : "-w-"
// eg. execute enable all disable : "--x"
// eg. all enabled : "rwx"
func NewAttributeUsingRwx(rwx string) Attribute {
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

// NewAttributeUsingByte
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
func NewAttributeUsingByte(v byte) Attribute {
	if ReadWriteExecute.IsGreaterThan(v) {
		msg := errcore.
			ShouldBeLessThanEqualMessage.
			Combine(
				"v byte should not be more than "+ReadWriteExecute.String(),
				v)

		panic(msg)
	}

	// TODO optimize logic in future.
	isRead := v >= ReadValue
	isWrite := (isRead && v >= ReadWriteValue) || (!isRead && v >= WriteValue)
	isExecute := (isWrite && isRead && v >= ReadWriteExecuteValue) ||
		(isRead && !isWrite && v >= ReadExecuteValue) ||
		(isWrite && !isRead && v >= WriteExecuteValue) ||
		(!isRead && !isWrite && v >= ExecuteValue)

	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

func NewAttributeUsingVariant(v AttrVariant) Attribute {
	return NewAttributeUsingByte(v.Value())
}
