package corefuncs

import "gitlab.com/evatix-go/core/coredata/corepayload"

type (
	ExecFunc               func()
	ActionFunc             func()
	IsApplyFunc            func() (isSuccess bool)
	InOutFunc              func(input interface{}) (output interface{})
	InOutErrFunc           func(input interface{}) (output interface{}, err error)
	InActionReturnsErrFunc func(input interface{}) (err error)
	NamedActionFunc        func(name string)
	ActionReturnsErrorFunc func() error
	IsSuccessFunc          func() (isSuccess bool)
	IsFailureFunc          func() (isFailed bool)
	// ResultDelegatingFunc
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set
	ResultDelegatingFunc           func(resultDelegatedTo interface{}) error
	NextReturnErrWrapperFunc       func(nextAction ActionReturnsErrorFunc) error
	NextVoidActionFunc             func(nextAction ExecFunc)
	PayloadProcessorFunc           func(payloads []byte) (err error)
	PayloadToPayloadWrapperFunc    func(payloads []byte) (payloadWrapper *corepayload.PayloadWrapper, err error)
	NextPayloadProcessorLinkerFunc func(nextLinkerFunc PayloadProcessorFunc) error
)
