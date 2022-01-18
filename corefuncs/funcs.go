package corefuncs

import "gitlab.com/evatix-go/core/coredata/corepayload"

type (
	VoidFunc      func()
	NamedVoidFunc func(name string)
	ReturnErrFunc func() error
	IsSuccessFunc func() (isSuccess bool)
	// ResultDelegatingFunc
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set
	ResultDelegatingFunc           func(resultDelegatedTo interface{}) error
	NextReturnErrWrapperFunc       func(nextAction ReturnErrFunc) error
	NextVoidActionFunc             func(nextAction VoidFunc)
	PayloadProcessorFunc           func(payload []byte) (err error)
	PayloadToPayloadWrapperFunc    func(payload []byte) (payloadWrapper *corepayload.PayloadWrapper, err error)
	NextPayloadProcessorLinkerFunc func(nextLinkerFunc PayloadProcessorFunc) error
)
