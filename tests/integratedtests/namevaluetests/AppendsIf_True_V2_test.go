package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── AppendsIf / PrependsIf ──

func Test_AppendsIf_True_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns non-empty -- true", actual)
}

func Test_AppendsIf_False_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns non-empty -- false", actual)
}

func Test_PrependsIf_True_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	result := namevalue.PrependsIf(true, items, namevalue.Instance[string, string]{Name: "a", Value: "1"})

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
	expected.ShouldBeEqual(t, 0, "PrependsIf returns non-empty -- true", actual)
}

func Test_PrependsIf_False_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	result := namevalue.PrependsIf(false, items, namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependsIf returns non-empty -- false", actual)
}

// ── NameValuesCollection constructors ──

func Test_NewNameValuesCollection_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewNameValuesCollection(10)

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"isEmpty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewNameValuesCollection returns non-empty -- with args", actual)
}

func Test_NewCollection_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewCollection()

	// Act
	actual := args.Map{"notNil": c != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewCollection returns correct value -- with args", actual)
}

func Test_EmptyNameValuesCollection_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.EmptyNameValuesCollection()

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyNameValuesCollection returns empty -- with args", actual)
}

func Test_NewNewNameValuesCollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}
	c := namevalue.NewNewNameValuesCollectionUsing(true, items...)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewNewNameValuesCollectionUsing returns non-empty -- with args", actual)
}

// ── Collection extended methods ──

func Test_Collection_JsonString(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JsonString()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyResult := emptyC.JsonString()

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"emptyResult": emptyResult,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"emptyResult": "",
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonString", actual)
}

func Test_Collection_Error_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	err := c.Error()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyErr := emptyC.Error()

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"emptyNil": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"emptyNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns error -- Error", actual)
}

func Test_Collection_ErrorUsingMessage_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	err := c.ErrorUsingMessage("prefix:")
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyErr := emptyC.ErrorUsingMessage("prefix:")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"emptyNil": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"emptyNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns error -- ErrorUsingMessage", actual)
}

func Test_Collection_CsvStrings_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.CsvStrings()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyResult := emptyC.CsvStrings()

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CsvStrings", actual)
}

func Test_Collection_JoinCsv_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinCsv", actual)
}

func Test_Collection_JoinCsvLine_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinCsvLine()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinCsvLine", actual)
}

func Test_Collection_JoinJsonStrings_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinJsonStrings(",")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinJsonStrings", actual)
}

func Test_Collection_JsonStrings_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JsonStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonStrings", actual)
}

func Test_Collection_Clear(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Clear()

	// Act
	actual := args.Map{"isEmpty": c.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Clear", actual)
}

func Test_Collection_Clear_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	result := c.Clear()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Clear nil", actual)
}

func Test_Collection_Dispose(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Dispose()

	// Act
	actual := args.Map{"isNil": c.Items == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Dispose", actual)
}

func Test_Collection_Dispose_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	c.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Dispose nil", actual)
}

func Test_Collection_ClonePtr_Nil_FromAppendsIfTrueV2(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	result := c.ClonePtr()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- ClonePtr nil", actual)
}

func Test_Collection_HasCompiledString(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})

	// Act
	actual := args.Map{"before": c.HasCompiledString()}
	c.CompiledLazyString()
	actual["after"] = c.HasCompiledString()

	// Assert
	expected := args.Map{
		"before": false,
		"after": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasCompiledString", actual)
}

func Test_Collection_HasCompiledString_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]

	// Act
	actual := args.Map{"result": c.HasCompiledString()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- HasCompiledString nil", actual)
}

func Test_Collection_InvalidateLazyString_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	c.InvalidateLazyString() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- InvalidateLazyString nil", actual)
}

func Test_Collection_String_WithCache(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	_ = c.CompiledLazyString() // populate cache
	result := c.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- String with cache", actual)
}

func Test_Collection_Length_Nil(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Length nil", actual)
}

// ── Instance JsonString invalid ──

func Test_Instance_JsonString_UnserializableValue(t *testing.T) {
	// Arrange
	ch := make(chan int)
	inst := namevalue.Instance[string, any]{Name: "ch", Value: ch}
	result := inst.JsonString()

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- JsonString unserializable", actual)
}
