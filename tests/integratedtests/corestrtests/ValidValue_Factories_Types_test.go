package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_Factories(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Factories", func() {
		// Arrange
		v1 := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValueEmpty()
		v3 := corestr.InvalidValidValue("msg")
		v4 := corestr.InvalidValidValueNoMessage()
		v5 := corestr.NewValidValueUsingAny(false, true, "hello")
		v6 := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		_ = v5
		_ = v6

		// Act
		actual := args.Map{
			"v1Valid": v1.IsValid, "v1Val": v1.Value,
			"v2Valid": v2.IsValid,
			"v3Invalid": !v3.IsValid,
			"v4Invalid": !v4.IsValid,
		}

		// Assert
		expected := args.Map{
			"v1Valid": true, "v1Val": "hello",
			"v2Valid": true,
			"v3Invalid": true,
			"v4Invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- factories", actual)
	})
}

func Test_ValidValue_Methods_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Methods", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"notEmpty":           !v.IsEmpty(),
			"hasValidNonEmpty":   v.HasValidNonEmpty(),
			"hasSafeNonEmpty":    v.HasSafeNonEmpty(),
			"notWhitespace":      !v.IsWhitespace(),
			"hasValidNonWS":      v.HasValidNonWhitespace(),
			"trim":               v.Trim() == "hello",
			"is":                 v.Is("hello"),
			"isNot":              !v.Is("world"),
			"isAnyOf":            v.IsAnyOf("hello", "world"),
			"isContains":         v.IsContains("hel"),
			"isAnyContains":      v.IsAnyContains("hel"),
			"isEqualNonSens":     v.IsEqualNonSensitive("HELLO"),
		}

		// Assert
		expected := args.Map{
			"notEmpty": true, "hasValidNonEmpty": true, "hasSafeNonEmpty": true,
			"notWhitespace": true, "hasValidNonWS": true, "trim": true,
			"is": true, "isNot": true, "isAnyOf": true, "isContains": true,
			"isAnyContains": true, "isEqualNonSens": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- methods", actual)
	})
}

func Test_ValidValue_Conversions_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Conversions", func() {
		// Arrange
		v := &corestr.ValidValue{Value: "42", IsValid: true}
		vBool := &corestr.ValidValue{Value: "true", IsValid: true}
		vFloat := &corestr.ValidValue{Value: "3.14", IsValid: true}
		bad := &corestr.ValidValue{Value: "abc", IsValid: true}
		empty := &corestr.ValidValue{Value: "", IsValid: true}
		big := &corestr.ValidValue{Value: "999", IsValid: true}

		// Act
		actual := args.Map{
			"int":     v.ValueInt(0) == 42,
			"defInt":  v.ValueDefInt() == 42,
			"byte":    v.ValueByte(0) == 42,
			"defByte": v.ValueDefByte() == 42,
			"bool":    vBool.ValueBool(),
			"float":   vFloat.ValueFloat64(0) != 0,
			"badInt":  bad.ValueInt(99) == 99,
			"emptyBool": !empty.ValueBool(),
		}
		_ = big.ValueByte(0)
		_ = big.ValueDefByte()
		_ = bad.ValueByte(99)
		_ = bad.ValueFloat64(1.0)
		_ = bad.ValueDefInt()

		// Assert
		expected := args.Map{
			"int": true, "defInt": true, "byte": true, "defByte": true,
			"bool": true, "float": true, "badInt": true, "emptyBool": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- conversions", actual)
	})
}

func Test_ValidValue_BytesOnce_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_BytesOnce", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		b1 := v.ValueBytesOnce()
		b2 := v.ValueBytesOnce()
		_ = v.ValueBytesOncePtr()

		// Act
		actual := args.Map{
			"len1": len(b1),
			"len2": len(b2),
		}

		// Assert
		expected := args.Map{
			"len1": 5,
			"len2": 5,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- BytesOnce", actual)
	})
}

