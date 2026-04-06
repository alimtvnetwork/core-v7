package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══ ValidValue ═══

func Test_C40_ValidValue_Constructors(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Constructors", func() {
		v1 := corestr.NewValidValue("hello")
		if v1.Value != "hello" || !v1.IsValid { t.Fatal() }
		v2 := corestr.NewValidValueEmpty()
		if v2.Value != "" || !v2.IsValid { t.Fatal() }
		v3 := corestr.InvalidValidValue("bad")
		if v3.IsValid { t.Fatal() }
		v4 := corestr.InvalidValidValueNoMessage()
		if v4.IsValid { t.Fatal() }
		v5 := corestr.NewValidValueUsingAny(false, true, "hello")
		if v5.Value == "" { t.Fatal() }
		v6 := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		if v6.Value == "" { t.Fatal() }
	})
}

func Test_C40_ValidValue_Methods(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Methods", func() {
		v := corestr.NewValidValue("hello")
		if v.IsEmpty() { t.Fatal() }
		if v.IsWhitespace() { t.Fatal() }
		if v.Trim() != "hello" { t.Fatal() }
		if !v.HasValidNonEmpty() { t.Fatal() }
		if !v.HasValidNonWhitespace() { t.Fatal() }
		if !v.HasSafeNonEmpty() { t.Fatal() }
		if !v.Is("hello") { t.Fatal() }
		if !v.IsAnyOf("hello") { t.Fatal() }
		if !v.IsAnyOf() { t.Fatal() }
		if v.IsAnyOf("xyz") { t.Fatal() }
		if !v.IsContains("ell") { t.Fatal() }
		if !v.IsAnyContains("ell") { t.Fatal() }
		if !v.IsAnyContains() { t.Fatal() }
		if v.IsAnyContains("xyz") { t.Fatal() }
		if !v.IsEqualNonSensitive("HELLO") { t.Fatal() }
	})
}

func Test_C40_ValidValue_NumericConversions(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_NumericConversions", func() {
		v := corestr.NewValidValue("42")
		if v.ValueInt(0) != 42 { t.Fatal() }
		if v.ValueDefInt() != 42 { t.Fatal() }
		if v.ValueByte(0) != 42 { t.Fatal() }
		if v.ValueDefByte() != 42 { t.Fatal() }
		if v.ValueFloat64(0) == 0 { t.Fatal() }
		if v.ValueDefFloat64() == 0 { t.Fatal() }
		// bool
		vb := corestr.NewValidValue("true")
		if !vb.ValueBool() { t.Fatal() }
		vbad := corestr.NewValidValue("xyz")
		if vbad.ValueBool() { t.Fatal() }
		if corestr.NewValidValue("").ValueBool() { t.Fatal() }
		// errors
		bad := corestr.NewValidValue("abc")
		if bad.ValueInt(99) != 99 { t.Fatal() }
		if bad.ValueByte(88) != 0 { t.Fatal() } // ValueByte returns 0 on error, not defVal
		// byte overflow
		big := corestr.NewValidValue("999")
		if big.ValueByte(0) != 255 { t.Fatal() }
		if big.ValueDefByte() != 255 { t.Fatal() }
		// negative byte
		neg := corestr.NewValidValue("-1")
		if neg.ValueByte(0) != 0 { t.Fatal() }
	})
}

func Test_C40_ValidValue_Regex(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Regex", func() {
		v := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		if !v.IsRegexMatches(re) { t.Fatal() }
		if v.IsRegexMatches(nil) { t.Fatal() }
		if v.RegexFindString(nil) != "" { t.Fatal() }
		if v.RegexFindString(re) != "123" { t.Fatal() }
		r, ok := v.RegexFindAllStringsWithFlag(re, -1)
		if !ok || len(r) == 0 { t.Fatal() }
		_, ok2 := v.RegexFindAllStringsWithFlag(nil, -1)
		if ok2 { t.Fatal() }
		if len(v.RegexFindAllStrings(re, -1)) == 0 { t.Fatal() }
		if len(v.RegexFindAllStrings(nil, -1)) != 0 { t.Fatal() }
	})
}

func Test_C40_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Split", func() {
		v := corestr.NewValidValue("a,b,c")
		if len(v.Split(",")) != 3 { t.Fatal() }
		_ = v.SplitNonEmpty(",")
		_ = v.SplitTrimNonWhitespace(",")
	})
}
func Test_C40_ValidValue_Clone_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_Clone_Clear_Dispose", func() {
		v := corestr.NewValidValue("hi")
		c := v.Clone()
		if c.Value != "hi" { t.Fatal() }
		var nilV *corestr.ValidValue
		if nilV.Clone() != nil { t.Fatal() }
		v.Clear()
		if v.Value != "" { t.Fatal() }
		v2 := corestr.NewValidValue("x")
		v2.Dispose()
		nilV.Clear()
		nilV.Dispose()
	})
}

