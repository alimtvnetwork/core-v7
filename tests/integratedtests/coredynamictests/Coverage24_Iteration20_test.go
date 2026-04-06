package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// TypeStatus — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I20_TypeStatus_IsValid_WithTypes(t *testing.T) {
	ts := &coredynamic.TypeStatus{
		IsSame: true,
		Left:   reflect.TypeOf(""),
		Right:  reflect.TypeOf(""),
	}
	actual := args.Map{"valid": ts.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- IsValid", actual)
}

func Test_I20_TypeStatus_IsValid_Nil(t *testing.T) {
	var ts *coredynamic.TypeStatus
	actual := args.Map{"valid": ts.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsValid nil", actual)
}

func Test_I20_TypeStatus_IsInvalid_Nil(t *testing.T) {
	var ts *coredynamic.TypeStatus
	actual := args.Map{"invalid": ts.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsInvalid nil", actual)
}

func Test_I20_TypeStatus_IsNotSame(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false}
	actual := args.Map{"notSame": ts.IsNotSame(), "notEqual": ts.IsNotEqualTypes()}
	expected := args.Map{"notSame": true, "notEqual": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsNotSame", actual)
}

func Test_I20_TypeStatus_IsAnyPointer(t *testing.T) {
	ts := coredynamic.TypeStatus{IsLeftPointer: true}
	actual := args.Map{"any": ts.IsAnyPointer(), "both": ts.IsBothPointer()}
	expected := args.Map{"any": true, "both": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsAnyPointer", actual)
}

func Test_I20_TypeStatus_IsBothPointer(t *testing.T) {
	ts := coredynamic.TypeStatus{IsLeftPointer: true, IsRightPointer: true}
	actual := args.Map{"both": ts.IsBothPointer()}
	expected := args.Map{"both": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsBothPointer", actual)
}

func Test_I20_TypeStatus_NonPointerLeft(t *testing.T) {
	ts := coredynamic.TypeStatus{
		Left:          reflect.TypeOf((*int)(nil)),
		Right:         reflect.TypeOf(0),
		IsLeftPointer: true,
	}
	np := ts.NonPointerLeft()
	actual := args.Map{"name": np.String()}
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NonPointerLeft", actual)
}

func Test_I20_TypeStatus_NonPointerRight(t *testing.T) {
	ts := coredynamic.TypeStatus{
		Left:           reflect.TypeOf(0),
		Right:          reflect.TypeOf((*int)(nil)),
		IsRightPointer: true,
	}
	np := ts.NonPointerRight()
	actual := args.Map{"name": np.String()}
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NonPointerRight", actual)
}

func Test_I20_TypeStatus_NonPointerLeft_NonPointer(t *testing.T) {
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(0), Right: reflect.TypeOf(0)}
	np := ts.NonPointerLeft()
	actual := args.Map{"name": np.String()}
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- NonPointerLeft non-ptr", actual)
}

func Test_I20_TypeStatus_IsSameRegardlessPointer_Same(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(0), Right: reflect.TypeOf(0)}
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsSameRegardlessPointer same", actual)
}

func Test_I20_TypeStatus_IsSameRegardlessPointer_PtrVsNonPtr(t *testing.T) {
	ts := coredynamic.TypeStatus{
		IsSame:        false,
		Left:          reflect.TypeOf((*int)(nil)),
		Right:         reflect.TypeOf(0),
		IsLeftPointer: true,
	}
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- IsSameRegardlessPointer ptr vs non-ptr", actual)
}

func Test_I20_TypeStatus_LeftName(t *testing.T) {
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"left": ts.LeftName(), "right": ts.RightName()}
	expected := args.Map{"left": "string", "right": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- LeftName/RightName", actual)
}

func Test_I20_TypeStatus_LeftFullName(t *testing.T) {
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"left": ts.LeftFullName(), "right": ts.RightFullName()}
	expected := args.Map{"left": "string", "right": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- LeftFullName/RightFullName", actual)
}

func Test_I20_TypeStatus_NotMatchMessage_Same(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	actual := args.Map{"empty": ts.NotMatchMessage("l", "r") == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchMessage same", actual)
}

func Test_I20_TypeStatus_NotMatchMessage_NotSame(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	msg := ts.NotMatchMessage("left", "right")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchMessage not same", actual)
}

func Test_I20_TypeStatus_NotMatchErr_Same(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true}
	actual := args.Map{"nil": ts.NotMatchErr("l", "r") == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotMatchErr same", actual)
}

func Test_I20_TypeStatus_NotMatchErr_NotSame(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"hasErr": ts.NotMatchErr("l", "r") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotMatchErr not same", actual)
}

func Test_I20_TypeStatus_ValidationError_Same(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true}
	actual := args.Map{"nil": ts.ValidationError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- ValidationError same", actual)
}

func Test_I20_TypeStatus_ValidationError_NotSame(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"hasErr": ts.ValidationError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- ValidationError not same", actual)
}

func Test_I20_TypeStatus_MustBeSame_NoPanic(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true}
	ts.MustBeSame() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame no panic", actual)
}

func Test_I20_TypeStatus_MustBeSame_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame panic", actual)
	}()
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	ts.MustBeSame()
}

