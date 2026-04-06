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
	d := coredynamic.NewDynamic("hello", true)
	dv := coredynamic.NewDynamicValid("valid")
	dp := coredynamic.NewDynamicPtr("ptr", true)
	inv := coredynamic.InvalidDynamic()
	invP := coredynamic.InvalidDynamicPtr()

	actual := args.Map{
		"dValid":   d.IsValid(),
		"dvValid":  dv.IsValid(),
		"dpValid":  dp.IsValid(),
		"invValid": inv.IsValid(),
		"invPNil":  invP.IsValid(),
	}
	expected := args.Map{
		"dValid": true, "dvValid": true, "dpValid": true,
		"invValid": false, "invPNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- constructors", actual)
}

func Test_C18_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	clone := d.Clone()
	cloneP := d.ClonePtr()
	_ = d.NonPtr()
	_ = d.Ptr()
	if !clone.IsValid() { t.Fatal("clone invalid") }
	if !cloneP.IsValid() { t.Fatal("cloneP invalid") }

	var nilD *coredynamic.Dynamic
	if nilD.ClonePtr() != nil { t.Fatal("expected nil") }
}

// ── DynamicGetters ──

func Test_C18_Dynamic_DataValue(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	actual := args.Map{
		"data":  d.Data(),
		"value": d.Value(),
	}
	expected := args.Map{"data": "hello", "value": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Data/Value", actual)
}

func Test_C18_Dynamic_Length(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	if d.Length() != 3 { t.Fatalf("expected 3, got %d", d.Length()) }
	nilD := coredynamic.NewDynamic(nil, false)
	if nilD.Length() != 0 { t.Fatal("expected 0 for nil") }
}

func Test_C18_Dynamic_StructString(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	s := d.String()
	if s == "" { t.Fatal("expected non-empty") }
	_ = d.StructString()
	_ = d.StructStringPtr()
	// call again for cache
	_ = d.StructStringPtr()
}

func Test_C18_Dynamic_IsNull_IsValid(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	nilD := coredynamic.NewDynamic(nil, false)
	actual := args.Map{
		"isNull":    d.IsNull(),
		"isValid":   d.IsValid(),
		"isInvalid": d.IsInvalid(),
		"nilNull":   nilD.IsNull(),
	}
	expected := args.Map{
		"isNull": false, "isValid": true, "isInvalid": false, "nilNull": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- IsNull/IsValid", actual)
}

func Test_C18_Dynamic_IsPointer_IsValueType(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	if d.IsPointer() { t.Fatal("string not a pointer") }
	if !d.IsValueType() { t.Fatal("string is value type") }
	// call again for cache
	_ = d.IsPointer()

	ptrD := coredynamic.NewDynamic(&struct{}{}, true)
	if !ptrD.IsPointer() { t.Fatal("ptr should be pointer") }
}

func Test_C18_Dynamic_IsStructStringChecks(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	nilD := coredynamic.NewDynamic(nil, false)
	actual := args.Map{
		"nullOrEmpty":      d.IsStructStringNullOrEmpty(),
		"nullOrEmptyOrWs":  d.IsStructStringNullOrEmptyOrWhitespace(),
		"nilNullOrEmpty":   nilD.IsStructStringNullOrEmpty(),
	}
	expected := args.Map{
		"nullOrEmpty": false, "nullOrEmptyOrWs": false, "nilNullOrEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsStructString*", actual)
}

func Test_C18_Dynamic_TypeChecks(t *testing.T) {
	strD := coredynamic.NewDynamic("hello", true)
	intD := coredynamic.NewDynamic(42, true)
	sliceD := coredynamic.NewDynamic([]int{1}, true)
	mapD := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	structD := coredynamic.NewDynamic(struct{}{}, true)
	funcD := coredynamic.NewDynamic(func() {}, true)

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
	expected := args.Map{
		"isPrimStr": true, "isNumInt": true, "isString": true,
		"isStruct": true, "isFunc": true, "isSlice": true,
		"isSliceMap": true, "isMap": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- type checks", actual)
}

func Test_C18_Dynamic_IntDefault(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	val, ok := d.IntDefault(0)
	if !ok || val != 42 { t.Fatal("expected 42") }

	badD := coredynamic.NewDynamic("notint", true)
	val2, ok2 := badD.IntDefault(99)
	if ok2 || val2 != 99 { t.Fatal("expected 99 default") }

	nilD := coredynamic.NewDynamic(nil, false)
	val3, ok3 := nilD.IntDefault(7)
	if ok3 || val3 != 7 { t.Fatal("expected 7") }
}

func Test_C18_Dynamic_Float64(t *testing.T) {
	d := coredynamic.NewDynamic(3.14, true)
	_, err := d.Float64()
	if err != nil { t.Fatalf("err: %v", err) }

	nilD := coredynamic.NewDynamic(nil, false)
	_, err2 := nilD.Float64()
	if err2 == nil { t.Fatal("expected error for nil") }

	badD := coredynamic.NewDynamic("notnum", true)
	_, err3 := badD.Float64()
	if err3 == nil { t.Fatal("expected error") }
}

func Test_C18_Dynamic_ValueCasts(t *testing.T) {
	intD := coredynamic.NewDynamic(42, true)
	uintD := coredynamic.NewDynamic(uint(10), true)
	stringsD := coredynamic.NewDynamic([]string{"a"}, true)
	boolD := coredynamic.NewDynamic(true, true)
	int64D := coredynamic.NewDynamic(int64(100), true)

	actual := args.Map{
		"valInt":     intD.ValueInt(),
		"valUInt":    uintD.ValueUInt(),
		"valStrings": len(stringsD.ValueStrings()),
		"valBool":    boolD.ValueBool(),
		"valInt64":   int64D.ValueInt64(),
	}
	expected := args.Map{
		"valInt": 42, "valUInt": uint(10), "valStrings": 1,
		"valBool": true, "valInt64": int64(100),
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueCasts", actual)

	// Wrong type casts
	badInt := coredynamic.NewDynamic("str", true)
	if badInt.ValueInt() != -1 { t.Fatal("expected -1") }
	if badInt.ValueUInt() != 0 { t.Fatal("expected 0") }
	if badInt.ValueStrings() != nil { t.Fatal("expected nil") }
	if badInt.ValueBool() { t.Fatal("expected false") }
	if badInt.ValueInt64() != -1 { t.Fatal("expected -1") }
}

func Test_C18_Dynamic_ValueNullErr(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	if d.ValueNullErr() != nil { t.Fatal("expected nil") }

	nilD := coredynamic.NewDynamic(nil, false)
	if nilD.ValueNullErr() == nil { t.Fatal("expected error") }

	var ptrD *coredynamic.Dynamic
	if ptrD.ValueNullErr() == nil { t.Fatal("expected error") }
}

func Test_C18_Dynamic_ValueString(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	if d.ValueString() != "hello" { t.Fatal("expected hello") }

	intD := coredynamic.NewDynamic(42, true)
	if intD.ValueString() == "" { t.Fatal("expected non-empty") }

	var nilD *coredynamic.Dynamic
	if nilD.ValueString() != "" { t.Fatal("expected empty for nil") }
}

func Test_C18_Dynamic_Bytes(t *testing.T) {
	d := coredynamic.NewDynamic([]byte{1, 2}, true)
	b, ok := d.Bytes()
	if !ok || len(b) != 2 { t.Fatal("expected bytes") }

	var nilD *coredynamic.Dynamic
	b2, ok2 := nilD.Bytes()
	if ok2 || b2 != nil { t.Fatal("expected nil") }
}

// ── DynamicReflect ──

func Test_C18_Dynamic_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	rv := d.ReflectValue()
	if rv == nil { t.Fatal("expected non-nil") }
	// cached
	rv2 := d.ReflectValue()
	if rv2 == nil { t.Fatal("expected cached") }
}

func Test_C18_Dynamic_ReflectKind_Type_TypeName(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	if d.ReflectKind() != reflect.String { t.Fatal("expected string") }
	rt := d.ReflectType()
	if rt == nil { t.Fatal("expected non-nil") }
	// cached
	_ = d.ReflectType()
	tn := d.ReflectTypeName()
	if tn == "" { t.Fatal("expected non-empty") }
}

func Test_C18_Dynamic_IsReflectTypeOf_IsReflectKind(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	if !d.IsReflectTypeOf(reflect.TypeOf("")) { t.Fatal("expected true") }
	if !d.IsReflectKind(reflect.String) { t.Fatal("expected true") }
}

func Test_C18_Dynamic_ItemUsingIndex_Key(t *testing.T) {
	sliceD := coredynamic.NewDynamic([]int{10, 20, 30}, true)
	item := sliceD.ItemUsingIndex(1)
	if item != 20 { t.Fatal("expected 20") }
	rv := sliceD.ItemReflectValueUsingIndex(0)
	if rv.Int() != 10 { t.Fatal("expected 10") }

	mapD := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	val := mapD.ItemUsingKey("a")
	if val != 1 { t.Fatal("expected 1") }
	_ = mapD.ItemReflectValueUsingKey("a")
}

func Test_C18_Dynamic_Loop(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return false
	})
	if count != 3 { t.Fatal("expected 3 iterations") }

	// with break
	count2 := 0
	d.Loop(func(i int, item any) bool {
		count2++
		return i == 0
	})
	if count2 != 1 { t.Fatal("expected 1 iteration") }

	// nil/invalid
	nilD := coredynamic.NewDynamic(nil, false)
	called := nilD.Loop(func(i int, item any) bool { return false })
	if called { t.Fatal("expected false for nil") }
}

func Test_C18_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamic(map[string]int{"a": 1, "b": 2}, true)
	count := 0
	d.LoopMap(func(i int, key, value any) bool {
		count++
		return false
	})
	if count != 2 { t.Fatal("expected 2") }

	// with break
	d.LoopMap(func(i int, key, value any) bool { return true })

	// nil
	nilD := coredynamic.NewDynamic(nil, false)
	if nilD.LoopMap(func(i int, k, v any) bool { return false }) { t.Fatal("expected false") }
}

