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
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns correct value -- basic", actual)
	})
}

func Test_S01_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueEmpty", func() {
		// Arrange
		v := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns correct value -- empty", actual)
	})
}

func Test_S01_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValidValue", func() {
		// Arrange
		v := corestr.InvalidValidValue("bad input")

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "bad input",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns correct value -- with message", actual)
	})
}

func Test_S01_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValidValueNoMessage", func() {
		// Arrange
		v := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns correct value -- no message", actual)
	})
}

func Test_S01_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAny", func() {
		// Arrange
		v := corestr.NewValidValueUsingAny(false, true, 42)

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"notEmpty": v.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny returns correct value -- int input", actual)
	})
}

func Test_S01_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAnyAutoValid", func() {
		// Arrange
		v := corestr.NewValidValueUsingAnyAutoValid(false, "hello")

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"notEmpty": v.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- non-empty", actual)
	})
}

func Test_S01_NewValidValueUsingAnyAutoValid_Empty(t *testing.T) {
	safeTest(t, "Test_S01_NewValidValueUsingAnyAutoValid_Empty", func() {
		// Arrange
		v := corestr.NewValidValueUsingAnyAutoValid(false, "")

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"val": v.Value,
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- empty input", actual)
	})
}

// ── Bool/Int/Byte/Float converters ──

func Test_S01_ValidValue_ValueBool_True(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_True", func() {
		// Arrange
		v := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- true string", actual)
	})
}

func Test_S01_ValidValue_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_Empty", func() {
		// Arrange
		v := corestr.NewValidValue("")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- empty string", actual)
	})
}

func Test_S01_ValidValue_ValueBool_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBool_Invalid", func() {
		// Arrange
		v := corestr.NewValidValue("notabool")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueInt", func() {
		// Arrange
		v := corestr.NewValidValue("42")

		// Act
		actual := args.Map{
			"int": v.ValueInt(0),
			"defInt": v.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"int": 42,
			"defInt": 42,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- valid int", actual)
	})
}

func Test_S01_ValidValue_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueInt_Invalid", func() {
		// Arrange
		v := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"int": v.ValueInt(99),
			"defInt": v.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"int": 99,
			"defInt": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte", func() {
		// Arrange
		v := corestr.NewValidValue("100")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(0),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(100),
			"defByte": byte(100),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- valid byte", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Overflow(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Overflow", func() {
		// Arrange
		v := corestr.NewValidValue("999")

		// Act
		actual := args.Map{"byte": v.ValueByte(5)}

		// Assert
		expected := args.Map{"byte": byte(255)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- overflow clamped to max", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Negative(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Negative", func() {
		// Arrange
		v := corestr.NewValidValue("-1")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(5),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(0),
			"defByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- negative", actual)
	})
}

func Test_S01_ValidValue_ValueByte_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueByte_Invalid", func() {
		// Arrange
		v := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(5),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(0),
			"defByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- invalid string", actual)
	})
}

func Test_S01_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueFloat64", func() {
		// Arrange
		v := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{
			"float": v.ValueFloat64(0),
			"defFloat": v.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"float": 3.14,
			"defFloat": 3.14,
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- valid float", actual)
	})
}

func Test_S01_ValidValue_ValueFloat64_Invalid(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueFloat64_Invalid", func() {
		// Arrange
		v := corestr.NewValidValue("xyz")

		// Act
		actual := args.Map{
			"float": v.ValueFloat64(1.5),
			"defFloat": v.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"float": 1.5,
			"defFloat": float64(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- invalid string", actual)
	})
}

// ── String checks ──

func Test_S01_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("")
		v2 := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"emptyIsEmpty": v.IsEmpty(),
			"helloIsEmpty": v2.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"emptyIsEmpty": true,
			"helloIsEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns correct value -- empty vs non-empty", actual)
	})
}

func Test_S01_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("   ")
		v2 := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{
			"wsIsWs": v.IsWhitespace(),
			"hiIsWs": v2.IsWhitespace(),
		}

		// Assert
		expected := args.Map{
			"wsIsWs": true,
			"hiIsWs": false,
		}
		expected.ShouldBeEqual(t, 0, "IsWhitespace returns correct value -- whitespace vs text", actual)
	})
}

func Test_S01_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Trim", func() {
		// Arrange
		v := corestr.NewValidValue("  hello  ")

		// Act
		actual := args.Map{"trimmed": v.Trim()}

		// Assert
		expected := args.Map{"trimmed": "hello"}
		expected.ShouldBeEqual(t, 0, "Trim returns correct value -- leading/trailing spaces", actual)
	})
}

func Test_S01_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("")
		v3 := corestr.InvalidValidValue("x")

		// Act
		actual := args.Map{
			"valid": v.HasValidNonEmpty(),
			"empty": v2.HasValidNonEmpty(),
			"invalid": v3.HasValidNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"empty": false,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns correct value -- various cases", actual)
	})
}

func Test_S01_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("   ")

		// Act
		actual := args.Map{
			"valid": v.HasValidNonWhitespace(),
			"ws": v2.HasValidNonWhitespace(),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"ws": false,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespace returns correct value -- text vs whitespace", actual)
	})
}

func Test_S01_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"safe": v.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"safe": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns correct value -- valid non-empty", actual)
	})
}

// ── Is / IsAnyOf / IsContains / IsAnyContains ──

func Test_S01_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Is", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"yes": v.Is("hello"),
			"no": v.Is("world"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "Is returns correct value -- match vs no match", actual)
	})
}

