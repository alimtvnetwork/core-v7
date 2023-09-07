package internalenuminf

import "gitlab.com/auk-go/core/coredata/corejson"

type StandardEnumer interface {
	BasicEnumer
	StringRangesGetter
	RangeValidateChecker
	corejson.JsonContractsBinder
}

type BasicEnumValuer interface {
	ValueByte() byte
	ValueInt() int
	ValueInt8() int8
	ValueInt16() int16
	ValueInt32() int32
	ValueString() string // value in string format
}

type BasicByteEnumer interface {
	UnmarshallEnumToValueByter
	MaxByte() byte
	MinByte() byte
	ValueByte() byte
	RangesByte() []byte
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