func Test_I20_TypeStatus_SrcDestinationMustBeSame_NoPanic(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: true}
	ts.SrcDestinationMustBeSame()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame no panic", actual)
}

func Test_I20_TypeStatus_SrcDestinationMustBeSame_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame panic", actual)
	}()
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	ts.SrcDestinationMustBeSame()
}

func Test_I20_TypeStatus_NotEqualSrcDestinationMessage(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	msg := ts.NotEqualSrcDestinationMessage()
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotEqualSrcDestinationMessage", actual)
}

func Test_I20_TypeStatus_NotEqualSrcDestinationErr(t *testing.T) {
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"hasErr": ts.NotEqualSrcDestinationErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotEqualSrcDestinationErr", actual)
}

func Test_I20_TypeStatus_IsEqual_BothNil(t *testing.T) {
	var a, b *coredynamic.TypeStatus
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsEqual both nil", actual)
}

func Test_I20_TypeStatus_IsEqual_OneNil(t *testing.T) {
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	actual := args.Map{"eq": a.IsEqual(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsEqual one nil", actual)
}

func Test_I20_TypeStatus_IsEqual_Same(t *testing.T) {
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual same", actual)
}

func Test_I20_TypeStatus_IsEqual_DiffIsSame(t *testing.T) {
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff IsSame", actual)
}

func Test_I20_TypeStatus_IsEqual_DiffLeft(t *testing.T) {
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(0), Right: reflect.TypeOf("")}
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff Left", actual)
}

func Test_I20_TypeStatus_IsEqual_DiffRight(t *testing.T) {
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff Right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedDynamic — GetAs*, Value*, Clone, Deserialize, conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_I20_TypedDynamic_GetAsString(t *testing.T) {
	td := coredynamic.NewTypedDynamic("hello", true)
	val, ok := td.GetAsString()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsString", actual)
}

func Test_I20_TypedDynamic_GetAsInt(t *testing.T) {
	td := coredynamic.NewTypedDynamic(42, true)
	val, ok := td.GetAsInt()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt", actual)
}

func Test_I20_TypedDynamic_GetAsInt64(t *testing.T) {
	td := coredynamic.NewTypedDynamic(int64(99), true)
	val, ok := td.GetAsInt64()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": int64(99), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt64", actual)
}

func Test_I20_TypedDynamic_GetAsUint(t *testing.T) {
	td := coredynamic.NewTypedDynamic(uint(7), true)
	val, ok := td.GetAsUint()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": uint(7), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsUint", actual)
}

