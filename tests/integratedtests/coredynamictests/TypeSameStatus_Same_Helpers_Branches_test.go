package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// TypeSameStatus
// =============================================================================

func Test_TypeSameStatus_Same_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus same type", actual)
}

func Test_TypeSameStatus_Different_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{
		"same": ts.IsSame,
		"notSame": ts.IsNotSame(),
	}

	// Assert
	expected := args.Map{
		"same": false,
		"notSame": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus different", actual)
}

func Test_TypeSameStatus_NilLeft_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, "b")

	// Act
	actual := args.Map{
		"leftNull": ts.IsLeftUnknownNull,
		"same": ts.IsSame,
	}

	// Assert
	expected := args.Map{
		"leftNull": true,
		"same": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus nil left", actual)
}

func Test_TypeSameStatus_Pointers_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "b")

	// Act
	actual := args.Map{
		"leftPtr": ts.IsLeftPointer,
		"rightPtr": ts.IsRightPointer,
		"anyPtr": ts.IsAnyPointer(),
	}

	// Assert
	expected := args.Map{
		"leftPtr": true,
		"rightPtr": false,
		"anyPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus pointers", actual)
}

func Test_TypeSameStatus_BothPointers(t *testing.T) {
	// Arrange
	a := "a"
	b := "b"
	ts := coredynamic.TypeSameStatus(&a, &b)

	// Act
	actual := args.Map{"both": ts.IsBothPointer()}

	// Assert
	expected := args.Map{"both": true}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus both pointers", actual)
}

// =============================================================================
// TypeStatus — methods
// =============================================================================

func Test_TypeStatus_IsValid_Nil(t *testing.T) {
	// Arrange
	var ts *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"valid": ts.IsValid(),
		"invalid": ts.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsValid nil", actual)
}

func Test_TypeStatus_IsValid_Valid(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"valid": ts.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsValid valid", actual)
}

func Test_TypeStatus_IsValid_Cached(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.IsValid() // first call caches

	// Act
	actual := args.Map{"valid": ts.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsValid cached", actual)
}

func Test_TypeStatus_IsNotEqualTypes(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{"r": ts.IsNotEqualTypes()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsNotEqualTypes", actual)
}

func Test_TypeStatus_NonPointerLeft(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"name": ts.NonPointerLeft().Name()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "TypeStatus NonPointerLeft", actual)
}

func Test_TypeStatus_NonPointerRight(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"name": ts.NonPointerRight().Name()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "TypeStatus NonPointerRight", actual)
}

func Test_TypeStatus_NonPointerLeft_Pointer(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "b")

	// Act
	actual := args.Map{"name": ts.NonPointerLeft().Name()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "TypeStatus NonPointerLeft pointer", actual)
}

func Test_TypeStatus_IsSameRegardlessPointer_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"r": ts.IsSameRegardlessPointer()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsSameRegardlessPointer same", actual)
}

func Test_TypeStatus_IsSameRegardlessPointer_PtrVsVal(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "b")

	// Act
	actual := args.Map{"r": ts.IsSameRegardlessPointer()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsSameRegardlessPointer ptr vs val", actual)
}

func Test_TypeStatus_LeftName_Nil(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, "b")

	// Act
	actual := args.Map{"r": ts.LeftName()}

	// Assert
	expected := args.Map{"r": "<nil>"}
	expected.ShouldBeEqual(t, 0, "TypeStatus LeftName nil", actual)
}

func Test_TypeStatus_RightName_Nil(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", nil)

	// Act
	actual := args.Map{"r": ts.RightName()}

	// Assert
	expected := args.Map{"r": "<nil>"}
	expected.ShouldBeEqual(t, 0, "TypeStatus RightName nil", actual)
}

func Test_TypeStatus_LeftFullName_Nil(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, "b")

	// Act
	actual := args.Map{"r": ts.LeftFullName()}

	// Assert
	expected := args.Map{"r": "<nil>"}
	expected.ShouldBeEqual(t, 0, "TypeStatus LeftFullName nil", actual)
}

func Test_TypeStatus_RightFullName_Nil(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", nil)

	// Act
	actual := args.Map{"r": ts.RightFullName()}

	// Assert
	expected := args.Map{"r": "<nil>"}
	expected.ShouldBeEqual(t, 0, "TypeStatus RightFullName nil", actual)
}

func Test_TypeStatus_LeftName_Valid(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"r": ts.LeftName()}

	// Assert
	expected := args.Map{"r": "string"}
	expected.ShouldBeEqual(t, 0, "TypeStatus LeftName valid", actual)
}

func Test_TypeStatus_NotMatchMessage_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"r": ts.NotMatchMessage("l", "r")}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotMatchMessage same", actual)
}

func Test_TypeStatus_NotMatchMessage_Diff(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)
	msg := ts.NotMatchMessage("left", "right")

	// Act
	actual := args.Map{"nonEmpty": msg != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotMatchMessage diff", actual)
}

func Test_TypeStatus_NotMatchErr_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"noErr": ts.NotMatchErr("l", "r") == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotMatchErr same", actual)
}

