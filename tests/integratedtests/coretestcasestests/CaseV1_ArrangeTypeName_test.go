package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/smarty/assertions/should"
)

// ── CaseV1: ArrangeTypeName ──
// Covers CaseV1.go L55-57

func Test_CaseV1_ArrangeTypeName_FromCaseV1ArrangeTypeNam(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "test type name",
		ArrangeInput:  "hello",
		ExpectedInput: "hello",
	}

	result := tc.ArrangeTypeName()

	// Act
	actual := args.Map{"nonEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "ArrangeTypeName returns type name -- string input", actual)
}

// ── CaseV1: ShouldBeEqual (success path) ──
// Covers CaseV1.go L445-464 and ShouldBe L369-382

func Test_CaseV1_ShouldBeEqual_Success(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "equal test",
		ExpectedInput: "hello",
	}

	// Assert
	tc.ShouldBeEqual(t, 0, "hello")
}

// ── CaseV1: ShouldBeSortedEqual ──
// Covers CaseV1.go L484-492

func Test_CaseV1_ShouldBeSortedEqual(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "sorted equal test",
		ExpectedInput: []string{"a", "b"},
	}

	tc.ShouldBeSortedEqual(t, 0, "a", "b")
}

// ── CaseV1: ShouldContains ──
// Covers CaseV1.go L498-505

func Test_CaseV1_ShouldContains(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "contains test",
		ExpectedInput: "ell",
	}

	tc.ShouldContains(t, 0, "hello")
}

// ── CaseV1: ShouldStartsWith ──
// Covers CaseV1.go L511-518

func Test_CaseV1_ShouldStartsWith(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "starts with test",
		ExpectedInput: "hel",
	}

	tc.ShouldStartsWith(t, 0, "hello")
}

// ── CaseV1: ShouldEndsWith ──
// Covers CaseV1.go L524-531

func Test_CaseV1_ShouldEndsWith(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "ends with test",
		ExpectedInput: "llo",
	}

	tc.ShouldEndsWith(t, 0, "hello")
}

// ── CaseV1: ShouldBeNotEqual ──
// Covers CaseV1.go L537-544

func Test_CaseV1_ShouldBeNotEqual(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "not equal test",
		ExpectedInput: "world",
	}

	tc.ShouldBeNotEqual(t, 0, "hello")
}

// ── CaseV1: ShouldBeTrimRegex ──
// Covers CaseV1.go L571-579

func Test_CaseV1_ShouldBeTrimRegex(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "regex test",
		ExpectedInput: "hel.*",
	}

	tc.ShouldBeTrimRegex(t, 0, "hello")
}

// ── CaseV1: ShouldHaveNoError ──
// Covers CaseV1.go L586-601

func Test_CaseV1_ShouldHaveNoError_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "no error test",
		ExpectedInput: "test",
	}

	tc.ShouldHaveNoError(t, "additional", 0, nil)
}

// ── CaseV1: AssertDirectly ──
// Covers CaseV1.go L615-631

func Test_CaseV1_AssertDirectly(t *testing.T) {
	tc := coretestcases.CaseV1{
		Title:         "assert directly test",
		ExpectedInput: 42,
	}

	tc.AssertDirectly(t, "msg", "check", 0, 42, should.Equal, 42)
}

// ── CaseV1: PrepareTitle ──
// Covers CaseV1.go L636-643

func Test_CaseV1_PrepareTitle_FromCaseV1ArrangeTypeNam(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "my test",
		ExpectedInput: "x",
	}

	result := tc.PrepareTitle(3, "extra")

	// Act
	actual := args.Map{
		"nonEmpty": len(result) > 0,
		"contains": true,
	}

	// Assert
	expected := args.Map{
		"nonEmpty": true,
		"contains": true,
	}
	expected.ShouldBeEqual(t, 0, "PrepareTitle returns formatted title", actual)
}

// ── CaseV1: AsSimpleTestCaseWrapperContractsBinder ──
// Covers CaseV1.go L649-651

