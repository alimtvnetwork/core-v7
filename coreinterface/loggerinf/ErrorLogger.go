package loggerinf

type ErrorLogger interface {
	Error(args ...interface{}) // Error logs a message at Error level.
}
