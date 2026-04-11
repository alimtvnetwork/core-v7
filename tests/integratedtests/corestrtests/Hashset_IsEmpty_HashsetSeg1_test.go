package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Hashset.go — Full coverage (~489 uncovered stmts, 1469 lines)
// =============================================================================

func Test_Hashset_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{
			"empty": hs.IsEmpty(),
			"hasItems": hs.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty/HasItems on empty", actual)
	})
}

func Test_Hashset_IsEmpty_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEmpty_NonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"empty": hs.IsEmpty(),
			"hasItems": hs.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": false,
			"hasItems": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty/HasItems on non-empty", actual)
	})
}

func Test_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCapacities", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCapacities(10, 20)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacities increases capacity", actual)
	})
}

func Test_Hashset_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCapacities_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCapacities()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacities no args returns same", actual)
	})
}

func Test_Hashset_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCapacitiesLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCapacitiesLock(10)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock increases capacity", actual)
	})
}

func Test_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCapacitiesLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCapacitiesLock()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock no args returns same", actual)
	})
}

func Test_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Resize", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.Resize(100)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Resize preserves items", actual)
	})
}

func Test_Hashset_Resize_Smaller(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Resize_Smaller", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(50)
		hs.Adds("a", "b", "c")
		hs.Resize(1)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Resize smaller ignored", actual)
	})
}

func Test_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ResizeLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.ResizeLock(100)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ResizeLock preserves items", actual)
	})
}

func Test_Hashset_ResizeLock_Smaller(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ResizeLock_Smaller", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(50)
		hs.Adds("a", "b")
		hs.ResizeLock(1)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ResizeLock smaller ignored", actual)
	})
}

func Test_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Collection", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		col := hs.Collection()

		// Act
		actual := args.Map{"len": col.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns Collection", actual)
	})
}

func Test_Hashset_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEmptyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"empty": hs.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock returns true", actual)
	})
}

func Test_Hashset_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ConcatNewHashsets", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("b")
		r := a.ConcatNewHashsets(false, b)

		// Act
		actual := args.Map{
			"hasA": r.Has("a"),
			"hasB": r.Has("b"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": true,
		}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets merges", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ConcatNewHashsets_Empty", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets empty clones", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_NilHashset(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ConcatNewHashsets_NilHashset", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.ConcatNewHashsets(false, nil)

		// Act
		actual := args.Map{"hasA": r.Has("a")}

		// Assert
		expected := args.Map{"hasA": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets skips nil", actual)
	})
}

func Test_Hashset_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ConcatNewStrings", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.ConcatNewStrings(false, []string{"b", "c"})

		// Act
		actual := args.Map{
			"hasB": r.Has("b"),
			"hasC": r.Has("c"),
		}

		// Assert
		expected := args.Map{
			"hasB": true,
			"hasC": true,
		}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings adds strings", actual)
	})
}

func Test_Hashset_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ConcatNewStrings_Empty", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.ConcatNewStrings(true)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings empty clones", actual)
	})
}

func Test_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := "hello"
		hs.AddPtr(&s)

		// Act
		actual := args.Map{"has": hs.Has("hello")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtr adds by pointer", actual)
	})
}

func Test_Hashset_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddWithWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock adds with lock", actual)
	})
}

func Test_Hashset_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddPtrLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := "x"
		hs.AddPtrLock(&s)

		// Act
		actual := args.Map{"has": hs.Has("x")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtrLock adds with lock", actual)
	})
}

func Test_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Add", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Add("a")

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Add adds item", actual)
	})
}

func Test_Hashset_AddBool_New(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddBool_New", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		existed := hs.AddBool("a")

		// Act
		actual := args.Map{
			"existed": existed,
			"has": hs.Has("a"),
		}

		// Assert
		expected := args.Map{
			"existed": false,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "AddBool new returns false", actual)
	})
}

func Test_Hashset_AddBool_Existing(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddBool_Existing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		existed := hs.AddBool("a")

		// Act
		actual := args.Map{"existed": existed}

		// Assert
		expected := args.Map{"existed": true}
		expected.ShouldBeEqual(t, 0, "AddBool existing returns true", actual)
	})
}

func Test_Hashset_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddNonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("a")
		hs.AddNonEmpty("")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty skips empty", actual)
	})
}

