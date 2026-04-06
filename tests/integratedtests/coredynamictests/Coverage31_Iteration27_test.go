package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

// ═══════════════════════════════════════════════════════════════════════
// DynamicGetters — value extraction and type checks
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_01_Dynamic_Data(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	// Act
	got := d.Data()
	// Assert
	if got != "hello" {
		t.Errorf("expected hello, got %v", got)
	}
}

func Test_C31_02_Dynamic_Value(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.Value() != 42 {
		t.Errorf("expected 42")
	}
}

func Test_C31_03_Dynamic_Length_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.Length() != 0 {
		t.Errorf("expected 0 for nil receiver")
	}
}

func Test_C31_04_Dynamic_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	if d.Length() != 3 {
		t.Errorf("expected 3, got %d", d.Length())
	}
}

func Test_C31_05_Dynamic_StructStringPtr_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.StructStringPtr() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C31_06_Dynamic_StructStringPtr_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr("test", true)
	ptr1 := d.StructStringPtr()
	ptr2 := d.StructStringPtr()
	if ptr1 != ptr2 {
		t.Errorf("expected cached pointer to be same")
	}
}

func Test_C31_07_Dynamic_String_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.String() != "" {
		t.Errorf("expected empty string")
	}
}

func Test_C31_08_Dynamic_StructString_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.StructString() != "" {
		t.Errorf("expected empty string")
	}
}

func Test_C31_09_Dynamic_IsNull(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	if !d.IsNull() {
		t.Errorf("expected IsNull true")
	}
}

func Test_C31_10_Dynamic_IsValid_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	if d.IsValid() {
		t.Errorf("expected invalid")
	}
	if !d.IsInvalid() {
		t.Errorf("expected IsInvalid true")
	}
}

func Test_C31_11_Dynamic_IsPointer_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsPointer() {
		t.Errorf("expected false for nil")
	}
}

func Test_C31_12_Dynamic_IsPointer_True(t *testing.T) {
	val := "hello"
	d := coredynamic.NewDynamicPtr(&val, true)
	if !d.IsPointer() {
		t.Errorf("expected IsPointer true")
	}
	// call again to test cached path
	if !d.IsPointer() {
		t.Errorf("expected IsPointer true on second call")
	}
}

func Test_C31_13_Dynamic_IsValueType_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsValueType() {
		t.Errorf("expected false")
	}
}

func Test_C31_14_Dynamic_IsValueType_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	if !d.IsValueType() {
		t.Errorf("expected IsValueType true")
	}
}

func Test_C31_15_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	var d *coredynamic.Dynamic
	if !d.IsStructStringNullOrEmpty() {
		t.Errorf("expected true for nil")
	}
	d2 := coredynamic.NewDynamicPtr(nil, false)
	if !d2.IsStructStringNullOrEmpty() {
		t.Errorf("expected true for null data")
	}
}

func Test_C31_16_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	var d *coredynamic.Dynamic
	if !d.IsStructStringNullOrEmptyOrWhitespace() {
		t.Errorf("expected true for nil")
	}
}

func Test_C31_17_Dynamic_IsPrimitive(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsPrimitive() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr(42, true)
	if !d2.IsPrimitive() {
		t.Errorf("expected true for int")
	}
}

func Test_C31_18_Dynamic_IsNumber(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsNumber() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr(3.14, true)
	if !d2.IsNumber() {
		t.Errorf("expected true for float64")
	}
}

func Test_C31_19_Dynamic_IsStringType(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsStringType() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr("hi", true)
	if !d2.IsStringType() {
		t.Errorf("expected true for string")
	}
	d3 := coredynamic.NewDynamicPtr(42, true)
	if d3.IsStringType() {
		t.Errorf("expected false for int")
	}
}

func Test_C31_20_Dynamic_IsStruct(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsStruct() {
		t.Errorf("expected false for nil")
	}
	type S struct{ X int }
	d2 := coredynamic.NewDynamicPtr(S{X: 1}, true)
	if !d2.IsStruct() {
		t.Errorf("expected true for struct")
	}
}

func Test_C31_21_Dynamic_IsFunc(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsFunc() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr(func() {}, true)
	if !d2.IsFunc() {
		t.Errorf("expected true for func")
	}
}

func Test_C31_22_Dynamic_IsSliceOrArray(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsSliceOrArray() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr([]int{1}, true)
	if !d2.IsSliceOrArray() {
		t.Errorf("expected true for slice")
	}
}

func Test_C31_23_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsSliceOrArrayOrMap() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	if !d2.IsSliceOrArrayOrMap() {
		t.Errorf("expected true for map")
	}
}

func Test_C31_24_Dynamic_IsMap(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsMap() {
		t.Errorf("expected false for nil")
	}
	d2 := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	if !d2.IsMap() {
		t.Errorf("expected true for map")
	}
}

