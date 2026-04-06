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

func Test_I24_CastTo_MatchingType(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid, "match": result.IsMatchingAcceptedType, "noErr": !result.HasError()}
	expected := args.Map{"valid": true, "match": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- matching", actual)
}

func Test_I24_CastTo_NonMatchingType(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))
	actual := args.Map{"match": result.IsMatchingAcceptedType, "hasErr": result.HasError()}
	expected := args.Map{"match": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns non-empty -- non-matching", actual)
}

func Test_I24_CastTo_PointerOutput(t *testing.T) {
	input := "hello"
	result := coredynamic.CastTo(true, &input, reflect.TypeOf((*string)(nil)))
	actual := args.Map{"notNil": result.Casted != nil, "match": result.IsMatchingAcceptedType}
	expected := args.Map{"notNil": true, "match": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- pointer output", actual)
}

func Test_I24_CastTo_MultipleAccepted(t *testing.T) {
	result := coredynamic.CastTo(false, 42, reflect.TypeOf(""), reflect.TypeOf(0))
	actual := args.Map{"match": result.IsMatchingAcceptedType}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- multiple accepted", actual)
}

func Test_I24_CastedResult_IsInvalid_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"invalid": cr.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns nil -- IsInvalid nil", actual)
}

func Test_I24_CastedResult_IsNotNull(t *testing.T) {
	cr := &coredynamic.CastedResult{IsNull: false}
	actual := args.Map{"notNull": cr.IsNotNull()}
	expected := args.Map{"notNull": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotNull", actual)
}

func Test_I24_CastedResult_IsNotNull_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"notNull": cr.IsNotNull()}
	expected := args.Map{"notNull": false}
	expected.ShouldBeEqual(t, 0, "CastedResult returns nil -- IsNotNull nil", actual)
}

func Test_I24_CastedResult_IsNotPointer(t *testing.T) {
	cr := &coredynamic.CastedResult{IsPointer: false}
	actual := args.Map{"notPtr": cr.IsNotPointer()}
	expected := args.Map{"notPtr": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotPointer", actual)
}

func Test_I24_CastedResult_IsNotMatchingAcceptedType(t *testing.T) {
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}
	actual := args.Map{"notMatch": cr.IsNotMatchingAcceptedType()}
	expected := args.Map{"notMatch": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsNotMatchingAcceptedType", actual)
}

func Test_I24_CastedResult_IsSourceKind(t *testing.T) {
	cr := &coredynamic.CastedResult{SourceKind: reflect.String}
	actual := args.Map{"isStr": cr.IsSourceKind(reflect.String), "isInt": cr.IsSourceKind(reflect.Int)}
	expected := args.Map{"isStr": true, "isInt": false}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- IsSourceKind", actual)
}

func Test_I24_CastedResult_HasAnyIssues_Valid(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}
	actual := args.Map{"issues": cr.HasAnyIssues()}
	expected := args.Map{"issues": false}
	expected.ShouldBeEqual(t, 0, "CastedResult returns non-empty -- HasAnyIssues valid", actual)
}

func Test_I24_CastedResult_HasAnyIssues_Invalid(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: false, IsNull: true, IsMatchingAcceptedType: false}
	actual := args.Map{"issues": cr.HasAnyIssues()}
	expected := args.Map{"issues": true}
	expected.ShouldBeEqual(t, 0, "CastedResult returns error -- HasAnyIssues invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItemDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_MapAnyItemDiff_Length(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Length", actual)
}

func Test_I24_MapAnyItemDiff_Length_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	actual := args.Map{"len": m.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Length nil", actual)
}

func Test_I24_MapAnyItemDiff_IsEmpty(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{}
	actual := args.Map{"empty": m.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- IsEmpty", actual)
}

func Test_I24_MapAnyItemDiff_HasAnyItem(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"has": m.HasAnyItem()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyItem", actual)
}

func Test_I24_MapAnyItemDiff_LastIndex(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	actual := args.Map{"idx": m.LastIndex()}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LastIndex", actual)
}

func Test_I24_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"b": 2, "a": 1}
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0], "second": keys[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- AllKeysSorted", actual)
}

