package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — IsEmpty / HasItems / Length / Count / Capacity
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"empty": c.IsEmpty(), "items": c.HasItems(), "len": c.Length(), "count": c.Count(), "hasAny": c.HasAnyItem()}
		expected := args.Map{"empty": true, "items": false, "len": 0, "count": 0, "hasAny": false}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- empty", actual)
	})
}

func Test_I31_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Length_Nil", func() {
		var c *corestr.Collection
		actual := args.Map{"len": c.Length(), "empty": c.IsEmpty()}
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil length", actual)
	})
}

func Test_I31_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"empty": c.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEmptyLock", actual)
	})
}

func Test_I31_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LengthLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		actual := args.Map{"len": c.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LengthLock", actual)
	})
}

func Test_I31_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"cap": c.Capacity() >= 10}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Capacity", actual)
	})
}

func Test_I31_Collection_Capacity_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Capacity_Nil", func() {
		c := corestr.New.Collection.Cap(0)
		actual := args.Map{"cap": c.Capacity()}
		expected := args.Map{"cap": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- Capacity nil items", actual)
	})
}

func Test_I31_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"li": c.LastIndex(), "hasIdx": c.HasIndex(1), "noIdx": c.HasIndex(5)}
		expected := args.Map{"li": 1, "hasIdx": true, "noIdx": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LastIndex/HasIndex", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Add variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_Add(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Add", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Add", actual)
	})
}

func Test_I31_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmpty", actual)
	})
}

func Test_I31_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmptyWhitespace", actual)
	})
}

func Test_I31_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddIf", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "a")
		c.AddIf(false, "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIf", actual)
	})
}

func Test_I31_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(true, "a", "b")
		c.AddIfMany(false, "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIfMany", actual)
	})
}

func Test_I31_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddError", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddError(nil)
		c.AddError(fmt.Errorf("oops"))
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddError", actual)
	})
}

func Test_I31_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFunc", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFunc(func() string { return "computed" })
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "computed"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddFunc", actual)
	})
}

func Test_I31_Collection_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFuncErr_NoErr", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "ok"}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddFuncErr no err", actual)
	})
}

func Test_I31_Collection_AddFuncErr_Err(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFuncErr_Err", func() {
		c := corestr.New.Collection.Cap(5)
		called := false
		c.AddFuncErr(func() (string, error) { return "", fmt.Errorf("fail") }, func(e error) { called = true })
		actual := args.Map{"empty": c.IsEmpty(), "called": called}
		expected := args.Map{"empty": true, "called": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddFuncErr err", actual)
	})
}

func Test_I31_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddLock", actual)
	})
}

func Test_I31_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Adds", func() {
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Adds", actual)
	})
}

func Test_I31_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsLock("a", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsLock", actual)
	})
}

func Test_I31_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddStrings", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStrings", actual)
	})
}

func Test_I31_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollection", func() {
		c := corestr.New.Collection.Cap(5)
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddCollection(other)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
	})
}

func Test_I31_Collection_AddCollection_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollection_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("x")
		other := corestr.New.Collection.Cap(5)
		c.AddCollection(other)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddCollection empty", actual)
	})
}

func Test_I31_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollections", func() {
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		empty := corestr.New.Collection.Cap(0)
		c.AddCollections(c1, empty, c2)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
	})
}

func Test_I31_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddWithWgLock", actual)
	})
}

func Test_I31_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsValues", actual)
	})
}

func Test_I31_Collection_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsValues_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsValues nil", actual)
	})
}

func Test_I31_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)
		actual := args.Map{"has": c.Has("k")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddHashmapsKeys", actual)
	})
}

func Test_I31_Collection_AddHashmapsKeys_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeys_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsKeys nil", actual)
	})
}

func Test_I31_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(hm)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsKeysValues", actual)
	})
}

func Test_I31_Collection_AddHashmapsKeysValues_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeysValues_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeysValues(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsKeysValues nil", actual)
	})
}

func Test_I31_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty("a", "", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddsNonEmpty", actual)
	})
}

func Test_I31_Collection_AddsNonEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmpty_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsNonEmpty nil", actual)
	})
}

func Test_I31_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Cap(5)
		s := "hello"
		empty := ""
		c.AddsNonEmptyPtrLock(&s, nil, &empty)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddsNonEmptyPtrLock", actual)
	})
}

func Test_I31_Collection_AddsNonEmptyPtrLock_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmptyPtrLock_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmptyPtrLock(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsNonEmptyPtrLock nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Indexing (First / Last / IndexAt / SafeIndexAt)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"at1": c.IndexAt(1)}
		expected := args.Map{"at1": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IndexAt", actual)
	})
}

