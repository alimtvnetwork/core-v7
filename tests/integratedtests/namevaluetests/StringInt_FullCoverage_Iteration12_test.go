package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ==========================================================================
// StringInt — full Collection[string, int] coverage
// ==========================================================================

func Test_StringInt_FullCoverage(t *testing.T) {
	// Arrange
	// Constructors
	c := namevalue.NewGenericCollection[string, int](5)
	cd := namevalue.NewGenericCollectionDefault[string, int]()
	ce := namevalue.EmptyGenericCollection[string, int]()

	// Act
	actual := args.Map{
		"c": c.IsEmpty(),
		"cd": cd.IsEmpty(),
		"ce": ce.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"c": true,
		"cd": true,
		"ce": true,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- constructors", actual)
}

func Test_StringInt_CollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.StringInt{{Name: "x", Value: 10}, {Name: "y", Value: 20}}
	c1 := namevalue.NewGenericCollectionUsing[string, int](true, items...)
	c2 := namevalue.NewGenericCollectionUsing[string, int](false, items...)
	c3 := namevalue.NewGenericCollectionUsing[string, int](true)

	// Act
	actual := args.Map{
		"c1": c1.Length(),
		"c2": c2.Length(),
		"c3": c3.Length(),
	}

	// Assert
	expected := args.Map{
		"c1": 2,
		"c2": 2,
		"c3": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- CollectionUsing", actual)
}

func Test_StringInt_AddMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Adds(namevalue.StringInt{Name: "b", Value: 2}, namevalue.StringInt{Name: "c", Value: 3})
	c.Adds() // empty
	c.Append(namevalue.StringInt{Name: "d", Value: 4})
	c.Append() // empty
	c.Prepend(namevalue.StringInt{Name: "z", Value: 0})
	c.Prepend() // empty
	c.AppendIf(true, namevalue.StringInt{Name: "e", Value: 5})
	c.AppendIf(false, namevalue.StringInt{Name: "f", Value: 6})
	c.AppendIf(true) // empty items
	c.PrependIf(true, namevalue.StringInt{Name: "y", Value: -1})
	c.PrependIf(false, namevalue.StringInt{Name: "w", Value: -2})
	c.PrependIf(true) // empty items

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- add methods", actual)
}

func Test_StringInt_FuncAppendPrepend(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	c.PrependUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "first", Value: -1}}
	})
	c.PrependUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "skip", Value: -2}}
	})
	c.PrependUsingFuncIf(true, nil) // nil func
	c.AppendUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "last", Value: 1}}
	})
	c.AppendUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "skip", Value: 2}}
	})
	c.AppendUsingFuncIf(true, nil) // nil func

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.Items[0].Name,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "first",
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- func append/prepend", actual)
}

func Test_StringInt_AppendPrependIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	pre := []namevalue.StringInt{{Name: "a", Value: -1}}
	app := []namevalue.StringInt{{Name: "z", Value: 1}}
	c.AppendPrependIf(true, pre, app)
	c.AppendPrependIf(false, pre, app)
	c.AppendPrependIf(true, nil, nil) // empty slices

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- AppendPrependIf", actual)
}

func Test_StringInt_AddsPtr(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	item := namevalue.StringInt{Name: "a", Value: 1}
	c.AddsPtr(&item, nil)
	c.AddsPtr() // empty

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- AddsPtr", actual)
}

func Test_StringInt_AddsIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.AddsIf(true, namevalue.StringInt{Name: "a", Value: 1})
	c.AddsIf(false, namevalue.StringInt{Name: "b", Value: 2})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- AddsIf", actual)
}

func Test_StringInt_QueryMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})

	var nilC *namevalue.Collection[string, int]

	// Act
	actual := args.Map{
		"len": c.Length(), "count": c.Count(), "empty": c.IsEmpty(),
		"hasAny": c.HasAnyItem(), "lastIdx": c.LastIndex(),
		"hasIdx0": c.HasIndex(0), "hasIdx5": c.HasIndex(5),
		"nilLen": nilC.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 2, "count": 2, "empty": false,
		"hasAny": true, "lastIdx": 1,
		"hasIdx0": true, "hasIdx5": false,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- query methods", actual)
}

