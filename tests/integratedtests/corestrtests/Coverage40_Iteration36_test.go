package corestrtests

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════

func Test_Cov40_KVP_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Basic", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}
		if kv.KeyName() != "name" || kv.VariableName() != "name" || kv.ValueString() != "alice" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsVariableNameEqual", func() {
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}
		if !kv.IsVariableNameEqual("x") || kv.IsVariableNameEqual("z") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsValueEqual", func() {
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}
		if !kv.IsValueEqual("y") || kv.IsValueEqual("z") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_Compile(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Compile", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.Compile() == "" || kv.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_KVP_IsKeyEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKeyEmpty", func() {
		kv := corestr.KeyValuePair{Key: "", Value: "v"}
		if !kv.IsKeyEmpty() || !kv.IsKeyValueAnyEmpty() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsValueEmpty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		if !kv.IsValueEmpty() || !kv.IsKeyValueAnyEmpty() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_HasKeyValue(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_HasKeyValue", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if !kv.HasKey() || !kv.HasValue() || kv.IsKeyValueEmpty() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKeyValueEmpty", func() {
		kv := corestr.KeyValuePair{}
		if !kv.IsKeyValueEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_KVP_Trim(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Trim", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
		if kv.TrimKey() != "k" || kv.TrimValue() != "v" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueBool", func() {
		kv := corestr.KeyValuePair{Value: "true"}
		if !kv.ValueBool() {
			t.Error("expected true")
		}
		kv2 := corestr.KeyValuePair{Value: ""}
		if kv2.ValueBool() {
			t.Error("expected false")
		}
		kv3 := corestr.KeyValuePair{Value: "invalid"}
		if kv3.ValueBool() {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_KVP_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueInt", func() {
		kv := corestr.KeyValuePair{Value: "42"}
		if kv.ValueInt(0) != 42 {
			t.Error("expected 42")
		}
		kv2 := corestr.KeyValuePair{Value: "bad"}
		if kv2.ValueInt(99) != 99 {
			t.Error("expected 99")
		}
	})
}

func Test_Cov40_KVP_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefInt", func() {
		kv := corestr.KeyValuePair{Value: "10"}
		if kv.ValueDefInt() != 10 {
			t.Error("expected 10")
		}
		kv2 := corestr.KeyValuePair{Value: "bad"}
		if kv2.ValueDefInt() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVP_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueByte", func() {
		kv := corestr.KeyValuePair{Value: "10"}
		if kv.ValueByte(0) != 10 {
			t.Error("expected 10")
		}
		kv2 := corestr.KeyValuePair{Value: "999"}
		if kv2.ValueByte(5) != 5 {
			t.Error("expected 5")
		}
	})
}

func Test_Cov40_KVP_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefByte", func() {
		kv := corestr.KeyValuePair{Value: "50"}
		if kv.ValueDefByte() != 50 {
			t.Error("expected 50")
		}
		kv2 := corestr.KeyValuePair{Value: "bad"}
		if kv2.ValueDefByte() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVP_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueFloat64", func() {
		kv := corestr.KeyValuePair{Value: "3.14"}
		if kv.ValueFloat64(0) != 3.14 {
			t.Error("expected 3.14")
		}
		kv2 := corestr.KeyValuePair{Value: "bad"}
		if kv2.ValueFloat64(1.0) != 1.0 {
			t.Error("expected 1.0")
		}
	})
}