func Test_I31_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_I31_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"safe": c.SafeIndexAtUsingLength("def", 2, 1), "oob": c.SafeIndexAtUsingLength("def", 2, 5)}
		expected := args.Map{"safe": "b", "oob": "def"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SafeIndexAtUsingLength", actual)
	})
}

func Test_I31_Collection_First_Last(t *testing.T) {
	safeTest(t, "Test_I31_Collection_First_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"first": c.First(), "last": c.Last()}
		expected := args.Map{"first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- First/Last", actual)
	})
}

func Test_I31_Collection_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_FirstOrDefault_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"val": c.FirstOrDefault()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- FirstOrDefault empty", actual)
	})
}

func Test_I31_Collection_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LastOrDefault_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"val": c.LastOrDefault()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- LastOrDefault empty", actual)
	})
}

func Test_I31_Collection_Single(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		actual := args.Map{"val": c.Single()}
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Single", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Take / Skip / Reverse
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_Take(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		t1 := c.Take(2)
		actual := args.Map{"len": t1.Length(), "first": t1.First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take", actual)
	})
}

func Test_I31_Collection_Take_All(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take_All", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		t1 := c.Take(5)
		actual := args.Map{"same": t1.Length()}
		expected := args.Map{"same": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take all", actual)
	})
}

func Test_I31_Collection_Take_Zero(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		t1 := c.Take(0)
		actual := args.Map{"empty": t1.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take zero", actual)
	})
}

func Test_I31_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		s := c.Skip(1)
		actual := args.Map{"len": s.Length(), "first": s.First()}
		expected := args.Map{"len": 2, "first": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip", actual)
	})
}

func Test_I31_Collection_Skip_Zero(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Skip_Zero", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.Skip(0)
		actual := args.Map{"len": s.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip zero", actual)
	})
}

func Test_I31_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		actual := args.Map{"first": c.First(), "last": c.Last()}
		expected := args.Map{"first": "c", "last": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse", actual)
	})
}

func Test_I31_Collection_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse_Two", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse two", actual)
	})
}

func Test_I31_Collection_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse_Single", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse single", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — RemoveAt / ChainRemoveAt / InsertAt
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)
		actual := args.Map{"ok": ok, "len": c.Length()}
		expected := args.Map{"ok": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
	})
}

func Test_I31_Collection_RemoveAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt_OutOfRange", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		ok := c.RemoveAt(5)
		actual := args.Map{"ok": ok}
		expected := args.Map{"ok": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt out of range", actual)
	})
}

func Test_I31_Collection_RemoveAt_Negative(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt_Negative", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		ok := c.RemoveAt(-1)
		actual := args.Map{"ok": ok}
		expected := args.Map{"ok": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt negative", actual)
	})
}

func Test_I31_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(0)
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ChainRemoveAt", actual)
	})
}

func Test_I31_Collection_InsertAt_First(t *testing.T) {
	safeTest(t, "Test_I31_Collection_InsertAt_First", func() {
		c := corestr.New.Collection.Cap(5)
		c.InsertAt(0, "a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- InsertAt first", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		actual := args.Map{"pages": c.GetPagesSize(2), "zero": c.GetPagesSize(0), "neg": c.GetPagesSize(-1)}
		expected := args.Map{"pages": 3, "zero": 0, "neg": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagesSize", actual)
	})
}

func Test_I31_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		p := c.GetSinglePageCollection(2, 2)
		actual := args.Map{"len": p.Length(), "first": p.First()}
		expected := args.Map{"len": 2, "first": "c"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection", actual)
	})
}

func Test_I31_Collection_GetSinglePageCollection_LastPage(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection_LastPage", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		p := c.GetSinglePageCollection(2, 3)
		actual := args.Map{"len": p.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection last page", actual)
	})
}

func Test_I31_Collection_GetSinglePageCollection_SmallList(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection_SmallList", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		p := c.GetSinglePageCollection(10, 1)
		actual := args.Map{"len": p.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection small", actual)
	})
}

func Test_I31_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagedCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		paged := c.GetPagedCollection(2)
		actual := args.Map{"notNil": paged != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection", actual)
	})
}

func Test_I31_Collection_GetPagedCollection_SmallList(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagedCollection_SmallList", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		paged := c.GetPagedCollection(10)
		actual := args.Map{"notNil": paged != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection small", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — EachItemSplitBy / ConcatNew
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_I31_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- EachItemSplitBy", actual)
	})
}

func Test_I31_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0, "b", "c")
		actual := args.Map{"len": cn.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ConcatNew", actual)
	})
}

func Test_I31_Collection_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ConcatNew_NoArgs", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0)
		actual := args.Map{"len": cn.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- ConcatNew no args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — IsEquals
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_Same", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"same": c.IsEquals(c)}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals same ptr", actual)
	})
}

