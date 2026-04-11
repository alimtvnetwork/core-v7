package coretestcasestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── CaseV1 basic getters ──

func Test_CaseV1_Input_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "test title",
		ArrangeInput:  "input",
		ExpectedInput: "expected",
	}

	// Act
	actual := args.Map{
		"input":     c.Input(),
		"expected":  c.Expected(),
		"title":     c.CaseTitle(),
		"typeName":  c.ArrangeTypeName(),
	}

	// Assert
	expected := args.Map{
		"input":     "input",
		"expected":  "expected",
		"title":     "test title",
		"typeName":  "string",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_Input returns correct value -- with args", actual)
}

func Test_CaseV1_ExpectedLines_String_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		ExpectedInput: "hello returns correct value -- with args",
	}

	lines := c.ExpectedLines()

	// Act
	actual := args.Map{
		"len":   len(lines),
		"first": lines[0],
	}

	// Assert
	expected := args.Map{
		"len":   1,
		"first": "hello returns correct value -- with args",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedLines_String returns correct value -- with args", actual)
}

func Test_CaseV1_ExpectedLines_Slice_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		ExpectedInput: []string{"a", "b"},
	}

	lines := c.ExpectedLines()

	// Act
	actual := args.Map{
		"len": len(lines),
	}

	// Assert
	expected := args.Map{
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedLines_Slice returns correct value -- with args", actual)
}

func Test_CaseV1_SetActual_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test", ActualInput: "result"}

	// Act
	actual := args.Map{
		"actual": c.Actual(),
	}

	// Assert
	expected := args.Map{
		"actual": "result",
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_SetActual returns correct value -- with args", actual)
}

func Test_CaseV1_AsSimpleTestCaseWrapper(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}
	wrapper := c.AsSimpleTestCaseWrapper()

	// Act
	actual := args.Map{
		"notNil": wrapper != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_AsSimpleTestCaseWrapper returns correct value -- with args", actual)
}

// ── CaseV1 ShouldBeEqual ──

func Test_CaseV1_ShouldBeEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqual test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	// Assert
	c.ShouldBeEqual(t, 0, "hello returns correct value -- with args")
}

func Test_CaseV1_ShouldBeEqualFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualFirst test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	// Assert
	c.ShouldBeEqualFirst(t, "hello returns correct value -- with args")
}

func Test_CaseV1_ShouldBeTrimEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeTrimEqual test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	c.ShouldBeTrimEqual(t, 0, "hello returns correct value -- with args")
}

func Test_CaseV1_ShouldBeTrimEqualFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeTrimEqualFirst test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	c.ShouldBeTrimEqualFirst(t, "hello returns correct value -- with args")
}

func Test_CaseV1_ShouldBeSortedEqual_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeSortedEqual test",
		ExpectedInput: []string{"a", "b"},
	}

	c.ShouldBeSortedEqual(t, 0, "a", "b")
}

func Test_CaseV1_ShouldBeSortedEqualFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeSortedEqualFirst test",
		ExpectedInput: []string{"a", "b"},
	}

	c.ShouldBeSortedEqualFirst(t, "a", "b")
}

func Test_CaseV1_ShouldContains_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "ShouldContains test",
		ExpectedInput: "hello",
	}

	err := c.VerifyError(0, stringcompareas.Contains, "hello world")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldContains passes -- hello in hello world", actual)
}

func Test_CaseV1_ShouldContainsFirst_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "ShouldContainsFirst test",
		ExpectedInput: "hello",
	}

	err := c.VerifyError(0, stringcompareas.Contains, "hello world")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldContainsFirst passes -- hello in hello world", actual)
}

func Test_CaseV1_ShouldStartsWith_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "ShouldStartsWith test",
		ExpectedInput: "hello",
	}

	err := c.VerifyError(0, stringcompareas.StartsWith, "hello world")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldStartsWith passes -- hello starts hello world", actual)
}

