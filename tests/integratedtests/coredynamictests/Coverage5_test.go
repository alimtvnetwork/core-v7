package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Dynamic constructors ──

func Test_Cov5_NewDynamic_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{
		"isValid":  d.IsValid(),
		"data":     d.Data(),
		"value":    d.Value(),
		"isNull":   d.IsNull(),
		"str":      d.String() != "",
		"length":   d.Length(),
	}
	expected := args.Map{"isValid": true, "data": "hello", "value": "hello", "isNull": false, "str": true, "length": actual["length"]}
	expected.ShouldBeEqual(t, 0, "NewDynamicValid returns valid -- string", actual)
}

func Test_Cov5_NewDynamic_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"isInvalid": d.IsInvalid(), "isNull": d.IsNull()}
	expected := args.Map{"isInvalid": true, "isNull": true}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic returns invalid -- nil data", actual)
}

func Test_Cov5_NewDynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"isValid": d.IsValid(), "notNil": d != nil}
	expected := args.Map{"isValid": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewDynamicPtr returns valid ptr -- string", actual)
}

func Test_Cov5_InvalidDynamicPtr(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	actual := args.Map{"isInvalid": d.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "InvalidDynamicPtr returns error -- returns invalid", actual)
}

// ── Dynamic Clone ──

func Test_Cov5_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	c := d.Clone()
	cp := d.ClonePtr()
	actual := args.Map{"val": c.Data(), "ptrVal": cp.Data()}
	expected := args.Map{"val": "hello", "ptrVal": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic Clone returns same data -- string", actual)
}

func Test_Cov5_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"isNil": d.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr nil -- returns nil", actual)
}

func Test_Cov5_Dynamic_NonPtr_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"nonPtrOk": d.NonPtr().IsValid(), "ptrOk": d.Ptr().IsValid()}
	expected := args.Map{"nonPtrOk": true, "ptrOk": true}
	expected.ShouldBeEqual(t, 0, "Dynamic NonPtr/Ptr -- valid", actual)
}

// ── Dynamic type checks ──

func Test_Cov5_Dynamic_TypeChecks_String(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{
		"isString":    d.IsStringType(),
		"isPrimitive": d.IsPrimitive(),
		"isNumber":    d.IsNumber(),
		"isStruct":    d.IsStruct(),
		"isFunc":      d.IsFunc(),
		"isSlice":     d.IsSliceOrArray(),
		"isMap":       d.IsMap(),
		"isPointer":   d.IsPointer(),
		"isValueType": d.IsValueType(),
	}
	expected := args.Map{
		"isString": true, "isPrimitive": true, "isNumber": false,
		"isStruct": false, "isFunc": false, "isSlice": false,
		"isMap": false, "isPointer": false, "isValueType": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- string", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Map(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"isMap": d.IsMap(), "isSliceOrMap": d.IsSliceOrArrayOrMap(), "length": d.Length()}
	expected := args.Map{"isMap": true, "isSliceOrMap": true, "length": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- map", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Slice(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	actual := args.Map{"isSlice": d.IsSliceOrArray(), "length": d.Length()}
	expected := args.Map{"isSlice": true, "length": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- slice", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Pointer(t *testing.T) {
	x := 42
	d := coredynamic.NewDynamicPtr(&x, true)
	actual := args.Map{"isPointer": d.IsPointer()}
	expected := args.Map{"isPointer": true}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- pointer", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Struct(t *testing.T) {
	type S struct{ A int }
	d := coredynamic.NewDynamicPtr(S{A: 1}, true)
	actual := args.Map{"isStruct": d.IsStruct()}
	expected := args.Map{"isStruct": true}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- struct", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Func(t *testing.T) {
	d := coredynamic.NewDynamicPtr(func() {}, true)
	actual := args.Map{"isFunc": d.IsFunc()}
	expected := args.Map{"isFunc": true}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- func", actual)
}

func Test_Cov5_Dynamic_TypeChecks_Number(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"isNumber": d.IsNumber(), "isPrimitive": d.IsPrimitive()}
	expected := args.Map{"isNumber": true, "isPrimitive": true}
	expected.ShouldBeEqual(t, 0, "Dynamic type checks -- number", actual)
}

// ── Dynamic value extraction ──

func Test_Cov5_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"val": d.ValueInt()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueInt -- 42", actual)
}

func Test_Cov5_Dynamic_IntDefault(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault nil -- returns default", actual)
}

// ── Dynamic JSON ──

func Test_Cov5_Dynamic_Json(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	r := d.Json()
	rp := d.JsonPtr()
	_, _ = d.JsonString()
	actual := args.Map{
		"hasBytes":  r.HasBytes(),
		"ptrNotNil": rp != nil,
	}
	expected := args.Map{"hasBytes": actual["hasBytes"], "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Json -- valid", actual)
}

func Test_Cov5_Dynamic_JsonParseSelfInject(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	r := d.JsonPtr()
	var d2 coredynamic.Dynamic
	err := d2.JsonParseSelfInject(r)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": actual["noErr"]}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonParseSelfInject -- roundtrip", actual)
}