func Test_I24_MapAnyItemDiff_Raw(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	raw := m.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Raw", actual)
}

func Test_I24_MapAnyItemDiff_Raw_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	raw := m.Raw()
	actual := args.Map{"len": len(raw)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Raw nil", actual)
}

func Test_I24_MapAnyItemDiff_Clear(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Clear", actual)
}

func Test_I24_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns nil -- Clear nil", actual)
}

func Test_I24_MapAnyItemDiff_IsRawEqual_Same(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 1})}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual same", actual)
}

func Test_I24_MapAnyItemDiff_IsRawEqual_Diff(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"eq": m.IsRawEqual(false, map[string]any{"a": 2})}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual diff", actual)
}

func Test_I24_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"changes": m.HasAnyChanges(false, map[string]any{"a": 2})}
	expected := args.Map{"changes": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyChanges", actual)
}

func Test_I24_MapAnyItemDiff_DiffRaw(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 99})
	actual := args.Map{"hasDiff": len(diff) > 0}
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffRaw", actual)
}

func Test_I24_MapAnyItemDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})
	actual := args.Map{"empty": diff.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- HashmapDiffUsingRaw no diff", actual)
}

func Test_I24_MapAnyItemDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 2})
	actual := args.Map{"hasDiff": diff.HasAnyItem()}
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HashmapDiffUsingRaw has diff", actual)
}

func Test_I24_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	mai := m.MapAnyItems()
	actual := args.Map{"notNil": mai != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- MapAnyItems", actual)
}

func Test_I24_MapAnyItemDiff_RawMapDiffer(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	d := m.RawMapDiffer()
	actual := args.Map{"notNil": d != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- RawMapDiffer", actual)
}

func Test_I24_MapAnyItemDiff_DiffJsonMessage(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- DiffJsonMessage", actual)
}

func Test_I24_MapAnyItemDiff_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ShouldDiffMessage", actual)
}

func Test_I24_MapAnyItemDiff_LogShouldDiffMessage(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LogShouldDiffMessage", actual)
}

func Test_I24_MapAnyItemDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	diff := m.DiffRaw(false, map[string]any{"a": 2})
	strs := m.ToStringsSliceOfDiffMap(diff)
	actual := args.Map{"hasItems": len(strs) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- ToStringsSliceOfDiffMap", actual)
}

func Test_I24_MapAnyItemDiff_Json(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	jr := m.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Json", actual)
}

func Test_I24_MapAnyItemDiff_JsonPtr(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	jr := m.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- JsonPtr", actual)
}

func Test_I24_MapAnyItemDiff_PrettyJsonString(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	s := m.PrettyJsonString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- PrettyJsonString", actual)
}

func Test_I24_MapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	m.LogPrettyJsonString() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- LogPrettyJsonString", actual)
}

func Test_I24_MapAnyItemDiff_LogPrettyJsonString_Empty(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{}
	m.LogPrettyJsonString() // empty path
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns empty -- LogPrettyJsonString empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_MapAsKeyValSlice_Valid(t *testing.T) {
	m := map[string]any{"a": 1, "b": 2}
	rv := reflect.ValueOf(m)
	kvc, err := coredynamic.MapAsKeyValSlice(rv)
	actual := args.Map{"noErr": err == nil, "len": kvc.Length()}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns non-empty -- valid", actual)
}

func Test_I24_MapAsKeyValSlice_NonMap(t *testing.T) {
	rv := reflect.ValueOf("hello")
	_, err := coredynamic.MapAsKeyValSlice(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns non-empty -- non-map", actual)
}

func Test_I24_MapAsKeyValSlice_Pointer(t *testing.T) {
	m := map[string]any{"a": 1}
	rv := reflect.ValueOf(&m)
	kvc, err := coredynamic.MapAsKeyValSlice(rv)
	actual := args.Map{"noErr": err == nil, "hasItems": kvc.Length() > 0}
	expected := args.Map{"noErr": false, "hasItems": false}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypeNotEqualErr + TypeMustBeSame + TypeSameStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_TypeNotEqualErr_Same(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", "b")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- same", actual)
}

func Test_I24_TypeNotEqualErr_Different(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", 1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- different", actual)
}

func Test_I24_TypeMustBeSame_NoPanic(t *testing.T) {
	coredynamic.TypeMustBeSame("a", "b")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- no panic", actual)
}

func Test_I24_TypeMustBeSame_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- panic", actual)
	}()
	coredynamic.TypeMustBeSame("a", 1)
}

