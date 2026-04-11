package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// Collection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "LengthLock", Expected: 2, Actual: c.LengthLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEquals_Same", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "IsEquals same", Expected: true, Actual: c1.IsEquals(c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEquals_Different(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEquals_Different", func() {
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"a", "c"})
		tc := caseV1Compat{Name: "IsEquals diff", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		c1 := corestr.New.Collection.Strings([]string{"ABC"})
		c2 := corestr.New.Collection.Strings([]string{"abc"})
		tc := caseV1Compat{Name: "IsEqualsWithSensitive insensitive", Expected: true, Actual: c1.IsEqualsWithSensitive(false, c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEqualsWithSensitive_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEqualsWithSensitive_BothNil", func() {
		var c1, c2 *corestr.Collection
		tc := caseV1Compat{Name: "IsEquals both nil", Expected: true, Actual: c1.IsEquals(c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEqualsWithSensitive_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEqualsWithSensitive_OneNil", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		var c2 *corestr.Collection
		tc := caseV1Compat{Name: "IsEquals one nil", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEqualsWithSensitive_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEqualsWithSensitive_DiffLength", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "IsEquals diff length", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "IsEmptyLock", Expected: true, Actual: c.IsEmptyLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "HasItems", Expected: true, Actual: c.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddLock", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddLock("hello")
		tc := caseV1Compat{Name: "AddLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmpty("")
		c.AddNonEmpty("x")
		tc := caseV1Compat{Name: "AddNonEmpty", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmptyWhitespace("  ")
		c.AddNonEmptyWhitespace("x")
		tc := caseV1Compat{Name: "AddNonEmptyWhitespace", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddError", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddError(nil)
		tc := caseV1Compat{Name: "AddError nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AsDefaultError_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AsDefaultError_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "AsDefaultError empty", Expected: true, Actual: c.AsDefaultError() == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AsError_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AsError_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		err := c.AsError(",")
		tc := caseV1Compat{Name: "AsError has items", Expected: true, Actual: err != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddIf_True", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddIf(true, "yes")
		tc := caseV1Compat{Name: "AddIf true", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddIf_False", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddIf(false, "no")
		tc := caseV1Compat{Name: "AddIf false", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddIfMany_True(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddIfMany_True", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddIfMany(true, "a", "b")
		tc := caseV1Compat{Name: "AddIfMany true", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddIfMany_False(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddIfMany_False", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddIfMany(false, "a", "b")
		tc := caseV1Compat{Name: "AddIfMany false", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddFunc", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddFunc(func() string { return "hello" })
		tc := caseV1Compat{Name: "AddFunc", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddFuncErr_NoError", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		tc := caseV1Compat{Name: "AddFuncErr no err", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_AddsLock", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddsLock("a", "b")
		tc := caseV1Compat{Name: "AddsLock", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")
		tc := caseV1Compat{Name: "EachItemSplitBy", Expected: 4, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_ConcatNew_NoAdding(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_ConcatNew_NoAdding", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.ConcatNew(0)
		tc := caseV1Compat{Name: "ConcatNew no adding", Expected: 1, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_ConcatNew_WithAdding(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_ConcatNew_WithAdding", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.ConcatNew(0, "b", "c")
		tc := caseV1Compat{Name: "ConcatNew with adding", Expected: 3, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_ToError_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_ToError_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "ToError empty", Expected: true, Actual: c.ToError(",") == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov46_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err"})
		tc := caseV1Compat{Name: "ToDefaultError", Expected: true, Actual: c.ToDefaultError() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// HashsetsCollection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_HashsetsCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_IsEmpty", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC IsEmpty", Expected: true, Actual: hc.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_HasItems", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC HasItems", Expected: true, Actual: hc.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_Add", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC Add", Expected: 1, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AddNonNil(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AddNonNil", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.AddNonNil(nil)
		hc.AddNonNil(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC AddNonNil", Expected: 1, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AddNonEmpty", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.AddNonEmpty(corestr.New.Hashset.Empty())
		hc.AddNonEmpty(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC AddNonEmpty", Expected: 1, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_Adds", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Adds(corestr.New.Hashset.StringsSpreadItems("a"), corestr.New.Hashset.StringsSpreadItems("b"))
		tc := caseV1Compat{Name: "HC Adds", Expected: 2, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_Adds_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_Adds_NilSkip", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Adds(nil)
		tc := caseV1Compat{Name: "HC Adds nil", Expected: 0, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_LastIndex", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		hc.Add(corestr.New.Hashset.StringsSpreadItems("b"))
		tc := caseV1Compat{Name: "HC LastIndex", Expected: 1, Actual: hc.LastIndex(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_StringsList", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		result := hc.StringsList()
		tc := caseV1Compat{Name: "HC StringsList", Expected: 1, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_StringsList_Empty", func() {
		hc := corestr.Empty.HashsetsCollection()
		result := hc.StringsList()
		tc := caseV1Compat{Name: "HC StringsList empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_ListPtr", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC ListPtr", Expected: true, Actual: hc.ListPtr() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_List(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_List", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC List", Expected: 0, Actual: len(hc.List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_ListDirectPtr", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("x"))
		result := hc.ListDirectPtr()
		tc := caseV1Compat{Name: "HC ListDirectPtr", Expected: true, Actual: result != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AddHashsetsCollection", func() {
		hc1 := corestr.Empty.HashsetsCollection()
		hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		hc2 := corestr.Empty.HashsetsCollection()
		hc2.AddHashsetsCollection(hc1)
		tc := caseV1Compat{Name: "HC AddHashsetsCollection", Expected: 1, Actual: hc2.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AddHashsetsCollection_Nil", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.AddHashsetsCollection(nil)
		tc := caseV1Compat{Name: "HC AddHC nil", Expected: 0, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_ConcatNew_NoArgs", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		result := hc.ConcatNew()
		tc := caseV1Compat{Name: "HC ConcatNew no args", Expected: 1, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_ConcatNew_WithArgs(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_ConcatNew_WithArgs", func() {
		hc1 := corestr.Empty.HashsetsCollection()
		hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(corestr.New.Hashset.StringsSpreadItems("b"))
		result := hc1.ConcatNew(hc2)
		tc := caseV1Compat{Name: "HC ConcatNew with args", Expected: 2, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_HasAll", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a", "b"))
		tc := caseV1Compat{Name: "HC HasAll", Expected: true, Actual: hc.HasAll("a", "b"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_HasAll_Empty", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC HasAll empty", Expected: false, Actual: hc.HasAll("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_IsEqual_Same", func() {
		hc1 := corestr.Empty.HashsetsCollection()
		hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		hc2 := corestr.Empty.HashsetsCollection()
		hc2.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC IsEqualPtr same", Expected: true, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_IsEqualPtr_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_IsEqualPtr_BothNil", func() {
		var hc1, hc2 *corestr.HashsetsCollection
		tc := caseV1Compat{Name: "HC IsEqualPtr both nil", Expected: true, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_IsEqualPtr_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_IsEqualPtr_OneNil", func() {
		hc1 := corestr.Empty.HashsetsCollection()
		var hc2 *corestr.HashsetsCollection
		tc := caseV1Compat{Name: "HC IsEqualPtr one nil", Expected: false, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_String", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		tc := caseV1Compat{Name: "HC String", Expected: true, Actual: len(hc.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_String_Empty", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC String empty", Expected: true, Actual: len(hc.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_Join", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		result := hc.Join(",")
		tc := caseV1Compat{Name: "HC Join", Expected: true, Actual: len(result) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_Serialize", func() {
		hc := corestr.Empty.HashsetsCollection()
		hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
		data, err := hc.Serialize()
		tc := caseV1Compat{Name: "HC Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AsJsonContractsBinder", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC AsJsonContractsBinder", Expected: true, Actual: hc.AsJsonContractsBinder() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AsJsoner", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC AsJsoner", Expected: true, Actual: hc.AsJsoner() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AsJsonParseSelfInjector", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC AsJsonParseSelfInjector", Expected: true, Actual: hc.AsJsonParseSelfInjector() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_HashsetsCollection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov46_HashsetsCollection_AsJsonMarshaller", func() {
		hc := corestr.Empty.HashsetsCollection()
		tc := caseV1Compat{Name: "HC AsJsonMarshaller", Expected: true, Actual: hc.AsJsonMarshaller() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CollectionsOfCollection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_CollectionsOfCollection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_IsEmpty", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{Name: "COC IsEmpty", Expected: true, Actual: coc.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_HasItems", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		tc := caseV1Compat{Name: "COC HasItems", Expected: true, Actual: coc.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_Length(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_Length", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		tc := caseV1Compat{Name: "COC Length", Expected: 1, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AllIndividualItemsLength", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))
		tc := caseV1Compat{Name: "COC AllIndividualItemsLength", Expected: 3, Actual: coc.AllIndividualItemsLength(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_Items(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_Items", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		tc := caseV1Compat{Name: "COC Items", Expected: 1, Actual: len(coc.Items()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_List(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_List", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		result := coc.List(0)
		tc := caseV1Compat{Name: "COC List", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_List_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_List_Empty", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		result := coc.List(0)
		tc := caseV1Compat{Name: "COC List empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_ToCollection", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		c := coc.ToCollection()
		tc := caseV1Compat{Name: "COC ToCollection", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AddStrings", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.AddStrings(false, []string{"a", "b"})
		tc := caseV1Compat{Name: "COC AddStrings", Expected: 1, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AddStrings_Empty", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.AddStrings(false, []string{})
		tc := caseV1Compat{Name: "COC AddStrings empty", Expected: 0, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
		tc := caseV1Compat{Name: "COC AddsStringsOfStrings", Expected: 2, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings_Nil", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.AddsStringsOfStrings(false)
		tc := caseV1Compat{Name: "COC AddsStringsOfStrings nil", Expected: 0, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_Adds(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_Adds", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.Adds(c)
		tc := caseV1Compat{Name: "COC Adds", Expected: 1, Actual: coc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_String(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_String", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		tc := caseV1Compat{Name: "COC String", Expected: true, Actual: len(coc.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AsJsonContractsBinder", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{Name: "COC AsJsonContractsBinder", Expected: true, Actual: coc.AsJsonContractsBinder() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AsJsoner", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{Name: "COC AsJsoner", Expected: true, Actual: coc.AsJsoner() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AsJsonParseSelfInjector", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{Name: "COC AsJsonParseSelfInjector", Expected: true, Actual: coc.AsJsonParseSelfInjector() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CollectionsOfCollection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov46_CollectionsOfCollection_AsJsonMarshaller", func() {
		coc := corestr.Empty.CollectionsOfCollection()
		tc := caseV1Compat{Name: "COC AsJsonMarshaller", Expected: true, Actual: coc.AsJsonMarshaller() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CharHashsetMap — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_GetChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM GetChar", Expected: byte('h'), Actual: chm.GetChar("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_GetChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_GetChar_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM GetChar empty", Expected: byte(0), Actual: chm.GetChar(""), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_GetCharOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM GetCharOf", Expected: byte('a'), Actual: chm.GetCharOf("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM Add", Expected: true, Actual: chm.Has("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_Add_SameChar(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_Add_SameChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("hi")
		tc := caseV1Compat{Name: "CHM Add same char", Expected: 1, Actual: chm.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_Has_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_Has_NotFound", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM Has not found", Expected: false, Actual: chm.Has("x"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HasWithHashset_Found(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HasWithHashset_Found", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		found, hs := chm.HasWithHashset("hello")
		tc := caseV1Compat{Name: "CHM HasWithHashset found", Expected: true, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "CHM HasWithHashset hs", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HasWithHashset_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		found, _ := chm.HasWithHashset("x")
		tc := caseV1Compat{Name: "CHM HasWithHashset empty", Expected: false, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("hi")
		tc := caseV1Compat{Name: "CHM LengthOf", Expected: 2, Actual: chm.LengthOf('h'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_LengthOf_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM LengthOf empty", Expected: 0, Actual: chm.LengthOf('x'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("abc")
		tc := caseV1Compat{Name: "CHM AllLengthsSum", Expected: 2, Actual: chm.AllLengthsSum(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_AllLengthsSum_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM AllLengthsSum empty", Expected: 0, Actual: chm.AllLengthsSum(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_List", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("abc")
		result := chm.List()
		tc := caseV1Compat{Name: "CHM List", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("banana")
		chm.Add("apple")
		result := chm.SortedListAsc()
		tc := caseV1Compat{Name: "CHM SortedListAsc first", Expected: "apple", Actual: result[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("banana")
		chm.Add("apple")
		result := chm.SortedListDsc()
		tc := caseV1Compat{Name: "CHM SortedListDsc first", Expected: "banana", Actual: result[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_IsEquals_Same", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm1.Add("hello")
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2.Add("hello")
		tc := caseV1Compat{Name: "CHM IsEquals same", Expected: true, Actual: chm1.IsEquals(chm2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_IsEquals_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM IsEquals nil", Expected: false, Actual: chm.IsEquals(nil), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_IsEquals_BothEmpty", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM IsEquals both empty", Expected: true, Actual: chm1.IsEquals(chm2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_IsEquals_DiffLength", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm1.Add("hello")
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2.Add("hello")
		chm2.Add("abc")
		tc := caseV1Compat{Name: "CHM IsEquals diff length", Expected: false, Actual: chm1.IsEquals(chm2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_GetMap", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		m := chm.GetMap()
		tc := caseV1Compat{Name: "CHM GetMap", Expected: true, Actual: len(m) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_String", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM String", Expected: true, Actual: len(chm.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM SummaryString", Expected: true, Actual: len(chm.SummaryString()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_Print_False(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_Print_False", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Print(false) // should not panic
		tc := caseV1Compat{Name: "CHM Print false", Expected: true, Actual: true, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("hi")
		tc := caseV1Compat{Name: "CHM LengthOfHashsetFromFirstChar", Expected: 2, Actual: chm.LengthOfHashsetFromFirstChar("h"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{"abc", "axy"})
		tc := caseV1Compat{Name: "CHM AddSameStartingCharItems", Expected: 2, Actual: chm.LengthOf('a'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_AddSameStartingCharItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddSameStartingCharItems('a', []string{})
		tc := caseV1Compat{Name: "CHM AddSameStartingCharItems empty", Expected: 0, Actual: chm.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hc := chm.HashsetsCollection()
		tc := caseV1Compat{Name: "CHM HashsetsCollection", Expected: true, Actual: hc.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HashsetsCollection_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hc := chm.HashsetsCollection()
		tc := caseV1Compat{Name: "CHM HashsetsCollection empty", Expected: true, Actual: hc.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("abc")
		hc := chm.HashsetsCollectionByChars('h')
		tc := caseV1Compat{Name: "CHM HashsetsCollectionByChars", Expected: 1, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov46_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov46_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hc := chm.HashsetsCollectionByStringsFirstChar("hello")
		tc := caseV1Compat{Name: "CHM HC ByStringsFirstChar", Expected: 1, Actual: hc.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
