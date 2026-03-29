package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — Segment 6a
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_CCM_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEmpty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"empty": ccm.IsEmpty(), "hasItems": ccm.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg6_CCM_Add(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Add", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado").Add("banana")
		actual := args.Map{"len": ccm.Length(), "allLen": ccm.AllLengthsSum()}
		expected := args.Map{"len": 2, "allLen": 3}
		expected.ShouldBeEqual(t, 0, "Add -- 2 char groups, 3 total items", actual)
	})
}

func Test_Seg6_CCM_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddLock("apple").AddLock("banana")
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock -- 2 char groups", actual)
	})
}

func Test_Seg6_CCM_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddStrings", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddStrings("apple", "avocado", "banana")
		actual := args.Map{"len": ccm.Length(), "allLen": ccm.AllLengthsSum()}
		expected := args.Map{"len": 2, "allLen": 3}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 groups 3 total", actual)
	})
}

func Test_Seg6_CCM_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddStrings_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddStrings()
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg6_CCM_GetChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"char": ccm.GetChar("abc"), "empty": ccm.GetChar("")}
		expected := args.Map{"char": byte('a'), "empty": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetChar -- first char or empty", actual)
	})
}

func Test_Seg6_CCM_Has(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Has", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"has": ccm.Has("apple"), "miss": ccm.Has("banana")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_Seg6_CCM_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Has_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"has": ccm.Has("apple")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_Seg6_CCM_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollection("apple")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": true, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollection -- found", actual)
	})
}

func Test_Seg6_CCM_HasWithCollection_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection_Miss", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollection("banana")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": false, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollection miss -- not found", actual)
	})
}

func Test_Seg6_CCM_HasWithCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		has, coll := ccm.HasWithCollection("apple")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": false, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollection empty -- not found", actual)
	})
}

func Test_Seg6_CCM_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollectionLock("apple")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": true, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock -- found", actual)
	})
}

func Test_Seg6_CCM_HasWithCollectionLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		has, coll := ccm.HasWithCollectionLock("apple")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": false, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock empty -- not found", actual)
	})
}

func Test_Seg6_CCM_HasWithCollectionLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock_Miss", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollectionLock("banana")
		actual := args.Map{"has": has, "collNotNil": coll != nil}
		expected := args.Map{"has": false, "collNotNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock miss -- not found", actual)
	})
}

func Test_Seg6_CCM_LengthOf(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOf", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado")
		actual := args.Map{"lenA": ccm.LengthOf(byte('a')), "lenZ": ccm.LengthOf(byte('z'))}
		expected := args.Map{"lenA": 2, "lenZ": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf -- 2 for 'a', 0 for 'z'", actual)
	})
}

func Test_Seg6_CCM_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOf_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"len": ccm.LengthOf(byte('a'))}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty -- 0", actual)
	})
}

func Test_Seg6_CCM_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"len": ccm.LengthOfLock(byte('a')), "miss": ccm.LengthOfLock(byte('z'))}
		expected := args.Map{"len": 1, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock -- found and missing", actual)
	})
}

func Test_Seg6_CCM_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfLock_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"len": ccm.LengthOfLock(byte('a'))}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty -- 0", actual)
	})
}

func Test_Seg6_CCM_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfCollectionFromFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado")
		actual := args.Map{"len": ccm.LengthOfCollectionFromFirstChar("abc"), "miss": ccm.LengthOfCollectionFromFirstChar("xyz")}
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfCollectionFromFirstChar -- 2 and 0", actual)
	})
}

func Test_Seg6_CCM_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AllLengthsSum", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado").Add("banana")
		actual := args.Map{"sum": ccm.AllLengthsSum()}
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum -- 3", actual)
	})
}

func Test_Seg6_CCM_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AllLengthsSumLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		actual := args.Map{"sum": ccm.AllLengthsSumLock()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock -- 2", actual)
	})
}

