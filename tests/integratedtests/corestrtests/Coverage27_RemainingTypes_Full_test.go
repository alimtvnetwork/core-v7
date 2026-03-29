package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =======================================================
// KeyValueCollection
// =======================================================

func Test_C27_KeyValueCollection_Add(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Add", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		if kvc.Length() != 2 {
			t.Errorf("expected 2 got %d", kvc.Length())
		}
	})
}

func Test_C27_KeyValueCollection_AddIf(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddIf", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddIf(false, "skip", "val")
		kvc.AddIf(true, "keep", "val")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Count", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if kvc.Count() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasAnyItem", func() {
		kvc := corestr.Empty.KeyValueCollection()
		if kvc.HasAnyItem() {
			t.Error("should not have items")
		}
		kvc.Add("k", "v")
		if !kvc.HasAnyItem() {
			t.Error("should have items")
		}
	})
}

func Test_C27_KeyValueCollection_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_LastIndex_HasIndex", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if kvc.LastIndex() != 0 {
			t.Error("expected 0")
		}
		if !kvc.HasIndex(0) {
			t.Error("should have index 0")
		}
	})
}

func Test_C27_KeyValueCollection_First_Last(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_First_Last", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		if kvc.First().Key != "k1" {
			t.Error("first should be k1")
		}
		if kvc.Last().Key != "k2" {
			t.Error("last should be k2")
		}
	})
}

func Test_C27_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_FirstOrDefault_Empty", func() {
		kvc := corestr.Empty.KeyValueCollection()
		if kvc.FirstOrDefault() != nil {
			t.Error("should be nil")
		}
	})
}

func Test_C27_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_LastOrDefault_Empty", func() {
		kvc := corestr.Empty.KeyValueCollection()
		if kvc.LastOrDefault() != nil {
			t.Error("should be nil")
		}
	})
}

func Test_C27_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Find", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		results := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "k1", false
		})
		if len(results) != 1 {
			t.Errorf("expected 1 got %d", len(results))
		}
	})
}

func Test_C27_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValueAt", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if kvc.SafeValueAt(0) != "v" {
			t.Error("expected v")
		}
		if kvc.SafeValueAt(5) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C27_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValuesAtIndexes", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		vals := kvc.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C27_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Strings", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.Strings()
		if len(strs) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_StringsUsingFormat", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.StringsUsingFormat("%s=%s")
		if len(strs) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_String", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.String()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Adds", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Adds(
			corestr.KeyValuePair{Key: "k1", Value: "v1"},
			corestr.KeyValuePair{Key: "k2", Value: "v2"},
		)
		if kvc.Length() != 2 {
			t.Errorf("expected 2 got %d", kvc.Length())
		}
	})
}

func Test_C27_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplit", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplit("=", "key=value")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplitTrim", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplitTrim("=", " key = value ")
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddMap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddMap(map[string]string{"k": "v"})
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashsetMap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddHashsetMap(map[string]bool{"a": true})
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashset", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		kvc.AddHashset(hs)
		if kvc.Length() != 2 {
			t.Errorf("expected 2 got %d", kvc.Length())
		}
	})
}

func Test_C27_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmap(hm)
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmaps", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmaps(hm)
		if kvc.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Hashmap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		hm := kvc.Hashmap()
		if hm.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_IsContains", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if !kvc.IsContains("k") {
			t.Error("should contain k")
		}
	})
}

func Test_C27_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Get", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		val, found := kvc.Get("k")
		if !found || val != "v" {
			t.Error("should find k=v")
		}
	})
}

func Test_C27_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasKey", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if !kvc.HasKey("k") {
			t.Error("should have key")
		}
	})
}

func Test_C27_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Map", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		m := kvc.Map()
		if m["k"] != "v" {
			t.Error("expected v")
		}
	})
}

func Test_C27_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeys", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		keys := kvc.AllKeys()
		if len(keys) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C27_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeysSorted", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("b", "v").Add("a", "v")
		keys := kvc.AllKeysSorted()
		if keys[0] != "a" {
			t.Error("first key should be a")
		}
	})
}

func Test_C27_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllValues", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		vals := kvc.AllValues()
		if len(vals) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Join", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Join(",")
		if result == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinKeys", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		result := kvc.JoinKeys(",")
		if result == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinValues", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.JoinValues(",")
		if result != "v" {
			t.Errorf("expected v got %s", result)
		}
	})
}

