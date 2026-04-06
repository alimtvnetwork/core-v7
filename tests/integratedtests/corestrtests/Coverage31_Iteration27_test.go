package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Factories
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_NewValidValue(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewValidValue", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "hello", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns non-empty -- with args", actual)
	})
}

func Test_I27_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewValidValueEmpty", func() {
		vv := corestr.NewValidValueEmpty()
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns empty -- with args", actual)
	})
}

func Test_I27_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_InvalidNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"valid": vv.IsValid, "msg": vv.Message}
		expected := args.Map{"valid": false, "msg": ""}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns error -- with args", actual)
	})
}

func Test_I27_ValidValue_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_InvalidWithMessage", func() {
		vv := corestr.InvalidValidValue("err")
		actual := args.Map{"valid": vv.IsValid, "msg": vv.Message}
		expected := args.Map{"valid": false, "msg": "err"}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns error -- with args", actual)
	})
}

func Test_I27_ValidValue_NewUsingAny(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewUsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, 42)
		actual := args.Map{"valid": vv.IsValid, "notEmpty": vv.Value != ""}
		expected := args.Map{"valid": true, "notEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny returns non-empty -- with args", actual)
	})
}

func Test_I27_ValidValue_NewUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewUsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, 42)
		actual := args.Map{"notEmpty": vv.Value != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Bytes, Checks, Trim
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("abc")
		b1 := vv.ValueBytesOnce()
		b2 := vv.ValueBytesOnce() // cached
		actual := args.Map{"len": len(b1), "same": &b1[0] == &b2[0]}
		expected := args.Map{"len": 3, "same": true}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- with args", actual)
	})
}
func Test_I27_ValidValue_IsEmpty_IsWhitespace_Trim(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsEmpty_IsWhitespace_Trim", func() {
		vv := corestr.NewValidValue("  hi  ")
		actual := args.Map{"empty": vv.IsEmpty(), "ws": vv.IsWhitespace(), "trim": vv.Trim()}
		expected := args.Map{"empty": false, "ws": false, "trim": "hi"}
		expected.ShouldBeEqual(t, 0, "IsEmpty/IsWhitespace/Trim returns empty -- with args", actual)
	})
}

func Test_I27_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		actual := args.Map{"hv": vv.HasValidNonEmpty(), "hvw": vv.HasValidNonWhitespace(), "safe": vv.HasSafeNonEmpty()}
		expected := args.Map{"hv": true, "hvw": true, "safe": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Type conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueBool", func() {
		vv1 := corestr.NewValidValue("true")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("")
		actual := args.Map{"t": vv1.ValueBool(), "f": vv2.ValueBool(), "e": vv3.ValueBool()}
		expected := args.Map{"t": true, "f": false, "e": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueBool", actual)
	})
}

func Test_I27_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueInt", func() {
		vv1 := corestr.NewValidValue("42")
		vv2 := corestr.NewValidValue("abc")
		actual := args.Map{"val": vv1.ValueInt(0), "def": vv2.ValueInt(99), "defInt": vv1.ValueDefInt()}
		expected := args.Map{"val": 42, "def": 99, "defInt": 42}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueInt", actual)
	})
}

func Test_I27_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueByte", func() {
		vv1 := corestr.NewValidValue("100")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("300")
		vv4 := corestr.NewValidValue("-1")
		actual := args.Map{"val": vv1.ValueByte(0), "err": vv2.ValueByte(7), "over": vv3.ValueByte(5), "neg": vv4.ValueByte(9)}
		expected := args.Map{"val": byte(100), "err": byte(0), "over": byte(255), "neg": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueByte", actual)
	})
}

func Test_I27_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueDefByte", func() {
		vv1 := corestr.NewValidValue("50")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("999")
		vv4 := corestr.NewValidValue("-5")
		actual := args.Map{"val": vv1.ValueDefByte(), "err": vv2.ValueDefByte(), "over": vv3.ValueDefByte(), "neg": vv4.ValueDefByte()}
		expected := args.Map{"val": byte(50), "err": byte(0), "over": byte(255), "neg": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueDefByte", actual)
	})
}

