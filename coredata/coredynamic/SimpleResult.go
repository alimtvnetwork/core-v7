package coredynamic

import (
	"errors"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
	"gitlab.com/evatix-go/core/msgtype"
)

type SimpleResult struct {
	Dynamic
	Result  interface{}
	Message string
	err     error
}

func InvalidSimpleResultNoMessage() *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: constants.EmptyString,
	}
}

func InvalidSimpleResult(
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: invalidMessage,
	}
}

func NewSimpleResult(
	result interface{},
	isValid bool,
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  result,
		Dynamic: NewDynamic(result, isValid),
		Message: invalidMessage,
	}
}

func (receiver *SimpleResult) GetErrorOnTypeMismatch(
	typeMatch reflect.Type,
	isIncludeInvalidMessage bool,
) error {
	if receiver.IsReflectTypeOf(typeMatch) {
		return nil
	}

	typeMismatchMessage := msgtype.CombineWithMsgType(
		msgtype.TypeMismatch,
		"Current type - ["+receiver.ReflectTypeName()+"], expected type",
		typeMatch) + constants.NewLineUnix

	if !isIncludeInvalidMessage {
		return errors.New(typeMismatchMessage)
	}

	return errors.New(typeMismatchMessage + receiver.Message)
}

func (receiver *SimpleResult) InvalidError() error {
	if receiver.err != nil {
		return receiver.err
	}

	if strutilinternal.IsEmptyOrWhitespace(receiver.Message) {
		return nil
	}

	if receiver.err == nil {
		receiver.err = errors.New(receiver.Message)
	}

	return receiver.err
}

func (receiver *SimpleResult) Clone() *SimpleResult {
	if receiver == nil {
		return nil
	}

	return &SimpleResult{
		Dynamic: receiver.Dynamic,
		Result:  receiver.Result,
		Message: receiver.Message,
	}
}
