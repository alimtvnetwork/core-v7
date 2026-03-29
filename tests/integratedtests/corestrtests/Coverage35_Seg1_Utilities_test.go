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
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil ptr", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_NilSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_NilSlice", func() {
		var s [][]string
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil inner slice", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_WithItems", func() {
		s := [][]string{{"a", "b"}, {"c"}}
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns total -- multi", actual)
	})
}

func Test_Seg1_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualStringsOfStringsLength_Empty", func() {
		s := [][]string{}
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- empty outer", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- no args", actual)
	})
}

func Test_Seg1_AllIndividualsLengthOfSimpleSlices_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_AllIndividualsLengthOfSimpleSlices_WithItems", func() {
		s1 := corestr.SimpleSlice{"a", "b"}
		s2 := corestr.SimpleSlice{"c"}
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(&s1, &s2)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns total -- multi slices", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_Empty", func() {
		actual := args.Map{"val": corestr.AnyToString(false, "")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty input", actual)
	})
}

func Test_Seg1_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_WithFieldName", func() {
		result := corestr.AnyToString(true, 42)
		actual := args.Map{"nonEmpty": result != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field name", actual)
	})
}

func Test_Seg1_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_Seg1_AnyToString_WithoutFieldName", func() {
		result := corestr.AnyToString(false, 42)
		actual := args.Map{"nonEmpty": result != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- without field name", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_Empty", func() {
		result := corestr.CloneSlice([]string{})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty input", actual)
	})
}

func Test_Seg1_CloneSlice_WithItems(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_WithItems", func() {
		src := []string{"a", "b", "c"}
		result := corestr.CloneSlice(src)
		actual := args.Map{"len": len(result), "eq": result[0] == "a" && result[2] == "c"}
		expected := args.Map{"len": 3, "eq": true}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns copy -- with items", actual)
	})
}

func Test_Seg1_CloneSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSlice_Nil", func() {
		result := corestr.CloneSlice(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- nil input", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_Clone", func() {
		result := corestr.CloneSliceIf(true, "a", "b")
		actual := args.Map{"len": len(result), "first": result[0]}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns clone -- isClone true", actual)
	})
}

func Test_Seg1_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_NoClone", func() {
		result := corestr.CloneSliceIf(false, "a", "b")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns same -- isClone false", actual)
	})
}

func Test_Seg1_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg1_CloneSliceIf_Empty", func() {
		result := corestr.CloneSliceIf(true)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- no items", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_HasLineNumber", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		actual := args.Map{"has": tw.HasLineNumber(), "invalid": tw.IsInvalidLineNumber()}
		expected := args.Map{"has": true, "invalid": false}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber HasLineNumber true -- valid line", actual)
	})
}

func Test_Seg1_TextWithLineNumber_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_Invalid", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}
		actual := args.Map{"has": tw.HasLineNumber(), "invalid": tw.IsInvalidLineNumber()}
		expected := args.Map{"has": false, "invalid": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsInvalidLineNumber -- invalid line", actual)
	})
}

func Test_Seg1_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_Length", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		actual := args.Map{"len": tw.Length()}
		expected := args.Map{"len": 5}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- 5 chars", actual)
	})
}

func Test_Seg1_TextWithLineNumber_NilLength(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_NilLength", func() {
		var tw *corestr.TextWithLineNumber
		actual := args.Map{"len": tw.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- nil", actual)
	})
}

func Test_Seg1_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_IsEmpty", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}
		actual := args.Map{"empty": tw.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- invalid line number", actual)
	})
}

func Test_Seg1_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_IsEmptyText", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		actual := args.Map{"emptyText": tw.IsEmptyText(), "emptyBoth": tw.IsEmptyTextLineBoth()}
		expected := args.Map{"emptyText": true, "emptyBoth": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmptyText -- empty text", actual)
	})
}