func Test_I27_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueFloat64", func() {
		vv1 := corestr.NewValidValue("3.14")
		vv2 := corestr.NewValidValue("abc")
		actual := args.Map{"close": vv1.ValueFloat64(0) > 3.1, "def": vv2.ValueFloat64(1.0), "defFloat": vv1.ValueDefFloat64() > 3.1}
		expected := args.Map{"close": true, "def": 1.0, "defFloat": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueFloat64", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — String matching
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Is", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"is": vv.Is("hello"), "isNot": vv.Is("world")}
		expected := args.Map{"is": true, "isNot": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Is", actual)
	})
}

func Test_I27_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		actual := args.Map{"found": vv.IsAnyOf("a", "b"), "notFound": vv.IsAnyOf("x", "y"), "empty": vv.IsAnyOf()}
		expected := args.Map{"found": true, "notFound": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsAnyOf", actual)
	})
}

func Test_I27_ValidValue_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsContains_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{
			"contains":    vv.IsContains("world"),
			"notContains": vv.IsContains("xyz"),
			"anyContains": vv.IsAnyContains("xyz", "hello"),
			"anyNone":     vv.IsAnyContains("xyz"),
			"anyEmpty":    vv.IsAnyContains(),
		}
		expected := args.Map{"contains": true, "notContains": false, "anyContains": true, "anyNone": false, "anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsContains/IsAnyContains", actual)
	})
}

func Test_I27_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		actual := args.Map{"eq": vv.IsEqualNonSensitive("hello"), "neq": vv.IsEqualNonSensitive("world")}
		expected := args.Map{"eq": true, "neq": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsEqualNonSensitive", actual)
	})
}

func Test_I27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"match": vv.IsRegexMatches(re), "nilRegex": vv.IsRegexMatches(nil)}
		expected := args.Map{"match": true, "nilRegex": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsRegexMatches", actual)
	})
}

func Test_I27_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindString", func() {
		vv := corestr.NewValidValue("abc123def")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"found": vv.RegexFindString(re), "nilRegex": vv.RegexFindString(nil)}
		expected := args.Map{"found": "123", "nilRegex": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindString", actual)
	})
}

func Test_I27_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items := vv.RegexFindAllStrings(re, -1)
		actual := args.Map{"len": len(items)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindAllStrings", actual)
	})
}

func Test_I27_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStrings_Nil", func() {
		vv := corestr.NewValidValue("a1b2")
		items := vv.RegexFindAllStrings(nil, -1)
		actual := args.Map{"len": len(items)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- RegexFindAllStrings nil", actual)
	})
}

func Test_I27_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
		actual := args.Map{"len": len(items), "hasAny": hasAny}
		expected := args.Map{"len": 3, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindAllStringsWithFlag", actual)
	})
}

func Test_I27_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		vv := corestr.NewValidValue("a1b2")
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{"len": len(items), "hasAny": hasAny}
		expected := args.Map{"len": 0, "hasAny": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- RegexFindAllStringsWithFlag nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Split
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		actual := args.Map{"len": len(vv.Split(","))}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Split", actual)
	})
}

func Test_I27_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_SplitNonEmpty", func() {
		vv := corestr.NewValidValue("a,,b")
		result := vv.SplitNonEmpty(",")
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- SplitNonEmpty", actual)
	})
}

func Test_I27_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_SplitTrimNonWhitespace", func() {
		vv := corestr.NewValidValue("a , , b")
		result := vv.SplitTrimNonWhitespace(",")
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- SplitTrimNonWhitespace", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Clone, String, JSON, Serialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()
		actual := args.Map{"val": cloned.Value, "notSame": cloned != vv}
		expected := args.Map{"val": "hello", "notSame": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clone", actual)
	})
}

func Test_I27_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clone_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"nil": vv.Clone() == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- Clone nil", actual)
	})
}

func Test_I27_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_String_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"val": vv.String()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- String nil", actual)
	})
}

