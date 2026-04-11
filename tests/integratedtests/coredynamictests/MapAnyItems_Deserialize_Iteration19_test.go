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

func Test_MapAnyItems_Deserialize_Success(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "hello"})
	var target string
	err := m.Deserialize("k", &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize", actual)
}

func Test_MapAnyItems_Deserialize_MissingKey(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.Deserialize("missing", &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Deserialize missing key", actual)
}

func Test_MapAnyItems_DeserializeMust(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": 42})
	var target int
	m.DeserializeMust("k", &target)

	// Act
	actual := args.Map{"val": target}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DeserializeMust", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallManyAt_Success(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "x", "b": "y"})
	var a, b string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &a},
		corejson.KeyAny{Key: "b", AnyInf: &b},
	)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"a": a,
		"b": b,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"a": "x",
		"b": "y",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetUsingUnmarshallManyAt", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallManyAt_Error(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var a string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "missing", AnyInf: &a},
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns error -- GetUsingUnmarshallManyAt error", actual)
}

func Test_MapAnyItems_GetItemRef_Success(t *testing.T) {
	// Arrange
	val := "hello"
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": &val})
	var target string
	err := m.GetItemRef("k", &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef", actual)
}

func Test_MapAnyItems_GetItemRef_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.GetItemRef("missing", &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef missing", actual)
}

func Test_MapAnyItems_GetItemRef_NilRef(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	err := m.GetItemRef("k", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- GetItemRef nil ref", actual)
}

func Test_MapAnyItems_GetItemRef_NotPointer(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	var target string
	err := m.GetItemRef("k", target) // not pointer

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetItemRef not pointer", actual)
}

func Test_MapAnyItems_GetManyItemsRefs_Empty_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.GetManyItemsRefs()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- GetManyItemsRefs empty", actual)
}

func Test_MapAnyItems_GetFieldsMap_Found(t *testing.T) {
	// Arrange
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})
	fm, err, found := m.GetFieldsMap("k")

	// Act
	actual := args.Map{
		"found": found,
		"noErr": err == nil,
		"notNil": fm != nil,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap", actual)
}

func Test_MapAnyItems_GetFieldsMap_NotFound(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	_, _, found := m.GetFieldsMap("missing")

	// Act
	actual := args.Map{"found": found}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetFieldsMap not found", actual)
}

func Test_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	// Arrange
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})
	fm, found := m.GetSafeFieldsMap("k")

	// Act
	actual := args.Map{
		"found": found,
		"notNil": fm != nil,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetSafeFieldsMap", actual)
}

func Test_MapAnyItems_AddKeyAny(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{
		"isNew": isNew,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"isNew": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddKeyAny", actual)
}

func Test_MapAnyItems_AddKeyAnyWithValidation_Match(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(""),
		corejson.KeyAny{Key: "k", AnyInf: "v"},
	)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddKeyAnyWithValidation match", actual)
}

func Test_MapAnyItems_AddKeyAnyWithValidation_Mismatch(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(0),
		corejson.KeyAny{Key: "k", AnyInf: "v"},
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddKeyAnyWithValidation mismatch", actual)
}

func Test_MapAnyItems_AddJsonResultPtr_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr("val")
	m.AddJsonResultPtr("k", jr)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddJsonResultPtr", actual)
}

func Test_MapAnyItems_AddJsonResultPtr_Nil(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.AddJsonResultPtr("k", nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- AddJsonResultPtr nil", actual)
}

func Test_MapAnyItems_AddMapResultOption_Override(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(true, map[string]any{"a": 2, "b": 3})

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns error -- AddMapResultOption override", actual)
}

func Test_MapAnyItems_AddMapResultOption_NoOverride(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.AddMapResultOption(false, map[string]any{"a": 2})
	v, _ := m.Get("a")

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- AddMapResultOption no override", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.AddManyMapResultsUsingOption(true, map[string]any{"a": 1}, map[string]any{"b": 2})

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddManyMapResultsUsingOption", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption_Empty_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.AddManyMapResultsUsingOption(true)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- AddManyMapResultsUsingOption empty", actual)
}

func Test_MapAnyItems_ReflectSetTo_Success(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "hello"})
	var target string
	err := m.ReflectSetTo("k", &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ReflectSetTo", actual)
}

func Test_MapAnyItems_ReflectSetTo_Missing_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var target string
	err := m.ReflectSetTo("missing", &target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ReflectSetTo missing", actual)
}

func Test_MapAnyItems_GetPagedCollection(t *testing.T) {
	// Arrange
	items := map[string]any{}
	for i := 0; i < 5; i++ {
		items[string(rune('a'+i))] = i
	}
	m := coredynamic.NewMapAnyItemsUsingItems(items)
	pages := m.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagedCollection", actual)
}

func Test_MapAnyItems_GetPagedCollection_SmallPage(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	pages := m.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagedCollection small", actual)
}

