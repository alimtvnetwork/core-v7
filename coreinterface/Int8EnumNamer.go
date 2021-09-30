package coreinterface

type Int8EnumNamer interface {
	ToNamer
	ValueInt8() int8
	Stringer
}
