package coreinterface

import "gitlab.com/evatix-go/core/coredata/corejson"

type BasicEnumer interface {
	ToNamer
	NameValuer
	ToNumberStringer
	Stringer
	IsValidChecker
	IsInvalidChecker
	corejson.JsonMarshaller
}

type BasicEnumContractsBinder interface {
	BasicEnumer
	TypeNameWithRangeNamesCsvGetter
	AsBasicEnumContractsBinder() BasicEnumContractsBinder
}

type StandardEnumer interface {
	BasicEnumer
	StringRangesGetter
	RangeValidateChecker
	corejson.JsonContractsBinder
}

type StandardEnumerContractsBinder interface {
	StandardEnumer
	AsStandardEnumerContractsBinder() StandardEnumerContractsBinder
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

type BasicIn16EnumContractsDelegateBinder interface {
	BasicEnumContractsBinder
	BasicInt16Enumer
	Int16ToEnumStringer
	AsBasicIn16EnumContractsDelegateBinder() BasicIn16EnumContractsDelegateBinder
}

type BasicInt8EnumContractsDelegateBinder interface {
	BasicEnumContractsBinder
	BasicInt8Enumer
	Int8ToEnumStringer
	AsBasicInt8EnumContractsDelegateBinder() BasicInt8EnumContractsDelegateBinder
}

type UnmarshallToValueByte interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (byte, error)
}

type UnmarshallToValueInt interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int, error)
}

type UnmarshallToValueInt8 interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int8, error)
}

type UnmarshallToValueInt16 interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int16, error)
}

type UnmarshallEnumToValueByte interface {
	UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error)
}

type BasicByteEnumer interface {
	UnmarshallEnumToValueByte
	MaxByte() byte
	MinByte() byte
	ValueByte() byte
	RangesByte() []byte
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

type BasicInt32Enumer interface {
	UnmarshallEnumToValueInt32(jsonUnmarshallingValue []byte) (int32, error)
	MaxInt32() int32
	MinInt32() int32
	ValueInt32() int32
	RangesInt32() []int32
	ToEnumString(input int32) string
}

type BasicInt16Enumer interface {
	UnmarshallEnumToValueInt16(jsonUnmarshallingValue []byte) (int16, error)
	MaxInt16() int16
	MinInt16() int16
	ValueInt16() int16
	RangesInt16() []int16
	ToEnumString(input int16) string
}

type BasicInt8Enumer interface {
	UnmarshallEnumToValueInt8(jsonUnmarshallingValue []byte) (int8, error)
	MaxInt8() int8
	MinInt8() int8
	ValueInt8() int8
	RangesInt8() []int8
	ToEnumString(input int8) string
}

type BasicIntEnumer interface {
	MaxInt() int
	MinInt() int
	ValueInt() int
	RangesInt() []int
	ToEnumString(input int) string
}

type BasicInt64Enumer interface {
	MaxInt64() int64
	MinInt64() int64
	ValueInt64() int64
	RangesInt64() []int64
}
