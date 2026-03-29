package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// CharHashsetMap.go — Full coverage (~421 uncovered stmts, 1193 lines)
// =============================================================================

// ── GetChar / GetCharOf ──

func Test_Cov63_CharHashsetMap_GetChar_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetChar_NonEmpty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"char": chm.GetChar("hello")}
		expected := args.Map{"char": byte('h')}
		expected.ShouldBeEqual(t, 0, "GetChar returns first byte", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetChar_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"char": chm.GetChar("")}
		expected := args.Map{"char": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetChar empty returns 0", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharOf_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharOf_NonEmpty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"char": chm.GetCharOf("abc")}
		expected := args.Map{"char": byte('a')}
		expected.ShouldBeEqual(t, 0, "GetCharOf returns first byte", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharOf_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"char": chm.GetCharOf("")}
		expected := args.Map{"char": byte(0)}
		expected.ShouldBeEqual(t, 0, "GetCharOf empty returns 0", actual)
	})
}

// ── IsEmpty / HasItems / Length ──

func Test_Cov63_CharHashsetMap_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmpty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"empty": chm.IsEmpty(), "hasItems": chm.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty on new", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmpty_Nil", func() {
		actual := args.Map{"empty": corestr.Empty.CharHashsetMap().IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty on nil items", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasItems_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasItems_NonEmpty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")
		actual := args.Map{"hasItems": chm.HasItems()}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HasItems after add", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEmptyLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"empty": chm.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_Length(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Length", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")
		chm.Add("world")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Length returns char groups", actual)
	})
}

func Test_Cov63_CharHashsetMap_Length_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Length_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("hello")
		actual := args.Map{"len": chm.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock", actual)
	})
}

// ── Add / AddStrings ──

func Test_Cov63_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		actual := args.Map{"len": chm.Length(), "allLen": chm.AllLengthsSum()}
		expected := args.Map{"len": 1, "allLen": 2}
		expected.ShouldBeEqual(t, 0, "Add same char groups", actual)
	})
}

func Test_Cov63_CharHashsetMap_Add_ExistingChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Add_ExistingChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("apricot")
		actual := args.Map{"allLen": chm.AllLengthsSum()}
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "Add to existing char bucket", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana", "avocado")
		actual := args.Map{"len": chm.Length(), "allLen": chm.AllLengthsSum()}
		expected := args.Map{"len": 2, "allLen": 3}
		expected.ShouldBeEqual(t, 0, "AddStrings", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStrings_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings()
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		chm.AddLock("avocado")
		actual := args.Map{"allLen": chm.AllLengthsSum()}
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddLock_NewChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddLock_NewChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		chm.AddLock("banana")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock new char", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock("apple", "banana")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsLock_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock()
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock nil", actual)
	})
}

// ── AddSameStartingCharItems ──

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})
		actual := args.Map{"allLen": chm.AllLengthsSum()}
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{})
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameStartingCharItems_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.AddSameStartingCharItems('a', []string{"avocado"})
		actual := args.Map{"allLen": chm.AllLengthsSum()}
		expected := args.Map{"allLen": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing", actual)
	})
}

// ── AddCollectionItems / AddCharCollectionMapItems ──

func Test_Cov63_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItems(nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddCharCollectionMapItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCharCollectionMapItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCharCollectionMapItems(nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharCollectionMapItems nil", actual)
	})
}

// ── AddHashsetItems ──

func Test_Cov63_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple", "banana")
		chm.AddHashsetItems(hs)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems", actual)
	})
}

// ── Has ──

func Test_Cov63_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"has": chm.Has("apple"), "miss": chm.Has("banana")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has", actual)
	})
}

func Test_Cov63_CharHashsetMap_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"has": chm.Has("x")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_Has_CharExistsButNotStr(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Has_CharExistsButNotStr", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"has": chm.Has("avocado")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has char exists but not str", actual)
	})
}

// ── HasWithHashset ──