func Test_C27_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Json", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Json()
		if result.HasError() {
			t.Error("should not error")
		}
	})
}

func Test_C27_KeyValueCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonPtr", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		if kvc.JsonPtr() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C27_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Serialize", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		_, err := kvc.Serialize()
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SerializeMust", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		data := kvc.SerializeMust()
		if len(data) == 0 {
			t.Error("should have data")
		}
	})
}

func Test_C27_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_ParseInjectUsingJson", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		_, err := kvc2.ParseInjectUsingJson(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonParseSelfInject", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		err := kvc2.JsonParseSelfInject(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyValueCollection_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AsJsonInterfaces", func() {
		kvc := corestr.Empty.KeyValueCollection()
		if kvc.AsJsonContractsBinder() == nil {
			t.Error("should not be nil")
		}
		if kvc.AsJsoner() == nil {
			t.Error("should not be nil")
		}
		if kvc.AsJsonParseSelfInjector() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C27_KeyValueCollection_Clear(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Clear", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		kvc.Clear()
		if kvc.Length() != 0 {
			t.Error("should be 0")
		}
	})
}

func Test_C27_KeyValueCollection_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Dispose", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		kvc.Dispose()
	})
}

func Test_C27_KeyValueCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Deserialize", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		var target corestr.KeyValueCollection
		err := kvc.Deserialize(&target)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyValueCollection_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Compile", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.Compile()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

// =======================================================
// KeyAnyValuePair
// =======================================================

func Test_C27_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Basic", func() {
		kav := corestr.KeyAnyValuePair{Key: "name", Value: "John"}
		if kav.KeyName() != "name" {
			t.Error("expected name")
		}
		if kav.VariableName() != "name" {
			t.Error("expected name")
		}
		if kav.ValueAny() != "John" {
			t.Error("expected John")
		}
		if !kav.IsVariableNameEqual("name") {
			t.Error("should be equal")
		}
	})
}

func Test_C27_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_ValueString", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kav.ValueString()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueNull", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		if !kav.IsValueNull() {
			t.Error("should be null")
		}
	})
}

func Test_C27_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_HasNonNull", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if !kav.HasNonNull() {
			t.Error("should have value")
		}
		if !kav.HasValue() {
			t.Error("should have value")
		}
	})
}

func Test_C27_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueEmptyString", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		if !kav.IsValueEmptyString() {
			t.Error("should be empty string")
		}
	})
}

func Test_C27_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueWhitespace", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		if !kav.IsValueWhitespace() {
			t.Error("should be whitespace")
		}
	})
}

func Test_C27_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_String", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		s := kav.String()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Compile", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kav.Compile() == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_SerializeMust", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		data := kav.SerializeMust()
		if len(data) == 0 {
			t.Error("should have data")
		}
	})
}

func Test_C27_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Serialize", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		_, err := kav.Serialize()
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Json", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		result := kav.Json()
		if result.HasError() {
			t.Error("should not error")
		}
	})
}

func Test_C27_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_JsonPtr", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kav.JsonPtr() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C27_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_ParseInjectUsingJson", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jsonResult := kav.Json()
		kav2 := &corestr.KeyAnyValuePair{}
		_, err := kav2.ParseInjectUsingJson(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_KeyAnyValuePair_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_AsJsonInterfaces", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		if kav.AsJsonContractsBinder() == nil {
			t.Error("should not be nil")
		}
		if kav.AsJsoner() == nil {
			t.Error("should not be nil")
		}
		if kav.AsJsonParseSelfInjector() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C27_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Clear_Dispose", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()
		if kav.Key != "" {
			t.Error("should be cleared")
		}
		kav2 := &corestr.KeyAnyValuePair{Key: "k2", Value: "v2"}
		kav2.Dispose()
	})
}

// =======================================================
// ValidValue
// =======================================================

func Test_C27_ValidValue_NewValidValue(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValue", func() {
		vv := corestr.NewValidValue("hello")
		if vv.Value != "hello" || !vv.IsValid {
			t.Error("unexpected state")
		}
	})
}

func Test_C27_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueEmpty", func() {
		vv := corestr.NewValidValueEmpty()
		if vv.Value != "" || !vv.IsValid {
			t.Error("unexpected state")
		}
	})
}

func Test_C27_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValue", func() {
		vv := corestr.InvalidValidValue("err")
		if vv.IsValid || vv.Message != "err" {
			t.Error("unexpected state")
		}
	})
}

