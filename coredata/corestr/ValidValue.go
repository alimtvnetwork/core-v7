package corestr

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

type ValueValid struct {
	Value      string
	valueBytes *[]byte
	IsValid    bool
	Message    string
}

func InvalidValueValidNoMessage() *ValueValid {
	return InvalidValueValid(constants.EmptyString)
}

func InvalidValueValid(message string) *ValueValid {
	return &ValueValid{
		Value:   constants.EmptyString,
		IsValid: false,
		Message: message,
	}
}

func (receiver *ValueValid) ValueBytesOnce() []byte {
	return *receiver.ValueBytesOncePtr()
}

func (receiver *ValueValid) ValueBytesOncePtr() *[]byte {
	if receiver.valueBytes == nil {
		valueBytes := []byte(receiver.Value)

		receiver.valueBytes = &valueBytes
	}

	return receiver.valueBytes
}

func (receiver *ValueValid) IsEmpty() bool {
	return receiver.Value == ""
}

func (receiver *ValueValid) IsWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(receiver.Value)
}

func (receiver *ValueValid) HasValidNonEmpty() bool {
	return receiver.IsValid && !receiver.IsEmpty()
}

func (receiver *ValueValid) HasValidNonWhitespace() bool {
	return receiver.IsValid && !receiver.IsWhitespace()
}

// HasSafeNonEmpty receiver.IsValid &&
//		!receiver.IsLeftEmpty() &&
//		!receiver.IsMiddleEmpty() &&
//		!receiver.IsRightEmpty()
func (receiver *ValueValid) HasSafeNonEmpty() bool {
	return receiver.IsValid &&
		!receiver.IsEmpty()
}

func (receiver *ValueValid) Is(val string) bool {
	return receiver.Value == val
}

// IsAnyOf if length of values are 0 then returns true
func (receiver *ValueValid) IsAnyOf(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if receiver.Value == value {
			return true
		}
	}

	return false
}

func (receiver *ValueValid) IsContains(val string) bool {
	return strings.Contains(receiver.Value, val)
}

// IsAnyContains if length of values are 0 then returns true
func (receiver *ValueValid) IsAnyContains(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if receiver.IsContains(value) {
			return true
		}
	}

	return false
}

func (receiver *ValueValid) IsEqualNonSensitive(val string) bool {
	return strings.EqualFold(receiver.Value, val)
}

func (receiver *ValueValid) IsRegexMatches(regexp *regexp.Regexp) bool {
	if regexp == nil {
		return false
	}

	return regexp.MatchString(receiver.Value)
}

func (receiver *ValueValid) Clone() *ValueValid {
	if receiver == nil {
		return nil
	}

	return &ValueValid{
		Value:   receiver.Value,
		IsValid: receiver.IsValid,
		Message: receiver.Message,
	}
}
