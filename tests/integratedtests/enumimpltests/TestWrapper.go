package enumimpltests

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
)

type TestWrapper struct {
	Header         string
	ExpectedMinMax corerange.MinMaxInt64
	EnumMap        enumimpl.DynamicMap
	EnumType       enumtype.Variant
}
