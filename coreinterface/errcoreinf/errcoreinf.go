package errcoreinf

import (
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type BasicErrWrapper interface {
	internalinterface.BasicErrWrapper
	ErrorTypeAsBasicErrorTyper() BasicErrorTyper
}

type AddErrorer interface {
	internalinterface.AddErrorer
}

type IsErrorsCollected interface {
	internalinterface.IsErrorsCollected
}

type BaseRawErrCollectionDefiner interface {
	internalinterface.BaseRawErrCollectionDefiner
}

type FullStringWithTracesGetter interface {
	internalinterface.FullStringWithTracesGetter
}

type FullStringWithTracesIfGetter interface {
	internalinterface.FullStringWithTracesIfGetter
}

type BaseErrorTyper interface {
	internalinterface.BaseErrorTyper
}

type BasicErrorTyper interface {
	BaseErrorTyper
	ErrorTypeAsBasicEnum() enuminf.BasicEnumer
}

type DyanmicLinqer interface {
	internalinterface.DyanmicLinqer
}

type AddManyErrorer interface {
	internalinterface.AddManyErrorer
}

type AddManyPointerErrorer interface {
	internalinterface.AddManyPointerErrorer
}

type ConditionalErrorAdder interface {
	internalinterface.ConditionalErrorAdder
}

type BaseErrorWrapperCollectionDefiner interface {
	internalinterface.BaseErrorWrapperCollectionDefiner
	AddErrorUsingBasicType(errType BasicErrorTyper, err error) BaseErrorWrapperCollectionDefiner
	AddBasicErrWrapper(basicErrWrapper BasicErrWrapper) BaseErrorWrapperCollectionDefiner
}
