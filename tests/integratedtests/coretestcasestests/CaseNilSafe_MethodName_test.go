package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── CaseNilSafe ──

func Test_Cov3_CaseNilSafe_MethodName(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Title: "test nil safe",
		Func:  (*coretests.DraftType).JsonString,
	}

	// Act
	name := tc.MethodName()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", name != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseNilSafe.MethodName returns function name -- DraftType.JsonString",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseNilSafe_CaseTitle_WithTitle(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Title: "my title",
		Func:  (*coretests.DraftType).JsonString,
	}

	// Act
	title := tc.CaseTitle()

	// Assert
	actual := args.Map{
		"title": title,
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseNilSafe.CaseTitle returns Title when set -- explicit title",
		ExpectedInput: args.Map{
			"title": "my title",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseNilSafe_CaseTitle_FallbackToMethodName(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Func: (*coretests.DraftType).JsonString,
	}

	// Act
	title := tc.CaseTitle()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", title != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseNilSafe.CaseTitle falls back to MethodName -- no title",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseNilSafe_Invoke(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Title: "invoke with valid receiver",
		Func:  (*coretests.DraftType).F1String,
	}
	receiver := &coretests.DraftType{SampleString1: "hello"}

	// Act
	result := tc.Invoke(receiver)

	// Assert
	actual := args.Map{
		"panicked": fmt.Sprintf("%v", result.Panicked),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseNilSafe.Invoke does not panic on valid receiver -- F1String",
		ExpectedInput: args.Map{
			"panicked": "false",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseNilSafe_InvokeNil(t *testing.T) {
	// Arrange
	tc := coretestcases.CaseNilSafe{
		Title: "invoke with nil receiver",
		Func:  (*coretests.DraftType).ClonePtr,
	}

	// Act
	result := tc.InvokeNil()

	// Assert
	actual := args.Map{
		"panicked": fmt.Sprintf("%v", result.Panicked),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseNilSafe.InvokeNil does not panic -- ClonePtr handles nil",
		ExpectedInput: args.Map{
			"panicked": "false",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

// ── CaseV1 additional methods ──

func Test_Cov3_CaseV1_ArrangeTypeName(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:        "type name test",
		ArrangeInput: args.Map{"key": "val"},
	}

	// Act
	typeName := c.ArrangeTypeName()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", typeName != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ArrangeTypeName returns type name -- args.Map input",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_AsBaseTestCase(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "base test",
		ArrangeInput:  "arrange",
		ExpectedInput: "expected",
	}

	// Act
	base := c.AsBaseTestCase()

	// Assert
	actual := args.Map{
		"title":    base.CaseTitle(),
		"input":    fmt.Sprintf("%v", base.Input()),
		"expected": fmt.Sprintf("%v", base.Expected()),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.AsBaseTestCase converts to BaseTestCase -- preserves fields",
		ExpectedInput: args.Map{
			"title":    "base test",
			"input":    "arrange",
			"expected": "expected",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_PrepareTitle(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "my case"}

	// Act
	result := c.PrepareTitle(3, "extra info")

	// Assert
	actual := args.Map{
		"notEmpty":     fmt.Sprintf("%v", result != ""),
		"containsCase": fmt.Sprintf("%v", len(result) > 10),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.PrepareTitle formats index + title + additional -- basic",
		ExpectedInput: args.Map{
			"notEmpty":     "true",
			"containsCase": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_ExpectedAsMap(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "map expected",
		ExpectedInput: args.Map{"key": "val"},
	}

	// Act
	m := c.ExpectedAsMap()

	// Assert
	actual := args.Map{
		"hasKey": fmt.Sprintf("%v", m["key"] == "val"),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ExpectedAsMap extracts args.Map -- valid map",
		ExpectedInput: args.Map{
			"hasKey": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_ExpectedAsMap_Panic(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "not a map",
		ExpectedInput: "string",
	}

	// Act
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		c.ExpectedAsMap()
	}()

	// Assert
	actual := args.Map{
		"panicked": fmt.Sprintf("%v", didPanic),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ExpectedAsMap panics on non-map -- string expected",
		ExpectedInput: args.Map{
			"panicked": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_AsSimpleTestCaseWrapperContractsBinder(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "test"}

	// Act
	binder := c.AsSimpleTestCaseWrapperContractsBinder()

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", binder != nil),
		"title":  binder.CaseTitle(),
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.AsSimpleTestCaseWrapperContractsBinder returns valid binder",
		ExpectedInput: args.Map{
			"notNil": "true",
			"title":  "test",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_ShouldStartsWithFirst(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "starts with test",
		ExpectedInput: "hel",
	}

	// Act / Assert
	c.ShouldStartsWithFirst(t, "hello world")
}

func Test_Cov3_CaseV1_ShouldEndsWithFirst(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "ends with test",
		ExpectedInput: "world",
	}

	// Act / Assert
	c.ShouldEndsWithFirst(t, "hello world")
}

func Test_Cov3_CaseV1_ShouldBeNotEqualFirst(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "not equal test",
		ExpectedInput: "xyz",
	}

	// Act / Assert
	c.ShouldBeNotEqualFirst(t, "hello")
}

func Test_Cov3_CaseV1_ShouldBeRegexFirst(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "regex test",
		ExpectedInput: "hel.*",
	}

	// Act / Assert
	c.ShouldBeRegexFirst(t, "hello")
}

func Test_Cov3_CaseV1_ShouldHaveNoError(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title: "no error test",
	}

	// Act / Assert
	c.ShouldHaveNoError(t, "additional", 0, nil)
}

// ── GenericGherkins CompareWith ──

func Test_Cov3_GenericGherkins_CompareWith_BothNil(t *testing.T) {
	// Arrange
	var a, b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", isEqual),
		"diff":    diff,
	}

	verify := coretestcases.CaseV1{
		Title:         "CompareWith returns true for both nil -- nil/nil",
		ExpectedInput: args.Map{
			"isEqual": "true",
			"diff":    "",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_GenericGherkins_CompareWith_OneNil(t *testing.T) {
	// Arrange
	a := &coretestcases.StringBoolGherkins{Title: "test"}
	var b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", isEqual),
		"diff":    diff,
	}

	verify := coretestcases.CaseV1{
		Title:         "CompareWith returns false when one is nil -- a!=nil, b=nil",
		ExpectedInput: args.Map{
			"isEqual": "false",
			"diff":    "one side is nil",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_GenericGherkins_CompareWith_Equal(t *testing.T) {
	// Arrange
	a := &coretestcases.StringBoolGherkins{
		Title:   "test",
		Feature: "f",
		Given:   "g",
		When:    "w",
		Then:    "t",
		Input:   "in",
		Expected: true,
	}
	b := &coretestcases.StringBoolGherkins{
		Title:   "test",
		Feature: "f",
		Given:   "g",
		When:    "w",
		Then:    "t",
		Input:   "in",
		Expected: true,
	}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", isEqual),
		"diff":    diff,
	}

	verify := coretestcases.CaseV1{
		Title:         "CompareWith returns true for equal structs -- all fields match",
		ExpectedInput: args.Map{
			"isEqual": "true",
			"diff":    "",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_GenericGherkins_CompareWith_TitleMismatch(t *testing.T) {
	// Arrange
	a := &coretestcases.StringBoolGherkins{Title: "a"}
	b := &coretestcases.StringBoolGherkins{Title: "b"}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actual := args.Map{
		"isEqual":      fmt.Sprintf("%v", isEqual),
		"hasDiff":      fmt.Sprintf("%v", diff != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "CompareWith returns false on Title mismatch -- a vs b",
		ExpectedInput: args.Map{
			"isEqual":      "false",
			"hasDiff":      "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_GenericGherkins_CompareWith_IsMatchingMismatch(t *testing.T) {
	// Arrange
	a := &coretestcases.StringBoolGherkins{IsMatching: true}
	b := &coretestcases.StringBoolGherkins{IsMatching: false}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", isEqual),
		"hasDiff": fmt.Sprintf("%v", diff != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "CompareWith returns false on IsMatching mismatch -- true vs false",
		ExpectedInput: args.Map{
			"isEqual": "false",
			"hasDiff": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

// ── GenericGherkins TypedWrapper ──

func Test_Cov3_GenericGherkins_TypedWrapper(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title:    "typed test",
		Input:    "input_val",
		Expected: true,
	}

	// Act
	g.SetTypedActual(false)

	// Assert
	actual := args.Map{
		"caseTitle":     g.CaseTitle(),
		"typedInput":    g.TypedInput(),
		"typedExpected": fmt.Sprintf("%v", g.TypedExpected()),
		"typedActual":   fmt.Sprintf("%v", g.TypedActual()),
	}

	verify := coretestcases.CaseV1{
		Title:         "GenericGherkins TypedWrapper methods work correctly -- StringBoolGherkins",
		ExpectedInput: args.Map{
			"caseTitle":     "typed test",
			"typedInput":    "input_val",
			"typedExpected": "true",
			"typedActual":   "false",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_GenericGherkins_AsTypedTestCaseWrapper(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title: "wrapper test",
		Input: "in",
	}

	// Act
	wrapper := g.AsTypedTestCaseWrapper()

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", wrapper != nil),
		"title":  wrapper.CaseTitle(),
		"input":  wrapper.TypedInput(),
	}

	verify := coretestcases.CaseV1{
		Title:         "GenericGherkins.AsTypedTestCaseWrapper returns valid wrapper -- basic",
		ExpectedInput: args.Map{
			"notNil": "true",
			"title":  "wrapper test",
			"input":  "in",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

// ── GenericGherkins ShouldMatchExpected ──

func Test_Cov3_GenericGherkins_ShouldMatchExpected_Pass(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title:    "match pass",
		Expected: true,
	}

	// Act / Assert - should not fail
	g.ShouldMatchExpected(t, 0, true)
}

func Test_Cov3_GenericGherkins_ShouldMatchExpectedFirst_Pass(t *testing.T) {
	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title:    "match first pass",
		Expected: false,
	}

	// Act / Assert - should not fail
	g.ShouldMatchExpectedFirst(t, false)
}

// ── GenericGherkins ShouldBeEqualArgs ──

func Test_Cov3_GenericGherkins_ShouldBeEqualArgs(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title:         "equal args test",
		ExpectedLines: []string{"a", "b"},
	}

	// Act / Assert
	g.ShouldBeEqualArgs(t, 0, "a", "b")
}

func Test_Cov3_GenericGherkins_ShouldBeEqualArgsFirst(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title:         "equal args first test",
		ExpectedLines: []string{"x"},
	}

	// Act / Assert
	g.ShouldBeEqualArgsFirst(t, "x")
}

// ── GenericGherkins ShouldBeEqualUsingExpected ──

func Test_Cov3_GenericGherkins_ShouldBeEqualUsingExpected(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title:         "using expected test",
		ExpectedLines: []string{"line1", "line2"},
	}

	// Act / Assert
	g.ShouldBeEqualUsingExpected(t, 0, []string{"line1", "line2"})
}

func Test_Cov3_GenericGherkins_ShouldBeEqualUsingExpectedFirst(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title:         "using expected first test",
		ExpectedLines: []string{"only"},
	}

	// Act / Assert
	g.ShouldBeEqualUsingExpectedFirst(t, []string{"only"})
}

// ── GenericGherkins ShouldBeEqual ──

func Test_Cov3_GenericGherkins_ShouldBeEqual(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title: "equal test",
	}

	// Act / Assert
	g.ShouldBeEqual(t, 0, []string{"a"}, []string{"a"})
}

func Test_Cov3_GenericGherkins_ShouldBeEqualFirst(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		Title: "equal first test",
	}

	// Act / Assert
	g.ShouldBeEqualFirst(t, []string{"b"}, []string{"b"})
}

// ── GenericGherkins ShouldBeEqualMap ──

func Test_Cov3_GenericGherkins_ShouldBeEqualMap(t *testing.T) {
	// Arrange
	g := &coretestcases.MapGherkins{
		Title: "map equal test",
		Expected: args.Map{
			"key1": "val1",
			"key2": "val2",
		},
	}

	// Act / Assert
	g.ShouldBeEqualMap(t, 0, args.Map{
		"key1": "val1",
		"key2": "val2",
	})
}

func Test_Cov3_GenericGherkins_ShouldBeEqualMapFirst(t *testing.T) {
	// Arrange
	g := &coretestcases.MapGherkins{
		Title: "map equal first test",
		Expected: args.Map{
			"k": "v",
		},
	}

	// Act / Assert
	g.ShouldBeEqualMapFirst(t, args.Map{
		"k": "v",
	})
}

// ── GenericGherkins CaseTitle fallback ──

func Test_Cov3_GenericGherkins_CaseTitle_WhenFallback(t *testing.T) {
	// Arrange
	g := &coretestcases.StringGherkins{
		When: "when fallback",
	}

	// Act
	title := g.CaseTitle()

	// Assert
	actual := args.Map{
		"title": title,
	}

	verify := coretestcases.CaseV1{
		Title:         "GenericGherkins.CaseTitle falls back to When -- no title",
		ExpectedInput: args.Map{
			"title": "when fallback",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

// ── CaseV1 ShouldBeEqualMap ──

func Test_Cov3_CaseV1_ShouldBeEqualMap(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title: "map assertion test",
		ExpectedInput: args.Map{
			"k1": "v1",
			"k2": "v2",
		},
	}

	// Act / Assert
	c.ShouldBeEqualMap(t, 0, args.Map{
		"k1": "v1",
		"k2": "v2",
	})
}

func Test_Cov3_CaseV1_ShouldBeEqualMapFirst(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title: "map first assertion test",
		ExpectedInput: args.Map{
			"only": "one",
		},
	}

	// Act / Assert
	c.ShouldBeEqualMapFirst(t, args.Map{
		"only": "one",
	})
}

// ── CaseV1 ExpectedLines with various types ──

func Test_Cov3_CaseV1_ExpectedLines_Int(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "int expected lines",
		ExpectedInput: 99,
	}

	// Act
	lines := c.ExpectedLines()

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(lines)),
		"line0":     lines[0],
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ExpectedLines converts int -- 99",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "99",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_ExpectedLines_Bool(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "bool expected lines",
		ExpectedInput: true,
	}

	// Act
	lines := c.ExpectedLines()

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(lines)),
		"line0":     lines[0],
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ExpectedLines converts bool -- true",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov3_CaseV1_ExpectedLines_MapStringInt(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title: "map[string]int expected lines",
		ExpectedInput: map[string]int{
			"age": 25,
		},
	}

	// Act
	lines := c.ExpectedLines()

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(lines)),
		"line0":     lines[0],
	}

	verify := coretestcases.CaseV1{
		Title:         "CaseV1.ExpectedLines converts map[string]int -- single entry",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "age : 25",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}
