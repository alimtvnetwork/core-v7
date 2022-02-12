package loggerinf

type persistentAllStacktraceLogger interface {
	// ErrorStackTraces
	//
	// Includes stack-traces
	ErrorStackTraces(err error) PersistentLogger
	// TitleErrorStackTraces
	//
	// Includes stack-traces
	TitleErrorStackTraces(title string, err error) PersistentLogger
	// DebugStackTraces
	//
	// Includes stack-traces
	DebugStackTraces(message string) PersistentLogger
	// DebugAttrStackTraces
	//
	// Includes stack-traces
	DebugAttrStackTraces(message, attr string) PersistentLogger
	// StackTracesIf
	//
	// Log StackTraces
	StackTracesIf(isCondition bool) PersistentLogger
	// StackTraces
	//
	// Log StackTraces as Info
	StackTraces() PersistentLogger
	// StackTracesSkip
	//
	// Log StackTraces
	StackTracesSkip(stackSkipIndex int) PersistentLogger
	// StackTracesSkipIf
	//
	// Log StackTraces
	StackTracesSkipIf(isCondition bool, stackSkipIndex int) PersistentLogger
	TitleStackTraces(title string) PersistentLogger
	TitleStackTracesSkip(stackSkipIndex int, title string) PersistentLogger
}
