package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// region HashsetsCollection Core

func Test_CovS23_01_HashsetsCollection_IsEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_01_HashsetsCollection_IsEmpty_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.IsEmpty()

		// Assert
		if !result {
			t.Errorf("IsEmpty on empty collection should be true")
		}
	})
}

func Test_CovS23_02_HashsetsCollection_HasItems_NonEmpty(t *testing.T) {
	safeTest(t, "Test_CovS23_02_HashsetsCollection_HasItems_NonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hc := corestr.New.HashsetsCollection.UsingHashsets(*hs)

		// Act
		result := hc.HasItems()

		// Assert
		if !result {
			t.Errorf("HasItems should be true for non-empty collection")
		}
	})
}

func Test_CovS23_03_HashsetsCollection_HasItems_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_03_HashsetsCollection_HasItems_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.HasItems()

		// Assert
		if result {
			t.Errorf("HasItems should be false for empty collection")
		}
	})
}

func Test_CovS23_04_HashsetsCollection_Length(t *testing.T) {
	safeTest(t, "Test_CovS23_04_HashsetsCollection_Length", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1, hs2)

		// Act
		length := hc.Length()

		// Assert
		if length != 2 {
			t.Errorf("Length expected 2, got %d", length)
		}
	})
}

func Test_CovS23_05_HashsetsCollection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_05_HashsetsCollection_Length_Nil", func() {
		// Arrange
		var hc *corestr.HashsetsCollection

		// Act
		length := hc.Length()

		// Assert
		if length != 0 {
			t.Errorf("Length on nil should be 0, got %d", length)
		}
	})
}

func Test_CovS23_06_HashsetsCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_CovS23_06_HashsetsCollection_LastIndex", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"x"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		lastIdx := hc.LastIndex()

		// Assert
		if lastIdx != 0 {
			t.Errorf("LastIndex expected 0, got %d", lastIdx)
		}
	})
}

func Test_CovS23_07_HashsetsCollection_LastIndex_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_07_HashsetsCollection_LastIndex_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		lastIdx := hc.LastIndex()

		// Assert
		if lastIdx != -1 {
			t.Errorf("LastIndex on empty expected -1, got %d", lastIdx)
		}
	})
}

// endregion

// region HashsetsCollection Add/Adds

func Test_CovS23_08_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_CovS23_08_HashsetsCollection_Add", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		hc.Add(hs)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("Add should result in length 1, got %d", hc.Length())
		}
	})
}

func Test_CovS23_09_HashsetsCollection_AddNonNil_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_09_HashsetsCollection_AddNonNil_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.AddNonNil(nil)

		// Assert
		if hc.Length() != 0 {
			t.Errorf("AddNonNil with nil should not add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_10_HashsetsCollection_AddNonNil_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_10_HashsetsCollection_AddNonNil_Valid", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"x"})

		// Act
		hc.AddNonNil(hs)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("AddNonNil with valid should add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_11_HashsetsCollection_AddNonEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_11_HashsetsCollection_AddNonEmpty_Empty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.Empty.Hashset()

		// Act
		hc.AddNonEmpty(hs)

		// Assert
		if hc.Length() != 0 {
			t.Errorf("AddNonEmpty with empty hashset should not add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_12_HashsetsCollection_AddNonEmpty_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_12_HashsetsCollection_AddNonEmpty_Valid", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hc.AddNonEmpty(hs)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("AddNonEmpty with valid should add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_13_HashsetsCollection_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_13_HashsetsCollection_Adds_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.Adds(nil)

		// Assert
		if hc.Length() != 0 {
			t.Errorf("Adds with nil should not add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_14_HashsetsCollection_Adds_SkipsEmpty(t *testing.T) {
	safeTest(t, "Test_CovS23_14_HashsetsCollection_Adds_SkipsEmpty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hsEmpty := corestr.Empty.Hashset()
		hsValid := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hc.Adds(hsEmpty, hsValid)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("Adds should skip empty hashsets, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_15_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_15_HashsetsCollection_AddHashsetsCollection_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.AddHashsetsCollection(nil)

		// Assert
		if hc.Length() != 0 {
			t.Errorf("AddHashsetsCollection nil should not add, got length %d", hc.Length())
		}
	})
}

func Test_CovS23_16_HashsetsCollection_AddHashsetsCollection_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_16_HashsetsCollection_AddHashsetsCollection_Valid", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))

		// Act
		hc1.AddHashsetsCollection(hc2)

		// Assert
		if hc1.Length() != 2 {
			t.Errorf("AddHashsetsCollection should merge, got length %d", hc1.Length())
		}
	})
}

// endregion

// region HashsetsCollection ConcatNew

func Test_CovS23_17_HashsetsCollection_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_CovS23_17_HashsetsCollection_ConcatNew_NoArgs", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		result := hc.ConcatNew()

		// Assert
		if result.Length() != 1 {
			t.Errorf("ConcatNew no args should clone, got length %d", result.Length())
		}
	})
}

