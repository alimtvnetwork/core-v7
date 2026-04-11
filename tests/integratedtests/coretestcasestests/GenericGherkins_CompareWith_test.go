package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── GenericGherkins — CompareWith ──

func Test_GenericGherkins_CompareWith_BothNil_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	var a, b *coretestcases.GenericGherkins[string, string]
	eq, diff := a.CompareWith(b)

	// Act
	actual := args.Map{
		"eq": eq,
		"diff": diff,
	}

	// Assert
	expected := args.Map{
		"eq": true,
		"diff": "",
	}
	expected.ShouldBeEqual(t, 0, "CompareWith returns nil -- both nil", actual)
}

func Test_GenericGherkins_CompareWith_OneNil_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Title: "x"}
	eq, diff := a.CompareWith(nil)

	// Act
	actual := args.Map{
		"eq": eq,
		"diff": diff,
	}

	// Assert
	expected := args.Map{
		"eq": false,
		"diff": "one side is nil",
	}
	expected.ShouldBeEqual(t, 0, "CompareWith returns nil -- one nil", actual)
}

func Test_GenericGherkins_CompareWith_Equal_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Title: "t", Feature: "f", Given: "g", When: "w", Then: "th", Input: "i", Expected: "e", Actual: "a", IsMatching: true}
	b := &coretestcases.GenericGherkins[string, string]{Title: "t", Feature: "f", Given: "g", When: "w", Then: "th", Input: "i", Expected: "e", Actual: "a", IsMatching: true}
	eq, diff := a.CompareWith(b)

	// Act
	actual := args.Map{
		"eq": eq,
		"diff": diff,
	}

	// Assert
	expected := args.Map{
		"eq": true,
		"diff": "",
	}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- equal", actual)
}

func Test_GenericGherkins_CompareWith_TitleDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Title: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Title: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- title diff", actual)
}

func Test_GenericGherkins_CompareWith_FeatureDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Feature: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Feature: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- feature diff", actual)
}

func Test_GenericGherkins_CompareWith_GivenDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Given: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Given: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- given diff", actual)
}

func Test_GenericGherkins_CompareWith_WhenDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{When: "a"}
	b := &coretestcases.GenericGherkins[string, string]{When: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- when diff", actual)
}

func Test_GenericGherkins_CompareWith_ThenDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Then: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Then: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- then diff", actual)
}

func Test_GenericGherkins_CompareWith_InputDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Input: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Input: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- input diff", actual)
}

func Test_GenericGherkins_CompareWith_ExpectedDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Expected: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Expected: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- expected diff", actual)
}

func Test_GenericGherkins_CompareWith_ActualDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{Actual: "a"}
	b := &coretestcases.GenericGherkins[string, string]{Actual: "b"}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- actual diff", actual)
}

func Test_GenericGherkins_CompareWith_IsMatchingDiff(t *testing.T) {
	// Arrange
	a := &coretestcases.GenericGherkins[string, string]{IsMatching: true}
	b := &coretestcases.GenericGherkins[string, string]{IsMatching: false}
	eq, _ := a.CompareWith(b)

	// Act
	actual := args.Map{"eq": eq}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "CompareWith returns non-empty -- ismatching diff", actual)
}

// ── GenericGherkins — Formatting ──

func Test_GenericGherkins_String(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Feature: "f", Given: "g", When: "w", Then: "th",
	}
	result := tc.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_GenericGherkins_ToString(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{Feature: "f", Given: "g", When: "w", Then: "th"}
	result := tc.ToString(5)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToString returns correct value -- with args", actual)
}

func Test_GenericGherkins_GetWithExpectation_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Feature: "f", Given: "g", When: "w", Then: "th",
		Actual: "act", Expected: "exp",
	}
	result := tc.GetWithExpectation(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetWithExpectation returns non-empty -- with args", actual)
}

func Test_GenericGherkins_GetMessageConditional_True(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Feature: "f", Given: "g", When: "w", Then: "th",
		Actual: "act", Expected: "exp",
	}
	result := tc.GetMessageConditional(true, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetMessageConditional returns non-empty -- true", actual)
}