func Test_C40_ValidValue_String_FullString(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_String_FullString", func() {
		v := corestr.NewValidValue("hi")
		if v.String() != "hi" { t.Fatal() }
		if v.FullString() == "" { t.Fatal() }
		var nilV *corestr.ValidValue
		if nilV.String() != "" { t.Fatal() }
		if nilV.FullString() != "" { t.Fatal() }
	})
}

func Test_C40_ValidValue_JSON(t *testing.T) {
	safeTest(t, "Test_C40_ValidValue_JSON", func() {
		v := corestr.NewValidValue("hi")
		j := v.Json()
		if j.HasError() { t.Fatal(j.Error) }
		jp := v.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
		_, err := v.Serialize()
		if err != nil { t.Fatal(err) }
	})
}

// ═══ ValidValues ═══

func Test_C40_ValidValues(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues", func() {
		vv := corestr.NewValidValues(5)
		vv.Add("a")
		vv.AddFull(true, "b", "msg")
		if vv.Length() != 2 { t.Fatal() }
		if vv.Count() != 2 { t.Fatal() }
		if !vv.HasAnyItem() { t.Fatal() }
		if vv.LastIndex() != 1 { t.Fatal() }
		if !vv.HasIndex(0) { t.Fatal() }
		if vv.IsEmpty() { t.Fatal() }
		if vv.SafeValueAt(0) != "a" { t.Fatal() }
		if vv.SafeValueAt(100) != "" { t.Fatal() }
		if vv.SafeValidValueAt(0) != "a" { t.Fatal() }
		_ = vv.SafeValuesAtIndexes(0, 1)
		_ = vv.SafeValidValuesAtIndexes(0, 1)
		_ = vv.Strings()
		_ = vv.FullStrings()
		_ = vv.String()
		// empty
		evv := corestr.EmptyValidValues()
		if evv.SafeValueAt(0) != "" { t.Fatal() }
		if evv.SafeValidValueAt(0) != "" { t.Fatal() }
		if len(evv.Strings()) != 0 { t.Fatal() }
		if len(evv.FullStrings()) != 0 { t.Fatal() }
	})
}

func Test_C40_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Add", func() {
		vv := corestr.NewValidValues(5)
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv.Adds(v1)
		if vv.Length() != 1 { t.Fatal() }
		vv.AddsPtr(corestr.NewValidValue("b"))
		if vv.Length() != 2 { t.Fatal() }
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
		vv := corestr.NewValidValues(2)
		vv.Add("a")
		cn := vv.ConcatNew(true)
		if cn.Length() != 1 { t.Fatal() }
		cn2 := vv.ConcatNew(false)
		if cn2 != vv { t.Fatal() }
		vv2 := corestr.NewValidValues(2)
		vv2.Add("b")
		cn3 := vv.ConcatNew(true, vv2)
		if cn3.Length() != 2 { t.Fatal() }
	})
}

func Test_C40_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Find", func() {
		vv := corestr.NewValidValues(3)
		vv.Add("a")
		vv.Add("b")
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "a", false
		})
		if len(found) != 1 { t.Fatal() }
		found2 := corestr.EmptyValidValues().Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})
		if len(found2) != 0 { t.Fatal() }
	})
}

func Test_C40_ValidValues_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_C40_ValidValues_Hashmap_Map", func() {
		vv := corestr.NewValidValues(2)
		vv.Add("a")
		hm := vv.Hashmap()
		if hm.Length() != 1 { t.Fatal() }
		m := vv.Map()
		if len(m) != 1 { t.Fatal() }
		evv := corestr.EmptyValidValues()
		if evv.Hashmap().Length() != 0 { t.Fatal() }
	})
}

// ═══ ValueStatus ═══

func Test_C40_ValueStatus(t *testing.T) {
	safeTest(t, "Test_C40_ValueStatus", func() {
		vs := corestr.InvalidValueStatus("bad")
		if vs.ValueValid.IsValid { t.Fatal() }
		vs2 := corestr.InvalidValueStatusNoMessage()
		if vs2.ValueValid.IsValid { t.Fatal() }
		c := vs.Clone()
		if c.ValueValid.IsValid { t.Fatal() }
	})
}

// ═══ KeyValuePair ═══

