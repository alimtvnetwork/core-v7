package coreinterface

type FormatLogger interface {
	FormatFatalLogger   // Fatal logs a message at Fatal level
	FormatErrorLogger   // Error logs a message at Error level.
	FormatWarningLogger // Warn logs a message at Warning level.
	FormatInfoLogger    // Info logs a message at Info level.
	FormatDebugLogger   // Debug logs a message at Debug level.
}
