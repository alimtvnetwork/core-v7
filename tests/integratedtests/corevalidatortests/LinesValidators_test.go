package corevalidatortests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// LinesValidators — collection basics
// ==========================================

func Test_LinesValidators_Count(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	actual := args.Map{"result": lv.Count() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinesValidators_LastIndex(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	lv.Add(corevalidator.LineValidator{})
	actual := args.Map{"result": lv.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinesValidators_Adds(t *testing.T) {
	lv := corevalidator.NewLinesValidators(3)
	lv.Adds(
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
	)
	actual := args.Map{"result": lv.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_LinesValidators_String(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:   "test",
			SearchAs: stringcompareas.Equal,
		},
	})
	s := lv.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
}

// ==========================================
// LinesValidators.IsMatch (with contents)
// ==========================================

func Test_LinesValidators_IsMatch_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	items := []corestr.TextWithLineNumber{
		{Text: "hello", LineNumber: 0},
	}
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty validators should match", actual)
}

func Test_LinesValidators_IsMatch_NoContentsSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	actual := args.Map{"result": lv.IsMatch(true, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "no contents with skip should match", actual)
}

func Test_LinesValidators_IsMatch_NoContentsNoSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	actual := args.Map{"result": lv.IsMatch(false, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no contents without skip should not match", actual)
}

func Test_LinesValidators_IsMatch_AllMatch(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all matching should return true", actual)
}

func Test_LinesValidators_IsMatch_OneFails(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "nope", LineNumber: 1},
	}
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one failing should return false", actual)
}

// ==========================================
// LinesValidators.VerifyFirstDefaultLineNumberError
// ==========================================

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_SkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip empty should return nil:", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_NoSkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty contents without skip should return error", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Pass(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "match should pass:", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Fail(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// ==========================================
// LinesValidators.AllVerifyError
// ==========================================

func Test_LinesValidators_AllVerifyError_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.AllVerifyError(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_LinesValidators_AllVerifyError_SkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.AllVerifyError(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip empty should return nil:", actual)
}

func Test_LinesValidators_AllVerifyError_NoSkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty contents without skip should return error", actual)
}

func Test_LinesValidators_AllVerifyError_Pass(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "match should pass:", actual)
}

func Test_LinesValidators_AllVerifyError_Fail(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

func Test_LineValidator_AllVerifyError_CollectsMultipleErrors(t *testing.T) {
	// Arrange: validator expects "ok", but all 3 inputs are wrong
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad1", LineNumber: 0},
		{Text: "bad2", LineNumber: 1},
		{Text: "bad3", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: error should be non-nil and contain all 3 failures
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllVerifyError should return error when all items fail", actual)

	errMsg := err.Error()
	for _, expected := range []string{"bad1", "bad2", "bad3"} {
		actual := args.Map{"result": strings.Contains(errMsg, expected)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "AllVerifyError should collect all errors, missing '' in:\n", actual)
	}
}

func Test_LineValidator_AllVerifyError_FirstFailOthersPass(t *testing.T) {
	// Arrange: validator expects "ok", first fails, rest pass
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
		{Text: "ok", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: should still report the one failure
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllVerifyError should return error when any item fails", actual)

	errMsg := err.Error()
	actual := args.Map{"result": strings.Contains(errMsg, "bad")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "error should mention the failed content 'bad', got:\n", actual)
}

// ==========================================
// LinesValidators.AsBasicSliceContractsBinder
// ==========================================

func Test_LinesValidators_AsBasicSliceContractsBinder(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	binder := lv.AsBasicSliceContractsBinder()
	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}
