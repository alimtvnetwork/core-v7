package loggerinf

// Logger
//
// logs message to io.Writer at various log levels.
type Logger interface {
	FatalLogger   // Fatal logs a message at Fatal level
	ErrorLogger   // Error logs a message at Error level.
	WarningLogger // Warn logs a message at Warning level.
	InfoLogger    // Info logs a message at Info level.
	DebugLogger   // Debug logs a message at Debug level.
}
