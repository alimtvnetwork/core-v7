package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════════════════════════════
// DynamicGetters — value extraction, type checks
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_01_Dynamic_Data(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.Data() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_02_Dynamic_Value(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	if d.Value() != "hello" {
		t.Error("expected hello")
	}
}

func Test_C34_03_Dynamic_Length_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_04_Dynamic_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	if d.Length() != 3 {
		t.Errorf("expected 3, got %d", d.Length())
	}
}

func Test_C34_05_Dynamic_StructStringPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	ptr := d.StructStringPtr()
	if ptr == nil || *ptr == "" {
		t.Error("expected non-empty string ptr")
	}
	// second call should return cached
	ptr2 := d.StructStringPtr()
	if ptr != ptr2 {
		t.Error("expected cached pointer")
	}
}

func Test_C34_06_Dynamic_StructStringPtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.StructStringPtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C34_07_Dynamic_String(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C34_08_Dynamic_String_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C34_09_Dynamic_StructString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.StructString() != "" {
		t.Error("expected empty")
	}
}

func Test_C34_10_Dynamic_IsNull(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	if !d.IsNull() {
		t.Error("expected null")
	}
}

func Test_C34_11_Dynamic_IsValid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if !d.IsValid() {
		t.Error("expected valid")
	}
}

func Test_C34_12_Dynamic_IsInvalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	if !d.IsInvalid() {
		t.Error("expected invalid")
	}
}

func Test_C34_13_Dynamic_IsPointer_True(t *testing.T) {
	val := 42
	d := coredynamic.NewDynamicValid(&val)
	if !d.IsPointer() {
		t.Error("expected pointer")
	}
}

func Test_C34_14_Dynamic_IsPointer_False(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.IsPointer() {
		t.Error("expected non-pointer")
	}
}

func Test_C34_15_Dynamic_IsPointer_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsPointer() {
		t.Error("expected false for nil")
	}
}

func Test_C34_16_Dynamic_IsValueType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if !d.IsValueType() {
		t.Error("expected value type")
	}
}

func Test_C34_17_Dynamic_IsValueType_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsValueType() {
		t.Error("expected false for nil")
	}
}

func Test_C34_18_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	if !d.IsStructStringNullOrEmpty() {
		t.Error("expected true for nil data")
	}
}

func Test_C34_19_Dynamic_IsStructStringNullOrEmpty_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if !d.IsStructStringNullOrEmpty() {
		t.Error("expected true for nil receiver")
	}
}

func Test_C34_20_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	if !d.IsStructStringNullOrEmptyOrWhitespace() {
		t.Error("expected true")
	}
}

func Test_C34_21_Dynamic_IsStructStringNullOrEmptyOrWhitespace_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if !d.IsStructStringNullOrEmptyOrWhitespace() {
		t.Error("expected true for nil")
	}
}

func Test_C34_22_Dynamic_IsPrimitive(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if !d.IsPrimitive() {
		t.Error("expected primitive")
	}
}

func Test_C34_23_Dynamic_IsPrimitive_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsPrimitive() {
		t.Error("expected false for nil")
	}
}

func Test_C34_24_Dynamic_IsNumber(t *testing.T) {
	d := coredynamic.NewDynamicValid(3.14)
	if !d.IsNumber() {
		t.Error("expected number")
	}
}

func Test_C34_25_Dynamic_IsNumber_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsNumber() {
		t.Error("expected false")
	}
}

func Test_C34_26_Dynamic_IsStringType(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	if !d.IsStringType() {
		t.Error("expected true")
	}
}

func Test_C34_27_Dynamic_IsStringType_NotString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.IsStringType() {
		t.Error("expected false")
	}
}

func Test_C34_28_Dynamic_IsStringType_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsStringType() {
		t.Error("expected false")
	}
}

func Test_C34_29_Dynamic_IsStruct(t *testing.T) {
	type s struct{ X int }
	d := coredynamic.NewDynamicValid(s{X: 1})
	if !d.IsStruct() {
		t.Error("expected struct")
	}
}

func Test_C34_30_Dynamic_IsStruct_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsStruct() {
		t.Error("expected false")
	}
}

func Test_C34_31_Dynamic_IsFunc(t *testing.T) {
	d := coredynamic.NewDynamicValid(func() {})
	if !d.IsFunc() {
		t.Error("expected func")
	}
}

