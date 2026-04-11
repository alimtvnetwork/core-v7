package coretesttests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/smartystreets/goconvey/convey"
)

type coverage2TestCaseMessenger struct {
	funcName string
	value    any
	expected any
	actual   any
}

func (it coverage2TestCaseMessenger) GetFuncName() string {
	return it.funcName
}

func (it coverage2TestCaseMessenger) Value() any {
	return it.value
}

func (it coverage2TestCaseMessenger) Expected() any {
	return it.expected
}

func (it coverage2TestCaseMessenger) Actual() any {
	return it.actual
}

func Test_BaseTestCase_ShouldAsserters(t *testing.T) {
	// Arrange
	base := &coretests.BaseTestCase{
		Title:         "base should asserter",
		ArrangeInput:  []string{"ok"},
		ExpectedInput: []string{"ok"},
		VerifyTypeOf:  coretests.NewVerifyTypeOf([]string{"seed"}),
	}

	base.ShouldBe(0, t, convey.ShouldResemble, []string{"ok"})
	base.ShouldBeExplicit(false, 1, t, "base explicit", []string{"ok"}, convey.ShouldResemble, []string{"ok"})

	disabled := &coretests.BaseTestCase{
		Title:         "disabled should",
		ExpectedInput: "done",
		IsEnable:      issetter.False,
	}
	disabled.ShouldBe(2, t, convey.ShouldEqual, "done")

	// Act
	actual := args.Map{
		"actualLen": len(base.ActualLines()),
		"actualVal": base.ActualLines()[0],
	}

	// Assert
	expected := args.Map{
		"actualLen": 1,
		"actualVal": "ok",
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- should asserters", actual)
}

func Test_BaseTestCase_TypeValidation(t *testing.T) {
	// Arrange
	mismatch := &coretests.BaseTestCase{
		Title:         "mismatch",
		ArrangeInput:  "arrange",
		ExpectedInput: true,
		VerifyTypeOf:  coretests.NewVerifyTypeOf([]string{"typed"}),
	}
	mismatch.SetActual(100)
	mismatchErr := mismatch.TypeValidationError()

	pass := &coretests.BaseTestCase{
		Title:         "pass",
		ArrangeInput:  []string{"arrange"},
		ExpectedInput: []string{"expected"},
		VerifyTypeOf:  coretests.NewVerifyTypeOf([]string{"typed"}),
	}
	pass.SetActual([]string{"actual"})
	pass.TypesValidationMustPasses(t)
	pass.TypeShouldMatch(t, 0, "type check")

	// Act
	actual := args.Map{
		"mismatchErr": mismatchErr != nil,
		"passErrNil":  pass.TypeValidationError() == nil,
	}

	// Assert
	expected := args.Map{
		"mismatchErr": true,
		"passErrNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns non-empty -- type validation", actual)
}

func Test_Compare_IsMatch(t *testing.T) {
	// Arrange
	instruction := &coretests.ComparingInstruction{}
	instruction.SetActual("alpha beta gamma")

	all := &coretests.Compare{StringContains: "beta alpha", MatchingLength: 0}
	partial := &coretests.Compare{StringContains: "alpha zeta beta", MatchingLength: 2}
	failed := &coretests.Compare{StringContains: "zeta", MatchingLength: 1}

	// Act
	actual := args.Map{
		"all":          all.IsMatch(false, 0, instruction),
		"partial":      partial.IsMatch(false, 1, instruction),
		"failed":       failed.IsMatch(false, 2, instruction),
		"sortedNotEmpty": all.SortedString() != "",
		"printHasIndex":  strings.Contains(all.GetPrintMessage(9), "Index:9"),
	}

	// Assert
	expected := args.Map{
		"all":          true,
		"partial":      true,
		"failed":       false,
		"sortedNotEmpty": true,
		"printHasIndex":  true,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsMatch", actual)
}

func Test_ComparingInstruction_IsMatch(t *testing.T) {
	// Arrange
	instruction := &coretests.ComparingInstruction{
		FunName:                    "matcher",
		Header:                     "when whitespace is sorted",
		TestCaseName:               "is match",
		MatchingAsEqualExpectation: "alpha beta",
		HasWhitespace:              true,
		IsMatchingAsEqual:          true,
		ComparingItems: []coretests.Compare{
			{StringContains: "alpha beta", MatchingLength: 0},
		},
	}

	instruction.SetActual("beta alpha")
	firstHash := instruction.ActualHashset()
	instruction.SetActual("beta alpha")
	secondHash := instruction.ActualHashset()

	isMatch := instruction.IsMatch(&coretests.CaseIndexPlusIsPrint{CaseIndex: 3, IsPrint: false})

	// Act
	actual := args.Map{
		"actual":    instruction.Actual(),
		"hashReset": firstHash != secondHash,
		"match":     isMatch,
	}

	// Assert
	expected := args.Map{
		"actual":    "beta alpha",
		"hashReset": true,
		"match":     true,
	}
	expected.ShouldBeEqual(t, 0, "ComparingInstruction returns correct value -- IsMatch", actual)
}

func Test_SimpleTestCase_ShouldAsserters(t *testing.T) {
	// Arrange
	testCase := coretests.SimpleTestCase{
		Title:         "simple should",
		ArrangeInput:  []string{"ok"},
		ExpectedInput: []string{"ok"},
	}

	testCase.ShouldBe(0, t, convey.ShouldResemble, []string{"ok"})
	testCase.ShouldBeExplicit(1, t, "simple explicit", []string{"ok"}, convey.ShouldResemble, []string{"ok"})

	// Act
	actual := args.Map{
		"stringNotEmpty":     testCase.String(0) != "",
		"linesStringNotEmpty": testCase.LinesString(0) != "",
	}

	// Assert
	expected := args.Map{
		"stringNotEmpty":     true,
		"linesStringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleTestCase returns correct value -- should asserters", actual)
}

func Test_SkipHelpers(t *testing.T) {
	// Arrange
	runCount := 0

	t.Run("skip-on-unix", func(subT *testing.T) {
		runCount++
		coretests.SkipOnUnix(subT)
	})

	t.Run("skip-on-windows", func(subT *testing.T) {
		runCount++
		coretests.SkipOnWindows(subT)
	})

	// Act
	actual := args.Map{"runCount": runCount}

	// Assert
	expected := args.Map{"runCount": 2}
	expected.ShouldBeEqual(t, 0, "Skip returns correct value -- helpers invoked", actual)
}

func Test_IsCompare_And_GetAssertHelpers(t *testing.T) {
	// Arrange
	messenger := coverage2TestCaseMessenger{
		funcName: "CompareFn",
		value:    "when comparing sorted",
		expected: "alpha beta",
		actual:   "beta alpha",
	}

	msg := coretests.GetAssertMessage(messenger, 5)
	header := coretests.GetTestHeader(messenger)

	nonWhiteSortMatch := coretests.IsStrMsgNonWhiteSortedEqual(
		false,
		"beta alpha",
		&errcore.ExpectationMessageDef{
			When:           "sorted",
			Expected:       "alpha beta",
			IsNonWhiteSort: true,
		},
	)

	trimMatch := coretests.IsStrMsgNonWhiteSortedEqual(
		false,
		"  hello  ",
		&errcore.ExpectationMessageDef{
			When:     "trim",
			Expected: "hello",
		},
	)

	errorMatch := coretests.IsErrorNonWhiteSortedEqual(
		false,
		errors.New("beta alpha"),
		&errcore.ExpectationMessageDef{
			When:           "error sorted",
			Expected:       "alpha beta",
			IsNonWhiteSort: true,
		},
	)

	emptyErrorMatch := coretests.IsErrorNonWhiteSortedEqual(
		false,
		nil,
		&errcore.ExpectationMessageDef{
			When:     "empty",
			Expected: "",
		},
	)

	// Act
	actual := args.Map{
		"msgNotEmpty":      msg != "",
		"header":           header,
		"nonWhiteSortMatch": nonWhiteSortMatch,
		"trimMatch":        trimMatch,
		"errorMatch":       errorMatch,
		"emptyErrorMatch":  emptyErrorMatch,
	}

	// Assert
	expected := args.Map{
		"msgNotEmpty":      true,
		"header":           "CompareMethod : [CompareFn]",
		"nonWhiteSortMatch": true,
		"trimMatch":        true,
		"errorMatch":       true,
		"emptyErrorMatch":  true,
	}
	expected.ShouldBeEqual(t, 0, "IsCompare returns correct value -- and getAssert helpers", actual)
}
