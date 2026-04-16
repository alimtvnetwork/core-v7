package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── CloneSlice ──

func Test_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_CloneSlice_Empty", func() {
		// Arrange
		tc := cloneSliceTestCases[0]

		// Act
		result := corestr.CloneSlice(nil)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": len(result),
		})
	})
}

func Test_CloneSlice_WithItems(t *testing.T) {
	safeTest(t, "Test_CloneSlice_WithItems", func() {
		// Arrange
		tc := cloneSliceTestCases[1]

		// Act
		result := corestr.CloneSlice([]string{"a", "b"})

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len":   len(result),
			"first": result[0],
		})
	})
}

// ── CloneSliceIf ──

func Test_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_Empty", func() {
		// Arrange
		tc := cloneSliceIfTestCases[0]

		// Act
		result := corestr.CloneSliceIf(true)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": len(result),
		})
	})
}

func Test_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_NoClone", func() {
		// Arrange
		tc := cloneSliceIfTestCases[1]

		// Act
		result := corestr.CloneSliceIf(false, "a", "b")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": len(result),
		})
	})
}

func Test_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_Clone", func() {
		// Arrange
		tc := cloneSliceIfTestCases[2]

		// Act
		result := corestr.CloneSliceIf(true, "a", "b")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": len(result),
		})
	})
}

// ── AnyToString ──

func Test_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_AnyToString_Empty", func() {
		// Arrange
		tc := anyToStringTestCases[0]

		// Act
		result := corestr.AnyToString(false, "")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty": result == "",
		})
	})
}

func Test_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_AnyToString_WithFieldName", func() {
		// Arrange
		tc := anyToStringTestCases[1]

		// Act
		result := corestr.AnyToString(true, "hello")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty": result == "",
		})
	})
}

func Test_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_AnyToString_WithoutFieldName", func() {
		// Arrange
		tc := anyToStringTestCases[2]

		// Act
		result := corestr.AnyToString(false, "hello")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty": result == "",
		})
	})
}

func Test_AnyToString_Ptr(t *testing.T) {
	safeTest(t, "Test_AnyToString_Ptr", func() {
		// Arrange
		tc := anyToStringTestCases[3]
		val := "hello"

		// Act
		result := corestr.AnyToString(false, &val)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"isEmpty": result == "",
		})
	})
}

// ── AllIndividualStringsOfStringsLength ──

func Test_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Nil", func() {
		// Arrange
		tc := allIndividualStringsOfStringsLengthTestCases[0]

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(nil)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": result,
		})
	})
}

func Test_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Empty", func() {
		// Arrange
		tc := allIndividualStringsOfStringsLengthTestCases[1]
		items := [][]string{}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": result,
		})
	})
}

func Test_AllIndividualStringsOfStringsLength_WithItems(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_WithItems", func() {
		// Arrange
		tc := allIndividualStringsOfStringsLengthTestCases[2]
		items := [][]string{{"a", "b"}, {"c"}}

		// Act
		result := corestr.AllIndividualStringsOfStringsLength(&items)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": result,
		})
	})
}

// ── AllIndividualsLengthOfSimpleSlices ──

func Test_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Arrange
		tc := allIndividualsLengthOfSimpleSlicesTestCases[0]

		// Act
		result := corestr.AllIndividualsLengthOfSimpleSlices()

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": result,
		})
	})
}

func Test_AllIndividualsLengthOfSimpleSlices_WithItems(t *testing.T) {
	safeTest(t, "Test_AllIndividualsLengthOfSimpleSlices_WithItems", func() {
		// Arrange
		tc := allIndividualsLengthOfSimpleSlicesTestCases[1]
		s1 := corestr.New.SimpleSlice.Lines("a", "b")
		s2 := corestr.New.SimpleSlice.Lines("c")

		// Act
		result := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"len": result,
		})
	})
}

// ── utils ──

