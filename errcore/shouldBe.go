package errcore

import (
	"encoding/json"
	"errors"
	"fmt"
)

type shouldBe struct{}

func (it shouldBe) StrEqMsg(actual, expecting string) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) StrEqErr(actual, expecting string) error {
	msg := it.StrEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) AnyEqMsg(actual, expecting interface{}) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) AnyEqErr(actual, expecting interface{}) error {
	msg := it.AnyEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) JsonEqMsg(actual, expecting interface{}) string {
	actualJson, err := json.Marshal(actual)
	MustBeEmpty(err)

	expectingJson, expectingErr := json.Marshal(expecting)
	MustBeEmpty(expectingErr)

	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actualJson,
		expectingJson)
}

func (it shouldBe) JsonEqErr(actual, expecting interface{}) error {
	msg := it.JsonEqMsg(expecting, actual)

	return errors.New(msg)
}
