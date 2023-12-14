package corerangestests

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
)

type testWrapper struct {
	coretests.BaseTestCase
	IsExpectingError bool
	HasPanic         bool
	Validator        corevalidator.SliceValidator
}

func (it testWrapper) Arrange() []corerange.MinMaxInt {
	return it.ArrangeInput.([]corerange.MinMaxInt)
}

func (it testWrapper) Expected() []int {
	return it.ExpectedInput.([]int)
}
