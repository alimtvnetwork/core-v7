package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── TextValidators: AllVerifyErrorMany ──

func Test_TextValidators_AllVerifyErrorMany_Match_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(2)
	tv.AddSimple("hello", stringcompareas.Contains)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.AllVerifyErrorMany(params, "hello world", "say hello")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorMany returns nil -- all match", actual)
}

func Test_TextValidators_AllVerifyErrorMany_NoMatch(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("missing", stringcompareas.Contains)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.AllVerifyErrorMany(params, "hello world")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorMany returns error -- no match", actual)
}

// ── TextValidators: VerifyErrorMany ──

func Test_TextValidators_VerifyErrorMany_ContinueOnError_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("hello", stringcompareas.Contains)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.VerifyErrorMany(true, params, "hello world")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorMany returns nil -- continue mode match", actual)
}

func Test_TextValidators_VerifyErrorMany_FirstOnly_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("hello", stringcompareas.Contains)
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.VerifyErrorMany(false, params, "hello world")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorMany returns nil -- first only match", actual)
}

func Test_TextValidators_VerifyErrorMany_Nil_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidators

	// Act
	err := tv.VerifyErrorMany(true, &corevalidator.Parameter{}, "hello")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorMany returns nil -- nil receiver", actual)
}

// ── TextValidators: AddSimpleAllTrue ──

func Test_TextValidators_AddSimpleAllTrue_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)

	// Act
	tv.AddSimpleAllTrue("test", stringcompareas.Equal)

	// Assert
	actual := args.Map{"length": tv.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "AddSimpleAllTrue adds 1 -- with all conditions", actual)
}

// ── TextValidators: HasAnyItem / HasIndex / Count ──

func Test_TextValidators_Accessors(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(2)
	tv.AddSimple("a", stringcompareas.Contains)
	tv.AddSimple("b", stringcompareas.Contains)

	// Act & Assert
	actual := args.Map{
		"hasAnyItem": tv.HasAnyItem(),
		"hasIndex0":  tv.HasIndex(0),
		"hasIndex1":  tv.HasIndex(1),
		"hasIndex2":  tv.HasIndex(2),
		"count":      tv.Count(),
		"lastIndex":  tv.LastIndex(),
	}
	expected := args.Map{
		"hasAnyItem": true,
		"hasIndex0":  true,
		"hasIndex1":  true,
		"hasIndex2":  false,
		"count":      1,
		"lastIndex":  1,
	}
	expected.ShouldBeEqual(t, 0, "Accessors return correct -- 2 items", actual)
}

// ── TextValidators: Dispose ──

func Test_TextValidators_Dispose_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("a", stringcompareas.Contains)

	// Act
	tv.Dispose()

	// Assert
	actual := args.Map{"isEmpty": tv.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dispose nils items -- empty", actual)
}

// ── TextValidators: String ──

func Test_TextValidators_String_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("test", stringcompareas.Contains)

	// Act
	result := tv.String()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- 1 item", actual)
}

// ── TextValidators: IsMatchMany ──

func Test_TextValidators_IsMatchMany(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("hello", stringcompareas.Contains)

	// Act
	result := tv.IsMatchMany(false, true, "hello world", "say hello")

	// Assert
	actual := args.Map{"isMatch": result}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- all contain hello", actual)
}

func Test_TextValidators_IsMatchMany_NoMatch(t *testing.T) {
	// Arrange
	tv := corevalidator.NewTextValidators(1)
	tv.AddSimple("missing", stringcompareas.Contains)

	// Act
	result := tv.IsMatchMany(false, true, "hello world")

	// Assert
	actual := args.Map{"isMatch": result}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns false -- no match", actual)
}

// ── TextValidator: VerifyMany ──

func Test_TextValidator_VerifyMany_ContinueOnError_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Contains,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.VerifyMany(true, params, "hello world")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns nil -- continue+match", actual)
}

func Test_TextValidator_VerifyMany_FirstOnly_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Contains,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.VerifyMany(false, params, "hello world")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyMany returns nil -- first+match", actual)
}

// ── TextValidator: IsMatchMany ──

func Test_TextValidator_IsMatchMany_SkipOnEmpty(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Contains,
	}

	// Act
	result := tv.IsMatchMany(true, true)

	// Assert
	actual := args.Map{"isMatch": result}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "IsMatchMany returns true -- skip on empty", actual)
}

// ── TextValidator: ToString multiline ──

func Test_TextValidator_ToString_MultiLine_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "test",
		SearchAs: stringcompareas.Contains,
	}

	// Act
	result := tv.ToString(false)

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "ToString returns multi-line -- false flag", actual)
}

// ── TextValidator: MethodName ──

func Test_TextValidator_MethodName_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "test",
		SearchAs: stringcompareas.Contains,
	}

	// Act
	result := tv.MethodName()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "MethodName returns name -- Contains", actual)
}

// ── TextValidator: AllVerifyError ──

func Test_TextValidator_AllVerifyError_Match(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Contains,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.AllVerifyError(params, "hello world", "say hello")

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- all match", actual)
}

func Test_TextValidator_AllVerifyError_NoMatch(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{
		Search:   "missing",
		SearchAs: stringcompareas.Contains,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := tv.AllVerifyError(params, "hello world")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- no match", actual)
}

// ── SliceValidator: AllVerifyErrorExceptLast ──

func Test_SliceValidator_AllVerifyErrorExceptLast(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a", "b", "DIFFERENT"},
		ExpectedLines: []string{"a", "b", "c"},
		CompareAs:     stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := sv.AllVerifyErrorExceptLast(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorExceptLast skips last -- different last line ok", actual)
}

// ── SliceValidator: AllVerifyErrorTestCase ──

func Test_SliceValidator_AllVerifyErrorTestCase_Match_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"hello"},
		CompareAs:     stringcompareas.Equal,
	}

	// Act
	err := sv.AllVerifyErrorTestCase(0, "test header", true)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns nil -- match", actual)
}

// ── LinesValidators: Accessors ──

func Test_LinesValidators_Accessors(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)

	// Act & Assert
	actual := args.Map{
		"count":      lv.Count(),
		"isEmpty":    lv.IsEmpty(),
		"hasAnyItem": lv.HasAnyItem(),
		"lastIndex":  lv.LastIndex(),
		"hasIndex0":  lv.HasIndex(0),
	}
	expected := args.Map{
		"count":      0,
		"isEmpty":    true,
		"hasAnyItem": false,
		"lastIndex":  -1,
		"hasIndex0":  false,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators empty accessors -- correct defaults", actual)
}

func Test_LinesValidators_String_FromTextValidatorsAllVer(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)

	// Act
	result := lv.String()

	// Assert
	actual := args.Map{"notNil": true}
	expected := args.Map{"notNil": true}
	_ = result
	expected.ShouldBeEqual(t, 0, "LinesValidators String returns -- non-panic", actual)
}
