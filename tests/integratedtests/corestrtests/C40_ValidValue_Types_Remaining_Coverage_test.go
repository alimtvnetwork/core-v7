package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══ ValidValue ═══

func Test_C40_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Constructors", func() {
		// Arrange
		v1 := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": v1.Value != "hello" || !v1.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2 := corestr.NewValidValueEmpty()
		actual = args.Map{"result": v2.Value != "" || !v2.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v3 := corestr.InvalidValidValue("bad")
		actual = args.Map{"result": v3.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v4 := corestr.InvalidValidValueNoMessage()
		actual = args.Map{"result": v4.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v5 := corestr.NewValidValueUsingAny(false, true, "hello")
		actual = args.Map{"result": v5.Value == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v6 := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		actual = args.Map{"result": v6.Value == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValue_Methods(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Methods", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": v.IsEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.Trim() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.HasValidNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.HasValidNonWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.Is("hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyOf("hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyOf("xyz")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsContains("ell")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyContains("ell")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsAnyContains("xyz")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsEqualNonSensitive("HELLO")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValue_NumericConversions(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_NumericConversions", func() {
		// Arrange
		v := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"result": v.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.ValueDefInt() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.ValueByte(0) != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.ValueDefByte() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.ValueFloat64(0) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.ValueDefFloat64() == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// bool
		vb := corestr.NewValidValue("true")
		actual = args.Map{"result": vb.ValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vbad := corestr.NewValidValue("xyz")
		actual = args.Map{"result": vbad.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.NewValidValue("").ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// errors
		bad := corestr.NewValidValue("abc")
		actual = args.Map{"result": bad.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": bad.ValueByte(88) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// byte overflow
		big := corestr.NewValidValue("999")
		actual = args.Map{"result": big.ValueByte(0) != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": big.ValueDefByte() != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// negative byte
		neg := corestr.NewValidValue("-1")
		actual = args.Map{"result": neg.ValueByte(0) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValue_Regex(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Regex", func() {
		// Arrange
		v := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": v.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.RegexFindString(re) != "123"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r, ok := v.RegexFindAllStringsWithFlag(re, -1)
		actual = args.Map{"result": ok || len(r) == 0}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_, ok2 := v.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(v.RegexFindAllStrings(re, -1)) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(v.RegexFindAllStrings(nil, -1)) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Split", func() {
		// Arrange
		v := corestr.NewValidValue("a,b,c")

		// Act
		actual := args.Map{"result": len(v.Split(",")) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")
	})
}

func Test_C40_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_ValueBytesOnce", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		b1 := v.ValueBytesOnce()
		b2 := v.ValueBytesOnce() // cached

		// Act
		actual := args.Map{"result": len(b1) != 2 || len(b2) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = v.ValueBytesOncePtr()
	})
}

func Test_C40_ValidValue_Clone_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Clone_Clear_Dispose", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		c := v.Clone()

		// Act
		actual := args.Map{"result": c.Value != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilV *corestr.ValidValue
		actual = args.Map{"result": nilV.Clone() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v.Clear()
		actual = args.Map{"result": v.Value != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2 := corestr.NewValidValue("x")
		v2.Dispose()
		nilV.Clear()
		nilV.Dispose()
	})
}

func Test_C40_ValidValue_String_FullString(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_String_FullString", func() {
		// Arrange
		v := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"result": v.String() != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v.FullString() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilV *corestr.ValidValue
		actual = args.Map{"result": nilV.String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": nilV.FullString() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValue_JSON(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_JSON", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		j := v.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		jp := v.JsonPtr()
		actual = args.Map{"hasError": jp.HasError()}
		expected = args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
		_, err := v.Serialize()
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ═══ ValidValues ═══

func Test_C40_ValidValues(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues", func() {
		// Arrange
		vv := corestr.NewValidValues(5)
		vv.Add("a")
		vv.AddFull(true, "b", "msg")

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.Count() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.LastIndex() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.SafeValueAt(0) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.SafeValueAt(100) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": vv.SafeValidValueAt(0) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = vv.SafeValuesAtIndexes(0, 1)
		_ = vv.SafeValidValuesAtIndexes(0, 1)
		_ = vv.Strings()
		_ = vv.FullStrings()
		_ = vv.String()
		// empty
		evv := corestr.EmptyValidValues()
		actual = args.Map{"result": evv.SafeValueAt(0) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": evv.SafeValidValueAt(0) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(evv.Strings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(evv.FullStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Add", func() {
		// Arrange
		vv := corestr.NewValidValues(5)
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv.Adds(v1)

		// Act
		actual := args.Map{"result": vv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vv.AddsPtr(corestr.NewValidValue("b"))
		actual = args.Map{"result": vv.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vv.AddHashsetMap(map[string]bool{"c": true})
		vv.AddHashsetMap(nil)
		hs := corestr.New.Hashset.StringsSpreadItems("d")
		vv.AddHashset(hs)
		vv.AddHashset(nil)
		vv.AddValidValues(corestr.NewValidValuesUsingValues(corestr.ValidValue{Value: "e", IsValid: true}))
		vv.AddValidValues(nil)
	})
}

func Test_C40_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_ConcatNew", func() {
		// Arrange
		vv := corestr.NewValidValues(2)
		vv.Add("a")
		cn := vv.ConcatNew(true)

		// Act
		actual := args.Map{"result": cn.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cn2 := vv.ConcatNew(false)
		actual = args.Map{"result": cn2 != vv}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vv2 := corestr.NewValidValues(2)
		vv2.Add("b")
		cn3 := vv.ConcatNew(true, vv2)
		actual = args.Map{"result": cn3.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Find", func() {
		// Arrange
		vv := corestr.NewValidValues(3)
		vv.Add("a")
		vv.Add("b")
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "a", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		found2 := corestr.EmptyValidValues().Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})
		actual = args.Map{"result": len(found2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_ValidValues_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Hashmap_Map", func() {
		// Arrange
		vv := corestr.NewValidValues(2)
		vv.Add("a")
		hm := vv.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		m := vv.Map()
		actual = args.Map{"result": len(m) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		evv := corestr.EmptyValidValues()
		actual = args.Map{"result": evv.Hashmap().Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

// ═══ ValueStatus ═══

func Test_C40_ValueStatus(t *testing.T) {
	safeTest(t, "Test_C40_ValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("bad")

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vs2 := corestr.InvalidValueStatusNoMessage()
		actual = args.Map{"result": vs2.ValueValid.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		c := vs.Clone()
		actual = args.Map{"result": c.ValueValid.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

// ═══ KeyValuePair ═══

func Test_C40_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.KeyName() != "k"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.VariableName() != "k"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueString() != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsVariableNameEqual("k")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsValueEqual("v")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.Compile() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.String() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsKeyEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsValueEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.HasKey()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.HasValue()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsKeyValueEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.TrimKey() != "k"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.TrimValue() != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.Is("k", "v")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsKey("k")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsVal("v")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.IsKeyValueAnyEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.FormatString("%s=%s") != "k=v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "")
		kv.Clear()
		kv.Dispose()
	})
}

func Test_C40_KeyValuePair_Numeric(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair_Numeric", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}

		// Act
		actual := args.Map{"result": kv.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueDefInt() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueByte(0) != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueDefByte() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueFloat64(0) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kv.ValueDefFloat64() == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		kvb := corestr.KeyValuePair{Key: "k", Value: "true"}
		actual = args.Map{"result": kvb.ValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		kvbad := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual = args.Map{"result": kvbad.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvbad.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_KeyValuePair_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair_JSON", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		_ = kv.JsonPtr()
		_, err := kv.Serialize()
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		_ = kv.SerializeMust()
	})
}

// ═══ KeyValueCollection ═══

func Test_C40_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(5)
		kvc.Add("a", "1")
		kvc.AddIf(true, "b", "2")
		kvc.AddIf(false, "skip", "skip")
		kvc.Adds(corestr.KeyValuePair{Key: "c", Value: "3"})

		// Act
		actual := args.Map{"result": kvc.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.Count() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.First().Key != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.FirstOrDefault() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.Last().Key != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.LastOrDefault() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.HasKey("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.IsContains("a")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v, ok := kvc.Get("a")
		actual = args.Map{"result": ok || v != "1"}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kvc.SafeValueAt(0) != "1"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = kvc.SafeValuesAtIndexes(0)
		_ = kvc.Strings()
		_ = kvc.StringsUsingFormat("%s=%s")
		_ = kvc.String()
		_ = kvc.AllKeys()
		_ = kvc.AllKeysSorted()
		_ = kvc.AllValues()
		_ = kvc.Join(",")
		_ = kvc.JoinKeys(",")
		_ = kvc.JoinValues(",")
		_ = kvc.Compile()
		// empty checks
		ekvc := corestr.Empty.KeyValueCollection()
		actual = args.Map{"result": ekvc.FirstOrDefault() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ekvc.LastOrDefault() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ekvc.SafeValueAt(0) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_KeyValueCollection_AddMethods(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection_AddMethods", func() {
		kvc := corestr.New.KeyValues.Cap(5)
		kvc.AddMap(map[string]string{"a": "1"})
		kvc.AddMap(nil)
		kvc.AddHashsetMap(map[string]bool{"b": true})
		kvc.AddHashsetMap(nil)
		kvc.AddHashset(corestr.New.Hashset.StringsSpreadItems("c"))
		kvc.AddHashset(nil)
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("d", "4")
		kvc.AddsHashmap(h)
		kvc.AddsHashmap(nil)
		kvc.AddsHashmaps(h)
		kvc.AddsHashmaps(nil)
		kvc.AddStringBySplit("=", "e=5")
		kvc.AddStringBySplitTrim("=", " f = 6 ")
		_ = kvc.Hashmap()
		_ = kvc.Map()
	})
}

func Test_C40_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection_Find", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(3)
		kvc.Add("a", "1")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_KeyValueCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection_JSON", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(2)
		kvc.Add("a", "1")
		j := kvc.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		_ = kvc.JsonPtr()
		_ = kvc.JsonModel()
		_ = kvc.JsonModelAny()
		b, err := kvc.MarshalJSON()
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		kvc2 := &corestr.KeyValueCollection{}
		err2 := kvc2.UnmarshalJSON(b)
		actual = args.Map{"result": err2}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
		_ = kvc.SerializeMust()
		_, err3 := kvc.Serialize()
		actual = args.Map{"result": err3}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err3", actual)
	})
}

// ═══ KeyAnyValuePair ═══

func Test_C40_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_C40_KeyAnyValuePair", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"result": kav.KeyName() != "k"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.VariableName() != "k"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.ValueAny() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.IsVariableNameEqual("k")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.IsValueNull()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.HasNonNull()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.HasValue()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.IsValueEmptyString()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.IsValueWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vs := kav.ValueString()
		actual = args.Map{"result": vs == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// call again for cached
		vs2 := kav.ValueString()
		actual = args.Map{"result": vs2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.Compile() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.String() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = kav.SerializeMust()
		kav.Clear()
		kav.Dispose()
		// nil
		var nilKAV *corestr.KeyAnyValuePair
		nilKAV.Clear()
		nilKAV.Dispose()
	})
}

func Test_C40_KeyAnyValuePair_NullValue(t *testing.T) {
	safeTest(t, "Test_C40_KeyAnyValuePair_NullValue", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{"result": kav.IsValueNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		vs := kav.ValueString()
		_ = vs // should be empty via GetOnce
	})
}

func Test_C40_KeyAnyValuePair_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyAnyValuePair_JSON", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		_ = kav.JsonPtr()
		actual = args.Map{"result": kav.AsJsonContractsBinder() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.AsJsoner() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": kav.AsJsonParseSelfInjector() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

// ═══ LeftRight ═══

func Test_C40_LeftRight(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("left", "right")

		// Act
		actual := args.Map{"result": lr.Left != "left"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.Right != "right"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsLeftEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsRightEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsLeftWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsRightWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.HasValidNonEmptyLeft()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.HasValidNonEmptyRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.HasValidNonWhitespaceLeft()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.HasValidNonWhitespaceRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsLeft("left")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsRight("right")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.Is("left", "right")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = lr.LeftBytes()
		_ = lr.RightBytes()
		_ = lr.LeftTrim()
		_ = lr.RightTrim()
		_ = lr.NonPtr()
		_ = lr.Ptr()
		lr.Clear()
		lr.Dispose()
		var nilLR *corestr.LeftRight
		nilLR.Clear()
		nilLR.Dispose()
	})
}

func Test_C40_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_IsEqual", func() {
		// Arrange
		a := corestr.NewLeftRight("a", "b")
		b := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": a.IsEqual(a)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilLR *corestr.LeftRight
		actual = args.Map{"result": nilLR.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": a.IsEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{"result": c.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_LeftRight_Regex(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_Regex", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello", "world")
		re := regexp.MustCompile(`^hel`)

		// Act
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(re)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_LeftRight_FromSlice(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_FromSlice", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		actual = args.Map{"result": lr2.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr3 := corestr.LeftRightUsingSlice(nil)
		actual = args.Map{"result": lr3.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr4 := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual = args.Map{"result": lr4.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr5 := corestr.LeftRightUsingSlicePtr(nil)
		actual = args.Map{"result": lr5.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr6 := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual = args.Map{"result": lr6.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr7 := corestr.LeftRightTrimmedUsingSlice(nil)
		actual = args.Map{"result": lr7.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr8 := corestr.LeftRightTrimmedUsingSlice([]string{})
		actual = args.Map{"result": lr8.IsValid}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lr9 := corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		actual = args.Map{"result": lr9.Left != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = corestr.InvalidLeftRight("msg")
		_ = corestr.InvalidLeftRightNoMessage()
	})
}

// ═══ LeftMiddleRight ═══

func Test_C40_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_C40_LeftMiddleRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{"result": lmr.Left != "l"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.Middle != "m"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.Right != "r"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsLeftEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsMiddleEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsRightEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsLeftWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsMiddleWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsRightWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonEmptyLeft()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonEmptyRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonEmptyMiddle()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonWhitespaceLeft()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonWhitespaceRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasValidNonWhitespaceMiddle()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.IsAll("l", "m", "r")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": lmr.Is("l", "r")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = lmr.LeftBytes()
		_ = lmr.MiddleBytes()
		_ = lmr.RightBytes()
		_ = lmr.LeftTrim()
		_ = lmr.MiddleTrim()
		_ = lmr.RightTrim()
		_ = lmr.Clone()
		_ = lmr.ToLeftRight()
		lmr.Clear()
		lmr.Dispose()
		var nilLMR *corestr.LeftMiddleRight
		nilLMR.Clear()
		nilLMR.Dispose()
		_ = corestr.InvalidLeftMiddleRight("msg")
		_ = corestr.InvalidLeftMiddleRightNoMessage()
	})
}

// ═══ Utility functions ═══

func Test_C40_CloneSlice(t *testing.T) {
	safeTest(t, "Test_C40_CloneSlice", func() {
		// Arrange
		r := corestr.CloneSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.CloneSlice(nil)
		actual = args.Map{"result": len(r2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_C40_CloneSliceIf", func() {
		// Arrange
		r := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.CloneSliceIf(false, "a")
		actual = args.Map{"result": len(r2) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r3 := corestr.CloneSliceIf(true)
		actual = args.Map{"result": len(r3) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_AnyToString(t *testing.T) {
	safeTest(t, "Test_C40_AnyToString", func() {
		// Arrange
		r := corestr.AnyToString(false, "hello")

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := corestr.AnyToString(true, "hello")
		actual = args.Map{"result": r2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r3 := corestr.AnyToString(false, "")
		actual = args.Map{"result": r3 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C40_AllIndividualStringsOfStringsLength", func() {
		// Arrange
		s := [][]string{{"a", "b"}, {"c"}}
		r := corestr.AllIndividualStringsOfStringsLength(&s)

		// Act
		actual := args.Map{"result": r != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.AllIndividualStringsOfStringsLength(nil) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C40_AllIndividualsLengthOfSimpleSlices", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("c")
		r := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)

		// Act
		actual := args.Map{"result": r != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.AllIndividualsLengthOfSimpleSlices() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C40_StringUtils(t *testing.T) {
	safeTest(t, "Test_C40_StringUtils", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapDouble("a") != `"a"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapSingle("a") != `'a'`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapTilda("a") != "`a`"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing(`"a"`) != `"a"`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing("a") != `"a"`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing("") != `""`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("'a'") != "'a'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("a") != "'a'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("") != "''"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}
