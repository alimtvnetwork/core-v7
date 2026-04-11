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

func Test_TypeStatus_IsValid_WithTypes(t *testing.T) {
	// Arrange
	ts := &coredynamic.TypeStatus{
		IsSame: true,
		Left:   reflect.TypeOf(""),
		Right:  reflect.TypeOf(""),
	}

	// Act
	actual := args.Map{"valid": ts.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- IsValid", actual)
}

func Test_TypeStatus_IsValid_Nil_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	var ts *coredynamic.TypeStatus

	// Act
	actual := args.Map{"valid": ts.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsValid nil", actual)
}

func Test_TypeStatus_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var ts *coredynamic.TypeStatus

	// Act
	actual := args.Map{"invalid": ts.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsInvalid nil", actual)
}

func Test_TypeStatus_IsNotSame(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false}

	// Act
	actual := args.Map{
		"notSame": ts.IsNotSame(),
		"notEqual": ts.IsNotEqualTypes(),
	}

	// Assert
	expected := args.Map{
		"notSame": true,
		"notEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsNotSame", actual)
}

func Test_TypeStatus_IsAnyPointer(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsLeftPointer: true}

	// Act
	actual := args.Map{
		"any": ts.IsAnyPointer(),
		"both": ts.IsBothPointer(),
	}

	// Assert
	expected := args.Map{
		"any": true,
		"both": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsAnyPointer", actual)
}

func Test_TypeStatus_IsBothPointer(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsLeftPointer: true, IsRightPointer: true}

	// Act
	actual := args.Map{"both": ts.IsBothPointer()}

	// Assert
	expected := args.Map{"both": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsBothPointer", actual)
}

func Test_TypeStatus_NonPointerLeft_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{
		Left:          reflect.TypeOf((*int)(nil)),
		Right:         reflect.TypeOf(0),
		IsLeftPointer: true,
	}
	np := ts.NonPointerLeft()

	// Act
	actual := args.Map{"name": np.String()}

	// Assert
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NonPointerLeft", actual)
}

func Test_TypeStatus_NonPointerRight_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{
		Left:           reflect.TypeOf(0),
		Right:          reflect.TypeOf((*int)(nil)),
		IsRightPointer: true,
	}
	np := ts.NonPointerRight()

	// Act
	actual := args.Map{"name": np.String()}

	// Assert
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NonPointerRight", actual)
}

func Test_TypeStatus_NonPointerLeft_NonPointer(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(0), Right: reflect.TypeOf(0)}
	np := ts.NonPointerLeft()

	// Act
	actual := args.Map{"name": np.String()}

	// Assert
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- NonPointerLeft non-ptr", actual)
}

func Test_TypeStatus_IsSameRegardlessPointer_Same_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(0), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsSameRegardlessPointer same", actual)
}

func Test_TypeStatus_IsSameRegardlessPointer_PtrVsNonPtr(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{
		IsSame:        false,
		Left:          reflect.TypeOf((*int)(nil)),
		Right:         reflect.TypeOf(0),
		IsLeftPointer: true,
	}

	// Act
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- IsSameRegardlessPointer ptr vs non-ptr", actual)
}

func Test_TypeStatus_LeftName(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{
		"left": ts.LeftName(),
		"right": ts.RightName(),
	}

	// Assert
	expected := args.Map{
		"left": "string",
		"right": "int",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- LeftName/RightName", actual)
}

func Test_TypeStatus_LeftFullName(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{
		"left": ts.LeftFullName(),
		"right": ts.RightFullName(),
	}

	// Assert
	expected := args.Map{
		"left": "string",
		"right": "int",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- LeftFullName/RightFullName", actual)
}

func Test_TypeStatus_NotMatchMessage_Same_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}

	// Act
	actual := args.Map{"empty": ts.NotMatchMessage("l", "r") == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchMessage same", actual)
}

func Test_TypeStatus_NotMatchMessage_NotSame(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	msg := ts.NotMatchMessage("left", "right")

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchMessage not same", actual)
}

func Test_TypeStatus_NotMatchErr_Same_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true}

	// Act
	actual := args.Map{"nil": ts.NotMatchErr("l", "r") == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotMatchErr same", actual)
}

