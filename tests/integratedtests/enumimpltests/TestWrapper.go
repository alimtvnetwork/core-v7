package enumimpltests

import (
	"gitlab.com/evatix-go/core/coredata/corerange"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl/enumtype"
)

type TestWrapper struct {
	Header         string
	ExpectedMinMax corerange.MinMaxInt64
	EnumMap        enumimpl.DynamicMap
	EnumType       enumtype.Variant
}