func Test_I27_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_String", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"val": vv.String()}
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- String", actual)
	})
}

func Test_I27_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"notEmpty": vv.FullString() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- FullString", actual)
	})
}

func Test_I27_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_FullString_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"val": vv.FullString()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- FullString nil", actual)
	})
}

func Test_I27_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clear_Dispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clear", actual)

		vv2 := corestr.NewValidValue("x")
		vv2.Dispose()
		(*corestr.ValidValue)(nil).Clear()
		(*corestr.ValidValue)(nil).Dispose()
	})
}

func Test_I27_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Json", func() {
		vv := corestr.NewValidValue("hello")
		jr := vv.Json()
		actual := args.Map{"noErr": !jr.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Json", actual)
	})
}

func Test_I27_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_JsonPtr", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"notNil": vv.JsonPtr() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- JsonPtr", actual)
	})
}

func Test_I27_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("hello")
		b, err := vv.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Serialize", actual)
	})
}

func Test_I27_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Deserialize", func() {
		vv := corestr.NewValidValue("hello")
		target := &corestr.ValidValue{}
		err := vv.Deserialize(target)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Deserialize", actual)
	})
}

func Test_I27_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ParseInjectUsingJson", func() {
		vv := &corestr.ValidValue{}
		jr := corejson.New(corestr.ValidValue{Value: "test", IsValid: true})
		parsed, err := vv.ParseInjectUsingJson(&jr)
		actual := args.Map{"noErr": err == nil, "notNil": parsed != nil}
		expected := args.Map{"noErr": true, "notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ParseInjectUsingJson", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Basic", func() {
		kv := &corestr.KeyAnyValuePair{Key: "name", Value: 42}
		actual := args.Map{
			"key": kv.KeyName(), "varName": kv.VariableName(),
			"valAny": kv.ValueAny(), "isVarEq": kv.IsVariableNameEqual("name"),
		}
		expected := args.Map{"key": "name", "varName": "name", "valAny": 42, "isVarEq": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- basic", actual)
	})
}

func Test_I27_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueNull", func() {
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		kv2 := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		var kv3 *corestr.KeyAnyValuePair
		actual := args.Map{"null": kv1.IsValueNull(), "notNull": kv2.IsValueNull(), "nilRcv": kv3.IsValueNull()}
		expected := args.Map{"null": true, "notNull": false, "nilRcv": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- IsValueNull", actual)
	})
}

func Test_I27_KeyAnyValuePair_HasNonNull_HasValue(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_HasNonNull_HasValue", func() {
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		kv2 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv3 *corestr.KeyAnyValuePair
		actual := args.Map{"has1": kv1.HasNonNull(), "has2": kv2.HasNonNull(), "hasNil": kv3.HasNonNull(), "hasVal": kv1.HasValue(), "hasValNil": kv3.HasValue()}
		expected := args.Map{"has1": true, "has2": false, "hasNil": false, "hasVal": true, "hasValNil": false}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- HasNonNull/HasValue", actual)
	})
}

func Test_I27_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueEmptyString", func() {
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv2 *corestr.KeyAnyValuePair
		actual := args.Map{"empty": kv1.IsValueEmptyString(), "nilRcv": kv2.IsValueEmptyString()}
		expected := args.Map{"empty": true, "nilRcv": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns empty -- IsValueEmptyString", actual)
	})
}

func Test_I27_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueWhitespace", func() {
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv2 *corestr.KeyAnyValuePair
		actual := args.Map{"ws": kv1.IsValueWhitespace(), "nilRcv": kv2.IsValueWhitespace()}
		expected := args.Map{"ws": true, "nilRcv": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- IsValueWhitespace", actual)
	})
}

func Test_I27_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_ValueString", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s1 := kv.ValueString()
		s2 := kv.ValueString() // cached path
		actual := args.Map{"notEmpty": s1 != "", "same": s1 == s2}
		expected := args.Map{"notEmpty": true, "same": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns non-empty -- ValueString", actual)
	})
}

