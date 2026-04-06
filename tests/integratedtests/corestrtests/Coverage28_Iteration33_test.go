package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══════════════════════════════════════════════════════════════════════
// KeyAnyValuePair — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C28_01_KeyAnyValuePair_KeyName(t *testing.T) {
	safeTest(t, "Test_C28_01_KeyAnyValuePair_KeyName", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if kv.KeyName() != "k" {
			t.Error("expected k")
		}
	})
}

func Test_C28_02_KeyAnyValuePair_VariableName(t *testing.T) {
	safeTest(t, "Test_C28_02_KeyAnyValuePair_VariableName", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if kv.VariableName() != "k" {
			t.Error("expected k")
		}
	})
}

func Test_C28_03_KeyAnyValuePair_ValueAny(t *testing.T) {
	safeTest(t, "Test_C28_03_KeyAnyValuePair_ValueAny", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if kv.ValueAny() != 42 {
			t.Error("expected 42")
		}
	})
}

func Test_C28_04_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C28_04_KeyAnyValuePair_IsVariableNameEqual", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		if !kv.IsVariableNameEqual("k") {
			t.Error("expected true")
		}
		if kv.IsVariableNameEqual("x") {
			t.Error("expected false")
		}
	})
}

func Test_C28_05_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	safeTest(t, "Test_C28_05_KeyAnyValuePair_IsValueNull_Nil", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		if !kv.IsValueNull() {
			t.Error("expected true")
		}
	})
}

func Test_C28_06_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C28_06_KeyAnyValuePair_IsValueNull_NilReceiver", func() {
		var kv *corestr.KeyAnyValuePair
		if !kv.IsValueNull() {
			t.Error("expected true")
		}
	})
}

func Test_C28_07_KeyAnyValuePair_IsValueNull_NonNil(t *testing.T) {
	safeTest(t, "Test_C28_07_KeyAnyValuePair_IsValueNull_NonNil", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if kv.IsValueNull() {
			t.Error("expected false")
		}
	})
}

func Test_C28_08_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C28_08_KeyAnyValuePair_HasNonNull", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		if !kv.HasNonNull() {
			t.Error("expected true")
		}
	})
}

func Test_C28_09_KeyAnyValuePair_HasNonNull_Nil(t *testing.T) {
	safeTest(t, "Test_C28_09_KeyAnyValuePair_HasNonNull_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		if kv.HasNonNull() {
			t.Error("expected false")
		}
	})
}

func Test_C28_10_KeyAnyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_C28_10_KeyAnyValuePair_HasValue", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if !kv.HasValue() {
			t.Error("expected true")
		}
	})
}

func Test_C28_11_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C28_11_KeyAnyValuePair_IsValueEmptyString", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		if !kv.IsValueEmptyString() {
			t.Error("expected true")
		}
	})
}

func Test_C28_12_KeyAnyValuePair_IsValueEmptyString_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C28_12_KeyAnyValuePair_IsValueEmptyString_NilReceiver", func() {
		var kv *corestr.KeyAnyValuePair
		if !kv.IsValueEmptyString() {
			t.Error("expected true")
		}
	})
}

func Test_C28_13_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C28_13_KeyAnyValuePair_IsValueWhitespace", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		if !kv.IsValueWhitespace() {
			t.Error("expected true")
		}
	})
}

func Test_C28_14_KeyAnyValuePair_IsValueWhitespace_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C28_14_KeyAnyValuePair_IsValueWhitespace_NilReceiver", func() {
		var kv *corestr.KeyAnyValuePair
		if !kv.IsValueWhitespace() {
			t.Error("expected true")
		}
	})
}

func Test_C28_15_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C28_15_KeyAnyValuePair_ValueString", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kv.ValueString()
		if s == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_16_KeyAnyValuePair_ValueString_Cached(t *testing.T) {
	safeTest(t, "Test_C28_16_KeyAnyValuePair_ValueString_Cached", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s1 := kv.ValueString()
		s2 := kv.ValueString()
		if s1 != s2 {
			t.Error("expected same")
		}
	})
}

func Test_C28_17_KeyAnyValuePair_ValueString_NilValue(t *testing.T) {
	safeTest(t, "Test_C28_17_KeyAnyValuePair_ValueString_NilValue", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k"}
		s := kv.ValueString()
		_ = s // should not panic
	})
}

func Test_C28_18_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C28_18_KeyAnyValuePair_String", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_19_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C28_19_KeyAnyValuePair_Compile", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.Compile() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_20_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C28_20_KeyAnyValuePair_SerializeMust", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()
		if len(b) == 0 {
			t.Error("expected bytes")
		}
	})
}

func Test_C28_21_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_C28_21_KeyAnyValuePair_Serialize", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()
		if err != nil || len(b) == 0 {
			t.Error("expected bytes")
		}
	})
}

func Test_C28_22_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_C28_22_KeyAnyValuePair_Json", func() {
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kv.Json()
		_ = j
	})
}

func Test_C28_23_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C28_23_KeyAnyValuePair_JsonPtr", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.JsonPtr() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_24_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C28_24_KeyAnyValuePair_ParseInjectUsingJson", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		_, err := kv2.ParseInjectUsingJson(jp)
		if err != nil {
			t.Errorf("unexpected: %v", err)
		}
	})
}

func Test_C28_25_KeyAnyValuePair_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C28_25_KeyAnyValuePair_ParseInjectUsingJsonMust", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		result := kv2.ParseInjectUsingJsonMust(jp)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_26_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C28_26_KeyAnyValuePair_AsJsonContractsBinder", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_27_KeyAnyValuePair_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C28_27_KeyAnyValuePair_AsJsoner", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_28_KeyAnyValuePair_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C28_28_KeyAnyValuePair_JsonParseSelfInject", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jp := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		err := kv2.JsonParseSelfInject(jp)
		if err != nil {
			t.Errorf("unexpected: %v", err)
		}
	})
}

