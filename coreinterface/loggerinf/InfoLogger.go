package loggerinf

type InfoLogger interface {
	Info(args ...interface{}) // Info logs a message at Info level.
}
