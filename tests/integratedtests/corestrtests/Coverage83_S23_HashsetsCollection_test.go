package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// region HashsetsCollection Core

func Test_CovS23_01_HashsetsCollection_IsEmpty_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_01_HashsetsCollection_IsEmpty_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.IsEmpty()

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty on empty collection should be true", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasItems should be true for non-empty collection", actual)
	})
}

func Test_CovS23_03_HashsetsCollection_HasItems_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_03_HashsetsCollection_HasItems_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.HasItems()

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasItems should be false for empty collection", actual)
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
		actual := args.Map{"result": length != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length expected 2", actual)
	})
}

func Test_CovS23_05_HashsetsCollection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_05_HashsetsCollection_Length_Nil", func() {
		// Arrange
		var hc *corestr.HashsetsCollection

		// Act
		length := hc.Length()

		// Assert
		actual := args.Map{"result": length != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Length on nil should be 0", actual)
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
		actual := args.Map{"result": lastIdx != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastIndex expected 0", actual)
	})
}

func Test_CovS23_07_HashsetsCollection_LastIndex_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_07_HashsetsCollection_LastIndex_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		lastIdx := hc.LastIndex()

		// Assert
		actual := args.Map{"result": lastIdx != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LastIndex on empty expected -1", actual)
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
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Add should result in length 1", actual)
	})
}

func Test_CovS23_09_HashsetsCollection_AddNonNil_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_09_HashsetsCollection_AddNonNil_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.AddNonNil(nil)

		// Assert
		actual := args.Map{"result": hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonNil with nil should not add, got length", actual)
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
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonNil with valid should add, got length", actual)
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
		actual := args.Map{"result": hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty with empty hashset should not add, got length", actual)
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
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty with valid should add, got length", actual)
	})
}

func Test_CovS23_13_HashsetsCollection_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_13_HashsetsCollection_Adds_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.Adds(nil)

		// Assert
		actual := args.Map{"result": hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Adds with nil should not add, got length", actual)
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
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Adds should skip empty hashsets, got length", actual)
	})
}

func Test_CovS23_15_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_15_HashsetsCollection_AddHashsetsCollection_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		hc.AddHashsetsCollection(nil)

		// Assert
		actual := args.Map{"result": hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection nil should not add, got length", actual)
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
		actual := args.Map{"result": hc1.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection should merge, got length", actual)
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
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew no args should clone, got length", actual)
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
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ConcatNew should merge, got length", actual)
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
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "List expected 1 item", actual)
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
		actual := args.Map{"result": listPtr == nil || len(*listPtr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ListPtr should return pointer to 1-item slice", actual)
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
		actual := args.Map{"result": listPtr == nil || len(*listPtr) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ListDirectPtr should return pointer to 1-item slice", actual)
	})
}

func Test_CovS23_22_HashsetsCollection_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_22_HashsetsCollection_StringsList_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		list := hc.StringsList()

		// Assert
		actual := args.Map{"result": len(list) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsList on empty expected 0", actual)
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
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsList expected 2", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IndexOf(0) should return valid hashset", actual)
	})
}

func Test_CovS23_25_HashsetsCollection_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_25_HashsetsCollection_IndexOf_Empty", func() {
		// Arrange
		hc := corestr.Empty.HashsetsCollection()

		// Act
		result := hc.IndexOf(0)

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IndexOf on empty should return nil", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll on empty should be false", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasAll should find all items", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasAll should be false when not all items present", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEqual should be true for two empty collections", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr same pointer should be true", actual)
	})
}

func Test_CovS23_31_HashsetsCollection_IsEqualPtr_Nil(t *testing.T) {
	safeTest(t, "Test_CovS23_31_HashsetsCollection_IsEqualPtr_Nil", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.IsEqualPtr(nil)

		// Assert
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr with nil should be false", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr different length should be false", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr different content should be false", actual)
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
		actual := args.Map{"result": result}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr same content should be true", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "String on empty should return NoElements indicator", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "String on non-empty should return content", actual)
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
		actual := args.Map{"result": result != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Join expected 'a', got ''", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Json should produce valid result, got error:", actual)
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
		actual := args.Map{"result": jsonResult == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr should not be nil", actual)
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
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonModel should not be nil", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonModelAny should not be nil", actual)
	})
}

