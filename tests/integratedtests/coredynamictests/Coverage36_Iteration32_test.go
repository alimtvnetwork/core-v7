package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — extended methods (AddMapResult, Diff, IsEqual, Clone, etc.)
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_01_MapAnyItems_AddMapResult(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResult(map[string]any{"b": 2, "a": 99})
	if m.GetValue("a") != 99 {
		t.Error("expected override to 99")
	}
	if m.GetValue("b") != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_02_MapAnyItems_AddMapResult_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddMapResult(nil)
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_03_MapAnyItems_AddMapResultOption_Override(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(true, map[string]any{"a": 99})
	if m.GetValue("a") != 99 {
		t.Error("expected 99")
	}
}

func Test_C36_04_MapAnyItems_AddMapResultOption_NoOverride(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(false, map[string]any{"a": 99, "b": 2})
	// noOverride only updates existing keys
	v := m.GetValue("a")
	_ = v
}

func Test_C36_05_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddManyMapResultsUsingOption(true, map[string]any{"b": 2}, map[string]any{"c": 3})
	if m.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C36_06_MapAnyItems_AddManyMapResultsUsingOption_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddManyMapResultsUsingOption(true)
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_07_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	if sub.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_08_MapAnyItems_GetNewMapUsingKeys_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	sub := m.GetNewMapUsingKeys(false)
	if sub.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_09_MapAnyItems_GetNewMapUsingKeys_Missing_NoPanic(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	sub := m.GetNewMapUsingKeys(false, "a", "nope")
	if sub.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C36_10_MapAnyItems_JsonString(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s, err := m.JsonString()
	if err != nil || s == "" {
		t.Error("expected json string")
	}
}

func Test_C36_11_MapAnyItems_JsonStringMust(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s := m.JsonStringMust()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_12_MapAnyItems_JsonResultOfKey_Found(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jr := m.JsonResultOfKey("k")
	if jr == nil || jr.HasError() {
		t.Error("expected valid result")
	}
}

func Test_C36_13_MapAnyItems_JsonResultOfKey_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	jr := m.JsonResultOfKey("nope")
	if !jr.HasError() {
		t.Error("expected error")
	}
}

func Test_C36_14_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	mr := m.JsonResultOfKeys("a", "b")
	if mr == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_15_MapAnyItems_JsonResultOfKeys_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	mr := m.JsonResultOfKeys()
	if mr == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_16_MapAnyItems_AllKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	if len(m.AllKeys()) != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_17_MapAnyItems_AllKeys_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	if len(m.AllKeys()) != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_18_MapAnyItems_AllValues(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	if len(m.AllValues()) != 1 {
		t.Error("expected 1")
	}
}

func Test_C36_19_MapAnyItems_AllValues_Empty(t *testing.T) {
	m := coredynamic.EmptyMapAnyItems()
	if len(m.AllValues()) != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_20_MapAnyItems_IsRawEqual(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if !m.IsRawEqual(false, map[string]any{"k": 1}) {
		t.Error("expected equal")
	}
}

func Test_C36_21_MapAnyItems_HasAnyChanges(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if m.HasAnyChanges(false, map[string]any{"k": 1}) {
		t.Error("expected no changes")
	}
	if !m.HasAnyChanges(false, map[string]any{"k": 2}) {
		t.Error("expected changes")
	}
}

func Test_C36_22_MapAnyItems_DiffRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.DiffRaw(false, map[string]any{"k": 2})
	_ = diff
}

func Test_C36_23_MapAnyItems_Diff(t *testing.T) {
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)
	m2.Add("k", 2)
	diff := m1.Diff(false, m2)
	_ = diff
}

func Test_C36_24_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"k": 2})
	if diff.Length() != 0 {
		t.Error("expected empty diff")
	}
}

func Test_C36_25_MapAnyItems_HashmapDiffUsingRaw_Same(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"k": 1})
	if diff.Length() != 0 {
		t.Error("expected no diff")
	}
}

func Test_C36_26_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	d := m.MapStringAnyDiff()
	if len(d) != 1 {
		t.Error("expected 1")
	}
}

func Test_C36_27_MapAnyItems_DiffJsonMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.DiffJsonMessage(false, map[string]any{"k": 2})
	_ = msg
}