func Test_CovS23_18_HashsetsCollection_ConcatNew_WithCollections(t *testing.T) {
	safeTest(t, "Test_CovS23_18_HashsetsCollection_ConcatNew_WithCollections", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))

		// Act
		result := hc1.ConcatNew(hc2)

		// Assert
		if result.Length() != 2 {
			t.Errorf("ConcatNew should merge, got length %d", result.Length())
		}
	})
}

// endregion

// region HashsetsCollection List/StringsList

func Test_CovS23_19_HashsetsCollection_List(t *testing.T) {
	safeTest(t, "Test_CovS23_19_HashsetsCollection_List", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		list := hc.List()

		// Assert
		if len(list) != 1 {
			t.Errorf("List expected 1 item, got %d", len(list))
		}
	})
}

func Test_CovS23_20_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_CovS23_20_HashsetsCollection_ListPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		listPtr := hc.ListPtr()

		// Assert
		if listPtr == nil || len(*listPtr) != 1 {
			t.Errorf("ListPtr should return pointer to 1-item slice")
		}
	})
}

func Test_CovS23_21_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_CovS23_21_HashsetsCollection_ListDirectPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		listPtr := hc.ListDirectPtr()

		// Assert
		if listPtr == nil || len(*listPtr) != 1 {
			t.Errorf("ListDirectPtr should return pointer to 1-item slice")
		}
	})
}

func Test_CovS23_22_HashsetsCollection_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_22_HashsetsCollection_StringsList_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		list := hc.StringsList()

		// Assert
		if len(list) != 0 {
			t.Errorf("StringsList on empty expected 0, got %d", len(list))
		}
	})
}

func Test_CovS23_23_HashsetsCollection_StringsList_NonEmpty(t *testing.T) {
	safeTest(t, "Test_CovS23_23_HashsetsCollection_StringsList_NonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"x", "y"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		list := hc.StringsList()

		// Assert
		if len(list) != 2 {
			t.Errorf("StringsList expected 2, got %d", len(list))
		}
	})
}

// endregion

// region HashsetsCollection IndexOf

func Test_CovS23_24_HashsetsCollection_IndexOf_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_24_HashsetsCollection_IndexOf_Valid", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.IndexOf(0)

		// Assert
		if result == nil {
			t.Errorf("IndexOf(0) should return valid hashset")
		}
	})
}

func Test_CovS23_25_HashsetsCollection_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_25_HashsetsCollection_IndexOf_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.IndexOf(0)

		// Assert
		if result != nil {
			t.Errorf("IndexOf on empty should return nil")
		}
	})
}

// endregion

// region HashsetsCollection HasAll

func Test_CovS23_26_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_26_HashsetsCollection_HasAll_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.HasAll("a")

		// Assert
		if result {
			t.Errorf("HasAll on empty should be false")
		}
	})
}

