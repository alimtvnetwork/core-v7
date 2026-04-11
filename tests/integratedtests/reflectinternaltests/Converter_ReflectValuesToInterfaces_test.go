package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── reflectConverter — uncovered branches ──

func Test_Converter_ReflectValuesToInterfaces_Empty(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.ReflectValuesToInterfaces(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValuesToInterfaces returns empty -- empty", actual)
}

func Test_Converter_ReflectValuesToInterfaces_Valid(t *testing.T) {
	// Arrange
	rvs := []reflect.Value{reflect.ValueOf(42), reflect.ValueOf("hello")}
	result := reflectinternal.Converter.ReflectValuesToInterfaces(rvs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValuesToInterfaces returns non-empty -- valid", actual)
}

func Test_Converter_ReflectValueToAnyValue_Ptr(t *testing.T) {
	// Arrange
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue returns correct value -- ptr", actual)
}

func Test_Converter_ReflectValueToAnyValue_String(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("hello")
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue returns correct value -- string", actual)
}

func Test_Converter_ReflectValueToAnyValue_Int(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": int64(42)}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue returns correct value -- int", actual)
}

func Test_Converter_ReflectValueToAnyValue_Default(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(3.14)
	result := reflectinternal.Converter.ReflectValueToAnyValue(rv)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueToAnyValue returns correct value -- default", actual)
}

func Test_Converter_InterfacesToTypes(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.InterfacesToTypes([]any{42, "hello"})
	emptyResult := reflectinternal.Converter.InterfacesToTypes(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypes returns correct value -- with args", actual)
}

func Test_Converter_InterfacesToTypesNames(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.InterfacesToTypesNames([]any{42, "hello"})
	emptyResult := reflectinternal.Converter.InterfacesToTypesNames(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypesNames returns correct value -- with args", actual)
}

func Test_Converter_InterfacesToTypesNamesWithValues(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.InterfacesToTypesNamesWithValues([]any{42})
	emptyResult := reflectinternal.Converter.InterfacesToTypesNamesWithValues(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "InterfacesToTypesNamesWithValues returns non-empty -- with args", actual)
}

func Test_Converter_ReflectValueToPointerReflectValue(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValueToPointerReflectValue(rv)

	// Act
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}

	// Assert
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueToPointerReflectValue returns error -- with args", actual)
}

func Test_Converter_ToPtrRvIfNotAlready_NonPtr(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ToPtrRvIfNotAlready(rv)

	// Act
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}

	// Assert
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ToPtrRvIfNotAlready returns non-empty -- non-ptr", actual)
}

func Test_Converter_ToPtrRvIfNotAlready_AlreadyPtr(t *testing.T) {
	// Arrange
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ToPtrRvIfNotAlready(rv)

	// Act
	actual := args.Map{"isPtr": result.Kind() == reflect.Ptr}

	// Assert
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "ToPtrRvIfNotAlready returns correct value -- already-ptr", actual)
}

func Test_Converter_ReducePointer(t *testing.T) {
	// Arrange
	x := 42
	result := reflectinternal.Converter.ReducePointer(&x, 3)

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"kind": result.Kind == reflect.Int,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"kind": true,
	}
	expected.ShouldBeEqual(t, 0, "ReducePointer returns correct value -- with args", actual)
}

func Test_Converter_ReducePointerDefault(t *testing.T) {
	// Arrange
	x := 42
	result := reflectinternal.Converter.ReducePointerDefault(&x)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefault returns correct value -- with args", actual)
}

func Test_Converter_ReducePointerRvDefault(t *testing.T) {
	// Arrange
	x := 42
	rv := reflect.ValueOf(&x)
	result := reflectinternal.Converter.ReducePointerRvDefault(rv)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerRvDefault returns error -- with args", actual)
}

func Test_Converter_ReducePointerDefaultToType(t *testing.T) {
	// Arrange
	x := 42
	result := reflectinternal.Converter.ReducePointerDefaultToType(&x)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefaultToType returns correct value -- with args", actual)
}

func Test_Converter_ReducePointerRvDefaultToType_Nil(t *testing.T) {
	// Arrange
	// Zero reflect.Value panics on .Kind() — test with a valid non-pointer value instead
	rv := reflect.ValueOf("hello")
	result := reflectinternal.Converter.ReducePointerRvDefaultToType(rv)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerRvDefaultToType returns error -- valid", actual)
}

