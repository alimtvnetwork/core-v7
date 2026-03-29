package corestrtests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =======================================================
// CharCollectionMap
// =======================================================

func Test_C26_CharCollectionMap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Empty", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmpty() {
			// Empty() creates an empty map — IsEmpty() should return true
			t.Error("expected empty internal state")
		}
		if ccm.Length() != 0 {
			t.Errorf("expected 0 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_CapSelfCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_CapSelfCap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(20, 15)
		if ccm == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_Items(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana", "avocado"})
		if ccm.Length() != 2 { // 'a' and 'b'
			t.Errorf("expected 2 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_Items_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Items_Empty", func() {
		ccm := corestr.New.CharCollectionMap.Items(nil)
		if ccm.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap", func() {
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, []string{"abc", "aef"})
		if ccm.Length() != 1 {
			t.Errorf("expected 1 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ItemsPtrWithCap_Empty", func() {
		ccm := corestr.New.CharCollectionMap.ItemsPtrWithCap(5, 10, nil)
		if ccm == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetChar", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		ch := ccm.GetChar("hello")
		if ch != 'h' {
			t.Errorf("expected h got %c", ch)
		}
		ch = ccm.GetChar("")
		if ch != 0 {
			t.Errorf("expected 0 got %d", ch)
		}
	})
}

func Test_C26_CharCollectionMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Add", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.Add("apple")
		ccm.Add("avocado")
		ccm.Add("banana")
		if ccm.Length() != 2 {
			t.Errorf("expected 2 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddLock("apple")
		if ccm.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharCollectionMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddStrings", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddStrings("apple", "banana", "cherry")
		if ccm.Length() != 3 {
			t.Errorf("expected 3 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameStartingCharItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddSameStartingCharItems('a', []string{"apple", "avocado"}, false)
		if ccm.LengthOf('a') != 2 {
			t.Errorf("expected 2 got %d", ccm.LengthOf('a'))
		}
		// add more to existing
		ccm.AddSameStartingCharItems('a', []string{"apricot"}, false)
		if ccm.LengthOf('a') != 3 {
			t.Errorf("expected 3 got %d", ccm.LengthOf('a'))
		}
	})
}

func Test_C26_CharCollectionMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Has", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		if !ccm.Has("apple") {
			t.Error("should have apple")
		}
		if ccm.Has("cherry") {
			t.Error("should not have cherry")
		}
	})
}

func Test_C26_CharCollectionMap_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollection("apple")
		if !has || col == nil {
			t.Error("should find apple")
		}
		has, _ = ccm.HasWithCollection("xyz")
		if has {
			t.Error("should not find xyz")
		}
	})
}

func Test_C26_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasWithCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		has, col := ccm.HasWithCollectionLock("apple")
		if !has || col == nil {
			t.Error("should find apple")
		}
	})
}

func Test_C26_CharCollectionMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOf", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})
		if ccm.LengthOf('a') != 2 {
			t.Errorf("expected 2 got %d", ccm.LengthOf('a'))
		}
		if ccm.LengthOf('z') != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C26_CharCollectionMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if ccm.LengthOfLock('a') != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthOfCollectionFromFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado"})
		if ccm.LengthOfCollectionFromFirstChar("apple") != 2 {
			t.Error("expected 2")
		}
		if ccm.LengthOfCollectionFromFirstChar("xyz") != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C26_CharCollectionMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSum", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})
		if ccm.AllLengthsSum() != 3 {
			t.Errorf("expected 3 got %d", ccm.AllLengthsSum())
		}
	})
}

func Test_C26_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AllLengthsSumLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if ccm.AllLengthsSumLock() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharCollectionMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_LengthLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if ccm.LengthLock() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharCollectionMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEmptyLock", func() {
		ccm := corestr.New.CharCollectionMap.Empty()
		if !ccm.IsEmptyLock() {
			// items map is non-nil but empty, so IsEmpty checks len
			t.Error("expected empty")
		}
	})
}

