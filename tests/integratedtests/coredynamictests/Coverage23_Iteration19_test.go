package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItems — Deserialize, GetItemRef, Paging, Diff, Clone, JSON model
// ══════════════════════════════════════════════════════════════════════════════

func Test_I19_MapAnyItems_Deserialize_Success(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "hello"})
	var target string
	err := m.Deserialize("k", &target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize", actual)
}

func Test_I19_MapAnyItems_Deserialize_MissingKey(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.Deserialize("missing", &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize missing key", actual)
}

func Test_I19_MapAnyItems_DeserializeMust(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": 42})
	var target int
	m.DeserializeMust("k", &target)
	actual := args.Map{"val": target}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DeserializeMust", actual)
}

func Test_I19_MapAnyItems_GetUsingUnmarshallManyAt_Success(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "x", "b": "y"})
	var a, b string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &a},
		corejson.KeyAny{Key: "b", AnyInf: &b},
	)
	actual := args.Map{"noErr": err == nil, "a": a, "b": b}
	expected := args.Map{"noErr": true, "a": "x", "b": "y"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetUsingUnmarshallManyAt", actual)
}

func Test_I19_MapAnyItems_GetUsingUnmarshallManyAt_Error(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var a string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "missing", AnyInf: &a},
	)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns error -- GetUsingUnmarshallManyAt error", actual)
}

func Test_I19_MapAnyItems_GetItemRef_Success(t *testing.T) {
	val := "hello"
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": &val})
	var target string
	err := m.GetItemRef("k", &target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef", actual)
}

func Test_I19_MapAnyItems_GetItemRef_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.GetItemRef("missing", &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef missing", actual)
}

func Test_I19_MapAnyItems_GetItemRef_NilRef(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	err := m.GetItemRef("k", nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- GetItemRef nil ref", actual)
}

func Test_I19_MapAnyItems_GetItemRef_NotPointer(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	var target string
	err := m.GetItemRef("k", target) // not pointer
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef not pointer", actual)
}

func Test_I19_MapAnyItems_GetManyItemsRefs_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	err := m.GetManyItemsRefs()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- GetManyItemsRefs empty", actual)
}

func Test_I19_MapAnyItems_GetFieldsMap_Found(t *testing.T) {
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})
	fm, err, found := m.GetFieldsMap("k")
	actual := args.Map{"found": found, "noErr": err == nil, "notNil": fm != nil}
	expected := args.Map{"found": true, "noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap", actual)
}

func Test_I19_MapAnyItems_GetFieldsMap_NotFound(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	_, _, found := m.GetFieldsMap("missing")
	actual := args.Map{"found": found}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap not found", actual)
}

func Test_I19_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})
	fm, found := m.GetSafeFieldsMap("k")
	actual := args.Map{"found": found, "notNil": fm != nil}
	expected := args.Map{"found": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetSafeFieldsMap", actual)
}

func Test_I19_MapAnyItems_AddKeyAny(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: "v"})
	actual := args.Map{"isNew": isNew, "len": m.Length()}
	expected := args.Map{"isNew": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddKeyAny", actual)
}

func Test_I19_MapAnyItems_AddKeyAnyWithValidation_Match(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(""),
		corejson.KeyAny{Key: "k", AnyInf: "v"},
	)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddKeyAnyWithValidation match", actual)
}

func Test_I19_MapAnyItems_AddKeyAnyWithValidation_Mismatch(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(0),
		corejson.KeyAny{Key: "k", AnyInf: "v"},
	)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddKeyAnyWithValidation mismatch", actual)
}

func Test_I19_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr("val")
	m.AddJsonResultPtr("k", jr)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddJsonResultPtr", actual)
}

func Test_I19_MapAnyItems_AddJsonResultPtr_Nil(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	m.AddJsonResultPtr("k", nil)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- AddJsonResultPtr nil", actual)
}

func Test_I19_MapAnyItems_AddMapResultOption_Override(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(true, map[string]any{"a": 2, "b": 3})
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns error -- AddMapResultOption override", actual)
}

func Test_I19_MapAnyItems_AddMapResultOption_NoOverride(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(false, map[string]any{"a": 2})
	v, _ := m.Get("a")
	actual := args.Map{"val": v}
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- AddMapResultOption no override", actual)
}

func Test_I19_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddManyMapResultsUsingOption", actual)
}

func Test_I19_MapAnyItems_AddManyMapResultsUsingOption_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	m.AddManyMapResultsUsingOption(true)
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- AddManyMapResultsUsingOption empty", actual)
}

func Test_I19_MapAnyItems_ReflectSetTo_Success(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "hello"})
	var target string
	err := m.ReflectSetTo("k", &target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ReflectSetTo", actual)
}

func Test_I19_MapAnyItems_ReflectSetTo_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.ReflectSetTo("missing", &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ReflectSetTo missing", actual)
}