func Test_I20_TypedDynamic_GetAsFloat64(t *testing.T) {
	td := coredynamic.NewTypedDynamic(3.14, true)
	val, ok := td.GetAsFloat64()
	actual := args.Map{"ok": ok, "close": val > 3.0}
	expected := args.Map{"ok": true, "close": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat64", actual)
}

func Test_I20_TypedDynamic_GetAsFloat32(t *testing.T) {
	td := coredynamic.NewTypedDynamic(float32(1.5), true)
	val, ok := td.GetAsFloat32()
	actual := args.Map{"ok": ok, "close": val > 1.0}
	expected := args.Map{"ok": true, "close": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat32", actual)
}

func Test_I20_TypedDynamic_GetAsBool(t *testing.T) {
	td := coredynamic.NewTypedDynamic(true, true)
	val, ok := td.GetAsBool()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBool", actual)
}

func Test_I20_TypedDynamic_GetAsBytes(t *testing.T) {
	td := coredynamic.NewTypedDynamic([]byte{1, 2}, true)
	val, ok := td.GetAsBytes()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBytes", actual)
}

func Test_I20_TypedDynamic_GetAsStrings(t *testing.T) {
	td := coredynamic.NewTypedDynamic([]string{"a", "b"}, true)
	val, ok := td.GetAsStrings()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsStrings", actual)
}

func Test_I20_TypedDynamic_ValueString(t *testing.T) {
	td := coredynamic.NewTypedDynamic("hello", true)
	actual := args.Map{"val": td.ValueString()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns non-empty -- ValueString", actual)
}

func Test_I20_TypedDynamic_ValueString_NonString(t *testing.T) {
	td := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"notEmpty": td.ValueString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns non-empty -- ValueString non-string", actual)
}

func Test_I20_TypedDynamic_ValueInt(t *testing.T) {
	td := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"val": td.ValueInt()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt", actual)
}

func Test_I20_TypedDynamic_ValueInt64(t *testing.T) {
	td := coredynamic.NewTypedDynamic(int64(999), true)
	actual := args.Map{"val": td.ValueInt64()}
	expected := args.Map{"val": int64(999)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt64", actual)
}

func Test_I20_TypedDynamic_ValueBool(t *testing.T) {
	td := coredynamic.NewTypedDynamic(true, true)
	actual := args.Map{"val": td.ValueBool()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool", actual)
}

func Test_I20_TypedDynamic_ValueBool_NotBool(t *testing.T) {
	td := coredynamic.NewTypedDynamic("nope", true)
	actual := args.Map{"val": td.ValueBool()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool not bool", actual)
}

func Test_I20_TypedDynamic_Clone(t *testing.T) {
	td := coredynamic.NewTypedDynamic("hello", true)
	cloned := td.Clone()
	actual := args.Map{"valid": cloned.IsValid(), "val": cloned.Data()}
	expected := args.Map{"valid": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}
	td := coredynamic.NewTypedDynamic("hello", true)
	d := td.ToDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ToDynamic", actual)
}

func Test_I20_TypedDynamic_Deserialize(t *testing.T) {
	td := coredynamic.NewTypedDynamicPtr("", true)
	err := td.Deserialize([]byte(`"world"`))
	actual := args.Map{"noErr": err == nil, "val": td.Data()}
	expected := args.Map{"noErr": true, "val": "world"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

func Test_I20_TypedDynamic_Deserialize_Nil(t *testing.T) {
	var td *coredynamic.TypedDynamic[string]
	err := td.Deserialize([]byte(`"x"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- Deserialize nil", actual)
}

func Test_I20_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	td := coredynamic.NewTypedDynamic([]byte{1, 2}, true)
	b, ok := td.Bytes()
	actual := args.Map{"ok": ok, "len": len(b)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes as bytes", actual)
}

func Test_I20_TypedDynamic_Bytes_Marshal(t *testing.T) {
	td := coredynamic.NewTypedDynamic("hello", true)
	b, ok := td.Bytes()
	actual := args.Map{"ok": ok, "hasBytes": len(b) > 0}
	expected := args.Map{"ok": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes marshal", actual)
}

func Test_I20_TypedDynamic_NonPtr(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	np := td.NonPtr()
	actual := args.Map{"val": np.Data()}
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- NonPtr", actual)
}

func Test_I20_TypedDynamic_JsonModel(t *testing.T) {
	td := coredynamic.NewTypedDynamic("val", true)
	actual := args.Map{"val": td.JsonModel()}
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModel", actual)
}

func Test_I20_TypedDynamic_JsonModelAny(t *testing.T) {
	td := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"val": td.JsonModelAny()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModelAny", actual)
}

func Test_I20_TypedDynamic_ValueMarshal(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.ValueMarshal()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueMarshal", actual)
}

func Test_I20_TypedDynamic_MarshalJSON(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- MarshalJSON", actual)
}

func Test_I20_TypedDynamic_UnmarshalJSON(t *testing.T) {
	td := coredynamic.NewTypedDynamicPtr("", true)
	err := td.UnmarshalJSON([]byte(`"abc"`))
	actual := args.Map{"noErr": err == nil, "valid": td.IsValid()}
	expected := args.Map{"noErr": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleRequest — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I20_TypedSimpleRequest_Clone(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	cloned := sr.Clone()
	actual := args.Map{"valid": cloned.IsValid(), "val": cloned.Data()}
	expected := args.Map{"valid": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone", actual)
}

func Test_I20_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	simple := sr.ToSimpleRequest()
	actual := args.Map{"valid": simple.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToSimpleRequest", actual)
}

func Test_I20_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	td := sr.ToTypedDynamic()
	actual := args.Map{"valid": td.IsValid(), "val": td.Data()}
	expected := args.Map{"valid": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToTypedDynamic", actual)
}

func Test_I20_TypedSimpleRequest_ToDynamic(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	d := sr.ToDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToDynamic", actual)
}

func Test_I20_TypedSimpleRequest_GetAsString(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	val, ok := sr.GetAsString()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsString", actual)
}

func Test_I20_TypedSimpleRequest_GetAsInt(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid(42)
	val, ok := sr.GetAsInt()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsInt", actual)
}

func Test_I20_TypedSimpleRequest_InvalidError_Cached(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleRequest[string]("err msg")
	e1 := sr.InvalidError()
	e2 := sr.InvalidError()
	actual := args.Map{"same": e1 == e2, "hasErr": e1 != nil}
	expected := args.Map{"same": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidError cached", actual)
}

func Test_I20_TypedSimpleRequest_JsonBytes(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	b, err := sr.JsonBytes()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonBytes", actual)
}

func Test_I20_TypedSimpleRequest_JsonModel(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	actual := args.Map{"val": sr.JsonModel()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModel", actual)
}

func Test_I20_TypedSimpleRequest_JsonModelAny(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	actual := args.Map{"val": sr.JsonModelAny()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModelAny", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleResult — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I20_TypedSimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	cloned := sr.Clone()
	actual := args.Map{"valid": cloned.IsValid(), "val": cloned.Data()}
	expected := args.Map{"valid": true, "val": "ok"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone", actual)
}
func Test_I20_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	simple := sr.ToSimpleResult()
	actual := args.Map{"valid": simple.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToSimpleResult", actual)
}

func Test_I20_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	td := sr.ToTypedDynamic()
	actual := args.Map{"valid": td.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToTypedDynamic", actual)
}

func Test_I20_TypedSimpleResult_ToDynamic(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	d := sr.ToDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToDynamic", actual)
}

func Test_I20_TypedSimpleResult_InvalidError_Cached(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleResult[string]("err")
	e1 := sr.InvalidError()
	e2 := sr.InvalidError()
	actual := args.Map{"same": e1 == e2, "hasErr": e1 != nil}
	expected := args.Map{"same": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidError cached", actual)
}

func Test_I20_TypedSimpleResult_GetAsString(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	val, ok := sr.GetAsString()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "ok", "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsString", actual)
}

func Test_I20_TypedSimpleResult_GetAsFloat64(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid(3.14)
	val, ok := sr.GetAsFloat64()
	actual := args.Map{"ok": ok, "close": val > 3.0}
	expected := args.Map{"ok": true, "close": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsFloat64", actual)
}

func Test_I20_TypedSimpleResult_GetAsBool(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid(true)
	val, ok := sr.GetAsBool()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsBool", actual)
}

func Test_I20_TypedSimpleResult_JsonBytes(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	b, err := sr.JsonBytes()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonBytes", actual)
}

func Test_I20_TypedSimpleResult_JsonPtr(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	jr := sr.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonPtr", actual)
}

func Test_I20_TypedSimpleResult_InvalidNoMessage(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleResultNoMessage[string]()
	actual := args.Map{"invalid": sr.IsInvalid(), "msg": sr.Message()}
	expected := args.Map{"invalid": true, "msg": ""}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidNoMessage", actual)
}

func Test_I20_TypedSimpleRequest_InvalidNoMessage(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()
	actual := args.Map{"invalid": sr.IsInvalid(), "msg": sr.Message()}
	expected := args.Map{"invalid": true, "msg": ""}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidNoMessage", actual)
}

func Test_I20_TypedSimpleRequest_String(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	actual := args.Map{"notEmpty": sr.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- String", actual)
}

func Test_I20_TypedSimpleResult_String(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	actual := args.Map{"notEmpty": sr.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- String", actual)
}

func Test_I20_TypedDynamic_Invalid(t *testing.T) {
	td := coredynamic.InvalidTypedDynamic[string]()
	actual := args.Map{"invalid": td.IsInvalid()}
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- Invalid", actual)
}

func Test_I20_TypedDynamic_InvalidPtr(t *testing.T) {
	td := coredynamic.InvalidTypedDynamicPtr[string]()
	actual := args.Map{"invalid": td.IsInvalid(), "notNil": td != nil}
	expected := args.Map{"invalid": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- InvalidPtr", actual)
}

func Test_I20_TypedDynamic_String(t *testing.T) {
	td := coredynamic.NewTypedDynamic("hello", true)
	actual := args.Map{"notEmpty": td.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- String", actual)
}

func Test_I20_TypedDynamic_JsonBytes(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.JsonBytes()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonBytes", actual)
}

func Test_I20_TypedDynamic_JsonResult(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.JsonResult()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonResult", actual)
}

func Test_I20_TypedDynamic_JsonString(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	s, err := td.JsonString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonString", actual)
}

func Test_I20_TypedDynamic_Json(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Json", actual)
}

func Test_I20_TypedDynamic_JsonPtr(t *testing.T) {
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonPtr", actual)
}

func Test_I20_TypedSimpleResult_MarshalJSON(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	b, err := sr.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- MarshalJSON", actual)
}

func Test_I20_TypedSimpleRequest_MarshalJSON(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("ok")
	b, err := sr.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- MarshalJSON", actual)
}

func Test_I20_TypedSimpleResult_GetAsBytes(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid([]byte{1, 2})
	val, ok := sr.GetAsBytes()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsBytes", actual)
}

func Test_I20_TypedSimpleResult_GetAsStrings(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid([]string{"a"})
	val, ok := sr.GetAsStrings()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsStrings", actual)
}

func Test_I20_TypedSimpleRequest_GetAsBytes(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid([]byte{3})
	val, ok := sr.GetAsBytes()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsBytes", actual)
}

func Test_I20_TypedSimpleRequest_GetAsStrings(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid([]string{"a", "b"})
	val, ok := sr.GetAsStrings()
	actual := args.Map{"ok": ok, "len": len(val)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsStrings", actual)
}

func Test_I20_TypedSimpleRequest_GetAsInt64(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid(int64(77))
	val, ok := sr.GetAsInt64()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": int64(77), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsInt64", actual)
}

func Test_I20_TypedSimpleRequest_GetAsFloat64(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid(1.5)
	val, ok := sr.GetAsFloat64()
	actual := args.Map{"ok": ok, "close": val > 1.0}
	expected := args.Map{"ok": true, "close": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsFloat64", actual)
}

func Test_I20_TypedSimpleRequest_GetAsFloat32(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid(float32(2.5))
	val, ok := sr.GetAsFloat32()
	actual := args.Map{"ok": ok, "close": val > 2.0}
	expected := args.Map{"ok": true, "close": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsFloat32", actual)
}

func Test_I20_TypedSimpleRequest_GetAsBool(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid(true)
	val, ok := sr.GetAsBool()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsBool", actual)
}

func Test_I20_TypedSimpleResult_GetAsInt(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid(42)
	val, ok := sr.GetAsInt()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsInt", actual)
}

func Test_I20_TypedSimpleResult_GetAsInt64(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid(int64(88))
	val, ok := sr.GetAsInt64()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": int64(88), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsInt64", actual)
}

func Test_I20_TypedSimpleResult_JsonModel(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	actual := args.Map{"val": sr.JsonModel()}
	expected := args.Map{"val": "ok"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModel", actual)
}

func Test_I20_TypedSimpleResult_JsonModelAny(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	actual := args.Map{"val": sr.JsonModelAny()}
	expected := args.Map{"val": "ok"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModelAny", actual)
}

func Test_I20_TypedSimpleRequest_JsonResult(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.JsonResult()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonResult", actual)
}

func Test_I20_TypedSimpleRequest_Json(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Json", actual)
}

func Test_I20_TypedSimpleRequest_JsonPtr(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonPtr", actual)
}

func Test_I20_TypedSimpleResult_JsonResult(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("x")
	jr := sr.JsonResult()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonResult", actual)
}

func Test_I20_TypedSimpleResult_Json(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid("x")
	jr := sr.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Json", actual)
}