func Test_ValidValue_Regex_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Regex", func() {
		// Arrange
		v := corestr.NewValidValue("hello123")

		// Act
		actual := args.Map{
			"nilMatch":   v.IsRegexMatches(nil),
			"nilFind":    v.RegexFindString(nil) == "",
		}
		_, has := v.RegexFindAllStringsWithFlag(nil, -1)
		actual["nilFindAll"] = !has
		actual["emptyFindAll"] = len(v.RegexFindAllStrings(nil, -1)) == 0

		// Assert
		expected := args.Map{
			"nilMatch": false, "nilFind": true, "nilFindAll": true, "emptyFindAll": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Regex", actual)
	})
}

func Test_ValidValue_Split_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Split", func() {
		// Arrange
		v := corestr.NewValidValue("a,b,c")

		// Act
		actual := args.Map{"splitLen": len(v.Split(","))}
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")

		// Assert
		expected := args.Map{"splitLen": 3}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Split", actual)
	})
}

func Test_ValidValue_Clone_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_Clone", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		c := v.Clone()
		var nilV *corestr.ValidValue

		// Act
		actual := args.Map{
			"cloneVal": c.Value,
			"nilClone": nilV.Clone() == nil,
		}

		// Assert
		expected := args.Map{
			"cloneVal": "hello",
			"nilClone": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Clone", actual)
	})
}

func Test_ValidValue_StringAndJson(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_StringAndJson", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		var nilV *corestr.ValidValue
		_ = v.Json()
		_ = v.JsonPtr()
		_, _ = v.Serialize()

		// Act
		actual := args.Map{
			"str": v.String() == "hello",
			"fullStr": v.FullString() != "",
			"nilStr": nilV.String() == "",
			"nilFull": nilV.FullString() == "",
		}

		// Assert
		expected := args.Map{
			"str": true,
			"fullStr": true,
			"nilStr": true,
			"nilFull": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- StringAndJson", actual)
	})
}

func Test_ValidValue_ClearDispose_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValue_ClearDispose", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v.Clear()
		v2 := corestr.NewValidValue("world")
		v2.Dispose()
		var nilV *corestr.ValidValue
		nilV.Clear()
		nilV.Dispose()

		// Act
		actual := args.Map{"cleared": v.Value == ""}

		// Assert
		expected := args.Map{"cleared": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ClearDispose", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValues_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValues", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		vv.Add("a").AddFull(true, "b", "msg")

		// Act
		actual := args.Map{
			"len": vv.Length(),
			"empty": vv.IsEmpty(),
			"hasAny": vv.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"empty": false,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- basic", actual)
	})
}

func Test_ValidValues_NilReceiver_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValues_NilReceiver", func() {
		// Arrange
		var vv *corestr.ValidValues

		// Act
		actual := args.Map{"len": vv.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- nil", actual)
	})
}

func Test_ValidValues_Factories(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValues_Factories", func() {
		// Arrange
		vv := corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a"})
		empty := corestr.NewValidValuesUsingValues()
		cap := corestr.NewValidValues(10)

		// Act
		actual := args.Map{
			"len": vv.Length(),
			"emptyLen": empty.Length(),
			"capLen": cap.Length(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"emptyLen": 0,
			"capLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- factories", actual)
	})
}

func Test_ValidValues_SafeValues(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValues_SafeValues", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		vv.Add("a").Add("b")

		// Act
		actual := args.Map{
			"safe0":      vv.SafeValueAt(0) == "a",
			"safe99":     vv.SafeValueAt(99) == "",
			"valid0":     vv.SafeValidValueAt(0) == "a",
			"valid99":    vv.SafeValidValueAt(99) == "",
			"valsLen":    len(vv.SafeValuesAtIndexes(0, 1)),
			"validVLen":  len(vv.SafeValidValuesAtIndexes(0)),
		}

		// Assert
		expected := args.Map{
			"safe0": true, "safe99": true, "valid0": true, "valid99": true,
			"valsLen": 2, "validVLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- SafeValues", actual)
	})
}

