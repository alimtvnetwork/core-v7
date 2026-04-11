package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// BaseLinesValidators
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseLinesValidators_LinesValidatorsLength_Nil(t *testing.T) {
	// Arrange
	var blv *corevalidator.BaseLinesValidators

	// Act
	actual := args.Map{"len": blv.LinesValidatorsLength()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesValidatorsLength returns nil -- nil", actual)
}

func Test_BaseLinesValidators_LinesValidatorsLength_Empty(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{}

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
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns empty -- empty", actual)
}

func Test_BaseLinesValidators_WithItems(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
		},
	}

	// Act
	actual := args.Map{
		"len": blv.LinesValidatorsLength(),
		"empty": blv.IsEmptyLinesValidators(),
		"has": blv.HasLinesValidators(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"empty": false,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLinesValidators returns non-empty -- with items", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_Empty(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{}
	lv := blv.ToLinesValidators()

	// Act
	actual := args.Map{
		"notNil": lv != nil,
		"len": lv.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToLinesValidators returns empty -- empty", actual)
}

func Test_BaseLinesValidators_ToLinesValidators_WithItems(t *testing.T) {
	// Arrange
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}},
			{TextValidator: corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal}},
		},
	}
	lv := blv.ToLinesValidators()

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToLinesValidators returns non-empty -- with items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseValidatorCoreCondition
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseValidatorCoreCondition_WithExisting(t *testing.T) {
	// Arrange
	cond := &corevalidator.Condition{IsTrimCompare: true}
	bv := corevalidator.BaseValidatorCoreCondition{ValidatorCoreCondition: cond}
	result := bv.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{"trim": result.IsTrimCompare}

	// Assert
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "ValidatorCoreConditionDefault returns non-empty -- existing", actual)
}

func Test_BaseValidatorCoreCondition_CreateDefault(t *testing.T) {
	// Arrange
	bv := corevalidator.BaseValidatorCoreCondition{}
	result := bv.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{
		"trim": result.IsTrimCompare,
		"notNil": bv.ValidatorCoreCondition != nil,
	}

	// Assert
	expected := args.Map{
		"trim": false,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ValidatorCoreConditionDefault returns non-empty -- new default", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LinesValidators — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinesValidators_NilReceiver(t *testing.T) {
	// Arrange
	var lv *corevalidator.LinesValidators

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"count": lv.Count(),
		"empty": lv.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"count": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns nil -- nil receiver", actual)
}

func Test_LinesValidators_NewAndBasicOps(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"empty": lv.IsEmpty(),
		"hasAny": lv.HasAnyItem(),
		"lastIdx": lv.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"hasAny": false,
		"lastIdx": -1,
	}
	expected.ShouldBeEqual(t, 0, "NewLinesValidators returns empty -- empty", actual)
}

func Test_LinesValidators_Add(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})

	// Act
	actual := args.Map{
		"len": lv.Length(),
		"count": lv.Count(),
		"hasAny": lv.HasAnyItem(),
		"hasIdx0": lv.HasIndex(0),
		"hasIdx5": lv.HasIndex(5),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"count": 1,
		"hasAny": true,
		"hasIdx0": true,
		"hasIdx5": false,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- Add", actual)
}

func Test_LinesValidators_AddPtr_Nil(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.AddPtr(nil)

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns nil -- AddPtr nil", actual)
}

func Test_LinesValidators_AddPtr_Valid(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	v := &corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}}
	lv.AddPtr(v)

	// Act
	actual := args.Map{"len": lv.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- AddPtr valid", actual)
}

func Test_LinesValidators_Adds(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- Adds", actual)
}

func Test_LinesValidators_String(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})

	// Act
	actual := args.Map{"notEmpty": lv.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- String", actual)
}

func Test_LinesValidators_AsBasicSliceContractsBinder(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)

	// Act
	actual := args.Map{"notNil": lv.AsBasicSliceContractsBinder() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LinesValidators returns non-empty -- AsBasicSliceContractsBinder", actual)
}

func Test_LinesValidators_IsMatchText_Empty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)

	// Act
	actual := args.Map{"match": lv.IsMatchText("anything", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatchText returns empty -- empty", actual)
}

func Test_LinesValidators_IsMatchText_Match(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})

	// Act
	actual := args.Map{
		"match": lv.IsMatchText("hello", true),
		"noMatch": lv.IsMatchText("world", true),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMatchText returns non-empty -- with validator", actual)
}