func Test_C34_32_Dynamic_IsFunc_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsFunc() {
		t.Error("expected false")
	}
}

func Test_C34_33_Dynamic_IsSliceOrArray(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1})
	if !d.IsSliceOrArray() {
		t.Error("expected true")
	}
}

func Test_C34_34_Dynamic_IsSliceOrArray_Array(t *testing.T) {
	d := coredynamic.NewDynamicValid([2]int{1, 2})
	if !d.IsSliceOrArray() {
		t.Error("expected true for array")
	}
}

func Test_C34_35_Dynamic_IsSliceOrArray_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsSliceOrArray() {
		t.Error("expected false")
	}
}

func Test_C34_36_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	if !d.IsSliceOrArrayOrMap() {
		t.Error("expected true")
	}
}

func Test_C34_37_Dynamic_IsSliceOrArrayOrMap_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsSliceOrArrayOrMap() {
		t.Error("expected false")
	}
}

func Test_C34_38_Dynamic_IsMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	if !d.IsMap() {
		t.Error("expected true")
	}
}

func Test_C34_39_Dynamic_IsMap_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.IsMap() {
		t.Error("expected false")
	}
}

func Test_C34_40_Dynamic_IntDefault_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	v, ok := d.IntDefault(0)
	if !ok || v != 42 {
		t.Errorf("expected 42/true, got %d/%v", v, ok)
	}
}

func Test_C34_41_Dynamic_IntDefault_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("not a number")
	v, ok := d.IntDefault(99)
	if ok || v != 99 {
		t.Errorf("expected 99/false, got %d/%v", v, ok)
	}
}

func Test_C34_42_Dynamic_IntDefault_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	v, ok := d.IntDefault(5)
	if ok || v != 5 {
		t.Error("expected default for nil")
	}
}

func Test_C34_43_Dynamic_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(3.14)
	v, err := d.Float64()
	if err != nil || v != 3.14 {
		t.Errorf("expected 3.14, got %v err=%v", v, err)
	}
}

func Test_C34_44_Dynamic_Float64_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("abc")
	_, err := d.Float64()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_45_Dynamic_Float64_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Float64()
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_C34_46_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.ValueInt() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_47_Dynamic_ValueInt_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("str")
	v := d.ValueInt()
	if v == 42 {
		t.Error("should not be 42")
	}
}

func Test_C34_48_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(10))
	if d.ValueUInt() != 10 {
		t.Error("expected 10")
	}
}

func Test_C34_49_Dynamic_ValueUInt_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("str")
	if d.ValueUInt() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_50_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	v := d.ValueStrings()
	if len(v) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_51_Dynamic_ValueStrings_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.ValueStrings() != nil {
		t.Error("expected nil")
	}
}

func Test_C34_52_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	if !d.ValueBool() {
		t.Error("expected true")
	}
}

func Test_C34_53_Dynamic_ValueBool_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.ValueBool() {
		t.Error("expected false")
	}
}

func Test_C34_54_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(100))
	if d.ValueInt64() != 100 {
		t.Error("expected 100")
	}
}

func Test_C34_55_Dynamic_ValueInt64_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	v := d.ValueInt64()
	_ = v // just ensure no panic
}

func Test_C34_56_Dynamic_ValueNullErr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ValueNullErr()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_57_Dynamic_ValueNullErr_NullData(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	err := d.ValueNullErr()
	if err == nil {
		t.Error("expected error for null data")
	}
}

func Test_C34_58_Dynamic_ValueNullErr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	err := d.ValueNullErr()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C34_59_Dynamic_ValueString_String(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	if d.ValueString() != "hello" {
		t.Error("expected hello")
	}
}

func Test_C34_60_Dynamic_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.ValueString()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C34_61_Dynamic_ValueString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.ValueString() != "" {
		t.Error("expected empty")
	}
}

func Test_C34_62_Dynamic_Bytes_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid([]byte{1, 2, 3})
	b, ok := d.Bytes()
	if !ok || len(b) != 3 {
		t.Error("expected bytes")
	}
}

func Test_C34_63_Dynamic_Bytes_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	_, ok := d.Bytes()
	if ok {
		t.Error("expected false")
	}
}