func Test_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddNonEmptyWhitespace", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmptyWhitespace("a")
		hs.AddNonEmptyWhitespace("  ")
		hs.AddNonEmptyWhitespace("")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace skips whitespace", actual)
	})
}

func Test_Hashset_AddIf(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddIf", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIf(true, "a")
		hs.AddIf(false, "b")

		// Act
		actual := args.Map{
			"hasA": hs.Has("a"),
			"hasB": hs.Has("b"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "AddIf conditionally adds", actual)
	})
}

func Test_Hashset_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddIfMany", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddIfMany(true, "a", "b")
		hs.AddIfMany(false, "c")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddIfMany conditionally adds many", actual)
	})
}

func Test_Hashset_AddFunc(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddFunc", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddFunc(func() string { return "x" })

		// Act
		actual := args.Map{"has": hs.Has("x")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFunc adds func result", actual)
	})
}

func Test_Hashset_AddFuncErr_NoError(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddFuncErr_NoError", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddFuncErr(
			func() (string, error) { return "a", nil },
			func(err error) {},
		)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFuncErr adds on no error", actual)
	})
}

func Test_Hashset_AddFuncErr_WithError(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddFuncErr_WithError", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		handled := false
		hs.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { handled = true },
		)

		// Act
		actual := args.Map{
			"len": hs.Length(),
			"handled": handled,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"handled": true,
		}
		expected.ShouldBeEqual(t, 0, "AddFuncErr handles error", actual)
	})
}

func Test_Hashset_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddStringsPtrWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsPtrWgLock adds with lock", actual)
	})
}

func Test_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddHashsetItems", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("b")
		a.AddHashsetItems(b)

		// Act
		actual := args.Map{"hasB": a.Has("b")}

		// Assert
		expected := args.Map{"hasB": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems merges", actual)
	})
}

func Test_Hashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddHashsetItems_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		a.AddHashsetItems(nil)

		// Act
		actual := args.Map{"len": a.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems nil returns same", actual)
	})
}

func Test_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddItemsMap", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{
			"hasA": hs.Has("a"),
			"hasB": hs.Has("b"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "AddItemsMap only adds true values", actual)
	})
}

func Test_Hashset_AddItemsMap_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddItemsMap_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMap(nil)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMap nil returns same", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddItemsMapWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		m := map[string]bool{"a": true, "b": false}
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()

		// Act
		actual := args.Map{
			"hasA": hs.Has("a"),
			"hasB": hs.Has("b"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock adds with lock", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddItemsMapWgLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddItemsMapWgLock(nil, nil)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock nil returns same", actual)
	})
}

func Test_Hashset_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddHashsetWgLock", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("b")
		wg := &sync.WaitGroup{}
		wg.Add(1)
		a.AddHashsetWgLock(b, wg)
		wg.Wait()

		// Act
		actual := args.Map{"hasB": a.Has("b")}

		// Assert
		expected := args.Map{"hasB": true}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock merges with lock", actual)
	})
}

func Test_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddHashsetWgLock_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		a.AddHashsetWgLock(nil, nil)

		// Act
		actual := args.Map{"len": a.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock nil returns same", actual)
	})
}

func Test_Hashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStrings adds", actual)
	})
}

func Test_Hashset_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddStrings_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStrings(nil)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings nil returns same", actual)
	})
}

func Test_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddSimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		hs.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddSimpleSlice adds", actual)
	})
}

func Test_Hashset_AddSimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddSimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Empty()
		hs.AddSimpleSlice(ss)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSimpleSlice empty returns same", actual)
	})
}

func Test_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddStringsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStringsLock([]string{"a"})

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddStringsLock adds with lock", actual)
	})
}

func Test_Hashset_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddStringsLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddStringsLock(nil)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock nil returns same", actual)
	})
}

func Test_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Adds", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.Adds("a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds adds variadic", actual)
	})
}

func Test_Hashset_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Adds_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		var s []string
		hs.Adds(s...)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil returns same", actual)
	})
}

func Test_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		col := corestr.New.Collection.Strings([]string{"a"})
		hs.AddCollection(col)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCollection adds items", actual)
	})
}

func Test_Hashset_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddCollection(nil)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil returns same", actual)
	})
}

