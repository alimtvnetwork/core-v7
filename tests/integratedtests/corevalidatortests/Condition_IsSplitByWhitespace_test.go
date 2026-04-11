package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// Condition
// ══════════════════════════════════════════════════════════════════════════════

func Test_Condition_IsSplitByWhitespace_AllFalse(t *testing.T) {
	// Arrange
	c := &corevalidator.Condition{}

	// Act
	actual := args.Map{"split": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"split": false}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns non-empty -- all false", actual)
}

func Test_Condition_IsSplitByWhitespace_UniqueWord_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	c := &corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"split": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns correct value -- unique word", actual)
}

func Test_Condition_IsSplitByWhitespace_NonEmpty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	c := &corevalidator.Condition{IsNonEmptyWhitespace: true}

	// Act
	actual := args.Map{"split": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns empty -- non-empty", actual)
}

func Test_Condition_IsSplitByWhitespace_Sort_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	c := &corevalidator.Condition{IsSortStringsBySpace: true}

	// Act
	actual := args.Map{"split": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns correct value -- sort", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Parameter
// ══════════════════════════════════════════════════════════════════════════════

func Test_Parameter_IsIgnoreCase_Sensitive(t *testing.T) {
	// Arrange
	p := corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	actual := args.Map{"ignoreCase": p.IsIgnoreCase()}

	// Assert
	expected := args.Map{"ignoreCase": false}
	expected.ShouldBeEqual(t, 0, "Parameter.IsIgnoreCase returns correct value -- sensitive", actual)
}

func Test_Parameter_IsIgnoreCase_Insensitive(t *testing.T) {
	// Arrange
	p := corevalidator.Parameter{IsCaseSensitive: false}

	// Act
	actual := args.Map{"ignoreCase": p.IsIgnoreCase()}

	// Assert
	expected := args.Map{"ignoreCase": true}
	expected.ShouldBeEqual(t, 0, "Parameter.IsIgnoreCase returns correct value -- insensitive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_LineNumber_HasLineNumber_Valid(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{"has": ln.HasLineNumber()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.HasLineNumber returns non-empty -- valid", actual)
}

func Test_LineNumber_HasLineNumber_Invalid_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: -1}

	// Act
	actual := args.Map{"has": ln.HasLineNumber()}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "LineNumber.HasLineNumber returns error -- invalid", actual)
}

func Test_LineNumber_IsMatch_BothInvalid_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: -1}

	// Act
	actual := args.Map{"match": ln.IsMatch(-1)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns error -- both invalid", actual)
}

func Test_LineNumber_IsMatch_Same(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{"match": ln.IsMatch(5)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns correct value -- same", actual)
}

func Test_LineNumber_IsMatch_Different(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{"match": ln.IsMatch(3)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns correct value -- different", actual)
}

func Test_LineNumber_IsMatch_InputInvalid_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{"match": ln.IsMatch(-1)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.IsMatch returns error -- input invalid", actual)
}

func Test_LineNumber_VerifyError_Match_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}
	err := ln.VerifyError(5)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.VerifyError returns error -- match", actual)
}

func Test_LineNumber_VerifyError_Mismatch_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	ln := &corevalidator.LineNumber{LineNumber: 5}
	err := ln.VerifyError(3)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineNumber.VerifyError returns error -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidator_ToString_SingleLine(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	result := tv.ToString(true)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.ToString returns non-empty -- single", actual)
}

func Test_TextValidator_ToString_MultiLine(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	result := tv.ToString(false)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.ToString returns non-empty -- multi", actual)
}

func Test_TextValidator_String(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}
	result := tv.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.String returns non-empty -- with args", actual)
}

func Test_TextValidator_SearchTextFinalized(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: " hello ", SearchAs: stringcompareas.Equal, Condition: corevalidator.Condition{IsTrimCompare: true}}
	result := tv.SearchTextFinalized()

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TextValidator.SearchTextFinalized returns non-empty -- trimmed", actual)
}

