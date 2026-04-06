package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════════════════════════════
// AnyCollection — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_01_AnyCollection_Basic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("a").Add("b").Add("c")
	if ac.Length() != 3 {
		t.Errorf("expected 3, got %d", ac.Length())
	}
	if ac.First() != "a" {
		t.Errorf("expected a")
	}
	if ac.Last() != "c" {
		t.Errorf("expected c")
	}
	if ac.At(1) != "b" {
		t.Errorf("expected b")
	}
}

func Test_C32_02_AnyCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	if !ac.IsEmpty() {
		t.Errorf("expected empty")
	}
	if ac.HasAnyItem() {
		t.Errorf("expected no items")
	}
	if ac.FirstOrDefault() != nil {
		t.Errorf("expected nil")
	}
	if ac.LastOrDefault() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_03_AnyCollection_NilReceiver(t *testing.T) {
	var ac *coredynamic.AnyCollection
	if ac.Length() != 0 {
		t.Errorf("expected 0")
	}
	if !ac.IsEmpty() {
		t.Errorf("expected true")
	}
}

func Test_C32_04_AnyCollection_FirstOrDefault_LastOrDefault(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	if ac.FirstOrDefault() != "x" {
		t.Errorf("expected x")
	}
	if ac.LastOrDefault() != "x" {
		t.Errorf("expected x")
	}
	if ac.FirstOrDefaultDynamic() == nil {
		t.Errorf("expected non-nil")
	}
	if ac.LastOrDefaultDynamic() == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_05_AnyCollection_FirstDynamic_LastDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")
	if ac.FirstDynamic() == nil {
		t.Errorf("expected non-nil")
	}
	if ac.LastDynamic() == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_06_AnyCollection_Skip_Take(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3).Add(4)
	if len(ac.Skip(2)) != 2 {
		t.Errorf("expected 2")
	}
	if len(ac.Take(2)) != 2 {
		t.Errorf("expected 2")
	}
	if len(ac.Limit(3)) != 3 {
		t.Errorf("expected 3")
	}
	if ac.SkipDynamic(1) == nil {
		t.Errorf("expected non-nil")
	}
	if ac.TakeDynamic(2) == nil {
		t.Errorf("expected non-nil")
	}
	if ac.LimitDynamic(2) == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_07_AnyCollection_SkipCollection_TakeCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3).Add(4)
	sc := ac.SkipCollection(2)
	if sc.Length() != 2 {
		t.Errorf("expected 2")
	}
	tc := ac.TakeCollection(2)
	if tc.Length() != 2 {
		t.Errorf("expected 2")
	}
	lc := ac.LimitCollection(3)
	if lc.Length() != 3 {
		t.Errorf("expected 3")
	}
	slc := ac.SafeLimitCollection(100)
	if slc.Length() != 4 {
		t.Errorf("expected 4")
	}
}

func Test_C32_08_AnyCollection_Count_LastIndex_HasIndex(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")
	if ac.Count() != 2 {
		t.Errorf("expected 2")
	}
	if ac.LastIndex() != 1 {
		t.Errorf("expected 1")
	}
	if !ac.HasIndex(1) {
		t.Errorf("expected true")
	}
	if ac.HasIndex(2) {
		t.Errorf("expected false")
	}
}

func Test_C32_09_AnyCollection_RemoveAt(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("a").Add("b").Add("c")
	if !ac.RemoveAt(1) {
		t.Errorf("expected success")
	}
	if ac.Length() != 2 {
		t.Errorf("expected 2")
	}
	if ac.RemoveAt(99) {
		t.Errorf("expected false for invalid")
	}
}

func Test_C32_10_AnyCollection_Items_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	items := ac.Items()
	if len(items) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_11_AnyCollection_Items_NonEmpty(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	if len(ac.Items()) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_12_AnyCollection_DynamicItems(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x").Add("y")
	di := ac.DynamicItems()
	if len(di) != 2 {
		t.Errorf("expected 2")
	}
	empty := coredynamic.EmptyAnyCollection()
	if len(empty.DynamicItems()) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_13_AnyCollection_DynamicCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	dc := ac.DynamicCollection()
	if dc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyAnyCollection()
	if empty.DynamicCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_14_AnyCollection_AtAsDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(42)
	d := ac.AtAsDynamic(0)
	if d.ValueInt() != 42 {
		t.Errorf("expected 42")
	}
}

func Test_C32_15_AnyCollection_ReflectSetAt(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("hello")
	var target string
	err := ac.ReflectSetAt(0, &target)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C32_16_AnyCollection_Loop_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})
	if count != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func Test_C32_17_AnyCollection_Loop_SyncBreak(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C32_18_AnyCollection_Loop_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	ac.Loop(true, func(i int, item any) bool {
		return false
	})
}

func Test_C32_19_AnyCollection_Loop_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Loop(false, func(i int, item any) bool {
		t.Errorf("should not be called")
		return false
	})
}

func Test_C32_20_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x").Add("y")
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return false
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C32_21_AnyCollection_LoopDynamic_SyncBreak(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return i == 0
	})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

func Test_C32_22_AnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.LoopDynamic(true, func(i int, item coredynamic.Dynamic) bool {
		return false
	})
}

func Test_C32_23_AnyCollection_LoopDynamic_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		t.Errorf("should not be called")
		return false
	})
}

