package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CastTo + CastedResult
// ══════════════════════════════════════════════════════════════════════════════

func Test_CastTo_MatchingType(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"match": result.IsMatchingAcceptedType,
		"noErr": !result.HasError(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"match": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- matching", actual)
}

func Test_CastTo_NonMatchingType(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{
		"match": result.IsMatchingAcceptedType,
		"hasErr": result.HasError(),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns non-empty -- non-matching", actual)
}

func Test_CastTo_PointerOutput(t *testing.T) {
	// Arrange
	input := "hello"
	result := coredynamic.CastTo(true, &input, reflect.TypeOf((*string)(nil)))

	// Act
	actual := args.Map{
		"notNil": result.Casted != nil,
		"match": result.IsMatchingAcceptedType,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"match": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- pointer output", actual)
}

func Test_CastTo_MultipleAccepted(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, 42, reflect.TypeOf(""), reflect.TypeOf(0))

	// Act
	actual := args.Map{"match": result.IsMatchingAcceptedType}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- multiple accepted", actual)
}

func Test_CastedResult_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"invalid": cr.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns nil -- IsInvalid nil", actual)
}

func Test_CastedResult_IsNotNull(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsNull: false}

	// Act
	actual := args.Map{"notNull": cr.IsNotNull()}

	// Assert
	expected := args.Map{"notNull": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotNull", actual)
}

func Test_CastedResult_IsNotNull_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"notNull": cr.IsNotNull()}

	// Assert
	expected := args.Map{"notNull": false}
	expected.ShouldBeEqual(t, 0, "CastedResult returns nil -- IsNotNull nil", actual)
}

func Test_CastedResult_IsNotPointer(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsPointer: false}

	// Act
	actual := args.Map{"notPtr": cr.IsNotPointer()}

	// Assert
	expected := args.Map{"notPtr": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotPointer", actual)
}

func Test_CastedResult_IsNotMatchingAcceptedType(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}

	// Act
	actual := args.Map{"notMatch": cr.IsNotMatchingAcceptedType()}

	// Assert
	expected := args.Map{"notMatch": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotMatchingAcceptedType", actual)
}

func Test_CastedResult_IsSourceKind(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{SourceKind: reflect.String}

	// Act
	actual := args.Map{
		"isStr": cr.IsSourceKind(reflect.String),
		"isInt": cr.IsSourceKind(reflect.Int),
	}

	// Assert
	expected := args.Map{
		"isStr": true,
		"isInt": false,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsSourceKind", actual)
}

func Test_CastedResult_HasAnyIssues_Valid(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}

	// Act
	actual := args.Map{"issues": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"issues": false}
	expected.ShouldBeEqual(t, 0, "CastedResult returns non-empty -- HasAnyIssues valid", actual)
}

func Test_CastedResult_HasAnyIssues_Invalid(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: false, IsNull: true, IsMatchingAcceptedType: false}

	// Act
	actual := args.Map{"issues": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"issues": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns error -- HasAnyIssues invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItemDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapAnyItemDiff_Length(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Length", actual)
}

func Test_MapAnyItemDiff_Length_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Length nil", actual)
}

func Test_MapAnyItemDiff_IsEmpty(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{}

	// Act
	actual := args.Map{"empty": m.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- IsEmpty", actual)
}

func Test_MapAnyItemDiff_HasAnyItem(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"has": m.HasAnyItem()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyItem", actual)
}

func Test_MapAnyItemDiff_LastIndex(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}

	// Act
	actual := args.Map{"idx": m.LastIndex()}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LastIndex", actual)
}

func Test_MapAnyItemDiff_AllKeysSorted_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"b": 2, "a": 1}
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
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- AllKeysSorted", actual)
}

func Test_MapAnyItemDiff_Raw_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"len": len(raw)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Raw", actual)
}

func Test_MapAnyItemDiff_Raw_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff
	raw := m.Raw()

	// Act
	actual := args.Map{"len": len(raw)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Raw nil", actual)
}

func Test_MapAnyItemDiff_Clear_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Clear", actual)
}

func Test_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Clear nil", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_Same(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual same", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_Diff(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual diff", actual)
}

func Test_MapAnyItemDiff_HasAnyChanges_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"changes": m.HasAnyChanges(false, map[string]any{"a": 2})}

	// Assert
	expected := args.Map{"changes": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyChanges", actual)
}

func Test_MapAnyItemDiff_DiffRaw_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 99})

	// Act
	actual := args.Map{"hasDiff": len(diff) > 0}

	// Assert
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffRaw", actual)
}

func Test_MapAnyItemDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"empty": diff.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- HashmapDiffUsingRaw no diff", actual)
}

func Test_MapAnyItemDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{"hasDiff": diff.HasAnyItem()}

	// Assert
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HashmapDiffUsingRaw has diff", actual)
}

func Test_MapAnyItemDiff_MapAnyItems_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	mai := m.MapAnyItems()

	// Act
	actual := args.Map{"notNil": mai != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- MapAnyItems", actual)
}