func Test_C26_CharCollectionMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HasItems", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if !ccm.HasItems() {
			t.Error("should have items")
		}
	})
}

func Test_C26_CharCollectionMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		if !ccm1.IsEquals(ccm2) {
			t.Error("should be equal")
		}
	})
}

func Test_C26_CharCollectionMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsLock", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if !ccm1.IsEqualsLock(ccm2) {
			t.Error("should be equal")
		}
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitive", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"Apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if ccm1.IsEqualsCaseSensitive(true, ccm2) {
			t.Error("should not be equal case-sensitive (different first chars)")
		}
	})
}

func Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEqualsCaseSensitiveLock", func() {
		ccm1 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm2 := corestr.New.CharCollectionMap.Items([]string{"apple"})
		if !ccm1.IsEqualsCaseSensitiveLock(true, ccm2) {
			t.Error("should be equal")
		}
	})
}

func Test_C26_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_IsEquals_Nil", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		if ccm.IsEquals(nil) {
			t.Error("should not equal nil")
		}
	})
}

func Test_C26_CharCollectionMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetMap", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetMap()
		if m == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCopyMapLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		m := ccm.GetCopyMapLock()
		if m == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_GetCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollection("abc", false)
		if col == nil {
			t.Error("should find 'a' bucket")
		}
		col2 := ccm.GetCollection("xyz", true)
		if col2 == nil {
			t.Error("should create on empty")
		}
		col3 := ccm.GetCollection("zzz", false)
		if col3 != nil {
			t.Error("should not create")
		}
	})
}

func Test_C26_CharCollectionMap_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionLock("abc", false)
		if col == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCollectionByChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := ccm.GetCollectionByChar('a')
		if col == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(col)
		if ccm.Length() != 2 {
			t.Errorf("expected 2 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCollectionItems_Nil", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		ccm.AddCollectionItems(nil)
		if ccm.Length() != 0 {
			t.Error("should be 0")
		}
	})
}

func Test_C26_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsValues", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k1", "apple")
		hm.AddOrUpdate("k2", "banana")
		ccm.AddHashmapsValues(hm)
		if ccm.Length() != 2 {
			t.Errorf("expected 2 got %d", ccm.Length())
		}
	})
}

func Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddHashmapsKeysValuesBoth", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)
		if ccm.Length() < 1 {
			t.Error("should have items")
		}
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
		if ccm.Length() < 1 {
			t.Error("should have items")
		}
	})
}

func Test_C26_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddCharHashsetMap", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		ccm.AddCharHashsetMap(chm)
		if ccm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", ccm.AllLengthsSum())
		}
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := ccm.AddSameCharsCollection("abc", col)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_Existing", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		col := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollection("abc", col)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollection_NilCol", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.AddSameCharsCollection("abc", nil)
		if result == nil {
			t.Error("should create new empty collection")
		}
	})
}

func Test_C26_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AddSameCharsCollectionLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := ccm.AddSameCharsCollectionLock("abc", col)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_Resize(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Resize", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Resize(100)
		if ccm.Length() != 1 {
			t.Error("should keep items")
		}
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
		if len(list) != 2 {
			t.Errorf("expected 2 got %d", len(list))
		}
	})
}

func Test_C26_CharCollectionMap_ListLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ListLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		list := ccm.ListLock()
		if len(list) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharCollectionMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SortedListAsc", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"cherry", "apple", "banana"})
		list := ccm.SortedListAsc()
		if len(list) != 3 {
			t.Errorf("expected 3 got %d", len(list))
		}
	})
}