func Test_C40_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.KeyName() != "k" { t.Fatal() }
		if kv.VariableName() != "k" { t.Fatal() }
		if kv.ValueString() != "v" { t.Fatal() }
		if !kv.IsVariableNameEqual("k") { t.Fatal() }
		if !kv.IsValueEqual("v") { t.Fatal() }
		if kv.Compile() == "" { t.Fatal() }
		if kv.String() == "" { t.Fatal() }
		if kv.IsKeyEmpty() { t.Fatal() }
		if kv.IsValueEmpty() { t.Fatal() }
		if !kv.HasKey() { t.Fatal() }
		if !kv.HasValue() { t.Fatal() }
		if kv.IsKeyValueEmpty() { t.Fatal() }
		if kv.TrimKey() != "k" { t.Fatal() }
		if kv.TrimValue() != "v" { t.Fatal() }
		if !kv.Is("k", "v") { t.Fatal() }
		if !kv.IsKey("k") { t.Fatal() }
		if !kv.IsVal("v") { t.Fatal() }
		if kv.IsKeyValueAnyEmpty() { t.Fatal() }
		if kv.FormatString("%s=%s") != "k=v" { t.Fatal() }
		_ = kv.ValueValid()
		_ = kv.ValueValidOptions(true, "")
		kv.Clear()
		kv.Dispose()
	})
}

func Test_C40_KeyValuePair_Numeric(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair_Numeric", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		if kv.ValueInt(0) != 42 { t.Fatal() }
		if kv.ValueDefInt() != 42 { t.Fatal() }
		if kv.ValueByte(0) != 42 { t.Fatal() }
		if kv.ValueDefByte() != 42 { t.Fatal() }
		if kv.ValueFloat64(0) == 0 { t.Fatal() }
		if kv.ValueDefFloat64() == 0 { t.Fatal() }
		kvb := corestr.KeyValuePair{Key: "k", Value: "true"}
		if !kvb.ValueBool() { t.Fatal() }
		kvbad := corestr.KeyValuePair{Key: "k", Value: "abc"}
		if kvbad.ValueBool() { t.Fatal() }
		if kvbad.ValueInt(99) != 99 { t.Fatal() }
	})
}

func Test_C40_KeyValuePair_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyValuePair_JSON", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()
		if j.HasError() { t.Fatal(j.Error) }
		_ = kv.JsonPtr()
		_, err := kv.Serialize()
		if err != nil { t.Fatal(err) }
		_ = kv.SerializeMust()
	})
}

// ═══ KeyValueCollection ═══

func Test_C40_KeyValueCollection(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection", func() {
		kvc := corestr.New.KeyValues.Cap(5)
		kvc.Add("a", "1")
		kvc.AddIf(true, "b", "2")
		kvc.AddIf(false, "skip", "skip")
		kvc.Adds(corestr.KeyValuePair{Key: "c", Value: "3"})
		if kvc.Length() != 3 { t.Fatal() }
		if kvc.Count() != 3 { t.Fatal() }
		if !kvc.HasAnyItem() { t.Fatal() }
		if !kvc.HasIndex(0) { t.Fatal() }
		if kvc.IsEmpty() { t.Fatal() }
		if kvc.First().Key != "a" { t.Fatal() }
		if kvc.FirstOrDefault() == nil { t.Fatal() }
		if kvc.Last().Key != "c" { t.Fatal() }
		if kvc.LastOrDefault() == nil { t.Fatal() }
		if !kvc.HasKey("a") { t.Fatal() }
		if !kvc.IsContains("a") { t.Fatal() }
		v, ok := kvc.Get("a")
		if !ok || v != "1" { t.Fatal() }
		if kvc.SafeValueAt(0) != "1" { t.Fatal() }
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
		if ekvc.FirstOrDefault() != nil { t.Fatal() }
		if ekvc.LastOrDefault() != nil { t.Fatal() }
		if ekvc.SafeValueAt(0) != "" { t.Fatal() }
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
		kvc := corestr.New.KeyValues.Cap(3)
		kvc.Add("a", "1")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})
		if len(found) != 1 { t.Fatal() }
	})
}

func Test_C40_KeyValueCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyValueCollection_JSON", func() {
		kvc := corestr.New.KeyValues.Cap(2)
		kvc.Add("a", "1")
		j := kvc.Json()
		if j.HasError() { t.Fatal(j.Error) }
		_ = kvc.JsonPtr()
		_ = kvc.JsonModel()
		_ = kvc.JsonModelAny()
		b, err := kvc.MarshalJSON()
		if err != nil { t.Fatal(err) }
		kvc2 := &corestr.KeyValueCollection{}
		err2 := kvc2.UnmarshalJSON(b)
		if err2 != nil { t.Fatal(err2) }
		_ = kvc.SerializeMust()
		_, err3 := kvc.Serialize()
		if err3 != nil { t.Fatal(err3) }
	})
}