func Test_MapAnyItems_JsonResultOfKey_Found(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	jr := m.JsonResultOfKey("k")

	// Act
	actual := args.Map{
		"notNil": jr != nil,
		"hasErr": jr.HasError(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey found", actual)
}

func Test_MapAnyItems_JsonResultOfKey_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := m.JsonResultOfKey("missing")

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKey missing", actual)
}

func Test_MapAnyItems_JsonResultOfKeys_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	mr := m.JsonResultOfKeys("a", "b")

	// Act
	actual := args.Map{"notNil": mr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultOfKeys", actual)
}

func Test_MapAnyItems_JsonResultOfKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	mr := m.JsonResultOfKeys()

	// Act
	actual := args.Map{"notNil": mr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonResultOfKeys empty", actual)
}

func Test_MapAnyItems_JsonMapResults_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	mr, err := m.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": mr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonMapResults", actual)
}

func Test_MapAnyItems_JsonMapResults_Empty_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	mr, err := m.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": mr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonMapResults empty", actual)
}

func Test_MapAnyItems_JsonResultsCollection_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsCollection", actual)
}

func Test_MapAnyItems_JsonResultsCollection_Empty_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	rc := m.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- JsonResultsCollection empty", actual)
}

func Test_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	rc := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_MapAnyItems_JsonModel(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	model := m.JsonModel()

	// Act
	actual := args.Map{"notNil": model != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonModel", actual)
}

func Test_MapAnyItems_JsonModelAny(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"notNil": m.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonModelAny", actual)
}

func Test_MapAnyItems_Json_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Json", actual)
}

func Test_MapAnyItems_JsonPtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	jr := m.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonPtr", actual)
}

func Test_MapAnyItems_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr(map[string]any{"a": 1})
	result, err := m.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ParseInjectUsingJson", actual)
}

func Test_MapAnyItems_JsonParseSelfInject_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr(map[string]any{"x": 2})
	err := m.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonParseSelfInject", actual)
}

func Test_MapAnyItems_Strings_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	strs := m.Strings()

	// Act
	actual := args.Map{"notEmpty": len(strs) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Strings", actual)
}

func Test_MapAnyItems_String_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"notEmpty": m.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- String", actual)
}

func Test_MapAnyItems_DeepClear(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.DeepClear()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- DeepClear", actual)
}

func Test_MapAnyItems_Dispose(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.Dispose()

	// Act
	actual := args.Map{"nil": m.Items == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Dispose", actual)
}

func Test_MapAnyItems_IsEqualRaw_Equal(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw equal", actual)
}

func Test_MapAnyItems_IsEqualRaw_NotEqual(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw not equal", actual)
}

func Test_MapAnyItems_IsEqualRaw_DiffLen(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"a": 1, "b": 2})}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw diff len", actual)
}

func Test_MapAnyItems_IsEqualRaw_MissingKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m.IsEqualRaw(map[string]any{"b": 1})}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqualRaw missing key", actual)
}

func Test_MapAnyItems_IsEqual_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"eq": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- IsEqual", actual)
}

func Test_MapAnyItems_IsEqual_BothNil_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	var m1, m2 *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"eq": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- IsEqual both nil", actual)
}

func Test_MapAnyItems_IsEqual_OneNil(t *testing.T) {
	// Arrange
	m1 := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"eq": m1.IsEqual(nil)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- IsEqual one nil", actual)
}

func Test_MapAnyItems_ClonePtr_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	cloned, err := m.ClonePtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": cloned != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- ClonePtr", actual)
}

func Test_MapAnyItems_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	_, err := m.ClonePtr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- ClonePtr nil", actual)
}

func Test_MapAnyItems_RawMapStringAnyDiff_Nil_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	diff := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"notNil": diff != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- RawMapStringAnyDiff nil", actual)
}

func Test_MapAnyItems_MapAnyItemsSelf_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"same": m.MapAnyItems() == m}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- MapAnyItems self", actual)
}

func Test_MapAnyItems_NewUsingAnyTypeMap_Success(t *testing.T) {
	// Arrange
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- NewUsingAnyTypeMap", actual)
}

func Test_MapAnyItems_NewUsingAnyTypeMap_Nil(t *testing.T) {
	// Arrange
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- NewUsingAnyTypeMap nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionMethods — AddIf, AddCollection, ConcatNew, Clone, Capacity, etc.
// ══════════════════════════════════════════════════════════════════════════════

func Test_CollectionMethods_AddIf_True(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c.AddIf(true, "x")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns non-empty -- AddIf true", actual)
}

func Test_CollectionMethods_AddIf_False(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c.AddIf(false, "x")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns non-empty -- AddIf false", actual)
}

func Test_CollectionMethods_AddManyIf_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c.AddManyIf(true, "a", "b")
	c.AddManyIf(false, "c")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddManyIf", actual)
}