func Test_I24_TypeSameStatus_Same(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	actual := args.Map{"same": ts.IsSame}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- same", actual)
}

func Test_I24_TypeSameStatus_Different(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 1)
	actual := args.Map{"same": ts.IsSame, "leftPtr": ts.IsLeftPointer}
	expected := args.Map{"same": false, "leftPtr": false}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- different", actual)
}

func Test_I24_TypeSameStatus_Pointer(t *testing.T) {
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "b")
	actual := args.Map{"leftPtr": ts.IsLeftPointer, "rightPtr": ts.IsRightPointer}
	expected := args.Map{"leftPtr": true, "rightPtr": false}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectKindValidation + ReflectTypeValidation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_ReflectKindValidation_Match(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.String, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns non-empty -- match", actual)
}

func Test_I24_ReflectKindValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.Int, "hello")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns non-empty -- mismatch", actual)
}

func Test_I24_ReflectTypeValidation_Match(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- match", actual)
}

func Test_I24_ReflectTypeValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- mismatch", actual)
}

func Test_I24_ReflectTypeValidation_NilNotAllowed(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil not allowed", actual)
}

func Test_I24_ReflectTypeValidation_NilAllowed(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- nil allowed but mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ZeroSet, ZeroSetAny, SafeZeroSet
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_ZeroSet(t *testing.T) {
	val := 42
	rv := reflect.ValueOf(&val)
	coredynamic.ZeroSet(rv)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSet returns correct value -- with args", actual)
}

func Test_I24_ZeroSetAny(t *testing.T) {
	val := "hello"
	coredynamic.ZeroSetAny(&val)
	actual := args.Map{"val": val}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns correct value -- with args", actual)
}

func Test_I24_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns nil -- nil", actual)
}

func Test_I24_SafeZeroSet(t *testing.T) {
	val := 42
	rv := reflect.ValueOf(&val)
	coredynamic.SafeZeroSet(rv)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns correct value -- with args", actual)
}

func Test_I24_SafeZeroSet_NilType(t *testing.T) {
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()

		coredynamic.SafeZeroSet(reflect.ValueOf(nil))
	}()

	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- nil reflect value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr + MustBeAcceptedTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_NotAcceptedTypesErr_Accepted(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- accepted", actual)
}

func Test_I24_NotAcceptedTypesErr_NotAccepted(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- not accepted", actual)
}

func Test_I24_MustBeAcceptedTypes_NoPanic(t *testing.T) {
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(""))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes panics -- no panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PointerOrNonPointer + UsingReflectValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_PointerOrNonPointer_NonPointer(t *testing.T) {
	val, _ := coredynamic.PointerOrNonPointer(false, "hello")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns non-empty -- non-ptr", actual)
}

func Test_I24_PointerOrNonPointer_PointerInput_NonPointerOutput(t *testing.T) {
	s := "hello"
	val, _ := coredynamic.PointerOrNonPointer(false, &s)
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns non-empty -- ptr->non-ptr", actual)
}

func Test_I24_PointerOrNonPointerUsingReflectValue_NonPtr(t *testing.T) {
	rv := reflect.ValueOf("hello")
	val, _ := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue returns non-empty -- non-ptr", actual)
}