func Test_Seg1_TextWithLineNumber_NilEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_TextWithLineNumber_NilEmpty", func() {
		var tw *corestr.TextWithLineNumber
		actual := args.Map{"empty": tw.IsEmpty(), "emptyText": tw.IsEmptyText()}
		expected := args.Map{"empty": true, "emptyText": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- nil receiver", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"notNil": vs != nil, "index": vs.Index}
		expected := args.Map{"notNil": true, "index": -1}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns valid struct -- default", actual)
	})
}

func Test_Seg1_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("err msg")
		actual := args.Map{"notNil": vs != nil, "index": vs.Index}
		expected := args.Map{"notNil": true, "index": -1}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns valid struct -- with message", actual)
	})
}

func Test_Seg1_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("clone me")
		cloned := vs.Clone()
		actual := args.Map{"notNil": cloned != nil, "sameIdx": cloned.Index == vs.Index}
		expected := args.Map{"notNil": true, "sameIdx": true}
		expected.ShouldBeEqual(t, 0, "ValueStatus Clone returns copy -- same index", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_KVP_BasicAccessors(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_BasicAccessors", func() {
		kvp := corestr.KeyValuePair{Key: "name", Value: "alice"}
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
		kvp := corestr.KeyValuePair{}
		actual := args.Map{
			"keyEmpty":    kvp.IsKeyEmpty(),
			"valEmpty":    kvp.IsValueEmpty(),
			"kvEmpty":     kvp.IsKeyValueEmpty(),
			"kvAnyEmpty":  kvp.IsKeyValueAnyEmpty(),
		}
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
		kvp := corestr.KeyValuePair{Key: " name ", Value: " 42 "}
		actual := args.Map{
			"trimKey": kvp.TrimKey(),
			"trimVal": kvp.TrimValue(),
		}
		expected := args.Map{
			"trimKey": "name",
			"trimVal": "42",
		}
		expected.ShouldBeEqual(t, 0, "KVP Trim -- whitespace removed", actual)
	})
}

func Test_Seg1_KVP_ValueBool(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueBool", func() {
		kvp := corestr.KeyValuePair{Key: "flag", Value: "true"}
		kvpEmpty := corestr.KeyValuePair{Key: "flag", Value: ""}
		kvpBad := corestr.KeyValuePair{Key: "flag", Value: "notabool"}
		actual := args.Map{"t": kvp.ValueBool(), "empty": kvpEmpty.ValueBool(), "bad": kvpBad.ValueBool()}
		expected := args.Map{"t": true, "empty": false, "bad": false}
		expected.ShouldBeEqual(t, 0, "KVP ValueBool -- various inputs", actual)
	})
}

func Test_Seg1_KVP_ValueInt(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueInt", func() {
		kvp := corestr.KeyValuePair{Key: "n", Value: "42"}
		kvpBad := corestr.KeyValuePair{Key: "n", Value: "abc"}
		actual := args.Map{"val": kvp.ValueInt(0), "def": kvp.ValueDefInt(), "bad": kvpBad.ValueInt(99)}
		expected := args.Map{"val": 42, "def": 42, "bad": 99}
		expected.ShouldBeEqual(t, 0, "KVP ValueInt -- valid and invalid", actual)
	})
}

func Test_Seg1_KVP_ValueByte(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueByte", func() {
		kvp := corestr.KeyValuePair{Key: "b", Value: "200"}
		kvpBad := corestr.KeyValuePair{Key: "b", Value: "abc"}
		kvpOverflow := corestr.KeyValuePair{Key: "b", Value: "999"}
		actual := args.Map{
			"val":      int(kvp.ValueByte(0)),
			"def":      int(kvp.ValueDefByte()),
			"bad":      int(kvpBad.ValueByte(5)),
			"overflow": int(kvpOverflow.ValueByte(7)),
		}
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
		kvp := corestr.KeyValuePair{Key: "f", Value: "3.14"}
		kvpBad := corestr.KeyValuePair{Key: "f", Value: "abc"}
		actual := args.Map{"val": kvp.ValueFloat64(0), "def": kvp.ValueDefFloat64(), "bad": kvpBad.ValueFloat64(1.5)}
		expected := args.Map{"val": 3.14, "def": 3.14, "bad": 1.5}
		expected.ShouldBeEqual(t, 0, "KVP ValueFloat64 -- valid and invalid", actual)
	})
}

