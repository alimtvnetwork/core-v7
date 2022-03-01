package errcore

import (
	"encoding/json"
	"errors"
	"fmt"
)

type shouldBe struct{}

func (it shouldBe) StrEqMsg(expecting, actual string) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) StrEqErr(expecting, actual string) error {
	msg := it.StrEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) AnyEqMsg(expecting, actual interface{}) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) AnyEqErr(expecting, actual interface{}) error {
	msg := it.AnyEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) JsonEqMsg(expecting, actual interface{}) string {
	actualJson, err := json.Marshal(actual)
	MustBeEmpty(err)

	expectingJson, expectingErr := json.Marshal(expecting)
	MustBeEmpty(expectingErr)

	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actualJson,
		expectingJson)
}

func (it shouldBe) JsonEqErr(expecting, actual interface{}) error {
	msg := it.JsonEqMsg(expecting, actual)

	return errors.New(msg)
}
