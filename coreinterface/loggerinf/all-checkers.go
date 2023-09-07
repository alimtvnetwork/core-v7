package loggerinf

type hasErrorChecker interface {
	HasError() bool
}

type HasParentChecker interface {
	HasParent() bool
}

type HasModelChecker interface {
	HasModel() bool
}

type StandardLoggerChecker interface {
	IsVerbose() bool
	IsProduction() bool
	IsTest() bool
	IsDebug() bool
	IsLog() bool
	IsStacktrace() bool
	IsJson() bool
}
