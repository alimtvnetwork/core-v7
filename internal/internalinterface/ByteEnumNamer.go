package internalinterface

type ByteEnumNamer interface {
	ToNamer
	ValueByte() byte
	Stringer
}
