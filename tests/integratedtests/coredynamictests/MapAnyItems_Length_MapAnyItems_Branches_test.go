package coredynamictests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// MapAnyItems — nil/empty receiver branches
// =============================================================================

func Test_MapAnyItems_Length_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Length empty", actual)
}

func Test_MapAnyItems_IsEmpty_True(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": m.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEmpty true", actual)
}

func Test_MapAnyItems_HasAnyItem_False(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": m.HasAnyItem()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasAnyItem false", actual)
}

func Test_MapAnyItems_HasAnyItem_True(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"r": m.HasAnyItem()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasAnyItem true", actual)
}

func Test_MapAnyItems_HasKey_Nil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"r": m.HasKey("x")}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasKey nil", actual)
}

func Test_MapAnyItems_HasKey_Missing_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": m.HasKey("x")}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasKey missing", actual)
}

func Test_MapAnyItems_HasKey_Found(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("x", 1)

	// Act
	actual := args.Map{"r": m.HasKey("x")}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasKey found", actual)
}

// =============================================================================
// MapAnyItems — Get/GetValue branches
// =============================================================================

func Test_MapAnyItems_GetValue_Found(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "val")

	// Act
	actual := args.Map{"r": m.GetValue("k")}

	// Assert
	expected := args.Map{"r": "val"}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetValue found", actual)
}

func Test_MapAnyItems_GetValue_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"isNil": m.GetValue("missing") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetValue missing", actual)
}

func Test_MapAnyItems_Get_Found(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", 42)
	val, has := m.Get("k")

	// Act
	actual := args.Map{
		"val": val,
		"has": has,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Get found", actual)
}

func Test_MapAnyItems_Get_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	val, has := m.Get("missing")

	// Act
	actual := args.Map{
		"isNil": val == nil,
		"has": has,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Get missing", actual)
}

// =============================================================================
// MapAnyItems — GetFieldsMap / GetSafeFieldsMap
// =============================================================================

func Test_MapAnyItems_GetFieldsMap_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	_, _, isFound := m.GetFieldsMap("missing")

	// Act
	actual := args.Map{"isFound": isFound}

	// Assert
	expected := args.Map{"isFound": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetFieldsMap missing", actual)
}

func Test_MapAnyItems_GetFieldsMap_Found_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", map[string]any{"a": 1})
	fieldMap, _, isFound := m.GetFieldsMap("k")

	// Act
	actual := args.Map{
		"isFound": isFound,
		"hasA": fieldMap != nil,
	}

	// Assert
	expected := args.Map{
		"isFound": true,
		"hasA": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetFieldsMap found", actual)
}

func Test_MapAnyItems_GetSafeFieldsMap_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	_, isFound := m.GetSafeFieldsMap("missing")

	// Act
	actual := args.Map{"isFound": isFound}

	// Assert
	expected := args.Map{"isFound": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetSafeFieldsMap missing", actual)
}

// =============================================================================
// MapAnyItems — Deserialize / GetUsingUnmarshallAt branches
// =============================================================================

func Test_MapAnyItems_GetUsingUnmarshallAt_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var out string
	err := m.GetUsingUnmarshallAt("missing", &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetUsingUnmarshallAt missing key", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallAt_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "hello")
	var out string
	err := m.GetUsingUnmarshallAt("k", &out)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetUsingUnmarshallAt valid", actual)
}

func Test_MapAnyItems_Deserialize_Missing_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var out int
	err := m.Deserialize("x", &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Deserialize missing", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallManyAt_Error_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var out string
	err := m.GetUsingUnmarshallManyAt(corejson.KeyAny{Key: "missing", AnyInf: &out})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetUsingUnmarshallManyAt error", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallManyAt_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k1", "v1")
	m.Add("k2", 42)
	var s string
	var n int
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "k1", AnyInf: &s},
		corejson.KeyAny{Key: "k2", AnyInf: &n},
	)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"s": s,
		"n": n,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"s": "v1",
		"n": 42,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetUsingUnmarshallManyAt valid", actual)
}

// =============================================================================
// MapAnyItems — GetItemRef branches
// =============================================================================

func Test_MapAnyItems_GetItemRef_Missing_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	var out string
	err := m.GetItemRef("missing", &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetItemRef missing", actual)
}

func Test_MapAnyItems_GetItemRef_NilRef_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "v")
	err := m.GetItemRef("k", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetItemRef nil ref", actual)
}