func Test_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Utils_WrapDouble", func() {
		// Arrange
		tc := utilsWrapTestCases[0]

		// Act
		result := corestr.StringUtils.WrapDouble("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Utils_WrapSingle", func() {
		// Arrange
		tc := utilsWrapTestCases[1]

		// Act
		result := corestr.StringUtils.WrapSingle("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Utils_WrapTilda", func() {
		// Arrange
		tc := utilsWrapTestCases[2]

		// Act
		result := corestr.StringUtils.WrapTilda("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Utils_WrapDoubleIfMissing_Empty", func() {
		// Arrange
		tc := utilsWrapTestCases[3]

		// Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Utils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange
		tc := utilsWrapTestCases[4]

		// Act
		result := corestr.StringUtils.WrapDoubleIfMissing(`"a"`)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Utils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange
		tc := utilsWrapTestCases[5]

		// Act
		result := corestr.StringUtils.WrapDoubleIfMissing("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Utils_WrapSingleIfMissing_Empty", func() {
		// Arrange
		tc := utilsWrapTestCases[6]

		// Act
		result := corestr.StringUtils.WrapSingleIfMissing("")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Utils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange
		tc := utilsWrapTestCases[7]

		// Act
		result := corestr.StringUtils.WrapSingleIfMissing("'a'")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

func Test_Utils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Utils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange
		tc := utilsWrapTestCases[8]

		// Act
		result := corestr.StringUtils.WrapSingleIfMissing("a")

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"result": result,
		})
	})
}

// ── Empty creators ──

func Test_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_EmptyCreator_All", func() {
		// Arrange
		tc := emptyCreatorTestCases[0]

		// Act
		allNonNil := corestr.Empty.Collection() != nil &&
			corestr.Empty.LinkedList() != nil &&
			corestr.Empty.SimpleSlice() != nil &&
			corestr.Empty.KeyAnyValuePair() != nil &&
			corestr.Empty.KeyValuePair() != nil &&
			corestr.Empty.KeyValueCollection() != nil &&
			corestr.Empty.LinkedCollections() != nil &&
			corestr.Empty.LeftRight() != nil &&
			corestr.Empty.SimpleStringOncePtr() != nil &&
			corestr.Empty.Hashset() != nil &&
			corestr.Empty.HashsetsCollection() != nil &&
			corestr.Empty.Hashmap() != nil &&
			corestr.Empty.CharCollectionMap() != nil &&
			corestr.Empty.KeyValuesCollection() != nil &&
			corestr.Empty.CollectionsOfCollection() != nil &&
			corestr.Empty.CharHashsetMap() != nil

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"allNonNil": allNonNil,
		})
	})
}

// ── DataModels ──

func Test_CharCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_CharCollectionDataModel", func() {
		// Arrange
		tc := dataModelTestCases[0]

		// Act
		dm := &corestr.CharCollectionDataModel{
			Items:                  map[byte]*corestr.Collection{},
			EachCollectionCapacity: 10,
		}
		ccm := corestr.NewCharCollectionMapUsingDataModel(dm)
		dm2 := corestr.NewCharCollectionMapDataModelUsing(ccm)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil": ccm != nil && dm2 != nil,
		})
	})
}

func Test_CharHashsetDataModel_HelpersModel(t *testing.T) {
	safeTest(t, "Test_CharHashsetDataModel", func() {
		// Arrange
		tc := dataModelTestCases[1]

		// Act
		dm := &corestr.CharHashsetDataModel{
			Items:               map[byte]*corestr.Hashset{},
			EachHashsetCapacity: 10,
		}
		chm := corestr.NewCharHashsetMapUsingDataModel(dm)
		dm2 := corestr.NewCharHashsetMapDataModelUsing(chm)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil": chm != nil && dm2 != nil,
		})
	})
}

func Test_HashmapDataModel_HelpersModel(t *testing.T) {
	safeTest(t, "Test_HashmapDataModel", func() {
		// Arrange
		tc := dataModelTestCases[2]

		// Act
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "b"}}
		hm := corestr.NewHashmapUsingDataModel(dm)
		dm2 := corestr.NewHashmapsDataModelUsing(hm)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil":   dm2 != nil,
			"nonEmpty": hm != nil && !hm.IsEmpty(),
		})
	})
}

func Test_HashsetDataModel_HelpersModel(t *testing.T) {
	safeTest(t, "Test_HashsetDataModel", func() {
		// Arrange
		tc := dataModelTestCases[3]

		// Act
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
		hs := corestr.NewHashsetUsingDataModel(dm)
		dm2 := corestr.NewHashsetsDataModelUsing(hs)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil":   dm2 != nil,
			"nonEmpty": hs != nil && !hs.IsEmpty(),
		})
	})
}

func Test_HashsetsCollectionDataModel_HelpersModel(t *testing.T) {
	safeTest(t, "Test_HashsetsCollectionDataModel", func() {
		// Arrange
		tc := dataModelTestCases[4]

		// Act
		dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{}}
		hc := corestr.NewHashsetsCollectionUsingDataModel(dm)
		dm2 := corestr.NewHashsetsCollectionDataModelUsing(hc)

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil": hc != nil && dm2 != nil,
		})
	})
}

// ── SimpleStringOnceModel ──

func Test_SimpleStringOnceModel(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnceModel", func() {
		// Arrange
		tc := simpleStringOnceModelTestCases[0]

		// Act
		m := corestr.SimpleStringOnceModel{Value: "hello", IsInitialize: true}

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"value": m.Value,
		})
	})
}

// ── CollectionsOfCollectionModel ──

func Test_CollectionsOfCollectionModel_HelpersModel(t *testing.T) {
	safeTest(t, "Test_CollectionsOfCollectionModel", func() {
		// Arrange
		tc := collectionsOfCollectionModelTestCases[0]

		// Act
		m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{}}

		// Assert
		tc.ShouldBeEqualMapFirst(t, args.Map{
			"nonNil": m.Items != nil,
		})
	})
}
