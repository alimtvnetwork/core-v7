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

func Test_ValidValue_NewValidValue(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewValidValue", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewValidValueEmpty", func() {
		// Arrange
		vv := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns empty -- with args", actual)
	})
}

func Test_ValidValue_InvalidNoMessage_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_InvalidNoMessage", func() {
		// Arrange
		vv := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns error -- with args", actual)
	})
}

func Test_ValidValue_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_InvalidWithMessage", func() {
		// Arrange
		vv := corestr.InvalidValidValue("err")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns error -- with args", actual)
	})
}

func Test_ValidValue_NewUsingAny_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewUsingAny", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAny(false, true, 42)

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"notEmpty": vv.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_NewUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_NewUsingAnyAutoValid", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAnyAutoValid(false, 42)

		// Act
		actual := args.Map{"notEmpty": vv.Value != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Bytes, Checks, Trim
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_ValueBytesOnce_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")
		b1 := vv.ValueBytesOnce()
		b2 := vv.ValueBytesOnce() // cached

		// Act
		actual := args.Map{
			"len": len(b1),
			"same": &b1[0] == &b2[0],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"same": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- with args", actual)
	})
}

func Test_ValidValue_ValueBytesOncePtr_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueBytesOncePtr", func() {
		// Arrange
		vv := corestr.NewValidValue("xy")

		// Act
		actual := args.Map{"len": len(vv.ValueBytesOncePtr())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValueBytesOncePtr returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsEmpty_IsWhitespace_Trim(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsEmpty_IsWhitespace_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue("  hi  ")

		// Act
		actual := args.Map{
			"empty": vv.IsEmpty(),
			"ws": vv.IsWhitespace(),
			"trim": vv.Trim(),
		}

		// Assert
		expected := args.Map{
			"empty": false,
			"ws": false,
			"trim": "hi",
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty/IsWhitespace/Trim returns empty -- with args", actual)
	})
}

func Test_ValidValue_HasValidNonEmpty_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		actual := args.Map{
			"hv": vv.HasValidNonEmpty(),
			"hvw": vv.HasValidNonWhitespace(),
			"safe": vv.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"hv": true,
			"hvw": true,
			"safe": true,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Type conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_ValueBool_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueBool", func() {
		// Arrange
		vv1 := corestr.NewValidValue("true")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("")

		// Act
		actual := args.Map{
			"t": vv1.ValueBool(),
			"f": vv2.ValueBool(),
			"e": vv3.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"t": true,
			"f": false,
			"e": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueBool", actual)
	})
}

func Test_ValidValue_ValueInt_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueInt", func() {
		// Arrange
		vv1 := corestr.NewValidValue("42")
		vv2 := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"val": vv1.ValueInt(0),
			"def": vv2.ValueInt(99),
			"defInt": vv1.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 99,
			"defInt": 42,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueInt", actual)
	})
}

func Test_ValidValue_ValueByte_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueByte", func() {
		// Arrange
		vv1 := corestr.NewValidValue("100")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("300")
		vv4 := corestr.NewValidValue("-1")

		// Act
		actual := args.Map{
			"val": vv1.ValueByte(0),
			"err": vv2.ValueByte(7),
			"over": vv3.ValueByte(5),
			"neg": vv4.ValueByte(9),
		}

		// Assert
		expected := args.Map{
			"val": byte(100),
			"err": byte(0),
			"over": byte(255),
			"neg": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueByte", actual)
	})
}

func Test_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueDefByte", func() {
		// Arrange
		vv1 := corestr.NewValidValue("50")
		vv2 := corestr.NewValidValue("abc")
		vv3 := corestr.NewValidValue("999")
		vv4 := corestr.NewValidValue("-5")

		// Act
		actual := args.Map{
			"val": vv1.ValueDefByte(),
			"err": vv2.ValueDefByte(),
			"over": vv3.ValueDefByte(),
			"neg": vv4.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(50),
			"err": byte(0),
			"over": byte(255),
			"neg": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueDefByte", actual)
	})
}

func Test_ValidValue_ValueFloat64_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ValueFloat64", func() {
		// Arrange
		vv1 := corestr.NewValidValue("3.14")
		vv2 := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"close": vv1.ValueFloat64(0) > 3.1,
			"def": vv2.ValueFloat64(1.0),
			"defFloat": vv1.ValueDefFloat64() > 3.1,
		}

		// Assert
		expected := args.Map{
			"close": true,
			"def": 1.0,
			"defFloat": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueFloat64", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — String matching
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_Is_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"is": vv.Is("hello"),
			"isNot": vv.Is("world"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isNot": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Is", actual)
	})
}

func Test_ValidValue_IsAnyOf_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{
			"found": vv.IsAnyOf("a", "b"),
			"notFound": vv.IsAnyOf("x", "y"),
			"empty": vv.IsAnyOf(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsAnyOf", actual)
	})
}

