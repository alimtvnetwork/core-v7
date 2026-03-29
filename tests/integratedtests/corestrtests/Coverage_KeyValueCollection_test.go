package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair + KeyValueCollection — Segment 17 Part 2
// ══════════════════════════════════════════════════════════════════════════════

// --- KeyValuePair ---

func Test_CovKVP_01_Basic(t *testing.T) {
	safeTest(t, "Test_CovKVP_01_Basic", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}
		if kv.KeyName() != "name" {
			t.Fatal("expected name")
		}
		if kv.VariableName() != "name" {
			t.Fatal("expected name")
		}
		if kv.ValueString() != "alice" {
			t.Fatal("expected alice")
		}
		if !kv.IsVariableNameEqual("name") {
			t.Fatal("expected true")
		}
		if !kv.IsValueEqual("alice") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovKVP_02_IsKey_IsVal_Is(t *testing.T) {
	safeTest(t, "Test_CovKVP_02_IsKey_IsVal_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if !kv.IsKey("k") {
			t.Fatal("expected true")
		}
		if !kv.IsVal("v") {
			t.Fatal("expected true")
		}
		if !kv.Is("k", "v") {
			t.Fatal("expected true")
		}
		if kv.Is("x", "v") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovKVP_03_IsEmpty_Has(t *testing.T) {
	safeTest(t, "Test_CovKVP_03_IsEmpty_Has", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.IsKeyEmpty() {
			t.Fatal("expected false")
		}
		if kv.IsValueEmpty() {
			t.Fatal("expected false")
		}
		if !kv.HasKey() {
			t.Fatal("expected true")
		}
		if !kv.HasValue() {
			t.Fatal("expected true")
		}
		if kv.IsKeyValueEmpty() {
			t.Fatal("expected false")
		}
		if kv.IsKeyValueAnyEmpty() {
			t.Fatal("expected false")
		}
		// empty
		kv2 := corestr.KeyValuePair{}
		if !kv2.IsKeyEmpty() {
			t.Fatal("expected true")
		}
		if !kv2.IsKeyValueEmpty() {
			t.Fatal("expected true")
		}
		if !kv2.IsKeyValueAnyEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_CovKVP_04_Trim(t *testing.T) {
	safeTest(t, "Test_CovKVP_04_Trim", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
		if kv.TrimKey() != "k" {
			t.Fatal("expected k")
		}
		if kv.TrimValue() != "v" {
			t.Fatal("expected v")
		}
	})
}

func Test_CovKVP_05_ValueBool(t *testing.T) {
	safeTest(t, "Test_CovKVP_05_ValueBool", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}
		if !kv.ValueBool() {
			t.Fatal("expected true")
		}
		kv2 := corestr.KeyValuePair{Key: "k", Value: ""}
		if kv2.ValueBool() {
			t.Fatal("expected false")
		}
		kv3 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		if kv3.ValueBool() {
			t.Fatal("expected false")
		}
	})
}

func Test_CovKVP_06_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovKVP_06_ValueInt_ValueDefInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		if kv.ValueInt(0) != 42 {
			t.Fatal("expected 42")
		}
		if kv.ValueDefInt() != 42 {
			t.Fatal("expected 42")
		}
		kv2 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		if kv2.ValueInt(99) != 99 {
			t.Fatal("expected 99")
		}
		if kv2.ValueDefInt() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVP_07_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovKVP_07_ValueByte_ValueDefByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		if kv.ValueByte(0) != 100 {
			t.Fatal("expected 100")
		}
		if kv.ValueDefByte() != 100 {
			t.Fatal("expected 100")
		}
		// out of range
		kv2 := corestr.KeyValuePair{Key: "k", Value: "999"}
		if kv2.ValueByte(5) != 5 {
			t.Fatal("expected 5")
		}
		// invalid
		kv3 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		if kv3.ValueByte(7) != 7 {
			t.Fatal("expected 7")
		}
	})
}

func Test_CovKVP_08_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovKVP_08_ValueFloat64_ValueDefFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		if kv.ValueFloat64(0) != 3.14 {
			t.Fatal("expected 3.14")
		}
		if kv.ValueDefFloat64() != 3.14 {
			t.Fatal("expected 3.14")
		}
		kv2 := corestr.KeyValuePair{Key: "k", Value: "abc"}
		if kv2.ValueFloat64(1.5) != 1.5 {
			t.Fatal("expected 1.5")
		}
	})
}

