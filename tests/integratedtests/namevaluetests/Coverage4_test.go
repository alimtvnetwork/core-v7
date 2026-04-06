package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── AppendsIf / PrependsIf ──

func Test_Cov4_AppendsIf_True(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns non-empty -- true", actual)
}

func Test_Cov4_AppendsIf_False(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns non-empty -- false", actual)
}

func Test_Cov4_PrependsIf_True(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	result := namevalue.PrependsIf(true, items, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": len(result), "first": result[0].Name}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "PrependsIf returns non-empty -- true", actual)
}

func Test_Cov4_PrependsIf_False(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	result := namevalue.PrependsIf(false, items, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependsIf returns non-empty -- false", actual)
}

// ── NameValuesCollection constructors ──

func Test_Cov4_NewNameValuesCollection(t *testing.T) {
	c := namevalue.NewNameValuesCollection(10)
	actual := args.Map{"notNil": c != nil, "isEmpty": c.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewNameValuesCollection returns non-empty -- with args", actual)
}

func Test_Cov4_NewCollection(t *testing.T) {
	c := namevalue.NewCollection()
	actual := args.Map{"notNil": c != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewCollection returns correct value -- with args", actual)
}

func Test_Cov4_EmptyNameValuesCollection(t *testing.T) {
	c := namevalue.EmptyNameValuesCollection()
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyNameValuesCollection returns empty -- with args", actual)
}

func Test_Cov4_NewNewNameValuesCollectionUsing(t *testing.T) {
	items := []namevalue.StringAny{
		{Name: "a", Value: 1},
	}
	c := namevalue.NewNewNameValuesCollectionUsing(true, items...)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewNewNameValuesCollectionUsing returns non-empty -- with args", actual)
}

// ── Collection extended methods ──

func Test_Cov4_Collection_JsonString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JsonString()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyResult := emptyC.JsonString()
	actual := args.Map{"notEmpty": result != "", "emptyResult": emptyResult}
	expected := args.Map{"notEmpty": true, "emptyResult": ""}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonString", actual)
}

func Test_Cov4_Collection_Error(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	err := c.Error()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyErr := emptyC.Error()
	actual := args.Map{"hasErr": err != nil, "emptyNil": emptyErr == nil}
	expected := args.Map{"hasErr": true, "emptyNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns error -- Error", actual)
}

func Test_Cov4_Collection_ErrorUsingMessage(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	err := c.ErrorUsingMessage("prefix:")
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyErr := emptyC.ErrorUsingMessage("prefix:")
	actual := args.Map{"hasErr": err != nil, "emptyNil": emptyErr == nil}
	expected := args.Map{"hasErr": true, "emptyNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns error -- ErrorUsingMessage", actual)
}

func Test_Cov4_Collection_CsvStrings(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.CsvStrings()
	emptyC := namevalue.NewGenericCollectionDefault[string, string]()
	emptyResult := emptyC.CsvStrings()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 1, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- CsvStrings", actual)
}

func Test_Cov4_Collection_JoinCsv(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinCsv()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinCsv", actual)
}

func Test_Cov4_Collection_JoinCsvLine(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinCsvLine()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinCsvLine", actual)
}

func Test_Cov4_Collection_JoinJsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JoinJsonStrings(",")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JoinJsonStrings", actual)
}

func Test_Cov4_Collection_JsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.JsonStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- JsonStrings", actual)
}

func Test_Cov4_Collection_Clear(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Clear()
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Clear", actual)
}

func Test_Cov4_Collection_Clear_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	result := c.Clear()
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Clear nil", actual)
}

func Test_Cov4_Collection_Dispose(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Dispose()
	actual := args.Map{"isNil": c.Items == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Dispose", actual)
}

func Test_Cov4_Collection_Dispose_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Dispose nil", actual)
}

func Test_Cov4_Collection_ClonePtr_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	result := c.ClonePtr()
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- ClonePtr nil", actual)
}

func Test_Cov4_Collection_HasCompiledString(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"before": c.HasCompiledString()}
	c.CompiledLazyString()
	actual["after"] = c.HasCompiledString()
	expected := args.Map{"before": false, "after": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasCompiledString", actual)
}

func Test_Cov4_Collection_HasCompiledString_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"result": c.HasCompiledString()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- HasCompiledString nil", actual)
}

func Test_Cov4_Collection_InvalidateLazyString_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.InvalidateLazyString() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- InvalidateLazyString nil", actual)
}

func Test_Cov4_Collection_String_WithCache(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	_ = c.CompiledLazyString() // populate cache
	result := c.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- String with cache", actual)
}

func Test_Cov4_Collection_Length_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns nil -- Length nil", actual)
}

// ── Instance JsonString invalid ──

func Test_Cov4_Instance_JsonString_UnserializableValue(t *testing.T) {
	ch := make(chan int)
	inst := namevalue.Instance[string, any]{Name: "ch", Value: ch}
	result := inst.JsonString()
	actual := args.Map{"empty": result}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- JsonString unserializable", actual)
}