func Test_ValidValue_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsContains_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"contains":    vv.IsContains("world"),
			"notContains": vv.IsContains("xyz"),
			"anyContains": vv.IsAnyContains("xyz", "hello"),
			"anyNone":     vv.IsAnyContains("xyz"),
			"anyEmpty":    vv.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"contains": true,
			"notContains": false,
			"anyContains": true,
			"anyNone": false,
			"anyEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsContains/IsAnyContains", actual)
	})
}

func Test_ValidValue_IsEqualNonSensitive_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{
			"eq": vv.IsEqualNonSensitive("hello"),
			"neq": vv.IsEqualNonSensitive("world"),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsEqualNonSensitive", actual)
	})
}

func Test_ValidValue_IsRegexMatches_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"match": vv.IsRegexMatches(re),
			"nilRegex": vv.IsRegexMatches(nil),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"nilRegex": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsRegexMatches", actual)
	})
}

func Test_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123def")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"found": vv.RegexFindString(re),
			"nilRegex": vv.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"found": "123",
			"nilRegex": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindString", actual)
	})
}

func Test_ValidValue_RegexFindAllStrings_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items := vv.RegexFindAllStrings(re, -1)

		// Act
		actual := args.Map{"len": len(items)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindAllStrings", actual)
	})
}

func Test_ValidValue_RegexFindAllStrings_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStrings_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		items := vv.RegexFindAllStrings(nil, -1)

		// Act
		actual := args.Map{"len": len(items)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- RegexFindAllStrings nil", actual)
	})
}

func Test_ValidValue_RegexFindAllStringsWithFlag_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)

		// Act
		actual := args.Map{
			"len": len(items),
			"hasAny": hasAny,
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- RegexFindAllStringsWithFlag", actual)
	})
}

func Test_ValidValue_RegexFindAllStringsWithFlag_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"len": len(items),
			"hasAny": hasAny,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- RegexFindAllStringsWithFlag nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Split
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_Split_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")

		// Act
		actual := args.Map{"len": len(vv.Split(","))}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Split", actual)
	})
}

func Test_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_SplitNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("a,,b")
		result := vv.SplitNonEmpty(",")

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- SplitNonEmpty", actual)
	})
}

func Test_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_SplitTrimNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("a , , b")
		result := vv.SplitTrimNonWhitespace(",")

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- SplitTrimNonWhitespace", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Clone, String, JSON, Serialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_Clone_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()

		// Act
		actual := args.Map{
			"val": cloned.Value,
			"notSame": cloned != vv,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"notSame": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clone", actual)
	})
}

func Test_ValidValue_Clone_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"nil": vv.Clone() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- Clone nil", actual)
	})
}

func Test_ValidValue_String_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"val": vv.String()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- String nil", actual)
	})
}

func Test_ValidValue_String_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_String", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"val": vv.String()}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- String", actual)
	})
}

func Test_ValidValue_FullString_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"notEmpty": vv.FullString() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- FullString", actual)
	})
}

func Test_ValidValue_FullString_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"val": vv.FullString()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "ValidValue returns nil -- FullString nil", actual)
	})
}

func Test_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Clear_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vv.Clear()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clear", actual)

		vv2 := corestr.NewValidValue("x")
		vv2.Dispose()
		(*corestr.ValidValue)(nil).Clear()
		(*corestr.ValidValue)(nil).Dispose()
	})
}

func Test_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Json", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		jr := vv.Json()

		// Act
		actual := args.Map{"noErr": !jr.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Json", actual)
	})
}

func Test_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_JsonPtr", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"notNil": vv.JsonPtr() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- JsonPtr", actual)
	})
}

func Test_ValidValue_Serialize_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Serialize", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		b, err := vv.Serialize()

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
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Serialize", actual)
	})
}

func Test_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_Deserialize", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		target := &corestr.ValidValue{}
		err := vv.Deserialize(target)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Deserialize", actual)
	})
}

func Test_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I27_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		vv := &corestr.ValidValue{}
		jr := corejson.New(corestr.ValidValue{Value: "test", IsValid: true})
		parsed, err := vv.ParseInjectUsingJson(&jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"notNil": parsed != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ParseInjectUsingJson", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyAnyValuePair_Basic_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Basic", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "name", Value: 42}

		// Act
		actual := args.Map{
			"key": kv.KeyName(), "varName": kv.VariableName(),
			"valAny": kv.ValueAny(), "isVarEq": kv.IsVariableNameEqual("name"),
		}

		// Assert
		expected := args.Map{
			"key": "name",
			"varName": "name",
			"valAny": 42,
			"isVarEq": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- basic", actual)
	})
}