func Test_CovS23_27_HashsetsCollection_HasAll_Found(t *testing.T) {
	safeTest(t, "Test_CovS23_27_HashsetsCollection_HasAll_Found", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.HasAll("a", "b")

		// Assert
		if !result {
			t.Errorf("HasAll should find all items")
		}
	})
}

func Test_CovS23_28_HashsetsCollection_HasAll_NotFound(t *testing.T) {
	safeTest(t, "Test_CovS23_28_HashsetsCollection_HasAll_NotFound", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.HasAll("a", "z")

		// Assert
		if result {
			t.Errorf("HasAll should be false when not all items present")
		}
	})
}

// endregion

// region HashsetsCollection IsEqual

func Test_CovS23_29_HashsetsCollection_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_CovS23_29_HashsetsCollection_IsEqual_BothEmpty", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc2 := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc1.IsEqual(*hc2)

		// Assert
		if !result {
			t.Errorf("IsEqual should be true for two empty collections")
		}
	})
}

func Test_CovS23_30_HashsetsCollection_IsEqualPtr_SamePtr(t *testing.T) {
	safeTest(t, "Test_CovS23_30_HashsetsCollection_IsEqualPtr_SamePtr", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		result := hc.IsEqualPtr(hc)

		// Assert
		if !result {
			t.Errorf("IsEqualPtr same pointer should be true")
		}
	})
}

func Test_CovS23_31_HashsetsCollection_IsEqualPtr_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_31_HashsetsCollection_IsEqualPtr_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.IsEqualPtr(nil)

		// Assert
		if result {
			t.Errorf("IsEqualPtr with nil should be false")
		}
	})
}

func Test_CovS23_32_HashsetsCollection_IsEqualPtr_DifferentLength(t *testing.T) {
	safeTest(t, "Test_CovS23_32_HashsetsCollection_IsEqualPtr_DifferentLength", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc1.IsEqualPtr(hc2)

		// Assert
		if result {
			t.Errorf("IsEqualPtr different length should be false")
		}
	})
}

func Test_CovS23_33_HashsetsCollection_IsEqualPtr_DifferentContent(t *testing.T) {
	safeTest(t, "Test_CovS23_33_HashsetsCollection_IsEqualPtr_DifferentContent", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"b"}))

		// Act
		result := hc1.IsEqualPtr(hc2)

		// Assert
		if result {
			t.Errorf("IsEqualPtr different content should be false")
		}
	})
}

func Test_CovS23_34_HashsetsCollection_IsEqualPtr_SameContent(t *testing.T) {
	safeTest(t, "Test_CovS23_34_HashsetsCollection_IsEqualPtr_SameContent", func() {
		// Arrange
		hc1 := corestr.New.HashsetsCollection.Empty()
		hc1.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))
		hc2 := corestr.New.HashsetsCollection.Empty()
		hc2.Add(corestr.New.Hashset.Strings([]string{"a", "b"}))

		// Act
		result := hc1.IsEqualPtr(hc2)

		// Assert
		if !result {
			t.Errorf("IsEqualPtr same content should be true")
		}
	})
}

// endregion

// region HashsetsCollection String/Join

func Test_CovS23_35_HashsetsCollection_String_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_35_HashsetsCollection_String_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.String()

		// Assert
		if result == "" {
			t.Errorf("String on empty should return NoElements indicator")
		}
	})
}

func Test_CovS23_36_HashsetsCollection_String_NonEmpty(t *testing.T) {
	safeTest(t, "Test_CovS23_36_HashsetsCollection_String_NonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.String()

		// Assert
		if result == "" {
			t.Errorf("String on non-empty should return content")
		}
	})
}

func Test_CovS23_37_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_CovS23_37_HashsetsCollection_Join", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.Join(",")

		// Assert
		if result != "a" {
			t.Errorf("Join expected 'a', got '%s'", result)
		}
	})
}

// endregion

// region HashsetsCollection JSON

