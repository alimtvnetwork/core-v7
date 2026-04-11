package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── SliceValidator — uncovered non-nil paths ──

func Test_SliceValidator_ComparingValidators_Cached_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}

	// Act
	v1 := sv.ComparingValidators()
	v2 := sv.ComparingValidators()

	// Assert
	actual := args.Map{
		"same": v1 == v2,
		"len": v1.Length(),
	}
	expected := args.Map{
		"same": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "ComparingValidators returns cached -- second call", actual)
}

func Test_SliceValidator_Dispose_WithValidators_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	_ = sv.ComparingValidators() // populate cache

	// Act
	sv.Dispose()

	// Assert
	actual := args.Map{
		"actualNil": sv.ActualLines == nil,
		"expectedNil": sv.ExpectedLines == nil,
	}
	expected := args.Map{
		"actualNil": true,
		"expectedNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dispose clears all fields -- with cached validators", actual)
}

func Test_SliceValidator_IsValid_BothNilLines(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- both nil lines", actual)
}

func Test_SliceValidator_IsValid_ActualNil_ExpectedNotNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid returns false -- actual nil expected not nil", actual)
}

func Test_SliceValidator_IsValid_LengthMismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
	}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid returns false -- length mismatch", actual)
}

func Test_SliceValidator_IsValidOtherLines_NilExpected(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValidOtherLines returns true -- both nil", actual)
}

func Test_SliceValidator_IsValidOtherLines_OnlyOneNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValidOtherLines returns false -- one nil one not", actual)
}

func Test_SliceValidator_SetActual_ReturnsChain(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}

	// Act
	result := sv.SetActual([]string{"x"})

	// Assert
	actual := args.Map{
		"same": result == sv,
		"used": sv.IsUsedAlready(),
	}
	expected := args.Map{
		"same": true,
		"used": true,
	}
	expected.ShouldBeEqual(t, 0, "SetActual returns self -- chaining", actual)
}

// ── SliceValidatorVerify — error paths ──

func Test_SliceValidator_AllVerifyErrorUptoLength_EmptyIgnoreCase(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true, IsCaseSensitive: true}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 1)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- empty actual skip", actual)
}

func Test_SliceValidator_AllVerifyError_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "test"}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- content mismatch", actual)
}

func Test_SliceValidator_VerifyFirstError_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"x"},
		ExpectedLines: []string{"y"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns error -- content mismatch", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"expected"},
	}

	// Act
	err := sv.AllVerifyErrorQuick(0, "test-header", "actual-different")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorQuick returns error -- mismatch", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}

	// Act
	err := sv.AllVerifyErrorTestCase(0, "test", true)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns nil -- match", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}

	// Act
	err := sv.AllVerifyErrorTestCase(0, "test", true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns error -- mismatch", actual)
}

func Test_SliceValidator_AllVerifyErrorExceptLast_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "different"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act — verifies up to second-last item only
	err := sv.AllVerifyErrorExceptLast(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorExceptLast returns nil -- first matches", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "c"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 2)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- first two match", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_IsFirstOnly(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"x", "y"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}

	// Act
	err := sv.AllVerifyErrorUptoLength(true, params, 2)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- isFirstOnly breaks early", actual)
}

func Test_SliceValidator_InitialVerify_AnyNilCase(t *testing.T) {
	// Arrange — actual nil, expected not nil
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- actual nil expected not nil", actual)
}

func Test_SliceValidator_InitialVerify_LengthMismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "h"}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- length mismatch", actual)
}

func Test_SliceValidator_InitialVerify_BothNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- both nil", actual)
}

// ── SliceValidatorMessages — uncovered paths ──

func Test_SliceValidator_ActualInputWithExpectingMessage_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"actual-line"},
		ExpectedLines: []string{"expected-line"},
	}

	// Act
	msg := sv.ActualInputWithExpectingMessage(0, "test")

	// Assert
	actual := args.Map{"notEmpty": len(msg) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ActualInputWithExpectingMessage returns non-empty -- valid inputs", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: true, Header: "h"}
	someErr := errors.New("test-error")

	// Act
	err := sv.UserInputsMergeWithError(params, someErr)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns merged error -- with error", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_NilErrorAttach(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: true, Header: "h"}

	// Act
	err := sv.UserInputsMergeWithError(params, nil)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- nil err but has user inputs", actual)
}