func Test_C18_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})
	if result.Length() != 2 { t.Fatalf("expected 2, got %d", result.Length()) }

	// with break
	result2 := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1
	})
	if result2.Length() != 2 { t.Fatal("expected 2") }

	// nil
	nilD := coredynamic.NewDynamic(nil, false)
	r := nilD.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) { return true, false })
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func Test_C18_Dynamic_ReflectSetTo(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	var target string
	_ = d.ReflectSetTo(&target)

	var nilD *coredynamic.Dynamic
	err := nilD.ReflectSetTo(&target)
	if err == nil { t.Fatal("expected error for nil") }
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
	d := coredynamic.NewDynamic(nil, false)
	b, err := d.JsonBytesPtr()
	if err != nil { t.Fatal("expected no error for null") }
	if len(b) != 0 { t.Fatal("expected empty bytes") }
}

func Test_C18_Dynamic_Deserialize(t *testing.T) {
	type testStruct struct{ Name string }
	target := &testStruct{}
	d := coredynamic.NewDynamic(target, true)
	_, _ = d.Deserialize([]byte(`{"Name":"test"}`))

	var nilD *coredynamic.Dynamic
	_, err := nilD.Deserialize([]byte(`{}`))
	if err == nil { t.Fatal("expected error") }
}

func Test_C18_Dynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	_ = d.UnmarshalJSON([]byte(`"world"`))

	var nilD *coredynamic.Dynamic
	err := nilD.UnmarshalJSON([]byte(`"x"`))
	if err == nil { t.Fatal("expected error for nil") }
}

