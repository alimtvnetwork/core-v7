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
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddIf(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddIf", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddIf(false, "skip", "val")
		kvc.AddIf(true, "keep", "val")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Count(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Count", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasAnyItem", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()

		// Act
		actual := args.Map{"result": kvc.HasAnyItem()}

		// Assert
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
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.LastIndex() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual := args.Map{"result": kvc.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have index 0", actual)
	})
}

func Test_C27_KeyValueCollection_First_Last(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_First_Last", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")

		// Act
		actual := args.Map{"result": kvc.First().Key != "k1"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be k1", actual)
		actual := args.Map{"result": kvc.Last().Key != "k2"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "last should be k2", actual)
	})
}

func Test_C27_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()

		// Act
		actual := args.Map{"result": kvc.FirstOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C27_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_LastOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()

		// Act
		actual := args.Map{"result": kvc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be nil", actual)
	})
}

func Test_C27_KeyValueCollection_Find(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Find", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		results := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "k1", false
		})

		// Act
		actual := args.Map{"result": len(results) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValueAt", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.SafeValueAt(0) != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
		actual := args.Map{"result": kvc.SafeValueAt(5) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		vals := kvc.SafeValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{"result": len(vals) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_Strings(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Strings", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.Strings()

		// Act
		actual := args.Map{"result": len(strs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_StringsUsingFormat", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		strs := kvc.StringsUsingFormat("%s=%s")

		// Act
		actual := args.Map{"result": len(strs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_String", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Adds", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Adds(
			corestr.KeyValuePair{Key: "k1", Value: "v1"},
			corestr.KeyValuePair{Key: "k2", Value: "v2"},
		)

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplit", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplit("=", "key=value")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddStringBySplitTrim", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddStringBySplitTrim("=", " key = value ")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddMap", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddMap(map[string]string{"k": "v"})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashsetMap", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.AddHashsetMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddHashset", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		kvc.AddHashset(hs)

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmap", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmap(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AddsHashmaps", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		kvc.AddsHashmaps(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Hashmap", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_IsContains", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.IsContains("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain k", actual)
	})
}

func Test_C27_KeyValueCollection_Get(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Get", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		val, found := kvc.Get("k")

		// Act
		actual := args.Map{"result": found || val != "v"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find k=v", actual)
	})
}

func Test_C27_KeyValueCollection_HasKey(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_HasKey", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.HasKey("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have key", actual)
	})
}

func Test_C27_KeyValueCollection_Map(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Map", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		m := kvc.Map()

		// Act
		actual := args.Map{"result": m["k"] != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_KeyValueCollection_AllKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeys", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		keys := kvc.AllKeys()

		// Act
		actual := args.Map{"result": len(keys) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_KeyValueCollection_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllKeysSorted", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("b", "v").Add("a", "v")
		keys := kvc.AllKeysSorted()

		// Act
		actual := args.Map{"result": keys[0] != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first key should be a", actual)
	})
}

func Test_C27_KeyValueCollection_AllValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AllValues", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		vals := kvc.AllValues()

		// Act
		actual := args.Map{"result": len(vals) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_KeyValueCollection_Join(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Join", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Join(",")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_JoinKeys(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinKeys", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k1", "v1").Add("k2", "v2")
		result := kvc.JoinKeys(",")

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyValueCollection_JoinValues(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JoinValues", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.JoinValues(",")

		// Act
		actual := args.Map{"result": result != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_KeyValueCollection_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Json", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		result := kvc.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_KeyValueCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonPtr", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyValueCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Serialize", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		_, err := kvc.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_SerializeMust", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		data := kvc.SerializeMust()

		// Act
		actual := args.Map{"result": len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C27_KeyValueCollection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_ParseInjectUsingJson", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		_, err := kvc2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_JsonParseSelfInject", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		jsonResult := kvc.Json()
		kvc2 := corestr.Empty.KeyValueCollection()
		err := kvc2.JsonParseSelfInject(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_AsJsonInterfaces", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()

		// Act
		actual := args.Map{"result": kvc.AsJsonContractsBinder() == nil}

		// Assert
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
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		kvc.Clear()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
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
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		var target corestr.KeyValueCollection
		err := kvc.Deserialize(&target)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyValueCollection_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyValueCollection_Compile", func() {
		// Arrange
		kvc := corestr.Empty.KeyValueCollection()
		kvc.Add("k", "v")
		s := kvc.Compile()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

// =======================================================
// KeyAnyValuePair
// =======================================================

func Test_C27_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Basic", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "name", Value: "John"}

		// Act
		actual := args.Map{"result": kav.KeyName() != "name"}

		// Assert
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
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kav.ValueString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueNull", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{"result": kav.IsValueNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be null", actual)
	})
}

func Test_C27_KeyAnyValuePair_HasNonNull(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_HasNonNull", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kav.HasNonNull()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
		actual := args.Map{"result": kav.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueEmptyString", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{"result": kav.IsValueEmptyString()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty string", actual)
	})
}

func Test_C27_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_IsValueWhitespace", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{"result": kav.IsValueWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be whitespace", actual)
	})
}