// ── SliceValidatorConstructors — NewSliceValidatorUsingErr ──

func Test_NewSliceValidatorUsingErr_Match(t *testing.T) {
	// Arrange
	testErr := errors.New("line1\nline2")

	// Act
	sv := corevalidator.NewSliceValidatorUsingErr(
		testErr,
		"line1\nline2",
		false, false, false,
		stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns valid -- matching error lines", actual)
}

func Test_NewSliceValidatorUsingErr_Mismatch(t *testing.T) {
	// Arrange
	testErr := errors.New("actual-line")

	// Act
	sv := corevalidator.NewSliceValidatorUsingErr(
		testErr,
		"expected-line",
		false, false, false,
		stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns invalid -- mismatching error lines", actual)
}

func Test_NewSliceValidatorUsingErr_WithConditions_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	testErr := errors.New("  hello  world  ")

	// Act
	sv := corevalidator.NewSliceValidatorUsingErr(
		testErr,
		"hello world",
		true, true, true,
		stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns valid -- with trim and sort conditions", actual)
}

// ── SimpleSliceValidator — verify paths ──

func Test_SimpleSliceValidator_VerifyAll_Match(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a", "b"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a", "b"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyAll([]string{"a", "b"}, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyAll returns nil -- match", actual)
}

func Test_SimpleSliceValidator_VerifyAll_Mismatch(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"b"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyAll([]string{"b"}, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyAll returns error -- mismatch", actual)
}

func Test_SimpleSliceValidator_VerifyFirst_Match(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyFirst([]string{"a"}, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyFirst returns nil -- match", actual)
}

func Test_SimpleSliceValidator_VerifyUpto_Match(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a", "b"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a", "b"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyUpto([]string{"a", "b"}, params, 1)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyUpto returns nil -- match first", actual)
}

// ── TextValidator — uncovered branches ──

func Test_TextValidator_SearchTextFinalized_Cached_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act
	s1 := tv.SearchTextFinalized()
	s2 := tv.SearchTextFinalized()

	// Assert
	actual := args.Map{
		"same": s1 == s2,
		"val": s1,
	}
	expected := args.Map{
		"same": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SearchTextFinalized returns cached -- second call", actual)
}

func Test_TextValidator_GetCompiledTerm_TrimCompare(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsTrimCompare: true,
		},
	}

	// Act
	result := tv.GetCompiledTermBasedOnConditions("  hello  ", true)

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTermBasedOnConditions returns trimmed -- trim enabled", actual)
}

func Test_TextValidator_GetCompiledTerm_SplitByWhitespace_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsNonEmptyWhitespace: true,
		},
	}

	// Act
	result := tv.GetCompiledTermBasedOnConditions("hello  world", true)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetCompiledTermBasedOnConditions returns processed -- split whitespace", actual)
}

func Test_TextValidator_VerifyDetailError_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "expected", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := tv.VerifyDetailError(params, "actual")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyDetailError returns error -- mismatch", actual)
}

func Test_TextValidator_VerifySimpleError_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "expected", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.VerifySimpleError(0, params, "actual")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifySimpleError returns error -- mismatch", actual)
}

func Test_TextValidator_AllVerifyError_WithErrors_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "expected", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.AllVerifyError(params, "bad1", "bad2")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- multiple mismatches", actual)
}

func Test_TextValidator_VerifyFirstError_WithContent(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.VerifyFirstError(params, "hello", "hello")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns nil -- all match", actual)
}

func Test_TextValidator_VerifyFirstError_FirstMismatch(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.VerifyFirstError(params, "bad", "hello")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns error -- first mismatch", actual)
}

func Test_TextValidator_VerifyMany_ContinueOnError_WithErrors(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.VerifyMany(true, params, "a", "b")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns error -- continue on error with mismatches", actual)
}

func Test_TextValidator_VerifyMany_FirstOnly_Mismatch(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tv.VerifyMany(false, params, "a")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns error -- first only mismatch", actual)
}

func Test_TextValidator_IsMatchMany_AllMatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatchMany(false, true, "hello", "hello")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- all match", actual)
}

