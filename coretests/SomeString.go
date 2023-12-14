package coretests

import "fmt"

type SomeString struct {
	Value string
}

func (it SomeString) String() string {
	return it.Value
}

func (it SomeString) AsStringer() fmt.Stringer {
	return it
}