func Test_Seg6_CCM_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEmptyLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"empty": ccm.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg6_CCM_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"len": ccm.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg6_CCM_GetMap(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetMap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"notNil": ccm.GetMap() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap -- not nil", actual)
	})
}

func Test_Seg6_CCM_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCopyMapLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"len": len(ccm.GetCopyMapLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock -- 1", actual)
	})
}

func Test_Seg6_CCM_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCopyMapLock_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"len": len(ccm.GetCopyMapLock())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty -- 0", actual)
	})
}

func Test_Seg6_CCM_List(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_List", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		actual := args.Map{"len": len(ccm.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List -- 2 items", actual)
	})
}

func Test_Seg6_CCM_List_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_List_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"len": len(ccm.List())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

func Test_Seg6_CCM_ListLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ListLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"len": len(ccm.ListLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListLock -- 1", actual)
	})
}

func Test_Seg6_CCM_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SortedListAsc", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("banana").Add("apple")
		sorted := ccm.SortedListAsc()
		actual := args.Map{"first": sorted[0]}
		expected := args.Map{"first": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- sorted", actual)
	})
}

func Test_Seg6_CCM_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SortedListAsc_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"len": len(ccm.SortedListAsc())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty -- 0", actual)
	})
}

func Test_Seg6_CCM_String(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_String", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"nonEmpty": ccm.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_CCM_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_StringLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"nonEmpty": ccm.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg6_CCM_SummaryString(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SummaryString", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"nonEmpty": ccm.SummaryString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_Seg6_CCM_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SummaryStringLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"nonEmpty": ccm.SummaryStringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock -- non-empty", actual)
	})
}

func Test_Seg6_CCM_Print(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Print", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Print(false) // skip print
		ccm.Print(true)  // actual print
	})
}

func Test_Seg6_CCM_PrintLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_PrintLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.PrintLock(false)
		ccm.PrintLock(true)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_Seg6_CCM_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")
		actual := args.Map{
			"eq":   ccm1.IsEquals(ccm2),
			"same": ccm1.IsEquals(ccm1),
			"nil":  ccm1.IsEquals(nil),
		}
		expected := args.Map{"eq": true, "same": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Seg6_CCM_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_BothEmpty", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Seg6_CCM_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_OneEmpty", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Seg6_CCM_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_DiffLen", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple").Add("banana")
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Seg6_CCM_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_DiffItems", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("avocado")
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff items -- false", actual)
	})
}

func Test_Seg6_CCM_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_MissingKey", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("banana")
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key -- false", actual)
	})
}

func Test_Seg6_CCM_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsLock", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")
		actual := args.Map{"eq": ccm1.IsEqualsLock(ccm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

func Test_Seg6_CCM_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsCaseSensitive", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")
		actual := args.Map{"eq": ccm1.IsEqualsCaseSensitive(true, ccm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsCaseSensitive -- true", actual)
	})
}

func Test_Seg6_CCM_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsCaseSensitiveLock", func() {
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")
		actual := args.Map{"eq": ccm1.IsEqualsCaseSensitiveLock(true, ccm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsCaseSensitiveLock -- true", actual)
	})
}

// ── AddSameStartingCharItems ────────────────────────────────────────────────

func Test_Seg6_CCM_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddSameStartingCharItems(byte('a'), []string{"apple", "avocado"}, false)
		actual := args.Map{"len": ccm.AllLengthsSum()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems -- 2 items", actual)
	})
}

func Test_Seg6_CCM_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems_Existing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.AddSameStartingCharItems(byte('a'), []string{"avocado"}, false)
		actual := args.Map{"len": ccm.AllLengthsSum()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing -- 2 items", actual)
	})
}

func Test_Seg6_CCM_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddSameStartingCharItems(byte('a'), []string{}, false)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty -- no change", actual)
	})
}

// ── AddHashmapsValues / AddHashmapsKeysValuesBoth ───────────────────────────

