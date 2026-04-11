package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── SliceValidator ──

func Test_SliceValidator_IsUsedAlready_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"used": sv.IsUsedAlready()}

	// Assert
	expected := args.Map{"used": false}
	expected.ShouldBeEqual(t, 0, "IsUsedAlready returns nil -- nil", actual)
}

func Test_SliceValidator_ActualLinesLength_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"len": sv.ActualLinesLength()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ActualLinesLength returns nil -- nil", actual)
}

func Test_SliceValidator_MethodName_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"notEmpty": sv.MethodName() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MethodName returns correct value -- with args", actual)
}

func Test_SliceValidator_SetActual_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{}
	sv.SetActual([]string{"a", "b"})

	// Act
	actual := args.Map{
		"used": sv.IsUsedAlready(),
		"len": sv.ActualLinesLength(),
	}

	// Assert
	expected := args.Map{
		"used": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "SetActual returns correct value -- with args", actual)
}

func Test_SliceValidator_SetActualVsExpected_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})

	// Act
	actual := args.Map{
		"used": sv.IsUsedAlready(),
		"actualLen": sv.ActualLinesLength(),
		"expectedLen": sv.ExpectingLinesLength(),
	}

	// Assert
	expected := args.Map{
		"used": true,
		"actualLen": 1,
		"expectedLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SetActualVsExpected returns correct value -- with args", actual)
}

func Test_SliceValidator_ActualLinesString(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ActualLines: []string{"line1", "line2"}}
	result := sv.ActualLinesString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ActualLinesString returns correct value -- with args", actual)
}

func Test_SliceValidator_ActualLinesString_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	result := sv.ActualLinesString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ActualLinesString returns nil -- nil", actual)
}

func Test_SliceValidator_ExpectingLinesString(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{ExpectedLines: []string{"line1"}}
	result := sv.ExpectingLinesString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesString returns correct value -- with args", actual)
}

func Test_SliceValidator_ExpectingLinesString_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	result := sv.ExpectingLinesString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesString returns nil -- nil", actual)
}

func Test_SliceValidator_ExpectingLinesLength_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"len": sv.ExpectingLinesLength()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesLength returns nil -- nil", actual)
}

func Test_SliceValidator_ComparingValidators(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}
	validators := sv.ComparingValidators()

	// Act
	actual := args.Map{"count": len(validators.Items)}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "ComparingValidators returns non-empty -- with args", actual)
}

func Test_SliceValidator_ComparingValidators_Cached_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	v1 := sv.ComparingValidators()
	v2 := sv.ComparingValidators()

	// Act
	actual := args.Map{"same": v1 == v2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "ComparingValidators returns non-empty -- cached", actual)
}

func Test_SliceValidator_IsValid_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns nil -- nil", actual)
}

func Test_SliceValidator_IsValid_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello", "world"},
		ExpectedLines: []string{"hello", "world"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns non-empty -- match", actual)
}

func Test_SliceValidator_IsValid_Mismatch_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid returns non-empty -- mismatch", actual)
}

func Test_SliceValidator_IsValid_DifferentLengths(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid returns non-empty -- different lengths", actual)
}

func Test_SliceValidator_IsValidOtherLines_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{
		"match":    sv.IsValidOtherLines(true, []string{"a", "b"}),
		"mismatch": sv.IsValidOtherLines(true, []string{"c", "d"}),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"mismatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidOtherLines returns non-empty -- with args", actual)
}

func Test_SliceValidator_IsValidLines_BothNil(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "isValidLines returns nil -- both nil", actual)
}

func Test_SliceValidator_IsValidLines_LinesNilExpectedNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ExpectedLines: nil,
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "isValidLines returns nil -- both nil (non-nil receiver)", actual)
}

func Test_SliceValidator_IsValidLines_OneNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "isValidLines returns nil -- one nil", actual)
}

func Test_SliceValidator_Dispose_FromSliceValidatorIsUsed(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_SliceValidator_Dispose_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	sv.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

// ── TextValidator ──

func Test_TextValidator_IsMatch_EqualMatch(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}

	// Act
	actual := args.Map{
		"match":    tv.IsMatch("hello", true),
		"mismatch": tv.IsMatch("world", true),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"mismatch": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- equal", actual)
}

func Test_TextValidator_IsMatch_Contains_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "ell",
		SearchAs: stringcompareas.Contains,
	}

	// Act
	actual := args.Map{"match": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- contains", actual)
}

func Test_TextValidator_IsMatch_StartsWith_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hel",
		SearchAs: stringcompareas.StartsWith,
	}

	// Act
	actual := args.Map{"match": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- startsWith", actual)
}

func Test_TextValidator_IsMatch_EndsWith_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "llo",
		SearchAs: stringcompareas.EndsWith,
	}

	// Act
	actual := args.Map{"match": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch returns non-empty -- endsWith", actual)
}

// ── TextValidators ──

func Test_TextValidators_Add(t *testing.T) {
	// Arrange
	validators := corevalidator.NewTextValidators(5)
	validators.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	validators.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{"count": len(validators.Items)}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "TextValidators.Add returns non-empty -- with args", actual)
}

