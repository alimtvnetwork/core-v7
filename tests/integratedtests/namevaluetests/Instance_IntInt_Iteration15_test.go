package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ==========================================================================
// Instance[int, int] — cover generic type path for non-string key types
// ==========================================================================

func Test_Instance_IntInt_String(t *testing.T) {
	// Arrange
	inst := namevalue.Instance[int, int]{Name: 10, Value: 20}
	s := inst.String()
	js := inst.JsonString()

	// Act
	actual := args.Map{
		"str": s != "",
		"json": js != "",
	}

	// Assert
	expected := args.Map{
		"str": true,
		"json": true,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IntInt String/JsonString", actual)
}

func Test_Instance_IntInt_IsNull(t *testing.T) {
	// Arrange
	inst := &namevalue.Instance[int, int]{Name: 1, Value: 2}
	var nilInst *namevalue.Instance[int, int]

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
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IntInt IsNull", actual)
}

func Test_Instance_IntInt_Dispose(t *testing.T) {
	// Arrange
	inst := &namevalue.Instance[int, int]{Name: 5, Value: 10}
	inst.Dispose()

	// Act
	actual := args.Map{
		"name": inst.Name,
		"val": inst.Value,
	}

	// Assert
	expected := args.Map{
		"name": 0,
		"val": 0,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- IntInt Dispose zeros", actual)
}

func Test_Instance_IntInt_Dispose_Nil(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[int, int]
	inst.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Instance returns nil -- IntInt Dispose nil no panic", actual)
}

// ==========================================================================
// Collection[int, int] — full generic type path coverage
// ==========================================================================

func Test_IntInt_Collection_Full(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollection[int, int](5)
	cd := namevalue.NewGenericCollectionDefault[int, int]()
	ce := namevalue.EmptyGenericCollection[int, int]()

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- constructors", actual)
}

func Test_IntInt_CollectionUsing(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[int, int]{{Name: 1, Value: 10}, {Name: 2, Value: 20}}
	c1 := namevalue.NewGenericCollectionUsing[int, int](true, items...)
	c2 := namevalue.NewGenericCollectionUsing[int, int](false, items...)
	c3 := namevalue.NewGenericCollectionUsing[int, int](true)

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- CollectionUsing", actual)
}

func Test_IntInt_AddMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})
	c.Adds(namevalue.Instance[int, int]{Name: 2, Value: 20}, namevalue.Instance[int, int]{Name: 3, Value: 30})
	c.Adds()
	c.Append(namevalue.Instance[int, int]{Name: 4, Value: 40})
	c.Append()
	c.Prepend(namevalue.Instance[int, int]{Name: 0, Value: 0})
	c.Prepend()
	c.AppendIf(true, namevalue.Instance[int, int]{Name: 5, Value: 50})
	c.AppendIf(false, namevalue.Instance[int, int]{Name: 6, Value: 60})
	c.AppendIf(true)
	c.PrependIf(true, namevalue.Instance[int, int]{Name: -1, Value: -10})
	c.PrependIf(false, namevalue.Instance[int, int]{Name: -2, Value: -20})
	c.PrependIf(true)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- add methods", actual)
}

func Test_IntInt_FuncAppendPrepend(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 0, Value: 0})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[int, int] {
		return []namevalue.Instance[int, int]{{Name: -1, Value: -10}}
	})
	c.PrependUsingFuncIf(false, func() []namevalue.Instance[int, int] {
		return []namevalue.Instance[int, int]{{Name: -2, Value: -20}}
	})
	c.PrependUsingFuncIf(true, nil)
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[int, int] {
		return []namevalue.Instance[int, int]{{Name: 1, Value: 10}}
	})
	c.AppendUsingFuncIf(false, func() []namevalue.Instance[int, int] {
		return []namevalue.Instance[int, int]{{Name: 2, Value: 20}}
	})
	c.AppendUsingFuncIf(true, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- func append/prepend", actual)
}

func Test_IntInt_AppendPrependIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 0, Value: 0})
	pre := []namevalue.Instance[int, int]{{Name: -1, Value: -10}}
	app := []namevalue.Instance[int, int]{{Name: 1, Value: 10}}
	c.AppendPrependIf(true, pre, app)
	c.AppendPrependIf(false, pre, app)
	c.AppendPrependIf(true, nil, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- AppendPrependIf", actual)
}

func Test_IntInt_AddsPtr(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	item := namevalue.Instance[int, int]{Name: 1, Value: 10}
	c.AddsPtr(&item, nil)
	c.AddsPtr()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- AddsPtr", actual)
}

func Test_IntInt_AddsIf(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.AddsIf(true, namevalue.Instance[int, int]{Name: 1, Value: 10})
	c.AddsIf(false, namevalue.Instance[int, int]{Name: 2, Value: 20})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- AddsIf", actual)
}

func Test_IntInt_QueryMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})
	c.Add(namevalue.Instance[int, int]{Name: 2, Value: 20})

	var nilC *namevalue.Collection[int, int]

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- query methods", actual)
}

