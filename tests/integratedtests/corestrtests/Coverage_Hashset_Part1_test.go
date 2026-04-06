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
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.IsEmpty()}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.IsEmptyLock()}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// with capacities
		r := hs.AddCapacities(10, 20)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.AddCapacitiesLock(10)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// capacity > length
		r := hs.Resize(100)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		r := hs.ResizeLock(100)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		col := hs.Collection()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovHS1_08_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_CovHS1_08_ConcatNewHashsets", func() {
		// Arrange
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty hashsets — clone
		r := a.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		// empty
		r := a.ConcatNewStrings(true)

		// Act
		actual := args.Map{"result": r.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := "hello"
		hs.AddPtr(&s)

		// Act
		actual := args.Map{"result": hs.Has("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has hello", actual)
	})
}

func Test_CovHS1_11_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_11_AddPtrLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := "x"
		hs.AddPtrLock(&s)

		// Act
		actual := args.Map{"result": hs.Has("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has x", actual)
	})
}

func Test_CovHS1_12_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_12_AddWithWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has a", actual)
	})
}

func Test_CovHS1_13_Add_AddBool(t *testing.T) {
	safeTest(t, "Test_CovHS1_13_Add_AddBool", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")
		// AddBool — new
		existed := hs.AddBool("b")

		// Act
		actual := args.Map{"result": existed}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("  ")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(false, "a", "b")

		// Act
		actual := args.Map{"result": hs.Length() != 0}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddFunc(func() string { return "generated" })

		// Act
		actual := args.Map{"result": hs.Has("generated")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has generated", actual)
	})
}

func Test_CovHS1_19_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_CovHS1_19_AddFuncErr", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		// success
		hs.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovHS1_21_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_CovHS1_21_AddHashsetItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		other := corestr.New.Hashset.Empty()
		other.Add("x")
		hs.AddHashsetItems(other)

		// Act
		actual := args.Map{"result": hs.Has("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected has x", actual)
		// nil
		hs.AddHashsetItems(nil)
	})
}

func Test_CovHS1_22_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_CovHS1_22_AddItemsMap", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"x": true, "y": false}
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": hs.Has("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		// nil
		hs.AddItemsMapWgLock(nil, wg)
	})
}

func Test_CovHS1_24_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_24_AddHashsetWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		other := corestr.New.Hashset.Empty()
		other.Add("z")
		wg.Add(1)
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()

		// Act
		actual := args.Map{"result": hs.Has("z")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected z", actual)
		// nil
		hs.AddHashsetWgLock(nil, wg)
	})
}

func Test_CovHS1_25_AddStrings_Adds(t *testing.T) {
	safeTest(t, "Test_CovHS1_25_AddStrings_Adds", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		hs.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		hs.AddSimpleSlice(corestr.Empty.SimpleSlice())
	})
}

func Test_CovHS1_27_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_CovHS1_27_AddStringsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStringsLock([]string{"a"})

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		hs.AddStringsLock(nil)
	})
}

func Test_CovHS1_28_AddCollection_AddCollections(t *testing.T) {
	safeTest(t, "Test_CovHS1_28_AddCollection_AddCollections", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCollection(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilter(filter, "a", nil, "b")

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, true, false
		}
		hs.AddsAnyUsingFilterLock(filter, "a", nil)

		// Act
		actual := args.Map{"result": hs.Length() != 1}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(str string, index int) (string, bool, bool) {
			return str, str != "skip", false
		}
		hs.AddsUsingFilter(filter, "a", "skip", "b")

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddLock("a")

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovHS1_33_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovHS1_33_HasAnyItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"result": hs.HasAnyItem()}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"result": hs.IsMissing("a")}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"result": hs.Has("a")}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": hs.HasAllStrings([]string{"a", "b"})}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")

		// Act
		actual := args.Map{"result": hs.HasAllCollectionItems(corestr.New.Collection.Strings([]string{"a", "b"}))}

		// Assert
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
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"result": hs.IsAllMissing("x", "y")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.IsAllMissing("a", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_39_HasAny(t *testing.T) {
	safeTest(t, "Test_CovHS1_39_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"result": hs.HasAny("x", "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.HasAny("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovHS1_40_IsEqual_IsEquals(t *testing.T) {
	safeTest(t, "Test_CovHS1_40_IsEqual_IsEquals", func() {
		// Arrange
		a := corestr.New.Hashset.Empty()
		a.Adds("a", "b")
		b := corestr.New.Hashset.Empty()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
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
		// Arrange
		a := corestr.New.Hashset.Empty()
		a.Add("a")
		b := corestr.New.Hashset.Empty()
		b.Add("a")

		// Act
		actual := args.Map{"result": a.IsEqualsLock(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_CovHS1_42_SortedList(t *testing.T) {
	safeTest(t, "Test_CovHS1_42_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("c", "a", "b")
		sorted := hs.SortedList()

		// Act
		actual := args.Map{"result": sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted asc", actual)
	})
}

func Test_CovHS1_43_Filter(t *testing.T) {
	safeTest(t, "Test_CovHS1_43_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("apple", "banana", "avocado")
		result := hs.Filter(func(s string) bool {
			return s[0] == 'a'
		})

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}