func Test_C18_Dynamic_ValueMarshal_Nil(t *testing.T) {
	var nilD *coredynamic.Dynamic
	_, err := nilD.ValueMarshal()
	if err == nil { t.Fatal("expected error") }
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
	ds := coredynamic.InvalidDynamicStatus("err")
	dsNoMsg := coredynamic.InvalidDynamicStatusNoMessage()
	clone := ds.Clone()
	cloneP := ds.ClonePtr()
	if clone.IsValid() { t.Fatal("expected invalid") }
	if cloneP == nil { t.Fatal("expected non-nil") }
	if dsNoMsg.IsValid() { t.Fatal("expected invalid") }

	var nilDS *coredynamic.DynamicStatus
	if nilDS.ClonePtr() != nil { t.Fatal("expected nil") }
}

// ── SimpleRequest ──

func Test_C18_SimpleRequest(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("data", true, "msg")
	srValid := coredynamic.NewSimpleRequestValid("data")
	srInv := coredynamic.InvalidSimpleRequest("err")
	srInvNoMsg := coredynamic.InvalidSimpleRequestNoMessage()

	actual := args.Map{
		"msg":      sr.Message(),
		"req":      sr.Request(),
		"val":      sr.Value(),
		"srValid":  srValid.IsValid(),
		"srInv":    srInv.IsValid(),
		"srInvNo":  srInvNoMsg.IsValid(),
	}
	expected := args.Map{
		"msg": "msg", "req": "data", "val": "data",
		"srValid": true, "srInv": false, "srInvNo": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- with args", actual)
}

func Test_C18_SimpleRequest_TypeMismatch(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err != nil { t.Fatal("expected nil for matching type") }
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err2 == nil { t.Fatal("expected error for mismatch") }
	err3 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	if err3 == nil { t.Fatal("expected error with message") }
}

func Test_C18_SimpleRequest_IsPointer(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "")
	if sr.IsPointer() { t.Fatal("expected false") }
	_ = sr.IsReflectKind(reflect.String)
}