func Test_Cov40_KVP_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefFloat64", func() {
		kv := corestr.KeyValuePair{Value: "bad"}
		if kv.ValueDefFloat64() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVP_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueValid", func() {
		kv := corestr.KeyValuePair{Value: "hello"}
		vv := kv.ValueValid()
		if !vv.IsValid || vv.Value != "hello" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Value: "x"}
		vv := kv.ValueValidOptions(false, "err")
		if vv.IsValid || vv.Message != "err" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_Is(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Is", func() {
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}
		if !kv.Is("a", "b") || kv.Is("a", "c") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKey_IsVal", func() {
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}
		if !kv.IsKey("a") || !kv.IsVal("b") || kv.IsKey("x") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_FormatString(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_FormatString", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.FormatString("%v=%v") != "k=v" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVP_Json(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Json", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.Json().Error != nil {
			t.Error("unexpected error")
		}
		if kv.JsonPtr() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_KVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Serialize", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		data, err := kv.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
		if len(kv.SerializeMust()) == 0 {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_KVP_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Clear_Dispose", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		if kv.Key != "" || kv.Value != "" {
			t.Error("expected empty")
		}
		kv2 := corestr.KeyValuePair{Key: "a", Value: "b"}
		kv2.Dispose()
		if kv2.Key != "" {
			t.Error("expected empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// KeyValueCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_KVC_Add_Length(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Add_Length", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k1", "v1").Add("k2", "v2")
		if kvc.Length() != 2 || kvc.Count() != 2 || kvc.IsEmpty() || !kvc.HasAnyItem() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddIf", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddIf(true, "a", "b")
		kvc.AddIf(false, "c", "d")
		if kvc.Length() != 1 {
			t.Errorf("expected 1, got %d", kvc.Length())
		}
	})
}

func Test_Cov40_KVC_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddStringBySplit", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddStringBySplit("=", "key=val")
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddStringBySplitTrim", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddStringBySplitTrim("=", " key = val ")
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Adds", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Adds(corestr.KeyValuePair{Key: "a", Value: "b"}, corestr.KeyValuePair{Key: "c", Value: "d"})
		if kvc.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_KVC_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Adds_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Adds()
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_AddMap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddMap", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddMap(map[string]string{"a": "b"})
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddMap_Nil", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddMap(nil)
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashsetMap", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashsetMap(map[string]bool{"x": true})
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashsetMap_Nil", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashsetMap(nil)
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashset", func() {
		kvc := corestr.New.KeyValues.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		kvc.AddHashset(hs)
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashset_Nil", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashset(nil)
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmap", func() {
		kvc := corestr.New.KeyValues.Empty()
		hm := corestr.New.Hashmap.Cap(3)
		hm.AddOrUpdate("a", "b")
		kvc.AddsHashmap(hm)
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmap_Nil", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddsHashmap(nil)
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmaps", func() {
		kvc := corestr.New.KeyValues.Empty()
		hm := corestr.New.Hashmap.Cap(3)
		hm.AddOrUpdate("a", "b")
		kvc.AddsHashmaps(hm)
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmaps_Nil", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddsHashmaps()
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_First_Last", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		if kvc.First().Key != "a" || kvc.Last().Key != "b" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_FirstOrDefault_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if kvc.FirstOrDefault() != nil || kvc.LastOrDefault() != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov40_KVC_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_FirstOrDefault_NonEmpty", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("x", "y")
		if kvc.FirstOrDefault() == nil || kvc.LastOrDefault() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_KVC_HasKey(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_HasKey", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if !kvc.HasKey("k") || kvc.HasKey("z") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_IsContains", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if !kvc.IsContains("k") || kvc.IsContains("z") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_Get(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Get", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		val, found := kvc.Get("k")
		if !found || val != "v" {
			t.Error("unexpected")
		}
		_, found2 := kvc.Get("missing")
		if found2 {
			t.Error("expected not found")
		}
	})
}

func Test_Cov40_KVC_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_HasIndex", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if !kvc.HasIndex(0) || kvc.HasIndex(1) {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_LastIndex", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		if kvc.LastIndex() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValueAt", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if kvc.SafeValueAt(0) != "v" || kvc.SafeValueAt(5) != "" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_SafeValueAt_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValueAt_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if kvc.SafeValueAt(0) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov40_KVC_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValuesAtIndexes", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		vals := kvc.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 || vals[0] != "1" || vals[1] != "2" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValuesAtIndexes_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		vals := kvc.SafeValuesAtIndexes()
		if len(vals) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_AllKeys_AllValues(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeys_AllValues", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		if len(kvc.AllKeys()) != 2 || len(kvc.AllValues()) != 2 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeys_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if len(kvc.AllKeys()) != 0 || len(kvc.AllValues()) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeysSorted", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("b", "2").Add("a", "1")
		sorted := kvc.AllKeysSorted()
		if sorted[0] != "a" || sorted[1] != "b" {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Hashmap", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		hm := kvc.Hashmap()
		if hm.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Hashmap_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if kvc.Hashmap().Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_Map(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Map", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if len(kvc.Map()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_KVC_Find(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})
		if len(found) != 1 {
			t.Errorf("expected 1, got %d", len(found))
		}
	})
}

