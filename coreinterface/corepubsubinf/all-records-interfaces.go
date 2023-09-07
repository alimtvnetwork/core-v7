package corepubsubinf

import (
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/loggerinf"
	"gitlab.com/auk-go/core/coreinterface/pathextendinf"
	"gitlab.com/auk-go/core/internal/internalinterface"
)

type IdAsStringer interface {
	internalinterface.IdAsStringer
}

type SubscriptionMainRecorder interface {
	IdAsStringer
	TableName() string

	IsEmpty() bool

	pathextendinf.PathExtenderGetter

	HasRecordError() bool
	SetRecordError() bool
	IsArchivedRecord() bool
	IsCompletedRecord() bool
	IsMigratedRecord() bool
	CompletionTyper() enuminf.CompletionStateTyper

	// DefaultDelayMillis
	//
	//  Default delay in milliseconds
	DefaultDelayMillis() int
}

type BaseLogModeler interface {
	enuminf.LoggerTyperGetter
	enuminf.EventTyperGetter
	errcoreinf.BasicErrorTyperGetter
	errcoreinf.ErrorStringGetter
	coreinterface.StackTracesBytesGetter
	coreinterface.JsonErrorBytesGetter
	IsEmpty() bool
	LogMessage() string
}

type CommunicateModeler interface {
	BaseLogModeler() BaseLogModeler
	PersistentId() uint
	IdAsStringer
	TableName() string

	SetCallerFileLineUsingStackSkip(
		stackSkip int,
	)

	loggerinf.SingleLogModeler
}

type SubscriptionRecorder interface {
	MainRecord() SubscriptionMainRecorder
	CommunicateRecord() CommunicateModeler
}
