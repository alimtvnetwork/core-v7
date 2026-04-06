package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════════════════════════════
// newCreator factories — New.Collection.String/Int/Byte/Any etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_01_NewCreator_String_Empty(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_02_NewCreator_String_Cap(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(10)
	if c.Capacity() < 10 {
		t.Error("expected cap >= 10")
	}
}

func Test_C35_03_NewCreator_String_From(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_04_NewCreator_String_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.Clone([]string{"a", "b"})
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_05_NewCreator_String_Items(t *testing.T) {
	c := coredynamic.New.Collection.String.Items("a", "b", "c")
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C35_06_NewCreator_String_Create(t *testing.T) {
	c := coredynamic.New.Collection.String.Create([]string{"x"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_07_NewCreator_String_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.String.LenCap(3, 10)
	if c.Length() != 3 || c.Capacity() < 10 {
		t.Error("expected len=3 cap>=10")
	}
}

func Test_C35_08_NewCreator_Int_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int.Empty()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_09_NewCreator_Int_Cap(t *testing.T) {
	c := coredynamic.New.Collection.Int.Cap(5)
	c.Add(42)
	if c.At(0) != 42 {
		t.Error("expected 42")
	}
}

func Test_C35_10_NewCreator_Int_Items(t *testing.T) {
	c := coredynamic.New.Collection.Int.Items(1, 2, 3)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C35_11_NewCreator_Int_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int.LenCap(2, 10)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_12_NewCreator_Int64_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int64.Empty()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_13_NewCreator_Int64_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int64.LenCap(1, 5)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_14_NewCreator_Byte_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Byte.Empty()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_15_NewCreator_Byte_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Byte.LenCap(2, 8)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_16_NewCreator_Any_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Any.Empty()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_17_NewCreator_Any_Items(t *testing.T) {
	c := coredynamic.New.Collection.Any.Items(1, "two", 3.0)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// KeyVal — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_18_KeyVal_KeyDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamic()
	if d.Value() != "k" {
		t.Error("expected k")
	}
}

func Test_C35_19_KeyVal_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	d := kv.ValueDynamic()
	if d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C35_20_KeyVal_KeyDynamicPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	dp := kv.KeyDynamicPtr()
	if dp == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_21_KeyVal_KeyDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.KeyDynamicPtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C35_22_KeyVal_ValueDynamicPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	dp := kv.ValueDynamicPtr()
	if dp == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_23_KeyVal_ValueDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.ValueDynamicPtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C35_24_KeyVal_IsKeyNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: 1}
	if !kv.IsKeyNull() {
		t.Error("expected true")
	}
}

func Test_C35_25_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: 1}
	if !kv.IsKeyNullOrEmptyString() {
		t.Error("expected true")
	}
}

func Test_C35_26_KeyVal_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	if !kv.IsValueNull() {
		t.Error("expected true")
	}
}

func Test_C35_27_KeyVal_String(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	s := kv.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_28_KeyVal_String_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C35_29_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	if rv.Kind() != reflect.Int {
		t.Error("expected int")
	}
}

func Test_C35_30_KeyVal_ValueInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	if kv.ValueInt() != 42 {
		t.Error("expected 42")
	}
}

func Test_C35_31_KeyVal_ValueInt_Wrong(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}
	v := kv.ValueInt()
	_ = v // just no panic
}

func Test_C35_32_KeyVal_ValueUInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(10)}
	if kv.ValueUInt() != 10 {
		t.Error("expected 10")
	}
}

func Test_C35_33_KeyVal_ValueStrings(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	if len(kv.ValueStrings()) != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_34_KeyVal_ValueStrings_Wrong(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	if kv.ValueStrings() != nil {
		t.Error("expected nil")
	}
}

func Test_C35_35_KeyVal_ValueBool(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	if !kv.ValueBool() {
		t.Error("expected true")
	}
}

func Test_C35_36_KeyVal_ValueInt64(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(100)}
	if kv.ValueInt64() != 100 {
		t.Error("expected 100")
	}
}

func Test_C35_37_KeyVal_CastKeyVal_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_C35_38_KeyVal_ReflectSetKey_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetKey(nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_39_KeyVal_ValueNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueNullErr()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_40_KeyVal_ValueNullErr_NullValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	err := kv.ValueNullErr()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_41_KeyVal_ValueNullErr_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	err := kv.ValueNullErr()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C35_42_KeyVal_KeyNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyNullErr()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_43_KeyVal_KeyNullErr_NullKey(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: 42}
	err := kv.KeyNullErr()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_44_KeyVal_KeyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "hello", Value: 1}
	if kv.KeyString() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_45_KeyVal_KeyString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.KeyString() != "" {
		t.Error("expected empty")
	}
}

