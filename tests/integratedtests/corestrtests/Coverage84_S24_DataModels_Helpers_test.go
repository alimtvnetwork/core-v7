package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// region HashmapDataModel

func Test_CovS24_01_NewHashmapUsingDataModel(t *testing.T) {
	safeTest(t, "Test_CovS24_01_NewHashmapUsingDataModel", func() {
		// Arrange
		model := &corestr.HashmapDataModel{
			Items: map[string]string{"key": "val"},
		}

		// Act
		hm := corestr.NewHashmapUsingDataModel(model)

		// Assert
		if hm == nil || hm.Length() != 1 {
			t.Errorf("NewHashmapUsingDataModel should create hashmap with 1 item")
		}
	})
}

func Test_CovS24_02_NewHashmapsDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_02_NewHashmapsDataModelUsing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		model := corestr.NewHashmapsDataModelUsing(hm)

		// Assert
		if model == nil || len(model.Items) != 2 {
			t.Errorf("NewHashmapsDataModelUsing should produce model with 2 items")
		}
	})
}

// endregion

// region HashsetDataModel

func Test_CovS24_03_NewHashsetUsingDataModel(t *testing.T) {
	safeTest(t, "Test_CovS24_03_NewHashsetUsingDataModel", func() {
		// Arrange
		model := &corestr.HashsetDataModel{
			Items: map[string]bool{"a": true, "b": true},
		}

		// Act
		hs := corestr.NewHashsetUsingDataModel(model)

		// Assert
		if hs == nil || hs.Length() != 2 {
			t.Errorf("NewHashsetUsingDataModel should create hashset with 2 items")
		}
	})
}

func Test_CovS24_04_NewHashsetsDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_04_NewHashsetsDataModelUsing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		model := corestr.NewHashsetsDataModelUsing(hs)

		// Assert
		if model == nil || len(model.Items) != 2 {
			t.Errorf("NewHashsetsDataModelUsing should produce model with 2 items")
		}
	})
}

// endregion

// region HashsetsCollectionDataModel

func Test_CovS24_05_NewHashsetsCollectionUsingDataModel(t *testing.T) {
	safeTest(t, "Test_CovS24_05_NewHashsetsCollectionUsingDataModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		model := &corestr.HashsetsCollectionDataModel{
			Items: []*corestr.Hashset{hs},
		}

		// Act
		hc := corestr.NewHashsetsCollectionUsingDataModel(model)

		// Assert
		if hc == nil || hc.Length() != 1 {
			t.Errorf("NewHashsetsCollectionUsingDataModel should create collection with 1 item")
		}
	})
}

func Test_CovS24_06_NewHashsetsCollectionDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_06_NewHashsetsCollectionDataModelUsing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs)

		// Act
		model := corestr.NewHashsetsCollectionDataModelUsing(hc)

		// Assert
		if model == nil || len(model.Items) != 1 {
			t.Errorf("NewHashsetsCollectionDataModelUsing should produce model with 1 item")
		}
	})
}

// endregion

// region CharCollectionDataModel

func Test_CovS24_07_NewCharCollectionMapUsingDataModel(t *testing.T) {
	safeTest(t, "Test_CovS24_07_NewCharCollectionMapUsingDataModel", func() {
		// Arrange
		coll := corestr.New.Collection.Strings([]string{"abc"})
		model := &corestr.CharCollectionDataModel{
			Items:                  map[byte]*corestr.Collection{'a': coll},
			EachCollectionCapacity: 10,
		}

		// Act
		ccm := corestr.NewCharCollectionMapUsingDataModel(model)

		// Assert
		if ccm == nil || ccm.Length() != 1 {
			t.Errorf("NewCharCollectionMapUsingDataModel should create map with 1 entry")
		}
	})
}

func Test_CovS24_08_NewCharCollectionMapDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_08_NewCharCollectionMapDataModelUsing", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()
		ccm.Add("alpha")

		// Act
		model := corestr.NewCharCollectionMapDataModelUsing(ccm)

		// Assert
		if model == nil {
			t.Errorf("NewCharCollectionMapDataModelUsing should produce model")
		}
	})
}

// endregion

// region CharHashsetDataModel

func Test_CovS24_09_NewCharHashsetMapUsingDataModel(t *testing.T) {
	safeTest(t, "Test_CovS24_09_NewCharHashsetMapUsingDataModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"abc"})
		model := &corestr.CharHashsetDataModel{
			Items:               map[byte]*corestr.Hashset{'a': hs},
			EachHashsetCapacity: 10,
		}

		// Act
		chm := corestr.NewCharHashsetMapUsingDataModel(model)

		// Assert
		if chm == nil || chm.Length() != 1 {
			t.Errorf("NewCharHashsetMapUsingDataModel should create map with 1 entry")
		}
	})
}

func Test_CovS24_10_NewCharHashsetMapDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_10_NewCharHashsetMapDataModelUsing", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 0)
		chm.Add("beta")

		// Act
		model := corestr.NewCharHashsetMapDataModelUsing(chm)

		// Assert
		if model == nil {
			t.Errorf("NewCharHashsetMapDataModelUsing should produce model")
		}
	})
}

// endregion

// region AllIndividualStringsOfStringsLength