func Test_TextValidators_Dispose_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	validators := corevalidator.NewTextValidators(5)
	validators.Add(corevalidator.TextValidator{Search: "a"})
	validators.Dispose()

	// Act
	actual := args.Map{"nil": validators.Items == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns non-empty -- with args", actual)
}

func Test_TextValidators_Dispose_Nil_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	var validators *corevalidator.TextValidators
	validators.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose returns nil -- nil", actual)
}

// ── Condition ──

func Test_Condition_IsSplitByWhitespace(t *testing.T) {
	// Arrange
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"split": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"split": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsSplitByWhitespace returns correct value -- with args", actual)
}

// ── Parameter ──

func Test_Parameter_Fields(t *testing.T) {
	// Arrange
	p := &corevalidator.Parameter{
		IsSkipCompareOnActualEmpty: true,
	}

	// Act
	actual := args.Map{"skip": p.IsSkipCompareOnActualEmpty}

	// Assert
	expected := args.Map{"skip": true}
	expected.ShouldBeEqual(t, 0, "Parameter returns correct value -- fields", actual)
}

// ── SimpleSliceValidator ──

func Test_SimpleSliceValidator_SetActual_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange

	// Assert
	expected := corestr.New.SimpleSlice.SpreadStrings("a", "b")
	sv := &corevalidator.SimpleSliceValidator{
		Expected: expected,
	}
	sv.SetActual([]string{"a", "b"})
	sliceV := sv.SliceValidator()

	// Act
	actual := args.Map{"notNil": sliceV != nil}
	expectedM := args.Map{"notNil": true}
	expectedM.ShouldBeEqual(t, 0, "SimpleSliceValidator.SetActual returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_VerifyAll_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange

	// Assert
	expected := corestr.New.SimpleSlice.SpreadStrings("a", "b")
	sv := &corevalidator.SimpleSliceValidator{
		Expected: expected,
	}
	sv.SetActual([]string{"a", "b"})
	sliceV := sv.SliceValidator()
	// Verify SliceValidator was created properly

	// Act
	actual := args.Map{
		"notNil": sliceV != nil,
		"actualLen": len(sliceV.ActualLines),
		"expectedLen": len(sliceV.ExpectedLines),
	}
	expectedM := args.Map{
		"notNil": true,
		"actualLen": 2,
		"expectedLen": 2,
	}
	expectedM.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyAll returns non-empty -- with args", actual)
}

func Test_SimpleSliceValidator_VerifyAll_Mismatch_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange

	// Assert
	expected := corestr.New.SimpleSlice.SpreadStrings("a", "b")
	sv := &corevalidator.SimpleSliceValidator{
		Expected: expected,
	}
	sv.SetActual([]string{"a", "c"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.VerifyAll([]string{"a", "c"}, params)

	// Act
	actual := args.Map{"hasErr": err != nil}
	expectedM := args.Map{"hasErr": true}
	expectedM.ShouldBeEqual(t, 0, "SimpleSliceValidator.VerifyAll returns non-empty -- mismatch", actual)
}

// ── LineNumber ──

func Test_LineNumber_Fields(t *testing.T) {
	// Arrange
	ln := corevalidator.LineNumber{
		LineNumber: 1,
	}

	// Act
	actual := args.Map{
		"num": ln.LineNumber,
		"hasLN": ln.HasLineNumber(),
	}

	// Assert
	expected := args.Map{
		"num": 1,
		"hasLN": true,
	}
	expected.ShouldBeEqual(t, 0, "LineNumber returns correct value -- fields", actual)
}

// ── HeaderSliceValidator ──

func Test_HeaderSliceValidator_IsValid(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidator{
		Header: "test-header",
		SliceValidator: corevalidator.SliceValidator{
			ActualLines:   []string{"a"},
			ExpectedLines: []string{"a"},
			CompareAs:     stringcompareas.Equal,
		},
	}

	// Act
	actual := args.Map{"valid": hsv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidator.IsValid returns non-empty -- with args", actual)
}

// ── HeaderSliceValidators ──

func Test_HeaderSliceValidators_IsEmpty_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	hsvs := &corevalidator.HeaderSliceValidators{}

	// Act
	actual := args.Map{"empty": hsvs.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsEmpty returns empty -- with args", actual)
}

func Test_HeaderSliceValidators_Length_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	hsvs := corevalidator.HeaderSliceValidators{
		{Header: "h1"},
	}

	// Act
	actual := args.Map{"count": hsvs.Length()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.Length returns non-empty -- with args", actual)
}

// ── SliceValidators ──

func Test_SliceValidators_IsEmpty(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{}

	// Act
	actual := args.Map{"empty": svs.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsEmpty returns empty -- with args", actual)
}

func Test_SliceValidators_Length(t *testing.T) {
	// Arrange
	svs := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}

	// Act
	actual := args.Map{"count": svs.Length()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "SliceValidators.Length returns non-empty -- with args", actual)
}

// ── RangesSegment ──

func Test_RangesSegment_FromSliceValidatorIsUsed(t *testing.T) {
	// Arrange
	rs := corevalidator.RangesSegment{
		ExpectedLines: []string{"a", "b"},
	}

	// Act
	actual := args.Map{"len": len(rs.ExpectedLines)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesSegment returns correct value -- with args", actual)
}
