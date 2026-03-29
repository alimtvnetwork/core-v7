package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — Segment 6b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg6_CHM_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEmpty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"empty": chm.IsEmpty(), "hasItems": chm.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg6_CHM_Add(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado").Add("banana")
		actual := args.Map{"len": chm.Length(), "allLen": chm.AllLengthsSum()}
		expected := args.Map{"len": 2, "allLen": 3}
		expected.ShouldBeEqual(t, 0, "Add -- 2 groups 3 total", actual)
	})
}

func Test_Seg6_CHM_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddLock("apple").AddLock("banana")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStrings("apple", "avocado", "banana")
		actual := args.Map{"len": chm.Length(), "sum": chm.AllLengthsSum()}
		expected := args.Map{"len": 2, "sum": 3}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 groups 3 total", actual)
	})
}

func Test_Seg6_CHM_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStrings_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStrings()
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg6_CHM_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStringsLock("apple", "banana")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddStringsLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddStringsLock()
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock empty -- no change", actual)
	})
}

func Test_Seg6_CHM_GetChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"char": chm.GetChar("abc"), "empty": chm.GetChar("")}
		expected := args.Map{"char": byte('a'), "empty": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetChar -- first char or empty", actual)
	})
}

func Test_Seg6_CHM_GetCharOf(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"char": chm.GetCharOf("abc"), "empty": chm.GetCharOf("")}
		expected := args.Map{"char": byte('a'), "empty": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetCharOf -- first char or empty", actual)
	})
}

func Test_Seg6_CHM_Has(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Has", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"has": chm.Has("apple"), "miss": chm.Has("banana")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_Seg6_CHM_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Has_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"has": chm.Has("apple")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": true, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset -- found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset_Miss", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("banana")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": false, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset miss -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashset_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		has, hs := chm.HasWithHashset("apple")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": false, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset empty -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": true, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock -- found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		has, hs := chm.HasWithHashsetLock("apple")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": false, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock empty -- not found", actual)
	})
}

func Test_Seg6_CHM_HasWithHashsetLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HasWithHashsetLock_Miss", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("banana")
		actual := args.Map{"has": has, "notNil": hs != nil}
		expected := args.Map{"has": false, "notNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock miss -- not found", actual)
	})
}

func Test_Seg6_CHM_LengthOf(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado")
		actual := args.Map{"lenA": chm.LengthOf(byte('a')), "lenZ": chm.LengthOf(byte('z'))}
		expected := args.Map{"lenA": 2, "lenZ": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf -- 2 for a, 0 for z", actual)
	})
}

func Test_Seg6_CHM_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOf_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"len": chm.LengthOf(byte('a'))}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty -- 0", actual)
	})
}

func Test_Seg6_CHM_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"len": chm.LengthOfLock(byte('a')), "miss": chm.LengthOfLock(byte('z'))}
		expected := args.Map{"len": 1, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock -- found and missing", actual)
	})
}

func Test_Seg6_CHM_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"len": chm.LengthOfLock(byte('a'))}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("avocado")
		actual := args.Map{"len": chm.LengthOfHashsetFromFirstChar("abc"), "miss": chm.LengthOfHashsetFromFirstChar("xyz")}
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfHashsetFromFirstChar -- 2 and 0", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum -- 2", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSum_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum empty -- 0", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSumLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"sum": chm.AllLengthsSumLock()}
		expected := args.Map{"sum": 1}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock -- 1", actual)
	})
}

func Test_Seg6_CHM_AllLengthsSumLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AllLengthsSumLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"sum": chm.AllLengthsSumLock()}
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEmptyLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"empty": chm.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg6_CHM_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_LengthLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"len": chm.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg6_CHM_GetMap(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetMap", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"notNil": chm.GetMap() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap -- not nil", actual)
	})
}

func Test_Seg6_CHM_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"len": len(chm.GetCopyMapLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock -- 1", actual)
	})
}

func Test_Seg6_CHM_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCopyMapLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"len": len(chm.GetCopyMapLock())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty -- 0", actual)
	})
}

func Test_Seg6_CHM_List(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_List", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		actual := args.Map{"len": len(chm.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List -- 2 items", actual)
	})
}

func Test_Seg6_CHM_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("banana").Add("apple")
		sorted := chm.SortedListAsc()
		actual := args.Map{"first": sorted[0]}
		expected := args.Map{"first": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- sorted", actual)
	})
}

