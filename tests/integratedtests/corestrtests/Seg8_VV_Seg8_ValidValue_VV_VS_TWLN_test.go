package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Segment 8b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg8_VV_NewValidValue(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_NewValidValue", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue -- valid", actual)
	})
}

func Test_Seg8_VV_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_NewValidValueEmpty", func() {
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
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty -- empty valid", actual)
	})
}

func Test_Seg8_VV_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_InvalidValidValue", func() {
		// Arrange
		vv := corestr.InvalidValidValue("err msg")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err msg",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue -- invalid with message", actual)
	})
}

func Test_Seg8_VV_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_InvalidValidValueNoMessage", func() {
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
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage -- invalid no msg", actual)
	})
}

func Test_Seg8_VV_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_NewValidValueUsingAny", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAny(false, true, "test")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"nonEmpty": vv.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny -- valid", actual)
	})
}

func Test_Seg8_VV_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_NewValidValueUsingAnyAutoValid", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")

		// Act
		actual := args.Map{"nonEmpty": vv.Value != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid -- non-empty", actual)
	})
}

func Test_Seg8_VV_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")
		b1 := vv.ValueBytesOnce()
		b2 := vv.ValueBytesOnce()

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
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce -- cached", actual)
	})
}

func Test_Seg8_VV_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueBytesOncePtr", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{"len": len(vv.ValueBytesOncePtr())}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ValueBytesOncePtr -- delegates", actual)
	})
}

func Test_Seg8_VV_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsEmpty_IsWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{
			"empty": vv.IsEmpty(),
			"ws": vv.IsWhitespace(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"ws": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty/IsWhitespace -- empty", actual)
	})
}

func Test_Seg8_VV_Trim(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue("  hi  ")

		// Act
		actual := args.Map{"val": vv.Trim()}

		// Assert
		expected := args.Map{"val": "hi"}
		expected.ShouldBeEqual(t, 0, "Trim -- trimmed", actual)
	})
}

func Test_Seg8_VV_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_HasValidNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("val")

		// Act
		actual := args.Map{
			"nonEmpty": vv.HasValidNonEmpty(),
			"nonWS": vv.HasValidNonWhitespace(),
			"safe": vv.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"nonEmpty": true,
			"nonWS": true,
			"safe": true,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty -- true", actual)
	})
}

func Test_Seg8_VV_ValueBool(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueBool", func() {
		// Arrange
		vv1 := corestr.NewValidValue("true")
		vv2 := corestr.NewValidValue("")
		vv3 := corestr.NewValidValue("xyz")

		// Act
		actual := args.Map{
			"true": vv1.ValueBool(),
			"empty": vv2.ValueBool(),
			"invalid": vv3.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"true": true,
			"empty": false,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "ValueBool -- various", actual)
	})
}

func Test_Seg8_VV_ValueInt(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{
			"val": vv.ValueInt(0),
			"def": vv.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 42,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt -- 42", actual)
	})
}

func Test_Seg8_VV_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueInt_Invalid", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"val": vv.ValueInt(99),
			"def": vv.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 99,
			"def": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt invalid -- defaults", actual)
	})
}

func Test_Seg8_VV_ValueByte(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueByte", func() {
		// Arrange
		vv := corestr.NewValidValue("100")

		// Act
		actual := args.Map{
			"val": vv.ValueByte(0),
			"def": vv.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(100),
			"def": byte(100),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte -- 100", actual)
	})
}

func Test_Seg8_VV_ValueByte_Negative(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueByte_Negative", func() {
		// Arrange
		vv := corestr.NewValidValue("-1")

		// Act
		actual := args.Map{"val": vv.ValueByte(5)}

		// Assert
		expected := args.Map{"val": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValueByte negative -- 0", actual)
	})
}

func Test_Seg8_VV_ValueByte_OverMax(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueByte_OverMax", func() {
		// Arrange
		vv := corestr.NewValidValue("300")

		// Act
		actual := args.Map{"val": vv.ValueByte(5)}

		// Assert
		expected := args.Map{"val": byte(255)} // MaxUnit8
		expected.ShouldBeEqual(t, 0, "ValueByte over max -- clamped to 255", actual)
	})
}

func Test_Seg8_VV_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{
			"val": vv.ValueFloat64(0.0),
			"def": vv.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"val": 3.14,
			"def": 3.14,
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 -- 3.14", actual)
	})
}

func Test_Seg8_VV_ValueFloat64_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ValueFloat64_Invalid", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{"val": vv.ValueFloat64(1.5)}

		// Assert
		expected := args.Map{"val": 1.5}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 invalid -- default", actual)
	})
}

