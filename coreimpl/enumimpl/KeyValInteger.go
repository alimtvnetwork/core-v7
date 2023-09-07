package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type KeyValInteger struct {
	Key          string
	ValueInteger int
}

func (it KeyValInteger) WrapKey() string {
	return toJsonName(it.Key)
}

func (it KeyValInteger) WrapValue() string {
	return toJsonName(it.ValueInteger)
}

func (it KeyValInteger) KeyAnyVal() KeyAnyVal {
	return KeyAnyVal{
		Key:      it.Key,
		AnyValue: it.ValueInteger,
	}
}

func (it KeyValInteger) IsString() bool {
	return it.ValueInteger == constants.MinInt
}

func (it KeyValInteger) String() string {
	if it.IsString() {
		// stringer
		return fmt.Sprintf(
			constants.StringEnumNameValueFormat,
			it.Key,
		)
	}

	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		it.Key,
		it.ValueInteger)
}
