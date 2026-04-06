package corevalidatortests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// LineValidator — enhanced tests with diff diagnostics
// ==========================================

func Test_LineValidator_VerifyError_LineAndTextMismatch_PrintsDiff(t *testing.T) {
	// Arrange: both line number and text are wrong
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "expected-text",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "LineValidator line+text mismatch",
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyError(params, 10, "actual-wrong-text")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for line+text mismatch", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    params.Header,
		Error:     err,
	}
	diag.PrintLineNumberMismatch(5, 10, "expected-text", "actual-wrong-text")
}

func Test_LineValidator_AllVerifyError_MultipleContents_PrintsDiff(t *testing.T) {
	// Arrange: 5 items, 3 fail
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "target",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "AllVerify multi-content diff",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "target", LineNumber: 0},
		{Text: "wrong-1", LineNumber: 1},
		{Text: "target", LineNumber: 2},
		{Text: "wrong-2", LineNumber: 3},
		{Text: "wrong-3", LineNumber: 4},
	}

	// Act
	err := v.AllVerifyError(params, items...)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for 3 mismatches", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    "LineValidator AllVerifyError Diff",
		Error:     err,
	}
	diag.PrintLineMatchDiagnostics(items, func(text string) bool {
		return text == "target"
	}, "target")

	errMsg := err.Error()
	for _, expected := range []string{"wrong-1", "wrong-2", "wrong-3"} {
		actual := args.Map{"result": strings.Contains(errMsg, expected)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "error should contain '', got:\n", actual)
	}
}

func Test_LineValidator_VerifyMany_CollectAll_PrintsDiff(t *testing.T) {
	// Arrange
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "VerifyMany collectAll diff",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "is ok here", LineNumber: 0},
		{Text: "no match", LineNumber: 1},
		{Text: "also ok", LineNumber: 2},
		{Text: "missing", LineNumber: 3},
	}

	// isContinueOnError = true -> collect all
	err := v.VerifyMany(true, params, items...)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected errors for lines 1 and 3", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    "VerifyMany (collectAll) Diff",
		Error:     err,
	}
	diag.PrintLineMatchDiagnostics(items, func(text string) bool {
		return strings.Contains(text, "ok")
	}, "Contains 'ok'")

	errMsg := err.Error()
	actual := args.Map{"result": strings.Contains(errMsg, "no match")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should mention 'no match':", actual)
}

// ==========================================
// LineValidator — VerifyFirstError with line number specifics
// ==========================================

func Test_LineValidator_VerifyFirstError_SpecificLineNumber_PrintsDiff(t *testing.T) {
	// Arrange: validator expects line 2 specifically
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 2},
		TextValidator: corevalidator.TextValidator{
			Search:    "expected",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "VerifyFirst specific line number",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "expected", LineNumber: 0}, // wrong line number
		{Text: "expected", LineNumber: 1}, // wrong line number
		{Text: "expected", LineNumber: 2}, // correct!
	}

	err := v.VerifyFirstError(params, items...)

	// The first item has LineNumber=0 but validator expects 2, so error
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "line 0 doesn't match expected line 2, should error", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    "Line number mismatch diagnostic",
		Error:     err,
	}
	diag.PrintLineNumberMismatch(2, 0, "expected", items[0].Text)
}

// ==========================================
// LinesValidators — enhanced multi-validator diff
// ==========================================

func Test_LinesValidators_AllVerifyError_MultiValidator_PrintsDiff(t *testing.T) {
	// Arrange: 2 validators, each checking different search terms
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "alpha",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "beta",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "Multi-validator AllVerifyError",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "contains alpha here", LineNumber: 0},
		{Text: "no match at all", LineNumber: 1},
		{Text: "alpha and beta together", LineNumber: 2},
	}

	err := lv.AllVerifyError(params, items...)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected errors: 'beta' not in line 0, 'alpha' not in line 1", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    "LinesValidators Multi-Validator Diff",
		Error:     err,
	}
	diag.PrintMultiValidatorDiagnostics(items, []string{"alpha", "beta"})
}

func Test_LinesValidators_IsMatchText_Multiple_PrintsDiff(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(3)
	lv.Adds(
		corevalidator.LineValidator{
			LineNumber: corevalidator.LineNumber{LineNumber: -1},
			TextValidator: corevalidator.TextValidator{
				Search:    "hello",
				SearchAs:  stringcompareas.Contains,
				Condition: corevalidator.DefaultDisabledCoreCondition,
			},
		},
		corevalidator.LineValidator{
			LineNumber: corevalidator.LineNumber{LineNumber: -1},
			TextValidator: corevalidator.TextValidator{
				Search:    "world",
				SearchAs:  stringcompareas.Contains,
				Condition: corevalidator.DefaultDisabledCoreCondition,
			},
		},
	)

	text := "hello universe"
	result := lv.IsMatchText(text, true)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "'world' is not in 'hello universe', should return false", actual)

	diag := ValidatorDiffDiagnostics{
		CaseIndex: 0,
		Header:    "IsMatchText Diagnostic",
	}
	diag.PrintTextMatchDiagnostics(text, map[string]bool{
		"hello": strings.Contains(text, "hello"),
		"world": strings.Contains(text, "world"),
	})
}

// ==========================================
// errcore.ErrorToLinesLineDiff tests
// ==========================================

func Test_ErrorToLinesLineDiff_NilError(t *testing.T) {
	// Arrange

	// Assert
	expected := []string{"line1", "line2"}
	result := errcore.ErrorToLinesLineDiff(0, "nil error test", nil, expected)

	// Act
	actual := args.Map{"result": strings.Contains(result, "MISSING EXPECTED")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil error vs expected lines should show missing, got:\n", actual)

	fmt.Print(result)
}

func Test_ErrorToLinesLineDiff_WithError(t *testing.T) {
	// Arrange
	err := fmt.Errorf("error line 1\nerror line 2\nerror line 3")

	// Assert
	expected := []string{"error line 1", "error line 2", "DIFFERENT"}

	result := errcore.ErrorToLinesLineDiff(0, "error diff test", err, expected)

	// Act
	actual := args.Map{"result": strings.Contains(result, "MISMATCH")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "line 2 should be mismatch, got:\n", actual)
	actual := args.Map{"result": strings.Contains(result, "Line")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain line number labels", actual)

	fmt.Print(result)
}
