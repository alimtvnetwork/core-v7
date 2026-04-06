package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection basics ──

func Test_C33_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_C33_Collection_JsonString", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		js := c.JsonString()
		if js == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C33_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_C33_Collection_JsonStringMust", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.JsonStringMust()
	})
}

func Test_C33_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C33_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		empty := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.HasAnyItem(), "empty": empty.HasAnyItem()}
		expected := args.Map{"has": true, "empty": false}
		expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- with args", actual)
	})
}

func Test_C33_Collection_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C33_Collection_LastIndex_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{
			"lastIdx":  c.LastIndex(),
			"hasIdx0":  c.HasIndex(0),
			"hasIdx5":  c.HasIndex(5),
			"hasIdxN1": c.HasIndex(-1),
		}
		expected := args.Map{"lastIdx": 1, "hasIdx0": true, "hasIdx5": false, "hasIdxN1": false}
		expected.ShouldBeEqual(t, 0, "LastIndex/HasIndex returns correct value -- with args", actual)
	})
}

func Test_C33_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_C33_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.ListStrings()) != 1 { t.Fatal("expected 1") }
		if len(c.ListStringsPtr()) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C33_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_C33_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.StringJSON()
	})
}

func Test_C33_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C33_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)
		if !ok { t.Fatal("expected success") }
		if c.Length() != 2 { t.Fatal("expected 2") }
		// invalid index
		if c.RemoveAt(-1) { t.Fatal("expected false") }
		if c.RemoveAt(99) { t.Fatal("expected false") }
	})
}

func Test_C33_Collection_Count_Capacity(t *testing.T) {
	safeTest(t, "Test_C33_Collection_Count_Capacity", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.Count() != 1 { t.Fatal("expected 1") }
		_ = c.Capacity()
	})
}

func Test_C33_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C33_Collection_Length_Nil", func() {
		var nilC *corestr.Collection
		if nilC.Length() != 0 { t.Fatal("expected 0") }
	})
}

func Test_C33_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_C33_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.LengthLock() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C33_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_C33_Collection_IsEquals", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c3 := corestr.New.Collection.Strings([]string{"a", "c"})
		actual := args.Map{
			"equal":    c1.IsEquals(c2),
			"notEqual": c1.IsEquals(c3),
		}
		expected := args.Map{"equal": true, "notEqual": false}
		expected.ShouldBeEqual(t, 0, "Collection.IsEquals returns correct value -- with args", actual)
	})
}

// ── ValidValue ──

func Test_C33_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Constructors", func() {
		v1 := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValueEmpty()
		v3 := corestr.InvalidValidValue("err")
		v4 := corestr.InvalidValidValueNoMessage()
		v5 := corestr.NewValidValueUsingAny(false, true, "test")
		v6 := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		actual := args.Map{
			"v1Valid": v1.IsValid, "v2Empty": v2.IsEmpty(),
			"v3Invalid": !v3.IsValid, "v4Invalid": !v4.IsValid,
			"v5NotNil": v5 != nil, "v6NotNil": v6 != nil,
		}
		expected := args.Map{
			"v1Valid": true, "v2Empty": true,
			"v3Invalid": true, "v4Invalid": true,
			"v5NotNil": true, "v6NotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- constructors", actual)
	})
}

func Test_C33_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ValueBytesOnce", func() {
		v := corestr.NewValidValue("hello")
		b1 := v.ValueBytesOnce()
		b2 := v.ValueBytesOnce() // cached
		if len(b1) != 5 || len(b2) != 5 { t.Fatal("expected 5") }
		_ = v.ValueBytesOncePtr()
	})
}

func Test_C33_ValidValue_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_IsEmpty_IsWhitespace", func() {
		v := corestr.NewValidValue("")
		v2 := corestr.NewValidValue("  ")
		v3 := corestr.NewValidValue("hello")
		actual := args.Map{
			"empty":      v.IsEmpty(),
			"ws":         v2.IsWhitespace(),
			"notEmpty":   v3.IsEmpty(),
			"trim":       v3.Trim(),
		}
		expected := args.Map{
			"empty": true, "ws": true, "notEmpty": false, "trim": "hello",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- empty/ws", actual)
	})
}

func Test_C33_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_HasValidNonEmpty", func() {
		v := corestr.NewValidValue("hello")
		vEmpty := corestr.NewValidValueEmpty()
		actual := args.Map{
			"nonEmpty":  v.HasValidNonEmpty(),
			"nonWs":     v.HasValidNonWhitespace(),
			"safe":      v.HasSafeNonEmpty(),
			"emptyFail": vEmpty.HasValidNonEmpty(),
		}
		expected := args.Map{
			"nonEmpty": true, "nonWs": true, "safe": true, "emptyFail": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns empty -- HasValidNonEmpty", actual)
	})
}

