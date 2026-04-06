package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Cover: Instance methods, Collection methods not hit by existing tests

func Test_Instance_IsNull_Nil_Cov2(t *testing.T) {
	var inst *namevalue.Instance[string, any]
	actual := args.Map{"result": inst.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_Instance_String_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, any]{Name: "key", Value: "val"}
	actual := args.Map{"result": inst.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_Instance_JsonString_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}
	actual := args.Map{"result": inst.JsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_Instance_Dispose_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}
	inst.Dispose()
	actual := args.Map{"result": inst.Name != "" || inst.Value != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be zeroed", actual)
}

func Test_Instance_Dispose_Nil_Cov2(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	inst.Dispose() // should not panic
}

func Test_Collection_NewGenericCollectionDefault_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c == nil || c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty collection", actual)
}

func Test_Collection_EmptyGenericCollection_Cov2(t *testing.T) {
	c := namevalue.EmptyGenericCollection[string, string]()
	actual := args.Map{"result": c.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Collection_NewGenericCollectionUsing_Nil_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionUsing[string, string](false, nil...)
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Collection_NewGenericCollectionUsing_NoClone_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{
		{Name: "a", Value: "1"},
	}
	c := namevalue.NewGenericCollectionUsing[string, string](false, items...)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_NewGenericCollectionUsing_Clone_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{
		{Name: "a", Value: "1"},
	}
	c := namevalue.NewGenericCollectionUsing[string, string](true, items...)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_Add_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_Adds_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Adds()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_Append_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_Append_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_AppendIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_AppendIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_Prepend_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.Prepend(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 2 || c.Items[0].Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
}

func Test_Collection_Prepend_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Prepend()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_PrependIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Items[0].Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
}

func Test_Collection_PrependIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_PrependUsingFuncIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	actual := args.Map{"result": c.Items[0].Name != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a' first", actual)
}

func Test_Collection_PrependUsingFuncIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependUsingFuncIf(false, nil)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_AppendUsingFuncIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_AppendUsingFuncIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(false, nil)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_AppendPrependIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	prepend := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	appnd := []namevalue.Instance[string, string]{{Name: "c", Value: "3"}}
	c.AppendPrependIf(true, prepend, appnd)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Collection_AppendPrependIf_Skip_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendPrependIf(false, nil, nil)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_AddsPtr_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(item, nil)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1, nil should be skipped", actual)
}

func Test_Collection_AddsPtr_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsPtr()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_CompiledLazyString_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	s1 := c.CompiledLazyString()
	s2 := c.CompiledLazyString() // should use cached
	actual := args.Map{"result": s1 != s2 || s1 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same cached string", actual)
}

func Test_Collection_CompiledLazyString_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"result": c.CompiledLazyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Collection_ConcatNew_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := c.ConcatNew(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original should be unchanged", actual)
}

func Test_Collection_ConcatNewPtr_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := &namevalue.Instance[string, string]{Name: "b", Value: "2"}
	c2 := c.ConcatNewPtr(item)
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_AddsIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_AddsIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_Count_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.Count() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Collection_HasAnyItem_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Collection_HasIndex_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.HasIndex(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.HasIndex(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Collection_IsEqualByString_Cov2(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := c1.Clone()
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal", actual)
}

func Test_Collection_IsEqualByString_DiffLen_Cov2(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_Collection_IsEqualByString_BothNil_Cov2(t *testing.T) {
	var c1, c2 *namevalue.Collection[string, string]
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Collection_IsEqualByString_OneNil_Cov2(t *testing.T) {
	var c1 *namevalue.Collection[string, string]
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal", actual)
}

func Test_Collection_JsonString_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JsonString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Collection_JsonString_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.JsonString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Collection_Error_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.Error() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.Error() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

func Test_Collection_ErrorUsingMessage_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.ErrorUsingMessage("msg") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.ErrorUsingMessage("msg") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
}

func Test_Collection_JoinJsonStrings_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinJsonStrings(",") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Collection_JoinCsv_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Collection_JoinCsvLine_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c.JoinCsvLine() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Collection_Clear_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Clear()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)
}

func Test_Collection_Clear_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.Clear() // should not panic
}

func Test_Collection_Dispose_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Dispose()
	actual := args.Map{"result": c.Items != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil items after dispose", actual)
}

func Test_Collection_Dispose_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.Dispose() // should not panic
}

func Test_Collection_ClonePtr_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	actual := args.Map{"result": c.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_AppendsIf_True_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_AppendsIf_False_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_PrependsIf_True_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.PrependsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": len(result) != 2 || result[0].Name != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'b' first", actual)
}

func Test_PrependsIf_False_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.PrependsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_NameValuesCollection_Aliases_Cov2(t *testing.T) {
	c := namevalue.NewNameValuesCollection(5)
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	c2 := namevalue.NewCollection()
	actual := args.Map{"result": c2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	c3 := namevalue.EmptyNameValuesCollection()
	actual := args.Map{"result": c3 == nil || !c3.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_NewNewNameValuesCollectionUsing_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, any]{{Name: "a", Value: 1}}
	c := namevalue.NewNewNameValuesCollectionUsing(true, items...)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
