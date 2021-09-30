package coreinterface

type Int16EnumNamer interface {
	ToNamer
	ValueInt16() int16
	Stringer
}