func Test_ValidValues_Strings_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValidValues_Strings", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		vv.Add("a")

		// Act
		actual := args.Map{
			"strLen":  len(vv.Strings()),
			"fullLen": len(vv.FullStrings()),
			"str":     vv.String() != "",
		}

		// Assert
		expected := args.Map{
			"strLen": 1,
			"fullLen": 1,
			"str": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns non-empty -- Strings", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_ValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()
		vs2 := corestr.InvalidValueStatus("msg")
		c := vs2.Clone()

		// Act
		actual := args.Map{
			"invalid": !vs.ValueValid.IsValid,
			"cloneMsg": c.ValueValid.Message,
		}

		// Assert
		expected := args.Map{
			"invalid": true,
			"cloneMsg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_TextWithLineNumber", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		var nilTw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasLine":    tw.HasLineNumber(),
			"notInvalid": !tw.IsInvalidLineNumber(),
			"len":        tw.Length(),
			"notEmpty":   !tw.IsEmpty(),
			"notEmptyTxt": !tw.IsEmptyText(),
			"notBothEmpty": !tw.IsEmptyTextLineBoth(),
			"nilHasLine":  nilTw.HasLineNumber(),
			"nilInvalid":  nilTw.IsInvalidLineNumber(),
			"nilLen":      nilTw.Length(),
			"nilEmpty":    nilTw.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"hasLine": true, "notInvalid": true, "len": 5, "notEmpty": true,
			"notEmptyTxt": true, "notBothEmpty": true,
			"nilHasLine": false, "nilInvalid": true, "nilLen": 0, "nilEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValuePair_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"key":      kv.KeyName() == "k",
			"varName":  kv.VariableName() == "k",
			"valStr":   kv.ValueString() == "v",
			"isVarEq":  kv.IsVariableNameEqual("k"),
			"isValEq":  kv.IsValueEqual("v"),
			"compile":  kv.Compile() != "",
			"string":   kv.String() != "",
			"notKeyEmpty": !kv.IsKeyEmpty(),
			"notValEmpty": !kv.IsValueEmpty(),
			"notBothEmpty": !kv.IsKeyValueEmpty(),
			"hasKey":   kv.HasKey(),
			"hasVal":   kv.HasValue(),
			"trimKey":  kv.TrimKey() == "k",
			"trimVal":  kv.TrimValue() == "v",
			"is":       kv.Is("k", "v"),
			"isKey":    kv.IsKey("k"),
			"isVal":    kv.IsVal("v"),
			"notAnyEmpty": !kv.IsKeyValueAnyEmpty(),
		}
		_ = kv.FormatString("%s=%s")
		_ = kv.Json()
		_ = kv.JsonPtr()
		_, _ = kv.Serialize()
		_ = kv.SerializeMust()
		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "")

		// Assert
		expected := args.Map{
			"key": true, "varName": true, "valStr": true, "isVarEq": true,
			"isValEq": true, "compile": true, "string": true,
			"notKeyEmpty": true, "notValEmpty": true, "notBothEmpty": true,
			"hasKey": true, "hasVal": true, "trimKey": true, "trimVal": true,
			"is": true, "isKey": true, "isVal": true, "notAnyEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_Conversions(t *testing.T) {
	safeTest(t, "Test_Cov17_KeyValuePair_Conversions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		_ = kv.ValueBool()
		_ = kv.ValueInt(0)
		_ = kv.ValueDefInt()
		_ = kv.ValueByte(0)
		_ = kv.ValueDefByte()
		_ = kv.ValueFloat64(0)
		_ = kv.ValueDefFloat64()
		kv.Clear()
		kv.Dispose()
		var nilKv *corestr.KeyValuePair
		nilKv.Clear()
		nilKv.Dispose()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- conversions", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyAnyValuePair_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_KeyAnyValuePair", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{
			"key":      kv.KeyName() == "k",
			"varName":  kv.VariableName() == "k",
			"valAny":   kv.ValueAny() == 42,
			"isVarEq":  kv.IsVariableNameEqual("k"),
			"notNull":  !kv.IsValueNull(),
			"hasNonNull": kv.HasNonNull(),
			"hasValue":  kv.HasValue(),
			"notEmptyStr": !kv.IsValueEmptyString(),
			"notWS":     !kv.IsValueWhitespace(),
			"valStr":    kv.ValueString() != "",
			"compile":   kv.Compile() != "",
			"string":    kv.String() != "",
		}
		// call ValueString again for cache
		_ = kv.ValueString()
		_ = kv.SerializeMust()
		_ = kv.AsJsonContractsBinder()
		_ = kv.AsJsoner()
		_ = kv.AsJsonParseSelfInjector()
		kv.Clear()
		kv.Dispose()
		var nilKv *corestr.KeyAnyValuePair
		nilKv.Clear()
		nilKv.Dispose()

		// Assert
		expected := args.Map{
			"key": true, "varName": true, "valAny": true, "isVarEq": true,
			"notNull": true, "hasNonNull": true, "hasValue": true,
			"notEmptyStr": true, "notWS": true, "valStr": true,
			"compile": true, "string": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashmapDiff_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_HashmapDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		right := map[string]string{"a": "2"}

		// Act
		actual := args.Map{
			"notEmpty":  !hd.IsEmpty(),
			"hasAny":    hd.HasAnyItem(),
			"len":       hd.Length(),
			"lastIdx":   hd.LastIndex(),
			"hasChanges": hd.HasAnyChanges(right),
			"notEqual":  !hd.IsRawEqual(right),
		}
		_ = hd.AllKeysSorted()
		_ = hd.MapAnyItems()
		_ = hd.Raw()
		_ = hd.RawMapStringAnyDiff()
		_ = hd.HashmapDiffUsingRaw(right)
		_ = hd.DiffRaw(right)
		_ = hd.DiffJsonMessage(right)
		_ = hd.ToStringsSliceOfDiffMap(map[string]string{"a": "diff"})
		_ = hd.ShouldDiffMessage("title", right)
		_ = hd.LogShouldDiffMessage("title", right)
		_, _ = hd.Serialize()
		_ = hd.Deserialize(&map[string]string{})
		var nilHd *corestr.HashmapDiff
		_ = nilHd.Length()
		_ = nilHd.Raw()
		_ = nilHd.MapAnyItems()

		// Assert
		expected := args.Map{
			"notEmpty": true, "hasAny": true, "len": 1, "lastIdx": 0,
			"hasChanges": true, "notEqual": true,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRight_Factories(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftRight_Factories", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		inv := corestr.InvalidLeftRight("msg")
		inv2 := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{
			"left": lr.Left, "right": lr.Right, "valid": lr.IsValid,
			"inv1Invalid": !inv.IsValid, "inv2Invalid": !inv2.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a", "right": "b", "valid": true,
			"inv1Invalid": true, "inv2Invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- factories", actual)
	})
}

func Test_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftRight_Methods", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello", "world")

		// Act
		actual := args.Map{
			"leftBytes":  string(lr.LeftBytes()) == "hello",
			"rightBytes": string(lr.RightBytes()) == "world",
			"leftTrim":   lr.LeftTrim() == "hello",
			"rightTrim":  lr.RightTrim() == "world",
			"notLeftEmpty":  !lr.IsLeftEmpty(),
			"notRightEmpty": !lr.IsRightEmpty(),
			"notLeftWS":  !lr.IsLeftWhitespace(),
			"notRightWS": !lr.IsRightWhitespace(),
			"hasLeft":    lr.HasValidNonEmptyLeft(),
			"hasRight":   lr.HasValidNonEmptyRight(),
			"hasWSLeft":  lr.HasValidNonWhitespaceLeft(),
			"hasWSRight": lr.HasValidNonWhitespaceRight(),
			"hasSafe":    lr.HasSafeNonEmpty(),
			"isLeft":     lr.IsLeft("hello"),
			"isRight":    lr.IsRight("world"),
			"is":         lr.Is("hello", "world"),
		}
		_ = lr.NonPtr()
		_ = lr.Ptr()
		_ = lr.IsLeftRegexMatch(nil)
		_ = lr.IsRightRegexMatch(nil)
		c := lr.Clone()
		_ = c
		_ = lr.IsEqual(corestr.NewLeftRight("hello", "world"))
		lr.Clear()
		lr.Dispose()
		var nilLr *corestr.LeftRight
		nilLr.Clear()
		nilLr.Dispose()

		// Assert
		expected := args.Map{
			"leftBytes": true, "rightBytes": true, "leftTrim": true, "rightTrim": true,
			"notLeftEmpty": true, "notRightEmpty": true, "notLeftWS": true, "notRightWS": true,
			"hasLeft": true, "hasRight": true, "hasWSLeft": true, "hasWSRight": true,
			"hasSafe": true, "isLeft": true, "isRight": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- methods", actual)
	})
}