// ═══ KeyAnyValuePair ═══

func Test_C40_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_C40_KeyAnyValuePair", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if kav.KeyName() != "k" { t.Fatal() }
		if kav.VariableName() != "k" { t.Fatal() }
		if kav.ValueAny() != 42 { t.Fatal() }
		if !kav.IsVariableNameEqual("k") { t.Fatal() }
		if kav.IsValueNull() { t.Fatal() }
		if !kav.HasNonNull() { t.Fatal() }
		if !kav.HasValue() { t.Fatal() }
		if kav.IsValueEmptyString() { t.Fatal() }
		if kav.IsValueWhitespace() { t.Fatal() }
		vs := kav.ValueString()
		if vs == "" { t.Fatal() }
		// call again for cached
		vs2 := kav.ValueString()
		if vs2 == "" { t.Fatal() }
		if kav.Compile() == "" { t.Fatal() }
		if kav.String() == "" { t.Fatal() }
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
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		if !kav.IsValueNull() { t.Fatal() }
		vs := kav.ValueString()
		_ = vs // should be empty via GetOnce
	})
}

func Test_C40_KeyAnyValuePair_JSON(t *testing.T) {
	safeTest(t, "Test_C40_KeyAnyValuePair_JSON", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()
		if j.HasError() { t.Fatal(j.Error) }
		_ = kav.JsonPtr()
		if kav.AsJsonContractsBinder() == nil { t.Fatal() }
		if kav.AsJsoner() == nil { t.Fatal() }
		if kav.AsJsonParseSelfInjector() == nil { t.Fatal() }
	})
}

// ═══ LeftRight ═══

func Test_C40_LeftRight(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight", func() {
		lr := corestr.NewLeftRight("left", "right")
		if lr.Left != "left" { t.Fatal() }
		if lr.Right != "right" { t.Fatal() }
		if lr.IsLeftEmpty() { t.Fatal() }
		if lr.IsRightEmpty() { t.Fatal() }
		if lr.IsLeftWhitespace() { t.Fatal() }
		if lr.IsRightWhitespace() { t.Fatal() }
		if !lr.HasValidNonEmptyLeft() { t.Fatal() }
		if !lr.HasValidNonEmptyRight() { t.Fatal() }
		if !lr.HasValidNonWhitespaceLeft() { t.Fatal() }
		if !lr.HasValidNonWhitespaceRight() { t.Fatal() }
		if !lr.HasSafeNonEmpty() { t.Fatal() }
		if !lr.IsLeft("left") { t.Fatal() }
		if !lr.IsRight("right") { t.Fatal() }
		if !lr.Is("left", "right") { t.Fatal() }
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
		a := corestr.NewLeftRight("a", "b")
		b := corestr.NewLeftRight("a", "b")
		if !a.IsEqual(b) { t.Fatal() }
		if !a.IsEqual(a) { t.Fatal() }
		var nilLR *corestr.LeftRight
		if !nilLR.IsEqual(nil) { t.Fatal() }
		if a.IsEqual(nil) { t.Fatal() }
	})
}

func Test_C40_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()
		if c.Left != "a" { t.Fatal() }
	})
}

func Test_C40_LeftRight_Regex(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_Regex", func() {
		lr := corestr.NewLeftRight("hello", "world")
		re := regexp.MustCompile(`^hel`)
		if !lr.IsLeftRegexMatch(re) { t.Fatal() }
		if lr.IsRightRegexMatch(re) { t.Fatal() }
		if lr.IsLeftRegexMatch(nil) { t.Fatal() }
		if lr.IsRightRegexMatch(nil) { t.Fatal() }
	})
}

func Test_C40_LeftRight_FromSlice(t *testing.T) {
	safeTest(t, "Test_C40_LeftRight_FromSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		if lr.Left != "a" { t.Fatal() }
		lr2 := corestr.LeftRightUsingSlice([]string{"a"})
		if lr2.Left != "a" { t.Fatal() }
		lr3 := corestr.LeftRightUsingSlice(nil)
		if lr3.IsValid { t.Fatal() }
		lr4 := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		if lr4.Left != "a" { t.Fatal() }
		lr5 := corestr.LeftRightUsingSlicePtr(nil)
		if lr5.IsValid { t.Fatal() }
		lr6 := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		if lr6.Left != "a" { t.Fatal() }
		lr7 := corestr.LeftRightTrimmedUsingSlice(nil)
		if lr7.IsValid { t.Fatal() }
		lr8 := corestr.LeftRightTrimmedUsingSlice([]string{})
		if lr8.IsValid { t.Fatal() }
		lr9 := corestr.LeftRightTrimmedUsingSlice([]string{"a"})
		if lr9.Left != "a" { t.Fatal() }
		_ = corestr.InvalidLeftRight("msg")
		_ = corestr.InvalidLeftRightNoMessage()
	})
}