func Test_C33_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ValueBool", func() {
		vTrue := corestr.NewValidValue("true")
		vFalse := corestr.NewValidValue("false")
		vBad := corestr.NewValidValue("notbool")
		vEmpty := corestr.NewValidValue("")
		actual := args.Map{
			"true": vTrue.ValueBool(), "false": vFalse.ValueBool(),
			"bad": vBad.ValueBool(), "empty": vEmpty.ValueBool(),
		}
		expected := args.Map{"true": true, "false": false, "bad": false, "empty": false}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueBool", actual)
	})
}

func Test_C33_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ValueInt", func() {
		v := corestr.NewValidValue("42")
		vBad := corestr.NewValidValue("abc")
		actual := args.Map{
			"good":   v.ValueInt(0),
			"bad":    vBad.ValueInt(99),
			"defInt": v.ValueDefInt(),
			"badDef": vBad.ValueDefInt(),
		}
		expected := args.Map{"good": 42, "bad": 99, "defInt": 42, "badDef": 0}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- ValueInt", actual)
	})
}

func Test_C33_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ValueByte", func() {
		v := corestr.NewValidValue("200")
		vBad := corestr.NewValidValue("abc")
		vNeg := corestr.NewValidValue("-1")
		vHigh := corestr.NewValidValue("999")
		_ = v.ValueByte(0)
		_ = v.ValueDefByte()
		_ = vBad.ValueByte(0)
		_ = vNeg.ValueByte(0)
		_ = vHigh.ValueByte(0)
	})
}

func Test_C33_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ValueFloat64", func() {
		v := corestr.NewValidValue("3.14")
		vBad := corestr.NewValidValue("abc")
		actual := args.Map{
			"good":   v.ValueFloat64(0),
			"bad":    vBad.ValueFloat64(1.0),
			"defF64": v.ValueDefFloat64(),
		}
		expected := args.Map{"good": 3.14, "bad": 1.0, "defF64": 3.14}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- float64", actual)
	})
}

func Test_C33_ValidValue_Is_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Is_IsAnyOf", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{
			"is":       v.Is("hello"),
			"isNot":    v.Is("world"),
			"anyOf":    v.IsAnyOf("a", "hello", "b"),
			"anyEmpty": v.IsAnyOf(),
			"anyNone":  v.IsAnyOf("x", "y"),
		}
		expected := args.Map{
			"is": true, "isNot": false, "anyOf": true,
			"anyEmpty": true, "anyNone": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Is/IsAnyOf", actual)
	})
}

func Test_C33_ValidValue_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_IsContains_IsAnyContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{
			"contains":    v.IsContains("world"),
			"notContains": v.IsContains("xyz"),
			"anyContains": v.IsAnyContains("xyz", "world"),
			"anyEmpty":    v.IsAnyContains(),
			"anyNone":     v.IsAnyContains("xyz", "abc"),
		}
		expected := args.Map{
			"contains": true, "notContains": false,
			"anyContains": true, "anyEmpty": true, "anyNone": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- IsContains", actual)
	})
}

func Test_C33_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_IsEqualNonSensitive", func() {
		v := corestr.NewValidValue("Hello")
		if !v.IsEqualNonSensitive("hello") { t.Fatal("expected true") }
	})
}

func Test_C33_ValidValue_Regex(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Regex", func() {
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{
			"matches":   v.IsRegexMatches(re),
			"nilMatch":  v.IsRegexMatches(nil),
			"find":      v.RegexFindString(re),
			"nilFind":   v.RegexFindString(nil),
		}
		expected := args.Map{
			"matches": true, "nilMatch": false,
			"find": "123", "nilFind": "",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns non-empty -- Regex", actual)

		items, hasAny := v.RegexFindAllStringsWithFlag(re, -1)
		if !hasAny || len(items) == 0 { t.Fatal("expected items") }
		_, noAny := v.RegexFindAllStringsWithFlag(nil, -1)
		if noAny { t.Fatal("expected false for nil regex") }

		all := v.RegexFindAllStrings(re, -1)
		if len(all) == 0 { t.Fatal("expected items") }
		nilAll := v.RegexFindAllStrings(nil, -1)
		if len(nilAll) != 0 { t.Fatal("expected empty") }
	})
}