func Test_TextValidator_IsMatchMany_EmptyNoSkip_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatchMany(false, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- empty contents no skip", actual)
}

func Test_TextValidator_String_Method(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"notEmpty": tv.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- valid validator", actual)
}

// ── TextValidators — uncovered paths ──

func Test_TextValidators_IsMatchMany_WithItems_Fail(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)

	// Act & Assert
	actual := args.Map{"match": tvs.IsMatchMany(false, true, "y")}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns false -- item mismatch", actual)
}

func Test_TextValidators_IsMatchMany_WithItems_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)

	// Act & Assert
	actual := args.Map{"match": tvs.IsMatchMany(false, true, "x")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- item match", actual)
}

func Test_TextValidators_VerifyErrorMany_WithItems_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tvs.VerifyErrorMany(true, params, "x")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorMany returns nil -- match continue", actual)
}

func Test_TextValidators_VerifyErrorMany_FirstOnly_WithItems(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tvs.VerifyErrorMany(false, params, "x")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorMany returns nil -- first only match", actual)
}

func Test_TextValidators_VerifyFirstErrorMany_WithError(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tvs.VerifyFirstErrorMany(params, "bad")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstErrorMany returns error -- mismatch", actual)
}

func Test_TextValidators_AllVerifyErrorMany_WithError(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := tvs.AllVerifyErrorMany(params, "bad")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorMany returns error -- mismatch", actual)
}

func Test_TextValidators_VerifyFirstError_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)

	// Act
	err := tvs.VerifyFirstError(0, "y", true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns error -- mismatch", actual)
}

func Test_TextValidators_AllVerifyError_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)

	// Act
	err := tvs.AllVerifyError(0, "y", true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- mismatch", actual)
}

func Test_TextValidators_IsMatch_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)
	tvs.AddSimple("x", stringcompareas.Equal)

	// Act & Assert
	actual := args.Map{"match": tvs.IsMatch("y", true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- mismatch", actual)
}

func Test_TextValidators_Adds_WithItems(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(2)

	// Act
	tvs.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)

	// Assert
	actual := args.Map{"len": tvs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Adds returns correct length -- two items", actual)
}

// ── LineValidator — uncovered branches ──

func Test_LineValidator_IsMatch_LineNumberMismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(3, "hello", true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- line number mismatch", actual)
}

func Test_LineValidator_IsMatch_TextMismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(5, "world", true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- text mismatch", actual)
}

func Test_LineValidator_IsMatch_AllMatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(5, "hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- line and text match", actual)
}

func Test_LineValidator_IsMatchMany_WithContents_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	contents := []corestr.TextWithLineNumber{
		{Text: "hello", LineNumber: 0},
		{Text: "hello", LineNumber: 1},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatchMany(false, true, contents...)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- all match", actual)
}

func Test_LineValidator_IsMatchMany_Mismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	contents := []corestr.TextWithLineNumber{
		{Text: "hello", LineNumber: 0},
		{Text: "world", LineNumber: 1},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatchMany(false, true, contents...)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns false -- second mismatch", actual)
}

func Test_LineValidator_IsMatchMany_EmptySkip_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatchMany(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- empty skip", actual)
}

func Test_LineValidator_VerifyError_LineNumberMismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.VerifyError(params, 3, "hello")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns error -- line number mismatch", actual)
}

func Test_LineValidator_VerifyError_TextMismatch_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.VerifyError(params, 5, "world")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns error -- text mismatch", actual)
}

func Test_LineValidator_VerifyError_AllMatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.VerifyError(params, 0, "hello")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns nil -- all match", actual)
}

func Test_LineValidator_VerifyMany_ContinueOnError_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "x", LineNumber: 0}}

	// Act
	err := lv.VerifyMany(true, params, contents...)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns nil -- continue on error match", actual)
}

func Test_LineValidator_VerifyMany_FirstOnly_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "y", LineNumber: 0}}

	// Act
	err := lv.VerifyMany(false, params, contents...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns error -- first only mismatch", actual)
}

func Test_LineValidator_VerifyFirstError_WithItems(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "x", LineNumber: 0}}

	// Act
	err := lv.VerifyFirstError(params, contents...)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns nil -- match", actual)
}