func Test_StringInt_StringMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "a", Value: 1})

	// Act
	actual := args.Map{
		"strings":    len(c.Strings()) > 0,
		"json":       len(c.JsonStrings()) > 0,
		"joinJson":   c.JoinJsonStrings(",") != "",
		"join":       c.Join(",") != "",
		"joinLines":  c.JoinLines() != "",
		"joinCsv":    c.JoinCsv() != "",
		"joinCsvLn":  c.JoinCsvLine() != "",
		"csvStrings": len(c.CsvStrings()) > 0,
		"jsonString": c.JsonString() != "",
		"string":     c.String() != "",
	}

	// Assert
	expected := args.Map{
		"strings": true, "json": true, "joinJson": true,
		"join": true, "joinLines": true, "joinCsv": true,
		"joinCsvLn": true, "csvStrings": true, "jsonString": true,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- string methods", actual)
}

func Test_StringInt_LazyString(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	s1 := c.CompiledLazyString()
	s2 := c.CompiledLazyString() // cached
	c.InvalidateLazyString()

	var nilC *namevalue.Collection[string, int]

	// Act
	actual := args.Map{
		"same":    s1 == s2,
		"hasPre":  c.HasCompiledString(),
		"nilComp": nilC.HasCompiledString(),
		"nilLazy": nilC.CompiledLazyString(),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"hasPre": false,
		"nilComp": false,
		"nilLazy": "",
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- lazy string", actual)

	nilC.InvalidateLazyString() // should not panic
}

func Test_StringInt_IsEqualByString(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, int]()
	c1.Add(namevalue.StringInt{Name: "a", Value: 1})
	c2 := namevalue.NewGenericCollectionDefault[string, int]()
	c2.Add(namevalue.StringInt{Name: "a", Value: 1})
	c3 := namevalue.NewGenericCollectionDefault[string, int]()
	c3.Add(namevalue.StringInt{Name: "b", Value: 2})
	var nilC *namevalue.Collection[string, int]

	// Act
	actual := args.Map{
		"equal":    c1.IsEqualByString(c2),
		"notEqual": c1.IsEqualByString(c3),
		"nilBoth":  nilC.IsEqualByString(nil),
		"nilOne":   c1.IsEqualByString(nil),
		"diffLen":  c1.IsEqualByString(namevalue.EmptyGenericCollection[string, int]()),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"nilBoth": true,
		"nilOne": false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- IsEqualByString", actual)
}

func Test_StringInt_ErrorMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	nilErr := c.Error()
	nilMsgErr := c.ErrorUsingMessage("prefix")
	c.Add(namevalue.StringInt{Name: "err", Value: 42})

	// Act
	actual := args.Map{
		"nilErr":    nilErr == nil,
		"nilMsgErr": nilMsgErr == nil,
		"hasErr":    c.Error() != nil,
		"hasMsgErr": c.ErrorUsingMessage("prefix") != nil,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"nilMsgErr": true,
		"hasErr": true,
		"hasMsgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns error -- error methods", actual)
}

func Test_StringInt_ConcatCloneClearDispose(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()
	c.Add(namevalue.StringInt{Name: "a", Value: 1})

	cn := c.ConcatNew(namevalue.StringInt{Name: "b", Value: 2})
	item := namevalue.StringInt{Name: "c", Value: 3}
	cnp := c.ConcatNewPtr(&item)
	cl := c.Clone()
	clp := c.ClonePtr()

	var nilC *namevalue.Collection[string, int]
	nilClp := nilC.ClonePtr()

	// Act
	actual := args.Map{
		"concatLen": cn.Length(), "concatPtrLen": cnp.Length(),
		"cloneLen": cl.Length(), "clonePtrLen": clp.Length(),
		"nilClone": nilClp == nil, "origLen": c.Length(),
	}

	// Assert
	expected := args.Map{
		"concatLen": 2, "concatPtrLen": 2,
		"cloneLen": 1, "clonePtrLen": 1,
		"nilClone": true, "origLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- concat/clone", actual)

	c.Clear()
	actual2 := args.Map{"afterClear": c.IsEmpty()}
	expected2 := args.Map{"afterClear": true}
	expected2.ShouldBeEqual(t, 0, "StringInt returns correct value -- clear", actual2)

	nilC.Clear()  // should not panic
	nilC.Dispose() // should not panic
}