func Test_C32_24_AnyCollection_AddAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.AddAny("val", true)
	if ac.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_25_AnyCollection_AddNonNull(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.AddNonNull(nil)
	if ac.Length() != 0 {
		t.Errorf("expected 0")
	}
	ac.AddNonNull("x")
	if ac.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_26_AnyCollection_AddNonNullDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.AddNonNullDynamic(nil, true)
	if ac.Length() != 0 {
		t.Errorf("expected 0")
	}
	ac.AddNonNullDynamic("x", true)
	if ac.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_27_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic("a", "b")
	if ac.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_28_AnyCollection_AddMany(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany("a", nil, "b")
	if ac.Length() != 2 {
		t.Errorf("expected 2 (nil skipped)")
	}
}

func Test_C32_29_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem([]string{"a", "b"})
	if ac.Length() < 1 {
		t.Errorf("expected items added")
	}
	ac.AddAnySliceFromSingleItem(nil)
}

func Test_C32_30_AnyCollection_ListStrings(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(42).Add("hello")
	strs := ac.ListStrings(false)
	if len(strs) != 2 {
		t.Errorf("expected 2")
	}
	strsWithName := ac.ListStrings(true)
	if len(strsWithName) != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_31_AnyCollection_Strings_String(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")
	strs := ac.Strings()
	if len(strs) != 2 {
		t.Errorf("expected 2")
	}
	s := ac.String()
	if s == "" {
		t.Errorf("expected non-empty")
	}
	empty := coredynamic.EmptyAnyCollection()
	if len(empty.Strings()) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_32_AnyCollection_JsonString(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	s, err := ac.JsonString()
	if err != nil || s == "" {
		t.Errorf("unexpected error or empty string")
	}
}

func Test_C32_33_AnyCollection_JsonStringMust(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	s := ac.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_34_AnyCollection_MarshalUnmarshalJSON(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("hello")
	b, err := ac.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Errorf("marshal failed")
	}
	ac2 := coredynamic.EmptyAnyCollection()
	err = ac2.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("unmarshal failed: %v", err)
	}
}

func Test_C32_35_AnyCollection_JsonModel_JsonModelAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	m := ac.JsonModel()
	if len(m) != 1 {
		t.Errorf("expected 1")
	}
	ma := ac.JsonModelAny()
	if ma == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_36_AnyCollection_Json_JsonPtr(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	j := ac.Json()
	_ = j
	jp := ac.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_37_AnyCollection_JsonResultsCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	rc := ac.JsonResultsCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyAnyCollection()
	if empty.JsonResultsCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_38_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	rc := ac.JsonResultsPtrCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyAnyCollection()
	if empty.JsonResultsPtrCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_39_AnyCollection_Paging(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	if ac.GetPagesSize(3) != 4 {
		t.Errorf("expected 4")
	}
	if ac.GetPagesSize(0) != 0 {
		t.Errorf("expected 0")
	}
	page := ac.GetSinglePageCollection(3, 2)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C32_40_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1).Add(2)
	page := ac.GetSinglePageCollection(5, 1)
	if page.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_41_AnyCollection_GetPagedCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(7)
	for i := 0; i < 7; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(3)
	if len(pages) != 3 {
		t.Errorf("expected 3 pages, got %d", len(pages))
	}
}

func Test_C32_42_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	pages := ac.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_43_AnyCollection_AddAnyItemsWithTypeValidation(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// type mismatch - stop on first error
	err = ac.AddAnyItemsWithTypeValidation(false, false, strType, 42)
	if err == nil {
		t.Errorf("expected type mismatch error")
	}
	// continue on error
	err = ac.AddAnyItemsWithTypeValidation(true, false, strType, "ok", 42)
	if err == nil {
		t.Errorf("expected error")
	}
	// empty items
	err = ac.AddAnyItemsWithTypeValidation(false, false, strType)
	if err != nil {
		t.Errorf("expected nil for empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_44_MapAnyItems_Basic(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	if m.Length() != 2 {
		t.Errorf("expected 2")
	}
	if m.IsEmpty() {
		t.Errorf("expected not empty")
	}
	if !m.HasAnyItem() {
		t.Errorf("expected has any")
	}
	if !m.HasKey("a") {
		t.Errorf("expected has key a")
	}
	if m.HasKey("z") {
		t.Errorf("expected false for z")
	}
}

func Test_C32_45_MapAnyItems_NilReceiver(t *testing.T) {
	var m *coredynamic.MapAnyItems
	if m.Length() != 0 {
		t.Errorf("expected 0")
	}
	if !m.IsEmpty() {
		t.Errorf("expected true")
	}
	if m.HasKey("x") {
		t.Errorf("expected false")
	}
}

func Test_C32_46_MapAnyItems_GetValue(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	if m.GetValue("a") != 1 {
		t.Errorf("expected 1")
	}
	if m.GetValue("z") != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_47_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "val"})
	v, has := m.Get("a")
	if !has || v != "val" {
		t.Errorf("expected val, has=true")
	}
	v2, has2 := m.Get("z")
	if has2 || v2 != nil {
		t.Errorf("expected nil, false")
	}
}

func Test_C32_48_MapAnyItems_EmptyItems(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(nil)
	if !m.IsEmpty() {
		t.Errorf("expected empty")
	}
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{})
	if !m2.IsEmpty() {
		t.Errorf("expected empty")
	}
}

