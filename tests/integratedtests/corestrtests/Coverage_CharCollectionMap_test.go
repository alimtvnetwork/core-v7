package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — Segment 14: Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovCCM_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_01_IsEmpty_HasItems", func() {
		ccm := corestr.Empty.CharCollectionMap()
		if !ccm.IsEmpty() {
			t.Fatal("expected empty")
		}
		if ccm.HasItems() {
			t.Fatal("expected no items")
		}
		ccm.Add("apple")
		if ccm.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !ccm.HasItems() {
			t.Fatal("expected items")
		}
	})
}

func Test_CovCCM_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_02_IsEmptyLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		if !ccm.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovCCM_03_GetChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_03_GetChar", func() {
		ccm := corestr.Empty.CharCollectionMap()
		c := ccm.GetChar("hello")
		if c != 'h' {
			t.Fatal("expected h")
		}
		c2 := ccm.GetChar("")
		if c2 != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCCM_04_Add_AddLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_04_Add_AddLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Add("ant")
		if ccm.Length() != 1 {
			t.Fatal("expected 1 char group")
		}
		ccm.Add("banana")
		if ccm.Length() != 2 {
			t.Fatal("expected 2 char groups")
		}
		// AddLock
		ccm.AddLock("cherry")
		if ccm.Length() != 3 {
			t.Fatal("expected 3")
		}
		// AddLock existing char
		ccm.AddLock("cat")
	})
}

func Test_CovCCM_05_AddStrings(t *testing.T) {
	safeTest(t, "Test_CovCCM_05_AddStrings", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "ant", "banana")
		if ccm.Length() != 2 {
			t.Fatal("expected 2")
		}
		ccm.AddStrings()
	})
}

func Test_CovCCM_06_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_06_AddSameStartingCharItems", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddSameStartingCharItems('a', []string{"apple", "ant"}, false)
		if ccm.Length() != 1 {
			t.Fatal("expected 1")
		}
		// add to existing
		ccm.AddSameStartingCharItems('a', []string{"axe"}, false)
		// empty
		ccm.AddSameStartingCharItems('b', nil, false)
	})
}

