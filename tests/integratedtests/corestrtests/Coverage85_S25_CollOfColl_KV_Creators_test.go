package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// region newCollectionsOfCollectionCreator

func Test_CovS25_01_CollOfCollCreator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS25_01_CollOfCollCreator_Cap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Cap should create non-nil CollectionsOfCollection", actual)
	})
}

func Test_CovS25_02_CollOfCollCreator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_02_CollOfCollCreator_Empty", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Assert
		actual := args.Map{"result": coc == nil || coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty should create empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_03_CollOfCollCreator_LenCap(t *testing.T) {
	safeTest(t, "Test_CovS25_03_CollOfCollCreator_LenCap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 10)

		// Assert
		actual := args.Map{"result": coc == nil || coc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LenCap(0,10) should create empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_04_CollOfCollCreator_Strings(t *testing.T) {
	safeTest(t, "Test_CovS25_04_CollOfCollCreator_Strings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Strings should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_05_CollOfCollCreator_CloneStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_05_CollOfCollCreator_CloneStrings", func() {
		// Arrange
		items := []string{"a", "b", "c"}

		// Act
		coc := corestr.New.CollectionsOfCollection.CloneStrings(items)

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CloneStrings should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_06_CollOfCollCreator_StringsOption(t *testing.T) {
	safeTest(t, "Test_CovS25_06_CollOfCollCreator_StringsOption", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"x"})

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsOption should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_07_CollOfCollCreator_StringsOptions(t *testing.T) {
	safeTest(t, "Test_CovS25_07_CollOfCollCreator_StringsOptions", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 3, []string{"y"})

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsOptions should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_08_CollOfCollCreator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_08_CollOfCollCreator_SpreadStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "SpreadStrings should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_09_CollOfCollCreator_StringsOfStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_09_CollOfCollCreator_StringsOfStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(
			false,
			[]string{"a", "b"},
			[]string{"c"},
		)

		// Assert
		actual := args.Map{"result": coc == nil || !coc.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsOfStrings should create non-empty CollectionsOfCollection", actual)
	})
}

func Test_CovS25_10_CollOfCollCreator_StringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_10_CollOfCollCreator_StringsOfStrings_Empty", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false)

		// Assert
		actual := args.Map{"result": coc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StringsOfStrings with no args should create CollectionsOfCollection", actual)
	})
}

// endregion

// region newKeyValuesCreator

func Test_CovS25_11_KeyValuesCreator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_11_KeyValuesCreator_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Empty()

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Empty should create empty KeyValueCollection", actual)
	})
}

func Test_CovS25_12_KeyValuesCreator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS25_12_KeyValuesCreator_Cap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Cap(10)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Cap should create empty KeyValueCollection with capacity", actual)
	})
}

func Test_CovS25_13_KeyValuesCreator_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovS25_13_KeyValuesCreator_UsingMap", func() {
		// Arrange
		m := map[string]string{"a": "1", "b": "2"}

		// Act
		kvc := corestr.New.KeyValues.UsingMap(m)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingMap should create KeyValueCollection with 2 items", actual)
	})
}

func Test_CovS25_14_KeyValuesCreator_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_14_KeyValuesCreator_UsingMap_Empty", func() {
		// Arrange
		m := map[string]string{}

		// Act
		kvc := corestr.New.KeyValues.UsingMap(m)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingMap empty should create empty KeyValueCollection", actual)
	})
}

func Test_CovS25_15_KeyValuesCreator_UsingKeyValuePairs(t *testing.T) {
	safeTest(t, "Test_CovS25_15_KeyValuesCreator_UsingKeyValuePairs", func() {
		// Arrange
		pair1 := corestr.KeyValuePair{Key: "a", Value: "1"}
		pair2 := corestr.KeyValuePair{Key: "b", Value: "2"}

		// Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs(pair1, pair2)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingKeyValuePairs should create collection with 2 items", actual)
	})
}

func Test_CovS25_16_KeyValuesCreator_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_16_KeyValuesCreator_UsingKeyValuePairs_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingKeyValuePairs with no args should be empty", actual)
	})
}

