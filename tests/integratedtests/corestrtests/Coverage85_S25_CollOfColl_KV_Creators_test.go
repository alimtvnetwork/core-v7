package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// region newCollectionsOfCollectionCreator

func Test_CovS25_01_CollOfCollCreator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS25_01_CollOfCollCreator_Cap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Assert
		if coc == nil {
			t.Errorf("Cap should create non-nil CollectionsOfCollection")
		}
	})
}

func Test_CovS25_02_CollOfCollCreator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_02_CollOfCollCreator_Empty", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Empty()

		// Assert
		if coc == nil || coc.Length() != 0 {
			t.Errorf("Empty should create empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_03_CollOfCollCreator_LenCap(t *testing.T) {
	safeTest(t, "Test_CovS25_03_CollOfCollCreator_LenCap", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 10)

		// Assert
		if coc == nil || coc.Length() != 0 {
			t.Errorf("LenCap(0,10) should create empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_04_CollOfCollCreator_Strings(t *testing.T) {
	safeTest(t, "Test_CovS25_04_CollOfCollCreator_Strings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.Strings([]string{"a", "b"})

		// Assert
		if coc == nil || !coc.HasItems() {
			t.Errorf("Strings should create non-empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_05_CollOfCollCreator_CloneStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_05_CollOfCollCreator_CloneStrings", func() {
		// Arrange
		items := []string{"a", "b", "c"}

		// Act
		coc := corestr.New.CollectionsOfCollection.CloneStrings(items)

		// Assert
		if coc == nil || !coc.HasItems() {
			t.Errorf("CloneStrings should create non-empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_06_CollOfCollCreator_StringsOption(t *testing.T) {
	safeTest(t, "Test_CovS25_06_CollOfCollCreator_StringsOption", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOption(true, 5, []string{"x"})

		// Assert
		if coc == nil || !coc.HasItems() {
			t.Errorf("StringsOption should create non-empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_07_CollOfCollCreator_StringsOptions(t *testing.T) {
	safeTest(t, "Test_CovS25_07_CollOfCollCreator_StringsOptions", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOptions(false, 3, []string{"y"})

		// Assert
		if coc == nil || !coc.HasItems() {
			t.Errorf("StringsOptions should create non-empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_08_CollOfCollCreator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_CovS25_08_CollOfCollCreator_SpreadStrings", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.SpreadStrings(false, "a", "b")

		// Assert
		if coc == nil || !coc.HasItems() {
			t.Errorf("SpreadStrings should create non-empty CollectionsOfCollection")
		}
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
		if coc == nil || !coc.HasItems() {
			t.Errorf("StringsOfStrings should create non-empty CollectionsOfCollection")
		}
	})
}

func Test_CovS25_10_CollOfCollCreator_StringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_10_CollOfCollCreator_StringsOfStrings_Empty", func() {
		// Arrange & Act
		coc := corestr.New.CollectionsOfCollection.StringsOfStrings(false)

		// Assert
		if coc == nil {
			t.Errorf("StringsOfStrings with no args should create CollectionsOfCollection")
		}
	})
}

// endregion

// region newKeyValuesCreator

func Test_CovS25_11_KeyValuesCreator_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_11_KeyValuesCreator_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Empty()

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Errorf("Empty should create empty KeyValueCollection")
		}
	})
}

func Test_CovS25_12_KeyValuesCreator_Cap(t *testing.T) {
	safeTest(t, "Test_CovS25_12_KeyValuesCreator_Cap", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.Cap(10)

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Errorf("Cap should create empty KeyValueCollection with capacity")
		}
	})
}

func Test_CovS25_13_KeyValuesCreator_UsingMap(t *testing.T) {
	safeTest(t, "Test_CovS25_13_KeyValuesCreator_UsingMap", func() {
		// Arrange
		m := map[string]string{"a": "1", "b": "2"}

		// Act
		kvc := corestr.New.KeyValues.UsingMap(m)

		// Assert
		if kvc == nil || kvc.Length() != 2 {
			t.Errorf("UsingMap should create KeyValueCollection with 2 items, got %d", kvc.Length())
		}
	})
}

func Test_CovS25_14_KeyValuesCreator_UsingMap_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_14_KeyValuesCreator_UsingMap_Empty", func() {
		// Arrange
		m := map[string]string{}

		// Act
		kvc := corestr.New.KeyValues.UsingMap(m)

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Errorf("UsingMap empty should create empty KeyValueCollection")
		}
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
		if kvc == nil || kvc.Length() != 2 {
			t.Errorf("UsingKeyValuePairs should create collection with 2 items, got %d", kvc.Length())
		}
	})
}