func Test_S01_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsAnyOf", func() {
		// Arrange
		v := corestr.NewValidValue("b")

		// Act
		actual := args.Map{
			"found": v.IsAnyOf("a", "b", "c"),
			"notFound": v.IsAnyOf("x", "y"),
			"empty": v.IsAnyOf(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns correct value -- found, not found, empty", actual)
	})
}

func Test_S01_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"yes": v.IsContains("world"),
			"no": v.IsContains("xyz"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- substring match", actual)
	})
}

func Test_S01_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsAnyContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"found": v.IsAnyContains("xyz", "world"),
			"notFound": v.IsAnyContains("abc"),
			"empty": v.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyContains returns correct value -- found, not found, empty", actual)
	})
}

func Test_S01_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		v := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{
			"yes": v.IsEqualNonSensitive("hello"),
			"no": v.IsEqualNonSensitive("world"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive returns correct value -- case insensitive", actual)
	})
}

// ── Regex ──

func Test_S01_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_IsRegexMatches", func() {
		// Arrange
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"matches": v.IsRegexMatches(re),
			"nilRegex": v.IsRegexMatches(nil),
		}

		// Assert
		expected := args.Map{
			"matches": true,
			"nilRegex": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches returns correct value -- match and nil", actual)
	})
}

func Test_S01_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindString", func() {
		// Arrange
		v := corestr.NewValidValue("abc123def")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"found": v.RegexFindString(re),
			"nil": v.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"found": "123",
			"nil": "",
		}
		expected.ShouldBeEqual(t, 0, "RegexFindString returns correct value -- match and nil regex", actual)
	})
}

func Test_S01_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items := v.RegexFindAllStrings(re, -1)
		nilItems := v.RegexFindAllStrings(nil, -1)

		// Act
		actual := args.Map{
			"count": len(items),
			"nilCount": len(nilItems),
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"nilCount": 0,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings returns correct value -- matches and nil", actual)
	})
}

func Test_S01_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := v.RegexFindAllStringsWithFlag(re, -1)
		nilItems, nilHas := v.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"count": len(items),
			"hasAny": hasAny,
			"nilCount": len(nilItems),
			"nilHas": nilHas,
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"hasAny": true,
			"nilCount": 0,
			"nilHas": false,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag returns correct value -- matches and nil", actual)
	})
}

// ── Split ──

func Test_S01_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Split", func() {
		// Arrange
		v := corestr.NewValidValue("a,b,c")
		parts := v.Split(",")

		// Act
		actual := args.Map{
			"count": len(parts),
			"first": parts[0],
			"last": parts[2],
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Split returns correct value -- comma separated", actual)
	})
}

func Test_S01_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_SplitTrimNonWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("a, ,b")
		parts := v.SplitTrimNonWhitespace(",")

		// Act
		actual := args.Map{"count": len(parts)}

		// Assert
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "SplitTrimNonWhitespace returns correct value -- trims whitespace", actual)
	})
}

// ── ValueBytesOnce ──

func Test_S01_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ValueBytesOnce", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		b := v.ValueBytesOnce()
		b2 := v.ValueBytesOncePtr()

		// Act
		actual := args.Map{
			"len": len(b),
			"lenPtr": len(b2),
			"eq": string(b) == string(b2),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"lenPtr": 2,
			"eq": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- caches bytes", actual)
	})
}

// ── Clone / Clear / Dispose / String / FullString ──

func Test_S01_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clone", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		c := v.Clone()

		// Act
		actual := args.Map{
			"val": c.Value,
			"valid": c.IsValid,
			"samePtr": v == c,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"samePtr": false,
		}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep copy", actual)
	})
}

func Test_S01_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clone_Nil", func() {
		// Arrange
		var v *corestr.ValidValue
		c := v.Clone()

		// Act
		actual := args.Map{"isNil": c == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- nil receiver", actual)
	})
}

func Test_S01_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Clear", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v.Clear()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear returns correct value -- resets all fields", actual)
	})
}

func Test_S01_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Dispose", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v.Dispose()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- clears", actual)
	})
}

func Test_S01_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Dispose_Nil", func() {
		// Arrange
		var v *corestr.ValidValue
		// Should not panic
		v.Dispose()

		// Act
		actual := args.Map{"isNil": v == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- nil receiver no panic", actual)
	})
}

func Test_S01_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_String", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"str": v.String()}

		// Assert
		expected := args.Map{"str": "hello"}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- value", actual)
	})
}

func Test_S01_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_String_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"str": v.String()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- nil receiver", actual)
	})
}

func Test_S01_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_FullString", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		s := v.FullString()

		// Act
		actual := args.Map{"notEmpty": s != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- non-empty", actual)
	})
}

func Test_S01_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_FullString_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"str": v.FullString()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- nil receiver", actual)
	})
}

// ── Json / Serialize ──

func Test_S01_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Json", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.Json()

		// Act
		actual := args.Map{"hasBytes": j.HasBytes()}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Json returns correct value -- has bytes", actual)
	})
}

func Test_S01_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_JsonPtr", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()

		// Act
		actual := args.Map{"notNil": j != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- not nil", actual)
	})
}

func Test_S01_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Serialize", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		b, err := v.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- no error", actual)
	})
}

func Test_S01_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()
		v2 := &corestr.ValidValue{}
		result, err := v2.ParseInjectUsingJson(j)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": result.Value,
			"valid": result.IsValid,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "test",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- round trip", actual)
	})
}

func Test_S01_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_S01_ValidValue_Deserialize", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		var v2 corestr.ValidValue
		err := v.Deserialize(&v2)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": v2.Value,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- round trip", actual)
	})
}
