package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DiffLeftRight: HasMismatch ──

func Test_DiffLeftRight_HasMismatch_Regardless_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: "1"}

	// Act
	actual := args.Map{"hasMismatch": d.HasMismatch(true)}

	// Assert
	expected := args.Map{"hasMismatch": false}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns false -- regardless same string", actual)
}

func Test_DiffLeftRight_HasMismatch_Strict(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}

	// Act
	actual := args.Map{"hasMismatch": d.HasMismatch(false)}

	// Assert
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns true -- strict different", actual)
}

func Test_DiffLeftRight_HasMismatchRegardlessOfType(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"hasMismatch": d.HasMismatchRegardlessOfType()}

	// Assert
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasMismatchRegardlessOfType returns true -- different", actual)
}

// ── DiffLeftRight: IsEqual ──

func Test_DiffLeftRight_IsEqual_Regardless(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 42, Right: 42}

	// Act
	actual := args.Map{"isEqual": d.IsEqual(true)}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- regardless same", actual)
}

func Test_DiffLeftRight_IsEqual_Strict(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}

	// Act
	actual := args.Map{"isEqual": d.IsEqual(false)}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- strict same", actual)
}

// ── DiffLeftRight: SpecificFullString ──

func Test_DiffLeftRight_SpecificFullString_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "hello", Right: 42}
	l, r := d.SpecificFullString()

	// Act
	actual := args.Map{
		"leftNotEmpty": l != "",
		"rightNotEmpty": r != "",
	}

	// Assert
	expected := args.Map{
		"leftNotEmpty": true,
		"rightNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SpecificFullString returns non-empty -- valid", actual)
}

// ── DiffLeftRight: DiffString ──

func Test_DiffLeftRight_DiffString_Same_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}

	// Act
	actual := args.Map{"empty": d.DiffString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffString returns empty -- same values", actual)
}

func Test_DiffLeftRight_DiffString_Different_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"notEmpty": d.DiffString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffString returns non-empty -- different", actual)
}

// ── DiffLeftRight: JsonString nil ──

func Test_DiffLeftRight_JsonString_Nil_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	var d *enumimpl.DiffLeftRight

	// Act
	actual := args.Map{"empty": d.JsonString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- nil receiver", actual)
}

// ── DiffLeftRight: Types / IsSameTypeSame ──

func Test_DiffLeftRight_IsSameTypeSame_True(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{"sameType": d.IsSameTypeSame()}

	// Assert
	expected := args.Map{"sameType": true}
	expected.ShouldBeEqual(t, 0, "IsSameTypeSame returns true -- both strings", actual)
}

func Test_DiffLeftRight_IsSameTypeSame_False(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}

	// Act
	actual := args.Map{"sameType": d.IsSameTypeSame()}

	// Assert
	expected := args.Map{"sameType": false}
	expected.ShouldBeEqual(t, 0, "IsSameTypeSame returns false -- string vs int", actual)
}

// ── differCheckerImpl ──

func Test_DifferChecker_GetSingleDiffResult_Left(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "L"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns left -- isLeft true", actual)
}

func Test_DifferChecker_GetSingleDiffResult_Right(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(false, "L", "R")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "R"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns right -- isLeft false", actual)
}

func Test_DifferChecker_GetResultOnKeyMissing_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "GetResultOnKeyMissing returns lVal -- key missing in right", actual)
}

func Test_DifferChecker_IsEqual_Regardless(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, "1")

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- regardless same string", actual)
}

func Test_DifferChecker_IsEqual_Strict(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.IsEqual(false, 1, 1)

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- strict same", actual)
}

func Test_DifferChecker_AsDifferChecker_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl.AsDifferChecker()

	// Act
	actual := args.Map{"notNil": checker != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsDifferChecker returns non-nil -- valid", actual)
}

// ── leftRightDiffCheckerImpl ──

func Test_LeftRightDiffChecker_GetSingleDiffResult_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight GetSingleDiffResult returns non-nil -- valid", actual)
}

func Test_LeftRightDiffChecker_GetResultOnKeyMissing_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight GetResultOnKeyMissing returns non-empty -- valid", actual)
}

func Test_LeftRightDiffChecker_IsEqual_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.IsEqual(true, "a", "a")

	// Act
	actual := args.Map{"isEqual": result}

	// Assert
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEqual returns true -- same", actual)
}

func Test_LeftRightDiffChecker_AsChecker_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	checker := enumimpl.LeftRightDiffCheckerImpl.AsChecker()

	// Act
	actual := args.Map{"notNil": checker != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight AsChecker returns non-nil -- valid", actual)
}

// ── Format / FormatUsingFmt ──

