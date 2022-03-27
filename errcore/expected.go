package errcore

import (
	"errors"
	"fmt"
	"reflect"
)

type expected struct{}

func (it expected) But(
	title interface{},
	wasExpecting,
	actualReceived interface{},
) error {
	return ExpectingErrorSimpleNoType(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButFoundAsMsg(
	title interface{},
	wasExpecting,
	actualReceived interface{},
) string {
	return ExpectingSimpleNoType(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButFoundWithTypeAsMsg(
	title interface{},
	wasExpecting,
	actualReceived interface{},
) string {
	return ExpectingSimple(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButUsingType(
	title interface{},
	wasExpecting,
	actualReceived interface{},
) error {
	return errors.New(ExpectingSimple(
		title,
		wasExpecting,
		actualReceived))
}

func (it expected) ReflectButFound(
	expected, found reflect.Kind,
) error {
	return fmt.Errorf(
		"expected [%v] but found [%v]",
		expected.String(), found.String())
}

func (it expected) PrimitiveButFound(
	found reflect.Kind,
) error {
	return fmt.Errorf(
		"expected [primitive] but found [%v]",
		found.String())
}

func (it expected) ValueHasNoElements(
	typ reflect.Kind,
) error {
	return fmt.Errorf(
		"generic value [%v] is nil or has no element",
		typ.String())
}