func Test_C35_46_KeyVal_ValueString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "world"}
	if kv.ValueString() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_47_KeyVal_ValueString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	if kv.ValueString() != "" {
		t.Error("expected empty")
	}
}

func Test_C35_48_KeyVal_KeyReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyReflectSet(nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_49_KeyVal_ValueReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueReflectSet(nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_50_KeyVal_ReflectSetTo_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetTo(nil)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_51_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	_ = j
}

func Test_C35_52_KeyVal_JsonPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	if kv.JsonPtr() == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_53_KeyVal_JsonModel(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	if kv.JsonModel() == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_54_KeyVal_Serialize(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C35_55_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jp := kv.JsonPtr()
	kv2 := &coredynamic.KeyVal{}
	_, err := kv2.ParseInjectUsingJson(jp)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C35_56_KeyVal_JsonParseSelfInject(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jp := kv.JsonPtr()
	kv2 := &coredynamic.KeyVal{}
	err := kv2.JsonParseSelfInject(jp)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValCollection — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_57_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	if kvc.Length() != 2 {
		t.Error("expected 2")
	}
	if !kvc.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_C35_58_KeyValCollection_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	if !kvc.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C35_59_KeyValCollection_AddPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	kvc.AddPtr(kv)
	kvc.AddPtr(nil)
	if kvc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_60_KeyValCollection_AddMany(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)
	if kvc.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_61_KeyValCollection_AddManyPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	k1 := &coredynamic.KeyVal{Key: "a", Value: 1}
	kvc.AddManyPtr(k1, nil)
	if kvc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_62_KeyValCollection_Items_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	if kvc.Items() != nil {
		t.Error("expected nil")
	}
}

func Test_C35_63_KeyValCollection_MapAnyItems(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	m := kvc.MapAnyItems()
	if m.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_64_KeyValCollection_MapAnyItems_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	m := kvc.MapAnyItems()
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_65_KeyValCollection_AllKeys(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeys()
	if len(keys) != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_66_KeyValCollection_AllKeysSorted(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeysSorted()
	if len(keys) != 2 || keys[0] > keys[1] {
		t.Error("expected sorted")
	}
}

func Test_C35_67_KeyValCollection_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	vals := kvc.AllValues()
	if len(vals) != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_68_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_69_KeyValCollection_String_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	if kvc.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C35_70_KeyValCollection_Json(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	j := kvc.Json()
	_ = j
}

func Test_C35_71_KeyValCollection_Serialize(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	b, err := kvc.Serialize()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C35_72_KeyValCollection_JsonString(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := kvc.JsonString()
	if err != nil || s == "" {
		t.Error("expected json string")
	}
}

func Test_C35_73_KeyValCollection_JsonStringMust(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.JsonStringMust()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_74_KeyValCollection_Clone(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	cloned := kvc.Clone()
	if cloned.Length() != 1 {
		t.Error("expected 1")
	}
}
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	if kvc.GetPagesSize(3) != 4 {
		t.Error("expected 4 pages")
	}
	if kvc.GetPagesSize(0) != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_78_KeyValCollection_GetSinglePageCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	page := kvc.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C35_79_KeyValCollection_GetPagedCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Errorf("expected 4, got %d", len(pages))
	}
}

func Test_C35_80_KeyValCollection_GetPagedCollection_Small(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	pages := kvc.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_81_KeyValCollection_JsonResultsCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	rc := kvc.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_82_KeyValCollection_JsonResultsPtrCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	rc := kvc.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C35_83_KeyValCollection_JsonMapResults(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	mr, err := kvc.JsonMapResults()
	if err != nil || mr == nil {
		t.Error("expected non-nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — Add, Get, Deserialize, validation, paging
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_84_MapAnyItems_Add(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.Add("key", 42)
	if !isNew {
		t.Error("expected new")
	}
	isNew2 := m.Add("key", 99)
	if isNew2 {
		t.Error("expected not new")
	}
}

func Test_C35_85_MapAnyItems_Set(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Set("k", 1)
	if m.GetValue("k") != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_86_MapAnyItems_HasKey(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	if !m.HasKey("k") {
		t.Error("expected true")
	}
	if m.HasKey("nope") {
		t.Error("expected false")
	}
}

func Test_C35_87_MapAnyItems_HasKey_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	if m.HasKey("k") {
		t.Error("expected false for nil")
	}
}

func Test_C35_88_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	v, has := m.Get("k")
	if !has || v != 42 {
		t.Error("expected 42")
	}
	_, has2 := m.Get("nope")
	if has2 {
		t.Error("expected not found")
	}
}

func Test_C35_89_MapAnyItems_GetValue(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "hello")
	if m.GetValue("k") != "hello" {
		t.Error("expected hello")
	}
	if m.GetValue("nope") != nil {
		t.Error("expected nil")
	}
}

func Test_C35_90_MapAnyItems_ReflectSetTo(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	err := m.ReflectSetTo("k", &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C35_91_MapAnyItems_ReflectSetTo_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.ReflectSetTo("nope", &target)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_92_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	err := m.Deserialize("k", &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if target != 42 {
		t.Errorf("expected 42, got %d", target)
	}
}

func Test_C35_93_MapAnyItems_Deserialize_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.Deserialize("nope", &target)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_94_MapAnyItems_AddKeyAny(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: 42})
	if !isNew {
		t.Error("expected new")
	}
}

func Test_C35_95_MapAnyItems_AddKeyAnyWithValidation_Valid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddKeyAnyWithValidation(reflect.TypeOf(0), corejson.KeyAny{Key: "k", AnyInf: 42})
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C35_96_MapAnyItems_AddKeyAnyWithValidation_Invalid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddKeyAnyWithValidation(reflect.TypeOf(0), corejson.KeyAny{Key: "k", AnyInf: "str"})
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_97_MapAnyItems_AddWithValidation_Valid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddWithValidation(reflect.TypeOf(0), "k", 42)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C35_98_MapAnyItems_AddWithValidation_Invalid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddWithValidation(reflect.TypeOf(0), "k", "str")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_99_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	jr := corejson.NewPtr(42)
	m.AddJsonResultPtr("k", jr)
	if !m.HasKey("k") {
		t.Error("expected key")
	}
}

func Test_C35_100_MapAnyItems_AddJsonResultPtr_Nil(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddJsonResultPtr("k", nil)
	if m.HasKey("k") {
		t.Error("expected no key")
	}
}

func Test_C35_101_MapAnyItems_GetPagesSize(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	if m.GetPagesSize(2) != 2 {
		t.Error("expected 2 pages")
	}
	if m.GetPagesSize(0) != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_102_MapAnyItems_GetFieldsMap(t *testing.T) {
	inner := map[string]any{"x": 1}
	b, _ := json.Marshal(inner)
	var stored any
	json.Unmarshal(b, &stored)
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", stored)
	fm, err, found := m.GetFieldsMap("k")
	if !found {
		t.Error("expected found")
	}
	_ = fm
	_ = err
}

func Test_C35_103_MapAnyItems_GetFieldsMap_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	_, _, found := m.GetFieldsMap("nope")
	if found {
		t.Error("expected not found")
	}
}

func Test_C35_104_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	_, found := m.GetSafeFieldsMap("nope")
	if found {
		t.Error("expected not found")
	}
}

func Test_C35_105_NewMapAnyItemsUsingItems(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": 1})
	if m.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_106_NewMapAnyItemsUsingItems_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(nil)
	if m.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_107_NewMapAnyItemsUsingAnyTypeMap(t *testing.T) {
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_C35_108_NewMapAnyItemsUsingAnyTypeMap_Valid(t *testing.T) {
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]any{"k": 1})
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if m.Length() != 1 {
		t.Error("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItemDiff
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_109_MapAnyItemDiff_Length(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	if d.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C35_110_MapAnyItemDiff_Length_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	if d.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_111_MapAnyItemDiff_IsEmpty(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{}
	if !d.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C35_112_MapAnyItemDiff_HasAnyItem(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"a": 1}
	if !d.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_C35_113_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"b": 1, "a": 2}
	keys := d.AllKeysSorted()
	if len(keys) != 2 || keys[0] != "a" {
		t.Error("expected sorted keys")
	}
}

func Test_C35_114_MapAnyItemDiff_Raw(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	r := d.Raw()
	if len(r) != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_115_MapAnyItemDiff_Raw_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	if len(d.Raw()) != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_116_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	m := d.MapAnyItems()
	if m.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_117_MapAnyItemDiff_RawMapDiffer(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	rd := d.RawMapDiffer()
	if len(rd) != 1 {
		t.Error("expected 1")
	}
}

func Test_C35_118_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	if !d.IsRawEqual(false, map[string]any{"k": 1}) {
		t.Error("expected equal")
	}
}

func Test_C35_119_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	if d.HasAnyChanges(false, map[string]any{"k": 1}) {
		t.Error("expected no changes")
	}
	if !d.HasAnyChanges(false, map[string]any{"k": 2}) {
		t.Error("expected changes")
	}
}

func Test_C35_120_MapAnyItemDiff_Clear(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	cleared := d.Clear()
	if len(cleared) != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_121_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	cleared := d.Clear()
	if len(cleared) != 0 {
		t.Error("expected 0")
	}
}

func Test_C35_122_MapAnyItemDiff_Json(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	j := d.Json()
	_ = j
}

func Test_C35_123_MapAnyItemDiff_PrettyJsonString(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	s := d.PrettyJsonString()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_124_MapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	d.LogPrettyJsonString() // just no panic
}

func Test_C35_125_MapAnyItemDiff_LogPrettyJsonString_Empty(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{}
	d.LogPrettyJsonString() // no panic
}

// ═══════════════════════════════════════════════════════════════════════
// TypeStatus — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_126_TypeStatus_IsValid(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if !st.IsValid() {
		t.Error("expected valid")
	}
}

func Test_C35_127_TypeStatus_IsValid_Nil(t *testing.T) {
	var st *coredynamic.TypeStatus
	if st.IsValid() {
		t.Error("expected false")
	}
}

func Test_C35_128_TypeStatus_IsInvalid(t *testing.T) {
	var st *coredynamic.TypeStatus
	if !st.IsInvalid() {
		t.Error("expected true")
	}
}

func Test_C35_129_TypeStatus_IsNotSame(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "hello")
	if !st.IsNotSame() {
		t.Error("expected not same")
	}
}

func Test_C35_130_TypeStatus_IsNotEqualTypes(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	if !st.IsNotEqualTypes() {
		t.Error("expected true")
	}
}

func Test_C35_131_TypeStatus_IsAnyPointer(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 42)
	if !st.IsAnyPointer() {
		t.Error("expected true")
	}
}

func Test_C35_132_TypeStatus_IsBothPointer(t *testing.T) {
	v1, v2 := 42, 100
	st := coredynamic.TypeSameStatus(&v1, &v2)
	if !st.IsBothPointer() {
		t.Error("expected true")
	}
}

func Test_C35_133_TypeStatus_NonPointerLeft(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 100)
	npl := st.NonPointerLeft()
	if npl.Kind() != reflect.Int {
		t.Error("expected int")
	}
}

func Test_C35_134_TypeStatus_NonPointerRight(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(100, &v)
	npr := st.NonPointerRight()
	if npr.Kind() != reflect.Int {
		t.Error("expected int")
	}
}

func Test_C35_135_TypeStatus_IsSameRegardlessPointer(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 100)
	if !st.IsSameRegardlessPointer() {
		t.Error("expected true")
	}
}

func Test_C35_136_TypeStatus_LeftName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.LeftName() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_137_TypeStatus_LeftName_Nil(t *testing.T) {
	st := coredynamic.TypeSameStatus(nil, 42)
	n := st.LeftName()
	if n == "" {
		t.Error("expected something")
	}
}

func Test_C35_138_TypeStatus_RightName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.RightName() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_139_TypeStatus_LeftFullName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.LeftFullName() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_140_TypeStatus_RightFullName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.RightFullName() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C35_141_TypeStatus_NotMatchMessage_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.NotMatchMessage("l", "r") != "" {
		t.Error("expected empty for same")
	}
}