func Test_TextValidator_SearchTextFinalized_Cached(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	r1 := tv.SearchTextFinalizedPtr()
	r2 := tv.SearchTextFinalizedPtr()

	// Act
	actual := args.Map{"same": r1 == r2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.SearchTextFinalized returns non-empty -- cached", actual)
}

func Test_TextValidator_GetCompiledTerm_NoConditions(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{}
	result := tv.GetCompiledTermBasedOnConditions("hello world", false)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns empty -- no conditions", actual)
}

func Test_TextValidator_GetCompiledTerm_Trim(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Condition: corevalidator.Condition{IsTrimCompare: true}}
	result := tv.GetCompiledTermBasedOnConditions("  hello  ", false)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns correct value -- trim", actual)
}

func Test_TextValidator_GetCompiledTerm_SplitByWhitespace(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Condition: corevalidator.Condition{IsNonEmptyWhitespace: true, IsSortStringsBySpace: true}}
	result := tv.GetCompiledTermBasedOnConditions("b a", false)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a b"}
	expected.ShouldBeEqual(t, 0, "GetCompiledTerm returns correct value -- split whitespace", actual)
}

func Test_TextValidator_IsMatch_Equal(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{
		"match": tv.IsMatch("hello", true),
		"noMatch": tv.IsMatch("world", true),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- equal", actual)
}

func Test_TextValidator_IsMatch_Contains(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "ell", SearchAs: stringcompareas.Contains}

	// Act
	actual := args.Map{"match": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- contains", actual)
}

func Test_TextValidator_IsMatchMany_NilReceiver(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	actual := args.Map{"match": tv.IsMatchMany(true, true, "a", "b")}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns nil -- nil", actual)
}

func Test_TextValidator_IsMatchMany_EmptySkip(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"match": tv.IsMatchMany(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns empty -- empty skip", actual)
}

func Test_TextValidator_IsMatchMany_EmptyNoSkip(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"match": tv.IsMatchMany(false, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns empty -- empty no skip", actual)
}

func Test_TextValidator_IsMatchMany_AllMatch(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"match": tv.IsMatchMany(false, true, "a", "a")}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns non-empty -- all match", actual)
}

func Test_TextValidator_IsMatchMany_OneFails(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"match": tv.IsMatchMany(false, true, "a", "b")}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatchMany returns non-empty -- one fails", actual)
}

func Test_TextValidator_MethodName(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{SearchAs: stringcompareas.StartsWith}

	// Act
	actual := args.Map{"name": tv.MethodName()}

	// Assert
	expected := args.Map{"name": "StartsWith"}
	expected.ShouldBeEqual(t, 0, "TextValidator.MethodName returns non-empty -- with args", actual)
}

func Test_TextValidator_VerifyDetailError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator
	err := tv.VerifyDetailError(&corevalidator.Parameter{}, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns nil -- nil", actual)
}

func Test_TextValidator_VerifyDetailError_Match(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifyDetailError(&corevalidator.Parameter{IsCaseSensitive: true}, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns error -- match", actual)
}

func Test_TextValidator_VerifyDetailError_Mismatch(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifyDetailError(&corevalidator.Parameter{IsCaseSensitive: true}, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyDetailError returns error -- mismatch", actual)
}

func Test_TextValidator_VerifySimpleError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{}, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns nil -- nil", actual)
}

func Test_TextValidator_VerifySimpleError_Match(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{IsCaseSensitive: true}, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns error -- match", actual)
}

func Test_TextValidator_VerifySimpleError_Mismatch(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	err := tv.VerifySimpleError(0, &corevalidator.Parameter{IsCaseSensitive: true}, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifySimpleError returns error -- mismatch", actual)
}

func Test_TextValidator_VerifyMany_ContinueOnError_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(true, params, "a", "b")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyMany returns non-empty -- continue", actual)
}

func Test_TextValidator_VerifyMany_StopOnFirst(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(false, params, "a", "b")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyMany returns non-empty -- stop first", actual)
}

func Test_TextValidator_VerifyFirstError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator
	err := tv.VerifyFirstError(&corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_TextValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.VerifyFirstError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.VerifyFirstError returns empty -- empty skip", actual)
}

func Test_TextValidator_AllVerifyError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator
	err := tv.AllVerifyError(&corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_TextValidator_AllVerifyError_EmptySkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.AllVerifyError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns empty -- empty skip", actual)
}

func Test_TextValidator_AllVerifyError_WithErrors(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}
	err := tv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true}, "a", "b", "c")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.AllVerifyError returns error -- with errors", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidators_NilLength(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- nil length", actual)
}

