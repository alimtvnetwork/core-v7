package corestrtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil ptr", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_NilSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_NilSlice", func() {
		// Arrange
		var s [][]string

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil inner slice", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_WithItems", func() {
		// Arrange
		s := [][]string{{"a", "b"}, {"c"}}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns total -- multi", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_Empty", func() {
		// Arrange
		s := [][]string{}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- empty outer", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- no args", actual)
	})
}

func Test_Seg1_AllIndividualsLengthOfSimpleSlices_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualsLengthOfSimpleSlices_WithItems", func() {
		// Arrange
		s1 := corestr.SimpleSlice{"a", "b"}
		s2 := corestr.SimpleSlice{"c"}

		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(&s1, &s2)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns total -- multi slices", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_Empty", func() {
		// Act
		actual := args.Map{"val": corestr.AnyToString(false, "")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty input", actual)
	})
}

func Test_Seg1_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_WithFieldName", func() {
		// Arrange
		result := corestr.AnyToString(true, 42)

		// Act
		actual := args.Map{"nonEmpty": result != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field name", actual)
	})
}

func Test_Seg1_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_WithoutFieldName", func() {
		// Arrange
		result := corestr.AnyToString(false, 42)

		// Act
		actual := args.Map{"nonEmpty": result != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- without field name", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_Empty", func() {
		// Arrange
		result := corestr.CloneSlice([]string{})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty input", actual)
	})
}

func Test_Seg1_CloneSlice_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_WithItems", func() {
		// Arrange
		src := []string{"a", "b", "c"}
		result := corestr.CloneSlice(src)

		// Act
		actual := args.Map{
			"len": len(result),
			"eq": result[0] == "a" && result[2] == "c",
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"eq": true,
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns copy -- with items", actual)
	})
}

func Test_Seg1_CloneSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_Nil", func() {
		// Arrange
		result := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- nil input", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_Clone", func() {
		// Arrange
		result := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns clone -- isClone true", actual)
	})
}

func Test_Seg1_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_NoClone", func() {
		// Arrange
		result := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns same -- isClone false", actual)
	})
}

func Test_Seg1_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_Empty", func() {
		// Arrange
		result := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- no items", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_HasLineNumber", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber HasLineNumber true -- valid line", actual)
	})
}

func Test_Seg1_TextWithLineNumber_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_Invalid", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsInvalidLineNumber -- invalid line", actual)
	})
}

func Test_Seg1_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_Length", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 5}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- 5 chars", actual)
	})
}

func Test_Seg1_TextWithLineNumber_NilLength(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_NilLength", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- nil", actual)
	})
}

func Test_Seg1_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_IsEmpty", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{"empty": tw.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- invalid line number", actual)
	})
}

func Test_Seg1_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_IsEmptyText", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{
			"emptyText": tw.IsEmptyText(),
			"emptyBoth": tw.IsEmptyTextLineBoth(),
		}

		// Assert
		expected := args.Map{
			"emptyText": true,
			"emptyBoth": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmptyText -- empty text", actual)
	})
}

func Test_Seg1_TextWithLineNumber_NilEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_NilEmpty", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"empty": tw.IsEmpty(),
			"emptyText": tw.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- nil receiver", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_InvalidNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns valid struct -- default", actual)
	})
}

func Test_Seg1_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_Invalid", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err msg")

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns valid struct -- with message", actual)
	})
}

func Test_Seg1_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_Clone", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("clone me")
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"notNil": cloned != nil,
			"sameIdx": cloned.Index == vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"sameIdx": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus Clone returns copy -- same index", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_KVP_BasicAccessors(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_BasicAccessors", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "name", Value: "alice"}

		// Act
		actual := args.Map{
			"key":      kvp.KeyName(),
			"varName":  kvp.VariableName(),
			"valStr":   kvp.ValueString(),
			"isVarEq":  kvp.IsVariableNameEqual("name"),
			"isValEq":  kvp.IsValueEqual("alice"),
			"compile":  kvp.Compile(),
			"hasKey":   kvp.HasKey(),
			"hasValue": kvp.HasValue(),
		}

		// Assert
		expected := args.Map{
			"key":      "name",
			"varName":  "name",
			"valStr":   "alice",
			"isVarEq":  true,
			"isValEq":  true,
			"compile":  kvp.String(),
			"hasKey":   true,
			"hasValue": true,
		}
		expected.ShouldBeEqual(t, 0, "KVP basic accessors -- happy path", actual)
	})
}