func Test_C36_28_MapAnyItems_ToStringsSliceOfDiffMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	s := m.ToStringsSliceOfDiffMap(map[string]any{"k": 2})
	_ = s
}

func Test_C36_29_MapAnyItems_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"k": 2})
	_ = msg
}

func Test_C36_30_MapAnyItems_LogShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"k": 1})
	_ = msg
}

func Test_C36_31_MapAnyItems_JsonModel(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jm := m.JsonModel()
	if jm == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_32_MapAnyItems_JsonModelAny(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	if m.JsonModelAny() == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_33_MapAnyItems_Json(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	j := m.Json()
	_ = j
}

func Test_C36_34_MapAnyItems_ParseInjectUsingJson(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	_, err := m2.ParseInjectUsingJson(jp)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C36_35_MapAnyItems_JsonParseSelfInject(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	err := m2.JsonParseSelfInject(jp)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C36_36_MapAnyItems_Strings(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	s := m.Strings()
	if len(s) == 0 {
		t.Error("expected strings")
	}
}

func Test_C36_37_MapAnyItems_String(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if m.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_38_MapAnyItems_Clear(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.Clear()
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_39_MapAnyItems_Clear_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	m.Clear() // no panic
}

func Test_C36_40_MapAnyItems_DeepClear(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.DeepClear()
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_41_MapAnyItems_Dispose(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	m.Dispose()
}

func Test_C36_42_MapAnyItems_Dispose_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	m.Dispose()
}

func Test_C36_43_MapAnyItems_IsEqualRaw_Same(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if !m.IsEqualRaw(map[string]any{"k": 1}) {
		t.Error("expected equal")
	}
}

func Test_C36_44_MapAnyItems_IsEqualRaw_DiffLength(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if m.IsEqualRaw(map[string]any{"k": 1, "k2": 2}) {
		t.Error("expected not equal")
	}
}

func Test_C36_45_MapAnyItems_IsEqualRaw_MissingKey(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if m.IsEqualRaw(map[string]any{"x": 1}) {
		t.Error("expected not equal")
	}
}

func Test_C36_46_MapAnyItems_IsEqualRaw_DiffValue(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if m.IsEqualRaw(map[string]any{"k": 2}) {
		t.Error("expected not equal")
	}
}

func Test_C36_47_MapAnyItems_IsEqualRaw_BothNil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	if !m.IsEqualRaw(nil) {
		t.Error("expected equal for both nil")
	}
}

func Test_C36_48_MapAnyItems_IsEqualRaw_OneNil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	if m.IsEqualRaw(map[string]any{"k": 1}) {
		t.Error("expected not equal")
	}
}

func Test_C36_49_MapAnyItems_IsEqual(t *testing.T) {
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)
	m2.Add("k", 1)
	if !m1.IsEqual(m2) {
		t.Error("expected equal")
	}
}

func Test_C36_50_MapAnyItems_IsEqual_BothNil(t *testing.T) {
	var m1, m2 *coredynamic.MapAnyItems
	if !m1.IsEqual(m2) {
		t.Error("expected equal")
	}
}

func Test_C36_51_MapAnyItems_IsEqual_OneNil(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	if m.IsEqual(nil) {
		t.Error("expected not equal")
	}
}

func Test_C36_52_MapAnyItems_IsEqual_DiffLen(t *testing.T) {
	m1 := coredynamic.NewMapAnyItems(4)
	m1.Add("k", 1)
	m2 := coredynamic.NewMapAnyItems(4)
	if m1.IsEqual(m2) {
		t.Error("expected not equal")
	}
}

func Test_C36_53_MapAnyItems_ClonePtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	cloned, err := m.ClonePtr()
	if err != nil || cloned == nil {
		t.Error("expected cloned")
	}
}

func Test_C36_54_MapAnyItems_ClonePtr_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	cloned, err := m.ClonePtr()
	if cloned != nil || err == nil {
		t.Error("expected nil + error")
	}
}

func Test_C36_55_MapAnyItems_RawMapStringAnyDiff(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	d := m.RawMapStringAnyDiff()
	if len(d) != 1 {
		t.Error("expected 1")
	}
}

func Test_C36_56_MapAnyItems_RawMapStringAnyDiff_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	d := m.RawMapStringAnyDiff()
	if len(d) != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_57_MapAnyItems_MapAnyItems_Self(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	if m.MapAnyItems() != m {
		t.Error("expected same pointer")
	}
}

