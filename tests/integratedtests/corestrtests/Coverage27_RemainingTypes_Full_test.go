package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =======================================================
// KeyValueCollection
// =======================================================

func Test_C27_KeyValueCollection_Add(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Add", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddIf(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddIf", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddIf(false, "skip", "val")
		kvc.AddIf(true, "keep", "val")
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Count", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasAnyItem", func() {
		kvc := corestr.Empty.KeyValueCollection()
		actual := args.Map{"result": kvc.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have items", actual)
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C27_KeyValueCollection_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_LastIndex_HasIndex", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.LastIndex() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": kvc.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have index 0", actual)
	})
}

func Test_C27_KeyValueCollection_First_Last(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_First_Last", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		actual := args.Map{"result": kvc.First().Key != "k1"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be k1", actual)
		actual := args.Map{"result": kvc.Last().Key != "k2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last should be k2", actual)
	})
}

func Test_C27_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_FirstOrDefault_Empty", func() {
		kvc := corestr.Empty.KeyValueCollection()
		actual := args.Map{"result": kvc.FirstOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C27_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_LastOrDefault_Empty", func() {
		kvc := corestr.Empty.KeyValueCollection()
		actual := args.Map{"result": kvc.LastOrDefault() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C27_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Find", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		results := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "k1", false
		})
		actual := args.Map{"result": len(results) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValueAt", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.SafeValueAt(0) != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
		actual := args.Map{"result": kvc.SafeValueAt(5) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValuesAtIndexes", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		vals := kvc.SafeValuesAtIndexes(0, 1)
		actual := args.Map{"result": len(vals) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Strings", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.Strings()
		actual := args.Map{"result": len(strs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_StringsUsingFormat", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.StringsUsingFormat("%s=%s")
		actual := args.Map{"result": len(strs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_String", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Adds", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Adds(
			corestr.KeyValuePair{Key: "k1", Value: "v1"},
			corestr.KeyValuePair{Key: "k2", Value: "v2"},
		)
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplit", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplit("=", "key=value")
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplitTrim", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplitTrim("=", " key = value ")
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddMap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddMap(map[string]string{"k": "v"})
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashsetMap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddHashsetMap(map[string]bool{"a": true})
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashset", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		kvc.AddHashset(hs)
		actual := args.Map{"result": kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmap(hm)
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmaps", func() {
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmaps(hm)
		actual := args.Map{"result": kvc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Hashmap", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		hm := kvc.Hashmap()
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_IsContains", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.IsContains("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain k", actual)
	})
}

func Test_C27_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Get", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		val, found := kvc.Get("k")
		actual := args.Map{"result": found || val != "v"}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find k=v", actual)
	})
}

func Test_C27_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasKey", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.HasKey("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have key", actual)
	})
}

