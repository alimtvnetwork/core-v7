package coretests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type ShouldAsserter interface {
	ShouldBe(
		caseIndex int,
		t *testing.T,
		assert convey.Assertion,
		actual interface{},
	)
	ShouldBeExplicit(
		isValidateType bool,
		caseIndex int,
		t *testing.T,
		title string,
		actual interface{},
		assert convey.Assertion,
		expected interface{},
	)
}