func Test_Cov63_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("apple")
		actual := args.Map{"has": has, "hsNonNil": hs != nil}
		expected := args.Map{"has": true, "hsNonNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset found", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		has, hs := chm.HasWithHashset("x")
		actual := args.Map{"has": has, "hsEmpty": hs.IsEmpty()}
		expected := args.Map{"has": false, "hsEmpty": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashset_MissChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashset_MissChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashset("banana")
		actual := args.Map{"has": has, "hsEmpty": hs.IsEmpty()}
		expected := args.Map{"has": false, "hsEmpty": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashset miss char", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, hs := chm.HasWithHashsetLock("apple")
		actual := args.Map{"has": has, "hsNonNil": hs != nil}
		expected := args.Map{"has": true, "hsNonNil": true}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		has, _ := chm.HasWithHashsetLock("x")
		actual := args.Map{"has": has}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HasWithHashsetLock_MissChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HasWithHashsetLock_MissChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		has, _ := chm.HasWithHashsetLock("banana")
		actual := args.Map{"has": has}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasWithHashsetLock miss char", actual)
	})
}

// ── LengthOf / LengthOfLock / LengthOfHashsetFromFirstChar ──

func Test_Cov63_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		actual := args.Map{"len": chm.LengthOf('a'), "miss": chm.LengthOf('z')}
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOf_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"len": chm.LengthOf('a')}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"len": chm.LengthOfLock('a')}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthOfLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"len": chm.LengthOfLock('a')}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		actual := args.Map{"len": chm.LengthOfHashsetFromFirstChar("abc"), "miss": chm.LengthOfHashsetFromFirstChar("xyz")}
		expected := args.Map{"len": 2, "miss": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfHashsetFromFirstChar", actual)
	})
}

// ── AllLengthsSum / AllLengthsSumLock ──

func Test_Cov63_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana", "avocado")
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSum_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSum_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"sum": chm.AllLengthsSum()}
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSumLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		actual := args.Map{"sum": chm.AllLengthsSumLock()}
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_AllLengthsSumLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AllLengthsSumLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"sum": chm.AllLengthsSumLock()}
		expected := args.Map{"sum": 0}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock empty", actual)
	})
}

// ── GetMap / GetCopyMapLock ──

func Test_Cov63_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetMap", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonNil": chm.GetMap() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		cp := chm.GetCopyMapLock()
		actual := args.Map{"len": len(cp)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCopyMapLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		cp := chm.GetCopyMapLock()
		actual := args.Map{"len": len(cp)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty", actual)
	})
}

// ── List / SortedListAsc / SortedListDsc ──

func Test_Cov63_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_List", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")
		actual := args.Map{"len": len(chm.List())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List", actual)
	})
}

func Test_Cov63_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")
		list := chm.SortedListAsc()
		actual := args.Map{"first": list[0], "second": list[1]}
		expected := args.Map{"first": "apple", "second": "banana"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc", actual)
	})
}

func Test_Cov63_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("banana", "apple")
		list := chm.SortedListDsc()
		actual := args.Map{"first": list[0], "second": list[1]}
		expected := args.Map{"first": "banana", "second": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc", actual)
	})
}

// ── String / SummaryString / StringLock / SummaryStringLock ──

func Test_Cov63_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_String", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String", actual)
	})
}

func Test_Cov63_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.SummaryString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString", actual)
	})
}

func Test_Cov63_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_StringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock", actual)
	})
}

func Test_Cov63_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonEmpty": chm.SummaryStringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock", actual)
	})
}

// ── Print / PrintLock ──

func Test_Cov63_CharHashsetMap_Print_True(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Print_True", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Print(true)
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Print true", actual)
	})
}

func Test_Cov63_CharHashsetMap_Print_False(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Print_False", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Print(false)
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Print false skips", actual)
	})
}

func Test_Cov63_CharHashsetMap_PrintLock_True(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_PrintLock_True", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.PrintLock(true)
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "PrintLock true", actual)
	})
}

func Test_Cov63_CharHashsetMap_PrintLock_False(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_PrintLock_False", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.PrintLock(false)
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "PrintLock false skips", actual)
	})
}

// ── IsEquals / IsEqualsLock ──

func Test_Cov63_CharHashsetMap_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_Same", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"eq": chm.IsEquals(chm)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same ptr", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"eq": chm.IsEquals(nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_BothEmpty", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_OneEmpty", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_DiffLen", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.AddStrings("apple", "banana")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("apple")
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_DiffContent", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("avocado")
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff content", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEquals_MissingKey(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEquals_MissingKey", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("banana")
		actual := args.Map{"eq": a.IsEquals(b)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key", actual)
	})
}

