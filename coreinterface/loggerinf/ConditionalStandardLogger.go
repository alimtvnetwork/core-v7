package loggerinf

type ConditionalStandardLogger interface {
	On(isCondition bool) StandardLogger
	OnErr(err error) StandardLogger
	OnString(expected, actual string) StandardLogger
	OnBytes(expectedRawBytes, actualBytes []byte) StandardLogger

	OnVerbose() SingleLogger
	OnProduction() SingleLogger
	OnTest() SingleLogger
	OnDebug() SingleLogger
	OnJson() SingleLogger
	OnStacktrace() SingleLogger

	OnFlag(name, value string) StandardLogger
	OnAnyFlag(name string, value interface{}) StandardLogger
	OnFunc(isLoggerFunc func(logger StandardLogger) bool) StandardLogger
	OnFlagEnabled(name string) StandardLogger
	OnFlagDisabled(name string) StandardLogger
	StackSkip(index int) StandardLogger
}
