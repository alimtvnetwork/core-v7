package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── Collection AppendPrependIf ──

func Test_AppendPrependIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.AppendPrependIf(true,
		[]namevalue.Instance[string, string]{{Name: "a", Value: "1"}},
		[]namevalue.Instance[string, string]{{Name: "c", Value: "3"}},
	)

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.Items[0].Name,
		"last": c.Items[2].Name,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf true -- both applied", actual)
}

func Test_AppendPrependIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.AppendPrependIf(false,
		[]namevalue.Instance[string, string]{{Name: "a", Value: "1"}},
		[]namevalue.Instance[string, string]{{Name: "c", Value: "3"}},
	)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf false -- skipped", actual)
}

// ── Collection AddsIf ──

func Test_AddsIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsIf true -- added", actual)
}

func Test_AddsIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf false -- skipped", actual)
}

// ── Collection PrependUsingFuncIf ──

func Test_PrependUsingFuncIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.Items[0].Name,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf returns non-empty -- true", actual)
}

func Test_PrependUsingFuncIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(false, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf returns non-empty -- false", actual)
}

func Test_PrependUsingFuncIf_NilFunc(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependUsingFuncIf(true, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf returns nil -- nil func", actual)
}

// ── Collection AppendUsingFuncIf ──

func Test_AppendUsingFuncIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf returns non-empty -- true", actual)
}

func Test_AppendUsingFuncIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(false, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf returns non-empty -- false", actual)
}

func Test_AppendUsingFuncIf_NilFunc(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf returns nil -- nil func", actual)
}

// ── Collection AddsPtr ──

func Test_AddsPtr(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(&item, nil) // nil should be skipped

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsPtr returns nil -- with nil skip", actual)
}

func Test_AddsPtr_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsPtr()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsPtr returns empty -- empty", actual)
}

// ── Collection ConcatNew / ConcatNewPtr ──

func Test_ConcatNew(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.ConcatNew(namevalue.Instance[string, string]{Name: "b", Value: "2"})

	// Act
	actual := args.Map{
		"origLen": c.Length(),
		"newLen": result.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 1,
		"newLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew -- immutable", actual)
}

func Test_ConcatNewPtr(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := namevalue.Instance[string, string]{Name: "b", Value: "2"}
	result := c.ConcatNewPtr(&item, nil)

	// Act
	actual := args.Map{"newLen": result.Length()}

	// Assert
	expected := args.Map{"newLen": 2}
	expected.ShouldBeEqual(t, 0, "ConcatNewPtr returns correct value -- with args", actual)
}

// ── Collection IsEqualByString ──

func Test_IsEqualByString_Equal(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"result": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns correct value -- equal", actual)
}

func Test_IsEqualByString_DiffLen(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()

	// Act
	actual := args.Map{"result": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns correct value -- diff len", actual)
}

func Test_IsEqualByString_DiffContent(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})

	// Act
	actual := args.Map{"result": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns correct value -- diff content", actual)
}

func Test_IsEqualByString_BothNil(t *testing.T) {
	// Arrange
	var c1, c2 *namevalue.Collection[string, string]

	// Act
	actual := args.Map{"result": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns nil -- both nil", actual)
}

func Test_IsEqualByString_OneNil(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()

	// Act
	actual := args.Map{"result": c.IsEqualByString(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns nil -- one nil", actual)
}

// ── NewGenericCollectionUsing no-clone ──

func Test_NewGenericCollectionUsing_NoClone(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	c := namevalue.NewGenericCollectionUsing[string, string](false, items...)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing returns empty -- no clone", actual)
}

func Test_NewGenericCollectionUsing_NilItems(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionUsing[string, string](true)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing returns nil -- nil items", actual)
}

// ── Collection HasIndex / LastIndex / HasAnyItem ──

func Test_Collection_IndexMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})

	// Act
	actual := args.Map{
		"lastIndex": c.LastIndex(),
		"hasIndex0": c.HasIndex(0),
		"hasIndex1": c.HasIndex(1),
		"hasIndex2": c.HasIndex(2),
		"hasAny":    c.HasAnyItem(),
		"count":     c.Count(),
	}

	// Assert
	expected := args.Map{
		"lastIndex": 1, "hasIndex0": true, "hasIndex1": true,
		"hasIndex2": false, "hasAny": true, "count": 2,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- index methods", actual)
}

// ── Collection CompiledLazyString nil ──

func Test_CompiledLazyString_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	result := c.CompiledLazyString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString nil -- empty", actual)
}

// ── Instance IsNull / Dispose ──

func Test_Instance_IsNull(t *testing.T) {
	// Arrange
	inst := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	var nilInst *namevalue.Instance[string, string]

	// Act
	actual := args.Map{
		"nonNil": inst.IsNull(),
		"nil": nilInst.IsNull(),
	}

	// Assert
	expected := args.Map{
		"nonNil": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IsNull", actual)
}

func Test_Instance_Dispose(t *testing.T) {
	// Arrange
	inst := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	inst.Dispose()

	// Act
	actual := args.Map{
		"name": inst.Name,
		"val": inst.Value,
	}

	// Assert
	expected := args.Map{
		"name": "",
		"val": "",
	}
	expected.ShouldBeEqual(t, 0, "Instance Dispose -- zeroed", actual)
}

func Test_Instance_Dispose_Nil(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[string, string]
	inst.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Instance Dispose nil -- no panic", actual)
}

// ── Append / Prepend / AppendIf / PrependIf empty ──

func Test_Append_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Append empty -- no-op", actual)
}

func Test_Prepend_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Prepend()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Prepend empty -- no-op", actual)
}

func Test_AppendIf_FalseSkip(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendIf false -- skipped", actual)
}

func Test_PrependIf_FalseSkip(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependIf false -- skipped", actual)
}

// ── Adds empty ──

func Test_Adds_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Adds()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Adds empty -- no-op", actual)
}

// ── Collection Join ──

func Test_Join(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	result := c.Join("; ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Join", actual)
}