func Test_C28_29_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C28_29_KeyAnyValuePair_AsJsonParseSelfInjector", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kv.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_30_KeyAnyValuePair_Clear(t *testing.T) {
	safeTest(t, "Test_C28_30_KeyAnyValuePair_Clear", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		if kv.Key != "" || kv.Value != nil {
			t.Error("expected cleared")
		}
	})
}

func Test_C28_31_KeyAnyValuePair_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C28_31_KeyAnyValuePair_Clear_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Clear() // no panic
	})
}

func Test_C28_32_KeyAnyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_C28_32_KeyAnyValuePair_Dispose", func() {
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Dispose()
	})
}

func Test_C28_33_KeyAnyValuePair_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C28_33_KeyAnyValuePair_Dispose_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Dispose()
	})
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValueCollection — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C28_34_KeyValueCollection_Add(t *testing.T) {
	safeTest(t, "Test_C28_34_KeyValueCollection_Add", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_35_KeyValueCollection_AddIf_True(t *testing.T) {
	safeTest(t, "Test_C28_35_KeyValueCollection_AddIf_True", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "k", "v")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_36_KeyValueCollection_AddIf_False(t *testing.T) {
	safeTest(t, "Test_C28_36_KeyValueCollection_AddIf_False", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(false, "k", "v")
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_37_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C28_37_KeyValueCollection_Adds", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)
		if kvc.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_38_KeyValueCollection_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C28_38_KeyValueCollection_Adds_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_39_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C28_39_KeyValueCollection_Count", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		if kvc.Count() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_40_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C28_40_KeyValueCollection_HasAnyItem", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		if !kvc.HasAnyItem() {
			t.Error("expected true")
		}
	})
}

func Test_C28_41_KeyValueCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C28_41_KeyValueCollection_LastIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.LastIndex() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_42_KeyValueCollection_HasIndex(t *testing.T) {
	safeTest(t, "Test_C28_42_KeyValueCollection_HasIndex", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if !kvc.HasIndex(0) {
			t.Error("expected true")
		}
		if kvc.HasIndex(5) {
			t.Error("expected false")
		}
	})
}

func Test_C28_43_KeyValueCollection_First(t *testing.T) {
	safeTest(t, "Test_C28_43_KeyValueCollection_First", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if kvc.First().Key != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C28_44_KeyValueCollection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C28_44_KeyValueCollection_FirstOrDefault", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.FirstOrDefault() != nil {
			t.Error("expected nil")
		}
		kvc.Add("a", "1")
		if kvc.FirstOrDefault().Key != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C28_45_KeyValueCollection_Last(t *testing.T) {
	safeTest(t, "Test_C28_45_KeyValueCollection_Last", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.Last().Key != "b" {
			t.Error("expected b")
		}
	})
}

func Test_C28_46_KeyValueCollection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C28_46_KeyValueCollection_LastOrDefault", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.LastOrDefault() != nil {
			t.Error("expected nil")
		}
		kvc.Add("a", "1")
		if kvc.LastOrDefault().Key != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C28_47_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C28_47_KeyValueCollection_Find", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "b", false
		})
		if len(found) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_48_KeyValueCollection_Find_Empty(t *testing.T) {
	safeTest(t, "Test_C28_48_KeyValueCollection_Find_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})
		if len(found) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_49_KeyValueCollection_Find_Break(t *testing.T) {
	safeTest(t, "Test_C28_49_KeyValueCollection_Find_Break", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		kvc.Add("c", "3")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, i == 0
		})
		if len(found) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_50_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C28_50_KeyValueCollection_SafeValueAt", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if kvc.SafeValueAt(0) != "1" {
			t.Error("expected 1")
		}
		if kvc.SafeValueAt(99) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C28_51_KeyValueCollection_SafeValueAt_Empty(t *testing.T) {
	safeTest(t, "Test_C28_51_KeyValueCollection_SafeValueAt_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.SafeValueAt(0) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C28_52_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C28_52_KeyValueCollection_SafeValuesAtIndexes", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		vals := kvc.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_53_KeyValueCollection_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C28_53_KeyValueCollection_SafeValuesAtIndexes_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		vals := kvc.SafeValuesAtIndexes()
		if len(vals) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_54_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_C28_54_KeyValueCollection_Strings", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		s := kvc.Strings()
		if len(s) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_55_KeyValueCollection_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C28_55_KeyValueCollection_Strings_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		s := kvc.Strings()
		if len(s) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_56_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_C28_56_KeyValueCollection_StringsUsingFormat", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		s := kvc.StringsUsingFormat("%s=%s")
		if len(s) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_57_KeyValueCollection_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_C28_57_KeyValueCollection_StringsUsingFormat_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		s := kvc.StringsUsingFormat("%s=%s")
		if len(s) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_58_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_C28_58_KeyValueCollection_String", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if kvc.String() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_59_KeyValueCollection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C28_59_KeyValueCollection_Length_Nil", func() {
		var kvc *corestr.KeyValueCollection
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_60_KeyValueCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C28_60_KeyValueCollection_IsEmpty", func() {
		kvc := &corestr.KeyValueCollection{}
		if !kvc.IsEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_C28_61_KeyValueCollection_Compile(t *testing.T) {
	safeTest(t, "Test_C28_61_KeyValueCollection_Compile", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("k", "v")
		if kvc.Compile() == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_62_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C28_62_KeyValueCollection_AddStringBySplit", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_63_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C28_63_KeyValueCollection_AddStringBySplitTrim", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_64_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C28_64_KeyValueCollection_AddMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1", "b": "2"})
		if kvc.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_65_KeyValueCollection_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_C28_65_KeyValueCollection_AddMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_66_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C28_66_KeyValueCollection_AddHashsetMap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_67_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C28_67_KeyValueCollection_AddHashsetMap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_68_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_C28_68_KeyValueCollection_AddHashset", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(hs)
		if kvc.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_69_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C28_69_KeyValueCollection_AddHashset_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_70_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_C28_70_KeyValueCollection_AddsHashmap", func() {
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("a", "1")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(hm)
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_71_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_C28_71_KeyValueCollection_AddsHashmap_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_72_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_C28_72_KeyValueCollection_Hashmap", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()
		if hm.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_73_KeyValueCollection_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_C28_73_KeyValueCollection_Hashmap_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		hm := kvc.Hashmap()
		if hm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_74_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_C28_74_KeyValueCollection_IsContains", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if !kvc.IsContains("a") {
			t.Error("expected true")
		}
		if kvc.IsContains("z") {
			t.Error("expected false")
		}
	})
}

func Test_C28_75_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_C28_75_KeyValueCollection_Get", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		v, ok := kvc.Get("a")
		if !ok || v != "1" {
			t.Error("expected 1")
		}
		_, ok2 := kvc.Get("z")
		if ok2 {
			t.Error("expected false")
		}
	})
}