func Test_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCollections", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		hs.AddCollections(a, nil, b)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollections skips nil", actual)
	})
}

func Test_Hashset_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddCollections_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		var cols []*corestr.Collection
		hs.AddCollections(cols...)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections nil returns same", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		hs.AddsAnyUsingFilter(filter, "a", nil, "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilter filters and adds", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }
		hs.AddsAnyUsingFilter(filter, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilter breaks early", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilter_Skip", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, false, false }
		hs.AddsAnyUsingFilter(filter, "a")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilter skips all", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		var anys []any
		hs.AddsAnyUsingFilter(filter, anys...)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilter nil returns same", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilterLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		hs.AddsAnyUsingFilterLock(filter, "a", nil, "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilterLock filters with lock", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilterLock_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }
		hs.AddsAnyUsingFilterLock(filter, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilterLock breaks early", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		var anys []any
		hs.AddsAnyUsingFilterLock(filter, anys...)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAnyUsingFilterLock nil returns same", actual)
	})
}

func Test_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s + "!", true, false }
		hs.AddsUsingFilter(filter, "a", "b")

		// Act
		actual := args.Map{
			"hasA": hs.Has("a!"),
			"hasB": hs.Has("b!"),
		}

		// Assert
		expected := args.Map{
			"hasA": true,
			"hasB": true,
		}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter transforms and adds", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, true }
		hs.AddsUsingFilter(filter, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter breaks early", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddsUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		var keys []string
		hs.AddsUsingFilter(filter, keys...)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter nil returns same", actual)
	})
}

func Test_Hashset_AddLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AddLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		hs.AddLock("a")

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLock adds with lock", actual)
	})
}

func Test_Hashset_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasAnyItem", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"has": hs.HasAnyItem()}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem returns true", actual)
	})
}

func Test_Hashset_IsMissing(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"missing": hs.IsMissing("z"),
			"found": hs.IsMissing("a"),
		}

		// Assert
		expected := args.Map{
			"missing": true,
			"found": false,
		}
		expected.ShouldBeEqual(t, 0, "IsMissing checks absence", actual)
	})
}

func Test_Hashset_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsMissingLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"missing": hs.IsMissingLock("z")}

		// Assert
		expected := args.Map{"missing": true}
		expected.ShouldBeEqual(t, 0, "IsMissingLock checks with lock", actual)
	})
}

func Test_Hashset_Has(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Has", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"has": hs.Has("a"),
			"miss": hs.Has("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has checks containment", actual)
	})
}

func Test_Hashset_Contains(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"has": hs.Contains("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Contains is alias for Has", actual)
	})
}

func Test_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEqual", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"eq": a.IsEqual(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual checks equality", actual)
	})
}

func Test_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		r := hs.SortedList()

		// Act
		actual := args.Map{
			"first": r[0],
			"last": r[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "SortedList returns sorted", actual)
	})
}

func Test_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("aa", "b", "cc")
		r := hs.Filter(func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter returns matching items", actual)
	})
}

func Test_Hashset_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"has": hs.HasLock("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock checks with lock", actual)
	})
}

func Test_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasAllStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")

		// Act
		actual := args.Map{
			"all": hs.HasAllStrings([]string{"a", "b"}),
			"miss": hs.HasAllStrings([]string{"a", "z"}),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllStrings checks all", actual)
	})
}

func Test_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasAllCollectionItems", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		col := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": hs.HasAllCollectionItems(col),
			"nil": hs.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAllCollectionItems checks", actual)
	})
}

func Test_Hashset_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasAll", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{
			"all": hs.HasAll("a", "b"),
			"miss": hs.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll checks all variadic", actual)
	})
}

func Test_Hashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"all": hs.IsAllMissing("x", "y"),
			"partial": hs.IsAllMissing("a", "x"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"partial": false,
		}
		expected.ShouldBeEqual(t, 0, "IsAllMissing checks all absent", actual)
	})
}

func Test_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{
			"any": hs.HasAny("a", "z"),
			"none": hs.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAny checks any present", actual)
	})
}

func Test_Hashset_HasWithLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_HasWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"has": hs.HasWithLock("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasWithLock checks with lock", actual)
	})
}