func Test_Seg1_KVP_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_EmptyChecks", func() {
		// Arrange
		kvp := corestr.KeyValuePair{}

		// Act
		actual := args.Map{
			"keyEmpty":    kvp.IsKeyEmpty(),
			"valEmpty":    kvp.IsValueEmpty(),
			"kvEmpty":     kvp.IsKeyValueEmpty(),
			"kvAnyEmpty":  kvp.IsKeyValueAnyEmpty(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty":    true,
			"valEmpty":    true,
			"kvEmpty":     true,
			"kvAnyEmpty":  true,
		}
		expected.ShouldBeEqual(t, 0, "KVP empty checks -- all empty", actual)
	})
}

func Test_Seg1_KVP_TrimAndConversions(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_TrimAndConversions", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: " name ", Value: " 42 "}

		// Act
		actual := args.Map{
			"trimKey": kvp.TrimKey(),
			"trimVal": kvp.TrimValue(),
		}

		// Assert
		expected := args.Map{
			"trimKey": "name",
			"trimVal": "42",
		}
		expected.ShouldBeEqual(t, 0, "KVP Trim -- whitespace removed", actual)
	})
}

func Test_Seg1_KVP_ValueBool(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueBool", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "flag", Value: "true"}
		kvpEmpty := corestr.KeyValuePair{Key: "flag", Value: ""}
		kvpBad := corestr.KeyValuePair{Key: "flag", Value: "notabool"}

		// Act
		actual := args.Map{
			"t": kvp.ValueBool(),
			"empty": kvpEmpty.ValueBool(),
			"bad": kvpBad.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"t": true,
			"empty": false,
			"bad": false,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueBool -- various inputs", actual)
	})
}

func Test_Seg1_KVP_ValueInt(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueInt", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "n", Value: "42"}
		kvpBad := corestr.KeyValuePair{Key: "n", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kvp.ValueInt(0),
			"def": kvp.ValueDefInt(),
			"bad": kvpBad.ValueInt(99),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 42,
			"bad": 99,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueInt -- valid and invalid", actual)
	})
}

func Test_Seg1_KVP_ValueByte(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueByte", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "b", Value: "200"}
		kvpBad := corestr.KeyValuePair{Key: "b", Value: "abc"}
		kvpOverflow := corestr.KeyValuePair{Key: "b", Value: "999"}

		// Act
		actual := args.Map{
			"val":      int(kvp.ValueByte(0)),
			"def":      int(kvp.ValueDefByte()),
			"bad":      int(kvpBad.ValueByte(5)),
			"overflow": int(kvpOverflow.ValueByte(7)),
		}

		// Assert
		expected := args.Map{
			"val":      200,
			"def":      200,
			"bad":      5,
			"overflow": 7,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueByte -- valid, invalid, overflow", actual)
	})
}

func Test_Seg1_KVP_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueFloat64", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "f", Value: "3.14"}
		kvpBad := corestr.KeyValuePair{Key: "f", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kvp.ValueFloat64(0),
			"def": kvp.ValueDefFloat64(),
			"bad": kvpBad.ValueFloat64(1.5),
		}

		// Assert
		expected := args.Map{
			"val": 3.14,
			"def": 3.14,
			"bad": 1.5,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueFloat64 -- valid and invalid", actual)
	})
}

func Test_Seg1_KVP_ValueValid(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueValid", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValid()

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"value": vv.Value,
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"value": "v",
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueValid -- default valid", actual)
	})
}

func Test_Seg1_KVP_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueValidOptions", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValidOptions(false, "err")

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
		expected.ShouldBeEqual(t, 0, "KVP ValueValidOptions -- invalid with message", actual)
	})
}

func Test_Seg1_KVP_Is(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Is", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"is":    kvp.Is("k", "v"),
			"isKey": kvp.IsKey("k"),
			"isVal": kvp.IsVal("v"),
			"notIs": kvp.Is("x", "y"),
		}

		// Assert
		expected := args.Map{
			"is":    true,
			"isKey": true,
			"isVal": true,
			"notIs": false,
		}
		expected.ShouldBeEqual(t, 0, "KVP Is/IsKey/IsVal -- match and no match", actual)
	})
}