func Test_I27_KeyAnyValuePair_ValueString_Null(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_ValueString_Null", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		s := kv.ValueString()
		_ = s // covers GetOnce path
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns non-empty -- ValueString null", actual)
	})
}

func Test_I27_KeyAnyValuePair_Compile_String(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Compile_String", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"compile": kv.Compile(), "str": kv.String()}
		expected := args.Map{"compile": "{k:v}", "str": "{k:v}"}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Compile/String", actual)
	})
}

func Test_I27_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_SerializeMust", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()
		actual := args.Map{"hasBytes": len(b) > 0}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- SerializeMust", actual)
	})
}

func Test_I27_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Serialize", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Serialize", actual)
	})
}

func Test_I27_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Json", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.Json()
		actual := args.Map{"noErr": !jr.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Json", actual)
	})
}

func Test_I27_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_JsonPtr", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"notNil": kv.JsonPtr() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- JsonPtr", actual)
	})
}

func Test_I27_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_AsJsonContractsBinder", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"notNil": kv.AsJsonContractsBinder() != nil, "jsoner": kv.AsJsoner() != nil, "selfInj": kv.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"notNil": true, "jsoner": true, "selfInj": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- As* methods", actual)
	})
}

func Test_I27_KeyAnyValuePair_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_JsonParseSelfInject", func() {
		kv := &corestr.KeyAnyValuePair{}
		jr := corejson.New(corestr.KeyAnyValuePair{Key: "test", Value: "val"})
		err := kv.JsonParseSelfInject(&jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- JsonParseSelfInject", actual)
	})
}

func Test_I27_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Clear_Dispose", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		actual := args.Map{"key": kv.Key, "valNull": kv.Value == nil}
		expected := args.Map{"key": "", "valNull": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Clear", actual)

		kv2 := &corestr.KeyAnyValuePair{Key: "x", Value: "y"}
		kv2.Dispose()
		(*corestr.KeyAnyValuePair)(nil).Clear()
		(*corestr.KeyAnyValuePair)(nil).Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues
// ══════════════════════════════════════════════════════════════════════════════

func Test_I27_ValidValues_NewEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewEmpty", func() {
		vvs := corestr.EmptyValidValues()
		actual := args.Map{"len": vvs.Length(), "empty": vvs.IsEmpty(), "count": vvs.Count()}
		expected := args.Map{"len": 0, "empty": true, "count": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- NewEmpty", actual)
	})
}

func Test_I27_ValidValues_NewWithCap(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewWithCap", func() {
		vvs := corestr.NewValidValues(5)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- NewWithCap", actual)
	})
}

func Test_I27_ValidValues_NewUsingValues(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewUsingValues", func() {
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)
		actual := args.Map{"len": vvs.Length(), "hasAny": vvs.HasAnyItem()}
		expected := args.Map{"len": 2, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- NewUsingValues", actual)
	})
}

func Test_I27_ValidValues_NewUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewUsingValues_Empty", func() {
		vvs := corestr.NewValidValuesUsingValues()
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- NewUsingValues empty", actual)
	})
}

func Test_I27_ValidValues_Add_AddFull(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Add_AddFull", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		vvs.AddFull(false, "c", "err")
		actual := args.Map{"len": vvs.Length(), "lastIdx": vvs.LastIndex()}
		expected := args.Map{"len": 3, "lastIdx": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Add/AddFull", actual)
	})
}

func Test_I27_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_HasIndex", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		actual := args.Map{"has0": vvs.HasIndex(0), "has1": vvs.HasIndex(1), "has2": vvs.HasIndex(2)}
		expected := args.Map{"has0": true, "has1": true, "has2": false}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- HasIndex", actual)
	})
}

func Test_I27_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValueAt", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("hello").Add("world")
		actual := args.Map{"v0": vvs.SafeValueAt(0), "v1": vvs.SafeValueAt(1), "oob": vvs.SafeValueAt(99)}
		expected := args.Map{"v0": "hello", "v1": "world", "oob": ""}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValueAt", actual)
	})
}

