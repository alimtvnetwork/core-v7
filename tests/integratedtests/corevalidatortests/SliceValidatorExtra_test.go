package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// SliceValidator.AllVerifyErrorExceptLast
// ==========================================

func Test_SliceValidator_AllVerifyErrorExceptLast_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "different"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.AllVerifyErrorExceptLast(params)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "except last should pass:", actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.AllVerifyErrorQuick
// ==========================================

func Test_SliceValidator_AllVerifyErrorQuick_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	err := v.AllVerifyErrorQuick(0, "test", "a", "b")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should pass:", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Fail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	err := v.AllVerifyErrorQuick(0, "test", "a", "x")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.AllVerifyErrorTestCase
// ==========================================

func Test_SliceValidator_AllVerifyErrorTestCase_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	err := v.AllVerifyErrorTestCase(0, "test", true)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should pass:", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Fail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	err := v.AllVerifyErrorTestCase(0, "test", true)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.ComparingValidators caching
// ==========================================

func Test_SliceValidator_ComparingValidators_Cached(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	first := v.ComparingValidators()
	second := v.ComparingValidators()
	actual := args.Map{"result": first != second}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return same cached instance", actual)
	actual := args.Map{"result": first.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 validators", actual)
}

// ==========================================
// SliceValidator.ActualLinesString / ExpectingLinesString
// ==========================================

// (nil receiver tests migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_ActualLinesString_NonEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines: []string{"hello", "world"},
	}
	s := v.ActualLinesString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_SliceValidator_ExpectingLinesString_NonEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		ExpectedLines: []string{"hello", "world"},
	}
	s := v.ExpectingLinesString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

// ==========================================
// SliceValidator.IsUsedAlready — nil receiver
// ==========================================

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// NewSliceValidatorUsingErr — with actual error
// ==========================================

