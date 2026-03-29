package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharHashsetMap — Segment 15: Full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovCHM_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_01_IsEmpty_HasItems", func() {
		chm := corestr.Empty.CharHashsetMap()
		if !chm.IsEmpty() {
			t.Fatal("expected empty")
		}
		if chm.HasItems() {
			t.Fatal("expected no items")
		}
		chm.Add("apple")
		if chm.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !chm.HasItems() {
			t.Fatal("expected items")
		}
	})
}

func Test_CovCHM_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_02_IsEmptyLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		if !chm.IsEmptyLock() {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovCHM_03_GetChar_GetCharOf(t *testing.T) {
	safeTest(t, "Test_CovCHM_03_GetChar_GetCharOf", func() {
		chm := corestr.Empty.CharHashsetMap()
		if chm.GetChar("hello") != 'h' {
			t.Fatal("expected h")
		}
		if chm.GetChar("") != 0 {
			t.Fatal("expected 0")
		}
		if chm.GetCharOf("hello") != 'h' {
			t.Fatal("expected h")
		}
		if chm.GetCharOf("") != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCHM_04_Add_AddLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_04_Add_AddLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		chm.Add("ant") // same char group
		if chm.Length() != 1 {
			t.Fatal("expected 1")
		}
		chm.Add("banana")
		if chm.Length() != 2 {
			t.Fatal("expected 2")
		}
		// AddLock
		chm.AddLock("cherry")
		if chm.Length() != 3 {
			t.Fatal("expected 3")
		}
		chm.AddLock("cat") // existing
	})
}

func Test_CovCHM_05_AddStrings_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_05_AddStrings_AddStringsLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		if chm.Length() != 2 {
			t.Fatal("expected 2")
		}
		chm.AddStrings()
		chm.AddStringsLock("cherry", "cat")
		chm.AddStringsLock()
	})
}

func Test_CovCHM_06_AddSameStartingCharItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_06_AddSameStartingCharItems", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddSameStartingCharItems('a', []string{"apple", "ant"})
		if chm.LengthOf('a') != 2 {
			t.Fatal("expected 2")
		}
		// existing
		chm.AddSameStartingCharItems('a', []string{"axe"})
		// empty
		chm.AddSameStartingCharItems('b', nil)
	})
}