func Test_C27_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValueNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		if vv.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "test")
		if vv.Value == "" {
			t.Error("should have value")
		}
	})
}

func Test_C27_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		if vv.Value == "" {
			t.Error("should have value")
		}
	})
}

func Test_C27_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEmpty", func() {
		vv := corestr.NewValidValue("")
		if !vv.IsEmpty() {
			t.Error("should be empty")
		}
	})
}

func Test_C27_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("   ")
		if !vv.IsWhitespace() {
			t.Error("should be whitespace")
		}
	})
}

func Test_C27_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Trim", func() {
		vv := corestr.NewValidValue("  hello  ")
		if vv.Trim() != "hello" {
			t.Error("expected trimmed")
		}
	})
}

func Test_C27_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.HasValidNonEmpty() {
			t.Error("should have valid non-empty")
		}
	})
}

func Test_C27_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonWhitespace", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.HasValidNonWhitespace() {
			t.Error("should be valid non-whitespace")
		}
	})
}

func Test_C27_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		if !vv.ValueBool() {
			t.Error("expected true")
		}
		vv2 := corestr.NewValidValue("invalid")
		if vv2.ValueBool() {
			t.Error("expected false")
		}
		vv3 := corestr.NewValidValue("")
		if vv3.ValueBool() {
			t.Error("expected false for empty")
		}
	})
}

func Test_C27_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		if vv.ValueInt(0) != 42 {
			t.Error("expected 42")
		}
		vv2 := corestr.NewValidValue("invalid")
		if vv2.ValueInt(99) != 99 {
			t.Error("expected 99")
		}
	})
}

func Test_C27_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefInt", func() {
		vv := corestr.NewValidValue("10")
		if vv.ValueDefInt() != 10 {
			t.Error("expected 10")
		}
	})
}

func Test_C27_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueByte", func() {
		vv := corestr.NewValidValue("200")
		if vv.ValueByte(0) != 200 {
			t.Errorf("expected 200 got %d", vv.ValueByte(0))
		}
		vv2 := corestr.NewValidValue("300")
		if vv2.ValueByte(0) != 255 {
			t.Errorf("expected 255 got %d", vv2.ValueByte(0))
		}
	})
}

func Test_C27_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefByte", func() {
		vv := corestr.NewValidValue("100")
		if vv.ValueDefByte() != 100 {
			t.Error("expected 100")
		}
	})
}

func Test_C27_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		if vv.ValueFloat64(0) != 3.14 {
			t.Error("expected 3.14")
		}
	})
}

func Test_C27_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefFloat64", func() {
		vv := corestr.NewValidValue("2.5")
		if vv.ValueDefFloat64() != 2.5 {
			t.Error("expected 2.5")
		}
	})
}

func Test_C27_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("hello")
		bytes := vv.ValueBytesOnce()
		if len(bytes) != 5 {
			t.Error("expected 5 bytes")
		}
		// call again for cached path
		bytes2 := vv.ValueBytesOnce()
		if len(bytes2) != 5 {
			t.Error("expected 5 bytes cached")
		}
	})
}

func Test_C27_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("hi")
		bytes := vv.ValueBytesOncePtr()
		if len(bytes) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C27_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.HasSafeNonEmpty() {
			t.Error("expected true")
		}
	})
}

func Test_C27_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Is", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.Is("hello") {
			t.Error("should match")
		}
	})
}

func Test_C27_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		if !vv.IsAnyOf("a", "b", "c") {
			t.Error("should match")
		}
		if !vv.IsAnyOf() {
			t.Error("empty list returns true")
		}
	})
}

func Test_C27_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		if !vv.IsContains("world") {
			t.Error("should contain")
		}
	})
}

func Test_C27_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		if !vv.IsAnyContains("xyz", "world") {
			t.Error("should contain")
		}
	})
}

func Test_C27_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		if !vv.IsEqualNonSensitive("hello") {
			t.Error("should be equal")
		}
	})
}

func Test_C27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		if !vv.IsRegexMatches(re) {
			t.Error("should match")
		}
		if vv.IsRegexMatches(nil) {
			t.Error("nil regex should return false")
		}
	})
}

func Test_C27_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindString", func() {
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		result := vv.RegexFindString(re)
		if result != "123" {
			t.Errorf("expected 123 got %s", result)
		}
		if vv.RegexFindString(nil) != "" {
			t.Error("nil regex should return empty")
		}
	})
}

