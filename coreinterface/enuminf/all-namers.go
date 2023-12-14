package enuminf

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

type Namer interface {
	Name() string
}

type ByteEnumNamer interface {
	enumNameStinger
	ValueByte() byte
}

// NameValuer / NameValue
//
//	should be a combined string output using name[value]
//
// Sample :
//   - "`EnumName[EnumValInteger]" -> `EnumName[2]`
type NameValuer interface {
	// NameValue
	//
	//   should be a combined string output using name[value]
	//
	// Sample :
	//  - "`EnumName[EnumValInteger]" -> `EnumName[2]`
	NameValue() string
}
