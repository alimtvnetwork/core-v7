package loggerinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
)

type MetaAttributesCompiler interface {
	coreinterface.Disposer

	StringFinalizer
	IfStringCompiler
	Compiler
	FmtCompiler
	// Committer
	//
	// logs and clears
	Committer
	CompileAnyTo(toPointer interface{}) error
	CompileAny() interface{}
	CompileStacks() []string
	ReflectSetter
	CompileMap() map[string]interface{}
	CompileToJsonResult() *corejson.Result

	CompiledAsBasicErr(
		basicErrTyper errcoreinf.BasicErrorTyper,
	) errcoreinf.BasicErrWrapper

	BytesCompiler
	BytesCompilerIf
	MustBytesCompiler
}