func Test_NewSliceValidatorUsingErr_WithError(t *testing.T) {
	err := errors.New("line1\nline2\nline3")
	v := corevalidator.NewSliceValidatorUsingErr(
		err, "line1\nline2\nline3",
		false, false, false,
		stringcompareas.Equal,
	)
	actual := args.Map{"result": v == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual := args.Map{"result": v.ActualLinesLength() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 actual lines", actual)
	actual := args.Map{"result": v.ExpectingLinesLength() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 expected lines", actual)
}

func Test_NewSliceValidatorUsingErr_WithConditions(t *testing.T) {
	err := errors.New("  hello  \n  world  ")
	v := corevalidator.NewSliceValidatorUsingErr(
		err, "hello\nworld",
		true, true, true,
		stringcompareas.Equal,
	)
	actual := args.Map{"result": v.IsTrimCompare}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsTrimCompare true", actual)
	actual := args.Map{"result": v.IsNonEmptyWhitespace}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsNonEmptyWhitespace true", actual)
	actual := args.Map{"result": v.IsSortStringsBySpace}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsSortStringsBySpace true", actual)
}

// ==========================================
// SliceValidator.UserInputsMergeWithError
// ==========================================

func Test_SliceValidator_UserInputsMergeWithError_NoAttach(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "test",
		IsAttachUserInputs: false,
	}
	testErr := errors.New("test error")
	result := v.UserInputsMergeWithError(params, testErr)
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
	actual := args.Map{"result": result.Error() != "test error"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "without attach, should return original error, got:", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_WithAttach(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "test",
		IsAttachUserInputs: true,
	}
	testErr := errors.New("test error")
	result := v.UserInputsMergeWithError(params, testErr)
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
	msg := result.Error()
	actual := args.Map{"result": msg == "test error"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "with attach, should include additional context", actual)
}

// ==========================================
// SliceValidator — isEmptyIgnoreCase boundary
// ==========================================

func Test_SliceValidator_AllVerifyError_EmptyActualNoSkip(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty actual without skip should return error", actual)
}

// ==========================================
// TextValidators.AddSimpleAllTrue
// ==========================================

func Test_TextValidators_AddSimpleAllTrue(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.AddSimpleAllTrue("hello", stringcompareas.Contains)
	actual := args.Map{"result": v.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add one validator", actual)
	item := v.Items[0]
	actual := args.Map{"result": item.IsTrimCompare}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsTrimCompare true", actual)
	actual := args.Map{"result": item.IsUniqueWordOnly}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsUniqueWordOnly true", actual)
	actual := args.Map{"result": item.IsNonEmptyWhitespace}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsNonEmptyWhitespace true", actual)
	actual := args.Map{"result": item.IsSortStringsBySpace}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have IsSortStringsBySpace true", actual)
}

// ==========================================
// TextValidators.AsBasicSliceContractsBinder
// ==========================================

func Test_TextValidators_AsBasicSliceContractsBinder(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	binder := v.AsBasicSliceContractsBinder()
	actual := args.Map{"result": binder == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ==========================================
// TextValidators.Count
// ==========================================

func Test_TextValidators_Count(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})
	if v.Count() != 1 { // Count = LastIndex = Length-1
		t.Errorf("expected Count=1 (LastIndex), got %d", v.Count())
	}
}

// ==========================================
// TextValidator.VerifySimpleError
// ==========================================

func Test_TextValidator_VerifySimpleError_Match(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifySimpleError(0, params, "hello")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "match should not error:", actual)
}

func Test_TextValidator_VerifySimpleError_Mismatch(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifySimpleError(0, params, "world")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.MethodName
// ==========================================

func Test_TextValidator_MethodName(t *testing.T) {
	v := corevalidator.TextValidator{SearchAs: stringcompareas.Contains}
	actual := args.Map{"result": v.MethodName() != "IsContains"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'IsContains', got ''", actual)
}

// ==========================================
// TextValidator.ToString
// ==========================================

func Test_TextValidator_ToString_SingleLine(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.ToString(true)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_TextValidator_ToString_MultiLine(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.ToString(false)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_TextValidator_String(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

// ==========================================
// TextValidator.GetCompiledTermBasedOnConditions
// ==========================================

func Test_TextValidator_GetCompiledTermBasedOnConditions_NoTrim(t *testing.T) {
	v := corevalidator.TextValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	result := v.GetCompiledTermBasedOnConditions("  hello  ", true)
	actual := args.Map{"result": result != "  hello  "}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no trim should return original, got ''", actual)
}

func Test_TextValidator_GetCompiledTermBasedOnConditions_WithTrim(t *testing.T) {
	v := corevalidator.TextValidator{
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	result := v.GetCompiledTermBasedOnConditions("  hello  ", true)
	actual := args.Map{"result": result != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "trim should return 'hello', got ''", actual)
}

// ==========================================
// TextValidators.VerifyFirstErrorMany
// ==========================================

func Test_TextValidators_VerifyFirstErrorMany_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.VerifyFirstErrorMany(params, "a")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty validators should return nil", actual)
}

func Test_TextValidators_VerifyFirstErrorMany_Pass(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "a",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyFirstErrorMany(params, "a")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should pass:", actual)
}

func Test_TextValidators_AllVerifyErrorMany_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.AllVerifyErrorMany(params, "a")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty validators should return nil", actual)
}

// ==========================================
// TextValidators.VerifyErrorMany — routing
// ==========================================

func Test_TextValidators_VerifyErrorMany_ContinueTrue(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyErrorMany(true, params, "a", "b")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatches should return error", actual)
}

func Test_TextValidators_VerifyErrorMany_ContinueFalse(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyErrorMany(false, params, "a", "b")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatches should return error", actual)
}

// ==========================================
// TextValidators.HasAnyItem
// ==========================================

func Test_TextValidators_HasAnyItem_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": v.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have items", actual)
}

func Test_TextValidators_HasAnyItem_NonEmpty(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	actual := args.Map{"result": v.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

// ==========================================
// TextValidators.String
// ==========================================

func Test_TextValidators_String(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	s := v.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}