func Test_CovS23_38_HashsetsCollection_Json(t *testing.T) {
	safeTest(t, "Test_CovS23_38_HashsetsCollection_Json", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		jsonResult := hc.Json()

		// Assert
		if jsonResult.HasError() {
			t.Errorf("Json should produce valid result, got error: %v", jsonResult.Error)
		}
	})
}

func Test_CovS23_39_HashsetsCollection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovS23_39_HashsetsCollection_JsonPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		jsonResult := hc.JsonPtr()

		// Assert
		if jsonResult == nil {
			t.Errorf("JsonPtr should not be nil")
		}
	})
}

func Test_CovS23_40_HashsetsCollection_JsonModel(t *testing.T) {
	safeTest(t, "Test_CovS23_40_HashsetsCollection_JsonModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		model := hc.JsonModel()

		// Assert
		if model == nil {
			t.Errorf("JsonModel should not be nil")
		}
	})
}

func Test_CovS23_41_HashsetsCollection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovS23_41_HashsetsCollection_JsonModelAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		result := hc.JsonModelAny()

		// Assert
		if result == nil {
			t.Errorf("JsonModelAny should not be nil")
		}
	})
}

func Test_CovS23_42_HashsetsCollection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovS23_42_HashsetsCollection_MarshalUnmarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		data, err := hc.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		hc2 := corestr.New.HashsetsCollection.Empty()
		err2 := hc2.UnmarshalJSON(data)

		// Assert
		if err2 != nil {
			t.Errorf("UnmarshalJSON error: %v", err2)
		}
	})
}

func Test_CovS23_43_HashsetsCollection_UnmarshalJSON_InvalidData(t *testing.T) {
	safeTest(t, "Test_CovS23_43_HashsetsCollection_UnmarshalJSON_InvalidData", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		err := hc.UnmarshalJSON([]byte("invalid"))

		// Assert
		if err == nil {
			t.Errorf("UnmarshalJSON with invalid data should return error")
		}
	})
}

func Test_CovS23_44_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_CovS23_44_HashsetsCollection_Serialize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		data, err := hc.Serialize()

		// Assert
		if err != nil || len(data) == 0 {
			t.Errorf("Serialize should produce bytes, err: %v", err)
		}
	})
}

func Test_CovS23_45_HashsetsCollection_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovS23_45_HashsetsCollection_Deserialize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		var target map[string]interface{}
		err := hc.Deserialize(&target)

		// Assert
		if err != nil {
			t.Errorf("Deserialize error: %v", err)
		}
	})
}

func Test_CovS23_46_HashsetsCollection_ParseInjectUsingJson_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_46_HashsetsCollection_ParseInjectUsingJson_Valid", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		jsonResult := hc.JsonPtr()

		// Act
		hc2 := corestr.New.HashsetsCollection.Empty()
		result, err := hc2.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil || result == nil {
			t.Errorf("ParseInjectUsingJson should succeed, err: %v", err)
		}
	})
}

func Test_CovS23_47_HashsetsCollection_ParseInjectUsingJson_Invalid(t *testing.T) {
	safeTest(t, "Test_CovS23_47_HashsetsCollection_ParseInjectUsingJson_Invalid", func() {
		// Arrange
		jsonResult := corejson.NewPtr("not a hashsets collection")

		// Act
		hc := corestr.New.HashsetsCollection.Empty()
		_, err := hc.ParseInjectUsingJson(jsonResult)

		// Assert
		if err == nil {
			t.Errorf("ParseInjectUsingJson with invalid JSON should error")
		}
	})
}

func Test_CovS23_48_HashsetsCollection_ParseInjectUsingJsonMust_Valid(t *testing.T) {
	safeTest(t, "Test_CovS23_48_HashsetsCollection_ParseInjectUsingJsonMust_Valid", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		jsonResult := hc.JsonPtr()

		// Act
		hc2 := corestr.New.HashsetsCollection.Empty()
		result := hc2.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result == nil {
			t.Errorf("ParseInjectUsingJsonMust should succeed")
		}
	})
}