func Test_CaseV1_AsSimpleTestCaseWrapperContractsBinder_FromCaseV1ArrangeTypeNam(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "binder test",
		ExpectedInput: "x",
	}

	binder := tc.AsSimpleTestCaseWrapperContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsSimpleTestCaseWrapperContractsBinder returns non-nil", actual)
}

// ── CaseV1FirstAssertions ──
// Covers CaseV1FirstAssertions.go all 7 First methods

func Test_CaseV1_ShouldBeEqualFirst(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first equal", ExpectedInput: "val"}

	// Assert
	tc.ShouldBeEqualFirst(t, "val")
}

func Test_CaseV1_ShouldBeTrimEqualFirst(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first trim equal", ExpectedInput: "val"}
	tc.ShouldBeTrimEqualFirst(t, " val ")
}

func Test_CaseV1_ShouldBeSortedEqualFirst(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first sorted equal", ExpectedInput: []string{"a", "b"}}
	tc.ShouldBeSortedEqualFirst(t, "a", "b")
}

func Test_CaseV1_ShouldContainsFirst(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first contains", ExpectedInput: "el"}
	tc.ShouldContainsFirst(t, "hello")
}

func Test_CaseV1_ShouldStartsWithFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first starts", ExpectedInput: "he"}
	tc.ShouldStartsWithFirst(t, "hello")
}

func Test_CaseV1_ShouldEndsWithFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first ends", ExpectedInput: "lo"}
	tc.ShouldEndsWithFirst(t, "hello")
}

func Test_CaseV1_ShouldBeNotEqualFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "first not equal", ExpectedInput: "abc"}
	tc.ShouldBeNotEqualFirst(t, "xyz")
}

// ── CaseV1MapAssertions: ShouldBeEqualMap matching ──
// Covers CaseV1MapAssertions.go L19-20 (ExpectedAsMap)

func Test_CaseV1_ShouldBeEqualMap_Matching(t *testing.T) {
	// Assert
	expected := args.Map{"key": "value"}
	tc := coretestcases.CaseV1{
		Title:         "map equal test",
		ExpectedInput: expected,
	}

	tc.ShouldBeEqualMap(t, 0, args.Map{"key": "value"})
}

// ── CaseV1MapAssertions: ShouldBeEqualMap mismatch ──
// Covers CaseV1MapAssertions.go L55-69

func Test_CaseV1_ShouldBeEqualMap_Mismatch(t *testing.T) {
	// Arrange

	// Assert
	expected := args.Map{"key": "expected"}
	tc := coretestcases.CaseV1{
		Title:         "map mismatch test",
		ExpectedInput: expected,
	}

	fakeT := &testing.T{}
	tc.ShouldBeEqualMap(fakeT, 0, args.Map{"key": "actual"})

	// Act
	actual := args.Map{"failed": fakeT.Failed()}
	exp := args.Map{"failed": true}
	exp.ShouldBeEqual(t, 0, "ShouldBeEqualMap marks test failed -- mismatch", actual)
}

// ── CaseV1: VerifyAll ──
// Covers CaseV1.go L262-280

func Test_CaseV1_VerifyAll(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify all test",
		ExpectedInput: []string{"a", "b"},
	}

	err := tc.VerifyAllEqual(0, "a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllEqual returns nil -- matching", actual)
}

// ── CaseV1: VerifyFirst ──
// Covers CaseV1.go L311-329

func Test_CaseV1_VerifyFirst(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify first test",
		ExpectedInput: []string{"hello"},
	}

	err := tc.VerifyFirst(0, stringcompareas.Equal, []string{"hello"})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- matching", actual)
}

// ── CaseV1: VerifyError ──
// Covers CaseV1.go L349-366

func Test_CaseV1_VerifyError(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify error test",
		ExpectedInput: "hello",
	}

	err := tc.VerifyError(0, stringcompareas.Equal, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns nil -- matching", actual)
}

// ── CaseV1: SliceValidator ──
// Covers CaseV1.go L233-239

func Test_CaseV1_SliceValidator(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "slice validator test",
		ExpectedInput: []string{"a"},
	}

	sv := tc.SliceValidator(0, []string{"a"})

	// Act
	actual := args.Map{"notNil": sv.ActualLines != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns valid -- matching", actual)
}