func Test_Seg8_VV_Is(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"is": vv.Is("hello"),
			"not": vv.Is("world"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"not": false,
		}
		expected.ShouldBeEqual(t, 0, "Is -- correct", actual)
	})
}

func Test_Seg8_VV_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{
			"found": vv.IsAnyOf("a", "b"),
			"not": vv.IsAnyOf("x"),
			"empty": vv.IsAnyOf(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"not": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyOf -- correct", actual)
	})
}

func Test_Seg8_VV_IsContains(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"yes": vv.IsContains("world"),
			"no": vv.IsContains("xyz"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains -- correct", actual)
	})
}

func Test_Seg8_VV_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"found": vv.IsAnyContains("xyz", "world"),
			"not": vv.IsAnyContains("abc"),
			"empty": vv.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"not": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyContains -- correct", actual)
	})
}

func Test_Seg8_VV_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"eq": vv.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive -- true", actual)
	})
}

func Test_Seg8_VV_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"match": vv.IsRegexMatches(re),
			"nil": vv.IsRegexMatches(nil),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches -- correct", actual)
	})
}

func Test_Seg8_VV_RegexFindString(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"found": vv.RegexFindString(re),
			"nil": vv.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"found": "123",
			"nil": "",
		}
		expected.ShouldBeEqual(t, 0, "RegexFindString -- correct", actual)
	})
}

func Test_Seg8_VV_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_RegexFindAllStrings", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)

		// Act
		actual := args.Map{
			"len": len(vv.RegexFindAllStrings(re, -1)),
			"nil": len(vv.RegexFindAllStrings(nil, -1)),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"nil": 0,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings -- correct", actual)
	})
}

func Test_Seg8_VV_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_RegexFindAllStringsWithFlag", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, has := vv.RegexFindAllStringsWithFlag(re, -1)
		nilItems, nilHas := vv.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"len": len(items),
			"has": has,
			"nilLen": len(nilItems),
			"nilHas": nilHas,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"has": true,
			"nilLen": 0,
			"nilHas": false,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag -- correct", actual)
	})
}

func Test_Seg8_VV_Split(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b")

		// Act
		actual := args.Map{"len": len(vv.Split(","))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Split -- 2", actual)
	})
}

func Test_Seg8_VV_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_SplitNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("a,,b")
		result := vv.SplitNonEmpty(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 3} // known: returns slice not nonEmptySlice
		expected.ShouldBeEqual(t, 0, "SplitNonEmpty -- returns original (known behavior)", actual)
	})
}

func Test_Seg8_VV_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_SplitTrimNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("a, ,b")
		result := vv.SplitTrimNonWhitespace(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 3} // known: returns slice not nonEmptySlice
		expected.ShouldBeEqual(t, 0, "SplitTrimNonWhitespace -- returns original (known behavior)", actual)
	})
}

func Test_Seg8_VV_Clone(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		c := vv.Clone()

		// Act
		actual := args.Map{
			"val": c.Value,
			"valid": c.IsValid,
			"diff": c != vv,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- copy", actual)
	})
}

func Test_Seg8_VV_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue
		c := vv.Clone()

		// Act
		actual := args.Map{"nil": c == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clone nil -- nil", actual)
	})
}

func Test_Seg8_VV_String(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_String", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"str": vv.String()}

		// Assert
		expected := args.Map{"str": "hello"}
		expected.ShouldBeEqual(t, 0, "String -- value", actual)
	})
}

func Test_Seg8_VV_String_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"str": vv.String()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "String nil -- empty", actual)
	})
}

func Test_Seg8_VV_FullString(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"nonEmpty": vv.FullString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "FullString -- non-empty", actual)
	})
}

func Test_Seg8_VV_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"str": vv.FullString()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "FullString nil -- empty", actual)
	})
}

func Test_Seg8_VV_Clear(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Clear", func() {
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
		expected.ShouldBeEqual(t, 0, "Clear -- reset", actual)
	})
}

func Test_Seg8_VV_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Clear_Nil", func() {
		var vv *corestr.ValidValue
		vv.Clear() // should not panic
	})
}

func Test_Seg8_VV_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vv.Dispose()

		// Act
		actual := args.Map{"val": vv.Value}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleared", actual)
	})
}

func Test_Seg8_VV_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Dispose_Nil", func() {
		var vv *corestr.ValidValue
		vv.Dispose() // should not panic
	})
}