func Test_IntInt_StringMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- string methods", actual)
}

func Test_IntInt_LazyString(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})
	s1 := c.CompiledLazyString()
	s2 := c.CompiledLazyString() // cached path
	c.InvalidateLazyString()

	var nilC *namevalue.Collection[int, int]

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- lazy string", actual)

	nilC.InvalidateLazyString() // no panic
}

func Test_IntInt_IsEqualByString(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[int, int]()
	c1.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})
	c2 := namevalue.NewGenericCollectionDefault[int, int]()
	c2.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})
	c3 := namevalue.NewGenericCollectionDefault[int, int]()
	c3.Add(namevalue.Instance[int, int]{Name: 2, Value: 20})
	var nilC *namevalue.Collection[int, int]

	// Act
	actual := args.Map{
		"equal":    c1.IsEqualByString(c2),
		"notEqual": c1.IsEqualByString(c3),
		"nilBoth":  nilC.IsEqualByString(nil),
		"nilOne":   c1.IsEqualByString(nil),
		"diffLen":  c1.IsEqualByString(namevalue.EmptyGenericCollection[int, int]()),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"nilBoth": true,
		"nilOne": false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- IsEqualByString", actual)
}

func Test_IntInt_ErrorMethods(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	nilErr := c.Error()
	nilMsgErr := c.ErrorUsingMessage("prefix")
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 42})

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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns error -- error methods", actual)
}

func Test_IntInt_ConcatCloneClearDispose(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()
	c.Add(namevalue.Instance[int, int]{Name: 1, Value: 10})

	cn := c.ConcatNew(namevalue.Instance[int, int]{Name: 2, Value: 20})
	item := namevalue.Instance[int, int]{Name: 3, Value: 30}
	cnp := c.ConcatNewPtr(&item)
	cl := c.Clone()
	clp := c.ClonePtr()

	var nilC *namevalue.Collection[int, int]
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
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- concat/clone", actual)

	c.Clear()
	actual2 := args.Map{"afterClear": c.IsEmpty()}
	expected2 := args.Map{"afterClear": true}
	expected2.ShouldBeEqual(t, 0, "IntInt Collection returns correct value -- clear", actual2)

	nilC.Clear()   // no panic
	nilC.Dispose() // no panic
}