func Test_CovCCM_07_Length_LengthLock_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_CovCCM_07_Length_LengthLock_AllLengthsSum", func() {
		ccm := corestr.Empty.CharCollectionMap()
		if ccm.Length() != 0 {
			t.Fatal("expected 0")
		}
		if ccm.LengthLock() != 0 {
			t.Fatal("expected 0")
		}
		if ccm.AllLengthsSum() != 0 {
			t.Fatal("expected 0")
		}
		ccm.AddStrings("apple", "ant", "banana")
		if ccm.AllLengthsSum() != 3 {
			t.Fatal("expected 3")
		}
		if ccm.AllLengthsSumLock() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovCCM_08_LengthOfCollectionFromFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_08_LengthOfCollectionFromFirstChar", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "ant")
		l := ccm.LengthOfCollectionFromFirstChar("a")
		if l != 2 {
			t.Fatal("expected 2")
		}
		l2 := ccm.LengthOfCollectionFromFirstChar("z")
		if l2 != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCCM_09_LengthOf_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_09_LengthOf_LengthOfLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		if ccm.LengthOf('a') != 1 {
			t.Fatal("expected 1")
		}
		if ccm.LengthOf('z') != 0 {
			t.Fatal("expected 0")
		}
		if ccm.LengthOfLock('a') != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.CharCollectionMap()
		if e.LengthOf('a') != 0 {
			t.Fatal("expected 0")
		}
		if e.LengthOfLock('a') != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCCM_10_Has(t *testing.T) {
	safeTest(t, "Test_CovCCM_10_Has", func() {
		ccm := corestr.Empty.CharCollectionMap()
		if ccm.Has("apple") {
			t.Fatal("expected false")
		}
		ccm.Add("apple")
		if !ccm.Has("apple") {
			t.Fatal("expected true")
		}
		if ccm.Has("ant") {
			t.Fatal("expected false")
		}
		// missing char group
		if ccm.Has("banana") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCCM_11_HasWithCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_11_HasWithCollection", func() {
		ccm := corestr.Empty.CharCollectionMap()
		// empty
		has, col := ccm.HasWithCollection("apple")
		if has || col == nil {
			t.Fatal("unexpected")
		}
		ccm.Add("apple")
		has2, col2 := ccm.HasWithCollection("apple")
		if !has2 || col2 == nil {
			t.Fatal("expected found")
		}
		// missing char
		has3, _ := ccm.HasWithCollection("banana")
		if has3 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCCM_12_HasWithCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_12_HasWithCollectionLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		has, _ := ccm.HasWithCollectionLock("x")
		if has {
			t.Fatal("expected false")
		}
		ccm.Add("apple")
		has2, col := ccm.HasWithCollectionLock("apple")
		if !has2 || col == nil {
			t.Fatal("expected found")
		}
		// missing char
		has3, _ := ccm.HasWithCollectionLock("banana")
		if has3 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCCM_13_IsEquals_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_13_IsEquals_IsEqualsLock", func() {
		a := corestr.Empty.CharCollectionMap()
		a.AddStrings("apple", "banana")
		b := corestr.Empty.CharCollectionMap()
		b.AddStrings("apple", "banana")
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
		if !a.IsEqualsLock(b) {
			t.Fatal("expected equal")
		}
		// nil
		if a.IsEquals(nil) {
			t.Fatal("expected false")
		}
		// same
		if !a.IsEquals(a) {
			t.Fatal("expected true")
		}
		// both empty
		e1 := corestr.Empty.CharCollectionMap()
		e2 := corestr.Empty.CharCollectionMap()
		if !e1.IsEquals(e2) {
			t.Fatal("expected true")
		}
		// one empty
		if a.IsEquals(e1) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.Empty.CharCollectionMap()
		c.Add("apple")
		if a.IsEquals(c) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCCM_14_IsEqualsCaseSensitive(t *testing.T) {
	safeTest(t, "Test_CovCCM_14_IsEqualsCaseSensitive", func() {
		a := corestr.Empty.CharCollectionMap()
		a.Add("Apple")
		b := corestr.Empty.CharCollectionMap()
		b.Add("Apple")
		if !a.IsEqualsCaseSensitive(true, b) {
			t.Fatal("expected equal")
		}
		if !a.IsEqualsCaseSensitiveLock(true, b) {
			t.Fatal("expected equal")
		}
		// missing key in right
		c := corestr.Empty.CharCollectionMap()
		c.Add("Banana")
		// same length but different keys - need same number of char groups
		d := corestr.Empty.CharCollectionMap()
		d.Add("Axe")
		if a.IsEqualsCaseSensitive(true, d) {
			t.Fatal("expected false - different content")
		}
	})
}

func Test_CovCCM_15_GetCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_15_GetCollection", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollection("a", false)
		if col == nil {
			t.Fatal("expected found")
		}
		// not found, no add
		col2 := ccm.GetCollection("z", false)
		if col2 != nil {
			t.Fatal("expected nil")
		}
		// not found, add new
		col3 := ccm.GetCollection("z", true)
		if col3 == nil {
			t.Fatal("expected new collection")
		}
	})
}

func Test_CovCCM_16_GetCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_16_GetCollectionLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollectionLock("a", false)
		if col == nil {
			t.Fatal("expected found")
		}
	})
}

func Test_CovCCM_17_GetCollectionByChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_17_GetCollectionByChar", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		col := ccm.GetCollectionByChar('a')
		if col == nil {
			t.Fatal("expected found")
		}
		col2 := ccm.GetCollectionByChar('z')
		if col2 != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovCCM_18_GetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_18_GetMap_GetCopyMapLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		m := ccm.GetMap()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
		m2 := ccm.GetCopyMapLock()
		if len(m2) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.CharCollectionMap()
		m3 := e.GetCopyMapLock()
		if len(m3) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCCM_19_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_CovCCM_19_GetCharsGroups", func() {
		ccm := corestr.Empty.CharCollectionMap()
		r := ccm.GetCharsGroups([]string{"apple", "ant", "banana"})
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		r2 := ccm.GetCharsGroups(nil)
		if r2 != ccm {
			t.Fatal("expected same")
		}
	})
}