// ═══ LeftMiddleRight ═══

func Test_C40_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_C40_LeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		if lmr.Left != "l" { t.Fatal() }
		if lmr.Middle != "m" { t.Fatal() }
		if lmr.Right != "r" { t.Fatal() }
		if lmr.IsLeftEmpty() { t.Fatal() }
		if lmr.IsMiddleEmpty() { t.Fatal() }
		if lmr.IsRightEmpty() { t.Fatal() }
		if lmr.IsLeftWhitespace() { t.Fatal() }
		if lmr.IsMiddleWhitespace() { t.Fatal() }
		if lmr.IsRightWhitespace() { t.Fatal() }
		if !lmr.HasValidNonEmptyLeft() { t.Fatal() }
		if !lmr.HasValidNonEmptyRight() { t.Fatal() }
		if !lmr.HasValidNonEmptyMiddle() { t.Fatal() }
		if !lmr.HasValidNonWhitespaceLeft() { t.Fatal() }
		if !lmr.HasValidNonWhitespaceRight() { t.Fatal() }
		if !lmr.HasValidNonWhitespaceMiddle() { t.Fatal() }
		if !lmr.HasSafeNonEmpty() { t.Fatal() }
		if !lmr.IsAll("l", "m", "r") { t.Fatal() }
		if !lmr.Is("l", "r") { t.Fatal() }
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
		r := corestr.CloneSlice([]string{"a", "b"})
		if len(r) != 2 { t.Fatal() }
		r2 := corestr.CloneSlice(nil)
		if len(r2) != 0 { t.Fatal() }
	})
}

func Test_C40_CloneSliceIf(t *testing.T) {
	safeTest(t, "Test_C40_CloneSliceIf", func() {
		r := corestr.CloneSliceIf(true, "a", "b")
		if len(r) != 2 { t.Fatal() }
		r2 := corestr.CloneSliceIf(false, "a")
		if len(r2) != 1 { t.Fatal() }
		r3 := corestr.CloneSliceIf(true)
		if len(r3) != 0 { t.Fatal() }
	})
}

func Test_C40_AnyToString(t *testing.T) {
	safeTest(t, "Test_C40_AnyToString", func() {
		r := corestr.AnyToString(false, "hello")
		if r == "" { t.Fatal() }
		r2 := corestr.AnyToString(true, "hello")
		if r2 == "" { t.Fatal() }
		r3 := corestr.AnyToString(false, "")
		if r3 != "" { t.Fatal() }
	})
}

func Test_C40_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C40_AllIndividualStringsOfStringsLength", func() {
		s := [][]string{{"a", "b"}, {"c"}}
		r := corestr.AllIndividualStringsOfStringsLength(&s)
		if r != 3 { t.Fatal() }
		if corestr.AllIndividualStringsOfStringsLength(nil) != 0 { t.Fatal() }
	})
}

func Test_C40_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C40_AllIndividualsLengthOfSimpleSlices", func() {
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("c")
		r := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
		if r != 3 { t.Fatal() }
		if corestr.AllIndividualsLengthOfSimpleSlices() != 0 { t.Fatal() }
	})
}

func Test_C40_StringUtils(t *testing.T) {
	safeTest(t, "Test_C40_StringUtils", func() {
		u := corestr.StringUtils
		if u.WrapDouble("a") != `"a"` { t.Fatal() }
		if u.WrapSingle("a") != `'a'` { t.Fatal() }
		if u.WrapTilda("a") != "`a`" { t.Fatal() }
		if u.WrapDoubleIfMissing(`"a"`) != `"a"` { t.Fatal() }
		if u.WrapDoubleIfMissing("a") != `"a"` { t.Fatal() }
		if u.WrapDoubleIfMissing("") != `""` { t.Fatal() }
		if u.WrapSingleIfMissing("'a'") != "'a'" { t.Fatal() }
		if u.WrapSingleIfMissing("a") != "'a'" { t.Fatal() }
		if u.WrapSingleIfMissing("") != "''" { t.Fatal() }
	})
}