func Test_C34_64_Dynamic_Bytes_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	if ok || b != nil {
		t.Error("expected nil/false")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicJson — serialization/deserialization
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_65_Dynamic_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`{}`))
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_C34_66_Dynamic_ValueMarshal(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.ValueMarshal()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_67_Dynamic_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_68_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b := d.JsonPayloadMust()
	if len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_69_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	b, err := d.JsonBytesPtr()
	if err != nil {
		t.Error("expected no error for null")
	}
	if len(b) != 0 {
		t.Error("expected empty bytes")
	}
}

func Test_C34_70_Dynamic_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.JsonBytesPtr()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_71_Dynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicValid("test")
	b, err := d.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_72_Dynamic_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`42`))
	if err == nil {
		t.Error("expected error for nil")
	}
}

func Test_C34_73_Dynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if d.JsonModel() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_74_Dynamic_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	if d.JsonModelAny() != "x" {
		t.Error("expected x")
	}
}

func Test_C34_75_Dynamic_Json(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	j := d.Json()
	_ = j
}

func Test_C34_76_Dynamic_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	if jp == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_77_Dynamic_JsonBytes(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.JsonBytes()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_78_Dynamic_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s, err := d.JsonString()
	if err != nil || s == "" {
		t.Error("expected json string")
	}
}

func Test_C34_79_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.JsonStringMust()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C34_80_Dynamic_ParseInjectUsingJson(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)
	_, err := d2.ParseInjectUsingJson(jp)
	if err == nil {
		t.Error("expected error from Dynamic json round-trip with untyped destination")
	}
}

func Test_C34_81_Dynamic_ParseInjectUsingJsonMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)

	didPanic := false
	func() {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()
		_ = d2.ParseInjectUsingJsonMust(jp)
	}()

	if !didPanic {
		t.Error("expected panic from ParseInjectUsingJsonMust when unmarshal fails")
	}
}

func Test_C34_82_Dynamic_JsonParseSelfInject(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)
	err := d2.JsonParseSelfInject(jp)
	if err == nil {
		t.Error("expected error from JsonParseSelfInject with untyped destination")
	}
}

func Test_C34_83_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	c := d.Clone()
	if c.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_84_Dynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	cp := d.ClonePtr()
	if cp == nil || cp.Value() != 42 {
		t.Error("expected cloned ptr")
	}
}

func Test_C34_85_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	if d.ClonePtr() != nil {
		t.Error("expected nil")
	}
}

