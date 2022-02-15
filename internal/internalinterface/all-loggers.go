package internalinterface

type VoidLogger interface {
	// Log
	//
	//  Prints the compiled error message with all types
	//  only not fatal or panic
	Log()
}

type VoidTracesLogger interface {
	// LogWithTraces
	//
	//  Prints the compiled error message with all types
	//  and stack-traces but not fatal or panic
	LogWithTraces()
}

type FatalVoidLogger interface {
	LogFatal()
}

type FatalTracesVoidLogger interface {
	LogFatalWithTraces()
}

type VoidIfLogger interface {
	LogIf(isLog bool)
}

type CompiledVoidLogger interface {
	VoidLogger
	VoidTracesLogger
	FatalVoidLogger
	FatalTracesVoidLogger
	VoidIfLogger
}