func Test_C31_25_Dynamic_IntDefault(t *testing.T) {
	var d *coredynamic.Dynamic
	val, ok := d.IntDefault(99)
	if ok || val != 99 {
		t.Errorf("expected default 99, ok=false")
	}
	d2 := coredynamic.NewDynamicPtr("42", true)
	val2, ok2 := d2.IntDefault(0)
	if !ok2 || val2 != 42 {
		t.Errorf("expected 42, ok=true")
	}
	d3 := coredynamic.NewDynamicPtr("abc", true)
	val3, ok3 := d3.IntDefault(7)
	if ok3 || val3 != 7 {
		t.Errorf("expected default 7, ok=false")
	}
}

func Test_C31_26_Dynamic_Float64(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Float64()
	if err == nil {
		t.Errorf("expected error for nil")
	}
	d2 := coredynamic.NewDynamicPtr("3.14", true)
	val, err2 := d2.Float64()
	if err2 != nil || val != 3.14 {
		t.Errorf("expected 3.14, got %v err=%v", val, err2)
	}
	d3 := coredynamic.NewDynamicPtr("notfloat", true)
	_, err3 := d3.Float64()
	if err3 == nil {
		t.Errorf("expected parse error")
	}
}

func Test_C31_27_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.ValueInt() != 42 {
		t.Errorf("expected 42")
	}
	d2 := coredynamic.NewDynamicValid("notint")
	if d2.ValueInt() == 42 {
		t.Errorf("expected invalid value")
	}
}

func Test_C31_28_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(10))
	if d.ValueUInt() != 10 {
		t.Errorf("expected 10")
	}
	d2 := coredynamic.NewDynamicValid("x")
	if d2.ValueUInt() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_29_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	if len(d.ValueStrings()) != 2 {
		t.Errorf("expected 2")
	}
	d2 := coredynamic.NewDynamicValid(42)
	if d2.ValueStrings() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C31_30_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	if !d.ValueBool() {
		t.Errorf("expected true")
	}
	d2 := coredynamic.NewDynamicValid("x")
	if d2.ValueBool() {
		t.Errorf("expected false")
	}
}

func Test_C31_31_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(100))
	if d.ValueInt64() != 100 {
		t.Errorf("expected 100")
	}
	d2 := coredynamic.NewDynamicValid("x")
	if d2.ValueInt64() == 100 {
		t.Errorf("expected invalid")
	}
}

func Test_C31_32_Dynamic_ValueNullErr(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.ValueNullErr() == nil {
		t.Errorf("expected error for nil receiver")
	}
	d2 := coredynamic.NewDynamicPtr(nil, false)
	if d2.ValueNullErr() == nil {
		t.Errorf("expected error for null data")
	}
	d3 := coredynamic.NewDynamicPtr("ok", true)
	if d3.ValueNullErr() != nil {
		t.Errorf("expected nil error")
	}
}

func Test_C31_33_Dynamic_ValueString(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.ValueString() != "" {
		t.Errorf("expected empty string")
	}
	d2 := coredynamic.NewDynamicPtr("hello", true)
	if d2.ValueString() != "hello" {
		t.Errorf("expected hello")
	}
	d3 := coredynamic.NewDynamicPtr(42, true)
	if d3.ValueString() == "" {
		t.Errorf("expected non-empty string")
	}
}

func Test_C31_34_Dynamic_Bytes(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	if ok || b != nil {
		t.Errorf("expected nil, false for nil receiver")
	}
	d2 := coredynamic.NewDynamicPtr([]byte{1, 2}, true)
	b2, ok2 := d2.Bytes()
	if !ok2 || len(b2) != 2 {
		t.Errorf("expected bytes")
	}
	d3 := coredynamic.NewDynamicPtr("str", true)
	_, ok3 := d3.Bytes()
	if ok3 {
		t.Errorf("expected false for string")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicJson
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_35_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	b, err := d.JsonBytesPtr()
	if err != nil || len(b) != 0 {
		t.Errorf("expected empty bytes for null")
	}
}

func Test_C31_36_Dynamic_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.JsonBytesPtr()
	if err != nil || len(b) == 0 {
		t.Errorf("expected json bytes")
	}
}

func Test_C31_37_Dynamic_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	s, err := d.JsonString()
	if err != nil || s != "42" {
		t.Errorf("expected '42', got '%s'", s)
	}
}

func Test_C31_38_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hi", true)
	s := d.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C31_39_Dynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamicValid(99)
	if d.JsonModel() != 99 {
		t.Errorf("expected 99")
	}
	if d.JsonModelAny() != 99 {
		t.Errorf("expected 99")
	}
}

func Test_C31_40_Dynamic_Json_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	j := d.Json()
	if j.HasError() {
		t.Errorf("unexpected error")
	}
	jp := d.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C31_41_Dynamic_ValueMarshal_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_C31_42_Dynamic_Deserialize_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`{}`))
	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_C31_43_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	b := d.JsonPayloadMust()
	if len(b) == 0 {
		t.Errorf("expected bytes")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicReflect
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_44_Dynamic_ReflectSetTo_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	if err == nil {
		t.Errorf("expected error for nil receiver")
	}
}