func Test_CovS25_16_KeyValuesCreator_UsingKeyValuePairs_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_16_KeyValuesCreator_UsingKeyValuePairs_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValuePairs()

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Errorf("UsingKeyValuePairs with no args should be empty")
		}
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
		if kvc == nil || kvc.Length() != 2 {
			t.Errorf("UsingKeyValueStrings should create collection with 2 items, got %d", kvc.Length())
		}
	})
}

func Test_CovS25_18_KeyValuesCreator_UsingKeyValueStrings_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_18_KeyValuesCreator_UsingKeyValueStrings_Empty", func() {
		// Arrange & Act
		kvc := corestr.New.KeyValues.UsingKeyValueStrings([]string{}, []string{})

		// Assert
		if kvc == nil || kvc.Length() != 0 {
			t.Errorf("UsingKeyValueStrings empty should create empty collection")
		}
	})
}

// endregion

// region funcs.go type definitions (compile-time coverage)

func Test_CovS25_19_ReturningBool_Fields(t *testing.T) {
	safeTest(t, "Test_CovS25_19_ReturningBool_Fields", func() {
		// Arrange
		rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}

		// Assert
		if !rb.IsBreak || rb.IsKeep {
			t.Errorf("ReturningBool fields mismatch")
		}
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
		if !r.IsKeep || r.IsBreak || r.Value != nil {
			t.Errorf("LinkedCollectionFilterResult fields mismatch")
		}
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
		if r.IsKeep || !r.IsBreak {
			t.Errorf("LinkedListFilterResult fields mismatch")
		}
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
		if p.Index != 5 || p.Node != nil {
			t.Errorf("LinkedCollectionFilterParameter fields mismatch")
		}
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
		if p.Index != 3 {
			t.Errorf("LinkedListFilterParameter fields mismatch")
		}
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
		if !p.IsFirstIndex || p.IsEndingIndex {
			t.Errorf("LinkedListProcessorParameter fields mismatch")
		}
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
		if p.IsFirstIndex || !p.IsEndingIndex {
			t.Errorf("LinkedCollectionProcessorParameter fields mismatch")
		}
	})
}

// endregion

// region Exported constants from consts.go

func Test_CovS25_26_RegularCollectionEfficiencyLimit(t *testing.T) {
	safeTest(t, "Test_CovS25_26_RegularCollectionEfficiencyLimit", func() {
		// Arrange & Act & Assert
		if corestr.RegularCollectionEfficiencyLimit != 1000 {
			t.Errorf("RegularCollectionEfficiencyLimit expected 1000, got %d", corestr.RegularCollectionEfficiencyLimit)
		}
	})
}

func Test_CovS25_27_DoubleLimit(t *testing.T) {
	safeTest(t, "Test_CovS25_27_DoubleLimit", func() {
		// Arrange & Act & Assert
		if corestr.DoubleLimit != 3000 {
			t.Errorf("DoubleLimit expected 3000, got %d", corestr.DoubleLimit)
		}
	})
}

func Test_CovS25_28_NoElements(t *testing.T) {
	safeTest(t, "Test_CovS25_28_NoElements", func() {
		// Arrange & Act & Assert
		if corestr.NoElements != " {No Element}" {
			t.Errorf("NoElements mismatch: '%s'", corestr.NoElements)
		}
	})
}

// endregion

// region vars.go exported variables