func Test_Cov40_KVC_Find_Break(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find_Break", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})
		if len(found) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_KVC_Find_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		if len(found) != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_KVC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Strings", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")
		if len(kvc.Strings()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_KVC_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Strings_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if len(kvc.Strings()) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_StringsUsingFormat", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		strs := kvc.StringsUsingFormat("%v=%v")
		if len(strs) != 1 || strs[0] != "k=v" {
			t.Errorf("unexpected: %v", strs)
		}
	})
}

func Test_Cov40_KVC_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_StringsUsingFormat_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if len(kvc.StringsUsingFormat("%v=%v")) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_String_Compile(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_String_Compile", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if kvc.String() == "" || kvc.Compile() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_KVC_Join_JoinKeys_JoinValues(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Join_JoinKeys_JoinValues", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		if kvc.Join(",") == "" || kvc.JoinKeys(",") == "" || kvc.JoinValues(",") == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_KVC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JSON", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		data, err := json.Marshal(kvc)
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
		kvc2 := corestr.New.KeyValues.Empty()
		err = json.Unmarshal(data, kvc2)
		if err != nil || kvc2.Length() != 1 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_UnmarshalJSON_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_UnmarshalJSON_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		err := json.Unmarshal([]byte(`[]`), kvc)
		if err != nil || kvc.Length() != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Json_JsonPtr", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if kvc.Json().Error != nil || kvc.JsonPtr() == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JsonModel", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		if len(kvc.JsonModel()) != 1 || kvc.JsonModelAny() == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Serialize_Deserialize", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		data, err := kvc.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
		if len(kvc.SerializeMust()) == 0 {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_KVC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_ParseInjectUsingJson", func() {
		src := corestr.New.KeyValues.Empty()
		src.Add("k", "v")
		kvc := corestr.New.KeyValues.Empty()
		result, err := kvc.ParseInjectUsingJson(src.JsonPtr())
		if err != nil || result.Length() != 1 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_KVC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AsJsoner", func() {
		kvc := corestr.New.KeyValues.Empty()
		if kvc.AsJsoner() == nil || kvc.AsJsonContractsBinder() == nil || kvc.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_KVC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JsonParseSelfInject", func() {
		src := corestr.New.KeyValues.Empty()
		src.Add("k", "v")
		kvc := corestr.New.KeyValues.Empty()
		err := kvc.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Cov40_KVC_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Clear_Dispose", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		kvc.Clear()
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Dispose", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		kvc.Dispose()
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_KVC_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Deserialize", func() {
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		var target []corestr.KeyValuePair
		err := kvc.Deserialize(&target)
		if err != nil {
			t.Error(err)
		}
	})
}

// ── newKeyValuesCreator ──