func Test_C32_49_MapAnyItems_Add_Set(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.Add("x", 1)
	if !isNew {
		t.Errorf("expected newly added")
	}
	isNew2 := m.Add("x", 2)
	if isNew2 {
		t.Errorf("expected not new (override)")
	}
	isNew3 := m.Set("y", 3)
	if !isNew3 {
		t.Errorf("expected newly added")
	}
}

func Test_C32_50_MapAnyItems_AddKeyAny(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: "v"})
	if !isNew {
		t.Errorf("expected new")
	}
}

func Test_C32_51_MapAnyItems_AddKeyAnyWithValidation(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	strType := reflect.TypeOf("")
	err := m.AddKeyAnyWithValidation(strType, corejson.KeyAny{Key: "k", AnyInf: "v"})
	if err != nil {
		t.Errorf("unexpected error")
	}
	err = m.AddKeyAnyWithValidation(strType, corejson.KeyAny{Key: "k2", AnyInf: 42})
	if err == nil {
		t.Errorf("expected type mismatch error")
	}
}

func Test_C32_52_MapAnyItems_AddWithValidation(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	strType := reflect.TypeOf("")
	err := m.AddWithValidation(strType, "k", "v")
	if err != nil {
		t.Errorf("unexpected error")
	}
	err = m.AddWithValidation(strType, "k2", 42)
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_C32_53_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	jr := corejson.NewPtr("hello")
	m.AddJsonResultPtr("k", jr)
	if m.Length() != 1 {
		t.Errorf("expected 1")
	}
	m.AddJsonResultPtr("k2", nil) // should skip
	if m.Length() != 1 {
		t.Errorf("expected still 1")
	}
}

func Test_C32_54_MapAnyItems_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("b", 2)
	m.Add("a", 1)
	keys := m.AllKeys()
	if len(keys) != 2 {
		t.Errorf("expected 2")
	}
	sortedKeys := m.AllKeysSorted()
	if sortedKeys[0] != "a" || sortedKeys[1] != "b" {
		t.Errorf("expected sorted")
	}
	vals := m.AllValues()
	if len(vals) != 2 {
		t.Errorf("expected 2")
	}
	empty := coredynamic.EmptyMapAnyItems()
	if len(empty.AllKeys()) != 0 {
		t.Errorf("expected 0")
	}
	if len(empty.AllKeysSorted()) != 0 {
		t.Errorf("expected 0")
	}
	if len(empty.AllValues()) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_55_MapAnyItems_ReflectSetTo(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "hello")
	var target string
	err := m.ReflectSetTo("k", &target)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = m.ReflectSetTo("missing", &target)
	if err == nil {
		t.Errorf("expected error for missing key")
	}
}

func Test_C32_56_MapAnyItems_GetUsingUnmarshallAt(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("name", "John")
	var result string
	err := m.GetUsingUnmarshallAt("name", &result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = m.GetUsingUnmarshallAt("missing", &result)
	if err == nil {
		t.Errorf("expected error for missing key")
	}
}

func Test_C32_57_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("val", 42)
	var result int
	err := m.Deserialize("val", &result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C32_58_MapAnyItems_GetUsingUnmarshallManyAt(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", "hello")
	m.Add("b", 42)
	var s string
	var n int
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &s},
		corejson.KeyAny{Key: "b", AnyInf: &n},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C32_59_MapAnyItems_GetFieldsMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("obj", map[string]any{"x": 1})
	fm, err, found := m.GetFieldsMap("obj")
	if !found || err != nil {
		t.Errorf("expected found=true, err=nil, got found=%v err=%v", found, err)
	}
	_ = fm
	_, _, found2 := m.GetFieldsMap("missing")
	if found2 {
		t.Errorf("expected not found")
	}
}

func Test_C32_60_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("obj", map[string]any{"x": 1})
	fm, found := m.GetSafeFieldsMap("obj")
	if !found {
		t.Errorf("expected found")
	}
	_ = fm
}

func Test_C32_61_MapAnyItems_GetItemRef(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	val := "hello"
	m.Add("k", &val)
	var target string
	err := m.GetItemRef("k", &target)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// missing key
	err = m.GetItemRef("missing", &target)
	if err == nil {
		t.Errorf("expected error for missing key")
	}
	// nil referenceOut
	err = m.GetItemRef("k", nil)
	if err == nil {
		t.Errorf("expected error for nil ref")
	}
	// non-pointer referenceOut
	err = m.GetItemRef("k", "not-a-pointer")
	if err == nil {
		t.Errorf("expected error for non-pointer")
	}
}

func Test_C32_62_MapAnyItems_GetManyItemsRefs(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	val := "hello"
	m.Add("k", &val)
	var target string
	err := m.GetManyItemsRefs(
		corejson.KeyAny{Key: "k", AnyInf: &target},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// empty
	err = m.GetManyItemsRefs()
	if err != nil {
		t.Errorf("expected nil for empty")
	}
}

func Test_C32_63_MapAnyItems_AddMapResult(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddMapResult(map[string]any{"a": 1, "b": 2})
	if m.Length() != 2 {
		t.Errorf("expected 2")
	}
	m.AddMapResult(nil)
	if m.Length() != 2 {
		t.Errorf("expected still 2")
	}
}

func Test_C32_64_MapAnyItems_AddMapResultOption(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(true, map[string]any{"a": 99, "b": 2})
	if m.Length() != 2 {
		t.Errorf("expected 2")
	}
	m.AddMapResultOption(false, map[string]any{"a": 100})
	m.AddMapResultOption(false, nil)
}

func Test_C32_65_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddManyMapResultsUsingOption(true,
		map[string]any{"a": 1},
		map[string]any{"b": 2},
	)
	if m.Length() != 2 {
		t.Errorf("expected 2")
	}
	m.AddManyMapResultsUsingOption(true)
}