func Test_CovCHM_07_Length_LengthLock_AllLengthsSum(t *testing.T) {
	safeTest(t, "Test_CovCHM_07_Length_LengthLock_AllLengthsSum", func() {
		chm := corestr.Empty.CharHashsetMap()
		if chm.Length() != 0 {
			t.Fatal("expected 0")
		}
		if chm.LengthLock() != 0 {
			t.Fatal("expected 0")
		}
		if chm.AllLengthsSum() != 0 {
			t.Fatal("expected 0")
		}
		chm.AddStrings("apple", "ant", "banana")
		if chm.AllLengthsSum() != 3 {
			t.Fatal("expected 3")
		}
		if chm.AllLengthsSumLock() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovCHM_08_LengthOfHashsetFromFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCHM_08_LengthOfHashsetFromFirstChar", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "ant")
		if chm.LengthOfHashsetFromFirstChar("a") != 2 {
			t.Fatal("expected 2")
		}
		if chm.LengthOfHashsetFromFirstChar("z") != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCHM_09_LengthOf_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_09_LengthOf_LengthOfLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		if chm.LengthOf('a') != 1 {
			t.Fatal("expected 1")
		}
		if chm.LengthOf('z') != 0 {
			t.Fatal("expected 0")
		}
		if chm.LengthOfLock('a') != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.Empty.CharHashsetMap()
		if e.LengthOf('a') != 0 {
			t.Fatal("expected 0")
		}
		if e.LengthOfLock('a') != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCHM_10_Has(t *testing.T) {
	safeTest(t, "Test_CovCHM_10_Has", func() {
		chm := corestr.Empty.CharHashsetMap()
		if chm.Has("apple") {
			t.Fatal("expected false")
		}
		chm.Add("apple")
		if !chm.Has("apple") {
			t.Fatal("expected true")
		}
		if chm.Has("banana") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCHM_11_HasWithHashset_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_11_HasWithHashset_Lock", func() {
		chm := corestr.Empty.CharHashsetMap()
		has, _ := chm.HasWithHashset("x")
		if has {
			t.Fatal("expected false")
		}
		chm.Add("apple")
		has2, hs := chm.HasWithHashset("apple")
		if !has2 || hs == nil {
			t.Fatal("expected found")
		}
		// missing char
		has3, _ := chm.HasWithHashset("banana")
		if has3 {
			t.Fatal("expected false")
		}
		// Lock variant
		has4, _ := chm.HasWithHashsetLock("apple")
		if !has4 {
			t.Fatal("expected found")
		}
		has5, _ := chm.HasWithHashsetLock("banana")
		if has5 {
			t.Fatal("expected false")
		}
		// empty lock
		e := corestr.Empty.CharHashsetMap()
		has6, _ := e.HasWithHashsetLock("x")
		if has6 {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCHM_12_IsEquals_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_12_IsEquals_IsEqualsLock", func() {
		a := corestr.Empty.CharHashsetMap()
		a.AddStrings("apple", "banana")
		b := corestr.Empty.CharHashsetMap()
		b.AddStrings("apple", "banana")
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
		if !a.IsEqualsLock(b) {
			t.Fatal("expected equal")
		}
		if a.IsEquals(nil) {
			t.Fatal("expected false")
		}
		if !a.IsEquals(a) {
			t.Fatal("expected true")
		}
		e1 := corestr.Empty.CharHashsetMap()
		e2 := corestr.Empty.CharHashsetMap()
		if !e1.IsEquals(e2) {
			t.Fatal("expected true")
		}
		if a.IsEquals(e1) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.Empty.CharHashsetMap()
		c.Add("apple")
		if a.IsEquals(c) {
			t.Fatal("expected false")
		}
		// same length, diff content
		d := corestr.Empty.CharHashsetMap()
		d.AddStrings("axe", "cherry")
		if a.IsEquals(d) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovCHM_13_GetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_13_GetMap_GetCopyMapLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		if len(chm.GetMap()) != 1 {
			t.Fatal("expected 1")
		}
		if len(chm.GetCopyMapLock()) != 1 {
			t.Fatal("expected 1")
		}
		e := corestr.Empty.CharHashsetMap()
		if len(e.GetCopyMapLock()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovCHM_14_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_CovCHM_14_GetCharsGroups", func() {
		chm := corestr.Empty.CharHashsetMap()
		r := chm.GetCharsGroups("apple", "ant", "banana")
		if r.Length() != 2 {
			t.Fatal("expected 2")
		}
		r2 := chm.GetCharsGroups()
		if r2 != chm {
			t.Fatal("expected same")
		}
	})
}

func Test_CovCHM_15_GetHashset_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_15_GetHashset_GetHashsetLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		hs := chm.GetHashset("a", false)
		if hs == nil {
			t.Fatal("expected found")
		}
		hs2 := chm.GetHashset("z", false)
		if hs2 != nil {
			t.Fatal("expected nil")
		}
		hs3 := chm.GetHashset("z", true)
		if hs3 == nil {
			t.Fatal("expected new")
		}
		hs4 := chm.GetHashsetLock(false, "a")
		if hs4 == nil {
			t.Fatal("expected found")
		}
	})
}

func Test_CovCHM_16_GetHashsetByChar_HashsetByChar_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_16_GetHashsetByChar_HashsetByChar_Lock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		if chm.GetHashsetByChar('a') == nil {
			t.Fatal("expected found")
		}
		if chm.HashsetByChar('a') == nil {
			t.Fatal("expected found")
		}
		hl := chm.HashsetByCharLock('a')
		if hl == nil {
			t.Fatal("expected found")
		}
		// missing
		hl2 := chm.HashsetByCharLock('z')
		if hl2 == nil {
			t.Fatal("expected empty hashset")
		}
	})
}

func Test_CovCHM_17_HashsetByStringFirstChar_Lock(t *testing.T) {
	safeTest(t, "Test_CovCHM_17_HashsetByStringFirstChar_Lock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		hs := chm.HashsetByStringFirstChar("a")
		if hs == nil {
			t.Fatal("expected found")
		}
		hs2 := chm.HashsetByStringFirstCharLock("a")
		if hs2 == nil {
			t.Fatal("expected found")
		}
	})
}

func Test_CovCHM_18_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_18_AddCollectionItems", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddCollectionItems(nil)
		col := corestr.New.Collection.Strings([]string{"apple", "banana"})
		chm.AddCollectionItems(col)
		if chm.AllLengthsSum() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovCHM_19_AddCharCollectionMapItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_19_AddCharCollectionMapItems", func() {
		chm := corestr.Empty.CharHashsetMap()
		ccm := corestr.Empty.CharCollectionMap()
		ccm.Add("apple")
		chm.AddCharCollectionMapItems(ccm)
		if chm.AllLengthsSum() != 1 {
			t.Fatal("expected 1")
		}
		chm.AddCharCollectionMapItems(nil)
	})
}