func Test_CovS24_11_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_CovS24_11_AllIndividualStringsOfStringsLength_Nil", func() {
		// Arrange & Act
		result := corestr.AllIndividualStringsOfStringsLength(nil)

		// Assert
		if result != 0 {
			t.Errorf("Expected 0 for nil, got %d", result)
		}
	})
}

func Test_CovS24_12_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	safeTest(t, "Test_CovS24_12_AllIndividualStringsOfStringsLength_Empty", func() {
		// Arrange
		items := [][]string{}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		if result != 0 {
			t.Errorf("Expected 0 for empty, got %d", result)
		}
	})
}

func Test_CovS24_13_AllIndividualStringsOfStringsLength_Multiple(t *testing.T) {
	safeTest(t, "Test_CovS24_13_AllIndividualStringsOfStringsLength_Multiple", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}, {"d", "e", "f"}}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		if result != 6 {
			t.Errorf("Expected 6, got %d", result)
		}
	})
}

// endregion

// region AllIndividualsLengthOfSimpleSlices

func Test_CovS24_14_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_CovS24_14_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Arrange & Act
		result := corestr.AllIndividualsLengthOfSimpleSlices()

		// Assert
		if result != 0 {
			t.Errorf("Expected 0 for nil, got %d", result)
		}
	})
}

func Test_CovS24_15_AllIndividualsLengthOfSimpleSlices_Multiple(t *testing.T) {
	safeTest(t, "Test_CovS24_15_AllIndividualsLengthOfSimpleSlices_Multiple", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss2 := corestr.New.SimpleSlice.Strings([]string{"c"})

		// Act
		result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)

		// Assert
		if result != 3 {
			t.Errorf("Expected 3, got %d", result)
		}
	})
}

// endregion

// region AnyToString (corestr)

func Test_CovS24_16_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_CovS24_16_AnyToString_EmptyString", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "")

		// Assert
		if result != "" {
			t.Errorf("AnyToString empty string should return empty, got '%s'", result)
		}
	})
}

func Test_CovS24_17_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_CovS24_17_AnyToString_WithFieldName", func() {
		// Arrange & Act
		result := corestr.AnyToString(true, "hello")

		// Assert
		if result == "" {
			t.Errorf("AnyToString with field name should return non-empty")
		}
	})
}

func Test_CovS24_18_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_CovS24_18_AnyToString_WithoutFieldName", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		if result == "" {
			t.Errorf("AnyToString without field name should return non-empty")
		}
	})
}

func Test_CovS24_19_AnyToString_Pointer(t *testing.T) {
	safeTest(t, "Test_CovS24_19_AnyToString_Pointer", func() {
		// Arrange
		val := "test"

		// Act
		result := corestr.AnyToString(false, &val)

		// Assert
		if result == "" {
			t.Errorf("AnyToString with pointer should return non-empty")
		}
	})
}

func Test_CovS24_20_AnyToString_NilInterface(t *testing.T) {
	safeTest(t, "Test_CovS24_20_AnyToString_NilInterface", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, nil)

		// Assert
		// nil will not match "" in the function, it goes through reflectInterfaceVal
		// reflectInterfaceVal returns nil for nil
		_ = result
	})
}

// endregion

// region reflectInterfaceVal (tested indirectly via AnyToString)

func Test_CovS24_21_AnyToString_Struct(t *testing.T) {
	safeTest(t, "Test_CovS24_21_AnyToString_Struct", func() {
		// Arrange
		type sample struct{ Name string }
		s := sample{Name: "test"}

		// Act
		result := corestr.AnyToString(true, s)

		// Assert
		if result == "" {
			t.Errorf("AnyToString with struct should return non-empty")
		}
	})
}

func Test_CovS24_22_AnyToString_StructPointer(t *testing.T) {
	safeTest(t, "Test_CovS24_22_AnyToString_StructPointer", func() {
		// Arrange
		type sample struct{ Name string }
		s := &sample{Name: "test"}

		// Act
		result := corestr.AnyToString(false, s)

		// Assert
		if result == "" {
			t.Errorf("AnyToString with struct pointer should return non-empty")
		}
	})
}

// endregion

// region CollectionsOfCollectionModel (just a data struct)

func Test_CovS24_23_CollectionsOfCollectionModel_Fields(t *testing.T) {
	safeTest(t, "Test_CovS24_23_CollectionsOfCollectionModel_Fields", func() {
		// Arrange
		coll := corestr.New.Collection.Strings([]string{"a"})
		model := corestr.CollectionsOfCollectionModel{
			Items: []*corestr.Collection{coll},
		}

		// Assert
		if len(model.Items) != 1 {
			t.Errorf("Model should have 1 item, got %d", len(model.Items))
		}
	})
}

// endregion

// region SimpleStringOnceModel (just a data struct)

func Test_CovS24_24_SimpleStringOnceModel_Fields(t *testing.T) {
	safeTest(t, "Test_CovS24_24_SimpleStringOnceModel_Fields", func() {
		// Arrange
		model := corestr.SimpleStringOnceModel{
			Value:        "hello",
			IsInitialize: true,
		}

		// Assert
		if model.Value != "hello" || !model.IsInitialize {
			t.Errorf("Model fields mismatch")
		}
	})
}

// endregion