func Test_C32_66_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	sub := m.GetNewMapUsingKeys(false, "a", "c")
	if sub.Length() != 2 {
		t.Errorf("expected 2")
	}
	empty := m.GetNewMapUsingKeys(false)
	if empty.Length() != 0 {
		t.Errorf("expected 0")
	}
	// not panic on missing when isPanicOnMissing=false
	sub2 := m.GetNewMapUsingKeys(false, "a", "missing")
	if sub2.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_67_MapAnyItems_JsonString(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s, err := m.JsonString()
	if err != nil || s == "" {
		t.Errorf("expected json string")
	}
}

func Test_C32_68_MapAnyItems_JsonStringMust(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s := m.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_69_MapAnyItems_JsonResultOfKey(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	jr := m.JsonResultOfKey("k")
	if jr == nil || jr.HasError() {
		t.Errorf("expected valid result")
	}
	jr2 := m.JsonResultOfKey("missing")
	if jr2 == nil || !jr2.HasError() {
		t.Errorf("expected error for missing key")
	}
}

func Test_C32_70_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	mr := m.JsonResultOfKeys("a", "b")
	if mr == nil {
		t.Errorf("expected non-nil")
	}
	mr2 := m.JsonResultOfKeys()
	_ = mr2
}

func Test_C32_71_MapAnyItems_Paging(t *testing.T) {
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 10; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}
	if m.GetPagesSize(3) != 4 {
		t.Errorf("expected 4")
	}
	if m.GetPagesSize(0) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_72_MapAnyItems_GetPagedCollection_Small(t *testing.T) {
	m := coredynamic.NewMapAnyItems(2)
	m.Add("a", 1)
	pages := m.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_73_MapAnyItems_GetSinglePageCollection_Small(t *testing.T) {
	m := coredynamic.NewMapAnyItems(2)
	m.Add("a", 1)
	page := m.GetSinglePageCollection(5, 1, m.AllKeysSorted())
	if page.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_74_MapAnyItems_IsEqualRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	if !m.IsEqualRaw(map[string]any{"a": 1, "b": 2}) {
		t.Errorf("expected equal")
	}
	if m.IsEqualRaw(map[string]any{"a": 1, "b": 3}) {
		t.Errorf("expected not equal")
	}
	if m.IsEqualRaw(map[string]any{"a": 1}) {
		t.Errorf("expected not equal (different length)")
	}
	if m.IsEqualRaw(map[string]any{"a": 1, "c": 2}) {
		t.Errorf("expected not equal (missing key)")
	}
	var nilM *coredynamic.MapAnyItems
	if !nilM.IsEqualRaw(nil) {
		t.Errorf("expected true for both nil")
	}
	if nilM.IsEqualRaw(map[string]any{"a": 1}) {
		t.Errorf("expected false for nil vs non-nil")
	}
}

func Test_C32_75_MapAnyItems_IsEqual(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	if !m1.IsEqual(m2) {
		t.Errorf("expected equal")
	}
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 2})
	if m1.IsEqual(m3) {
		t.Errorf("expected not equal")
	}
	var nilM *coredynamic.MapAnyItems
	if !nilM.IsEqual(nil) {
		t.Errorf("expected true")
	}
	if nilM.IsEqual(m1) {
		t.Errorf("expected false")
	}
	m4 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	if m1.IsEqual(m4) {
		t.Errorf("expected not equal (different length)")
	}
}

func Test_C32_76_MapAnyItems_Clear_DeepClear_Dispose(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Clear()
	if m.Length() != 0 {
		t.Errorf("expected 0")
	}
	m.Add("b", 2)
	m.DeepClear()
	if m.Length() != 0 {
		t.Errorf("expected 0")
	}
	m.Dispose()
	var nilM *coredynamic.MapAnyItems
	nilM.Clear()
	nilM.DeepClear()
	nilM.Dispose()
}