func Test_CovKVP_09_ValueValid_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_CovKVP_09_ValueValid_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		if !vv.IsValid || vv.Value != "v" {
			t.Fatal("expected valid")
		}
		vv2 := kv.ValueValidOptions(false, "msg")
		if vv2.IsValid || vv2.Message != "msg" {
			t.Fatal("expected invalid with msg")
		}
	})
}

func Test_CovKVP_10_FormatString_String_Compile(t *testing.T) {
	safeTest(t, "Test_CovKVP_10_FormatString_String_Compile", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.FormatString("%s=%s") != "k=v" {
			t.Fatal("expected k=v")
		}
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
		if err != nil {
			t.Fatal("unexpected error")
		}
		_ = kv.SerializeMust()
	})
}

func Test_CovKVP_12_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovKVP_12_Clear_Dispose", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		if kv.Key != "" || kv.Value != "" {
			t.Fatal("expected cleared")
		}
		kv2 := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv2.Dispose()
	})
}

// --- KeyValueCollection ---

func Test_CovKVC_01_IsEmpty_Length_Count(t *testing.T) {
	safeTest(t, "Test_CovKVC_01_IsEmpty_Length_Count", func() {
		kvc := &corestr.KeyValueCollection{}
		if !kvc.IsEmpty() {
			t.Fatal("expected empty")
		}
		if kvc.Length() != 0 {
			t.Fatal("expected 0")
		}
		if kvc.Count() != 0 {
			t.Fatal("expected 0")
		}
		kvc.Add("k", "v")
		if kvc.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovKVC_02_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovKVC_02_HasAnyItem_LastIndex_HasIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.HasAnyItem() {
			t.Fatal("expected false")
		}
		kvc.Add("k", "v")
		if !kvc.HasAnyItem() {
			t.Fatal("expected true")
		}
		if kvc.LastIndex() != 0 {
			t.Fatal("expected 0")
		}
		if !kvc.HasIndex(0) {
			t.Fatal("expected true")
		}
		if kvc.HasIndex(1) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovKVC_03_First_Last_OrDefault(t *testing.T) {
	safeTest(t, "Test_CovKVC_03_First_Last_OrDefault", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.FirstOrDefault() != nil {
			t.Fatal("expected nil")
		}
		if kvc.LastOrDefault() != nil {
			t.Fatal("expected nil")
		}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.First().Key != "a" {
			t.Fatal("expected a")
		}
		if kvc.Last().Key != "b" {
			t.Fatal("expected b")
		}
		if kvc.FirstOrDefault().Key != "a" {
			t.Fatal("expected a")
		}
		if kvc.LastOrDefault().Key != "b" {
			t.Fatal("expected b")
		}
	})
}

func Test_CovKVC_04_Add_AddIf(t *testing.T) {
	safeTest(t, "Test_CovKVC_04_Add_AddIf", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.AddIf(false, "b", "2")
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
		kvc.AddIf(true, "b", "2")
		if kvc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovKVC_05_AddStringBySplit_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_CovKVC_05_AddStringBySplit_AddStringBySplitTrim", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
		kvc.AddStringBySplitTrim("=", " key = value ")
		if kvc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovKVC_06_Adds(t *testing.T) {
	safeTest(t, "Test_CovKVC_06_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		if kvc.Length() != 2 {
			t.Fatal("expected 2")
		}
		kvc.Adds()
	})
}

func Test_CovKVC_07_AddMap_AddHashsetMap_AddHashset(t *testing.T) {
	safeTest(t, "Test_CovKVC_07_AddMap_AddHashsetMap_AddHashset", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1"})
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
		kvc.AddMap(nil)
		kvc.AddHashsetMap(map[string]bool{"b": true})
		if kvc.Length() != 2 {
			t.Fatal("expected 2")
		}
		kvc.AddHashsetMap(nil)
		hs := corestr.New.Hashset.Strings([]string{"c"})
		kvc.AddHashset(hs)
		if kvc.Length() != 3 {
			t.Fatal("expected 3")
		}
		kvc.AddHashset(nil)
	})
}

func Test_CovKVC_08_AddsHashmap_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_CovKVC_08_AddsHashmap_AddsHashmaps", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("a", "1")
		kvc.AddsHashmap(hm)
		if kvc.Length() != 1 {
			t.Fatal("expected 1")
		}
		kvc.AddsHashmap(nil)
		kvc.AddsHashmaps(hm, nil)
		kvc.AddsHashmaps()
	})
}