func Test_C34_86_Dynamic_NonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	np := d.NonPtr()
	if np.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_87_Dynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	p := d.Ptr()
	if p == nil {
		t.Error("expected non-nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Collection[T] — base, Lock, Search, Distinct, GroupBy, Map
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_88_CollectionFrom(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C34_89_CollectionFrom_Nil(t *testing.T) {
	c := coredynamic.CollectionFrom[int](nil)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_90_CollectionClone(t *testing.T) {
	c := coredynamic.CollectionClone([]string{"a", "b"})
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_91_Collection_FirstOrDefault_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	v, ok := c.FirstOrDefault()
	if ok || v != nil {
		t.Error("expected nil/false")
	}
}

func Test_C34_92_Collection_FirstOrDefault_NonEmpty(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	v, ok := c.FirstOrDefault()
	if !ok || *v != 42 {
		t.Error("expected 42/true")
	}
}

func Test_C34_93_Collection_LastOrDefault_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	v, ok := c.LastOrDefault()
	if ok || v != nil {
		t.Error("expected nil/false")
	}
}

func Test_C34_94_Collection_LastOrDefault_NonEmpty(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	v, ok := c.LastOrDefault()
	if !ok || *v != 2 {
		t.Error("expected 2/true")
	}
}

func Test_C34_95_Collection_Items_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	items := c.Items()
	if len(items) != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_96_Collection_Count(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if c.Count() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_97_Collection_HasAnyItem(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	if c.HasAnyItem() {
		t.Error("expected false")
	}
	c.Add(1)
	if !c.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_C34_98_Collection_HasIndex(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if !c.HasIndex(1) {
		t.Error("expected true")
	}
	if c.HasIndex(2) {
		t.Error("expected false")
	}
	if c.HasIndex(-1) {
		t.Error("expected false for -1")
	}
}

func Test_C34_99_Collection_Skip(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	s := c.Skip(1)
	if len(s) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_100_Collection_Take(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	s := c.Take(2)
	if len(s) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_101_Collection_Limit(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if len(c.Limit(1)) != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_102_Collection_SkipCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sc := c.SkipCollection(2)
	if sc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_103_Collection_TakeCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	tc := c.TakeCollection(2)
	if tc.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_104_Collection_LimitCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	lc := c.LimitCollection(1)
	if lc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_105_Collection_SafeLimitCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	lc := c.SafeLimitCollection(100)
	if lc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_106_Collection_AddMany(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddMany(1, 2, 3)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C34_107_Collection_AddNonNil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	val := 42
	c.AddNonNil(&val)
	c.AddNonNil(nil)
	if c.Length() != 1 {
		t.Errorf("expected 1, got %d", c.Length())
	}
}

func Test_C34_108_Collection_RemoveAt_Valid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	ok := c.RemoveAt(1)
	if !ok || c.Length() != 2 {
		t.Error("expected success")
	}
}

func Test_C34_109_Collection_RemoveAt_Invalid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	if c.RemoveAt(5) {
		t.Error("expected false")
	}
}

func Test_C34_110_Collection_Clear(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	c.Clear()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_111_Collection_Dispose(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.Dispose()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_112_Collection_Loop(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sum := 0
	c.Loop(func(i int, v int) bool {
		sum += v
		return false
	})
	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}

func Test_C34_113_Collection_Loop_Break(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	count := 0
	c.Loop(func(i int, v int) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C34_114_Collection_Loop_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.Loop(func(i int, v int) bool {
		t.Error("should not be called")
		return false
	})
}

func Test_C34_115_Collection_LoopAsync(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	var mu sync.Mutex
	sum := 0
	c.LoopAsync(func(i int, v int) {
		mu.Lock()
		sum += v
		mu.Unlock()
	})
	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}

func Test_C34_116_Collection_LoopAsync_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, v int) {
		t.Error("should not be called")
	})
}

func Test_C34_117_Collection_Filter(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3).Add(4)
	evens := c.Filter(func(v int) bool { return v%2 == 0 })
	if evens.Length() != 2 {
		t.Errorf("expected 2, got %d", evens.Length())
	}
}

func Test_C34_118_Collection_Filter_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := c.Filter(func(v int) bool { return true })
	if result.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_119_Collection_GetPagesSize(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	if c.GetPagesSize(3) != 4 {
		t.Error("expected 4")
	}
}

func Test_C34_120_Collection_GetPagesSize_Zero(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	if c.GetPagesSize(0) != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_121_Collection_GetSinglePageCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	page := c.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C34_122_Collection_GetSinglePageCollection_Small(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	page := c.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Error("expected same")
	}
}

func Test_C34_123_Collection_GetPagedCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	pages := c.GetPagedCollection(3)
	if len(pages) != 4 {
		t.Errorf("expected 4, got %d", len(pages))
	}
}

func Test_C34_124_Collection_GetPagedCollection_Small(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	pages := c.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_125_Collection_MarshalJSON(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	b, err := c.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_126_Collection_UnmarshalJSON(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	err := c.UnmarshalJSON([]byte(`[1,2,3]`))
	if err != nil || c.Length() != 3 {
		t.Error("expected 3 items")
	}
}

func Test_C34_127_Collection_JsonString(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s, err := c.JsonString()
	if err != nil || s == "" {
		t.Error("expected json string")
	}
}

func Test_C34_128_Collection_JsonStringMust(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s := c.JsonStringMust()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C34_129_Collection_Strings(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	strs := c.Strings()
	if len(strs) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_130_Collection_String(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s := c.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionLock — thread-safe variants
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_131_Collection_LengthLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if c.LengthLock() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_132_Collection_IsEmptyLock(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	if !c.IsEmptyLock() {
		t.Error("expected true")
	}
}

func Test_C34_133_Collection_AddLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddLock(42)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_134_Collection_AddsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddsLock(1, 2, 3)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C34_135_Collection_AddManyLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyLock(1, 2)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_136_Collection_AddCollectionLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c2 := coredynamic.NewCollection[int](4)
	c2.Add(2).Add(3)
	c.AddCollectionLock(c2)
	if c.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C34_137_Collection_AddCollectionLock_Nil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.AddCollectionLock(nil)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_138_Collection_AddCollectionsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c2 := coredynamic.NewCollection[int](4)
	c2.Add(1)
	c3 := coredynamic.NewCollection[int](4)
	c3.Add(2)
	c.AddCollectionsLock(c2, nil, c3)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_139_Collection_AddIfLock_True(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(true, 42)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_140_Collection_AddIfLock_False(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(false, 42)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_141_Collection_RemoveAtLock_Valid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if !c.RemoveAtLock(0) {
		t.Error("expected true")
	}
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_142_Collection_RemoveAtLock_Invalid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	if c.RemoveAtLock(5) {
		t.Error("expected false")
	}
}

func Test_C34_143_Collection_ClearLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	c.ClearLock()
	if c.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_144_Collection_ItemsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	items := c.ItemsLock()
	if len(items) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_145_Collection_FirstLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	if c.FirstLock() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_146_Collection_LastLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if c.LastLock() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_147_Collection_AddWithWgLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, 42)
	wg.Wait()
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_148_Collection_LoopLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	sum := 0
	c.LoopLock(func(i int, v int) bool {
		sum += v
		return false
	})
	if sum != 3 {
		t.Errorf("expected 3, got %d", sum)
	}
}

func Test_C34_149_Collection_LoopLock_Break(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	count := 0
	c.LoopLock(func(i int, v int) bool {
		count++
		return true
	})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

func Test_C34_150_Collection_FilterLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	result := c.FilterLock(func(v int) bool { return v > 1 })
	if result.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_151_Collection_StringsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	strs := c.StringsLock()
	if len(strs) != 2 {
		t.Error("expected 2")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSearch — Contains, IndexOf, Has, HasAll, LastIndexOf, Count
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_152_Contains(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	if !coredynamic.Contains(c, 2) {
		t.Error("expected true")
	}
	if coredynamic.Contains(c, 99) {
		t.Error("expected false")
	}
}

func Test_C34_153_IndexOf(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("c")
	if coredynamic.IndexOf(c, "b") != 1 {
		t.Error("expected 1")
	}
	if coredynamic.IndexOf(c, "z") != -1 {
		t.Error("expected -1")
	}
}

func Test_C34_154_Has(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	if !coredynamic.Has(c, 42) {
		t.Error("expected true")
	}
}

func Test_C34_155_HasAll(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	if !coredynamic.HasAll(c, 1, 3) {
		t.Error("expected true")
	}
	if coredynamic.HasAll(c, 1, 99) {
		t.Error("expected false")
	}
}

func Test_C34_156_HasAll_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	if coredynamic.HasAll(c, 1) {
		t.Error("expected false for empty")
	}
}

func Test_C34_157_LastIndexOf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	if coredynamic.LastIndexOf(c, 1) != 2 {
		t.Error("expected 2")
	}
	if coredynamic.LastIndexOf(c, 99) != -1 {
		t.Error("expected -1")
	}
}

func Test_C34_158_Count(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	if coredynamic.Count(c, 1) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_159_ContainsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	if !coredynamic.ContainsLock(c, 42) {
		t.Error("expected true")
	}
}

func Test_C34_160_IndexOfLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if coredynamic.IndexOfLock(c, 2) != 1 {
		t.Error("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_161_Distinct(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1).Add(3).Add(2)
	d := coredynamic.Distinct(c)
	if d.Length() != 3 {
		t.Errorf("expected 3, got %d", d.Length())
	}
}

func Test_C34_162_Distinct_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	d := coredynamic.Distinct(c)
	if d.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_163_Unique(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("a")
	u := coredynamic.Unique(c)
	if u.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_164_DistinctLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(1).Add(2)
	d := coredynamic.DistinctLock(c)
	if d.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_165_DistinctCount(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	if coredynamic.DistinctCount(c) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_166_DistinctCount_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	if coredynamic.DistinctCount(c) != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_167_IsDistinct_True(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	if !coredynamic.IsDistinct(c) {
		t.Error("expected true")
	}
}

func Test_C34_168_IsDistinct_False(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(1)
	if coredynamic.IsDistinct(c) {
		t.Error("expected false")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_169_GroupBy(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3).Add(4)
	groups := coredynamic.GroupBy(c, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if len(groups) != 2 {
		t.Errorf("expected 2 groups, got %d", len(groups))
	}
	if groups["even"].Length() != 2 {
		t.Error("expected 2 evens")
	}
}

func Test_C34_170_GroupBy_Nil(t *testing.T) {
	groups := coredynamic.GroupBy[int, string](nil, func(v int) string { return "" })
	if len(groups) != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_171_GroupBy_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	groups := coredynamic.GroupBy(c, func(v int) string { return "" })
	if len(groups) != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_172_GroupByLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	groups := coredynamic.GroupByLock(c, func(v int) int { return v % 2 })
	if len(groups) != 2 {
		t.Error("expected 2 groups")
	}
}

func Test_C34_173_GroupByCount(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("a")
	counts := coredynamic.GroupByCount(c, func(s string) string { return s })
	if counts["a"] != 2 {
		t.Errorf("expected 2, got %d", counts["a"])
	}
}

func Test_C34_174_GroupByCount_Nil(t *testing.T) {
	counts := coredynamic.GroupByCount[string, string](nil, func(s string) string { return s })
	if len(counts) != 0 {
		t.Error("expected empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMap — Map, FlatMap, Reduce
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_175_Map(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	result := coredynamic.Map(c, func(v int) int { return v * 2 })
	if result.Length() != 3 || result.At(0) != 2 {
		t.Error("expected doubled")
	}
}

func Test_C34_176_Map_Nil(t *testing.T) {
	result := coredynamic.Map[int, int](nil, func(v int) int { return v })
	if result.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_177_Map_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := coredynamic.Map(c, func(v int) int { return v })
	if result.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_178_FlatMap(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a,b").Add("c,d")
	result := coredynamic.FlatMap(c, func(s string) []string {
		return []string{s + "1", s + "2"}
	})
	if result.Length() != 4 {
		t.Errorf("expected 4, got %d", result.Length())
	}
}

func Test_C34_179_FlatMap_Nil(t *testing.T) {
	result := coredynamic.FlatMap[int, int](nil, func(v int) []int { return nil })
	if result.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C34_180_Reduce(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sum := coredynamic.Reduce(c, 0, func(acc int, v int) int { return acc + v })
	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}

func Test_C34_181_Reduce_Nil(t *testing.T) {
	result := coredynamic.Reduce[int, int](nil, 42, func(acc int, v int) int { return acc + v })
	if result != 42 {
		t.Error("expected initial value")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// AnyCollection — extended methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_182_AnyCollection_AtAsDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	d := ac.AtAsDynamic(0)
	if d.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C34_183_AnyCollection_DynamicItems(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	items := ac.DynamicItems()
	if len(items) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_184_AnyCollection_DynamicItems_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	items := ac.DynamicItems()
	if len(items) != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_185_AnyCollection_DynamicCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	dc := ac.DynamicCollection()
	if dc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_186_AnyCollection_DynamicCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	dc := ac.DynamicCollection()
	if dc.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_187_AnyCollection_ReflectSetAt(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	var target int
	err := ac.ReflectSetAt(0, &target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C34_188_AnyCollection_ListStrings(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("hello").Add(42)
	strs := ac.ListStrings(false)
	if len(strs) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_189_AnyCollection_ListStringsPtr(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("x")
	strs := ac.ListStringsPtr(true)
	if len(strs) != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_190_AnyCollection_Loop_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})
	if count != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_191_AnyCollection_Loop_Sync_Break(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return true
	})
	if count != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_192_AnyCollection_Loop_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.Loop(true, func(i int, item any) bool {
		return false
	})
	// just ensure no panic
}

func Test_C34_193_AnyCollection_Loop_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Loop(false, func(i int, item any) bool {
		t.Error("should not be called")
		return false
	})
}

func Test_C34_194_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		count++
		return false
	})
	if count != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_195_AnyCollection_LoopDynamic_Sync_Break(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		count++
		return true
	})
	if count != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_196_AnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.LoopDynamic(true, func(i int, d coredynamic.Dynamic) bool {
		return false
	})
}

func Test_C34_197_AnyCollection_LoopDynamic_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		t.Error("should not be called")
		return false
	})
}

func Test_C34_198_AnyCollection_AddAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAny(42, true)
	if ac.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_199_AnyCollection_AddNonNull(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddNonNull(nil)
	ac.AddNonNull(42)
	if ac.Length() != 1 {
		t.Errorf("expected 1, got %d", ac.Length())
	}
}

func Test_C34_200_AnyCollection_AddNonNullDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddNonNullDynamic(nil, true)
	ac.AddNonNullDynamic(42, true)
	if ac.Length() != 1 {
		t.Errorf("expected 1, got %d", ac.Length())
	}
}

func Test_C34_201_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic(1, 2, 3)
	if ac.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C34_202_AnyCollection_AddAnyManyDynamic_Nil(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic(nil)
	// nil variadic should skip
}

func Test_C34_203_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem([]int{1, 2, 3})
	if ac.Length() != 3 {
		t.Errorf("expected 3, got %d", ac.Length())
	}
}

func Test_C34_204_AnyCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem(nil)
	if ac.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_205_AnyCollection_AddMany(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany(1, nil, 3)
	if ac.Length() != 2 {
		t.Errorf("expected 2 (nil skipped), got %d", ac.Length())
	}
}

func Test_C34_206_AnyCollection_AddMany_NilVariadic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany(nil)
	// nil variadic should skip
}

func Test_C34_207_AnyCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(0), 42)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C34_208_AnyCollection_AddAnyWithTypeValidation_Invalid(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(0), "str")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_209_AnyCollection_AddAnyItemsWithTypeValidation_Continue(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), 1, "bad", 3)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_210_AnyCollection_AddAnyItemsWithTypeValidation_Stop(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0), 1, "bad", 3)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_211_AnyCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0))
	if err != nil {
		t.Error("expected nil")
	}
}

func Test_C34_212_AnyCollection_Paging(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	if ac.GetPagesSize(3) != 4 {
		t.Error("expected 4 pages")
	}
	if ac.GetPagesSize(0) != 0 {
		t.Error("expected 0 for zero page size")
	}
}

func Test_C34_213_AnyCollection_GetSinglePageCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	page := ac.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C34_214_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	page := ac.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Error("expected same")
	}
}

func Test_C34_215_AnyCollection_GetPagedCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(3)
	if len(pages) == 0 {
		t.Error("expected pages")
	}
}

func Test_C34_216_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	pages := ac.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_217_AnyCollection_Json(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	j := ac.Json()
	_ = j
}

func Test_C34_218_AnyCollection_JsonPtr(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	if jp == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_219_AnyCollection_JsonModel(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	m := ac.JsonModel()
	if len(m) != 1 {
		t.Error("expected 1")
	}
}

func Test_C34_220_AnyCollection_JsonModelAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	v := ac.JsonModelAny()
	if v == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_221_AnyCollection_MarshalJSON(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	b, err := ac.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C34_222_AnyCollection_UnmarshalJSON(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.UnmarshalJSON([]byte(`[1,2,3]`))
	if err != nil {
		t.Error("expected success")
	}
}

func Test_C34_223_AnyCollection_UnmarshalJSON_Bad(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.UnmarshalJSON([]byte(`not json`))
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_224_AnyCollection_JsonResultsCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	rc := ac.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_225_AnyCollection_JsonResultsCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_226_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	rc := ac.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_227_AnyCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C34_228_AnyCollection_ParseInjectUsingJson(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	_, err := ac2.ParseInjectUsingJson(jp)
	if err == nil {
		t.Error("expected error for AnyCollection JSON payload {}")
	}
}

func Test_C34_229_AnyCollection_ParseInjectUsingJson_Bad(t *testing.T) {
	badJson := corejson.NewPtr("not an any collection")
	ac := coredynamic.NewAnyCollection(4)
	_, err := ac.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C34_230_AnyCollection_ParseInjectUsingJsonMust(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	panicked := false

	func() {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()

		_ = ac2.ParseInjectUsingJsonMust(jp)
	}()

	if !panicked {
		t.Error("expected panic for AnyCollection JSON payload {}")
	}
}

func Test_C34_231_AnyCollection_JsonParseSelfInject(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	err := ac2.JsonParseSelfInject(jp)
	if err == nil {
		t.Error("expected error for AnyCollection JSON payload {}")
	}
}

func Test_C34_232_AnyCollection_Strings(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add("x")
	strs := ac.Strings()
	if len(strs) != 2 {
		t.Error("expected 2")
	}
}

func Test_C34_233_AnyCollection_Strings_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	strs := ac.Strings()
	if len(strs) != 0 {
		t.Error("expected 0")
	}
}

func Test_C34_234_AnyCollection_String(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	s := ac.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C34_235_AnyCollection_GetPagingInfo(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	info := ac.GetPagingInfo(3, 1)
	if info.TotalPages != 4 {
		t.Error("expected 4 total pages")
	}
}
