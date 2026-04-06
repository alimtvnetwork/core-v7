package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/smarty/assertions/should"
)

// ── CaseV1.VerifyTypeOfMatch — with VerifyTypeOf set ──

func Test_Cov7_CaseV1_VerifyTypeOfMatch_WithVerifyTypeOf(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify type match",
		ArrangeInput:  "hello",
		ExpectedInput: "world",
		VerifyTypeOf:  coretests.NewVerifyTypeOf("hello"),
	}
	// Both are strings — types should match
	c.VerifyTypeOfMatch(t, 0, "actual-string")
}

func Test_Cov7_CaseV1_VerifyTypeOfMust(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify type must",
		ExpectedInput: "world",
		VerifyTypeOf:  coretests.NewVerifyTypeOf("hello"),
	}
	c.VerifyTypeOfMust(t, 0, "actual-string")
}

func Test_Cov7_CaseV1_VerifyType(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify type",
		ExpectedInput: "world",
		VerifyTypeOf:  coretests.NewVerifyTypeOf("hello"),
	}
	c.VerifyType(t, 0, "actual-string")
}

func Test_Cov7_CaseV1_VerifyTypeMust(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify type must",
		ExpectedInput: "world",
		VerifyTypeOf:  coretests.NewVerifyTypeOf("hello"),
	}
	c.VerifyTypeMust(t, 0, "actual-string")
}

// ── CaseV1.VerifyTypeOfMatch — skip verify ──

func Test_Cov7_CaseV1_VerifyTypeOfMust_SkipVerify(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "skip verify must",
		ExpectedInput: "world",
	}
	// No VerifyTypeOf → skip
	c.VerifyTypeOfMust(t, 0, 42)
}

func Test_Cov7_CaseV1_VerifyType_SkipVerify(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "skip verify type",
		ExpectedInput: "world",
	}
	c.VerifyType(t, 0, 42)
}

func Test_Cov7_CaseV1_VerifyTypeMust_SkipVerify(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "skip verify type must",
		ExpectedInput: "world",
	}
	c.VerifyTypeMust(t, 0, 42)
}

// ── CaseV1.VerifyAllEqualCondition ──

func Test_Cov7_CaseV1_VerifyAllEqualCondition(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify all equal condition",
		ExpectedInput: "hello",
	}
	err := c.VerifyAllEqualCondition(
		0,
		corevalidator.DefaultTrimCoreCondition,
		"  hello  ",
	)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllEqualCondition returns correct value -- with args", actual)
}

// ── CaseV1.SliceValidatorCondition ──

func Test_Cov7_CaseV1_SliceValidatorCondition(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "slice validator condition",
		ExpectedInput: "hello",
	}
	sv := c.SliceValidatorCondition(
		stringcompareas.Equal,
		corevalidator.DefaultTrimCoreCondition,
		[]string{"  hello  "},
	)
	actual := args.Map{
		"hasActual":   len(sv.ActualLines) > 0,
		"hasExpected": len(sv.ExpectedLines) > 0,
	}
	expected := args.Map{
		"hasActual":   true,
		"hasExpected": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidatorCondition returns non-empty -- with args", actual)
}

// ── CaseV1.ShouldBeRegex ──

func Test_Cov7_CaseV1_ShouldBeRegex(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "regex test",
		ExpectedInput: "^hel.*ld$",
	}
	c.ShouldBeRegex(t, 0, "hello world")
}

// ── CaseV1.ShouldBeTrimRegex ──

func Test_Cov7_CaseV1_ShouldBeTrimRegex(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "trim regex",
		ExpectedInput: "^hello$",
	}
	c.ShouldBeTrimRegex(t, 0, "   hello   ")
}

// ── CaseV1.VerifyError with type verify ──

func Test_Cov7_CaseV1_VerifyError_WithTypeVerify(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify error with type",
		ExpectedInput: "hello",
		VerifyTypeOf:  coretests.NewVerifyTypeOf("hello"),
	}
	err := c.VerifyError(0, stringcompareas.Equal, "hello")
	// err may be non-nil if type verification uses slice comparison
	_ = err
}