func Test_Cov63_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_IsEqualsLock", func() {
		a := corestr.New.CharHashsetMap.Cap(10, 10)
		a.Add("apple")
		b := corestr.New.CharHashsetMap.Cap(10, 10)
		b.Add("apple")
		actual := args.Map{"eq": a.IsEqualsLock(b)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock", actual)
	})
}

// ── GetHashset / GetHashsetLock ──

func Test_Cov63_CharHashsetMap_GetHashset_Found(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_Found", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.GetHashset("avocado", false)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset found", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashset_MissNoAdd(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_MissNoAdd", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashset("x", false)
		actual := args.Map{"nil": hs == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss no add", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashset_MissAddNew(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashset_MissAddNew", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashset("x", true)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashset miss add new", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.GetHashsetLock(true, "x")
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetLock", actual)
	})
}

// ── GetHashsetByChar / HashsetByChar / HashsetByCharLock ──

func Test_Cov63_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.GetHashsetByChar('a')
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "GetHashsetByChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByChar('a')
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByCharLock_Found(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByCharLock_Found", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByCharLock('a')
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock found", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByCharLock_Miss(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByCharLock_Miss", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.HashsetByCharLock('z')
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss returns empty", actual)
	})
}

// ── HashsetByStringFirstChar / HashsetByStringFirstCharLock ──

func Test_Cov63_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("avocado")
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.HashsetByStringFirstCharLock("avocado")
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock", actual)
	})
}

// ── HashsetsCollection / HashsetsCollectionByChars / HashsetsCollectionByStringsFirstChar ──

func Test_Cov63_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollection()
		actual := args.Map{"nonNil": hsc != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollection_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollection()
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'b')
		actual := args.Map{"nonNil": hsc != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollectionByChars('a')
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByChars_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hsc := chm.HashsetsCollectionByChars('a', 'z')
		actual := args.Map{"nonNil": hsc != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars skips nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("avocado", "berry")
		actual := args.Map{"nonNil": hsc != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar", actual)
	})
}

func Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_HashsetsCollectionByStringsFirstChar_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hsc := chm.HashsetsCollectionByStringsFirstChar("x")
		actual := args.Map{"empty": hsc.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringsFirstChar empty", actual)
	})
}

// ── GetCharsGroups ──

func Test_Cov63_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.GetCharsGroups("apple", "banana", "avocado")
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups", actual)
	})
}

func Test_Cov63_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_GetCharsGroups_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.GetCharsGroups()
		actual := args.Map{"same": r == chm}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty returns self", actual)
	})
}

// ── AddSameCharsCollection ──

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		hs := chm.AddSameCharsCollection("a", col)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollection("a", col)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_NilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.AddSameCharsCollection("a", nil)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollection_ExistingNilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollection_ExistingNilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.AddSameCharsCollection("a", nil)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing nil col", actual)
	})
}

// ── AddSameCharsHashset ──

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple")
		r := chm.AddSameCharsHashset("a", hs)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := corestr.New.Hashset.StringsSpreadItems("avocado")
		r := chm.AddSameCharsHashset("a", hs)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.AddSameCharsHashset("a", nil)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsHashset_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsHashset_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.AddSameCharsHashset("a", nil)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsHashset existing nil", actual)
	})
}

// ── AddSameCharsCollectionLock ──

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		hs := chm.AddSameCharsCollectionLock("a", col)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		col := corestr.New.Collection.Strings([]string{"avocado"})
		hs := chm.AddSameCharsCollectionLock("a", col)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_NilCol", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := chm.AddSameCharsCollectionLock("a", nil)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock nil", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddSameCharsCollectionLock_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := chm.AddSameCharsCollectionLock("a", nil)
		actual := args.Map{"nonNil": hs != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil", actual)
	})
}

// ── AddHashsetLock ──

