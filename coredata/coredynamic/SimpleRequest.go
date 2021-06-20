package coredynamic

import (
	"errors"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/core/msgtype"
)

type SimpleRequest struct {
	Dynamic
	message string
	err     error
}

func InvalidSimpleRequestNoMessage() *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(nil, false),
		message: constants.EmptyString,
	}
}

func InvalidSimpleRequest(
	message string,
) *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(nil, false),
		message: message,
	}
}

func NewSimpleRequest(
	request interface{},
	isValid bool,
	message string,
) *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(request, isValid),
		message: message,
	}
}

func (receiver *SimpleRequest) Message() string {
	return receiver.message
}

func (receiver *SimpleRequest) Request() interface{} {
	return receiver.Dynamic.Data()
}

func (receiver *SimpleRequest) Value() interface{} {
	return receiver.Dynamic.Data()
}

func (receiver *SimpleRequest) GetErrorOnTypeMismatch(
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

	return errors.New(typeMismatchMessage + receiver.message)
}

func (receiver *SimpleRequest) IsReflectKind(checkingKind reflect.Kind) bool {
	return receiver.ReflectKind() == checkingKind
}

func (receiver *SimpleRequest) IsPointer() bool {
	if receiver.isPointer.IsUninitialized() {
		receiver.isPointer = issetter.GetBool(
			receiver.IsReflectKind(reflect.Ptr))
	}

	return receiver.isPointer.IsTrue()
}

func (receiver *SimpleRequest) InvalidError() error {
	if receiver.err != nil {
		return receiver.err
	}

	if strutilinternal.IsEmptyOrWhitespace(receiver.message) {
		return nil
	}

	if receiver.err == nil {
		receiver.err = errors.New(receiver.message)
	}

	return receiver.err
}
