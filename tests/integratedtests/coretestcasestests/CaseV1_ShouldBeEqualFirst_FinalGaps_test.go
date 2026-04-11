package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage11 — coretests/coretestcases final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- CaseV1FirstAssertions ---

func Test_CaseV1_ShouldBeEqualFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "first assertion test",
		ExpectedInput: "hello",
	}

	// Act & Assert (no panic = success)
	tc.ShouldBeEqualFirst(t, "hello")
}

func Test_CaseV1_ShouldBeTrimEqualFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "trim equal first test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.ShouldBeTrimEqualFirst(t, "hello")
}

func Test_CaseV1_ShouldBeSortedEqualFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "sorted equal first test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.ShouldBeSortedEqualFirst(t, "hello")
}

func Test_CaseV1_ShouldContainsFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "contains first test",
		ExpectedInput: "hell",
	}

	// Act & Assert
	tc.ShouldContainsFirst(t, "hello world")
}

func Test_CaseV1_ShouldStartsWithFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "starts with first test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.ShouldStartsWithFirst(t, "hello world")
}

func Test_CaseV1_ShouldEndsWithFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "ends with first test",
		ExpectedInput: "world",
	}

	// Act & Assert
	tc.ShouldEndsWithFirst(t, "hello world")
}

func Test_CaseV1_ShouldBeNotEqualFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "not equal first test",
		ExpectedInput: "different",
	}

	// Act & Assert
	tc.ShouldBeNotEqualFirst(t, "hello")
}

func Test_CaseV1_ShouldBeRegexFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "regex first test",
		ExpectedInput: "hel.*",
	}

	// Act & Assert
	tc.ShouldBeRegexFirst(t, "hello")
}

// --- CaseV1 other methods ---

func Test_CaseV1_Input_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		ArrangeInput: "test-input",
	}

	// Act & Assert
	convey.Convey("CaseV1.Input returns ArrangeInput", t, func() {
		convey.So(tc.Input(), convey.ShouldEqual, "test-input")
	})
}

func Test_CaseV1_Actual_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		ActualInput: "actual-val",
	}

	// Act & Assert
	convey.Convey("CaseV1.Actual returns ActualInput", t, func() {
		convey.So(tc.Actual(), convey.ShouldEqual, "actual-val")
	})
}

func Test_CaseV1_ArrangeTypeName_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		ArrangeInput: "hello",
	}

	// Act
	result := tc.ArrangeTypeName()

	// Assert
	convey.Convey("CaseV1.ArrangeTypeName returns type name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_CaseV1_CaseTitle_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "my-title",
	}

	// Act & Assert
	convey.Convey("CaseV1.CaseTitle returns Title", t, func() {
		convey.So(tc.CaseTitle(), convey.ShouldEqual, "my-title")
	})
}

func Test_CaseV1_PrepareTitle_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "base",
	}

	// Act
	result := tc.PrepareTitle(0, "extra")

	// Assert
	convey.Convey("CaseV1.PrepareTitle formats title", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "base")
		convey.So(result, convey.ShouldContainSubstring, "extra")
	})
}

func Test_CaseV1_AsBaseTestCase_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "test",
	}

	// Act
	base := tc.AsBaseTestCase()

	// Assert
	convey.Convey("CaseV1.AsBaseTestCase returns base", t, func() {
		convey.So(base.Title, convey.ShouldEqual, "test")
	})
}

func Test_CaseV1_AsSimpleTestCaseWrapper_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "test",
	}

	// Act & Assert
	convey.Convey("CaseV1.AsSimpleTestCaseWrapper", t, func() {
		convey.So(tc.AsSimpleTestCaseWrapper(), convey.ShouldNotBeNil)
	})
}

