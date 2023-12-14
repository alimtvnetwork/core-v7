package isanytests

import (
	"gitlab.com/auk-go/core/coretests"
)

type testWrapper struct {
	coretests.BaseTestCase
}

func (it testWrapper) Arrange() []string {
	return it.ArrangeInput.([]string)
}

func (it testWrapper) Expected() []string {
	return it.ExpectedInput.([]string)
}
