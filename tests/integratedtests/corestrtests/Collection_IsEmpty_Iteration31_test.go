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

func Test_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{
			"empty": c.IsEmpty(),
			"items": c.HasItems(),
			"len": c.Length(),
			"count": c.Count(),
			"hasAny": c.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"items": false,
			"len": 0,
			"count": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- empty", actual)
	})
}

func Test_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Length_Nil", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{
			"len": c.Length(),
			"empty": c.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- nil length", actual)
	})
}

func Test_Collection_IsEmptyLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEmptyLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"empty": c.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEmptyLock", actual)
	})
}

func Test_Collection_LengthLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LengthLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{"len": c.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LengthLock", actual)
	})
}

func Test_Collection_Capacity_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Capacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 10}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Capacity", actual)
	})
}

func Test_Collection_Capacity_Nil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Capacity_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(0)

		// Act
		actual := args.Map{"cap": c.Capacity()}

		// Assert
		expected := args.Map{"cap": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- Capacity nil items", actual)
	})
}

func Test_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LastIndex", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"li": c.LastIndex(),
			"hasIdx": c.HasIndex(1),
			"noIdx": c.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"li": 1,
			"hasIdx": true,
			"noIdx": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LastIndex/HasIndex", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Add variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Add_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Add", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Add", actual)
	})
}

func Test_Collection_AddNonEmpty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmpty("")
		c.AddNonEmpty("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmpty", actual)
	})
}

func Test_Collection_AddNonEmptyWhitespace_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddNonEmptyWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddNonEmptyWhitespace", actual)
	})
}

func Test_Collection_AddIf_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddIf", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "a")
		c.AddIf(false, "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIf", actual)
	})
}

func Test_Collection_AddIfMany_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddIfMany", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(true, "a", "b")
		c.AddIfMany(false, "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddIfMany", actual)
	})
}

func Test_Collection_AddError_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddError", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddError(nil)
		c.AddError(fmt.Errorf("oops"))

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddError", actual)
	})
}

func Test_Collection_AddFunc_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFunc", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "computed"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddFunc", actual)
	})
}

func Test_Collection_AddFuncErr_NoErr_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFuncErr_NoErr", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "ok"}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddFuncErr no err", actual)
	})
}

func Test_Collection_AddFuncErr_Err(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddFuncErr_Err", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		called := false
		c.AddFuncErr(func() (string, error) { return "", fmt.Errorf("fail") }, func(e error) { called = true })

		// Act
		actual := args.Map{
			"empty": c.IsEmpty(),
			"called": called,
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"called": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AddFuncErr err", actual)
	})
}

func Test_Collection_AddLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddLock", actual)
	})
}

func Test_Collection_Adds_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Adds", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Adds", actual)
	})
}

func Test_Collection_AddsLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsLock("a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsLock", actual)
	})
}

func Test_Collection_AddStrings_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddStrings", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddStrings", actual)
	})
}

func Test_Collection_AddCollection_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollection", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddCollection(other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
	})
}

func Test_Collection_AddCollection_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollection_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("x")
		other := corestr.New.Collection.Cap(5)
		c.AddCollection(other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddCollection empty", actual)
	})
}

func Test_Collection_AddCollections_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCollections", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		empty := corestr.New.Collection.Cap(0)
		c.AddCollections(c1, empty, c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
	})
}

func Test_Collection_AddWithWgLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddWithWgLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddWithWgLock", actual)
	})
}

func Test_Collection_AddHashmapsValues_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsValues", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsValues", actual)
	})
}

func Test_Collection_AddHashmapsValues_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsValues_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsValues nil", actual)
	})
}

func Test_Collection_AddHashmapsKeys_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)

		// Act
		actual := args.Map{"has": c.Has("k")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddHashmapsKeys", actual)
	})
}

func Test_Collection_AddHashmapsKeys_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeys_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsKeys nil", actual)
	})
}

func Test_Collection_AddHashmapsKeysValues_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeysValues", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(hm)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddHashmapsKeysValues", actual)
	})
}

func Test_Collection_AddHashmapsKeysValues_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddHashmapsKeysValues_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeysValues(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddHashmapsKeysValues nil", actual)
	})
}

func Test_Collection_AddsNonEmpty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddsNonEmpty", actual)
	})
}

func Test_Collection_AddsNonEmpty_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmpty_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsNonEmpty nil", actual)
	})
}

func Test_Collection_AddsNonEmptyPtrLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmptyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		s := "hello"
		empty := ""
		c.AddsNonEmptyPtrLock(&s, nil, &empty)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AddsNonEmptyPtrLock", actual)
	})
}

func Test_Collection_AddsNonEmptyPtrLock_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddsNonEmptyPtrLock_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmptyPtrLock(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddsNonEmptyPtrLock nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Indexing (First / Last / IndexAt / SafeIndexAt)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_IndexAt_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IndexAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"at1": c.IndexAt(1)}

		// Assert
		expected := args.Map{"at1": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IndexAt", actual)
	})
}