func Test_MapAnyItems_GetItemRef_NotPointer_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "v")
	err := m.GetItemRef("k", "notpointer")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetItemRef not pointer", actual)
}

func Test_MapAnyItems_GetManyItemsRefs_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.GetManyItemsRefs()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetManyItemsRefs empty", actual)
}

func Test_MapAnyItems_GetManyItemsRefs_Error(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.GetManyItemsRefs(corejson.KeyAny{Key: "missing", AnyInf: nil})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetManyItemsRefs error", actual)
}

// =============================================================================
// MapAnyItems — Add/Set branches
// =============================================================================

func Test_MapAnyItems_Add_New(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.Add("k", 1)

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Add new", actual)
}

func Test_MapAnyItems_Add_Override(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", 1)
	isNew := m.Add("k", 2)

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Add override", actual)
}

func Test_MapAnyItems_Set_New(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.Set("k", 1)

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Set new", actual)
}

func Test_MapAnyItems_AddKeyAny_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: 1})

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddKeyAny", actual)
}

func Test_MapAnyItems_AddKeyAnyWithValidation_Mismatch_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(""),
		corejson.KeyAny{Key: "k", AnyInf: 42},
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddKeyAnyWithValidation mismatch", actual)
}

func Test_MapAnyItems_AddKeyAnyWithValidation_Match_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddKeyAnyWithValidation(
		reflect.TypeOf(""),
		corejson.KeyAny{Key: "k", AnyInf: "val"},
	)

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
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddKeyAnyWithValidation match", actual)
}

func Test_MapAnyItems_AddWithValidation_Mismatch_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddWithValidation(reflect.TypeOf(0), "k", "notint")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddWithValidation mismatch", actual)
}

func Test_MapAnyItems_AddWithValidation_Match_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddWithValidation(reflect.TypeOf(0), "k", 42)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddWithValidation match", actual)
}

// =============================================================================
// MapAnyItems — AddJsonResultPtr branches
// =============================================================================

func Test_MapAnyItems_AddJsonResultPtr_Nil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddJsonResultPtr("k", nil)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddJsonResultPtr nil", actual)
}

func Test_MapAnyItems_AddJsonResultPtr_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := corejson.NewPtr("hello")
	r := m.AddJsonResultPtr("k", jr)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddJsonResultPtr valid", actual)
}

// =============================================================================
// MapAnyItems — AddMapResult / AddMapResultOption / AddManyMapResultsUsingOption
// =============================================================================

func Test_MapAnyItems_AddMapResult_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddMapResult(nil)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddMapResult empty", actual)
}

func Test_MapAnyItems_AddMapResult_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddMapResult(map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddMapResult valid", actual)
}

func Test_MapAnyItems_AddMapResultOption_Override_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.AddMapResultOption(true, map[string]any{"a": 99})

	// Act
	actual := args.Map{"val": r.GetValue("a")}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddMapResultOption override", actual)
}

func Test_MapAnyItems_AddMapResultOption_NoOverride_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddMapResultOption(false, nil)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddMapResultOption no override empty", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddManyMapResultsUsingOption(true)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddManyMapResultsUsingOption empty", actual)
}

func Test_MapAnyItems_AddManyMapResultsUsingOption_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.AddManyMapResultsUsingOption(true,
		map[string]any{"a": 1},
		map[string]any{"b": 2},
	)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AddManyMapResultsUsingOption valid", actual)
}

// =============================================================================
// MapAnyItems — Paging branches
// =============================================================================

func Test_MapAnyItems_GetPagesSize_Zero_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": m.GetPagesSize(0)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetPagesSize zero", actual)
}

func Test_MapAnyItems_GetPagesSize_Negative(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": m.GetPagesSize(-1)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetPagesSize negative", actual)
}

func Test_MapAnyItems_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)

	// Act
	actual := args.Map{"r": m.GetPagesSize(2)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetPagesSize valid", actual)
}

func Test_MapAnyItems_GetPagedCollection_SmallData(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	pages := m.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetPagedCollection small data", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetNewMapUsingKeys empty", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_NoPanic(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	m.Add("b", 2)
	r := m.GetNewMapUsingKeys(false, "a", "missing")

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems GetNewMapUsingKeys noPanic", actual)
}

// =============================================================================
// MapAnyItems — JSON branches
// =============================================================================

func Test_MapAnyItems_JsonString_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "v")
	s, err := m.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nonEmpty": len(s) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonString valid", actual)
}

