package coretests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type V2ShouldAsserter interface {
	ShouldBeSimpleAsserter
	V2ShouldBeExplicitAsserter
}

type ShouldBeSimpleAsserter interface {
	ShouldBe(
		caseIndex int,
		t *testing.T,
		assert convey.Assertion,
		actual interface{},
	)
}

type ShouldBeEqualAsserter interface {
	ShouldBeEqual(
		caseIndex int,
		t *testing.T,
		actual interface{},
	)
}

type ShouldHaveNoErrorAsserter interface {
	ShouldHaveNoError(
		caseIndex int,
		t *testing.T,
		err error,
	)
}

type ShouldContainsAsserter interface {
	ShouldContains(
		caseIndex int,
		t *testing.T,
		actual interface{},
	)
}

type V2ShouldBeExplicitAsserter interface {
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

type V1ShouldBeExplicitAsserter interface {
	ShouldBeExplicit(
		caseIndex int,
		t *testing.T,
		title string,
		actual interface{},
		assert convey.Assertion,
		expected interface{},
	)
}

type V1ShouldAsserter interface {
	ShouldBeSimpleAsserter
	V1ShouldBeExplicitAsserter
}