func Test_CaseV1_AsSimpleTestCaseWrapperContractsBinder_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "test",
	}

	// Act & Assert
	convey.Convey("CaseV1.AsSimpleTestCaseWrapperContractsBinder", t, func() {
		convey.So(tc.AsSimpleTestCaseWrapperContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- CaseV1MapAssertions ---

func Test_CaseV1_ExpectedAsMap_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "map test",
		ExpectedInput: args.Map{
			"key": "value",
		},
	}

	// Act
	result := tc.ExpectedAsMap()

	// Assert
	convey.Convey("CaseV1.ExpectedAsMap returns map", t, func() {
		convey.So(result["key"], convey.ShouldEqual, "value")
	})
}

func Test_CaseV1_ExpectedAsMap_Panics(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "not-a-map",
		ExpectedInput: "string-value",
	}

	// Act & Assert
	convey.Convey("CaseV1.ExpectedAsMap panics on non-map", t, func() {
		convey.So(func() { tc.ExpectedAsMap() }, convey.ShouldPanic)
	})
}

func Test_CaseV1_ShouldBeEqualMap_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "map equal test",
		ExpectedInput: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"key": "value",
	})
}

func Test_CaseV1_ShouldBeEqualMapFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "map equal first test",
		ExpectedInput: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"key": "value",
	})
}

// --- CaseV1 assertion methods ---

func Test_CaseV1_ShouldContains_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "contains test",
		ExpectedInput: "hell",
	}

	// Act & Assert
	tc.ShouldContains(t, 0, "hello")
}

func Test_CaseV1_ShouldStartsWith_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "starts with test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.ShouldStartsWith(t, 0, "hello world")
}

func Test_CaseV1_ShouldEndsWith_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "ends with test",
		ExpectedInput: "world",
	}

	// Act & Assert
	tc.ShouldEndsWith(t, 0, "hello world")
}

func Test_CaseV1_ShouldBeNotEqual_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "not equal test",
		ExpectedInput: "different",
	}

	// Act & Assert
	tc.ShouldBeNotEqual(t, 0, "hello")
}

func Test_CaseV1_ShouldBeRegex(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "regex test",
		ExpectedInput: "hel.*",
	}

	// Act & Assert
	tc.ShouldBeRegex(t, 0, "hello")
}

func Test_CaseV1_ShouldBeTrimRegex_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "trim regex test",
		ExpectedInput: "hel.*",
	}

	// Act & Assert
	tc.ShouldBeTrimRegex(t, 0, "hello")
}

func Test_CaseV1_ShouldHaveNoError_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "no error test",
	}

	// Act & Assert
	tc.ShouldHaveNoError(t, "additional", 0, nil)
}

func Test_CaseV1_AssertDirectly_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title: "direct assert test",
	}

	// Act & Assert
	tc.AssertDirectly(
		t,
		"extra",
		"message",
		0,
		"actual",
		convey.ShouldEqual,
		"actual",
	)
}

// --- CaseV1 VerifyTypeOfMatch, VerifyType, etc. ---

func Test_CaseV1_VerifyTypeOfMatch(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "type match test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.VerifyTypeOfMatch(t, 0, "world")
}

func Test_CaseV1_VerifyTypeOfMust(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "type must test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.VerifyTypeOfMust(t, 0, "world")
}

func Test_CaseV1_VerifyType(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify type test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.VerifyType(t, 0, "world")
}

func Test_CaseV1_VerifyTypeMust(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify type must test",
		ExpectedInput: "hello",
	}

	// Act & Assert
	tc.VerifyTypeMust(t, 0, "world")
}

func Test_CaseV1_TypeShouldMatch(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "type should match test",
		ExpectedInput: "hello",
		ActualInput:   "world",
	}

	// Act
	err := tc.TypeShouldMatch(t)

	// Assert
	convey.Convey("TypeShouldMatch returns nil for matching types", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- CaseV1 VerifyAll and related ---

func Test_CaseV1_VerifyAllEqual_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify all equal",
		ExpectedInput: "hello",
	}

	// Act
	err := tc.VerifyAllEqual(0, "hello")

	// Assert
	convey.Convey("VerifyAllEqual returns nil on match", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_CaseV1_VerifyError_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify error test",
		ExpectedInput: "hello",
	}

	// Act
	err := tc.VerifyError(0, 1, "hello") // stringcompareas.Equal = 1

	// Assert
	convey.Convey("VerifyError returns nil on match", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_CaseV1_VerifyFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "verify first test",
		ExpectedInput: "hello",
	}

	// Act
	err := tc.VerifyFirst(0, 1, []string{"hello"})

	// Assert
	convey.Convey("VerifyFirst returns nil on match", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_CaseV1_SliceValidator_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseV1{
		Title:         "slice validator test",
		ExpectedInput: "hello",
	}

	// Act
	sv := tc.SliceValidator(1, []string{"hello"})

	// Assert
	convey.Convey("SliceValidator creates validator", t, func() {
		convey.So(sv.ActualLines, convey.ShouldNotBeNil)
	})
}