func Test_Format_Valid(t *testing.T) {
	// Arrange
	result := enumimpl.Format("MyEnum", "Invalid", "0", "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "Enum of MyEnum - Invalid - 0"}
	expected.ShouldBeEqual(t, 0, "Format returns correct string -- valid template", actual)
}

type cov17Formatter struct{}

func (f cov17Formatter) TypeName() string    { return "TestEnum" }
func (f cov17Formatter) Name() string        { return "Active" }
func (f cov17Formatter) ValueString() string { return "1" }

func Test_FormatUsingFmt_Valid(t *testing.T) {
	// Arrange
	result := enumimpl.FormatUsingFmt(cov17Formatter{}, "{type-name}.{name}={value}")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "TestEnum.Active=1"}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt returns correct string -- valid formatter", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_PrependJoin_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "PrependJoin returns correct string -- dot joiner", actual)
}

func Test_JoinPrependUsingDot_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.JoinPrependUsingDot("root", "child")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "root.child"}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot returns correct string -- root.child", actual)
}

// ── OnlySupportedErr ──

func Test_OnlySupportedErr_AllSupported_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- all supported", actual)
}

func Test_OnlySupportedErr_HasUnsupported_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- unsupported names", actual)
}

func Test_OnlySupportedErr_EmptyAll_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- empty allNames", actual)
}

// ── UnsupportedNames ──

func Test_UnsupportedNames_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a", "c")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns unsupported -- b only", actual)
}

// ── KeyAnyVal methods ──

func Test_KeyAnyVal_IsString_True_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: "hello"}

	// Act
	actual := args.Map{"isString": kv.IsString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString returns true -- string value", actual)
}

func Test_KeyAnyVal_IsString_False(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: 42}

	// Act
	actual := args.Map{"isString": kv.IsString()}

	// Assert
	expected := args.Map{"isString": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString returns false -- int value", actual)
}

func Test_KeyAnyVal_String_String(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "status", AnyValue: "active"}

	// Act
	actual := args.Map{"notEmpty": kv.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal String returns non-empty -- string enum", actual)
}

func Test_KeyAnyVal_String_Int(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "priority", AnyValue: 1}

	// Act
	actual := args.Map{"notEmpty": kv.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal String returns non-empty -- int enum", actual)
}

func Test_KeyAnyVal_Accessors(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "k", AnyValue: 42}

	// Act
	actual := args.Map{
		"keyString":    kv.KeyString(),
		"anyValString": kv.AnyValString(),
		"wrapKey":      kv.WrapKey(),
		"wrapValue":    kv.WrapValue(),
		"valInt":       kv.ValInt(),
	}

	// Assert
	expected := args.Map{
		"keyString":    "k",
		"anyValString": "42",
		"wrapKey":      "\"k\"",
		"wrapValue":    "\"%!s(int=42)\"",
		"valInt":       42,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal accessors return correct -- all methods", actual)
}

func Test_KeyAnyVal_KeyValInteger_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	kv := enumimpl.KeyAnyVal{Key: "x", AnyValue: 5}
	kvi := kv.KeyValInteger()

	// Act
	actual := args.Map{
		"key": kvi.Key,
		"val": kvi.ValueInteger,
	}

	// Assert
	expected := args.Map{
		"key": "x",
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger returns correct -- from KeyAnyVal", actual)
}

// ── KeyValInteger methods ──

func Test_KeyValInteger_Accessors(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "k", ValueInteger: 10}
	kav := kvi.KeyAnyVal()

	// Act
	actual := args.Map{
		"wrapKey":   kvi.WrapKey(),
		"wrapValue": kvi.WrapValue(),
		"isString":  kvi.IsString(),
		"kavKey":    kav.Key,
	}

	// Assert
	expected := args.Map{
		"wrapKey":   "\"k\"",
		"wrapValue": "\"%!s(int=10)\"",
		"isString":  false,
		"kavKey":    "k",
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger accessors return correct -- all methods", actual)
}

func Test_KeyValInteger_String_Int(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "priority", ValueInteger: 1}

	// Act
	actual := args.Map{"notEmpty": kvi.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValInteger String returns non-empty -- int", actual)
}

// ── AllNameValues ──

func Test_AllNameValues_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	names := []string{"a", "b"}
	values := []int{1, 2}
	result := enumimpl.AllNameValues(names, values)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns 2 -- matching slices", actual)
}

// ── KeyAnyValues ──

func Test_KeyAnyValues_Empty_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{}, []int{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns empty -- empty input", actual)
}

func Test_KeyAnyValues_Valid(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"firstKey": result[0].Key,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstKey": "a",
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns 2 -- valid input", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_IntegersRangesOfAnyVal_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]int{3, 1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal returns sorted -- valid slice", actual)
}

// ── NameWithValue ──

func Test_NameWithValue_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.NameWithValue("Active")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns non-empty -- valid", actual)
}

// ── ConvEnumAnyValToInteger: string type ──

func Test_ConvEnumAnyValToInteger_String_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger("hello")

	// Act
	actual := args.Map{"isMinInt": result < 0}

	// Assert
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string", actual)
}

func Test_ConvEnumAnyValToInteger_Int_FromDiffLeftRightHasMism(t *testing.T) {
	// Arrange
	result := enumimpl.ConvEnumAnyValToInteger(42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int", actual)
}
