package internalenuminf

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