// --- CaseNilSafe ---

func Test_CaseNilSafe_CaseTitle_Fallback(t *testing.T) {
	// Arrange — empty title, uses MethodName fallback
	tc := coretestcases.CaseNilSafe{
		Func: func() {},
	}

	// Act
	title := tc.CaseTitle()

	// Assert
	convey.Convey("CaseNilSafe.CaseTitle falls back to MethodName", t, func() {
		convey.So(title, convey.ShouldNotBeEmpty)
	})
}

func Test_CaseNilSafe_MethodName_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Func: func() {},
	}

	// Act
	name := tc.MethodName()

	// Assert
	convey.Convey("CaseNilSafe.MethodName returns non-empty", t, func() {
		convey.So(name, convey.ShouldNotBeEmpty)
	})
}

// --- GenericGherkins ---

func Test_GenericGherkins_CompareWith_BothNil_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	var g1 *coretestcases.AnyGherkins
	var g2 *coretestcases.AnyGherkins

	// Act
	isEqual, diff := g1.CompareWith(g2)

	// Assert
	convey.Convey("CompareWith both nil returns true", t, func() {
		convey.So(isEqual, convey.ShouldBeTrue)
		convey.So(diff, convey.ShouldBeEmpty)
	})
}

func Test_GenericGherkins_CompareWith_OneNil_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g1 := &coretestcases.AnyGherkins{Title: "test"}
	var g2 *coretestcases.AnyGherkins

	// Act
	isEqual, diff := g1.CompareWith(g2)

	// Assert
	convey.Convey("CompareWith one nil returns false", t, func() {
		convey.So(isEqual, convey.ShouldBeFalse)
		convey.So(diff, convey.ShouldContainSubstring, "nil")
	})
}

func Test_GenericGherkins_CompareWith_AllDiffs(t *testing.T) {
	// Arrange
	g1 := &coretestcases.AnyGherkins{
		Title:      "a",
		Feature:    "fa",
		Given:      "ga",
		When:       "wa",
		Then:       "ta",
		Input:      "ia",
		Expected:   "ea",
		Actual:     "aa",
		IsMatching: true,
	}
	g2 := &coretestcases.AnyGherkins{
		Title:      "b",
		Feature:    "fb",
		Given:      "gb",
		When:       "wb",
		Then:       "tb",
		Input:      "ib",
		Expected:   "eb",
		Actual:     "ab",
		IsMatching: false,
	}

	// Act
	isEqual, diff := g1.CompareWith(g2)

	// Assert
	convey.Convey("CompareWith all diffs returns false with diff string", t, func() {
		convey.So(isEqual, convey.ShouldBeFalse)
		convey.So(diff, convey.ShouldContainSubstring, "Title")
		convey.So(diff, convey.ShouldContainSubstring, "Feature")
		convey.So(diff, convey.ShouldContainSubstring, "Given")
		convey.So(diff, convey.ShouldContainSubstring, "When")
		convey.So(diff, convey.ShouldContainSubstring, "Then")
		convey.So(diff, convey.ShouldContainSubstring, "Input")
		convey.So(diff, convey.ShouldContainSubstring, "Expected")
		convey.So(diff, convey.ShouldContainSubstring, "Actual")
		convey.So(diff, convey.ShouldContainSubstring, "IsMatching")
	})
}

func Test_GenericGherkins_CompareWith_Equal_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g1 := &coretestcases.AnyGherkins{
		Title: "same",
	}
	g2 := &coretestcases.AnyGherkins{
		Title: "same",
	}

	// Act
	isEqual, diff := g1.CompareWith(g2)

	// Assert
	convey.Convey("CompareWith equal returns true", t, func() {
		convey.So(isEqual, convey.ShouldBeTrue)
		convey.So(diff, convey.ShouldBeEmpty)
	})
}

