package internalinterface

type LogTypeChecker interface {
	IsSilent() bool
	HasNoLog() bool
	IsSkip() bool

	IsSuccess() bool

	IsInfo() bool
	IsTrace() bool
	IsDebug() bool
	IsError() bool
	IsFatal() bool
	IsPanic() bool
	IsInvalid() bool

	IsCustom() bool
	IsFile() bool
	HasPattern(pattern string) bool
	IsErrorLogical() bool
	IsErrorFatalLogical() bool
	IsErrorFatalPanicLogical() bool
}
