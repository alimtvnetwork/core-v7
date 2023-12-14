package coretests

import (
	"reflect"

	"gitlab.com/auk-go/core/issetter"
)

// VerifyTypeOf
//
// # Verify type of
//
//   - BaseTestCase.ArrangeInput : What we set for initial input
//   - BaseTestCase.ActualInput : must set BaseTestCase.SetActual
//   - BaseTestCase.ExpectedInput : What we expect in the test case
type VerifyTypeOf struct {
	IsVerify      issetter.Value // Only false makes it verify, creating and attaching it makes it verify true by default.
	ArrangeInput  reflect.Type   // Verify type for the BaseTestCase.ArrangeInput
	ActualInput   reflect.Type   // Verify type for the BaseTestCase.ActualInput, must set BaseTestCase.SetActual
	ExpectedInput reflect.Type   // Verify type for the BaseTestCase.ExpectedInput
}

func NewVerifyTypeOf(arrange interface{}) *VerifyTypeOf {
	return &VerifyTypeOf{
		IsVerify:      issetter.True,
		ArrangeInput:  reflect.TypeOf(arrange),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}
}

func (it *VerifyTypeOf) IsDefined() bool {
	return it != nil
}

func (it *VerifyTypeOf) IsInvalid() bool {
	return it != nil
}

func (it *VerifyTypeOf) IsInvalidOrSkipVerify() bool {
	return it == nil || it.IsVerify.IsFalse()
}