func Test_Seg1_KVP_ValueValid(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueValid", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValid()
		actual := args.Map{"valid": vv.IsValid, "value": vv.Value}
		expected := args.Map{"valid": true, "value": "v"}
		expected.ShouldBeEqual(t, 0, "KVP ValueValid -- default valid", actual)
	})
}

func Test_Seg1_KVP_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ValueValidOptions", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValidOptions(false, "err")
		actual := args.Map{"valid": vv.IsValid, "msg": vv.Message}
		expected := args.Map{"valid": false, "msg": "err"}
		expected.ShouldBeEqual(t, 0, "KVP ValueValidOptions -- invalid with message", actual)
	})
}

func Test_Seg1_KVP_Is(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Is", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"is":    kvp.Is("k", "v"),
			"isKey": kvp.IsKey("k"),
			"isVal": kvp.IsVal("v"),
			"notIs": kvp.Is("x", "y"),
		}
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
		var kvp *corestr.KeyValuePair
		actual := args.Map{"anyEmpty": kvp.IsKeyValueAnyEmpty()}
		expected := args.Map{"anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KVP nil IsKeyValueAnyEmpty -- nil receiver", actual)
	})
}

func Test_Seg1_KVP_FormatString(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_FormatString", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"fmt": kvp.FormatString("%v=%v")}
		expected := args.Map{"fmt": "k=v"}
		expected.ShouldBeEqual(t, 0, "KVP FormatString -- custom format", actual)
	})
}

func Test_Seg1_KVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Json", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kvp.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KVP Json -- no error", actual)
	})
}

func Test_Seg1_KVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_Serialize", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kvp.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KVP Serialize -- success", actual)
	})
}

func Test_Seg1_KVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_SerializeMust", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kvp.SerializeMust()
		actual := args.Map{"hasBytes": len(b) > 0}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KVP SerializeMust -- success", actual)
	})
}

func Test_Seg1_KVP_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_KVP_ClearDispose", func() {
		kvp := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvp.Clear()
		actual := args.Map{"keyEmpty": kvp.Key == "", "valEmpty": kvp.Value == ""}
		expected := args.Map{"keyEmpty": true, "valEmpty": true}
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
		kav := &corestr.KeyAnyValuePair{Key: "name", Value: 42}
		actual := args.Map{
			"key":      kav.KeyName(),
			"varName":  kav.VariableName(),
			"valAny":   kav.ValueAny(),
			"isVarEq":  kav.IsVariableNameEqual("name"),
			"isNull":   kav.IsValueNull(),
			"hasValue": kav.HasValue(),
			"hasNon":   kav.HasNonNull(),
		}
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
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}
		actual := args.Map{
			"isNull":  kav.IsValueNull(),
			"emptyStr": kav.IsValueEmptyString(),
			"ws":       kav.IsValueWhitespace(),
		}
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
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
		actual := args.Map{"valStr": kav.ValueString()}
		expected := args.Map{"valStr": "hello"}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString -- string value", actual)
	})
}

func Test_Seg1_KAVP_ValueStringCached(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ValueStringCached", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: 99}
		// First call initializes, second returns cached
		v1 := kav.ValueString()
		v2 := kav.ValueString()
		actual := args.Map{"same": v1 == v2}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString cached -- same on second call", actual)
	})
}

func Test_Seg1_KAVP_Json(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Json", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP Json -- no error", actual)
	})
}

func Test_Seg1_KAVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Serialize", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kav.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KAVP Serialize -- success", actual)
	})
}

func Test_Seg1_KAVP_SerializeMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_SerializeMust", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kav.SerializeMust()
		actual := args.Map{"hasBytes": len(b) > 0}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KAVP SerializeMust -- success", actual)
	})
}

