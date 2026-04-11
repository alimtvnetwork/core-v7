package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — IsEmpty / HasItems / Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_Hashset_IsEmpty_New(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEmpty_New", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{
			"empty": hs.IsEmpty(),
			"items": hs.HasItems(),
			"len": hs.Length(),
			"hasAny": hs.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"items": false,
			"len": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- empty", actual)
	})
}

func Test_Hashset_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Length_Nil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		actual := args.Map{
			"len": hs.Length(),
			"empty": hs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- nil length", actual)
	})
}

func Test_Hashset_IsEmptyLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEmptyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"empty": hs.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEmptyLock", actual)
	})
}

func Test_Hashset_LengthLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_LengthLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")

		// Act
		actual := args.Map{"len": hs.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- LengthLock", actual)
	})
}

func Test_Hashset_Add_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Add", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")

		// Act
		actual := args.Map{
			"has": hs.Has("a"),
			"len": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Add", actual)
	})
}

func Test_Hashset_AddBool_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddBool", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		existed1 := hs.AddBool("a")
		existed2 := hs.AddBool("a")

		// Act
		actual := args.Map{
			"existed1": existed1,
			"existed2": existed2,
		}

		// Assert
		expected := args.Map{
			"existed1": false,
			"existed2": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddBool", actual)
	})
}

func Test_Hashset_AddNonEmpty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddNonEmpty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddNonEmpty("")
		hs.AddNonEmpty("a")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddNonEmpty", actual)
	})
}

func Test_Hashset_AddNonEmptyWhitespace_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddNonEmptyWhitespace", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddNonEmptyWhitespace("   ")
		hs.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddNonEmptyWhitespace", actual)
	})
}

func Test_Hashset_AddIf_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddIf", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
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
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIf", actual)
	})
}

func Test_Hashset_AddIfMany_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddIfMany", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddIfMany(true, "a", "b")
		hs.AddIfMany(false, "c")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIfMany", actual)
	})
}

func Test_Hashset_AddFunc_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFunc", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"has": hs.Has("computed")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddFunc", actual)
	})
}

func Test_Hashset_AddFuncErr_NoErr_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFuncErr_NoErr", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"has": hs.Has("ok")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddFuncErr no err", actual)
	})
}

func Test_Hashset_AddFuncErr_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFuncErr_Err", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		called := false
		hs.AddFuncErr(func() (string, error) { return "", fmt.Errorf("fail") }, func(e error) { called = true })

		// Act
		actual := args.Map{
			"empty": hs.IsEmpty(),
			"called": called,
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"called": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- AddFuncErr err", actual)
	})
}

func Test_Hashset_AddLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddLock("a")

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddLock", actual)
	})
}

func Test_Hashset_AddPtr_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		s := "hello"
		hs.AddPtr(&s)

		// Act
		actual := args.Map{"has": hs.Has("hello")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtr", actual)
	})
}

func Test_Hashset_AddPtrLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddPtrLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		s := "hello"
		hs.AddPtrLock(&s)

		// Act
		actual := args.Map{"has": hs.Has("hello")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtrLock", actual)
	})
}

func Test_Hashset_Adds_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Adds", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds("a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Adds", actual)
	})
}

func Test_Hashset_Adds_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Adds_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds(nil...)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Adds nil", actual)
	})
}

func Test_Hashset_AddStrings_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStrings", actual)
	})
}

func Test_Hashset_AddStrings_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStrings_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStrings(nil)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddStrings nil", actual)
	})
}

func Test_Hashset_AddStringsLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock([]string{"a"})

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStringsLock", actual)
	})
}

func Test_Hashset_AddStringsLock_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock(nil)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddStringsLock nil", actual)
	})
}

func Test_Hashset_AddHashsetItems_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.AddHashsetItems(other)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashsetItems", actual)
	})
}

func Test_Hashset_AddHashsetItems_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetItems_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddHashsetItems(nil)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetItems nil", actual)
	})
}