func Test_KeyAnyValuePair_IsValueNull_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueNull", func() {
		// Arrange
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		kv2 := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		var kv3 *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{
			"null": kv1.IsValueNull(),
			"notNull": kv2.IsValueNull(),
			"nilRcv": kv3.IsValueNull(),
		}

		// Assert
		expected := args.Map{
			"null": true,
			"notNull": false,
			"nilRcv": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- IsValueNull", actual)
	})
}

func Test_KeyAnyValuePair_HasNonNull_HasValue(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_HasNonNull_HasValue", func() {
		// Arrange
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		kv2 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv3 *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{
			"has1": kv1.HasNonNull(),
			"has2": kv2.HasNonNull(),
			"hasNil": kv3.HasNonNull(),
			"hasVal": kv1.HasValue(),
			"hasValNil": kv3.HasValue(),
		}

		// Assert
		expected := args.Map{
			"has1": true,
			"has2": false,
			"hasNil": false,
			"hasVal": true,
			"hasValNil": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- HasNonNull/HasValue", actual)
	})
}

func Test_KeyAnyValuePair_IsValueEmptyString_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueEmptyString", func() {
		// Arrange
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv2 *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{
			"empty": kv1.IsValueEmptyString(),
			"nilRcv": kv2.IsValueEmptyString(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"nilRcv": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns empty -- IsValueEmptyString", actual)
	})
}

func Test_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_IsValueWhitespace", func() {
		// Arrange
		kv1 := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		var kv2 *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{
			"ws": kv1.IsValueWhitespace(),
			"nilRcv": kv2.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"ws": true,
			"nilRcv": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- IsValueWhitespace", actual)
	})
}

func Test_KeyAnyValuePair_ValueString_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_ValueString", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s1 := kv.ValueString()
		s2 := kv.ValueString() // cached path

		// Act
		actual := args.Map{
			"notEmpty": s1 != "",
			"same": s1 == s2,
		}

		// Assert
		expected := args.Map{
			"notEmpty": true,
			"same": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns non-empty -- ValueString", actual)
	})
}

func Test_KeyAnyValuePair_ValueString_Null(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_ValueString_Null", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		s := kv.ValueString()
		_ = s // covers GetOnce path

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns non-empty -- ValueString null", actual)
	})
}

func Test_KeyAnyValuePair_Compile_String(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Compile_String", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"compile": kv.Compile(),
			"str": kv.String(),
		}

		// Assert
		expected := args.Map{
			"compile": "{k:v}",
			"str": "{k:v}",
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Compile/String", actual)
	})
}

func Test_KeyAnyValuePair_SerializeMust_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- SerializeMust", actual)
	})
}

func Test_KeyAnyValuePair_Serialize_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Serialize", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

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
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Serialize", actual)
	})
}

func Test_KeyAnyValuePair_Json_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Json", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.Json()

		// Act
		actual := args.Map{"noErr": !jr.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- Json", actual)
	})
}

func Test_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_JsonPtr", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"notNil": kv.JsonPtr() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- JsonPtr", actual)
	})
}

func Test_KeyAnyValuePair_AsJsonContractsBinder_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_AsJsonContractsBinder", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"notNil": kv.AsJsonContractsBinder() != nil,
			"jsoner": kv.AsJsoner() != nil,
			"selfInj": kv.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"jsoner": true,
			"selfInj": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- As* methods", actual)
	})
}

func Test_KeyAnyValuePair_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_JsonParseSelfInject", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{}
		jr := corejson.New(corestr.KeyAnyValuePair{Key: "test", Value: "val"})
		err := kv.JsonParseSelfInject(&jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- JsonParseSelfInject", actual)
	})
}

func Test_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I27_KeyAnyValuePair_Clear_Dispose", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"valNull": kv.Value == nil,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"valNull": true,
		}
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

func Test_ValidValues_NewEmpty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewEmpty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"empty": vvs.IsEmpty(),
			"count": vvs.Count(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
			"count": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- NewEmpty", actual)
	})
}

func Test_ValidValues_NewWithCap(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewWithCap", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- NewWithCap", actual)
	})
}

func Test_ValidValues_NewUsingValues_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewUsingValues", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"hasAny": vvs.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- NewUsingValues", actual)
	})
}

func Test_ValidValues_NewUsingValues_Empty_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_NewUsingValues_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues()

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- NewUsingValues empty", actual)
	})
}

func Test_ValidValues_Add_AddFull(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Add_AddFull", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")
		vvs.AddFull(false, "c", "err")

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"lastIdx": vvs.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"lastIdx": 2,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Add/AddFull", actual)
	})
}

func Test_ValidValues_HasIndex_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_HasIndex", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{
			"has0": vvs.HasIndex(0),
			"has1": vvs.HasIndex(1),
			"has2": vvs.HasIndex(2),
		}

		// Assert
		expected := args.Map{
			"has0": true,
			"has1": true,
			"has2": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- HasIndex", actual)
	})
}