func Test_GenericGherkins_GetMessageConditional_False(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Feature: "f", Given: "g", When: "w", Then: "th",
	}
	result := tc.GetMessageConditional(false, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetMessageConditional returns non-empty -- false", actual)
}

func Test_GenericGherkins_FullString_Nil_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]
	result := tc.FullString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "<nil GenericGherkins>"}
	expected.ShouldBeEqual(t, 0, "FullString returns nil -- nil", actual)
}

func Test_GenericGherkins_FullString_Valid(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Title: "t", Feature: "f", Given: "g", When: "w", Then: "th",
		Input: "i", Expected: "e", Actual: "a", IsMatching: true,
	}
	result := tc.FullString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns non-empty -- valid", actual)
}

func Test_GenericGherkins_FullString_WithExtraArgs_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{
		Title:     "t",
		ExtraArgs: args.Map{"k": "v"},
	}
	result := tc.FullString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns non-empty -- with extras", actual)
}

// ── GenericGherkins — Getters ──

func Test_GenericGherkins_IsFailedToMatch_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{IsMatching: true}

	// Act
	actual := args.Map{"val": tc.IsFailedToMatch()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsFailedToMatch returns correct value -- with args", actual)
}

func Test_GenericGherkins_HasExtraArgs_Empty(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{}

	// Act
	actual := args.Map{"val": tc.HasExtraArgs()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasExtraArgs returns empty -- empty", actual)
}

func Test_GenericGherkins_HasExtraArgs_WithData(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{ExtraArgs: args.Map{"k": "v"}}

	// Act
	actual := args.Map{"val": tc.HasExtraArgs()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasExtraArgs returns non-empty -- with data", actual)
}

func Test_GenericGherkins_HasExtraArgs_Nil(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]

	// Act
	actual := args.Map{"val": tc.HasExtraArgs()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasExtraArgs returns nil -- nil", actual)
}

func Test_GenericGherkins_GetExtra_Valid(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{ExtraArgs: args.Map{"k": "v"}}

	// Act
	actual := args.Map{"val": fmt.Sprintf("%v", tc.GetExtra("k"))}

	// Assert
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "GetExtra returns non-empty -- valid", actual)
}

func Test_GenericGherkins_GetExtra_Nil(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]

	// Act
	actual := args.Map{"nil": tc.GetExtra("k") == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetExtra returns nil -- nil", actual)
}

func Test_GenericGherkins_GetExtra_NilExtraArgs(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{}

	// Act
	actual := args.Map{"nil": tc.GetExtra("k") == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetExtra returns nil -- nil extraargs", actual)
}

func Test_GenericGherkins_GetExtraAsString_Valid(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{ExtraArgs: args.Map{"k": "v"}}
	val, ok := tc.GetExtraAsString("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsString returns non-empty -- valid", actual)
}

func Test_GenericGherkins_GetExtraAsString_Nil(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]
	val, ok := tc.GetExtraAsString("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "",
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsString returns nil -- nil", actual)
}

func Test_GenericGherkins_GetExtraAsBool_Valid(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{ExtraArgs: args.Map{"k": true}}
	val, ok := tc.GetExtraAsBool("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBool returns non-empty -- valid", actual)
}

func Test_GenericGherkins_GetExtraAsBool_Nil(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]
	val, ok := tc.GetExtraAsBool("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": false,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBool returns nil -- nil", actual)
}

func Test_GenericGherkins_GetExtraAsBoolDefault_Valid(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{ExtraArgs: args.Map{"k": false}}
	val := tc.GetExtraAsBoolDefault("k", true)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBoolDefault returns non-empty -- valid", actual)
}

func Test_GenericGherkins_GetExtraAsBoolDefault_Nil(t *testing.T) {
	// Arrange
	var tc *coretestcases.GenericGherkins[string, string]
	val := tc.GetExtraAsBoolDefault("k", true)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBoolDefault returns nil -- nil", actual)
}

func Test_GenericGherkins_GetExtraAsBoolDefault_NilExtraArgs(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{}
	val := tc.GetExtraAsBoolDefault("k", true)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "GetExtraAsBoolDefault returns nil -- nil extraargs", actual)
}

// ── GenericGherkins — TypedWrapper ──

func Test_GenericGherkins_AsTypedTestCaseWrapper_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{Title: "wrapper"}
	w := tc.AsTypedTestCaseWrapper()

	// Act
	actual := args.Map{
		"notNil": w != nil,
		"title": w.CaseTitle(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"title": "wrapper",
	}
	expected.ShouldBeEqual(t, 0, "AsTypedTestCaseWrapper returns correct value -- with args", actual)
}

// ── GenericGherkins — ShouldMatchExpected match ──

func Test_GenericGherkins_ShouldMatchExpected_Match_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, bool]{Title: "match test", Expected: true}
	// Should not error — actual matches expected
	tc.ShouldMatchExpected(t, 0, true)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ShouldMatchExpected returns correct value -- match", actual)
}