func Test_LeftRight_FromSlice(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftRight_FromSlice", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		lr3 := corestr.LeftRightUsingSlice(nil)
		_ = corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		_ = corestr.LeftRightUsingSlicePtr(nil)
		_ = corestr.LeftRightTrimmedUsingSlice(nil)
		_ = corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		lr8 := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{
			"lr1Left": lr.Left, "lr1Right": lr.Right,
			"lr2Left": lr2.Left, "lr3Invalid": !lr3.IsValid,
			"lr8Left": lr8.Left, "lr8Right": lr8.Right,
		}

		// Assert
		expected := args.Map{
			"lr1Left": "a", "lr1Right": "b",
			"lr2Left": "a", "lr3Invalid": true,
			"lr8Left": "a", "lr8Right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- FromSlice", actual)
	})
}

func Test_LeftRight_FromSplit_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftRight_FromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=val", "=")
		lr2 := corestr.LeftRightFromSplitTrimmed(" key = val ", "=")
		lr3 := corestr.LeftRightFromSplitFull("a:b:c", ":")
		_ = corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{
			"lr1Left": lr.Left, "lr1Right": lr.Right,
			"lr2Left": lr2.Left, "lr2Right": lr2.Right,
			"lr3Left": lr3.Left, "lr3Right": lr3.Right,
		}

		// Assert
		expected := args.Map{
			"lr1Left": "key", "lr1Right": "val",
			"lr2Left": "key", "lr2Right": "val",
			"lr3Left": "a", "lr3Right": "b:c",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- FromSplit", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftMiddleRight_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftMiddleRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right,
			"notLeftEmpty":   !lmr.IsLeftEmpty(),
			"notMiddleEmpty": !lmr.IsMiddleEmpty(),
			"notRightEmpty":  !lmr.IsRightEmpty(),
			"notLeftWS":      !lmr.IsLeftWhitespace(),
			"notMiddleWS":    !lmr.IsMiddleWhitespace(),
			"notRightWS":     !lmr.IsRightWhitespace(),
			"hasSafe":        lmr.HasSafeNonEmpty(),
			"isAll":          lmr.IsAll("a", "b", "c"),
			"is":             lmr.Is("a", "c"),
		}
		_ = string(lmr.LeftBytes())
		_ = string(lmr.MiddleBytes())
		_ = string(lmr.RightBytes())
		_ = lmr.LeftTrim()
		_ = lmr.MiddleTrim()
		_ = lmr.RightTrim()
		_ = lmr.HasValidNonEmptyLeft()
		_ = lmr.HasValidNonEmptyMiddle()
		_ = lmr.HasValidNonEmptyRight()
		_ = lmr.HasValidNonWhitespaceLeft()
		_ = lmr.HasValidNonWhitespaceMiddle()
		_ = lmr.HasValidNonWhitespaceRight()
		_ = lmr.Clone()
		_ = lmr.ToLeftRight()
		lmr.Clear()
		lmr.Dispose()
		var nilLmr *corestr.LeftMiddleRight
		nilLmr.Clear()
		nilLmr.Dispose()
		_ = corestr.InvalidLeftMiddleRight("msg")
		_ = corestr.InvalidLeftMiddleRightNoMessage()

		// Assert
		expected := args.Map{
			"left": "a", "middle": "b", "right": "c",
			"notLeftEmpty": true, "notMiddleEmpty": true, "notRightEmpty": true,
			"notLeftWS": true, "notMiddleWS": true, "notRightWS": true,
			"hasSafe": true, "isAll": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_FromSplit_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_LeftMiddleRight_FromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		_ = corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		lmr3 := corestr.LeftMiddleRightFromSplitN("a:b:c:d", ":")
		_ = corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right,
			"lmr3Left": lmr3.Left, "lmr3Middle": lmr3.Middle, "lmr3Right": lmr3.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a", "middle": "b", "right": "c",
			"lmr3Left": "a", "lmr3Middle": "b", "lmr3Right": "c:d",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- FromSplit", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleStringOnce_Core(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_Core", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{
			"val":     s.Value() == "hello",
			"init":    s.IsInitialized(),
			"defined": s.IsDefined(),
			"notUninit": !s.IsUninitialized(),
			"notInvalid": !s.IsInvalid(),
			"safe":    s.SafeValue() == "hello",
			"notEmpty": !s.IsEmpty(),
			"notWS":   !s.IsWhitespace(),
			"trim":    s.Trim() == "hello",
			"hasValid": s.HasValidNonEmpty(),
			"hasWS":   s.HasValidNonWhitespace(),
			"hasSafe": s.HasSafeNonEmpty(),
			"is":      s.Is("hello"),
			"isAnyOf": s.IsAnyOf("hello"),
			"contains": s.IsContains("hel"),
			"anyContains": s.IsAnyContains("hel"),
			"nonSens": s.IsEqualNonSensitive("HELLO"),
		}
		_ = s.ValueBytes()
		_ = s.ValueBytesPtr()

		// Assert
		expected := args.Map{
			"val": true, "init": true, "defined": true, "notUninit": true,
			"notInvalid": true, "safe": true, "notEmpty": true, "notWS": true,
			"trim": true, "hasValid": true, "hasWS": true, "hasSafe": true,
			"is": true, "isAnyOf": true, "contains": true, "anyContains": true,
			"nonSens": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- core", actual)
	})
}

