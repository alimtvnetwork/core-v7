package coredynamictests

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Collection — constructors / accessors / size
// =============================================================================

func Test_Col_NewCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)

	// Act
	actual := args.Map{
		"len": c.Length(),
		"empty": c.IsEmpty(),
		"hasAny": c.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"hasAny": false,
	}
	expected.ShouldBeEqual(t, 0, "NewCollection", actual)
}

func Test_Col_EmptyCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()

	// Act
	actual := args.Map{
		"len": c.Length(),
		"count": c.Count(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"count": 0,
	}
	expected.ShouldBeEqual(t, 0, "EmptyCollection", actual)
}

func Test_Col_CollectionFrom_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string](nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionFrom nil", actual)
}

func Test_Col_CollectionFrom_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionFrom valid", actual)
}

func Test_Col_CollectionClone(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionClone([]int{1, 2, 3})

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "CollectionClone", actual)
}

func Test_Col_Length_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]

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
	expected.ShouldBeEqual(t, 0, "Length nil", actual)
}

func Test_Col_At(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"r": c.At(1)}

	// Assert
	expected := args.Map{"r": "b"}
	expected.ShouldBeEqual(t, 0, "At", actual)
}

func Test_Col_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	_, ok := c.FirstOrDefault()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault empty", actual)
}

func Test_Col_FirstOrDefault_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"x"})
	v, ok := c.FirstOrDefault()

	// Act
	actual := args.Map{
		"ok": ok,
		"val": *v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"val": "x",
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault valid", actual)
}

func Test_Col_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	_, ok := c.LastOrDefault()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "LastOrDefault empty", actual)
}

func Test_Col_LastOrDefault_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	v, ok := c.LastOrDefault()

	// Act
	actual := args.Map{
		"ok": ok,
		"val": *v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"val": "b",
	}
	expected.ShouldBeEqual(t, 0, "LastOrDefault valid", actual)
}

func Test_Col_Items_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]

	// Act
	actual := args.Map{"len": len(c.Items())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Items nil", actual)
}

func Test_Col_LastIndex(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": c.LastIndex()}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "LastIndex", actual)
}

func Test_Col_HasIndex(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{
		"valid": c.HasIndex(1),
		"invalid": c.HasIndex(5),
		"neg": c.HasIndex(-1),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
		"neg": false,
	}
	expected.ShouldBeEqual(t, 0, "HasIndex", actual)
}

// =============================================================================
// Slicing
// =============================================================================

func Test_Col_Skip(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})

	// Act
	actual := args.Map{"len": len(c.Skip(2))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Skip", actual)
}

func Test_Col_Take(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})

	// Act
	actual := args.Map{"len": len(c.Take(2))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Take", actual)
}

func Test_Col_Limit(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": len(c.Limit(1))}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Limit", actual)
}

func Test_Col_SkipCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": c.SkipCollection(1).Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SkipCollection", actual)
}

func Test_Col_TakeCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": c.TakeCollection(2).Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TakeCollection", actual)
}

func Test_Col_LimitCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": c.LimitCollection(1).Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LimitCollection", actual)
}

func Test_Col_SafeLimitCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"len": c.SafeLimitCollection(100).Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeLimitCollection capped", actual)
}

// =============================================================================
// Mutators
// =============================================================================

func Test_Col_Add(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.Add("a").Add("b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Add", actual)
}

func Test_Col_AddMany(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddMany(1, 2, 3)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AddMany", actual)
}

func Test_Col_AddNonNil_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddNonNil(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddNonNil nil", actual)
}

func Test_Col_AddNonNil_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	s := "hello"
	c.AddNonNil(&s)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddNonNil valid", actual)
}

func Test_Col_RemoveAt_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	ok := c.RemoveAt(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": c.Length(),
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAt valid", actual)
}

func Test_Col_RemoveAt_Invalid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"ok": c.RemoveAt(5)}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "RemoveAt invalid", actual)
}

func Test_Col_Clear(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	c.Clear()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear", actual)
}

func Test_Col_Dispose(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	c.Dispose()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)
}

// =============================================================================
// Loop / LoopAsync / Filter
// =============================================================================

func Test_Col_Loop_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	count := 0
	c.Loop(func(i int, item int) bool { count++; return false })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "Loop empty", actual)
}

func Test_Col_Loop_Break(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	count := 0
	c.Loop(func(i int, item int) bool { count++; return item == 2 })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Loop break", actual)
}

func Test_Col_LoopAsync(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	var mu sync.Mutex
	sum := 0
	c.LoopAsync(func(i int, item int) {
		mu.Lock()
		sum += item
		mu.Unlock()
	})

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "LoopAsync", actual)
}

func Test_Col_LoopAsync_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, item int) {})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LoopAsync empty", actual)
}

func Test_Col_Filter_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	f := c.Filter(func(i int) bool { return true })

	// Act
	actual := args.Map{"len": f.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Filter empty", actual)
}

func Test_Col_Filter_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	f := c.Filter(func(i int) bool { return i%2 == 0 })

	// Act
	actual := args.Map{"len": f.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Filter valid", actual)
}

// =============================================================================
// Paging
// =============================================================================

func Test_Col_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": c.GetPagesSize(0)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "GetPagesSize zero", actual)
}

func Test_Col_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})

	// Act
	actual := args.Map{"r": c.GetPagesSize(2)}

	// Assert
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "GetPagesSize valid", actual)
}

func Test_Col_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	p := c.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": p.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GetSinglePage small", actual)
}

func Test_Col_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	pages := c.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection small", actual)
}

func Test_Col_GetPagedCollection_Multi(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	pages := c.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection multi", actual)
}

// =============================================================================
// JSON
// =============================================================================

func Test_Col_MarshalJSON(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	b, err := json.Marshal(c)

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
	expected.ShouldBeEqual(t, 0, "MarshalJSON", actual)
}

func Test_Col_UnmarshalJSON(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	err := json.Unmarshal([]byte(`[1,2,3]`), c)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
}

func Test_Col_JsonString(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	s, err := c.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "[1]",
	}
	expected.ShouldBeEqual(t, 0, "JsonString", actual)
}

func Test_Col_JsonStringMust(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	s := c.JsonStringMust()

	// Act
	actual := args.Map{"r": s}

	// Assert
	expected := args.Map{"r": `["a"]`}
	expected.ShouldBeEqual(t, 0, "JsonStringMust", actual)
}

// =============================================================================
// Strings / String
// =============================================================================

func Test_Col_Strings(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	s := c.Strings()

	// Act
	actual := args.Map{
		"len": len(s),
		"first": s[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "1",
	}
	expected.ShouldBeEqual(t, 0, "Strings", actual)
}

func Test_Col_String(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"r": c.String()}

	// Assert
	expected := args.Map{"r": "1"}
	expected.ShouldBeEqual(t, 0, "String", actual)
}