// --- GenericGherkins Formatting ---

func Test_GenericGherkins_FullString_Nil(t *testing.T) {
	// Arrange
	var g *coretestcases.AnyGherkins

	// Act
	result := g.FullString()

	// Assert
	convey.Convey("FullString nil returns <nil> marker", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "nil")
	})
}

func Test_GenericGherkins_FullString_WithExtraArgs(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:    "test",
		Feature:  "feature",
		Given:    "given",
		When:     "when",
		Then:     "then",
		Input:    "input",
		Expected: "expected",
		Actual:   "actual",
		ExtraArgs: args.Map{
			"extra": "value",
		},
	}

	// Act
	result := g.FullString()

	// Assert
	convey.Convey("FullString includes all fields and extra args", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "Title")
		convey.So(result, convey.ShouldContainSubstring, "ExtraArgs")
	})
}

func Test_GenericGherkins_String_ToString(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Feature: "feat",
		Given:   "given",
		When:    "when",
		Then:    "then",
	}

	// Act
	str := g.String()
	toStr := g.ToString(1)

	// Assert
	convey.Convey("String and ToString return non-empty", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
		convey.So(toStr, convey.ShouldNotBeEmpty)
	})
}

func Test_GenericGherkins_GetWithExpectation(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Feature:  "feat",
		Expected: "exp",
		Actual:   "act",
	}

	// Act
	result := g.GetWithExpectation(0)

	// Assert
	convey.Convey("GetWithExpectation returns formatted string", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_GenericGherkins_GetMessageConditional(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Feature:  "feat",
		Expected: "exp",
	}

	// Act
	withExp := g.GetMessageConditional(true, 0)
	withoutExp := g.GetMessageConditional(false, 0)

	// Assert
	convey.Convey("GetMessageConditional switches between formats", t, func() {
		convey.So(withExp, convey.ShouldNotBeEmpty)
		convey.So(withoutExp, convey.ShouldNotBeEmpty)
	})
}

// --- GenericGherkins Getters ---

func Test_GenericGherkins_IsFailedToMatch(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		IsMatching: true,
	}

	// Act & Assert
	convey.Convey("IsFailedToMatch is inverse of IsMatching", t, func() {
		convey.So(g.IsFailedToMatch(), convey.ShouldBeFalse)
	})
}

func Test_GenericGherkins_ExtraArgs(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		ExtraArgs: args.Map{
			"str":  "hello",
			"bool": true,
		},
	}

	// Act & Assert
	convey.Convey("GenericGherkins extra arg getters", t, func() {
		convey.So(g.HasExtraArgs(), convey.ShouldBeTrue)
		convey.So(g.GetExtra("str"), convey.ShouldEqual, "hello")

		str, ok := g.GetExtraAsString("str")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(str, convey.ShouldEqual, "hello")

		b, ok := g.GetExtraAsBool("bool")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(b, convey.ShouldBeTrue)

		defVal := g.GetExtraAsBoolDefault("missing", true)
		convey.So(defVal, convey.ShouldBeTrue)
	})
}

func Test_GenericGherkins_ExtraArgs_NilReceiver(t *testing.T) {
	// Arrange
	var g *coretestcases.AnyGherkins

	// Act & Assert
	convey.Convey("GenericGherkins nil receiver extra args", t, func() {
		convey.So(g.HasExtraArgs(), convey.ShouldBeFalse)
		convey.So(g.GetExtra("key"), convey.ShouldBeNil)

		_, ok := g.GetExtraAsString("key")
		convey.So(ok, convey.ShouldBeFalse)

		_, ok = g.GetExtraAsBool("key")
		convey.So(ok, convey.ShouldBeFalse)

		convey.So(g.GetExtraAsBoolDefault("key", false), convey.ShouldBeFalse)
	})
}

func Test_GenericGherkins_ExtraArgs_NilMap(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{}

	// Act & Assert
	convey.Convey("GenericGherkins nil ExtraArgs map", t, func() {
		convey.So(g.GetExtra("key"), convey.ShouldBeNil)

		_, ok := g.GetExtraAsString("key")
		convey.So(ok, convey.ShouldBeFalse)

		_, ok = g.GetExtraAsBool("key")
		convey.So(ok, convey.ShouldBeFalse)

		convey.So(g.GetExtraAsBoolDefault("key", true), convey.ShouldBeTrue)
	})
}

