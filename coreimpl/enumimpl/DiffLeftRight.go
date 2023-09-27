package enumimpl

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

type DiffLeftRight struct {
	Left, Right interface{}
}

func (it *DiffLeftRight) Types() (l, r reflect.Type) {
	l = reflect.TypeOf(it.Left)
	r = reflect.TypeOf(it.Right)

	return l, r
}

func (it *DiffLeftRight) IsSameTypeSame() bool {
	l, r := it.Types()

	return l == r
}

func (it *DiffLeftRight) IsSame() bool {
	return reflect.DeepEqual(it.Left, it.Right)
}

func (it *DiffLeftRight) IsSameRegardlessOfType() bool {
	leftString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Left)
	rightString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Right)

	return leftString == rightString
}

func (it *DiffLeftRight) IsEqual(isRegardless bool) bool {
	if isRegardless {
		return it.IsSameRegardlessOfType()
	}

	return it.IsSame()
}

func (it *DiffLeftRight) HasMismatch(isRegardless bool) bool {
	if isRegardless {
		return it.HasMismatchRegardlessOfType()
	}

	return it.IsNotEqual()
}

func (it *DiffLeftRight) IsNotEqual() bool {
	return !it.IsSame()
}

func (it *DiffLeftRight) HasMismatchRegardlessOfType() bool {
	return !it.IsSameRegardlessOfType()
}

func (it *DiffLeftRight) String() string {
	return it.JsonString()
}

func (it *DiffLeftRight) JsonString() string {
	if it == nil {
		return ""
	}

	b, e := json.Marshal(*it)

	if e != nil {
		return "error : " + e.Error()
	}

	return string(b)
}

func (it *DiffLeftRight) SpecificFullString() (l, r string) {
	l = fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Left)
	r = fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Right)

	return l, r
}

func (it *DiffLeftRight) DiffString() string {
	if it.IsSameRegardlessOfType() {
		return ""
	}

	return it.String()
}