func Test_C28_76_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_C28_76_KeyValueCollection_Map", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		m := kvc.Map()
		if len(m) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_77_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_C28_77_KeyValueCollection_HasKey", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if !kvc.HasKey("a") {
			t.Error("expected true")
		}
		if kvc.HasKey("z") {
			t.Error("expected false")
		}
	})
}

func Test_C28_78_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C28_78_KeyValueCollection_AllKeysSorted", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2")
		kvc.Add("a", "1")
		keys := kvc.AllKeysSorted()
		if keys[0] != "a" || keys[1] != "b" {
			t.Error("expected sorted")
		}
	})
}

func Test_C28_79_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_C28_79_KeyValueCollection_AllKeys", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		keys := kvc.AllKeys()
		if len(keys) != 1 || keys[0] != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C28_80_KeyValueCollection_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_C28_80_KeyValueCollection_AllKeys_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		keys := kvc.AllKeys()
		if len(keys) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_81_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_C28_81_KeyValueCollection_AllValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		vals := kvc.AllValues()
		if len(vals) != 1 || vals[0] != "1" {
			t.Error("expected 1")
		}
	})
}

func Test_C28_82_KeyValueCollection_AllValues_Empty(t *testing.T) {
	safeTest(t, "Test_C28_82_KeyValueCollection_AllValues_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		vals := kvc.AllValues()
		if len(vals) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_83_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_C28_83_KeyValueCollection_Join", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		j := kvc.Join(",")
		if j == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_84_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C28_84_KeyValueCollection_JoinKeys", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.JoinKeys(",") == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_85_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_C28_85_KeyValueCollection_JoinValues", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Add("b", "2")
		if kvc.JoinValues(",") == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C28_86_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_C28_86_KeyValueCollection_AddsHashmaps", func() {
		hm1 := corestr.New.Hashmap.Cap(2)
		hm1.AddOrUpdate("a", "1")
		hm2 := corestr.New.Hashmap.Cap(2)
		hm2.AddOrUpdate("b", "2")
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps(hm1, hm2)
		if kvc.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_87_KeyValueCollection_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_C28_87_KeyValueCollection_AddsHashmaps_Nil", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps()
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_88_KeyValueCollection_JsonModel(t *testing.T) {
	safeTest(t, "Test_C28_88_KeyValueCollection_JsonModel", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if len(kvc.JsonModel()) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_89_KeyValueCollection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C28_89_KeyValueCollection_JsonModelAny", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if kvc.JsonModelAny() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_90_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C28_90_KeyValueCollection_Serialize", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.Serialize()
		if err != nil || len(b) == 0 {
			t.Error("expected bytes")
		}
	})
}

func Test_C28_91_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C28_91_KeyValueCollection_SerializeMust", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b := kvc.SerializeMust()
		if len(b) == 0 {
			t.Error("expected bytes")
		}
	})
}

func Test_C28_92_KeyValueCollection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C28_92_KeyValueCollection_MarshalJSON", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.MarshalJSON()
		if err != nil || len(b) == 0 {
			t.Error("expected bytes")
		}
	})
}

func Test_C28_93_KeyValueCollection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C28_93_KeyValueCollection_UnmarshalJSON", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, _ := kvc.MarshalJSON()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.UnmarshalJSON(b)
		if err != nil || kvc2.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C28_94_KeyValueCollection_UnmarshalJSON_Empty(t *testing.T) {
	safeTest(t, "Test_C28_94_KeyValueCollection_UnmarshalJSON_Empty", func() {
		kvc := &corestr.KeyValueCollection{}
		err := kvc.UnmarshalJSON([]byte("[]"))
		if err != nil {
			t.Error("expected nil err")
		}
	})
}

func Test_C28_95_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_C28_95_KeyValueCollection_Json", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		j := kvc.Json()
		_ = j
	})
}

func Test_C28_96_KeyValueCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C28_96_KeyValueCollection_JsonPtr", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		if kvc.JsonPtr() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_97_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C28_97_KeyValueCollection_ParseInjectUsingJson", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jp := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		_, err := kvc2.ParseInjectUsingJson(jp)
		if err != nil {
			t.Errorf("unexpected: %v", err)
		}
	})
}