func Test_Seg6_CHM_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		sorted := chm.SortedListDsc()
		actual := args.Map{"first": sorted[0]}
		expected := args.Map{"first": "banana"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc -- descending", actual)
	})
}

func Test_Seg6_CHM_String(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_String", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg6_CHM_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_StringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg6_CHM_SummaryString(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.SummaryString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_Seg6_CHM_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.SummaryStringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock -- non-empty", actual)
	})
}

func Test_Seg6_CHM_Print(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Print", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.Print(false)
		chm.Print(true)
	})
}

func Test_Seg6_CHM_PrintLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_PrintLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_Seg6_CHM_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple")
		actual := args.Map{
			"eq":   chm1.IsEquals(chm2),
			"same": chm1.IsEquals(chm1),
			"nil":  chm1.IsEquals(nil),
		}
		expected := args.Map{"eq": true, "same": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Seg6_CHM_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_BothEmpty", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(0, 4)
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"eq": chm1.IsEquals(chm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Seg6_CHM_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_OneEmpty", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		actual := args.Map{"eq": chm1.IsEquals(chm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_DiffLen", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple").Add("banana")
		actual := args.Map{"eq": chm1.IsEquals(chm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_DiffItems", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("avocado")
		actual := args.Map{"eq": chm1.IsEquals(chm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff items -- false", actual)
	})
}

func Test_Seg6_CHM_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEquals_MissingKey", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("banana")
		actual := args.Map{"eq": chm1.IsEquals(chm2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key -- false", actual)
	})
}

func Test_Seg6_CHM_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_IsEqualsLock", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm1.Add("apple")
		chm2 := corestr.New.CharHashsetMap.Cap(4, 4)
		chm2.Add("apple")
		actual := args.Map{"eq": chm1.IsEqualsLock(chm2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

// ── AddSameStartingCharItems ────────────────────────────────────────────────

func Test_Seg6_CHM_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddSameStartingCharItems(byte('a'), []string{"apple", "avocado"})
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.AddSameStartingCharItems(byte('a'), []string{"avocado"})
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameStartingCharItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddSameStartingCharItems(byte('a'), []string{})
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty -- no change", actual)
	})
}

// ── AddCollectionItems / AddCharCollectionMapItems ──────────────────────────

func Test_Seg6_CHM_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(c)
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddCollectionItems(nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil -- no change", actual)
	})
}

func Test_Seg6_CHM_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCharCollectionMapItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)
		actual := args.Map{"has": chm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems -- added", actual)
	})
}

func Test_Seg6_CHM_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddCharCollectionMapItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.AddCharCollectionMapItems(nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems nil -- no change", actual)
	})
}

func Test_Seg6_CHM_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		chm.AddHashsetItems(hs)
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems -- 2 items", actual)
	})
}

func Test_Seg6_CHM_AddHashsetItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Empty()
		chm.AddHashsetItems(hs)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems empty -- no change", actual)
	})
}

// ── GetHashset / GetHashsetLock ─────────────────────────────────────────────

func Test_Seg6_CHM_GetHashset(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.GetHashset("a", false)
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset -- found", actual)
	})
}

func Test_Seg6_CHM_GetHashset_AddNew(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset_AddNew", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.GetHashset("z", true)
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset addNew -- created", actual)
	})
}

func Test_Seg6_CHM_GetHashset_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashset_Miss", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.GetHashset("z", false)
		actual := args.Map{"nil": hs == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss -- nil", actual)
	})
}

func Test_Seg6_CHM_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.GetHashsetLock(false, "a")
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock -- found", actual)
	})
}

func Test_Seg6_CHM_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"notNil": chm.GetHashsetByChar(byte('a')) != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetByChar -- found", actual)
	})
}

// ── HashsetByChar / HashsetByStringFirstChar ────────────────────────────────

func Test_Seg6_CHM_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByChar(byte('a'))
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByCharLock(byte('a'))
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByCharLock_Miss(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByCharLock_Miss", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := chm.HashsetByCharLock(byte('z'))
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("avocado")
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar -- found", actual)
	})
}

func Test_Seg6_CHM_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("avocado")
		actual := args.Map{"notNil": hs != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock -- found", actual)
	})
}