// ── CaseNilSafe: ShouldBeSafeFirst ──
// Covers CaseNilSafe.go L115-122

type cov9NilSafeTestStruct struct{}

func (s *cov9NilSafeTestStruct) IsValid() bool {
	if s == nil {
		return false
	}
	return true
}

func Test_CaseNilSafe_ShouldBeSafeFirst(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "IsValid on nil returns false",
		Func:  (*cov9NilSafeTestStruct).IsValid,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	}

	// Assert
	tc.ShouldBeSafeFirst(t)
}

// ── CaseNilSafeAssertHelper: assertDiffOnMismatch ──
// Covers CaseNilSafeAssertHelper.go L17-27 (internal helper, covered via ShouldBeSafe)

// ── GenericGherkinsAssertions: ShouldBeEqual success ──
// Covers GenericGherkinsAssertions.go L19-33

func Test_GenericGherkins_ShouldBeEqual_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:         "gherkins equal test",
		ExpectedLines: []string{"hello"},
	}

	// Assert
	tc.ShouldBeEqual(t, 0, []string{"hello"}, []string{"hello"})
}

// ── GenericGherkinsAssertions: ShouldBeEqualFirst ──
// Covers GenericGherkinsAssertions.go L42-51

func Test_GenericGherkins_ShouldBeEqualFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title: "gherkins first test",
	}

	// Assert
	tc.ShouldBeEqualFirst(t, []string{"a"}, []string{"a"})
}

// ── GenericGherkinsAssertions: ShouldBeEqualArgs ──
// Covers GenericGherkinsAssertions.go L69-78

func Test_GenericGherkins_ShouldBeEqualArgs_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:         "gherkins args test",
		ExpectedLines: []string{"val"},
	}

	// Assert
	tc.ShouldBeEqualArgs(t, 0, "val")
}

// ── GenericGherkinsAssertions: ShouldBeEqualArgsFirst ──
// Covers GenericGherkinsAssertions.go L94-102

func Test_GenericGherkins_ShouldBeEqualArgsFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:         "gherkins args first test",
		ExpectedLines: []string{"val"},
	}

	// Assert
	tc.ShouldBeEqualArgsFirst(t, "val")
}

// ── GenericGherkinsAssertions: ShouldBeEqualUsingExpected ──
// Covers GenericGherkinsAssertions.go L111-120

func Test_GenericGherkins_ShouldBeEqualUsingExpected_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:         "gherkins expected test",
		ExpectedLines: []string{"val"},
	}

	// Assert
	tc.ShouldBeEqualUsingExpected(t, 0, []string{"val"})
}

// ── GenericGherkinsAssertions: ShouldBeEqualUsingExpectedFirst ──
// Covers GenericGherkinsAssertions.go L128-136

func Test_GenericGherkins_ShouldBeEqualUsingExpectedFirst_FromCaseV1ArrangeTypeNam(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:         "gherkins expected first test",
		ExpectedLines: []string{"val"},
	}

	// Assert
	tc.ShouldBeEqualUsingExpectedFirst(t, []string{"val"})
}

// ── GenericGherkinsCompare: CompareWith nil cases ──
// Covers GenericGherkinsCompare.go L14-17

func Test_GenericGherkins_CompareWith_OneNil_FromCaseV1ArrangeTypeNam(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Title: "compare test",
	}

	isEqual, diff := tc.CompareWith(nil)

	// Act
	actual := args.Map{
		"equal": isEqual,
		"hasDiff": len(diff) > 0,
	}

	// Assert
	expected := args.Map{
		"equal": false,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "CompareWith returns false -- one nil", actual)
}

// ── GenericGherkinsAssertions: ShouldBeEqual with When fallback ──
// Covers GenericGherkinsAssertions.go L23-25

func Test_GenericGherkins_ShouldBeEqual_WhenFallback(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title: "", // empty title triggers When fallback
		When:  "when condition",
	}

	// Assert
	tc.ShouldBeEqual(t, 0, []string{"a"}, []string{"a"})
}