func Test_C27_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		results := vv.RegexFindAllStrings(re, -1)
		if len(results) != 3 {
			t.Errorf("expected 3 got %d", len(results))
		}
	})
}

func Test_C27_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
		if !hasAny || len(items) != 2 {
			t.Error("expected 2 items")
		}
	})
}

func Test_C27_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		if len(parts) != 3 {
			t.Error("expected 3")
		}
	})
}

func Test_C27_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_SplitNonEmpty", func() {
		vv := corestr.NewValidValue("a,,b")
		parts := vv.SplitNonEmpty(",")
		_ = parts
	})
}

func Test_C27_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_SplitTrimNonWhitespace", func() {
		vv := corestr.NewValidValue("a , , b")
		parts := vv.SplitTrimNonWhitespace(",")
		_ = parts
	})
}

func Test_C27_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()
		if cloned.Value != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C27_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_String", func() {
		vv := corestr.NewValidValue("hello")
		if vv.String() != "hello" {
			t.Error("expected hello")
		}
	})
}

func Test_C27_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("hello")
		s := vv.FullString()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Clear_Dispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		if vv.Value != "" {
			t.Error("should be cleared")
		}
		vv2 := corestr.NewValidValue("test")
		vv2.Dispose()
	})
}

func Test_C27_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Json", func() {
		vv := corestr.NewValidValue("hello")
		result := vv.Json()
		if result.HasError() {
			t.Error("should not error")
		}
	})
}

func Test_C27_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("hello")
		_, err := vv.Serialize()
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C27_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ParseInjectUsingJson", func() {
		vv := corestr.NewValidValue("hello")
		jsonResult := vv.Json()
		vv2 := &corestr.ValidValue{}
		_, err := vv2.ParseInjectUsingJson(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

// =======================================================
// ValidValues
// =======================================================

func Test_C27_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		if !vvs.IsEmpty() {
			t.Error("should be empty")
		}
	})
}

func Test_C27_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		if vvs.Length() != 2 {
			t.Errorf("expected 2 got %d", vvs.Length())
		}
	})
}

func Test_C27_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddFull", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "val", "msg")
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_Count_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Count_HasAnyItem", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		if vvs.Count() != 1 {
			t.Error("expected 1")
		}
		if !vvs.HasAnyItem() {
			t.Error("should have items")
		}
	})
}

func Test_C27_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Find", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		results := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "a", false
		})
		if len(results) != 1 {
			t.Errorf("expected 1 got %d", len(results))
		}
	})
}

func Test_C27_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		if vvs.SafeValueAt(0) != "a" {
			t.Error("expected a")
		}
		if vvs.SafeValueAt(5) != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C27_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		if vvs.SafeValidValueAt(0) != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C27_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		vals := vvs.SafeValuesAtIndexes(0, 1)
		if len(vals) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C27_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vals := vvs.SafeValidValuesAtIndexes(0)
		if len(vals) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.Strings()
		if len(strs) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_FullStrings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.FullStrings()
		if len(strs) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_String", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		if vvs.String() == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C27_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Adds", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Adds(corestr.ValidValue{Value: "a", IsValid: true})
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddsPtr", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddValidValues", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)
		if vvs1.Length() != 2 {
			t.Errorf("expected 2 got %d", vvs1.Length())
		}
	})
}

func Test_C27_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_ConcatNew", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs.ConcatNew(false, vvs2)
		if result.Length() != 2 {
			t.Errorf("expected 2 got %d", result.Length())
		}
	})
}

func Test_C27_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashsetMap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true})
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashset", func() {
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Hashmap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		hm := vvs.Hashmap()
		if hm.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Map", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		m := vvs.Map()
		if len(m) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C27_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_NewValidValuesUsingValues", func() {
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)
		if vvs.Length() != 2 {
			t.Errorf("expected 2 got %d", vvs.Length())
		}
	})
}

// =======================================================
// LeftRight
// =======================================================

func Test_C27_LeftRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_New", func() {
		lr := corestr.NewLeftRight("left", "right")
		if lr.Left != "left" || lr.Right != "right" {
			t.Error("unexpected values")
		}
	})
}

func Test_C27_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Invalid", func() {
		lr := corestr.InvalidLeftRight("err")
		if lr.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		if lr.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		if lr.Left != "a" || lr.Right != "b" {
			t.Error("unexpected values")
		}
	})
}

func Test_C27_LeftRight_UsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Single", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		if lr.Left != "a" || lr.Right != "" {
			t.Error("unexpected values")
		}
	})
}

