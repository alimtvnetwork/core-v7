package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Segment 18 Part 2
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovVV_01_Constructors(t *testing.T) {
	safeTest(t, "Test_CovVV_01_Constructors", func() {
		vv := corestr.NewValidValue("hello")
		if vv.Value != "hello" || !vv.IsValid {
			t.Fatal("expected valid hello")
		}
		ve := corestr.NewValidValueEmpty()
		if ve.Value != "" || !ve.IsValid {
			t.Fatal("expected valid empty")
		}
		iv := corestr.InvalidValidValue("msg")
		if iv.IsValid || iv.Message != "msg" {
			t.Fatal("expected invalid with msg")
		}
		ivn := corestr.InvalidValidValueNoMessage()
		if ivn.IsValid || ivn.Message != "" {
			t.Fatal("expected invalid no msg")
		}
	})
}

func Test_CovVV_02_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_CovVV_02_NewValidValueUsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "test")
		if !vv.IsValid {
			t.Fatal("expected valid")
		}
	})
}

func Test_CovVV_03_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_CovVV_03_NewValidValueUsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		_ = vv
	})
}
func Test_CovVV_05_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_CovVV_05_IsEmpty_IsWhitespace", func() {
		vv := corestr.NewValidValue("")
		if !vv.IsEmpty() {
			t.Fatal("expected empty")
		}
		if !vv.IsWhitespace() {
			t.Fatal("expected whitespace")
		}
		vv2 := corestr.NewValidValue("hi")
		if vv2.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovVV_06_Trim(t *testing.T) {
	safeTest(t, "Test_CovVV_06_Trim", func() {
		vv := corestr.NewValidValue("  hi  ")
		if vv.Trim() != "hi" {
			t.Fatal("expected hi")
		}
	})
}

func Test_CovVV_07_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovVV_07_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("hi")
		if !vv.HasValidNonEmpty() {
			t.Fatal("expected true")
		}
		if !vv.HasValidNonWhitespace() {
			t.Fatal("expected true")
		}
		if !vv.HasSafeNonEmpty() {
			t.Fatal("expected true")
		}
		iv := corestr.InvalidValidValue("")
		if iv.HasValidNonEmpty() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovVV_08_ValueBool(t *testing.T) {
	safeTest(t, "Test_CovVV_08_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		if !vv.ValueBool() {
			t.Fatal("expected true")
		}
		vv2 := corestr.NewValidValue("")
		if vv2.ValueBool() {
			t.Fatal("expected false")
		}
		vv3 := corestr.NewValidValue("abc")
		if vv3.ValueBool() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovVV_09_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovVV_09_ValueInt_ValueDefInt", func() {
		vv := corestr.NewValidValue("42")
		if vv.ValueInt(0) != 42 {
			t.Fatal("expected 42")
		}
		if vv.ValueDefInt() != 42 {
			t.Fatal("expected 42")
		}
		vv2 := corestr.NewValidValue("abc")
		if vv2.ValueInt(99) != 99 {
			t.Fatal("expected 99")
		}
	})
}

func Test_CovVV_10_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovVV_10_ValueByte_ValueDefByte", func() {
		vv := corestr.NewValidValue("100")
		if vv.ValueByte(0) != 100 {
			t.Fatal("expected 100")
		}
		if vv.ValueDefByte() != 100 {
			t.Fatal("expected 100")
		}
		// out of range
		vv2 := corestr.NewValidValue("999")
		if vv2.ValueByte(5) != 255 {
			t.Fatal("expected 255 (max)")
		}
		// negative
		vv3 := corestr.NewValidValue("-1")
		if vv3.ValueByte(5) != 0 {
			t.Fatal("expected 0")
		}
		// invalid
		vv4 := corestr.NewValidValue("abc")
		if vv4.ValueByte(7) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovVV_11_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovVV_11_ValueFloat64_ValueDefFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		if vv.ValueFloat64(0) != 3.14 {
			t.Fatal("expected 3.14")
		}
		if vv.ValueDefFloat64() != 3.14 {
			t.Fatal("expected 3.14")
		}
		vv2 := corestr.NewValidValue("abc")
		if vv2.ValueFloat64(1.5) != 1.5 {
			t.Fatal("expected 1.5")
		}
	})
}

