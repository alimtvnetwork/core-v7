package coreinterface

type BasicEnumContractsBinder interface {
	StringRangesGetter
	RangeValidateChecker
	AsBasicEnumContractsBinder() BasicEnumContractsBinder
}

type BasicByteEnumContractsBinder interface {
	BasicEnumContractsBinder
	BasicByteEnumer
	AsBasicByteEnumContractsBinder() BasicByteEnumContractsBinder
}

type BasicByteEnumContractsDelegateBinder interface {
	BasicByteEnumContractsBinder
	ByteToEnumStringer
	AsBasicByteEnumContractsDelegateBinder() BasicByteEnumContractsDelegateBinder
}

type BasicByteEnumer interface {
	MaxByte() byte
	MinByte() byte
	ValueByte() byte
	RangesByte() *[]byte
}

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

type BasicInt16Enumer interface {
	MaxInt16() int8
	MinInt16() int8
	ValueInt16() int8
	RangesInt16() *[]int8
	ToEnumString(input int8) string
}

type BasicInt8Enumer interface {
	MaxInt8() int16
	MinInt8() int16
	ValueInt8() int16
	RangesInt8() *[]int16
	ToEnumString(input int16) string
}

type BasicIntEnumer interface {
	MaxInt() int
	MinInt() int
	ValueInt() int
	RangesInt() *[]int
	ToEnumString(input int) string
}

type BasicInt64Enumer interface {
	MaxInt64() int64
	MinInt64() int64
	ValueInt64() int64
	RangesInt64() *[]int64
}