func Test_Seg1_KAVP_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ParseInjectUsingJson", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result, err := kav2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "notNil": result != nil}
		expected := args.Map{"noErr": true, "notNil": true}
		expected.ShouldBeEqual(t, 0, "KAVP ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg1_KAVP_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ParseInjectUsingJsonMust", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result := kav2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"notNil": result != nil}
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
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		err := kav2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg1_KAVP_Interfaces(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Interfaces", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"binder":   kav.AsJsonContractsBinder() != nil,
			"jsoner":   kav.AsJsoner() != nil,
			"injector": kav.AsJsonParseSelfInjector() != nil,
		}
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
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"nonEmpty": kav.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "KAVP String -- non-empty", actual)
	})
}

func Test_Seg1_KAVP_Compile(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_Compile", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"eq": kav.Compile() == kav.String()}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "KAVP Compile equals String -- same result", actual)
	})
}

func Test_Seg1_KAVP_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_KAVP_ClearDispose", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()
		actual := args.Map{"keyEmpty": kav.Key == "", "valNil": kav.Value == nil}
		expected := args.Map{"keyEmpty": true, "valNil": true}
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
		var kav *corestr.KeyAnyValuePair
		actual := args.Map{"isNull": kav.IsValueNull()}
		expected := args.Map{"isNull": true}
		expected.ShouldBeEqual(t, 0, "KAVP nil receiver IsValueNull -- true", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_LeftRight_Creators(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_Creators", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewLeftRight -- valid pair", actual)
	})
}

func Test_Seg1_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_Invalid", func() {
		lr := corestr.InvalidLeftRight("err")
		actual := args.Map{"valid": lr.IsValid, "msg": lr.Message}
		expected := args.Map{"valid": false, "msg": "err"}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight -- invalid with message", actual)
	})
}

func Test_Seg1_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage -- invalid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- 2 items valid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSliceEmpty", func() {
		lr := corestr.LeftRightUsingSlice([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- empty slice invalid", actual)
	})
}

func Test_Seg1_LeftRight_UsingSliceSingle(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSliceSingle", func() {
		lr := corestr.LeftRightUsingSlice([]string{"only"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "only", "right": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- single item", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlicePtr", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual := args.Map{"left": lr.Left, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- delegates to UsingSlice", actual)
	})
}

func Test_Seg1_LeftRight_UsingSlicePtrEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_UsingSlicePtrEmpty", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- empty", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- trimmed", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceNil(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceNil", func() {
		lr := corestr.LeftRightTrimmedUsingSlice(nil)
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- nil", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceEmpty", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- empty", actual)
	})
}

func Test_Seg1_LeftRight_TrimmedUsingSliceSingle(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_TrimmedUsingSliceSingle", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})
		actual := args.Map{"left": lr.Left, "valid": lr.IsValid}
		expected := args.Map{"left": "only", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- single item not trimmed", actual)
	})
}

func Test_Seg1_LeftRight_StringMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_StringMethods", func() {
		lr := corestr.NewLeftRight(" hello ", " world ")
		actual := args.Map{
			"leftBytes":  string(lr.LeftBytes()),
			"rightBytes": string(lr.RightBytes()),
			"leftTrim":   lr.LeftTrim(),
			"rightTrim":  lr.RightTrim(),
		}
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
		lr := corestr.NewLeftRight("", "")
		actual := args.Map{
			"leftEmpty":  lr.IsLeftEmpty(),
			"rightEmpty": lr.IsRightEmpty(),
			"leftWS":     lr.IsLeftWhitespace(),
			"rightWS":    lr.IsRightWhitespace(),
		}
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
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{
			"validLeft":    lr.HasValidNonEmptyLeft(),
			"validRight":   lr.HasValidNonEmptyRight(),
			"validWSLeft":  lr.HasValidNonWhitespaceLeft(),
			"validWSRight": lr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lr.HasSafeNonEmpty(),
		}
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
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		p := lr.Ptr()
		actual := args.Map{"nonPtrLeft": np.Left, "ptrLeft": p.Left}
		expected := args.Map{"nonPtrLeft": "a", "ptrLeft": "a"}
		expected.ShouldBeEqual(t, 0, "LeftRight NonPtr/Ptr -- same values", actual)
	})
}