func Test_C33_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Split", func() {
		v := corestr.NewValidValue("a,b,c")
		s := v.Split(",")
		if len(s) != 3 { t.Fatal("expected 3") }
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")
	})
}

func Test_C33_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Clone", func() {
		v := corestr.NewValidValue("hello")
		clone := v.Clone()
		if clone.Value != "hello" { t.Fatal("clone mismatch") }
		var nilV *corestr.ValidValue
		if nilV.Clone() != nil { t.Fatal("expected nil") }
	})
}

func Test_C33_ValidValue_String_FullString(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_String_FullString", func() {
		v := corestr.NewValidValue("hello")
		_ = v.String()
		_ = v.FullString()
		var nilV *corestr.ValidValue
		_ = nilV.String()
		_ = nilV.FullString()
	})
}

func Test_C33_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Clear_Dispose", func() {
		v := corestr.NewValidValue("hello")
		v.Clear()
		if v.Value != "" { t.Fatal("expected empty after clear") }
		v2 := corestr.NewValidValue("x")
		v2.Dispose()
		var nilV *corestr.ValidValue
		nilV.Clear()
		nilV.Dispose()
	})
}

func Test_C33_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Json", func() {
		v := corestr.NewValidValue("hello")
		j := v.Json()
		_ = j.JsonString()
		_ = v.JsonPtr()
		_, _ = v.Serialize()
	})
}

func Test_C33_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_ParseInjectUsingJson", func() {
		v := corestr.NewValidValue("hello")
		r := corejson.New(corestr.ValidValue{Value: "world", IsValid: true})
		_, _ = v.ParseInjectUsingJson(&r)
	})
}

func Test_C33_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_C33_ValidValue_Deserialize", func() {
		v := corestr.NewValidValue("hello")
		var target corestr.ValidValue
		_ = v.Deserialize(&target)
	})
}

// ── LeftRight ──

func Test_C33_LeftRight_Constructors(t *testing.T) {
	safeTest(t, "Test_C33_LeftRight_Constructors", func() {
		lr := corestr.NewLeftRight("L", "R")
		inv := corestr.InvalidLeftRight("err")
		invNoMsg := corestr.InvalidLeftRightNoMessage()
		fromSlice := corestr.LeftRightUsingSlice([]string{"a", "b"})
		fromSlice1 := corestr.LeftRightUsingSlice([]string{"a"})
		fromSlice0 := corestr.LeftRightUsingSlice([]string{})
		fromSlicePtr := corestr.LeftRightUsingSlicePtr([]string{"x", "y"})
		fromSlicePtrEmpty := corestr.LeftRightUsingSlicePtr([]string{})
		trimmed := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		trimmed1 := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		trimmedNil := corestr.LeftRightTrimmedUsingSlice(nil)
		trimmedEmpty := corestr.LeftRightTrimmedUsingSlice([]string{})

		actual := args.Map{
			"lrValid": lr.IsValid, "invInvalid": !inv.IsValid,
			"invNoMsg": !invNoMsg.IsValid,
			"sliceValid": fromSlice.IsValid, "slice1Invalid": !fromSlice1.IsValid,
			"slice0Invalid": !fromSlice0.IsValid,
			"ptrValid": fromSlicePtr.IsValid, "ptrEmptyInv": !fromSlicePtrEmpty.IsValid,
			"trimValid": trimmed.IsValid, "trim1Inv": !trimmed1.IsValid,
			"trimNilInv": !trimmedNil.IsValid, "trimEmptyInv": !trimmedEmpty.IsValid,
		}
		expected := args.Map{
			"lrValid": true, "invInvalid": true, "invNoMsg": true,
			"sliceValid": true, "slice1Invalid": true, "slice0Invalid": true,
			"ptrValid": true, "ptrEmptyInv": true,
			"trimValid": true, "trim1Inv": true,
			"trimNilInv": true, "trimEmptyInv": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- constructors", actual)
	})
}

func Test_C33_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_C33_LeftRight_Methods", func() {
		lr := corestr.NewLeftRight("left", "right")
		_ = lr.LeftBytes()
		_ = lr.RightBytes()
		_ = lr.LeftTrim()
		_ = lr.RightTrim()

		actual := args.Map{
			"isLeftEmpty":  lr.IsLeftEmpty(),
			"isRightEmpty": lr.IsRightEmpty(),
			"isLeftWs":     lr.IsLeftWhitespace(),
			"isRightWs":    lr.IsRightWhitespace(),
			"hasValidL":    lr.HasValidNonEmptyLeft(),
			"hasValidR":    lr.HasValidNonEmptyRight(),
			"hasValidWsL":  lr.HasValidNonWhitespaceLeft(),
			"hasValidWsR":  lr.HasValidNonWhitespaceRight(),
			"hasSafe":      lr.HasSafeNonEmpty(),
			"isLeft":       lr.IsLeft("left"),
			"isRight":      lr.IsRight("right"),
			"is":           lr.Is("left", "right"),
		}
		expected := args.Map{
			"isLeftEmpty": false, "isRightEmpty": false,
			"isLeftWs": false, "isRightWs": false,
			"hasValidL": true, "hasValidR": true,
			"hasValidWsL": true, "hasValidWsR": true,
			"hasSafe": true, "isLeft": true, "isRight": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- methods", actual)
	})
}

