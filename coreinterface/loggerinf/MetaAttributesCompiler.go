package loggerinf

import "gitlab.com/evatix-go/core/coreinterface"

type MetaAttributesCompiler interface {
	coreinterface.Disposer
	StringFinalizer
	IfStringCompiler
	Compiler
	FmtCompiler
	Comitter
	CompileAnyTo(toPointer interface{}) error
	CompileAny() interface{}
	CompileStacks() []string
	ReflectSetter
	CompileMap() map[string]interface{}
}