func Test_C27_KeyAnyValuePair_String(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_String", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		s := kav.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_Compile(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Compile", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kav.Compile() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_KeyAnyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_SerializeMust", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		data := kav.SerializeMust()

		// Act
		actual := args.Map{"result": len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C27_KeyAnyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Serialize", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		_, err := kav.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyAnyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_Json", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		result := kav.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_KeyAnyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_JsonPtr", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kav.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C27_KeyAnyValuePair_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_ParseInjectUsingJson", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jsonResult := kav.Json()
		kav2 := &corestr.KeyAnyValuePair{}
		_, err := kav2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_KeyAnyValuePair_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C27_KeyAnyValuePair_AsJsonInterfaces", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kav.AsJsonContractsBinder() == nil}

		// Assert
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
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()

		// Act
		actual := args.Map{"result": kav.Key != ""}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueEmpty", func() {
		// Arrange
		vv := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{"result": vv.Value != "" || !vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValue", func() {
		// Arrange
		vv := corestr.InvalidValidValue("err")

		// Act
		actual := args.Map{"result": vv.IsValid || vv.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected state", actual)
	})
}

func Test_C27_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_InvalidValidValueNoMessage", func() {
		// Arrange
		vv := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{"result": vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAny", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAny(false, true, "test")

		// Act
		actual := args.Map{"result": vv.Value == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")

		// Act
		actual := args.Map{"result": vv.Value == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have value", actual)
	})
}

func Test_C27_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("")

		// Act
		actual := args.Map{"result": vv.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("   ")

		// Act
		actual := args.Map{"result": vv.IsWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be whitespace", actual)
	})
}

func Test_C27_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue("  hello  ")

		// Act
		actual := args.Map{"result": vv.Trim() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have valid non-empty", actual)
	})
}

func Test_C27_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.HasValidNonWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be valid non-whitespace", actual)
	})
}

func Test_C27_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBool", func() {
		// Arrange
		vv := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"result": vv.ValueBool()}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"result": vv.ValueInt(0) != 42}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("10")

		// Act
		actual := args.Map{"result": vv.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueByte", func() {
		// Arrange
		vv := corestr.NewValidValue("200")

		// Act
		actual := args.Map{"result": vv.ValueByte(0) != 200}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("100")

		// Act
		actual := args.Map{"result": vv.ValueDefByte() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{"result": vv.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	})
}

func Test_C27_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueDefFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("2.5")

		// Act
		actual := args.Map{"result": vv.ValueDefFloat64() != 2.5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		bytes := vv.ValueBytesOnce()

		// Act
		actual := args.Map{"result": len(bytes) != 5}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("hi")
		bytes := vv.ValueBytesOncePtr()

		// Act
		actual := args.Map{"result": len(bytes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.Is("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C27_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{"result": vv.IsAnyOf("a", "b", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty list returns true", actual)
	})
}

func Test_C27_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": vv.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C27_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": vv.IsAnyContains("xyz", "world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C27_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": vv.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C27_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("hello123")
		re := regexp.MustCompile(`\d+`)
		result := vv.RegexFindString(re)

		// Act
		actual := args.Map{"result": result != "123"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual := args.Map{"result": vv.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return empty", actual)
	})
}

func Test_C27_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		results := vv.RegexFindAllStrings(re, -1)

		// Act
		actual := args.Map{"result": len(results) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)

		// Act
		actual := args.Map{"result": hasAny || len(items) != 2}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
	})
}

func Test_C27_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")

		// Act
		actual := args.Map{"result": len(parts) != 3}

		// Assert
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
		// Arrange
		vv := corestr.NewValidValue("hello")
		cloned := vv.Clone()

		// Act
		actual := args.Map{"result": cloned.Value != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_String", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.String() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		s := vv.FullString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Clear_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vv.Clear()

		// Act
		actual := args.Map{"result": vv.Value != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be cleared", actual)
		vv2 := corestr.NewValidValue("test")
		vv2.Dispose()
	})
}

