package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 6: Basic ops, Add variants, Has/Missing, Filter (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovHS1_01_IsEmpty_HasItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_01_IsEmpty_HasItems", func() {
		hs := corestr.New.Hashset.Empty()
		actual := args.Map{"result": hs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": hs.HasItems()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		hs.Add("a")
		actual := args.Map{"result": hs.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual := args.Map{"result": hs.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
	})
}

func Test_CovHS1_02_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_02_IsEmptyLock", func() {
		hs := corestr.New.Hashset.Empty()
		actual := args.Map{"result": hs.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		hs.Add("x")
		actual := args.Map{"result": hs.IsEmptyLock()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovHS1_03_AddCapacities(t *testing.T) {
	safeTest(t, "Test_CovHS1_03_AddCapacities", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// with capacities
		r := hs.AddCapacities(10, 20)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty capacities
		r2 := hs.AddCapacities()
		actual := args.Map{"result": r2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_04_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_04_AddCapacitiesLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.AddCapacitiesLock(10)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		r2 := hs.AddCapacitiesLock()
		actual := args.Map{"result": r2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_05_Resize(t *testing.T) {
	safeTest(t, "Test_CovHS1_05_Resize", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// capacity > length
		r := hs.Resize(100)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// capacity < length (no-op)
		r2 := hs.Resize(0)
		actual := args.Map{"result": r2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_06_ResizeLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_06_ResizeLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.ResizeLock(100)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		r2 := hs.ResizeLock(0)
		actual := args.Map{"result": r2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_07_Collection(t *testing.T) {
	safeTest(t, "Test_CovHS1_07_Collection", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		col := hs.Collection()
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_08_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_CovHS1_08_ConcatNewHashsets", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty hashsets — clone
		r := a.ConcatNewHashsets(true)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with hashsets
		b := corestr.New.Hashset.Empty()
		b.Add("b")
		r2 := a.ConcatNewHashsets(false, b, nil)
		actual := args.Map{"result": r2.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovHS1_09_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_CovHS1_09_ConcatNewStrings", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty
		r := a.ConcatNewStrings(true)
		actual := args.Map{"result": r.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// with strings
		r2 := a.ConcatNewStrings(false, []string{"b", "c"})
		actual := args.Map{"result": r2.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovHS1_10_AddPtr(t *testing.T) {
	safeTest(t, "Test_CovHS1_10_AddPtr", func() {
		hs := corestr.New.Hashset.Empty()
		s := "hello"
		hs.AddPtr(&s)
		actual := args.Map{"result": hs.Has("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has hello", actual)
	})
}

func Test_CovHS1_11_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_11_AddPtrLock", func() {
		hs := corestr.New.Hashset.Empty()
		s := "x"
		hs.AddPtrLock(&s)
		actual := args.Map{"result": hs.Has("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has x", actual)
	})
}

func Test_CovHS1_12_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_12_AddWithWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
	})
}

func Test_CovHS1_13_Add_AddBool(t *testing.T) {
	safeTest(t, "Test_CovHS1_13_Add_AddBool", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// AddBool — new
		existed := hs.AddBool("b")
		actual := args.Map{"result": existed}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for new", actual)
		// AddBool — existing
		existed2 := hs.AddBool("a")
		actual := args.Map{"result": existed2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for existing", actual)
	})
}

func Test_CovHS1_14_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovHS1_14_AddNonEmpty", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.AddNonEmpty("a")
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_15_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_CovHS1_15_AddNonEmptyWhitespace", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("  ")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_16_AddIf(t *testing.T) {
	safeTest(t, "Test_CovHS1_16_AddIf", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(false, "skip")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.AddIf(true, "keep")
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_17_AddIfMany(t *testing.T) {
	safeTest(t, "Test_CovHS1_17_AddIfMany", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(false, "a", "b")
		actual := args.Map{"result": hs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		hs.AddIfMany(true, "a", "b")
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS1_18_AddFunc(t *testing.T) {
	safeTest(t, "Test_CovHS1_18_AddFunc", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddFunc(func() string { return "generated" })
		actual := args.Map{"result": hs.Has("generated")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has generated", actual)
	})
}

func Test_CovHS1_19_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovHS1_19_AddFuncErr", func() {
		hs := corestr.New.Hashset.Empty()
		// success
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)
		actual := args.Map{"result": hs.Has("ok")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has ok", actual)
		// error
		hs.AddFuncErr(
			func() (string, error) { return "", fmt.Errorf("err") },
			func(err error) {},
		)
	})
}

func Test_CovHS1_20_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_20_AddStringsPtrWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS1_21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_21_AddHashsetItems", func() {
		hs := corestr.New.Hashset.Empty()
		other := corestr.New.Hashset.Empty()
		other.Add("x")
		hs.AddHashsetItems(other)
		actual := args.Map{"result": hs.Has("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has x", actual)
		// nil
		hs.AddHashsetItems(nil)
	})
}

func Test_CovHS1_22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_CovHS1_22_AddItemsMap", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
		actual := args.Map{"result": hs.Has("b")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have b (disabled)", actual)
		// nil
		hs.AddItemsMap(nil)
	})
}

func Test_CovHS1_23_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_23_AddItemsMapWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"x": true, "y": false}
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()
		actual := args.Map{"result": hs.Has("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		// nil
		hs.AddItemsMapWgLock(nil, wg)
	})
}

func Test_CovHS1_24_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_24_AddHashsetWgLock", func() {
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		other := corestr.New.Hashset.Empty()
		other.Add("z")
		wg.Add(1)
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()
		actual := args.Map{"result": hs.Has("z")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected z", actual)
		// nil
		hs.AddHashsetWgLock(nil, wg)
	})
}

func Test_CovHS1_25_AddStrings_Adds(t *testing.T) {
	safeTest(t, "Test_CovHS1_25_AddStrings_Adds", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddStrings([]string{"a", "b"})
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		hs.AddStrings(nil)
		hs.Adds("c")
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		hs.Adds()
	})
}

func Test_CovHS1_26_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovHS1_26_AddSimpleSlice", func() {
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		hs.AddSimpleSlice(ss)
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		hs.AddSimpleSlice(corestr.Empty.SimpleSlice())
	})
}

func Test_CovHS1_27_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_27_AddStringsLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddStringsLock([]string{"a"})
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hs.AddStringsLock(nil)
	})
}

