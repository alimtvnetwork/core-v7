package errcore

import (
	"errors"
	"fmt"
)

type ExpectingRecord struct {
	ExpectingTitle string
	WasExpecting   any
}

// Message
// Expecting
//
// returns
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (it *ExpectingRecord) Message(actual any) string {
	if it == nil {
		return ""
	}

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
func (it *ExpectingRecord) MessageSimple(actual any) string {
	if it == nil {
		return ""
	}

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
func (it *ExpectingRecord) MessageSimpleNoType(actual any) string {
	if it == nil {
		return ""
	}

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
func (it *ExpectingRecord) Error(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.Message(actual))
}

// ErrorSimple
// Expecting
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func (it *ExpectingRecord) ErrorSimple(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.MessageSimple(actual))
}

// ErrorSimpleNoType
// Expecting
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Actual"
func (it *ExpectingRecord) ErrorSimpleNoType(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.MessageSimpleNoType(actual))
}