func Test_LinesValidators_IsMatch_Empty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)

	// Act
	actual := args.Map{"match": lv.IsMatch(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns empty -- empty validators", actual)
}

func Test_LinesValidators_IsMatch_NoContentsSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})

	// Act
	actual := args.Map{"match": lv.IsMatch(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns empty -- no contents skip", actual)
}

func Test_LinesValidators_IsMatch_NoContentsNoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})

	// Act
	actual := args.Map{"match": lv.IsMatch(false, true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns empty -- no contents no skip", actual)
}

func Test_LinesValidators_IsMatch_WithContents(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}

	// Act
	actual := args.Map{"match": lv.IsMatch(false, true, twl)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns non-empty -- with contents", actual)
}

func Test_LinesValidators_IsMatch_Mismatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}

	// Act
	actual := args.Map{"match": lv.IsMatch(false, true, twl)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct value -- mismatch", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Empty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns empty -- empty", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_NoContentsSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns empty -- no contents skip", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_NoContentsNoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: false}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns empty -- no contents no skip", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Match(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.VerifyFirstDefaultLineNumberError(params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns error -- match", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Mismatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}
	err := lv.VerifyFirstDefaultLineNumberError(params, twl)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns error -- mismatch", actual)
}

func Test_LinesValidators_AllVerifyError_Empty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns empty -- empty", actual)
}

func Test_LinesValidators_AllVerifyError_NoContentsSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns empty -- no contents skip", actual)
}

func Test_LinesValidators_AllVerifyError_NoContentsNoSkip(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: false}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns empty -- no contents no skip", actual)
}