func Test_CovHS1_28_AddCollection_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovHS1_28_AddCollection_AddCollections", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hs.AddCollection(nil)
		hs.AddCollection(corestr.Empty.Collection())

		hs.AddCollections(corestr.New.Collection.Strings([]string{"b"}), nil)
		actual := args.Map{"result": hs.Has("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		hs.AddCollections()
	})
}

func Test_CovHS1_29_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHS1_29_AddsAnyUsingFilter", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilter(filter, "a", nil, "b")
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil anys
		hs.AddsAnyUsingFilter(filter)
		// break
		hs2 := corestr.New.Hashset.Empty()
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs2.AddsAnyUsingFilter(breakFilter, "a", "b")
		actual := args.Map{"result": hs2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break)", actual)
	})
}

func Test_CovHS1_30_AddsAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_30_AddsAnyUsingFilterLock", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilterLock(filter, "a", nil)
		actual := args.Map{"result": hs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// nil
		hs.AddsAnyUsingFilterLock(filter)
		// break
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs.AddsAnyUsingFilterLock(breakFilter, "x", "y")
	})
}

func Test_CovHS1_31_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovHS1_31_AddsUsingFilter", func() {
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, str != "skip", false
		}
		hs.AddsUsingFilter(filter, "a", "skip", "b")
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		hs.AddsUsingFilter(filter)
		// break
		breakFilter := func(str string, index int) (string, bool, bool) {
			return str, true, true
		}
		hs2 := corestr.New.Hashset.Empty()
		hs2.AddsUsingFilter(breakFilter, "a", "b")
		actual := args.Map{"result": hs2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_32_AddLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_32_AddLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.AddLock("a")
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovHS1_33_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovHS1_33_HasAnyItem", func() {
		hs := corestr.New.Hashset.Empty()
		actual := args.Map{"result": hs.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		hs.Add("a")
		actual := args.Map{"result": hs.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovHS1_34_IsMissing_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_34_IsMissing_IsMissingLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		actual := args.Map{"result": hs.IsMissing("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		actual := args.Map{"result": hs.IsMissing("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected missing", actual)
		actual := args.Map{"result": hs.IsMissingLock("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
	})
}

func Test_CovHS1_35_Has_Contains_HasLock_HasWithLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_35_Has_Contains_HasLock_HasWithLock", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		actual := args.Map{"result": hs.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.Contains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasWithLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.Has("missing")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_36_HasAllStrings_HasAll(t *testing.T) {
	safeTest(t, "Test_CovHS1_36_HasAllStrings_HasAll", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")
		actual := args.Map{"result": hs.HasAllStrings([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAllStrings([]string{"a", "x"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": hs.HasAll("a", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_37_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_37_HasAllCollectionItems", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")
		actual := args.Map{"result": hs.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a", "b"}))}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAllCollectionItems(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		actual := args.Map{"result": hs.HasAllCollectionItems(corestr.Empty.Collection())}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_CovHS1_38_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_CovHS1_38_IsAllMissing", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		actual := args.Map{"result": hs.IsAllMissing("x", "y")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.IsAllMissing("a", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_39_HasAny(t *testing.T) {
	safeTest(t, "Test_CovHS1_39_HasAny", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		actual := args.Map{"result": hs.HasAny("x", "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_40_IsEqual_IsEquals(t *testing.T) {
	safeTest(t, "Test_CovHS1_40_IsEqual_IsEquals", func() {
		a := corestr.New.Hashset.Empty()
		a.Adds("a", "b")
		b := corestr.New.Hashset.Empty()
		b.Adds("a", "b")
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// same ptr
		actual := args.Map{"result": a.IsEqual(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal to self", actual)
		// nil
		actual := args.Map{"result": a.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		// both empty
		e1 := corestr.New.Hashset.Empty()
		e2 := corestr.New.Hashset.Empty()
		actual := args.Map{"result": e1.IsEquals(e2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty == empty", actual)
		// one empty
		actual := args.Map{"result": a.IsEquals(e1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.New.Hashset.Empty()
		c.Add("a")
		actual := args.Map{"result": a.IsEquals(c)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for diff length", actual)
		// same length, diff keys
		d := corestr.New.Hashset.Empty()
		d.Adds("a", "z")
		actual := args.Map{"result": a.IsEquals(d)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for diff keys", actual)
	})
}

func Test_CovHS1_41_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_41_IsEqualsLock", func() {
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		b := corestr.New.Hashset.Empty()
		b.Add("a")
		actual := args.Map{"result": a.IsEqualsLock(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_CovHS1_42_SortedList(t *testing.T) {
	safeTest(t, "Test_CovHS1_42_SortedList", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		sorted := hs.SortedList()
		actual := args.Map{"result": sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted asc", actual)
	})
}

func Test_CovHS1_43_Filter(t *testing.T) {
	safeTest(t, "Test_CovHS1_43_Filter", func() {
		hs := corestr.New.Hashset.Empty()
		hs.Adds("apple", "banana", "avocado")
		result := hs.Filter(func(s string) bool {
			return s[0] == 'a'
		})
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}