func Test_Collection_SafeIndexAtUsingLength_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_SafeIndexAtUsingLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"safe": c.SafeIndexAtUsingLength("def", 2, 1),
			"oob": c.SafeIndexAtUsingLength("def", 2, 5),
		}

		// Assert
		expected := args.Map{
			"safe": "b",
			"oob": "def",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SafeIndexAtUsingLength", actual)
	})
}

func Test_Collection_First_Last_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_First_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- First/Last", actual)
	})
}

func Test_Collection_FirstOrDefault_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_FirstOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.FirstOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- FirstOrDefault empty", actual)
	})
}

func Test_Collection_LastOrDefault_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_LastOrDefault_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"val": c.LastOrDefault()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- LastOrDefault empty", actual)
	})
}

func Test_Collection_Single_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"only"})

		// Act
		actual := args.Map{"val": c.Single()}

		// Assert
		expected := args.Map{"val": "only"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Single", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Take / Skip / Reverse
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Take_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		t1 := c.Take(2)

		// Act
		actual := args.Map{
			"len": t1.Length(),
			"first": t1.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take", actual)
	})
}

func Test_Collection_Take_All(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take_All", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		t1 := c.Take(5)

		// Act
		actual := args.Map{"same": t1.Length()}

		// Assert
		expected := args.Map{"same": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take all", actual)
	})
}

func Test_Collection_Take_Zero_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Take_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		t1 := c.Take(0)

		// Act
		actual := args.Map{"empty": t1.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Take zero", actual)
	})
}

func Test_Collection_Skip_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		s := c.Skip(1)

		// Act
		actual := args.Map{
			"len": s.Length(),
			"first": s.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "b",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip", actual)
	})
}

func Test_Collection_Skip_Zero_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Skip_Zero", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.Skip(0)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip zero", actual)
	})
}

func Test_Collection_Reverse_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()

		// Act
		actual := args.Map{
			"first": c.First(),
			"last": c.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse", actual)
	})
}

func Test_Collection_Reverse_Two_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse_Two", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse two", actual)
	})
}

func Test_Collection_Reverse_Single_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Reverse_Single", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Reverse single", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — RemoveAt / ChainRemoveAt / InsertAt
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_RemoveAt_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)

		// Act
		actual := args.Map{
			"ok": ok,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
	})
}

func Test_Collection_RemoveAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt_OutOfRange", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		ok := c.RemoveAt(5)

		// Act
		actual := args.Map{"ok": ok}

		// Assert
		expected := args.Map{"ok": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt out of range", actual)
	})
}

func Test_Collection_RemoveAt_Negative(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveAt_Negative", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		ok := c.RemoveAt(-1)

		// Act
		actual := args.Map{"ok": ok}

		// Assert
		expected := args.Map{"ok": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt negative", actual)
	})
}

func Test_Collection_ChainRemoveAt_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(0)

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ChainRemoveAt", actual)
	})
}

func Test_Collection_InsertAt_First_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_InsertAt_First", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.InsertAt(0, "a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- InsertAt first", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_GetPagesSize_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagesSize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})

		// Act
		actual := args.Map{
			"pages": c.GetPagesSize(2),
			"zero": c.GetPagesSize(0),
			"neg": c.GetPagesSize(-1),
		}

		// Assert
		expected := args.Map{
			"pages": 3,
			"zero": 0,
			"neg": 0,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagesSize", actual)
	})
}

func Test_Collection_GetSinglePageCollection_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		p := c.GetSinglePageCollection(2, 2)

		// Act
		actual := args.Map{
			"len": p.Length(),
			"first": p.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "c",
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection", actual)
	})
}

func Test_Collection_GetSinglePageCollection_LastPage_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection_LastPage", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		p := c.GetSinglePageCollection(2, 3)

		// Act
		actual := args.Map{"len": p.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection last page", actual)
	})
}

func Test_Collection_GetSinglePageCollection_SmallList_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetSinglePageCollection_SmallList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		p := c.GetSinglePageCollection(10, 1)

		// Act
		actual := args.Map{"len": p.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetSinglePageCollection small", actual)
	})
}

func Test_Collection_GetPagedCollection_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagedCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		paged := c.GetPagedCollection(2)

		// Act
		actual := args.Map{"notNil": paged != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection", actual)
	})
}

func Test_Collection_GetPagedCollection_SmallList_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_GetPagedCollection_SmallList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		paged := c.GetPagedCollection(10)

		// Act
		actual := args.Map{"notNil": paged != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- GetPagedCollection small", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — EachItemSplitBy / ConcatNew
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_EachItemSplitBy_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_EachItemSplitBy", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- EachItemSplitBy", actual)
	})
}

func Test_Collection_ConcatNew_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ConcatNew", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0, "b", "c")

		// Act
		actual := args.Map{"len": cn.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ConcatNew", actual)
	})
}

func Test_Collection_ConcatNew_NoArgs(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ConcatNew_NoArgs", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0)

		// Act
		actual := args.Map{"len": cn.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- ConcatNew no args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — IsEquals
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_Same", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"same": c.IsEquals(c)}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals same ptr", actual)
	})
}

