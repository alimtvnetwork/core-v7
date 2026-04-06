package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Condition
// =============================================================================

func Test_Condition_IsSplitByWhitespace_AllFalse_Cov(t *testing.T) {
	c := corevalidator.Condition{}
	actual := args.Map{"result": c.IsSplitByWhitespace()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all false should return false", actual)
}

func Test_Condition_IsSplitByWhitespace_UniqueWord(t *testing.T) {
	c := corevalidator.Condition{IsUniqueWordOnly: true}
	actual := args.Map{"result": c.IsSplitByWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unique word should return true", actual)
}

func Test_Condition_IsSplitByWhitespace_NonEmpty(t *testing.T) {
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}
	actual := args.Map{"result": c.IsSplitByWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "non-empty whitespace should return true", actual)
}

func Test_Condition_IsSplitByWhitespace_Sort(t *testing.T) {
	c := corevalidator.Condition{IsSortStringsBySpace: true}
	actual := args.Map{"result": c.IsSplitByWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "sort should return true", actual)
}

// =============================================================================
// Parameter
// =============================================================================

func Test_Parameter_IsIgnoreCase_Cov(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: true}
	actual := args.Map{"result": p.IsIgnoreCase()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case sensitive should not ignore case", actual)
	p2 := corevalidator.Parameter{IsCaseSensitive: false}
	actual := args.Map{"result": p2.IsIgnoreCase()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "not case sensitive should ignore case", actual)
}

// =============================================================================
// LineNumber
// =============================================================================

func Test_LineNumber_HasLineNumber_Cov(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"result": ln.HasLineNumber()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have line number", actual)
}

func Test_LineNumber_HasLineNumber_Invalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	actual := args.Map{"result": ln.HasLineNumber()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "invalid should not have line number", actual)
}

func Test_LineNumber_IsMatch_BothInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	actual := args.Map{"result": ln.IsMatch(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both invalid should match", actual)
}

func Test_LineNumber_IsMatch_InputInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"result": ln.IsMatch(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "invalid input should match", actual)
}

