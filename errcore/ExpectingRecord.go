package errcore

import (
	"errors"
	"fmt"
)

type ExpectingRecord struct {
	ExpectingTitle string
	WasExpecting   interface{}
}

// Message
// Expecting
//
// returns
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (it *ExpectingRecord) Message(actual interface{}) string {
	return fmt.Sprintf(
		expectingMessageFormat,
		it.ExpectingTitle,
		it.WasExpecting, it.WasExpecting,
		actual, actual)
}

// MessageSimple
// Expecting
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func (it *ExpectingRecord) MessageSimple(actual interface{}) string {
	return ExpectingSimple(
		it.ExpectingTitle,
		it.WasExpecting,
		actual)
}

// MessageSimpleNoType
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Actual"
func (it *ExpectingRecord) MessageSimpleNoType(actual interface{}) string {
	return ExpectingSimpleNoType(
		it.ExpectingTitle,
		it.WasExpecting,
		actual)
}

// Error
// Expecting
//
// returns
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (it *ExpectingRecord) Error(actual interface{}) error {
	return errors.New(it.Message(actual))
}

// ErrorSimple
// Expecting
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func (it *ExpectingRecord) ErrorSimple(actual interface{}) error {
	return errors.New(it.MessageSimple(actual))
}

// ErrorSimpleNoType
// Expecting
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Actual"
func (it *ExpectingRecord) ErrorSimpleNoType(actual interface{}) error {
	return errors.New(it.MessageSimpleNoType(actual))
}