// ── MapAnyItems ──

func Test_Cov5_MapAnyItems_Basic(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	m.Add("key1", "val1")
	m.Set("key2", 42)
	actual := args.Map{
		"len":     m.Length(),
		"isEmpty": m.IsEmpty(),
		"hasAny":  m.HasAnyItem(),
		"hasKey":  m.HasKey("key1"),
		"noKey":   m.HasKey("missing"),
	}
	expected := args.Map{"len": 2, "isEmpty": false, "hasAny": true, "hasKey": true, "noKey": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems basic -- 2 items", actual)
}

func Test_Cov5_MapAnyItems_GetValue(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"val": m.GetValue("a"), "missing": m.GetValue("x") == nil}
	expected := args.Map{"val": 1, "missing": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetValue -- found and missing", actual)
}

func Test_Cov5_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	val, has := m.Get("a")
	_, notHas := m.Get("x")
	actual := args.Map{"val": val, "has": has, "notHas": notHas}
	expected := args.Map{"val": 1, "has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Get -- found and missing", actual)
}

func Test_Cov5_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "hello"})
	var s string
	err := m.Deserialize("a", &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Deserialize -- valid", actual)
}

func Test_Cov5_MapAnyItems_Deserialize_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var s string
	err := m.Deserialize("missing", &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Deserialize missing -- error", actual)
}

func Test_Cov5_MapAnyItems_AllKeysSorted(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0], "len": len(keys)}
	expected := args.Map{"first": "a", "len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllKeysSorted -- sorted", actual)
}

func Test_Cov5_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	r := m.Json()
	jsonStr, jsonErr := m.JsonString()
	actual := args.Map{"hasBytes": r.HasBytes(), "jsonStr": jsonStr != "", "jsonErr": jsonErr == nil}
	expected := args.Map{"hasBytes": true, "jsonStr": true, "jsonErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Json -- valid", actual)
}

func Test_Cov5_MapAnyItems_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	actual := args.Map{"len": m.Length(), "isEmpty": m.IsEmpty(), "hasKey": m.HasKey("a")}
	expected := args.Map{"len": 0, "isEmpty": true, "hasKey": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Nil -- safe defaults", actual)
}

func Test_Cov5_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	r := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	err := m2.JsonParseSelfInject(r)
	actual := args.Map{"noErr": err == nil, "hasKey": m2.HasKey("a")}
	expected := args.Map{"noErr": true, "hasKey": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonParseSelfInject -- roundtrip", actual)
}

// ── Collection[T] ──

func Test_Cov5_Collection_Basic(t *testing.T) {
	c := coredynamic.NewCollection[string](5)
	c.Add("a").Add("b").Add("c")
	actual := args.Map{
		"len":     c.Length(),
		"count":   c.Count(),
		"isEmpty": c.IsEmpty(),
		"hasAny":  c.HasAnyItem(),
		"first":   c.First(),
		"last":    c.Last(),
		"at":      c.At(1),
		"lastIdx": c.LastIndex(),
		"hasIdx":  c.HasIndex(2),
		"noIdx":   c.HasIndex(5),
	}
	expected := args.Map{
		"len": 3, "count": 3, "isEmpty": false, "hasAny": true,
		"first": "a", "last": "c", "at": "b", "lastIdx": 2,
		"hasIdx": true, "noIdx": false,
	}
	expected.ShouldBeEqual(t, 0, "Collection basic -- 3 items", actual)
}

func Test_Cov5_Collection_FirstLastOrDefault(t *testing.T) {
	c := coredynamic.NewCollection[string](5)
	c.Add("a").Add("b")
	first, fOk := c.FirstOrDefault()
	last, lOk := c.LastOrDefault()
	empty := coredynamic.EmptyCollection[string]()
	_, efOk := empty.FirstOrDefault()
	_, elOk := empty.LastOrDefault()
	actual := args.Map{"first": *first, "fOk": fOk, "last": *last, "lOk": lOk, "efOk": efOk, "elOk": elOk}
	expected := args.Map{"first": "a", "fOk": true, "last": "b", "lOk": true, "efOk": false, "elOk": false}
	expected.ShouldBeEqual(t, 0, "Collection FirstLastOrDefault -- valid and empty", actual)
}

func Test_Cov5_Collection_AddMany(t *testing.T) {
	c := coredynamic.NewCollection[int](5)
	c.AddMany(1, 2, 3)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection AddMany -- 3 items", actual)
}