func Test_Seg6_CCM_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsValues", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("k", "apple")
		ccm.AddHashmapsValues(hm)
		actual := args.Map{"has": ccm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues -- added values", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsValues_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsValues(nil)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues nil -- no change", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysValuesBoth", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)
		actual := args.Map{"has": ccm.AllLengthsSum() > 0}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth -- added", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsKeysValuesBoth_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysValuesBoth_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsKeysValuesBoth(nil)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth nil -- no change", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("key", "val")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			hm,
		)
		actual := args.Map{"has": ccm.Has("val")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter -- added", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter nil -- no change", actual)
	})
}

func Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Break", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(4)
		hm.Set("a", "1")
		hm.Set("b", "2")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, true
			},
			hm,
		)
		actual := args.Map{"hasItems": ccm.AllLengthsSum() > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter break -- stops", actual)
	})
}

// ── AddCollectionItems / AddCharHashsetMap ──────────────────────────────────

func Test_Seg6_CCM_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCollectionItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(c)
		actual := args.Map{"sum": ccm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems -- 2 items", actual)
	})
}

func Test_Seg6_CCM_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCollectionItems_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddCollectionItems(nil)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil -- no change", actual)
	})
}

func Test_Seg6_CCM_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCharHashsetMap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		ccm.AddCharHashsetMap(chm)
		actual := args.Map{"has": ccm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCharHashsetMap -- added", actual)
	})
}

func Test_Seg6_CCM_AddCharHashsetMap_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCharHashsetMap_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		ccm.AddCharHashsetMap(chm)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharHashsetMap empty -- no change", actual)
	})
}

// ── GetCollection / GetCollectionLock ────────────────────────────────────────

func Test_Seg6_CCM_GetCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		coll := ccm.GetCollection("a", false)
		actual := args.Map{"notNil": coll != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection -- found", actual)
	})
}

func Test_Seg6_CCM_GetCollection_AddNew(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection_AddNew", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		coll := ccm.GetCollection("z", true)
		actual := args.Map{"notNil": coll != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection addNew -- created", actual)
	})
}

func Test_Seg6_CCM_GetCollection_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection_Miss", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		coll := ccm.GetCollection("z", false)
		actual := args.Map{"nil": coll == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection miss -- nil", actual)
	})
}

func Test_Seg6_CCM_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		coll := ccm.GetCollectionLock("a", false)
		actual := args.Map{"notNil": coll != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollectionLock -- found", actual)
	})
}

func Test_Seg6_CCM_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollectionByChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"notNil": ccm.GetCollectionByChar(byte('a')) != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollectionByChar -- found", actual)
	})
}

// ── AddSameCharsCollection ──────────────────────────────────────────────────

func Test_Seg6_CCM_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_Existing", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollection("a", c)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing -- merged", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollection_NilCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NilCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		result := ccm.AddSameCharsCollection("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil -- existing returned", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollection_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NewChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := ccm.AddSameCharsCollection("b", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new char -- added", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollection_NewCharNilColl(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NewCharNilColl", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.AddSameCharsCollection("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new char nil coll -- created empty", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollectionLock("a", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock -- merged", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollectionLock_NewCharNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_NewCharNil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.AddSameCharsCollectionLock("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new nil -- created", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollectionLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_NewChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := ccm.AddSameCharsCollectionLock("b", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new char -- added", actual)
	})
}

func Test_Seg6_CCM_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_ExistingNil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		result := ccm.AddSameCharsCollectionLock("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil -- returned existing", actual)
	})
}

// ── Hashset conversions ─────────────────────────────────────────────────────

func Test_Seg6_CCM_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByChar(byte('a'))
		actual := args.Map{"notNil": hs != nil, "has": hs.Has("apple")}
		expected := args.Map{"notNil": true, "has": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar -- found", actual)
	})
}

func Test_Seg6_CCM_HashsetByChar_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByChar_Miss", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		actual := args.Map{"nil": ccm.HashsetByChar(byte('z')) == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar miss -- nil", actual)
	})
}