func Test_Seg1_KVP_NilChecks(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_NilChecks", func() {
		// Arrange
		var kvp *corestr.KeyValuePair

		// Act
		actual := args.Map{"anyEmpty": kvp.IsKeyValueAnyEmpty()}

		// Assert
		expected := args.Map{"anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KVP nil IsKeyValueAnyEmpty -- nil receiver", actual)
	})
}

func Test_Seg1_KVP_FormatString(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_FormatString", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"fmt": kvp.FormatString("%v=%v")}

		// Assert
		expected := args.Map{"fmt": "k=v"}
		expected.ShouldBeEqual(t, 0, "KVP FormatString -- custom format", actual)
	})
}

func Test_Seg1_KVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Json", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kvp.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KVP Json -- no error", actual)
	})
}

func Test_Seg1_KVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Serialize", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kvp.Serialize()

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
		expected.ShouldBeEqual(t, 0, "KVP Serialize -- success", actual)
	})
}

func Test_Seg1_KVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_SerializeMust", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kvp.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KVP SerializeMust -- success", actual)
	})
}

func Test_Seg1_KVP_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ClearDispose", func() {
		// Arrange
		kvp := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvp.Clear()

		// Act
		actual := args.Map{
			"keyEmpty": kvp.Key == "",
			"valEmpty": kvp.Value == "",
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"valEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KVP Clear -- fields emptied", actual)
	})
}

func Test_Seg1_KVP_DisposeNil(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_DisposeNil", func() {
		var kvp *corestr.KeyValuePair
		kvp.Dispose() // should not panic
		kvp.Clear()   // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_KAVP_BasicAccessors(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_BasicAccessors", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "name", Value: 42}

		// Act
		actual := args.Map{
			"key":      kav.KeyName(),
			"varName":  kav.VariableName(),
			"valAny":   kav.ValueAny(),
			"isVarEq":  kav.IsVariableNameEqual("name"),
			"isNull":   kav.IsValueNull(),
			"hasValue": kav.HasValue(),
			"hasNon":   kav.HasNonNull(),
		}

		// Assert
		expected := args.Map{
			"key":      "name",
			"varName":  "name",
			"valAny":   42,
			"isVarEq":  true,
			"isNull":   false,
			"hasValue": true,
			"hasNon":   true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP basic accessors -- happy path", actual)
	})
}

func Test_Seg1_KAVP_NilValue(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_NilValue", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{
			"isNull":  kav.IsValueNull(),
			"emptyStr": kav.IsValueEmptyString(),
			"ws":       kav.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"isNull":  true,
			"emptyStr": true,
			"ws":       true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP nil value -- all empty checks true", actual)
	})
}

func Test_Seg1_KAVP_ValueString(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ValueString", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "hello"}

		// Act
		actual := args.Map{"valStr": kav.ValueString()}

		// Assert
		expected := args.Map{"valStr": "hello"}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString -- string value", actual)
	})
}

func Test_Seg1_KAVP_ValueStringCached(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ValueStringCached", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: 99}
		// First call initializes, second returns cached
		v1 := kav.ValueString()
		v2 := kav.ValueString()

		// Act
		actual := args.Map{"same": v1 == v2}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString cached -- same on second call", actual)
	})
}

func Test_Seg1_KAVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Json", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP Json -- no error", actual)
	})
}

func Test_Seg1_KAVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Serialize", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kav.Serialize()

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
		expected.ShouldBeEqual(t, 0, "KAVP Serialize -- success", actual)
	})
}

func Test_Seg1_KAVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_SerializeMust", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kav.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KAVP SerializeMust -- success", actual)
	})
}

func Test_Seg1_KAVP_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ParseInjectUsingJson", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result, err := kav2.ParseInjectUsingJson(jr)

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
		expected.ShouldBeEqual(t, 0, "KAVP ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg1_KAVP_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ParseInjectUsingJsonMust", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result := kav2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KAVP ParseInjectUsingJsonMust -- no panic", actual)
	})
}

func Test_Seg1_KAVP_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		kav := &corestr.KeyAnyValuePair{}
		badJson := &corejson.Result{}
		_ = kav.ParseInjectUsingJsonMust(badJson)
	})
}

