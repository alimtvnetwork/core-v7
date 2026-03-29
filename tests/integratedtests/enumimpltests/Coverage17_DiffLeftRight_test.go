package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DiffLeftRight: HasMismatch ──

func Test_Cov17_DiffLeftRight_HasMismatch_Regardless(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: "1"}
	actual := args.Map{"hasMismatch": d.HasMismatch(true)}
	expected := args.Map{"hasMismatch": false}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns false -- regardless same string", actual)
}

func Test_Cov17_DiffLeftRight_HasMismatch_Strict(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	actual := args.Map{"hasMismatch": d.HasMismatch(false)}
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns true -- strict different", actual)
}

func Test_Cov17_DiffLeftRight_HasMismatchRegardlessOfType(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	actual := args.Map{"hasMismatch": d.HasMismatchRegardlessOfType()}
	expected := args.Map{"hasMismatch": true}
	expected.ShouldBeEqual(t, 0, "HasMismatchRegardlessOfType returns true -- different", actual)
}

// ── DiffLeftRight: IsEqual ──

func Test_Cov17_DiffLeftRight_IsEqual_Regardless(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 42, Right: 42}
	actual := args.Map{"isEqual": d.IsEqual(true)}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- regardless same", actual)
}

func Test_Cov17_DiffLeftRight_IsEqual_Strict(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	actual := args.Map{"isEqual": d.IsEqual(false)}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- strict same", actual)
}

// ── DiffLeftRight: SpecificFullString ──

func Test_Cov17_DiffLeftRight_SpecificFullString(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "hello", Right: 42}
	l, r := d.SpecificFullString()
	actual := args.Map{"leftNotEmpty": l != "", "rightNotEmpty": r != ""}
	expected := args.Map{"leftNotEmpty": true, "rightNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "SpecificFullString returns non-empty -- valid", actual)
}

// ── DiffLeftRight: DiffString ──

func Test_Cov17_DiffLeftRight_DiffString_Same(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	actual := args.Map{"empty": d.DiffString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffString returns empty -- same values", actual)
}

func Test_Cov17_DiffLeftRight_DiffString_Different(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	actual := args.Map{"notEmpty": d.DiffString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffString returns non-empty -- different", actual)
}

// ── DiffLeftRight: JsonString nil ──

func Test_Cov17_DiffLeftRight_JsonString_Nil(t *testing.T) {
	var d *enumimpl.DiffLeftRight
	actual := args.Map{"empty": d.JsonString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- nil receiver", actual)
}

// ── DiffLeftRight: Types / IsSameTypeSame ──

func Test_Cov17_DiffLeftRight_IsSameTypeSame_True(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	actual := args.Map{"sameType": d.IsSameTypeSame()}
	expected := args.Map{"sameType": true}
	expected.ShouldBeEqual(t, 0, "IsSameTypeSame returns true -- both strings", actual)
}

func Test_Cov17_DiffLeftRight_IsSameTypeSame_False(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}
	actual := args.Map{"sameType": d.IsSameTypeSame()}
	expected := args.Map{"sameType": false}
	expected.ShouldBeEqual(t, 0, "IsSameTypeSame returns false -- string vs int", actual)
}

// ── differCheckerImpl ──

func Test_Cov17_DifferChecker_GetSingleDiffResult_Left(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "L"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns left -- isLeft true", actual)
}

func Test_Cov17_DifferChecker_GetSingleDiffResult_Right(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(false, "L", "R")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "R"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns right -- isLeft false", actual)
}

func Test_Cov17_DifferChecker_GetResultOnKeyMissing(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "GetResultOnKeyMissing returns lVal -- key missing in right", actual)
}

func Test_Cov17_DifferChecker_IsEqual_Regardless(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, "1")
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- regardless same string", actual)
}

func Test_Cov17_DifferChecker_IsEqual_Strict(t *testing.T) {
	result := enumimpl.DefaultDiffCheckerImpl.IsEqual(false, 1, 1)
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- strict same", actual)
}

func Test_Cov17_DifferChecker_AsDifferChecker(t *testing.T) {
	checker := enumimpl.DefaultDiffCheckerImpl.AsDifferChecker()
	actual := args.Map{"notNil": checker != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsDifferChecker returns non-nil -- valid", actual)
}

// ── leftRightDiffCheckerImpl ──

func Test_Cov17_LeftRightDiffChecker_GetSingleDiffResult(t *testing.T) {
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "L", "R")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight GetSingleDiffResult returns non-nil -- valid", actual)
}

func Test_Cov17_LeftRightDiffChecker_GetResultOnKeyMissing(t *testing.T) {
	result := enumimpl.LeftRightDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LeftRight GetResultOnKeyMissing returns non-empty -- valid", actual)
}

func Test_Cov17_LeftRightDiffChecker_IsEqual(t *testing.T) {
	result := enumimpl.LeftRightDiffCheckerImpl.IsEqual(true, "a", "a")
	actual := args.Map{"isEqual": result}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEqual returns true -- same", actual)
}

