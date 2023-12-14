package trydo

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

type WrappedErr struct {
	Error     error
	Exception Exception
	HasThrown bool
	HasError  bool
}

func (it *WrappedErr) ExceptionType() reflect.Type {
	if it.IsInvalidException() {
		return nil
	}

	return reflect.TypeOf(it.Exception)
}

func (it *WrappedErr) IsDefined() bool {
	return it != nil
}

func (it *WrappedErr) IsInvalid() bool {
	return it == nil
}

func (it *WrappedErr) IsInvalidException() bool {
	if it.IsInvalid() {
		return true
	}

	if it.HasThrown {
		return it.Exception != nil
	}

	return true
}

func (it *WrappedErr) HasErrorOrException() bool {
	return it != nil &&
		it.HasError ||
		it.HasThrown
}

func (it *WrappedErr) IsBothPresent() bool {
	return it != nil &&
		it.HasError &&
		it.HasThrown &&
		it.Exception != nil
}

func (it *WrappedErr) HasException() bool {
	return it != nil &&
		it.Exception != nil
}

func (it *WrappedErr) ErrorString() string {
	if it.IsInvalid() {
		return ""
	}

	if it.Error != nil {
		return it.Error.Error()
	}

	return ""
}

func (it *WrappedErr) ExceptionString() string {
	if it.IsInvalidException() {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		it.Exception)
}

func (it *WrappedErr) String() string {
	if it.IsInvalid() {
		return ""
	}

	if it.IsBothPresent() {
		return it.ErrorString() +
			constants.CommaSpace +
			it.ExceptionString()
	}

	if it.HasError {
		return it.ErrorString()
	}

	if it.HasException() {
		return it.ExceptionString()
	}

	return ""
}
