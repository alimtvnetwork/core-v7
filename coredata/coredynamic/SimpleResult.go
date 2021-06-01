package coredynamic

import (
	"errors"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/stringutil"
	"gitlab.com/evatix-go/core/msgtype"
)

type SimpleResult struct {
	Result         interface{}
	InvalidMessage string
	err            error
	Dynamic
}

func NewSimpleResult(
	result interface{},
	isValid bool,
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:         result,
		Dynamic:        NewDynamic(result, isValid),
		InvalidMessage: invalidMessage,
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

	return errors.New(typeMismatchMessage + receiver.InvalidMessage)
}

func (receiver *SimpleResult) InvalidError() error {
	if receiver.err != nil {
		return receiver.err
	}

	if stringutil.IsEmptyOrWhitespace(receiver.InvalidMessage) {
		return nil
	}

	if receiver.err == nil {
		receiver.err = errors.New(receiver.InvalidMessage)
	}

	return receiver.err
}
