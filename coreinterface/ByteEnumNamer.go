package coreinterface

type ByteEnumNamer interface {
	ToNamer
	ValueByte() byte
	Stringer
}