func Test_Seg1_KAVP_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_JsonParseSelfInject", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		err := kav2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg1_KAVP_Interfaces(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Interfaces", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"binder":   kav.AsJsonContractsBinder() != nil,
			"jsoner":   kav.AsJsoner() != nil,
			"injector": kav.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"binder":   true,
			"jsoner":   true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP interface casts -- all non-nil", actual)
	})
}

func Test_Seg1_KAVP_String(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_String", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kav.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "KAVP String -- non-empty", actual)
	})
}

func Test_Seg1_KAVP_Compile(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Compile", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"eq": kav.Compile() == kav.String()}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "KAVP Compile equals String -- same result", actual)
	})
}

func Test_Seg1_KAVP_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ClearDispose", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()

		// Act
		actual := args.Map{
			"keyEmpty": kav.Key == "",
			"valNil": kav.Value == nil,
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"valNil": true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP Clear -- fields emptied", actual)
	})
}

func Test_Seg1_KAVP_DisposeNil(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_DisposeNil", func() {
		var kav *corestr.KeyAnyValuePair
		kav.Dispose() // should not panic
		kav.Clear()   // should not panic
	})
}

func Test_Seg1_KAVP_NilIsValueNull(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_NilIsValueNull", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"isNull": kav.IsValueNull()}

		// Assert
		expected := args.Map{"isNull": true}
		expected.ShouldBeEqual(t, 0, "KAVP nil receiver IsValueNull -- true", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_LeftRight_Creators(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_Creators", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight -- valid pair", actual)
	})
}

func Test_Seg1_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_Invalid", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("err")

		// Act
		actual := args.Map{
			"valid": lr.IsValid,
			"msg": lr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight -- invalid with message", actual)
	})
}

func Test_Seg1_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_InvalidNoMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage -- invalid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- 2 items valid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSliceEmpty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- empty slice invalid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSliceSingle(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSliceSingle", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"only"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"right": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- single item", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlicePtr", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- delegates to UsingSlice", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlicePtrEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlicePtrEmpty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- empty", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- trimmed", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceNil(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceNil", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- nil", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceEmpty", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- empty", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceSingle(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceSingle", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- single item not trimmed", actual)
	})
}

func Test_Seg1_LeftRight_StringMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_StringMethods", func() {
		// Arrange
		lr := corestr.NewLeftRight(" hello ", " world ")

		// Act
		actual := args.Map{
			"leftBytes":  string(lr.LeftBytes()),
			"rightBytes": string(lr.RightBytes()),
			"leftTrim":   lr.LeftTrim(),
			"rightTrim":  lr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"leftBytes":  " hello ",
			"rightBytes": " world ",
			"leftTrim":   "hello",
			"rightTrim":  "world",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight string methods -- bytes and trim", actual)
	})
}

func Test_Seg1_LeftRight_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_EmptyChecks", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "")

		// Act
		actual := args.Map{
			"leftEmpty":  lr.IsLeftEmpty(),
			"rightEmpty": lr.IsRightEmpty(),
			"leftWS":     lr.IsLeftWhitespace(),
			"rightWS":    lr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty":  true,
			"rightEmpty": true,
			"leftWS":     true,
			"rightWS":    true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight empty checks -- empty strings", actual)
	})
}

func Test_Seg1_LeftRight_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_ValidNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"validLeft":    lr.HasValidNonEmptyLeft(),
			"validRight":   lr.HasValidNonEmptyRight(),
			"validWSLeft":  lr.HasValidNonWhitespaceLeft(),
			"validWSRight": lr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft":    true,
			"validRight":   true,
			"validWSLeft":  true,
			"validWSRight": true,
			"safeNonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight HasValidNonEmpty -- all valid", actual)
	})
}

func Test_Seg1_LeftRight_NonPtrPtr(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_NonPtrPtr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		p := lr.Ptr()

		// Act
		actual := args.Map{
			"nonPtrLeft": np.Left,
			"ptrLeft": p.Left,
		}

		// Assert
		expected := args.Map{
			"nonPtrLeft": "a",
			"ptrLeft": "a",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight NonPtr/Ptr -- same values", actual)
	})
}