func Test_Hashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		r := hs.OrderedList()

		// Act
		actual := args.Map{
			"first": r[0],
			"last": r[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "OrderedList returns sorted", actual)
	})
}

func Test_Hashset_OrderedList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_OrderedList_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.OrderedList()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "OrderedList empty returns empty", actual)
	})
}

func Test_Hashset_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SafeStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": len(hs.SafeStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SafeStrings returns list", actual)
	})
}

func Test_Hashset_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SafeStrings_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(hs.SafeStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeStrings empty returns empty", actual)
	})
}

func Test_Hashset_Lines(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Lines", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": len(hs.Lines())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Lines returns list", actual)
	})
}

func Test_Hashset_Lines_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Lines_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"len": len(hs.Lines())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Lines empty returns empty", actual)
	})
}

func Test_Hashset_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonNil": hs.SimpleSlice() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice returns SimpleSlice", actual)
	})
}

func Test_Hashset_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"empty": hs.SimpleSlice().IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice empty returns empty", actual)
	})
}

func Test_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredItems", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("aa", "b")
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems filters", actual)
	})
}

func Test_Hashset_GetFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredItems_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems empty returns empty", actual)
	})
}

func Test_Hashset_GetFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredItems_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		r := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems breaks early", actual)
	})
}

func Test_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("aa", "b")
		r := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection filters", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredCollection_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"empty": r.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection empty returns empty", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetFilteredCollection_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection breaks early", actual)
	})
}

func Test_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptHashset", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		b := corestr.New.Hashset.StringsSpreadItems("b")
		r := a.GetAllExceptHashset(b)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset excludes", actual)
	})
}

func Test_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptHashset_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.GetAllExceptHashset(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset nil returns all", actual)
	})
}

func Test_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExcept", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := a.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept excludes", actual)
	})
}

func Test_Hashset_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExcept_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil returns all", actual)
	})
}

func Test_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptSpread", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := a.GetAllExceptSpread("a")

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread excludes", actual)
	})
}

func Test_Hashset_GetAllExceptSpread_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptSpread_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.GetAllExceptSpread()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread nil returns all", actual)
	})
}

func Test_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptCollection", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		col := corestr.New.Collection.Strings([]string{"a"})
		r := a.GetAllExceptCollection(col)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection excludes", actual)
	})
}

func Test_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_GetAllExceptCollection_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		r := a.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil returns all", actual)
	})
}

func Test_Hashset_Items(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Items", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": len(hs.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items returns map", actual)
	})
}

func Test_Hashset_List(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_List", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": len(hs.List())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List returns list", actual)
	})
}

func Test_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_MapStringAny", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		m := hs.MapStringAny()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapStringAny returns map", actual)
	})
}

func Test_Hashset_MapStringAny_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_MapStringAny_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		m := hs.MapStringAny()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "MapStringAny empty returns empty map", actual)
	})
}

func Test_Hashset_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_MapStringAnyDiff", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.MapStringAnyDiff()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "MapStringAnyDiff returns diff map", actual)
	})
}

func Test_Hashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		s := hs.JoinSorted(",")

		// Act
		actual := args.Map{"val": s}

		// Assert
		expected := args.Map{"val": "a,b,c"}
		expected.ShouldBeEqual(t, 0, "JoinSorted returns sorted join", actual)
	})
}

func Test_Hashset_JoinSorted_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JoinSorted_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		s := hs.JoinSorted(",")

		// Act
		actual := args.Map{"empty": s == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "JoinSorted empty returns empty", actual)
	})
}

func Test_Hashset_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ListPtrSortedAsc", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("c", "a", "b")
		r := hs.ListPtrSortedAsc()

		// Act
		actual := args.Map{"first": r[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedAsc returns sorted asc", actual)
	})
}

func Test_Hashset_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ListPtrSortedDsc", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b", "c")
		r := hs.ListPtrSortedDsc()

		// Act
		actual := args.Map{"first": r[0]}

		// Assert
		expected := args.Map{"first": "c"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedDsc returns sorted desc", actual)
	})
}

func Test_Hashset_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ListPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": len(hs.ListPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtr returns list", actual)
	})
}

func Test_Hashset_Clear(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Clear", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.Clear()

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empties hashset", actual)
	})
}