func Test_C31_45_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	kv, err := d.MapToKeyVal()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if kv.Length() != 2 {
		t.Errorf("expected 2 entries, got %d", kv.Length())
	}
}

func Test_C31_46_Dynamic_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	name := d.ReflectTypeName()
	if name == "" {
		t.Errorf("expected non-empty type name")
	}
}

func Test_C31_47_Dynamic_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	if !d.IsReflectTypeOf(reflect.TypeOf("")) {
		t.Errorf("expected true for string type")
	}
}

func Test_C31_48_Dynamic_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]string{"a", "b", "c"}, true)
	if d.ItemUsingIndex(1) != "b" {
		t.Errorf("expected b")
	}
	rv := d.ItemReflectValueUsingIndex(0)
	if rv.String() != "a" {
		t.Errorf("expected a")
	}
}

func Test_C31_49_Dynamic_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"x": 5}, true)
	if d.ItemUsingKey("x") != 5 {
		t.Errorf("expected 5")
	}
	rv := d.ItemReflectValueUsingKey("x")
	if rv.Interface() != 5 {
		t.Errorf("expected 5")
	}
}

func Test_C31_50_Dynamic_Loop_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.Loop(func(i int, item any) bool { return false })
	if called {
		t.Errorf("expected false for nil data")
	}
}

func Test_C31_51_Dynamic_Loop_WithBreak(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2 iterations, got %d", count)
	}
}

func Test_C31_52_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	if !called {
		t.Errorf("expected called")
	}
}

func Test_C31_53_Dynamic_LoopMap_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	if called {
		t.Errorf("expected false")
	}
}

func Test_C31_54_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4, 5}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt()%2 == 0, false
	})
	if result.Length() != 2 {
		t.Errorf("expected 2 even numbers, got %d", result.Length())
	}
}

func Test_C31_55_Dynamic_FilterAsDynamicCollection_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	if result.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_56_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4, 5}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 2
	})
	if result.Length() != 3 {
		t.Errorf("expected 3, got %d", result.Length())
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Dynamic Clone / Ptr / NonPtr
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_57_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	c := d.Clone()
	if c.Value() != "hello" {
		t.Errorf("expected hello")
	}
}
func Test_C31_59_Dynamic_NonPtr_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	np := d.NonPtr()
	if np.Value() != 42 {
		t.Errorf("expected 42")
	}
	p := d.Ptr()
	if p == nil {
		t.Errorf("expected non-nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Collection[T] — generic
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_60_Collection_Basic(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20).Add(30)
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
	if c.First() != 10 || c.Last() != 30 {
		t.Errorf("wrong first/last")
	}
	if c.At(1) != 20 {
		t.Errorf("expected 20 at index 1")
	}
}

func Test_C31_61_Collection_FirstOrDefault_LastOrDefault(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	f, ok := c.FirstOrDefault()
	if ok || f != nil {
		t.Errorf("expected nil, false for empty")
	}
	l, ok := c.LastOrDefault()
	if ok || l != nil {
		t.Errorf("expected nil, false for empty")
	}
	c.Add(5)
	f2, ok2 := c.FirstOrDefault()
	if !ok2 || *f2 != 5 {
		t.Errorf("expected 5")
	}
}

func Test_C31_62_Collection_Items_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	items := c.Items()
	if len(items) != 0 {
		t.Errorf("expected empty")
	}
}

func Test_C31_63_Collection_Skip_Take_Limit(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	if len(c.Skip(2)) != 3 {
		t.Errorf("expected 3")
	}
	if len(c.Take(2)) != 2 {
		t.Errorf("expected 2")
	}
	if len(c.Limit(3)) != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_64_Collection_SkipCollection_TakeCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	sc := c.SkipCollection(3)
	if sc.Length() != 2 {
		t.Errorf("expected 2, got %d", sc.Length())
	}
	tc := c.TakeCollection(2)
	if tc.Length() != 2 {
		t.Errorf("expected 2")
	}
	lc := c.LimitCollection(3)
	if lc.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_65_Collection_SafeLimitCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sl := c.SafeLimitCollection(10)
	if sl.Length() != 3 {
		t.Errorf("expected 3 (capped)")
	}
}

func Test_C31_66_Collection_AddMany_AddNonNil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddMany(1, 2, 3)
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
	val := 99
	c.AddNonNil(&val)
	c.AddNonNil(nil)
	if c.Length() != 4 {
		t.Errorf("expected 4")
	}
}

func Test_C31_67_Collection_RemoveAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !c.RemoveAt(1) {
		t.Errorf("expected success")
	}
	if c.Length() != 2 {
		t.Errorf("expected 2")
	}
	if c.RemoveAt(99) {
		t.Errorf("expected false for invalid index")
	}
}

func Test_C31_68_Collection_Clear_Dispose(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Clear()
	if c.Length() != 0 {
		t.Errorf("expected 0")
	}
	c.Add(1)
	c.Dispose()
	if c.Length() != 0 {
		t.Errorf("expected 0 after dispose")
	}
}

