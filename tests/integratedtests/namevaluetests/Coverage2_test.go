package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
)

// Cover: Instance methods, Collection methods not hit by existing tests

func Test_Instance_IsNull_Nil_Cov2(t *testing.T) {
	var inst *namevalue.Instance[string, any]
	if !inst.IsNull() {
		t.Error("nil should be null")
	}
}

func Test_Instance_String_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, any]{Name: "key", Value: "val"}
	if inst.String() == "" {
		t.Error("should not be empty")
	}
}

func Test_Instance_JsonString_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}
	if inst.JsonString() == "" {
		t.Error("should not be empty")
	}
}

func Test_Instance_Dispose_Cov2(t *testing.T) {
	inst := namevalue.Instance[string, string]{Name: "key", Value: "val"}
	inst.Dispose()
	if inst.Name != "" || inst.Value != "" {
		t.Error("should be zeroed")
	}
}

func Test_Instance_Dispose_Nil_Cov2(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	inst.Dispose() // should not panic
}

func Test_Collection_NewGenericCollectionDefault_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c == nil || c.Length() != 0 {
		t.Error("expected empty collection")
	}
}

func Test_Collection_EmptyGenericCollection_Cov2(t *testing.T) {
	c := namevalue.EmptyGenericCollection[string, string]()
	if !c.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_Collection_NewGenericCollectionUsing_Nil_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionUsing[string, string](false, nil...)
	if c == nil {
		t.Error("expected non-nil")
	}
}

func Test_Collection_NewGenericCollectionUsing_NoClone_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{
		{Name: "a", Value: "1"},
	}
	c := namevalue.NewGenericCollectionUsing[string, string](false, items...)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_NewGenericCollectionUsing_Clone_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{
		{Name: "a", Value: "1"},
	}
	c := namevalue.NewGenericCollectionUsing[string, string](true, items...)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_Add_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_Adds_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Adds()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_Append_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_Append_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_AppendIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_AppendIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_Prepend_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.Prepend(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 2 || c.Items[0].Name != "a" {
		t.Error("expected 'a' first")
	}
}

func Test_Collection_Prepend_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Prepend()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_PrependIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Items[0].Name != "a" {
		t.Error("expected 'a' first")
	}
}

func Test_Collection_PrependIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_PrependUsingFuncIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	if c.Items[0].Name != "a" {
		t.Error("expected 'a' first")
	}
}

func Test_Collection_PrependUsingFuncIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependUsingFuncIf(false, nil)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_AppendUsingFuncIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_AppendUsingFuncIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(false, nil)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_AppendPrependIf_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	prepend := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	appnd := []namevalue.Instance[string, string]{{Name: "c", Value: "3"}}
	c.AppendPrependIf(true, prepend, appnd)
	if c.Length() != 3 {
		t.Errorf("expected 3, got %d", c.Length())
	}
}

func Test_Collection_AppendPrependIf_Skip_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendPrependIf(false, nil, nil)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_AddsPtr_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(item, nil)
	if c.Length() != 1 {
		t.Error("expected 1, nil should be skipped")
	}
}

func Test_Collection_AddsPtr_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsPtr()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_CompiledLazyString_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	s1 := c.CompiledLazyString()
	s2 := c.CompiledLazyString() // should use cached
	if s1 != s2 || s1 == "" {
		t.Error("expected same cached string")
	}
}

func Test_Collection_CompiledLazyString_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	if c.CompiledLazyString() != "" {
		t.Error("nil should be empty")
	}
}

func Test_Collection_ConcatNew_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := c.ConcatNew(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if c2.Length() != 2 {
		t.Error("expected 2")
	}
	if c.Length() != 1 {
		t.Error("original should be unchanged")
	}
}

func Test_Collection_ConcatNewPtr_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := &namevalue.Instance[string, string]{Name: "b", Value: "2"}
	c2 := c.ConcatNewPtr(item)
	if c2.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Collection_AddsIf_True_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Collection_AddsIf_False_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_Count_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.Count() != 0 {
		t.Error("expected 0")
	}
}

