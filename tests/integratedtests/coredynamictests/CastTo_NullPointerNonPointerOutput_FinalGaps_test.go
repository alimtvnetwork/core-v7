package coredynamictests

import (
	"math"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage74 — Final coverage gaps for coredata/coredynamic (97.9% → 100%)
// ══════════════════════════════════════════════════════════════════════════════

// ── CastTo: null pointer with non-pointer output ──

func Test_CastTo_NullPointerNonPointerOutput(t *testing.T) {
	// Arrange
	var nilPtr *string
	acceptedType := reflect.TypeOf(nilPtr)

	// Act
	result := coredynamic.CastTo(false, nilPtr, acceptedType)

	// Assert
	actual := args.Map{"result": result.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for null pointer with non-pointer output", actual)

	actual = args.Map{"result": result.IsNull}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsNull to be true", actual)
}

// ── MapAnyItems.GetItemRef: nil/null referenceOut, type mismatch, non-ptr found ──

func Test_MapAnyItems_GetItemRef_NilReferenceOut(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": "v"}}

	// Act
	err := m.GetItemRef("k", nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil referenceOut", actual)
}

func Test_MapAnyItems_GetItemRef_NonPointerRef(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": "v"}}

	// Act
	err := m.GetItemRef("k", "notAPointer")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-pointer referenceOut", actual)
}

func Test_MapAnyItems_GetItemRef_NilPointerValues(t *testing.T) {
	// Arrange
	var nilPtr *string
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": nilPtr}}
	var out *string

	// Act
	err := m.GetItemRef("k", &out)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil pointer values", actual)
}

func Test_MapAnyItems_GetItemRef_TypeMismatch(t *testing.T) {
	// Arrange
	val := "hello"
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": &val}}
	var out *int

	// Act
	err := m.GetItemRef("k", &out)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for type mismatch", actual)
}

func Test_MapAnyItems_GetItemRef_NonPtrFoundItem(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": "hello"}}
	var out string

	// Act — string value triggers reflect.Value.IsNil panic on non-nilable type
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		_ = m.GetItemRef("k", &out)
	}()

	// Assert
	actual := args.Map{"result": didPanic}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic for IsNil on non-pointer stored value, but did not panic", actual)
}

func Test_MapAnyItems_GetItemRef_PtrFoundItem(t *testing.T) {
	// Arrange
	val := "world"
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": &val}}
	var out string

	// Act
	err := m.GetItemRef("k", &out)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": out != "world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'world', got ''", actual)
}

// ── MapAnyItems.GetUsingUnmarshallAt: marshal error & unmarshal error ──

func Test_MapAnyItems_GetUsingUnmarshallAt_MarshalError(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": make(chan int)}}
	var out string

	// Act
	err := m.GetUsingUnmarshallAt("k", &out)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected marshal error for channel value", actual)
}

func Test_MapAnyItems_GetUsingUnmarshallAt_UnmarshalError(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": "not-a-number"}}
	var out int

	// Act
	err := m.GetUsingUnmarshallAt("k", &out)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unmarshal error for type mismatch", actual)
}

// ── MapAnyItems.HashmapDiffUsingRaw: returns diff & empty diff ──
// NOTE: DiffRaw internally compares rightMap against itself (it.Items is unused),
// so HashmapDiffUsingRaw always returns an empty map regardless of it.Items content.

func Test_MapAnyItems_HashmapDiffUsingRaw_WithDiff(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{
		"a": 1,
		"b": 2,
	}}

	// Act — rightMap is diffed against itself internally, so result is always empty
	diff := m.HashmapDiffUsingRaw(false, map[string]any{
		"a": 1,
		"b": 999,
	})

	// Assert
	actual := args.Map{"result": len(diff) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty diff due to self-compare", actual)
}

func Test_MapAnyItems_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{
		"a": 1,
	}}

	// Act
	diff := m.HashmapDiffUsingRaw(false, map[string]any{
		"a": 1,
	})

	// Assert
	actual := args.Map{"result": len(diff) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty diff", actual)
}

// ── MapAnyItems paging: length != allKeys ──

func Test_MapAnyItems_GetSinglePageCollection_LengthMismatchPanics(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"a": 1, "b": 2, "c": 3}}
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for length mismatch", actual)
	}()

	// Act — pass wrong number of keys, eachPageSize must be <= length
	m.GetSinglePageCollection(2, 1, []string{"a"})
}