func Test_CaseV1_ShouldStartsWithFirst_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "ShouldStartsWithFirst test",
		ExpectedInput: "hello",
	}

	err := c.VerifyError(0, stringcompareas.StartsWith, "hello world")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldStartsWithFirst passes -- hello starts hello world", actual)
}

func Test_CaseV1_ShouldEndsWith_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldEndsWith test",
		ExpectedInput: "world",
	}

	c.ShouldEndsWith(t, 0, "hello world")
}

func Test_CaseV1_ShouldEndsWithFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldEndsWithFirst test",
		ExpectedInput: "world",
	}

	c.ShouldEndsWithFirst(t, "hello world")
}

func Test_CaseV1_ShouldBeNotEqual_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeNotEqual test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	c.ShouldBeNotEqual(t, 0, "world")
}

func Test_CaseV1_ShouldBeNotEqualFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeNotEqualFirst test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	c.ShouldBeNotEqualFirst(t, "world")
}

// ── CaseV1 VerifyAll / VerifyError ──

func Test_CaseV1_VerifyAllEqual(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "VerifyAllEqual test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	err := c.VerifyAllEqual(0, "hello returns correct value -- with args")

	// Act
	actual := args.Map{
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyAllEqual returns correct value -- with args", actual)
}

func Test_CaseV1_VerifyError_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "VerifyError test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	err := c.VerifyError(0, stringcompareas.Equal, "hello returns correct value -- with args")

	// Act
	actual := args.Map{
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyError returns error -- with args", actual)
}

func Test_CaseV1_VerifyFirst_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "VerifyFirst test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	err := c.VerifyFirst(0, stringcompareas.Equal, []string{"hello returns correct value -- with args"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_VerifyFirst returns correct value -- with args", actual)
}

func Test_CaseV1_SliceValidator_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "SliceValidator test",
		ExpectedInput: "hello returns correct value -- with args",
	}

	sv := c.SliceValidator(stringcompareas.Equal, []string{"hello returns correct value -- with args"})

	// Act
	actual := args.Map{
		"hasActual":   len(sv.ActualLines) > 0,
		"hasExpected": len(sv.ExpectedLines) > 0,
	}

	// Assert
	expected := args.Map{
		"hasActual":   true,
		"hasExpected": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_SliceValidator returns non-empty -- with args", actual)
}

// ── CaseV1 Map Assertions ──

func Test_CaseV1_ShouldBeEqualMap_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualMap test",
		ExpectedInput: args.Map{"key": "value"},
	}

	// Assert
	c.ShouldBeEqualMap(t, 0, args.Map{"key": "value"})
}

func Test_CaseV1_ShouldBeEqualMapFirst_FromCaseV1Input(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ShouldBeEqualMapFirst test",
		ExpectedInput: args.Map{"key": "value"},
	}

	// Assert
	c.ShouldBeEqualMapFirst(t, args.Map{"key": "value"})
}

func Test_CaseV1_ExpectedAsMap_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		ExpectedInput: args.Map{"key": "value"},
	}

	m := c.ExpectedAsMap()

	// Act
	actual := args.Map{
		"len": len(m),
	}

	// Assert
	expected := args.Map{
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "CaseV1_ExpectedAsMap returns correct value -- with args", actual)
}

// ── GenericGherkins CompareWith ──

func Test_GenericGherkins_CompareWith_Equal_FromCaseV1Input(t *testing.T) {
	// Arrange
	g1 := &coretestcases.StringBoolGherkins{Title: "t", Feature: "f", When: "w"}
	g2 := &coretestcases.StringBoolGherkins{Title: "t", Feature: "f", When: "w"}

	isEqual, diff := g1.CompareWith(g2)

	// Act
	actual := args.Map{
		"isEqual":   isEqual,
		"diffEmpty": diff == "",
	}

	// Assert
	expected := args.Map{
		"isEqual":   true,
		"diffEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_Equal returns non-empty -- with args", actual)
}