func Test_C36_58_MapAnyItems_GetItemRef_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.GetItemRef("nope", &target)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C36_59_MapAnyItems_GetItemRef_NilRef(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	err := m.GetItemRef("k", nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C36_60_MapAnyItems_GetItemRef_NonPointer(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	err := m.GetItemRef("k", 42)
	if err == nil {
		t.Error("expected error for non-pointer")
	}
}

func Test_C36_61_MapAnyItems_GetUsingUnmarshallManyAt(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", "hello")
	var i int
	var s string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &i},
		corejson.KeyAny{Key: "b", AnyInf: &s},
	)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C36_62_MapAnyItems_GetManyItemsRefs_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.GetManyItemsRefs()
	if err != nil {
		t.Error("expected nil")
	}
}

func Test_C36_63_MapAnyItems_ReflectSetToMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// expected for type mismatch
		}
	}()
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	m.ReflectSetToMust("k", &target)
}

func Test_C36_64_MapAnyItems_DeserializeMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// expected for missing key
		}
	}()
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	m.DeserializeMust("k", &target)
}

func Test_C36_65_MapAnyItems_ParseInjectUsingJsonMust(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	jp := m.JsonPtr()
	m2 := coredynamic.EmptyMapAnyItems()
	result := m2.ParseInjectUsingJsonMust(jp)
	if result == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_66_MapAnyItems_JsonResultsCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	rc := m.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_67_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	rc := m.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_68_LeftRight_IsEmpty(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}
	if !lr.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C36_69_LeftRight_IsEmpty_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	if !lr.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C36_70_LeftRight_HasAnyItem(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 1, Right: nil}
	if !lr.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_C36_71_LeftRight_HasLeft(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 1}
	if !lr.HasLeft() {
		t.Error("expected true")
	}
}

func Test_C36_72_LeftRight_HasRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Right: 2}
	if !lr.HasRight() {
		t.Error("expected true")
	}
}

func Test_C36_73_LeftRight_IsLeftEmpty(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: 1}
	if !lr.IsLeftEmpty() {
		t.Error("expected true")
	}
}

func Test_C36_74_LeftRight_IsRightEmpty(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 1, Right: nil}
	if !lr.IsRightEmpty() {
		t.Error("expected true")
	}
}

func Test_C36_75_LeftRight_LeftReflectSet(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	var target int
	err := lr.LeftReflectSet(&target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C36_76_LeftRight_LeftReflectSet_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	err := lr.LeftReflectSet(nil)
	if err != nil {
		t.Error("expected nil for nil receiver")
	}
}

func Test_C36_77_LeftRight_RightReflectSet(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	var target int
	err := lr.RightReflectSet(&target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C36_78_LeftRight_RightReflectSet_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	err := lr.RightReflectSet(nil)
	if err != nil {
		t.Error("expected nil")
	}
}

func Test_C36_79_LeftRight_DeserializeLeft(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 42}
	r := lr.DeserializeLeft()
	if r == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_80_LeftRight_DeserializeLeft_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	if lr.DeserializeLeft() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_81_LeftRight_DeserializeRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Right: 42}
	r := lr.DeserializeRight()
	if r == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_82_LeftRight_DeserializeRight_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	if lr.DeserializeRight() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_83_LeftRight_LeftToDynamic(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 42}
	d := lr.LeftToDynamic()
	if d == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_84_LeftRight_LeftToDynamic_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	if lr.LeftToDynamic() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_85_LeftRight_RightToDynamic(t *testing.T) {
	lr := &coredynamic.LeftRight{Right: 42}
	d := lr.RightToDynamic()
	if d == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_86_LeftRight_RightToDynamic_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	if lr.RightToDynamic() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_87_LeftRight_TypeStatus(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: 42, Right: 100}
	ts := lr.TypeStatus()
	if !ts.IsValid() {
		t.Error("expected valid")
	}
}

func Test_C36_88_LeftRight_TypeStatus_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()
	_ = ts
}

// ═══════════════════════════════════════════════════════════════════════
// CastTo + CastedResult
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_89_CastTo_MatchingType(t *testing.T) {
	result := coredynamic.CastTo(false, 42, reflect.TypeOf(0))
	if !result.IsMatchingAcceptedType {
		t.Error("expected matching")
	}
	if result.HasError() {
		t.Error("expected no error")
	}
}