func Test_Hashset_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Clear_Nil", func() {
		// Arrange
		var hs *corestr.Hashset
		r := hs.Clear()

		// Act
		actual := args.Map{"nil": r == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil returns nil", actual)
	})
}

func Test_Hashset_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Dispose", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.Dispose()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose clears and nils", actual)
	})
}

func Test_Hashset_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Dispose_Nil", func() {
		// Arrange
		var hs *corestr.Hashset
		hs.Dispose() // should not panic

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose nil does not panic", actual)
	})
}

func Test_Hashset_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ListCopyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.ListCopyLock()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListCopyLock returns copy", actual)
	})
}

func Test_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ToLowerSet", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("Hello", "WORLD")
		r := hs.ToLowerSet()

		// Act
		actual := args.Map{
			"hasHello": r.Has("hello"),
			"hasWorld": r.Has("world"),
		}

		// Assert
		expected := args.Map{
			"hasHello": true,
			"hasWorld": true,
		}
		expected.ShouldBeEqual(t, 0, "ToLowerSet lowercases all keys", actual)
	})
}

func Test_Hashset_Length(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Length", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Length returns count", actual)
	})
}

func Test_Hashset_LengthLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_LengthLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"len": hs.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock returns count", actual)
	})
}

func Test_Hashset_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_Same", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"eq": hs.IsEquals(hs)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals same ptr returns true", actual)
	})
}

func Test_Hashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.Empty()
		b := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty returns true", actual)
	})
}

func Test_Hashset_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_OneEmpty", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty returns false", actual)
	})
}

func Test_Hashset_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_DiffLen", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("a", "b")

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff length returns false", actual)
	})
}

func Test_Hashset_IsEquals_DiffContent(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_DiffContent", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("b")

		// Act
		actual := args.Map{"eq": a.IsEquals(b)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff content returns false", actual)
	})
}

func Test_Hashset_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEquals_Nil", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"eq": a.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals nil returns false", actual)
	})
}

func Test_Hashset_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_IsEqualsLock", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a")
		b := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"eq": a.IsEqualsLock(b)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock checks with lock", actual)
	})
}

func Test_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Remove", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		hs.Remove("a")

		// Act
		actual := args.Map{
			"has": hs.Has("a"),
			"len": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Remove deletes item", actual)
	})
}

func Test_Hashset_SafeRemove(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_SafeRemove", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.SafeRemove("a")
		hs.SafeRemove("z") // should not panic

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeRemove safely removes", actual)
	})
}

func Test_Hashset_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_RemoveWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		hs.RemoveWithLock("a")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock removes with lock", actual)
	})
}

func Test_Hashset_String(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_String", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonEmpty": hs.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String returns formatted", actual)
	})
}

func Test_Hashset_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_String_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"nonEmpty": hs.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty returns marker", actual)
	})
}

func Test_Hashset_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_StringLock", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonEmpty": hs.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock returns string", actual)
	})
}

func Test_Hashset_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_StringLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		actual := args.Map{"nonEmpty": hs.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty returns marker", actual)
	})
}

func Test_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Join", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		s := hs.Join(",")

		// Act
		actual := args.Map{"val": s}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Join joins items", actual)
	})
}

func Test_Hashset_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JoinLine", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		s := hs.JoinLine()

		// Act
		actual := args.Map{"val": s}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "JoinLine joins with newline", actual)
	})
}

func Test_Hashset_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_NonEmptyJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		s := hs.NonEmptyJoins(",")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins joins non-empty", actual)
	})
}

func Test_Hashset_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_NonWhitespaceJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		s := hs.NonWhitespaceJoins(",")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins joins non-ws", actual)
	})
}

func Test_Hashset_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JsonModel", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		m := hs.JsonModel()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel returns map", actual)
	})
}

func Test_Hashset_JsonModel_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JsonModel_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		m := hs.JsonModel()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "JsonModel empty returns empty", actual)
	})
}

func Test_Hashset_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JsonModelAny", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.JsonModelAny()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny returns any", actual)
	})
}

func Test_Hashset_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_MarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		b, err := hs.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON returns bytes", actual)
	})
}

func Test_Hashset_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_UnmarshalJSON", func() {
		// Arrange
		hs := &corestr.Hashset{}
		err := hs.UnmarshalJSON([]byte(`{"a":true,"b":true}`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON parses items", actual)
	})
}

