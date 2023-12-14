package coretests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type BaseTestCaseWrapper interface {
	SimpleTestCaseWrapper
	V2ShouldAsserter
	TypeValidationError() error
	TypesValidationMustPasses(t *testing.T)
	ArrangeString() string
	ActualString() string
	SetActual(actual interface{})
	FormTitle(caseIndex int) string
	CustomTitle(caseIndex int, title string) string
	noPrintAssert(
		caseIndex int,
		t *testing.T,
		assert convey.Assertion,
		actual interface{},
	)

	LinesString(caseIndex int) string

	String(caseIndex int) string
	IsDisabled() bool
	AsSimpleTestCaseWrapper() SimpleTestCaseWrapper
}