func Test_Seg6_CCM_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByCharLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByCharLock(byte('a'))
		actual := args.Map{"has": hs.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock -- found", actual)
	})
}

func Test_Seg6_CCM_HashsetByCharLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByCharLock_Miss", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hs := ccm.HashsetByCharLock(byte('z'))
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss -- empty", actual)
	})
}

func Test_Seg6_CCM_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstChar("avocado")
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar -- found", actual)
	})
}

func Test_Seg6_CCM_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByStringFirstCharLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstCharLock("avocado")
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock -- found", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		hsc := ccm.HashsetsCollection()
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection -- 2 hashsets", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollection_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollection()
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty -- empty", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByChars", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		hsc := ccm.HashsetsCollectionByChars(byte('a'))
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars -- 1 hashset", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByChars_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollectionByChars(byte('a'))
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty -- empty", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hsc := ccm.HashsetsCollectionByStringFirstChar("avocado")
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringFirstChar -- 1", actual)
	})
}

func Test_Seg6_CCM_HashsetsCollectionByStringFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByStringFirstChar_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollectionByStringFirstChar("a")
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringFirstChar empty -- empty", actual)
	})
}

// ── GetCharsGroups ──────────────────────────────────────────────────────────

func Test_Seg6_CCM_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCharsGroups", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.GetCharsGroups([]string{"apple", "banana"})
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups -- 2 groups", actual)
	})
}

func Test_Seg6_CCM_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCharsGroups_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.GetCharsGroups([]string{})
		actual := args.Map{"same": result == ccm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty -- returns self", actual)
	})
}

// ── Resize / AddLength ──────────────────────────────────────────────────────

func Test_Seg6_CCM_Resize(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Resize", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.Add("apple")
		ccm.Resize(10)
		actual := args.Map{"has": ccm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Resize -- preserved items", actual)
	})
}

func Test_Seg6_CCM_Resize_SmallerThanLen(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Resize_SmallerThanLen", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		ccm.Resize(1)
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_Seg6_CCM_AddLength(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLength", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.Add("apple")
		ccm.AddLength(10, 20)
		actual := args.Map{"has": ccm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLength -- preserved", actual)
	})
}

func Test_Seg6_CCM_AddLength_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLength_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.AddLength()
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddLength empty -- no change", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg6_CCM_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Json", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		j := ccm.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_CCM_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_MarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		b, err := ccm.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg6_CCM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_UnmarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		b, _ := ccm.MarshalJSON()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm2.UnmarshalJSON(b)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg6_CCM_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_UnmarshalJSON_Invalid", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg6_CCM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonModel", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		actual := args.Map{"notNil": ccm.JsonModel() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_Seg6_CCM_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonModelAny", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		actual := args.Map{"notNil": ccm.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg6_CCM_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_InterfaceCasts", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		actual := args.Map{
			"jsoner":   ccm.AsJsoner() != nil,
			"binder":   ccm.AsJsonContractsBinder() != nil,
			"injector": ccm.AsJsonParseSelfInjector() != nil,
			"marsh":    ccm.AsJsonMarshaller() != nil,
		}
		expected := args.Map{"jsoner": true, "binder": true, "injector": true, "marsh": true}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_CCM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ParseInjectUsingJson", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		_, err := ccm2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_CCM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ParseInjectUsingJsonMust", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		result := ccm2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg6_CCM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonParseSelfInject", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Seg6_CCM_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Clear", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Clear()
		actual := args.Map{"len": ccm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_CCM_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Clear_Empty", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		result := ccm.Clear()
		actual := args.Map{"same": result == ccm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_Seg6_CCM_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Dispose", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Dispose()
		actual := args.Map{"empty": ccm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Seg6_CCM_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Dispose_Nil", func() {
		var ccm *corestr.CharCollectionMap
		ccm.Dispose() // should not panic
	})
}