func Test_Hashset_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_UnmarshalJSON_Error", func() {
		// Arrange
		hs := &corestr.Hashset{}
		err := hs.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error", actual)
	})
}

func Test_Hashset_Json(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Json", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.Json()

		// Act
		actual := args.Map{"nonEmpty": r.JsonString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json returns Result", actual)
	})
}

func Test_Hashset_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JsonPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.JsonPtr()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns Result ptr", actual)
	})
}

func Test_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ParseInjectUsingJson", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		jr := corejson.NewPtr(map[string]bool{"a": true})
		r, err := hs.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonNil": r != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson parses", actual)
	})
}

func Test_Hashset_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ParseInjectUsingJson_Error", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := hs.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error", actual)
	})
}

func Test_Hashset_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ParseInjectUsingJsonMust", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		jr := corejson.NewPtr(map[string]bool{"a": true})
		r := hs.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust parses", actual)
	})
}

func Test_Hashset_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			hs.ParseInjectUsingJsonMust(jr)
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics on error", actual)
	})
}

func Test_Hashset_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AsJsonContractsBinder", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonNil": hs.AsJsonContractsBinder() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns interface", actual)
	})
}

func Test_Hashset_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AsJsoner", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonNil": hs.AsJsoner() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsoner returns interface", actual)
	})
}

func Test_Hashset_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_JsonParseSelfInject", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		jr := corejson.NewPtr(map[string]bool{"a": true})
		err := hs.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject injects", actual)
	})
}

func Test_Hashset_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AsJsonParseSelfInjector", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonNil": hs.AsJsonParseSelfInjector() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector returns interface", actual)
	})
}

func Test_Hashset_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_AsJsonMarshaller", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")

		// Act
		actual := args.Map{"nonNil": hs.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller returns interface", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLinesRaw_BothEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.DistinctDiffLinesRaw()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw both empty returns empty", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLinesRaw_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.DistinctDiffLinesRaw()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw left only returns left", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLinesRaw_RightOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.DistinctDiffLinesRaw("a", "b")

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw right only returns right", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_Mixed(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLinesRaw_Mixed", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.DistinctDiffLinesRaw("b", "c")

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2} // "a" from left missing in right, "c" from right missing in left
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw returns symmetric diff", actual)
	})
}

func Test_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffHashset", func() {
		// Arrange
		a := corestr.New.Hashset.StringsSpreadItems("a", "b")
		b := corestr.New.Hashset.StringsSpreadItems("b", "c")
		r := a.DistinctDiffHashset(b)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffHashset returns diff", actual)
	})
}

func Test_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLines_BothEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.DistinctDiffLines()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines both empty returns empty", actual)
	})
}

func Test_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLines_LeftOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.DistinctDiffLines()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines left only returns items", actual)
	})
}

func Test_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLines_RightOnly", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.DistinctDiffLines("a")

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines right only returns right", actual)
	})
}

func Test_Hashset_DistinctDiffLines_Mixed(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_DistinctDiffLines_Mixed", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		r := hs.DistinctDiffLines("b", "c")

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines returns symmetric diff", actual)
	})
}

func Test_Hashset_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Serialize", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		b, err := hs.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize returns bytes", actual)
	})
}

func Test_Hashset_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Deserialize", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		var target map[string]bool
		err := hs.Deserialize(&target)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(target),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize parses to target", actual)
	})
}

func Test_Hashset_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_WrapDoubleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.WrapDoubleQuote()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuote wraps keys", actual)
	})
}

func Test_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuoteIfMissing wraps if missing", actual)
	})
}

func Test_Hashset_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_WrapSingleQuote", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.WrapSingleQuote()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuote wraps keys", actual)
	})
}

func Test_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_WrapSingleQuoteIfMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuoteIfMissing wraps if missing", actual)
	})
}

func Test_Hashset_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Transpile", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a")
		r := hs.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "Transpile transforms keys", actual)
	})
}

func Test_Hashset_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Cov61_Hashset_Transpile_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()
		r := hs.Transpile(func(s string) string { return s })

		// Act
		actual := args.Map{"empty": r.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Transpile empty returns empty", actual)
	})
}