func Test_Hashset_AddItemsMap_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMap", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
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
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMap", actual)
	})
}

func Test_Hashset_AddItemsMap_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMap_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddItemsMap(nil)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddItemsMap nil", actual)
	})
}

func Test_Hashset_AddCollection_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		coll := corestr.New.Collection.Strings([]string{"a"})
		hs.AddCollection(coll)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCollection", actual)
	})
}

func Test_Hashset_AddCollection_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddCollection(nil)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddCollection nil", actual)
	})
}

func Test_Hashset_AddCollections_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollections", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		hs.AddCollections(c1, nil, c2)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCollections", actual)
	})
}

func Test_Hashset_AddCollections_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollections_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddCollections(nil...)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddCollections nil", actual)
	})
}

func Test_Hashset_AddWithWgLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddWithWgLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- AddWithWgLock", actual)
	})
}

func Test_Hashset_AddSimpleSlice_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddSimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.SimpleSlice([]string{"a", "b"})
		hs.AddSimpleSlice(&ss)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddSimpleSlice", actual)
	})
}

func Test_Hashset_Has_Contains(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Has_Contains", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": hs.Has("a"),
			"contains": hs.Contains("a"),
			"missing": hs.IsMissing("b"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"contains": true,
			"missing": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Has/Contains/IsMissing", actual)
	})
}

func Test_Hashset_HasLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"hl": hs.HasLock("a"),
			"hwl": hs.HasWithLock("a"),
			"ml": hs.IsMissingLock("z"),
		}

		// Assert
		expected := args.Map{
			"hl": true,
			"hwl": true,
			"ml": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- lock variants", actual)
	})
}

func Test_Hashset_HasAllStrings_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAllStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": hs.HasAllStrings([]string{"a", "b"}),
			"miss": hs.HasAllStrings([]string{"a", "c"}),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAllStrings", actual)
	})
}

func Test_Hashset_HasAll_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAll", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": hs.HasAll("a", "b"),
			"miss": hs.HasAll("a", "c"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAll", actual)
	})
}

func Test_Hashset_HasAny_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"any": hs.HasAny("z", "a"),
			"none": hs.HasAny("x", "y"),
		}

		// Assert
		expected := args.Map{
			"any": true,
			"none": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAny", actual)
	})
}

func Test_Hashset_IsAllMissing_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsAllMissing", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"allMiss": hs.IsAllMissing("x", "y"),
			"notAll": hs.IsAllMissing("a", "x"),
		}

		// Assert
		expected := args.Map{
			"allMiss": true,
			"notAll": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsAllMissing", actual)
	})
}

func Test_Hashset_HasAllCollectionItems_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAllCollectionItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": hs.HasAllCollectionItems(coll),
			"nil": hs.HasAllCollectionItems(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAllCollectionItems", actual)
	})
}

func Test_Hashset_List_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_List", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(hs.List())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- List", actual)
	})
}

func Test_Hashset_ListPtr_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListPtr", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(hs.ListPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtr", actual)
	})
}

func Test_Hashset_ListCopyLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListCopyLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(hs.ListCopyLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListCopyLock", actual)
	})
}

func Test_Hashset_Items_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Items", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(hs.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Items", actual)
	})
}

func Test_Hashset_Collection_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Collection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := hs.Collection()

		// Act
		actual := args.Map{"len": coll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Collection", actual)
	})
}

func Test_Hashset_SortedList_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SortedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		sorted := hs.SortedList()

		// Act
		actual := args.Map{
			"first": sorted[0],
			"second": sorted[1],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"second": "b",
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SortedList", actual)
	})
}

func Test_Hashset_OrderedList_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_OrderedList", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		ol := hs.OrderedList()

		// Act
		actual := args.Map{"len": len(ol)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- OrderedList", actual)
	})
}

func Test_Hashset_OrderedList_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_OrderedList_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		ol := hs.OrderedList()

		// Act
		actual := args.Map{"len": len(ol)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- OrderedList empty", actual)
	})
}

