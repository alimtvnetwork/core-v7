package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// Collection — index access, paging, filtering
// ═══════════════════════════════════════════════════════════════

func Test_Cov47_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := caseV1Compat{Name: "IndexAt", Expected: "b", Actual: c.IndexAt(1), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_SafeIndexAtUsingLength_InRange(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_SafeIndexAtUsingLength_InRange", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "SafeIndexAt in range", Expected: "b", Actual: c.SafeIndexAtUsingLength("def", 2, 1), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_SafeIndexAtUsingLength_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_SafeIndexAtUsingLength_OutOfRange", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "SafeIndexAt oob", Expected: "def", Actual: c.SafeIndexAtUsingLength("def", 1, 5), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_First(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		tc := caseV1Compat{Name: "First", Expected: "x", Actual: c.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Last(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"x", "y"})
		tc := caseV1Compat{Name: "Last", Expected: "y", Actual: c.Last(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_LastOrDefault_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_LastOrDefault_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "LastOrDefault has", Expected: "b", Actual: c.LastOrDefault(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_LastOrDefault_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "LastOrDefault empty", Expected: "", Actual: c.LastOrDefault(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_FirstOrDefault_HasItems(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_FirstOrDefault_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "FirstOrDefault has", Expected: "a", Actual: c.FirstOrDefault(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_FirstOrDefault_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "FirstOrDefault empty", Expected: "", Actual: c.FirstOrDefault(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Single(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		tc := caseV1Compat{Name: "Single", Expected: "only", Actual: c.Single(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Take_LessThanLength(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Take_LessThanLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Take(2)
		tc := caseV1Compat{Name: "Take 2", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Take_MoreThanLength(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Take_MoreThanLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.Take(5)
		tc := caseV1Compat{Name: "Take more", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Take_Zero(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Take_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.Take(0)
		tc := caseV1Compat{Name: "Take zero", Expected: true, Actual: result.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Skip_Zero(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Skip_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.Skip(0)
		tc := caseV1Compat{Name: "Skip zero", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Skip_Some(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Skip_Some", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Skip(1)
		tc := caseV1Compat{Name: "Skip 1", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Reverse_Multiple(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Reverse_Multiple", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		tc := caseV1Compat{Name: "Reverse first", Expected: "c", Actual: c.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Reverse_Two", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()
		tc := caseV1Compat{Name: "Reverse two", Expected: "b", Actual: c.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_Reverse_Single", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()
		tc := caseV1Compat{Name: "Reverse single", Expected: "a", Actual: c.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		tc := caseV1Compat{Name: "GetPagesSize", Expected: 2, Actual: c.GetPagesSize(3), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetPagesSize_Zero(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetPagesSize_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "GetPagesSize zero", Expected: 0, Actual: c.GetPagesSize(0), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		page := c.GetSinglePageCollection(2, 1)
		tc := caseV1Compat{Name: "GetSinglePage", Expected: 2, Actual: page.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetSinglePageCollection_LastPage", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		page := c.GetSinglePageCollection(2, 3)
		tc := caseV1Compat{Name: "GetSinglePage last", Expected: 1, Actual: page.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetSinglePageCollection_SmallList(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetSinglePageCollection_SmallList", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		page := c.GetSinglePageCollection(5, 1)
		tc := caseV1Compat{Name: "GetSinglePage small", Expected: 1, Actual: page.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetPagedCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		paged := c.GetPagedCollection(2)
		tc := caseV1Compat{Name: "GetPagedCollection", Expected: 3, Actual: paged.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_GetPagedCollection_SmallList(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_GetPagedCollection_SmallList", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		paged := c.GetPagedCollection(5)
		tc := caseV1Compat{Name: "GetPagedCollection small", Expected: 1, Actual: paged.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_InsertAt_First(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_InsertAt_First", func() {
		c := corestr.New.Collection.Cap(0)
		c.InsertAt(0, "a")
		tc := caseV1Compat{Name: "InsertAt first", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
		tc := caseV1Compat{Name: "ChainRemoveAt", Expected: 2, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddCollection", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b", "c"})
		c1.AddCollection(c2)
		tc := caseV1Compat{Name: "AddCollection", Expected: 3, Actual: c1.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddCollection_Empty", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Cap(0)
		c1.AddCollection(c2)
		tc := caseV1Compat{Name: "AddCollection empty", Expected: 1, Actual: c1.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})
		c.AddCollections(c2, c3)
		tc := caseV1Compat{Name: "AddCollections", Expected: 3, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k1", "v1")
		c.AddHashmapsValues(hm)
		tc := caseV1Compat{Name: "AddHashmapsValues", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsValues_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil)
		tc := caseV1Compat{Name: "AddHashmapsValues nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k1", "v1")
		c.AddHashmapsKeys(hm)
		tc := caseV1Compat{Name: "AddHashmapsKeys", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeys_Nil(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeys_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil)
		tc := caseV1Compat{Name: "AddHashmapsKeys nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k1", "v1")
		c.AddHashmapsKeysValues(hm)
		tc := caseV1Compat{Name: "AddHashmapsKeysValues", Expected: 2, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeysValues_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeysValues(nil)
		tc := caseV1Compat{Name: "AddHashmapsKeysValues nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k1", "v1")
		c.AddHashmapsKeysValuesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return kv.Key + "=" + kv.Value, true, false
		}, hm)
		tc := caseV1Compat{Name: "AddHMKVUsingFilter", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter_Break", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(2)
		hm.AddOrUpdate("k1", "v1")
		hm.AddOrUpdate("k2", "v2")
		c.AddHashmapsKeysValuesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return kv.Key, true, true // break after first
		}, hm)
		tc := caseV1Compat{Name: "AddHMKVUsingFilter break", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov47_Collection_AddHashmapsKeysValuesUsingFilter_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeysValuesUsingFilter(nil, nil)
		tc := caseV1Compat{Name: "AddHMKVUsingFilter nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// CharHashsetMap — AddStrings, GetHashset, AddSameChars, lock methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov47_CharHashsetMap_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddStrings", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings("hello", "abc", "hi")
		tc := caseV1Compat{Name: "CHM AddStrings", Expected: 2, Actual: chm.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddStrings_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStrings()
		tc := caseV1Compat{Name: "CHM AddStrings empty", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetHashset_Found(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetHashset_Found", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.GetHashset("hello", false)
		tc := caseV1Compat{Name: "CHM GetHashset found", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetHashset_NotFound_NoCreate(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetHashset_NotFound_NoCreate", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.GetHashset("x", false)
		tc := caseV1Compat{Name: "CHM GetHashset not found no create", Expected: true, Actual: hs == nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetHashset_NotFound_Create(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetHashset_NotFound_Create", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.GetHashset("x", true)
		tc := caseV1Compat{Name: "CHM GetHashset not found create", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.GetHashsetLock(false, "hello")
		tc := caseV1Compat{Name: "CHM GetHashsetLock", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("hello")
		tc := caseV1Compat{Name: "CHM AddLock", Expected: true, Actual: chm.Has("hello"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddLock_ExistingChar(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddLock_ExistingChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddLock("hello")
		chm.AddLock("hi")
		tc := caseV1Compat{Name: "CHM AddLock existing", Expected: 2, Actual: chm.LengthOf('h'), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsCollection_New(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsCollection_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"abc", "axy"})
		hs := chm.AddSameCharsCollection("a", col)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollection new", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsCollection_Existing(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsCollection_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		col := corestr.New.Collection.Strings([]string{"axy"})
		hs := chm.AddSameCharsCollection("a", col)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollection existing", Expected: true, Actual: hs.Has("axy"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsCollection_NilCollection(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsCollection_NilCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := chm.AddSameCharsCollection("a", nil)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollection nil", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsCollection_ExistingNilCollection(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsCollection_ExistingNilCollection", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		hs := chm.AddSameCharsCollection("a", nil)
		tc := caseV1Compat{Name: "CHM AddSameCharsCollection existing nil col", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsHashset_New(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsHashset_New", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.StringsSpreadItems("abc", "axy")
		result := chm.AddSameCharsHashset("a", hs)
		tc := caseV1Compat{Name: "CHM AddSameCharsHashset new", Expected: true, Actual: result.Has("abc"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsHashset_Existing(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsHashset_Existing", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("abc")
		hs := corestr.New.Hashset.StringsSpreadItems("axy")
		result := chm.AddSameCharsHashset("a", hs)
		tc := caseV1Compat{Name: "CHM AddSameCharsHashset existing", Expected: true, Actual: result.Has("axy"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddSameCharsHashset_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddSameCharsHashset_NilHashset", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.AddSameCharsHashset("a", nil)
		tc := caseV1Compat{Name: "CHM AddSameCharsHashset nil", Expected: true, Actual: result != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddHashsetItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.StringsSpreadItems("hello", "abc")
		chm.AddHashsetItems(hs)
		tc := caseV1Compat{Name: "CHM AddHashsetItems", Expected: 2, Actual: chm.AllLengthsSum(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddHashsetItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddHashsetItems_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		hs := corestr.New.Hashset.Empty()
		chm.AddHashsetItems(hs)
		tc := caseV1Compat{Name: "CHM AddHashsetItems empty", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddCollectionItems", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		col := corestr.New.Collection.Strings([]string{"hello", "abc"})
		chm.AddCollectionItems(col)
		tc := caseV1Compat{Name: "CHM AddCollectionItems", Expected: 2, Actual: chm.AllLengthsSum(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddCollectionItems_Nil", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddCollectionItems(nil)
		tc := caseV1Compat{Name: "CHM AddCollectionItems nil", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddStringsLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock("hello", "abc")
		tc := caseV1Compat{Name: "CHM AddStringsLock", Expected: 2, Actual: chm.AllLengthsSum(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AddStringsLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AddStringsLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.AddStringsLock()
		tc := caseV1Compat{Name: "CHM AddStringsLock empty", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_LengthLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM LengthLock", Expected: 1, Actual: chm.LengthLock(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_IsEmptyLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM IsEmptyLock", Expected: true, Actual: chm.IsEmptyLock(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_IsEqualsLock", func() {
		chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm1.Add("hello")
		chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
		chm2.Add("hello")
		tc := caseV1Compat{Name: "CHM IsEqualsLock", Expected: true, Actual: chm1.IsEqualsLock(chm2), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_AllLengthsSumLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		chm.Add("abc")
		tc := caseV1Compat{Name: "CHM AllLengthsSumLock", Expected: 2, Actual: chm.AllLengthsSumLock(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_LengthOfLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_LengthOfLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		tc := caseV1Compat{Name: "CHM LengthOfLock", Expected: 1, Actual: chm.LengthOfLock('h'), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_LengthOfLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_LengthOfLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		tc := caseV1Compat{Name: "CHM LengthOfLock empty", Expected: 0, Actual: chm.LengthOfLock('x'), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_HasWithHashsetLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		found, hs := chm.HasWithHashsetLock("hello")
		tc := caseV1Compat{Name: "CHM HasWithHashsetLock", Expected: true, Actual: found, Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "CHM HasWithHashsetLock hs", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_HasWithHashsetLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_HasWithHashsetLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		found, _ := chm.HasWithHashsetLock("x")
		tc := caseV1Compat{Name: "CHM HasWithHashsetLock empty", Expected: false, Actual: found, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetCopyMapLock", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		m := chm.GetCopyMapLock()
		tc := caseV1Compat{Name: "CHM GetCopyMapLock", Expected: true, Actual: len(m) > 0, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetCopyMapLock_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		m := chm.GetCopyMapLock()
		tc := caseV1Compat{Name: "CHM GetCopyMapLock empty", Expected: 0, Actual: len(m), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetCharsGroups(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetCharsGroups", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.GetCharsGroups("hello", "abc", "hi")
		tc := caseV1Compat{Name: "CHM GetCharsGroups", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_GetCharsGroups_Empty", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		result := chm.GetCharsGroups()
		tc := caseV1Compat{Name: "CHM GetCharsGroups empty", Expected: true, Actual: result != nil, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov47_CharHashsetMap_HashsetByChar(t *testing.T) {
	safeTest(t, "Test_Cov47_CharHashsetMap_HashsetByChar", func() {
		chm := corestr.New.CharHashsetMap.Cap(10, 5)
		chm.Add("hello")
		hs := chm.HashsetByChar('h')
		tc := caseV1Compat{Name: "CHM HashsetByChar", Expected: true, Actual: hs != nil && hs.Has("hello"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}
