package coreinterface

type TypeWriter interface {
	TypeWrite(typeNameGetter TypeNameGetter, content string) error
	TypeWriteMust(typeNameGetter TypeNameGetter, content string)
}