func Test_Hashset_ListPtrSortedAsc_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListPtrSortedAsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		sorted := hs.ListPtrSortedAsc()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtrSortedAsc", actual)
	})
}

func Test_Hashset_ListPtrSortedDsc_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListPtrSortedDsc", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		sorted := hs.ListPtrSortedDsc()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtrSortedDsc", actual)
	})
}

func Test_Hashset_SafeStrings_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SafeStrings", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"len": len(hs.SafeStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- SafeStrings empty", actual)
	})
}

func Test_Hashset_Lines_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Lines", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"len": len(hs.Lines())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- Lines empty", actual)
	})
}

func Test_Hashset_SimpleSlice_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SimpleSlice", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		ss := hs.SimpleSlice()

		// Act
		actual := args.Map{"len": ss.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SimpleSlice", actual)
	})
}

func Test_Hashset_SimpleSlice_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SimpleSlice_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		ss := hs.SimpleSlice()

		// Act
		actual := args.Map{"empty": ss.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- SimpleSlice empty", actual)
	})
}

func Test_Hashset_MapStringAny_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		m := hs.MapStringAny()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MapStringAny", actual)
	})
}

func Test_Hashset_MapStringAny_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAny_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		m := hs.MapStringAny()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- MapStringAny empty", actual)
	})
}

func Test_Hashset_MapStringAnyDiff_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAnyDiff", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		d := hs.MapStringAnyDiff()

		// Act
		actual := args.Map{"notNil": d != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MapStringAnyDiff", actual)
	})
}

func Test_Hashset_Resize_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Resize", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Resize(100)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize", actual)
	})
}

func Test_Hashset_Resize_AlreadyLarger_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Resize_AlreadyLarger", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hs.Resize(1)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize already larger", actual)
	})
}

func Test_Hashset_ResizeLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ResizeLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.ResizeLock(100)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ResizeLock", actual)
	})
}

func Test_Hashset_ResizeLock_AlreadyLarger_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ResizeLock_AlreadyLarger", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hs.ResizeLock(1)

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ResizeLock already larger", actual)
	})
}

func Test_Hashset_AddCapacities_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacities", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacities(10, 20)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCapacities", actual)
	})
}

func Test_Hashset_AddCapacities_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacities_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacities()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddCapacities empty", actual)
	})
}

func Test_Hashset_AddCapacitiesLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacitiesLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacitiesLock(10)

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCapacitiesLock", actual)
	})
}

func Test_Hashset_AddCapacitiesLock_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacitiesLock_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacitiesLock()

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddCapacitiesLock empty", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_NoArgs_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewHashsets_NoArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewHashsets(true)

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewHashsets no args", actual)
	})
}

func Test_Hashset_ConcatNewHashsets_WithArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewHashsets_WithArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})
		c := hs.ConcatNewHashsets(true, other, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- ConcatNewHashsets with args", actual)
	})
}

func Test_Hashset_ConcatNewStrings_NoArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewStrings_NoArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewStrings(true)

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewStrings no args", actual)
	})
}

func Test_Hashset_ConcatNewStrings_WithArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewStrings_WithArgs", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewStrings(true, []string{"b", "c"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- ConcatNewStrings with args", actual)
	})
}

func Test_Hashset_Filter_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Filter", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"abc", "x"})
		filtered := hs.Filter(func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"len": filtered.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Filter", actual)
	})
}

func Test_Hashset_GetFilteredItems_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredItems", actual)
	})
}

func Test_Hashset_GetFilteredItems_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- GetFilteredItems empty", actual)
	})
}

func Test_Hashset_GetFilteredItems_Break_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, true })

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredItems break", actual)
	})
}

func Test_Hashset_GetFilteredCollection_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"len": coll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredCollection", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })

		// Act
		actual := args.Map{"empty": coll.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- GetFilteredCollection empty", actual)
	})
}