func Test_Seg1_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_RegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello123", "world456")
		re := regexp.MustCompile(`[0-9]+`)

		// Act
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilRegex":   lr.IsLeftRegexMatch(nil),
		}

		// Assert
		expected := args.Map{
			"leftMatch":  true,
			"rightMatch": true,
			"nilRegex":   false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight regex match -- valid and nil regex", actual)
	})
}

func Test_Seg1_LeftRight_IsComparisons(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_IsComparisons", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"isLeft":  lr.IsLeft("a"),
			"isRight": lr.IsRight("b"),
			"is":      lr.Is("a", "b"),
		}

		// Assert
		expected := args.Map{
			"isLeft":  true,
			"isRight": true,
			"is":      true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Is comparisons -- match", actual)
	})
}

func Test_Seg1_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")

		// Act
		actual := args.Map{
			"eq":      lr1.IsEqual(lr2),
			"neq":     lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight IsEqual -- equal, not equal, nil cases", actual)
	})
}

func Test_Seg1_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"right": c.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Clone -- same values", actual)
	})
}

func Test_Seg1_LeftRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_ClearDispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Clear -- emptied", actual)
	})
}

func Test_Seg1_LeftRight_DisposeNil(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_DisposeNil", func() {
		var lr *corestr.LeftRight
		lr.Dispose() // should not panic
		lr.Clear()   // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_LMR_Creators(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_Creators", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
			"valid": lmr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight -- valid triple", actual)
	})
}

func Test_Seg1_LMR_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_Invalid", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{
			"valid": lmr.IsValid,
			"msg": lmr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight -- invalid with message", actual)
	})
}

func Test_Seg1_LMR_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage -- invalid", actual)
	})
}

func Test_Seg1_LMR_BytesMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_BytesMethods", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		actual := args.Map{
			"leftB":  string(lmr.LeftBytes()),
			"midB":   string(lmr.MiddleBytes()),
			"rightB": string(lmr.RightBytes()),
		}

		// Assert
		expected := args.Map{
			"leftB":  "L",
			"midB":   "M",
			"rightB": "R",
		}
		expected.ShouldBeEqual(t, 0, "LMR Bytes methods -- correct bytes", actual)
	})
}

func Test_Seg1_LMR_TrimMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_TrimMethods", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" L ", " M ", " R ")

		// Act
		actual := args.Map{
			"leftTrim":  lmr.LeftTrim(),
			"midTrim":   lmr.MiddleTrim(),
			"rightTrim": lmr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"leftTrim":  "L",
			"midTrim":   "M",
			"rightTrim": "R",
		}
		expected.ShouldBeEqual(t, 0, "LMR Trim methods -- trimmed", actual)
	})
}

func Test_Seg1_LMR_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_EmptyChecks", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "", "")

		// Act
		actual := args.Map{
			"leftEmpty":  lmr.IsLeftEmpty(),
			"midEmpty":   lmr.IsMiddleEmpty(),
			"rightEmpty": lmr.IsRightEmpty(),
			"leftWS":     lmr.IsLeftWhitespace(),
			"midWS":      lmr.IsMiddleWhitespace(),
			"rightWS":    lmr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty":  true,
			"midEmpty":   true,
			"rightEmpty": true,
			"leftWS":     true,
			"midWS":      true,
			"rightWS":    true,
		}
		expected.ShouldBeEqual(t, 0, "LMR empty checks -- all empty", actual)
	})
}

func Test_Seg1_LMR_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_ValidNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"validLeft":    lmr.HasValidNonEmptyLeft(),
			"validMid":     lmr.HasValidNonEmptyMiddle(),
			"validRight":   lmr.HasValidNonEmptyRight(),
			"validWSLeft":  lmr.HasValidNonWhitespaceLeft(),
			"validWSMid":   lmr.HasValidNonWhitespaceMiddle(),
			"validWSRight": lmr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lmr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft":    true,
			"validMid":     true,
			"validRight":   true,
			"validWSLeft":  true,
			"validWSMid":   true,
			"validWSRight": true,
			"safeNonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmpty -- all valid", actual)
	})
}

func Test_Seg1_LMR_IsAll(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"isAll": lmr.IsAll("a", "b", "c"),
			"is":    lmr.Is("a", "c"),
		}

		// Assert
		expected := args.Map{
			"isAll": true,
			"is":    true,
		}
		expected.ShouldBeEqual(t, 0, "LMR IsAll and Is -- match", actual)
	})
}