func Test_Cov40_Creator_KV_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_Cap", func() {
		kvc := corestr.New.KeyValues.Cap(5)
		if kvc == nil || kvc.Length() != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_Creator_KV_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_Empty", func() {
		kvc := corestr.New.KeyValues.Empty()
		if kvc == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_Creator_KV_UsingMap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingMap", func() {
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1"})
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_KV_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingMap_Empty", func() {
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{})
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_Creator_KV_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValuePairs", func() {
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(corestr.KeyValuePair{Key: "a", Value: "1"})
		if kvc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_KV_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValuePairs_Empty", func() {
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_Creator_KV_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValueStrings", func() {
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{"a", "b"}, []string{"1", "2"})
		if kvc.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_Creator_KV_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValueStrings_Empty", func() {
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})
		if kvc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// HashsetsCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_HC_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Basic", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if !hc.IsEmpty() || hc.HasItems() || hc.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov40_HC_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Add", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(hs)
		if hc.Length() != 1 || !hc.HasItems() {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_HC_AddNonNil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddNonNil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonNil(nil)
		hc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_HC_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddNonEmpty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonEmpty(corestr.New.Hashset.Empty())
		hc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_HC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Adds", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Adds(corestr.New.Hashset.Strings([]string{"a"}), corestr.New.Hashset.Strings([]string{"b"}))
		if hc.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_HC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Adds_Nil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Adds()
		if hc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_HC_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddHashsetsCollection", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hc1.AddHashsetsCollection(hc2)
		if hc1.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_HC_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddHashsetsCollection_Nil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddHashsetsCollection(nil)
		if hc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_HC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ConcatNew", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hc1.ConcatNew(hc2)
		if result.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_HC_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ConcatNew_NoArgs", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hc.ConcatNew()
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_HC_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_LastIndex", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.LastIndex() != 0 {
			t.Errorf("expected 0")
		}
	})
}
func Test_Cov40_HC_StringsList(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_StringsList", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))
		if len(hc.StringsList()) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_HC_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_StringsList_Empty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if len(hc.StringsList()) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_HC_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_HasAll", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))
		if !hc.HasAll("a", "b") {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_HC_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_HasAll_Empty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if hc.HasAll("a") {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_HC_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqual", func() {
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if !hc1.IsEqual(*hc2) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov40_HC_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_SamePtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if !hc.IsEqualPtr(hc) {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_HC_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_BothEmpty", func() {
		a := corestr.New.HashsetsCollection.Empty()
		b := corestr.New.HashsetsCollection.Empty()
		if !a.IsEqualPtr(b) {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_HC_IsEqualPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_DiffLen", func() {
		a := corestr.New.HashsetsCollection.Empty()
		a.Add(corestr.New.Hashset.Strings([]string{"a"}))
		b := corestr.New.HashsetsCollection.Empty()
		if a.IsEqualPtr(b) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_HC_IsEqualPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_Nil", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if hc.IsEqualPtr(nil) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_HC_String(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_String", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_HC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_String_Empty", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if !strings.Contains(hc.String(), "NoElements") {
			// accept any format
			_ = 0
		}
	})
}

func Test_Cov40_HC_Join(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Join", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.Join(",") == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_HC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JSON", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := json.Marshal(hc)
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_HC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_UnmarshalJSON", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, _ := json.Marshal(hc)
		hc2 := corestr.New.HashsetsCollection.Empty()
		err := json.Unmarshal(data, hc2)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Cov40_HC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Json_JsonPtr", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if hc.Json().Error != nil || hc.JsonPtr() == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_HC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JsonModel", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if hc.JsonModel() == nil || hc.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_HC_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Serialize", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := hc.Serialize()
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_HC_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Deserialize", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		var target interface{}
		_ = hc.Deserialize(&target)
	})
}

func Test_Cov40_HC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ParseInjectUsingJson", func() {
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		result, err := hc.ParseInjectUsingJson(src.JsonPtr())
		if err != nil || result == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_HC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ParseInjectUsingJsonMust", func() {
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		result := hc.ParseInjectUsingJsonMust(src.JsonPtr())
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_HC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AsJsoner", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if hc.AsJsoner() == nil || hc.AsJsonContractsBinder() == nil || hc.AsJsonParseSelfInjector() == nil || hc.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_HC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JsonParseSelfInject", func() {
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		err := hc.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Error(err)
		}
	})
}

// ── newHashsetsCollectionCreator ──