func Test_Seg1_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_RegexMatch", func() {
		lr := corestr.NewLeftRight("hello123", "world456")
		re := regexp.MustCompile(`[0-9]+`)
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilRegex":   lr.IsLeftRegexMatch(nil),
		}
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
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{
			"isLeft":  lr.IsLeft("a"),
			"isRight": lr.IsRight("b"),
			"is":      lr.Is("a", "b"),
		}
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
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")
		actual := args.Map{
			"eq":      lr1.IsEqual(lr2),
			"neq":     lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}
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
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()
		actual := args.Map{"left": c.Left, "right": c.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRight Clone -- same values", actual)
	})
}

func Test_Seg1_LeftRight_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_LeftRight_ClearDispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "", "right": ""}
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
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right, "valid": lmr.IsValid}
		expected := args.Map{"left": "a", "mid": "b", "right": "c", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight -- valid triple", actual)
	})
}

func Test_Seg1_LMR_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		actual := args.Map{"valid": lmr.IsValid, "msg": lmr.Message}
		expected := args.Map{"valid": false, "msg": "err"}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight -- invalid with message", actual)
	})
}

func Test_Seg1_LMR_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"valid": lmr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage -- invalid", actual)
	})
}

func Test_Seg1_LMR_BytesMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_BytesMethods", func() {
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")
		actual := args.Map{
			"leftB":  string(lmr.LeftBytes()),
			"midB":   string(lmr.MiddleBytes()),
			"rightB": string(lmr.RightBytes()),
		}
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
		lmr := corestr.NewLeftMiddleRight(" L ", " M ", " R ")
		actual := args.Map{
			"leftTrim":  lmr.LeftTrim(),
			"midTrim":   lmr.MiddleTrim(),
			"rightTrim": lmr.RightTrim(),
		}
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
		lmr := corestr.NewLeftMiddleRight("", "", "")
		actual := args.Map{
			"leftEmpty":  lmr.IsLeftEmpty(),
			"midEmpty":   lmr.IsMiddleEmpty(),
			"rightEmpty": lmr.IsRightEmpty(),
			"leftWS":     lmr.IsLeftWhitespace(),
			"midWS":      lmr.IsMiddleWhitespace(),
			"rightWS":    lmr.IsRightWhitespace(),
		}
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
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{
			"validLeft":    lmr.HasValidNonEmptyLeft(),
			"validMid":     lmr.HasValidNonEmptyMiddle(),
			"validRight":   lmr.HasValidNonEmptyRight(),
			"validWSLeft":  lmr.HasValidNonWhitespaceLeft(),
			"validWSMid":   lmr.HasValidNonWhitespaceMiddle(),
			"validWSRight": lmr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lmr.HasSafeNonEmpty(),
		}
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
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{
			"isAll": lmr.IsAll("a", "b", "c"),
			"is":    lmr.Is("a", "c"),
		}
		expected := args.Map{
			"isAll": true,
			"is":    true,
		}
		expected.ShouldBeEqual(t, 0, "LMR IsAll and Is -- match", actual)
	})
}

func Test_Seg1_LMR_Clone(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()
		actual := args.Map{"left": c.Left, "mid": c.Middle, "right": c.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LMR Clone -- same values", actual)
	})
}

func Test_Seg1_LMR_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "c", "valid": true}
		expected.ShouldBeEqual(t, 0, "LMR ToLeftRight -- left and right preserved", actual)
	})
}

func Test_Seg1_LMR_ClearDispose(t *testing.T) {
	safeTest(t, "Test_Seg1_LMR_ClearDispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "", "mid": "", "right": ""}
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
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit -- 3 parts", actual)
	})
}

func Test_Seg1_LMRFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a : b : c ", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LMRFromSplitN(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c:d:e"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN -- remainder in right", actual)
	})
}

