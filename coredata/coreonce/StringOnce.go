package coreonce

import (
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type StringOnce struct {
	innerData       string
	initializerFunc func() string
	isInitialized   bool
}

func NewStringOnce(initializerFunc func() string) StringOnce {
	return StringOnce{
		initializerFunc: initializerFunc,
	}
}

func NewStringOncePtr(initializerFunc func() string) *StringOnce {
	return &StringOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *StringOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *StringOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *StringOnce) ValuePtr() *string {
	val := it.Value()

	return &val
}

func (it *StringOnce) Value() string {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *StringOnce) Execute() string {
	return it.Value()
}

func (it *StringOnce) IsEqual(equalString string) bool {
	return it.Value() == equalString
}

func (it *StringOnce) HasPrefix(prefix string) bool {
	return strings.HasPrefix(
		it.Value(), prefix)
}

func (it *StringOnce) IsStartsWith(startsWith string) bool {
	return strings.HasPrefix(
		it.Value(), startsWith)
}

func (it *StringOnce) HasSuffix(suffix string) bool {
	return strings.HasSuffix(
		it.Value(), suffix)
}

func (it *StringOnce) IsEndsWith(
	endsWith string,
) bool {
	return strings.HasSuffix(
		it.Value(), endsWith)
}

func (it *StringOnce) SplitBy(
	splitter string,
) []string {
	return strings.Split(it.Value(), splitter)
}

func (it *StringOnce) SplitLeftRightTrim(
	splitter string,
) (left, right string) {
	left, right = it.SplitLeftRight(splitter)

	return strings.TrimSpace(left), strings.TrimSpace(right)
}

func (it *StringOnce) SplitLeftRight(
	splitter string,
) (left, right string) {
	items := strings.SplitN(
		it.Value(),
		splitter,
		constants.Two)

	if len(items) == 2 {
		return items[0], items[1]
	}

	if len(items) > 2 {
		return items[0], items[len(items)]
	}

	// len <= 1
	return items[0], ""
}

func (it *StringOnce) IsContains(equalString string) bool {
	return strings.Contains(it.Value(), equalString)
}

func (it *StringOnce) IsEmpty() bool {
	return it.Value() == ""
}

func (it *StringOnce) IsEmptyOrWhitespace() bool {
	return strings.TrimSpace(it.Value()) == ""
}

func (it *StringOnce) Bytes() []byte {
	return []byte(it.Value())
}

func (it *StringOnce) Error() error {
	return errors.New(it.Value())
}

func (it *StringOnce) String() string {
	return it.Value()
}

func (it *StringOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
