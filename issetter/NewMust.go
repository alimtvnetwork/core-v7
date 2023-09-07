package issetter

import (
	"gitlab.com/auk-go/core/errcore"
)

func NewMust(name string) Value {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