func Test_I19_MapAnyItems_GetPagedCollection(t *testing.T) {
	items := map[string]any{}
	for i := 0; i < 5; i++ {
		items[string(rune('a'+i))] = i
	}
	m := coredynamic.NewMapAnyItemsUsingItems(items)
	pages := m.GetPagedCollection(2)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagedCollection", actual)
}

func Test_I19_MapAnyItems_GetPagedCollection_SmallPage(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	pages := m.GetPagedCollection(10)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagedCollection small", actual)
}

func Test_I19_MapAnyItems_JsonResultOfKey_Found(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	jr := m.JsonResultOfKey("k")
	actual := args.Map{"notNil": jr != nil, "hasErr": jr.HasError()}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey found", actual)
}

func Test_I19_MapAnyItems_JsonResultOfKey_Missing(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	jr := m.JsonResultOfKey("missing")
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey missing", actual)
}

func Test_I19_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	mr := m.JsonResultOfKeys("a", "b")
	actual := args.Map{"notNil": mr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKeys", actual)
}

func Test_I19_MapAnyItems_JsonResultOfKeys_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	mr := m.JsonResultOfKeys()
	actual := args.Map{"notNil": mr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonResultOfKeys empty", actual)
}

func Test_I19_MapAnyItems_JsonMapResults(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	mr, err := m.JsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonMapResults", actual)
}

func Test_I19_MapAnyItems_JsonMapResults_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	mr, err := m.JsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonMapResults empty", actual)
}

func Test_I19_MapAnyItems_JsonResultsCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsCollection", actual)
}

func Test_I19_MapAnyItems_JsonResultsCollection_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	rc := m.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonResultsCollection empty", actual)
}

func Test_I19_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_I19_MapAnyItems_JsonModel(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	model := m.JsonModel()
	actual := args.Map{"notNil": model != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonModel", actual)
}

func Test_I19_MapAnyItems_JsonModelAny(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"notNil": m.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonModelAny", actual)
}

func Test_I19_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_I19_MapAnyItems_JsonPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonPtr", actual)
}

func Test_I19_MapAnyItems_ParseInjectUsingJson(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr(map[string]any{"a": 1})
	result, err := m.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ParseInjectUsingJson", actual)
}

func Test_I19_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr(map[string]any{"x": 2})
	err := m.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonParseSelfInject", actual)
}

func Test_I19_MapAnyItems_Strings(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()
	actual := args.Map{"notEmpty": len(strs) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Strings", actual)
}

func Test_I19_MapAnyItems_String(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"notEmpty": m.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- String", actual)
}

func Test_I19_MapAnyItems_DeepClear(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.DeepClear()
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DeepClear", actual)
}

func Test_I19_MapAnyItems_Dispose(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.Dispose()
	actual := args.Map{"nil": m.Items == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Dispose", actual)
}

func Test_I19_MapAnyItems_IsEqualRaw_Equal(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 1})}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw equal", actual)
}

func Test_I19_MapAnyItems_IsEqualRaw_NotEqual(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 2})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw not equal", actual)
}

func Test_I19_MapAnyItems_IsEqualRaw_DiffLen(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 1, "b": 2})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw diff len", actual)
}

func Test_I19_MapAnyItems_IsEqualRaw_MissingKey(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"b": 1})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw missing key", actual)
}