func Test_Seg1_LMR_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"mid": c.Middle,
			"right": c.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clone -- same values", actual)
	})
}

func Test_Seg1_LMR_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LMR ToLeftRight -- left and right preserved", actual)
	})
}

func Test_Seg1_LMR_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_ClearDispose", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"mid": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clear -- emptied", actual)
	})
}

func Test_Seg1_LMR_DisposeNil(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_DisposeNil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose() // should not panic
		lmr.Clear()   // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit / LeftRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_LMRFromSplit(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit -- 3 parts", actual)
	})
}

func Test_Seg1_LMRFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LMRFromSplitN(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitN", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c:d:e",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN -- remainder in right", actual)
	})
}

func Test_Seg1_LMRFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitNTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LRFromSplit(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit -- 2 parts", actual)
	})
}

func Test_Seg1_LRFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LRFromSplitFull(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitFull", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b:c:d",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull -- remainder in right", actual)
	})
}

func Test_Seg1_LRFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitFullTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{"left": lr.Left}

		// Assert
		expected := args.Map{"left": "a"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed -- trimmed", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — first ~200 statements (basic operations)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_BasicOps", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"len":      c.Length(),
			"count":    c.Count(),
			"hasAny":   c.HasAnyItem(),
			"hasItems": c.HasItems(),
			"isEmpty":  c.IsEmpty(),
			"lastIdx":  c.LastIndex(),
			"hasIdx0":  c.HasIndex(0),
			"hasIdx5":  c.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"len":      3,
			"count":    3,
			"hasAny":   true,
			"hasItems": true,
			"isEmpty":  false,
			"lastIdx":  2,
			"hasIdx0":  true,
			"hasIdx5":  false,
		}
		expected.ShouldBeEqual(t, 0, "Collection basic ops -- 3 items", actual)
	})
}

func Test_Seg1_Collection_NilLength(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_NilLength", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{
			"len": c.Length(),
			"empty": c.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection nil receiver -- length 0", actual)
	})
}

func Test_Seg1_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_Capacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 10}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection Capacity -- at least 10", actual)
	})
}

func Test_Seg1_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")
		ok := c.RemoveAt(1)

		// Act
		actual := args.Map{
			"ok": ok,
			"len": c.Length(),
			"first": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- middle removed", actual)
	})
}

func Test_Seg1_Collection_RemoveAt_OutOfBounds(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_RemoveAt_OutOfBounds", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{
			"neg": c.RemoveAt(-1),
			"over": c.RemoveAt(5),
		}

		// Assert
		expected := args.Map{
			"neg": false,
			"over": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- out of bounds false", actual)
	})
}

func Test_Seg1_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ListStringsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{"len": len(c.ListStringsPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection ListStringsPtr -- returns items", actual)
	})
}

func Test_Seg1_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmpty("a").AddNonEmpty("").AddNonEmpty("b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmpty -- skips empty", actual)
	})
}

func Test_Seg1_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("   ").AddNonEmptyWhitespace("b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmptyWhitespace -- skips whitespace", actual)
	})
}

func Test_Seg1_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddError(errors.New("err1")).AddError(nil).AddError(errors.New("err2"))

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddError -- skips nil error", actual)
	})
}

func Test_Seg1_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AsError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.AsError("; ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- non-nil error", actual)
	})
}

func Test_Seg1_Collection_AsErrorEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AsErrorEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{
			"nil": c.AsError("; ") == nil,
			"defNil": c.AsDefaultError() == nil,
		}

		// Assert
		expected := args.Map{
			"nil": true,
			"defNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- nil when empty", actual)
	})
}

func Test_Seg1_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes").AddIf(false, "no")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "yes",
		}
		expected.ShouldBeEqual(t, 0, "Collection AddIf -- only true added", actual)
	})
}

func Test_Seg1_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(true, "a", "b").AddIfMany(false, "c", "d")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddIfMany -- only true batch added", actual)
	})
}

func Test_Seg1_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{
			"len": c.Length(),
			"val": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Collection AddFunc -- adds func result", actual)
	})
}

func Test_Seg1_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFuncErr_Success", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- success path", actual)
	})
}

func Test_Seg1_Collection_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFuncErr_Error", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		called := false
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { called = true },
		)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"called": called,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"called": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- error path", actual)
	})
}