func Test_GenericGherkins_ShouldMatchExpectedFirst_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, bool]{Title: "match first", Expected: true}
	tc.ShouldMatchExpectedFirst(t, true)

	// Act
	actual := args.Map{"passed": true}

	// Assert
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ShouldMatchExpectedFirst returns correct value -- with args", actual)
}

// ── GenericGherkins — ShouldBeEqualMap ──

func Test_MapGherkins_ShouldBeEqualMap(t *testing.T) {
	tc := &coretestcases.MapGherkins{
		Title:    "map equal test",
		Expected: args.Map{"k": "v"},
	}

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{"k": "v"})
}

func Test_MapGherkins_ShouldBeEqualMapFirst(t *testing.T) {
	tc := &coretestcases.MapGherkins{
		Title:    "map equal first test",
		Expected: args.Map{"k": "v"},
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"k": "v"})
}

func Test_MapGherkins_ShouldBeEqualMap_WhenFallback_FromGenericGherkinsCompa(t *testing.T) {
	tc := &coretestcases.MapGherkins{
		When:     "when title",
		Expected: args.Map{"k": "v"},
	}

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{"k": "v"})
}

// ── CaseV1 — ArrangeTypeName ──

func Test_CaseV1_ArrangeTypeName_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ArrangeInput: "hello"}

	// Act
	actual := args.Map{"val": c.ArrangeTypeName()}

	// Assert
	expected := args.Map{"val": "string"}
	expected.ShouldBeEqual(t, 0, "ArrangeTypeName returns correct value -- with args", actual)
}

// ── CaseV1 — AsSimpleTestCaseWrapper ──

func Test_CaseV1_AsSimpleTestCaseWrapper_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}
	w := c.AsSimpleTestCaseWrapper()

	// Act
	actual := args.Map{"title": w.CaseTitle()}

	// Assert
	expected := args.Map{"title": "test"}
	expected.ShouldBeEqual(t, 0, "AsSimpleTestCaseWrapper returns correct value -- with args", actual)
}

// ── CaseV1 — AsSimpleTestCaseWrapperContractsBinder ──

func Test_CaseV1_AsSimpleTestCaseWrapperContractsBinder_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}
	w := c.AsSimpleTestCaseWrapperContractsBinder()

	// Act
	actual := args.Map{"notNil": w != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsSimpleTestCaseWrapperContractsBinder returns correct value -- with args", actual)
}

// ── CaseV1 — AsBaseTestCase ──

func Test_CaseV1_AsBaseTestCase_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}
	b := c.AsBaseTestCase()

	// Act
	actual := args.Map{"title": b.Title}

	// Assert
	expected := args.Map{"title": "test"}
	expected.ShouldBeEqual(t, 0, "AsBaseTestCase returns correct value -- with args", actual)
}

// ── CaseV1 — PrepareTitle ──