func Test_C31_69_Collection_Loop(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	sum := 0
	c.Loop(func(i int, item int) bool {
		sum += item
		return false
	})
	if sum != 60 {
		t.Errorf("expected 60, got %d", sum)
	}
}

func Test_C31_70_Collection_Loop_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	called := false
	c.Loop(func(i int, item int) bool {
		called = true
		return false
	})
	if called {
		t.Errorf("should not be called on empty")
	}
}

func Test_C31_71_Collection_Loop_Break(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	count := 0
	c.Loop(func(i int, item int) bool {
		count++
		return i == 2
	})
	if count != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func Test_C31_72_Collection_Filter(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6})
	evens := c.Filter(func(v int) bool { return v%2 == 0 })
	if evens.Length() != 3 {
		t.Errorf("expected 3 evens, got %d", evens.Length())
	}
}

func Test_C31_73_Collection_Filter_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := c.Filter(func(v int) bool { return true })
	if result.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_74_Collection_LoopAsync(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.LoopAsync(func(i int, item int) {
		// just ensure it doesn't panic
	})
}

func Test_C31_75_Collection_LoopAsync_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, item int) {
		t.Errorf("should not be called")
	})
}

// ═══════════════════════════════════════════════════════════════════════
// Collection paging
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_76_Collection_GetPagesSize(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	if c.GetPagesSize(2) != 3 {
		t.Errorf("expected 3 pages")
	}
	if c.GetPagesSize(0) != 0 {
		t.Errorf("expected 0 for invalid page size")
	}
	if c.GetPagesSize(-1) != 0 {
		t.Errorf("expected 0 for negative page size")
	}
}

func Test_C31_77_Collection_GetSinglePageCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	page := c.GetSinglePageCollection(3, 2)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C31_78_Collection_GetSinglePageCollection_SmallCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	page := c.GetSinglePageCollection(5, 1)
	if page.Length() != 2 {
		t.Errorf("expected 2, got %d", page.Length())
	}
}

func Test_C31_79_Collection_GetPagedCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6, 7})
	pages := c.GetPagedCollection(3)
	if len(pages) != 3 {
		t.Errorf("expected 3 pages, got %d", len(pages))
	}
}

func Test_C31_80_Collection_GetPagedCollection_SmallCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	pages := c.GetPagedCollection(5)
	if len(pages) != 1 {
		t.Errorf("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Collection serialization
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_81_Collection_MarshalUnmarshalJSON(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	b, err := c.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Errorf("marshal failed")
	}
	c2 := coredynamic.EmptyCollection[int]()
	err = c2.UnmarshalJSON(b)
	if err != nil || c2.Length() != 3 {
		t.Errorf("unmarshal failed")
	}
}

func Test_C31_82_Collection_JsonString(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	s, err := c.JsonString()
	if err != nil || s != "[1,2]" {
		t.Errorf("expected [1,2], got %s", s)
	}
}

func Test_C31_83_Collection_JsonStringMust(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	s := c.JsonStringMust()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C31_84_Collection_Strings(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	strs := c.Strings()
	if len(strs) != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_85_Collection_String(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	s := c.String()
	if s == "" {
		t.Errorf("expected non-empty string")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMethods
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_86_Collection_AddIf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIf(true, 1)
	c.AddIf(false, 2)
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C31_87_Collection_AddManyIf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true, 1, 2, 3)
	c.AddManyIf(false, 4, 5)
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_88_Collection_AddCollection(t *testing.T) {
	c1 := coredynamic.CollectionFrom([]int{1, 2})
	c2 := coredynamic.CollectionFrom([]int{3, 4})
	c1.AddCollection(c2)
	if c1.Length() != 4 {
		t.Errorf("expected 4")
	}
	c1.AddCollection(nil)
	if c1.Length() != 4 {
		t.Errorf("expected still 4")
	}
}

func Test_C31_89_Collection_AddCollections(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2, 3})
	c3 := coredynamic.CollectionFrom([]int{4})
	c.AddCollections(c2, nil, c3)
	if c.Length() != 4 {
		t.Errorf("expected 4")
	}
}

func Test_C31_90_Collection_ConcatNew(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	c2 := c.ConcatNew(3, 4)
	if c2.Length() != 4 || c.Length() != 2 {
		t.Errorf("ConcatNew should not mutate original")
	}
}

func Test_C31_91_Collection_Clone(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c2 := c.Clone()
	if c2.Length() != 3 {
		t.Errorf("expected 3")
	}
	var nilC *coredynamic.Collection[int]
	c3 := nilC.Clone()
	if c3.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_92_Collection_Capacity(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	if c.Capacity() < 10 {
		t.Errorf("expected capacity >= 10")
	}
	var nilC *coredynamic.Collection[int]
	if nilC.Capacity() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_93_Collection_AddCapacity(t *testing.T) {
	c := coredynamic.NewCollection[int](5)
	c.AddCapacity(10)
	if c.Capacity() < 15 {
		t.Errorf("expected >= 15")
	}
	c.AddCapacity(0)
	c.AddCapacity(-1)
}

func Test_C31_94_Collection_Resize(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Resize(100)
	if c.Capacity() < 100 {
		t.Errorf("expected >= 100")
	}
	// no-op resize
	c.Resize(5)
}

func Test_C31_95_Collection_Reverse(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	c.Reverse()
	if c.At(0) != 4 || c.At(3) != 1 {
		t.Errorf("expected reversed")
	}
	// single element
	c2 := coredynamic.CollectionFrom([]int{1})
	c2.Reverse()
	if c2.At(0) != 1 {
		t.Errorf("single element unchanged")
	}
}

func Test_C31_96_Collection_InsertAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 5})
	c.InsertAt(2, 3, 4)
	if c.Length() != 5 || c.At(2) != 3 || c.At(3) != 4 {
		t.Errorf("InsertAt failed")
	}
	// no items
	c.InsertAt(0)
}

func Test_C31_97_Collection_IndexOfFunc_ContainsFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	idx := c.IndexOfFunc(func(v int) bool { return v == 20 })
	if idx != 1 {
		t.Errorf("expected 1")
	}
	if !c.ContainsFunc(func(v int) bool { return v == 30 }) {
		t.Errorf("expected contains")
	}
	if c.ContainsFunc(func(v int) bool { return v == 99 }) {
		t.Errorf("expected not contains")
	}
}

