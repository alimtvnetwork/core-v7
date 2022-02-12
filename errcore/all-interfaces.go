package errcore

type errorGetter interface {
	Error() error
}

type compiledErrorGetter interface {
	CompiledError() error
}

type compiledErrorWithTracesGetter interface {
	CompiledErrorWithStackTraces() error
}

type fullStringWithTracesGetter interface {
	FullStringWithTraces() error
}

type lengthGetter interface {
	Length() int
}