func Test_C33_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C33_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")
		var nilLR *corestr.LeftRight
		actual := args.Map{
			"equal":    lr1.IsEqual(lr2),
			"notEqual": lr1.IsEqual(lr3),
			"nilBoth":  nilLR.IsEqual(nil),
			"nilLeft":  nilLR.IsEqual(lr1),
		}
		expected := args.Map{"equal": true, "notEqual": false, "nilBoth": true, "nilLeft": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsEqual", actual)
	})
}

func Test_C33_LeftRight_Regex(t *testing.T) {
	safeTest(t, "Test_C33_LeftRight_Regex", func() {
		lr := corestr.NewLeftRight("abc123", "xyz456")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{
			"leftMatch":   lr.IsLeftRegexMatch(re),
			"rightMatch":  lr.IsRightRegexMatch(re),
			"nilLeft":     lr.IsLeftRegexMatch(nil),
			"nilRight":    lr.IsRightRegexMatch(nil),
		}
		expected := args.Map{
			"leftMatch": true, "rightMatch": true,
			"nilLeft": false, "nilRight": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- regex", actual)
	})
}

func Test_C33_LeftRight_Clone_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C33_LeftRight_Clone_Clear_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		clone := lr.Clone()
		if !clone.IsEqual(lr) { t.Fatal("clone mismatch") }
		_ = lr.NonPtr()
		_ = lr.Ptr()
		lr.Clear()
		lr.Dispose()
		var nilLR *corestr.LeftRight
		nilLR.Clear()
		nilLR.Dispose()
	})
}

// ── LeftMiddleRight ──

func Test_C33_LeftMiddleRight_Constructors(t *testing.T) {
	safeTest(t, "Test_C33_LeftMiddleRight_Constructors", func() {
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")
		inv := corestr.InvalidLeftMiddleRight("err")
		invNoMsg := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{
			"valid": lmr.IsValid, "inv": !inv.IsValid, "invNoMsg": !invNoMsg.IsValid,
		}
		expected := args.Map{"valid": true, "inv": true, "invNoMsg": true}
		expected.ShouldBeEqual(t, 0, "LMR returns correct value -- constructors", actual)
	})
}

func Test_C33_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_C33_LeftMiddleRight_Methods", func() {
		lmr := corestr.NewLeftMiddleRight("left", "mid", "right")
		_ = lmr.LeftBytes()
		_ = lmr.RightBytes()
		_ = lmr.MiddleBytes()
		_ = lmr.LeftTrim()
		_ = lmr.RightTrim()
		_ = lmr.MiddleTrim()

		actual := args.Map{
			"isLeftEmpty": lmr.IsLeftEmpty(), "isRightEmpty": lmr.IsRightEmpty(),
			"isMidEmpty": lmr.IsMiddleEmpty(),
			"isLeftWs": lmr.IsLeftWhitespace(), "isRightWs": lmr.IsRightWhitespace(),
			"isMidWs": lmr.IsMiddleWhitespace(),
			"hasValidL": lmr.HasValidNonEmptyLeft(), "hasValidR": lmr.HasValidNonEmptyRight(),
			"hasValidM": lmr.HasValidNonEmptyMiddle(),
			"hasValidWsL": lmr.HasValidNonWhitespaceLeft(),
			"hasValidWsR": lmr.HasValidNonWhitespaceRight(),
			"hasValidWsM": lmr.HasValidNonWhitespaceMiddle(),
			"hasSafe": lmr.HasSafeNonEmpty(),
			"isAll": lmr.IsAll("left", "mid", "right"),
			"is": lmr.Is("left", "right"),
		}
		expected := args.Map{
			"isLeftEmpty": false, "isRightEmpty": false, "isMidEmpty": false,
			"isLeftWs": false, "isRightWs": false, "isMidWs": false,
			"hasValidL": true, "hasValidR": true, "hasValidM": true,
			"hasValidWsL": true, "hasValidWsR": true, "hasValidWsM": true,
			"hasSafe": true, "isAll": true, "is": true,
		}
		expected.ShouldBeEqual(t, 0, "LMR returns correct value -- methods", actual)
	})
}