func Test_C31_98_Collection_SafeAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if c.SafeAt(1) != 2 {
		t.Errorf("expected 2")
	}
	if c.SafeAt(99) != 0 {
		t.Errorf("expected zero value for invalid index")
	}
}

func Test_C31_99_Collection_SprintItems(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.SprintItems("[%v]")
	if items[0] != "[1]" || items[1] != "[2]" {
		t.Errorf("format mismatch")
	}
}

func Test_C31_100_Collection_HasIndex(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !c.HasIndex(2) {
		t.Errorf("expected true")
	}
	if c.HasIndex(3) {
		t.Errorf("expected false")
	}
	if c.HasIndex(-1) {
		t.Errorf("expected false for negative")
	}
}

func Test_C31_101_Collection_Count_Alias(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	if c.Count() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_102_Collection_HasAnyItem(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	if c.HasAnyItem() {
		t.Errorf("expected false")
	}
	c.Add(1)
	if !c.HasAnyItem() {
		t.Errorf("expected true")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionLock
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_103_CollectionLock_LengthLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	if c.LengthLock() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_104_CollectionLock_IsEmptyLock(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	if !c.IsEmptyLock() {
		t.Errorf("expected true")
	}
}

func Test_C31_105_CollectionLock_AddLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddLock(1).AddLock(2)
	if c.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_106_CollectionLock_AddsLock_AddManyLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddsLock(1, 2, 3)
	c.AddManyLock(4, 5)
	if c.Length() != 5 {
		t.Errorf("expected 5")
	}
}

func Test_C31_107_CollectionLock_AddCollectionLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2, 3})
	c.AddCollectionLock(c2)
	c.AddCollectionLock(nil)
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_108_CollectionLock_AddCollectionsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2})
	c3 := coredynamic.CollectionFrom([]int{3})
	c.AddCollectionsLock(c2, nil, c3)
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_109_CollectionLock_AddIfLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(true, 1)
	c.AddIfLock(false, 2)
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C31_110_CollectionLock_RemoveAtLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !c.RemoveAtLock(1) {
		t.Errorf("expected success")
	}
	if c.RemoveAtLock(99) {
		t.Errorf("expected false")
	}
}

func Test_C31_111_CollectionLock_ClearLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.ClearLock()
	if c.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_112_CollectionLock_ItemsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.ItemsLock()
	if len(items) != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_113_CollectionLock_FirstLock_LastLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	if c.FirstLock() != 10 || c.LastLock() != 30 {
		t.Errorf("wrong first/last lock")
	}
}

func Test_C31_114_CollectionLock_LoopLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sum := 0
	c.LoopLock(func(i int, item int) bool {
		sum += item
		return false
	})
	if sum != 6 {
		t.Errorf("expected 6")
	}
}

func Test_C31_115_CollectionLock_LoopLock_Break(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.LoopLock(func(i int, item int) bool {
		count++
		return i == 0
	})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

func Test_C31_116_CollectionLock_FilterLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	evens := c.FilterLock(func(v int) bool { return v%2 == 0 })
	if evens.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_117_CollectionLock_StringsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	strs := c.StringsLock()
	if len(strs) != 2 {
		t.Errorf("expected 2")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSearch
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_118_Contains_IndexOf(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	if !coredynamic.Contains(c, "b") {
		t.Errorf("expected contains b")
	}
	if coredynamic.IndexOf(c, "d") != -1 {
		t.Errorf("expected -1")
	}
}

func Test_C31_119_Has_HasAll(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !coredynamic.Has(c, 2) {
		t.Errorf("expected has")
	}
	if !coredynamic.HasAll(c, 1, 2, 3) {
		t.Errorf("expected has all")
	}
	if coredynamic.HasAll(c, 1, 4) {
		t.Errorf("expected false")
	}
	empty := coredynamic.EmptyCollection[int]()
	if coredynamic.HasAll(empty, 1) {
		t.Errorf("expected false for empty")
	}
}

func Test_C31_120_LastIndexOf(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 2, 1})
	if coredynamic.LastIndexOf(c, 2) != 3 {
		t.Errorf("expected 3")
	}
	if coredynamic.LastIndexOf(c, 99) != -1 {
		t.Errorf("expected -1")
	}
}