func Test_MapAnyItems_JsonStringMust_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "v")
	s := m.JsonStringMust()

	// Act
	actual := args.Map{"nonEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonStringMust valid", actual)
}

func Test_MapAnyItems_JsonResultOfKey_Found_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("k", "v")
	jr := m.JsonResultOfKey("k")

	// Act
	actual := args.Map{"noErr": !jr.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultOfKey found", actual)
}

func Test_MapAnyItems_JsonResultOfKey_Missing_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := m.JsonResultOfKey("missing")

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultOfKey missing", actual)
}

func Test_MapAnyItems_JsonResultOfKeys_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.JsonResultOfKeys()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultOfKeys empty", actual)
}

func Test_MapAnyItems_JsonResultOfKeys_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.JsonResultOfKeys("a")

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultOfKeys valid", actual)
}

func Test_MapAnyItems_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := m.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems ParseInjectUsingJson error", actual)
}

func Test_MapAnyItems_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		m.ParseInjectUsingJsonMust(jr)
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems ParseInjectUsingJsonMust panics", actual)
}

func Test_MapAnyItems_JsonParseSelfInject_Error(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	jr := &corejson.Result{Error: errors.New("fail")}
	err := m.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonParseSelfInject error", actual)
}

// =============================================================================
// MapAnyItems — AllKeys / AllKeysSorted / AllValues
// =============================================================================

func Test_MapAnyItems_AllKeys_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"len": len(m.AllKeys())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllKeys empty", actual)
}

func Test_MapAnyItems_AllKeys_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("b", 2)
	m.Add("a", 1)

	// Act
	actual := args.Map{"len": len(m.AllKeys())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllKeys valid", actual)
}

func Test_MapAnyItems_AllKeysSorted_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"len": len(m.AllKeysSorted())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllKeysSorted empty", actual)
}

func Test_MapAnyItems_AllKeysSorted_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("b", 2)
	m.Add("a", 1)
	keys := m.AllKeysSorted()

	// Act
	actual := args.Map{
		"first": keys[0],
		"second": keys[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllKeysSorted valid", actual)
}

func Test_MapAnyItems_AllValues_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"len": len(m.AllValues())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllValues empty", actual)
}

func Test_MapAnyItems_AllValues_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"len": len(m.AllValues())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems AllValues valid", actual)
}

// =============================================================================
// MapAnyItems — IsEqual / IsEqualRaw branches
// =============================================================================

func Test_MapAnyItems_IsEqual_BothNil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	var a *coredynamic.MapAnyItems
	var b *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqual both nil", actual)
}

func Test_MapAnyItems_IsEqual_LeftNil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	var a *coredynamic.MapAnyItems
	b := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqual left nil", actual)
}

func Test_MapAnyItems_IsEqual_RightNil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": a.IsEqual(nil)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqual right nil", actual)
}

func Test_MapAnyItems_IsEqual_DiffLength(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()
	a.Add("a", 1)
	b := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqual diff length", actual)
}

func Test_MapAnyItems_IsEqual_Same(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()
	a.Add("a", 1)
	b := coredynamic.EmptyMapAnyItems()
	b.Add("a", 1)

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqual same", actual)
}

func Test_MapAnyItems_IsEqualRaw_BothNil(t *testing.T) {
	// Arrange
	var a *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"r": a.IsEqualRaw(nil)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqualRaw both nil", actual)
}

func Test_MapAnyItems_IsEqualRaw_LeftNil(t *testing.T) {
	// Arrange
	var a *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"r": a.IsEqualRaw(map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqualRaw left nil", actual)
}

func Test_MapAnyItems_IsEqualRaw_DiffLength(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()
	a.Add("a", 1)

	// Act
	actual := args.Map{"r": a.IsEqualRaw(map[string]any{})}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqualRaw diff length", actual)
}

func Test_MapAnyItems_IsEqualRaw_MissingKey_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()
	a.Add("a", 1)

	// Act
	actual := args.Map{"r": a.IsEqualRaw(map[string]any{"b": 1})}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqualRaw missing key", actual)
}