func Test_C28_98_KeyValueCollection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C28_98_KeyValueCollection_AsJsonContractsBinder", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_99_KeyValueCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C28_99_KeyValueCollection_AsJsoner", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_100_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C28_100_KeyValueCollection_JsonParseSelfInject", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jp := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jp)
		if err != nil {
			t.Errorf("unexpected: %v", err)
		}
	})
}

func Test_C28_101_KeyValueCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C28_101_KeyValueCollection_AsJsonParseSelfInjector", func() {
		kvc := &corestr.KeyValueCollection{}
		if kvc.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C28_102_KeyValueCollection_Clear(t *testing.T) {
	safeTest(t, "Test_C28_102_KeyValueCollection_Clear", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()
		if kvc.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_103_KeyValueCollection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C28_103_KeyValueCollection_Clear_Nil", func() {
		var kvc *corestr.KeyValueCollection
		kvc.Clear()
	})
}

func Test_C28_104_KeyValueCollection_Dispose(t *testing.T) {
	safeTest(t, "Test_C28_104_KeyValueCollection_Dispose", func() {
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Dispose()
	})
}

func Test_C28_105_KeyValueCollection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C28_105_KeyValueCollection_Dispose_Nil", func() {
		var kvc *corestr.KeyValueCollection
		kvc.Dispose()
	})
}

func Test_C28_106_KeyValueCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C28_106_KeyValueCollection_Deserialize", func() {
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		var target corestr.KeyValueCollection
		err := kvc.Deserialize(&target)
		if err != nil {
			t.Errorf("unexpected: %v", err)
		}
	})
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleStringOnce — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C28_107_SimpleStringOnce_Value(t *testing.T) {
	safeTest(t, "Test_C28_107_SimpleStringOnce_Value", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if s.Value() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_108_SimpleStringOnce_IsInitialized(t *testing.T) {
	safeTest(t, "Test_C28_108_SimpleStringOnce_IsInitialized", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if !s.IsInitialized() {
			t.Error("expected true")
		}
	})
}

func Test_C28_109_SimpleStringOnce_IsDefined(t *testing.T) {
	safeTest(t, "Test_C28_109_SimpleStringOnce_IsDefined", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if !s.IsDefined() {
			t.Error("expected true")
		}
	})
}

func Test_C28_110_SimpleStringOnce_IsUninitialized(t *testing.T) {
	safeTest(t, "Test_C28_110_SimpleStringOnce_IsUninitialized", func() {
		s := &corestr.SimpleStringOnce{}
		if !s.IsUninitialized() {
			t.Error("expected true")
		}
	})
}

func Test_C28_111_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_C28_111_SimpleStringOnce_Invalidate", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Invalidate()
		if s.IsInitialized() {
			t.Error("expected false")
		}
	})
}

func Test_C28_112_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_C28_112_SimpleStringOnce_Reset", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Reset()
		if s.IsInitialized() || s.Value() != "" {
			t.Error("expected reset")
		}
	})
}

func Test_C28_113_SimpleStringOnce_IsInvalid(t *testing.T) {
	safeTest(t, "Test_C28_113_SimpleStringOnce_IsInvalid", func() {
		s := &corestr.SimpleStringOnce{}
		if !s.IsInvalid() {
			t.Error("expected true")
		}
	})
}

func Test_C28_114_SimpleStringOnce_IsInvalid_Valid(t *testing.T) {
	safeTest(t, "Test_C28_114_SimpleStringOnce_IsInvalid_Valid", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if s.IsInvalid() {
			t.Error("expected false")
		}
	})
}

func Test_C28_115_SimpleStringOnce_ValueBytes(t *testing.T) {
	safeTest(t, "Test_C28_115_SimpleStringOnce_ValueBytes", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if string(s.ValueBytes()) != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_116_SimpleStringOnce_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_C28_116_SimpleStringOnce_ValueBytesPtr", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if string(s.ValueBytesPtr()) != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_117_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C28_117_SimpleStringOnce_SetOnUninitialized", func() {
		s := &corestr.SimpleStringOnce{}
		err := s.SetOnUninitialized("hello")
		if err != nil || s.Value() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_118_SimpleStringOnce_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C28_118_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		err := s.SetOnUninitialized("world")
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_C28_119_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C28_119_SimpleStringOnce_GetSetOnce", func() {
		s := &corestr.SimpleStringOnce{}
		v := s.GetSetOnce("hello")
		if v != "hello" {
			t.Error("expected hello")
		}
		v2 := s.GetSetOnce("world")
		if v2 != "hello" {
			t.Error("expected hello still")
		}
	})
}

func Test_C28_120_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_C28_120_SimpleStringOnce_GetOnce", func() {
		s := &corestr.SimpleStringOnce{}
		v := s.GetOnce()
		if v != "" {
			t.Error("expected empty")
		}
		if !s.IsInitialized() {
			t.Error("expected initialized")
		}
	})
}

func Test_C28_121_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C28_121_SimpleStringOnce_GetOnceFunc", func() {
		s := &corestr.SimpleStringOnce{}
		v := s.GetOnceFunc(func() string { return "hello" })
		if v != "hello" {
			t.Error("expected hello")
		}
		v2 := s.GetOnceFunc(func() string { return "world" })
		if v2 != "hello" {
			t.Error("expected hello still")
		}
	})
}

func Test_C28_122_SimpleStringOnce_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_C28_122_SimpleStringOnce_SetOnceIfUninitialized", func() {
		s := &corestr.SimpleStringOnce{}
		if !s.SetOnceIfUninitialized("hello") {
			t.Error("expected true")
		}
		if s.SetOnceIfUninitialized("world") {
			t.Error("expected false")
		}
	})
}

func Test_C28_123_SimpleStringOnce_SetInitialize(t *testing.T) {
	safeTest(t, "Test_C28_123_SimpleStringOnce_SetInitialize", func() {
		s := &corestr.SimpleStringOnce{}
		s.SetInitialize()
		if !s.IsInitialized() {
			t.Error("expected true")
		}
	})
}