func Test_CovS25_17_KeyValuesCreator_UsingKeyValueStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_17_KeyValuesCreator_UsingKeyValueStrings", func() {
		// Arrange
		keys := []string{"a", "b"}
		values := []string{"1", "2"}

		// Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings(keys, values)

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingKeyValueStrings should create collection with 2 items", actual)
	})
}

func Test_CovS25_18_KeyValuesCreator_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_18_KeyValuesCreator_UsingKeyValueStrings_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Assert
		actual := args.Map{"result": kvc == nil || kvc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingKeyValueStrings empty should create empty collection", actual)
	})
}

// endregion

// region funcs.go type definitions (compile-time coverage)

func Test_CovS25_19_ReturningBool_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_19_ReturningBool_Fields", func() {
		// Arrange
		rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}

		// Assert
		actual := args.Map{"result": rb.IsBreak || rb.IsKeep}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "ReturningBool fields mismatch", actual)
	})
}

func Test_CovS25_20_LinkedCollectionFilterResult_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_20_LinkedCollectionFilterResult_Fields", func() {
		// Arrange
		r := corestr.LinkedCollectionFilterResult{
			Value:   nil,
			IsKeep:  true,
			IsBreak: false,
		}

		// Assert
		actual := args.Map{"result": r.IsKeep || r.IsBreak || r.Value != nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "LinkedCollectionFilterResult fields mismatch", actual)
	})
}

func Test_CovS25_21_LinkedListFilterResult_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_21_LinkedListFilterResult_Fields", func() {
		// Arrange
		r := corestr.LinkedListFilterResult{
			Value:   nil,
			IsKeep:  false,
			IsBreak: true,
		}

		// Assert
		actual := args.Map{"result": r.IsKeep || !r.IsBreak}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LinkedListFilterResult fields mismatch", actual)
	})
}

func Test_CovS25_22_LinkedCollectionFilterParameter_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_22_LinkedCollectionFilterParameter_Fields", func() {
		// Arrange
		p := corestr.LinkedCollectionFilterParameter{
			Node:  nil,
			Index: 5,
		}

		// Assert
		actual := args.Map{"result": p.Index != 5 || p.Node != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LinkedCollectionFilterParameter fields mismatch", actual)
	})
}

func Test_CovS25_23_LinkedListFilterParameter_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_23_LinkedListFilterParameter_Fields", func() {
		// Arrange
		p := corestr.LinkedListFilterParameter{
			Node:  nil,
			Index: 3,
		}

		// Assert
		actual := args.Map{"result": p.Index != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LinkedListFilterParameter fields mismatch", actual)
	})
}

func Test_CovS25_24_LinkedListProcessorParameter_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_24_LinkedListProcessorParameter_Fields", func() {
		// Arrange
		p := corestr.LinkedListProcessorParameter{
			Index:          0,
			CurrentNode:    nil,
			PrevNode:       nil,
			IsFirstIndex:   true,
			IsEndingIndex:  false,
		}

		// Assert
		actual := args.Map{"result": p.IsFirstIndex || p.IsEndingIndex}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "LinkedListProcessorParameter fields mismatch", actual)
	})
}

func Test_CovS25_25_LinkedCollectionProcessorParameter_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_25_LinkedCollectionProcessorParameter_Fields", func() {
		// Arrange
		p := corestr.LinkedCollectionProcessorParameter{
			Index:          2,
			CurrentNode:    nil,
			PrevNode:       nil,
			IsFirstIndex:   false,
			IsEndingIndex:  true,
		}

		// Assert
		actual := args.Map{"result": p.IsFirstIndex || !p.IsEndingIndex}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LinkedCollectionProcessorParameter fields mismatch", actual)
	})
}

// endregion

// region Exported constants from consts.go

func Test_CovS25_26_RegularCollectionEfficiencyLimit(t *testing.T) {
	safeTest(t, "Test_CovS25_26_RegularCollectionEfficiencyLimit", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.RegularCollectionEfficiencyLimit != 1000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "RegularCollectionEfficiencyLimit expected 1000", actual)
	})
}