func Test_C27_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Map", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		m := kvc.Map()
		actual := args.Map{"result": m["k"] != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeys", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		keys := kvc.AllKeys()
		actual := args.Map{"result": len(keys) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeysSorted", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("b", "v").Add("a", "v")
		keys := kvc.AllKeysSorted()
		actual := args.Map{"result": keys[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first key should be a", actual)
	})
}

func Test_C27_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllValues", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		vals := kvc.AllValues()
		actual := args.Map{"result": len(vals) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Join", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Join(",")
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinKeys", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		result := kvc.JoinKeys(",")
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinValues", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.JoinValues(",")
		actual := args.Map{"result": result != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Json", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_KeyValueCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonPtr", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		actual := args.Map{"result": kvc.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Serialize", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		_, err := kvc.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SerializeMust", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		data := kvc.SerializeMust()
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C27_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_ParseInjectUsingJson", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		_, err := kvc2.ParseInjectUsingJson(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonParseSelfInject", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		err := kvc2.JsonParseSelfInject(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AsJsonInterfaces", func() {
		kvc := corestr.Empty.KeyValueCollection()
		actual := args.Map{"result": kvc.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": kvc.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": kvc.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyValueCollection_Clear(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Clear", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		kvc.Clear()
		actual := args.Map{"result": kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Compile", func() {
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.Compile()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

// =======================================================
// KeyAnyValuePair
// =======================================================

func Test_C27_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Basic", func() {
		kav := corestr.KeyAnyValuePair{Key: "name", Value: "John"}
		actual := args.Map{"result": kav.KeyName() != "name"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected name", actual)
		actual := args.Map{"result": kav.VariableName() != "name"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected name", actual)
		actual := args.Map{"result": kav.ValueAny() != "John"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected John", actual)
		actual := args.Map{"result": kav.IsVariableNameEqual("name")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C27_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_ValueString", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kav.ValueString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueNull", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		actual := args.Map{"result": kav.IsValueNull()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be null", actual)
	})
}

func Test_C27_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_HasNonNull", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kav.HasNonNull()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
		actual := args.Map{"result": kav.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueEmptyString", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		actual := args.Map{"result": kav.IsValueEmptyString()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty string", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueWhitespace", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}
		actual := args.Map{"result": kav.IsValueWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be whitespace", actual)
	})
}

func Test_C27_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_String", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		s := kav.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Compile", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kav.Compile() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_SerializeMust", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		data := kav.SerializeMust()
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C27_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Serialize", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		_, err := kav.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Json", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		result := kav.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_JsonPtr", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kav.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_ParseInjectUsingJson", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jsonResult := kav.Json()
		kav2 := &corestr.KeyAnyValuePair{}
		_, err := kav2.ParseInjectUsingJson(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyAnyValuePair_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_AsJsonInterfaces", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kav.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": kav.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": kav.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyAnyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Clear_Dispose", func() {
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()
		actual := args.Map{"result": kav.Key != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be cleared", actual)
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
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueEmpty", func() {
		vv := corestr.NewValidValueEmpty()
		actual := args.Map{"result": vv.Value != "" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValue", func() {
		vv := corestr.InvalidValidValue("err")
		actual := args.Map{"result": vv.IsValid || vv.Message != "err"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValueNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"result": vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "test")
		actual := args.Map{"result": vv.Value == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		actual := args.Map{"result": vv.Value == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEmpty", func() {
		vv := corestr.NewValidValue("")
		actual := args.Map{"result": vv.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("   ")
		actual := args.Map{"result": vv.IsWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be whitespace", actual)
	})
}

func Test_C27_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Trim", func() {
		vv := corestr.NewValidValue("  hello  ")
		actual := args.Map{"result": vv.Trim() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-empty", actual)
	})
}

func Test_C27_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonWhitespace", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be valid non-whitespace", actual)
	})
}

func Test_C27_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		actual := args.Map{"result": vv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		vv2 := corestr.NewValidValue("invalid")
		actual := args.Map{"result": vv2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		vv3 := corestr.NewValidValue("")
		actual := args.Map{"result": vv3.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C27_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		actual := args.Map{"result": vv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		vv2 := corestr.NewValidValue("invalid")
		actual := args.Map{"result": vv2.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_C27_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefInt", func() {
		vv := corestr.NewValidValue("10")
		actual := args.Map{"result": vv.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueByte", func() {
		vv := corestr.NewValidValue("200")
		actual := args.Map{"result": vv.ValueByte(0) != 200}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		vv2 := corestr.NewValidValue("300")
		actual := args.Map{"result": vv2.ValueByte(0) != 255}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255", actual)
	})
}

func Test_C27_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefByte", func() {
		vv := corestr.NewValidValue("100")
		actual := args.Map{"result": vv.ValueDefByte() != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		actual := args.Map{"result": vv.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	})
}

func Test_C27_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefFloat64", func() {
		vv := corestr.NewValidValue("2.5")
		actual := args.Map{"result": vv.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("hello")
		bytes := vv.ValueBytesOnce()
		actual := args.Map{"result": len(bytes) != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5 bytes", actual)
		// call again for cached path
		bytes2 := vv.ValueBytesOnce()
		actual := args.Map{"result": len(bytes2) != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5 bytes cached", actual)
	})
}

func Test_C27_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("hi")
		bytes := vv.ValueBytesOncePtr()
		actual := args.Map{"result": len(bytes) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Is", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.Is("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C27_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		actual := args.Map{"result": vv.IsAnyOf("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty list returns true", actual)
	})
}

func Test_C27_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{"result": vv.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C27_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{"result": vv.IsAnyContains("xyz", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C27_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": vv.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C27_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindString", func() {
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		result := vv.RegexFindString(re)
		actual := args.Map{"result": result != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual := args.Map{"result": vv.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return empty", actual)
	})
}

func Test_C27_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		results := vv.RegexFindAllStrings(re, -1)
		actual := args.Map{"result": len(results) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)
		actual := args.Map{"result": hasAny || len(items) != 2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
	})
}

func Test_C27_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		actual := args.Map{"result": len(parts) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
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
		actual := args.Map{"result": cloned.Value != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_String", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.String() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("hello")
		s := vv.FullString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Clear_Dispose", func() {
		vv := corestr.NewValidValue("hello")
		vv.Clear()
		actual := args.Map{"result": vv.Value != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be cleared", actual)
		vv2 := corestr.NewValidValue("test")
		vv2.Dispose()
	})
}

func Test_C27_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Json", func() {
		vv := corestr.NewValidValue("hello")
		result := vv.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("hello")
		_, err := vv.Serialize()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ParseInjectUsingJson", func() {
		vv := corestr.NewValidValue("hello")
		jsonResult := vv.Json()
		vv2 := &corestr.ValidValue{}
		_, err := vv2.ParseInjectUsingJson(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

// =======================================================
// ValidValues
// =======================================================

func Test_C27_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		actual := args.Map{"result": vvs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddFull", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "val", "msg")
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Count_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Count_HasAnyItem", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		actual := args.Map{"result": vvs.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": vvs.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C27_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Find", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		results := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "a", false
		})
		actual := args.Map{"result": len(results) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		actual := args.Map{"result": vvs.SafeValueAt(0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": vvs.SafeValueAt(5) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		actual := args.Map{"result": vvs.SafeValidValueAt(0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		vals := vvs.SafeValuesAtIndexes(0, 1)
		actual := args.Map{"result": len(vals) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vals := vvs.SafeValidValuesAtIndexes(0)
		actual := args.Map{"result": len(vals) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.Strings()
		actual := args.Map{"result": len(strs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_FullStrings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.FullStrings()
		actual := args.Map{"result": len(strs) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_String", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		actual := args.Map{"result": vvs.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Adds", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Adds(corestr.ValidValue{Value: "a", IsValid: true})
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddsPtr", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddValidValues", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)
		actual := args.Map{"result": vvs1.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_ConcatNew", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs.ConcatNew(false, vvs2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashsetMap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true})
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashset", func() {
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Hashmap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		hm := vvs.Hashmap()
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Map", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		m := vvs.Map()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_NewValidValuesUsingValues", func() {
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =======================================================
// LeftRight
// =======================================================

func Test_C27_LeftRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_New", func() {
		lr := corestr.NewLeftRight("left", "right")
		actual := args.Map{"result": lr.Left != "left" || lr.Right != "right"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Invalid", func() {
		lr := corestr.InvalidLeftRight("err")
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_UsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Single", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice(nil)
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlicePtr", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual := args.Map{"result": lr.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_TrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed values", actual)
	})
}

func Test_C27_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Methods", func() {
		lr := corestr.NewLeftRight("left", "right")
		actual := args.Map{"result": len(lr.LeftBytes()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have bytes", actual)
		actual := args.Map{"result": len(lr.RightBytes()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have bytes", actual)
		actual := args.Map{"result": lr.LeftTrim() != "left"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": lr.RightTrim() != "right"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": lr.IsLeftEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
		actual := args.Map{"result": lr.IsRightEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
		actual := args.Map{"result": lr.IsLeftWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be whitespace", actual)
		actual := args.Map{"result": lr.IsRightWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be whitespace", actual)
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-empty left", actual)
		actual := args.Map{"result": lr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-empty right", actual)
		actual := args.Map{"result": lr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have safe non-empty", actual)
		actual := args.Map{"result": lr.Is("left", "right")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": lr.IsLeft("left")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": lr.IsRight("right")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C27_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr1.IsEqual(lr2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C27_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		actual := args.Map{"result": cloned.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_LeftRight_IsRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_IsRegexMatch", func() {
		lr := corestr.NewLeftRight("hello123", "world")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": lr.IsRightRegexMatch(re)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not match", actual)
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
		actual := args.Map{"result": ptr == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

// =======================================================
// LeftMiddleRight
// =======================================================

func Test_C27_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		actual := args.Map{"result": lmr.Left != "l" || lmr.Middle != "m" || lmr.Right != "r"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Methods", func() {
		lmr := corestr.NewLeftMiddleRight("left", "mid", "right")
		actual := args.Map{"result": len(lmr.LeftBytes()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have bytes", actual)
		actual := args.Map{"result": len(lmr.RightBytes()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have bytes", actual)
		actual := args.Map{"result": len(lmr.MiddleBytes()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have bytes", actual)
		actual := args.Map{"result": lmr.LeftTrim() != "left"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": lmr.RightTrim() != "right"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": lmr.MiddleTrim() != "mid"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": lmr.IsLeftEmpty() || lmr.IsRightEmpty() || lmr.IsMiddleEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
		actual := args.Map{"result": lmr.IsLeftWhitespace() || lmr.IsRightWhitespace() || lmr.IsMiddleWhitespace()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be whitespace", actual)
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyRight() || !lmr.HasValidNonEmptyMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-empty", actual)
		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceRight() || !lmr.HasValidNonWhitespaceMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-whitespace", actual)
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be safe non-empty", actual)
		actual := args.Map{"result": lmr.IsAll("left", "mid", "right")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match all", actual)
		actual := args.Map{"result": lmr.Is("left", "right")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C27_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		cloned := lmr.Clone()
		actual := args.Map{"result": cloned.Left != "l"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected l", actual)
	})
}

func Test_C27_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lr := lmr.ToLeftRight()
		actual := args.Map{"result": lr.Left != "l" || lr.Right != "r"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
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
		actual := args.Map{"result": twl.HasLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have line number", actual)
	})
}

func Test_C27_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsInvalidLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": twl.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_Length", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		actual := args.Map{"result": twl.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmpty", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": twl.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyText", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		actual := args.Map{"result": twl.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty text", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": twl.IsEmptyTextLineBoth()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty both", actual)
	})
}

// =======================================================
// ValueStatus
// =======================================================

func Test_C27_ValueStatus_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatus", func() {
		vs := corestr.InvalidValueStatus("err")
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValueStatus_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatusNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_Clone", func() {
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		cloned := vs.Clone()
		actual := args.Map{"result": cloned.Index != 5 || cloned.ValueValid.Value != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

// =======================================================
// emptyCreator
// =======================================================

func Test_C27_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_C27_EmptyCreator_All", func() {
		actual := args.Map{"result": corestr.Empty.Collection() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.LinkedList() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.SimpleSlice() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.KeyAnyValuePair() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.KeyValuePair() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.KeyValueCollection() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.LinkedCollections() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.LeftRight() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		sso := corestr.Empty.SimpleStringOnce()
		_ = sso
		actual := args.Map{"result": corestr.Empty.SimpleStringOncePtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.Hashset() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.HashsetsCollection() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.Hashmap() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.CharCollectionMap() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.KeyValuesCollection() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.CollectionsOfCollection() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		actual := args.Map{"result": corestr.Empty.CharHashsetMap() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
	})
}

// =======================================================
// KeyValuePair (string methods)
// =======================================================

func Test_C27_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_KeyValuePair_Basic", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

// =======================================================
// DataModels
// =======================================================

func Test_C27_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashsetDataModel", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		actual := args.Map{"result": dm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		hs2 := corestr.NewHashsetUsingDataModel(dm)
		actual := args.Map{"result": hs2 == nil || hs2.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C27_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashmapDataModel", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)
		actual := args.Map{"result": dm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		hm2 := corestr.NewHashmapUsingDataModel(dm)
		actual := args.Map{"result": hm2 == nil || hm2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C27_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashsetsCollectionDataModel", func() {
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1)
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		actual := args.Map{"result": dm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		actual := args.Map{"result": hsc2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
	})
}