func Test_C28_124_SimpleStringOnce_SetUnInit(t *testing.T) {
	safeTest(t, "Test_C28_124_SimpleStringOnce_SetUnInit", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		s.SetUnInit()
		if s.IsInitialized() {
			t.Error("expected false")
		}
	})
}

func Test_C28_125_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C28_125_SimpleStringOnce_ConcatNew", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		s2 := s.ConcatNew(" world")
		if s2.Value() != "hello world" {
			t.Error("expected hello world")
		}
	})
}

func Test_C28_126_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C28_126_SimpleStringOnce_ConcatNewUsingStrings", func() {
		s := corestr.New.SimpleStringOnce.Init("a")
		s2 := s.ConcatNewUsingStrings(",", "b", "c")
		if s2.Value() != "a,b,c" {
			t.Error("expected a,b,c")
		}
	})
}

func Test_C28_127_SimpleStringOnce_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C28_127_SimpleStringOnce_IsEmpty", func() {
		s := &corestr.SimpleStringOnce{}
		if !s.IsEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_C28_128_SimpleStringOnce_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C28_128_SimpleStringOnce_IsWhitespace", func() {
		s := corestr.New.SimpleStringOnce.Init("   ")
		if !s.IsWhitespace() {
			t.Error("expected true")
		}
	})
}

func Test_C28_129_SimpleStringOnce_Trim(t *testing.T) {
	safeTest(t, "Test_C28_129_SimpleStringOnce_Trim", func() {
		s := corestr.New.SimpleStringOnce.Init("  hello  ")
		if s.Trim() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_130_SimpleStringOnce_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C28_130_SimpleStringOnce_HasValidNonEmpty", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if !s.HasValidNonEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_C28_131_SimpleStringOnce_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C28_131_SimpleStringOnce_HasValidNonWhitespace", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if !s.HasValidNonWhitespace() {
			t.Error("expected true")
		}
	})
}

func Test_C28_132_SimpleStringOnce_IsValueBool(t *testing.T) {
	safeTest(t, "Test_C28_132_SimpleStringOnce_IsValueBool", func() {
		s := corestr.New.SimpleStringOnce.Init("true")
		if !s.IsValueBool() {
			t.Error("expected true")
		}
	})
}

func Test_C28_133_SimpleStringOnce_SafeValue(t *testing.T) {
	safeTest(t, "Test_C28_133_SimpleStringOnce_SafeValue", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if s.SafeValue() != "x" {
			t.Error("expected x")
		}
	})
}

func Test_C28_134_SimpleStringOnce_SafeValue_Uninit(t *testing.T) {
	safeTest(t, "Test_C28_134_SimpleStringOnce_SafeValue_Uninit", func() {
		s := &corestr.SimpleStringOnce{}
		if s.SafeValue() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C28_135_SimpleStringOnce_Uint16(t *testing.T) {
	safeTest(t, "Test_C28_135_SimpleStringOnce_Uint16", func() {
		s := corestr.New.SimpleStringOnce.Init("100")
		v, ok := s.Uint16()
		if !ok || v != 100 {
			t.Error("expected 100")
		}
	})
}

func Test_C28_136_SimpleStringOnce_Uint32(t *testing.T) {
	safeTest(t, "Test_C28_136_SimpleStringOnce_Uint32", func() {
		s := corestr.New.SimpleStringOnce.Init("1000")
		v, ok := s.Uint32()
		if !ok || v != 1000 {
			t.Error("expected 1000")
		}
	})
}

func Test_C28_137_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_C28_137_SimpleStringOnce_WithinRange", func() {
		s := corestr.New.SimpleStringOnce.Init("50")
		v, ok := s.WithinRange(true, 0, 100)
		if !ok || v != 50 {
			t.Error("expected 50")
		}
	})
}

func Test_C28_138_SimpleStringOnce_WithinRange_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C28_138_SimpleStringOnce_WithinRange_OutOfRange", func() {
		s := corestr.New.SimpleStringOnce.Init("200")
		v, ok := s.WithinRange(true, 0, 100)
		if ok || v != 100 {
			t.Error("expected 100 bounded")
		}
	})
}

func Test_C28_139_SimpleStringOnce_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C28_139_SimpleStringOnce_WithinRange_NoBoundary", func() {
		s := corestr.New.SimpleStringOnce.Init("200")
		v, ok := s.WithinRange(false, 0, 100)
		if ok || v != 200 {
			t.Error("expected 200 no boundary")
		}
	})
}

func Test_C28_140_SimpleStringOnce_WithinRange_Below(t *testing.T) {
	safeTest(t, "Test_C28_140_SimpleStringOnce_WithinRange_Below", func() {
		s := corestr.New.SimpleStringOnce.Init("-5")
		v, ok := s.WithinRange(true, 0, 100)
		if ok || v != 0 {
			t.Error("expected 0 bounded")
		}
	})
}

func Test_C28_141_SimpleStringOnce_WithinRange_ParseErr(t *testing.T) {
	safeTest(t, "Test_C28_141_SimpleStringOnce_WithinRange_ParseErr", func() {
		s := corestr.New.SimpleStringOnce.Init("abc")
		v, ok := s.WithinRange(true, 0, 100)
		if ok || v != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_142_SimpleStringOnce_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_C28_142_SimpleStringOnce_WithinRangeDefault", func() {
		s := corestr.New.SimpleStringOnce.Init("50")
		v, ok := s.WithinRangeDefault(0, 100)
		if !ok || v != 50 {
			t.Error("expected 50")
		}
	})
}

