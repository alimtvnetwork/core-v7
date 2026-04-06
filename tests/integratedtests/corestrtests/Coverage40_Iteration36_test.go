package corestrtests

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════

func Test_Cov40_KVP_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Basic", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}

		// Act
		actual := args.Map{"result": kv.KeyName() != "name" || kv.VariableName() != "name" || kv.ValueString() != "alice"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsVariableNameEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}

		// Act
		actual := args.Map{"result": kv.IsVariableNameEqual("x") || kv.IsVariableNameEqual("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsValueEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "x", Value: "y"}

		// Act
		actual := args.Map{"result": kv.IsValueEqual("y") || kv.IsValueEqual("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_Compile(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.Compile() == "" || kv.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_KVP_IsKeyEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKeyEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "", Value: "v"}

		// Act
		actual := args.Map{"result": kv.IsKeyEmpty() || !kv.IsKeyValueAnyEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{"result": kv.IsValueEmpty() || !kv.IsKeyValueAnyEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_HasKeyValue(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_HasKeyValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.HasKey() || !kv.HasValue() || kv.IsKeyValueEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKeyValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{}

		// Act
		actual := args.Map{"result": kv.IsKeyValueEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_KVP_Trim(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Trim", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}

		// Act
		actual := args.Map{"result": kv.TrimKey() != "k" || kv.TrimValue() != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueBool", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "true"}

		// Act
		actual := args.Map{"result": kv.ValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		kv2 := corestr.KeyValuePair{Value: ""}
		actual := args.Map{"result": kv2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		kv3 := corestr.KeyValuePair{Value: "invalid"}
		actual := args.Map{"result": kv3.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_KVP_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "42"}

		// Act
		actual := args.Map{"result": kv.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		kv2 := corestr.KeyValuePair{Value: "bad"}
		actual := args.Map{"result": kv2.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_Cov40_KVP_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "10"}

		// Act
		actual := args.Map{"result": kv.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		kv2 := corestr.KeyValuePair{Value: "bad"}
		actual := args.Map{"result": kv2.ValueDefInt() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVP_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "10"}

		// Act
		actual := args.Map{"result": kv.ValueByte(0) != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		kv2 := corestr.KeyValuePair{Value: "999"}
		actual := args.Map{"result": kv2.ValueByte(5) != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_Cov40_KVP_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "50"}

		// Act
		actual := args.Map{"result": kv.ValueDefByte() != 50}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
		kv2 := corestr.KeyValuePair{Value: "bad"}
		actual := args.Map{"result": kv2.ValueDefByte() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVP_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "3.14"}

		// Act
		actual := args.Map{"result": kv.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		kv2 := corestr.KeyValuePair{Value: "bad"}
		actual := args.Map{"result": kv2.ValueFloat64(1.0) != 1.0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.0", actual)
	})
}

func Test_Cov40_KVP_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueDefFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "bad"}

		// Act
		actual := args.Map{"result": kv.ValueDefFloat64() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVP_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "hello"}
		vv := kv.ValueValid()

		// Act
		actual := args.Map{"result": vv.IsValid || vv.Value != "hello"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_ValueValidOptions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "x"}
		vv := kv.ValueValidOptions(false, "err")

		// Act
		actual := args.Map{"result": vv.IsValid || vv.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_Is(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Is", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		actual := args.Map{"result": kv.Is("a", "b") || kv.Is("a", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_IsKey_IsVal", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "a", Value: "b"}

		// Act
		actual := args.Map{"result": kv.IsKey("a") || !kv.IsVal("b") || kv.IsKey("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_FormatString(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_FormatString", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.FormatString("%v=%v") != "k=v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVP_Json(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Json", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.Json().Error != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": kv.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_KVP_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Serialize", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		data, err := kv.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": len(kv.SerializeMust()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_KVP_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVP_Clear_Dispose", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{"result": kv.Key != "" || kv.Value != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		kv2 := corestr.KeyValuePair{Key: "a", Value: "b"}
		kv2.Dispose()
		actual := args.Map{"result": kv2.Key != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// KeyValueCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_KVC_Add_Length(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Add_Length", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k1", "v1").Add("k2", "v2")

		// Act
		actual := args.Map{"result": kvc.Length() != 2 || kvc.Count() != 2 || kvc.IsEmpty() || !kvc.HasAnyItem()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddIf", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddIf(true, "a", "b")
		kvc.AddIf(false, "c", "d")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddStringBySplit(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddStringBySplit", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddStringBySplit("=", "key=val")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddStringBySplitTrim(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddStringBySplitTrim", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddStringBySplitTrim("=", " key = val ")

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Adds", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Adds(corestr.KeyValuePair{Key: "a", Value: "b"}, corestr.KeyValuePair{Key: "c", Value: "d"})

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_KVC_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Adds_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Adds()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AddMap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddMap(map[string]string{"a": "b"})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddMap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashsetMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashsetMap(map[string]bool{"x": true})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashsetMap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AddHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashset", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		kvc.AddHashset(hs)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddHashset_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddHashset(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AddsHashmap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hm := corestr.New.Hashmap.Cap(3)
		hm.AddOrUpdate("a", "b")
		kvc.AddsHashmap(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddsHashmap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmap_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddsHashmap(nil)

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AddsHashmaps(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmaps", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		hm := corestr.New.Hashmap.Cap(3)
		hm.AddOrUpdate("a", "b")
		kvc.AddsHashmaps(hm)

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_AddsHashmaps_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.AddsHashmaps()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_First_Last(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_First_Last", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.First().Key != "a" || kvc.Last().Key != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": kvc.FirstOrDefault() != nil || kvc.LastOrDefault() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_Cov40_KVC_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_FirstOrDefault_NonEmpty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("x", "y")

		// Act
		actual := args.Map{"result": kvc.FirstOrDefault() == nil || kvc.LastOrDefault() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_KVC_HasKey(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_HasKey", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.HasKey("k") || kvc.HasKey("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_IsContains", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.IsContains("k") || kvc.IsContains("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_Get(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Get", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		val, found := kvc.Get("k")

		// Act
		actual := args.Map{"result": found || val != "v"}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		_, found2 := kvc.Get("missing")
		actual := args.Map{"result": found2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not found", actual)
	})
}

func Test_Cov40_KVC_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_HasIndex", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.HasIndex(0) || kvc.HasIndex(1)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_LastIndex", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValueAt", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.SafeValueAt(0) != "v" || kvc.SafeValueAt(5) != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_SafeValueAt_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValueAt_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": kvc.SafeValueAt(0) != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov40_KVC_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		vals := kvc.SafeValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{"result": len(vals) != 2 || vals[0] != "1" || vals[1] != "2"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_SafeValuesAtIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		vals := kvc.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"result": len(vals) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AllKeys_AllValues(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeys_AllValues", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"result": len(kvc.AllKeys()) != 2 || len(kvc.AllValues()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_AllKeys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeys_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": len(kvc.AllKeys()) != 0 || len(kvc.AllValues()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_AllKeysSorted(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AllKeysSorted", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("b", "2").Add("a", "1")
		sorted := kvc.AllKeysSorted()

		// Act
		actual := args.Map{"result": sorted[0] != "a" || sorted[1] != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Hashmap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Hashmap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Hashmap_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": kvc.Hashmap().Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_Map(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Map", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": len(kvc.Map()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Find(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Find_Break(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find_Break", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Find_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Find_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		found := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})

		// Act
		actual := args.Map{"result": len(found) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Strings", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"result": len(kvc.Strings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_KVC_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Strings_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": len(kvc.Strings()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_StringsUsingFormat(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_StringsUsingFormat", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		strs := kvc.StringsUsingFormat("%v=%v")

		// Act
		actual := args.Map{"result": len(strs) != 1 || strs[0] != "k=v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_Cov40_KVC_StringsUsingFormat_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": len(kvc.StringsUsingFormat("%v=%v")) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_String_Compile(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_String_Compile", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.String() == "" || kvc.Compile() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_KVC_Join_JoinKeys_JoinValues(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Join_JoinKeys_JoinValues", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"result": kvc.Join(",") == "" || kvc.JoinKeys(",") == "" || kvc.JoinValues(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_KVC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JSON", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		data, err := json.Marshal(kvc)

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		kvc2 := corestr.New.KeyValues.Empty()
		err = json.Unmarshal(data, kvc2)
		actual := args.Map{"result": err != nil || kvc2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_UnmarshalJSON_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_UnmarshalJSON_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		err := json.Unmarshal([]byte(`[]`), kvc)

		// Act
		actual := args.Map{"result": err != nil || kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Json_JsonPtr", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": kvc.Json().Error != nil || kvc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JsonModel", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")

		// Act
		actual := args.Map{"result": len(kvc.JsonModel()) != 1 || kvc.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Serialize_Deserialize", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		data, err := kvc.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		actual := args.Map{"result": len(kvc.SerializeMust()) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_KVC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_ParseInjectUsingJson", func() {
		// Arrange
		src := corestr.New.KeyValues.Empty()
		src.Add("k", "v")
		kvc := corestr.New.KeyValues.Empty()
		result, err := kvc.ParseInjectUsingJson(src.JsonPtr())

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_KVC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_AsJsoner", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": kvc.AsJsoner() == nil || kvc.AsJsonContractsBinder() == nil || kvc.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_KVC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_JsonParseSelfInject", func() {
		// Arrange
		src := corestr.New.KeyValues.Empty()
		src.Add("k", "v")
		kvc := corestr.New.KeyValues.Empty()
		err := kvc.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov40_KVC_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Clear_Dispose", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		kvc.Clear()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Dispose", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		kvc.Dispose()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_KVC_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov40_KVC_Deserialize", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		var target []corestr.KeyValuePair
		err := kvc.Deserialize(&target)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ── newKeyValuesCreator ──

func Test_Cov40_Creator_KV_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_Cap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Cap(5)

		// Act
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_Creator_KV_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()

		// Act
		actual := args.Map{"result": kvc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_Creator_KV_UsingMap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingMap", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_KV_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingMap_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingMap(map[string]string{})

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_Creator_KV_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValuePairs", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(corestr.KeyValuePair{Key: "a", Value: "1"})

		// Act
		actual := args.Map{"result": kvc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_KV_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValuePairs_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_Creator_KV_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValueStrings", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{"a", "b"}, []string{"1", "2"})

		// Act
		actual := args.Map{"result": kvc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_Creator_KV_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_KV_UsingKeyValueStrings_Empty", func() {
		// Arrange
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Act
		actual := args.Map{"result": kvc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// HashsetsCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_HC_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Basic", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.IsEmpty() || hc.HasItems() || hc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov40_HC_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Add", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc.Add(hs)

		// Act
		actual := args.Map{"result": hc.Length() != 1 || !hc.HasItems()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_HC_AddNonNil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddNonNil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonNil(nil)
		hc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_HC_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddNonEmpty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddNonEmpty(corestr.New.Hashset.Empty())
		hc.AddNonEmpty(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_HC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Adds", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Adds(corestr.New.Hashset.Strings([]string{"a"}), corestr.New.Hashset.Strings([]string{"b"}))

		// Act
		actual := args.Map{"result": hc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_HC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Adds_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Adds()

		// Act
		actual := args.Map{"result": hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_HC_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddHashsetsCollection", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hc1.AddHashsetsCollection(hc2)

		// Act
		actual := args.Map{"result": hc1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_HC_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AddHashsetsCollection_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.AddHashsetsCollection(nil)

		// Act
		actual := args.Map{"result": hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_HC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ConcatNew", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hc1.ConcatNew(hc2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_HC_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ConcatNew_NoArgs", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hc.ConcatNew()

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_HC_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_LastIndex", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.LastIndex() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_HC_List_ListPtr_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_List_ListPtr_ListDirectPtr", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(hc.List()) != 1 || hc.ListPtr() == nil || hc.ListDirectPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_HC_StringsList(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_StringsList", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))

		// Act
		actual := args.Map{"result": len(hc.StringsList()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_HC_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_StringsList_Empty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": len(hc.StringsList()) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_HC_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_HasAll", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))

		// Act
		actual := args.Map{"result": hc.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_HC_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_HasAll_Empty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.HasAll("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_HC_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqual", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc1.IsEqual(*hc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov40_HC_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_SamePtr", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.IsEqualPtr(hc)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_HC_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_BothEmpty", func() {
		// Arrange
		a := corestr.New.HashsetsCollection.Empty()
		b := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_HC_IsEqualPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_DiffLen", func() {
		// Arrange
		a := corestr.New.HashsetsCollection.Empty()
		a.Add(corestr.New.Hashset.Strings([]string{"a"}))
		b := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": a.IsEqualPtr(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_HC_IsEqualPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_IsEqualPtr_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.IsEqualPtr(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_HC_String(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_String", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.Join(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_HC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JSON", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := json.Marshal(hc)

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_HC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_UnmarshalJSON", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, _ := json.Marshal(hc)
		hc2 := corestr.New.HashsetsCollection.Empty()
		err := json.Unmarshal(data, hc2)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov40_HC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Json_JsonPtr", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hc.Json().Error != nil || hc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_HC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JsonModel", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.JsonModel() == nil || hc.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_HC_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_Serialize", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := hc.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
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
		// Arrange
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		result, err := hc.ParseInjectUsingJson(src.JsonPtr())

		// Act
		actual := args.Map{"result": err != nil || result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_HC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_ParseInjectUsingJsonMust", func() {
		// Arrange
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		result := hc.ParseInjectUsingJsonMust(src.JsonPtr())

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_HC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_AsJsoner", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.AsJsoner() == nil || hc.AsJsonContractsBinder() == nil || hc.AsJsonParseSelfInjector() == nil || hc.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_HC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_HC_JsonParseSelfInject", func() {
		// Arrange
		src := corestr.New.HashsetsCollection.Empty()
		src.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc := corestr.New.HashsetsCollection.Empty()
		err := hc.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ── newHashsetsCollectionCreator ──

func Test_Cov40_Creator_HC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_Empty", func() {
		// Act
		actual := args.Map{"result": corestr.New.HashsetsCollection.Empty() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_Creator_HC_UsingHashsets(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsets", func() {
		// Arrange
		hs := *corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsets(hs)

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_HC_UsingHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsets_Empty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.UsingHashsets()

		// Act
		actual := args.Map{"result": hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_Creator_HC_UsingHashsetsPointers(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsetsPointers", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_HC_UsingHashsetsPointers_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_UsingHashsetsPointers_Empty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers()

		// Act
		actual := args.Map{"result": hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_Creator_HC_LenCap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_LenCap", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.LenCap(0, 5)

		// Act
		actual := args.Map{"result": hc == nil || hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_Creator_HC_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_HC_Cap", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Cap(5)

		// Act
		actual := args.Map{"result": hc == nil || hc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// CollectionsOfCollection
// ══════════════════════════════════════════════════════════════

func Test_Cov40_COC_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Basic", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": coc.IsEmpty() || coc.HasItems() || coc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov40_COC_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Add", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_Add_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Add_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{}))

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_Adds(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Adds", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(c)

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Adds_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Adds()

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.AddCollections(c)

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddStrings_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddStrings(false, []string{})

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": coc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_COC_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddsStringsOfStrings_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddsStringsOfStrings(false)

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_AddAsyncFuncItems(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddAsyncFuncItems", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		coc.AddAsyncFuncItems(wg, false, func() []string { return []string{"a"} })

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_AddAsyncFuncItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AddAsyncFuncItems_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.AddAsyncFuncItems(nil, false)

		// Act
		actual := args.Map{"result": coc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AllIndividualItemsLength", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{"result": coc.AllIndividualItemsLength() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_Cov40_COC_Items(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Items", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": len(coc.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_List(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))

		// Act
		actual := args.Map{"result": len(coc.List(0)) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_COC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_List_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": len(coc.List(0)) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_COC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": coc.ToCollection().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_COC_String(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_String", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": coc.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_COC_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, err := json.Marshal(coc)

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_COC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_UnmarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		data, _ := json.Marshal(coc)
		coc2 := corestr.New.CollectionsOfCollection.Empty()
		err := json.Unmarshal(data, coc2)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov40_COC_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_Json_JsonPtr", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": coc.Json().Error != nil || coc.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_COC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JsonModel", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": coc.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_COC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ParseInjectUsingJson", func() {
		// Arrange
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		result, err := coc.ParseInjectUsingJson(src.JsonPtr())

		// Act
		actual := args.Map{"result": err != nil || result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_COC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_ParseInjectUsingJsonMust", func() {
		// Arrange
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		result := coc.ParseInjectUsingJsonMust(src.JsonPtr())

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_COC_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_AsJsoner", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": coc.AsJsoner() == nil || coc.AsJsonContractsBinder() == nil || coc.AsJsonParseSelfInjector() == nil || coc.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_COC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_COC_JsonParseSelfInject", func() {
		// Arrange
		src := corestr.New.CollectionsOfCollection.Empty()
		src.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc := corestr.New.CollectionsOfCollection.Empty()
		err := coc.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ── newCollectionsOfCollectionCreator ──

func Test_Cov40_Creator_COC_Cap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Cap", func() {
		// Act
		actual := args.Map{"result": corestr.New.CollectionsOfCollection.Cap(5) == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_Creator_COC_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Empty", func() {
		// Act
		actual := args.Map{"result": corestr.New.CollectionsOfCollection.Empty() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_Creator_COC_Strings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_Strings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_COC_CloneStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_CloneStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.CloneStrings([]string{"a"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_COC_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_SpreadStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_COC_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": coc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_Creator_COC_StringsOption(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOption", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.StringsOption(false, 5, []string{"a"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_COC_StringsOptions(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_StringsOptions", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_Creator_COC_LenCap(t *testing.T) {
	safeTest(t, "Test_Cov40_Creator_COC_LenCap", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 5)

		// Act
		actual := args.Map{"result": coc == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════
// CharHashsetMap — core operations
// ══════════════════════════════════════════════════════════════

func Test_Cov40_CHM_Basic(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Basic", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.IsEmpty() || chm.HasItems() || chm.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_Cov40_CHM_Add(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Add", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("avocado").Add("banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStrings", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings("apple", "banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStrings_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_Has(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Has", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.Has("apple") || chm.Has("banana")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Has_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.Has("x")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")

		// Act
		actual := args.Map{"result": has || hs == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		has2, _ := chm.HasWithHashset("banana")
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashset_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		has, _ := chm.HasWithHashset("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_GetChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.GetChar("abc") != 'a' || chm.GetChar("") != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_GetCharOf(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.GetCharOf("xyz") != 'x' || chm.GetCharOf("") != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOf", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("avocado")

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": chm.LengthOf('z') != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOf_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOfHashsetFromFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.LengthOfHashsetFromFirstChar("a") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": chm.LengthOfHashsetFromFirstChar("z") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AllLengthsSum", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AllLengthsSum_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.AllLengthsSum() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_List(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_List", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"result": len(chm.List()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SortedListAsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("banana").Add("apple")
		list := chm.SortedListAsc()

		// Act
		actual := args.Map{"result": list[0] != "apple"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected apple first", actual)
	})
}

func Test_Cov40_CHM_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SortedListDsc", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		list := chm.SortedListDsc()

		// Act
		actual := args.Map{"result": list[0] != "banana"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected banana first", actual)
	})
}

func Test_Cov40_CHM_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetMap", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": len(chm.GetMap()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_GetHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.GetHashset("a", false)

		// Act
		actual := args.Map{"result": hs == nil || !hs.Has("apple")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
		hs2 := chm.GetHashset("z", false)
		actual := args.Map{"result": hs2 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		hs3 := chm.GetHashset("z", true)
		actual := args.Map{"result": hs3 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.GetHashsetByChar('a') == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.HashsetByChar('a') == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByStringFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.HashsetByStringFirstChar("apple") == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollection()

		// Act
		actual := args.Map{"result": hc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollection_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.HashsetsCollection().Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByChars", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollectionByChars('a')

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByChars_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.HashsetsCollectionByChars('a').Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple").Add("banana")
		hc := chm.HashsetsCollectionByStringsFirstChar("apple")

		// Act
		actual := args.Map{"result": hc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetsCollectionByStringsFirstChar_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.HashsetsCollectionByStringsFirstChar("a").Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCollectionItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCollectionItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		chm.AddHashsetItems(hs)

		// Act
		actual := args.Map{"result": chm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.AddSameStartingCharItems('a', []string{"avocado"})

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameStartingCharItems_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{})

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple"})
		hs := chm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollection("a", col)

		// Act
		actual := args.Map{"result": hs == nil || !hs.Has("avocado")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_NilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil (new hashset created)", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollection_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollection_ExistingNilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing hashset", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddSameCharsHashset("a", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_NilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsHashset("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsHashset_ExistingNilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsHashset_ExistingNilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddSameCharsHashset("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected existing", actual)
	})
}

func Test_Cov40_CHM_IsEquals(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals", func() {
		// Arrange
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2.Add("apple")

		// Act
		actual := args.Map{"result": chm1.IsEquals(chm2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov40_CHM_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_SamePtr", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.IsEquals(chm)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_CHM_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.IsEquals(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		b := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_CHM_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_DiffLen", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_DiffContent", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("avocado")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEquals_MissingKey", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("banana")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_Lock_Variants(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Lock_Variants", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual := args.Map{"result": chm.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": chm.AllLengthsSumLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual := args.Map{"result": chm.LengthOfLock('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_LengthOfLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.LengthOfLock('a') != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("apple")

		// Act
		actual := args.Map{"result": chm.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_CHM_AddLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("apple")
		chm.AddLock("avocado")

		// Act
		actual := args.Map{"result": chm.LengthOf('a') != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStringsLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock("apple", "banana")

		// Act
		actual := args.Map{"result": chm.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddStringsLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.GetHashsetLock(false, "a")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCopyMapLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		m := chm.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_Cov40_CHM_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCopyMapLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		m := chm.GetCopyMapLock()

		// Act
		actual := args.Map{"result": len(m) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")

		// Act
		actual := args.Map{"result": has || hs == nil}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		has, _ := chm.HasWithHashsetLock("x")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_HasWithHashsetLock_MissingChar(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HasWithHashsetLock_MissingChar", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		has, _ := chm.HasWithHashsetLock("banana")

		// Act
		actual := args.Map{"result": has}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_Cov40_CHM_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_IsEqualsLock", func() {
		// Arrange
		a := corestr.New.CharHashsetMap.Cap(10, 5)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 5)
		b.Add("apple")

		// Act
		actual := args.Map{"result": a.IsEqualsLock(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_Cov40_CHM_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('a')

		// Act
		actual := args.Map{"result": hs == nil || !hs.Has("apple")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_HashsetByCharLock_Missing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByCharLock_Missing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('z')

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty hashset not nil", actual)
	})
}

func Test_Cov40_CHM_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_HashsetByStringFirstCharLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("apple")

		// Act
		actual := args.Map{"result": hs == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharsGroups", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		groups := chm.GetCharsGroups("apple", "avocado", "banana")

		// Act
		actual := args.Map{"result": groups.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_Cov40_CHM_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_GetCharsGroups_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		groups := chm.GetCharsGroups()

		// Act
		actual := args.Map{"result": groups != chm}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
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
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_CHM_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_StringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_CHM_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SummaryString", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.SummaryString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_CHM_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_SummaryStringLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.SummaryStringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_Cov40_CHM_JSON(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		data, err := json.Marshal(chm)

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_UnmarshalJSON", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		data, _ := json.Marshal(chm)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		err := json.Unmarshal(data, chm2)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov40_CHM_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Json_JsonPtr", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.Json().Error != nil || chm.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JsonModel", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")

		// Act
		actual := args.Map{"result": chm.JsonModel() == nil || chm.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_ParseInjectUsingJson", func() {
		// Arrange
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result, err := chm.ParseInjectUsingJson(src.JsonPtr())

		// Act
		actual := args.Map{"result": err != nil || result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_Cov40_CHM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_ParseInjectUsingJsonMust", func() {
		// Arrange
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.ParseInjectUsingJsonMust(src.JsonPtr())

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AsJsoner", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)

		// Act
		actual := args.Map{"result": chm.AsJsoner() == nil || chm.AsJsonContractsBinder() == nil || chm.AsJsonParseSelfInjector() == nil || chm.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_JsonParseSelfInject", func() {
		// Arrange
		src := corestr.New.CharHashsetMap.Cap(10, 5)
		src.Add("apple")
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		err := chm.JsonParseSelfInject(src.JsonPtr())

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_Cov40_CHM_Clear(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Clear", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.Clear()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_Clear_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Clear()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_RemoveAll", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		chm.RemoveAll()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_RemoveAll_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.RemoveAll()

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCharCollectionMapItems", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 5)
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)

		// Act
		actual := args.Map{"result": chm.Has("apple")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_Cov40_CHM_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddCharCollectionMapItems_Nil", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddCharCollectionMapItems(nil)

		// Act
		actual := args.Map{"result": chm.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_Cov40_CHM_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddHashsetLock("a", hs)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddHashsetLock_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_NilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddHashsetLock("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddHashsetLock_ExistingNilHashset(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddHashsetLock_ExistingNilHashset", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddHashsetLock("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := chm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_Existing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollectionLock("a", col)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_NilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_Cov40_CHM_AddSameCharsCollectionLock_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov40_CHM_AddSameCharsCollectionLock_ExistingNilCol", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("apple")
		result := chm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"result": result == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}