func Test_I27_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("hello")
		vvs.AddFull(false, "bad", "err")
		actual := args.Map{"v0": vvs.SafeValidValueAt(0), "v1": vvs.SafeValidValueAt(1), "oob": vvs.SafeValidValueAt(99)}
		expected := args.Map{"v0": "hello", "v1": "", "oob": ""}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValidValueAt", actual)
	})
}

func Test_I27_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		result := vvs.SafeValuesAtIndexes(0, 2)
		actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
		expected := args.Map{"len": 2, "first": "a", "second": "c"}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValuesAtIndexes", actual)
	})
}

func Test_I27_ValidValues_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValuesAtIndexes_Empty", func() {
		vvs := corestr.EmptyValidValues()
		result := vvs.SafeValuesAtIndexes()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- SafeValuesAtIndexes empty", actual)
	})
}

func Test_I27_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValidValuesAtIndexes", actual)
	})
}

func Test_I27_ValidValues_Strings_FullStrings_String(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Strings_FullStrings_String", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		actual := args.Map{"strLen": len(vvs.Strings()), "fullLen": len(vvs.FullStrings()), "strNotEmpty": vvs.String() != ""}
		expected := args.Map{"strLen": 2, "fullLen": 2, "strNotEmpty": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Strings/FullStrings/String", actual)
	})
}

func Test_I27_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Strings_Empty", func() {
		vvs := corestr.EmptyValidValues()
		actual := args.Map{"strLen": len(vvs.Strings()), "fullLen": len(vvs.FullStrings())}
		expected := args.Map{"strLen": 0, "fullLen": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Strings empty", actual)
	})
}

func Test_I27_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})
		actual := args.Map{"len": len(found)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Find", actual)
	})
}

func Test_I27_ValidValues_Find_Break(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find_Break", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, index == 0
		})
		actual := args.Map{"len": len(found)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Find break", actual)
	})
}

func Test_I27_ValidValues_Find_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find_Empty", func() {
		vvs := corestr.EmptyValidValues()
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})
		actual := args.Map{"len": len(found)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Find empty", actual)
	})
}

func Test_I27_ValidValues_Adds_AddsPtr(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Adds_AddsPtr", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})
		vvs.AddsPtr(corestr.NewValidValue("c"))
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Adds/AddsPtr", actual)
	})
}

func Test_I27_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Adds_Empty", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Adds()
		vvs.AddsPtr()
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Adds empty", actual)
	})
}

func Test_I27_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddValidValues", func() {
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b").Add("c")
		vvs1.AddValidValues(vvs2)
		actual := args.Map{"len": vvs1.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- AddValidValues", actual)
	})
}

func Test_I27_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddValidValues_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs.AddValidValues(nil)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddValidValues nil", actual)
	})
}

func Test_I27_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew", func() {
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- ConcatNew", actual)
	})
}

func Test_I27_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew_EmptyClone", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(true)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty clone", actual)
	})
}

func Test_I27_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew_EmptyNoClone", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(false)
		actual := args.Map{"same": result == vvs}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty no clone", actual)
	})
}

func Test_I27_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashsetMap", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- AddHashsetMap", actual)
	})
}

func Test_I27_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashsetMap_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(nil)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddHashsetMap nil", actual)
	})
}

func Test_I27_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashset_Nil", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashset(nil)
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddHashset nil", actual)
	})
}

func Test_I27_ValidValues_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Hashmap_Map", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddFull(true, "key", "val")
		hm := vvs.Hashmap()
		m := vvs.Map()
		actual := args.Map{"hmNotNil": hm != nil, "mapLen": len(m)}
		expected := args.Map{"hmNotNil": true, "mapLen": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Hashmap/Map", actual)
	})
}

func Test_I27_ValidValues_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Hashmap_Empty", func() {
		vvs := corestr.EmptyValidValues()
		hm := vvs.Hashmap()
		actual := args.Map{"len": hm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Hashmap empty", actual)
	})
}

func Test_I27_ValidValues_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Length_Nil", func() {
		var vvs *corestr.ValidValues
		actual := args.Map{"len": vvs.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- Length nil", actual)
	})
}