func Test_C35_142_TypeStatus_NotMatchMessage_Different(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	msg := st.NotMatchMessage("l", "r")
	if msg == "" {
		t.Error("expected message")
	}
}

func Test_C35_143_TypeStatus_NotMatchErr_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.NotMatchErr("l", "r") != nil {
		t.Error("expected nil")
	}
}

func Test_C35_144_TypeStatus_NotMatchErr_Different(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	if st.NotMatchErr("l", "r") == nil {
		t.Error("expected error")
	}
}

func Test_C35_145_TypeStatus_MustBeSame_Same(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not panic")
		}
	}()
	st := coredynamic.TypeSameStatus(42, 100)
	st.MustBeSame()
}

func Test_C35_146_TypeStatus_MustBeSame_Different(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	st := coredynamic.TypeSameStatus(42, "x")
	st.MustBeSame()
}

func Test_C35_147_TypeStatus_ValidationError(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	if st.ValidationError() == nil {
		t.Error("expected error")
	}
}

func Test_C35_148_TypeStatus_ValidationError_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.ValidationError() != nil {
		t.Error("expected nil")
	}
}

func Test_C35_149_TypeStatus_NotEqualSrcDestinationMessage(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	msg := st.NotEqualSrcDestinationMessage()
	if msg == "" {
		t.Error("expected message")
	}
}