func Test_Collection_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_BothNil", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{"eq": c.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- IsEquals both nil", actual)
	})
}

func Test_Collection_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_OneNil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"eq": c.IsEquals(nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- IsEquals one nil", actual)
	})
}

func Test_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_BothEmpty", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"eq": c1.IsEquals(c2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEquals both empty", actual)
	})
}

func Test_Collection_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_DiffLen", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"eq": c1.IsEquals(c2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals diff len", actual)
	})
}

func Test_Collection_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEquals_DiffItems", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})

		// Act
		actual := args.Map{"eq": c1.IsEquals(c2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- IsEquals diff items", actual)
	})
}

func Test_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_I31_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"Hello"})
		c2 := corestr.New.Collection.Strings([]string{"hello"})

		// Act
		actual := args.Map{
			"eq": c1.IsEqualsWithSensitive(false, c2),
			"neq": c1.IsEqualsWithSensitive(true, c2),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- IsEqualsWithSensitive", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — AsError / ToError
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsError_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{"nil": c.AsError(",") == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AsError empty", actual)
	})
}

func Test_Collection_AsError_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsError(", ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsError", actual)
	})
}

func Test_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AsDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"err1"})
		err := c.AsDefaultError()

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- AsDefaultError", actual)
	})
}

func Test_Collection_ToError_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ToError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		err := c.ToError(", ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToError", actual)
	})
}

func Test_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ToDefaultError", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		err := c.ToDefaultError()

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection returns error -- ToDefaultError", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Append variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_AppendCollectionPtr_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendCollectionPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendCollectionPtr", actual)
	})
}

func Test_Collection_AppendCollections_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendCollections_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendCollections empty", actual)
	})
}

func Test_Collection_AppendAnys_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys(42, nil, "hello")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnys", actual)
	})
}

func Test_Collection_AppendAnys_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnys_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendAnys empty", actual)
	})
}

func Test_Collection_AppendAnysLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnysLock", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock(42)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AppendAnysLock", actual)
	})
}

func Test_Collection_AppendAnysLock_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendAnysLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendAnysLock empty", actual)
	})
}

func Test_Collection_AppendNonEmptyAnys_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys("hello", nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- AppendNonEmptyAnys", actual)
	})
}

func Test_Collection_AppendNonEmptyAnys_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AppendNonEmptyAnys_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys(nil...)

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AppendNonEmptyAnys nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — RemoveItemsIndexes
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_RemoveItemsIndexes_NilIgnore_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexes_NilIgnore", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.RemoveItemsIndexes(true, nil...)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexes nil ignore", actual)
	})
}

func Test_Collection_RemoveItemsIndexesPtr(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexesPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexesPtr(true, []int{1})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveItemsIndexesPtr", actual)
	})
}

func Test_Collection_RemoveItemsIndexesPtr_NilIndexes(t *testing.T) {
	safeTest(t, "Test_I31_Collection_RemoveItemsIndexesPtr_NilIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexesPtr(true, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- RemoveItemsIndexesPtr nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Resize / AddCapacity
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_Resize_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Resize", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		c.Resize(100)

		// Act
		actual := args.Map{
			"cap": c.Capacity() >= 100,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"cap": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Resize", actual)
	})
}

func Test_Collection_Resize_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_I31_Collection_Resize_AlreadyLarger", func() {
		// Arrange
		c := corestr.New.Collection.Cap(100)
		c.Resize(5)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Resize already larger", actual)
	})
}

func Test_Collection_AddCapacity_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCapacity", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(50, 50)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCapacity", actual)
	})
}

func Test_Collection_AddCapacity_Nil_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_AddCapacity_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddCapacity(nil...)

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Collection returns nil -- AddCapacity nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — List / ListStrings / Items
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_List_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_List", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"lLen": len(c.List()),
			"lsLen": len(c.ListStrings()),
			"lspLen": len(c.ListStringsPtr()),
			"iLen": len(c.Items()),
			"lpLen": len(c.ListPtr()),
		}

		// Assert
		expected := args.Map{
			"lLen": 1,
			"lsLen": 1,
			"lspLen": 1,
			"iLen": 1,
			"lpLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- List variants", actual)
	})
}

func Test_Collection_ListCopyPtrLock_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ListCopyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		cp := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(cp)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ListCopyPtrLock", actual)
	})
}

func Test_Collection_ListCopyPtrLock_Empty_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_ListCopyPtrLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		cp := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(cp)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection returns empty -- ListCopyPtrLock empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// Collection — JsonString / StringJSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_JsonString_FromCollectionIsEmptyIte(t *testing.T) {
	safeTest(t, "Test_I31_Collection_JsonString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"notEmpty": c.JsonString() != "",
			"must": c.JsonStringMust() != "",
			"sj": c.StringJSON() != "",
		}

		// Assert
		expected := args.Map{
			"notEmpty": true,
			"must": true,
			"sj": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonString variants", actual)
	})
}