func Test_CovS23_42_HashsetsCollection_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovS23_42_HashsetsCollection_MarshalUnmarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		data, err := hc.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
		hc2 := corestr.New.HashsetsCollection.Empty()
		err2 := hc2.UnmarshalJSON(data)

		// Assert
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
	})
}

func Test_CovS23_43_HashsetsCollection_UnmarshalJSON_InvalidData(t *testing.T) {
	safeTest(t, "Test_CovS23_43_HashsetsCollection_UnmarshalJSON_InvalidData", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		err := hc.UnmarshalJSON([]byte("invalid"))

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON with invalid data should return error", actual)
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
		actual := args.Map{"result": err != nil || len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Serialize should produce bytes, err:", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Deserialize error:", actual)
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
		actual := args.Map{"result": err != nil || result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson should succeed, err:", actual)
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
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson with invalid JSON should error", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust should succeed", actual)
	})
}

func Test_CovS23_49_HashsetsCollection_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_CovS23_49_HashsetsCollection_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		jsonResult := corejson.NewPtr("bad data")

		// Act & Assert
		defer func() {
			r := recover()
			actual := args.Map{"result": r == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust should panic on invalid data", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject should succeed, err:", actual)
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
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder should not be nil", actual)
	})
}

func Test_CovS23_52_HashsetsCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_CovS23_52_HashsetsCollection_AsJsoner", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsoner()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsJsoner should not be nil", actual)
	})
}

func Test_CovS23_53_HashsetsCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_CovS23_53_HashsetsCollection_AsJsonParseSelfInjector", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsonParseSelfInjector()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector should not be nil", actual)
	})
}

func Test_CovS23_54_HashsetsCollection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_CovS23_54_HashsetsCollection_AsJsonMarshaller", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		result := hc.AsJsonMarshaller()

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller should not be nil", actual)
	})
}

// endregion

// region newHashsetsCollectionCreator

func Test_CovS23_55_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_55_Creator_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.Empty()

		// Assert
		actual := args.Map{"result": hc == nil || !hc.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty() should create empty collection", actual)
	})
}

func Test_CovS23_56_Creator_UsingHashsets(t *testing.T) {
	safeTest(t, "Test_CovS23_56_Creator_UsingHashsets", func() {
		// Arrange
		hs := *corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		hc := corestr.New.HashsetsCollection.UsingHashsets(hs)

		// Assert
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingHashsets expected 1", actual)
	})
}

func Test_CovS23_57_Creator_UsingHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_57_Creator_UsingHashsets_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.UsingHashsets()

		// Assert
		actual := args.Map{"result": hc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "UsingHashsets() with no args should be empty", actual)
	})
}

func Test_CovS23_58_Creator_UsingHashsetsPointers(t *testing.T) {
	safeTest(t, "Test_CovS23_58_Creator_UsingHashsetsPointers", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Assert
		actual := args.Map{"result": hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingHashsetsPointers expected 1", actual)
	})
}

func Test_CovS23_59_Creator_UsingHashsetsPointers_Empty(t *testing.T) {
	safeTest(t, "Test_CovS23_59_Creator_UsingHashsetsPointers_Empty", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers()

		// Assert
		actual := args.Map{"result": hc.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "UsingHashsetsPointers() with no args should be empty", actual)
	})
}

func Test_CovS23_60_Creator_LenCap(t *testing.T) {
	safeTest(t, "Test_CovS23_60_Creator_LenCap", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.LenCap(0, 10)

		// Assert
		actual := args.Map{"result": hc == nil || hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LenCap should create empty collection with capacity", actual)
	})
}

func Test_CovS23_61_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS23_61_Creator_Cap", func() {
		// Arrange & Act
		hc := corestr.New.HashsetsCollection.Cap(5)

		// Assert
		actual := args.Map{"result": hc == nil || hc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Cap should create empty collection with capacity", actual)
	})
}

// endregion