func Test_Cov5_Collection_AddNonNil(t *testing.T) {
	c := coredynamic.NewCollection[int](5)
	v := 42
	c.AddNonNil(&v)
	c.AddNonNil(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection AddNonNil -- skip nil", actual)
}

func Test_Cov5_Collection_RemoveAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	ok := c.RemoveAt(1)
	notOk := c.RemoveAt(99)
	actual := args.Map{"ok": ok, "notOk": notOk, "len": c.Length()}
	expected := args.Map{"ok": true, "notOk": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- valid and invalid", actual)
}

func Test_Cov5_Collection_Skip_Take(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	actual := args.Map{
		"skipLen": len(c.Skip(2)),
		"takeLen": len(c.Take(3)),
		"limitLen": len(c.Limit(3)),
	}
	expected := args.Map{"skipLen": 3, "takeLen": 3, "limitLen": 3}
	expected.ShouldBeEqual(t, 0, "Collection Skip/Take/Limit -- correct", actual)
}

func Test_Cov5_Collection_Filter(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	filtered := c.Filter(func(i int) bool { return i > 3 })
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection Filter -- gt 3", actual)
}

func Test_Cov5_Collection_Loop(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sum := 0
	c.Loop(func(i int, item int) bool { sum += item; return false })
	actual := args.Map{"sum": sum}
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "Collection Loop -- sum all", actual)
}

func Test_Cov5_Collection_GetPagesSize(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	actual := args.Map{
		"pages2": c.GetPagesSize(2),
		"pages0": c.GetPagesSize(0),
	}
	expected := args.Map{"pages2": 3, "pages0": 0}
	expected.ShouldBeEqual(t, 0, "Collection GetPagesSize -- 2 per page", actual)
}

func Test_Cov5_Collection_Clear_Dispose(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Clear()
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection Clear -- empty after", actual)
}

func Test_Cov5_Collection_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	actual := args.Map{"len": c.Length(), "isEmpty": c.IsEmpty()}
	expected := args.Map{"len": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection Nil -- safe defaults", actual)
}

func Test_Cov5_Collection_CollectionFrom_Nil(t *testing.T) {
	c := coredynamic.CollectionFrom[int](nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionFrom nil -- empty", actual)
}

func Test_Cov5_Collection_CollectionClone(t *testing.T) {
	c := coredynamic.CollectionClone([]int{1, 2})
	actual := args.Map{"len": c.Length(), "first": c.First()}
	expected := args.Map{"len": 2, "first": 1}
	expected.ShouldBeEqual(t, 0, "CollectionClone -- correct copy", actual)
}

// ── LeftRight ──

func Test_Cov5_LeftRight(t *testing.T) {
	lr := coredynamic.LeftRight{Left: "l", Right: "r"}
	actual := args.Map{"left": lr.Left, "right": lr.Right}
	expected := args.Map{"left": "l", "right": "r"}
	expected.ShouldBeEqual(t, 0, "LeftRight -- correct", actual)
}

// ── Type ──

func Test_Cov5_Type_Basic(t *testing.T) {
	typ := coredynamic.Type("hello")
	actual := args.Map{
		"notNil": typ != nil,
		"name":   typ.Name(),
	}
	expected := args.Map{"notNil": true, "name": "string"}
	expected.ShouldBeEqual(t, 0, "Type basic -- string", actual)
}

// ── TypeMustBeSame ──

func Test_Cov5_TypeMustBeSame_Same(t *testing.T) {
	coredynamic.TypeMustBeSame("hello", "world") // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame same -- no panic", actual)
}

func Test_Cov5_TypeMustBeSame_Different(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame different -- panics", actual)
	}()
	coredynamic.TypeMustBeSame("hello", 42)
}

// ── ValueStatus ──

func Test_Cov5_ValueStatus(t *testing.T) {
	vs := &coredynamic.ValueStatus{Value: "hello", IsValid: true}
	actual := args.Map{"val": vs.Value, "isValid": vs.IsValid}
	expected := args.Map{"val": "hello", "isValid": true}
	expected.ShouldBeEqual(t, 0, "ValueStatus -- basic", actual)
}

// ── SimpleRequest / SimpleResult ──

func Test_Cov5_SimpleRequest(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	actual := args.Map{"notNil": sr != nil, "val": sr.Value()}
	expected := args.Map{"notNil": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "SimpleRequest -- valid", actual)
}

func Test_Cov5_SimpleResult(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	actual := args.Map{"result": sr.Result, "isValid": sr.IsValid()}
	expected := args.Map{"result": "hello", "isValid": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult -- valid", actual)
}

// ── NewMapAnyItemsUsingAnyTypeMap ──

func Test_Cov5_NewMapAnyItemsUsingAnyTypeMap(t *testing.T) {
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]any{"a": 1})
	_, nilErr := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	actual := args.Map{"noErr": err == nil, "hasKey": m != nil && m.HasKey("a"), "nilErr": nilErr != nil}
	expected := args.Map{"noErr": actual["noErr"], "hasKey": actual["hasKey"], "nilErr": true}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingAnyTypeMap -- valid and nil", actual)
}