// --- GenericGherkinsTypedWrapper ---

func Test_GenericGherkins_CaseTitle_Fallback(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		When: "fallback-when",
	}

	// Act
	title := g.CaseTitle()

	// Assert
	convey.Convey("CaseTitle falls back to When", t, func() {
		convey.So(title, convey.ShouldEqual, "fallback-when")
	})
}

func Test_GenericGherkins_TypedInput_Expected_Actual(t *testing.T) {
	// Arrange
	g := &coretestcases.GenericGherkins[string, int]{
		Input:    "input",
		Expected: 42,
		Actual:   99,
	}

	// Act & Assert
	convey.Convey("TypedInput, TypedExpected, TypedActual", t, func() {
		convey.So(g.TypedInput(), convey.ShouldEqual, "input")
		convey.So(g.TypedExpected(), convey.ShouldEqual, 42)
		convey.So(g.TypedActual(), convey.ShouldEqual, 99)
	})
}

func Test_GenericGherkins_SetTypedActual(t *testing.T) {
	// Arrange
	g := &coretestcases.GenericGherkins[string, int]{
		Input: "test",
	}

	// Act
	g.SetTypedActual(100)

	// Assert
	convey.Convey("SetTypedActual sets actual value", t, func() {
		convey.So(g.TypedActual(), convey.ShouldEqual, 100)
	})
}

func Test_GenericGherkins_AsTypedTestCaseWrapper_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.GenericGherkins[string, int]{
		Title: "wrapper test",
	}

	// Act
	wrapper := g.AsTypedTestCaseWrapper()

	// Assert
	convey.Convey("AsTypedTestCaseWrapper returns interface", t, func() {
		convey.So(wrapper, convey.ShouldNotBeNil)
	})
}

// --- GenericGherkinsAssertions ---

func Test_GenericGherkins_ShouldBeEqual_WithWhenFallback(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		When:          "when-title",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqual(t, 0, []string{"hello"}, []string{"hello"})
}

func Test_GenericGherkins_ShouldBeEqualFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:         "first test",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqualFirst(t, []string{"hello"}, []string{"hello"})
}

func Test_GenericGherkins_ShouldBeEqualArgs_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:         "args test",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqualArgs(t, 0, "hello")
}

func Test_GenericGherkins_ShouldBeEqualArgsFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:         "args first test",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqualArgsFirst(t, "hello")
}

func Test_GenericGherkins_ShouldBeEqualUsingExpected_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:         "using expected test",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqualUsingExpected(t, 0, []string{"hello"})
}

func Test_GenericGherkins_ShouldBeEqualUsingExpectedFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.AnyGherkins{
		Title:         "using expected first test",
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	g.ShouldBeEqualUsingExpectedFirst(t, []string{"hello"})
}

// --- GenericGherkinsMapAssertions ---

func Test_GenericGherkins_ShouldBeEqualMap_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.MapGherkins{
		Title: "map assertion test",
		Expected: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	g.ShouldBeEqualMap(t, 0, args.Map{
		"key": "value",
	})
}

func Test_GenericGherkins_ShouldBeEqualMapFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.MapGherkins{
		Title: "map first test",
		Expected: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	g.ShouldBeEqualMapFirst(t, args.Map{
		"key": "value",
	})
}

// --- GenericGherkinsTypedAssertions ---

func Test_GenericGherkins_ShouldMatchExpected_Match(t *testing.T) {
	// Arrange
	g := &coretestcases.GenericGherkins[string, bool]{
		Title:    "match test",
		Expected: true,
	}

	// Act & Assert (no failure = pass)
	g.ShouldMatchExpected(t, 0, true)
}

func Test_GenericGherkins_ShouldMatchExpectedFirst_FromCaseV1ShouldBeEqualF(t *testing.T) {
	// Arrange
	g := &coretestcases.GenericGherkins[string, string]{
		Title:    "match first test",
		Expected: "hello",
	}

	// Act & Assert
	g.ShouldMatchExpectedFirst(t, "hello")
}