func Test_Seg8_VV_Json(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Json", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		j := vv.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg8_VV_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_ParseInjectUsingJson", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		jr := vv.JsonPtr()
		vv2 := &corestr.ValidValue{}
		result, err := vv2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"notNil": result != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg8_VV_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Serialize", func() {
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
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg8_VV_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg8_VV_Deserialize", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		var target corestr.ValidValue
		err := vv.Deserialize(&target)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": target.Value,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues — Segment 8c
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg8_VVS_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{
			"empty": vvs.IsEmpty(),
			"len": vvs.Length(),
			"count": vvs.Count(),
			"hasAny": vvs.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
			"count": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "EmptyValidValues -- empty", actual)
	})
}

func Test_Seg8_VVS_NewValidValues(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_NewValidValues", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)

		// Act
		actual := args.Map{"empty": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValues -- empty with capacity", actual)
	})
}

func Test_Seg8_VVS_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_NewValidValuesUsingValues", func() {
		// Arrange
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		v2 := corestr.ValidValue{Value: "b", IsValid: true}
		vvs := corestr.NewValidValuesUsingValues(v1, v2)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NewValidValuesUsingValues -- 2 items", actual)
	})
}

func Test_Seg8_VVS_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_NewValidValuesUsingValues_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues()

		// Act
		actual := args.Map{"empty": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "NewValidValuesUsingValues empty -- empty", actual)
	})
}

func Test_Seg8_VVS_Add(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Add", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"lastIdx": vvs.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 items", actual)
	})
}

func Test_Seg8_VVS_AddFull(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddFull", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddFull(false, "val", "msg")

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFull -- 1 item", actual)
	})
}

func Test_Seg8_VVS_Adds(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Adds", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		v2 := corestr.ValidValue{Value: "b", IsValid: true}
		vvs.Adds(v1, v2)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2", actual)
	})
}

func Test_Seg8_VVS_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Adds_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Adds()

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_Seg8_VVS_AddsPtr(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddsPtr", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		v := corestr.NewValidValue("a")
		vvs.AddsPtr(v)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsPtr -- 1", actual)
	})
}

func Test_Seg8_VVS_AddsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddsPtr_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddsPtr()

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsPtr empty -- no change", actual)
	})
}

func Test_Seg8_VVS_HasIndex(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_HasIndex", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
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
		expected.ShouldBeEqual(t, 0, "HasIndex -- correct", actual)
	})
}

func Test_Seg8_VVS_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{
			"at0": vvs.SafeValueAt(0),
			"at1": vvs.SafeValueAt(1),
			"out": vvs.SafeValueAt(5),
		}

		// Assert
		expected := args.Map{
			"at0": "a",
			"at1": "b",
			"out": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValueAt -- correct", actual)
	})
}

func Test_Seg8_VVS_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a").AddFull(false, "b", "")

		// Act
		actual := args.Map{
			"at0": vvs.SafeValidValueAt(0),
			"at1": vvs.SafeValidValueAt(1),
		}

		// Assert
		expected := args.Map{
			"at0": "a",
			"at1": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValidValueAt -- only valid returned", actual)
	})
}

func Test_Seg8_VVS_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
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
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes -- correct", actual)
	})
}

func Test_Seg8_VVS_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a")
		result := vvs.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes empty -- 0", actual)
	})
}

func Test_Seg8_VVS_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValidValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a").AddFull(false, "b", "")
		result := vvs.SafeValidValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{
			"first": result[0],
			"second": result[1],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"second": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValidValuesAtIndexes -- correct", actual)
	})
}

func Test_Seg8_VVS_SafeValidValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_SafeValidValuesAtIndexes_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		result := vvs.SafeValidValuesAtIndexes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeValidValuesAtIndexes empty -- 0", actual)
	})
}

func Test_Seg8_VVS_Find(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Find", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a").Add("b").Add("c")
		result := vvs.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "b", false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find -- found b", actual)
	})
}

func Test_Seg8_VVS_Find_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Find_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		result := vvs.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Find empty -- 0", actual)
	})
}

func Test_Seg8_VVS_Find_Break(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Find_Break", func() {
		// Arrange
		vvs := corestr.NewValidValues(3)
		vvs.Add("a").Add("b").Add("c")
		result := vvs.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, true
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find break -- 1", actual)
	})
}

func Test_Seg8_VVS_Strings(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Strings", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{"len": len(vvs.Strings())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Strings -- 2", actual)
	})
}

func Test_Seg8_VVS_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Strings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"len": len(vvs.Strings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Strings empty -- 0", actual)
	})
}

func Test_Seg8_VVS_FullStrings(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_FullStrings", func() {
		// Arrange
		vvs := corestr.NewValidValues(1)
		vvs.Add("a")

		// Act
		actual := args.Map{"len": len(vvs.FullStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FullStrings -- 1", actual)
	})
}

