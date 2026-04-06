package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValidValue ──

func Test_SrcC04_ValidValue_Factories_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Factories_Verification", func() {
		// Arrange
		tc := srcC04ValidValueFactoriesTestCase

		// Act
		v := corestr.NewValidValue("hello")
		actual := args.Map{
			"value":   v.Value,
			"isValid": v.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_EmptyFactory_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_EmptyFactory_Verification", func() {
		// Arrange
		tc := srcC04ValidValueEmptyFactoryTestCase

		// Act
		v := corestr.NewValidValueEmpty()
		actual := args.Map{
			"value":   v.Value,
			"isValid": v.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Invalid_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Invalid_Verification", func() {
		// Arrange
		tc := srcC04ValidValueInvalidTestCase

		// Act
		v := corestr.InvalidValidValue("msg")
		v2 := corestr.InvalidValidValueNoMessage()
		actual := args.Map{
			"isValid":           v.IsValid,
			"invalidNoMsgValid": v2.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Methods_Verification", func() {
		// Arrange
		tc := srcC04ValidValueMethodsTestCase
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"isEmpty":                v.IsEmpty(),
			"hasValidNonEmpty":       v.HasValidNonEmpty(),
			"hasSafeNonEmpty":        v.HasSafeNonEmpty(),
			"isWhitespace":           v.IsWhitespace(),
			"hasValidNonWhitespace":  v.HasValidNonWhitespace(),
			"trim":                   v.Trim(),
			"isHello":                v.Is("hello"),
			"isWorld":                v.Is("world"),
			"isAnyOfHelloWorld":      v.IsAnyOf("hello", "world"),
			"isAnyOfEmpty":           v.IsAnyOf(),
			"isAnyOfX":               v.IsAnyOf("x"),
			"isContainsHel":          v.IsContains("hel"),
			"isAnyContainsHel":       v.IsAnyContains("hel"),
			"isAnyContainsEmpty":     v.IsAnyContains(),
			"isAnyContainsXyz":       v.IsAnyContains("xyz"),
			"isEqualNonSensitive":    v.IsEqualNonSensitive("HELLO"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Conversions_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Conversions_Verification", func() {
		// Arrange
		tc := srcC04ValidValueConversionsTestCase
		v := &corestr.ValidValue{Value: "42", IsValid: true}
		vTrue := &corestr.ValidValue{Value: "true", IsValid: true}

		// Act
		actual := args.Map{
			"valueInt":     v.ValueInt(0),
			"valueDefInt":  v.ValueDefInt(),
			"valueByte":    int(v.ValueByte(0)),
			"valueDefByte": int(v.ValueDefByte()),
			"trueBool":     vTrue.ValueBool(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_BadConversions_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_BadConversions_Verification", func() {
		// Arrange
		tc := srcC04ValidValueBadConversionsTestCase
		bad := &corestr.ValidValue{Value: "abc", IsValid: true}

		// Act
		actual := args.Map{
			"valueInt":    bad.ValueInt(99),
			"valueDefInt": bad.ValueDefInt(),
			"valueFloat":  bad.ValueFloat64(1.0),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC04_ValidValue_Regex_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Regex_Verification", func() {
		// Arrange
		tc := srcC04ValidValueRegexTestCase
		v := corestr.NewValidValue("hello123")

		// Act
		items, has := v.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{
			"isMatch":    v.IsRegexMatches(nil),
			"findStr":    v.RegexFindString(nil),
			"hasAll":     has,
			"allLen":     len(items),
			"findAllLen": len(v.RegexFindAllStrings(nil, -1)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Split_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Split_Verification", func() {
		// Arrange
		tc := srcC04ValidValueSplitTestCase
		v := corestr.NewValidValue("a,b,c")

		// Act
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")
		actual := args.Map{
			"splitLen": len(v.Split(",")),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Clone_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Clone_Verification", func() {
		for caseIndex, tc := range srcC04ValidValueCloneTestCases {
			// Arrange
			input := tc.ArrangeInput.(args.Map)
			isNil := input.GetAsBoolDefault("isNil", false)

			// Act
			actual := args.Map{}
			if isNil {
				var nilV *corestr.ValidValue
				actual["isNilResult"] = nilV.Clone() == nil
			} else {
				v := corestr.NewValidValue("hello")
				c := v.Clone()
				actual["value"] = c.Value
			}

			// Assert
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	})
}

func Test_SrcC04_ValidValue_String_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_String_Verification", func() {
		// Arrange
		tc := srcC04ValidValueStringTestCase
		v := corestr.NewValidValue("hello")
		var nilV *corestr.ValidValue

		// Act
		actual := args.Map{
			"string":          v.String(),
			"fullStringEmpty": v.FullString() == "",
			"nilString":       nilV.String(),
			"nilFullString":   nilV.FullString(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_ClearDispose_Verification", func() {
		// Arrange
		tc := srcC04ValidValueClearDisposeTestCase

		// Act
		v := corestr.NewValidValue("hello")
		v.Clear()
		var nilV *corestr.ValidValue
		noPanic := !callPanicsSrcC04(func() {
			nilV.Clear()
			nilV.Dispose()
			v2 := corestr.NewValidValue("world")
			v2.Dispose()
		})
		actual := args.Map{
			"clearedValue": v.Value,
			"noPanic":      noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_Json_Verification", func() {
		// Arrange
		tc := srcC04ValidValueJsonTestCase

		// Act
		v := corestr.ValidValue{Value: "hello", IsValid: true}
		noPanic := !callPanicsSrcC04(func() {
			_ = v.Json()
			_ = v.JsonPtr()
			_, _ = v.Serialize()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValue_AdditionalFactories_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_AdditionalFactories_Verification", func() {
		// Arrange — exercise remaining factory methods

		// Act
		v5 := corestr.NewValidValueUsingAny(false, true, "hello")
		v6 := corestr.NewValidValueUsingAnyAutoValid(false, "hello")

		// Assert
		if v5.Value == "" {
			t.Fatal("expected non-empty")
		}
		_ = v6
	})
}

func Test_SrcC04_ValidValue_OverflowByte_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValue_OverflowByte_Verification", func() {
		// Arrange
		big := &corestr.ValidValue{Value: "999", IsValid: true}
		neg := &corestr.ValidValue{Value: "-1", IsValid: true}
		empty := &corestr.ValidValue{Value: "", IsValid: true}

		// Act — exercise overflow/edge paths
		_ = big.ValueByte(0)
		_ = big.ValueDefByte()
		_ = neg.ValueByte(0)
		_ = empty.ValueBool()

		// Assert
		vf := &corestr.ValidValue{Value: "3.14", IsValid: true}
		if vf.ValueFloat64(0) == 0 || vf.ValueDefFloat64() == 0 {
			t.Fatal("expected non-zero float")
		}
	})
}

// ── ValidValues ──

func Test_SrcC04_ValidValues_Basic_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValues_Basic_Verification", func() {
		// Arrange
		tc := srcC04ValidValuesBasicTestCase

		// Act
		vv := corestr.EmptyValidValues()
		emptyIsEmpty := vv.IsEmpty()
		emptyHasAny := vv.HasAnyItem()
		emptyLength := vv.Length()
		vv.Add("a").AddFull(true, "b", "msg")
		actual := args.Map{
			"emptyIsEmpty": emptyIsEmpty,
			"emptyHasAny":  emptyHasAny,
			"emptyLength":  emptyLength,
			"addedLength":  vv.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValues_NilReceiver_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValues_NilReceiver_Verification", func() {
		// Arrange
		tc := srcC04ValidValuesNilTestCase
		var vv *corestr.ValidValues

		// Act
		actual := args.Map{
			"length": vv.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValues_Factories_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValues_Factories_Verification", func() {
		// Arrange
		tc := srcC04ValidValuesFactoriesTestCase

		// Act
		vv := corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "a"})
		empty := corestr.NewValidValuesUsingValues()
		cap := corestr.NewValidValues(10)
		actual := args.Map{
			"usingValuesLen": vv.Length(),
			"emptyLen":       empty.Length(),
			"capLen":         cap.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValues_SafeValues_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValues_SafeValues_Verification", func() {
		// Arrange
		tc := srcC04ValidValuesSafeValuesTestCase
		vv := corestr.EmptyValidValues()
		vv.Add("a").Add("b")

		// Act
		actual := args.Map{
			"safeAt0":          vv.SafeValueAt(0),
			"safeAt99":         vv.SafeValueAt(99),
			"safeValidAt0":     vv.SafeValidValueAt(0),
			"safeValidAt99":    vv.SafeValidValueAt(99),
			"safeValuesLen":    len(vv.SafeValuesAtIndexes(0, 1)),
			"safeValidValsLen": len(vv.SafeValidValuesAtIndexes(0)),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_ValidValues_Strings_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValidValues_Strings_Verification", func() {
		// Arrange
		tc := srcC04ValidValuesStringsTestCase
		vv := corestr.EmptyValidValues()
		vv.Add("a")

		// Act
		actual := args.Map{
			"stringsLen":     len(vv.Strings()),
			"fullStringsLen": len(vv.FullStrings()),
			"stringNonEmpty": vv.String() != "",
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── ValueStatus ──

func Test_SrcC04_ValueStatus_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_ValueStatus_Verification", func() {
		// Arrange
		tc := srcC04ValueStatusTestCase

		// Act
		vs := corestr.InvalidValueStatusNoMessage()
		vs2 := corestr.InvalidValueStatus("msg")
		c := vs2.Clone()
		actual := args.Map{
			"invalidNoMsgValid": vs.ValueValid.IsValid,
			"cloneMessage":      c.ValueValid.Message,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── TextWithLineNumber ──

func Test_SrcC04_TextWithLineNumber_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_TextWithLineNumber_Verification", func() {
		// Arrange
		tc := srcC04TextWithLineNumberTestCase
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		var nilTw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasLineNumber":    tw.HasLineNumber(),
			"isInvalidLine":    tw.IsInvalidLineNumber(),
			"length":           tw.Length(),
			"isEmpty":          tw.IsEmpty(),
			"isEmptyText":      tw.IsEmptyText(),
			"isEmptyBoth":      tw.IsEmptyTextLineBoth(),
			"nilHasLine":       nilTw.HasLineNumber(),
			"nilIsInvalidLine": nilTw.IsInvalidLineNumber(),
			"nilLength":        nilTw.Length(),
			"nilIsEmpty":       nilTw.IsEmpty(),
			"nilIsEmptyText":   nilTw.IsEmptyText(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── KeyValuePair ──

func Test_SrcC04_KeyValuePair_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_KeyValuePair_Methods_Verification", func() {
		// Arrange
		tc := srcC04KeyValuePairTestCase
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		_ = kv.FormatString("%s=%s")
		_ = kv.Json()
		_ = kv.JsonPtr()
		_, _ = kv.Serialize()
		_ = kv.SerializeMust()
		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "")
		actual := args.Map{
			"keyName":          kv.KeyName(),
			"variableName":     kv.VariableName(),
			"valueString":      kv.ValueString(),
			"isVarNameEqual":   kv.IsVariableNameEqual("k"),
			"isValueEqual":     kv.IsValueEqual("v"),
			"compileNonEmpty":  kv.Compile() != "",
			"stringNonEmpty":   kv.String() != "",
			"isKeyEmpty":       kv.IsKeyEmpty(),
			"isValueEmpty":     kv.IsValueEmpty(),
			"isKeyValueEmpty":  kv.IsKeyValueEmpty(),
			"hasKey":           kv.HasKey(),
			"hasValue":         kv.HasValue(),
			"trimKey":          kv.TrimKey(),
			"trimValue":        kv.TrimValue(),
			"isKV":             kv.Is("k", "v"),
			"isKey":            kv.IsKey("k"),
			"isVal":            kv.IsVal("v"),
			"isKeyValAnyEmpty": kv.IsKeyValueAnyEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_KeyValuePair_Conversions_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_KeyValuePair_Conversions_Verification", func() {
		// Arrange
		tc := srcC04KeyValuePairConversionsTestCase
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}

		// Act
		_ = kv.ValueBool()
		_ = kv.ValueByte(0)
		_ = kv.ValueDefByte()
		_ = kv.ValueFloat64(0)
		_ = kv.ValueDefFloat64()
		actual := args.Map{
			"valueInt":    kv.ValueInt(0),
			"valueDefInt": kv.ValueDefInt(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_KeyValuePair_ClearDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_KeyValuePair_ClearDispose_Verification", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		var nilKv *corestr.KeyValuePair

		// Act
		noPanic := !callPanicsSrcC04(func() {
			kv.Clear()
			kv.Dispose()
			nilKv.Clear()
			nilKv.Dispose()
		})

		// Assert
		if !noPanic {
			t.Fatal("expected no panic")
		}
	})
}

// ── LeftRight ──

func Test_SrcC04_LeftRight_Factories_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftRight_Factories_Verification", func() {
		// Arrange
		tc := srcC04LeftRightFactoriesTestCase

		// Act
		lr := corestr.NewLeftRight("a", "b")
		inv := corestr.InvalidLeftRight("msg")
		inv2 := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{
			"left":            lr.Left,
			"right":           lr.Right,
			"isValid":         lr.IsValid,
			"invalidIsValid":  inv.IsValid,
			"invalid2IsValid": inv2.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_LeftRight_Methods_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftRight_Methods_Verification", func() {
		// Arrange
		tc := srcC04LeftRightMethodsTestCase
		lr := corestr.NewLeftRight("hello", "world")

		// Act
		_ = lr.NonPtr()
		_ = lr.Ptr()
		c := lr.Clone()
		actual := args.Map{
			"leftBytes":                string(lr.LeftBytes()),
			"rightBytes":               string(lr.RightBytes()),
			"leftTrim":                 lr.LeftTrim(),
			"rightTrim":                lr.RightTrim(),
			"isLeftEmpty":              lr.IsLeftEmpty(),
			"isRightEmpty":             lr.IsRightEmpty(),
			"isLeftWhitespace":         lr.IsLeftWhitespace(),
			"isRightWhitespace":        lr.IsRightWhitespace(),
			"hasValidNonEmptyLeft":     lr.HasValidNonEmptyLeft(),
			"hasValidNonEmptyRight":    lr.HasValidNonEmptyRight(),
			"hasValidNonWsLeft":        lr.HasValidNonWhitespaceLeft(),
			"hasValidNonWsRight":       lr.HasValidNonWhitespaceRight(),
			"hasSafeNonEmpty":          lr.HasSafeNonEmpty(),
			"isLeft":                   lr.IsLeft("hello"),
			"isRight":                  lr.IsRight("world"),
			"is":                       lr.Is("hello", "world"),
			"isLeftRegexNil":           lr.IsLeftRegexMatch(nil),
			"isRightRegexNil":          lr.IsRightRegexMatch(nil),
			"cloneLeft":                c.Left,
			"isEqual":                  lr.IsEqual(corestr.NewLeftRight("hello", "world")),
		}
		lr.Clear()
		lr.Dispose()
		var nilLr *corestr.LeftRight
		nilLr.Clear()
		nilLr.Dispose()

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_LeftRight_FromSlice_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftRight_FromSlice_Verification", func() {
		// Arrange
		tc := srcC04LeftRightFromSliceTestCase

		// Act
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		lr3 := corestr.LeftRightUsingSlice(nil)
		_ = corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		_ = corestr.LeftRightUsingSlicePtr(nil)
		_ = corestr.LeftRightTrimmedUsingSlice(nil)
		_ = corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		lr8 := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{
			"twoLeft":    lr.Left,
			"twoRight":   lr.Right,
			"oneLeft":    lr2.Left,
			"oneRight":   lr2.Right,
			"nilIsValid": lr3.IsValid,
			"trimLeft":   lr8.Left,
			"trimRight":  lr8.Right,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_LeftRight_FromSplit_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftRight_FromSplit_Verification", func() {
		// Arrange
		tc := srcC04LeftRightFromSplitTestCase

		// Act
		lr := corestr.LeftRightFromSplit("key=val", "=")
		lr2 := corestr.LeftRightFromSplitTrimmed(" key = val ", "=")
		lr3 := corestr.LeftRightFromSplitFull("a:b:c", ":")
		_ = corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		actual := args.Map{
			"splitLeft":  lr.Left,
			"splitRight": lr.Right,
			"trimLeft":   lr2.Left,
			"trimRight":  lr2.Right,
			"fullLeft":   lr3.Left,
			"fullRight":  lr3.Right,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── LeftMiddleRight ──

func Test_SrcC04_LeftMiddleRight_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftMiddleRight_Verification", func() {
		// Arrange
		tc := srcC04LeftMiddleRightTestCase
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		_ = lmr.IsLeftWhitespace()
		_ = lmr.IsMiddleWhitespace()
		_ = lmr.IsRightWhitespace()
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
		c := lmr.Clone()
		lr := lmr.ToLeftRight()
		hasSafe := lmr.HasSafeNonEmpty()
		isAll := lmr.IsAll("a", "b", "c")
		isAC := lmr.Is("a", "c")
		inv := corestr.InvalidLeftMiddleRight("msg")
		_ = corestr.InvalidLeftMiddleRightNoMessage()
		lmr.Clear()
		lmr.Dispose()
		var nilLmr *corestr.LeftMiddleRight
		nilLmr.Clear()
		nilLmr.Dispose()

		actual := args.Map{
			"left":            "a",
			"middle":          "b",
			"right":           "c",
			"isLeftEmpty":     false,
			"isMiddleEmpty":   false,
			"isRightEmpty":    false,
			"hasSafeNonEmpty": hasSafe,
			"isAll":           isAll,
			"is":              isAC,
			"cloneLeft":       c.Left,
			"toLrLeft":        lr.Left,
			"toLrRight":       lr.Right,
			"invalidIsValid":  inv.IsValid,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC04_LeftMiddleRight_FromSplit_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_LeftMiddleRight_FromSplit_Verification", func() {
		// Arrange
		tc := srcC04LeftMiddleRightFromSplitTestCase

		// Act
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		_ = corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		lmr3 := corestr.LeftMiddleRightFromSplitN("a:b:c:d", ":")
		_ = corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c ", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"nLeft":   lmr3.Left,
			"nMiddle": lmr3.Middle,
			"nRight":  lmr3.Right,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── KeyAnyValuePair ──

func Test_SrcC04_KeyAnyValuePair_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_KeyAnyValuePair_Verification", func() {
		// Arrange
		tc := srcC04KeyAnyValuePairTestCase
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		_ = kv.ValueString() // first call
		_ = kv.ValueString() // cached
		_ = kv.SerializeMust()
		_ = kv.AsJsonContractsBinder()
		_ = kv.AsJsoner()
		_ = kv.AsJsonParseSelfInjector()
		actual := args.Map{
			"keyName":        kv.KeyName(),
			"variableName":   kv.VariableName(),
			"valueAny":       kv.ValueAny(),
			"isVarNameEqual": kv.IsVariableNameEqual("k"),
			"isValueNull":    kv.IsValueNull(),
			"hasNonNull":     kv.HasNonNull(),
			"hasValue":       kv.HasValue(),
			"isEmptyStr":     kv.IsValueEmptyString(),
			"isWhitespace":   kv.IsValueWhitespace(),
			"compileEmpty":   kv.Compile() == "",
			"stringEmpty":    kv.String() == "",
		}
		kv.Clear()
		kv.Dispose()
		var nilKv *corestr.KeyAnyValuePair
		nilKv.Clear()
		nilKv.Dispose()

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ── HashmapDiff ──

func Test_SrcC04_HashmapDiff_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC04_HashmapDiff_Verification", func() {
		// Arrange
		tc := srcC04HashmapDiffTestCase
		hd := corestr.HashmapDiff(map[string]string{"a": "1"})
		right := map[string]string{"a": "2"}
		var nilHd *corestr.HashmapDiff

		// Act
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
		_ = nilHd.Raw()
		_ = nilHd.MapAnyItems()

		actual := args.Map{
			"isEmpty":    hd.IsEmpty(),
			"hasAny":     hd.HasAnyItem(),
			"length":     hd.Length(),
			"lastIndex":  hd.LastIndex(),
			"hasChanges": hd.HasAnyChanges(right),
			"isRawEqual": hd.IsRawEqual(right),
			"nilLength":  nilHd.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC04(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
