package corefuncs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corepayload"
)

type (
	ExecFunc                    func()
	StringerActionFunc          func() (result string)
	StringerWithErrorActionFunc func() (result string, err error)
	ActionFunc                  func()
	IsApplyFunc                 func() (isSuccess bool)
	InOutFunc                   func(input interface{}) (output interface{})
	InOutErrFunc                func(input interface{}) (output interface{}, err error)
	SerializeOutputFunc         func(input interface{}) (serializedBytes []byte, err error)
	SerializerVoidFunc          func() (serializedBytes []byte, err error)
	InActionReturnsErrFunc      func(input interface{}) (err error)
	NamedActionFunc             func(name string)
	ActionReturnsErrorFunc      func() error
	IsSuccessFunc               func() (isSuccess bool)
	IsFailureFunc               func() (isFailed bool)
	// ResultDelegatingFunc
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set
	ResultDelegatingFunc                 func(resultDelegatedTo interface{}) error
	NextReturnErrWrapperFunc             func(nextAction ActionReturnsErrorFunc) error
	NextVoidActionFunc                   func(nextAction ExecFunc)
	PayloadProcessorFunc                 func(payloads []byte) (err error)
	MultiPayloadsProcessorFunc           func(multiPayloads ...[]byte) (err error)
	BytesCollectionPayloadsProcessorFunc func(collectionOfBytes *corejson.BytesCollection) (err error)
	PayloadToPayloadWrapperFunc          func(payloads []byte) (payloadWrapper *corepayload.PayloadWrapper, err error)
	NextPayloadProcessorLinkerFunc       func(nextLinkerFunc PayloadProcessorFunc) error
)
