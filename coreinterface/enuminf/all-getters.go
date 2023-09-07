package enuminf

import "gitlab.com/auk-go/core/internal/internalinterface"

type SplitNameValueByteGetter interface {
	enumNameStinger
	Value() byte
}

type SplitNameValueIntegerGetter interface {
	enumNameStinger
	Value() int
}

type SplitNameValueInteger8Getter interface {
	enumNameStinger
	Value() int8
}

type SplitNameValueInteger16Getter interface {
	enumNameStinger
	Value() int16
}

type SplitNameValueInteger32Getter interface {
	enumNameStinger
	Value() int32
}

type TypeNameGetter interface {
	TypeName() string
}

type StringRangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type TypeNameWithRangeNamesCsvGetter interface {
	StringRangeNamesCsvGetter
	TypeNameGetter
}

type ByteTypeEnumGetter interface {
	TypeEnum() BasicByteEnumContractsBinder
}

type StringRangesGetter interface {
	StringRangesPtr() *[]string
	StringRanges() []string
}

type BasicEnumerGetter interface {
	TypeBasicEnum() BasicEnumer
}

type RangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type RangesIntegerStringMapGetter interface {
	RangesIntegerStringMap() map[int]string
}

type RangesDynamicMapGetter interface {
	RangesDynamicMap() map[string]interface{}
}

type IntegerEnumRangesGetter interface {
	IntegerEnumRanges() []int
}

type LoggerTyperGetter interface {
	LoggerTyper() LoggerTyper
}

type EventTyperGetter interface {
	EventTyper() EventTyper
}

type ErrorStringGetter interface {
	internalinterface.ErrorStringGetter
}