func Test_C26_CharCollectionMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_GetCharsGroups", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm.GetCharsGroups([]string{"apple", "banana"})
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByChar('a')
		if hs == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByCharLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByCharLock('a')
		if hs == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstChar("abc")
		if hs == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetByStringFirstCharLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		hs := ccm.HashsetByStringFirstCharLock("abc")
		if hs == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollection", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollection()
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByChars", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByChars('a', 'b')
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_HashsetsCollectionByStringFirstChar", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		hsc := ccm.HashsetsCollectionByStringFirstChar("abc", "bcd")
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryString", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryString()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharCollectionMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_SummaryStringLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.SummaryStringLock()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharCollectionMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_String", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.String()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharCollectionMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_StringLock", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		s := ccm.StringLock()
		if s == "" {
			t.Error("should not be empty")
		}
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
		if model == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonModelAny", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		if ccm.JsonModelAny() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_MarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, err := json.Marshal(ccm)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(data) == 0 {
			t.Error("should have data")
		}
	})
}

func Test_C26_CharCollectionMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_UnmarshalJSON", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		data, _ := json.Marshal(ccm)
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := json.Unmarshal(data, ccm2)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharCollectionMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Json", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		result := ccm.Json()
		if result.HasError() {
			t.Error("should not error")
		}
	})
}

func Test_C26_CharCollectionMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonPtr", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ptr := ccm.JsonPtr()
		if ptr == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJson", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		_, err := ccm2.ParseInjectUsingJson(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_ParseInjectUsingJsonMust", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		result := ccm2.ParseInjectUsingJsonMust(&jsonResult)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_JsonParseSelfInject", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		jsonResult := ccm.Json()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(10, 10)
		err := ccm2.JsonParseSelfInject(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharCollectionMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_AsJsonInterfaces", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"a"})
		if ccm.AsJsonContractsBinder() == nil {
			t.Error("should not be nil")
		}
		if ccm.AsJsoner() == nil {
			t.Error("should not be nil")
		}
		if ccm.AsJsonMarshaller() == nil {
			t.Error("should not be nil")
		}
		if ccm.AsJsonParseSelfInjector() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharCollectionMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharCollectionMap_Clear", func() {
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple"})
		ccm.Clear()
		if ccm.Length() != 0 {
			t.Error("should be 0")
		}
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
		if dm == nil {
			t.Error("should not be nil")
		}
		ccm2 := corestr.NewCharCollectionMapUsingDataModel(dm)
		if ccm2 == nil {
			t.Error("should not be nil")
		}
	})
}

// =======================================================
// CharHashsetMap
// =======================================================

func Test_C26_CharHashsetMap_Cap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Cap", func() {
		chm := corestr.New.CharHashsetMap.Cap(20, 10)
		if chm == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_CapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_CapItems", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		if chm.Length() != 2 {
			t.Errorf("expected 2 got %d", chm.Length())
		}
	})
}

func Test_C26_CharHashsetMap_Strings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, []string{"apple", "banana", "avocado"})
		if chm.AllLengthsSum() != 3 {
			t.Errorf("expected 3 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_Strings_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Strings_Nil", func() {
		chm := corestr.New.CharHashsetMap.Strings(10, nil)
		if chm == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_GetChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		if chm.GetChar("hello") != 'h' {
			t.Error("expected h")
		}
		if chm.GetChar("") != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C26_CharHashsetMap_GetCharOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharOf", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		if chm.GetCharOf("hello") != 'h' {
			t.Error("expected h")
		}
		if chm.GetCharOf("") != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C26_CharHashsetMap_Add(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Add", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.Add("apple")
		chm.Add("avocado")
		chm.Add("banana")
		if chm.Length() != 2 {
			t.Errorf("expected 2 got %d", chm.Length())
		}
	})
}

func Test_C26_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddLock("apple")
		if chm.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStrings("apple", "banana")
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddStringsLock("apple", "banana")
		if chm.AllLengthsSum() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C26_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameStartingCharItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddSameStartingCharItems('a', []string{"apple", "avocado"})
		if chm.LengthOf('a') != 2 {
			t.Errorf("expected 2 got %d", chm.LengthOf('a'))
		}
		// Add to existing
		chm.AddSameStartingCharItems('a', []string{"apricot"})
		if chm.LengthOf('a') != 3 {
			t.Errorf("expected 3 got %d", chm.LengthOf('a'))
		}
	})
}