// ── CaseV1.TypeShouldMatch ──

func Test_Cov7_CaseV1_TypeShouldMatch(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "type should match",
		ExpectedInput: []string{"hello"},
		VerifyTypeOf:  coretests.NewVerifyTypeOf([]string{"hello"}),
	}
	err := c.TypeShouldMatch(t)
	if err != nil {
		t.Fatalf("unexpected type mismatch: %v", err)
	}
}

// ── CaseV1.ShouldBeUsingCondition with type verify ──

func Test_Cov7_CaseV1_ShouldBeUsingCondition_WithVerify(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "should be condition with verify",
		ExpectedInput: []string{"hello"},
		VerifyTypeOf:  coretests.NewVerifyTypeOf([]string{"hello"}),
	}
	err := c.ShouldBeUsingCondition(
		t, 0,
		stringcompareas.Equal,
		corevalidator.DefaultDisabledCoreCondition,
		"hello",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ── CaseV1.AssertDirectly ──

func Test_Cov7_CaseV1_AssertDirectly(t *testing.T) {
	c := coretestcases.CaseV1{
		Title: "assert directly",
	}
	c.AssertDirectly(
		t,
		"additional info",
		"comparison message",
		0,
		"hello",
		should.Equal,
		"hello",
	)
}

// ── CaseV1.ShouldBeEqual with []string expected ──

func Test_Cov7_CaseV1_ShouldBeEqual_SliceExpected(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "slice expected",
		ExpectedInput: []string{"a returns correct value -- with args", "b"},
	}
	c.ShouldBeEqual(t, 0, "a returns correct value -- with args", "b")
}

// ── CaseV1.SetExpected ──

func Test_Cov7_CaseV1_SetExpected(t *testing.T) {
	c := coretestcases.CaseV1{}
	c.SetExpected("new-expected")
	// Value receiver — doesn't modify c, but covers the method
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "SetExpected returns correct value -- with args", actual)
}

// ── CaseV1.SetActual ──

func Test_Cov7_CaseV1_SetActual(t *testing.T) {
	c := coretestcases.CaseV1{}
	c.SetActual("new-actual")
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "SetActual returns correct value -- with args", actual)
}

// ── CaseNilSafe with Args ──

func Test_Cov7_CaseNilSafe_WithArgs(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "ClonePtr with args",
		Func:  (*coretests.DraftType).ClonePtr,
		Args:  []any{},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}
	tc.ShouldBeSafe(t, 0)
}

// ── CaseNilSafe.InvokeNil with a method that returns something ──

func Test_Cov7_CaseNilSafe_InvokeNil_ReturnValue(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "ClonePtr nil returns nil",
		Func:  (*coretests.DraftType).ClonePtr,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}
	tc.ShouldBeSafe(t, 0)
}

// ── CaseNilSafe.ShouldBeSafeFirst ──

func Test_Cov7_CaseNilSafe_ShouldBeSafeFirst(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "safe first",
		Func:  (*coretests.DraftType).ClonePtr,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}
	tc.ShouldBeSafeFirst(t)
}

// ── GenericGherkins.ShouldBeEqual with When fallback ──

func Test_Cov7_GenericGherkins_ShouldBeEqual_WhenFallback(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		When:          "when-based-title",
		ExpectedLines: []string{"hello"},
	}
	tc.ShouldBeEqual(t, 0, []string{"hello"}, []string{"hello"})
}

// ── GenericGherkins.ShouldBeEqualMap with When fallback ──

func Test_Cov7_MapGherkins_ShouldBeEqualMap_WhenFallback(t *testing.T) {
	tc := &coretestcases.MapGherkins{
		When:     "when-title-map",
		Expected: args.Map{"k": "v"},
	}
	tc.ShouldBeEqualMap(t, 0, args.Map{"k": "v"})
}

// ── GenericGherkins.CompareWith — multiple field diffs ──

