package internalenuminf

type Int8EnumNamer interface {
	enumNameStinger
	ValueInt8() int8
}

type Int16EnumNamer interface {
	enumNameStinger
	ValueInt16() int16
}

type Int32EnumNamer interface {
	enumNameStinger
	ValueInt32() int32
}

type namer interface {
	Name() string
}

type ByteEnumNamer interface {
	enumNameStinger
	ValueByte() byte
}