func Test_CovVV_12_Is_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_CovVV_12_Is_IsAnyOf", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.Is("hello") {
			t.Fatal("expected true")
		}
		if vv.Is("world") {
			t.Fatal("expected false")
		}
		if !vv.IsAnyOf("a", "hello") {
			t.Fatal("expected true")
		}
		if vv.IsAnyOf("a", "b") {
			t.Fatal("expected false")
		}
		if !vv.IsAnyOf() {
			t.Fatal("expected true for empty")
		}
	})
}

func Test_CovVV_13_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_CovVV_13_IsContains_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		if !vv.IsContains("world") {
			t.Fatal("expected true")
		}
		if !vv.IsAnyContains("xyz", "world") {
			t.Fatal("expected true")
		}
		if vv.IsAnyContains("xyz", "abc") {
			t.Fatal("expected false")
		}
		if !vv.IsAnyContains() {
			t.Fatal("expected true for empty")
		}
	})
}

func Test_CovVV_14_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_CovVV_14_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		if !vv.IsEqualNonSensitive("hello") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovVV_15_Regex(t *testing.T) {
	safeTest(t, "Test_CovVV_15_Regex", func() {
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		if !vv.IsRegexMatches(re) {
			t.Fatal("expected true")
		}
		if vv.IsRegexMatches(nil) {
			t.Fatal("expected false")
		}
		if vv.RegexFindString(re) != "123" {
			t.Fatal("expected 123")
		}
		if vv.RegexFindString(nil) != "" {
			t.Fatal("expected empty")
		}
		items := vv.RegexFindAllStrings(re, -1)
		if len(items) != 1 {
			t.Fatal("expected 1")
		}
		if len(vv.RegexFindAllStrings(nil, -1)) != 0 {
			t.Fatal("expected 0")
		}
		items2, has := vv.RegexFindAllStringsWithFlag(re, -1)
		if !has || len(items2) != 1 {
			t.Fatal("expected found")
		}
		_, has2 := vv.RegexFindAllStringsWithFlag(nil, -1)
		if has2 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovVV_16_Split_SplitNonEmpty_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovVV_16_Split_SplitNonEmpty_SplitTrimNonWhitespace", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		if len(parts) != 3 {
			t.Fatal("expected 3")
		}
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
		vv := corestr.NewValidValue("hello")
		c := vv.Clone()
		if c.Value != "hello" {
			t.Fatal("expected hello")
		}
		// nil clone
		var nilVV *corestr.ValidValue
		if nilVV.Clone() != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovVV_18_String_FullString(t *testing.T) {
	safeTest(t, "Test_CovVV_18_String_FullString", func() {
		vv := corestr.NewValidValue("hello")
		if vv.String() != "hello" {
			t.Fatal("expected hello")
		}
		_ = vv.FullString()
		// nil
		var nilVV *corestr.ValidValue
		if nilVV.String() != "" {
			t.Fatal("expected empty")
		}
		if nilVV.FullString() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovVV_19_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovVV_19_Clear_Dispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		if vv.Value != "" {
			t.Fatal("expected empty")
		}
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
		vv := corestr.NewValidValue("hello")
		_ = vv.Json()
		jr := vv.JsonPtr()
		vv2 := &corestr.ValidValue{}
		r, err := vv2.ParseInjectUsingJson(jr)
		if err != nil || r == nil {
			t.Fatal("unexpected error")
		}
		_, err2 := vv.Serialize()
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		target := &corestr.ValidValue{}
		err3 := vv.Deserialize(target)
		if err3 != nil {
			t.Fatal("unexpected error")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues — Segment 19 Part 1
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovVVs_01_Constructors(t *testing.T) {
	safeTest(t, "Test_CovVVs_01_Constructors", func() {
		vvs := corestr.EmptyValidValues()
		if vvs.Length() != 0 {
			t.Fatal("expected 0")
		}
		vvs2 := corestr.NewValidValues(5)
		if vvs2.Length() != 0 {
			t.Fatal("expected 0")
		}
		vv1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv2 := corestr.ValidValue{Value: "b", IsValid: true}
		vvs3 := corestr.NewValidValuesUsingValues(vv1, vv2)
		if vvs3.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty values
		vvs4 := corestr.NewValidValuesUsingValues()
		if vvs4.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovVVs_02_Count_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovVVs_02_Count_HasAnyItem_LastIndex_HasIndex", func() {
		vvs := corestr.EmptyValidValues()
		if vvs.Count() != 0 {
			t.Fatal("expected 0")
		}
		if vvs.HasAnyItem() {
			t.Fatal("expected false")
		}
		vvs.Add("a")
		if vvs.Count() != 1 {
			t.Fatal("expected 1")
		}
		if !vvs.HasAnyItem() {
			t.Fatal("expected true")
		}
		if vvs.LastIndex() != 0 {
			t.Fatal("expected 0")
		}
		if !vvs.HasIndex(0) {
			t.Fatal("expected true")
		}
		if vvs.HasIndex(1) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovVVs_03_Add_AddFull(t *testing.T) {
	safeTest(t, "Test_CovVVs_03_Add_AddFull", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
		vvs.AddFull(false, "b", "msg")
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovVVs_04_Adds_AddsPtr(t *testing.T) {
	safeTest(t, "Test_CovVVs_04_Adds_AddsPtr", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
		vvs.Adds()
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
		vv := corestr.NewValidValue("c")
		vvs.AddsPtr(vv)
		if vvs.Length() != 3 {
			t.Fatal("expected 3")
		}
		vvs.AddsPtr()
		if vvs.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovVVs_05_AddValidValues(t *testing.T) {
	safeTest(t, "Test_CovVVs_05_AddValidValues", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		vvs.AddValidValues(vvs2)
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
		vvs.AddValidValues(nil)
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
		vvs.AddValidValues(corestr.EmptyValidValues())
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovVVs_06_AddHashsetMap_AddHashset(t *testing.T) {
	safeTest(t, "Test_CovVVs_06_AddHashsetMap_AddHashset", func() {
		vvs := corestr.EmptyValidValues()
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
		vvs.AddHashsetMap(nil)
		hs := corestr.New.Hashset.Strings([]string{"c"})
		vvs.AddHashset(hs)
		if vvs.Length() != 3 {
			t.Fatal("expected 3")
		}
		vvs.AddHashset(nil)
	})
}

func Test_CovVVs_07_ConcatNew(t *testing.T) {
	safeTest(t, "Test_CovVVs_07_ConcatNew", func() {
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		// no args, clone
		c := vvs.ConcatNew(true)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		// no args, no clone
		c2 := vvs.ConcatNew(false)
		if c2.Length() != 1 {
			t.Fatal("expected 1")
		}
		// with args
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")
		c3 := vvs.ConcatNew(true, vvs2)
		if c3.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovVVs_08_Find(t *testing.T) {
	safeTest(t, "Test_CovVVs_08_Find", func() {
		vvs := corestr.EmptyValidValues()
		// empty find
		r := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		vvs.Add("a")
		vvs.Add("b")
		// find all
		r2 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
		// break
		r3 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, true
		})
		if len(r3) != 1 {
			t.Fatal("expected 1")
		}
		// skip
		r4 := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, false, false
		})
		if len(r4) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovVVs_09_SafeValueAt_SafeValidValueAt_SafeIndexes(t *testing.T) {
	safeTest(t, "Test_CovVVs_09_SafeValueAt_SafeValidValueAt_SafeIndexes", func() {
		vvs := corestr.EmptyValidValues()
		if vvs.SafeValueAt(0) != "" {
			t.Fatal("expected empty")
		}
		if vvs.SafeValidValueAt(0) != "" {
			t.Fatal("expected empty")
		}
		vvs.Add("a")
		vvs.AddFull(false, "b", "")
		if vvs.SafeValueAt(0) != "a" {
			t.Fatal("expected a")
		}
		if vvs.SafeValueAt(99) != "" {
			t.Fatal("expected empty")
		}
		if vvs.SafeValidValueAt(0) != "a" {
			t.Fatal("expected a")
		}
		// invalid valid value returns empty
		if vvs.SafeValidValueAt(1) != "" {
			t.Fatal("expected empty for invalid")
		}
		vals := vvs.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 {
			t.Fatal("expected 2")
		}
		vals2 := vvs.SafeValuesAtIndexes()
		if len(vals2) != 0 {
			t.Fatal("expected 0")
		}
		vals3 := vvs.SafeValidValuesAtIndexes(0)
		if len(vals3) != 1 {
			t.Fatal("expected 1")
		}
		vals4 := vvs.SafeValidValuesAtIndexes()
		if len(vals4) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovVVs_10_Strings_FullStrings_String(t *testing.T) {
	safeTest(t, "Test_CovVVs_10_Strings_FullStrings_String", func() {
		vvs := corestr.EmptyValidValues()
		if len(vvs.Strings()) != 0 {
			t.Fatal("expected 0")
		}
		if len(vvs.FullStrings()) != 0 {
			t.Fatal("expected 0")
		}
		vvs.Add("a")
		if len(vvs.Strings()) != 1 {
			t.Fatal("expected 1")
		}
		if len(vvs.FullStrings()) != 1 {
			t.Fatal("expected 1")
		}
		_ = vvs.String()
	})
}

func Test_CovVVs_11_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_CovVVs_11_Hashmap_Map", func() {
		vvs := corestr.EmptyValidValues()
		hm := vvs.Hashmap()
		if hm.Length() != 0 {
			t.Fatal("expected 0")
		}
		vvs.Add("a")
		hm2 := vvs.Hashmap()
		if hm2.Length() != 1 {
			t.Fatal("expected 1")
		}
		m := vvs.Map()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovVVs_12_IsEmpty(t *testing.T) {
	safeTest(t, "Test_CovVVs_12_IsEmpty", func() {
		vvs := corestr.EmptyValidValues()
		if !vvs.IsEmpty() {
			t.Fatal("expected empty")
		}
		vvs.Add("a")
		if vvs.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollectionNode + NonChainedLinkedCollectionNodes — Segment 19 Part 2
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovLCN_01_IsEmpty_HasElement_HasNext(t *testing.T) {
	safeTest(t, "Test_CovLCN_01_IsEmpty_HasElement_HasNext", func() {
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		if node.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !node.HasElement() {
			t.Fatal("expected has element")
		}
		if node.HasNext() {
			t.Fatal("expected no next")
		}
	})
}

func Test_CovLCN_02_EndOfChain(t *testing.T) {
	safeTest(t, "Test_CovLCN_02_EndOfChain", func() {
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		end, length := node.EndOfChain()
		if end != node || length != 1 {
			t.Fatal("expected self, length 1")
		}
	})
}

func Test_CovLCN_03_Clone(t *testing.T) {
	safeTest(t, "Test_CovLCN_03_Clone", func() {
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		c := node.Clone()
		if c.HasNext() {
			t.Fatal("expected no next")
		}
	})
}

func Test_CovLCN_04_IsEqual_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_CovLCN_04_IsEqual_IsChainEqual", func() {
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n1 := &corestr.LinkedCollectionNode{Element: col}
		n2 := &corestr.LinkedCollectionNode{Element: col}
		if !n1.IsEqual(n2) {
			t.Fatal("expected equal")
		}
		if !n1.IsChainEqual(n2) {
			t.Fatal("expected chain equal")
		}
		// same ptr
		if !n1.IsChainEqual(n1) {
			t.Fatal("expected equal same ptr")
		}
		// nil
		var nilN *corestr.LinkedCollectionNode
		if !nilN.IsEqual(nil) {
			t.Fatal("expected equal nil")
		}
		if nilN.IsEqual(n1) {
			t.Fatal("expected not equal")
		}
		if n1.IsChainEqual(nil) {
			t.Fatal("expected not equal")
		}
	})
}

func Test_CovLCN_05_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_CovLCN_05_IsEqualValue", func() {
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n := &corestr.LinkedCollectionNode{Element: col}
		if !n.IsEqualValue(col) {
			t.Fatal("expected equal")
		}
		if n.IsEqualValue(nil) {
			t.Fatal("expected not equal")
		}
	})
}
func Test_CovLCN_07_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_CovLCN_07_CreateLinkedList", func() {
		col := corestr.New.Collection.Strings([]string{"false", "a"})
		n := &corestr.LinkedCollectionNode{Element: col}
		ll := n.CreateLinkedList()
		if ll.Length() != 1 {
			t.Fatalf("expected 1, got %d", ll.Length())
		}
	})
}

// --- NonChainedLinkedCollectionNodes ---

func Test_CovNCLCN_01_Basic(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_01_Basic", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		if !nc.IsEmpty() {
			t.Fatal("expected empty")
		}
		if nc.HasItems() {
			t.Fatal("expected no items")
		}
		if nc.Length() != 0 {
			t.Fatal("expected 0")
		}
		if nc.IsChainingApplied() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovNCLCN_02_Adds_First_Last(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_02_Adds_First_Last", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)
		if nc.Length() != 2 {
			t.Fatal("expected 2")
		}
		if nc.First() != n1 {
			t.Fatal("expected n1")
		}
		if nc.Last() != n2 {
			t.Fatal("expected n2")
		}
		nc.Adds(nil)
		if nc.Items() == nil {
			t.Fatal("expected non-nil items")
		}
	})
}

func Test_CovNCLCN_03_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_03_FirstOrDefault_LastOrDefault", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		if nc.FirstOrDefault() != nil {
			t.Fatal("expected nil")
		}
		if nc.LastOrDefault() != nil {
			t.Fatal("expected nil")
		}
		n := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"false", "a"})}
		nc.Adds(n)
		if nc.FirstOrDefault() != n {
			t.Fatal("expected n")
		}
		if nc.LastOrDefault() != n {
			t.Fatal("expected n")
		}
	})
}

func Test_CovNCLCN_04_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_04_ApplyChaining", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		// empty apply
		nc.ApplyChaining()
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)
		nc.ApplyChaining()
		if !nc.IsChainingApplied() {
			t.Fatal("expected true")
		}
		if !n1.HasNext() {
			t.Fatal("expected n1 has next")
		}
		// re-apply should be no-op
		nc.ApplyChaining()
	})
}

func Test_CovNCLCN_05_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_CovNCLCN_05_ToChainedNodes", func() {
		nc := corestr.NewNonChainedLinkedCollectionNodes(5)
		// empty
		cn := nc.ToChainedNodes()
		if cn == nil {
			t.Fatal("expected non-nil")
		}
		col1 := corestr.New.Collection.Strings([]string{"false", "a"})
		col2 := corestr.New.Collection.Strings([]string{"false", "b"})
		n1 := &corestr.LinkedCollectionNode{Element: col1}
		n2 := &corestr.LinkedCollectionNode{Element: col2}
		nc.Adds(n1, n2)
		cn2 := nc.ToChainedNodes()
		if cn2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}