func Test_C27_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice(nil)
		if lr.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlicePtr", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		if lr.Left != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C27_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_TrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		if lr.Left != "a" || lr.Right != "b" {
			t.Error("expected trimmed values")
		}
	})
}

func Test_C27_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Methods", func() {
		lr := corestr.NewLeftRight("left", "right")
		if len(lr.LeftBytes()) == 0 {
			t.Error("should have bytes")
		}
		if len(lr.RightBytes()) == 0 {
			t.Error("should have bytes")
		}
		if lr.LeftTrim() != "left" {
			t.Error("unexpected")
		}
		if lr.RightTrim() != "right" {
			t.Error("unexpected")
		}
		if lr.IsLeftEmpty() {
			t.Error("should not be empty")
		}
		if lr.IsRightEmpty() {
			t.Error("should not be empty")
		}
		if lr.IsLeftWhitespace() {
			t.Error("should not be whitespace")
		}
		if lr.IsRightWhitespace() {
			t.Error("should not be whitespace")
		}
		if !lr.HasValidNonEmptyLeft() {
			t.Error("should have valid non-empty left")
		}
		if !lr.HasValidNonEmptyRight() {
			t.Error("should have valid non-empty right")
		}
		if !lr.HasSafeNonEmpty() {
			t.Error("should have safe non-empty")
		}
		if !lr.Is("left", "right") {
			t.Error("should match")
		}
		if !lr.IsLeft("left") {
			t.Error("should match")
		}
		if !lr.IsRight("right") {
			t.Error("should match")
		}
	})
}

func Test_C27_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		if !lr1.IsEqual(lr2) {
			t.Error("should be equal")
		}
	})
}

func Test_C27_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		if cloned.Left != "a" {
			t.Error("expected a")
		}
	})
}

func Test_C27_LeftRight_IsRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_IsRegexMatch", func() {
		lr := corestr.NewLeftRight("hello123", "world")
		re := regexp.MustCompile(`\d+`)
		if !lr.IsLeftRegexMatch(re) {
			t.Error("should match")
		}
		if lr.IsRightRegexMatch(re) {
			t.Error("should not match")
		}
	})
}

func Test_C27_LeftRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Clear_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		lr2 := corestr.NewLeftRight("c", "d")
		lr2.Dispose()
	})
}

func Test_C27_LeftRight_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_NonPtr_Ptr", func() {
		lr := corestr.NewLeftRight("a", "b")
		nonPtr := lr.NonPtr()
		_ = nonPtr
		ptr := lr.Ptr()
		if ptr == nil {
			t.Error("should not be nil")
		}
	})
}

// =======================================================
// LeftMiddleRight
// =======================================================

func Test_C27_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		if lmr.Left != "l" || lmr.Middle != "m" || lmr.Right != "r" {
			t.Error("unexpected values")
		}
	})
}

func Test_C27_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		if lmr.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		if lmr.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Methods", func() {
		lmr := corestr.NewLeftMiddleRight("left", "mid", "right")
		if len(lmr.LeftBytes()) == 0 {
			t.Error("should have bytes")
		}
		if len(lmr.RightBytes()) == 0 {
			t.Error("should have bytes")
		}
		if len(lmr.MiddleBytes()) == 0 {
			t.Error("should have bytes")
		}
		if lmr.LeftTrim() != "left" {
			t.Error("unexpected")
		}
		if lmr.RightTrim() != "right" {
			t.Error("unexpected")
		}
		if lmr.MiddleTrim() != "mid" {
			t.Error("unexpected")
		}
		if lmr.IsLeftEmpty() || lmr.IsRightEmpty() || lmr.IsMiddleEmpty() {
			t.Error("should not be empty")
		}
		if lmr.IsLeftWhitespace() || lmr.IsRightWhitespace() || lmr.IsMiddleWhitespace() {
			t.Error("should not be whitespace")
		}
		if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyRight() || !lmr.HasValidNonEmptyMiddle() {
			t.Error("should have valid non-empty")
		}
		if !lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceRight() || !lmr.HasValidNonWhitespaceMiddle() {
			t.Error("should have valid non-whitespace")
		}
		if !lmr.HasSafeNonEmpty() {
			t.Error("should be safe non-empty")
		}
		if !lmr.IsAll("left", "mid", "right") {
			t.Error("should match all")
		}
		if !lmr.Is("left", "right") {
			t.Error("should match")
		}
	})
}