func Test_Seg1_LMRFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LMRFromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle}
		expected := args.Map{"left": "a", "mid": "b"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LRFromSplit(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit -- 2 parts", actual)
	})
}

func Test_Seg1_LRFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg1_LRFromSplitFull(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b:c:d"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull -- remainder in right", actual)
	})
}

func Test_Seg1_LRFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg1_LRFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		actual := args.Map{"left": lr.Left}
		expected := args.Map{"left": "a"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed -- trimmed", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — first ~200 statements (basic operations)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg1_Collection_BasicOps(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_BasicOps", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")
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
		var c *corestr.Collection
		actual := args.Map{"len": c.Length(), "empty": c.IsEmpty()}
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "Collection nil receiver -- length 0", actual)
	})
}

func Test_Seg1_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"cap": c.Capacity() >= 10}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection Capacity -- at least 10", actual)
	})
}

func Test_Seg1_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")
		ok := c.RemoveAt(1)
		actual := args.Map{"ok": ok, "len": c.Length(), "first": c.ListStrings()[0]}
		expected := args.Map{"ok": true, "len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- middle removed", actual)
	})
}

func Test_Seg1_Collection_RemoveAt_OutOfBounds(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_RemoveAt_OutOfBounds", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		actual := args.Map{"neg": c.RemoveAt(-1), "over": c.RemoveAt(5)}
		expected := args.Map{"neg": false, "over": false}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- out of bounds false", actual)
	})
}

func Test_Seg1_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		actual := args.Map{"len": len(c.ListStringsPtr())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection ListStringsPtr -- returns items", actual)
	})
}

func Test_Seg1_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmpty("a").AddNonEmpty("").AddNonEmpty("b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmpty -- skips empty", actual)
	})
}

func Test_Seg1_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("   ").AddNonEmptyWhitespace("b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmptyWhitespace -- skips whitespace", actual)
	})
}

func Test_Seg1_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddError", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddError(errors.New("err1")).AddError(nil).AddError(errors.New("err2"))
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddError -- skips nil error", actual)
	})
}

func Test_Seg1_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AsError", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.AsError("; ")
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- non-nil error", actual)
	})
}

func Test_Seg1_Collection_AsErrorEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AsErrorEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"nil": c.AsError("; ") == nil, "defNil": c.AsDefaultError() == nil}
		expected := args.Map{"nil": true, "defNil": true}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- nil when empty", actual)
	})
}

func Test_Seg1_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddIf", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes").AddIf(false, "no")
		actual := args.Map{"len": c.Length(), "first": c.ListStrings()[0]}
		expected := args.Map{"len": 1, "first": "yes"}
		expected.ShouldBeEqual(t, 0, "Collection AddIf -- only true added", actual)
	})
}

func Test_Seg1_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(true, "a", "b").AddIfMany(false, "c", "d")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddIfMany -- only true batch added", actual)
	})
}

func Test_Seg1_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFunc", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFunc(func() string { return "hello" })
		actual := args.Map{"len": c.Length(), "val": c.ListStrings()[0]}
		expected := args.Map{"len": 1, "val": "hello"}
		expected.ShouldBeEqual(t, 0, "Collection AddFunc -- adds func result", actual)
	})
}

func Test_Seg1_Collection_AddFuncErr_Success(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFuncErr_Success", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { t.Fatal("should not be called") },
		)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- success path", actual)
	})
}

func Test_Seg1_Collection_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddFuncErr_Error", func() {
		c := corestr.New.Collection.Cap(5)
		called := false
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { called = true },
		)
		actual := args.Map{"len": c.Length(), "called": called}
		expected := args.Map{"len": 0, "called": true}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- error path", actual)
	})
}

func Test_Seg1_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a,b").Add("c,d")
		result := c.EachItemSplitBy(",")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection EachItemSplitBy -- 4 items", actual)
	})
}

func Test_Seg1_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_Adds", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection Adds -- 3 items", actual)
	})
}

func Test_Seg1_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddStrings", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddStrings -- 2 items", actual)
	})
}