func Test_Collection_HasAnyItem_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.HasAnyItem() {
		t.Error("expected false")
	}
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if !c.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_Collection_HasIndex_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.HasIndex(0) {
		t.Error("expected false")
	}
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if !c.HasIndex(0) {
		t.Error("expected true")
	}
}

func Test_Collection_IsEqualByString_Cov2(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := c1.Clone()
	if !c1.IsEqualByString(c2) {
		t.Error("should be equal")
	}
}

func Test_Collection_IsEqualByString_DiffLen_Cov2(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	if c1.IsEqualByString(c2) {
		t.Error("should not be equal")
	}
}

func Test_Collection_IsEqualByString_BothNil_Cov2(t *testing.T) {
	var c1, c2 *namevalue.Collection[string, string]
	if !c1.IsEqualByString(c2) {
		t.Error("both nil should be equal")
	}
}

func Test_Collection_IsEqualByString_OneNil_Cov2(t *testing.T) {
	var c1 *namevalue.Collection[string, string]
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c1.IsEqualByString(c2) {
		t.Error("should not be equal")
	}
}

func Test_Collection_JsonString_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JsonString() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Collection_JsonString_Empty_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.JsonString() != "" {
		t.Error("expected empty")
	}
}

func Test_Collection_Error_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.Error() != nil {
		t.Error("empty should return nil")
	}
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.Error() == nil {
		t.Error("should return error")
	}
}

func Test_Collection_ErrorUsingMessage_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	if c.ErrorUsingMessage("msg") != nil {
		t.Error("empty should return nil")
	}
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.ErrorUsingMessage("msg") == nil {
		t.Error("should return error")
	}
}

func Test_Collection_JoinJsonStrings_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinJsonStrings(",") == "" {
		t.Error("expected non-empty")
	}
}

func Test_Collection_JoinCsv_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinCsv() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Collection_JoinCsvLine_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	if c.JoinCsvLine() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Collection_Clear_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Clear()
	if c.Length() != 0 {
		t.Error("expected 0 after clear")
	}
}

func Test_Collection_Clear_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.Clear() // should not panic
}

func Test_Collection_Dispose_Cov2(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Dispose()
	if c.Items != nil {
		t.Error("expected nil items after dispose")
	}
}

func Test_Collection_Dispose_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	c.Dispose() // should not panic
}

func Test_Collection_ClonePtr_Nil_Cov2(t *testing.T) {
	var c *namevalue.Collection[string, string]
	if c.ClonePtr() != nil {
		t.Error("expected nil")
	}
}

func Test_AppendsIf_True_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if len(result) != 2 {
		t.Error("expected 2")
	}
}

func Test_AppendsIf_False_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.AppendsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if len(result) != 1 {
		t.Error("expected 1")
	}
}

func Test_PrependsIf_True_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.PrependsIf(true, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if len(result) != 2 || result[0].Name != "b" {
		t.Error("expected 'b' first")
	}
}

func Test_PrependsIf_False_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	result := namevalue.PrependsIf(false, items, namevalue.Instance[string, string]{Name: "b", Value: "2"})
	if len(result) != 1 {
		t.Error("expected 1")
	}
}

func Test_NameValuesCollection_Aliases_Cov2(t *testing.T) {
	c := namevalue.NewNameValuesCollection(5)
	if c == nil {
		t.Error("expected non-nil")
	}
	c2 := namevalue.NewCollection()
	if c2 == nil {
		t.Error("expected non-nil")
	}
	c3 := namevalue.EmptyNameValuesCollection()
	if c3 == nil || !c3.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_NewNewNameValuesCollectionUsing_Cov2(t *testing.T) {
	items := []namevalue.Instance[string, any]{{Name: "a", Value: 1}}
	c := namevalue.NewNewNameValuesCollectionUsing(true, items...)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}