func Test_TypeStatus_NotMatchErr_NotSame(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{"hasErr": ts.NotMatchErr("l", "r") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotMatchErr not same", actual)
}

func Test_TypeStatus_ValidationError_Same_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true}

	// Act
	actual := args.Map{"nil": ts.ValidationError() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- ValidationError same", actual)
}

func Test_TypeStatus_ValidationError_NotSame(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{"hasErr": ts.ValidationError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- ValidationError not same", actual)
}

func Test_TypeStatus_MustBeSame_NoPanic(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true}
	ts.MustBeSame() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame no panic", actual)
}

func Test_TypeStatus_MustBeSame_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame panic", actual)
	}()
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	ts.MustBeSame()
}

func Test_TypeStatus_SrcDestinationMustBeSame_NoPanic(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: true}
	ts.SrcDestinationMustBeSame()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame no panic", actual)
}

func Test_TypeStatus_SrcDestinationMustBeSame_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame panic", actual)
	}()
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	ts.SrcDestinationMustBeSame()
}

func Test_TypeStatus_NotEqualSrcDestinationMessage_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}
	msg := ts.NotEqualSrcDestinationMessage()

	// Act
	actual := args.Map{"notEmpty": msg != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotEqualSrcDestinationMessage", actual)
}

func Test_TypeStatus_NotEqualSrcDestinationErr_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{"hasErr": ts.NotEqualSrcDestinationErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- NotEqualSrcDestinationErr", actual)
}

func Test_TypeStatus_IsEqual_BothNil_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	var a, b *coredynamic.TypeStatus

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsEqual both nil", actual)
}

func Test_TypeStatus_IsEqual_OneNil_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}

	// Act
	actual := args.Map{"eq": a.IsEqual(nil)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- IsEqual one nil", actual)
}

func Test_TypeStatus_IsEqual_Same_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual same", actual)
}

func Test_TypeStatus_IsEqual_DiffIsSame(t *testing.T) {
	// Arrange
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: false, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff IsSame", actual)
}

func Test_TypeStatus_IsEqual_DiffLeft(t *testing.T) {
	// Arrange
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(0), Right: reflect.TypeOf("")}

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff Left", actual)
}

func Test_TypeStatus_IsEqual_DiffRight_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	a := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf("")}
	b := &coredynamic.TypeStatus{IsSame: true, Left: reflect.TypeOf(""), Right: reflect.TypeOf(0)}

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual diff Right", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedDynamic — GetAs*, Value*, Clone, Deserialize, conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypedDynamic_GetAsString(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)
	val, ok := td.GetAsString()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsString", actual)
}

func Test_TypedDynamic_GetAsInt(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(42, true)
	val, ok := td.GetAsInt()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt", actual)
}

func Test_TypedDynamic_GetAsInt64(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(int64(99), true)
	val, ok := td.GetAsInt64()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": int64(99),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt64", actual)
}

func Test_TypedDynamic_GetAsUint(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(uint(7), true)
	val, ok := td.GetAsUint()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": uint(7),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsUint", actual)
}