func Test_C33_LeftMiddleRight_Clone_ToLeftRight_Clear(t *testing.T) {
	safeTest(t, "Test_C33_LeftMiddleRight_Clone_ToLeftRight_Clear", func() {
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")
		clone := lmr.Clone()
		if !clone.IsAll("L", "M", "R") { t.Fatal("clone mismatch") }
		lr := lmr.ToLeftRight()
		if !lr.Is("L", "R") { t.Fatal("ToLeftRight mismatch") }
		lmr.Clear()
		lmr.Dispose()
		var nilLMR *corestr.LeftMiddleRight
		nilLMR.Clear()
		nilLMR.Dispose()
	})
}

// ── Hashmap basics ──

func Test_C33_Hashmap_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_C33_Hashmap_IsEmpty_HasItems", func() {
		hm := corestr.New.Hashmap.Empty()
		actual := args.Map{"empty": hm.IsEmpty(), "hasItems": hm.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- empty", actual)
	})
}

func Test_C33_Hashmap_AddOrUpdate(t *testing.T) {
	safeTest(t, "Test_C33_Hashmap_AddOrUpdate", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdateKeyStrValInt("k2", 42)
		hm.AddOrUpdateKeyStrValFloat("k3", 3.14)
		hm.AddOrUpdateKeyStrValFloat64("k4", 2.71)
		hm.AddOrUpdateKeyStrValAny("k5", true)
		if hm.Length() != 5 { t.Fatalf("expected 5, got %d", hm.Length()) }
	})
}

func Test_C33_Hashmap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C33_Hashmap_IsEmptyLock", func() {
		hm := corestr.New.Hashmap.Empty()
		if !hm.IsEmptyLock() { t.Fatal("expected empty") }
	})
}

func Test_C33_Hashmap_Collection(t *testing.T) {
	safeTest(t, "Test_C33_Hashmap_Collection", func() {
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c := hm.Collection()
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

// ── Hashset basics ──

func Test_C33_Hashset_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_C33_Hashset_IsEmpty_HasItems", func() {
		hs := corestr.New.Hashset.Empty()
		actual := args.Map{"empty": hs.IsEmpty(), "hasItems": hs.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- empty", actual)
	})
}

func Test_C33_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_C33_Hashset_AddCapacities", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		hs.AddCapacities(10)
		hs.AddCapacities()
		hs.AddCapacitiesLock(5)
		hs.AddCapacitiesLock()
	})
}

func Test_C33_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_C33_Hashset_Resize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.Resize(10)
		hs.Resize(0) // smaller, should not resize
		hs.ResizeLock(20)
		hs.ResizeLock(0) // smaller
	})
}

// ── SimpleSlice additional ──

func Test_C33_SS_InsertAt(t *testing.T) {
	safeTest(t, "Test_C33_SS_InsertAt", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")
		if ss.Length() != 3 { t.Fatal("expected 3") }
	})
}

func Test_C33_SS_AddsIf(t *testing.T) {
	safeTest(t, "Test_C33_SS_AddsIf", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddsIf(true, "a", "b")
		ss.AddsIf(false, "c")
		if ss.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C33_SS_FirstLast(t *testing.T) {
	safeTest(t, "Test_C33_SS_FirstLast", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		actual := args.Map{
			"first": ss.First(), "last": ss.Last(),
			"firstD": ss.FirstDynamic(), "lastD": ss.LastDynamic(),
			"firstOrDef": ss.FirstOrDefault(), "lastOrDef": ss.LastOrDefault(),
		}
		expected := args.Map{
			"first": "a", "last": "c",
			"firstD": "a", "lastD": "c",
			"firstOrDef": "a", "lastOrDef": "c",
		}
		expected.ShouldBeEqual(t, 0, "SS returns correct value -- First/Last", actual)
	})
}

func Test_C33_SS_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C33_SS_FirstOrDefault_Empty", func() {
		ss := corestr.New.SimpleSlice.Empty()
		if ss.FirstOrDefault() != "" { t.Fatal("expected empty") }
		if ss.LastOrDefault() != "" { t.Fatal("expected empty") }
		_ = ss.FirstOrDefaultDynamic()
		_ = ss.LastOrDefaultDynamic()
	})
}