func Test_C31_121_Count(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 2, 3, 2})
	if coredynamic.Count(c, 2) != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_122_ContainsLock_IndexOfLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	if !coredynamic.ContainsLock(c, 20) {
		t.Errorf("expected true")
	}
	if coredynamic.IndexOfLock(c, 40) != -1 {
		t.Errorf("expected -1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSort
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_123_SortFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })
	if c.At(0) != 1 || c.At(2) != 3 {
		t.Errorf("sort failed")
	}
}

func Test_C31_124_SortAsc_SortDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAsc(c)
	if c.At(0) != 1 {
		t.Errorf("expected 1 first")
	}
	coredynamic.SortDesc(c)
	if c.At(0) != 3 {
		t.Errorf("expected 3 first")
	}
}

func Test_C31_125_SortAscLock_SortDescLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAscLock(c)
	if c.At(0) != 1 {
		t.Errorf("expected 1")
	}
	coredynamic.SortDescLock(c)
	if c.At(0) != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_126_SortedAsc_SortedDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	asc := coredynamic.SortedAsc(c)
	if asc.At(0) != 1 {
		t.Errorf("expected 1")
	}
	if c.At(0) != 3 {
		t.Errorf("original should be unchanged")
	}
	desc := coredynamic.SortedDesc(c)
	if desc.At(0) != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_127_IsSorted_IsSortedAsc_IsSortedDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !coredynamic.IsSortedAsc(c) {
		t.Errorf("expected sorted asc")
	}
	if coredynamic.IsSortedDesc(c) {
		t.Errorf("expected not sorted desc")
	}
	single := coredynamic.CollectionFrom([]int{1})
	if !single.IsSorted(func(a, b int) bool { return a < b }) {
		t.Errorf("single element is always sorted")
	}
}

func Test_C31_128_SortFuncLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{5, 3, 1})
	c.SortFuncLock(func(a, b int) bool { return a < b })
	if c.At(0) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C31_129_SortedFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	sorted := c.SortedFunc(func(a, b int) bool { return a < b })
	if sorted.At(0) != 1 || c.At(0) != 3 {
		t.Errorf("SortedFunc should not mutate original")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMap
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_130_Map(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	mapped := coredynamic.Map(c, func(v int) string {
		return "x"
	})
	if mapped.Length() != 3 {
		t.Errorf("expected 3")
	}
}

func Test_C31_131_Map_Nil(t *testing.T) {
	result := coredynamic.Map[int, string](nil, func(v int) string { return "" })
	if result.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_132_FlatMap(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	result := coredynamic.FlatMap(c, func(v int) []string {
		return []string{"a", "b"}
	})
	if result.Length() != 4 {
		t.Errorf("expected 4, got %d", result.Length())
	}
}

func Test_C31_133_FlatMap_Nil(t *testing.T) {
	result := coredynamic.FlatMap[int, string](nil, func(v int) []string { return nil })
	if result.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_134_Reduce(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	sum := coredynamic.Reduce(c, 0, func(acc int, item int) int {
		return acc + item
	})
	if sum != 10 {
		t.Errorf("expected 10, got %d", sum)
	}
}

func Test_C31_135_Reduce_Nil(t *testing.T) {
	result := coredynamic.Reduce[int, int](nil, 42, func(acc int, item int) int { return acc })
	if result != 42 {
		t.Errorf("expected initial value 42")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_136_Distinct(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 2, 3, 3, 3})
	d := coredynamic.Distinct(c)
	if d.Length() != 3 {
		t.Errorf("expected 3, got %d", d.Length())
	}
}

func Test_C31_137_Distinct_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	d := coredynamic.Distinct(c)
	if d.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_138_Unique(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "a"})
	u := coredynamic.Unique(c)
	if u.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_139_DistinctLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 1, 2})
	d := coredynamic.DistinctLock(c)
	if d.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_140_DistinctCount(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 1, 2, 3, 3})
	if coredynamic.DistinctCount(c) != 3 {
		t.Errorf("expected 3")
	}
	empty := coredynamic.EmptyCollection[int]()
	if coredynamic.DistinctCount(empty) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_141_IsDistinct(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if !coredynamic.IsDistinct(c) {
		t.Errorf("expected distinct")
	}
	c2 := coredynamic.CollectionFrom([]int{1, 2, 2})
	if coredynamic.IsDistinct(c2) {
		t.Errorf("expected not distinct")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_142_GroupBy(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6})
	groups := coredynamic.GroupBy(c, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if len(groups) != 2 {
		t.Errorf("expected 2 groups")
	}
	if groups["even"].Length() != 3 {
		t.Errorf("expected 3 evens")
	}
}

func Test_C31_143_GroupBy_Nil(t *testing.T) {
	groups := coredynamic.GroupBy[int, string](nil, func(v int) string { return "" })
	if len(groups) != 0 {
		t.Errorf("expected empty map")
	}
}

func Test_C31_144_GroupByLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	groups := coredynamic.GroupByLock(c, func(v int) int { return v % 2 })
	if len(groups) != 2 {
		t.Errorf("expected 2 groups")
	}
}

