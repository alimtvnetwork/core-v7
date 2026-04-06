package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/namevalue"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Instance ──

func Test_C6_Instance_IsNull(t *testing.T) {
	inst := namevalue.StringAny{Name: "k", Value: "v"}
	var nilInst *namevalue.StringAny
	actual := args.Map{"notNull": !inst.IsNull(), "nilIsNull": nilInst.IsNull()}
	expected := args.Map{"notNull": true, "nilIsNull": true}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IsNull", actual)
}

func Test_C6_Instance_String(t *testing.T) {
	inst := namevalue.StringAny{Name: "key", Value: "val"}
	s := inst.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_C6_Instance_JsonString(t *testing.T) {
	inst := namevalue.StringAny{Name: "key", Value: "val"}
	js := inst.JsonString()
	actual := args.Map{"result": js == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
}

func Test_C6_Instance_Dispose(t *testing.T) {
	inst := namevalue.StringAny{Name: "key", Value: "val"}
	inst.Dispose()
	actual := args.Map{"result": inst.Name != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty after dispose", actual)

	// nil dispose should not panic
	var nilInst *namevalue.StringAny
	nilInst.Dispose()
}

// ── Collection constructors ──

func Test_C6_NewGenericCollection(t *testing.T) {
	c := namevalue.NewGenericCollection[string, any](5)
	actual := args.Map{"len": c.Length(), "empty": c.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "NewGenericCollection returns correct value -- with args", actual)
}

func Test_C6_NewGenericCollectionDefault(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, any]()
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C6_EmptyGenericCollection(t *testing.T) {
	c := namevalue.EmptyGenericCollection[string, any]()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C6_NewGenericCollectionUsing(t *testing.T) {
	items := []namevalue.StringAny{
		{Name: "a", Value: 1},
		{Name: "b", Value: 2},
	}
	// with clone
	c1 := namevalue.NewGenericCollectionUsing[string, any](true, items...)
	// without clone
	c2 := namevalue.NewGenericCollectionUsing[string, any](false, items...)
	// nil items
	c3 := namevalue.NewGenericCollectionUsing[string, any](false)
	actual := args.Map{"c1Len": c1.Length(), "c2Len": c2.Length(), "c3Len": c3.Length()}
	expected := args.Map{"c1Len": 2, "c2Len": 2, "c3Len": 0}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing returns correct value -- with args", actual)
}

// ── NameValuesCollection constructors ──

func Test_C6_NewNameValuesCollection(t *testing.T) {
	c := namevalue.NewNameValuesCollection(5)
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C6_NewCollection(t *testing.T) {
	c := namevalue.NewCollection()
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C6_NewNewNameValuesCollectionUsing(t *testing.T) {
	c := namevalue.NewNewNameValuesCollectionUsing(true, namevalue.StringAny{Name: "x", Value: 1})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_EmptyNameValuesCollection(t *testing.T) {
	c := namevalue.EmptyNameValuesCollection()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ── Collection Add/Adds/Append/Prepend ──

func Test_C6_Collection_Add(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_Adds(t *testing.T) {
	c := namevalue.NewCollection()
	c.Adds(namevalue.StringAny{Name: "a", Value: 1}, namevalue.StringAny{Name: "b", Value: 2})
	c.Adds() // empty
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C6_Collection_Append(t *testing.T) {
	c := namevalue.NewCollection()
	c.Append(namevalue.StringAny{Name: "a", Value: 1})
	c.Append() // empty
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_AppendIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.AppendIf(true, namevalue.StringAny{Name: "a", Value: 1})
	c.AppendIf(false, namevalue.StringAny{Name: "b", Value: 2})
	c.AppendIf(true) // empty items
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_Prepend(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "b", Value: 2})
	c.Prepend(namevalue.StringAny{Name: "a", Value: 1})
	c.Prepend() // empty
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C6_Collection_PrependIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "b", Value: 2})
	c.PrependIf(true, namevalue.StringAny{Name: "a", Value: 1})
	c.PrependIf(false, namevalue.StringAny{Name: "c", Value: 3})
	c.PrependIf(true) // empty items
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C6_Collection_PrependUsingFuncIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "b", Value: 2})
	c.PrependUsingFuncIf(true, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "a", Value: 1}}
	})
	c.PrependUsingFuncIf(false, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "c", Value: 3}}
	})
	c.PrependUsingFuncIf(true, nil) // nil func
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C6_Collection_AppendUsingFuncIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.AppendUsingFuncIf(true, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "a", Value: 1}}
	})
	c.AppendUsingFuncIf(false, func() []namevalue.StringAny {
		return []namevalue.StringAny{{Name: "b", Value: 2}}
	})
	c.AppendUsingFuncIf(true, nil) // nil func
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_AppendPrependIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "mid", Value: 0})
	prepend := []namevalue.StringAny{{Name: "first", Value: 1}}
	appnd := []namevalue.StringAny{{Name: "last", Value: 2}}
	c.AppendPrependIf(true, prepend, appnd)
	c.AppendPrependIf(false, prepend, appnd) // skip
	// Also test with empty slices
	c.AppendPrependIf(true, nil, nil)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C6_Collection_AddsPtr(t *testing.T) {
	c := namevalue.NewCollection()
	item := namevalue.StringAny{Name: "a", Value: 1}
	c.AddsPtr(&item, nil) // nil should be skipped
	c.AddsPtr()           // empty
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_AddsIf(t *testing.T) {
	c := namevalue.NewCollection()
	c.AddsIf(true, namevalue.StringAny{Name: "a", Value: 1})
	c.AddsIf(false, namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ── Collection query methods ──

func Test_C6_Collection_LengthCountEmpty(t *testing.T) {
	c := namevalue.NewCollection()
	var nilC *namevalue.Collection[string, any]
	actual := args.Map{
		"len":      c.Length(),
		"count":    c.Count(),
		"empty":    c.IsEmpty(),
		"hasAny":   c.HasAnyItem(),
		"nilLen":   nilC.Length(),
	}
	expected := args.Map{
		"len":      0,
		"count":    0,
		"empty":    true,
		"hasAny":   false,
		"nilLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "Length/Count/Empty returns empty -- with args", actual)
}

func Test_C6_Collection_LastIndex_HasIndex(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	c.Add(namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"lastIdx": c.LastIndex(), "hasIdx0": c.HasIndex(0), "hasIdx5": c.HasIndex(5)}
	expected := args.Map{"lastIdx": 1, "hasIdx0": true, "hasIdx5": false}
	expected.ShouldBeEqual(t, 0, "LastIndex/HasIndex returns correct value -- with args", actual)
}

// ── Collection string methods ──

func Test_C6_Collection_Strings(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	strs := c.Strings()
	actual := args.Map{"result": len(strs) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_JsonStrings(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	strs := c.JsonStrings()
	actual := args.Map{"result": len(strs) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_JoinJsonStrings(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	s := c.JoinJsonStrings(",")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C6_Collection_Join(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	s := c.Join(",")
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C6_Collection_JoinLines(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	_ = c.JoinLines()
}

func Test_C6_Collection_JoinCsv(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	_ = c.JoinCsv()
}

func Test_C6_Collection_JoinCsvLine(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	_ = c.JoinCsvLine()
}

func Test_C6_Collection_CsvStrings(t *testing.T) {
	c := namevalue.NewCollection()
	// empty
	csv := c.CsvStrings()
	actual := args.Map{"result": len(csv) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// with items
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	csv = c.CsvStrings()
	actual := args.Map{"result": len(csv) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ── Collection IsEqualByString ──

func Test_C6_Collection_IsEqualByString(t *testing.T) {
	c1 := namevalue.NewCollection()
	c1.Add(namevalue.StringAny{Name: "a", Value: 1})
	c2 := namevalue.NewCollection()
	c2.Add(namevalue.StringAny{Name: "a", Value: 1})
	c3 := namevalue.NewCollection()
	c3.Add(namevalue.StringAny{Name: "b", Value: 2})
	var nilC *namevalue.Collection[string, any]

	actual := args.Map{
		"equal":     c1.IsEqualByString(c2),
		"notEqual":  c1.IsEqualByString(c3),
		"nilBoth":   nilC.IsEqualByString(nil),
		"nilLeft":   nilC.IsEqualByString(c1),
		"nilRight":  c1.IsEqualByString(nil),
		"diffLen":   c1.IsEqualByString(namevalue.NewCollection()),
	}
	expected := args.Map{
		"equal":     true,
		"notEqual":  false,
		"nilBoth":   true,
		"nilLeft":   false,
		"nilRight":  false,
		"diffLen":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualByString returns correct value -- with args", actual)
}

// ── Collection JsonString / String ──

func Test_C6_Collection_JsonString(t *testing.T) {
	c := namevalue.NewCollection()
	empty := c.JsonString()
	actual := args.Map{"result": empty != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty collection", actual)
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	js := c.JsonString()
	actual := args.Map{"result": js == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
}

func Test_C6_Collection_String(t *testing.T) {
	c := namevalue.NewCollection()
	empty := c.String()
	actual := args.Map{"result": empty != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty collection", actual)
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	s := c.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_C6_Collection_HasCompiledString_CompiledLazyString(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	// First call compiles
	s1 := c.CompiledLazyString()
	actual := args.Map{"result": s1 == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": c.HasCompiledString()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected compiled", actual)
	// Second call returns cached
	s2 := c.CompiledLazyString()
	actual := args.Map{"result": s1 != s2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
	// Invalidate
	c.InvalidateLazyString()
	actual := args.Map{"result": c.HasCompiledString()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not compiled", actual)

	// nil receiver
	var nilC *namevalue.Collection[string, any]
	nilS := nilC.CompiledLazyString()
	actual := args.Map{"result": nilS != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	nilC.InvalidateLazyString() // should not panic
}

// ── Collection Error ──

func Test_C6_Collection_Error(t *testing.T) {
	c := namevalue.NewCollection()
	actual := args.Map{"result": c.Error() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	c.Add(namevalue.StringAny{Name: "err", Value: "msg"})
	actual := args.Map{"result": c.Error() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C6_Collection_ErrorUsingMessage(t *testing.T) {
	c := namevalue.NewCollection()
	actual := args.Map{"result": c.ErrorUsingMessage("prefix") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
	c.Add(namevalue.StringAny{Name: "err", Value: "msg"})
	err := c.ErrorUsingMessage("prefix")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Collection ConcatNew / ConcatNewPtr ──

func Test_C6_Collection_ConcatNew(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	c2 := c.ConcatNew(namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original should be unchanged", actual)
}

func Test_C6_Collection_ConcatNewPtr(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	item := namevalue.StringAny{Name: "b", Value: 2}
	c2 := c.ConcatNewPtr(&item)
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ── Collection Clone / ClonePtr / Clear / Dispose ──

func Test_C6_Collection_Clone(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	clone := c.Clone()
	actual := args.Map{"result": clone.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_Collection_ClonePtr(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	clone := c.ClonePtr()
	actual := args.Map{"result": clone.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	var nilC *namevalue.Collection[string, any]
	nilClone := nilC.ClonePtr()
	actual := args.Map{"result": nilClone != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil ptr", actual)
}

func Test_C6_Collection_Clear(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	c.Clear()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 after clear", actual)

	var nilC *namevalue.Collection[string, any]
	result := nilC.Clear()
	actual := args.Map{"result": result != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C6_Collection_Dispose(t *testing.T) {
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	c.Dispose()
	actual := args.Map{"result": c.Items != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil after dispose", actual)

	var nilC *namevalue.Collection[string, any]
	nilC.Dispose() // should not panic
}

// ── AppendsIf / PrependsIf ──

func Test_C6_AppendsIf(t *testing.T) {
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(true, items, namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	result2 := namevalue.AppendsIf(false, items, namevalue.StringAny{Name: "c", Value: 3})
	actual := args.Map{"result": len(result2)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	result3 := namevalue.AppendsIf[string, any](true, items)
	actual := args.Map{"result": len(result3)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C6_PrependsIf(t *testing.T) {
	items := []namevalue.StringAny{{Name: "b", Value: 2}}
	result := namevalue.PrependsIf(true, items, namevalue.StringAny{Name: "a", Value: 1})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	result2 := namevalue.PrependsIf(false, items, namevalue.StringAny{Name: "c", Value: 3})
	actual := args.Map{"result": len(result2)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	result3 := namevalue.PrependsIf[string, any](true, items)
	actual := args.Map{"result": len(result3)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}
