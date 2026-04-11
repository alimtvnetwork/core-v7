package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ============================================================================
// AppendsIf
// ============================================================================

func Test_AppendsIf_True_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(true, items, namevalue.StringAny{Name: "b", Value: 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[1].Name,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendsIf appends -- isAdd true", actual)
}

func Test_AppendsIf_False_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(false, items, namevalue.StringAny{Name: "b", Value: 2})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf no-op -- isAdd false", actual)
}

func Test_AppendsIf_EmptyAppend(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(true, items)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf no-op -- empty appending items", actual)
}

// ============================================================================
// PrependsIf
// ============================================================================

func Test_PrependsIf_True_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{{Name: "b", Value: 2}}
	result := namevalue.PrependsIf(true, items, namevalue.StringAny{Name: "a", Value: 1})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0].Name,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "PrependsIf prepends -- isAdd true", actual)
}

func Test_PrependsIf_False_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{{Name: "b", Value: 2}}
	result := namevalue.PrependsIf(false, items, namevalue.StringAny{Name: "a", Value: 1})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependsIf no-op -- isAdd false", actual)
}

// ============================================================================
// NewNameValuesCollection / EmptyNameValuesCollection
// ============================================================================

func Test_NewNameValuesCollection_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewNameValuesCollection(5)

	// Act
	actual := args.Map{
		"len": c.Length(),
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewNameValuesCollection creates empty -- cap 5", actual)
}

func Test_EmptyNameValuesCollection_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.EmptyNameValuesCollection()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyNameValuesCollection creates empty -- cap 0", actual)
}

func Test_NewNewNameValuesCollectionUsing_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewNewNameValuesCollectionUsing(true,
		namevalue.StringAny{Name: "a", Value: 1},
		namevalue.StringAny{Name: "b", Value: 2},
	)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewNewNameValuesCollectionUsing creates with items -- clone", actual)
}

// ============================================================================
// Collection — Count, HasIndex
// ============================================================================

func Test_Collection_Count(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})

	// Act
	actual := args.Map{"count": c.Count()}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Count returns item count -- two items", actual)
}

func Test_Collection_HasIndex_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{"result": c.HasIndex(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasIndex returns true -- index 0 exists", actual)
}

func Test_Collection_HasIndex_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)

	// Act
	actual := args.Map{"result": c.HasIndex(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- empty collection", actual)
}

// ============================================================================
// Collection — AppendIf / PrependIf
// ============================================================================

func Test_Collection_AppendIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendIf(true, namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendIf appends -- isAppend true", actual)
}

func Test_Collection_AppendIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendIf(false, namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendIf no-op -- isAppend false", actual)
}

func Test_Collection_PrependIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependIf(true, namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{
		"first": c.Items[0].Name,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "PrependIf prepends -- isPrepend true", actual)
}

func Test_Collection_PrependIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependIf(false, namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependIf no-op -- isPrepend false", actual)
}

// ============================================================================
// Collection — AddsIf
// ============================================================================

func Test_Collection_AddsIf_True(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AddsIf(true, namevalue.StringInt{Name: "a", Value: 1}, namevalue.StringInt{Name: "b", Value: 2})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddsIf adds items -- isAdd true", actual)
}

func Test_Collection_AddsIf_False(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AddsIf(false, namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf no-op -- isAdd false", actual)
}

// ============================================================================
// Collection — AddsPtr
// ============================================================================

func Test_Collection_AddsPtr_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	item := namevalue.StringInt{Name: "a", Value: 1}
	c.AddsPtr(&item, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsPtr adds non-nil skips nil -- mixed", actual)
}

// ============================================================================
// Collection — PrependUsingFuncIf / AppendUsingFuncIf
// ============================================================================

func Test_Collection_PrependUsingFuncIf_True_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})

	// Act
	actual := args.Map{
		"first": c.Items[0].Name,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf prepends -- true", actual)
}

func Test_Collection_PrependUsingFuncIf_False_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.PrependUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf no-op -- false", actual)
}

func Test_Collection_PrependUsingFuncIf_NilFunc_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.PrependUsingFuncIf(true, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf no-op -- nil func", actual)
}

func Test_Collection_AppendUsingFuncIf_True_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf appends -- true", actual)
}

func Test_Collection_AppendUsingFuncIf_False_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf no-op -- false", actual)
}

// ============================================================================
// Collection — AppendPrependIf
// ============================================================================

func Test_Collection_AppendPrependIf_True_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	c.AppendPrependIf(true,
		[]namevalue.StringInt{{Name: "first", Value: 1}},
		[]namevalue.StringInt{{Name: "last", Value: 2}},
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
		"first": "first",
		"last": "last",
	}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf prepends and appends -- true", actual)
}

func Test_Collection_AppendPrependIf_False_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	c.AppendPrependIf(false,
		[]namevalue.StringInt{{Name: "first", Value: 1}},
		[]namevalue.StringInt{{Name: "last", Value: 2}},
	)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf no-op -- false", actual)
}

// ============================================================================
// Collection — CompiledLazyString
// ============================================================================

func Test_Collection_CompiledLazyString_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	first := c.CompiledLazyString()
	second := c.CompiledLazyString()

	// Act
	actual := args.Map{
		"same": first == second,
		"hasContent": len(first) > 0,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"hasContent": true,
	}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString caches result -- second call same", actual)
}

func Test_Collection_CompiledLazyString_Nil_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, int]
	result := c.CompiledLazyString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString returns empty -- nil receiver", actual)
}

// ============================================================================
// Collection — ConcatNewPtr
// ============================================================================

func Test_Collection_ConcatNewPtr_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	item := namevalue.StringInt{Name: "b", Value: 2}
	result := c.ConcatNewPtr(&item)

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
	expected.ShouldBeEqual(t, 0, "ConcatNewPtr creates new collection with added ptr -- clone + add", actual)
}

// ============================================================================
// Collection — JsonStrings / JoinJsonStrings / JoinCsv / JoinCsvLine
// ============================================================================

func Test_Collection_JsonStrings_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JsonStrings()

	// Act
	actual := args.Map{
		"len": len(result),
		"hasContent": len(result[0]) > 0,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasContent": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonStrings returns json strings -- one item", actual)
}

func Test_Collection_JoinJsonStrings_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	result := c.JoinJsonStrings(",")

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinJsonStrings returns joined json -- comma joiner", actual)
}

func Test_Collection_JoinCsv_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JoinCsv()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinCsv returns csv string -- one item", actual)
}

func Test_Collection_JoinCsvLine_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	result := c.JoinCsvLine()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinCsvLine returns csv lines -- two items", actual)
}

// ============================================================================
// Collection — JsonString / CsvStrings
// ============================================================================

func Test_Collection_JsonString_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JsonString()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns json -- one item", actual)
}

func Test_Collection_JsonString_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	result := c.JsonString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- empty collection", actual)
}

func Test_Collection_CsvStrings_Empty_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	result := c.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns empty slice -- empty collection", actual)
}

func Test_Collection_CsvStrings_NonEmpty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns quoted strings -- one item", actual)
}

// ============================================================================
// EmptyGenericCollection
// ============================================================================

func Test_EmptyGenericCollection_FromAppendsIfTrue(t *testing.T) {
	// Arrange
	c := namevalue.EmptyGenericCollection[string, int]()

	// Act
	actual := args.Map{
		"len": c.Length(),
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "EmptyGenericCollection creates empty -- zero length", actual)
}