func Test_CaseV1_PrepareTitle_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}
	result := c.PrepareTitle(3, "additional")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrepareTitle returns correct value -- with args", actual)
}

// ── CaseV1 — ShouldStartsWithFirst / ShouldEndsWithFirst ──

func Test_CaseV1_ShouldStartsWithFirst_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "starts test", ExpectedInput: "hel"}
	c.ShouldStartsWithFirst(t, "hello")
}

func Test_CaseV1_ShouldEndsWithFirst_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "ends test", ExpectedInput: "llo"}
	c.ShouldEndsWithFirst(t, "hello")
}

// ── CaseV1 — ShouldBeNotEqualFirst ──

func Test_CaseV1_ShouldBeNotEqualFirst_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "not equal first", ExpectedInput: "abc"}
	c.ShouldBeNotEqualFirst(t, "xyz")
}

// ── CaseV1 — ShouldBeRegexFirst ──

func Test_CaseV1_ShouldBeRegexFirst_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "regex first", ExpectedInput: "^hel.*"}
	c.ShouldBeRegexFirst(t, "hello")
}

// ── CaseV1 — ShouldBeTrimRegex ──

func Test_CaseV1_ShouldBeTrimRegex_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "trim regex", ExpectedInput: "^hello$"}
	c.ShouldBeTrimRegex(t, 0, "  hello  ")
}

// ── CaseV1 — ShouldHaveNoError ──

func Test_CaseV1_ShouldHaveNoError_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "no error test"}
	c.ShouldHaveNoError(t, "additional", 0, nil)
}

// ── CaseV1 — ExpectedAsMap ──

func Test_CaseV1_ExpectedAsMap_Valid(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ExpectedInput: args.Map{"k": "v"}}
	m := c.ExpectedAsMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ExpectedAsMap returns non-empty -- valid", actual)
}

func Test_CaseV1_ExpectedAsMap_Panic_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{ExpectedInput: "not a map"}
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		c.ExpectedAsMap()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ExpectedAsMap panics -- panic", actual)
}

// ── CaseV1 — ShouldBeEqualMap / ShouldBeEqualMapFirst ──

func Test_CaseV1_ShouldBeEqualMap_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "map test", ExpectedInput: args.Map{"k": "v"}}

	// Assert
	c.ShouldBeEqualMap(t, 0, args.Map{"k": "v"})
}

func Test_CaseV1_ShouldBeEqualMapFirst_FromGenericGherkinsCompa(t *testing.T) {
	c := coretestcases.CaseV1{Title: "map first", ExpectedInput: args.Map{"k": "v"}}

	// Assert
	c.ShouldBeEqualMapFirst(t, args.Map{"k": "v"})
}

// ── CaseNilSafe — MethodName / CaseTitle ──

func Test_CaseNilSafe_MethodName_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{Func: fmt.Sprint}

	// Act
	actual := args.Map{"notEmpty": tc.MethodName() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe returns nil -- MethodName", actual)
}

func Test_CaseNilSafe_CaseTitle_WithTitle_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{Title: "my title", Func: fmt.Sprint}

	// Act
	actual := args.Map{"val": tc.CaseTitle()}

	// Assert
	expected := args.Map{"val": "my title"}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe returns nil -- CaseTitle with title", actual)
}

func Test_CaseNilSafe_CaseTitle_Fallback_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{Func: fmt.Sprint}

	// Act
	actual := args.Map{"notEmpty": tc.CaseTitle() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe returns nil -- CaseTitle fallback", actual)
}

// ── GenericGherkins — CaseTitle fallback ──

func Test_GenericGherkins_CaseTitle_WhenFallback_FromGenericGherkinsCompa(t *testing.T) {
	// Arrange
	tc := &coretestcases.GenericGherkins[string, string]{When: "when-val"}

	// Act
	actual := args.Map{"val": tc.CaseTitle()}

	// Assert
	expected := args.Map{"val": "when-val"}
	expected.ShouldBeEqual(t, 0, "CaseTitle returns correct value -- when fallback", actual)
}