func Test_C28_143_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_C28_143_SimpleStringOnce_Int", func() {
		s := corestr.New.SimpleStringOnce.Init("42")
		if s.Int() != 42 {
			t.Error("expected 42")
		}
	})
}

func Test_C28_144_SimpleStringOnce_Int_Error(t *testing.T) {
	safeTest(t, "Test_C28_144_SimpleStringOnce_Int_Error", func() {
		s := corestr.New.SimpleStringOnce.Init("abc")
		if s.Int() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_145_SimpleStringOnce_Byte(t *testing.T) {
	safeTest(t, "Test_C28_145_SimpleStringOnce_Byte", func() {
		s := corestr.New.SimpleStringOnce.Init("200")
		if s.Byte() != 200 {
			t.Error("expected 200")
		}
	})
}

func Test_C28_146_SimpleStringOnce_Byte_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C28_146_SimpleStringOnce_Byte_OutOfRange", func() {
		s := corestr.New.SimpleStringOnce.Init("300")
		if s.Byte() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_147_SimpleStringOnce_Byte_Error(t *testing.T) {
	safeTest(t, "Test_C28_147_SimpleStringOnce_Byte_Error", func() {
		s := corestr.New.SimpleStringOnce.Init("abc")
		if s.Byte() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_148_SimpleStringOnce_Int16(t *testing.T) {
	safeTest(t, "Test_C28_148_SimpleStringOnce_Int16", func() {
		s := corestr.New.SimpleStringOnce.Init("100")
		if s.Int16() != 100 {
			t.Error("expected 100")
		}
	})
}

func Test_C28_149_SimpleStringOnce_Int16_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C28_149_SimpleStringOnce_Int16_OutOfRange", func() {
		s := corestr.New.SimpleStringOnce.Init("40000")
		if s.Int16() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_150_SimpleStringOnce_Int32(t *testing.T) {
	safeTest(t, "Test_C28_150_SimpleStringOnce_Int32", func() {
		s := corestr.New.SimpleStringOnce.Init("1000")
		if s.Int32() != 1000 {
			t.Error("expected 1000")
		}
	})
}

func Test_C28_151_SimpleStringOnce_Int32_OutOfRange(t *testing.T) {
	safeTest(t, "Test_C28_151_SimpleStringOnce_Int32_OutOfRange", func() {
		s := corestr.New.SimpleStringOnce.Init("3000000000")
		if s.Int32() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_152_SimpleStringOnce_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_C28_152_SimpleStringOnce_BooleanDefault", func() {
		s := corestr.New.SimpleStringOnce.Init("true")
		if !s.BooleanDefault() {
			t.Error("expected true")
		}
	})
}

func Test_C28_153_SimpleStringOnce_Boolean_Yes(t *testing.T) {
	safeTest(t, "Test_C28_153_SimpleStringOnce_Boolean_Yes", func() {
		s := corestr.New.SimpleStringOnce.Init("yes")
		if !s.Boolean(false) {
			t.Error("expected true")
		}
	})
}

func Test_C28_154_SimpleStringOnce_Boolean_Y(t *testing.T) {
	safeTest(t, "Test_C28_154_SimpleStringOnce_Boolean_Y", func() {
		s := corestr.New.SimpleStringOnce.Init("y")
		if !s.Boolean(false) {
			t.Error("expected true")
		}
	})
}

func Test_C28_155_SimpleStringOnce_Boolean_1(t *testing.T) {
	safeTest(t, "Test_C28_155_SimpleStringOnce_Boolean_1", func() {
		s := corestr.New.SimpleStringOnce.Init("1")
		if !s.Boolean(false) {
			t.Error("expected true")
		}
	})
}

func Test_C28_156_SimpleStringOnce_Boolean_YES(t *testing.T) {
	safeTest(t, "Test_C28_156_SimpleStringOnce_Boolean_YES", func() {
		s := corestr.New.SimpleStringOnce.Init("YES")
		if !s.Boolean(false) {
			t.Error("expected true")
		}
	})
}

func Test_C28_157_SimpleStringOnce_Boolean_CapY(t *testing.T) {
	safeTest(t, "Test_C28_157_SimpleStringOnce_Boolean_CapY", func() {
		s := corestr.New.SimpleStringOnce.Init("Y")
		if !s.Boolean(false) {
			t.Error("expected true")
		}
	})
}

func Test_C28_158_SimpleStringOnce_Boolean_Bad(t *testing.T) {
	safeTest(t, "Test_C28_158_SimpleStringOnce_Boolean_Bad", func() {
		s := corestr.New.SimpleStringOnce.Init("xyz")
		if s.Boolean(false) {
			t.Error("expected false")
		}
	})
}

func Test_C28_159_SimpleStringOnce_Boolean_ConsiderInit_Uninit(t *testing.T) {
	safeTest(t, "Test_C28_159_SimpleStringOnce_Boolean_ConsiderInit_Uninit", func() {
		s := &corestr.SimpleStringOnce{}
		if s.Boolean(true) {
			t.Error("expected false")
		}
	})
}

func Test_C28_160_SimpleStringOnce_IsSetter(t *testing.T) {
	safeTest(t, "Test_C28_160_SimpleStringOnce_IsSetter", func() {
		s := corestr.New.SimpleStringOnce.Init("yes")
		sv := s.IsSetter(false)
		if !sv.IsTrue() {
			t.Error("expected true")
		}
	})
}

func Test_C28_161_SimpleStringOnce_IsSetter_False(t *testing.T) {
	safeTest(t, "Test_C28_161_SimpleStringOnce_IsSetter_False", func() {
		s := &corestr.SimpleStringOnce{}
		sv := s.IsSetter(true)
		if sv.IsTrue() {
			t.Error("expected false")
		}
	})
}

func Test_C28_162_SimpleStringOnce_IsSetter_Bad(t *testing.T) {
	safeTest(t, "Test_C28_162_SimpleStringOnce_IsSetter_Bad", func() {
		s := corestr.New.SimpleStringOnce.Init("xyz")
		sv := s.IsSetter(false)
		_ = sv
	})
}

func Test_C28_163_SimpleStringOnce_ValueInt(t *testing.T) {
	safeTest(t, "Test_C28_163_SimpleStringOnce_ValueInt", func() {
		s := corestr.New.SimpleStringOnce.Init("42")
		if s.ValueInt(0) != 42 {
			t.Error("expected 42")
		}
	})
}

func Test_C28_164_SimpleStringOnce_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C28_164_SimpleStringOnce_ValueDefInt", func() {
		s := corestr.New.SimpleStringOnce.Init("10")
		if s.ValueDefInt() != 10 {
			t.Error("expected 10")
		}
	})
}

func Test_C28_165_SimpleStringOnce_ValueByte(t *testing.T) {
	safeTest(t, "Test_C28_165_SimpleStringOnce_ValueByte", func() {
		s := corestr.New.SimpleStringOnce.Init("200")
		if s.ValueByte(0) != 200 {
			t.Error("expected 200")
		}
	})
}

func Test_C28_166_SimpleStringOnce_ValueByte_Overflow(t *testing.T) {
	safeTest(t, "Test_C28_166_SimpleStringOnce_ValueByte_Overflow", func() {
		s := corestr.New.SimpleStringOnce.Init("300")
		if s.ValueByte(99) != 99 {
			t.Error("expected 99")
		}
	})
}

func Test_C28_167_SimpleStringOnce_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C28_167_SimpleStringOnce_ValueDefByte", func() {
		s := corestr.New.SimpleStringOnce.Init("100")
		if s.ValueDefByte() != 100 {
			t.Error("expected 100")
		}
	})
}

