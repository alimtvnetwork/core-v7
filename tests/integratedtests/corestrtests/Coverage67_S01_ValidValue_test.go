package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Constructors ──

func Test_S01_NewValidValue(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValue", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{"val": v.Value, "valid": v.IsValid, "msg": v.Message}
		expected := args.Map{"val": "hello", "valid": true, "msg": ""}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns correct value -- basic", actual)
	})
}

func Test_S01_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueEmpty", func() {
		v := corestr.NewValidValueEmpty()
		actual := args.Map{"val": v.Value, "valid": v.IsValid}
		expected := args.Map{"val": "", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns correct value -- empty", actual)
	})
}

func Test_S01_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValidValue", func() {
		v := corestr.InvalidValidValue("bad input")
		actual := args.Map{"val": v.Value, "valid": v.IsValid, "msg": v.Message}
		expected := args.Map{"val": "", "valid": false, "msg": "bad input"}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns correct value -- with message", actual)
	})
}

func Test_S01_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValidValueNoMessage", func() {
		v := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"val": v.Value, "valid": v.IsValid, "msg": v.Message}
		expected := args.Map{"val": "", "valid": false, "msg": ""}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns correct value -- no message", actual)
	})
}

func Test_S01_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAny", func() {
		v := corestr.NewValidValueUsingAny(false, true, 42)
		actual := args.Map{"valid": v.IsValid, "notEmpty": v.Value != ""}
		expected := args.Map{"valid": true, "notEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny returns correct value -- int input", actual)
	})
}

func Test_S01_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAnyAutoValid", func() {
		v := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		actual := args.Map{"valid": v.IsValid, "notEmpty": v.Value != ""}
		expected := args.Map{"valid": false, "notEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- non-empty", actual)
	})
}

func Test_S01_NewValidValueUsingAnyAutoValid_Empty(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAnyAutoValid_Empty", func() {
		v := corestr.NewValidValueUsingAnyAutoValid(false, "")
		actual := args.Map{"valid": v.IsValid, "val": v.Value}
		expected := args.Map{"valid": true, "val": ""}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- empty input", actual)
	})
}

// ── Bool/Int/Byte/Float converters ──

func Test_S01_ValidValue_ValueBool_True(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_True", func() {
		v := corestr.NewValidValue("true")
		actual := args.Map{"bool": v.ValueBool()}
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- true string", actual)
	})
}

func Test_S01_ValidValue_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_Empty", func() {
		v := corestr.NewValidValue("")
		actual := args.Map{"bool": v.ValueBool()}
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- empty string", actual)
	})
}

func Test_S01_ValidValue_ValueBool_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_Invalid", func() {
		v := corestr.NewValidValue("notabool")
		actual := args.Map{"bool": v.ValueBool()}
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueInt", func() {
		v := corestr.NewValidValue("42")
		actual := args.Map{"int": v.ValueInt(0), "defInt": v.ValueDefInt()}
		expected := args.Map{"int": 42, "defInt": 42}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- valid int", actual)
	})
}

func Test_S01_ValidValue_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueInt_Invalid", func() {
		v := corestr.NewValidValue("abc")
		actual := args.Map{"int": v.ValueInt(99), "defInt": v.ValueDefInt()}
		expected := args.Map{"int": 99, "defInt": 0}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte", func() {
		v := corestr.NewValidValue("100")
		actual := args.Map{"byte": v.ValueByte(0), "defByte": v.ValueDefByte()}
		expected := args.Map{"byte": byte(100), "defByte": byte(100)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- valid byte", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Overflow(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Overflow", func() {
		v := corestr.NewValidValue("999")
		actual := args.Map{"byte": v.ValueByte(5)}
		expected := args.Map{"byte": byte(255)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- overflow clamped to max", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Negative(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Negative", func() {
		v := corestr.NewValidValue("-1")
		actual := args.Map{"byte": v.ValueByte(5), "defByte": v.ValueDefByte()}
		expected := args.Map{"byte": byte(0), "defByte": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- negative", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Invalid", func() {
		v := corestr.NewValidValue("abc")
		actual := args.Map{"byte": v.ValueByte(5), "defByte": v.ValueDefByte()}
		expected := args.Map{"byte": byte(0), "defByte": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueFloat64", func() {
		v := corestr.NewValidValue("3.14")
		actual := args.Map{"float": v.ValueFloat64(0), "defFloat": v.ValueDefFloat64()}
		expected := args.Map{"float": 3.14, "defFloat": 3.14}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- valid float", actual)
	})
}

func Test_S01_ValidValue_ValueFloat64_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueFloat64_Invalid", func() {
		v := corestr.NewValidValue("xyz")
		actual := args.Map{"float": v.ValueFloat64(1.5), "defFloat": v.ValueDefFloat64()}
		expected := args.Map{"float": 1.5, "defFloat": float64(0)}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- invalid string", actual)
	})
}

// ── String checks ──

func Test_S01_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsEmpty", func() {
		v := corestr.NewValidValue("")
		v2 := corestr.NewValidValue("hello")
		actual := args.Map{"emptyIsEmpty": v.IsEmpty(), "helloIsEmpty": v2.IsEmpty()}
		expected := args.Map{"emptyIsEmpty": true, "helloIsEmpty": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns correct value -- empty vs non-empty", actual)
	})
}