func Test_Hashset_GetFilteredCollection_Break_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, true })

		// Act
		actual := args.Map{"len": coll.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredCollection break", actual)
	})
}

func Test_Hashset_GetAllExceptHashset_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		exc := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptHashset(exc)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptHashset", actual)
	})
}

func Test_Hashset_GetAllExceptHashset_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptHashset_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptHashset(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptHashset nil", actual)
	})
}

func Test_Hashset_GetAllExcept_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExcept", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExcept", actual)
	})
}

func Test_Hashset_GetAllExcept_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExcept_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExcept nil", actual)
	})
}

func Test_Hashset_GetAllExceptSpread_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptSpread", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetAllExceptSpread("a")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptSpread", actual)
	})
}

func Test_Hashset_GetAllExceptSpread_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptSpread_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptSpread(nil...)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptSpread nil", actual)
	})
}

func Test_Hashset_GetAllExceptCollection_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptCollection", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		coll := corestr.New.Collection.Strings([]string{"a"})
		result := hs.GetAllExceptCollection(coll)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptCollection", actual)
	})
}

func Test_Hashset_GetAllExceptCollection_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptCollection_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptCollection nil", actual)
	})
}

func Test_Hashset_AddsUsingFilter_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsUsingFilter", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsUsingFilter nil", actual)
	})
}

func Test_Hashset_AddsUsingFilter_Break_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsUsingFilter break", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(nil, nil...)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilter nil", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_NilItem_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_NilItem", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, nil, "hello")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilter nil item", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilter_Break_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsAnyUsingFilter break", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilterLock_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilterLock(nil, nil...)

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilterLock nil", actual)
	})
}

func Test_Hashset_AddsAnyUsingFilterLock_Break_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilterLock_Break", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilterLock(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsAnyUsingFilterLock break", actual)
	})
}

func Test_Hashset_Remove_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Remove", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.Remove("a")

		// Act
		actual := args.Map{"has": hs.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Remove", actual)
	})
}

func Test_Hashset_SafeRemove_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SafeRemove", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.SafeRemove("a")
		hs.SafeRemove("missing")

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SafeRemove", actual)
	})
}

func Test_Hashset_RemoveWithLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_RemoveWithLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.RemoveWithLock("a")

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- RemoveWithLock", actual)
	})
}

func Test_Hashset_Clear_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Clear", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Clear()

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Clear", actual)
	})
}

func Test_Hashset_Clear_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Clear_Nil", func() {
		// Arrange
		var hs *corestr.Hashset
		result := hs.Clear()

		// Act
		actual := args.Map{"nil": result == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Clear nil", actual)
	})
}

func Test_Hashset_Dispose_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Dispose", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Dispose()

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Dispose", actual)
	})
}

func Test_Hashset_Dispose_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Dispose_Nil", func() {
		// Arrange
		var hs *corestr.Hashset
		hs.Dispose() // should not panic

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Dispose nil", actual)
	})
}

func Test_Hashset_IsEquals_Same_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_Same", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"same": hs.IsEquals(hs)}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals same ptr", actual)
	})
}

func Test_Hashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_BothNil", func() {
		// Arrange
		var hs *corestr.Hashset

		// Act
		actual := args.Map{"eq": hs.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- IsEquals both nil", actual)
	})
}

func Test_Hashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_OneNil", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": hs.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- IsEquals one nil", actual)
	})
}

func Test_Hashset_IsEquals_BothEmpty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_BothEmpty", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Cap(5)
		hs2 := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"eq": hs1.IsEquals(hs2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEquals both empty", actual)
	})
}

func Test_Hashset_IsEquals_DiffLen_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_DiffLen", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"eq": hs1.IsEquals(hs2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals diff len", actual)
	})
}

func Test_Hashset_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_DiffItems", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b"})

		// Act
		actual := args.Map{"eq": hs1.IsEquals(hs2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals diff items", actual)
	})
}