func Test_Cov17_LeftRightDiffChecker_AsChecker(t *testing.T) {
	checker := enumimpl.LeftRightDiffCheckerImpl.AsChecker()
	actual := args.Map{"notNil": checker != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight AsChecker returns non-nil -- valid", actual)
}

// ── Format / FormatUsingFmt ──

func Test_Cov17_Format_Valid(t *testing.T) {
	result := enumimpl.Format("MyEnum", "Invalid", "0", "Enum of {type-name} - {name} - {value}")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Enum of MyEnum - Invalid - 0"}
	expected.ShouldBeEqual(t, 0, "Format returns correct string -- valid template", actual)
}

type cov17Formatter struct{}

func (f cov17Formatter) TypeName() string    { return "TestEnum" }
func (f cov17Formatter) Name() string        { return "Active" }
func (f cov17Formatter) ValueString() string { return "1" }

func Test_Cov17_FormatUsingFmt_Valid(t *testing.T) {
	result := enumimpl.FormatUsingFmt(cov17Formatter{}, "{type-name}.{name}={value}")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "TestEnum.Active=1"}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt returns correct string -- valid formatter", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_Cov17_PrependJoin(t *testing.T) {
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "PrependJoin returns correct string -- dot joiner", actual)
}

func Test_Cov17_JoinPrependUsingDot(t *testing.T) {
	result := enumimpl.JoinPrependUsingDot("root", "child")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root.child"}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot returns correct string -- root.child", actual)
}

// ── OnlySupportedErr ──

func Test_Cov17_OnlySupportedErr_AllSupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- all supported", actual)
}

func Test_Cov17_OnlySupportedErr_HasUnsupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- unsupported names", actual)
}

func Test_Cov17_OnlySupportedErr_EmptyAll(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{}, "a")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- empty allNames", actual)
}

// ── UnsupportedNames ──

func Test_Cov17_UnsupportedNames(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a", "c")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "b"}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns unsupported -- b only", actual)
}

// ── KeyAnyVal methods ──

func Test_Cov17_KeyAnyVal_IsString_True(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: "hello"}
	actual := args.Map{"isString": kv.IsString()}
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString returns true -- string value", actual)
}

func Test_Cov17_KeyAnyVal_IsString_False(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "name", AnyValue: 42}
	actual := args.Map{"isString": kv.IsString()}
	expected := args.Map{"isString": false}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString returns false -- int value", actual)
}

func Test_Cov17_KeyAnyVal_String_String(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "status", AnyValue: "active"}
	actual := args.Map{"notEmpty": kv.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal String returns non-empty -- string enum", actual)
}

func Test_Cov17_KeyAnyVal_String_Int(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "priority", AnyValue: 1}
	actual := args.Map{"notEmpty": kv.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal String returns non-empty -- int enum", actual)
}

func Test_Cov17_KeyAnyVal_Accessors(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "k", AnyValue: 42}
	actual := args.Map{
		"keyString":    kv.KeyString(),
		"anyValString": kv.AnyValString(),
		"wrapKey":      kv.WrapKey(),
		"wrapValue":    kv.WrapValue(),
		"valInt":       kv.ValInt(),
	}
	expected := args.Map{
		"keyString":    "k",
		"anyValString": "42",
		"wrapKey":      "\"k\"",
		"wrapValue":    "\"%!s(int=42)\"",
		"valInt":       42,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal accessors return correct -- all methods", actual)
}

func Test_Cov17_KeyAnyVal_KeyValInteger(t *testing.T) {
	kv := enumimpl.KeyAnyVal{Key: "x", AnyValue: 5}
	kvi := kv.KeyValInteger()
	actual := args.Map{"key": kvi.Key, "val": kvi.ValueInteger}
	expected := args.Map{"key": "x", "val": 5}
	expected.ShouldBeEqual(t, 0, "KeyValInteger returns correct -- from KeyAnyVal", actual)
}

// ── KeyValInteger methods ──

func Test_Cov17_KeyValInteger_Accessors(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "k", ValueInteger: 10}
	kav := kvi.KeyAnyVal()
	actual := args.Map{
		"wrapKey":   kvi.WrapKey(),
		"wrapValue": kvi.WrapValue(),
		"isString":  kvi.IsString(),
		"kavKey":    kav.Key,
	}
	expected := args.Map{
		"wrapKey":   "\"k\"",
		"wrapValue": "\"%!s(int=10)\"",
		"isString":  false,
		"kavKey":    "k",
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger accessors return correct -- all methods", actual)
}

func Test_Cov17_KeyValInteger_String_Int(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "priority", ValueInteger: 1}
	actual := args.Map{"notEmpty": kvi.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyValInteger String returns non-empty -- int", actual)
}

// ── AllNameValues ──

func Test_Cov17_AllNameValues(t *testing.T) {
	names := []string{"a", "b"}
	values := []int{1, 2}
	result := enumimpl.AllNameValues(names, values)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns 2 -- matching slices", actual)
}

// ── KeyAnyValues ──

func Test_Cov17_KeyAnyValues_Empty(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{}, []int{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns empty -- empty input", actual)
}

func Test_Cov17_KeyAnyValues_Valid(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})
	actual := args.Map{"len": len(result), "firstKey": result[0].Key}
	expected := args.Map{"len": 2, "firstKey": "a"}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns 2 -- valid input", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_Cov17_IntegersRangesOfAnyVal(t *testing.T) {
	result := enumimpl.IntegersRangesOfAnyVal([]int{3, 1, 2})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": 1, "last": 3}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal returns sorted -- valid slice", actual)
}

// ── NameWithValue ──

func Test_Cov17_NameWithValue(t *testing.T) {
	result := enumimpl.NameWithValue("Active")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns non-empty -- valid", actual)
}

// ── ConvEnumAnyValToInteger: string type ──

func Test_Cov17_ConvEnumAnyValToInteger_String(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger("hello")
	actual := args.Map{"isMinInt": result < 0}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string", actual)
}

func Test_Cov17_ConvEnumAnyValToInteger_Int(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int", actual)
}
