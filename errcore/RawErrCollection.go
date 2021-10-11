package errcore

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type RawErrCollection struct {
	Items []error
}

func (it *RawErrCollection) Add(err error) {
	if err == nil {
		return
	}

	it.Items = append(it.Items, err)
}

func (it *RawErrCollection) AddWithTraceRef(
	err error,
	traces []string,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	it.Items = append(
		it.Items,
		ErrorWithTracesRefToError(err, traces, referenceItem))
}

func (it *RawErrCollection) AddWithCompiledTraceRef(
	err error,
	compiledTrace string,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	compiledErr := ErrorWithCompiledTraceRefToError(
		err,
		compiledTrace,
		referenceItem)

	it.Items = append(
		it.Items,
		compiledErr)
}

func (it *RawErrCollection) AddWithRef(
	err error,
	referenceItem interface{},
) {
	if err == nil {
		return
	}

	compiledErr := ErrorWithRefToError(
		err,
		referenceItem)

	it.Items = append(
		it.Items,
		compiledErr)
}

func (it *RawErrCollection) Adds(errorItems ...error) {
	if len(errorItems) == 0 {
		return
	}

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		it.Items = append(it.Items, err)
	}
}

func (it *RawErrCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *RawErrCollection) IsEmpty() bool {
	return it == nil || len(it.Items) == 0
}

func (it *RawErrCollection) HasError() bool {
	return it != nil && len(it.Items) > 0
}

func (it *RawErrCollection) Clear() {
	if it.IsEmpty() {
		return
	}

	tempItems := it.Items
	clearFunc := func() {
		for i := 0; i < len(tempItems); i++ {
			tempItems[i] = nil
		}
	}

	go clearFunc()
	it.Items = []error{}
}

func (it *RawErrCollection) Dispose() {
	if it.IsEmpty() {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it RawErrCollection) Strings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, err := range it.Items {
		slice[i] = err.Error()
	}

	return slice
}

func (it RawErrCollection) StringUsingJoiner(joiner string) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(
		it.Strings(),
		joiner)
}

func (it RawErrCollection) StringUsingJoinerAdditional(joiner, additionalMessage string) string {
	if it.IsEmpty() {
		return ""
	}

	return strings.Join(
		it.Strings(),
		joiner) + additionalMessage
}

func (it RawErrCollection) String() string {
	if it.IsEmpty() {
		return ""
	}

	return it.StringUsingJoiner(constants.NewLineUnix)
}

func (it RawErrCollection) CompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.String()

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingJoiner(joiner string) error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.StringUsingJoiner(joiner)

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingJoinerAdditionalMessage(joiner, additionalMessage string) error {
	if it.IsEmpty() {
		return nil
	}

	toString := it.StringUsingJoinerAdditional(
		joiner,
		additionalMessage)

	return errors.New(toString)
}

func (it RawErrCollection) CompiledErrorUsingStackTraces(joiner string, stackTraces []string) error {
	if it.IsEmpty() {
		return nil
	}

	return ErrorWithTracesRefToError(
		it.CompiledErrorUsingJoiner(joiner),
		stackTraces,
		nil)
}

func (it RawErrCollection) StringWithAdditionalMessage(additionalMessage string) string {
	if it.IsEmpty() {
		return ""
	}

	return it.StringUsingJoiner(constants.NewLineUnix) + additionalMessage
}