func Test_MapAnyItemDiff_RawMapDiffer_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	d := m.RawMapDiffer()

	// Act
	actual := args.Map{"notNil": d != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- RawMapDiffer", actual)
}

func Test_MapAnyItemDiff_DiffJsonMessage_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffJsonMessage", actual)
}

func Test_MapAnyItemDiff_ShouldDiffMessage_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ShouldDiffMessage", actual)
}

func Test_MapAnyItemDiff_LogShouldDiffMessage_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LogShouldDiffMessage", actual)
}

func Test_MapAnyItemDiff_ToStringsSliceOfDiffMap_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.DiffRaw(false, map[string]any{"a": 2})
	strs := m.ToStringsSliceOfDiffMap(diff)

	// Act
	actual := args.Map{"hasItems": len(strs) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ToStringsSliceOfDiffMap", actual)
}

func Test_MapAnyItemDiff_Json_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	jr := m.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Json", actual)
}

func Test_MapAnyItemDiff_JsonPtr(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	jr := m.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- JsonPtr", actual)
}

func Test_MapAnyItemDiff_PrettyJsonString(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	s := m.PrettyJsonString()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- PrettyJsonString", actual)
}

func Test_MapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	m.LogPrettyJsonString() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LogPrettyJsonString", actual)
}

func Test_MapAnyItemDiff_LogPrettyJsonString_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{}
	m.LogPrettyJsonString() // empty path

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- LogPrettyJsonString empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapAsKeyValSlice_Valid(t *testing.T) {
	// Arrange
	m := map[string]any{"a": 1, "b": 2}
	rv := reflect.ValueOf(m)
	kvc, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": kvc.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns non-empty -- valid", actual)
}

func Test_MapAsKeyValSlice_NonMap(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")
	_, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns non-empty -- non-map", actual)
}

func Test_MapAsKeyValSlice_Pointer(t *testing.T) {
	// Arrange
	m := map[string]any{"a": 1}
	rv := reflect.ValueOf(&m)
	kvc, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasItems": kvc.Length() > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": false,
		"hasItems": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypeNotEqualErr + TypeMustBeSame + TypeSameStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypeNotEqualErr_Same(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- same", actual)
}

func Test_TypeNotEqualErr_Different(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", 1)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- different", actual)
}

func Test_TypeMustBeSame_NoPanic(t *testing.T) {
	// Arrange
	coredynamic.TypeMustBeSame("a", "b")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- no panic", actual)
}

func Test_TypeMustBeSame_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- panic", actual)
	}()
	coredynamic.TypeMustBeSame("a", 1)
}

func Test_TypeSameStatus_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- same", actual)
}

func Test_TypeSameStatus_Different(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 1)

	// Act
	actual := args.Map{
		"same": ts.IsSame,
		"leftPtr": ts.IsLeftPointer,
	}

	// Assert
	expected := args.Map{
		"same": false,
		"leftPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- different", actual)
}

func Test_TypeSameStatus_Pointer(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "b")

	// Act
	actual := args.Map{
		"leftPtr": ts.IsLeftPointer,
		"rightPtr": ts.IsRightPointer,
	}

	// Assert
	expected := args.Map{
		"leftPtr": true,
		"rightPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectKindValidation + ReflectTypeValidation
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectKindValidation_Match(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.String, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns non-empty -- match", actual)
}

func Test_ReflectKindValidation_Mismatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.Int, "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns non-empty -- mismatch", actual)
}

func Test_ReflectTypeValidation_Match(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- match", actual)
}

func Test_ReflectTypeValidation_Mismatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- mismatch", actual)
}

func Test_ReflectTypeValidation_NilNotAllowed_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil not allowed", actual)
}

func Test_ReflectTypeValidation_NilAllowed_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil allowed but mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ZeroSet, ZeroSetAny, SafeZeroSet
// ══════════════════════════════════════════════════════════════════════════════

func Test_ZeroSet(t *testing.T) {
	// Arrange
	val := 42
	rv := reflect.ValueOf(&val)
	coredynamic.ZeroSet(rv)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSet returns correct value -- with args", actual)
}

func Test_ZeroSetAny(t *testing.T) {
	// Arrange
	val := "hello"
	coredynamic.ZeroSetAny(&val)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns correct value -- with args", actual)
}

