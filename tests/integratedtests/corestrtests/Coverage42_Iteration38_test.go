package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ─── Collection: AddHashmapsKeysValues ──────

func Test_Cov42_Collection_AddHashmapsKeysValues_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddHashmapsKeysValues_Valid", func() {
		col := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		col.AddHashmapsKeysValues(hm)
		tc := coretestcases.CaseV1{
			Title:         "AddHashmapsKeysValues adds both",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddHashmapsKeysValues_Nil", func() {
		col := corestr.New.Collection.Empty()
		col.AddHashmapsKeysValues(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddHashmapsKeysValues nil",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddHashmapsKeysValuesUsingFilter ──────

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Accept(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Accept", func() {
		col := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		}
		col.AddHashmapsKeysValuesUsingFilter(filter, hm)
		tc := coretestcases.CaseV1{
			Title:         "AddHashmapsKeysValuesUsingFilter accept",
			ExpectedInput: true,
			ActualInput:   col.Has("v1"),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		col := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.KeyValues(corestr.KeyValuePair{Key: "k1", Value: "v1"})
		filter := func(pair corestr.KeyValuePair) (string, bool, bool) {
			return pair.Value, false, true
		}
		col.AddHashmapsKeysValuesUsingFilter(filter, hm)
		tc := coretestcases.CaseV1{
			Title:         "AddHashmapsKeysValuesUsingFilter break",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		col := corestr.New.Collection.Empty()
		col.AddHashmapsKeysValuesUsingFilter(nil, nil)
		tc := coretestcases.CaseV1{
			Title:         "AddHashmapsKeysValuesUsingFilter nil",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddPointerCollectionsLock ──────

func Test_Cov42_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddPointerCollectionsLock", func() {
		col := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a", "b"})
		col.AddPointerCollectionsLock(other)
		tc := coretestcases.CaseV1{
			Title:         "AddPointerCollectionsLock",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AppendCollectionPtr ──────

func Test_Cov42_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendCollectionPtr", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		other := corestr.New.Collection.Strings([]string{"y", "z"})
		col.AppendCollectionPtr(other)
		tc := coretestcases.CaseV1{
			Title:         "AppendCollectionPtr",
			ExpectedInput: 3,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: Single ──────

func Test_Cov42_Collection_Single_OneItem(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Single_OneItem", func() {
		col := corestr.New.Collection.Strings([]string{"only"})
		tc := coretestcases.CaseV1{
			Title:         "Single with one item",
			ExpectedInput: "only",
			ActualInput:   col.Single(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: SortedListDsc ──────

func Test_Cov42_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_SortedListDsc", func() {
		col := corestr.New.Collection.Strings([]string{"apple", "cherry", "banana"})
		sorted := col.SortedListDsc()
		tc := coretestcases.CaseV1{
			Title:         "SortedListDsc first item",
			ExpectedInput: "cherry",
			ActualInput:   sorted[0],
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: HasUsingSensitivity ──────

func Test_Cov42_Collection_HasUsingSensitivity_CaseSensitive(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_HasUsingSensitivity_CaseSensitive", func() {
		col := corestr.New.Collection.Strings([]string{"Hello"})
		tc := coretestcases.CaseV1{
			Title:         "HasUsingSensitivity case sensitive miss",
			ExpectedInput: false,
			ActualInput:   col.HasUsingSensitivity("hello", true),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_HasUsingSensitivity_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_HasUsingSensitivity_CaseInsensitive", func() {
		col := corestr.New.Collection.Strings([]string{"Hello"})
		tc := coretestcases.CaseV1{
			Title:         "HasUsingSensitivity case insensitive match",
			ExpectedInput: true,
			ActualInput:   col.HasUsingSensitivity("hello", false),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddStringsAsync ──────

func Test_Cov42_Collection_AddStringsAsync_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddStringsAsync_Empty", func() {
		col := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		col.AddStringsAsync(wg, []string{})
		tc := coretestcases.CaseV1{
			Title:         "AddStringsAsync empty",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddNonEmptyStrings / AddNonEmptyStringsSlice ──────

func Test_Cov42_Collection_AddNonEmptyStrings_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddNonEmptyStrings_Valid", func() {
		col := corestr.New.Collection.Empty()
		col.AddNonEmptyStrings("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmptyStrings valid",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddNonEmptyStrings_Empty", func() {
		col := corestr.New.Collection.Empty()
		col.AddNonEmptyStrings()
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmptyStrings no args",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddNonEmptyStringsSlice_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddNonEmptyStringsSlice_Valid", func() {
		col := corestr.New.Collection.Empty()
		col.AddNonEmptyStringsSlice([]string{"x", "y"})
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmptyStringsSlice valid",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddNonEmptyStringsSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddNonEmptyStringsSlice_Empty", func() {
		col := corestr.New.Collection.Empty()
		col.AddNonEmptyStringsSlice([]string{})
		tc := coretestcases.CaseV1{
			Title:         "AddNonEmptyStringsSlice empty",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddFuncResult ──────

func Test_Cov42_Collection_AddFuncResult_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddFuncResult_Valid", func() {
		col := corestr.New.Collection.Empty()
		col.AddFuncResult(func() string { return "hello" })
		tc := coretestcases.CaseV1{
			Title:         "AddFuncResult valid",
			ExpectedInput: "hello",
			ActualInput:   col.First(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AddFuncResult_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddFuncResult_Nil", func() {
		col := corestr.New.Collection.Empty()
		col.AddFuncResult(nil)
		tc := coretestcases.CaseV1{
			Title:         "AddFuncResult nil",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AddStringsByFuncChecking ──────

func Test_Cov42_Collection_AddStringsByFuncChecking_Filter(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AddStringsByFuncChecking_Filter", func() {
		col := corestr.New.Collection.Empty()
		col.AddStringsByFuncChecking(
			[]string{"apple", "ban", "cherry"},
			func(line string) bool { return len(line) > 3 },
		)
		tc := coretestcases.CaseV1{
			Title:         "AddStringsByFuncChecking filters",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: ExpandSlicePlusAdd ──────

func Test_Cov42_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_ExpandSlicePlusAdd", func() {
		col := corestr.New.Collection.Empty()
		col.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string {
				return []string{line}
			},
		)
		tc := coretestcases.CaseV1{
			Title:         "ExpandSlicePlusAdd",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: MergeSlicesOfSlice ──────

func Test_Cov42_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_MergeSlicesOfSlice", func() {
		col := corestr.New.Collection.Empty()
		col.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})
		tc := coretestcases.CaseV1{
			Title:         "MergeSlicesOfSlice",
			ExpectedInput: 3,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: CharCollectionMap ──────

func Test_Cov42_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_CharCollectionMap", func() {
		col := corestr.New.Collection.Strings([]string{"apple", "banana", "avocado"})
		ccm := col.CharCollectionMap()
		tc := coretestcases.CaseV1{
			Title:         "CharCollectionMap groups by first char",
			ExpectedInput: 2,
			ActualInput:   ccm.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: CsvLines / CsvLinesOptions / Csv / CsvOptions ──────

func Test_Cov42_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_CsvLines", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lines := col.CsvLines()
		tc := coretestcases.CaseV1{
			Title:         "CsvLines",
			ExpectedInput: 2,
			ActualInput:   len(lines),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_CsvLinesOptions", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		lines := col.CsvLinesOptions(true)
		tc := coretestcases.CaseV1{
			Title:         "CsvLinesOptions",
			ExpectedInput: 1,
			ActualInput:   len(lines),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_Csv_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Csv_Empty", func() {
		col := corestr.New.Collection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "Csv empty",
			ExpectedInput: "",
			ActualInput:   col.Csv(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_Csv_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Csv_NonEmpty", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		csv := col.Csv()
		tc := coretestcases.CaseV1{
			Title:         "Csv non-empty",
			ExpectedInput: true,
			ActualInput:   len(csv) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_CsvOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_CsvOptions_Empty", func() {
		col := corestr.New.Collection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "CsvOptions empty",
			ExpectedInput: "",
			ActualInput:   col.CsvOptions(true),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_CsvOptions_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_CsvOptions_NonEmpty", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		csv := col.CsvOptions(false)
		tc := coretestcases.CaseV1{
			Title:         "CsvOptions non-empty",
			ExpectedInput: true,
			ActualInput:   len(csv) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: IsContainsPtr ──────

func Test_Cov42_Collection_IsContainsPtr_Found(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsPtr_Found", func() {
		col := corestr.New.Collection.Strings([]string{"hello"})
		s := "hello"
		tc := coretestcases.CaseV1{
			Title:         "IsContainsPtr found",
			ExpectedInput: true,
			ActualInput:   col.IsContainsPtr(&s),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsPtr_Nil", func() {
		col := corestr.New.Collection.Strings([]string{"hello"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsPtr nil",
			ExpectedInput: false,
			ActualInput:   col.IsContainsPtr(nil),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsPtr_Empty", func() {
		col := corestr.New.Collection.Empty()
		s := "hello"
		tc := coretestcases.CaseV1{
			Title:         "IsContainsPtr empty collection",
			ExpectedInput: false,
			ActualInput:   col.IsContainsPtr(&s),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: GetHashsetPlusHasAll ──────

func Test_Cov42_Collection_GetHashsetPlusHasAll_True(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetHashsetPlusHasAll_True", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		hs, hasAll := col.GetHashsetPlusHasAll([]string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "GetHashsetPlusHasAll true",
			ExpectedInput: true,
			ActualInput:   hasAll && hs != nil,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_GetHashsetPlusHasAll_False(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetHashsetPlusHasAll_False", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		_, hasAll := col.GetHashsetPlusHasAll([]string{"a", "c"})
		tc := coretestcases.CaseV1{
			Title:         "GetHashsetPlusHasAll false",
			ExpectedInput: false,
			ActualInput:   hasAll,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: IsContainsAllSlice ──────

func Test_Cov42_Collection_IsContainsAllSlice_True(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllSlice_True", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllSlice true",
			ExpectedInput: true,
			ActualInput:   col.IsContainsAllSlice([]string{"a", "b"}),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsAllSlice_False(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllSlice_False", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllSlice false",
			ExpectedInput: false,
			ActualInput:   col.IsContainsAllSlice([]string{"a", "c"}),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsAllSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllSlice_Empty", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllSlice empty items",
			ExpectedInput: false,
			ActualInput:   col.IsContainsAllSlice([]string{}),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsAllSlice_EmptyCollection(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllSlice_EmptyCollection", func() {
		col := corestr.New.Collection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllSlice empty collection",
			ExpectedInput: false,
			ActualInput:   col.IsContainsAllSlice([]string{"a"}),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: IsContainsAll / IsContainsAllLock ──────

func Test_Cov42_Collection_IsContainsAll_True(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAll_True", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAll true",
			ExpectedInput: true,
			ActualInput:   col.IsContainsAll("a", "b"),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsAllLock_True(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllLock_True", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllLock true",
			ExpectedInput: true,
			ActualInput:   col.IsContainsAllLock("a", "b"),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_IsContainsAllLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_IsContainsAllLock_Nil", func() {
		col := corestr.New.Collection.Empty()
		tc := coretestcases.CaseV1{
			Title:         "IsContainsAllLock nil",
			ExpectedInput: false,
			ActualInput:   col.IsContainsAllLock("x"),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: New (instance method) ──────

func Test_Cov42_Collection_New_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_New_Empty", func() {
		col := corestr.New.Collection.Empty()
		newCol := col.New()
		tc := coretestcases.CaseV1{
			Title:         "Collection.New empty",
			ExpectedInput: 0,
			ActualInput:   newCol.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_New_WithArgs(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_New_WithArgs", func() {
		col := corestr.New.Collection.Empty()
		newCol := col.New("a", "b")
		tc := coretestcases.CaseV1{
			Title:         "Collection.New with args",
			ExpectedInput: 2,
			ActualInput:   newCol.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── newCollectionCreator: CloneStrings / LineUsingSep / LineDefault / StringsPlusCap / CapStrings / LenCap ──────

func Test_Cov42_newCollectionCreator_CloneStrings(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_CloneStrings", func() {
		items := []string{"a", "b"}
		col := corestr.New.Collection.CloneStrings(items)
		items[0] = "changed"
		tc := coretestcases.CaseV1{
			Title:         "CloneStrings is independent",
			ExpectedInput: "a",
			ActualInput:   col.First(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_LineUsingSep(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_LineUsingSep", func() {
		col := corestr.New.Collection.LineUsingSep(",", "a,b,c")
		tc := coretestcases.CaseV1{
			Title:         "LineUsingSep",
			ExpectedInput: 3,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_LineDefault(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_LineDefault", func() {
		col := corestr.New.Collection.LineDefault("a\nb")
		tc := coretestcases.CaseV1{
			Title:         "LineDefault",
			ExpectedInput: true,
			ActualInput:   col.Length() >= 1,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_StringsPlusCap_ZeroCap(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_StringsPlusCap_ZeroCap", func() {
		col := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		tc := coretestcases.CaseV1{
			Title:         "StringsPlusCap zero cap",
			ExpectedInput: 1,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_StringsPlusCap_WithCap(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_StringsPlusCap_WithCap", func() {
		col := corestr.New.Collection.StringsPlusCap(10, []string{"a", "b"})
		tc := coretestcases.CaseV1{
			Title:         "StringsPlusCap with cap",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_CapStrings_ZeroCap(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_CapStrings_ZeroCap", func() {
		col := corestr.New.Collection.CapStrings(0, []string{"x"})
		tc := coretestcases.CaseV1{
			Title:         "CapStrings zero cap",
			ExpectedInput: 1,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_CapStrings_WithCap(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_CapStrings_WithCap", func() {
		col := corestr.New.Collection.CapStrings(5, []string{"x"})
		tc := coretestcases.CaseV1{
			Title:         "CapStrings with cap",
			ExpectedInput: 1,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_LenCap(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_LenCap", func() {
		col := corestr.New.Collection.LenCap(3, 10)
		tc := coretestcases.CaseV1{
			Title:         "LenCap creates with length",
			ExpectedInput: 3,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_Create(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_Create", func() {
		col := corestr.New.Collection.Create([]string{"a"})
		tc := coretestcases.CaseV1{
			Title:         "Create wraps slice",
			ExpectedInput: "a",
			ActualInput:   col.First(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_StringsOptions_Clone", func() {
		items := []string{"x"}
		col := corestr.New.Collection.StringsOptions(true, items)
		items[0] = "changed"
		tc := coretestcases.CaseV1{
			Title:         "StringsOptions clone",
			ExpectedInput: "x",
			ActualInput:   col.First(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_newCollectionCreator_StringsOptions_NoClone_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_newCollectionCreator_StringsOptions_NoClone_Empty", func() {
		col := corestr.New.Collection.StringsOptions(false, []string{})
		tc := coretestcases.CaseV1{
			Title:         "StringsOptions no clone empty",
			ExpectedInput: true,
			ActualInput:   col.IsEmpty(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AppendAnys with nil items ──────

func Test_Cov42_Collection_AppendAnys_WithNilItem(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnys_WithNilItem", func() {
		col := corestr.New.Collection.Empty()
		col.AppendAnys("hello", nil, "world")
		tc := coretestcases.CaseV1{
			Title:         "AppendAnys skips nil",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AppendAnysUsingFilter ──────

func Test_Cov42_Collection_AppendAnysUsingFilter_Accept(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnysUsingFilter_Accept", func() {
		col := corestr.New.Collection.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		col.AppendAnysUsingFilter(filter, "a", "b")
		tc := coretestcases.CaseV1{
			Title:         "AppendAnysUsingFilter accept",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnysUsingFilter_Break", func() {
		col := corestr.New.Collection.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		col.AppendAnysUsingFilter(filter, "a", "b")
		tc := coretestcases.CaseV1{
			Title:         "AppendAnysUsingFilter break after first",
			ExpectedInput: 1,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AppendAnysUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnysUsingFilter_Skip", func() {
		col := corestr.New.Collection.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, false, false
		}
		col.AppendAnysUsingFilter(filter, "a")
		tc := coretestcases.CaseV1{
			Title:         "AppendAnysUsingFilter skip",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AppendAnysUsingFilterLock ──────

func Test_Cov42_Collection_AppendAnysUsingFilterLock_Accept(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnysUsingFilterLock_Accept", func() {
		col := corestr.New.Collection.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		col.AppendAnysUsingFilterLock(filter, "x")
		tc := coretestcases.CaseV1{
			Title:         "AppendAnysUsingFilterLock accept",
			ExpectedInput: 1,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AppendAnysUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendAnysUsingFilterLock_Nil", func() {
		col := corestr.New.Collection.Empty()
		col.AppendAnysUsingFilterLock(nil, nil)
		tc := coretestcases.CaseV1{
			Title:         "AppendAnysUsingFilterLock nil args",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: AppendNonEmptyAnys ──────

func Test_Cov42_Collection_AppendNonEmptyAnys_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendNonEmptyAnys_Valid", func() {
		col := corestr.New.Collection.Empty()
		col.AppendNonEmptyAnys("hello", nil, "world")
		tc := coretestcases.CaseV1{
			Title:         "AppendNonEmptyAnys skips nil",
			ExpectedInput: 2,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_AppendNonEmptyAnys_Nil", func() {
		col := corestr.New.Collection.Empty()
		col.AppendNonEmptyAnys(nil)
		tc := coretestcases.CaseV1{
			Title:         "AppendNonEmptyAnys nil args",
			ExpectedInput: 0,
			ActualInput:   col.Length(),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: GetAllExceptCollection / GetAllExcept ──────

func Test_Cov42_Collection_GetAllExceptCollection_WithExclude(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetAllExceptCollection_WithExclude", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		exclude := corestr.New.Collection.Strings([]string{"b"})
		result := col.GetAllExceptCollection(exclude)
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection excludes",
			ExpectedInput: 2,
			ActualInput:   len(result),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_GetAllExceptCollection_NilExclude(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetAllExceptCollection_NilExclude", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		result := col.GetAllExceptCollection(nil)
		tc := coretestcases.CaseV1{
			Title:         "GetAllExceptCollection nil returns copy",
			ExpectedInput: 2,
			ActualInput:   len(result),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_GetAllExcept_Valid(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetAllExcept_Valid", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := col.GetAllExcept([]string{"c"})
		tc := coretestcases.CaseV1{
			Title:         "GetAllExcept excludes",
			ExpectedInput: 2,
			ActualInput:   len(result),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_GetAllExcept_Nil", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		result := col.GetAllExcept(nil)
		tc := coretestcases.CaseV1{
			Title:         "GetAllExcept nil returns copy",
			ExpectedInput: 1,
			ActualInput:   len(result),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: SummaryString / SummaryStringWithHeader ──────

func Test_Cov42_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_SummaryString", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		s := col.SummaryString(1)
		tc := coretestcases.CaseV1{
			Title:         "SummaryString",
			ExpectedInput: true,
			ActualInput:   len(s) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_SummaryStringWithHeader_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_SummaryStringWithHeader_NonEmpty", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		s := col.SummaryStringWithHeader("header")
		tc := coretestcases.CaseV1{
			Title:         "SummaryStringWithHeader non-empty",
			ExpectedInput: true,
			ActualInput:   len(s) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_SummaryStringWithHeader_Empty", func() {
		col := corestr.New.Collection.Empty()
		s := col.SummaryStringWithHeader("header")
		tc := coretestcases.CaseV1{
			Title:         "SummaryStringWithHeader empty",
			ExpectedInput: true,
			ActualInput:   len(s) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: Joins with extra items ──────

func Test_Cov42_Collection_Joins_WithExtra(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Joins_WithExtra", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		result := col.Joins(",", "c")
		tc := coretestcases.CaseV1{
			Title:         "Joins with extra items",
			ExpectedInput: true,
			ActualInput:   len(result) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_Joins_NoExtra(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Joins_NoExtra", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		result := col.Joins(",")
		tc := coretestcases.CaseV1{
			Title:         "Joins no extra",
			ExpectedInput: "a,b",
			ActualInput:   result,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

// ─── Collection: Serialize / Deserialize ──────

func Test_Cov42_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_Serialize", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		bytes, err := col.Serialize()
		tc := coretestcases.CaseV1{
			Title:         "Serialize success",
			ExpectedInput: true,
			ActualInput:   err == nil && len(bytes) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}
func Test_Cov42_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_NonEmptyList", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		list := col.NonEmptyList()
		tc := coretestcases.CaseV1{
			Title:         "NonEmptyList filters empty",
			ExpectedInput: 2,
			ActualInput:   len(list),
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}
func Test_Cov42_Collection_StringLock_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_StringLock_NonEmpty", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		s := col.StringLock()
		tc := coretestcases.CaseV1{
			Title:         "StringLock non-empty",
			ExpectedInput: true,
			ActualInput:   len(s) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}

func Test_Cov42_Collection_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov42_Collection_StringLock_Empty", func() {
		col := corestr.New.Collection.Empty()
		s := col.StringLock()
		tc := coretestcases.CaseV1{
			Title:         "StringLock empty",
			ExpectedInput: true,
			ActualInput:   len(s) > 0,
		}
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", tc.Actual()))
	})
}