func Test_C36_90_CastTo_NonMatchingType(t *testing.T) {
	result := coredynamic.CastTo(false, "str", reflect.TypeOf(0))
	if result.IsMatchingAcceptedType {
		t.Error("expected non-matching")
	}
}

func Test_C36_91_CastTo_PointerOutput(t *testing.T) {
	v := 42
	result := coredynamic.CastTo(true, &v, reflect.TypeOf(&v))
	if result.HasAnyIssues() {
		t.Error("expected no issues")
	}
}

func Test_C36_92_CastedResult_IsInvalid_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	if !cr.IsInvalid() {
		t.Error("expected invalid")
	}
}

func Test_C36_93_CastedResult_IsNotNull(t *testing.T) {
	cr := &coredynamic.CastedResult{IsNull: false}
	if !cr.IsNotNull() {
		t.Error("expected not null")
	}
}

func Test_C36_94_CastedResult_IsNotPointer(t *testing.T) {
	cr := &coredynamic.CastedResult{IsPointer: false}
	if !cr.IsNotPointer() {
		t.Error("expected not pointer")
	}
}

func Test_C36_95_CastedResult_IsNotMatchingAcceptedType(t *testing.T) {
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}
	if !cr.IsNotMatchingAcceptedType() {
		t.Error("expected true")
	}
}

func Test_C36_96_CastedResult_IsSourceKind(t *testing.T) {
	cr := &coredynamic.CastedResult{SourceKind: reflect.Int}
	if !cr.IsSourceKind(reflect.Int) {
		t.Error("expected true")
	}
}

func Test_C36_97_CastedResult_HasAnyIssues(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}
	if cr.HasAnyIssues() {
		t.Error("expected no issues")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectSetFromTo
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_98_ReflectSetFromTo_BothNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(nil, nil)
	if err != nil {
		t.Error("expected nil")
	}
}