func Test_C27_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		cloned := lmr.Clone()
		if cloned.Left != "l" {
			t.Error("expected l")
		}
	})
}

func Test_C27_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lr := lmr.ToLeftRight()
		if lr.Left != "l" || lr.Right != "r" {
			t.Error("unexpected values")
		}
	})
}

func Test_C27_LeftMiddleRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Clear_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lmr.Clear()
		lmr2 := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr2.Dispose()
	})
}

// =======================================================
// TextWithLineNumber
// =======================================================

func Test_C27_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_HasLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		if !twl.HasLineNumber() {
			t.Error("should have line number")
		}
	})
}

func Test_C27_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsInvalidLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		if !twl.IsInvalidLineNumber() {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_Length", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		if twl.Length() != 5 {
			t.Errorf("expected 5 got %d", twl.Length())
		}
	})
}

func Test_C27_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmpty", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		if !twl.IsEmpty() {
			t.Error("should be empty")
		}
	})
}

func Test_C27_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyText", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		if !twl.IsEmptyText() {
			t.Error("should be empty text")
		}
	})
}

func Test_C27_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		if !twl.IsEmptyTextLineBoth() {
			t.Error("should be empty both")
		}
	})
}

// =======================================================
// ValueStatus
// =======================================================

func Test_C27_ValueStatus_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatus", func() {
		vs := corestr.InvalidValueStatus("err")
		if vs.ValueValid.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_ValueStatus_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatusNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		if vs.ValueValid.IsValid {
			t.Error("should be invalid")
		}
	})
}

func Test_C27_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_Clone", func() {
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		cloned := vs.Clone()
		if cloned.Index != 5 || cloned.ValueValid.Value != "hello" {
			t.Error("clone mismatch")
		}
	})
}

// =======================================================
// emptyCreator
// =======================================================

func Test_C27_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_C27_EmptyCreator_All", func() {
		if corestr.Empty.Collection() == nil {
			t.Error("nil")
		}
		if corestr.Empty.LinkedList() == nil {
			t.Error("nil")
		}
		if corestr.Empty.SimpleSlice() == nil {
			t.Error("nil")
		}
		if corestr.Empty.KeyAnyValuePair() == nil {
			t.Error("nil")
		}
		if corestr.Empty.KeyValuePair() == nil {
			t.Error("nil")
		}
		if corestr.Empty.KeyValueCollection() == nil {
			t.Error("nil")
		}
		if corestr.Empty.LinkedCollections() == nil {
			t.Error("nil")
		}
		if corestr.Empty.LeftRight() == nil {
			t.Error("nil")
		}
		sso := corestr.Empty.SimpleStringOnce()
		_ = sso
		if corestr.Empty.SimpleStringOncePtr() == nil {
			t.Error("nil")
		}
		if corestr.Empty.Hashset() == nil {
			t.Error("nil")
		}
		if corestr.Empty.HashsetsCollection() == nil {
			t.Error("nil")
		}
		if corestr.Empty.Hashmap() == nil {
			t.Error("nil")
		}
		if corestr.Empty.CharCollectionMap() == nil {
			t.Error("nil")
		}
		if corestr.Empty.KeyValuesCollection() == nil {
			t.Error("nil")
		}
		if corestr.Empty.CollectionsOfCollection() == nil {
			t.Error("nil")
		}
		if corestr.Empty.CharHashsetMap() == nil {
			t.Error("nil")
		}
	})
}

// =======================================================
// KeyValuePair (string methods)
// =======================================================

func Test_C27_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_KeyValuePair_Basic", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.String() == "" {
			t.Error("should not be empty")
		}
	})
}

// =======================================================
// DataModels
// =======================================================

func Test_C27_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashsetDataModel", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		if dm == nil {
			t.Error("nil")
		}
		hs2 := corestr.NewHashsetUsingDataModel(dm)
		if hs2 == nil || hs2.Length() != 2 {
			t.Error("unexpected")
		}
	})
}

func Test_C27_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashmapDataModel", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)
		if dm == nil {
			t.Error("nil")
		}
		hm2 := corestr.NewHashmapUsingDataModel(dm)
		if hm2 == nil || hm2.Length() != 1 {
			t.Error("unexpected")
		}
	})
}

func Test_C27_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashsetsCollectionDataModel", func() {
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1)
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		if dm == nil {
			t.Error("nil")
		}
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		if hsc2 == nil {
			t.Error("nil")
		}
	})
}
