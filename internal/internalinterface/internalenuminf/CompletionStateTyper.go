package internalenuminf

type CompletionStateTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsStartState() bool

	IsRunning() bool
	IsSuccess() bool
	IsSuccessWithWarning() bool
	IsFailedMiddleWithError() bool
	IsCompleteWithError() bool

	IsEndState() bool

	IsCompletedSuccess() bool
	IsCompletedWithIssues() bool
}