func Test_TextValidators_Count(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{"count": tvs.Count()}

	// Assert
	expected := args.Map{"count": 0} // Count = LastIndex = Length - 1
	expected.ShouldBeEqual(t, 0, "TextValidators.Count returns non-empty -- with args", actual)
}

func Test_TextValidators_Adds(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TextValidators.Adds returns non-empty -- with args", actual)
}

func Test_TextValidators_Adds_Empty(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds()

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators.Adds returns empty -- empty", actual)
}

func Test_TextValidators_AddSimple(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TextValidators.AddSimple returns non-empty -- with args", actual)
}

func Test_TextValidators_AddSimpleAllTrue(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimpleAllTrue("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{
		"len": tvs.Length(),
		"hasAny": tvs.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators.AddSimpleAllTrue returns non-empty -- with args", actual)
}

func Test_TextValidators_HasIndex(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a"})

	// Act
	actual := args.Map{
		"has0": tvs.HasIndex(0),
		"has1": tvs.HasIndex(1),
	}

	// Assert
	expected := args.Map{
		"has0": true,
		"has1": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators.HasIndex returns non-empty -- with args", actual)
}

func Test_TextValidators_String(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	result := tvs.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.String returns non-empty -- with args", actual)
}

func Test_TextValidators_IsMatch_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)

	// Act
	actual := args.Map{"match": tvs.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns empty -- empty", actual)
}

func Test_TextValidators_IsMatch_AllPass(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hel", SearchAs: stringcompareas.StartsWith})
	tvs.Add(corevalidator.TextValidator{Search: "llo", SearchAs: stringcompareas.EndsWith})

	// Act
	actual := args.Map{"match": tvs.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns non-empty -- all pass", actual)
}

func Test_TextValidators_IsMatch_OneFails(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	tvs.Add(corevalidator.TextValidator{Search: "world", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{"match": tvs.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatch returns non-empty -- one fails", actual)
}

func Test_TextValidators_IsMatchMany_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)

	// Act
	actual := args.Map{"match": tvs.IsMatchMany(true, true, "a")}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.IsMatchMany returns empty -- empty", actual)
}

func Test_TextValidators_VerifyFirstError_Empty(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.VerifyFirstError(0, "hello", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns empty -- empty", actual)
}

func Test_TextValidators_VerifyFirstError_Match_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyFirstError(0, "hello", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns error -- match", actual)
}

func Test_TextValidators_VerifyFirstError_Mismatch_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyFirstError(0, "world", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstError returns error -- mismatch", actual)
}

func Test_TextValidators_VerifyErrorMany_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators
	err := tvs.VerifyErrorMany(true, &corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns nil -- nil", actual)
}

func Test_TextValidators_VerifyErrorMany_Continue(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyErrorMany(true, &corevalidator.Parameter{IsCaseSensitive: true}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns error -- continue", actual)
}

func Test_TextValidators_VerifyErrorMany_StopFirst(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	err := tvs.VerifyErrorMany(false, &corevalidator.Parameter{IsCaseSensitive: true}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyErrorMany returns error -- stop first", actual)
}

func Test_TextValidators_VerifyFirstErrorMany_Empty(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.VerifyFirstErrorMany(&corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.VerifyFirstErrorMany returns empty -- empty", actual)
}

func Test_TextValidators_AllVerifyErrorMany_Empty(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.AllVerifyErrorMany(&corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyErrorMany returns empty -- empty", actual)
}

func Test_TextValidators_AllVerifyError_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	err := tvs.AllVerifyError(0, "hello", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyError returns empty -- empty", actual)
}

func Test_TextValidators_AllVerifyError_WithErrors(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal})
	err := tvs.AllVerifyError(0, "y", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AllVerifyError returns error -- with errors", actual)
}

func Test_TextValidators_Dispose(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "a"})
	tvs.Dispose()

	// Act
	actual := args.Map{"nil": tvs.Items == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns non-empty -- with args", actual)
}

func Test_TextValidators_Dispose_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators
	tvs.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns nil -- nil", actual)
}

func Test_TextValidators_AsBasicSliceContractsBinder(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	binder := tvs.AsBasicSliceContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.AsBasicSliceContractsBinder returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LineValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_LineValidator_IsMatch_LineAndText(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{
		"matchBoth":  lv.IsMatch(5, "hello", true),
		"wrongLine":  lv.IsMatch(3, "hello", true),
		"wrongText":  lv.IsMatch(5, "world", true),
		"anyLine":    lv.IsMatch(-1, "hello", true),
	}

	// Assert
	expected := args.Map{
		"matchBoth": true, "wrongLine": false, "wrongText": false, "anyLine": true,
	}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatch returns non-empty -- with args", actual)
}

func Test_LineValidator_IsMatchMany_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator

	// Act
	actual := args.Map{"match": lv.IsMatchMany(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns nil -- nil", actual)
}

func Test_LineValidator_IsMatchMany_EmptySkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns empty -- empty skip", actual)
}

func Test_LineValidator_IsMatchMany_AllMatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "a", LineNumber: 0},
		{Text: "a", LineNumber: 1},
	}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(false, true, items...)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns non-empty -- all match", actual)
}

func Test_LineValidator_IsMatchMany_OneFails(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "a", LineNumber: 0},
		{Text: "b", LineNumber: 1},
	}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(false, true, items...)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineValidator.IsMatchMany returns non-empty -- one fails", actual)
}

func Test_LineValidator_VerifyError_Match(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 0, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- match", actual)
}

func Test_LineValidator_VerifyError_LineMismatch_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 3, "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- line mismatch", actual)
}