func Test_C26_CharHashsetMap_Has(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Has", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		if !chm.Has("apple") {
			t.Error("should have apple")
		}
		if chm.Has("cherry") {
			t.Error("should not have cherry")
		}
	})
}

func Test_C26_CharHashsetMap_HasWithHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashset", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashset("apple")
		if !has || hs == nil {
			t.Error("should find apple")
		}
		has, _ = chm.HasWithHashset("xyz")
		if has {
			t.Error("should not find xyz")
		}
	})
}

func Test_C26_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		has, hs := chm.HasWithHashsetLock("apple")
		if !has || hs == nil {
			t.Error("should find apple")
		}
	})
}

func Test_C26_CharHashsetMap_LengthOf(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOf", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")
		if chm.LengthOf('a') != 2 {
			t.Errorf("expected 2 got %d", chm.LengthOf('a'))
		}
	})
}

func Test_C26_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if chm.LengthOfLock('a') != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "avocado")
		if chm.LengthOfHashsetFromFirstChar("abc") != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C26_CharHashsetMap_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSum", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AllLengthsSumLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if chm.AllLengthsSumLock() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_LengthLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if chm.LengthLock() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C26_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEmptyLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		if !chm.IsEmptyLock() {
			t.Error("should be empty")
		}
	})
}

func Test_C26_CharHashsetMap_HasItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HasItems", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if !chm.HasItems() {
			t.Error("should have items")
		}
	})
}

func Test_C26_CharHashsetMap_IsEquals(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals", func() {
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		if !chm1.IsEquals(chm2) {
			t.Error("should be equal")
		}
	})
}

func Test_C26_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEqualsLock", func() {
		chm1 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm2 := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if !chm1.IsEqualsLock(chm2) {
			t.Error("should be equal")
		}
	})
}

func Test_C26_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_IsEquals_Nil", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		if chm.IsEquals(nil) {
			t.Error("should not equal nil")
		}
	})
}

func Test_C26_CharHashsetMap_GetMap(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetMap", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if chm.GetMap() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		if chm.GetCopyMapLock() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_GetHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashset", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashset("abc", false)
		if hs == nil {
			t.Error("should find 'a' bucket")
		}
		hs2 := chm.GetHashset("xyz", true)
		if hs2 == nil {
			t.Error("should create on empty")
		}
	})
}

func Test_C26_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetLock(false, "abc")
		if hs == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetHashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.GetHashsetByChar('a')
		if hs == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByChar('a')
		if hs == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByCharLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByCharLock('a')
		if hs == nil {
			t.Error("should not be nil")
		}
		hs2 := chm.HashsetByCharLock('z')
		if hs2 == nil {
			t.Error("should return empty, not nil")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstChar("abc")
		if hs == nil {
			t.Error("should find")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetByStringFirstCharLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		hs := chm.HashsetByStringFirstCharLock("abc")
		if hs == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		chm.AddCollectionItems(nil)
		if chm.Length() != 0 {
			t.Error("should be 0")
		}
	})
}

func Test_C26_CharHashsetMap_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddCharCollectionMapItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		ccm := corestr.New.CharCollectionMap.Items([]string{"apple", "banana"})
		chm.AddCharCollectionMapItems(ccm)
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple", "banana"})
		chm.AddHashsetItems(hs)
		if chm.AllLengthsSum() != 2 {
			t.Errorf("expected 2 got %d", chm.AllLengthsSum())
		}
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple", "avocado"})
		result := chm.AddSameCharsCollection("abc", col)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddSameCharsHashset("abc", hs)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		result := chm.AddHashsetLock("abc", hs)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AddSameCharsCollectionLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		col := corestr.New.Collection.Strings([]string{"apple"})
		result := chm.AddSameCharsCollectionLock("abc", col)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollection", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollection()
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByChars", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'b')
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("abc", "bcd")
		if hsc == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm.GetCharsGroups("apple", "banana")
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_List(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_List", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
		list := chm.List()
		if len(list) != 2 {
			t.Errorf("expected 2 got %d", len(list))
		}
	})
}