func Test_C28_168_SimpleStringOnce_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C28_168_SimpleStringOnce_ValueFloat64", func() {
		s := corestr.New.SimpleStringOnce.Init("3.14")
		if s.ValueFloat64(0) != 3.14 {
			t.Error("expected 3.14")
		}
	})
}

func Test_C28_169_SimpleStringOnce_ValueFloat64_Error(t *testing.T) {
	safeTest(t, "Test_C28_169_SimpleStringOnce_ValueFloat64_Error", func() {
		s := corestr.New.SimpleStringOnce.Init("abc")
		if s.ValueFloat64(1.0) != 1.0 {
			t.Error("expected 1.0")
		}
	})
}

func Test_C28_170_SimpleStringOnce_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C28_170_SimpleStringOnce_ValueDefFloat64", func() {
		s := corestr.New.SimpleStringOnce.Init("2.5")
		if s.ValueDefFloat64() != 2.5 {
			t.Error("expected 2.5")
		}
	})
}

func Test_C28_171_SimpleStringOnce_NonPtr(t *testing.T) {
	safeTest(t, "Test_C28_171_SimpleStringOnce_NonPtr", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		np := s.NonPtr()
		_ = np
	})
}

func Test_C28_172_SimpleStringOnce_Ptr(t *testing.T) {
	safeTest(t, "Test_C28_172_SimpleStringOnce_Ptr", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if s.Ptr() == nil {
			t.Error("expected same")
		}
	})
}

func Test_C28_173_SimpleStringOnce_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C28_173_SimpleStringOnce_HasSafeNonEmpty", func() {
		s := corestr.New.SimpleStringOnce.Init("x")
		if !s.HasSafeNonEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_C28_174_SimpleStringOnce_Is(t *testing.T) {
	safeTest(t, "Test_C28_174_SimpleStringOnce_Is", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		if !s.Is("hello") {
			t.Error("expected true")
		}
	})
}

func Test_C28_175_SimpleStringOnce_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C28_175_SimpleStringOnce_IsAnyOf", func() {
		s := corestr.New.SimpleStringOnce.Init("b")
		if !s.IsAnyOf("a", "b", "c") {
			t.Error("expected true")
		}
		if !s.IsAnyOf() {
			t.Error("expected true for empty")
		}
		if s.IsAnyOf("x", "y") {
			t.Error("expected false")
		}
	})
}

func Test_C28_176_SimpleStringOnce_IsContains(t *testing.T) {
	safeTest(t, "Test_C28_176_SimpleStringOnce_IsContains", func() {
		s := corestr.New.SimpleStringOnce.Init("hello world")
		if !s.IsContains("world") {
			t.Error("expected true")
		}
	})
}

func Test_C28_177_SimpleStringOnce_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C28_177_SimpleStringOnce_IsAnyContains", func() {
		s := corestr.New.SimpleStringOnce.Init("hello world")
		if !s.IsAnyContains("xyz", "world") {
			t.Error("expected true")
		}
		if !s.IsAnyContains() {
			t.Error("expected true for empty")
		}
	})
}

func Test_C28_178_SimpleStringOnce_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C28_178_SimpleStringOnce_IsEqualNonSensitive", func() {
		s := corestr.New.SimpleStringOnce.Init("Hello")
		if !s.IsEqualNonSensitive("hello") {
			t.Error("expected true")
		}
	})
}

func Test_C28_179_SimpleStringOnce_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C28_179_SimpleStringOnce_IsRegexMatches", func() {
		s := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)
		if !s.IsRegexMatches(re) {
			t.Error("expected true")
		}
		if s.IsRegexMatches(nil) {
			t.Error("expected false")
		}
	})
}