func Test_LineValidator_VerifyError_TextMismatch_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, 0, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyError returns error -- text mismatch", actual)
}

func Test_LineValidator_VerifyMany_Continue(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	items := []corestr.TextWithLineNumber{{Text: "a", LineNumber: 0}}
	err := lv.VerifyMany(true, params, items...)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyMany returns non-empty -- continue", actual)
}

func Test_LineValidator_VerifyMany_StopFirst(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	items := []corestr.TextWithLineNumber{{Text: "a", LineNumber: 0}}
	err := lv.VerifyMany(false, params, items...)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyMany returns non-empty -- stop first", actual)
}

func Test_LineValidator_VerifyFirstError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	err := lv.VerifyFirstError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_LineValidator_VerifyFirstError_EmptySkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	err := lv.VerifyFirstError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.VerifyFirstError returns empty -- empty skip", actual)
}

func Test_LineValidator_AllVerifyError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	err := lv.AllVerifyError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_LineValidator_AllVerifyError_EmptySkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	err := lv.AllVerifyError(&corevalidator.Parameter{IsSkipCompareOnActualEmpty: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator.AllVerifyError returns empty -- empty skip", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LinesValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinesValidators_NilLength(t *testing.T) {
	// Arrange
	var lv *corevalidator.LinesValidators

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"empty": lv.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns nil -- nil", actual)
}

func Test_LinesValidators_Basic(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"empty": lv.IsEmpty(),
		"count": lv.Count(),
		"hasAny": lv.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"count": 0,
		"hasAny": false,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- basic", actual)
}

func Test_LinesValidators_Add_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	})

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"hasAny": lv.HasAnyItem(),
		"lastIndex": lv.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasAny": true,
		"lastIndex": 0,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators.Add returns non-empty -- with args", actual)
}