func Test_I24_PointerOrNonPointerUsingReflectValue_PtrInput_NonPtrOut(t *testing.T) {
	s := "hello"
	rv := reflect.ValueOf(&s)
	val, _ := coredynamic.PointerOrNonPointerUsingReflectValue(false, rv)
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointerUsingReflectValue returns non-empty -- ptr->non-ptr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_LengthOfReflect_Slice(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- slice", actual)
}

func Test_I24_LengthOfReflect_Array(t *testing.T) {
	rv := reflect.ValueOf([3]int{1, 2, 3})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- array", actual)
}

func Test_I24_LengthOfReflect_Map(t *testing.T) {
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- map", actual)
}

func Test_I24_LengthOfReflect_Other(t *testing.T) {
	rv := reflect.ValueOf("hello")
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- other", actual)
}

func Test_I24_LengthOfReflect_Pointer(t *testing.T) {
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsAnyTypesOf + TypesIndexOf
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_IsAnyTypesOf_Found(t *testing.T) {
	actual := args.Map{"found": coredynamic.IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(0), reflect.TypeOf(""))}
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- found", actual)
}

func Test_I24_IsAnyTypesOf_NotFound(t *testing.T) {
	actual := args.Map{"found": coredynamic.IsAnyTypesOf(reflect.TypeOf(""), reflect.TypeOf(0))}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- not found", actual)
}

func Test_I24_TypesIndexOf_Found(t *testing.T) {
	idx := coredynamic.TypesIndexOf(reflect.TypeOf(""), reflect.TypeOf(0), reflect.TypeOf(""))
	actual := args.Map{"idx": idx}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- found", actual)
}

func Test_I24_TypesIndexOf_NotFound(t *testing.T) {
	idx := coredynamic.TypesIndexOf(reflect.TypeOf(""), reflect.TypeOf(0))
	actual := args.Map{"idx": idx}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- not found", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeTypeName + ReflectInterfaceVal + ValueStatus + DynamicStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_I24_SafeTypeName_String(t *testing.T) {
	actual := args.Map{"name": coredynamic.SafeTypeName("hello")}
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- string", actual)
}

func Test_I24_SafeTypeName_Nil(t *testing.T) {
	actual := args.Map{"name": coredynamic.SafeTypeName(nil)}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns nil -- nil", actual)
}

func Test_I24_ReflectInterfaceVal_NonPointer(t *testing.T) {
	val := coredynamic.ReflectInterfaceVal("hello")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns non-empty -- non-ptr", actual)
}

func Test_I24_ReflectInterfaceVal_Pointer(t *testing.T) {
	s := "hello"
	val := coredynamic.ReflectInterfaceVal(&s)
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- pointer", actual)
}

func Test_I24_ValueStatus_Invalid(t *testing.T) {
	vs := coredynamic.InvalidValueStatus("bad")
	actual := args.Map{"valid": vs.IsValid, "msg": vs.Message}
	expected := args.Map{"valid": false, "msg": "bad"}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns error -- invalid", actual)
}

func Test_I24_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := coredynamic.InvalidValueStatusNoMessage()
	actual := args.Map{"valid": vs.IsValid, "msg": vs.Message}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns empty -- invalid no msg", actual)
}

func Test_I24_DynamicStatus_Invalid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("bad")
	actual := args.Map{"valid": ds.IsValid(), "msg": ds.Message}
	expected := args.Map{"valid": false, "msg": "bad"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns error -- invalid", actual)
}

func Test_I24_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	actual := args.Map{"valid": ds.IsValid(), "msg": ds.Message}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns empty -- invalid no msg", actual)
}

func Test_I24_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.DynamicStatus{
		Dynamic: coredynamic.NewDynamic("hello", true),
		Message: "test",
	}
	cloned := ds.Clone()
	actual := args.Map{"valid": cloned.IsValid(), "msg": cloned.Message}
	expected := args.Map{"valid": true, "msg": "test"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- Clone", actual)
}

func Test_I24_DynamicStatus_ClonePtr(t *testing.T) {
	ds := &coredynamic.DynamicStatus{
		Dynamic: coredynamic.NewDynamic("hello", true),
		Message: "test",
	}
	cloned := ds.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "msg": cloned.Message}
	expected := args.Map{"notNil": true, "msg": "test"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- ClonePtr", actual)
}

func Test_I24_DynamicStatus_ClonePtr_Nil(t *testing.T) {
	var ds *coredynamic.DynamicStatus
	actual := args.Map{"nil": ds.ClonePtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns nil -- ClonePtr nil", actual)
}

func Test_I24_MapAnyItemDiff_IsRawEqual_RegardlessType(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"eq": m.IsRawEqual(true, map[string]any{"a": 1})}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual regardless type", actual)
}