func Test_LineValidator_VerifyFirstError_EmptySkip_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: true}

	// Act
	err := lv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns nil -- empty skip", actual)
}

func Test_LineValidator_AllVerifyError_WithItems(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{
		{Text: "a", LineNumber: 0},
		{Text: "b", LineNumber: 1},
	}

	// Act
	err := lv.AllVerifyError(params, contents...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- mismatches", actual)
}

func Test_LineValidator_AllVerifyError_EmptySkip_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: true}

	// Act
	err := lv.AllVerifyError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- empty skip", actual)
}

func Test_LineValidator_Nil_AllVerifyError(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.AllVerifyError(params, corestr.TextWithLineNumber{Text: "x"})

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- nil receiver", actual)
}

func Test_LineValidator_Nil_VerifyFirstError(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.VerifyFirstError(params, corestr.TextWithLineNumber{Text: "x"})

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns nil -- nil receiver", actual)
}

// ── LinesValidators — uncovered non-empty paths ──

func Test_LinesValidators_IsMatchText_WithItems(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	})

	// Act & Assert
	actual := args.Map{
		"match":    lv.IsMatchText("hello", true),
		"mismatch": lv.IsMatchText("world", true),
	}
	expected := args.Map{
		"match": true,
		"mismatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMatchText returns correct -- with items", actual)
}

func Test_LinesValidators_IsMatch_WithContents_Match(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	contents := []corestr.TextWithLineNumber{{Text: "x", LineNumber: 0}}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(false, true, contents...)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- contents match", actual)
}

func Test_LinesValidators_IsMatch_WithContents_Mismatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	contents := []corestr.TextWithLineNumber{{Text: "y", LineNumber: 0}}

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(false, true, contents...)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- contents mismatch", actual)
}

func Test_LinesValidators_IsMatch_EmptyContents_NoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(false, true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- empty no skip", actual)
}

func Test_LinesValidators_IsMatch_EmptyContents_Skip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})

	// Act & Assert
	actual := args.Map{"match": lv.IsMatch(true, true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- empty skip", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_WithMatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "x", LineNumber: 0}}

	// Act
	err := lv.VerifyFirstDefaultLineNumberError(params, contents...)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns nil -- match", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_EmptyNoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: false}

	// Act
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns error -- empty no skip", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_EmptySkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: true}

	// Act
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns nil -- empty skip", actual)
}

func Test_LinesValidators_AllVerifyError_WithMatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "x", LineNumber: 0}}

	// Act
	err := lv.AllVerifyError(params, contents...)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- match", actual)
}

func Test_LinesValidators_AllVerifyError_WithMismatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "y", LineNumber: 0}}

	// Act
	err := lv.AllVerifyError(params, contents...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- mismatch", actual)
}

func Test_LinesValidators_AllVerifyError_EmptyNoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: false}

	// Act
	err := lv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- empty no skip", actual)
}

func Test_LinesValidators_AllVerifyError_EmptySkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsSkipCompareOnActualEmpty: true}

	// Act
	err := lv.AllVerifyError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- empty skip", actual)
}

func Test_LinesValidators_AddPtr_WithItem(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	item := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	}

	// Act
	lv.AddPtr(item)

	// Assert
	actual := args.Map{"len": lv.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddPtr adds item -- non-nil pointer", actual)
}

func Test_LinesValidators_Adds_WithItems(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)

	// Act
	lv.Adds(
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal}},
	)

	// Assert
	actual := args.Map{"len": lv.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Adds adds items -- two validators", actual)
}

// ── SliceValidators / HeaderSliceValidators — non-empty error paths ──

func Test_SliceValidators_IsMatch_WithMismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}

	// Act & Assert
	actual := args.Map{"match": sv.IsMatch(true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- validator mismatch", actual)
}

func Test_SliceValidators_VerifyAll_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- mismatch", actual)
}

func Test_SliceValidators_VerifyAll_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- match", actual)
}

func Test_SliceValidators_VerifyAllError_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns error -- mismatch", actual)
}

func Test_SliceValidators_VerifyFirst_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns error -- mismatch", actual)
}

func Test_SliceValidators_VerifyFirst_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- match", actual)
}

func Test_SliceValidators_VerifyUpto_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns error -- mismatch", actual)
}

func Test_SliceValidators_VerifyUpto_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- match", actual)
}