func Test_CovS23_49_HashsetsCollection_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_CovS23_49_HashsetsCollection_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		jsonResult := corejson.NewPtr("bad data")

		// Act & Assert
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("ParseInjectUsingJsonMust should panic on invalid data")
			}
		}()
		hc := corestr.New.HashsetsCollection.Empty()
		hc.ParseInjectUsingJsonMust(jsonResult)
	})
}

func Test_CovS23_50_HashsetsCollection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovS23_50_HashsetsCollection_JsonParseSelfInject", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)
		jsonResult := hc.JsonPtr()

		// Act
		hc2 := corestr.New.HashsetsCollection.Empty()
		err := hc2.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("JsonParseSelfInject should succeed, err: %v", err)
		}
	})
}

// endregion

// region HashsetsCollection Interface Casts

func Test_CovS23_51_HashsetsCollection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_CovS23_51_HashsetsCollection_AsJsonContractsBinder", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsonContractsBinder()

		// Assert
		if result == nil {
			t.Errorf("AsJsonContractsBinder should not be nil")
		}
	})
}

func Test_CovS23_52_HashsetsCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_CovS23_52_HashsetsCollection_AsJsoner", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsoner()

		// Assert
		if result == nil {
			t.Errorf("AsJsoner should not be nil")
		}
	})
}

func Test_CovS23_53_HashsetsCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_CovS23_53_HashsetsCollection_AsJsonParseSelfInjector", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsonParseSelfInjector()

		// Assert
		if result == nil {
			t.Errorf("AsJsonParseSelfInjector should not be nil")
		}
	})
}

func Test_CovS23_54_HashsetsCollection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_CovS23_54_HashsetsCollection_AsJsonMarshaller", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsonMarshaller()

		// Assert
		if result == nil {
			t.Errorf("AsJsonMarshaller should not be nil")
		}
	})
}

// endregion

// region newHashsetsCollectionCreator

func Test_CovS23_55_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_55_Creator_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.Empty()

		// Assert
		if hc == nil || !hc.IsEmpty() {
			t.Errorf("Empty() should create empty collection")
		}
	})
}

func Test_CovS23_56_Creator_UsingHashsets(t *testing.T) {
	safeTest(t, "Test_CovS23_56_Creator_UsingHashsets", func() {
		// Arrange
		hs := *corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		hc := corestr.New.HashsetsCollection.UsingHashsets(hs)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("UsingHashsets expected 1, got %d", hc.Length())
		}
	})
}

func Test_CovS23_57_Creator_UsingHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_57_Creator_UsingHashsets_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.UsingHashsets()

		// Assert
		if !hc.IsEmpty() {
			t.Errorf("UsingHashsets() with no args should be empty")
		}
	})
}

func Test_CovS23_58_Creator_UsingHashsetsPointers(t *testing.T) {
	safeTest(t, "Test_CovS23_58_Creator_UsingHashsetsPointers", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Assert
		if hc.Length() != 1 {
			t.Errorf("UsingHashsetsPointers expected 1, got %d", hc.Length())
		}
	})
}

func Test_CovS23_59_Creator_UsingHashsetsPointers_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_59_Creator_UsingHashsetsPointers_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers()

		// Assert
		if !hc.IsEmpty() {
			t.Errorf("UsingHashsetsPointers() with no args should be empty")
		}
	})
}

func Test_CovS23_60_Creator_LenCap(t *testing.T) {
	safeTest(t, "Test_CovS23_60_Creator_LenCap", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.LenCap(0, 10)

		// Assert
		if hc == nil || hc.Length() != 0 {
			t.Errorf("LenCap should create empty collection with capacity")
		}
	})
}

func Test_CovS23_61_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS23_61_Creator_Cap", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.Cap(5)

		// Assert
		if hc == nil || hc.Length() != 0 {
			t.Errorf("Cap should create empty collection with capacity")
		}
	})
}

// endregion