func Test_C18_SimpleRequest_InvalidError(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("data", true, "")
	if sr.InvalidError() != nil { t.Fatal("expected nil for empty msg") }
	sr2 := coredynamic.InvalidSimpleRequest("some error")
	err := sr2.InvalidError()
	if err == nil { t.Fatal("expected error") }
	// cached
	err2 := sr2.InvalidError()
	if err2 == nil { t.Fatal("expected cached error") }
}

// ── SimpleResult ──

func Test_C18_SimpleResult(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "")
	srValid := coredynamic.NewSimpleResultValid("data")
	srInv := coredynamic.InvalidSimpleResult("err")
	srInvNoMsg := coredynamic.InvalidSimpleResultNoMessage()

	actual := args.Map{
		"result":   sr.Result,
		"valid":    srValid.IsValid(),
		"inv":      srInv.IsValid(),
		"invNoMsg": srInvNoMsg.IsValid(),
	}
	expected := args.Map{
		"result": "data", "valid": true, "inv": false, "invNoMsg": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- with args", actual)
}

func Test_C18_SimpleResult_TypeMismatch(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err != nil { t.Fatal("expected nil") }
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C18_SimpleResult_InvalidError(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("data")
	if sr.InvalidError() != nil { t.Fatal("expected nil") }
	sr2 := coredynamic.InvalidSimpleResult("err msg")
	if sr2.InvalidError() == nil { t.Fatal("expected error") }
	// cached
	if sr2.InvalidError() == nil { t.Fatal("expected cached") }
}

func Test_C18_SimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("data")
	clone := sr.Clone()
	cloneP := sr.ClonePtr()
	if !clone.IsValid() { t.Fatal("clone invalid") }
	if cloneP == nil { t.Fatal("expected non-nil") }

	var nilSR *coredynamic.SimpleResult
	if nilSR.ClonePtr() != nil { t.Fatal("expected nil") }
}

// ── TypedDynamic ──

func Test_C18_TypedDynamic(t *testing.T) {
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	tdValid := coredynamic.NewTypedDynamicValid[string]("world")
	tdPtr := coredynamic.NewTypedDynamicPtr[int](42, true)
	invTd := coredynamic.InvalidTypedDynamic[string]()
	invTdP := coredynamic.InvalidTypedDynamicPtr[string]()

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
	td := coredynamic.NewTypedDynamic[string]("", false)
	err := td.UnmarshalJSON([]byte(`"hello"`))
	if err != nil { t.Fatal("expected nil") }
	if !td.IsValid() { t.Fatal("expected valid after unmarshal") }
}