func Test_IntInt_CsvStrings_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()

	// Act
	actual := args.Map{"len": len(c.CsvStrings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns empty -- CsvStrings empty", actual)
}

func Test_IntInt_JsonString_Empty(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[int, int]()

	// Act
	actual := args.Map{"val": c.JsonString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "IntInt Collection returns empty -- JsonString empty", actual)
}

// ==========================================================================
// Collection.String() — cached lazyToString return path
// ==========================================================================

func Test_StringAny_String_CachedPath(t *testing.T) {
	// Arrange
	c := namevalue.NewCollection()
	c.Add(namevalue.StringAny{Name: "a", Value: 1})
	// compile the lazy string
	_ = c.CompiledLazyString()
	// now String() should use the cached path (line 397-398)
	s := c.String()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection.String returns correct value -- cached lazy path", actual)
}

// ==========================================================================
// AppendsIf / PrependsIf — int,int type coverage
// ==========================================================================

func Test_AppendsIf_IntInt(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[int, int]{{Name: 1, Value: 10}}
	r1 := namevalue.AppendsIf(true, items, namevalue.Instance[int, int]{Name: 2, Value: 20})
	r2 := namevalue.AppendsIf(false, items, namevalue.Instance[int, int]{Name: 3, Value: 30})
	r3 := namevalue.AppendsIf[int, int](true, items)

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
	expected.ShouldBeEqual(t, 0, "AppendsIf returns correct value -- IntInt", actual)
}

func Test_PrependsIf_IntInt(t *testing.T) {
	// Arrange
	items := []namevalue.Instance[int, int]{{Name: 2, Value: 20}}
	r1 := namevalue.PrependsIf(true, items, namevalue.Instance[int, int]{Name: 1, Value: 10})
	r2 := namevalue.PrependsIf(false, items, namevalue.Instance[int, int]{Name: 3, Value: 30})
	r3 := namevalue.PrependsIf[int, int](true, items)

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
	expected.ShouldBeEqual(t, 0, "PrependsIf returns correct value -- IntInt", actual)
}

// ==========================================================================
// Instance[string, bool] — another type permutation
// ==========================================================================

func Test_Instance_StringBool(t *testing.T) {
	// Arrange
	inst := namevalue.Instance[string, bool]{Name: "flag", Value: true}
	s := inst.String()
	js := inst.JsonString()

	// Act
	actual := args.Map{
		"str": s != "",
		"json": js != "",
	}

	// Assert
	expected := args.Map{
		"str": true,
		"json": true,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringBool", actual)
}

func Test_Instance_StringBool_Dispose(t *testing.T) {
	// Arrange
	inst := &namevalue.Instance[string, bool]{Name: "flag", Value: true}
	inst.Dispose()

	// Act
	actual := args.Map{
		"name": inst.Name,
		"val": inst.Value,
	}

	// Assert
	expected := args.Map{
		"name": "",
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringBool Dispose", actual)
}

// ==========================================================================
// Collection[string, bool] — full generic path
// ==========================================================================

func Test_StringBool_Collection(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, bool]()
	c.Add(namevalue.Instance[string, bool]{Name: "a", Value: true})
	c.Add(namevalue.Instance[string, bool]{Name: "b", Value: false})

	// Act
	actual := args.Map{
		"len":      c.Length(),
		"strings":  len(c.Strings()) > 0,
		"json":     c.JsonString() != "",
		"join":     c.Join(",") != "",
		"csvLen":   len(c.CsvStrings()) > 0,
		"joinCsv":  c.JoinCsv() != "",
		"joinLine": c.JoinCsvLine() != "",
	}

	// Assert
	expected := args.Map{
		"len": 2, "strings": true, "json": true,
		"join": true, "csvLen": true, "joinCsv": true,
		"joinLine": true,
	}
	expected.ShouldBeEqual(t, 0, "StringBool Collection returns correct value -- full", actual)
}

func Test_StringBool_Clone_IsEqual(t *testing.T) {
	// Arrange
	c1 := namevalue.NewGenericCollectionDefault[string, bool]()
	c1.Add(namevalue.Instance[string, bool]{Name: "a", Value: true})
	c2 := c1.Clone()

	// Act
	actual := args.Map{
		"eq": c1.IsEqualByString(c2),
		"len": c2.Length(),
	}

	// Assert
	expected := args.Map{
		"eq": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringBool Collection returns correct value -- clone equal", actual)
}

func Test_StringBool_Error(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, bool]()
	c.Add(namevalue.Instance[string, bool]{Name: "err", Value: true})

	// Act
	actual := args.Map{
		"hasErr": c.Error() != nil,
		"hasMsg": c.ErrorUsingMessage("p") != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"hasMsg": true,
	}
	expected.ShouldBeEqual(t, 0, "StringBool Collection returns error -- error methods", actual)
}

func Test_StringBool_Dispose(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, bool]()
	c.Add(namevalue.Instance[string, bool]{Name: "a", Value: true})
	c.Dispose()

	// Act
	actual := args.Map{"nil": c.Items == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "StringBool Collection returns nil -- dispose", actual)
}

// ==========================================================================
// Instance[string, float64] — float type permutation
// ==========================================================================

func Test_Instance_StringFloat64(t *testing.T) {
	// Arrange
	inst := namevalue.Instance[string, float64]{Name: "pi", Value: 3.14}
	s := inst.String()
	js := inst.JsonString()

	// Act
	actual := args.Map{
		"str": s != "",
		"json": js != "",
	}

	// Assert
	expected := args.Map{
		"str": true,
		"json": true,
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringFloat64", actual)
}

func Test_StringFloat64_Collection(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, float64]()
	c.Add(namevalue.Instance[string, float64]{Name: "pi", Value: 3.14})
	c.Add(namevalue.Instance[string, float64]{Name: "e", Value: 2.71})

	// Act
	actual := args.Map{
		"len": c.Length(), "join": c.Join(",") != "",
		"json": c.JsonString() != "", "string": c.String() != "",
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"join": true,
		"json": true,
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "StringFloat64 Collection returns correct value -- full", actual)
}

// ==========================================================================
// Instance[string, []string] — slice value type permutation
// ==========================================================================

func Test_Instance_StringSlice(t *testing.T) {
	// Arrange
	inst := namevalue.Instance[string, []string]{Name: "tags", Value: []string{"a", "b"}}
	s := inst.String()
	js := inst.JsonString()
	inst.Dispose()

	// Act
	actual := args.Map{
		"str": s != "",
		"json": js != "",
		"name": inst.Name,
	}

	// Assert
	expected := args.Map{
		"str": true,
		"json": true,
		"name": "",
	}
	expected.ShouldBeEqual(t, 0, "Instance returns correct value -- StringSlice", actual)
}

func Test_StringSlice_Collection(t *testing.T) {
	// Arrange
	c := namevalue.NewGenericCollectionDefault[string, []string]()
	c.Add(namevalue.Instance[string, []string]{Name: "tags", Value: []string{"x"}})

	// Act
	actual := args.Map{
		"len": c.Length(), "string": c.String() != "",
		"json": c.JsonString() != "",
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"string": true,
		"json": true,
	}
	expected.ShouldBeEqual(t, 0, "StringSlice Collection returns correct value -- full", actual)

	cl := c.Clone()
	actual2 := args.Map{"cloneLen": cl.Length()}
	expected2 := args.Map{"cloneLen": 1}
	expected2.ShouldBeEqual(t, 0, "StringSlice Collection returns correct value -- clone", actual2)
}