func Test_C33_SS_Skip_Take_Limit(t *testing.T) {
	safeTest(t, "Test_C33_SS_Skip_Take_Limit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		skip := ss.Skip(1)
		if len(skip) != 2 { t.Fatal("expected 2") }
		skipAll := ss.Skip(99)
		if len(skipAll) != 0 { t.Fatal("expected 0") }
		_ = ss.SkipDynamic(1)

		take := ss.Take(2)
		if len(take) != 2 { t.Fatal("expected 2") }
		takeAll := ss.Take(99)
		if len(takeAll) != 3 { t.Fatal("expected 3") }
		_ = ss.TakeDynamic(2)

		limit := ss.Limit(2)
		if len(limit) != 2 { t.Fatal("expected 2") }
		_ = ss.LimitDynamic(2)
	})
}

func Test_C33_SS_IsContains_IndexOf(t *testing.T) {
	safeTest(t, "Test_C33_SS_IsContains_IndexOf", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		actual := args.Map{
			"contains":    ss.IsContains("a"),
			"notContains": ss.IsContains("z"),
			"indexOf":     ss.IndexOf("b"),
			"notFound":    ss.IndexOf("z"),
			"hasAny":      ss.HasAnyItem(),
			"lastIdx":     ss.LastIndex(),
		}
		expected := args.Map{
			"contains": true, "notContains": false,
			"indexOf": 1, "notFound": -1,
			"hasAny": true, "lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "SS returns correct value -- IsContains/IndexOf", actual)
	})
}

func Test_C33_SS_CountFunc(t *testing.T) {
	safeTest(t, "Test_C33_SS_CountFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		if count != 2 { t.Fatal("expected 2") }
	})
}

func Test_C33_SS_IsContainsFunc_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C33_SS_IsContainsFunc_IndexOfFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("Hello", "World")
		found := ss.IsContainsFunc("hello", func(item, searching string) bool {
			return item == "Hello"
		})
		if !found { t.Fatal("expected true") }

		idx := ss.IndexOfFunc("World", func(item, searching string) bool {
			return item == searching
		})
		if idx != 1 { t.Fatal("expected 1") }
	})
}

func Test_C33_SS_AsDefaultError_AsError(t *testing.T) {
	safeTest(t, "Test_C33_SS_AsDefaultError_AsError", func() {
		ss := corestr.New.SimpleSlice.Lines("err1", "err2")
		err := ss.AsDefaultError()
		if err == nil { t.Fatal("expected error") }
		err2 := ss.AsError(",")
		if err2 == nil { t.Fatal("expected error") }

		empty := corestr.New.SimpleSlice.Empty()
		if empty.AsDefaultError() != nil { t.Fatal("expected nil") }
		var nilSS *corestr.SimpleSlice
		if nilSS.AsError(",") != nil { t.Fatal("expected nil for nil") }
	})
}

func Test_C33_SS_AddStruct_AddPointer(t *testing.T) {
	safeTest(t, "Test_C33_SS_AddStruct_AddPointer", func() {
		ss := corestr.New.SimpleSlice.Empty()
		type sample struct{ Name string }
		s := sample{Name: "test"}
		ss.AddStruct(true, s)
		ss.AddPointer(true, &s)
	})
}

func Test_C33_SS_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C33_SS_AddAsTitleValue", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddAsTitleValue("key", "val")
		ss.AddAsTitleValueIf(true, "k2", "v2")
		ss.AddAsTitleValueIf(false, "k3", "v3")
		if ss.Length() != 2 { t.Fatal("expected 2") }
	})
}

// ── CloneSlice ──

func Test_C33_CloneSlice(t *testing.T) {
	safeTest(t, "Test_C33_CloneSlice", func() {
		result := corestr.CloneSlice([]string{"a", "b"})
		if len(result) != 2 { t.Fatal("expected 2") }
		nilResult := corestr.CloneSlice(nil)
		if len(nilResult) != 0 { t.Fatal("expected empty for nil") }
	})
}

func Test_C33_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_C33_CloneSliceIf", func() {
		result := corestr.CloneSliceIf(true, []string{"a"}...)
		if len(result) != 1 { t.Fatal("expected 1") }
		noClone := corestr.CloneSliceIf(false, []string{"a"}...)
		if len(noClone) != 1 { t.Fatal("expected original") }
	})
}