func Test_S01_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsWhitespace", func() {
		v := corestr.NewValidValue("   ")
		v2 := corestr.NewValidValue("hi")
		actual := args.Map{"wsIsWs": v.IsWhitespace(), "hiIsWs": v2.IsWhitespace()}
		expected := args.Map{"wsIsWs": true, "hiIsWs": false}
		expected.ShouldBeEqual(t, 0, "IsWhitespace returns correct value -- whitespace vs text", actual)
	})
}

func Test_S01_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Trim", func() {
		v := corestr.NewValidValue("  hello  ")
		actual := args.Map{"trimmed": v.Trim()}
		expected := args.Map{"trimmed": "hello"}
		expected.ShouldBeEqual(t, 0, "Trim returns correct value -- leading/trailing spaces", actual)
	})
}

func Test_S01_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasValidNonEmpty", func() {
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("")
		v3 := corestr.InvalidValidValue("x")
		actual := args.Map{"valid": v.HasValidNonEmpty(), "empty": v2.HasValidNonEmpty(), "invalid": v3.HasValidNonEmpty()}
		expected := args.Map{"valid": true, "empty": false, "invalid": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns correct value -- various cases", actual)
	})
}

func Test_S01_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasValidNonWhitespace", func() {
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("   ")
		actual := args.Map{"valid": v.HasValidNonWhitespace(), "ws": v2.HasValidNonWhitespace()}
		expected := args.Map{"valid": true, "ws": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespace returns correct value -- text vs whitespace", actual)
	})
}

func Test_S01_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasSafeNonEmpty", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{"safe": v.HasSafeNonEmpty()}
		expected := args.Map{"safe": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns correct value -- valid non-empty", actual)
	})
}

// ── Is / IsAnyOf / IsContains / IsAnyContains ──

func Test_S01_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Is", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{"yes": v.Is("hello"), "no": v.Is("world")}
		expected := args.Map{"yes": true, "no": false}
		expected.ShouldBeEqual(t, 0, "Is returns correct value -- match vs no match", actual)
	})
}

func Test_S01_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsAnyOf", func() {
		v := corestr.NewValidValue("b")
		actual := args.Map{"found": v.IsAnyOf("a", "b", "c"), "notFound": v.IsAnyOf("x", "y"), "empty": v.IsAnyOf()}
		expected := args.Map{"found": true, "notFound": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns correct value -- found, not found, empty", actual)
	})
}

func Test_S01_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{"yes": v.IsContains("world"), "no": v.IsContains("xyz")}
		expected := args.Map{"yes": true, "no": false}
		expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- substring match", actual)
	})
}

func Test_S01_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsAnyContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{"found": v.IsAnyContains("xyz", "world"), "notFound": v.IsAnyContains("abc"), "empty": v.IsAnyContains()}
		expected := args.Map{"found": true, "notFound": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "IsAnyContains returns correct value -- found, not found, empty", actual)
	})
}

func Test_S01_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsEqualNonSensitive", func() {
		v := corestr.NewValidValue("Hello")
		actual := args.Map{"yes": v.IsEqualNonSensitive("hello"), "no": v.IsEqualNonSensitive("world")}
		expected := args.Map{"yes": true, "no": false}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive returns correct value -- case insensitive", actual)
	})
}

// ── Regex ──

func Test_S01_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsRegexMatches", func() {
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"matches": v.IsRegexMatches(re), "nilRegex": v.IsRegexMatches(nil)}
		expected := args.Map{"matches": true, "nilRegex": false}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches returns correct value -- match and nil", actual)
	})
}

func Test_S01_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindString", func() {
		v := corestr.NewValidValue("abc123def")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"found": v.RegexFindString(re), "nil": v.RegexFindString(nil)}
		expected := args.Map{"found": "123", "nil": ""}
		expected.ShouldBeEqual(t, 0, "RegexFindString returns correct value -- match and nil regex", actual)
	})
}

func Test_S01_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindAllStrings", func() {
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items := v.RegexFindAllStrings(re, -1)
		nilItems := v.RegexFindAllStrings(nil, -1)
		actual := args.Map{"count": len(items), "nilCount": len(nilItems)}
		expected := args.Map{"count": 3, "nilCount": 0}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings returns correct value -- matches and nil", actual)
	})
}

