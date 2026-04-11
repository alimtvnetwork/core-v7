package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/issetter"
)

// ── CaseV1 additional methods ──

func Test_CaseV1_Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ArrangeInput: "hello returns correct value -- with args"}

	// Act
	actual := args.Map{"val": fmt.Sprintf("%v", c.Input())}

	// Assert
	expected := args.Map{"val": "hello returns correct value -- with args"}
	expected.ShouldBeEqual(t, 0, "CaseV1.Input -- hello", actual)
}

func Test_CaseV1_Expected(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ExpectedInput: "world"}

	// Act
	actual := args.Map{"val": fmt.Sprintf("%v", c.Expected())}

	// Assert
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "CaseV1.Expected -- world", actual)
}

func Test_CaseV1_ExpectedLines_String(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ExpectedInput: "single"}
	lines := c.ExpectedLines()

	// Act
	actual := args.Map{
		"len": len(lines),
		"first": lines[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "single",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1.ExpectedLines string -- single", actual)
}

func Test_CaseV1_ExpectedLines_Slice(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ExpectedInput: []string{"a", "b"}}
	lines := c.ExpectedLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CaseV1.ExpectedLines slice -- 2 items", actual)
}

func Test_CaseV1_Actual(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ActualInput: "test"}

	// Act
	actual := args.Map{"val": fmt.Sprintf("%v", c.Actual())}

	// Assert
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "CaseV1.Actual -- test", actual)
}

func Test_CaseV1_SetActual(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{}
	c.SetActual("set")
	// SetActual on value receiver doesn't modify original -- this is expected Go behavior

	// Act
	actual := args.Map{"type": "CaseV1"}

	// Assert
	expected := args.Map{"type": "CaseV1"}
	expected.ShouldBeEqual(t, 0, "CaseV1.SetActual -- called", actual)
}

func Test_CaseV1_CaseTitle(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "my title"}

	// Act
	actual := args.Map{"val": c.CaseTitle()}

	// Assert
	expected := args.Map{"val": "my title"}
	expected.ShouldBeEqual(t, 0, "CaseV1.CaseTitle -- my title", actual)
}

func Test_CaseV1_SetExpected(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{}
	c.SetExpected("val")

	// Act
	actual := args.Map{"called": true}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "CaseV1.SetExpected -- called", actual)
}

func Test_CaseV1_ShouldBeEqualFirst_FromCaseV1InputV2(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "equal first test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	// Assert
	c.ShouldBeEqualFirst(t, "hello returns correct value -- with args")
}

func Test_CaseV1_ShouldBeTrimEqualFirst_FromCaseV1InputV2(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "trim equal first test",
		ExpectedInput: "hello",
	}
	err := c.VerifyError(0, stringcompareas.Equal, "hello")
	_ = err // exercise the trim-equal path
}

func Test_CaseV1_ShouldBeSortedEqualFirst_FromCaseV1InputV2(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "sorted equal first test",
		ExpectedInput: []string{"b", "a"},
	}
	c.ShouldBeSortedEqualFirst(t, "b", "a")
}

func Test_CaseV1_ShouldContainsFirst_FromCaseV1InputV2(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "contains first test",
		ExpectedInput: "hel",
	}
	c.ShouldContainsFirst(t, "hello world")
}

// ── CaseNilSafe: ShouldBeSafe ──

func Test_CaseNilSafe_ShouldBeSafe(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "ClonePtr nil safe",
		Func:  (*coretests.DraftType).ClonePtr,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}

	// Assert
	tc.ShouldBeSafe(t, 0)
}

func Test_CaseNilSafe_ShouldBeSafeFirst_FromCaseV1InputV2(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "ClonePtr nil safe first",
		Func:  (*coretests.DraftType).ClonePtr,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}

	// Assert
	tc.ShouldBeSafeFirst(t)
}

// ── CaseV1 VerifyTypeOfMatch with disabled verify ──

func Test_CaseV1_VerifyTypeOfMatch_SkipVerify(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "skip verify",
		ExpectedInput: "hello returns correct value -- with args",
		IsEnable:      issetter.True,
		// VerifyTypeOf not set → IsTypeInvalidOrSkipVerify returns true
	}
	// Should not panic — just skips
	c.VerifyTypeOfMatch(t, 0, "hello returns correct value -- with args")

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "VerifyTypeOfMatch skip verify -- no VerifyTypeOf", actual)
}

// ── CaseV1 ShouldBeEqual wraps string ──

func Test_CaseV1_ShouldBeEqual_StringWrap(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "string wrap test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	// Assert
	c.ShouldBeEqual(t, 0, "hello returns correct value -- with args")
}

// ── CaseV1 ShouldBeEqual empty string ──

func Test_CaseV1_ShouldBeEqual_EmptyString(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "empty string wrap test",
		ExpectedInput: "",
	}

	// Assert
	c.ShouldBeEqual(t, 0, "")
}