func Test_CollectionMethods_AddCollection_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c1 := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := coredynamic.New.Collection.String.From([]string{"b", "c"})
	c1.AddCollection(c2)

	// Act
	actual := args.Map{"len": c1.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCollection", actual)
}

func Test_CollectionMethods_AddCollections_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c1 := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := coredynamic.New.Collection.String.From([]string{"b"})
	c.AddCollections(c1, nil, c2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCollections", actual)
}

func Test_CollectionMethods_ConcatNew_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a"})
	c2 := c.ConcatNew("b", "c")

	// Act
	actual := args.Map{
		"origLen": c.Length(),
		"newLen": c2.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 1,
		"newLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- ConcatNew", actual)
}

func Test_CollectionMethods_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	cloned := c.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Clone", actual)
}

func Test_CollectionMethods_Clone_Nil_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]
	cloned := c.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns nil -- Clone nil", actual)
}

func Test_CollectionMethods_Capacity_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Cap(10)

	// Act
	actual := args.Map{"cap": c.Capacity() >= 10}

	// Assert
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Capacity", actual)
}

func Test_CollectionMethods_AddCapacity_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c.AddCapacity(5)

	// Act
	actual := args.Map{"cap": c.Capacity() >= 5}

	// Assert
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- AddCapacity", actual)
}

func Test_CollectionMethods_Resize_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()
	c.Resize(20)

	// Act
	actual := args.Map{"cap": c.Capacity() >= 20}

	// Assert
	expected := args.Map{"cap": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Resize", actual)
}

func Test_CollectionMethods_Reverse_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})
	c.Reverse()

	// Act
	actual := args.Map{
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- Reverse", actual)
}

func Test_CollectionMethods_InsertAt_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "c"})
	c.InsertAt(1, "b")
	items := c.Items()

	// Act
	actual := args.Map{
		"len": len(items),
		"mid": items[1],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"mid": "b",
	}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- InsertAt", actual)
}

func Test_CollectionMethods_IndexOfFunc_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})
	idx := c.IndexOfFunc(func(s string) bool { return s == "b" })

	// Act
	actual := args.Map{"idx": idx}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- IndexOfFunc", actual)
}

func Test_CollectionMethods_IndexOfFunc_NotFound(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a"})
	idx := c.IndexOfFunc(func(s string) bool { return s == "z" })

	// Act
	actual := args.Map{"idx": idx}

	// Assert
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- IndexOfFunc not found", actual)
}

func Test_CollectionMethods_ContainsFunc(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})

	// Act
	actual := args.Map{"has": c.ContainsFunc(func(s string) bool { return s == "a" })}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- ContainsFunc", actual)
}

func Test_CollectionMethods_SafeAt(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})

	// Act
	actual := args.Map{
		"valid": c.SafeAt(0),
		"invalid": c.SafeAt(99),
	}

	// Assert
	expected := args.Map{
		"valid": "a",
		"invalid": "",
	}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- SafeAt", actual)
}

func Test_CollectionMethods_SprintItems_FromMapAnyItemsDeseriali(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	strs := c.SprintItems("[%s]")

	// Act
	actual := args.Map{"first": strs[0]}

	// Assert
	expected := args.Map{"first": "[a]"}
	expected.ShouldBeEqual(t, 0, "CollectionMethods returns correct value -- SprintItems", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — additional paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectSetFromTo_BothNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(nil, nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- both nil", actual)
}

func Test_ReflectSetFromTo_SamePointerType(t *testing.T) {
	// Arrange
	x := "hello"
	src := &x
	var dst string
	err := coredynamic.ReflectSetFromTo(src, &dst)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": dst,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same pointer type", actual)
}

func Test_ReflectSetFromTo_NonPointerToPointer(t *testing.T) {
	// Arrange
	var dst int
	err := coredynamic.ReflectSetFromTo(42, &dst)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": dst,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns non-empty -- non-ptr to ptr", actual)
}

func Test_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	// Arrange
	type s struct{ X int }
	var dst s
	err := coredynamic.ReflectSetFromTo([]byte(`{"X":5}`), &dst)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": dst.X,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 5,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bytes to struct", actual)
}

func Test_ReflectSetFromTo_StructToBytes(t *testing.T) {
	// Arrange
	type s struct{ X int }
	src := s{X: 7}
	var dst []byte
	err := coredynamic.ReflectSetFromTo(src, &dst)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(dst) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- struct to bytes", actual)
}

func Test_ReflectSetFromTo_DestNotPointer(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", "not a pointer")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- dest not pointer", actual)
}

func Test_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	// Arrange
	var dst int
	err := coredynamic.ReflectSetFromTo("hello", &dst)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- type mismatch", actual)
}

func Test_ReflectSetFromTo_DestNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- dest nil", actual)
}
