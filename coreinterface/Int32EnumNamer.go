package coreinterface

type Int32EnumNamer interface {
	ToNamer
	ValueInt32() int32
	Stringer
}