func Test_CovS25_29_StaticJsonError(t *testing.T) {
	safeTest(t, "Test_CovS25_29_StaticJsonError", func() {
		// Arrange & Act & Assert
		if corestr.StaticJsonError == nil {
			t.Errorf("StaticJsonError should not be nil")
		}
	})
}

func Test_CovS25_30_ExpectingLengthForLeftRight(t *testing.T) {
	safeTest(t, "Test_CovS25_30_ExpectingLengthForLeftRight", func() {
		// Arrange & Act & Assert
		if corestr.ExpectingLengthForLeftRight != 2 {
			t.Errorf("ExpectingLengthForLeftRight expected 2, got %d", corestr.ExpectingLengthForLeftRight)
		}
	})
}

func Test_CovS25_31_LeftRightExpectingLengthMessager(t *testing.T) {
	safeTest(t, "Test_CovS25_31_LeftRightExpectingLengthMessager", func() {
		// Arrange & Act & Assert
		if corestr.LeftRightExpectingLengthMessager == nil {
			t.Errorf("LeftRightExpectingLengthMessager should not be nil")
		}
	})
}

func Test_CovS25_32_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_CovS25_32_StringUtils_WrapDouble", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDouble("test")

		// Assert
		if result != "\"test\"" {
			t.Errorf("WrapDouble expected '\"test\"', got '%s'", result)
		}
	})
}

func Test_CovS25_33_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_CovS25_33_StringUtils_WrapSingle", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingle("test")

		// Assert
		if result != "'test'" {
			t.Errorf("WrapSingle expected \"'test'\", got '%s'", result)
		}
	})
}

func Test_CovS25_34_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_CovS25_34_StringUtils_WrapTilda", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapTilda("test")

		// Assert
		if result != "`test`" {
			t.Errorf("WrapTilda expected \"`test`\", got '%s'", result)
		}
	})
}

func Test_CovS25_35_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_35_StringUtils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"test\"")

		// Assert
		if result != "\"test\"" {
			t.Errorf("WrapDoubleIfMissing already wrapped should return same, got '%s'", result)
		}
	})
}

func Test_CovS25_36_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_36_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("test")

		// Assert
		if result != "\"test\"" {
			t.Errorf("WrapDoubleIfMissing should wrap, got '%s'", result)
		}
	})
}

func Test_CovS25_37_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_37_StringUtils_WrapDoubleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")

		// Assert
		if result != "\"\"" {
			t.Errorf("WrapDoubleIfMissing empty should return \"\\\"\\\"\", got '%s'", result)
		}
	})
}

func Test_CovS25_38_StringUtils_WrapDoubleIfMissing_QuotedEmpty(t *testing.T) {
	safeTest(t, "Test_CovS25_38_StringUtils_WrapDoubleIfMissing_QuotedEmpty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("\"\"")

		// Assert
		if result != "\"\"" {
			t.Errorf("WrapDoubleIfMissing '\"\"' should return same, got '%s'", result)
		}
	})
}

func Test_CovS25_39_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_39_StringUtils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("'test'")

		// Assert
		if result != "'test'" {
			t.Errorf("WrapSingleIfMissing already wrapped should return same, got '%s'", result)
		}
	})
}

func Test_CovS25_40_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_CovS25_40_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("test")

		// Assert
		if result != "'test'" {
			t.Errorf("WrapSingleIfMissing should wrap, got '%s'", result)
		}
	})
}

func Test_CovS25_41_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_CovS25_41_StringUtils_WrapSingleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("")

		// Assert
		if result != "''" {
			t.Errorf("WrapSingleIfMissing empty should return '', got '%s'", result)
		}
	})
}

func Test_CovS25_42_StringUtils_WrapSingleIfMissing_QuotedEmpty(t *testing.T) {
	safeTest(t, "Test_CovS25_42_StringUtils_WrapSingleIfMissing_QuotedEmpty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("''")

		// Assert
		if result != "''" {
			t.Errorf("WrapSingleIfMissing '' should return same, got '%s'", result)
		}
	})
}

// endregion