func Test_Hashset_IsEqual_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEqual", func() {
		// Arrange
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": hs1.IsEqual(hs2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEqual", actual)
	})
}

func Test_Hashset_IsEqualsLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEqualsLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": hs.IsEqualsLock(hs)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEqualsLock", actual)
	})
}

func Test_Hashset_ToLowerSet_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ToLowerSet", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"ABC"})
		lower := hs.ToLowerSet()

		// Act
		actual := args.Map{"has": lower.Has("abc")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ToLowerSet", actual)
	})
}

func Test_Hashset_String_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_String_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"notEmpty": hs.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- String empty", actual)
	})
}

func Test_Hashset_String_WithItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_String_WithItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": hs.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- String with items", actual)
	})
}

func Test_Hashset_StringLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_StringLock", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"notEmpty": hs.StringLock() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- StringLock empty", actual)
	})
}

func Test_Hashset_StringLock_WithItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_StringLock_WithItems", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": hs.StringLock() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- StringLock with items", actual)
	})
}

func Test_Hashset_Join_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Join", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"val": hs.Join(",")}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Join", actual)
	})
}

func Test_Hashset_JoinLine_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinLine", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"val": hs.JoinLine()}

		// Assert
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JoinLine", actual)
	})
}

func Test_Hashset_JoinSorted_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinSorted_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{"val": hs.JoinSorted(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- JoinSorted empty", actual)
	})
}

func Test_Hashset_JoinSorted_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinSorted", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"val": hs.JoinSorted(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JoinSorted", actual)
	})
}

func Test_Hashset_NonEmptyJoins_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_NonEmptyJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": hs.NonEmptyJoins(",") != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- NonEmptyJoins", actual)
	})
}

func Test_Hashset_NonWhitespaceJoins_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_NonWhitespaceJoins", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notEmpty": hs.NonWhitespaceJoins(",") != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- NonWhitespaceJoins", actual)
	})
}

func Test_Hashset_JsonModel_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jm := hs.JsonModel()

		// Act
		actual := args.Map{"len": len(jm)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonModel", actual)
	})
}

func Test_Hashset_JsonModel_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModel_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		jm := hs.JsonModel()

		// Act
		actual := args.Map{"len": len(jm)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- JsonModel empty", actual)
	})
}

func Test_Hashset_JsonModelAny_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModelAny", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})

		// Act
		actual := args.Map{"notNil": hs.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonModelAny", actual)
	})
}

func Test_Hashset_MarshalJSON_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		b, err := hs.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MarshalJSON", actual)
	})
}

func Test_Hashset_UnmarshalJSON_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_UnmarshalJSON", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		err := hs.UnmarshalJSON([]byte(`{"a":true}`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"has": hs.Has("a"),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- UnmarshalJSON", actual)
	})
}