func Test_MapAnyItems_IsEqualRaw_DiffValue(t *testing.T) {
	// Arrange
	a := coredynamic.EmptyMapAnyItems()
	a.Add("a", 1)

	// Act
	actual := args.Map{"r": a.IsEqualRaw(map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsEqualRaw diff value", actual)
}

// =============================================================================
// MapAnyItems — Clear / DeepClear / Dispose
// =============================================================================

func Test_MapAnyItems_Clear_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	m.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Clear nil", actual)
}

func Test_MapAnyItems_Clear_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	m.Clear()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Clear valid", actual)
}

func Test_MapAnyItems_DeepClear_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	m.DeepClear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems DeepClear nil", actual)
}

func Test_MapAnyItems_DeepClear_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	m.DeepClear()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems DeepClear valid", actual)
}

func Test_MapAnyItems_Dispose_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	m.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Dispose nil", actual)
}

func Test_MapAnyItems_Dispose_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	m.Dispose()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Dispose valid", actual)
}

// =============================================================================
// MapAnyItems — ClonePtr
// =============================================================================

func Test_MapAnyItems_ClonePtr_Nil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	_, err := m.ClonePtr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems ClonePtr nil", actual)
}

func Test_MapAnyItems_ClonePtr_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	c, err := m.ClonePtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems ClonePtr valid", actual)
}

// =============================================================================
// MapAnyItems — Misc
// =============================================================================

func Test_MapAnyItems_RawMapStringAnyDiff_Nil_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	r := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems RawMapStringAnyDiff nil", actual)
}

func Test_MapAnyItems_String_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"nonEmpty": len(m.String()) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems String valid", actual)
}

func Test_MapAnyItems_Strings_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"nonEmpty": len(m.Strings()) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Strings valid", actual)
}

func Test_MapAnyItems_MapAnyItemsSelf_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"same": m.MapAnyItems() == m}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems MapAnyItems self", actual)
}

func Test_MapAnyItems_JsonModel_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.JsonModel()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonModel empty", actual)
}

func Test_MapAnyItems_JsonModel_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.JsonModel()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonModel valid", actual)
}

func Test_MapAnyItems_JsonModelAny_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"notNil": m.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonModelAny", actual)
}

func Test_MapAnyItems_Json_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems Json", actual)
}

func Test_MapAnyItems_JsonPtr_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.JsonPtr()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonPtr", actual)
}

func Test_MapAnyItems_JsonMapResults_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r, err := m.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": r != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonMapResults empty", actual)
}

func Test_MapAnyItems_JsonMapResults_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r, err := m.JsonMapResults()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": r != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonMapResults valid", actual)
}

func Test_MapAnyItems_JsonResultsCollection_Empty_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultsCollection empty", actual)
}

func Test_MapAnyItems_JsonResultsCollection_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultsCollection valid", actual)
}

func Test_MapAnyItems_JsonResultsPtrCollection_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	r := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultsPtrCollection empty", actual)
}

func Test_MapAnyItems_JsonResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	r := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems JsonResultsPtrCollection valid", actual)
}

// =============================================================================
// MapAnyItems — NewMapAnyItemsUsingAnyTypeMap
// =============================================================================

func Test_MapAnyItems_NewMapAnyItemsUsingAnyTypeMap_Nil(t *testing.T) {
	// Arrange
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingAnyTypeMap nil", actual)
}

func Test_MapAnyItems_NewMapAnyItemsUsingAnyTypeMap_Valid(t *testing.T) {
	// Arrange
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]any{"a": 1})

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
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingAnyTypeMap valid", actual)
}

func Test_MapAnyItems_NewMapAnyItemsUsingItems_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingItems empty", actual)
}

func Test_MapAnyItems_NewMapAnyItemsUsingItems_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewMapAnyItemsUsingItems valid", actual)
}

// =============================================================================
// MapAnyItems — Diff branches
// =============================================================================

func Test_MapAnyItems_IsRawEqual_True(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"r": m.IsRawEqual(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems IsRawEqual true", actual)
}

func Test_MapAnyItems_HasAnyChanges_True(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"r": m.HasAnyChanges(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasAnyChanges true", actual)
}

func Test_MapAnyItems_HasAnyChanges_False(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)

	// Act
	actual := args.Map{"r": m.HasAnyChanges(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HasAnyChanges false", actual)
}

func Test_MapAnyItems_HashmapDiffUsingRaw_NoDiff_FromMapAnyItemsLengthMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.Add("a", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"len": len(diff)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems HashmapDiffUsingRaw no diff", actual)
}