func Test_ValidValues_SafeValueAt_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("hello").Add("world")

		// Act
		actual := args.Map{
			"v0": vvs.SafeValueAt(0),
			"v1": vvs.SafeValueAt(1),
			"oob": vvs.SafeValueAt(99),
		}

		// Assert
		expected := args.Map{
			"v0": "hello",
			"v1": "world",
			"oob": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValueAt", actual)
	})
}

func Test_ValidValues_SafeValidValueAt_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("hello")
		vvs.AddFull(false, "bad", "err")

		// Act
		actual := args.Map{
			"v0": vvs.SafeValidValueAt(0),
			"v1": vvs.SafeValidValueAt(1),
			"oob": vvs.SafeValidValueAt(99),
		}

		// Assert
		expected := args.Map{
			"v0": "hello",
			"v1": "",
			"oob": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValidValueAt", actual)
	})
}

func Test_ValidValues_SafeValuesAtIndexes_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		result := vvs.SafeValuesAtIndexes(0, 2)

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
			"second": result[1],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
			"second": "c",
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValuesAtIndexes", actual)
	})
}

func Test_ValidValues_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		result := vvs.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- SafeValuesAtIndexes empty", actual)
	})
}

func Test_ValidValues_SafeValidValuesAtIndexes_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_SafeValidValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValidValuesAtIndexes", actual)
	})
}

func Test_ValidValues_Strings_FullStrings_String(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Strings_FullStrings_String", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{
			"strLen": len(vvs.Strings()),
			"fullLen": len(vvs.FullStrings()),
			"strNotEmpty": vvs.String() != "",
		}

		// Assert
		expected := args.Map{
			"strLen": 2,
			"fullLen": 2,
			"strNotEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Strings/FullStrings/String", actual)
	})
}

func Test_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Strings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{
			"strLen": len(vvs.Strings()),
			"fullLen": len(vvs.FullStrings()),
		}

		// Assert
		expected := args.Map{
			"strLen": 0,
			"fullLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Strings empty", actual)
	})
}

func Test_ValidValues_Find_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})

		// Act
		actual := args.Map{"len": len(found)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Find", actual)
	})
}

func Test_ValidValues_Find_Break(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find_Break", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, index == 0
		})

		// Act
		actual := args.Map{"len": len(found)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Find break", actual)
	})
}

func Test_ValidValues_Find_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Find_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		found := vvs.Find(func(index int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})

		// Act
		actual := args.Map{"len": len(found)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Find empty", actual)
	})
}

func Test_ValidValues_Adds_AddsPtr(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Adds_AddsPtr", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})
		vvs.AddsPtr(corestr.NewValidValue("c"))

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Adds/AddsPtr", actual)
	})
}

func Test_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Adds_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Adds()
		vvs.AddsPtr()

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Adds empty", actual)
	})
}

func Test_ValidValues_AddValidValues_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddValidValues", func() {
		// Arrange
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b").Add("c")
		vvs1.AddValidValues(vvs2)

		// Act
		actual := args.Map{"len": vvs1.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- AddValidValues", actual)
	})
}

func Test_ValidValues_AddValidValues_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddValidValues_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs.AddValidValues(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddValidValues nil", actual)
	})
}

func Test_ValidValues_ConcatNew_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew", func() {
		// Arrange
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- ConcatNew", actual)
	})
}

func Test_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew_EmptyClone", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(true)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty clone", actual)
	})
}

func Test_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_ConcatNew_EmptyNoClone", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		result := vvs.ConcatNew(false)

		// Act
		actual := args.Map{"same": result == vvs}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- ConcatNew empty no clone", actual)
	})
}

func Test_ValidValues_AddHashsetMap_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- AddHashsetMap", actual)
	})
}

func Test_ValidValues_AddHashsetMap_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashsetMap_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddHashsetMap nil", actual)
	})
}

func Test_ValidValues_AddHashset_Nil_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_AddHashset_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.AddHashset(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- AddHashset nil", actual)
	})
}

func Test_ValidValues_Hashmap_Map_FromValidValueNewValidVa(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Hashmap_Map", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.AddFull(true, "key", "val")
		hm := vvs.Hashmap()
		m := vvs.Map()

		// Act
		actual := args.Map{
			"hmNotNil": hm != nil,
			"mapLen": len(m),
		}

		// Assert
		expected := args.Map{
			"hmNotNil": true,
			"mapLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Hashmap/Map", actual)
	})
}

func Test_ValidValues_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Hashmap_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		hm := vvs.Hashmap()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns empty -- Hashmap empty", actual)
	})
}

func Test_ValidValues_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I27_ValidValues_Length_Nil", func() {
		// Arrange
		var vvs *corestr.ValidValues

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- Length nil", actual)
	})
}