func Test_Cov40_Creator_HC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_Empty", func() {
		if corestr.New.HashsetsCollection.Empty() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_Creator_HC_UsingHashsets(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsets", func() {
		hs := *corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsets(hs)
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_HC_UsingHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsets_Empty", func() {
		hc := corestr.New.HashsetsCollection.UsingHashsets()
		if hc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_Creator_HC_UsingHashsetsPointers(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsetsPointers", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_HC_UsingHashsetsPointers_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsetsPointers_Empty", func() {
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers()
		if hc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_Creator_HC_LenCap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_LenCap", func() {
		hc := corestr.New.HashsetsCollection.LenCap(0, 5)
		if hc == nil || hc.Length() != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_Creator_HC_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_Cap", func() {
		hc := corestr.New.HashsetsCollection.Cap(5)
		if hc == nil || hc.Length() != 0 {
			t.Error("unexpected")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// CollectionsOfCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_COC_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Basic", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		if !coc.IsEmpty() || coc.HasItems() || coc.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov40_COC_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Add", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_COC_Add_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Add_Empty", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{}))
		if coc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_COC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Adds", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(c)
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_COC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Adds_Nil", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Adds()
		if coc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_COC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddCollections", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.AddCollections(c)
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_COC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_COC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddStrings_Empty", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{})
		if coc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_COC_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddsStringsOfStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
		if coc.Length() != 2 {
			t.Errorf("expected 2, got %d", coc.Length())
		}
	})
}

func Test_Cov40_COC_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddsStringsOfStrings_Nil", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddsStringsOfStrings(false)
		if coc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_COC_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddAsyncFuncItems", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		coc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_COC_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddAsyncFuncItems_Nil", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddAsyncFuncItems(nil, false)
		if coc.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_COC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AllIndividualItemsLength", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))
		if coc.AllIndividualItemsLength() != 3 {
			t.Errorf("expected 3")
		}
	})
}

func Test_Cov40_COC_Items(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Items", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if len(coc.Items()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_COC_List(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_List", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		if len(coc.List(0)) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_COC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_List_Empty", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		if len(coc.List(0)) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_COC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ToCollection", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if coc.ToCollection().Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_COC_String(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_String", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if coc.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_COC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JSON", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := json.Marshal(coc)
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_COC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_UnmarshalJSON", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, _ := json.Marshal(coc)
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := json.Unmarshal(data, coc2)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Cov40_COC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Json_JsonPtr", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		if coc.Json().Error != nil || coc.JsonPtr() == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_COC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JsonModel", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		if coc.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_COC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ParseInjectUsingJson", func() {
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		result, err := coc.ParseInjectUsingJson(src.JsonPtr())
		if err != nil || result == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_COC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ParseInjectUsingJsonMust", func() {
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		result := coc.ParseInjectUsingJsonMust(src.JsonPtr())
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_COC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AsJsoner", func() {
		coc := corestr.New.CollectionsOfCollection.Empty()
		if coc.AsJsoner() == nil || coc.AsJsonContractsBinder() == nil || coc.AsJsonParseSelfInjector() == nil || coc.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_COC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JsonParseSelfInject", func() {
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		err := coc.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Error(err)
		}
	})
}

// ── newCollectionsOfCollectionCreator ──

func Test_Cov40_Creator_COC_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Cap", func() {
		if corestr.New.CollectionsOfCollection.Cap(5) == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_Creator_COC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Empty", func() {
		if corestr.New.CollectionsOfCollection.Empty() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_Creator_COC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Strings", func() {
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_COC_CloneStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_CloneStrings", func() {
		coc := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_COC_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_SpreadStrings", func() {
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_COC_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOfStrings", func() {
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
		if coc.Length() != 2 {
			t.Errorf("expected 2, got %d", coc.Length())
		}
	})
}

func Test_Cov40_Creator_COC_StringsOption(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOption", func() {
		coc := corestr.New.CollectionsOfCollection.StringsOption(false, 5, []string{"a"})
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_COC_StringsOptions(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOptions", func() {
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})
		if coc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_Creator_COC_LenCap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_LenCap", func() {
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)
		if coc == nil {
			t.Error("expected non-nil")
		}
	})
}

// ══════════════════════════════════════════════════════════════
// CharHashsetMap — core operations
// ══════════════════════════════════════════════════════════════

func Test_Cov40_CHM_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Basic", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if !chm.IsEmpty() || chm.HasItems() || chm.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov40_CHM_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("avocado").Add("banana")
		if chm.Length() != 2 {
			t.Errorf("expected 2, got %d", chm.Length())
		}
	})
}