func Test_C35_150_TypeStatus_NotEqualSrcDestinationErr(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	if st.NotEqualSrcDestinationErr() == nil {
		t.Error("expected error")
	}
}

func Test_C35_151_TypeStatus_SrcDestinationMustBeSame_Same(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not panic")
		}
	}()
	st := coredynamic.TypeSameStatus(42, 100)
	st.SrcDestinationMustBeSame()
}

func Test_C35_152_TypeStatus_IsEqual_BothNil(t *testing.T) {
	var a, b *coredynamic.TypeStatus
	if !a.IsEqual(b) {
		t.Error("expected true")
	}
}

func Test_C35_153_TypeStatus_IsEqual_OneNil(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if st.IsEqual(nil) {
		t.Error("expected false")
	}
}

func Test_C35_154_TypeStatus_IsEqual_Same(t *testing.T) {
	st1 := coredynamic.TypeSameStatus(42, 100)
	st2 := coredynamic.TypeSameStatus(42, 100)
	if !st1.IsEqual(&st2) {
		t.Error("expected true")
	}
}

func Test_C35_155_TypeStatus_IsEqual_Different(t *testing.T) {
	st1 := coredynamic.TypeSameStatus(42, 100)
	st2 := coredynamic.TypeSameStatus(42, "x")
	if st1.IsEqual(&st2) {
		t.Error("expected false")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicStatus, ValueStatus, CastedResult
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_156_DynamicStatus_Invalid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	if ds.IsValid() {
		t.Error("expected invalid")
	}
	if ds.Message != "msg" {
		t.Error("expected msg")
	}
}

func Test_C35_157_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	if ds.IsValid() {
		t.Error("expected invalid")
	}
}