func Test_CovCCM_20_List_ListLock_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_CovCCM_20_List_ListLock_SortedListAsc", func() {
		ccm := corestr.Empty.CharCollectionMap()
		if len(ccm.List()) != 0 {
			t.Fatal("expected 0")
		}
		ccm.AddStrings("banana", "apple")
		list := ccm.List()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
		ll := ccm.ListLock()
		if len(ll) != 2 {
			t.Fatal("expected 2")
		}
		sorted := ccm.SortedListAsc()
		if len(sorted) != 2 {
			t.Fatal("expected 2")
		}
		// empty sorted
		e := corestr.Empty.CharCollectionMap()
		if len(e.SortedListAsc()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCCM_21_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovCCM_21_AddCollectionItems", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddCollectionItems(nil)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(col)
		if ccm.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovCCM_22_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_CovCCM_22_AddHashmapsValues", func() {
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("k", "apple")
		ccm.AddHashmapsValues(hm, nil)
		if ccm.AllLengthsSum() != 1 {
			t.Fatal("expected 1")
		}
		ccm.AddHashmapsValues(nil)
	})
}

func Test_CovCCM_23_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovCCM_23_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		filter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, false
		}
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(filter, hm, nil)
		if ccm.AllLengthsSum() != 1 {
			t.Fatal("expected 1")
		}
		// nil
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(filter, nil)
		// break
		breakFilter := func(p corestr.KeyValuePair) (string, bool, bool) {
			return p.Value, true, true
		}
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(breakFilter, hm)
	})
}

func Test_CovCCM_24_AddHashmapsKeysValuesBoth(t *testing.T) {
	safeTest(t, "Test_CovCCM_24_AddHashmapsKeysValuesBoth", func() {
		ccm := corestr.Empty.CharCollectionMap()
		hm := corestr.Empty.Hashmap()
		hm.AddOrUpdate("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)
		if ccm.AllLengthsSum() < 2 {
			t.Fatal("expected at least 2")
		}
		ccm.AddHashmapsKeysValuesBoth(nil)
	})
}

func Test_CovCCM_25_AddCharHashsetMap(t *testing.T) {
	safeTest(t, "Test_CovCCM_25_AddCharHashsetMap", func() {
		ccm := corestr.Empty.CharCollectionMap()
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		ccm.AddCharHashsetMap(chm)
		if ccm.AllLengthsSum() != 1 {
			t.Fatal("expected 1")
		}
		// empty
		ccm.AddCharHashsetMap(corestr.Empty.CharHashsetMap())
	})
}

func Test_CovCCM_26_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_26_AddSameCharsCollection", func() {
		ccm := corestr.Empty.CharCollectionMap()
		col := corestr.New.Collection.Strings([]string{"apple", "ant"})
		r := ccm.AddSameCharsCollection("a", col)
		if r == nil {
			t.Fatal("expected collection")
		}
		// add more to existing
		col2 := corestr.New.Collection.Strings([]string{"axe"})
		r2 := ccm.AddSameCharsCollection("a", col2)
		if r2 == nil {
			t.Fatal("expected collection")
		}
		// nil collection, existing char
		r3 := ccm.AddSameCharsCollection("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// nil collection, new char
		r4 := ccm.AddSameCharsCollection("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
	})
}

func Test_CovCCM_27_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_27_AddSameCharsCollectionLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		col := corestr.New.Collection.Strings([]string{"apple"})
		r := ccm.AddSameCharsCollectionLock("a", col)
		if r == nil {
			t.Fatal("expected collection")
		}
		// existing, add more
		col2 := corestr.New.Collection.Strings([]string{"ant"})
		r2 := ccm.AddSameCharsCollectionLock("a", col2)
		if r2 == nil {
			t.Fatal("expected collection")
		}
		// existing, nil collection
		r3 := ccm.AddSameCharsCollectionLock("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// new char, nil
		r4 := ccm.AddSameCharsCollectionLock("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
		// new char, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := ccm.AddSameCharsCollectionLock("b", col3)
		if r5 == nil {
			t.Fatal("expected collection")
		}
	})
}

func Test_CovCCM_28_Resize_AddLength(t *testing.T) {
	safeTest(t, "Test_CovCCM_28_Resize_AddLength", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Resize(10)
		ccm.Resize(0) // no-op, current >= new
		ccm.AddLength(5, 3)
		ccm.AddLength()
	})
}