func Test_GenericGherkins_CompareWith_Diff(t *testing.T) {
	// Arrange
	g1 := &coretestcases.StringBoolGherkins{Title: "a"}
	g2 := &coretestcases.StringBoolGherkins{Title: "b"}

	isEqual, diff := g1.CompareWith(g2)

	// Act
	actual := args.Map{
		"isEqual":    isEqual,
		"hasDiff":    diff != "",
	}

	// Assert
	expected := args.Map{
		"isEqual":    false,
		"hasDiff":    true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_Diff returns non-empty -- with args", actual)
}

func Test_GenericGherkins_CompareWith_BothNil_FromCaseV1Input(t *testing.T) {
	// Arrange
	var g1, g2 *coretestcases.StringBoolGherkins

	isEqual, _ := g1.CompareWith(g2)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_BothNil returns nil -- with args", actual)
}

func Test_GenericGherkins_CompareWith_OneNil_FromCaseV1Input(t *testing.T) {
	// Arrange
	g1 := &coretestcases.StringBoolGherkins{Title: "a"}
	var g2 *coretestcases.StringBoolGherkins

	isEqual, diff := g1.CompareWith(g2)

	// Act
	actual := args.Map{
		"isEqual": isEqual,
		"hasDiff": diff != "",
	}

	// Assert
	expected := args.Map{
		"isEqual": false,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_CompareWith_OneNil returns nil -- with args", actual)
}

// ── GenericGherkins Typed Assertions ──

func Test_GenericGherkins_ShouldMatchExpected(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title:    "ShouldMatchExpected test",
		Expected: true,
	}

	g.ShouldMatchExpected(t, 0, true)
}

func Test_GenericGherkins_ShouldMatchExpectedFirst(t *testing.T) {
	g := &coretestcases.StringBoolGherkins{
		Title:    "ShouldMatchExpectedFirst test",
		Expected: true,
	}

	g.ShouldMatchExpectedFirst(t, true)
}

// ── GenericGherkins TypedWrapper ──

func Test_GenericGherkins_TypedWrapper_FromCaseV1Input(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title:    "wrapper",
		Input:    "input",
		Expected: true,
	}

	g.SetTypedActual(false)

	// Act
	actual := args.Map{
		"caseTitle":    g.CaseTitle(),
		"typedInput":   g.TypedInput(),
		"typedExpect":  g.TypedExpected(),
		"typedActual":  g.TypedActual(),
		"wrapperNotNil": g.AsTypedTestCaseWrapper() != nil,
	}

	// Assert
	expected := args.Map{
		"caseTitle":    "wrapper",
		"typedInput":   "input",
		"typedExpect":  true,
		"typedActual":  false,
		"wrapperNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GenericGherkins_TypedWrapper returns correct value -- with args", actual)
}

// ── CaseNilSafe ──

func Test_CaseNilSafe_CaseTitle_Empty(t *testing.T) {
	// Arrange
	c := coretestcases.CaseNilSafe{
		Func: (*testing.T).Name,
	}

	title := c.CaseTitle()

	// Act
	actual := args.Map{
		"hasTitle": title != "",
	}

	// Assert
	expected := args.Map{
		"hasTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe_CaseTitle_Empty returns nil -- with args", actual)
}

func Test_CaseNilSafe_MethodName_FromCaseV1Input(t *testing.T) {
	// Arrange
	c := coretestcases.CaseNilSafe{
		Title: "explicit title",
		Func:  (*testing.T).Name,
	}

	// Act
	actual := args.Map{
		"title":      c.CaseTitle(),
		"methodName": c.MethodName() != "",
	}

	// Assert
	expected := args.Map{
		"title":      "explicit title",
		"methodName": true,
	}
	expected.ShouldBeEqual(t, 0, "CaseNilSafe_MethodName returns nil -- with args", actual)
}
