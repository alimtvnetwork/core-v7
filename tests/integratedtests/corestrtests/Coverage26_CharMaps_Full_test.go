package corestrtests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =======================================================
// CharCollectionMap
// =======================================================

func Test_C26_CharCollectionMap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Empty", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmpty() {
			// Empty() creates an empty map — IsEmpty() should return true
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected empty internal state", actual)
		}
		actual := args.Map{"result": ccm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_CapSelfCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_CapSelfCap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(20, 15)
		actual := args.Map{"result": ccm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Items(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana", "avocado"})
		actual := args.Map{"length": ccm.Length()}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap.Items returns 2 -- a and b groups", actual)
	})
}

func Test_C26_CharCollectionMap_Items_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items_Empty", func() {
		ccm := corestr.New.CharCollectionMap.Items(nil)
		actual := args.Map{"result": ccm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap", func() {
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"abc", "aef"})
		actual := args.Map{"result": ccm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty", func() {
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, nil)
		actual := args.Map{"result": ccm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetChar", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		ch := ccm.GetChar("hello")
		actual := args.Map{"result": ch != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		ch = ccm.GetChar("")
		actual := args.Map{"result": ch != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Add", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.Add("apple")
		ccm.Add("avocado")
		ccm.Add("banana")
		actual := args.Map{"result": ccm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddLock("apple")
		actual := args.Map{"result": ccm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddStrings", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddStrings("apple", "banana", "cherry")
		actual := args.Map{"result": ccm.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameStartingCharItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddSameStartingCharItems('a', []string{"apple", "avocado"}, false)
		actual := args.Map{"result": ccm.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// add more to existing
		ccm.AddSameStartingCharItems('a', []string{"apricot"}, false)
		actual := args.Map{"result": ccm.LengthOf('a') != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Has", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		actual := args.Map{"result": ccm.Has("apple")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": ccm.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C26_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollection("apple")
		actual := args.Map{"result": has || col == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
		has, _ = ccm.HasWithCollection("xyz")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find xyz", actual)
	})
}

func Test_C26_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollectionLock("apple")
		actual := args.Map{"result": has || col == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOf", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})
		actual := args.Map{"result": ccm.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": ccm.LengthOf('z') != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm.LengthOfLock('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})
		actual := args.Map{"result": ccm.LengthOfCollectionFromFirstChar("apple") != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": ccm.LengthOfCollectionFromFirstChar("xyz") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSum", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})
		actual := args.Map{"result": ccm.AllLengthsSum() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSumLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm.AllLengthsSumLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEmptyLock", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmptyLock() {
			// items map is non-nil but empty, so IsEmpty checks len
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected empty", actual)
		}
	})
}

func Test_C26_CharCollectionMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasItems", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		actual := args.Map{"result": ccm1.IsEquals(ccm2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsLock", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm1.IsEqualsLock(ccm2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitive", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm1.IsEqualsCaseSensitive(true, ccm2)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be equal case-sensitive (different first chars)", actual)
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		actual := args.Map{"result": ccm1.IsEqualsCaseSensitiveLock(true, ccm2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals_Nil", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		actual := args.Map{"result": ccm.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not equal nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetMap", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetMap()
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCopyMapLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetCopyMapLock()
		actual := args.Map{"result": m == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollection("abc", false)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find 'a' bucket", actual)
		col2 := ccm.GetCollection("xyz", true)
		actual := args.Map{"result": col2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create on empty", actual)
		col3 := ccm.GetCollection("zzz", false)
		actual := args.Map{"result": col3 != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not create", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionLock("abc", false)
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionByChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionByChar('a')
		actual := args.Map{"result": col == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(col)
		actual := args.Map{"result": ccm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddCollectionItems(nil)
		actual := args.Map{"result": ccm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsValues", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "apple")
		hm.AddOrUpdate("k2", "banana")
		ccm.AddHashmapsValues(hm)
		actual := args.Map{"result": ccm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)
		actual := args.Map{"result": ccm.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			hm,
		)
		actual := args.Map{"result": ccm.Length() < 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCharHashsetMap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		ccm.AddCharHashsetMap(chm)
		actual := args.Map{"result": ccm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := ccm.AddSameCharsCollection("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_Existing", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollection("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.AddSameCharsCollection("abc", nil)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create new empty collection", actual)
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := ccm.AddSameCharsCollectionLock("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Resize", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Resize(100)
		actual := args.Map{"result": ccm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should keep items", actual)
	})
}

func Test_C26_CharCollectionMap_AddLength(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddLength", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.AddLength(10, 20)
		// just verifying no panic
	})
}

func Test_C26_CharCollectionMap_List(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_List", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		list := ccm.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ListLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		list := ccm.ListLock()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SortedListAsc", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"cherry", "apple", "banana"})
		list := ccm.SortedListAsc()
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharCollectionMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCharsGroups", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.GetCharsGroups([]string{"apple", "banana"})
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByChar('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByCharLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByCharLock('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstChar("abc")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstCharLock("abc")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollection()
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByChars", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByChars('a', 'b')
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByStringFirstChar("abc", "bcd")
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryString", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryStringLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryStringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_String", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_StringLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharCollectionMap_Print(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Print", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Print(false) // no output
		ccm.Print(true)  // prints
	})
}

func Test_C26_CharCollectionMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_PrintLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.PrintLock(false)
		ccm.PrintLock(true)
	})
}

func Test_C26_CharCollectionMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonModel", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		model := ccm.JsonModel()
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonModelAny", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		actual := args.Map{"result": ccm.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_MarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, err := json.Marshal(ccm)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C26_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_UnmarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, _ := json.Marshal(ccm)
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := json.Unmarshal(data, ccm2)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Json", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		result := ccm.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C26_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonPtr", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ptr := ccm.JsonPtr()
		actual := args.Map{"result": ptr == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJson", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		_, err := ccm2.ParseInjectUsingJson(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJsonMust", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm2.ParseInjectUsingJsonMust(&jsonResult)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonParseSelfInject", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := ccm2.JsonParseSelfInject(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharCollectionMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AsJsonInterfaces", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		actual := args.Map{"result": ccm.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": ccm.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Clear", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Clear()
		actual := args.Map{"result": ccm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharCollectionMap_Dispose(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Dispose", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Dispose()
	})
}

func Test_C26_CharCollectionMap_DataModel(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_DataModel", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		dm := corestr.NewCharCollectionMapDataModelUsing(ccm)
		actual := args.Map{"result": dm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		ccm2 := corestr.NewCharCollectionMapUsingDataModel(dm)
		actual := args.Map{"result": ccm2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

// =======================================================
// CharHashsetMap
// =======================================================

func Test_C26_CharHashsetMap_Cap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Cap", func() {
		chm := corestr.New.CharHashsetMap.Cap(20, 10)
		actual := args.Map{"result": chm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_CapItems", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"result": chm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana", "avocado"})
		actual := args.Map{"result": chm.AllLengthsSum() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings_Nil", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, nil)
		actual := args.Map{"result": chm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"result": chm.GetChar("hello") != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": chm.GetChar("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"result": chm.GetCharOf("hello") != 'h'}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected h", actual)
		actual := args.Map{"result": chm.GetCharOf("") != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C26_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		chm.Add("banana")
		actual := args.Map{"result": chm.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		actual := args.Map{"result": chm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock("apple", "banana")
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})
		actual := args.Map{"result": chm.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// Add to existing
		chm.AddSameStartingCharItems('a', []string{"apricot"})
		actual := args.Map{"result": chm.LengthOf('a') != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C26_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Has", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"result": chm.Has("apple")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have apple", actual)
		actual := args.Map{"result": chm.Has("cherry")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have cherry", actual)
	})
}

func Test_C26_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashset", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashset("apple")
		actual := args.Map{"result": has || hs == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
		has, _ = chm.HasWithHashset("xyz")
		actual := args.Map{"result": has}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not find xyz", actual)
	})
}

func Test_C26_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashsetLock("apple")
		actual := args.Map{"result": has || hs == nil}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should find apple", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")
		actual := args.Map{"result": chm.LengthOf('a') != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.LengthOfLock('a') != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")
		actual := args.Map{"result": chm.LengthOfHashsetFromFirstChar("abc") != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSumLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.AllLengthsSumLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C26_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEmptyLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		actual := args.Map{"result": chm.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be empty", actual)
	})
}

func Test_C26_CharHashsetMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasItems", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have items", actual)
	})
}

func Test_C26_CharHashsetMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals", func() {
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		actual := args.Map{"result": chm1.IsEquals(chm2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEqualsLock", func() {
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm1.IsEqualsLock(chm2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C26_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals_Nil", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		actual := args.Map{"result": chm.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not equal nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetMap", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.GetMap() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		actual := args.Map{"result": chm.GetCopyMapLock() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashset", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashset("abc", false)
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find 'a' bucket", actual)
		hs2 := chm.GetHashset("xyz", true)
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should create on empty", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetLock(false, "abc")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetByChar('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByChar('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByCharLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByCharLock('a')
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		hs2 := chm.HashsetByCharLock('z')
		actual := args.Map{"result": hs2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return empty, not nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstChar("abc")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should find", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstCharLock("abc")
		actual := args.Map{"result": hs == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItems(nil)
		actual := args.Map{"result": chm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCharCollectionMapItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		chm.AddCharCollectionMapItems(ccm)
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		chm.AddHashsetItems(hs)
		actual := args.Map{"result": chm.AllLengthsSum() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := chm.AddSameCharsCollection("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddSameCharsHashset("abc", hs)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddHashsetLock("abc", hs)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollectionLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := chm.AddSameCharsCollectionLock("abc", col)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollection()
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'b')
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("abc", "bcd")
		actual := args.Map{"result": hsc == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm.GetCharsGroups("apple", "banana")
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_List", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := chm.List()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C26_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "cherry", "apple", "banana")
		list := chm.SortedListAsc()
		actual := args.Map{"result": len(list) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": list[0] != "apple"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be apple", actual)
	})
}

func Test_C26_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "cherry")
		list := chm.SortedListDsc()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual := args.Map{"result": list[0] != "cherry"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "first should be cherry", actual)
	})
}

func Test_C26_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryStringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_String", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_StringLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.StringLock()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	})
}

func Test_C26_CharHashsetMap_Print(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Print", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.Print(false)
		chm.Print(true)
	})
}

func Test_C26_CharHashsetMap_PrintLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_PrintLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

func Test_C26_CharHashsetMap_JsonModel(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonModel", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		model := chm.JsonModel()
		actual := args.Map{"result": model == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonModelAny", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		actual := args.Map{"result": chm.JsonModelAny() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_MarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, err := json.Marshal(chm)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should have data", actual)
	})
}

func Test_C26_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_UnmarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, _ := json.Marshal(chm)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := json.Unmarshal(data, chm2)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Json", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		result := chm.Json()
		actual := args.Map{"result": result.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not error", actual)
	})
}

