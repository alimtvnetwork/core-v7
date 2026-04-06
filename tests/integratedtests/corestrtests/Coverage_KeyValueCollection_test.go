package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair + KeyValueCollection — Segment 17 Part 2
// ══════════════════════════════════════════════════════════════════════════════

// --- KeyValuePair ---

func Test_CovKVP_01_Basic(t *testing.T) {
	safeTest(t, "Test_CovKVP_01_Basic", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}
		actual := args.Map{"result": kv.KeyName() != "name"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected name", actual)
		actual := args.Map{"result": kv.VariableName() != "name"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected name", actual)
		actual := args.Map{"result": kv.ValueString() != "alice"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected alice", actual)
		actual := args.Map{"result": kv.IsVariableNameEqual("name")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.IsValueEqual("alice")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovKVP_02_IsKey_IsVal_Is(t *testing.T) {
	safeTest(t, "Test_CovKVP_02_IsKey_IsVal_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.IsKey("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.IsVal("v")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.Is("k", "v")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.Is("x", "v")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovKVP_03_IsEmpty_Has(t *testing.T) {
	safeTest(t, "Test_CovKVP_03_IsEmpty_Has", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.IsKeyEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": kv.IsValueEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": kv.HasKey()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv.IsKeyValueEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": kv.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// empty
		kv2 := corestr.KeyValuePair{}
		actual := args.Map{"result": kv2.IsKeyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv2.IsKeyValueEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kv2.IsKeyValueAnyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovKVP_04_Trim(t *testing.T) {
	safeTest(t, "Test_CovKVP_04_Trim", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
		actual := args.Map{"result": kv.TrimKey() != "k"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
		actual := args.Map{"result": kv.TrimValue() != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_CovKVP_05_ValueBool(t *testing.T) {
	safeTest(t, "Test_CovKVP_05_ValueBool", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}
		actual := args.Map{"result": kv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: ""}
		actual := args.Map{"result": kv2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		kv3 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{"result": kv3.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovKVP_06_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovKVP_06_ValueInt_ValueDefInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		actual := args.Map{"result": kv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual := args.Map{"result": kv.ValueDefInt() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{"result": kv2.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
		actual := args.Map{"result": kv2.ValueDefInt() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVP_07_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovKVP_07_ValueByte_ValueDefByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		actual := args.Map{"result": kv.ValueByte(0) != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual := args.Map{"result": kv.ValueDefByte() != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		// out of range
		kv2 := corestr.KeyValuePair{Key: "k", Value: "999"}
		actual := args.Map{"result": kv2.ValueByte(5) != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
		// invalid
		kv3 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{"result": kv3.ValueByte(7) != 7}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7", actual)
	})
}

func Test_CovKVP_08_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovKVP_08_ValueFloat64_ValueDefFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		actual := args.Map{"result": kv.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual := args.Map{"result": kv.ValueDefFloat64() != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		actual := args.Map{"result": kv2.ValueFloat64(1.5) != 1.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_CovKVP_09_ValueValid_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_CovKVP_09_ValueValid_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		actual := args.Map{"result": vv.IsValid || vv.Value != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
		vv2 := kv.ValueValidOptions(false, "msg")
		actual := args.Map{"result": vv2.IsValid || vv2.Message != "msg"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with msg", actual)
	})
}

func Test_CovKVP_10_FormatString_String_Compile(t *testing.T) {
	safeTest(t, "Test_CovKVP_10_FormatString_String_Compile", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.FormatString("%s=%s") != "k=v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k=v", actual)
		_ = kv.String()
		_ = kv.Compile()
	})
}

func Test_CovKVP_11_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_CovKVP_11_Json_Serialize", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		_ = kv.Json()
		_ = kv.JsonPtr()
		_, err := kv.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		_ = kv.SerializeMust()
	})
}

func Test_CovKVP_12_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovKVP_12_Clear_Dispose", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		actual := args.Map{"result": kv.Key != "" || kv.Value != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv2.Dispose()
	})
}

// --- KeyValueCollection ---

func Test_CovKVC_01_IsEmpty_Length_Count(t *testing.T) {
	safeTest(t, "Test_CovKVC_01_IsEmpty_Length_Count", func() {
		kvc := &corestr.KeyValueCollection{}
		actual := args.Map{"result": kvc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": kvc.Count() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovKVC_02_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovKVC_02_HasAnyItem_LastIndex_HasIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		actual := args.Map{"result": kvc.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kvc.LastIndex() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": kvc.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kvc.HasIndex(1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovKVC_03_First_Last_OrDefault(t *testing.T) {
	safeTest(t, "Test_CovKVC_03_First_Last_OrDefault", func() {
		kvc := &corestr.KeyValueCollection{}
		actual := args.Map{"result": kvc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual := args.Map{"result": kvc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		actual := args.Map{"result": kvc.First().Key != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": kvc.Last().Key != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual := args.Map{"result": kvc.FirstOrDefault().Key != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": kvc.LastOrDefault().Key != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_CovKVC_04_Add_AddIf(t *testing.T) {
	safeTest(t, "Test_CovKVC_04_Add_AddIf", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.AddIf(false, "b", "2")
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kvc.AddIf(true, "b", "2")
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovKVC_05_AddStringBySplit_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_CovKVC_05_AddStringBySplit_AddStringBySplitTrim", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kvc.AddStringBySplitTrim("=", " key = value ")
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovKVC_06_Adds(t *testing.T) {
	safeTest(t, "Test_CovKVC_06_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		kvc.Adds()
	})
}

func Test_CovKVC_07_AddMap_AddHashsetMap_AddHashset(t *testing.T) {
	safeTest(t, "Test_CovKVC_07_AddMap_AddHashsetMap_AddHashset", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1"})
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kvc.AddMap(nil)
		kvc.AddHashsetMap(map[string]bool{"b": true})
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		kvc.AddHashsetMap(nil)
		hs := corestr.New.Hashset.Strings([]string{"c"})
		kvc.AddHashset(hs)
		actual := args.Map{"result": kvc.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		kvc.AddHashset(nil)
	})
}

func Test_CovKVC_08_AddsHashmap_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_CovKVC_08_AddsHashmap_AddsHashmaps", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		kvc.AddsHashmap(hm)
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		kvc.AddsHashmap(nil)
		kvc.AddsHashmaps(hm, nil)
		kvc.AddsHashmaps()
	})
}

func Test_CovKVC_09_HasKey_IsContains_Get(t *testing.T) {
	safeTest(t, "Test_CovKVC_09_HasKey_IsContains_Get", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		actual := args.Map{"result": kvc.HasKey("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": kvc.HasKey("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": kvc.IsContains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		v, ok := kvc.Get("a")
		actual := args.Map{"result": ok || v != "1"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		_, ok2 := kvc.Get("z")
		actual := args.Map{"result": ok2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovKVC_10_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	safeTest(t, "Test_CovKVC_10_AllKeys_AllKeysSorted_AllValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2")
		kvc.Add("a", "1")
		keys := kvc.AllKeys()
		actual := args.Map{"result": len(keys) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		sorted := kvc.AllKeysSorted()
		actual := args.Map{"result": sorted[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		vals := kvc.AllValues()
		actual := args.Map{"result": len(vals) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovKVC_11_SafeValueAt_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_CovKVC_11_SafeValueAt_SafeValuesAtIndexes", func() {
		kvc := &corestr.KeyValueCollection{}
		actual := args.Map{"result": kvc.SafeValueAt(0) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		actual := args.Map{"result": kvc.SafeValueAt(0) != "1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": kvc.SafeValueAt(99) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		vals := kvc.SafeValuesAtIndexes(0, 1)
		actual := args.Map{"result": len(vals) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		vals2 := kvc.SafeValuesAtIndexes()
		actual := args.Map{"result": len(vals2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVC_12_Find(t *testing.T) {
	safeTest(t, "Test_CovKVC_12_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		// empty
		r := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		actual := args.Map{"result": len(r) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		// find all
		r2 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		actual := args.Map{"result": len(r2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// break
		r3 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})
		actual := args.Map{"result": len(r3) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// skip
		r4 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, false, false
		})
		actual := args.Map{"result": len(r4) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovKVC_13_Strings_StringsUsingFormat_String(t *testing.T) {
	safeTest(t, "Test_CovKVC_13_Strings_StringsUsingFormat_String", func() {
		kvc := &corestr.KeyValueCollection{}
		actual := args.Map{"result": len(kvc.Strings()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": len(kvc.StringsUsingFormat("%s=%s")) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		kvc.Add("a", "1")
		actual := args.Map{"result": len(kvc.Strings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": len(kvc.StringsUsingFormat("%s=%s")) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = kvc.String()
		_ = kvc.Compile()
	})
}

func Test_CovKVC_14_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_CovKVC_14_Hashmap_Map", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		m := kvc.Map()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovKVC_15_Join_JoinKeys_JoinValues(t *testing.T) {
	safeTest(t, "Test_CovKVC_15_Join_JoinKeys_JoinValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		_ = kvc.Join(",")
		_ = kvc.JoinKeys(",")
		_ = kvc.JoinValues(",")
	})
}

func Test_CovKVC_16_SerializeMust(t *testing.T) {
	safeTest(t, "Test_CovKVC_16_SerializeMust", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b := kvc.SerializeMust()
		actual := args.Map{"result": len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_CovKVC_17_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovKVC_17_JsonModel_MarshalUnmarshal", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		_ = kvc.JsonModel()
		_ = kvc.JsonModelAny()
		data, err := kvc.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		kvc2 := &corestr.KeyValueCollection{}
		err2 := kvc2.UnmarshalJSON(data)
		actual := args.Map{"result": err2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// empty unmarshal
		err3 := kvc2.UnmarshalJSON([]byte("[]"))
		actual := args.Map{"result": err3 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// invalid
		err4 := kvc2.UnmarshalJSON([]byte("bad"))
		actual := args.Map{"result": err4 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovKVC_18_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovKVC_18_Json_JsonPtr_ParseInject", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		_ = kvc.Json()
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		r, err := kvc2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err != nil || r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovKVC_19_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovKVC_19_JsonParseSelfInject_AsInterfaces", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jr)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		_ = kvc.AsJsonContractsBinder()
		_ = kvc.AsJsoner()
		_ = kvc.AsJsonParseSelfInjector()
	})
}

func Test_CovKVC_20_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovKVC_20_Serialize_Deserialize", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		_, err := kvc.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := &corestr.KeyValueCollection{}
		err2 := kvc.Deserialize(target)
		actual := args.Map{"result": err2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovKVC_21_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovKVC_21_Clear_Dispose", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		kvc2 := &corestr.KeyValueCollection{}
		kvc2.Add("b", "2")
		kvc2.Dispose()
	})
}