func Test_C26_CharHashsetMap_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListAsc", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "cherry", "apple", "banana")
		list := chm.SortedListAsc()
		if len(list) != 3 {
			t.Errorf("expected 3 got %d", len(list))
		}
		if list[0] != "apple" {
			t.Errorf("first should be apple got %s", list[0])
		}
	})
}

func Test_C26_CharHashsetMap_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SortedListDsc", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "cherry")
		list := chm.SortedListDsc()
		if len(list) != 2 {
			t.Errorf("expected 2 got %d", len(list))
		}
		if list[0] != "cherry" {
			t.Errorf("first should be cherry got %s", list[0])
		}
	})
}

func Test_C26_CharHashsetMap_SummaryString(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryString", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryString()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharHashsetMap_SummaryStringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_SummaryStringLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.SummaryStringLock()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharHashsetMap_String(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_String", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.String()
		if s == "" {
			t.Error("should not be empty")
		}
	})
}

func Test_C26_CharHashsetMap_StringLock(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_StringLock", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		s := chm.StringLock()
		if s == "" {
			t.Error("should not be empty")
		}
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
		if model == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonModelAny", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		if chm.JsonModelAny() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_MarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, err := json.Marshal(chm)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(data) == 0 {
			t.Error("should have data")
		}
	})
}

func Test_C26_CharHashsetMap_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_UnmarshalJSON", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		data, _ := json.Marshal(chm)
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := json.Unmarshal(data, chm2)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharHashsetMap_Json(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Json", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		result := chm.Json()
		if result.HasError() {
			t.Error("should not error")
		}
	})
}

func Test_C26_CharHashsetMap_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonPtr", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		ptr := chm.JsonPtr()
		if ptr == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJson", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		_, err := chm2.ParseInjectUsingJson(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_ParseInjectUsingJsonMust", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		result := chm2.ParseInjectUsingJsonMust(&jsonResult)
		if result == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_JsonParseSelfInject", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		jsonResult := chm.Json()
		chm2 := corestr.New.CharHashsetMap.Cap(10, 10)
		err := chm2.JsonParseSelfInject(&jsonResult)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C26_CharHashsetMap_AsJsonInterfaces(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_AsJsonInterfaces", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "a")
		if chm.AsJsonContractsBinder() == nil {
			t.Error("should not be nil")
		}
		if chm.AsJsoner() == nil {
			t.Error("should not be nil")
		}
		if chm.AsJsonMarshaller() == nil {
			t.Error("should not be nil")
		}
		if chm.AsJsonParseSelfInjector() == nil {
			t.Error("should not be nil")
		}
	})
}

func Test_C26_CharHashsetMap_RemoveAll(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_RemoveAll", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.RemoveAll()
		time.Sleep(10 * time.Millisecond) // allow goroutine cleanup
		if chm.Length() != 0 {
			t.Error("should be 0")
		}
	})
}

func Test_C26_CharHashsetMap_Clear(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_Clear", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		chm.Clear()
		time.Sleep(10 * time.Millisecond)
		if chm.Length() != 0 {
			t.Error("should be 0")
		}
	})
}

func Test_C26_CharHashsetMap_DataModel(t *testing.T) {
	safeTest(t, "Test_C26_CharHashsetMap_DataModel", func() {
		chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple")
		dm := corestr.NewCharHashsetMapDataModelUsing(chm)
		if dm == nil {
			t.Error("should not be nil")
		}
		chm2 := corestr.NewCharHashsetMapUsingDataModel(dm)
		if chm2 == nil {
			t.Error("should not be nil")
		}
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
			t.Error("timeout")
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
			t.Error("timeout")
		}
	})
}