func Test_CovCHM_20_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_CovCHM_20_AddHashsetItems", func() {
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		chm.AddHashsetItems(hs)
		if chm.AllLengthsSum() != 1 {
			t.Fatal("expected 1")
		}
		chm.AddHashsetItems(corestr.New.Hashset.Empty())
	})
}

func Test_CovCHM_21_AddSameCharsCollection(t *testing.T) {
	safeTest(t, "Test_CovCHM_21_AddSameCharsCollection", func() {
		chm := corestr.Empty.CharHashsetMap()
		col := corestr.New.Collection.Strings([]string{"apple", "ant"})
		r := chm.AddSameCharsCollection("a", col)
		if r == nil {
			t.Fatal("expected hashset")
		}
		// existing, add more
		col2 := corestr.New.Collection.Strings([]string{"axe"})
		r2 := chm.AddSameCharsCollection("a", col2)
		if r2 == nil {
			t.Fatal("expected hashset")
		}
		// existing, nil
		r3 := chm.AddSameCharsCollection("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// new, nil
		r4 := chm.AddSameCharsCollection("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
		// new, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := chm.AddSameCharsCollection("b", col3)
		if r5 == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_CovCHM_22_AddSameCharsHashset(t *testing.T) {
	safeTest(t, "Test_CovCHM_22_AddSameCharsHashset", func() {
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		r := chm.AddSameCharsHashset("a", hs)
		if r == nil {
			t.Fatal("expected hashset")
		}
		// existing, add
		hs2 := corestr.New.Hashset.Strings([]string{"ant"})
		r2 := chm.AddSameCharsHashset("a", hs2)
		if r2 == nil {
			t.Fatal("expected hashset")
		}
		// existing, nil
		r3 := chm.AddSameCharsHashset("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// new, nil
		r4 := chm.AddSameCharsHashset("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
		// new, with items
		hs3 := corestr.New.Hashset.Strings([]string{"banana"})
		r5 := chm.AddSameCharsHashset("b", hs3)
		if r5 == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_CovCHM_23_AddSameCharsCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_23_AddSameCharsCollectionLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		col := corestr.New.Collection.Strings([]string{"apple"})
		r := chm.AddSameCharsCollectionLock("a", col)
		if r == nil {
			t.Fatal("expected hashset")
		}
		// existing, add
		col2 := corestr.New.Collection.Strings([]string{"ant"})
		r2 := chm.AddSameCharsCollectionLock("a", col2)
		if r2 == nil {
			t.Fatal("expected hashset")
		}
		// existing, nil
		r3 := chm.AddSameCharsCollectionLock("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// new, nil
		r4 := chm.AddSameCharsCollectionLock("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
		// new, with items
		col3 := corestr.New.Collection.Strings([]string{"banana"})
		r5 := chm.AddSameCharsCollectionLock("b", col3)
		if r5 == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_CovCHM_24_AddHashsetLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_24_AddHashsetLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		hs := corestr.New.Hashset.Strings([]string{"apple"})
		r := chm.AddHashsetLock("a", hs)
		if r == nil {
			t.Fatal("expected hashset")
		}
		// existing, add
		hs2 := corestr.New.Hashset.Strings([]string{"ant"})
		r2 := chm.AddHashsetLock("a", hs2)
		if r2 == nil {
			t.Fatal("expected hashset")
		}
		// existing, nil
		r3 := chm.AddHashsetLock("a", nil)
		if r3 == nil {
			t.Fatal("expected existing")
		}
		// new, nil
		r4 := chm.AddHashsetLock("z", nil)
		if r4 == nil {
			t.Fatal("expected new")
		}
		// new, with items
		hs3 := corestr.New.Hashset.Strings([]string{"banana"})
		r5 := chm.AddHashsetLock("b", hs3)
		if r5 == nil {
			t.Fatal("expected hashset")
		}
	})
}

func Test_CovCHM_25_List_SortedListAsc_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_CovCHM_25_List_SortedListAsc_SortedListDsc", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("banana", "apple")
		list := chm.List()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
		asc := chm.SortedListAsc()
		if len(asc) != 2 {
			t.Fatal("expected 2")
		}
		dsc := chm.SortedListDsc()
		if len(dsc) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovCHM_26_HashsetsCollectionByChars(t *testing.T) {
	safeTest(t, "Test_CovCHM_26_HashsetsCollectionByChars", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByChars('a', 'z')
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		// empty
		e := corestr.Empty.CharHashsetMap()
		hsc2 := e.HashsetsCollectionByChars('a')
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCHM_27_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	safeTest(t, "Test_CovCHM_27_HashsetsCollectionByStringsFirstChar", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		hsc := chm.HashsetsCollectionByStringsFirstChar("apple", "cherry")
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		e := corestr.Empty.CharHashsetMap()
		hsc2 := e.HashsetsCollectionByStringsFirstChar("x")
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCHM_28_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_CovCHM_28_HashsetsCollection", func() {
		chm := corestr.Empty.CharHashsetMap()
		hsc := chm.HashsetsCollection()
		if hsc == nil {
			t.Fatal("expected non-nil")
		}
		chm.Add("apple")
		hsc2 := chm.HashsetsCollection()
		if hsc2 == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovCHM_29_String_StringLock_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovCHM_29_String_StringLock_SummaryString", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		if chm.String() == "" {
			t.Fatal("expected non-empty")
		}
		if chm.StringLock() == "" {
			t.Fatal("expected non-empty")
		}
		if chm.SummaryString() == "" {
			t.Fatal("expected non-empty")
		}
		if chm.SummaryStringLock() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovCHM_30_Print_PrintLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_30_Print_PrintLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		chm.Print(false)
		chm.Print(true)
		chm.PrintLock(false)
		chm.PrintLock(true)
	})
}

func Test_CovCHM_31_JsonModel_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovCHM_31_JsonModel_JsonModelAny", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		_ = chm.JsonModel()
		_ = chm.JsonModelAny()
	})
}

func Test_CovCHM_32_MarshalJSON_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovCHM_32_MarshalJSON_UnmarshalJSON", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		data, err := chm.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		chm2 := corestr.Empty.CharHashsetMap()
		err2 := chm2.UnmarshalJSON(data)
		if err2 != nil {
			t.Fatal("unexpected error")
		}
		err3 := chm2.UnmarshalJSON([]byte("bad"))
		if err3 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovCHM_33_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovCHM_33_Json_JsonPtr_ParseInject", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		_ = chm.Json()
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		r, err := chm2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if r.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovCHM_34_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovCHM_34_ParseInjectUsingJsonMust", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		r := chm2.ParseInjectUsingJsonMust(jr)
		if r.IsEmpty() {
			t.Fatal("expected not empty")
		}
	})
}

func Test_CovCHM_35_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovCHM_35_JsonParseSelfInject", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.Add("apple")
		jr := chm.JsonPtr()
		chm2 := corestr.Empty.CharHashsetMap()
		err := chm2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovCHM_36_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovCHM_36_AsInterfaces", func() {
		chm := corestr.Empty.CharHashsetMap()
		_ = chm.AsJsonContractsBinder()
		_ = chm.AsJsoner()
		_ = chm.AsJsonMarshaller()
		_ = chm.AsJsonParseSelfInjector()
	})
}