func Test_Seg1_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollection", func() {
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b").Add("c")
		c1.AddCollection(c2)
		actual := args.Map{"len": c1.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection -- merged", actual)
	})
}

func Test_Seg1_Collection_AddCollectionEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollectionEmpty", func() {
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c1.AddCollection(c2)
		actual := args.Map{"len": c1.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection empty -- no change", actual)
	})
}

func Test_Seg1_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddCollections", func() {
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b")
		empty := corestr.New.Collection.Cap(5)
		c.AddCollections(c1, empty, c2)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddCollections -- skips empty", actual)
	})
}

func Test_Seg1_Collection_LockMethods(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_LockMethods", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")
		c.AddsLock("b", "c")
		actual := args.Map{"len": c.LengthLock(), "emptyLock": c.IsEmptyLock()}
		expected := args.Map{"len": 3, "emptyLock": false}
		expected.ShouldBeEqual(t, 0, "Collection lock methods -- 3 items", actual)
	})
}

func Test_Seg1_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEquals", func() {
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("a", "b")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("a", "b")
		c3 := corestr.New.Collection.Cap(5)
		c3.Adds("a", "c")
		actual := args.Map{
			"eq":   c1.IsEquals(c2),
			"neq":  c1.IsEquals(c3),
		}
		expected := args.Map{
			"eq":   true,
			"neq":  false,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEquals -- equal and not equal", actual)
	})
}

func Test_Seg1_Collection_IsEqualsInsensitive(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEqualsInsensitive", func() {
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("Hello", "World")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("hello", "world")
		actual := args.Map{
			"sensitive":   c1.IsEqualsWithSensitive(true, c2),
			"insensitive": c1.IsEqualsWithSensitive(false, c2),
		}
		expected := args.Map{
			"sensitive":   false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEqualsWithSensitive -- case comparison", actual)
	})
}

func Test_Seg1_Collection_IsEqualsEdge(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_IsEqualsEdge", func() {
		var nilC *corestr.Collection
		emptyC := corestr.New.Collection.Cap(0)
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		actual := args.Map{
			"bothNil":   nilC.IsEquals(nil),
			"nilVsNon":  nilC.IsEquals(c),
			"emptyBoth": emptyC.IsEquals(corestr.New.Collection.Cap(0)),
			"diffLen":   c.IsEquals(emptyC),
		}
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
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.ToError("; ")
		defErr := c.ToDefaultError()
		actual := args.Map{"notNil": err != nil, "defNotNil": defErr != nil}
		expected := args.Map{"notNil": true, "defNotNil": true}
		expected.ShouldBeEqual(t, 0, "Collection ToError -- non-nil", actual)
	})
}

func Test_Seg1_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0, "c", "d")
		actual := args.Map{"origLen": c.Length(), "newLen": c2.Length()}
		expected := args.Map{"origLen": 2, "newLen": 4}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- new collection with all items", actual)
	})
}

func Test_Seg1_Collection_ConcatNewEmpty(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_ConcatNewEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0)
		actual := args.Map{"newLen": c2.Length()}
		expected := args.Map{"newLen": 2}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- empty adds returns copy", actual)
	})
}

func Test_Seg1_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_JsonString", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		s := c.JsonString()
		s2 := c.JsonStringMust()
		s3 := c.StringJSON()
		actual := args.Map{"nonEmpty": s != "", "eq": s == s2, "eq2": s == s3}
		expected := args.Map{"nonEmpty": true, "eq": true, "eq2": true}
		expected.ShouldBeEqual(t, 0, "Collection JsonString -- all variants match", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsValues(h)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues -- 2 values added", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsValuesNil(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsValuesNil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues nil -- no items", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsKeys(h)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys -- 2 keys added", actual)
	})
}

func Test_Seg1_Collection_AddHashmapsKeysNil(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddHashmapsKeysNil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys nil -- no items", actual)
	})
}

func Test_Seg1_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Seg1_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("a")
		c.AddPointerCollectionsLock(c2)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddPointerCollectionsLock -- 1 item", actual)
	})
}