func Test_C18_TypedDynamic_Deserialize(t *testing.T) {
	td := coredynamic.NewTypedDynamic[string]("", false)
	err := td.Deserialize([]byte(`"hello"`))
	if err != nil { t.Fatal("expected nil") }

	var nilTD *coredynamic.TypedDynamic[string]
	err2 := nilTD.Deserialize([]byte(`"x"`))
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C18_TypedDynamic_GetAs(t *testing.T) {
	strTD := coredynamic.NewTypedDynamic[string]("hello", true)
	intTD := coredynamic.NewTypedDynamic[int](42, true)
	boolTD := coredynamic.NewTypedDynamic[bool](true, true)

	s, sok := strTD.GetAsString()
	if !sok || s != "hello" { t.Fatal("GetAsString failed") }
	i, iok := intTD.GetAsInt()
	if !iok || i != 42 { t.Fatal("GetAsInt failed") }
	b, bok := boolTD.GetAsBool()
	if !bok || !b { t.Fatal("GetAsBool failed") }

	_, i64ok := intTD.GetAsInt64()
	if i64ok { t.Fatal("expected false for int->int64") }
	_, uok := intTD.GetAsUint()
	if uok { t.Fatal("expected false") }
	_, f64ok := intTD.GetAsFloat64()
	if f64ok { t.Fatal("expected false") }
	_, f32ok := intTD.GetAsFloat32()
	if f32ok { t.Fatal("expected false") }
	_, bok2 := intTD.GetAsBytes()
	if bok2 { t.Fatal("expected false") }
	_, sok2 := intTD.GetAsStrings()
	if sok2 { t.Fatal("expected false") }
}

func Test_C18_TypedDynamic_Value_Methods(t *testing.T) {
	strTD := coredynamic.NewTypedDynamic[string]("hello", true)
	intTD := coredynamic.NewTypedDynamic[int](42, true)
	boolTD := coredynamic.NewTypedDynamic[bool](true, true)
	int64TD := coredynamic.NewTypedDynamic[int64](int64(100), true)

	if strTD.ValueString() != "hello" { t.Fatal("expected hello") }
	if intTD.ValueInt() != 42 { t.Fatal("expected 42") }
	if !boolTD.ValueBool() { t.Fatal("expected true") }
	if int64TD.ValueInt64() != 100 { t.Fatal("expected 100") }

	// wrong type
	if strTD.ValueInt() != -1 { t.Fatal("expected -1") }
	if strTD.ValueBool() { t.Fatal("expected false") }
	if strTD.ValueInt64() != -1 { t.Fatal("expected -1") }
	if intTD.ValueString() == "" { t.Fatal("expected non-empty via sprintf") }
}

func Test_C18_TypedDynamic_Clone(t *testing.T) {
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	clone := td.Clone()
	cloneP := td.ClonePtr()
	_ = td.NonPtr()
	_ = td.Ptr()
	_ = td.ToDynamic()
	if clone.Data() != "hello" { t.Fatal("clone mismatch") }
	if cloneP == nil { t.Fatal("expected non-nil") }

	var nilTD *coredynamic.TypedDynamic[string]
	if nilTD.ClonePtr() != nil { t.Fatal("expected nil") }
}

func Test_C18_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	bytesTD := coredynamic.NewTypedDynamic[[]byte]([]byte{1, 2}, true)
	b, ok := bytesTD.Bytes()
	if !ok || len(b) != 2 { t.Fatal("expected bytes") }

	_, bok := bytesTD.GetAsBytes()
	if !bok { t.Fatal("expected true") }
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
	kvInt := coredynamic.KeyVal{Key: "k", Value: 42}
	kvBool := coredynamic.KeyVal{Key: "k", Value: true}
	kvStr := coredynamic.KeyVal{Key: "k", Value: []string{"a"}}
	kvUInt := coredynamic.KeyVal{Key: "k", Value: uint(5)}
	kvI64 := coredynamic.KeyVal{Key: "k", Value: int64(99)}

	if kvInt.ValueInt() != 42 { t.Fatal("expected 42") }
	if !kvBool.ValueBool() { t.Fatal("expected true") }
	if len(kvStr.ValueStrings()) != 1 { t.Fatal("expected 1") }
	if kvUInt.ValueUInt() != 5 { t.Fatal("expected 5") }
	if kvI64.ValueInt64() != 99 { t.Fatal("expected 99") }

	// wrong type
	kvBad := coredynamic.KeyVal{Key: "k", Value: "str"}
	if kvBad.ValueInt() != -1 { t.Fatal("expected -1") }
	if kvBad.ValueUInt() != 0 { t.Fatal("expected 0") }
	if kvBad.ValueBool() { t.Fatal("expected false") }
}