func Test_Seg8_VVS_FullStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_FullStrings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"len": len(vvs.FullStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FullStrings empty -- 0", actual)
	})
}

func Test_Seg8_VVS_String(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_String", func() {
		// Arrange
		vvs := corestr.NewValidValues(1)
		vvs.Add("a")

		// Act
		actual := args.Map{"nonEmpty": vvs.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg8_VVS_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Length_Nil", func() {
		// Arrange
		var vvs *corestr.ValidValues

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Seg8_VVS_AddValidValues(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddValidValues", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(2)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(2)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)

		// Act
		actual := args.Map{"len": vvs1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddValidValues -- merged", actual)
	})
}

func Test_Seg8_VVS_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddValidValues_Nil", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a")
		vvs.AddValidValues(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddValidValues nil -- no change", actual)
	})
}

func Test_Seg8_VVS_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_ConcatNew", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(2)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(2)
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- combined", actual)
	})
}

func Test_Seg8_VVS_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_ConcatNew_EmptyClone", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a")
		result := vvs.ConcatNew(true)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty clone -- cloned", actual)
	})
}

func Test_Seg8_VVS_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_ConcatNew_EmptyNoClone", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.Add("a")
		result := vvs.ConcatNew(false)

		// Act
		actual := args.Map{"same": result == vvs}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty no clone -- returns self", actual)
	})
}

func Test_Seg8_VVS_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap -- 2", actual)
	})
}

func Test_Seg8_VVS_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddHashsetMap_Nil", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddHashsetMap(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap nil -- no change", actual)
	})
}

func Test_Seg8_VVS_AddHashset(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddHashset", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		vvs.AddHashset(hs)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashset -- 1", actual)
	})
}

func Test_Seg8_VVS_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_AddHashset_Nil", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddHashset(nil)

		// Act
		actual := args.Map{"len": vvs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashset nil -- no change", actual)
	})
}

func Test_Seg8_VVS_Hashmap(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Hashmap", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddFull(true, "key", "value")
		hm := vvs.Hashmap()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap -- 1", actual)
	})
}

func Test_Seg8_VVS_Map(t *testing.T) {
	safeTest(t, "Test_Seg8_VVS_Map", func() {
		// Arrange
		vvs := corestr.NewValidValues(2)
		vvs.AddFull(true, "key", "value")
		m := vvs.Map()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Map -- 1", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — Segment 8d
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg8_VS_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_Seg8_VS_InvalidValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"msg": vs.ValueValid.Message,
			"idx": vs.Index,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
			"idx": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus -- invalid with message", actual)
	})
}

func Test_Seg8_VS_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg8_VS_InvalidValueStatusNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"idx": vs.Index,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"idx": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage -- invalid", actual)
	})
}

func Test_Seg8_VS_Clone(t *testing.T) {
	safeTest(t, "Test_Seg8_VS_Clone", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")
		c := vs.Clone()

		// Act
		actual := args.Map{
			"msg": c.ValueValid.Message,
			"idx": c.Index,
			"diff": c != vs,
		}

		// Assert
		expected := args.Map{
			"msg": "err",
			"idx": -1,
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- copy", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — Segment 8e
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg8_TWLN_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_HasLineNumber", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "HasLineNumber -- valid", actual)
	})
}

func Test_Seg8_TWLN_InvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_InvalidLineNumber", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "InvalidLineNumber -- invalid", actual)
	})
}

func Test_Seg8_TWLN_Length(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_Length", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}

		// Act
		actual := args.Map{"len": twln.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Length -- 3", actual)
	})
}

func Test_Seg8_TWLN_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_Length_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": twln.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Seg8_TWLN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmpty", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- false", actual)
	})
}

func Test_Seg8_TWLN_IsEmpty_True(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmpty_True", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true (no text + invalid line)", actual)
	})
}

func Test_Seg8_TWLN_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmpty_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty nil -- true", actual)
	})
}

func Test_Seg8_TWLN_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmptyText", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{"emptyText": twln.IsEmptyText()}

		// Assert
		expected := args.Map{"emptyText": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText -- true", actual)
	})
}

func Test_Seg8_TWLN_IsEmptyText_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmptyText_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"emptyText": twln.IsEmptyText()}

		// Assert
		expected := args.Map{"emptyText": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText nil -- true", actual)
	})
}

func Test_Seg8_TWLN_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_IsEmptyTextLineBoth", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"both": twln.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"both": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth -- delegates to IsEmpty", actual)
	})
}

func Test_Seg8_TWLN_HasLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_TWLN_HasLineNumber_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "HasLineNumber nil -- false/true", actual)
	})
}