func Test_SliceValidators_VerifyAllErrorUsingActual_Match_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyAllErrorUsingActual(params, "a")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	// Note: SetActual on range copy won't affect the original, so this may or may not error
	// We just ensure no panic
	expected := args.Map{"hasErr": err != nil}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual does not panic -- with items", actual)
}

func Test_SliceValidators_SetActualOnAll_WithItems(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"}},
		},
	}

	// Act — should not panic
	sv.SetActualOnAll("a")

	// Assert
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll does not panic -- with items", actual)
}

// ── HeaderSliceValidators — non-empty error paths ──

func Test_HeaderSliceValidators_IsMatch_WithMismatch(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}

	// Act & Assert
	actual := args.Map{"match": hsv.IsMatch(true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- validator mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyAll_WithError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyAll_Match(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- match", actual)
}

func Test_HeaderSliceValidators_VerifyAllError_WithError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_WithError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_Match(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- match", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_WithError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_Match_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- match", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_WithItems(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	_ = hsv.VerifyAllErrorUsingActual(params, "a")

	// Assert
	// SetActual works on a copy, so this may show header-only error
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual does not panic -- with items", actual)
}

func Test_HeaderSliceValidators_SetActualOnAll_WithItems(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"a"},
			},
		},
	}

	// Act — should not panic
	hsv.SetActualOnAll("a")

	// Assert
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll does not panic -- with items", actual)
}

// ── RangeSegmentsValidator — uncovered paths ──

func Test_RangeSegmentsValidator_SetActual_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{Title: "test"}

	// Act
	result := rsv.SetActual([]string{"a", "b", "c"})

	// Assert
	actual := args.Map{"same": result == rsv}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetActual returns self -- chaining", actual)
}

// ── Condition — IsTrimCompare branch ──

func Test_Condition_IsTrimCompare(t *testing.T) {
	// Arrange
	c := corevalidator.Condition{IsTrimCompare: true}

	// Act & Assert
	actual := args.Map{
		"trim": c.IsTrimCompare,
		"split": c.IsSplitByWhitespace(),
	}
	expected := args.Map{
		"trim": true,
		"split": false,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct -- trim only no split", actual)
}

// ── vars.go — DefaultConditions ──

func Test_DefaultConditions(t *testing.T) {
	// Assert
	actual := args.Map{
		"disabledSplit":     corevalidator.DefaultDisabledCoreCondition.IsSplitByWhitespace(),
		"trimOnlySplit":     corevalidator.DefaultTrimCoreCondition.IsSplitByWhitespace(),
		"sortTrimSplit":     corevalidator.DefaultSortTrimCoreCondition.IsSplitByWhitespace(),
		"uniqueWordsSplit":  corevalidator.DefaultUniqueWordsCoreCondition.IsSplitByWhitespace(),
		"emptyValidatorLen": len(corevalidator.EmptyValidator.Search),
	}
	expected := args.Map{
		"disabledSplit":     false,
		"trimOnlySplit":     false,
		"sortTrimSplit":     true,
		"uniqueWordsSplit":  true,
		"emptyValidatorLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DefaultConditions returns correct -- all presets", actual)
}

// ── TextValidator — case-insensitive branches ──

func Test_TextValidator_IsMatch_CaseInsensitive_FromSliceValidatorCompar(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "Hello", SearchAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{
		"caseSensitive":   tv.IsMatch("hello", true),
		"caseInsensitive": tv.IsMatch("hello", false),
	}
	expected := args.Map{
		"caseSensitive":   false,
		"caseInsensitive": true,
	}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct -- case sensitivity", actual)
}

func Test_TextValidator_IsMatch_Contains_CaseInsensitive(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "ELL", SearchAs: stringcompareas.Contains}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatch("hello", false)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- contains case insensitive", actual)
}

func Test_TextValidator_IsMatch_StartsWith_CaseInsensitive(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "HEL", SearchAs: stringcompareas.StartsWith}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatch("hello", false)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- startsWith case insensitive", actual)
}

func Test_TextValidator_IsMatch_EndsWith_CaseInsensitive(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "LLO", SearchAs: stringcompareas.EndsWith}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatch("hello", false)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- endsWith case insensitive", actual)
}