func Test_LinesValidators_AllVerifyError_Match(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.AllVerifyError(params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- match", actual)
}

func Test_LinesValidators_AllVerifyError_Mismatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(5)
	lv.Add(corevalidator.LineValidator{TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}
	err := lv.AllVerifyError(params, twl)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LineValidator — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_LineValidator_IsMatch_NoLineCheck(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{"match": lv.IsMatch(-1, "hello", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns empty -- IsMatch no line check", actual)
}

func Test_LineValidator_IsMatch_LineMatches(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{
		"match": lv.IsMatch(5, "hello", true),
		"lineMismatch": lv.IsMatch(3, "hello", true),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"lineMismatch": false,
	}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- IsMatch line matches", actual)
}

func Test_LineValidator_IsMatch_TextMismatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{"match": lv.IsMatch(-1, "world", true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- IsMatch text mismatch", actual)
}

func Test_LineValidator_IsMatchMany_Nil(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator

	// Act
	actual := args.Map{"match": lv.IsMatchMany(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns nil -- IsMatchMany nil", actual)
}

func Test_LineValidator_IsMatchMany_EmptySkip(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(true, true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns empty -- IsMatchMany empty skip", actual)
}

func Test_LineValidator_IsMatchMany_WithContents(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(false, true, twl)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- IsMatchMany with contents", actual)
}

func Test_LineValidator_IsMatchMany_Mismatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}

	// Act
	actual := args.Map{"match": lv.IsMatchMany(false, true, twl)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- IsMatchMany mismatch", actual)
}

func Test_LineValidator_VerifyError_LineMatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, -1, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- VerifyError match", actual)
}

func Test_LineValidator_VerifyError_LineMismatch(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- VerifyError line mismatch", actual)
}

func Test_LineValidator_VerifyError_TextMismatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := lv.VerifyError(params, -1, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- VerifyError text mismatch", actual)
}

func Test_LineValidator_VerifyMany_ContinueOnError(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.VerifyMany(true, params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- VerifyMany continue", actual)
}

func Test_LineValidator_VerifyMany_StopOnFirst(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.VerifyMany(false, params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns non-empty -- VerifyMany stop first", actual)
}

func Test_LineValidator_VerifyFirstError_Nil(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	params := &corevalidator.Parameter{}
	err := lv.VerifyFirstError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns nil -- VerifyFirstError nil", actual)
}

func Test_LineValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := lv.VerifyFirstError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns empty -- VerifyFirstError empty skip", actual)
}

func Test_LineValidator_VerifyFirstError_Match(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.VerifyFirstError(params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- VerifyFirstError match", actual)
}

func Test_LineValidator_VerifyFirstError_Mismatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}
	err := lv.VerifyFirstError(params, twl)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- VerifyFirstError mismatch", actual)
}

func Test_LineValidator_AllVerifyError_Nil(t *testing.T) {
	// Arrange
	var lv *corevalidator.LineValidator
	params := &corevalidator.Parameter{}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns nil -- AllVerifyError nil", actual)
}

func Test_LineValidator_AllVerifyError_EmptySkip(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns empty -- AllVerifyError empty skip", actual)
}

func Test_LineValidator_AllVerifyError_Match(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "hello", LineNumber: -1}
	err := lv.AllVerifyError(params, twl)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- AllVerifyError match", actual)
}

func Test_LineValidator_AllVerifyError_Mismatch(t *testing.T) {
	// Arrange
	lv := &corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	twl := corestr.TextWithLineNumber{Text: "world", LineNumber: -1}
	err := lv.AllVerifyError(params, twl)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LineValidator returns error -- AllVerifyError mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RangeSegmentsValidator — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_RangeSegmentsValidator_SetActual(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{Title: "test"}
	rsv.SetActual([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"segLen": rsv.LengthOfVerifierSegments()}

	// Assert
	expected := args.Map{"segLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- SetActual", actual)
}

func Test_RangeSegmentsValidator_Validators(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	rsv.SetActual([]string{"a", "b", "c"})
	validators := rsv.Validators()

	// Act
	actual := args.Map{"len": validators.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- Validators", actual)
}

func Test_RangeSegmentsValidator_VerifyAll_Match(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := rsv.VerifyAll("header", []string{"a", "b", "c"}, params, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyAll match", actual)
}

func Test_RangeSegmentsValidator_VerifySimple(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := rsv.VerifySimple([]string{"a", "b", "c"}, params, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifySimple match", actual)
}

func Test_RangeSegmentsValidator_VerifyFirst(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := rsv.VerifyFirst("header", []string{"a", "b", "c"}, params, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyFirst match", actual)
}

func Test_RangeSegmentsValidator_VerifyUpto(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := rsv.VerifyUpto("header", []string{"a", "b", "c"}, params, 2, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyUpto match", actual)
}

func Test_RangeSegmentsValidator_VerifyFirstDefault(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := rsv.VerifyFirstDefault([]string{"a", "b", "c"}, params, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyFirstDefault match", actual)
}

func Test_RangeSegmentsValidator_VerifyUptoDefault(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := rsv.VerifyUptoDefault([]string{"a", "b", "c"}, params, 2, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyUptoDefault match", actual)
}

func Test_RangeSegmentsValidator_VerifyAll_Mismatch(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"x", "y"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := rsv.VerifyAll("header", []string{"a", "b", "c"}, params, true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns non-empty -- VerifyAll mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSliceValidator — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleSliceValidator_VerifyFirst_Mismatch(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"b"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := ssv.VerifyFirst([]string{"b"}, params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator returns non-empty -- VerifyFirst mismatch", actual)
}

func Test_SimpleSliceValidator_VerifyUpto_Mismatch(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"a", "b"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"x", "y"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := ssv.VerifyUpto([]string{"x", "y"}, params, 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator returns non-empty -- VerifyUpto mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidatorConstructors — NewSliceValidatorUsingAny
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewSliceValidatorUsingAny_Match(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"line1\nline2",
		"line1\nline2",
		false, false, false,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns non-empty -- match", actual)
}

func Test_NewSliceValidatorUsingAny_Mismatch(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"actual",
		"expected",
		false, false, false,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns non-empty -- mismatch", actual)
}

func Test_NewSliceValidatorUsingAny_WithConditions(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"  hello  world  ",
		"hello world",
		true, true, true,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns non-empty -- with conditions", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator — case-insensitive and whitespace-condition branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidator_IsMatch_CaseInsensitive(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "Hello", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"match": tv.IsMatch("hello", false)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns non-empty -- IsMatch case insensitive", actual)
}

func Test_TextValidator_IsMatch_WithUniqueWord(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "a b c",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsTrimCompare:        true,
			IsUniqueWordOnly:     true,
			IsNonEmptyWhitespace: true,
			IsSortStringsBySpace: true,
		},
	}

	// Act
	actual := args.Map{"match": tv.IsMatch("  c  b  a  ", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns non-empty -- IsMatch unique word sort", actual)
}

func Test_TextValidator_VerifyDetailError_CaseInsensitive(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "Hello", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: false}
	err := tv.VerifyDetailError(params, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns error -- VerifyDetailError case insensitive", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators — VerifyFirstErrorMany and AllVerifyErrorMany non-empty paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidators_VerifyFirstErrorMany_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyFirstErrorMany(params, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyFirstErrorMany match", actual)
}

func Test_TextValidators_VerifyFirstErrorMany_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyFirstErrorMany(params, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyFirstErrorMany mismatch", actual)
}

func Test_TextValidators_AllVerifyErrorMany_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.AllVerifyErrorMany(params, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- AllVerifyErrorMany match", actual)
}

func Test_TextValidators_AllVerifyErrorMany_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.AllVerifyErrorMany(params, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- AllVerifyErrorMany mismatch", actual)
}

func Test_TextValidators_VerifyErrorMany_ContinueOnError(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyErrorMany(true, params, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyErrorMany continue", actual)
}

func Test_TextValidators_VerifyErrorMany_StopOnFirst(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyErrorMany(false, params, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyErrorMany stop first", actual)
}

func Test_TextValidators_VerifyFirstError_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "hello", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyFirstError match", actual)
}

func Test_TextValidators_VerifyFirstError_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "world", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyFirstError mismatch", actual)
}

func Test_TextValidators_AllVerifyError_Match(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "hello", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- AllVerifyError match", actual)
}

func Test_TextValidators_AllVerifyError_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "world", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- AllVerifyError mismatch", actual)
}

func Test_TextValidators_IsMatch_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{"match": tvs.IsMatch("world", true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidators returns non-empty -- IsMatch mismatch", actual)
}

func Test_TextValidators_IsMatchMany_Mismatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{"match": tvs.IsMatchMany(false, true, "world")}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidators returns non-empty -- IsMatchMany mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidators & HeaderSliceValidators — non-empty verify paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidators_VerifyAllError_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"a"}},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := sv.VerifyAllError(params)
	// VerifyAllError always inserts header

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns error -- VerifyAllError match", actual)
}

func Test_SliceValidators_VerifyAllErrorUsingActual_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"}},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := sv.VerifyAllErrorUsingActual(params, "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns error -- VerifyAllErrorUsingActual", actual)
}

func Test_SliceValidators_IsMatch_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"a"}},
		},
	}

	// Act
	actual := args.Map{"match": sv.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns non-empty -- IsMatch match", actual)
}

func Test_SliceValidators_IsMatch_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"b"}},
		},
	}

	// Act
	actual := args.Map{"match": sv.IsMatch(true)}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns non-empty -- IsMatch mismatch", actual)
}

func Test_SliceValidators_IsValid_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"a"}},
		},
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators returns non-empty -- IsValid match", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Match(t *testing.T) {
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
	err := hsv.VerifyAllErrorUsingActual(params, "a")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns error -- VerifyAllErrorUsingActual", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Mismatch(t *testing.T) {
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
	err := hsv.VerifyAllErrorUsingActual(params, "b")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns error -- VerifyAllErrorUsingActual mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_Match(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a", "b"},
				ExpectedLines: []string{"a", "b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}
	err := hsv.VerifyUpto(false, false, 2, params)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns non-empty -- VerifyUpto match", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_PrintError(t *testing.T) {
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
	err := hsv.VerifyUpto(true, false, 1, params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators returns error -- VerifyUpto print error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// vars.go — DefaultDisabledCoreCondition, EmptyValidator, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_Vars_DefaultDisabledCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultDisabledCoreCondition

	// Act
	actual := args.Map{
		"trim": c.IsTrimCompare,
		"unique": c.IsUniqueWordOnly,
		"split": c.IsSplitByWhitespace(),
	}

	// Assert
	expected := args.Map{
		"trim": false,
		"unique": false,
		"split": false,
	}
	expected.ShouldBeEqual(t, 0, "DefaultDisabledCoreCondition returns correct value -- with args", actual)
}

func Test_Vars_DefaultTrimCoreCondition(t *testing.T) {
	// Arrange
	c := corevalidator.DefaultTrimCoreCondition

	// Act
	actual := args.Map{"trim": c.IsTrimCompare}

	// Assert
	expected := args.Map{"trim": true}
	expected.ShouldBeEqual(t, 0, "DefaultTrimCoreCondition returns correct value -- with args", actual)
}

func Test_Vars_DefaultSortTrimCoreCondition(t *testing.T) {
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

func Test_Vars_DefaultUniqueWordsCoreCondition(t *testing.T) {
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

func Test_Vars_EmptyValidator(t *testing.T) {
	// Arrange
	ev := corevalidator.EmptyValidator

	// Act
	actual := args.Map{
		"search": ev.Search,
		"match": ev.IsMatch("", true),
	}

	// Assert
	expected := args.Map{
		"search": "",
		"match": true,
	}
	expected.ShouldBeEqual(t, 0, "EmptyValidator returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — additional branches for isEmptyIgnoreCase, isLengthOkay
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_AllVerifyErrorQuick_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	err := sv.AllVerifyErrorQuick(0, "test", "a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorQuick returns error -- match", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_CaseInsensitive(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"Hello"},
		ExpectedLines: []string{"hello"},
	}
	err := sv.AllVerifyErrorTestCase(0, "test", false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns error -- case insensitive", actual)
}

func Test_SliceValidator_LengthVerifyError_UptoWithAttach(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "h"}
	err := sv.AllVerifyErrorUptoLength(false, params, 5)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LengthVerifyError returns error -- upto exceeds with attach", actual)
}

func Test_SliceValidator_CompactMismatch_EmptyActual_OneExpected(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   nil,
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CompactMismatch returns nil -- nil actual one expected", actual)
}

func Test_SliceValidator_CompactMismatch_EmptyExpected(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: nil,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CompactMismatch returns nil -- empty actual nil expected", actual)
}