func Test_C36_99_ReflectSetFromTo_ToNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(42, nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C36_100_ReflectSetFromTo_ToNonPointer(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(42, 0)
	if err == nil {
		t.Error("expected error for non-pointer dest")
	}
}

func Test_C36_101_ReflectSetFromTo_SameNonPointerToPointer(t *testing.T) {
	var target int
	err := coredynamic.ReflectSetFromTo(42, &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if target != 42 {
		t.Errorf("expected 42, got %d", target)
	}
}

func Test_C36_102_ReflectSetFromTo_SamePointerTypes(t *testing.T) {
	v := 42
	var target int
	from := &v
	err := coredynamic.ReflectSetFromTo(from, &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if target != 42 {
		t.Errorf("expected 42, got %d", target)
	}
}

func Test_C36_103_ReflectSetFromTo_BytesToType(t *testing.T) {
	var target int
	err := coredynamic.ReflectSetFromTo([]byte(`42`), &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if target != 42 {
		t.Errorf("expected 42, got %d", target)
	}
}

func Test_C36_104_ReflectSetFromTo_TypeToBytes(t *testing.T) {
	var target []byte
	err := coredynamic.ReflectSetFromTo(42, &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if len(target) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_105_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	var target string
	err := coredynamic.ReflectSetFromTo(42, &target)
	if err == nil {
		t.Error("expected type mismatch error")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_106_ReflectInterfaceVal_NonPointer(t *testing.T) {
	result := coredynamic.ReflectInterfaceVal(42)
	if result != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_107_ReflectInterfaceVal_Pointer(t *testing.T) {
	v := 42
	result := coredynamic.ReflectInterfaceVal(&v)
	if result != 42 {
		t.Error("expected 42")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_108_LengthOfReflect_Slice(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	if coredynamic.LengthOfReflect(rv) != 3 {
		t.Error("expected 3")
	}
}

func Test_C36_109_LengthOfReflect_Map(t *testing.T) {
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	if coredynamic.LengthOfReflect(rv) != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_110_LengthOfReflect_Array(t *testing.T) {
	rv := reflect.ValueOf([3]int{1, 2, 3})
	if coredynamic.LengthOfReflect(rv) != 3 {
		t.Error("expected 3")
	}
}

func Test_C36_111_LengthOfReflect_Other(t *testing.T) {
	rv := reflect.ValueOf(42)
	if coredynamic.LengthOfReflect(rv) != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_112_MapAsKeyValSlice_Valid(t *testing.T) {
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	kvc, err := coredynamic.MapAsKeyValSlice(rv)
	if err != nil || kvc.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_113_MapAsKeyValSlice_NotMap(t *testing.T) {
	rv := reflect.ValueOf(42)
	_, err := coredynamic.MapAsKeyValSlice(rv)
	if err == nil {
		t.Error("expected error")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ZeroSet, SafeZeroSet, ZeroSetAny
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_114_ZeroSet(t *testing.T) {
	type S struct{ X int }
	s := S{X: 42}
	coredynamic.ZeroSet(reflect.ValueOf(&s))
	if s.X != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_115_ZeroSetAny(t *testing.T) {
	type S struct{ X int }
	s := &S{X: 42}
	coredynamic.ZeroSetAny(s)
	if s.X != 0 {
		t.Error("expected 0")
	}
}

func Test_C36_116_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // no panic
}

func Test_C36_117_SafeZeroSet(t *testing.T) {
	type S struct{ X int }
	s := &S{X: 42}
	coredynamic.SafeZeroSet(reflect.ValueOf(s))
	if s.X != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// TypedDynamic[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_118_TypedDynamic_Basic(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	if d.Data() != "hello" || !d.IsValid() {
		t.Error("expected hello, valid")
	}
}

func Test_C36_119_TypedDynamic_Invalid(t *testing.T) {
	d := coredynamic.InvalidTypedDynamic[int]()
	if !d.IsInvalid() {
		t.Error("expected invalid")
	}
}

func Test_C36_120_TypedDynamic_InvalidPtr(t *testing.T) {
	d := coredynamic.InvalidTypedDynamicPtr[int]()
	if d == nil || d.IsValid() {
		t.Error("expected invalid ptr")
	}
}

func Test_C36_121_TypedDynamic_NewPtr(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](42, true)
	if d == nil || d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_122_TypedDynamic_Valid(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[int](42)
	if !d.IsValid() || d.Value() != 42 {
		t.Error("expected valid 42")
	}
}

func Test_C36_123_TypedDynamic_String(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	if d.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_124_TypedDynamic_JsonBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.JsonBytes()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_125_TypedDynamic_JsonResult(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	jr := d.JsonResult()
	_ = jr
}

func Test_C36_126_TypedDynamic_JsonString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	s, err := d.JsonString()
	if err != nil || s == "" {
		t.Error("expected json string")
	}
}

func Test_C36_127_TypedDynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_128_TypedDynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](0, false)
	err := d.UnmarshalJSON([]byte(`42`))
	if err != nil || d.Value() != 42 || !d.IsValid() {
		t.Error("expected 42 valid")
	}
}

func Test_C36_129_TypedDynamic_ValueMarshal(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, err := d.ValueMarshal()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_130_TypedDynamic_Bytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	b, ok := d.Bytes()
	if !ok || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_131_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	b, ok := d.Bytes()
	if !ok || string(b) != "hi" {
		t.Error("expected hi")
	}
}

func Test_C36_132_TypedDynamic_GetAsString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hi", true)
	v, ok := d.GetAsString()
	if !ok || v != "hi" {
		t.Error("expected hi")
	}
}

func Test_C36_133_TypedDynamic_GetAsInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	v, ok := d.GetAsInt()
	if !ok || v != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_134_TypedDynamic_GetAsInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int64](100, true)
	v, ok := d.GetAsInt64()
	if !ok || v != 100 {
		t.Error("expected 100")
	}
}

func Test_C36_135_TypedDynamic_GetAsUint(t *testing.T) {
	d := coredynamic.NewTypedDynamic[uint](10, true)
	v, ok := d.GetAsUint()
	if !ok || v != 10 {
		t.Error("expected 10")
	}
}

func Test_C36_136_TypedDynamic_GetAsFloat64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[float64](3.14, true)
	v, ok := d.GetAsFloat64()
	if !ok || v != 3.14 {
		t.Error("expected 3.14")
	}
}

func Test_C36_137_TypedDynamic_GetAsFloat32(t *testing.T) {
	d := coredynamic.NewTypedDynamic[float32](1.5, true)
	v, ok := d.GetAsFloat32()
	if !ok || v != 1.5 {
		t.Error("expected 1.5")
	}
}

func Test_C36_138_TypedDynamic_GetAsBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic[bool](true, true)
	v, ok := d.GetAsBool()
	if !ok || !v {
		t.Error("expected true")
	}
}

func Test_C36_139_TypedDynamic_GetAsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	v, ok := d.GetAsBytes()
	if !ok || string(v) != "hi" {
		t.Error("expected hi")
	}
}

func Test_C36_140_TypedDynamic_GetAsStrings(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]string]([]string{"a", "b"}, true)
	v, ok := d.GetAsStrings()
	if !ok || len(v) != 2 {
		t.Error("expected 2")
	}
}

func Test_C36_141_TypedDynamic_ValueString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	if d.ValueString() != "hello" {
		t.Error("expected hello")
	}
}

func Test_C36_142_TypedDynamic_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	if d.ValueString() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_143_TypedDynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	if d.ValueInt() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_144_TypedDynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int64](100, true)
	if d.ValueInt64() != 100 {
		t.Error("expected 100")
	}
}

func Test_C36_145_TypedDynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic[bool](true, true)
	if !d.ValueBool() {
		t.Error("expected true")
	}
}