func Test_CovCCM_29_HashsetByChar_HashsetByCharLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_29_HashsetByChar_HashsetByCharLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hs := ccm.HashsetByChar('a')
		if hs == nil {
			t.Fatal("expected hashset")
		}
		hs2 := ccm.HashsetByChar('z')
		if hs2 != nil {
			t.Fatal("expected nil")
		}
		hs3 := ccm.HashsetByCharLock('a')
		if hs3 == nil {
			t.Fatal("expected hashset")
		}
		// nil collection
		hs4 := ccm.HashsetByCharLock('z')
		if hs4 == nil {
			t.Fatal("expected empty hashset")
		}
	})
}

func Test_CovCCM_30_HashsetByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_30_HashsetByStringFirstChar", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstChar("a")
		if hs == nil {
			t.Fatal("expected hashset")
		}
		hs2 := ccm.HashsetByStringFirstCharLock("a")
		if hs2 == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_CovCCM_31_HashsetsCollectionByStringFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCCM_31_HashsetsCollectionByStringFirstChar", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "banana")
		hsc := ccm.HashsetsCollectionByStringFirstChar("apple", "banana", "cherry")
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		// empty
		e := corestr.Empty.CharCollectionMap()
		hsc2 := e.HashsetsCollectionByStringFirstChar("x")
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCCM_32_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovCCM_32_HashsetsCollection", func() {
		ccm := corestr.Empty.CharCollectionMap()
		hsc := ccm.HashsetsCollection()
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		ccm.Add("apple")
		hsc2 := ccm.HashsetsCollection()
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCCM_33_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_CovCCM_33_HashsetsCollectionByChars", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		hsc := ccm.HashsetsCollectionByChars('a', 'z')
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		// empty
		e := corestr.Empty.CharCollectionMap()
		hsc2 := e.HashsetsCollectionByChars('a')
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCCM_34_String_StringLock_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovCCM_34_String_StringLock_SummaryString", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		s := ccm.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		sl := ccm.StringLock()
		if sl == "" {
			t.Fatal("expected non-empty")
		}
		ss := ccm.SummaryString()
		if ss == "" {
			t.Fatal("expected non-empty")
		}
		ssl := ccm.SummaryStringLock()
		if ssl == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovCCM_35_Print_PrintLock(t *testing.T) {
	safeTest(t, "Test_CovCCM_35_Print_PrintLock", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		ccm.Print(false) // skip
		ccm.Print(true)
		ccm.PrintLock(false) // skip
		ccm.PrintLock(true)
	})
}

func Test_CovCCM_36_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovCCM_36_JsonModel_JsonModelAny", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		_ = ccm.JsonModel()
		_ = ccm.JsonModelAny()
	})
}

func Test_CovCCM_37_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovCCM_37_MarshalJSON_UnmarshalJSON", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		data, err := ccm.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		ccm2 := corestr.Empty.CharCollectionMap()
		err2 := ccm2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		// invalid
		err3 := ccm2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovCCM_38_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCCM_38_Json_JsonPtr_ParseInject", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		_ = ccm.Json()
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		r, err := ccm2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovCCM_39_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovCCM_39_ParseInjectUsingJsonMust", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		r := ccm2.ParseInjectUsingJsonMust(jr)
		if r.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovCCM_40_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovCCM_40_JsonParseSelfInject", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.Empty.CharCollectionMap()
		err := ccm2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovCCM_41_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCCM_41_AsInterfaces", func() {
		ccm := corestr.Empty.CharCollectionMap()
		_ = ccm.AsJsonContractsBinder()
		_ = ccm.AsJsoner()
		_ = ccm.AsJsonMarshaller()
		_ = ccm.AsJsonParseSelfInjector()
	})
}

func Test_CovCCM_42_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_CovCCM_42_Clear_Dispose", func() {
		ccm := corestr.Empty.CharCollectionMap()
		ccm.AddStrings("apple", "banana")
		ccm.Clear()
		if !ccm.IsEmpty() {
			t.Fatal("expected empty")
		}
		// clear already empty
		ccm.Clear()
		// dispose
		ccm2 := corestr.Empty.CharCollectionMap()
		ccm2.Add("x")
		ccm2.Dispose()
	})
}