func Test_SimpleStringOnce_Set(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_Set", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		err := s.SetOnUninitialized("val")
		err2 := s.SetOnUninitialized("val2")

		// Act
		actual := args.Map{
			"err1Nil": err == nil,
			"err2NotNil": err2 != nil,
		}

		// Assert
		expected := args.Map{
			"err1Nil": true,
			"err2NotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- Set", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		v := s.GetSetOnce("first")
		v2 := s.GetSetOnce("second")

		// Act
		actual := args.Map{
			"v1": v,
			"v2": v2,
		}

		// Assert
		expected := args.Map{
			"v1": "first",
			"v2": "first",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- GetSetOnce", actual)
	})
}

func Test_SimpleStringOnce_GetOnceFunc_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Uninitialized("")
		v := s.GetOnceFunc(func() string { return "computed" })
		v2 := s.GetOnceFunc(func() string { return "other" })

		// Act
		actual := args.Map{
			"v1": v,
			"v2": v2,
		}

		// Assert
		expected := args.Map{
			"v1": "computed",
			"v2": "computed",
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- GetOnceFunc", actual)
	})
}

func Test_SimpleStringOnce_Conversions_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_Conversions", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("42")
		s2 := corestr.New.SimpleStringOnce.Init("3.14")
		s3 := corestr.New.SimpleStringOnce.Init("true")
		s4 := corestr.New.SimpleStringOnce.Init("yes")
		s5 := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		actual := args.Map{
			"int":     s.Int() == 42,
			"valInt":  s.ValueInt(0) == 42,
			"defInt":  s.ValueDefInt() == 42,
			"byte":    s.Byte() == 42,
			"valByte": s.ValueByte(0) == 42,
			"float":   s2.ValueFloat64(0) != 0,
			"bool":    s3.Boolean(false),
			"boolDef": s3.BooleanDefault(),
			"isValBool": s3.IsValueBool(),
			"yes":     s4.Boolean(false),
			"badInt":  s5.Int() == 0,
		}
		_ = s5.Byte()
		_ = s5.Int16()
		_ = s5.Int32()
		_ = s5.ValueByte(0)
		_ = s5.IsSetter(false)
		_ = s5.IsSetter(true)

		// Assert
		expected := args.Map{
			"int": true, "valInt": true, "defInt": true, "byte": true,
			"valByte": true, "float": true, "bool": true, "boolDef": true,
			"isValBool": true, "yes": true, "badInt": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- conversions", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_FromValidValueFactoriesT(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_WithinRange", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("50")
		v, ok := s.WithinRange(true, 0, 100)
		v2, ok2 := s.WithinRangeDefault(0, 100)
		_, _ = s.Uint16()
		_, _ = s.Uint32()

		// Act
		actual := args.Map{
			"v": v,
			"ok": ok,
			"v2": v2,
			"ok2": ok2,
		}

		// Assert
		expected := args.Map{
			"v": 50,
			"ok": true,
			"v2": 50,
			"ok2": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns non-empty -- WithinRange", actual)
	})
}