func Test_C35_158_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	cloned := ds.Clone()
	if cloned.Message != "msg" {
		t.Error("expected msg")
	}
}
	vs := coredynamic.InvalidValueStatus("msg")
	if vs.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C35_162_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := coredynamic.InvalidValueStatusNoMessage()
	if vs.IsValid {
		t.Error("expected invalid")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// BytesConverter — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_163_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	if bc.SafeCastString() != "hello" {
		t.Error("expected hello")
	}
}

func Test_C35_164_BytesConverter_SafeCastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	if bc.SafeCastString() != "" {
		t.Error("expected empty")
	}
}

func Test_C35_165_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()
	if err != nil || s != "hello" {
		t.Error("expected hello")
	}
}

func Test_C35_166_BytesConverter_CastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	_, err := bc.CastString()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_167_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	v, err := bc.ToBool()
	if err != nil || !v {
		t.Error("expected true")
	}
}

func Test_C35_168_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	var target int
	err := bc.Deserialize(&target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
	if target != 42 {
		t.Errorf("expected 42, got %d", target)
	}
}

func Test_C35_169_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	if err != nil || s != "hello" {
		t.Errorf("expected hello, got %s err=%v", s, err)
	}
}

func Test_C35_170_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	strs, err := bc.ToStrings()
	if err != nil || len(strs) != 2 {
		t.Error("expected 2 strings")
	}
}

func Test_C35_171_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`100`))
	v, err := bc.ToInt64()
	if err != nil || v != 100 {
		t.Errorf("expected 100, got %d", v)
	}
}

func Test_C35_172_BytesConverter_ToHashmap(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":"1"}`))
	hm, err := bc.ToHashmap()
	if err != nil || hm == nil {
		t.Error("expected hashmap")
	}
}

func Test_C35_173_BytesConverter_ToHashmap_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashmap()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_174_BytesConverter_ToHashset(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":true,"b":true}`))
	hs, err := bc.ToHashset()
	if err != nil || hs == nil || hs.Length() != 2 {
		t.Error("expected hashset with 2 items")
	}
}

func Test_C35_175_BytesConverter_ToHashset_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashset()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_176_BytesConverter_ToCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	c, err := bc.ToCollection()
	if err != nil || c == nil {
		t.Error("expected collection")
	}
}

func Test_C35_177_BytesConverter_ToCollection_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToCollection()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_178_BytesConverter_ToSimpleSlice(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x","y"]`))
	ss, err := bc.ToSimpleSlice()
	if err != nil || ss == nil {
		t.Error("expected simple slice")
	}
}

func Test_C35_179_BytesConverter_ToSimpleSlice_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToSimpleSlice()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C35_180_NewBytesConverterUsingJsonResult(t *testing.T) {
	jr := corejson.NewPtr(42)
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(jr)
	if err != nil || bc == nil {
		t.Error("expected converter")
	}
}
