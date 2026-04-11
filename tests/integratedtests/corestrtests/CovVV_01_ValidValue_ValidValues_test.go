package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Segment 18 Part 2
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovVV_01_Constructors(t *testing.T) {
	safeTest(t, "Test_CovVV_01_Constructors", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid hello", actual)
		ve := corestr.NewValidValueEmpty()
		actual = args.Map{"result": ve.Value != "" || !ve.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid empty", actual)
		iv := corestr.InvalidValidValue("msg")
		actual = args.Map{"result": iv.IsValid || iv.Message != "msg"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with msg", actual)
		ivn := corestr.InvalidValidValueNoMessage()
		actual = args.Map{"result": ivn.IsValid || ivn.Message != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid no msg", actual)
	})
}

func Test_CovVV_02_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_CovVV_02_NewValidValueUsingAny", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAny(false, true, "test")

		// Act
		actual := args.Map{"result": vv.IsValid}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
	})
}

func Test_CovVV_03_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_CovVV_03_NewValidValueUsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		_ = vv
	})
}

func Test_CovVV_04_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_CovVV_04_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")
		b := vv.ValueBytesOnce()

		// Act
		actual := args.Map{"result": len(b) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// second call returns cached
		b2 := vv.ValueBytesOnce()
		actual = args.Map{"result": len(b2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// deprecated alias
		b3 := vv.ValueBytesOncePtr()
		actual = args.Map{"result": len(b3) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovVV_05_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_CovVV_05_IsEmpty_IsWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("")

		// Act
		actual := args.Map{"result": vv.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": vv.IsWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
		vv2 := corestr.NewValidValue("hi")
		actual = args.Map{"result": vv2.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovVV_06_Trim(t *testing.T) {
	safeTest(t, "Test_CovVV_06_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue("  hi  ")

		// Act
		actual := args.Map{"result": vv.Trim() != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_CovVV_07_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovVV_07_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"result": vv.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.HasValidNonWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		iv := corestr.InvalidValidValue("")
		actual = args.Map{"result": iv.HasValidNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovVV_08_ValueBool(t *testing.T) {
	safeTest(t, "Test_CovVV_08_ValueBool", func() {
		// Arrange
		vv := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"result": vv.ValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		vv2 := corestr.NewValidValue("")
		actual = args.Map{"result": vv2.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		vv3 := corestr.NewValidValue("abc")
		actual = args.Map{"result": vv3.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovVV_09_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovVV_09_ValueInt_ValueDefInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"result": vv.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": vv.ValueDefInt() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		vv2 := corestr.NewValidValue("abc")
		actual = args.Map{"result": vv2.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_CovVV_10_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovVV_10_ValueByte_ValueDefByte", func() {
		// Arrange
		vv := corestr.NewValidValue("100")

		// Act
		actual := args.Map{"result": vv.ValueByte(0) != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": vv.ValueDefByte() != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		// out of range
		vv2 := corestr.NewValidValue("999")
		actual = args.Map{"result": vv2.ValueByte(5) != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255 (max)", actual)
		// negative
		vv3 := corestr.NewValidValue("-1")
		actual = args.Map{"result": vv3.ValueByte(5) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// invalid
		vv4 := corestr.NewValidValue("abc")
		actual = args.Map{"result": vv4.ValueByte(7) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovVV_11_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovVV_11_ValueFloat64_ValueDefFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{"result": vv.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual = args.Map{"result": vv.ValueDefFloat64() != 3.14}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		vv2 := corestr.NewValidValue("abc")
		actual = args.Map{"result": vv2.ValueFloat64(1.5) != 1.5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_CovVV_12_Is_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_CovVV_12_Is_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.Is("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.Is("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.IsAnyOf("a", "hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsAnyOf("a", "b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_CovVV_13_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_CovVV_13_IsContains_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": vv.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsAnyContains("xyz", "world")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsAnyContains("xyz", "abc")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_CovVV_14_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_CovVV_14_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovVV_15_Regex(t *testing.T) {
	safeTest(t, "Test_CovVV_15_Regex", func() {
		// Arrange
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": vv.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.RegexFindString(re) != "123"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual = args.Map{"result": vv.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		items := vv.RegexFindAllStrings(re, -1)
		actual = args.Map{"result": len(items) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(vv.RegexFindAllStrings(nil, -1)) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		items2, has := vv.RegexFindAllStringsWithFlag(re, -1)
		actual = args.Map{"result": has || len(items2) != 1}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		_, has2 := vv.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": has2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovVV_16_Split_SplitNonEmpty_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovVV_16_Split_SplitNonEmpty_SplitTrimNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")

		// Act
		actual := args.Map{"result": len(parts) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		vv2 := corestr.NewValidValue("a,,b")
		parts2 := vv2.SplitNonEmpty(",")
		_ = parts2
		vv3 := corestr.NewValidValue("a, ,b")
		parts3 := vv3.SplitTrimNonWhitespace(",")
		_ = parts3
	})
}

func Test_CovVV_17_Clone(t *testing.T) {
	safeTest(t, "Test_CovVV_17_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		c := vv.Clone()

		// Act
		actual := args.Map{"result": c.Value != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		// nil clone
		var nilVV *corestr.ValidValue
		actual = args.Map{"result": nilVV.Clone() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovVV_18_String_FullString(t *testing.T) {
	safeTest(t, "Test_CovVV_18_String_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.String() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		_ = vv.FullString()
		// nil
		var nilVV *corestr.ValidValue
		actual = args.Map{"result": nilVV.String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": nilVV.FullString() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovVV_19_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovVV_19_Clear_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vv.Clear()

		// Act
		actual := args.Map{"result": vv.Value != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		vv2 := corestr.NewValidValue("hello")
		vv2.Dispose()
		// nil clear
		var nilVV *corestr.ValidValue
		nilVV.Clear()
		nilVV.Dispose()
	})
}

func Test_CovVV_20_Json_ParseInject_Serialize(t *testing.T) {
	safeTest(t, "Test_CovVV_20_Json_ParseInject_Serialize", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		_ = vv.Json()
		jr := vv.JsonPtr()
		vv2 := &corestr.ValidValue{}
		r, err := vv2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil || r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		_, err2 := vv.Serialize()
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := &corestr.ValidValue{}
		err3 := vv.Deserialize(target)
		actual = args.Map{"result": err3 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues — Segment 19 Part 1
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovVVs_01_Constructors(t *testing.T) {
	safeTest(t, "Test_CovVVs_01_Constructors", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vvs2 := corestr.NewValidValues(5)
		actual = args.Map{"result": vvs2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vv1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv2 := corestr.ValidValue{Value: "b", IsValid: true}
		vvs3 := corestr.NewValidValuesUsingValues(vv1, vv2)
		actual = args.Map{"result": vvs3.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty values
		vvs4 := corestr.NewValidValuesUsingValues()
		actual = args.Map{"result": vvs4.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovVVs_02_Count_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovVVs_02_Count_HasAnyItem_LastIndex_HasIndex", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.Count() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": vvs.HasAnyItem()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		vvs.Add("a")
		actual = args.Map{"result": vvs.Count() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": vvs.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vvs.LastIndex() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": vvs.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vvs.HasIndex(1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovVVs_03_Add_AddFull(t *testing.T) {
	safeTest(t, "Test_CovVVs_03_Add_AddFull", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		vvs.AddFull(false, "b", "msg")
		actual = args.Map{"result": vvs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovVVs_04_Adds_AddsPtr(t *testing.T) {
	safeTest(t, "Test_CovVVs_04_Adds_AddsPtr", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vvs.Adds()
		actual = args.Map{"result": vvs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vv := corestr.NewValidValue("c")
		vvs.AddsPtr(vv)
		actual = args.Map{"result": vvs.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		vvs.AddsPtr()
		actual = args.Map{"result": vvs.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovVVs_05_AddValidValues(t *testing.T) {
	safeTest(t, "Test_CovVVs_05_AddValidValues", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		vvs.AddValidValues(vvs2)

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vvs.AddValidValues(nil)
		actual = args.Map{"result": vvs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vvs.AddValidValues(corestr.EmptyValidValues())
		actual = args.Map{"result": vvs.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovVVs_06_AddHashsetMap_AddHashset(t *testing.T) {
	safeTest(t, "Test_CovVVs_06_AddHashsetMap_AddHashset", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vvs.AddHashsetMap(nil)
		hs := corestr.New.Hashset.Strings([]string{"c"})
		vvs.AddHashset(hs)
		actual = args.Map{"result": vvs.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		vvs.AddHashset(nil)
	})
}

func Test_CovVVs_07_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovVVs_07_ConcatNew", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		// no args, clone
		c := vvs.ConcatNew(true)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// no args, no clone
		c2 := vvs.ConcatNew(false)
		actual = args.Map{"result": c2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with args
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		c3 := vvs.ConcatNew(true, vvs2)
		actual = args.Map{"result": c3.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovVVs_08_Find(t *testing.T) {
	safeTest(t, "Test_CovVVs_08_Find", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		// empty find
		r := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})

		// Act
		actual := args.Map{"result": len(r) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vvs.Add("a")
		vvs.Add("b")
		// find all
		r2 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})
		actual = args.Map{"result": len(r2) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// break
		r3 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, true
		})
		actual = args.Map{"result": len(r3) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// skip
		r4 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, false, false
		})
		actual = args.Map{"result": len(r4) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovVVs_09_SafeValueAt_SafeValidValueAt_SafeIndexes(t *testing.T) {
	safeTest(t, "Test_CovVVs_09_SafeValueAt_SafeValidValueAt_SafeIndexes", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.SafeValueAt(0) != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": vvs.SafeValidValueAt(0) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		vvs.Add("a")
		vvs.AddFull(false, "b", "")
		actual = args.Map{"result": vvs.SafeValueAt(0) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": vvs.SafeValueAt(99) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": vvs.SafeValidValueAt(0) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		// invalid valid value returns empty
		actual = args.Map{"result": vvs.SafeValidValueAt(1) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
		vals := vvs.SafeValuesAtIndexes(0, 1)
		actual = args.Map{"result": len(vals) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vals2 := vvs.SafeValuesAtIndexes()
		actual = args.Map{"result": len(vals2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vals3 := vvs.SafeValidValuesAtIndexes(0)
		actual = args.Map{"result": len(vals3) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		vals4 := vvs.SafeValidValuesAtIndexes()
		actual = args.Map{"result": len(vals4) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovVVs_10_Strings_FullStrings_String(t *testing.T) {
	safeTest(t, "Test_CovVVs_10_Strings_FullStrings_String", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": len(vvs.Strings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": len(vvs.FullStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vvs.Add("a")
		actual = args.Map{"result": len(vvs.Strings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(vvs.FullStrings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = vvs.String()
	})
}

func Test_CovVVs_11_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_CovVVs_11_Hashmap_Map", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		hm := vvs.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		vvs.Add("a")
		hm2 := vvs.Hashmap()
		actual = args.Map{"result": hm2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		m := vvs.Map()
		actual = args.Map{"result": len(m) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovVVs_12_IsEmpty(t *testing.T) {
	safeTest(t, "Test_CovVVs_12_IsEmpty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		vvs.Add("a")
		actual = args.Map{"result": vvs.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollectionNode + NonChainedLinkedCollectionNodes — Segment 19 Part 2
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLCN_01_IsEmpty_HasElement_HasNext(t *testing.T) {
	safeTest(t, "Test_CovLCN_01_IsEmpty_HasElement_HasNext", func() {
		// Arrange
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}

		// Act
		actual := args.Map{"result": node.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": node.HasElement()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has element", actual)
		actual = args.Map{"result": node.HasNext()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)
	})
}

func Test_CovLCN_02_EndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLCN_02_EndOfChain", func() {
		// Arrange
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		end, length := node.EndOfChain()

		// Act
		actual := args.Map{"result": end != node || length != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected self, length 1", actual)
	})
}

func Test_CovLCN_03_Clone(t *testing.T) {
	safeTest(t, "Test_CovLCN_03_Clone", func() {
		// Arrange
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		c := node.Clone()

		// Act
		actual := args.Map{"result": c.HasNext()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no next", actual)
	})
}

func Test_CovLCN_04_IsEqual_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_CovLCN_04_IsEqual_IsChainEqual", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n1 := &corestr.LinkedCollectionNode{Element: col}
		n2 := &corestr.LinkedCollectionNode{Element: col}

		// Act
		actual := args.Map{"result": n1.IsEqual(n2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": n1.IsChainEqual(n2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected chain equal", actual)
		// same ptr
		actual = args.Map{"result": n1.IsChainEqual(n1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal same ptr", actual)
		// nil
		var nilN *corestr.LinkedCollectionNode
		actual = args.Map{"result": nilN.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal nil", actual)
		actual = args.Map{"result": nilN.IsEqual(n1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		actual = args.Map{"result": n1.IsChainEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_CovLCN_05_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_CovLCN_05_IsEqualValue", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n := &corestr.LinkedCollectionNode{Element: col}

		// Act
		actual := args.Map{"result": n.IsEqualValue(col)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": n.IsEqualValue(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_CovLCN_06_String_List_ListPtr_Join_StringList(t *testing.T) {
	safeTest(t, "Test_CovLCN_06_String_List_ListPtr_Join_StringList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"false", "a", "b"})
		n := &corestr.LinkedCollectionNode{Element: col}
		_ = n.String()
		list := n.List()

		// Act
		actualListLen := args.Map{"result": len(list)}

		// Assert
		expectedListLen := args.Map{"result": 3}
		expectedListLen.ShouldBeEqual(t, 0, "List returns 3 items -- Collection has 3 items", actualListLen)
		lp := n.ListPtr()
		actual := args.Map{"result": len(*lp) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_ = n.Join(",")
		_ = n.StringList("Header: ")
	})
}

func Test_CovLCN_07_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_CovLCN_07_CreateLinkedList", func() {
		// Arrange
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n := &corestr.LinkedCollectionNode{Element: col}
		ll := n.CreateLinkedList()

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// --- NonChainedLinkedCollectionNodes ---

func Test_CovNCLCN_01_Basic(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_01_Basic", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		actual := args.Map{"result": nc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": nc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		actual = args.Map{"result": nc.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": nc.IsChainingApplied()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovNCLCN_02_Adds_First_Last(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_02_Adds_First_Last", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)

		// Act
		actual := args.Map{"result": nc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": nc.First() != n1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected n1", actual)
		actual = args.Map{"result": nc.Last() != n2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected n2", actual)
		nc.Adds(nil)
		actual = args.Map{"result": nc.Items() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil items", actual)
	})
}

func Test_CovNCLCN_03_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_03_FirstOrDefault_LastOrDefault", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)

		// Act
		actual := args.Map{"result": nc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": nc.LastOrDefault() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		nc.Adds(n)
		actual = args.Map{"result": nc.FirstOrDefault() != n}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected n", actual)
		actual = args.Map{"result": nc.LastOrDefault() != n}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected n", actual)
	})
}

func Test_CovNCLCN_04_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_04_ApplyChaining", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		// empty apply
		nc.ApplyChaining()
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)
		nc.ApplyChaining()

		// Act
		actual := args.Map{"result": nc.IsChainingApplied()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": n1.HasNext()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected n1 has next", actual)
		// re-apply should be no-op
		nc.ApplyChaining()
	})
}

func Test_CovNCLCN_05_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_05_ToChainedNodes", func() {
		// Arrange
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		// empty
		cn := nc.ToChainedNodes()

		// Act
		actual := args.Map{"result": cn == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)
		cn2 := nc.ToChainedNodes()
		actual = args.Map{"result": cn2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}