func Test_I19_MapAnyItems_IsEqual(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	actual := args.Map{"eq": m1.IsEqual(m2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqual", actual)
}

func Test_I19_MapAnyItems_IsEqual_BothNil(t *testing.T) {
	var m1, m2 *coredynamic.MapAnyItems
	actual := args.Map{"eq": m1.IsEqual(m2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- IsEqual both nil", actual)
}

func Test_I19_MapAnyItems_IsEqual_OneNil(t *testing.T) {
	m1 := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"eq": m1.IsEqual(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- IsEqual one nil", actual)
}
	var m *coredynamic.MapAnyItems
	diff := m.RawMapStringAnyDiff()
	actual := args.Map{"notNil": diff != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- RawMapStringAnyDiff nil", actual)
}

func Test_I19_MapAnyItems_MapAnyItemsSelf(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"same": m.MapAnyItems() == m}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- MapAnyItems self", actual)
}

func Test_I19_MapAnyItems_NewUsingAnyTypeMap_Success(t *testing.T) {
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]int{"a": 1})
	actual := args.Map{"noErr": err == nil, "len": m.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- NewUsingAnyTypeMap", actual)
}

func Test_I19_MapAnyItems_NewUsingAnyTypeMap_Nil(t *testing.T) {
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- NewUsingAnyTypeMap nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionMethods — AddIf, AddCollection, ConcatNew, Clone, Capacity, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_I19_CollectionMethods_AddIf_True(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c.AddIf(true, "x")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns non-empty -- AddIf true", actual)
}

func Test_I19_CollectionMethods_AddIf_False(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c.AddIf(false, "x")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns non-empty -- AddIf false", actual)
}

func Test_I19_CollectionMethods_AddManyIf(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c.AddManyIf(true, "a", "b")
	c.AddManyIf(false, "c")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddManyIf", actual)
}

func Test_I19_CollectionMethods_AddCollection(t *testing.T) {
	c1 := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := coredynamic.New.Collection.String.From([]string{"b", "c"})
	c1.AddCollection(c2)
	actual := args.Map{"len": c1.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCollection", actual)
}

func Test_I19_CollectionMethods_AddCollections(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c1 := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := coredynamic.New.Collection.String.From([]string{"b"})
	c.AddCollections(c1, nil, c2)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCollections", actual)
}

func Test_I19_CollectionMethods_ConcatNew(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := c.ConcatNew("b", "c")
	actual := args.Map{"origLen": c.Length(), "newLen": c2.Length()}
	expected := args.Map{"origLen": 1, "newLen": 3}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- ConcatNew", actual)
}

func Test_I19_CollectionMethods_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	cloned := c.Clone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Clone", actual)
}

func Test_I19_CollectionMethods_Clone_Nil(t *testing.T) {
	var c *coredynamic.Collection[string]
	cloned := c.Clone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns nil -- Clone nil", actual)
}

func Test_I19_CollectionMethods_Capacity(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(10)
	actual := args.Map{"cap": c.Capacity() >= 10}
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Capacity", actual)
}

func Test_I19_CollectionMethods_AddCapacity(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c.AddCapacity(5)
	actual := args.Map{"cap": c.Capacity() >= 5}
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCapacity", actual)
}

func Test_I19_CollectionMethods_Resize(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	c.Resize(20)
	actual := args.Map{"cap": c.Capacity() >= 20}
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Resize", actual)
}

func Test_I19_CollectionMethods_Reverse(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})
	c.Reverse()
	actual := args.Map{"first": c.First(), "last": c.Last()}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Reverse", actual)
}

func Test_I19_CollectionMethods_InsertAt(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "c"})
	c.InsertAt(1, "b")
	items := c.Items()
	actual := args.Map{"len": len(items), "mid": items[1]}
	expected := args.Map{"len": 3, "mid": "b"}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- InsertAt", actual)
}

func Test_I19_CollectionMethods_IndexOfFunc(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})
	idx := c.IndexOfFunc(func(s string) bool { return s == "b" })
	actual := args.Map{"idx": idx}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- IndexOfFunc", actual)
}

func Test_I19_CollectionMethods_IndexOfFunc_NotFound(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a"})
	idx := c.IndexOfFunc(func(s string) bool { return s == "z" })
	actual := args.Map{"idx": idx}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- IndexOfFunc not found", actual)
}

func Test_I19_CollectionMethods_ContainsFunc(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"has": c.ContainsFunc(func(s string) bool { return s == "a" })}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- ContainsFunc", actual)
}

func Test_I19_CollectionMethods_SafeAt(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"valid": c.SafeAt(0), "invalid": c.SafeAt(99)}
	expected := args.Map{"valid": "a", "invalid": ""}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- SafeAt", actual)
}

func Test_I19_CollectionMethods_SprintItems(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	strs := c.SprintItems("[%s]")
	actual := args.Map{"first": strs[0]}
	expected := args.Map{"first": "[a]"}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- SprintItems", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — additional paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I19_ReflectSetFromTo_BothNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(nil, nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- both nil", actual)
}

func Test_I19_ReflectSetFromTo_SamePointerType(t *testing.T) {
	x := "hello"
	src := &x
	var dst string
	err := coredynamic.ReflectSetFromTo(src, &dst)
	actual := args.Map{"noErr": err == nil, "val": dst}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same pointer type", actual)
}

func Test_I19_ReflectSetFromTo_NonPointerToPointer(t *testing.T) {
	var dst int
	err := coredynamic.ReflectSetFromTo(42, &dst)
	actual := args.Map{"noErr": err == nil, "val": dst}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns non-empty -- non-ptr to ptr", actual)
}

func Test_I19_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	type s struct{ X int }
	var dst s
	err := coredynamic.ReflectSetFromTo([]byte(`{"X":5}`), &dst)
	actual := args.Map{"noErr": err == nil, "val": dst.X}
	expected := args.Map{"noErr": true, "val": 5}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bytes to struct", actual)
}

func Test_I19_ReflectSetFromTo_StructToBytes(t *testing.T) {
	type s struct{ X int }
	src := s{X: 7}
	var dst []byte
	err := coredynamic.ReflectSetFromTo(src, &dst)
	actual := args.Map{"noErr": err == nil, "hasBytes": len(dst) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- struct to bytes", actual)
}

func Test_I19_ReflectSetFromTo_DestNotPointer(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", "not a pointer")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- dest not pointer", actual)
}

func Test_I19_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	var dst int
	err := coredynamic.ReflectSetFromTo("hello", &dst)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- type mismatch", actual)
}

func Test_I19_ReflectSetFromTo_DestNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- dest nil", actual)
}