func Test_C31_145_GroupByCount(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "a", "c", "b", "a"})
	counts := coredynamic.GroupByCount(c, func(v string) string { return v })
	if counts["a"] != 3 || counts["b"] != 2 || counts["c"] != 1 {
		t.Errorf("wrong counts")
	}
}

func Test_C31_146_GroupByCount_Nil(t *testing.T) {
	counts := coredynamic.GroupByCount[string, string](nil, func(v string) string { return v })
	if len(counts) != 0 {
		t.Errorf("expected empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_147_LeftRight_IsEmpty_HasAnyItem(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}
	if !lr.IsEmpty() {
		t.Errorf("expected empty")
	}
	lr2 := &coredynamic.LeftRight{Left: "a", Right: nil}
	if lr2.IsEmpty() {
		t.Errorf("expected not empty")
	}
	if !lr2.HasAnyItem() {
		t.Errorf("expected has any")
	}
}

func Test_C31_148_LeftRight_HasLeft_HasRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "x", Right: nil}
	if !lr.HasLeft() {
		t.Errorf("expected HasLeft true")
	}
	if lr.HasRight() {
		t.Errorf("expected HasRight false")
	}
}

func Test_C31_149_LeftRight_IsLeftEmpty_IsRightEmpty(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "y"}
	if !lr.IsLeftEmpty() {
		t.Errorf("expected left empty")
	}
	if lr.IsRightEmpty() {
		t.Errorf("expected right not empty")
	}
}

func Test_C31_150_LeftRight_LeftToDynamic_RightToDynamic(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ld := lr.LeftToDynamic()
	if ld.Value() != "a" {
		t.Errorf("expected a")
	}
	rd := lr.RightToDynamic()
	if rd.Value() != "b" {
		t.Errorf("expected b")
	}
}

func Test_C31_151_LeftRight_NilReceiver(t *testing.T) {
	var lr *coredynamic.LeftRight
	if !lr.IsEmpty() {
		t.Errorf("expected empty for nil")
	}
	if lr.HasAnyItem() {
		t.Errorf("expected false for nil")
	}
	if lr.HasLeft() {
		t.Errorf("expected false for nil")
	}
	if lr.HasRight() {
		t.Errorf("expected false for nil")
	}
	if !lr.IsLeftEmpty() {
		t.Errorf("expected true for nil")
	}
	if !lr.IsRightEmpty() {
		t.Errorf("expected true for nil")
	}
	if lr.LeftToDynamic() != nil {
		t.Errorf("expected nil")
	}
	if lr.RightToDynamic() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C31_152_LeftRight_DeserializeLeft_DeserializeRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "hello", Right: 42}
	l := lr.DeserializeLeft()
	if l == nil {
		t.Errorf("expected non-nil")
	}
	r := lr.DeserializeRight()
	if r == nil {
		t.Errorf("expected non-nil")
	}
	var nilLR *coredynamic.LeftRight
	if nilLR.DeserializeLeft() != nil {
		t.Errorf("expected nil")
	}
	if nilLR.DeserializeRight() != nil {
		t.Errorf("expected nil")
	}
}

func Test_C31_153_LeftRight_LeftReflectSet_RightReflectSet(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "hello", Right: "world"}
	var leftTarget string
	err := lr.LeftReflectSet(&leftTarget)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var rightTarget string
	err = lr.RightReflectSet(&rightTarget)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C31_154_LeftRight_ReflectSet_NilReceiver(t *testing.T) {
	var lr *coredynamic.LeftRight
	if lr.LeftReflectSet(nil) != nil {
		t.Errorf("expected nil for nil receiver")
	}
	if lr.RightReflectSet(nil) != nil {
		t.Errorf("expected nil for nil receiver")
	}
}

