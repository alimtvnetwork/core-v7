package internalenuminf

import "fmt"

type ByteToEnumStringer interface {
	ToByteEnumString(input byte) string
}

type IntToEnumStringer interface {
	ToIntEnumString(input int) string
}

type Int8ToEnumStringer interface {
	ToInt8EnumString(input int8) string
}

type Int16ToEnumStringer interface {
	ToInt16EnumString(input int16) string
}

type Int32ToEnumStringer interface {
	ToInt32EnumString(input int32) string
}

type enumNameStinger interface {
	namer
	fmt.Stringer
}

// ToNumberStringer
//
//	It returns string number value.
//
// Examples:
//   - ToNumberString() -> "1"  if the value is 1
//   - ToNumberString() -> "10" if the value is 10
type toNumberStringer interface {
	// ToNumberString
	//
	//  It returns string number value.
	//
	// Examples:
	//  - ToNumberString() -> "1"  if the value is 1
	//  - ToNumberString() -> "10" if the value is 10
	ToNumberString() string
}