func Test_TypeStatus_NotMatchErr_Diff(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{"hasErr": ts.NotMatchErr("l", "r") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotMatchErr diff", actual)
}

func Test_TypeStatus_ValidationError_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"noErr": ts.ValidationError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus ValidationError same", actual)
}

func Test_TypeStatus_ValidationError_Diff(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{"hasErr": ts.ValidationError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus ValidationError diff", actual)
}

func Test_TypeStatus_MustBeSame_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.MustBeSame() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus MustBeSame same", actual)
}

func Test_TypeStatus_MustBeSame_Panics(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus MustBeSame panics", actual)
	}()
	ts.MustBeSame()
}

func Test_TypeStatus_SrcDestinationMustBeSame_Same(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.SrcDestinationMustBeSame() // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus SrcDestMustBeSame same", actual)
}

func Test_TypeStatus_SrcDestinationMustBeSame_Panics(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus SrcDestMustBeSame panics", actual)
	}()
	ts.SrcDestinationMustBeSame()
}

func Test_TypeStatus_NotEqualSrcDestinationMessage(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)
	msg := ts.NotEqualSrcDestinationMessage()

	// Act
	actual := args.Map{"nonEmpty": msg != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotEqualSrcDestMsg", actual)
}

func Test_TypeStatus_NotEqualSrcDestinationErr(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{"hasErr": ts.NotEqualSrcDestinationErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus NotEqualSrcDestErr", actual)
}

func Test_TypeStatus_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *coredynamic.TypeStatus

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual both nil", actual)
}

func Test_TypeStatus_IsEqual_OneNil(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"r": ts.IsEqual(nil)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual one nil", actual)
}

func Test_TypeStatus_IsEqual_Same(t *testing.T) {
	// Arrange
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("c", "d")

	// Act
	actual := args.Map{"r": ts1.IsEqual(&ts2)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual same", actual)
}

func Test_TypeStatus_IsEqual_DiffSame(t *testing.T) {
	// Arrange
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", 42)

	// Act
	actual := args.Map{"r": ts1.IsEqual(&ts2)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual diff IsSame", actual)
}

func Test_TypeStatus_IsEqual_DiffRight(t *testing.T) {
	// Arrange
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", 42)
	ts2.IsSame = true // force same flag

	// Act
	actual := args.Map{"r": ts1.IsEqual(&ts2)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual diff Right", actual)
}

// =============================================================================
// ValueStatus
// =============================================================================

func Test_ValueStatus_Invalid_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("bad")

	// Act
	actual := args.Map{
		"valid": vs.IsValid,
		"msg": vs.Message,
		"idx": vs.Index,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "bad",
		"idx": -1,
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus Invalid", actual)
}

func Test_ValueStatus_InvalidNoMessage_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatusNoMessage()

	// Act
	actual := args.Map{
		"valid": vs.IsValid,
		"msg": vs.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus InvalidNoMessage", actual)
}

// =============================================================================
// SafeTypeName
// =============================================================================

func Test_SafeTypeName_Nil_FromTypeSameStatusSameHe(t *testing.T) {
	// Act
	actual := args.Map{"r": coredynamic.SafeTypeName(nil)}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "SafeTypeName nil", actual)
}

func Test_SafeTypeName_String_FromTypeSameStatusSameHe(t *testing.T) {
	// Act
	actual := args.Map{"r": coredynamic.SafeTypeName("hello")}

	// Assert
	expected := args.Map{"r": "string"}
	expected.ShouldBeEqual(t, 0, "SafeTypeName string", actual)
}

// =============================================================================
// LengthOfReflect
// =============================================================================

func Test_LengthOfReflect_Slice_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	v := reflect.ValueOf([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": coredynamic.LengthOfReflect(v)}

	// Assert
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect slice", actual)
}

func Test_LengthOfReflect_Array_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	v := reflect.ValueOf([2]int{1, 2})

	// Act
	actual := args.Map{"r": coredynamic.LengthOfReflect(v)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect array", actual)
}

func Test_LengthOfReflect_Map_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	v := reflect.ValueOf(map[string]int{"a": 1})

	// Act
	actual := args.Map{"r": coredynamic.LengthOfReflect(v)}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect map", actual)
}

func Test_LengthOfReflect_Default(t *testing.T) {
	// Arrange
	v := reflect.ValueOf(42)

	// Act
	actual := args.Map{"r": coredynamic.LengthOfReflect(v)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect default", actual)
}

// =============================================================================
// IsAnyTypesOf
// =============================================================================

func Test_IsAnyTypesOf_Found_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{"r": coredynamic.IsAnyTypesOf(strType, intType, strType)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf found", actual)
}

func Test_IsAnyTypesOf_NotFound_FromTypeSameStatusSameHe(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	boolType := reflect.TypeOf(true)

	// Act
	actual := args.Map{"r": coredynamic.IsAnyTypesOf(strType, intType, boolType)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf not found", actual)
}