func Test_LinesValidators_AddPtr_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.AddPtr(nil)

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AddPtr returns nil -- nil", actual)
}

func Test_LinesValidators_AddPtr_Valid_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	v := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	lv.AddPtr(v)

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AddPtr returns non-empty -- valid", actual)
}

func Test_LinesValidators_Adds_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Adds(
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal}},
	)

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinesValidators.Adds returns non-empty -- with args", actual)
}

func Test_LinesValidators_HasIndex_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})

	// Act
	actual := args.Map{
		"has0": lv.HasIndex(0),
		"has1": lv.HasIndex(1),
	}

	// Assert
	expected := args.Map{
		"has0": true,
		"has1": false,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators.HasIndex returns non-empty -- with args", actual)
}

func Test_LinesValidators_String_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})

	// Act
	actual := args.Map{"notEmpty": lv.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.String returns non-empty -- with args", actual)
}

func Test_LinesValidators_AsBasicSliceContractsBinder_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	binder := lv.AsBasicSliceContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.AsBasicSliceContractsBinder returns non-empty -- with args", actual)
}

func Test_LinesValidators_IsMatchText_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{"match": lv.IsMatchText("anything", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns empty -- empty", actual)
}

func Test_LinesValidators_IsMatchText_Match_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	})

	// Act
	actual := args.Map{"match": lv.IsMatchText("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns non-empty -- match", actual)
}

