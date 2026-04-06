package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// CharCollectionMap
// ═══════════════════════════════════════════════════════════════

func Test_Cov52_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Add", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("hello")
		tc := caseV1Compat{Name: "CCM Add", Expected: true, Actual: ccm.Has("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddLock("world")
		tc := caseV1Compat{Name: "CCM AddLock", Expected: true, Actual: ccm.Has("world"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddStrings", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddStrings("abc", "axy", "bcd")
		tc := caseV1Compat{Name: "CCM AddStrings", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddStrings_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddStrings()
		tc := caseV1Compat{Name: "CCM AddStrings empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEmpty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM IsEmpty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HasItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("x")
		tc := caseV1Compat{Name: "CCM HasItems", Expected: true, Actual: ccm.HasItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEmptyLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM IsEmptyLock", Expected: true, Actual: ccm.IsEmptyLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Length(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Length", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Add("bcd")
		tc := caseV1Compat{Name: "CCM Length", Expected: 2, Actual: ccm.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_LengthLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM LengthLock", Expected: 1, Actual: ccm.LengthLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Has_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Has_NotFound", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM Has not found", Expected: false, Actual: ccm.Has("z"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HasWithCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		found, col := ccm.HasWithCollection("abc")
		tc := caseV1Compat{Name: "CCM HasWithCollection", Expected: true, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "CCM HasWithCollection col", Expected: true, Actual: col != nil, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HasWithCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HasWithCollection_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		found, _ := ccm.HasWithCollection("z")
		tc := caseV1Compat{Name: "CCM HasWithCollection empty", Expected: false, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HasWithCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		found, _ := ccm.HasWithCollectionLock("abc")
		tc := caseV1Compat{Name: "CCM HasWithCollectionLock", Expected: true, Actual: found, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_LengthOf", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Add("axy")
		tc := caseV1Compat{Name: "CCM LengthOf", Expected: 2, Actual: ccm.LengthOf('a'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_LengthOf_Missing(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_LengthOf_Missing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM LengthOf missing", Expected: 0, Actual: ccm.LengthOf('z'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_LengthOfLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM LengthOfLock", Expected: 1, Actual: ccm.LengthOfLock('a'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM LengthOfCollFromFirstChar", Expected: 1, Actual: ccm.LengthOfCollectionFromFirstChar("axy"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AllLengthsSum", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Add("bcd")
		tc := caseV1Compat{Name: "CCM AllLengthsSum", Expected: 2, Actual: ccm.AllLengthsSum(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AllLengthsSumLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM AllLengthsSumLock", Expected: 1, Actual: ccm.AllLengthsSumLock(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM GetChar", Expected: byte('h'), Actual: ccm.GetChar("hello"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetChar_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM GetChar empty", Expected: byte(0), Actual: ccm.GetChar(""), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetMap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM GetMap", Expected: true, Actual: ccm.GetMap() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCopyMapLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM GetCopyMapLock", Expected: true, Actual: ccm.GetCopyMapLock() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCollection(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		col := ccm.GetCollection("axy", false)
		tc := caseV1Compat{Name: "CCM GetCollection", Expected: true, Actual: col != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCollection_AddNew(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCollection_AddNew", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		col := ccm.GetCollection("xyz", true)
		tc := caseV1Compat{Name: "CCM GetCollection addNew", Expected: true, Actual: col != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCollection_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		col := ccm.GetCollection("xyz", false)
		tc := caseV1Compat{Name: "CCM GetCollection nil", Expected: true, Actual: col == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		col := ccm.GetCollectionLock("axy", false)
		tc := caseV1Compat{Name: "CCM GetCollectionLock", Expected: true, Actual: col != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCollectionByChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		col := ccm.GetCollectionByChar('a')
		tc := caseV1Compat{Name: "CCM GetCollectionByChar", Expected: true, Actual: col != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEquals_Same", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM IsEquals same", Expected: true, Actual: ccm.IsEquals(ccm), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEquals_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM IsEquals nil", Expected: false, Actual: ccm.IsEquals(nil), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEquals_Equal(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEquals_Equal", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm1.Add("abc")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm2.Add("abc")
		tc := caseV1Compat{Name: "CCM IsEquals equal", Expected: true, Actual: ccm1.IsEquals(ccm2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEqualsLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM IsEqualsLock", Expected: true, Actual: ccm.IsEqualsLock(ccm), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_IsEqualsCaseSensitive", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm1.Add("abc")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm2.Add("abc")
		tc := caseV1Compat{Name: "CCM IsEqualsCaseSensitive", Expected: true, Actual: ccm1.IsEqualsCaseSensitive(true, ccm2), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_List(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_List", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Add("bcd")
		tc := caseV1Compat{Name: "CCM List", Expected: 2, Actual: len(ccm.List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_ListLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM ListLock", Expected: 1, Actual: len(ccm.ListLock()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_SortedListAsc", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("bcd")
		ccm.Add("abc")
		list := ccm.SortedListAsc()
		tc := caseV1Compat{Name: "CCM SortedListAsc", Expected: "abc", Actual: list[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_SortedListAsc_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM SortedListAsc empty", Expected: 0, Actual: len(ccm.SortedListAsc()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_String", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM String", Expected: true, Actual: len(ccm.String()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_StringLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM StringLock", Expected: true, Actual: len(ccm.StringLock()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_SummaryString", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM SummaryString", Expected: true, Actual: len(ccm.SummaryString()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_SummaryStringLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM SummaryStringLock", Expected: true, Actual: len(ccm.SummaryStringLock()) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameStartingCharItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddSameStartingCharItems('a', []string{"abc", "axy"}, false)
		tc := caseV1Compat{Name: "CCM AddSameStartingCharItems", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Existing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.AddSameStartingCharItems('a', []string{"axy"}, false)
		tc := caseV1Compat{Name: "CCM AddSameStartingCharItems existing", Expected: 2, Actual: ccm.LengthOf('a'), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameStartingCharItems_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddSameStartingCharItems('a', []string{}, false)
		tc := caseV1Compat{Name: "CCM AddSameStartingCharItems empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddCollectionItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		col := corestr.New.Collection.Strings([]string{"abc", "bcd"})
		ccm.AddCollectionItems(col)
		tc := caseV1Compat{Name: "CCM AddCollectionItems", Expected: 2, Actual: ccm.AllLengthsSum(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddCollectionItems_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddCollectionItems(nil)
		tc := caseV1Compat{Name: "CCM AddCollectionItems nil", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddHashmapsValues", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k", "abc")
		ccm.AddHashmapsValues(hm)
		tc := caseV1Compat{Name: "CCM AddHashmapsValues", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddHashmapsValues_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.AddHashmapsValues(nil)
		tc := caseV1Compat{Name: "CCM AddHashmapsValues nil", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetByChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		hs := ccm.HashsetByChar('a')
		tc := caseV1Compat{Name: "CCM HashsetByChar", Expected: true, Actual: hs != nil && hs.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetByChar_Missing(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetByChar_Missing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		hs := ccm.HashsetByChar('z')
		tc := caseV1Compat{Name: "CCM HashsetByChar missing", Expected: true, Actual: hs == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetByCharLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		hs := ccm.HashsetByCharLock('a')
		tc := caseV1Compat{Name: "CCM HashsetByCharLock", Expected: true, Actual: hs != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		hs := ccm.HashsetByStringFirstChar("axy")
		tc := caseV1Compat{Name: "CCM HashsetByStringFirstChar", Expected: true, Actual: hs != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetsCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Add("bcd")
		hsc := ccm.HashsetsCollection()
		tc := caseV1Compat{Name: "CCM HashsetsCollection", Expected: true, Actual: hsc != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetsCollection_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		hsc := ccm.HashsetsCollection()
		tc := caseV1Compat{Name: "CCM HashsetsCollection empty", Expected: true, Actual: hsc != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetsCollectionByChars", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		hsc := ccm.HashsetsCollectionByChars('a')
		tc := caseV1Compat{Name: "CCM HashsetsCollByChars", Expected: true, Actual: hsc != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		hsc := ccm.HashsetsCollectionByStringFirstChar("abc")
		tc := caseV1Compat{Name: "CCM HashsetsCollByStringFirstChar", Expected: true, Actual: hsc != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Resize", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 2)
		ccm.Add("abc")
		ccm.Resize(10)
		tc := caseV1Compat{Name: "CCM Resize", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Resize_Smaller(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Resize_Smaller", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.Add("abc")
		ccm.Resize(1)
		tc := caseV1Compat{Name: "CCM Resize smaller", Expected: true, Actual: ccm.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddLength(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddLength", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 2)
		ccm.AddLength(5, 3)
		tc := caseV1Compat{Name: "CCM AddLength", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddLength_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddLength_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 2)
		ccm.AddLength()
		tc := caseV1Compat{Name: "CCM AddLength empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_JsonModel", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		dm := ccm.JsonModel()
		tc := caseV1Compat{Name: "CCM JsonModel", Expected: true, Actual: dm != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_JsonModelAny", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		tc := caseV1Compat{Name: "CCM JsonModelAny", Expected: true, Actual: ccm.JsonModelAny() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_MarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		data, err := ccm.MarshalJSON()
		tc := caseV1Compat{Name: "CCM MarshalJSON", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_UnmarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		data, _ := ccm.MarshalJSON()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		err := ccm2.UnmarshalJSON(data)
		tc := caseV1Compat{Name: "CCM UnmarshalJSON", Expected: true, Actual: err == nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Json", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		j := ccm.Json()
		tc := caseV1Compat{Name: "CCM Json", Expected: true, Actual: j.HasSafeItems(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_JsonPtr", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		j := ccm.JsonPtr()
		tc := caseV1Compat{Name: "CCM JsonPtr", Expected: true, Actual: j != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AsJsonMarshaller", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM AsJsonMarshaller", Expected: true, Actual: ccm.AsJsonMarshaller() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AsJsoner", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM AsJsoner", Expected: true, Actual: ccm.AsJsoner() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AsJsonContractsBinder", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM AsJsonContractsBinder", Expected: true, Actual: ccm.AsJsonContractsBinder() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AsJsonParseSelfInjector", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		tc := caseV1Compat{Name: "CCM AsJsonParseSelfInjector", Expected: true, Actual: ccm.AsJsonParseSelfInjector() != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameCharsCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		col := corestr.New.Collection.Strings([]string{"abc", "axy"})
		result := ccm.AddSameCharsCollection("abc", col)
		tc := caseV1Compat{Name: "CCM AddSameCharsCollection", Expected: true, Actual: result != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameCharsCollection_NilCol", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		result := ccm.AddSameCharsCollection("abc", nil)
		tc := caseV1Compat{Name: "CCM AddSameCharsCollection nil col", Expected: true, Actual: result != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_AddSameCharsCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		col := corestr.New.Collection.Strings([]string{"abc"})
		result := ccm.AddSameCharsCollectionLock("abc", col)
		tc := caseV1Compat{Name: "CCM AddSameCharsCollectionLock", Expected: true, Actual: result != nil, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Clear", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Clear()
		tc := caseV1Compat{Name: "CCM Clear", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Clear_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Clear()
		tc := caseV1Compat{Name: "CCM Clear empty", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Dispose", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Dispose()
		tc := caseV1Compat{Name: "CCM Dispose", Expected: true, Actual: ccm.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_Print(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_Print", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.Print(false)
		tc := caseV1Compat{Name: "CCM Print skip", Expected: true, Actual: true, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_PrintLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		ccm.Add("abc")
		ccm.PrintLock(false)
		tc := caseV1Compat{Name: "CCM PrintLock skip", Expected: true, Actual: true, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCharsGroups", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		result := ccm.GetCharsGroups([]string{"abc", "axy", "bcd"})
		tc := caseV1Compat{Name: "CCM GetCharsGroups", Expected: true, Actual: result != nil && result.Has("abc"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov52_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov52_CharCollectionMap_GetCharsGroups_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(5, 5)
		result := ccm.GetCharsGroups([]string{})
		tc := caseV1Compat{Name: "CCM GetCharsGroups empty", Expected: true, Actual: result == ccm, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
