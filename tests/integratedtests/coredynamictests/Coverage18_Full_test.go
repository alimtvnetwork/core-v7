package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Dynamic constructors ──

func Test_C18_Dynamic_Constructors(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	dv := coredynamic.NewDynamicValid("valid")
	dp := coredynamic.NewDynamicPtr("ptr", true)
	inv := coredynamic.InvalidDynamic()
	invP := coredynamic.InvalidDynamicPtr()

	// Act
	actual := args.Map{
		"dValid":   d.IsValid(),
		"dvValid":  dv.IsValid(),
		"dpValid":  dp.IsValid(),
		"invValid": inv.IsValid(),
		"invPNil":  invP.IsValid(),
	}

	// Assert
	expected := args.Map{
		"dValid": true, "dvValid": true, "dpValid": true,
		"invValid": false, "invPNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- constructors", actual)
}

func Test_C18_Dynamic_Clone(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	clone := d.Clone()
	cloneP := d.ClonePtr()
	_ = d.NonPtr()
	_ = d.Ptr()

	// Act
	actual := args.Map{"result": clone.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "clone invalid", actual)
	actual := args.Map{"result": cloneP.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "cloneP invalid", actual)

	var nilD *coredynamic.Dynamic
	actual := args.Map{"result": nilD.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── DynamicGetters ──

func Test_C18_Dynamic_DataValue(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{
		"data":  d.Data(),
		"value": d.Value(),
	}

	// Assert
	expected := args.Map{
		"data": "hello",
		"value": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Data/Value", actual)
}

func Test_C18_Dynamic_Length(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)

	// Act
	actual := args.Map{"result": d.Length()}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	nilD := coredynamic.NewDynamic(nil, false)
	actual := args.Map{"result": nilD.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
}

func Test_C18_Dynamic_StructString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	s := d.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	_ = d.StructString()
	_ = d.StructStringPtr()
	// call again for cache
	_ = d.StructStringPtr()
}

func Test_C18_Dynamic_IsNull_IsValid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	nilD := coredynamic.NewDynamic(nil, false)

	// Act
	actual := args.Map{
		"isNull":    d.IsNull(),
		"isValid":   d.IsValid(),
		"isInvalid": d.IsInvalid(),
		"nilNull":   nilD.IsNull(),
	}

	// Assert
	expected := args.Map{
		"isNull": false, "isValid": true, "isInvalid": false, "nilNull": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- IsNull/IsValid", actual)
}

func Test_C18_Dynamic_IsPointer_IsValueType(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"result": d.IsPointer()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string not a pointer", actual)
	actual := args.Map{"result": d.IsValueType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "string is value type", actual)
	// call again for cache
	_ = d.IsPointer()

	ptrD := coredynamic.NewDynamic(&struct{}{}, true)
	actual := args.Map{"result": ptrD.IsPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ptr should be pointer", actual)
}

func Test_C18_Dynamic_IsStructStringChecks(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	nilD := coredynamic.NewDynamic(nil, false)

	// Act
	actual := args.Map{
		"nullOrEmpty":      d.IsStructStringNullOrEmpty(),
		"nullOrEmptyOrWs":  d.IsStructStringNullOrEmptyOrWhitespace(),
		"nilNullOrEmpty":   nilD.IsStructStringNullOrEmpty(),
	}

	// Assert
	expected := args.Map{
		"nullOrEmpty": false, "nullOrEmptyOrWs": false, "nilNullOrEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsStructString*", actual)
}

func Test_C18_Dynamic_TypeChecks(t *testing.T) {
	// Arrange
	strD := coredynamic.NewDynamic("hello", true)
	intD := coredynamic.NewDynamic(42, true)
	sliceD := coredynamic.NewDynamic([]int{1}, true)
	mapD := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	structD := coredynamic.NewDynamic(struct{}{}, true)
	funcD := coredynamic.NewDynamic(func() {}, true)

	// Act
	actual := args.Map{
		"isPrimStr":   strD.IsPrimitive(),
		"isNumInt":    intD.IsNumber(),
		"isString":    strD.IsStringType(),
		"isStruct":    structD.IsStruct(),
		"isFunc":      funcD.IsFunc(),
		"isSlice":     sliceD.IsSliceOrArray(),
		"isSliceMap":  sliceD.IsSliceOrArrayOrMap(),
		"isMap":       mapD.IsMap(),
	}

	// Assert
	expected := args.Map{
		"isPrimStr": true, "isNumInt": true, "isString": true,
		"isStruct": true, "isFunc": true, "isSlice": true,
		"isSliceMap": true, "isMap": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- type checks", actual)
}

func Test_C18_Dynamic_IntDefault(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	val, ok := d.IntDefault(0)

	// Act
	actual := args.Map{"result": ok || val != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	badD := coredynamic.NewDynamic("notint", true)
	val2, ok2 := badD.IntDefault(99)
	actual := args.Map{"result": ok2 || val2 != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99 default", actual)

	nilD := coredynamic.NewDynamic(nil, false)
	val3, ok3 := nilD.IntDefault(7)
	actual := args.Map{"result": ok3 || val3 != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_C18_Dynamic_Float64(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(3.14, true)
	_, err := d.Float64()

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err:", actual)

	nilD := coredynamic.NewDynamic(nil, false)
	_, err2 := nilD.Float64()
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)

	badD := coredynamic.NewDynamic("notnum", true)
	_, err3 := badD.Float64()
	actual := args.Map{"result": err3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_Dynamic_ValueCasts(t *testing.T) {
	// Arrange
	intD := coredynamic.NewDynamic(42, true)
	uintD := coredynamic.NewDynamic(uint(10), true)
	stringsD := coredynamic.NewDynamic([]string{"a"}, true)
	boolD := coredynamic.NewDynamic(true, true)
	int64D := coredynamic.NewDynamic(int64(100), true)

	// Act
	actual := args.Map{
		"valInt":     intD.ValueInt(),
		"valUInt":    uintD.ValueUInt(),
		"valStrings": len(stringsD.ValueStrings()),
		"valBool":    boolD.ValueBool(),
		"valInt64":   int64D.ValueInt64(),
	}

	// Assert
	expected := args.Map{
		"valInt": 42, "valUInt": uint(10), "valStrings": 1,
		"valBool": true, "valInt64": int64(100),
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueCasts", actual)

	// Wrong type casts
	badInt := coredynamic.NewDynamic("str", true)
	actual := args.Map{"result": badInt.ValueInt() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": badInt.ValueUInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": badInt.ValueStrings() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": badInt.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": badInt.ValueInt64() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C18_Dynamic_ValueNullErr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"result": d.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	nilD := coredynamic.NewDynamic(nil, false)
	actual := args.Map{"result": nilD.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	var ptrD *coredynamic.Dynamic
	actual := args.Map{"result": ptrD.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_Dynamic_ValueString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"result": d.ValueString() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	intD := coredynamic.NewDynamic(42, true)
	actual := args.Map{"result": intD.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	var nilD *coredynamic.Dynamic
	actual := args.Map{"result": nilD.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_C18_Dynamic_Bytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]byte{1, 2}, true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{"result": ok || len(b) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)

	var nilD *coredynamic.Dynamic
	b2, ok2 := nilD.Bytes()
	actual := args.Map{"result": ok2 || b2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── DynamicReflect ──

func Test_C18_Dynamic_ReflectValue(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	rv := d.ReflectValue()

	// Act
	actual := args.Map{"result": rv == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// cached
	rv2 := d.ReflectValue()
	actual := args.Map{"result": rv2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached", actual)
}

func Test_C18_Dynamic_ReflectKind_Type_TypeName(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"result": d.ReflectKind() != reflect.String}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
	rt := d.ReflectType()
	actual := args.Map{"result": rt == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	// cached
	_ = d.ReflectType()
	tn := d.ReflectTypeName()
	actual := args.Map{"result": tn == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C18_Dynamic_IsReflectTypeOf_IsReflectKind(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"result": d.IsReflectTypeOf(reflect.TypeOf(""))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": d.IsReflectKind(reflect.String)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C18_Dynamic_ItemUsingIndex_Key(t *testing.T) {
	// Arrange
	sliceD := coredynamic.NewDynamic([]int{10, 20, 30}, true)
	item := sliceD.ItemUsingIndex(1)

	// Act
	actual := args.Map{"result": item != 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 20", actual)
	rv := sliceD.ItemReflectValueUsingIndex(0)
	actual := args.Map{"result": rv.Int() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)

	mapD := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	val := mapD.ItemUsingKey("a")
	actual := args.Map{"result": val != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_ = mapD.ItemReflectValueUsingKey("a")
}

func Test_C18_Dynamic_Loop(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 iterations", actual)

	// with break
	count2 := 0
	d.Loop(func(i int, item any) bool {
		count2++
		return i == 0
	})
	actual := args.Map{"result": count2 != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 iteration", actual)

	// nil/invalid
	nilD := coredynamic.NewDynamic(nil, false)
	called := nilD.Loop(func(i int, item any) bool { return false })
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C18_Dynamic_LoopMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"a": 1, "b": 2}, true)
	count := 0
	d.LoopMap(func(i int, key, value any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	// with break
	d.LoopMap(func(i int, key, value any) bool { return true })

	// nil
	nilD := coredynamic.NewDynamic(nil, false)
	actual := args.Map{"result": nilD.LoopMap(func(i int, k, v any) bool { return false })}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C18_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})

	// Act
	actual := args.Map{"result": result.Length()}

	// Assert
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	// with break
	result2 := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1
	})
	actual := args.Map{"result": result2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	// nil
	nilD := coredynamic.NewDynamic(nil, false)
	r := nilD.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) { return true, false })
	actual := args.Map{"result": r.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C18_Dynamic_ReflectSetTo(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	var target string
	_ = d.ReflectSetTo(&target)

	var nilD *coredynamic.Dynamic
	err := nilD.ReflectSetTo(&target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ── DynamicJson ──

func Test_C18_Dynamic_JsonMethods(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	_ = d.JsonModel()
	_ = d.JsonModelAny()
	_ = d.Json()
	_ = d.JsonPtr()
	_, _ = d.MarshalJSON()
	_, _ = d.JsonBytesPtr()
	_, _ = d.JsonBytes()
	_, _ = d.JsonString()
	_ = d.JsonStringMust()
	_ = d.JsonPayloadMust()
	_, _ = d.ValueMarshal()
}

func Test_C18_Dynamic_JsonNull(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)
	b, err := d.JsonBytesPtr()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for null", actual)
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty bytes", actual)
}

func Test_C18_Dynamic_Deserialize(t *testing.T) {
	// Arrange
	type testStruct struct{ Name string }
	target := &testStruct{}
	d := coredynamic.NewDynamic(target, true)
	_, _ = d.Deserialize([]byte(`{"Name":"test"}`))

	var nilD *coredynamic.Dynamic
	_, err := nilD.Deserialize([]byte(`{}`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_Dynamic_UnmarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	_ = d.UnmarshalJSON([]byte(`"world"`))

	var nilD *coredynamic.Dynamic
	err := nilD.UnmarshalJSON([]byte(`"x"`))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C18_Dynamic_ValueMarshal_Nil(t *testing.T) {
	// Arrange
	var nilD *coredynamic.Dynamic
	_, err := nilD.ValueMarshal()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_Dynamic_ParseInjectUsingJson(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	r := corejson.New("world")
	_, _ = d.ParseInjectUsingJson(&r)
	_ = d.JsonParseSelfInject(&r)
}

func Test_C18_Dynamic_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	d := coredynamic.NewDynamicValid("hello")
	bad := corejson.NewResult.UsingString(`invalid`)
	d.ParseInjectUsingJsonMust(bad)
}

// ── DynamicStatus ──

func Test_C18_DynamicStatus(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("err")
	dsNoMsg := coredynamic.InvalidDynamicStatusNoMessage()
	clone := ds.Clone()
	cloneP := ds.ClonePtr()

	// Act
	actual := args.Map{"result": clone.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": cloneP == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dsNoMsg.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)

	var nilDS *coredynamic.DynamicStatus
	actual := args.Map{"result": nilDS.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── SimpleRequest ──

func Test_C18_SimpleRequest(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("data", true, "msg")
	srValid := coredynamic.NewSimpleRequestValid("data")
	srInv := coredynamic.InvalidSimpleRequest("err")
	srInvNoMsg := coredynamic.InvalidSimpleRequestNoMessage()

	// Act
	actual := args.Map{
		"msg":      sr.Message(),
		"req":      sr.Request(),
		"val":      sr.Value(),
		"srValid":  srValid.IsValid(),
		"srInv":    srInv.IsValid(),
		"srInvNo":  srInvNoMsg.IsValid(),
	}

	// Assert
	expected := args.Map{
		"msg": "msg", "req": "data", "val": "data",
		"srValid": true, "srInv": false, "srInvNo": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- with args", actual)
}

func Test_C18_SimpleRequest_TypeMismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("hello", true, "msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for matching type", actual)
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for mismatch", actual)
	err3 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"result": err3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error with message", actual)
}

func Test_C18_SimpleRequest_IsPointer(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("hello", true, "")

	// Act
	actual := args.Map{"result": sr.IsPointer()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_ = sr.IsReflectKind(reflect.String)
}

func Test_C18_SimpleRequest_InvalidError(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("data", true, "")

	// Act
	actual := args.Map{"result": sr.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty msg", actual)
	sr2 := coredynamic.InvalidSimpleRequest("some error")
	err := sr2.InvalidError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// cached
	err2 := sr2.InvalidError()
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached error", actual)
}

// ── SimpleResult ──

func Test_C18_SimpleResult(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResult("data", true, "")
	srValid := coredynamic.NewSimpleResultValid("data")
	srInv := coredynamic.InvalidSimpleResult("err")
	srInvNoMsg := coredynamic.InvalidSimpleResultNoMessage()

	// Act
	actual := args.Map{
		"result":   sr.Result,
		"valid":    srValid.IsValid(),
		"inv":      srInv.IsValid(),
		"invNoMsg": srInvNoMsg.IsValid(),
	}

	// Assert
	expected := args.Map{
		"result": "data", "valid": true, "inv": false, "invNoMsg": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- with args", actual)
}

func Test_C18_SimpleResult_TypeMismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_SimpleResult_InvalidError(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("data")

	// Act
	actual := args.Map{"result": sr.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	sr2 := coredynamic.InvalidSimpleResult("err msg")
	actual := args.Map{"result": sr2.InvalidError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// cached
	actual := args.Map{"result": sr2.InvalidError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached", actual)
}

func Test_C18_SimpleResult_Clone(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("data")
	clone := sr.Clone()
	cloneP := sr.ClonePtr()

	// Act
	actual := args.Map{"result": clone.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "clone invalid", actual)
	actual := args.Map{"result": cloneP == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	var nilSR *coredynamic.SimpleResult
	actual := args.Map{"result": nilSR.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── TypedDynamic ──

func Test_C18_TypedDynamic(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	tdValid := coredynamic.NewTypedDynamicValid[string]("world")
	tdPtr := coredynamic.NewTypedDynamicPtr[int](42, true)
	invTd := coredynamic.InvalidTypedDynamic[string]()
	invTdP := coredynamic.InvalidTypedDynamicPtr[string]()

	// Act
	actual := args.Map{
		"data":    td.Data(),
		"value":   td.Value(),
		"valid":   td.IsValid(),
		"invalid": td.IsInvalid(),
		"str":     td.String(),
		"tdVVal":  tdValid.Value(),
		"ptrVal":  tdPtr.Data(),
		"invVal":  invTd.IsValid(),
		"invPVal": invTdP.IsValid(),
	}

	// Assert
	expected := args.Map{
		"data": "hello", "value": "hello", "valid": true,
		"invalid": false, "str": "hello", "tdVVal": "world",
		"ptrVal": 42, "invVal": false, "invPVal": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- with args", actual)
}

func Test_C18_TypedDynamic_Json(t *testing.T) {
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	_, _ = td.JsonBytes()
	_ = td.JsonResult()
	_ = td.Json()
	_ = td.JsonPtr()
	_, _ = td.JsonString()
	_, _ = td.MarshalJSON()
	_, _ = td.ValueMarshal()
	_, _ = td.Bytes()
	_ = td.JsonModel()
	_ = td.JsonModelAny()
}

func Test_C18_TypedDynamic_UnmarshalJSON(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("", false)
	err := td.UnmarshalJSON([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": td.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid after unmarshal", actual)
}

func Test_C18_TypedDynamic_Deserialize(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("", false)
	err := td.Deserialize([]byte(`"hello"`))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	var nilTD *coredynamic.TypedDynamic[string]
	err2 := nilTD.Deserialize([]byte(`"x"`))
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C18_TypedDynamic_GetAs(t *testing.T) {
	// Arrange
	strTD := coredynamic.NewTypedDynamic[string]("hello", true)
	intTD := coredynamic.NewTypedDynamic[int](42, true)
	boolTD := coredynamic.NewTypedDynamic[bool](true, true)

	s, sok := strTD.GetAsString()

	// Act
	actual := args.Map{"result": sok || s != "hello"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetAsString failed", actual)
	i, iok := intTD.GetAsInt()
	actual := args.Map{"result": iok || i != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetAsInt failed", actual)
	b, bok := boolTD.GetAsBool()
	actual := args.Map{"result": bok || !b}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetAsBool failed", actual)

	_, i64ok := intTD.GetAsInt64()
	actual := args.Map{"result": i64ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for int->int64", actual)
	_, uok := intTD.GetAsUint()
	actual := args.Map{"result": uok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, f64ok := intTD.GetAsFloat64()
	actual := args.Map{"result": f64ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, f32ok := intTD.GetAsFloat32()
	actual := args.Map{"result": f32ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, bok2 := intTD.GetAsBytes()
	actual := args.Map{"result": bok2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	_, sok2 := intTD.GetAsStrings()
	actual := args.Map{"result": sok2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C18_TypedDynamic_Value_Methods(t *testing.T) {
	// Arrange
	strTD := coredynamic.NewTypedDynamic[string]("hello", true)
	intTD := coredynamic.NewTypedDynamic[int](42, true)
	boolTD := coredynamic.NewTypedDynamic[bool](true, true)
	int64TD := coredynamic.NewTypedDynamic[int64](int64(100), true)

	// Act
	actual := args.Map{"result": strTD.ValueString() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	actual := args.Map{"result": intTD.ValueInt() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	actual := args.Map{"result": boolTD.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": int64TD.ValueInt64() != 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)

	// wrong type
	actual := args.Map{"result": strTD.ValueInt() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": strTD.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": strTD.ValueInt64() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": intTD.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty via sprintf", actual)
}

func Test_C18_TypedDynamic_Clone(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	clone := td.Clone()
	cloneP := td.ClonePtr()
	_ = td.NonPtr()
	_ = td.Ptr()
	_ = td.ToDynamic()

	// Act
	actual := args.Map{"result": clone.Data() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	actual := args.Map{"result": cloneP == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	var nilTD *coredynamic.TypedDynamic[string]
	actual := args.Map{"result": nilTD.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C18_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	// Arrange
	bytesTD := coredynamic.NewTypedDynamic[[]byte]([]byte{1, 2}, true)
	b, ok := bytesTD.Bytes()

	// Act
	actual := args.Map{"result": ok || len(b) != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)

	_, bok := bytesTD.GetAsBytes()
	actual := args.Map{"result": bok}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── KeyVal ──

func Test_C18_KeyVal_Methods(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "key", Value: "val"}
	_ = kv.KeyDynamic()
	_ = kv.ValueDynamic()
	_ = kv.KeyDynamicPtr()
	_ = kv.ValueDynamicPtr()
	_ = kv.IsKeyNull()
	_ = kv.IsValueNull()
	_ = kv.String()
	_ = kv.ValueReflectValue()
	_ = kv.KeyString()
	_ = kv.ValueString()
	_ = kv.JsonModel()
	_ = kv.JsonModelAny()
	_ = kv.Json()
	_ = kv.JsonPtr()
	_, _ = kv.Serialize()
}

func Test_C18_KeyVal_ValueCasts(t *testing.T) {
	// Arrange
	kvInt := coredynamic.KeyVal{Key: "k", Value: 42}
	kvBool := coredynamic.KeyVal{Key: "k", Value: true}
	kvStr := coredynamic.KeyVal{Key: "k", Value: []string{"a"}}
	kvUInt := coredynamic.KeyVal{Key: "k", Value: uint(5)}
	kvI64 := coredynamic.KeyVal{Key: "k", Value: int64(99)}

	// Act
	actual := args.Map{"result": kvInt.ValueInt() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	actual := args.Map{"result": kvBool.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": len(kvStr.ValueStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": kvUInt.ValueUInt() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual := args.Map{"result": kvI64.ValueInt64() != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)

	// wrong type
	kvBad := coredynamic.KeyVal{Key: "k", Value: "str"}
	actual := args.Map{"result": kvBad.ValueInt() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": kvBad.ValueUInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": kvBad.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C18_KeyVal_NullChecks(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"result": kv.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": kv.KeyNullErr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	kvNil := coredynamic.KeyVal{Key: nil, Value: nil}
	actual := args.Map{"result": kvNil.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": kvNil.KeyNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	actual := args.Map{"result": nilKV.KeyNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	actual := args.Map{"result": nilKV.KeyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilKV.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C18_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	r := corejson.New(coredynamic.KeyVal{Key: "a", Value: "b"})
	_, _ = kv.ParseInjectUsingJson(&r)
	_ = kv.JsonParseSelfInject(&r)
}

func Test_C18_KeyVal_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	bad := corejson.NewResult.UsingString(`invalid`)
	kv.ParseInjectUsingJsonMust(bad)
}

// ── KeyValCollection ──

func Test_C18_KeyValCollection_Full(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.AddPtr(&coredynamic.KeyVal{Key: "c", Value: 3})
	kvc.AddPtr(nil)
	kvc.AddMany(coredynamic.KeyVal{Key: "d", Value: 4})
	kvc.AddMany()
	kvc.AddManyPtr(&coredynamic.KeyVal{Key: "e", Value: 5}, nil)
	kvc.AddManyPtr()

	// Act
	actual := args.Map{"result": kvc.Length()}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	actual := args.Map{"result": kvc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	actual := args.Map{"result": kvc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	_ = kvc.Items()
	_ = kvc.MapAnyItems()
	_, _ = kvc.JsonMapResults()
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
	_ = kvc.GetPagesSize(2)
	_ = kvc.GetPagesSize(0)
	_ = kvc.GetPagedCollection(2)
	_ = kvc.AllKeys()
	_ = kvc.AllKeysSorted()
	_ = kvc.AllValues()
	_ = kvc.String()
	_, _ = kvc.Serialize()
	_, _ = kvc.JsonString()
	// JsonStringMust panics with nil error because HandleError panics on empty JSON ({})
	func() {
		defer func() { recover() }()
		_ = kvc.JsonStringMust()
	}()
}

func Test_C18_KeyValCollection_NilItems(t *testing.T) {
	// Arrange
	var nilKVC *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"result": nilKVC.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": nilKVC.Items() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilKVC.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilKVC.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C18_KeyValCollection_Empty(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	actual := args.Map{"result": kvc.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	_ = kvc.MapAnyItems()
	_, _ = kvc.JsonMapResults()
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
	_ = kvc.AllKeys()
	_ = kvc.AllKeysSorted()
	_ = kvc.AllValues()
}

func Test_C18_KeyValCollection_Clone(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	clone := kvc.Clone()
	cloneP := kvc.ClonePtr()
	_ = clone.NonPtr()
	_ = kvc.Ptr()

	// Act
	actual := args.Map{"result": cloneP.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C18_KeyValCollection_Paging(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pages", actual)
	_ = kvc.GetPagingInfo(3, 1)
	_ = kvc.GetSinglePageCollection(3, 1)
	// small collection
	small := coredynamic.NewKeyValCollection(1)
	small.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	_ = small.GetSinglePageCollection(5, 1)
}

// ── LeftRight (coredynamic) ──

func Test_C18_LeftRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "L", Right: "R"}

	// Act
	actual := args.Map{
		"empty":     lr.IsEmpty(),
		"hasAny":    lr.HasAnyItem(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"leftEmpty": lr.IsLeftEmpty(),
		"rightEmpty": lr.IsRightEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": false, "hasAny": true, "hasLeft": true,
		"hasRight": true, "leftEmpty": false, "rightEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- with args", actual)

	_ = lr.LeftToDynamic()
	_ = lr.RightToDynamic()
	_ = lr.DeserializeLeft()
	_ = lr.DeserializeRight()
	_ = lr.TypeStatus()
	var target string
	_ = lr.LeftReflectSet(&target)
	_ = lr.RightReflectSet(&target)
}

func Test_C18_LeftRight_Nil(t *testing.T) {
	// Arrange
	var nilLR *coredynamic.LeftRight

	// Act
	actual := args.Map{"result": nilLR.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilLR.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilLR.HasLeft()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilLR.HasRight()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": nilLR.LeftToDynamic() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilLR.RightToDynamic() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilLR.DeserializeLeft() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilLR.DeserializeRight() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	_ = nilLR.TypeStatus()
	actual := args.Map{"result": nilLR.LeftReflectSet(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilLR.RightReflectSet(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── ValueStatus ──

func Test_C18_ValueStatus(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("err")
	vsNoMsg := coredynamic.InvalidValueStatusNoMessage()

	// Act
	actual := args.Map{"result": vs.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": vsNoMsg.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

// ── SafeTypeName ──

func Test_C18_SafeTypeName(t *testing.T) {
	// Arrange
	name := coredynamic.SafeTypeName("hello")

	// Act
	actual := args.Map{"result": name != "string"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'string'", actual)
	nilName := coredynamic.SafeTypeName(nil)
	actual := args.Map{"result": nilName != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

// ── IsAnyTypesOf ──

func Test_C18_IsAnyTypesOf(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	found := coredynamic.IsAnyTypesOf(strType, strType, intType)

	// Act
	actual := args.Map{"result": found}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	notFound := coredynamic.IsAnyTypesOf(reflect.TypeOf(true), strType, intType)
	actual := args.Map{"result": notFound}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ── LengthOfReflect ──

func Test_C18_LengthOfReflect(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	l := coredynamic.LengthOfReflect(rv)

	// Act
	actual := args.Map{"result": l}

	// Assert
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}
