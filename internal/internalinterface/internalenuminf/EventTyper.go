package internalenuminf

type EventTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsLog() bool
	IsSuccess() bool
	IsError() bool
	IsFailure() bool
	IsDebug() bool
}