func Test_ZeroSetAny_Nil(t *testing.T) {
	// Arrange
	coredynamic.ZeroSetAny(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns nil -- nil", actual)
}

func Test_SafeZeroSet(t *testing.T) {
	// Arrange
	val := 42
	rv := reflect.ValueOf(&val)
	coredynamic.SafeZeroSet(rv)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns correct value -- with args", actual)
}

func Test_SafeZeroSet_NilType(t *testing.T) {
	// Arrange
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		coredynamic.SafeZeroSet(reflect.ValueOf(nil))
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- nil reflect value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr + MustBeAcceptedTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_NotAcceptedTypesErr_Accepted(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- accepted", actual)
}

func Test_NotAcceptedTypesErr_NotAccepted(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- not accepted", actual)
}

func Test_MustBeAcceptedTypes_NoPanic(t *testing.T) {
	// Arrange
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes panics -- no panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PointerOrNonPointer + UsingReflectValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_PointerOrNonPointer_NonPointer(t *testing.T) {
	// Arrange
	val, _ := coredynamic.PointerOrNonPointer(false, "hello")

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns non-empty -- non-ptr", actual)
}

func Test_PointerOrNonPointer_PointerInput_NonPointerOutput(t *testing.T) {
	// Arrange
	s := "hello"
	val, _ := coredynamic.PointerOrNonPointer(false, &s)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns non-empty -- ptr->non-ptr", actual)
}

func Test_PointerOrNonPointerUsingReflectValue_NonPtr(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")
	val, _ := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue returns non-empty -- non-ptr", actual)
}

func Test_PointerOrNonPointerUsingReflectValue_PtrInput_NonPtrOut(t *testing.T) {
	// Arrange
	s := "hello"
	rv := reflect.ValueOf(&s)
	val, _ := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue returns non-empty -- ptr->non-ptr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ══════════════════════════════════════════════════════════════════════════════

func Test_LengthOfReflect_Slice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- slice", actual)
}

func Test_LengthOfReflect_Array(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([3]int{1, 2, 3})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- array", actual)
}

func Test_LengthOfReflect_Map(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- map", actual)
}

func Test_LengthOfReflect_Other(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- other", actual)
}

func Test_LengthOfReflect_Pointer(t *testing.T) {
	// Arrange
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsAnyTypesOf + TypesIndexOf
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsAnyTypesOf_Found(t *testing.T) {
	// Act
	actual := args.Map{"found": coredynamic.IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(0), reflect.TypeOf(""))}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- found", actual)
}

func Test_IsAnyTypesOf_NotFound(t *testing.T) {
	// Act
	actual := args.Map{"found": coredynamic.IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(0))}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- not found", actual)
}

func Test_TypesIndexOf_Found(t *testing.T) {
	// Arrange
	idx := coredynamic.TypesIndexOf(reflect.TypeOf(""), reflect.TypeOf(0), reflect.TypeOf(""))

	// Act
	actual := args.Map{"idx": idx}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- found", actual)
}

func Test_TypesIndexOf_NotFound(t *testing.T) {
	// Arrange
	idx := coredynamic.TypesIndexOf(reflect.TypeOf(""), reflect.TypeOf(0))

	// Act
	actual := args.Map{"idx": idx}

	// Assert
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- not found", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeTypeName + ReflectInterfaceVal + ValueStatus + DynamicStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_SafeTypeName_String(t *testing.T) {
	// Act
	actual := args.Map{"name": coredynamic.SafeTypeName("hello")}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- string", actual)
}

func Test_SafeTypeName_Nil(t *testing.T) {
	// Act
	actual := args.Map{"name": coredynamic.SafeTypeName(nil)}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns nil -- nil", actual)
}

func Test_ReflectInterfaceVal_NonPointer_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	val := coredynamic.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns non-empty -- non-ptr", actual)
}

func Test_ReflectInterfaceVal_Pointer_FromCastToMatchingTypeIt(t *testing.T) {
	// Arrange
	s := "hello"
	val := coredynamic.ReflectInterfaceVal(&s)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- pointer", actual)
}

func Test_ValueStatus_Invalid(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("bad")

	// Act
	actual := args.Map{
		"valid": vs.IsValid,
		"msg": vs.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "bad",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns error -- invalid", actual)
}

func Test_ValueStatus_InvalidNoMessage(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatusNoMessage()

	// Act
	actual := args.Map{
		"valid": vs.IsValid,
		"msg": vs.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns empty -- invalid no msg", actual)
}

func Test_DynamicStatus_Invalid(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("bad")

	// Act
	actual := args.Map{
		"valid": ds.IsValid(),
		"msg": ds.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "bad",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns error -- invalid", actual)
}

func Test_DynamicStatus_InvalidNoMessage(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatusNoMessage()

	// Act
	actual := args.Map{
		"valid": ds.IsValid(),
		"msg": ds.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns empty -- invalid no msg", actual)
}

func Test_DynamicStatus_Clone(t *testing.T) {
	// Arrange
	ds := coredynamic.DynamicStatus{
		Dynamic: coredynamic.NewDynamic("hello", true),
		Message: "test",
	}
	cloned := ds.Clone()

	// Act
	actual := args.Map{
		"valid": cloned.IsValid(),
		"msg": cloned.Message,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"msg": "test",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- Clone", actual)
}

func Test_DynamicStatus_ClonePtr(t *testing.T) {
	// Arrange
	ds := &coredynamic.DynamicStatus{
		Dynamic: coredynamic.NewDynamic("hello", true),
		Message: "test",
	}
	cloned := ds.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"msg": cloned.Message,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"msg": "test",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- ClonePtr", actual)
}

func Test_DynamicStatus_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var ds *coredynamic.DynamicStatus

	// Act
	actual := args.Map{"nil": ds.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns nil -- ClonePtr nil", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_RegardlessType(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{"eq": m.IsRawEqual(true, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual regardless type", actual)
}