func Test_C18_KeyVal_NullChecks(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	if kv.ValueNullErr() != nil { t.Fatal("expected nil") }
	if kv.KeyNullErr() != nil { t.Fatal("expected nil") }

	kvNil := coredynamic.KeyVal{Key: nil, Value: nil}
	if kvNil.ValueNullErr() == nil { t.Fatal("expected error") }
	if kvNil.KeyNullErr() == nil { t.Fatal("expected error") }

	var nilKV *coredynamic.KeyVal
	if nilKV.ValueNullErr() == nil { t.Fatal("expected error for nil") }
	if nilKV.KeyNullErr() == nil { t.Fatal("expected error for nil") }
	if nilKV.KeyString() != "" { t.Fatal("expected empty") }
	if nilKV.ValueString() != "" { t.Fatal("expected empty") }
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
	kvc := coredynamic.NewKeyValCollection(5)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.AddPtr(&coredynamic.KeyVal{Key: "c", Value: 3})
	kvc.AddPtr(nil)
	kvc.AddMany(coredynamic.KeyVal{Key: "d", Value: 4})
	kvc.AddMany()
	kvc.AddManyPtr(&coredynamic.KeyVal{Key: "e", Value: 5}, nil)
	kvc.AddManyPtr()

	if kvc.Length() != 5 { t.Fatalf("expected 5, got %d", kvc.Length()) }
	if kvc.IsEmpty() { t.Fatal("expected not empty") }
	if !kvc.HasAnyItem() { t.Fatal("expected true") }
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
	var nilKVC *coredynamic.KeyValCollection
	if nilKVC.Length() != 0 { t.Fatal("expected 0") }
	if nilKVC.Items() != nil { t.Fatal("expected nil") }
	if nilKVC.String() != "" { t.Fatal("expected empty") }
	if nilKVC.ClonePtr() != nil { t.Fatal("expected nil") }
}

func Test_C18_KeyValCollection_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	if !kvc.IsEmpty() { t.Fatal("expected empty") }
	_ = kvc.MapAnyItems()
	_, _ = kvc.JsonMapResults()
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
	_ = kvc.AllKeys()
	_ = kvc.AllKeysSorted()
	_ = kvc.AllValues()
}

func Test_C18_KeyValCollection_Clone(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	clone := kvc.Clone()
	cloneP := kvc.ClonePtr()
	_ = clone.NonPtr()
	_ = kvc.Ptr()
	if cloneP.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C18_KeyValCollection_Paging(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)
	if len(pages) == 0 { t.Fatal("expected pages") }
	_ = kvc.GetPagingInfo(3, 1)
	_ = kvc.GetSinglePageCollection(3, 1)
	// small collection
	small := coredynamic.NewKeyValCollection(1)
	small.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	_ = small.GetSinglePageCollection(5, 1)
}

// ── LeftRight (coredynamic) ──

func Test_C18_LeftRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "L", Right: "R"}
	actual := args.Map{
		"empty":     lr.IsEmpty(),
		"hasAny":    lr.HasAnyItem(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"leftEmpty": lr.IsLeftEmpty(),
		"rightEmpty": lr.IsRightEmpty(),
	}
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
	var nilLR *coredynamic.LeftRight
	if !nilLR.IsEmpty() { t.Fatal("expected empty") }
	if nilLR.HasAnyItem() { t.Fatal("expected false") }
	if nilLR.HasLeft() { t.Fatal("expected false") }
	if nilLR.HasRight() { t.Fatal("expected false") }
	if nilLR.LeftToDynamic() != nil { t.Fatal("expected nil") }
	if nilLR.RightToDynamic() != nil { t.Fatal("expected nil") }
	if nilLR.DeserializeLeft() != nil { t.Fatal("expected nil") }
	if nilLR.DeserializeRight() != nil { t.Fatal("expected nil") }
	_ = nilLR.TypeStatus()
	if nilLR.LeftReflectSet(nil) != nil { t.Fatal("expected nil") }
	if nilLR.RightReflectSet(nil) != nil { t.Fatal("expected nil") }
}

// ── ValueStatus ──

func Test_C18_ValueStatus(t *testing.T) {
	vs := coredynamic.InvalidValueStatus("err")
	vsNoMsg := coredynamic.InvalidValueStatusNoMessage()
	if vs.IsValid { t.Fatal("expected invalid") }
	if vsNoMsg.IsValid { t.Fatal("expected invalid") }
}

// ── SafeTypeName ──

func Test_C18_SafeTypeName(t *testing.T) {
	name := coredynamic.SafeTypeName("hello")
	if name != "string" { t.Fatal("expected 'string'") }
	nilName := coredynamic.SafeTypeName(nil)
	if nilName != "" { t.Fatal("expected empty for nil") }
}

// ── IsAnyTypesOf ──

func Test_C18_IsAnyTypesOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	found := coredynamic.IsAnyTypesOf(strType, strType, intType)
	if !found { t.Fatal("expected true") }
	notFound := coredynamic.IsAnyTypesOf(reflect.TypeOf(true), strType, intType)
	if notFound { t.Fatal("expected false") }
}

// ── LengthOfReflect ──

func Test_C18_LengthOfReflect(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	l := coredynamic.LengthOfReflect(rv)
	if l != 3 { t.Fatalf("expected 3, got %d", l) }
}