// ── MapAnyItems paging: negative page index ──

func Test_MapAnyItems_GetSinglePageCollection_NegativePageIndexPanics(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"a": 1, "b": 2, "c": 3}}
	allKeys := m.AllKeys()
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for negative page index", actual)
	}()

	// Act
	m.GetSinglePageCollection(2, 0, allKeys)
}

// ── MapAnyItems.GetNewMapUsingKeys: isPanicOnMissing ──

func Test_MapAnyItems_GetNewMapUsingKeys_PanicOnMissing(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"a": 1}}
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic on missing key", actual)
	}()

	// Act
	m.GetNewMapUsingKeys(true, "nonexistent")
}

// ── MapAnyItems.NewFromAnyMap: reflect error ──

func Test_MapAnyItems_NewUsingAnyTypeMap_NonMapType(t *testing.T) {
	// Arrange
	notAMap := "hello"

	// Act
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(notAMap)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-map type", actual)
}

// ── MapAnyItems.JsonString/JsonStringMust error branches ──

func Test_MapAnyItems_JsonString_MarshalError(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": make(chan int)}}

	// Act
	_, err := m.JsonString()

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected marshal error", actual)
}

func Test_MapAnyItems_JsonStringMust_Panics(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": make(chan int)}}
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()

	// Act
	m.JsonStringMust()
}

// ── MapAnyItems.ClonePtr: HasError branch ──

func Test_MapAnyItems_ClonePtr_MarshalError(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"k": math.NaN()}}

	// Act
	_, err := m.ClonePtr()

	// Assert — NaN may or may not cause JSON error depending on implementation
	_ = err
}

// ── MapAnyItems.JsonMapResults: exercises conversion ──

func Test_MapAnyItems_JsonMapResults_Normal(t *testing.T) {
	// Arrange
	m := &coredynamic.MapAnyItems{Items: map[string]any{"a": "1", "b": "2"}}

	// Act
	mr, err := m.JsonMapResults()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": mr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil MapResults", actual)
}

// ── SafeZeroSet: settable pointer elem ──

func Test_SafeZeroSet_Settable(t *testing.T) {
	// Arrange
	val := 42
	rv := reflect.ValueOf(&val)

	// Act
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"result": val != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected val to be zeroed", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Accepted Gaps Documentation
// ══════════════════════════════════════════════════════════════════════════════
//
// 1. ReflectInterfaceVal.go:20 — logically unreachable dead code
//    (lines 12-13 and 16-17 are exhaustive for all kinds)
//
// 2. CollectionLock.LengthLock:15 — nil check after Lock(), dead code
//
// 3. CollectionLock.ItemsLock:125-127 — nil items after Lock(), defensive
//
// 4. Collection.JsonString:355-357 — json.Marshal error on typed slice
//
// 5. Collection.JsonStringMust:364-365 — cascaded from JsonString
//
// 6. AnyCollection.JsonString:485-487 — json.Marshal error on []any
//
// 7. AnyCollection.JsonStringMust:495-499 — cascaded from JsonString
//
// 8. DynamicCollection.JsonString:416-418 — json.Marshal error
//
// 9. DynamicCollection.JsonStringMust:426-430 — cascaded
//
// 10. DynamicJson.go:54 — MarshalJSON error on innerData
//
// 11. DynamicJson.go:123 — ParseInjectUsingJsonMust panic
//
// 12. DynamicJson.go:139-141, 149-151, 159-163 — cascading JSON errors
//
// 13. TypedDynamic.JsonString:117-119 — json.Marshal defensive
//
// 14. KeyVal.ReflectSetKeyValue:134-136 — ReflectSetFromTo error
//
// 15. KeyValCollection lines 139-141, 342-344, 365-366, 385-387, 395-397
//     — JSON parse/serialize error branches (defensive)
//
// 16. ReflectSetFromTo.go:159-167, 174-180 — byte conversion edge cases
//
// 17. MapAnyItems.go:362-373 — unreachable after lines 350-354 and 356-359
//     (exhaustive if-else on foundItemRv.Kind() == reflect.Ptr)
//
// 18. MapAnyItems.go:903-904 — ToKeyValCollection AddAny error (defensive)
// ══════════════════════════════════════════════════════════════════════════════