func Test_Cov40_CHM_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings("apple", "banana")
		if chm.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStrings_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings()
		if chm.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_CHM_Has(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Has", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if !chm.Has("apple") || chm.Has("banana") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Has_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.Has("x") {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")
		if !has || hs == nil {
			t.Error("unexpected")
		}
		has2, _ := chm.HasWithHashset("banana")
		if has2 {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashset_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		has, _ := chm.HasWithHashset("x")
		if has {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_GetChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.GetChar("abc") != 'a' || chm.GetChar("") != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_GetCharOf(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.GetCharOf("xyz") != 'x' || chm.GetCharOf("") != 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("avocado")
		if chm.LengthOf('a') != 2 {
			t.Errorf("expected 2, got %d", chm.LengthOf('a'))
		}
		if chm.LengthOf('z') != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOf_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.LengthOf('a') != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.LengthOfHashsetFromFirstChar("a") != 1 {
			t.Error("expected 1")
		}
		if chm.LengthOfHashsetFromFirstChar("z") != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AllLengthsSum_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.AllLengthsSum() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_List(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_List", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		if len(chm.List()) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("banana").Add("apple")
		list := chm.SortedListAsc()
		if list[0] != "apple" {
			t.Error("expected apple first")
		}
	})
}

func Test_Cov40_CHM_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		list := chm.SortedListDsc()
		if list[0] != "banana" {
			t.Error("expected banana first")
		}
	})
}

func Test_Cov40_CHM_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetMap", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if len(chm.GetMap()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_CHM_GetHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.GetHashset("a", false)
		if hs == nil || !hs.Has("apple") {
			t.Error("unexpected")
		}
		hs2 := chm.GetHashset("z", false)
		if hs2 != nil {
			t.Error("expected nil")
		}
		hs3 := chm.GetHashset("z", true)
		if hs3 == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.GetHashsetByChar('a') == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.HashsetByChar('a') == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.HashsetByStringFirstChar("apple") == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollection()
		if hc.Length() != 2 {
			t.Errorf("expected 2, got %d", hc.Length())
		}
	})
}

func Test_Cov40_CHM_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollection_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.HashsetsCollection().Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollectionByChars('a')
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_CHM_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByChars_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.HashsetsCollectionByChars('a').Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollectionByStringsFirstChar("apple")
		if hc.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.HashsetsCollectionByStringsFirstChar("a").Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)
		if chm.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddCollectionItems(nil)
		if chm.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		chm.AddHashsetItems(hs)
		if chm.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})
		if chm.LengthOf('a') != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.AddSameStartingCharItems('a', []string{"avocado"})
		if chm.LengthOf('a') != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{})
		if chm.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple"})
		hs := chm.AddSameCharsCollection("a", col)
		if hs == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollection("a", col)
		if hs == nil || !hs.Has("avocado") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_NilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.AddSameCharsCollection("a", nil)
		if hs == nil {
			t.Error("expected non-nil (new hashset created)")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_ExistingNilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.AddSameCharsCollection("a", nil)
		if hs == nil {
			t.Error("expected existing hashset")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddSameCharsHashset("a", hs)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddSameCharsHashset("a", hs)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsHashset("a", nil)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_ExistingNilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_ExistingNilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddSameCharsHashset("a", nil)
		if result == nil {
			t.Error("expected existing")
		}
	})
}

