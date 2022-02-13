package internalinterface

type LogTypeChecker interface {
	IsSilent() bool
	HasNoLog() bool
	IsSkip() bool
	IsInfo() bool
	IsTrace() bool
	IsDebug() bool
	IsError() bool
	IsFatal() bool
	IsPanic() bool
	IsSuccess() bool
	IsInvalid() bool
	IsErrorLogical() bool
	IsErrorFatalLogical() bool
	IsErrorFatalPanicLogical() bool
}
