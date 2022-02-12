package internalinterface

type ErrWrapperLogger interface {
	// Log
	//
	//  Prints the compiled error message with all types
	//  only not fatal or panic
	Log()
	// LogWithTraces
	//
	//  Prints the compiled error message with all types
	//  and stack-traces but not fatal or panic
	LogWithTraces()
	LogFatal()
	LogFatalWithTraces()
	LogIf(isLog bool)
}
