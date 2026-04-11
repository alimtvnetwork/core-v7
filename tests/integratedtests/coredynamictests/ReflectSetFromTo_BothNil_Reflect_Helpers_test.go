package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// ReflectSetFromTo
// =============================================================================

func Test_ReflectSetFromTo_BothNil_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(nil, nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo both nil", actual)
}

func Test_ReflectSetFromTo_ToNil_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to nil", actual)
}

func Test_ReflectSetFromTo_ToNotPointer(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to not pointer", actual)
}

func Test_ReflectSetFromTo_SameNonPointerToPointer(t *testing.T) {
	// Arrange
	var dest string
	err := coredynamic.ReflectSetFromTo("hello", &dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"dest": dest,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"dest": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo non-ptr to ptr", actual)
}

func Test_ReflectSetFromTo_SamePointerTypes(t *testing.T) {
	// Arrange
	src := "hello"
	var dest string
	err := coredynamic.ReflectSetFromTo(&src, &dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"dest": dest,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"dest": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo same ptr types", actual)
}

func Test_ReflectSetFromTo_BytesToStruct_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	type Simple struct{ Name string }
	b := []byte(`{"Name":"test"}`)
	var dest Simple
	err := coredynamic.ReflectSetFromTo(b, &dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": dest.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo bytes to struct", actual)
}

func Test_ReflectSetFromTo_StructToBytes_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	type Simple struct{ Name string }
	src := Simple{Name: "test"}
	var dest []byte
	err := coredynamic.ReflectSetFromTo(src, &dest)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(dest) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo struct to bytes", actual)
}

func Test_ReflectSetFromTo_TypeMismatch_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	var dest int
	err := coredynamic.ReflectSetFromTo("hello", &dest)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo type mismatch", actual)
}

// =============================================================================
// ReflectTypeValidation
// =============================================================================

func Test_ReflectTypeValidation_NilNotExpected(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation nil not expected", actual)
}

func Test_ReflectTypeValidation_Match_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation match", actual)
}

func Test_ReflectTypeValidation_Mismatch_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation mismatch", actual)
}

// =============================================================================
// ReflectKindValidation
// =============================================================================

func Test_ReflectKindValidation_Match_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.String, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation match", actual)
}

func Test_ReflectKindValidation_Mismatch_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.Int, "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation mismatch", actual)
}

// =============================================================================
// ReflectInterfaceVal
// =============================================================================

func Test_ReflectInterfaceVal_Value_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	r := coredynamic.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"r": r}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal value", actual)
}

func Test_ReflectInterfaceVal_Pointer_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	s := "hello"
	r := coredynamic.ReflectInterfaceVal(&s)

	// Act
	actual := args.Map{"r": r}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal pointer", actual)
}

// =============================================================================
// PointerOrNonPointer
// =============================================================================

func Test_PointerOrNonPointer_NonPointerOutput(t *testing.T) {
	// Arrange
	s := "hello"
	out, _ := coredynamic.PointerOrNonPointer(false, &s)

	// Act
	actual := args.Map{"r": out}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer non-ptr output", actual)
}

func Test_PointerOrNonPointer_ValuePassthrough(t *testing.T) {
	// Arrange
	out, _ := coredynamic.PointerOrNonPointer(false, "hello")

	// Act
	actual := args.Map{"r": out}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer value passthrough", actual)
}

// =============================================================================
// AnyToReflectVal
// =============================================================================

func Test_AnyToReflectVal_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	rv := coredynamic.AnyToReflectVal("hello")

	// Act
	actual := args.Map{
		"valid": rv.IsValid(),
		"kind": rv.Kind().String(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"kind": "string",
	}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal", actual)
}

// =============================================================================
// CastTo
// =============================================================================

func Test_CastTo_Matching_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	r := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"valid": r.IsValid,
		"matching": r.IsMatchingAcceptedType,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"matching": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo matching", actual)
}

func Test_CastTo_NotMatching_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	r := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{
		"matching": r.IsMatchingAcceptedType,
		"hasErr": r.HasError(),
	}

	// Assert
	expected := args.Map{
		"matching": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo not matching", actual)
}

// =============================================================================
// NotAcceptedTypesErr / MustBeAcceptedTypes
// =============================================================================

func Test_NotAcceptedTypesErr_Accepted_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr accepted", actual)
}

func Test_NotAcceptedTypesErr_NotAccepted_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr not accepted", actual)
}

func Test_MustBeAcceptedTypes_Valid(t *testing.T) {
	// Arrange
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes valid", actual)
}

func Test_MustBeAcceptedTypes_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes panics", actual)
	}()
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(0))
}

// =============================================================================
// TypeNotEqualErr / TypeMustBeSame
// =============================================================================

func Test_TypeNotEqualErr_Same_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr same", actual)
}

func Test_TypeNotEqualErr_Different_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr different", actual)
}

func Test_TypeMustBeSame_Same_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	coredynamic.TypeMustBeSame("a", "b")

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame same", actual)
}

func Test_TypeMustBeSame_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics", actual)
	}()
	coredynamic.TypeMustBeSame("a", 42)
}

// =============================================================================
// TypesIndexOf
// =============================================================================

func Test_TypesIndexOf_Found_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{"r": coredynamic.TypesIndexOf(strType, intType, strType)}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf found", actual)
}

func Test_TypesIndexOf_NotFound_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{"r": coredynamic.TypesIndexOf(strType, intType)}

	// Assert
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf not found", actual)
}

// =============================================================================
// Type
// =============================================================================

func Test_Type(t *testing.T) {
	// Arrange
	rt := coredynamic.Type("hello")

	// Act
	actual := args.Map{"name": rt.Name()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "Type", actual)
}

// =============================================================================
// ZeroSet / ZeroSetAny / SafeZeroSet
// =============================================================================

func Test_ZeroSet_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.ZeroSet(reflect.ValueOf(&s))

	// Act
	actual := args.Map{"name": s.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSet", actual)
}

func Test_ZeroSetAny_Valid(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.ZeroSetAny(&s)

	// Act
	actual := args.Map{"name": s.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny valid", actual)
}

func Test_ZeroSetAny_Nil_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	coredynamic.ZeroSetAny(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny nil", actual)
}

func Test_SafeZeroSet_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.SafeZeroSet(reflect.ValueOf(&s))

	// Act
	actual := args.Map{"name": s.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet", actual)
}

// =============================================================================
// LengthOfReflect — pointer case
// =============================================================================

func Test_LengthOfReflect_Ptr_FromReflectSetFromToBoth(t *testing.T) {
	// Arrange
	s := []int{1, 2, 3}
	rv := reflect.ValueOf(&s)

	// Act
	actual := args.Map{"r": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect ptr", actual)
}