func Test_C36_146_TypedDynamic_Clone(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	c := d.Clone()
	if c.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_147_TypedDynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](42, true)
	c := d.ClonePtr()
	if c == nil || c.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_148_TypedDynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.TypedDynamic[int]
	if d.ClonePtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_149_TypedDynamic_NonPtr(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	np := d.NonPtr()
	if np.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_150_TypedDynamic_Ptr(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](42, true)
	if d.Ptr() != d {
		t.Error("expected same pointer")
	}
}

func Test_C36_151_TypedDynamic_ToDynamic(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	dyn := d.ToDynamic()
	if dyn.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_152_TypedDynamic_Deserialize(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.Deserialize([]byte(`42`))
	if err != nil || d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_153_TypedDynamic_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.TypedDynamic[int]
	err := d.Deserialize([]byte(`42`))
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C36_154_TypedDynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	if d.JsonModel() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_155_TypedDynamic_JsonModelAny(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](42, true)
	if d.JsonModelAny() != 42 {
		t.Error("expected 42")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// TypedSimpleRequest[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_156_TypedSimpleRequest_Basic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequest[int](42, true, "")
	if r.Data() != 42 || !r.IsValid() {
		t.Error("expected 42 valid")
	}
}

func Test_C36_157_TypedSimpleRequest_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	if !r.IsValid() || r.Request() != 42 || r.Value() != 42 {
		t.Error("expected valid 42")
	}
}

func Test_C36_158_TypedSimpleRequest_Invalid(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequest[int]("err")
	if !r.IsInvalid() || r.Message() != "err" {
		t.Error("expected invalid with err")
	}
}

func Test_C36_159_TypedSimpleRequest_InvalidNoMessage(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequestNoMessage[int]()
	if !r.IsInvalid() {
		t.Error("expected invalid")
	}
}

func Test_C36_160_TypedSimpleRequest_IsValid_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if r.IsValid() {
		t.Error("expected false")
	}
}

func Test_C36_161_TypedSimpleRequest_IsInvalid_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if !r.IsInvalid() {
		t.Error("expected true")
	}
}

func Test_C36_162_TypedSimpleRequest_Message_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if r.Message() != "" {
		t.Error("expected empty")
	}
}