func Test_CovS25_27_DoubleLimit(t *testing.T) {
	safeTest(t, "Test_CovS25_27_DoubleLimit", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.DoubleLimit != 3000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "DoubleLimit expected 3000", actual)
	})
}

func Test_CovS25_28_NoElements(t *testing.T) {
	safeTest(t, "Test_CovS25_28_NoElements", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.NoElements != " {No Element}"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NoElements mismatch: ''", actual)
	})
}

// endregion

// region vars.go exported variables

func Test_CovS25_29_StaticJsonError(t *testing.T) {
	safeTest(t, "Test_CovS25_29_StaticJsonError", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.StaticJsonError == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "StaticJsonError should not be nil", actual)
	})
}

func Test_CovS25_30_ExpectingLengthForLeftRight(t *testing.T) {
	safeTest(t, "Test_CovS25_30_ExpectingLengthForLeftRight", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.ExpectingLengthForLeftRight != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ExpectingLengthForLeftRight expected 2", actual)
	})
}

func Test_CovS25_31_LeftRightExpectingLengthMessager(t *testing.T) {
	safeTest(t, "Test_CovS25_31_LeftRightExpectingLengthMessager", func() {
		// Arrange & Act & Assert
		actual := args.Map{"result": corestr.LeftRightExpectingLengthMessager == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "LeftRightExpectingLengthMessager should not be nil", actual)
	})
}

func Test_CovS25_32_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_CovS25_32_StringUtils_WrapDouble", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDouble("test")

		// Assert
		actual := args.Map{"result": result != "\"test\""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDouble expected '\"test\"', got ''", actual)
	})
}

func Test_CovS25_33_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_CovS25_33_StringUtils_WrapSingle", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingle("test")

		// Assert
		actual := args.Map{"result": result != "'test'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingle expected \"'test'\"", actual)
	})
}

func Test_CovS25_34_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_CovS25_34_StringUtils_WrapTilda", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapTilda("test")

		// Assert
		actual := args.Map{"result": result != "`test`"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapTilda expected \"`test`\"", actual)
	})
}

func Test_CovS25_35_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_35_StringUtils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"test\"")

		// Assert
		actual := args.Map{"result": result != "\"test\""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing already wrapped should return same, got ''", actual)
	})
}

func Test_CovS25_36_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_36_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("test")

		// Assert
		actual := args.Map{"result": result != "\"test\""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing should wrap, got ''", actual)
	})
}

func Test_CovS25_37_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_37_StringUtils_WrapDoubleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")

		// Assert
		actual := args.Map{"result": result != "\"\""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing empty should return \"\\\"\\\"\"", actual)
	})
}

func Test_CovS25_38_StringUtils_WrapDoubleIfMissing_QuotedEmpty(t *testing.T) {
	safeTest(t, "Test_CovS25_38_StringUtils_WrapDoubleIfMissing_QuotedEmpty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"\"")

		// Assert
		actual := args.Map{"result": result != "\"\""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing '\"\"' should return same, got ''", actual)
	})
}

func Test_CovS25_39_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_39_StringUtils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("'test'")

		// Assert
		actual := args.Map{"result": result != "'test'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing already wrapped should return same, got ''", actual)
	})
}

func Test_CovS25_40_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_40_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("test")

		// Assert
		actual := args.Map{"result": result != "'test'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing should wrap, got ''", actual)
	})
}

func Test_CovS25_41_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_41_StringUtils_WrapSingleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("")

		// Assert
		actual := args.Map{"result": result != "''"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing empty should return '', got ''", actual)
	})
}

func Test_CovS25_42_StringUtils_WrapSingleIfMissing_QuotedEmpty(t *testing.T) {
	safeTest(t, "Test_CovS25_42_StringUtils_WrapSingleIfMissing_QuotedEmpty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("''")

		// Assert
		actual := args.Map{"result": result != "''"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing '' should return same, got ''", actual)
	})
}

// endregion