func Test_CovCHM_37_Clear_RemoveAll(t *testing.T) {
	safeTest(t, "Test_CovCHM_37_Clear_RemoveAll", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddStrings("apple", "banana")
		chm.Clear()
		if !chm.IsEmpty() {
			t.Fatal("expected empty")
		}
		chm.Clear() // already empty
		chm.Add("x")
		chm.RemoveAll()
		if !chm.IsEmpty() {
			t.Fatal("expected empty")
		}
		// RemoveAll on empty
		e := corestr.Empty.CharHashsetMap()
		e.RemoveAll()
	})
}

func Test_CovCHM_38_AddCollectionItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_38_AddCollectionItemsAsyncLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		// nil
		chm.AddCollectionItemsAsyncLock(nil, nil)
		// empty
		chm.AddCollectionItemsAsyncLock(corestr.Empty.Collection(), nil)
	})
}

func Test_CovCHM_39_AddHashsetItemsAsyncLock(t *testing.T) {
	safeTest(t, "Test_CovCHM_39_AddHashsetItemsAsyncLock", func() {
		chm := corestr.Empty.CharHashsetMap()
		chm.AddHashsetItemsAsyncLock(nil, nil)
		chm.AddHashsetItemsAsyncLock(corestr.New.Hashset.Empty(), nil)
	})
}