func Test_StringInt_CsvStrings_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()

	// Act
	actual := args.Map{"len": len(c.CsvStrings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringInt returns empty -- CsvStrings empty", actual)
}

func Test_StringInt_JsonString_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, int]()

	// Act
	actual := args.Map{"val": c.JsonString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "StringInt returns empty -- JsonString empty", actual)
}

// ==========================================================================
// StringMapAny — Collection[string, map[string]any] coverage
// ==========================================================================

func Test_StringMapAny_FullCoverage(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	ce := namevalue.EmptyGenericCollection[string, map[string]any]()
	c5 := namevalue.NewGenericCollection[string, map[string]any](5)

	// Act
	actual := args.Map{
		"c": c.IsEmpty(),
		"ce": ce.IsEmpty(),
		"c5": c5.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"c": true,
		"ce": true,
		"c5": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- constructors", actual)
}

func Test_StringMapAny_AddAndQuery(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	c.Add(namevalue.StringMapAny{Name: "cfg", Value: map[string]any{"k": "v"}})
	c.Adds(namevalue.StringMapAny{Name: "cfg2", Value: map[string]any{"k2": "v2"}})
	c.Adds() // empty
	c.Append(namevalue.StringMapAny{Name: "cfg3", Value: map[string]any{"k3": "v3"}})
	c.Append()                                                                          // empty
	c.Prepend(namevalue.StringMapAny{Name: "cfg0", Value: map[string]any{"k0": "v0"}})
	c.Prepend()                                                                          // empty
	c.AppendIf(true, namevalue.StringMapAny{Name: "cfg4", Value: map[string]any{"k4": "v4"}})
	c.AppendIf(false, namevalue.StringMapAny{Name: "skip", Value: nil})
	c.PrependIf(true, namevalue.StringMapAny{Name: "cfgP", Value: map[string]any{"kP": "vP"}})
	c.PrependIf(false, namevalue.StringMapAny{Name: "skip2", Value: nil})

	// Act
	actual := args.Map{
		"len": c.Length(), "count": c.Count(), "empty": c.IsEmpty(),
		"hasAny": c.HasAnyItem(), "lastIdx": c.LastIndex(),
		"hasIdx0": c.HasIndex(0),
	}

	// Assert
	expected := args.Map{
		"len": 6, "count": 6, "empty": false,
		"hasAny": true, "lastIdx": 5,
		"hasIdx0": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- add and query", actual)
}

func Test_StringMapAny_StringMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	c.Add(namevalue.StringMapAny{Name: "cfg", Value: map[string]any{"k": "v"}})

	// Act
	actual := args.Map{
		"strings":  len(c.Strings()) > 0,
		"join":     c.Join(",") != "",
		"string":   c.String() != "",
	}

	// Assert
	expected := args.Map{
		"strings": true,
		"join": true,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- string methods", actual)
}

func Test_StringMapAny_CloneDispose(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	c.Add(namevalue.StringMapAny{Name: "a", Value: map[string]any{"k": "v"}})
	cl := c.Clone()
	c.Dispose()

	// Act
	actual := args.Map{
		"cloneLen": cl.Length(),
		"disposed": c.Items == nil,
	}

	// Assert
	expected := args.Map{
		"cloneLen": 1,
		"disposed": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- clone/dispose", actual)
}

func Test_StringMapAny_CollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.StringMapAny{{Name: "a", Value: map[string]any{"k": 1}}}
	c1 := namevalue.NewGenericCollectionUsing[string, map[string]any](true, items...)
	c2 := namevalue.NewGenericCollectionUsing[string, map[string]any](false, items...)

	// Act
	actual := args.Map{
		"c1": c1.Length(),
		"c2": c2.Length(),
	}

	// Assert
	expected := args.Map{
		"c1": 1,
		"c2": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- CollectionUsing", actual)
}

func Test_StringMapAny_FuncIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	c.PrependUsingFuncIf(true, func() []namevalue.StringMapAny {
		return []namevalue.StringMapAny{{Name: "a", Value: map[string]any{"k": "v"}}}
	})
	c.AppendUsingFuncIf(true, func() []namevalue.StringMapAny {
		return []namevalue.StringMapAny{{Name: "b", Value: map[string]any{"k": "v"}}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns correct value -- FuncIf", actual)
}

// ==========================================================================
// StringMapString — Collection[string, map[string]string] coverage
// ==========================================================================

func Test_StringMapString_FullCoverage(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.Add(namevalue.StringMapString{Name: "cfg", Value: map[string]string{"k": "v"}})
	c.Adds(namevalue.StringMapString{Name: "cfg2", Value: map[string]string{"k2": "v2"}})
	c.Append(namevalue.StringMapString{Name: "cfg3", Value: map[string]string{"k3": "v3"}})
	c.Prepend(namevalue.StringMapString{Name: "cfg0", Value: map[string]string{"k0": "v0"}})

	// Act
	actual := args.Map{
		"len": c.Length(), "count": c.Count(),
		"lastIdx": c.LastIndex(), "hasAny": c.HasAnyItem(),
		"strings": len(c.Strings()) > 0, "join": c.Join(",") != "",
	}

	// Assert
	expected := args.Map{
		"len": 4, "count": 4,
		"lastIdx": 3, "hasAny": true,
		"strings": true, "join": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- full coverage", actual)
}

func Test_StringMapString_CloneAndConcat(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.Add(namevalue.StringMapString{Name: "a", Value: map[string]string{"k": "v"}})
	cl := c.Clone()
	cn := c.ConcatNew(namevalue.StringMapString{Name: "b", Value: map[string]string{"k2": "v2"}})
	item := namevalue.StringMapString{Name: "c", Value: map[string]string{"k3": "v3"}}
	cnp := c.ConcatNewPtr(&item)

	// Act
	actual := args.Map{
		"clone": cl.Length(),
		"concat": cn.Length(),
		"concatPtr": cnp.Length(),
	}

	// Assert
	expected := args.Map{
		"clone": 1,
		"concat": 2,
		"concatPtr": 2,
	}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- clone/concat", actual)
}

func Test_StringMapString_AddsIf_AddsPtr(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.AddsIf(true, namevalue.StringMapString{Name: "a", Value: map[string]string{"k": "v"}})
	c.AddsIf(false, namevalue.StringMapString{Name: "b", Value: map[string]string{"k2": "v2"}})
	item := namevalue.StringMapString{Name: "c", Value: map[string]string{"k3": "v3"}}
	c.AddsPtr(&item, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- AddsIf/AddsPtr", actual)
}

func Test_StringMapString_LazyString(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.Add(namevalue.StringMapString{Name: "a", Value: map[string]string{"k": "v"}})
	s1 := c.CompiledLazyString()
	s2 := c.CompiledLazyString()
	c.InvalidateLazyString()

	// Act
	actual := args.Map{
		"same": s1 == s2,
		"invalidated": !c.HasCompiledString(),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"invalidated": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- lazy string", actual)
}

// ==========================================================================
// Instance — type-specific coverage for StringInt, StringMapAny, StringMapString
// ==========================================================================

func Test_Instance_StringInt(t *testing.T) {
	// Arrange
	inst := namevalue.StringInt{Name: "count", Value: 42}
	s := inst.String()
	js := inst.JsonString()
	inst.Dispose()

	// Act
	actual := args.Map{
		"string": s != "",
		"json": js != "",
		"disposed": inst.Name,
	}

	// Assert
	expected := args.Map{
		"string": true,
		"json": true,
		"disposed": "",
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringInt", actual)
}

func Test_Instance_StringMapAny(t *testing.T) {
	// Arrange
	inst := namevalue.StringMapAny{Name: "cfg", Value: map[string]any{"k": "v"}}
	s := inst.String()
	js := inst.JsonString()
	inst.Dispose()

	// Act
	actual := args.Map{
		"string": s != "",
		"json": js != "",
		"disposed": inst.Name,
	}

	// Assert
	expected := args.Map{
		"string": true,
		"json": true,
		"disposed": "",
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringMapAny", actual)
}

func Test_Instance_StringMapString(t *testing.T) {
	// Arrange
	inst := namevalue.StringMapString{Name: "cfg", Value: map[string]string{"k": "v"}}
	s := inst.String()
	js := inst.JsonString()
	inst.Dispose()

	// Act
	actual := args.Map{
		"string": s != "",
		"json": js != "",
		"disposed": inst.Name,
	}

	// Assert
	expected := args.Map{
		"string": true,
		"json": true,
		"disposed": "",
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringMapString", actual)
}

// ==========================================================================
// AppendsIf / PrependsIf — type-specific coverage
// ==========================================================================

func Test_AppendsIf_StringInt(t *testing.T) {
	// Arrange
	items := []namevalue.StringInt{{Name: "a", Value: 1}}
	r1 := namevalue.AppendsIf(true, items, namevalue.StringInt{Name: "b", Value: 2})
	r2 := namevalue.AppendsIf(false, items, namevalue.StringInt{Name: "c", Value: 3})
	r3 := namevalue.AppendsIf[string, int](true, items)

	// Act
	actual := args.Map{
		"r1": len(r1),
		"r2": len(r2),
		"r3": len(r3),
	}

	// Assert
	expected := args.Map{
		"r1": 2,
		"r2": 1,
		"r3": 1,
	}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns correct value -- StringInt", actual)
}

func Test_PrependsIf_StringInt(t *testing.T) {
	// Arrange
	items := []namevalue.StringInt{{Name: "b", Value: 2}}
	r1 := namevalue.PrependsIf(true, items, namevalue.StringInt{Name: "a", Value: 1})
	r2 := namevalue.PrependsIf(false, items, namevalue.StringInt{Name: "c", Value: 3})
	r3 := namevalue.PrependsIf[string, int](true, items)

	// Act
	actual := args.Map{
		"r1": len(r1),
		"r2": len(r2),
		"r3": len(r3),
	}

	// Assert
	expected := args.Map{
		"r1": 2,
		"r2": 1,
		"r3": 1,
	}
	expected.ShouldBeEqual(t, 0, "PrependsIf returns correct value -- StringInt", actual)
}

func Test_AppendsIf_StringMapAny(t *testing.T) {
	// Arrange
	items := []namevalue.StringMapAny{{Name: "a", Value: map[string]any{"k": 1}}}
	r := namevalue.AppendsIf(true, items, namevalue.StringMapAny{Name: "b", Value: map[string]any{"k": 2}})

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendsIf returns correct value -- StringMapAny", actual)
}

func Test_PrependsIf_StringMapString(t *testing.T) {
	// Arrange
	items := []namevalue.StringMapString{{Name: "b", Value: map[string]string{"k": "v"}}}
	r := namevalue.PrependsIf(true, items, namevalue.StringMapString{Name: "a", Value: map[string]string{"k": "v"}})

	// Act
	actual := args.Map{
		"len": len(r),
		"first": r[0].Name,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "PrependsIf returns correct value -- StringMapString", actual)
}

// ==========================================================================
// IsEqualByString — type-specific for StringInt, StringMapString
// ==========================================================================

func Test_StringInt_IsEqualByString_Clone(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, int]()
	c1.Add(namevalue.StringInt{Name: "a", Value: 1})
	c2 := c1.Clone()

	// Act
	actual := args.Map{"equal": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "StringInt returns correct value -- IsEqualByString clone", actual)
}

func Test_StringMapString_IsEqualByString(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c1.Add(namevalue.StringMapString{Name: "a", Value: map[string]string{"k": "v"}})
	c2 := c1.Clone()

	// Act
	actual := args.Map{"equal": c1.IsEqualByString(c2)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- IsEqualByString", actual)
}

// ==========================================================================
// Error/ErrorUsingMessage — type-specific
// ==========================================================================

func Test_StringInt_ErrorMethods_Empty(t *testing.T) {
	// Arrange
	c := namevalue.EmptyGenericCollection[string, int]()

	// Act
	actual := args.Map{
		"err": c.Error() == nil,
		"msgErr": c.ErrorUsingMessage("p") == nil,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringInt returns empty -- Error empty", actual)
}

func Test_StringMapAny_ErrorMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]any]()
	c.Add(namevalue.StringMapAny{Name: "e", Value: map[string]any{"k": "v"}})

	// Act
	actual := args.Map{
		"err": c.Error() != nil,
		"msgErr": c.ErrorUsingMessage("p") != nil,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"msgErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringMapAny returns error -- Error", actual)
}

// ==========================================================================
// AppendPrependIf, CollectionUsing — type-specific for StringMapString
// ==========================================================================

func Test_StringMapString_AppendPrependIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.Add(namevalue.StringMapString{Name: "mid", Value: map[string]string{"k": "v"}})
	pre := []namevalue.StringMapString{{Name: "first", Value: map[string]string{"k": "v"}}}
	app := []namevalue.StringMapString{{Name: "last", Value: map[string]string{"k": "v"}}}
	c.AppendPrependIf(true, pre, app)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- AppendPrependIf", actual)
}

func Test_StringMapString_FuncIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, map[string]string]()
	c.PrependUsingFuncIf(true, func() []namevalue.StringMapString {
		return []namevalue.StringMapString{{Name: "a", Value: map[string]string{"k": "v"}}}
	})
	c.AppendUsingFuncIf(true, func() []namevalue.StringMapString {
		return []namevalue.StringMapString{{Name: "b", Value: map[string]string{"k": "v"}}}
	})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- FuncIf", actual)
}

func Test_StringMapString_CollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.StringMapString{{Name: "a", Value: map[string]string{"k": "v"}}}
	c1 := namevalue.NewGenericCollectionUsing[string, map[string]string](true, items...)
	c2 := namevalue.NewGenericCollectionUsing[string, map[string]string](false, items...)
	c3 := namevalue.NewGenericCollectionUsing[string, map[string]string](true)

	// Act
	actual := args.Map{
		"c1": c1.Length(),
		"c2": c2.Length(),
		"c3": c3.Length(),
	}

	// Assert
	expected := args.Map{
		"c1": 1,
		"c2": 1,
		"c3": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringMapString returns correct value -- CollectionUsing", actual)
}

// ==========================================================================
// Nil Instance IsNull — type-specific
// ==========================================================================

func Test_Instance_IsNull_StringInt(t *testing.T) {
	// Arrange
	var nilInst *namevalue.StringInt
	inst := &namevalue.StringInt{Name: "a", Value: 1}

	// Act
	actual := args.Map{
		"nil": nilInst.IsNull(),
		"nonNil": inst.IsNull(),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"nonNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IsNull StringInt", actual)
}

func Test_Instance_Dispose_NilStringInt(t *testing.T) {
	// Arrange
	var nilInst *namevalue.StringInt
	nilInst.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Instance returns nil -- Dispose nil StringInt", actual)
}