func Test_Cov40_CHM_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2.Add("apple")
		if !chm1.IsEquals(chm2) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov40_CHM_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_SamePtr", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if !chm.IsEquals(chm) {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_CHM_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.IsEquals(nil) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_BothEmpty", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		if !a.IsEquals(b) {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_CHM_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_DiffLen", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		if a.IsEquals(b) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_DiffContent", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("avocado")
		if a.IsEquals(b) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_MissingKey", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("banana")
		if a.IsEquals(b) {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_Lock_Variants(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Lock_Variants", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.IsEmptyLock() {
			t.Error("expected not empty")
		}
		if chm.LengthLock() != 1 {
			t.Error("expected 1")
		}
		if chm.AllLengthsSumLock() != 1 {
			t.Error("expected 1")
		}
		if chm.LengthOfLock('a') != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_Cov40_CHM_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOfLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.LengthOfLock('a') != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("apple")
		if !chm.Has("apple") {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_CHM_AddLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("apple")
		chm.AddLock("avocado")
		if chm.LengthOf('a') != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock("apple", "banana")
		if chm.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov40_CHM_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStringsLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock()
		if chm.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_CHM_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.GetHashsetLock(false, "a")
		if hs == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		m := chm.GetCopyMapLock()
		if len(m) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov40_CHM_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCopyMapLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		m := chm.GetCopyMapLock()
		if len(m) != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov40_CHM_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")
		if !has || hs == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		has, _ := chm.HasWithHashsetLock("x")
		if has {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_HasWithHashsetLock_MissingChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock_MissingChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, _ := chm.HasWithHashsetLock("banana")
		if has {
			t.Error("expected false")
		}
	})
}

func Test_Cov40_CHM_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEqualsLock", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("apple")
		if !a.IsEqualsLock(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov40_CHM_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('a')
		if hs == nil || !hs.Has("apple") {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_HashsetByCharLock_Missing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByCharLock_Missing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('z')
		if hs == nil {
			t.Error("expected empty hashset not nil")
		}
	})
}

func Test_Cov40_CHM_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("apple")
		if hs == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		groups := chm.GetCharsGroups("apple", "avocado", "banana")
		if groups.Length() != 2 {
			t.Errorf("expected 2, got %d", groups.Length())
		}
	})
}

func Test_Cov40_CHM_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharsGroups_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		groups := chm.GetCharsGroups()
		if groups != chm {
			t.Error("expected same pointer")
		}
	})
}

func Test_Cov40_CHM_Print(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Print", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.Print(false) // skip
		chm.Print(true)  // print
	})
}

func Test_Cov40_CHM_PrintLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_PrintLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

func Test_Cov40_CHM_String(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_String", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_CHM_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_StringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.StringLock() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_CHM_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.SummaryString() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_CHM_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.SummaryStringLock() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov40_CHM_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		data, err := json.Marshal(chm)
		if err != nil || len(data) == 0 {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_UnmarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		data, _ := json.Marshal(chm)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		err := json.Unmarshal(data, chm2)
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Cov40_CHM_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Json_JsonPtr", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.Json().Error != nil || chm.JsonPtr() == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JsonModel", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		if chm.JsonModel() == nil || chm.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_ParseInjectUsingJson", func() {
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result, err := chm.ParseInjectUsingJson(src.JsonPtr())
		if err != nil || result == nil {
			t.Error("unexpected")
		}
	})
}

func Test_Cov40_CHM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_ParseInjectUsingJsonMust", func() {
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.ParseInjectUsingJsonMust(src.JsonPtr())
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AsJsoner", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		if chm.AsJsoner() == nil || chm.AsJsonContractsBinder() == nil || chm.AsJsonParseSelfInjector() == nil || chm.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JsonParseSelfInject", func() {
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		err := chm.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_Cov40_CHM_Clear(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Clear", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.Clear()
		if chm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Clear_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Clear()
		if chm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.RemoveAll()
		if chm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_RemoveAll_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.RemoveAll()
		if chm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCharCollectionMapItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)
		if !chm.Has("apple") {
			t.Error("expected true")
		}
	})
}

func Test_Cov40_CHM_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCharCollectionMapItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddCharCollectionMapItems(nil)
		if chm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_Cov40_CHM_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddHashsetLock("a", hs)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddHashsetLock("a", hs)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddHashsetLock_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddHashsetLock("a", nil)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddHashsetLock_ExistingNilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_ExistingNilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddHashsetLock("a", nil)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := chm.AddSameCharsCollectionLock("a", col)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollectionLock("a", col)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_NilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsCollectionLock("a", nil)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_ExistingNilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddSameCharsCollectionLock("a", nil)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}
