package coreinterface

type SpecificTypeWriter interface {
	SpecificTypeWrite(typeName string, content string) error
	SpecificTypeWriteMust(typeName string, content string)
}