func Test_TypedDynamic_GetAsFloat64(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(3.14, true)
	val, ok := td.GetAsFloat64()

	// Act
	actual := args.Map{
		"ok": ok,
		"close": val > 3.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat64", actual)
}

func Test_TypedDynamic_GetAsFloat32(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(float32(1.5), true)
	val, ok := td.GetAsFloat32()

	// Act
	actual := args.Map{
		"ok": ok,
		"close": val > 1.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat32", actual)
}

func Test_TypedDynamic_GetAsBool(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(true, true)
	val, ok := td.GetAsBool()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBool", actual)
}

func Test_TypedDynamic_GetAsBytes(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic([]byte{1, 2}, true)
	val, ok := td.GetAsBytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBytes", actual)
}

func Test_TypedDynamic_GetAsStrings(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic([]string{"a", "b"}, true)
	val, ok := td.GetAsStrings()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsStrings", actual)
}

func Test_TypedDynamic_ValueString(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)

	// Act
	actual := args.Map{"val": td.ValueString()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns non-empty -- ValueString", actual)
}

func Test_TypedDynamic_ValueString_NonString(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"notEmpty": td.ValueString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns non-empty -- ValueString non-string", actual)
}

func Test_TypedDynamic_ValueInt(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"val": td.ValueInt()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt", actual)
}

func Test_TypedDynamic_ValueInt64(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(int64(999), true)

	// Act
	actual := args.Map{"val": td.ValueInt64()}

	// Assert
	expected := args.Map{"val": int64(999)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt64", actual)
}

func Test_TypedDynamic_ValueBool(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(true, true)

	// Act
	actual := args.Map{"val": td.ValueBool()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool", actual)
}

func Test_TypedDynamic_ValueBool_NotBool(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("nope", true)

	// Act
	actual := args.Map{"val": td.ValueBool()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool not bool", actual)
}

func Test_TypedDynamic_Clone(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)
	cloned := td.Clone()

	// Act
	actual := args.Map{
		"valid": cloned.IsValid(),
		"val": cloned.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}

func Test_TypedDynamic_ClonePtr(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamicPtr("hello", true)
	cloned := td.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"val": cloned.Data(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ClonePtr", actual)
}

func Test_TypedDynamic_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var td *coredynamic.TypedDynamic[string]

	// Act
	actual := args.Map{"nil": td.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- ClonePtr nil", actual)
}

func Test_TypedDynamic_ToDynamic(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)
	d := td.ToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ToDynamic", actual)
}

func Test_TypedDynamic_Deserialize_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamicPtr("", true)
	err := td.Deserialize([]byte(`"world"`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": td.Data(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "world",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

func Test_TypedDynamic_Deserialize_Nil(t *testing.T) {
	// Arrange
	var td *coredynamic.TypedDynamic[string]
	err := td.Deserialize([]byte(`"x"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- Deserialize nil", actual)
}

func Test_TypedDynamic_Bytes_AsBytes(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic([]byte{1, 2}, true)
	b, ok := td.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(b),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes as bytes", actual)
}

func Test_TypedDynamic_Bytes_Marshal(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)
	b, ok := td.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes marshal", actual)
}

func Test_TypedDynamic_NonPtr(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	np := td.NonPtr()

	// Act
	actual := args.Map{"val": np.Data()}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- NonPtr", actual)
}

func Test_TypedDynamic_JsonModel(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("val", true)

	// Act
	actual := args.Map{"val": td.JsonModel()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModel", actual)
}

func Test_TypedDynamic_JsonModelAny(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"val": td.JsonModelAny()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModelAny", actual)
}

func Test_TypedDynamic_ValueMarshal(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.ValueMarshal()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueMarshal", actual)
}

func Test_TypedDynamic_MarshalJSON(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- MarshalJSON", actual)
}

func Test_TypedDynamic_UnmarshalJSON(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamicPtr("", true)
	err := td.UnmarshalJSON([]byte(`"abc"`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"valid": td.IsValid(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleRequest — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypedSimpleRequest_Clone(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	cloned := sr.Clone()

	// Act
	actual := args.Map{
		"valid": cloned.IsValid(),
		"val": cloned.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone", actual)
}

func Test_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	simple := sr.ToSimpleRequest()

	// Act
	actual := args.Map{"valid": simple.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToSimpleRequest", actual)
}

func Test_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	td := sr.ToTypedDynamic()

	// Act
	actual := args.Map{
		"valid": td.IsValid(),
		"val": td.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToTypedDynamic", actual)
}

func Test_TypedSimpleRequest_ToDynamic(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	d := sr.ToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToDynamic", actual)
}

func Test_TypedSimpleRequest_GetAsString(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	val, ok := sr.GetAsString()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsString", actual)
}

func Test_TypedSimpleRequest_GetAsInt(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid(42)
	val, ok := sr.GetAsInt()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsInt", actual)
}

func Test_TypedSimpleRequest_InvalidError_Cached(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidTypedSimpleRequest[string]("err msg")
	e1 := sr.InvalidError()
	e2 := sr.InvalidError()

	// Act
	actual := args.Map{
		"same": e1 == e2,
		"hasErr": e1 != nil,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidError cached", actual)
}

func Test_TypedSimpleRequest_JsonBytes(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")
	b, err := sr.JsonBytes()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonBytes", actual)
}

func Test_TypedSimpleRequest_JsonModel(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")

	// Act
	actual := args.Map{"val": sr.JsonModel()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModel", actual)
}

func Test_TypedSimpleRequest_JsonModelAny(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")

	// Act
	actual := args.Map{"val": sr.JsonModelAny()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModelAny", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleResult — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypedSimpleResult_Clone(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	cloned := sr.Clone()

	// Act
	actual := args.Map{
		"valid": cloned.IsValid(),
		"val": cloned.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "ok",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone", actual)
}

func Test_TypedSimpleResult_ClonePtr(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	cloned := sr.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"val": cloned.Data(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": "ok",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ClonePtr", actual)
}

func Test_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	simple := sr.ToSimpleResult()

	// Act
	actual := args.Map{"valid": simple.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToSimpleResult", actual)
}

func Test_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	td := sr.ToTypedDynamic()

	// Act
	actual := args.Map{"valid": td.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToTypedDynamic", actual)
}

func Test_TypedSimpleResult_ToDynamic(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	d := sr.ToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToDynamic", actual)
}

func Test_TypedSimpleResult_InvalidError_Cached(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidTypedSimpleResult[string]("err")
	e1 := sr.InvalidError()
	e2 := sr.InvalidError()

	// Act
	actual := args.Map{
		"same": e1 == e2,
		"hasErr": e1 != nil,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidError cached", actual)
}

func Test_TypedSimpleResult_GetAsString(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	val, ok := sr.GetAsString()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "ok",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsString", actual)
}

func Test_TypedSimpleResult_GetAsFloat64(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid(3.14)
	val, ok := sr.GetAsFloat64()

	// Act
	actual := args.Map{
		"ok": ok,
		"close": val > 3.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsFloat64", actual)
}

func Test_TypedSimpleResult_GetAsBool(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid(true)
	val, ok := sr.GetAsBool()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsBool", actual)
}

func Test_TypedSimpleResult_JsonBytes(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	b, err := sr.JsonBytes()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonBytes", actual)
}

func Test_TypedSimpleResult_JsonPtr(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	jr := sr.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonPtr", actual)
}

func Test_TypedSimpleResult_InvalidNoMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidTypedSimpleResultNoMessage[string]()

	// Act
	actual := args.Map{
		"invalid": sr.IsInvalid(),
		"msg": sr.Message(),
	}

	// Assert
	expected := args.Map{
		"invalid": true,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidNoMessage", actual)
}

func Test_TypedSimpleRequest_InvalidNoMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()

	// Act
	actual := args.Map{
		"invalid": sr.IsInvalid(),
		"msg": sr.Message(),
	}

	// Assert
	expected := args.Map{
		"invalid": true,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidNoMessage", actual)
}

func Test_TypedSimpleRequest_String(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("hello")

	// Act
	actual := args.Map{"notEmpty": sr.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- String", actual)
}

func Test_TypedSimpleResult_String(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")

	// Act
	actual := args.Map{"notEmpty": sr.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- String", actual)
}

func Test_TypedDynamic_Invalid(t *testing.T) {
	// Arrange
	td := coredynamic.InvalidTypedDynamic[string]()

	// Act
	actual := args.Map{"invalid": td.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- Invalid", actual)
}

func Test_TypedDynamic_InvalidPtr(t *testing.T) {
	// Arrange
	td := coredynamic.InvalidTypedDynamicPtr[string]()

	// Act
	actual := args.Map{
		"invalid": td.IsInvalid(),
		"notNil": td != nil,
	}

	// Assert
	expected := args.Map{
		"invalid": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- InvalidPtr", actual)
}

func Test_TypedDynamic_String(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("hello", true)

	// Act
	actual := args.Map{"notEmpty": td.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- String", actual)
}

func Test_TypedDynamic_JsonBytes(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	b, err := td.JsonBytes()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonBytes", actual)
}

func Test_TypedDynamic_JsonResult(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.JsonResult()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonResult", actual)
}

func Test_TypedDynamic_JsonString_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	s, err := td.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonString", actual)
}

func Test_TypedDynamic_Json_FromTypeStatusIsValidIte(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Json", actual)
}

func Test_TypedDynamic_JsonPtr(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic("x", true)
	jr := td.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonPtr", actual)
}

func Test_TypedSimpleResult_MarshalJSON(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")
	b, err := sr.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- MarshalJSON", actual)
}

func Test_TypedSimpleRequest_MarshalJSON(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("ok")
	b, err := sr.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- MarshalJSON", actual)
}

func Test_TypedSimpleResult_GetAsBytes(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid([]byte{1, 2})
	val, ok := sr.GetAsBytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsBytes", actual)
}

func Test_TypedSimpleResult_GetAsStrings(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid([]string{"a"})
	val, ok := sr.GetAsStrings()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsStrings", actual)
}

func Test_TypedSimpleRequest_GetAsBytes(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid([]byte{3})
	val, ok := sr.GetAsBytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsBytes", actual)
}

func Test_TypedSimpleRequest_GetAsStrings(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid([]string{"a", "b"})
	val, ok := sr.GetAsStrings()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsStrings", actual)
}

func Test_TypedSimpleRequest_GetAsInt64(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid(int64(77))
	val, ok := sr.GetAsInt64()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": int64(77),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsInt64", actual)
}

func Test_TypedSimpleRequest_GetAsFloat64(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid(1.5)
	val, ok := sr.GetAsFloat64()

	// Act
	actual := args.Map{
		"ok": ok,
		"close": val > 1.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsFloat64", actual)
}

func Test_TypedSimpleRequest_GetAsFloat32(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid(float32(2.5))
	val, ok := sr.GetAsFloat32()

	// Act
	actual := args.Map{
		"ok": ok,
		"close": val > 2.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsFloat32", actual)
}

func Test_TypedSimpleRequest_GetAsBool(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid(true)
	val, ok := sr.GetAsBool()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAsBool", actual)
}

func Test_TypedSimpleResult_GetAsInt(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid(42)
	val, ok := sr.GetAsInt()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsInt", actual)
}

func Test_TypedSimpleResult_GetAsInt64(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid(int64(88))
	val, ok := sr.GetAsInt64()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": int64(88),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAsInt64", actual)
}

func Test_TypedSimpleResult_JsonModel(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")

	// Act
	actual := args.Map{"val": sr.JsonModel()}

	// Assert
	expected := args.Map{"val": "ok"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModel", actual)
}

func Test_TypedSimpleResult_JsonModelAny(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("ok")

	// Act
	actual := args.Map{"val": sr.JsonModelAny()}

	// Assert
	expected := args.Map{"val": "ok"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModelAny", actual)
}

func Test_TypedSimpleRequest_JsonResult(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.JsonResult()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonResult", actual)
}

func Test_TypedSimpleRequest_Json(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Json", actual)
}

func Test_TypedSimpleRequest_JsonPtr(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleRequestValid("x")
	jr := sr.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonPtr", actual)
}

func Test_TypedSimpleResult_JsonResult(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("x")
	jr := sr.JsonResult()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonResult", actual)
}

func Test_TypedSimpleResult_Json(t *testing.T) {
	// Arrange
	sr := coredynamic.NewTypedSimpleResultValid("x")
	jr := sr.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Json", actual)
}