func Test_C28_180_SimpleStringOnce_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C28_180_SimpleStringOnce_RegexFindString", func() {
		s := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)
		if s.RegexFindString(re) != "123" {
			t.Error("expected 123")
		}
		if s.RegexFindString(nil) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C28_181_SimpleStringOnce_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C28_181_SimpleStringOnce_RegexFindAllStringsWithFlag", func() {
		s := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, has := s.RegexFindAllStringsWithFlag(re, -1)
		if !has || len(items) != 3 {
			t.Error("expected 3")
		}
		_, has2 := s.RegexFindAllStringsWithFlag(nil, -1)
		if has2 {
			t.Error("expected false")
		}
	})
}

func Test_C28_182_SimpleStringOnce_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C28_182_SimpleStringOnce_RegexFindAllStrings", func() {
		s := corestr.New.SimpleStringOnce.Init("a1b2")
		re := regexp.MustCompile(`\d`)
		items := s.RegexFindAllStrings(re, -1)
		if len(items) != 2 {
			t.Error("expected 2")
		}
		items2 := s.RegexFindAllStrings(nil, -1)
		if len(items2) != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C28_183_SimpleStringOnce_LinesSimpleSlice(t *testing.T) {
	safeTest(t, "Test_C28_183_SimpleStringOnce_LinesSimpleSlice", func() {
		s := corestr.New.SimpleStringOnce.Init("a\nb")
		sl := s.LinesSimpleSlice()
		if sl.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_184_SimpleStringOnce_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_C28_184_SimpleStringOnce_SimpleSlice", func() {
		s := corestr.New.SimpleStringOnce.Init("a,b,c")
		sl := s.SimpleSlice(",")
		if sl.Length() != 3 {
			t.Error("expected 3")
		}
	})
}

func Test_C28_185_SimpleStringOnce_Split(t *testing.T) {
	safeTest(t, "Test_C28_185_SimpleStringOnce_Split", func() {
		s := corestr.New.SimpleStringOnce.Init("a,b")
		sp := s.Split(",")
		if len(sp) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C28_186_SimpleStringOnce_SplitLeftRight(t *testing.T) {
	safeTest(t, "Test_C28_186_SimpleStringOnce_SplitLeftRight", func() {
		s := corestr.New.SimpleStringOnce.Init("key=value")
		left, right := s.SplitLeftRight("=")
		if left != "key" || right != "value" {
			t.Error("expected key, value")
		}
	})
}

func Test_C28_187_SimpleStringOnce_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_C28_187_SimpleStringOnce_SplitLeftRight_NoSep", func() {
		s := corestr.New.SimpleStringOnce.Init("nosep")
		left, right := s.SplitLeftRight("=")
		if left != "nosep" || right != "" {
			t.Error("expected nosep, empty")
		}
	})
}

func Test_C28_188_SimpleStringOnce_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_C28_188_SimpleStringOnce_SplitLeftRightTrim", func() {
		s := corestr.New.SimpleStringOnce.Init(" key = value ")
		left, right := s.SplitLeftRightTrim("=")
		if left != "key" || right != "value" {
			t.Error("expected trimmed key, value")
		}
	})
}

func Test_C28_189_SimpleStringOnce_SplitLeftRightTrim_NoSep(t *testing.T) {
	safeTest(t, "Test_C28_189_SimpleStringOnce_SplitLeftRightTrim_NoSep", func() {
		s := corestr.New.SimpleStringOnce.Init(" nosep ")
		left, right := s.SplitLeftRightTrim("=")
		if left != "nosep" || right != "" {
			t.Error("expected trimmed nosep, empty")
		}
	})
}

func Test_C28_190_SimpleStringOnce_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C28_190_SimpleStringOnce_SplitNonEmpty", func() {
		s := corestr.New.SimpleStringOnce.Init("a,,b")
		sp := s.SplitNonEmpty(",")
		_ = sp
	})
}

func Test_C28_191_SimpleStringOnce_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C28_191_SimpleStringOnce_SplitTrimNonWhitespace", func() {
		s := corestr.New.SimpleStringOnce.Init("a , b , c")
		sp := s.SplitTrimNonWhitespace(",")
		_ = sp
	})
}
	safeTest(t, "Test_C28_194_SimpleStringOnce_Clone", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		c := s.Clone()
		if c.Value() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C28_195_SimpleStringOnce_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_C28_195_SimpleStringOnce_CloneUsingNewVal", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		c := s.CloneUsingNewVal("world")
		if c.Value() != "world" {
			t.Error("expected world")
		}
	})
}

func Test_C28_196_SimpleStringOnce_Dispose(t *testing.T) {
	safeTest(t, "Test_C28_196_SimpleStringOnce_Dispose", func() {
		s := corestr.New.SimpleStringOnce.Init("hello")
		s.Dispose()
	})
}

func Test_C28_197_SimpleStringOnce_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C28_197_SimpleStringOnce_Dispose_Nil", func() {
		var s *corestr.SimpleStringOnce
		s.Dispose()
	})
}

func Test_C28_198_SimpleStringOnce_String(t *testing.T) {
	safeTest(t, "Test_C28_198_SimpleStringOnce_String", func() {
		s := corestr.New.SimpleStringOnce.Init("hi")
		if s.String() != "hi" {
			t.Error("expected hi")
		}
	})
}

func Test_C28_199_SimpleStringOnce_String_Nil(t *testing.T) {
	safeTest(t, "Test_C28_199_SimpleStringOnce_String_Nil", func() {
		var s *corestr.SimpleStringOnce
		if s.String() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C28_200_SimpleStringOnce_StringPtr(t *testing.T) {
	safeTest(t, "Test_C28_200_SimpleStringOnce_StringPtr", func() {
		s := corestr.New.SimpleStringOnce.Init("hi")
		if *s.StringPtr() != "hi" {
			t.Error("expected hi")
		}
	})
}
