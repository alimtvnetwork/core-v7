package loggerinf

import "io"

type Committer interface {
	Commit()
}

type BytesCompiler interface {
	CompileBytes() ([]byte, error)
}

type BytesCompilerIf interface {
	CompileBytesIf(isCompile bool) ([]byte, error)
}

type MustBytesCompiler interface {
	CompileBytesMust() []byte
}

type StringFinalizer interface {
	Finalize() string
}

type IfStringCompiler interface {
	CompileIf(isCompile bool) string
}

type Compiler interface {
	Compile() string
}

type FmtCompiler interface {
	CompileFmt(formatter string, v ...interface{}) string
}

type Serializer interface {
	Serialize() ([]byte, error)
}

type NewGeneralWriter interface {
	NewGeneralWriter(writeConfigurer WriterConfigurer) io.Writer
}

type Configurer interface {
	LoggerTyper() LoggerTyper
	StackSkipIndex() int
}

type WriterConfigurer interface {
	Configurer
	AdditionalConfigProcessor
}

type AdditionalConfigProcessor interface {
	AdditionalConfigBytes() []byte
	AdditionalConfigProcess() error
}