func Test_I31_Collection_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_BothNil", func() {
		var c *corestr.Collection
		actual := args.Map{"eq": c.IsEquals(nil)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- IsEquals both nil", actual)
	})
}

func Test_I31_Collection_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_OneNil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"eq": c.IsEquals(nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- IsEquals one nil", actual)
	})
}

func Test_I31_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_BothEmpty", func() {
		c1 := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Cap(5)
		actual := args.Map{"eq": c1.IsEquals(c2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEquals both empty", actual)
	})
}

func Test_I31_Collection_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_DiffLen", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"eq": c1.IsEquals(c2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals diff len", actual)
	})
}

func Test_I31_Collection_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_DiffItems", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		actual := args.Map{"eq": c1.IsEquals(c2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals diff items", actual)
	})
}

func Test_I31_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		c1 := corestr.New.Collection.Strings([]string{"Hello"})
		c2 := corestr.New.Collection.Strings([]string{"hello"})
		actual := args.Map{"eq": c1.IsEqualsWithSensitive(false, c2), "neq": c1.IsEqualsWithSensitive(true, c2)}
		expected := args.Map{"eq": true, "neq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- IsEqualsWithSensitive", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — AsError / ToError
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsError_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		actual := args.Map{"nil": c.AsError(",") == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AsError empty", actual)
	})
}

func Test_I31_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsError", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsError(", ")
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsError", actual)
	})
}

func Test_I31_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		err := c.AsDefaultError()
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsDefaultError", actual)
	})
}

func Test_I31_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		err := c.ToError(", ")
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToError", actual)
	})
}

func Test_I31_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		err := c.ToDefaultError()
		actual := args.Map{"notNil": err != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToDefaultError", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Append variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(other)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendCollectionPtr", actual)
	})
}

func Test_I31_Collection_AppendCollections_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendCollections_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendCollections empty", actual)
	})
}

func Test_I31_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys(42, nil, "hello")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnys", actual)
	})
}

func Test_I31_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnys_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendAnys empty", actual)
	})
}

func Test_I31_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock(42)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysLock", actual)
	})
}

func Test_I31_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnysLock_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendAnysLock empty", actual)
	})
}

func Test_I31_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys("hello", nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendNonEmptyAnys", actual)
	})
}

func Test_I31_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendNonEmptyAnys_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys(nil...)
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AppendNonEmptyAnys nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — RemoveItemsIndexes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexes_NilIgnore", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.RemoveItemsIndexes(true, nil...)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexes nil ignore", actual)
	})
}

func Test_I31_Collection_RemoveItemsIndexesPtr(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexesPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexesPtr(true, []int{1})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveItemsIndexesPtr", actual)
	})
}

func Test_I31_Collection_RemoveItemsIndexesPtr_NilIndexes(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexesPtr_NilIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexesPtr(true, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexesPtr nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Resize / AddCapacity
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Resize(100)
		actual := args.Map{"cap": c.Capacity() >= 100, "len": c.Length()}
		expected := args.Map{"cap": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Resize", actual)
	})
}

func Test_I31_Collection_Resize_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Resize_AlreadyLarger", func() {
		c := corestr.New.Collection.Cap(100)
		c.Resize(5)
		actual := args.Map{"cap": c.Capacity() >= 100}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Resize already larger", actual)
	})
}

func Test_I31_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(50, 50)
		actual := args.Map{"cap": c.Capacity() >= 100}
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCapacity", actual)
	})
}

func Test_I31_Collection_AddCapacity_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCapacity_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(nil...)
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddCapacity nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — List / ListStrings / Items
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_List(t *testing.T) {
	safeTest(t, "Test_I31_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"lLen": len(c.List()), "lsLen": len(c.ListStrings()), "lspLen": len(c.ListStringsPtr()), "iLen": len(c.Items()), "lpLen": len(c.ListPtr())}
		expected := args.Map{"lLen": 1, "lsLen": 1, "lspLen": 1, "iLen": 1, "lpLen": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- List variants", actual)
	})
}

func Test_I31_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		cp := c.ListCopyPtrLock()
		actual := args.Map{"len": len(cp)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ListCopyPtrLock", actual)
	})
}

func Test_I31_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ListCopyPtrLock_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		cp := c.ListCopyPtrLock()
		actual := args.Map{"len": len(cp)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- ListCopyPtrLock empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — JsonString / StringJSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_I31_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_I31_Collection_JsonString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"notEmpty": c.JsonString() != "", "must": c.JsonStringMust() != "", "sj": c.StringJSON() != ""}
		expected := args.Map{"notEmpty": true, "must": true, "sj": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonString variants", actual)
	})
}