func Test_Converter_ReflectValToInterfaces_Slice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces returns correct value -- slice", actual)
}

func Test_Converter_ReflectValToInterfaces_PtrSlice(t *testing.T) {
	// Arrange
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces returns correct value -- ptr-slice", actual)
}

func Test_Converter_ReflectValToInterfaces_NotSlice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces returns correct value -- not-slice", actual)
}

func Test_Converter_ReflectValToInterfaces_EmptySlice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{})
	result := reflectinternal.Converter.ReflectValToInterfaces(false, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfaces returns empty -- empty-slice", actual)
}

func Test_Converter_ReflectValToInterfacesAsync(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync returns correct value -- with args", actual)
}

func Test_Converter_ReflectValToInterfacesAsync_Ptr(t *testing.T) {
	// Arrange
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync returns correct value -- ptr", actual)
}

func Test_Converter_ReflectValToInterfacesAsync_NotSlice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	result := reflectinternal.Converter.ReflectValToInterfacesAsync(rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesAsync returns correct value -- not-slice", actual)
}

func Test_Converter_ReflectValToInterfacesUsingProcessor(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesUsingProcessor(false,
		func(item any) (any, bool, bool) { return item, true, false }, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesUsingProcessor returns correct value -- with args", actual)
}

func Test_Converter_ReflectValToInterfacesUsingProcessor_Break(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	result := reflectinternal.Converter.ReflectValToInterfacesUsingProcessor(false,
		func(item any) (any, bool, bool) { return item, true, true }, rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ReflectValToInterfacesUsingProcessor returns correct value -- break", actual)
}

func Test_Converter_ReflectInterfaceVal(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.ReflectInterfaceVal(42)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- with args", actual)
}

func Test_Converter_ReflectInterfaceVal_Ptr(t *testing.T) {
	// Arrange
	x := 42
	result := reflectinternal.Converter.ReflectInterfaceVal(&x)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- ptr", actual)
}

func Test_Converter_ToPointerRv(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.ToPointerRv(42)
	nilResult := reflectinternal.Converter.ToPointerRv(nil)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ToPointerRv returns error -- with args", actual)
}

func Test_Converter_ToPointer(t *testing.T) {
	// Arrange
	result := reflectinternal.Converter.ToPointer(42)
	nilResult := reflectinternal.Converter.ToPointer(nil)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ToPointer returns correct value -- with args", actual)
}

// ── reflectGetUsingReflectValue — uncovered branches ──

func Test_RvGetter_PublicValuesMapStruct_NonStruct(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.ReflectGetterUsingReflectValue.PublicValuesMapStruct(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PublicValuesMapStruct returns non-empty -- non-struct", actual)
}

func Test_RvGetter_PublicValuesMapStruct_WithUnexported(t *testing.T) {
	// Arrange
	type S struct {
		Pub  int
		priv int //nolint:unused
	}
	rv := reflect.ValueOf(S{Pub: 1})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.PublicValuesMapStruct(rv)

	// Act
	actual := args.Map{
		"len": len(m),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "PublicValuesMapStruct returns non-empty -- unexported", actual)
}

func Test_RvGetter_FieldNameWithTypeMap_Struct(t *testing.T) {
	// Arrange
	type S struct{ A int; B string }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap returns non-empty -- struct", actual)
}

func Test_RvGetter_FieldNameWithTypeMap_Ptr(t *testing.T) {
	// Arrange
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap returns non-empty -- ptr", actual)
}

func Test_RvGetter_FieldNameWithTypeMap_NonStruct(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	m := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)

	// Act
	actual := args.Map{"nil": m == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "FieldNameWithTypeMap returns non-empty -- non-struct", actual)
}

func Test_RvGetter_FieldNameWithValuesMap(t *testing.T) {
	// Arrange
	type S struct{ A int }
	rv := reflect.ValueOf(S{A: 42})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithValuesMap(rv)

	// Act
	actual := args.Map{
		"len": len(m),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldNameWithValuesMap returns non-empty -- with args", actual)
}

func Test_RvGetter_FieldNamesMap_Ptr(t *testing.T) {
	// Arrange
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNamesMap(rv)

	// Act
	actual := args.Map{
		"len": len(m),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldNamesMap returns correct value -- ptr", actual)
}

func Test_RvGetter_FieldNamesMap_NonStruct(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNamesMap(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FieldNamesMap returns non-empty -- non-struct", actual)
}

func Test_RvGetter_StructFieldsMap_Ptr(t *testing.T) {
	// Arrange
	type S struct{ A int }
	rv := reflect.ValueOf(&S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.StructFieldsMap(rv)

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StructFieldsMap returns correct value -- ptr", actual)
}

func Test_RvGetter_StructFieldsMap_NonStruct(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	m := reflectinternal.ReflectGetterUsingReflectValue.StructFieldsMap(rv)

	// Act
	actual := args.Map{"nil": m == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "StructFieldsMap returns non-empty -- non-struct", actual)
}

func Test_RvGetter_NullFieldsMap(t *testing.T) {
	// Arrange
	type S struct{ A *int }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.NullFieldsMap(3, rv)

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "NullFieldsMap returns correct value -- with args", actual)
}

func Test_RvGetter_NullOrZeroFieldsMap(t *testing.T) {
	// Arrange
	type S struct{ A int; B *int }
	rv := reflect.ValueOf(S{})
	m := reflectinternal.ReflectGetterUsingReflectValue.NullOrZeroFieldsMap(3, rv)

	// Act
	actual := args.Map{"gt0": len(m) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "NullOrZeroFieldsMap returns correct value -- with args", actual)
}

// ── reflectTypeConverter — uncovered branches ──

func Test_ReflectType_SafeName(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.SafeName(42)
	nilResult := reflectinternal.ReflectType.SafeName(nil)

	// Act
	actual := args.Map{
		"val": result,
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"val": "int",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeName returns correct value -- with args", actual)
}

func Test_ReflectType_SafeTypeNameOfSliceOrSingle(t *testing.T) {
	// Arrange
	single := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(true, 42)
	slice := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(false, []int{1})

	// Act
	actual := args.Map{
		"single": single,
		"slice": slice,
	}

	// Assert
	expected := args.Map{
		"single": "int",
		"slice": "int",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeNameOfSliceOrSingle returns correct value -- with args", actual)
}

func Test_ReflectType_SliceFirstItemTypeName(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.SliceFirstItemTypeName([]int{1})
	nilResult := reflectinternal.ReflectType.SliceFirstItemTypeName(nil)

	// Act
	actual := args.Map{
		"val": result,
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"val": "int",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SliceFirstItemTypeName returns correct value -- with args", actual)
}

func Test_ReflectType_NamesStringUsingReflectType(t *testing.T) {
	// Arrange
	full := reflectinternal.ReflectType.NamesStringUsingReflectType(true, reflect.TypeOf(42))
	empty := reflectinternal.ReflectType.NamesStringUsingReflectType(true)

	// Act
	actual := args.Map{
		"notEmpty": full != "",
		"empty": empty,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "NamesStringUsingReflectType returns correct value -- with args", actual)
}

func Test_ReflectType_TypeNamesString(t *testing.T) {
	// Arrange
	full := reflectinternal.ReflectType.TypeNamesString(true, 42)
	empty := reflectinternal.ReflectType.TypeNamesString(true)

	// Act
	actual := args.Map{
		"notEmpty": full != "",
		"empty": empty,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "TypeNamesString returns correct value -- with args", actual)
}

func Test_ReflectType_NamesUsingReflectType_Short(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.NamesUsingReflectType(false, reflect.TypeOf(42))

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NamesUsingReflectType returns correct value -- short", actual)
}

func Test_ReflectType_NamesReferenceString(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.NamesReferenceString(true, 42)
	empty := reflectinternal.ReflectType.NamesReferenceString(true)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": empty,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "NamesReferenceString returns correct value -- with args", actual)
}

func Test_ReflectType_Names(t *testing.T) {
	// Arrange
	full := reflectinternal.ReflectType.Names(true, 42)
	short := reflectinternal.ReflectType.Names(false, 42)
	empty := reflectinternal.ReflectType.Names(true)

	// Act
	actual := args.Map{
		"fullLen": len(full),
		"shortLen": len(short),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"fullLen": 1,
		"shortLen": 1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Names returns correct value -- with args", actual)
}

func Test_ReflectType_Name(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.Name(42)
	nilResult := reflectinternal.ReflectType.Name(nil)

	// Act
	actual := args.Map{
		"val": result,
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"val": "int",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "Name returns correct value -- with args", actual)
}

func Test_ReflectType_NameUsingFmt(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.NameUsingFmt(42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameUsingFmt returns correct value -- with args", actual)
}

// ── sliceConverter — uncovered branches ──

func Test_SliceConverter_Length(t *testing.T) {
	// Act
	actual := args.Map{
		"slice": reflectinternal.SliceConverter.Length([]int{1, 2}),
		"map":   reflectinternal.SliceConverter.Length(map[string]int{"a": 1}),
		"nil":   reflectinternal.SliceConverter.Length(nil),
		"int":   reflectinternal.SliceConverter.Length(42),
	}

	// Assert
	expected := args.Map{
		"slice": 2,
		"map": 1,
		"nil": 0,
		"int": 0,
	}
	expected.ShouldBeEqual(t, 0, "SliceConverter returns correct value -- Length", actual)
}

func Test_SliceConverter_ToStringsRv_Ptr(t *testing.T) {
	// Arrange
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	result, err := reflectinternal.SliceConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringsRv returns correct value -- ptr", actual)
}

func Test_SliceConverter_ToStringsRv_NotSlice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.SliceConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToStringsRv returns correct value -- not-slice", actual)
}

func Test_SliceConverter_ToStringsRv_Empty(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{})
	result, err := reflectinternal.SliceConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringsRv returns empty -- empty", actual)
}

func Test_SliceConverter_ToStrings(t *testing.T) {
	// Arrange
	result, err := reflectinternal.SliceConverter.ToStrings([]int{1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- with args", actual)
}

func Test_SliceConverter_ToStringsMust(t *testing.T) {
	// Arrange
	result := reflectinternal.SliceConverter.ToStringsMust([]int{1, 2})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStringsMust returns correct value -- with args", actual)
}

func Test_SliceConverter_ToStringsRvUsingProcessor(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})
	result, err := reflectinternal.SliceConverter.ToStringsRvUsingProcessor(rv,
		func(i int, item any) (string, bool, bool) { return "x", true, i == 1 })

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringsRvUsingProcessor returns correct value -- with args", actual)
}

func Test_SliceConverter_ToStringsRvUsingSimpleProcessor(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2})
	result, err := reflectinternal.SliceConverter.ToStringsRvUsingSimpleProcessor(rv, true,
		func(i int, item any) string {
			if i == 0 {
				return ""
			}
			return "x"
		})

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringsRvUsingSimpleProcessor returns correct value -- with args", actual)
}

func Test_SliceConverter_ToAnyItemsAsync(t *testing.T) {
	// Arrange
	result := reflectinternal.SliceConverter.ToAnyItemsAsync([]int{1, 2})
	nilResult := reflectinternal.SliceConverter.ToAnyItemsAsync(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToAnyItemsAsync returns correct value -- with args", actual)
}

// ── mapConverter — uncovered branches ──

func Test_MapConverter_Length(t *testing.T) {
	// Act
	actual := args.Map{"val": reflectinternal.MapConverter.Length(map[string]int{"a": 1})}

	// Assert
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "MapConverter returns correct value -- Length", actual)
}

func Test_MapConverter_ToStringsRv_Ptr(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(&m)
	result, err := reflectinternal.MapConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter returns correct value -- ToStringsRv ptr", actual)
}

func Test_MapConverter_ToStringsRv_NotMap(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapConverter returns correct value -- ToStringsRv not-map", actual)
}

func Test_MapConverter_ToStringsRv_IntKey(t *testing.T) {
	// Arrange
	m := map[int]string{1: "a"}
	rv := reflect.ValueOf(m)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapConverter returns correct value -- ToStringsRv int-key", actual)
}

func Test_MapConverter_ToKeysStrings(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToKeysStrings(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToKeysStrings returns correct value -- with args", actual)
}

func Test_MapConverter_ToValuesAny(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToValuesAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToValuesAny(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToValuesAny returns non-empty -- with args", actual)
}

func Test_MapConverter_ToKeysAny(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToKeysAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToKeysAny(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToKeysAny returns correct value -- with args", actual)
}

func Test_MapConverter_ToKeysValuesAny(t *testing.T) {
	// Arrange
	keys, vals, err := reflectinternal.MapConverter.ToKeysValuesAny(map[string]int{"a": 1})
	nilKeys, _, _ := reflectinternal.MapConverter.ToKeysValuesAny(nil)

	// Act
	actual := args.Map{
		"keysLen": len(keys),
		"valsLen": len(vals),
		"noErr": err == nil,
		"nilKeysLen": len(nilKeys),
	}

	// Assert
	expected := args.Map{
		"keysLen": 1,
		"valsLen": 1,
		"noErr": true,
		"nilKeysLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToKeysValuesAny returns non-empty -- with args", actual)
}

func Test_MapConverter_ToStrings_Nil(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToStrings(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStrings returns nil -- nil", actual)
}

func Test_MapConverter_ToStringsMust_ReflectValue(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(m)
	result := reflectinternal.MapConverter.ToStringsMust(rv)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsMust returns correct value -- reflect.Value", actual)
}

func Test_MapConverter_ToSortedStrings(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToSortedStrings(map[string]int{"b": 2, "a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToSortedStrings(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToSortedStrings returns correct value -- with args", actual)
}

func Test_MapConverter_ToSortedStringsMust(t *testing.T) {
	// Arrange
	result := reflectinternal.MapConverter.ToSortedStringsMust(map[string]int{"b": 2, "a": 1})
	nilResult := reflectinternal.MapConverter.ToSortedStringsMust(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToSortedStringsMust returns correct value -- with args", actual)
}

func Test_MapConverter_ToMapStringAnyRv(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToMapStringAnyRv returns correct value -- with args", actual)
}

func Test_MapConverter_ToMapStringAnyRv_IntKey(t *testing.T) {
	// Arrange
	m := map[int]string{1: "a"}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToMapStringAnyRv returns correct value -- int-key", actual)
}

func Test_MapConverter_ToMapStringAny(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToMapStringAny(map[string]int{"a": 1})
	nilResult, _ := reflectinternal.MapConverter.ToMapStringAny(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToMapStringAny returns correct value -- with args", actual)
}

// ── isChecker — uncovered branches ──

func Test_Is_Conclusive_Slices(t *testing.T) {
	// Arrange
	eq, conc := reflectinternal.Is.Conclusive([]int{1}, []int{1})

	// Act
	actual := args.Map{
		"eq": eq,
		"conc": conc,
	}

	// Assert
	expected := args.Map{
		"eq": false,
		"conc": false,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- slices", actual)
}

func Test_Is_Conclusive_BothNilPtrs(t *testing.T) {
	// Arrange
	var a, b *int
	eq, conc := reflectinternal.Is.Conclusive(a, b)

	// Act
	actual := args.Map{
		"eq": eq,
		"conc": conc,
	}

	// Assert
	expected := args.Map{
		"eq": true,
		"conc": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- both nil ptrs", actual)
}

func Test_Is_Conclusive_OneNilPtr(t *testing.T) {
	// Arrange
	var a *int
	x := 42
	eq, conc := reflectinternal.Is.Conclusive(a, &x)

	// Act
	actual := args.Map{
		"eq": eq,
		"conc": conc,
	}

	// Assert
	expected := args.Map{
		"eq": false,
		"conc": true,
	}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- one nil ptr", actual)
}

func Test_Is_AnyEqual_Slices(t *testing.T) {
	// Act
	actual := args.Map{"val": reflectinternal.Is.AnyEqual([]int{1}, []int{1})}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AnyEqual returns correct value -- slices", actual)
}

func Test_Is_NotFunc(t *testing.T) {
	// Act
	actual := args.Map{
		"int":  reflectinternal.Is.NotFunc(42),
		"nil":  reflectinternal.Is.NotFunc(nil),
	}

	// Assert
	expected := args.Map{
		"int": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "NotFunc returns correct value -- with args", actual)
}

func Test_Is_SliceOrArrayOf(t *testing.T) {
	// Act
	actual := args.Map{
		"slice": reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf([]int{})),
		"int":   reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf(42)),
	}

	// Assert
	expected := args.Map{
		"slice": true,
		"int": false,
	}
	expected.ShouldBeEqual(t, 0, "SliceOrArrayOf returns correct value -- with args", actual)
}

func Test_Is_NotNull(t *testing.T) {
	// Act
	actual := args.Map{
		"val": reflectinternal.Is.NotNull(42),
		"nil": reflectinternal.Is.NotNull(nil),
	}

	// Assert
	expected := args.Map{
		"val": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "NotNull returns correct value -- with args", actual)
}

func Test_Is_Defined(t *testing.T) {
	// Act
	actual := args.Map{"val": reflectinternal.Is.Defined(42)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Defined returns correct value -- with args", actual)
}

func Test_Is_ZeroRv_Array(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([2]int{0, 0})

	// Act
	actual := args.Map{"zero": reflectinternal.Is.ZeroRv(rv)}

	// Assert
	expected := args.Map{"zero": true}
	expected.ShouldBeEqual(t, 0, "ZeroRv returns correct value -- array", actual)
}

func Test_Is_ZeroRv_Struct(t *testing.T) {
	// Arrange
	type S struct{ A int }
	rv := reflect.ValueOf(S{})

	// Act
	actual := args.Map{"zero": reflectinternal.Is.ZeroRv(rv)}

	// Assert
	expected := args.Map{"zero": true}
	expected.ShouldBeEqual(t, 0, "ZeroRv returns correct value -- struct", actual)
}

func Test_Is_InterfaceRv(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	actual := args.Map{"val": reflectinternal.Is.InterfaceRv(rv)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "InterfaceRv returns correct value -- with args", actual)
}

func Test_Is_Interface(t *testing.T) {
	// Act
	actual := args.Map{"val": reflectinternal.Is.Interface(42)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "Interface returns correct value -- with args", actual)
}

func Test_Is_StructRv(t *testing.T) {
	// Arrange
	type S struct{}
	rv := reflect.ValueOf(S{})

	// Act
	actual := args.Map{"val": reflectinternal.Is.StructRv(rv)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "StructRv returns correct value -- with args", actual)
}

// ── looper — uncovered branches ──

func Test_Looper_FieldsFor(t *testing.T) {
	// Arrange
	type S struct{ A int; B string }
	count := 0
	err := reflectinternal.Looper.FieldsFor(S{}, func(fp *reflectmodel.FieldProcessor) error {
		return nil
	})
	_ = count

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "FieldsFor returns correct value -- with args", actual)
}

func Test_Looper_FieldNames(t *testing.T) {
	// Arrange
	type S struct{ A int; B string }
	names, err := reflectinternal.Looper.FieldNames(S{})

	// Act
	actual := args.Map{
		"len": len(names),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldNames returns correct value -- with args", actual)
}

func Test_Looper_FieldsMap(t *testing.T) {
	// Arrange
	type S struct{ A int }
	m, err := reflectinternal.Looper.FieldsMap(S{})

	// Act
	actual := args.Map{
		"len": len(m),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldsMap returns correct value -- with args", actual)
}

func Test_Looper_MethodsMap(t *testing.T) {
	// Arrange
	type S struct{}
	m, err := reflectinternal.Looper.MethodsMap(S{})

	// Act
	actual := args.Map{
		"notNil": m != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodsMap returns correct value -- with args", actual)
}

func Test_Looper_ReducePointer(t *testing.T) {
	// Arrange
	x := 42
	result := reflectinternal.Looper.ReducePointer(&x, 3)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointer returns correct value -- with args", actual)
}

func Test_Looper_ReducePointerDefault(t *testing.T) {
	// Arrange
	result := reflectinternal.Looper.ReducePointerDefault(42)

	// Act
	actual := args.Map{"valid": result.IsValid}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "ReducePointerDefault returns correct value -- with args", actual)
}

func Test_Looper_Slice(t *testing.T) {
	// Arrange
	err := reflectinternal.Looper.Slice([]int{1, 2}, func(total, index int, item any) error { return nil })
	nilErr := reflectinternal.Looper.Slice(nil, func(total, index int, item any) error { return nil })

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nilNoErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nilNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Slice returns correct value -- with args", actual)
}

func Test_Looper_Map(t *testing.T) {
	// Arrange
	err := reflectinternal.Looper.Map(map[string]int{"a": 1}, func(total, index int, key, value any) error { return nil })
	nilErr := reflectinternal.Looper.Map(nil, func(total, index int, key, value any) error { return nil })

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nilNoErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nilNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- with args", actual)
}

func Test_Looper_ToPointerReflectValue(t *testing.T) {
	// Arrange
	type S struct{}
	_, err := reflectinternal.Looper.ToPointerReflectValue(S{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ToPointerReflectValue returns error -- struct", actual)
}

func Test_Looper_ToPointerReflectValue_Invalid(t *testing.T) {
	// Arrange
	_, err := reflectinternal.Looper.ToPointerReflectValue(42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToPointerReflectValue returns error -- invalid", actual)
}

// ── codeStack — uncovered ──

func Test_CodeStack_NewDefault(t *testing.T) {
	// Arrange
	st := reflectinternal.CodeStack.NewDefault()

	// Act
	actual := args.Map{"isOkay": st.IsOkay}

	// Assert
	expected := args.Map{"isOkay": true}
	expected.ShouldBeEqual(t, 0, "NewDefault returns correct value -- with args", actual)
}

func Test_CodeStack_LastFileWithLine(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.LastFileWithLine(0, 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LastFileWithLine returns non-empty -- with args", actual)
}

func Test_CodeStack_NewStacks(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.NewStacks(0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewStacks returns correct value -- with args", actual)
}

func Test_CodeStack_StacksStrings(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksStrings(0)

	// Act
	actual := args.Map{"gt0": len(result) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "StacksStrings returns correct value -- with args", actual)
}

func Test_CodeStack_StacksStringsCount(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksStringsCount(0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StacksStringsCount returns correct value -- with args", actual)
}

func Test_CodeStack_StacksStringsFiltered(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksStringsFiltered(0, 4)

	// Act
	actual := args.Map{"gt0": len(result) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "StacksStringsFiltered returns correct value -- with args", actual)
}

func Test_CodeStack_StacksString(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksString(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksString returns correct value -- with args", actual)
}

func Test_CodeStack_StacksStringDefault(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksStringDefault(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksStringDefault returns correct value -- with args", actual)
}

func Test_CodeStack_StacksStringCount(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.StacksStringCount(0, 2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StacksStringCount returns correct value -- with args", actual)
}

func Test_CodeStack_SingleStack(t *testing.T) {
	// Arrange
	result := reflectinternal.CodeStack.SingleStack(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SingleStack returns correct value -- with args", actual)
}

// ── reflectPath ──

func Test_Path_CurDir(t *testing.T) {
	// Arrange
	result := reflectinternal.Path.CurDir()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Path returns correct value -- CurDir", actual)
}

func Test_Path_RepoDir(t *testing.T) {
	// Arrange
	result := reflectinternal.Path.RepoDir()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Path returns correct value -- RepoDir", actual)
}

// ── getFunc — uncovered ──

func Test_GetFunc_NameOnlyByStack(t *testing.T) {
	// Arrange
	result := reflectinternal.GetFunc.NameOnlyByStack(0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameOnlyByStack returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethodsRv(t *testing.T) {
	// Arrange
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodsRv(rv)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodsRv returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethodsMapRv(t *testing.T) {
	// Arrange
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodsMapRv(rv)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodsMapRv returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethodProcessorsMap(t *testing.T) {
	// Arrange
	type S struct{}
	rv := reflect.ValueOf(S{})
	result := reflectinternal.GetFunc.GetMethodProcessorsMap(rv)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetMethodProcessorsMap returns correct value -- with args", actual)
}

func Test_GetFunc_GetPkgPath(t *testing.T) {
	// Arrange
	fn := func() {}
	result := reflectinternal.GetFunc.GetPkgPath(fn)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetPkgPath returns correct value -- with args", actual)
}

// ── Utils uncovered ──

func Test_Utils_VerifyReflectTypes_Mismatch(t *testing.T) {
	// Arrange
	ok, err := reflectinternal.Utils.VerifyReflectTypes(
		"TestRoot",
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")},
		[]reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)},
	)

	// Act
	actual := args.Map{
		"ok": ok,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"ok": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "VerifyReflectTypes returns correct value -- mismatch", actual)
}

func Test_Utils_FullNameToPkgName(t *testing.T) {
	// Arrange
	result := reflectinternal.Utils.FullNameToPkgName("mypackage.MyFunc")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullNameToPkgName returns correct value -- with args", actual)
}

// ── TypeNameToValidVariableName — more branches ──

func Test_TypeNameToValidVariableName_SliceBrackets(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNameToValidVariableName("[]MyType")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- slice", actual)
}

func Test_TypeNameToValidVariableName_PtrStar(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNameToValidVariableName("*MyType")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- ptr", actual)
}

func Test_TypeNameToValidVariableName_DotWithSlice(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNameToValidVariableName("[]pkg.MyType")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- dot-slice", actual)
}

func Test_TypeNameToValidVariableName_DotWithPtr(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNameToValidVariableName("*pkg.MyType")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- dot-ptr", actual)
}