func Test_C32_77_MapAnyItems_Strings_String(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	strs := m.Strings()
	if len(strs) == 0 {
		t.Errorf("expected non-empty")
	}
	s := m.String()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_78_MapAnyItems_MapAnyItems_Self(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	if m.MapAnyItems() != m {
		t.Errorf("expected self reference")
	}
}

func Test_C32_79_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	diff := m.MapStringAnyDiff()
	if diff == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_80_MapAnyItems_RawMapStringAnyDiff_NilReceiver(t *testing.T) {
	var m *coredynamic.MapAnyItems
	diff := m.RawMapStringAnyDiff()
	if len(diff) != 0 {
		t.Errorf("expected empty map for nil receiver")
	}
}

func Test_C32_81_MapAnyItems_IsRawEqual(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	if !m.IsRawEqual(false, map[string]any{"a": 1}) {
		t.Errorf("expected equal")
	}
}

func Test_C32_82_MapAnyItems_HasAnyChanges(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	if m.HasAnyChanges(false, map[string]any{"a": 1}) {
		t.Errorf("expected no changes")
	}
	if !m.HasAnyChanges(false, map[string]any{"a": 2}) {
		t.Errorf("expected changes")
	}
}

func Test_C32_83_MapAnyItems_Json_JsonPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	j := m.Json()
	_ = j
	jp := m.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_84_MapAnyItems_JsonModel_JsonModelAny(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	model := m.JsonModel()
	if model == nil {
		t.Errorf("expected non-nil")
	}
	if m.JsonModelAny() == nil {
		t.Errorf("expected non-nil")
	}
	empty := coredynamic.EmptyMapAnyItems()
	em := empty.JsonModel()
	_ = em
}

func Test_C32_85_MapAnyItems_JsonMapResults(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	mr, err := m.JsonMapResults()
	if err != nil || mr == nil {
		t.Errorf("unexpected error")
	}
	empty := coredynamic.EmptyMapAnyItems()
	mr2, _ := empty.JsonMapResults()
	_ = mr2
}

func Test_C32_86_MapAnyItems_JsonResultsCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	rc := m.JsonResultsCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyMapAnyItems()
	if empty.JsonResultsCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_87_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	rc := m.JsonResultsPtrCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyMapAnyItems()
	if empty.JsonResultsPtrCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_88_MapAnyItems_ClonePtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	cloned, err := m.ClonePtr()
	if err != nil || cloned == nil {
		t.Errorf("unexpected error or nil: %v", err)
	}
	var nilM *coredynamic.MapAnyItems
	_, err = nilM.ClonePtr()
	if err == nil {
		t.Errorf("expected error for nil receiver")
	}
}

func Test_C32_89_MapAnyItems_NewUsingAnyTypeMap(t *testing.T) {
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]int{"a": 1})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if m.Length() != 1 {
		t.Errorf("expected 1")
	}
	_, err = coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	if err == nil {
		t.Errorf("expected error for nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_90_ValueStatus(t *testing.T) {
	vs := coredynamic.InvalidValueStatus("test")
	if vs.IsValid {
		t.Errorf("expected invalid")
	}
	if vs.Message != "test" {
		t.Errorf("expected test")
	}
	vs2 := coredynamic.InvalidValueStatusNoMessage()
	if vs2.Message != "" {
		t.Errorf("expected empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// KeyVal — additional coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_91_KeyVal_KeyDynamic_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "name", Value: 42}
	kd := kv.KeyDynamic()
	if kd.ValueString() != "name" {
		t.Errorf("expected name")
	}
	vd := kv.ValueDynamic()
	if vd.ValueInt() != 42 {
		t.Errorf("expected 42")
	}
}

func Test_C32_92_KeyVal_KeyDynamicPtr_ValueDynamicPtr(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	if kv.KeyDynamicPtr() == nil {
		t.Errorf("expected non-nil")
	}
	if kv.ValueDynamicPtr() == nil {
		t.Errorf("expected non-nil")
	}
	var nilKV *coredynamic.KeyVal
	if nilKV.KeyDynamicPtr() != nil {
		t.Errorf("expected nil")
	}
	if nilKV.ValueDynamicPtr() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_93_KeyVal_IsKeyNull_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "a", Value: nil}
	if kv.IsKeyNull() {
		t.Errorf("expected false")
	}
	if !kv.IsValueNull() {
		t.Errorf("expected true")
	}
}

func Test_C32_94_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: nil}
	if !kv.IsKeyNullOrEmptyString() {
		t.Errorf("expected true")
	}
	kv2 := coredynamic.KeyVal{Key: "x", Value: nil}
	if kv2.IsKeyNullOrEmptyString() {
		t.Errorf("expected false")
	}
}

func Test_C32_95_KeyVal_ValueInt_ValueUInt_ValueBool_ValueInt64_ValueStrings(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	if kv.ValueInt() != 42 {
		t.Errorf("expected 42")
	}
	kv2 := coredynamic.KeyVal{Key: "k", Value: uint(10)}
	if kv2.ValueUInt() != 10 {
		t.Errorf("expected 10")
	}
	kv3 := coredynamic.KeyVal{Key: "k", Value: true}
	if !kv3.ValueBool() {
		t.Errorf("expected true")
	}
	kv4 := coredynamic.KeyVal{Key: "k", Value: int64(99)}
	if kv4.ValueInt64() != 99 {
		t.Errorf("expected 99")
	}
	kv5 := coredynamic.KeyVal{Key: "k", Value: []string{"a"}}
	if len(kv5.ValueStrings()) != 1 {
		t.Errorf("expected 1")
	}
	// mismatches
	kvBad := coredynamic.KeyVal{Key: "k", Value: "str"}
	if kvBad.ValueInt() != -1 {
		t.Errorf("expected invalid")
	}
	if kvBad.ValueUInt() != 0 {
		t.Errorf("expected 0")
	}
	if kvBad.ValueBool() {
		t.Errorf("expected false")
	}
	if kvBad.ValueInt64() != -1 {
		t.Errorf("expected invalid")
	}
	if kvBad.ValueStrings() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_96_KeyVal_String_NilReceiver(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.String() != "" {
		t.Errorf("expected empty")
	}
}

func Test_C32_97_KeyVal_String(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "name", Value: "val"}
	s := kv.String()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_98_KeyVal_KeyString_ValueString(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	if kv.KeyString() == "" {
		t.Errorf("expected non-empty")
	}
	if kv.ValueString() == "" {
		t.Errorf("expected non-empty")
	}
	var nilKV *coredynamic.KeyVal
	if nilKV.KeyString() != "" {
		t.Errorf("expected empty")
	}
	if nilKV.ValueString() != "" {
		t.Errorf("expected empty")
	}
	kvNilKey := &coredynamic.KeyVal{Key: nil, Value: nil}
	if kvNilKey.KeyString() != "" {
		t.Errorf("expected empty for nil key")
	}
	if kvNilKey.ValueString() != "" {
		t.Errorf("expected empty for nil value")
	}
}

func Test_C32_99_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	if rv.Int() != 42 {
		t.Errorf("expected 42")
	}
}