func Test_Cov63_CharHashsetMap_AddHashsetLock_New(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.StringsSpreadItems("apple")
		r := chm.AddHashsetLock("a", hs)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock new", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_Existing(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		hs := corestr.New.Hashset.StringsSpreadItems("avocado")
		r := chm.AddHashsetLock("a", hs)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm.AddHashsetLock("a", nil)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock nil creates empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddHashsetLock_ExistingNil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetLock_ExistingNil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.AddHashsetLock("a", nil)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetLock existing nil", actual)
	})
}

// ── JSON ──

func Test_Cov63_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonModel", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		actual := args.Map{"nonNil": chm.JsonModel() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonModelAny", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"nonNil": chm.JsonModelAny() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
	})
}

func Test_Cov63_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_MarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		b, err := chm.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
	})
}

func Test_Cov63_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_UnmarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		b, _ := chm.MarshalJSON()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.UnmarshalJSON(b)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
	})
}

func Test_Cov63_CharHashsetMap_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_UnmarshalJSON_Error", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm.UnmarshalJSON([]byte("invalid"))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON error", actual)
	})
}

func Test_Cov63_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Json", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.Json()
		actual := args.Map{"nonEmpty": r.JsonString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonPtr", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		r := chm.JsonPtr()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJson", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		r, err := chm2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "nonNil": r != nil}
		expected := args.Map{"noErr": true, "nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJson_Error", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := chm.ParseInjectUsingJson(jr)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson error", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		r := chm2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust", actual)
	})
}

func Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_ParseInjectUsingJsonMust_Panics", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			chm.ParseInjectUsingJsonMust(jr)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics", actual)
	})
}

func Test_Cov63_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_JsonParseSelfInject", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
	})
}

// ── Interface casts ──

func Test_Cov63_CharHashsetMap_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsoner", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"nonNil": chm.AsJsoner() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsoner", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonContractsBinder", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"nonNil": chm.AsJsonContractsBinder() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonMarshaller", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"nonNil": chm.AsJsonMarshaller() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller", actual)
	})
}

func Test_Cov63_CharHashsetMap_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AsJsonParseSelfInjector", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"nonNil": chm.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector", actual)
	})
}

// ── RemoveAll / Clear ──

func Test_Cov63_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		chm.RemoveAll()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll", actual)
	})
}

func Test_Cov63_CharHashsetMap_RemoveAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_RemoveAll_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.RemoveAll()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll on empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Clear", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		chm.Clear()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear", actual)
	})
}

func Test_Cov63_CharHashsetMap_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Clear_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Clear()
		actual := args.Map{"empty": chm.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empty", actual)
	})
}

// ── Constructors ──

func Test_Cov63_CharHashsetMap_Cap_MinEnforced(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Cap_MinEnforced", func() {
		chm := corestr.New.CharHashsetMap.Cap(1, 1)
		actual := args.Map{"nonNil": chm != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Cap enforces minimum", actual)
	})
}

func Test_Cov63_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_CapItems", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CapItems", actual)
	})
}

func Test_Cov63_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Strings", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana"})
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Strings", actual)
	})
}

func Test_Cov63_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_Strings_Nil", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, nil)
		actual := args.Map{"nonNil": chm != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Strings nil", actual)
	})
}

// ── AddStringsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Small(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Small", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		called := false
		chm.AddStringsAsyncLock([]string{"apple", "banana"}, func(c *corestr.CharHashsetMap) {
			called = true
		})
		actual := args.Map{"called": called, "len": chm.Length()}
		expected := args.Map{"called": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock small", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsAsyncLock([]string{}, nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock empty", actual)
	})
}

func Test_Cov63_CharHashsetMap_AddStringsAsyncLock_NilCallback(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddStringsAsyncLock_NilCallback", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsAsyncLock([]string{"apple"}, nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsAsyncLock nil callback", actual)
	})
}

// ── AddCollectionItemsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddCollectionItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddCollectionItemsAsyncLock_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItemsAsyncLock(nil, nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItemsAsyncLock nil", actual)
	})
}

// ── AddHashsetItemsAsyncLock ──

func Test_Cov63_CharHashsetMap_AddHashsetItemsAsyncLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov63_CharHashsetMap_AddHashsetItemsAsyncLock_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddHashsetItemsAsyncLock(nil, nil)
		actual := args.Map{"len": chm.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItemsAsyncLock nil", actual)
	})
}