func Test_LinesValidators_IsMatchText_NoMatch_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	})

	// Act
	actual := args.Map{"match": lv.IsMatchText("world", true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatchText returns empty -- no match", actual)
}

func Test_LinesValidators_IsMatch_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{"match": lv.IsMatch(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- empty", actual)
}

func Test_LinesValidators_IsMatch_NoContentsSkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})

	// Act
	actual := args.Map{"match": lv.IsMatch(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- no contents skip", actual)
}

func Test_LinesValidators_IsMatch_NoContentsNoSkip_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a"}})

	// Act
	actual := args.Map{"match": lv.IsMatch(false, true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LinesValidators.IsMatch returns empty -- no contents no skip", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseLinesValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseLinesValidators_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var blv *corevalidator.BaseLinesValidators

	// Act
	actual := args.Map{
		"len": blv.LinesValidatorsLength(),
		"empty": blv.IsEmptyLinesValidators(),
		"has": blv.HasLinesValidators(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns nil -- nil", actual)
}

func Test_BaseLinesValidators_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{}

	// Act
	actual := args.Map{
		"len": blv.LinesValidatorsLength(),
		"empty": blv.IsEmptyLinesValidators(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns empty -- empty", actual)
}

func Test_BaseLinesValidators_WithItems_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		},
	}

	// Act
	actual := args.Map{
		"len": blv.LinesValidatorsLength(),
		"has": blv.HasLinesValidators(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns non-empty -- with items", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_Empty_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{}
	lv := blv.ToLinesValidators()

	// Act
	actual := args.Map{
		"notNil": lv != nil,
		"empty": lv.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators.ToLinesValidators returns empty -- empty", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_WithItems_FromConditionIsSplitByWh_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		},
	}
	lv := blv.ToLinesValidators()

	// Act
	actual := args.Map{
		"notNil": lv != nil,
		"len": lv.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators.ToLinesValidators returns non-empty -- with items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseValidatorCoreCondition
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseValidatorCoreCondition_WithCondition(t *testing.T) {
	// Arrange
	cond := &corevalidator.Condition{IsTrimCompare: true}
	bvc := &corevalidator.BaseValidatorCoreCondition{ValidatorCoreCondition: cond}
	result := bvc.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{"trim": result.IsTrimCompare}

	// Assert
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns non-empty -- with condition", actual)
}

func Test_BaseValidatorCoreCondition_NilCondition(t *testing.T) {
	// Arrange
	bvc := &corevalidator.BaseValidatorCoreCondition{}
	result := bvc.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{"trim": result.IsTrimCompare}

	// Assert
	expected := args.Map{"trim": false}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns nil -- nil condition", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Messages
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_ActualInputWithExpectingMessage(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
		CompareAs:     stringcompareas.Equal,
	}
	result := sv.ActualInputWithExpectingMessage(0, "test")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.ActualInputWithExpectingMessage returns non-empty -- with args", actual)
}

func Test_SliceValidator_ActualInputMessage(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ActualLines: []string{"hello"}}
	result := sv.ActualInputMessage(0, "test")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.ActualInputMessage returns non-empty -- with args", actual)
}

func Test_SliceValidator_UserExpectingMessage(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ExpectedLines: []string{"hello"}}
	result := sv.UserExpectingMessage(0, "test")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.UserExpectingMessage returns non-empty -- with args", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_NoAttach(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: false}
	testErr := errors.New("test error")
	result := sv.UserInputsMergeWithError(params, testErr)

	// Act
	actual := args.Map{"sameErr": result.Error() == "test error"}

	// Assert
	expected := args.Map{"sameErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns empty -- no attach", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_Attach_NilErr(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: true}
	result := sv.UserInputsMergeWithError(params, nil)

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns nil -- attach nil err", actual)
}

func Test_SliceValidator_UserInputsMergeWithError_Attach_WithErr(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}}
	params := &corevalidator.Parameter{IsAttachUserInputs: true}
	testErr := errors.New("base")
	result := sv.UserInputsMergeWithError(params, testErr)

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"containsBase": len(result.Error()) > 4,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"containsBase": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- attach with err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewSliceValidatorUsingErr(t *testing.T) {
	// Arrange
	testErr := errors.New("line1\nline2")
	sv := corevalidator.NewSliceValidatorUsingErr(testErr, "line1\nline2", true, false, false, stringcompareas.Equal)

	// Act
	actual := args.Map{
		"notNil": sv != nil,
		"actualLen": sv.ActualLinesLength(),
		"expectedLen": sv.ExpectingLinesLength(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"actualLen": 2,
		"expectedLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns error -- with args", actual)
}

func Test_NewSliceValidatorUsingAny_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny("hello\nworld", "hello\nworld", false, false, false, stringcompareas.Equal)

	// Act
	actual := args.Map{
		"notNil": sv != nil,
		"actualLen": sv.ActualLinesLength(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"actualLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Verify
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_VerifyFirstError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.VerifyFirstError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.VerifyFirstError returns nil -- nil", actual)
}

func Test_SliceValidator_AllVerifyError_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns nil -- nil", actual)
}

func Test_SliceValidator_AllVerifyError_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello", "world"},
		ExpectedLines: []string{"hello", "world"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns error -- match", actual)
}

func Test_SliceValidator_AllVerifyError_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyError(&corevalidator.Parameter{IsCaseSensitive: true})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyError returns error -- mismatch", actual)
}

func Test_SliceValidator_AllVerifyErrorExceptLast_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator.AllVerifyErrorExceptLast returns nil -- nil", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_EmptyIgnore(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}, 1)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns empty -- empty ignore", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_BothNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 0)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- both nil", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_OneNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines: []string{"a"},
		CompareAs:   stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- one nil", actual)
}

func Test_SliceValidator_AllVerifyErrorUptoLength_LengthMismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 2)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- length mismatch", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorTestCase(0, "test", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns nil -- nil", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Nil_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorQuick(0, "test", "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorQuick returns nil -- nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSliceValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleSliceValidator_SetActual(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	result := ssv.SetActual([]string{"a"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.SetActual returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_SliceValidator(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	sv := ssv.SliceValidator()

	// Act
	actual := args.Map{"notNil": sv != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.SliceValidator returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_VerifyAll(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	err := ssv.VerifyAll([]string{"a"}, &corevalidator.Parameter{IsCaseSensitive: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyAll returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_VerifyFirst(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a"})
	err := ssv.VerifyFirst([]string{"a"}, &corevalidator.Parameter{IsCaseSensitive: true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyFirst returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_VerifyUpto(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a", "b"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"a", "b"})
	err := ssv.VerifyUpto([]string{"a", "b"}, &corevalidator.Parameter{IsCaseSensitive: true}, 1)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyUpto returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RangesSegment / RangeSegmentsValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_RangesSegment_Fields(t *testing.T) {
	// Arrange
	rs := corevalidator.RangesSegment{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"len": len(rs.ExpectedLines)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RangesSegment returns correct value -- fields", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Vars (predefined conditions)
// ══════════════════════════════════════════════════════════════════════════════

func Test_DefaultDisabledCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultDisabledCoreCondition

	// Act
	actual := args.Map{
		"trim": c.IsTrimCompare,
		"unique": c.IsUniqueWordOnly,
		"nonEmpty": c.IsNonEmptyWhitespace,
		"sort": c.IsSortStringsBySpace,
	}

	// Assert
	expected := args.Map{
		"trim": false,
		"unique": false,
		"nonEmpty": false,
		"sort": false,
	}
	expected.ShouldBeEqual(t, 0, "DefaultDisabledCoreCondition returns correct value -- with args", actual)
}

func Test_DefaultTrimCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultTrimCoreCondition

	// Act
	actual := args.Map{"trim": c.IsTrimCompare}

	// Assert
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "DefaultTrimCoreCondition returns correct value -- with args", actual)
}

func Test_DefaultSortTrimCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultSortTrimCoreCondition

	// Act
	actual := args.Map{
		"trim": c.IsTrimCompare,
		"nonEmpty": c.IsNonEmptyWhitespace,
		"sort": c.IsSortStringsBySpace,
	}

	// Assert
	expected := args.Map{
		"trim": true,
		"nonEmpty": true,
		"sort": true,
	}
	expected.ShouldBeEqual(t, 0, "DefaultSortTrimCoreCondition returns correct value -- with args", actual)
}

func Test_DefaultUniqueWordsCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultUniqueWordsCoreCondition

	// Act
	actual := args.Map{
		"trim": c.IsTrimCompare,
		"unique": c.IsUniqueWordOnly,
		"nonEmpty": c.IsNonEmptyWhitespace,
		"sort": c.IsSortStringsBySpace,
	}

	// Assert
	expected := args.Map{
		"trim": true,
		"unique": true,
		"nonEmpty": true,
		"sort": true,
	}
	expected.ShouldBeEqual(t, 0, "DefaultUniqueWordsCoreCondition returns correct value -- with args", actual)
}

func Test_EmptyValidator(t *testing.T) {
	// Arrange
	v := corevalidator.EmptyValidator

	// Act
	actual := args.Map{
		"search": v.Search,
		"method": v.SearchAs.Name(),
		"trim": v.IsTrimCompare,
	}

	// Assert
	expected := args.Map{
		"search": "",
		"method": "Equal",
		"trim": true,
	}
	expected.ShouldBeEqual(t, 0, "EmptyValidator returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidators_NilLength(t *testing.T) {
	// Arrange
	var svs *corevalidator.SliceValidators

	// Act
	actual := args.Map{
		"len": svs.Length(),
		"empty": svs.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns nil -- nil", actual)
}

func Test_SliceValidators_IsMatch_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}

	// Act
	actual := args.Map{"match": svs.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns empty -- empty", actual)
}

func Test_SliceValidators_IsMatch_AllPass(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{ActualLines: []string{"a"}, ExpectedLines: []string{"a"}, CompareAs: stringcompareas.Equal},
		},
	}

	// Act
	actual := args.Map{"match": svs.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns non-empty -- all pass", actual)
}

func Test_SliceValidators_IsMatch_Fail(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{ActualLines: []string{"a"}, ExpectedLines: []string{"b"}, CompareAs: stringcompareas.Equal},
		},
	}

	// Act
	actual := args.Map{"match": svs.IsMatch(true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsMatch returns non-empty -- fail", actual)
}

func Test_SliceValidators_IsValid(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}

	// Act
	actual := args.Map{"valid": svs.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsValid returns non-empty -- with args", actual)
}

func Test_SliceValidators_SetActualOnAll_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	svs.SetActualOnAll("a", "b") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.SetActualOnAll returns empty -- empty", actual)
}

func Test_SliceValidators_VerifyAll_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAll("header", &corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAll returns empty -- empty", actual)
}

func Test_SliceValidators_VerifyAllError_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAllError(&corevalidator.Parameter{Header: "test"})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAllError returns empty -- empty", actual)
}

func Test_SliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyAllErrorUsingActual(&corevalidator.Parameter{Header: "test"}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyAllErrorUsingActual returns empty -- empty", actual)
}

func Test_SliceValidators_VerifyFirst_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyFirst(&corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyFirst returns empty -- empty", actual)
}

func Test_SliceValidators_VerifyUpto_Empty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}
	err := svs.VerifyUpto(false, false, 1, &corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.VerifyUpto returns empty -- empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// HeaderSliceValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_HeaderSliceValidators_NilLength(t *testing.T) {
	// Arrange
	var hsvs corevalidator.HeaderSliceValidators

	// Act
	actual := args.Map{
		"len": hsvs.Length(),
		"empty": hsvs.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns nil -- nil", actual)
}

func Test_HeaderSliceValidators_IsMatch_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}

	// Act
	actual := args.Map{"match": hsvs.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsMatch returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_IsValid(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}

	// Act
	actual := args.Map{"valid": hsvs.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsValid returns non-empty -- with args", actual)
}

func Test_HeaderSliceValidators_SetActualOnAll_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	hsvs.SetActualOnAll("a") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.SetActualOnAll returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_VerifyAll_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAll("header", &corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAll returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_VerifyAllError_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAllError(&corevalidator.Parameter{Header: "test"})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAllError returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyAllErrorUsingActual(&corevalidator.Parameter{Header: "test"}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyAllErrorUsingActual returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyFirst(&corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyFirst returns empty -- empty", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_Empty(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{}
	err := hsvs.VerifyUpto(false, false, 1, &corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.VerifyUpto returns empty -- empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RangeSegmentsValidator
// ══════════════════════════════════════════════════════════════════════════════

func Test_RangeSegmentsValidator_LengthOfVerifierSegments(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		VerifierSegments: []corevalidator.RangesSegment{{}, {}},
	}

	// Act
	actual := args.Map{"len": rsv.LengthOfVerifierSegments()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator.LengthOfVerifierSegments returns non-empty -- with args", actual)
}

func Test_RangeSegmentsValidator_SetActual_FromConditionIsSplitByWh(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{Title: "test"}
	result := rsv.SetActual([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator.SetActual returns non-empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — isLengthOkay / isEmptyIgnoreCase
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_Dispose_WithValidators(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	_ = sv.ComparingValidators() // force lazy init
	sv.Dispose()

	// Act
	actual := args.Map{
		"actualNil": sv.ActualLines == nil,
		"expectedNil": sv.ExpectedLines == nil,
	}

	// Assert
	expected := args.Map{
		"actualNil": true,
		"expectedNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator.Dispose returns non-empty -- with validators", actual)
}