func Test_C31_155_LeftRight_TypeStatus(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	if !ts.IsSame {
		t.Errorf("expected same type")
	}
	var nilLR *coredynamic.LeftRight
	ts2 := nilLR.TypeStatus()
	_ = ts2
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_156_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("test msg")
	c := ds.Clone()
	if c.Message != "test msg" {
		t.Errorf("expected test msg")
	}
}
func Test_C31_158_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	if ds.Message != "" {
		t.Errorf("expected empty message")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionTypes — factory functions
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_159_CollectionTypes_Factories(t *testing.T) {
	sc := coredynamic.NewStringCollection(2)
	sc.Add("a")
	if sc.Length() != 1 {
		t.Errorf("expected 1")
	}
	es := coredynamic.EmptyStringCollection()
	if es.Length() != 0 {
		t.Errorf("expected 0")
	}
	ic := coredynamic.NewIntCollection(2)
	ic.Add(1)
	if ic.Length() != 1 {
		t.Errorf("expected 1")
	}
	eic := coredynamic.EmptyIntCollection()
	if eic.Length() != 0 {
		t.Errorf("expected 0")
	}
	i64c := coredynamic.NewInt64Collection(2)
	i64c.Add(int64(1))
	bc := coredynamic.NewByteCollection(2)
	bc.Add(byte(1))
	boolC := coredynamic.NewBoolCollection(2)
	boolC.Add(true)
	f64c := coredynamic.NewFloat64Collection(2)
	f64c.Add(1.0)
	amc := coredynamic.NewAnyMapCollection(2)
	amc.Add(map[string]any{"a": 1})
	smc := coredynamic.NewStringMapCollection(2)
	smc.Add(map[string]string{"a": "b"})
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionFrom / CollectionClone
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_160_CollectionFrom_NilSlice(t *testing.T) {
	c := coredynamic.CollectionFrom[int](nil)
	if c.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_161_CollectionClone(t *testing.T) {
	original := []int{1, 2, 3}
	c := coredynamic.CollectionClone(original)
	original[0] = 99
	if c.At(0) != 1 {
		t.Errorf("expected deep copy, got %d", c.At(0))
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Collection AddWithWgLock
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_162_CollectionLock_AddWithWgLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, 42)
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_163_DynamicCollection_Strings_String(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.Add(coredynamic.NewDynamicValid("a"))
	dc.Add(coredynamic.NewDynamicValid("b"))
	strs := dc.Strings()
	if len(strs) != 2 {
		t.Errorf("expected 2")
	}
	s := dc.String()
	if s == "" {
		t.Errorf("expected non-empty")
	}
}

func Test_C31_164_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAnyNonNull(nil, true)
	if dc.Length() != 0 {
		t.Errorf("expected 0 for nil")
	}
	dc.AddAnyNonNull("hello", true)
	if dc.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C31_165_DynamicCollection_AddPtr_AddManyPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	d1 := coredynamic.NewDynamicPtr("a", true)
	dc.AddPtr(d1)
	dc.AddPtr(nil)
	if dc.Length() != 1 {
		t.Errorf("expected 1")
	}
	d2 := coredynamic.NewDynamicPtr("b", true)
	dc.AddManyPtr(d2, nil)
	if dc.Length() != 2 {
		t.Errorf("expected 2")
	}
}

func Test_C31_166_DynamicCollection_RemoveAt(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	if !dc.RemoveAt(1) {
		t.Errorf("expected success")
	}
	if dc.RemoveAt(99) {
		t.Errorf("expected false")
	}
}

func Test_C31_167_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(42, true)
	strs := dc.ListStrings()
	if len(strs) != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C31_168_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	items := dc.AnyItems()
	if len(items) != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyDynamicCollection()
	if len(empty.AnyItems()) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_169_DynamicCollection_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	ac := dc.AnyItemsCollection()
	if ac.Length() != 1 {
		t.Errorf("expected 1")
	}
	empty := coredynamic.EmptyDynamicCollection()
	if empty.AnyItemsCollection().Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_170_DynamicCollection_AddAnyMany(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany("a", "b", "c")
	if dc.Length() != 3 {
		t.Errorf("expected 3")
	}
	dc.AddAnyMany()
}

func Test_C31_171_DynamicCollection_JsonModel_JsonModelAny(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	m := dc.JsonModel()
	if len(m.Items) != 1 {
		t.Errorf("expected 1 item in model")
	}
	ma := dc.JsonModelAny()
	if ma == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C31_172_DynamicCollection_Json_JsonPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	j := dc.Json()
	_ = j
	jp := dc.JsonPtr()
	if jp == nil {
		t.Errorf("expected non-nil")
	}
}

func Test_C31_173_DynamicCollection_Paging(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagesSize(3)
	if pages != 4 {
		t.Errorf("expected 4 pages, got %d", pages)
	}
	if dc.GetPagesSize(0) != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_174_DynamicCollection_Items_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	items := dc.Items()
	if len(items) != 0 {
		t.Errorf("expected empty")
	}
}

func Test_C31_175_DynamicCollection_Length_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	if dc.Length() != 0 {
		t.Errorf("expected 0")
	}
}

func Test_C31_176_DynamicCollection_IsEmpty_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	if !dc.IsEmpty() {
		t.Errorf("expected true")
	}
}

func Test_C31_177_DynamicCollection_Loop(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("a", true).AddAny("b", true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C31_178_DynamicCollection_Loop_Break(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C31_179_DynamicCollection_HasIndex(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	if !dc.HasIndex(0) {
		t.Errorf("expected true")
	}
	if dc.HasIndex(1) {
		t.Errorf("expected false")
	}
}