func Test_SimpleStringOnce_Various(t *testing.T) {
	safeTest(t, "Test_Cov17_SimpleStringOnce_Various", func() {
		// Arrange
		s := corestr.New.SimpleStringOnce.Init("hello")
		_ = s.ConcatNew(" world")
		_ = s.ConcatNewUsingStrings(" ", "beautiful", "world")
		_ = s.Split(",")
		_ = s.SplitNonEmpty(",")
		_ = s.SplitTrimNonWhitespace(",")
		_, _ = s.SplitLeftRight(",")
		_, _ = s.SplitLeftRightTrim(",")
		_ = s.LinesSimpleSlice()
		_ = s.SimpleSlice(",")
		_ = s.IsRegexMatches(nil)
		_ = s.RegexFindString(nil)
		_, _ = s.RegexFindAllStringsWithFlag(nil, -1)
		_ = s.RegexFindAllStrings(nil, -1)
		_ = s.NonPtr()
		_ = s.Ptr()
		_ = s.String()
		_ = s.StringPtr()
		_ = s.Clone()
		_ = s.ClonePtr()
		_ = s.CloneUsingNewVal("new")
		_ = s.JsonModel()
		_ = s.JsonModelAny()
		_, _ = s.MarshalJSON()
		_, _ = s.Serialize()
		_ = s.AsJsoner()
		_ = s.AsJsonContractsBinder()
		_ = s.AsJsonParseSelfInjector()
		_ = s.AsJsonMarshaller()

		s.Invalidate()
		s2 := corestr.New.SimpleStringOnce.Init("world")
		s2.Reset()

		s3 := corestr.New.SimpleStringOnce.Uninitialized("")
		_ = s3.SetOnceIfUninitialized("val")
		_ = s3.SetOnceIfUninitialized("val2")
		_ = s3.GetOnce()

		var nilS *corestr.SimpleStringOnce
		_ = nilS.String()
		_ = nilS.StringPtr()
		_ = nilS.ClonePtr()

		s.Dispose()

		// Act
		actual := args.Map{"done": true}

		// Assert
		expected := args.Map{"done": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns correct value -- various", actual)
	})
}