func Test_C32_100_KeyVal_ValueNullErr_KeyNullErr(t *testing.T) {
	var nilKV *coredynamic.KeyVal
	if nilKV.ValueNullErr() == nil {
		t.Errorf("expected error")
	}
	if nilKV.KeyNullErr() == nil {
		t.Errorf("expected error")
	}
	kvNullVal := &coredynamic.KeyVal{Key: "k", Value: nil}
	if kvNullVal.ValueNullErr() == nil {
		t.Errorf("expected error for null value")
	}
	kvNullKey := &coredynamic.KeyVal{Key: nil, Value: "v"}
	if kvNullKey.KeyNullErr() == nil {
		t.Errorf("expected error for null key")
	}
	kvOk := &coredynamic.KeyVal{Key: "k", Value: "v"}
	if kvOk.ValueNullErr() != nil {
		t.Errorf("expected nil")
	}
	if kvOk.KeyNullErr() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_101_KeyVal_CastKeyVal(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var nilKV *coredynamic.KeyVal
	if nilKV.CastKeyVal(&k, &v) == nil {
		t.Errorf("expected error")
	}
}

func Test_C32_102_KeyVal_ReflectSetKey(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var target string
	err := kv.ReflectSetKey(&target)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var nilKV *coredynamic.KeyVal
	if nilKV.ReflectSetKey(&target) == nil {
		t.Errorf("expected error")
	}
}

func Test_C32_103_KeyVal_KeyReflectSet_ValueReflectSet_ReflectSetTo(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var k, v string
	if kv.KeyReflectSet(&k) != nil {
		t.Errorf("unexpected error")
	}
	if kv.ValueReflectSet(&v) != nil {
		t.Errorf("unexpected error")
	}
	if kv.ReflectSetTo(&v) != nil {
		t.Errorf("unexpected error")
	}
	var nilKV *coredynamic.KeyVal
	if nilKV.KeyReflectSet(&k) == nil {
		t.Errorf("expected error")
	}
	if nilKV.ValueReflectSet(&v) == nil {
		t.Errorf("expected error")
	}
	if nilKV.ReflectSetTo(&v) == nil {
		t.Errorf("expected error")
	}
}

func Test_C32_104_KeyVal_ReflectSetToMust(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	var target string
	kv.ReflectSetToMust(&target)
	if target != "v" {
		t.Errorf("expected v")
	}
}

func Test_C32_105_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	_ = j
	jp := kv.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
	if kv.JsonModel() == nil {
		t.Errorf("expected non-nil")
	}
	if kv.JsonModelAny() == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_106_KeyVal_Serialize(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	if err != nil || len(b) == 0 {
		t.Errorf("serialize failed")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValCollection — additional coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_107_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	if kvc.Length() != 2 {
		t.Errorf("expected 2")
	}
	if kvc.IsEmpty() {
		t.Errorf("expected not empty")
	}
	if !kvc.HasAnyItem() {
		t.Errorf("expected has any")
	}
}

func Test_C32_108_KeyValCollection_AddPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kv := &coredynamic.KeyVal{Key: "a", Value: 1}
	kvc.AddPtr(kv)
	kvc.AddPtr(nil)
	if kvc.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_109_KeyValCollection_AddMany_AddManyPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)
	if kvc.Length() != 2 {
		t.Errorf("expected 2")
	}
	kv := &coredynamic.KeyVal{Key: "c", Value: 3}
	kvc.AddManyPtr(kv, nil)
	if kvc.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C32_110_KeyValCollection_Items_NilReceiver(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	if kvc.Items() != nil {
		t.Errorf("expected nil")
	}
	if kvc.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_111_KeyValCollection_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()
	if len(keys) != 2 {
		t.Errorf("expected 2")
	}
	sorted := kvc.AllKeysSorted()
	if sorted[0] != "a" {
		t.Errorf("expected sorted")
	}
	vals := kvc.AllValues()
	if len(vals) != 2 {
		t.Errorf("expected 2")
	}
	empty := coredynamic.EmptyKeyValCollection()
	if len(empty.AllKeys()) != 0 {
		t.Errorf("expected 0")
	}
	if len(empty.AllKeysSorted()) != 0 {
		t.Errorf("expected 0")
	}
	if len(empty.AllValues()) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_112_KeyValCollection_MapAnyItems(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	m := kvc.MapAnyItems()
	if m.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyKeyValCollection()
	if empty.MapAnyItems().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_113_KeyValCollection_String_NilReceiver(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	if kvc.String() != "" {
		t.Errorf("expected empty")
	}
}

func Test_C32_114_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_115_KeyValCollection_Clone_ClonePtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c := kvc.Clone()
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
	cp := kvc.ClonePtr()
	if cp.Length() != 1 {
		t.Errorf("expected 1")
	}
	var nilKVC *coredynamic.KeyValCollection
	if nilKVC.ClonePtr() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_116_KeyValCollection_NonPtr_Ptr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	np := kvc.NonPtr()
	_ = np
	p := kvc.Ptr()
	if p == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_117_KeyValCollection_JsonString(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := kvc.JsonString()
	if err != nil || s == "" {
		t.Errorf("json string failed")
	}
}

func Test_C32_118_KeyValCollection_JsonStringMust(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_119_KeyValCollection_Serialize(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	b, err := kvc.Serialize()
	if err != nil || len(b) == 0 {
		t.Errorf("serialize failed")
	}
}

func Test_C32_120_KeyValCollection_Paging(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	if kvc.GetPagesSize(3) != 4 {
		t.Errorf("expected 4")
	}
	if kvc.GetPagesSize(0) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_121_KeyValCollection_GetSinglePageCollection_Small(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	page := kvc.GetSinglePageCollection(5, 1)
	if page.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_122_KeyValCollection_GetPagedCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(7)
	for i := 0; i < 7; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)
	if len(pages) != 3 {
		t.Errorf("expected 3 pages")
	}
}

func Test_C32_123_KeyValCollection_GetPagedCollection_Small(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	pages := kvc.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_124_KeyValCollection_JsonMapResults(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	mr, err := kvc.JsonMapResults()
	if err != nil || mr == nil {
		t.Errorf("error: %v", err)
	}
	empty := coredynamic.EmptyKeyValCollection()
	mr2, _ := empty.JsonMapResults()
	_ = mr2
}

func Test_C32_125_KeyValCollection_JsonResultsCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyKeyValCollection()
	if empty.JsonResultsCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_126_KeyValCollection_JsonResultsPtrCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsPtrCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyKeyValCollection()
	if empty.JsonResultsPtrCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_127_KeyValCollection_Json_JsonPtr(t *testing.T) {
	kvc := coredynamic.KeyValCollection{}
	j := kvc.Json()
	_ = j
	jp := kvc.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
	if kvc.JsonModel() == nil {
		t.Errorf("expected non-nil")
	}
	if kvc.JsonModelAny() == nil {
		t.Errorf("expected non-nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// newCreator factories
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_128_NewCreator_Collection_String(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	if c.Length() != 0 {
		t.Errorf("expected 0")
	}
	c2 := coredynamic.New.Collection.String.Cap(5)
	if c2.Capacity() < 5 {
		t.Errorf("expected cap >= 5")
	}
	c3 := coredynamic.New.Collection.String.From([]string{"a", "b"})
	if c3.Length() != 2 {
		t.Errorf("expected 2")
	}
	c4 := coredynamic.New.Collection.String.Clone([]string{"x"})
	if c4.Length() != 1 {
		t.Errorf("expected 1")
	}
	c5 := coredynamic.New.Collection.String.Items("a", "b", "c")
	if c5.Length() != 3 {
		t.Errorf("expected 3")
	}
	c6 := coredynamic.New.Collection.String.Create([]string{"x"})
	if c6.Length() != 1 {
		t.Errorf("expected 1")
	}
	c7 := coredynamic.New.Collection.String.LenCap(3, 10)
	if c7.Length() != 3 || c7.Capacity() < 10 {
		t.Errorf("expected len=3, cap>=10")
	}
}

func Test_C32_129_NewCreator_Collection_Int(t *testing.T) {
	c := coredynamic.New.Collection.Int.Empty()
	c.Add(1)
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
	c2 := coredynamic.New.Collection.Int.LenCap(2, 5)
	if c2.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_130_NewCreator_Collection_Int64(t *testing.T) {
	c := coredynamic.New.Collection.Int64.Cap(5)
	c.Add(int64(1))
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
	c2 := coredynamic.New.Collection.Int64.LenCap(2, 5)
	if c2.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_131_NewCreator_Collection_Byte(t *testing.T) {
	c := coredynamic.New.Collection.Byte.Cap(5)
	c.Add(byte(1))
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
	c2 := coredynamic.New.Collection.Byte.LenCap(2, 5)
	if c2.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_132_NewCreator_Collection_Any(t *testing.T) {
	c := coredynamic.New.Collection.Any.Empty()
	c.Add(42)
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
	c2 := coredynamic.New.Collection.Any.Items("a", 1, true)
	if c2.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C32_133_NewCreator_Collection_Others(t *testing.T) {
	bs := coredynamic.New.Collection.ByteSlice.Empty()
	bs.Add([]byte{1})
	if bs.Length() != 1 {
		t.Errorf("expected 1")
	}
	b := coredynamic.New.Collection.Bool.Cap(5)
	b.Add(true)
	if b.Length() != 1 {
		t.Errorf("expected 1")
	}
	f32 := coredynamic.New.Collection.Float32.Empty()
	f32.Add(1.0)
	f64 := coredynamic.New.Collection.Float64.Empty()
	f64.Add(2.0)
	am := coredynamic.New.Collection.AnyMap.Empty()
	am.Add(map[string]any{"a": 1})
	sm := coredynamic.New.Collection.StringMap.Empty()
	sm.Add(map[string]string{"a": "b"})
	im := coredynamic.New.Collection.IntMap.Empty()
	im.Add(map[string]int{"a": 1})
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection — additional methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_134_DynCol_AddAnyItemsWithTypeValidation(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	err = dc.AddAnyItemsWithTypeValidation(false, false, strType, 42)
	if err == nil {
		t.Errorf("expected error")
	}
	err = dc.AddAnyItemsWithTypeValidation(true, false, strType, "ok", 42)
	if err == nil {
		t.Errorf("expected error")
	}
	err = dc.AddAnyItemsWithTypeValidation(false, false, strType)
	if err != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_135_DynCol_AddAnySliceFromSingleItem(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, []string{"a", "b"})
	if dc.Length() < 1 {
		t.Errorf("expected items")
	}
	dc.AddAnySliceFromSingleItem(true, nil)
}

func Test_C32_136_DynCol_JsonResultsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	rc := dc.JsonResultsCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyDynamicCollection()
	if empty.JsonResultsCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_137_DynCol_JsonResultsPtrCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	rc := dc.JsonResultsPtrCollection()
	if rc.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyDynamicCollection()
	if empty.JsonResultsPtrCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C32_138_DynCol_JsonString(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(42, true)
	s, err := dc.JsonString()
	if err != nil || s == "" {
		t.Errorf("json string failed")
	}
}

func Test_C32_139_DynCol_JsonStringMust(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	s := dc.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C32_140_DynCol_MarshalUnmarshalJSON(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("hello", true)
	b, err := dc.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Errorf("marshal failed")
	}
}

func Test_C32_141_DynCol_SafeLimitCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	slc := dc.SafeLimitCollection(100)
	if slc.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C32_142_DynCol_At_First_Last_Accessors(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	atVal := dc.At(1)
	if atVal.ValueString() == "" {
		t.Errorf("expected non-empty")
	}
	f := dc.First()
	if f.ValueString() == "" {
		t.Errorf("expected non-empty")
	}
	l := dc.Last()
	if l.ValueString() == "" {
		t.Errorf("expected non-empty")
	}
	fd := dc.FirstDynamic()
	if fd == nil {
		t.Errorf("expected non-nil")
	}
	ld := dc.LastDynamic()
	if ld == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_143_DynCol_FirstOrDefault_LastOrDefault(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	if dc.FirstOrDefault() == nil {
		t.Errorf("expected non-nil")
	}
	if dc.LastOrDefault() == nil {
		t.Errorf("expected non-nil")
	}
	if dc.FirstOrDefaultDynamic() == nil {
		t.Errorf("expected non-nil")
	}
	if dc.LastOrDefaultDynamic() == nil {
		t.Errorf("expected non-nil")
	}
	empty := coredynamic.EmptyDynamicCollection()
	if empty.FirstOrDefault() != nil {
		t.Errorf("expected nil")
	}
	if empty.LastOrDefault() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C32_144_DynCol_Skip_Take_Limit(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true).AddAny(4, true)
	if len(dc.Skip(2)) != 2 {
		t.Errorf("expected 2")
	}
	if len(dc.Take(2)) != 2 {
		t.Errorf("expected 2")
	}
	if len(dc.Limit(3)) != 3 {
		t.Errorf("expected 3")
	}
	if dc.SkipDynamic(1) == nil {
		t.Errorf("expected non-nil")
	}
	if dc.TakeDynamic(2) == nil {
		t.Errorf("expected non-nil")
	}
	if dc.LimitDynamic(2) == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C32_145_DynCol_SkipCollection_TakeCollection_LimitCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	sc := dc.SkipCollection(1)
	if sc.Length() != 2 {
		t.Errorf("expected 2")
	}
	tc := dc.TakeCollection(2)
	if tc.Length() != 2 {
		t.Errorf("expected 2")
	}
	lc := dc.LimitCollection(2)
	if lc.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C32_146_DynCol_GetPagedCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Errorf("expected 4 pages, got %d", len(pages))
	}
}

func Test_C32_147_DynCol_GetPagedCollection_Small(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	pages := dc.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C32_148_DynCol_GetSinglePageCollection_Small(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	page := dc.GetSinglePageCollection(5, 1)
	if page.Length() != 1 {
		t.Errorf("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection AddWithWgLock equivalent not exists;
// Additional AddAny with sync test
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_149_Collection_AddWithWgLock_Proper(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go c.AddWithWgLock(wg, 1)
	go c.AddWithWgLock(wg, 2)
	go c.AddWithWgLock(wg, 3)
	wg.Wait()
	if c.Length() != 3 {
		t.Errorf("expected 3, got %d", c.Length())
	}
}

func Test_C32_150_MapAnyItems_DiffRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 3})
	_ = diff
}

func Test_C32_151_MapAnyItems_Diff(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	diff := m1.Diff(false, m2)
	_ = diff
}

func Test_C32_152_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})
	_ = diff
	diff2 := m.HashmapDiffUsingRaw(false, map[string]any{"a": 2})
	_ = diff2
}

func Test_C32_153_MapAnyItems_DiffJsonMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})
	_ = msg
}

func Test_C32_154_MapAnyItems_ToStringsSliceOfDiffMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	strs := m.ToStringsSliceOfDiffMap(map[string]any{"a": 2})
	_ = strs
}

func Test_C32_155_MapAnyItems_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	_ = msg
}

func Test_C32_156_MapAnyItems_LogShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})
	_ = msg
}