func Test_CovKVC_09_HasKey_IsContains_Get(t *testing.T) {
	safeTest(t, "Test_CovKVC_09_HasKey_IsContains_Get", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if !kvc.HasKey("a") {
			t.Fatal("expected true")
		}
		if kvc.HasKey("z") {
			t.Fatal("expected false")
		}
		if !kvc.IsContains("a") {
			t.Fatal("expected true")
		}
		v, ok := kvc.Get("a")
		if !ok || v != "1" {
			t.Fatal("expected found")
		}
		_, ok2 := kvc.Get("z")
		if ok2 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovKVC_10_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	safeTest(t, "Test_CovKVC_10_AllKeys_AllKeysSorted_AllValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2")
		kvc.Add("a", "1")
		keys := kvc.AllKeys()
		if len(keys) != 2 {
			t.Fatal("expected 2")
		}
		sorted := kvc.AllKeysSorted()
		if sorted[0] != "a" {
			t.Fatal("expected a first")
		}
		vals := kvc.AllValues()
		if len(vals) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovKVC_11_SafeValueAt_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_CovKVC_11_SafeValueAt_SafeValuesAtIndexes", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.SafeValueAt(0) != "" {
			t.Fatal("expected empty")
		}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.SafeValueAt(0) != "1" {
			t.Fatal("expected 1")
		}
		if kvc.SafeValueAt(99) != "" {
			t.Fatal("expected empty")
		}
		vals := kvc.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 {
			t.Fatal("expected 2")
		}
		vals2 := kvc.SafeValuesAtIndexes()
		if len(vals2) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVC_12_Find(t *testing.T) {
	safeTest(t, "Test_CovKVC_12_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		// empty
		r := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		// find all
		r2 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		if len(r2) != 2 {
			t.Fatal("expected 2")
		}
		// break
		r3 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})
		if len(r3) != 1 {
			t.Fatal("expected 1")
		}
		// skip
		r4 := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, false, false
		})
		if len(r4) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovKVC_13_Strings_StringsUsingFormat_String(t *testing.T) {
	safeTest(t, "Test_CovKVC_13_Strings_StringsUsingFormat_String", func() {
		kvc := &corestr.KeyValueCollection{}
		if len(kvc.Strings()) != 0 {
			t.Fatal("expected 0")
		}
		if len(kvc.StringsUsingFormat("%s=%s")) != 0 {
			t.Fatal("expected 0")
		}
		kvc.Add("a", "1")
		if len(kvc.Strings()) != 1 {
			t.Fatal("expected 1")
		}
		if len(kvc.StringsUsingFormat("%s=%s")) != 1 {
			t.Fatal("expected 1")
		}
		_ = kvc.String()
		_ = kvc.Compile()
	})
}

func Test_CovKVC_14_Hashmap_Map(t *testing.T) {
	safeTest(t, "Test_CovKVC_14_Hashmap_Map", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
		m := kvc.Map()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
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
		if len(b) == 0 {
			t.Fatal("expected bytes")
		}
	})
}

func Test_CovKVC_17_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovKVC_17_JsonModel_MarshalUnmarshal", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		_ = kvc.JsonModel()
		_ = kvc.JsonModelAny()
		data, err := kvc.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		kvc2 := &corestr.KeyValueCollection{}
		err2 := kvc2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		// empty unmarshal
		err3 := kvc2.UnmarshalJSON([]byte("[]"))
		if err3 != nil {
			t.Fatal("unexpected error")
		}
		// invalid
		err4 := kvc2.UnmarshalJSON([]byte("bad"))
		if err4 == nil {
			t.Fatal("expected error")
		}
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
		if err != nil || r == nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovKVC_19_JsonParseSelfInject_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovKVC_19_JsonParseSelfInject_AsInterfaces", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
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
		if err != nil {
			t.Fatal("unexpected error")
		}
		target := &corestr.KeyValueCollection{}
		err2 := kvc.Deserialize(target)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovKVC_21_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovKVC_21_Clear_Dispose", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()
		if kvc.Length() != 0 {
			t.Fatal("expected 0")
		}
		kvc2 := &corestr.KeyValueCollection{}
		kvc2.Add("b", "2")
		kvc2.Dispose()
	})
}