func Test_Seg1_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a,b").Add("c,d")
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection EachItemSplitBy -- 4 items", actual)
	})
}

func Test_Seg1_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection Adds -- 3 items", actual)
	})
}

func Test_Seg1_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddStrings -- 2 items", actual)
	})
}

func Test_Seg1_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollection", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b").Add("c")
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection -- merged", actual)
	})
}

func Test_Seg1_Collection_AddCollectionEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollectionEmpty", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection empty -- no change", actual)
	})
}

func Test_Seg1_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b")
		empty := corestr.New.Collection.Cap(5)
		c.AddCollections(c1, empty, c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddCollections -- skips empty", actual)
	})
}

func Test_Seg1_Collection_LockMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_LockMethods", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")
		c.AddsLock("b", "c")

		// Act
		actual := args.Map{
			"len": c.LengthLock(),
			"emptyLock": c.IsEmptyLock(),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"emptyLock": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection lock methods -- 3 items", actual)
	})
}

func Test_Seg1_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEquals", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("a", "b")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("a", "b")
		c3 := corestr.New.Collection.Cap(5)
		c3.Adds("a", "c")

		// Act
		actual := args.Map{
			"eq":   c1.IsEquals(c2),
			"neq":  c1.IsEquals(c3),
		}

		// Assert
		expected := args.Map{
			"eq":   true,
			"neq":  false,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEquals -- equal and not equal", actual)
	})
}

func Test_Seg1_Collection_IsEqualsInsensitive(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEqualsInsensitive", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("Hello", "World")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("hello", "world")

		// Act
		actual := args.Map{
			"sensitive":   c1.IsEqualsWithSensitive(true, c2),
			"insensitive": c1.IsEqualsWithSensitive(false, c2),
		}

		// Assert
		expected := args.Map{
			"sensitive":   false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEqualsWithSensitive -- case comparison", actual)
	})
}

func Test_Seg1_Collection_IsEqualsEdge(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEqualsEdge", func() {
		// Arrange
		var nilC *corestr.Collection
		emptyC := corestr.New.Collection.Cap(0)
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{
			"bothNil":   nilC.IsEquals(nil),
			"nilVsNon":  nilC.IsEquals(c),
			"emptyBoth": emptyC.IsEquals(corestr.New.Collection.Cap(0)),
			"diffLen":   c.IsEquals(emptyC),
		}

		// Assert
		expected := args.Map{
			"bothNil":   true,
			"nilVsNon":  false,
			"emptyBoth": true,
			"diffLen":   false,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEquals edge -- nil and empty", actual)
	})
}

func Test_Seg1_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.ToError("; ")
		defErr := c.ToDefaultError()

		// Act
		actual := args.Map{
			"notNil": err != nil,
			"defNotNil": defErr != nil,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"defNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection ToError -- non-nil", actual)
	})
}

func Test_Seg1_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0, "c", "d")

		// Act
		actual := args.Map{
			"origLen": c.Length(),
			"newLen": c2.Length(),
		}

		// Assert
		expected := args.Map{
			"origLen": 2,
			"newLen": 4,
		}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- new collection with all items", actual)
	})
}

func Test_Seg1_Collection_ConcatNewEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ConcatNewEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0)

		// Act
		actual := args.Map{"newLen": c2.Length()}

		// Assert
		expected := args.Map{"newLen": 2}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- empty adds returns copy", actual)
	})
}

func Test_Seg1_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		s := c.JsonString()
		s2 := c.JsonStringMust()
		s3 := c.StringJSON()

		// Act
		actual := args.Map{
			"nonEmpty": s != "",
			"eq": s == s2,
			"eq2": s == s3,
		}

		// Assert
		expected := args.Map{
			"nonEmpty": true,
			"eq": true,
			"eq2": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection JsonString -- all variants match", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsValues(h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues -- 2 values added", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsValuesNil(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsValuesNil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues nil -- no items", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsKeys(h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys -- 2 keys added", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsKeysNil(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsKeysNil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys nil -- no items", actual)
	})
}

func Test_Seg1_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddPointerCollectionsLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("a")
		c.AddPointerCollectionsLock(c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddPointerCollectionsLock -- 1 item", actual)
	})
}