func Test_Hashset_UnmarshalJSON_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_UnmarshalJSON_Err", func() {
		// Arrange
		hs := corestr.New.Hashset.Cap(5)
		err := hs.UnmarshalJSON([]byte(`{invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- UnmarshalJSON err", actual)
	})
}

func Test_Hashset_Json_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Json", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		j := hs.Json()
		actual := args.Map{"hasBytes": j.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Json", actual)
	})
}

func Test_Hashset_JsonPtr_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonPtr", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jp := hs.JsonPtr()
		actual := args.Map{"notNil": jp != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonPtr", actual)
	})
}

func Test_Hashset_Serialize_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Serialize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		b, err := hs.Serialize()
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Serialize", actual)
	})
}

func Test_Hashset_Deserialize_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Deserialize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		target := map[string]bool{}
		err := hs.Deserialize(&target)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Deserialize", actual)
	})
}

func Test_Hashset_ParseInjectUsingJson_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ParseInjectUsingJson", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Cap(5)
		result, err := hs2.ParseInjectUsingJson(jr)
		actual := args.Map{
			"noErr": err == nil,
			"has": result.Has("a"),
		}
		expected := args.Map{
			"noErr": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ParseInjectUsingJson", actual)
	})
}

func Test_Hashset_ParseInjectUsingJson_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ParseInjectUsingJson_Err", func() {
		hs := corestr.New.Hashset.Cap(5)
		badJson := corejson.NewPtr(42)
		_, err := hs.ParseInjectUsingJson(badJson)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- ParseInjectUsingJson err", actual)
	})
}

func Test_Hashset_JsonParseSelfInject_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonParseSelfInject", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Cap(5)
		err := hs2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonParseSelfInject", actual)
	})
}

func Test_Hashset_AsJsoner_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsoner", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsoner() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsoner", actual)
	})
}

func Test_Hashset_AsJsonContractsBinder_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonContractsBinder", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonContractsBinder() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonContractsBinder", actual)
	})
}

func Test_Hashset_AsJsonParseSelfInjector_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonParseSelfInjector", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonParseSelfInjector", actual)
	})
}

func Test_Hashset_AsJsonMarshaller_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonMarshaller", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonMarshaller() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonMarshaller", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_BothEmpty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_BothEmpty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- DistinctDiffLinesRaw both empty", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_LeftOnly_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_LeftOnly", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw left only", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_RightOnly_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_RightOnly", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLinesRaw("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw right only", actual)
	})
}

func Test_Hashset_DistinctDiffLinesRaw_Both_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_Both", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.DistinctDiffLinesRaw("b", "c")
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw both", actual)
	})
}

func Test_Hashset_DistinctDiffLines_BothEmpty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_BothEmpty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLines()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- DistinctDiffLines both empty", actual)
	})
}

func Test_Hashset_DistinctDiffLines_LeftOnly_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_LeftOnly", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.DistinctDiffLines()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines left only", actual)
	})
}

func Test_Hashset_DistinctDiffLines_RightOnly_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_RightOnly", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLines("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines right only", actual)
	})
}

func Test_Hashset_DistinctDiffLines_Both(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_Both", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.DistinctDiffLines("b", "c")
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines both", actual)
	})
}

func Test_Hashset_DistinctDiffHashset_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffHashset", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		other := corestr.New.Hashset.Strings([]string{"b", "c"})
		result := hs.DistinctDiffHashset(other)
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffHashset", actual)
	})
}

func Test_Hashset_WrapDoubleQuote_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapDoubleQuote", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapDoubleQuote()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapDoubleQuote", actual)
	})
}

func Test_Hashset_WrapDoubleQuoteIfMissing_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapDoubleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapDoubleQuoteIfMissing()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapDoubleQuoteIfMissing", actual)
	})
}

func Test_Hashset_WrapSingleQuote_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapSingleQuote", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapSingleQuote()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapSingleQuote", actual)
	})
}

func Test_Hashset_WrapSingleQuoteIfMissing_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapSingleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapSingleQuoteIfMissing()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapSingleQuoteIfMissing", actual)
	})
}

func Test_Hashset_Transpile_Empty_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Transpile_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.Transpile(func(s string) string { return s })
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- Transpile empty", actual)
	})
}

func Test_Hashset_AddStringsPtrWgLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsPtrWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStringsPtrWgLock", actual)
	})
}

func Test_Hashset_AddHashsetWgLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		other := corestr.New.Hashset.Strings([]string{"a"})
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashsetWgLock", actual)
	})
}

func Test_Hashset_AddHashsetWgLock_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetWgLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		hs.AddHashsetWgLock(nil, wg)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetWgLock nil", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMapWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		m := map[string]bool{"a": true, "b": false}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()
		actual := args.Map{
			"hasA": hs.Has("a"),
			"hasB": hs.Has("b"),
		}
		expected := args.Map{
			"hasA": true,
			"hasB": false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMapWgLock", actual)
	})
}

func Test_Hashset_AddItemsMapWgLock_Nil_FromHashsetIsEmptyIterat(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMapWgLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		hs.AddItemsMapWgLock(nil, wg)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddItemsMapWgLock nil", actual)
	})
}
