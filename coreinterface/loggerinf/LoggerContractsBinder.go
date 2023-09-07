package loggerinf

type LoggerContractsBinder interface {
	Logger
	AsLogger() Logger
}