func Test_S01_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindAllStringsWithFlag", func() {
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := v.RegexFindAllStringsWithFlag(re, -1)
		nilItems, nilHas := v.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{"count": len(items), "hasAny": hasAny, "nilCount": len(nilItems), "nilHas": nilHas}
		expected := args.Map{"count": 3, "hasAny": true, "nilCount": 0, "nilHas": false}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag returns correct value -- matches and nil", actual)
	})
}

// ── Split ──

func Test_S01_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Split", func() {
		v := corestr.NewValidValue("a,b,c")
		parts := v.Split(",")
		actual := args.Map{"count": len(parts), "first": parts[0], "last": parts[2]}
		expected := args.Map{"count": 3, "first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "Split returns correct value -- comma separated", actual)
	})
}

func Test_S01_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_SplitTrimNonWhitespace", func() {
		v := corestr.NewValidValue("a, ,b")
		parts := v.SplitTrimNonWhitespace(",")
		actual := args.Map{"count": len(parts)}
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "SplitTrimNonWhitespace returns correct value -- trims whitespace", actual)
	})
}

// ── ValueBytesOnce ──

func Test_S01_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBytesOnce", func() {
		v := corestr.NewValidValue("hi")
		b := v.ValueBytesOnce()
		b2 := v.ValueBytesOncePtr()
		actual := args.Map{"len": len(b), "lenPtr": len(b2), "eq": string(b) == string(b2)}
		expected := args.Map{"len": 2, "lenPtr": 2, "eq": true}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- caches bytes", actual)
	})
}

// ── Clone / Clear / Dispose / String / FullString ──

func Test_S01_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clone", func() {
		v := corestr.NewValidValue("hello")
		c := v.Clone()
		actual := args.Map{"val": c.Value, "valid": c.IsValid, "samePtr": v == c}
		expected := args.Map{"val": "hello", "valid": true, "samePtr": false}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep copy", actual)
	})
}

func Test_S01_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clone_Nil", func() {
		var v *corestr.ValidValue
		c := v.Clone()
		actual := args.Map{"isNil": c == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- nil receiver", actual)
	})
}

func Test_S01_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clear", func() {
		v := corestr.NewValidValue("hello")
		v.Clear()
		actual := args.Map{"val": v.Value, "valid": v.IsValid, "msg": v.Message}
		expected := args.Map{"val": "", "valid": false, "msg": ""}
		expected.ShouldBeEqual(t, 0, "Clear returns correct value -- resets all fields", actual)
	})
}

func Test_S01_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Dispose", func() {
		v := corestr.NewValidValue("hello")
		v.Dispose()
		actual := args.Map{"val": v.Value, "valid": v.IsValid}
		expected := args.Map{"val": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- clears", actual)
	})
}

func Test_S01_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Dispose_Nil", func() {
		var v *corestr.ValidValue
		// Should not panic
		v.Dispose()
		actual := args.Map{"isNil": v == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- nil receiver no panic", actual)
	})
}

func Test_S01_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_String", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{"str": v.String()}
		expected := args.Map{"str": "hello"}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- value", actual)
	})
}

func Test_S01_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_String_Nil", func() {
		var v *corestr.ValidValue
		actual := args.Map{"str": v.String()}
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- nil receiver", actual)
	})
}

func Test_S01_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_FullString", func() {
		v := corestr.NewValidValue("hello")
		s := v.FullString()
		actual := args.Map{"notEmpty": s != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- non-empty", actual)
	})
}

func Test_S01_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_FullString_Nil", func() {
		var v *corestr.ValidValue
		actual := args.Map{"str": v.FullString()}
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- nil receiver", actual)
	})
}

// ── Json / Serialize ──

func Test_S01_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Json", func() {
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.Json()
		actual := args.Map{"hasBytes": j.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Json returns correct value -- has bytes", actual)
	})
}

func Test_S01_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_JsonPtr", func() {
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()
		actual := args.Map{"notNil": j != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- not nil", actual)
	})
}

func Test_S01_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Serialize", func() {
		v := corestr.NewValidValue("hello")
		b, err := v.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- no error", actual)
	})
}

func Test_S01_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ParseInjectUsingJson", func() {
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()
		v2 := &corestr.ValidValue{}
		result, err := v2.ParseInjectUsingJson(j)
		actual := args.Map{"noErr": err == nil, "val": result.Value, "valid": result.IsValid}
		expected := args.Map{"noErr": true, "val": "test", "valid": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- round trip", actual)
	})
}

func Test_S01_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Deserialize", func() {
		v := corestr.NewValidValue("hello")
		var v2 corestr.ValidValue
		err := v.Deserialize(&v2)
		actual := args.Map{"noErr": err == nil, "val": v2.Value}
		expected := args.Map{"noErr": true, "val": "hello"}
		expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- round trip", actual)
	})
}