func Test_LineNumber_IsMatch_Exact(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"result": ln.IsMatch(5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match", actual)
	actual := args.Map{"result": ln.IsMatch(3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}

func Test_LineNumber_VerifyError_Match(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"result": ln.VerifyError(5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return nil", actual)
}

func Test_LineNumber_VerifyError_Mismatch(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	actual := args.Map{"result": ln.VerifyError(3) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// =============================================================================
// TextValidator — uncovered branches
// =============================================================================

func Test_TextValidator_ToString_MultiLine_Cov(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "test",
		SearchAs: stringcompareas.Equal,
	}
	str := tv.ToString(false)
	actual := args.Map{"result": str == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TextValidator_IsMatchMany_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	actual := args.Map{"result": tv.IsMatchMany(false, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should return true", actual)
}

func Test_TextValidator_IsMatchMany_EmptySkip_Cov(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	actual := args.Map{"result": tv.IsMatchMany(true, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty contents with skip should return true", actual)
}

func Test_TextValidator_IsMatchMany_Fail(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	actual := args.Map{"result": tv.IsMatchMany(false, true, "y")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return false", actual)
}

func Test_TextValidator_VerifyDetailError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	actual := args.Map{"result": tv.VerifyDetailError(params, "content") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidator_VerifySimpleError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	actual := args.Map{"result": tv.VerifySimpleError(0, params, "content") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidator_VerifyMany_FirstOnly_Cov(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(false, params, "x", "y")
	// first only, stops on first error if any
	_ = err
}

func Test_TextValidator_VerifyMany_ContinueOnError(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(true, params, "x")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return nil", actual)
}

func Test_TextValidator_VerifyFirstError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	actual := args.Map{"result": tv.VerifyFirstError(params, "x") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidator_VerifyFirstError_EmptySkip_Cov(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	actual := args.Map{"result": tv.VerifyFirstError(params) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty with skip should return nil", actual)
}

func Test_TextValidator_AllVerifyError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	actual := args.Map{"result": tv.AllVerifyError(params, "x") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidator_AllVerifyError_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	actual := args.Map{"result": tv.AllVerifyError(params) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty with skip should return nil", actual)
}

func Test_TextValidator_MethodName_Cov(t *testing.T) {
	tv := corevalidator.TextValidator{SearchAs: stringcompareas.StartsWith}
	actual := args.Map{"result": tv.MethodName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return method name", actual)
}

// =============================================================================
// TextValidators — uncovered branches
// =============================================================================

func Test_TextValidators_Count_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	tvs.Add(corevalidator.TextValidator{})
	if tvs.Count() != 1 { // Count returns LastIndex
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "count should equal lastindex", actual)
	}
}

func Test_TextValidators_Adds_Empty_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds()
	actual := args.Map{"result": tvs.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_TextValidators_AddSimple_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("test", stringcompareas.Equal)
	actual := args.Map{"result": tvs.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1", actual)
}

func Test_TextValidators_AddSimpleAllTrue_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimpleAllTrue("test", stringcompareas.Equal)
	actual := args.Map{"result": tvs.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1", actual)
}

func Test_TextValidators_HasAnyItem(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	actual := args.Map{"result": tvs.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
	tvs.Add(corevalidator.TextValidator{})
	actual := args.Map{"result": tvs.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_TextValidators_HasIndex_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	actual := args.Map{"result": tvs.HasIndex(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have index 0", actual)
	actual := args.Map{"result": tvs.HasIndex(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have index 5", actual)
}

func Test_TextValidators_String_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal})
	actual := args.Map{"result": tvs.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_TextValidators_IsMatch_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.IsMatch("anything", true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

func Test_TextValidators_IsMatchMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.IsMatchMany(true, true, "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

func Test_TextValidators_VerifyFirstError_Empty_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.VerifyFirstError(0, "x", true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_TextValidators_VerifyErrorMany_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	actual := args.Map{"result": tvs.VerifyErrorMany(true, &corevalidator.Parameter{}, "x") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidators_VerifyErrorMany_FirstOnly(t *testing.T) {
	var tvs *corevalidator.TextValidators
	actual := args.Map{"result": tvs.VerifyErrorMany(false, &corevalidator.Parameter{}, "x") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TextValidators_VerifyFirstErrorMany_Empty_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.VerifyFirstErrorMany(&corevalidator.Parameter{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_TextValidators_AllVerifyErrorMany_Empty_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.AllVerifyErrorMany(&corevalidator.Parameter{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_TextValidators_AllVerifyError_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.AllVerifyError(0, "x", true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_TextValidators_Dispose_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	tvs.Dispose()
	actual := args.Map{"result": tvs.Items != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should dispose", actual)
}

func Test_TextValidators_Dispose_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	tvs.Dispose() // should not panic
}

func Test_TextValidators_Length_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	actual := args.Map{"result": tvs.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_TextValidators_AsBasicSliceContractsBinder_Cov(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	actual := args.Map{"result": tvs.AsBasicSliceContractsBinder() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return self", actual)
}

// =============================================================================
// SliceValidator — uncovered branches
// =============================================================================

func Test_SliceValidator_IsUsedAlready_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.IsUsedAlready()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_SliceValidator_ActualLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.ActualLinesLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_SliceValidator_ActualLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.ActualLinesString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_SliceValidator_ExpectingLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.ExpectingLinesString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_SliceValidator_ExpectingLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.ExpectingLinesLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_SliceValidator_IsValid_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should return true", actual)
}

func Test_SliceValidator_Dispose_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	sv.Dispose() // should not panic
}

func Test_SliceValidator_SetActualVsExpected_Cov(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	actual := args.Map{"result": sv.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_SliceValidator_MethodName_Cov(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.StartsWith}
	actual := args.Map{"result": sv.MethodName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return method name", actual)
}

func Test_SliceValidator_VerifyFirstError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.VerifyFirstError(&corevalidator.Parameter{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_AllVerifyError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.AllVerifyError(&corevalidator.Parameter{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.AllVerifyErrorQuick(0, "header", "a") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_AllVerifyErrorExceptLast_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{}) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.AllVerifyErrorTestCase(0, "header", true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_VerifyFirstLengthUptoError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"result": sv.VerifyFirstLengthUptoError(&corevalidator.Parameter{}, 5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SliceValidator_IsValidOtherLines(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"hello"},
	}
	actual := args.Map{"result": sv.IsValidOtherLines(true, []string{"hello"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

// =============================================================================
// SliceValidators — uncovered branches
// =============================================================================

func Test_SliceValidators_Length_Nil(t *testing.T) {
	var svs *corevalidator.SliceValidators
	actual := args.Map{"result": svs.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_SliceValidators_IsEmpty_Nil(t *testing.T) {
	var svs *corevalidator.SliceValidators
	actual := args.Map{"result": svs.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_SliceValidators_SetActualOnAll_Empty_Cov(t *testing.T) {
	var svs *corevalidator.SliceValidators
	svs.SetActualOnAll("a") // should not panic
}

func Test_SliceValidators_IsValid_Empty_Cov(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	actual := args.Map{"result": svs.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// =============================================================================
// HeaderSliceValidators — uncovered branches
// =============================================================================

func Test_HeaderSliceValidators_Length_Nil(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	actual := args.Map{"result": hsv.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_HeaderSliceValidators_IsEmpty_Nil(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	actual := args.Map{"result": hsv.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_HeaderSliceValidators_SetActualOnAll_Empty_Cov(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	hsv.SetActualOnAll("a") // should not panic
}

func Test_HeaderSliceValidators_IsValid_Empty(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	actual := args.Map{"result": hsv.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// =============================================================================
// LinesValidators — uncovered branches
// =============================================================================

func Test_LinesValidators_Length_Nil(t *testing.T) {
	var lv *corevalidator.LinesValidators
	actual := args.Map{"result": lv.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_LinesValidators_Count_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"result": lv.Count() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return 0", actual)
}

func Test_LinesValidators_HasAnyItem(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"result": lv.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have items", actual)
}

func Test_LinesValidators_HasIndex_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	actual := args.Map{"result": lv.HasIndex(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have index 0", actual)
}

func Test_LinesValidators_AddPtr_Nil_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.AddPtr(nil)
	actual := args.Map{"result": lv.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not add", actual)
}

func Test_LinesValidators_String_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	_ = lv.String() // should not panic
}

func Test_LinesValidators_AsBasicSliceContractsBinder_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	actual := args.Map{"result": lv.AsBasicSliceContractsBinder() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return self", actual)
}

func Test_LinesValidators_IsMatchText_Empty_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	actual := args.Map{"result": lv.IsMatchText("test", true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

func Test_LinesValidators_IsMatch_Empty_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	actual := args.Map{"result": lv.IsMatch(false, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
}

// =============================================================================
// BaseLinesValidators — uncovered branches
// =============================================================================

func Test_BaseLinesValidators_Nil(t *testing.T) {
	var blv *corevalidator.BaseLinesValidators
	actual := args.Map{"result": blv.LinesValidatorsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_BaseLinesValidators_IsEmpty(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	actual := args.Map{"result": blv.IsEmptyLinesValidators()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_BaseLinesValidators_HasLinesValidators(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	actual := args.Map{"result": blv.HasLinesValidators()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have validators", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_Empty_Cov(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	lv := blv.ToLinesValidators()
	actual := args.Map{"result": lv.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return empty", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_WithItems(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}},
		},
	}
	lv := blv.ToLinesValidators()
	actual := args.Map{"result": lv.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 1", actual)
}

// =============================================================================
// BaseValidatorCoreCondition — uncovered branches
// =============================================================================

func Test_BaseValidatorCoreCondition_Default_NilCondition(t *testing.T) {
	bvc := &corevalidator.BaseValidatorCoreCondition{}
	c := bvc.ValidatorCoreConditionDefault()
	_ = c // should not panic
}

func Test_BaseValidatorCoreCondition_Default_NonNilCondition(t *testing.T) {
	cond := &corevalidator.Condition{IsTrimCompare: true}
	bvc := &corevalidator.BaseValidatorCoreCondition{ValidatorCoreCondition: cond}
	c := bvc.ValidatorCoreConditionDefault()
	actual := args.Map{"result": c.IsTrimCompare}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should use existing condition", actual)
}

// =============================================================================
// NewSliceValidatorUsingErr / NewSliceValidatorUsingAny
// =============================================================================

func Test_NewSliceValidatorUsingAny(t *testing.T) {
	sv := corevalidator.NewSliceValidatorUsingAny(
		"hello",
		"hello",
		false, false, false,
		stringcompareas.Equal,
	)
	actual := args.Map{"result": sv.IsValid(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

// =============================================================================
// SimpleSliceValidator — uncovered branches
// =============================================================================

func Test_SimpleSliceValidator_VerifyFirst_Cov(t *testing.T) {
	// Exercise the VerifyFirst path
	sv := &corevalidator.SimpleSliceValidator{
		CompareAs: stringcompareas.Equal,
	}
	// set expected via the Expected field is needed but requires corestr import
	// Just exercise the SetActual path
	sv.SetActual([]string{"a"})
	_ = sv
}

// =============================================================================
// RangeSegmentsValidator — uncovered branches
// =============================================================================

func Test_RangeSegmentsValidator_LengthOfVerifierSegments_Cov(t *testing.T) {
	rsv := &corevalidator.RangeSegmentsValidator{}
	actual := args.Map{"result": rsv.LengthOfVerifierSegments() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 0", actual)
}
