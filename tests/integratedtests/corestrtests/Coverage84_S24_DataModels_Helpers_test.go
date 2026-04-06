package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": hm == nil || hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashmapUsingDataModel should create hashmap with 1 item", actual)
	})
}

func Test_CovS24_02_NewHashmapsDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_02_NewHashmapsDataModelUsing", func() {
		// Arrange
		hm := corestr.New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})

		// Act
		model := corestr.NewHashmapsDataModelUsing(hm)

		// Assert
		actual := args.Map{"result": model == nil || len(model.Items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashmapsDataModelUsing should produce model with 2 items", actual)
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
		actual := args.Map{"result": hs == nil || hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashsetUsingDataModel should create hashset with 2 items", actual)
	})
}

func Test_CovS24_04_NewHashsetsDataModelUsing(t *testing.T) {
	safeTest(t, "Test_CovS24_04_NewHashsetsDataModelUsing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		model := corestr.NewHashsetsDataModelUsing(hs)

		// Assert
		actual := args.Map{"result": model == nil || len(model.Items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashsetsDataModelUsing should produce model with 2 items", actual)
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
		actual := args.Map{"result": hc == nil || hc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionUsingDataModel should create collection with 1 item", actual)
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
		actual := args.Map{"result": model == nil || len(model.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionDataModelUsing should produce model with 1 item", actual)
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
		actual := args.Map{"result": ccm == nil || ccm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewCharCollectionMapUsingDataModel should create map with 1 entry", actual)
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
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewCharCollectionMapDataModelUsing should produce model", actual)
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
		actual := args.Map{"result": chm == nil || chm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewCharHashsetMapUsingDataModel should create map with 1 entry", actual)
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
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewCharHashsetMapDataModelUsing should produce model", actual)
	})
}

// endregion

// region AllIndividualStringsOfStringsLength

func Test_CovS24_11_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_CovS24_11_AllIndividualStringsOfStringsLength_Nil", func() {
		// Arrange & Act
		result := corestr.AllIndividualStringsOfStringsLength(nil)

		// Assert
		actual := args.Map{"result": result != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 0 for nil", actual)
	})
}

func Test_CovS24_12_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	safeTest(t, "Test_CovS24_12_AllIndividualStringsOfStringsLength_Empty", func() {
		// Arrange
		items := [][]string{}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		actual := args.Map{"result": result != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 0 for empty", actual)
	})
}

func Test_CovS24_13_AllIndividualStringsOfStringsLength_Multiple(t *testing.T) {
	safeTest(t, "Test_CovS24_13_AllIndividualStringsOfStringsLength_Multiple", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}, {"d", "e", "f"}}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		actual := args.Map{"result": result != 6}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 6", actual)
	})
}

// endregion

// region AllIndividualsLengthOfSimpleSlices

func Test_CovS24_14_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_CovS24_14_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Arrange & Act
		result := corestr.AllIndividualsLengthOfSimpleSlices()

		// Assert
		actual := args.Map{"result": result != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 0 for nil", actual)
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
		actual := args.Map{"result": result != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Expected 3", actual)
	})
}

// endregion

// region AnyToString (corestr)

func Test_CovS24_16_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_CovS24_16_AnyToString_EmptyString", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "")

		// Assert
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString empty string should return empty, got ''", actual)
	})
}

func Test_CovS24_17_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_CovS24_17_AnyToString_WithFieldName", func() {
		// Arrange & Act
		result := corestr.AnyToString(true, "hello")

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString with field name should return non-empty", actual)
	})
}

func Test_CovS24_18_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_CovS24_18_AnyToString_WithoutFieldName", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString without field name should return non-empty", actual)
	})
}

func Test_CovS24_19_AnyToString_Pointer(t *testing.T) {
	safeTest(t, "Test_CovS24_19_AnyToString_Pointer", func() {
		// Arrange
		val := "test"

		// Act
		result := corestr.AnyToString(false, &val)

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString with pointer should return non-empty", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString with struct should return non-empty", actual)
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
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "AnyToString with struct pointer should return non-empty", actual)
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
		actual := args.Map{"result": len(model.Items) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Model should have 1 item", actual)
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
		actual := args.Map{"result": model.Value != "hello" || !model.IsInitialize}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Model fields mismatch", actual)
	})
}

// endregion