// ── HashsetsCollection ──────────────────────────────────────────────────────

func Test_Seg6_CHM_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		hsc := chm.HashsetsCollection()
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection -- 2 hashsets", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollection_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollection()
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple").Add("banana")
		hsc := chm.HashsetsCollectionByChars(byte('a'))
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars -- 1 hashset", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByChars_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollectionByChars(byte('a'))
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty -- empty", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hsc := chm.HashsetsCollectionByStringsFirstChar("avocado")
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar -- 1", actual)
	})
}

func Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_HashsetsCollectionByStringsFirstChar_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		hsc := chm.HashsetsCollectionByStringsFirstChar("a")
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar empty -- empty", actual)
	})
}

// ── GetCharsGroups ──────────────────────────────────────────────────────────

func Test_Seg6_CHM_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.GetCharsGroups("apple", "banana")
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups -- 2 groups", actual)
	})
}

func Test_Seg6_CHM_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_GetCharsGroups_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.GetCharsGroups()
		actual := args.Map{"same": result == chm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty -- returns self", actual)
	})
}

// ── AddSameCharsCollection / AddSameCharsHashset ────────────────────────────

func Test_Seg6_CHM_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollection("a", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NilColl(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NilColl", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsCollection("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil -- returns existing", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NewChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := chm.AddSameCharsCollection("b", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollection_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollection_NewNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsCollection("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddSameCharsHashset("a", hs)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_NewChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"banana"})
		result := chm.AddSameCharsHashset("b", hs)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_NewNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsHashset("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsHashset_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsHashset_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsHashset("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing nil -- returns existing", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := chm.AddSameCharsCollectionLock("a", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock -- merged", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_NewNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddSameCharsCollectionLock("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new nil -- created", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_NewChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := chm.AddSameCharsCollectionLock("b", c)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new char -- added", actual)
	})
}

func Test_Seg6_CHM_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddSameCharsCollectionLock_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddSameCharsCollectionLock("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil -- returned", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		hs := corestr.New.Hashset.Strings([]string{"avocado"})
		result := chm.AddHashsetLock("a", hs)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing -- merged", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_NewChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		hs := corestr.New.Hashset.Strings([]string{"banana"})
		result := chm.AddHashsetLock("b", hs)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new char -- added", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_NewNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_NewNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		result := chm.AddHashsetLock("z", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new nil -- created empty", actual)
	})
}

func Test_Seg6_CHM_AddHashsetLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_AddHashsetLock_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		result := chm.AddHashsetLock("a", nil)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing nil -- returned", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg6_CHM_Json(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Json", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		j := chm.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg6_CHM_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_MarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		b, err := chm.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg6_CHM_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_UnmarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		b, _ := chm.MarshalJSON()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm2.UnmarshalJSON(b)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg6_CHM_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_UnmarshalJSON_Invalid", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg6_CHM_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonModel", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		actual := args.Map{"notNil": chm.JsonModel() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_Seg6_CHM_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonModelAny", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		actual := args.Map{"notNil": chm.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg6_CHM_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_InterfaceCasts", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		actual := args.Map{
			"jsoner":   chm.AsJsoner() != nil,
			"binder":   chm.AsJsonContractsBinder() != nil,
			"injector": chm.AsJsonParseSelfInjector() != nil,
			"marsh":    chm.AsJsonMarshaller() != nil,
		}
		expected := args.Map{"jsoner": true, "binder": true, "injector": true, "marsh": true}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg6_CHM_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_ParseInjectUsingJson", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		_, err := chm2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg6_CHM_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_ParseInjectUsingJsonMust", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg6_CHM_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_JsonParseSelfInject", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(0, 4)
		err := chm2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / RemoveAll ───────────────────────────────────────────────────────

func Test_Seg6_CHM_Clear(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Clear", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.Clear()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg6_CHM_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_Clear_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm.Clear()
		actual := args.Map{"same": result == chm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_Seg6_CHM_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		chm.RemoveAll()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll -- emptied", actual)
	})
}

func Test_Seg6_CHM_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Seg6_CHM_RemoveAll_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		result := chm.RemoveAll()
		actual := args.Map{"same": result == chm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll empty -- returns self", actual)
	})
}