func Test_C27_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Json", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		result := vv.Json()

		// Act
		actual := args.Map{"result": result.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C27_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_Serialize", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		_, err := vv.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C27_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		jsonResult := vv.Json()
		vv2 := &corestr.ValidValue{}
		_, err := vv2.ParseInjectUsingJson(&jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

// =======================================================
// ValidValues
// =======================================================

func Test_C27_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Add", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddFull", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "val", "msg")

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Count_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Count_HasAnyItem", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": vvs.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C27_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Find", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		results := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "a", false
		})

		// Act
		actual := args.Map{"result": len(results) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.SafeValueAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": vvs.SafeValueAt(5) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.SafeValidValueAt(0) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		vals := vvs.SafeValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{"result": len(vals) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_SafeValidValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vals := vvs.SafeValidValuesAtIndexes(0)

		// Act
		actual := args.Map{"result": len(vals) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Strings", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.Strings()

		// Act
		actual := args.Map{"result": len(strs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_FullStrings", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		strs := vvs.FullStrings()

		// Act
		actual := args.Map{"result": len(strs) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_String", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C27_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Adds", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Adds(corestr.ValidValue{Value: "a", IsValid: true})

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddsPtr", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddValidValues", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)

		// Act
		actual := args.Map{"result": vvs1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_ConcatNew", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs.ConcatNew(false, vvs2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_AddHashset", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Hashmap", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		hm := vvs.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_Map", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		m := vvs.Map()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C27_ValidValues_NewValidValuesUsingValues", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

// =======================================================
// LeftRight
// =======================================================

func Test_C27_LeftRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_New", func() {
		// Arrange
		lr := corestr.NewLeftRight("left", "right")

		// Act
		actual := args.Map{"result": lr.Left != "left" || lr.Right != "right"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Invalid", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("err")

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_InvalidNoMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_UsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Single", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a"})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice(nil)

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftRight_UsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_UsingSlicePtr", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_TrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed values", actual)
	})
}

func Test_C27_LeftRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Methods", func() {
		// Arrange
		lr := corestr.NewLeftRight("left", "right")

		// Act
		actual := args.Map{"result": len(lr.LeftBytes()) == 0}

		// Assert
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
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr1.IsEqual(lr2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C27_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()

		// Act
		actual := args.Map{"result": cloned.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_LeftRight_IsRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_LeftRight_IsRegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello123", "world")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}

		// Assert
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
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		nonPtr := lr.NonPtr()
		_ = nonPtr
		ptr := lr.Ptr()

		// Act
		actual := args.Map{"result": ptr == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

// =======================================================
// LeftMiddleRight
// =======================================================

func Test_C27_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{"result": lmr.Left != "l" || lmr.Middle != "m" || lmr.Right != "r"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected values", actual)
	})
}

func Test_C27_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Invalid", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_LeftMiddleRight_Methods(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_Methods", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("left", "mid", "right")

		// Act
		actual := args.Map{"result": len(lmr.LeftBytes()) == 0}

		// Assert
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
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		cloned := lmr.Clone()

		// Act
		actual := args.Map{"result": cloned.Left != "l"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected l", actual)
	})
}

func Test_C27_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{"result": lr.Left != "l" || lr.Right != "r"}

		// Assert
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
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{"result": twl.HasLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have line number", actual)
	})
}

func Test_C27_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsInvalidLineNumber", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": twl.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_Length", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act
		actual := args.Map{"result": twl.Length() != 5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmpty", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": twl.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyText", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{"result": twl.IsEmptyText()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty text", actual)
	})
}

func Test_C27_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C27_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		// Arrange
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": twl.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty both", actual)
	})
}

// =======================================================
// ValueStatus
// =======================================================

func Test_C27_ValueStatus_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValueStatus_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_InvalidValueStatusNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C27_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C27_ValueStatus_Clone", func() {
		// Arrange
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		cloned := vs.Clone()

		// Act
		actual := args.Map{"result": cloned.Index != 5 || cloned.ValueValid.Value != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

// =======================================================
// emptyCreator
// =======================================================

func Test_C27_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_C27_EmptyCreator_All", func() {
		// Act
		actual := args.Map{"result": corestr.Empty.Collection() == nil}

		// Assert
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
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

// =======================================================
// DataModels
// =======================================================

func Test_C27_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C27_HashsetDataModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		dm := corestr.NewHashsetsDataModelUsing(hs)

		// Act
		actual := args.Map{"result": dm == nil}

		// Assert
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
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)

		// Act
		actual := args.Map{"result": dm == nil}

		// Assert
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
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1)
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)

		// Act
		actual := args.Map{"result": dm == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		actual := args.Map{"result": hsc2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil", actual)
	})
}
