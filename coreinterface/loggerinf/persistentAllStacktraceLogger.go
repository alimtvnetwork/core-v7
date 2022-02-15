package loggerinf

type persistentAllStacktraceLogger interface {
	// ErrorStackTraces
	//
	// Includes stack-traces
	ErrorStackTraces(err error) BasePersistentLogger
	// TitleErrorStackTraces
	//
	// Includes stack-traces
	TitleErrorStackTraces(title string, err error) BasePersistentLogger
	// DebugStackTraces
	//
	// Includes stack-traces
	DebugStackTraces(message string) BasePersistentLogger
	// DebugAttrStackTraces
	//
	// Includes stack-traces
	DebugAttrStackTraces(message, attr string) BasePersistentLogger
	// StackTracesIf
	//
	// Log StackTraces
	StackTracesIf(isCondition bool) BasePersistentLogger
	// StackTraces
	//
	// Log StackTraces as Info
	StackTraces() BasePersistentLogger
	// StackTracesSkip
	//
	// Log StackTraces
	StackTracesSkip(stackSkipIndex int) BasePersistentLogger
	// StackTracesSkipIf
	//
	// Log StackTraces
	StackTracesSkipIf(isCondition bool, stackSkipIndex int) BasePersistentLogger
	TitleStackTraces(title string) BasePersistentLogger
	TitleStackTracesSkip(stackSkipIndex int, title string) BasePersistentLogger
}