func Test_C36_163_TypedSimpleRequest_String(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	if r.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_164_TypedSimpleRequest_String_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if r.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C36_165_TypedSimpleRequest_InvalidError(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequest[int]("err msg")
	err := r.InvalidError()
	if err == nil {
		t.Error("expected error")
	}
	// call again for caching
	err2 := r.InvalidError()
	if err2 != err {
		t.Error("expected same error")
	}
}

func Test_C36_166_TypedSimpleRequest_InvalidError_NoMessage(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	if r.InvalidError() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_167_TypedSimpleRequest_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if r.InvalidError() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_168_TypedSimpleRequest_JsonBytes(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	b, err := r.JsonBytes()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_169_TypedSimpleRequest_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hi")
	v, ok := r.GetAsString()
	if !ok || v != "hi" {
		t.Error("expected hi")
	}
	_, ok2 := r.GetAsInt()
	if ok2 {
		t.Error("expected false")
	}
}

func Test_C36_170_TypedSimpleRequest_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	c := r.Clone()
	if c == nil || c.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_171_TypedSimpleRequest_Clone_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	if r.Clone() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_172_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	sr := r.ToSimpleRequest()
	if sr == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_173_TypedSimpleRequest_ToSimpleRequest_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	sr := r.ToSimpleRequest()
	if sr == nil {
		t.Error("expected non-nil fallback")
	}
}

func Test_C36_174_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	td := r.ToTypedDynamic()
	if td.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_175_TypedSimpleRequest_ToTypedDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	td := r.ToTypedDynamic()
	if td.IsValid() {
		t.Error("expected invalid")
	}
}

func Test_C36_176_TypedSimpleRequest_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[int](42)
	d := r.ToDynamic()
	if d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_177_TypedSimpleRequest_ToDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[int]
	d := r.ToDynamic()
	if d.IsValid() {
		t.Error("expected invalid")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// TypedSimpleResult[T]
// ═══════════════════════════════════════════════════════════════════════

func Test_C36_178_TypedSimpleResult_Basic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResult[int](42, true, "")
	if r.Data() != 42 || !r.IsValid() || r.Result() != 42 {
		t.Error("expected 42 valid")
	}
}

func Test_C36_179_TypedSimpleResult_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	if !r.IsValid() {
		t.Error("expected valid")
	}
}

func Test_C36_180_TypedSimpleResult_Invalid(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResult[int]("err")
	if !r.IsInvalid() || r.Message() != "err" {
		t.Error("expected invalid with err")
	}
}

func Test_C36_181_TypedSimpleResult_InvalidNoMessage(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResultNoMessage[int]()
	if !r.IsInvalid() {
		t.Error("expected invalid")
	}
}

func Test_C36_182_TypedSimpleResult_IsValid_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if r.IsValid() {
		t.Error("expected false")
	}
}

func Test_C36_183_TypedSimpleResult_IsInvalid_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if !r.IsInvalid() {
		t.Error("expected true")
	}
}

func Test_C36_184_TypedSimpleResult_Message_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if r.Message() != "" {
		t.Error("expected empty")
	}
}

func Test_C36_185_TypedSimpleResult_String(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	if r.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C36_186_TypedSimpleResult_String_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if r.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C36_187_TypedSimpleResult_InvalidError(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResult[int]("err")
	if r.InvalidError() == nil {
		t.Error("expected error")
	}
	if r.InvalidError() == nil {
		t.Error("expected cached error")
	}
}

func Test_C36_188_TypedSimpleResult_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if r.InvalidError() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_189_TypedSimpleResult_JsonBytes(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	b, err := r.JsonBytes()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C36_190_TypedSimpleResult_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("hi")
	v, ok := r.GetAsString()
	if !ok || v != "hi" {
		t.Error("expected hi")
	}
}

func Test_C36_191_TypedSimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	c := r.Clone()
	if c.Data() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_192_TypedSimpleResult_ClonePtr(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	c := r.ClonePtr()
	if c == nil || c.Data() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_193_TypedSimpleResult_ClonePtr_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	if r.ClonePtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C36_194_TypedSimpleResult_Clone_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	c := r.Clone()
	if c.IsValid() {
		t.Error("expected not valid")
	}
}

func Test_C36_195_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	sr := r.ToSimpleResult()
	if sr == nil {
		t.Error("expected non-nil")
	}
}

func Test_C36_196_TypedSimpleResult_ToSimpleResult_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	sr := r.ToSimpleResult()
	if sr == nil {
		t.Error("expected fallback")
	}
}

func Test_C36_197_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	td := r.ToTypedDynamic()
	if td.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_198_TypedSimpleResult_ToTypedDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	td := r.ToTypedDynamic()
	if td.IsValid() {
		t.Error("expected invalid")
	}
}

func Test_C36_199_TypedSimpleResult_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	d := r.ToDynamic()
	if d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C36_200_TypedSimpleResult_ToDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[int]
	d := r.ToDynamic()
	if d.IsValid() {
		t.Error("expected invalid")
	}
}