func Test_C26_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonPtr", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		ptr := chm.JsonPtr()
		actual := args.Map{"result": ptr == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJson", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		_, err := chm2.ParseInjectUsingJson(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm2.ParseInjectUsingJsonMust(&jsonResult)
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonParseSelfInject", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.JsonParseSelfInject(&jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C26_CharHashsetMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AsJsonInterfaces", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		actual := args.Map{"result": chm.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		actual := args.Map{"result": chm.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.RemoveAll()
		time.Sleep(10 * time.Millisecond) // allow goroutine cleanup
		actual := args.Map{"result": chm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Clear", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.Clear()
		time.Sleep(10 * time.Millisecond)
		actual := args.Map{"result": chm.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be 0", actual)
	})
}

func Test_C26_CharHashsetMap_DataModel(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_DataModel", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		dm := corestr.NewCharHashsetMapDataModelUsing(chm)
		actual := args.Map{"result": dm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
		chm2 := corestr.NewCharHashsetMapUsingDataModel(dm)
		actual := args.Map{"result": chm2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	})
}

func Test_C26_CharHashsetMap_AddCollectionItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItemsAsyncLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		done := make(chan bool, 1)
		chm.AddCollectionItemsAsyncLock(col, func(charHashset *corestr.CharHashsetMap) {
			done <- true
		})
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "timeout", actual)
		}
	})
}

func Test_C26_CharHashsetMap_AddHashsetItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetItemsAsyncLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		done := make(chan bool, 1)
		chm.AddHashsetItemsAsyncLock(hs, func(charHashset *corestr.CharHashsetMap) {
			done <- true
		})
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "timeout", actual)
		}
	})
}