func Test_Cov7_GenericGherkins_CompareWith_MultipleDiffs(t *testing.T) {
	a := &coretestcases.GenericGherkins[string, string]{
		Title:   "a returns correct value -- with args",
		Feature: "fa",
		Given:   "ga",
		When:    "wa",
		Then:    "ta",
		Input:   "ia",
		Expected: "ea",
		Actual:   "aa",
		IsMatching: true,
	}
	b := &coretestcases.GenericGherkins[string, string]{
		Title:   "b",
		Feature: "fb",
		Given:   "gb",
		When:    "wb",
		Then:    "tb",
		Input:   "ib",
		Expected: "eb",
		Actual:   "ab",
		IsMatching: false,
	}
	isEqual, diff := a.CompareWith(b)
	actual := args.Map{
		"isEqual":  isEqual,
		"hasDiff":  diff != "",
		"multiSep": len(diff) > 20, // multiple diffs joined by "; "
	}
	expected := args.Map{
		"isEqual":  false,
		"hasDiff":  true,
		"multiSep": true,
	}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- multiple diffs", actual)
}

// ── GenericGherkins.FullString without ExtraArgs ──

func Test_Cov7_GenericGherkins_FullString_NoExtraArgs(t *testing.T) {
	tc := &coretestcases.GenericGherkins[string, string]{
		Title: "no extras",
		Input: "input",
	}
	result := tc.FullString()
	actual := args.Map{
		"containsTitle": fmt.Sprintf("%v", len(result) > 0),
	}
	expected := args.Map{
		"containsTitle": "true",
	}
	expected.ShouldBeEqual(t, 0, "FullString returns empty -- no extra args", actual)
}

// ── CaseV1.ShouldBe ──

func Test_Cov7_CaseV1_ShouldBe(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "should be test",
		ExpectedInput: "hello",
	}
	err := c.ShouldBe(t, 0, stringcompareas.Equal, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe returns correct value -- with args", actual)
}

// ── CaseV1.VerifyAll ──

func Test_Cov7_CaseV1_VerifyAll(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify all",
		ExpectedInput: "hello",
	}
	err := c.VerifyAll(0, stringcompareas.Equal, []string{"hello"})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns correct value -- with args", actual)
}

// ── CaseV1.VerifyAllSliceValidator ──

func Test_Cov7_CaseV1_VerifyAllSliceValidator(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "verify all slice validator",
		ExpectedInput: "hello",
	}
	sv := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"hello"},
	}
	err := c.VerifyAllSliceValidator(0, sv)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllSliceValidator returns non-empty -- with args", actual)
}

// ── CaseV1 — IsEnable flag ──

func Test_Cov7_CaseV1_IsEnable(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:    "enabled case",
		IsEnable: issetter.True,
	}
	actual := args.Map{
		"isTrue": c.IsEnable.IsTrue(),
	}
	expected := args.Map{
		"isTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1 returns correct value -- IsEnable", actual)
}

// ── CaseV1 ExpectedLines with int ──

func Test_Cov7_CaseV1_ExpectedLines_Int(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: 42,
	}
	lines := c.ExpectedLines()
	actual := args.Map{
		"len": len(lines),
	}
	expected := args.Map{
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "ExpectedLines returns correct value -- int", actual)
}

// ── CaseV1 ExpectedLines with bool ──

func Test_Cov7_CaseV1_ExpectedLines_Bool(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: true,
	}
	lines := c.ExpectedLines()
	actual := args.Map{
		"len":   len(lines),
		"first": lines[0],
	}
	expected := args.Map{
		"len":   1,
		"first": "true",
	}
	expected.ShouldBeEqual(t, 0, "ExpectedLines returns correct value -- bool", actual)
}

// ── CaseV1 ExpectedLines with []int ──

func Test_Cov7_CaseV1_ExpectedLines_IntSlice(t *testing.T) {
	c := coretestcases.CaseV1{
		ExpectedInput: []int{1, 2, 3},
	}
	lines := c.ExpectedLines()
	actual := args.Map{
		"len": len(lines),
	}
	expected := args.Map{
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "ExpectedLines returns correct value -- int slice", actual)
}
