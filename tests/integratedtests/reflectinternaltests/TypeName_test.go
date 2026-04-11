package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ── TypeName ──

func Test_TypeName_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"int":     reflectinternal.TypeName(42),
		"string":  reflectinternal.TypeName("hello"),
		"nil":     reflectinternal.TypeName(nil),
		"slice":   reflectinternal.TypeName([]int{1, 2}),
	}
	expected := args.Map{
		"int":     "int",
		"string":  "string",
		"nil":     "",
		"slice":   "[]int",
	}
	expected.ShouldBeEqual(t, 0, "TypeName returns correct value -- with args", actual)
}

// ── TypeNames ──

func Test_TypeNames_FromTypeName(t *testing.T) {
	// Act
	fullNames := reflectinternal.TypeNames(true, 42, "hello")
	shortNames := reflectinternal.TypeNames(false, 42, "hello")

	// Assert
	actual := args.Map{
		"fullLen":   len(fullNames),
		"shortLen":  len(shortNames),
		"fullFirst": fullNames[0],
	}
	expected := args.Map{
		"fullLen":   2,
		"shortLen":  2,
		"fullFirst": "int",
	}
	expected.ShouldBeEqual(t, 0, "TypeNames returns correct value -- with args", actual)
}

// ── TypeNameToValidVariableName ──

func Test_TypeNameToValidVariableName_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":   reflectinternal.TypeNameToValidVariableName("MyType"),
		"empty":    reflectinternal.TypeNameToValidVariableName(""),
		"withDot":  reflectinternal.TypeNameToValidVariableName("pkg.MyType"),
		"withPtr":  reflectinternal.TypeNameToValidVariableName("*MyType"),
		"withSlice": reflectinternal.TypeNameToValidVariableName("[]MyType"),
	}
	expected := args.Map{
		"simple":   "MyType",
		"empty":    "",
		"withDot":  reflectinternal.TypeNameToValidVariableName("pkg.MyType"),
		"withPtr":  reflectinternal.TypeNameToValidVariableName("*MyType"),
		"withSlice": reflectinternal.TypeNameToValidVariableName("[]MyType"),
	}
	expected.ShouldBeEqual(t, 0, "TypeNameToValidVariableName returns non-empty -- with args", actual)
}

// ── isChecker ──

func Test_Is_Null(t *testing.T) {
	// Arrange
	var nilPtr *int
	val := 42

	// Act
	actual := args.Map{
		"nilIsNull":     reflectinternal.Is.Null(nil),
		"nilPtrIsNull":  reflectinternal.Is.Null(nilPtr),
		"valueIsNull":   reflectinternal.Is.Null(val),
		"nilIsNotNull":  reflectinternal.Is.NotNull(nil),
		"nilIsDefined":  reflectinternal.Is.Defined(nil),
		"valIsDefined":  reflectinternal.Is.Defined(val),
	}
	expected := args.Map{
		"nilIsNull":     true,
		"nilPtrIsNull":  true,
		"valueIsNull":   false,
		"nilIsNotNull":  false,
		"nilIsDefined":  false,
		"valIsDefined":  true,
	}
	expected.ShouldBeEqual(t, 0, "Is_Null returns correct value -- with args", actual)
}

func Test_Is_NullRv(t *testing.T) {
	// Arrange
	var nilSlice []int
	slice := []int{1, 2, 3}

	// Act
	actual := args.Map{
		"nilSlice": reflectinternal.Is.NullRv(reflect.ValueOf(nilSlice)),
		"slice":    reflectinternal.Is.NullRv(reflect.ValueOf(slice)),
		"int":      reflectinternal.Is.NullRv(reflect.ValueOf(42)),
	}
	expected := args.Map{
		"nilSlice": true,
		"slice":    false,
		"int":      false,
	}
	expected.ShouldBeEqual(t, 0, "Is_NullRv returns correct value -- with args", actual)
}

func Test_Is_TypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"numberInt":     reflectinternal.Is.Number(42),
		"numberFloat":   reflectinternal.Is.Number(3.14),
		"numberString":  reflectinternal.Is.Number("hello"),
		"stringCheck":   reflectinternal.Is.String("hello"),
		"stringInt":     reflectinternal.Is.String(42),
		"boolCheck":     reflectinternal.Is.Boolean(true),
		"boolInt":       reflectinternal.Is.Boolean(42),
		"primitiveInt":  reflectinternal.Is.Primitive(42),
		"primitiveStr":  reflectinternal.Is.Primitive("hello"),
		"primitiveBool": reflectinternal.Is.Primitive(true),
	}
	expected := args.Map{
		"numberInt":     true,
		"numberFloat":   true,
		"numberString":  false,
		"stringCheck":   true,
		"stringInt":     false,
		"boolCheck":     true,
		"boolInt":       false,
		"primitiveInt":  true,
		"primitiveStr":  true,
		"primitiveBool": true,
	}
	expected.ShouldBeEqual(t, 0, "Is_TypeChecks returns correct value -- with args", actual)
}

func Test_Is_Pointer_FromTypeName(t *testing.T) {
	// Arrange
	val := 42

	// Act
	actual := args.Map{
		"ptrIsPtr":   reflectinternal.Is.Pointer(&val),
		"valNotPtr":  reflectinternal.Is.Pointer(val),
	}
	expected := args.Map{
		"ptrIsPtr":   true,
		"valNotPtr":  false,
	}
	expected.ShouldBeEqual(t, 0, "Is_Pointer returns correct value -- with args", actual)
}

func Test_Is_Function_FromTypeName(t *testing.T) {
	// Arrange
	fn := func() {}

	// Act
	actual := args.Map{
		"funcIsFunc": reflectinternal.Is.Function(fn),
		"intNotFunc": reflectinternal.Is.Function(42),
		"isFunc":     reflectinternal.Is.Func(fn),
		"isNotFunc":  reflectinternal.Is.NotFunc(42),
		"nilIsFunc":  reflectinternal.Is.Func(nil),
	}
	expected := args.Map{
		"funcIsFunc": true,
		"intNotFunc": false,
		"isFunc":     true,
		"isNotFunc":  true,
		"nilIsFunc":  true,
	}
	expected.ShouldBeEqual(t, 0, "Is_Function returns correct value -- with args", actual)
}

func Test_Is_FuncTypeOf(t *testing.T) {
	// Act
	actual := args.Map{
		"funcType":   reflectinternal.Is.FuncTypeOf(reflect.TypeOf(func() {})),
		"intType":    reflectinternal.Is.FuncTypeOf(reflect.TypeOf(42)),
	}
	expected := args.Map{
		"funcType":   true,
		"intType":    false,
	}
	expected.ShouldBeEqual(t, 0, "Is_FuncTypeOf returns correct value -- with args", actual)
}

func Test_Is_SliceOrArrayOf_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"sliceType": reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf([]int{})),
		"intType":   reflectinternal.Is.SliceOrArrayOf(reflect.TypeOf(42)),
	}
	expected := args.Map{
		"sliceType": true,
		"intType":   false,
	}
	expected.ShouldBeEqual(t, 0, "Is_SliceOrArrayOf returns correct value -- with args", actual)
}

func Test_Is_NumberKind(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    reflectinternal.Is.NumberKind(reflect.Int),
		"float":  reflectinternal.Is.NumberKind(reflect.Float64),
		"string": reflectinternal.Is.NumberKind(reflect.String),
		"bool":   reflectinternal.Is.NumberKind(reflect.Bool),
	}
	expected := args.Map{
		"int":    true,
		"float":  true,
		"string": false,
		"bool":   false,
	}
	expected.ShouldBeEqual(t, 0, "Is_NumberKind returns correct value -- with args", actual)
}

func Test_Is_PrimitiveKind(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    reflectinternal.Is.PrimitiveKind(reflect.Int),
		"string": reflectinternal.Is.PrimitiveKind(reflect.String),
		"bool":   reflectinternal.Is.PrimitiveKind(reflect.Bool),
		"struct": reflectinternal.Is.PrimitiveKind(reflect.Struct),
	}
	expected := args.Map{
		"int":    true,
		"string": true,
		"bool":   true,
		"struct": false,
	}
	expected.ShouldBeEqual(t, 0, "Is_PrimitiveKind returns correct value -- with args", actual)
}

func Test_Is_Zero_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"zeroInt":    reflectinternal.Is.Zero(0),
		"nonZeroInt": reflectinternal.Is.Zero(42),
		"emptyStr":   reflectinternal.Is.Zero(""),
		"nonEmptyStr": reflectinternal.Is.Zero("hello"),
		"nilVal":     reflectinternal.Is.Zero(nil),
	}
	expected := args.Map{
		"zeroInt":    true,
		"nonZeroInt": false,
		"emptyStr":   true,
		"nonEmptyStr": false,
		"nilVal":     true,
	}
	expected.ShouldBeEqual(t, 0, "Is_Zero returns correct value -- with args", actual)
}

func Test_Is_ZeroRv_Struct_FromTypeName(t *testing.T) {
	// Arrange
	type testStruct struct{ X int }

	// Act
	actual := args.Map{
		"zeroStruct":    reflectinternal.Is.ZeroRv(reflect.ValueOf(testStruct{})),
		"nonZeroStruct": reflectinternal.Is.ZeroRv(reflect.ValueOf(testStruct{X: 1})),
	}
	expected := args.Map{
		"zeroStruct":    true,
		"nonZeroStruct": false,
	}
	expected.ShouldBeEqual(t, 0, "Is_ZeroRv_Struct returns correct value -- with args", actual)
}

func Test_Is_ZeroRv_Array_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"zeroArray":    reflectinternal.Is.ZeroRv(reflect.ValueOf([3]int{})),
		"nonZeroArray": reflectinternal.Is.ZeroRv(reflect.ValueOf([3]int{1, 0, 0})),
	}
	expected := args.Map{
		"zeroArray":    true,
		"nonZeroArray": false,
	}
	expected.ShouldBeEqual(t, 0, "Is_ZeroRv_Array returns correct value -- with args", actual)
}

func Test_Is_Struct_FromTypeName(t *testing.T) {
	// Arrange
	type testStruct struct{ X int }
	s := testStruct{X: 1}

	// Act
	actual := args.Map{
		"structVal":  reflectinternal.Is.Struct(s),
		"structPtr":  reflectinternal.Is.Struct(&s),
		"intVal":     reflectinternal.Is.Struct(42),
		"rvStruct":   reflectinternal.Is.StructRv(reflect.ValueOf(s)),
		"rvInt":      reflectinternal.Is.StructRv(reflect.ValueOf(42)),
	}
	expected := args.Map{
		"structVal":  true,
		"structPtr":  true,
		"intVal":     false,
		"rvStruct":   true,
		"rvInt":      false,
	}
	expected.ShouldBeEqual(t, 0, "Is_Struct returns correct value -- with args", actual)
}

func Test_Is_Conclusive_FromTypeName(t *testing.T) {
	// Act
	eq1, conc1 := reflectinternal.Is.Conclusive(42, 42)
	eq2, conc2 := reflectinternal.Is.Conclusive(nil, nil)
	eq3, conc3 := reflectinternal.Is.Conclusive(42, nil)
	eq4, conc4 := reflectinternal.Is.Conclusive(42, "hello")

	// Assert
	actual := args.Map{
		"sameEq":     eq1,
		"sameConc":   conc1,
		"nilBothEq":  eq2,
		"nilBothConc": conc2,
		"nilOneEq":   eq3,
		"nilOneConc": conc3,
		"diffTypeEq": eq4,
		"diffTypeConc": conc4,
	}
	expected := args.Map{
		"sameEq":     true,
		"sameConc":   true,
		"nilBothEq":  true,
		"nilBothConc": true,
		"nilOneEq":   false,
		"nilOneConc": true,
		"diffTypeEq": false,
		"diffTypeConc": true,
	}
	expected.ShouldBeEqual(t, 0, "Is_Conclusive returns correct value -- with args", actual)
}

func Test_Is_AnyEqual_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"sameInt":    reflectinternal.Is.AnyEqual(42, 42),
		"diffInt":    reflectinternal.Is.AnyEqual(42, 43),
		"sameSlice":  reflectinternal.Is.AnyEqual([]int{1, 2}, []int{1, 2}),
		"diffSlice":  reflectinternal.Is.AnyEqual([]int{1, 2}, []int{3, 4}),
		"nilBoth":    reflectinternal.Is.AnyEqual(nil, nil),
	}
	expected := args.Map{
		"sameInt":    true,
		"diffInt":    false,
		"sameSlice":  true,
		"diffSlice":  false,
		"nilBoth":    true,
	}
	expected.ShouldBeEqual(t, 0, "Is_AnyEqual returns correct value -- with args", actual)
}

// ── Converter ──

func Test_Converter_ArgsToReflectValues(t *testing.T) {
	// Act
	result := reflectinternal.Converter.ArgsToReflectValues([]any{1, "hello", true})
	emptyResult := reflectinternal.Converter.ArgsToReflectValues([]any{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(emptyResult),
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "Converter_ArgsToReflectValues returns non-empty -- with args", actual)
}

func Test_Converter_ReflectValuesToInterfaces(t *testing.T) {
	// Arrange
	rvs := []reflect.Value{reflect.ValueOf(42), reflect.ValueOf("hello")}

	// Act
	result := reflectinternal.Converter.ReflectValuesToInterfaces(rvs)
	emptyResult := reflectinternal.Converter.ReflectValuesToInterfaces([]reflect.Value{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(emptyResult),
	}
	expected := args.Map{
		"resultLen": 2,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "Converter_ReflectValuesToInterfaces returns non-empty -- with args", actual)
}

func Test_Converter_InterfacesToTypes_FromTypeName(t *testing.T) {
	// Act
	types := reflectinternal.Converter.InterfacesToTypes([]any{42, "hello"})
	names := reflectinternal.Converter.InterfacesToTypesNames([]any{42, "hello"})
	namesVals := reflectinternal.Converter.InterfacesToTypesNamesWithValues([]any{42})
	emptyTypes := reflectinternal.Converter.InterfacesToTypes([]any{})

	// Assert
	actual := args.Map{
		"typesLen":    len(types),
		"namesLen":    len(names),
		"namesVLen":   len(namesVals),
		"emptyLen":    len(emptyTypes),
	}
	expected := args.Map{
		"typesLen":    2,
		"namesLen":    2,
		"namesVLen":   1,
		"emptyLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "Converter_InterfacesToTypes returns correct value -- with args", actual)
}

func Test_Converter_ReflectValToInterfaces(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3}

	// Act
	result := reflectinternal.Converter.ReflectValToInterfaces(false, reflect.ValueOf(slice))
	ptrResult := reflectinternal.Converter.ReflectValToInterfaces(false, reflect.ValueOf(&slice))
	notSlice := reflectinternal.Converter.ReflectValToInterfaces(false, reflect.ValueOf(42))
	emptySlice := reflectinternal.Converter.ReflectValToInterfaces(false, reflect.ValueOf([]int{}))

	// Assert
	actual := args.Map{
		"resultLen":   len(result),
		"ptrLen":      len(ptrResult),
		"notSliceLen": len(notSlice),
		"emptyLen":    len(emptySlice),
	}
	expected := args.Map{
		"resultLen":   3,
		"ptrLen":      3,
		"notSliceLen": 0,
		"emptyLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "Converter_ReflectValToInterfaces returns correct value -- with args", actual)
}

func Test_Converter_ReflectInterfaceVal_FromTypeName(t *testing.T) {
	// Arrange
	val := 42

	// Act
	result := reflectinternal.Converter.ReflectInterfaceVal(val)
	ptrResult := reflectinternal.Converter.ReflectInterfaceVal(&val)

	// Assert
	actual := args.Map{
		"result":    result,
		"ptrResult": ptrResult,
	}
	expected := args.Map{
		"result":    int64(42),
		"ptrResult": 42,
	}
	expected.ShouldBeEqual(t, 0, "Converter_ReflectInterfaceVal returns correct value -- with args", actual)
}

// ── SliceConverter ──

func Test_SliceConverter_Length_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"sliceLen": reflectinternal.SliceConverter.Length([]int{1, 2, 3}),
		"mapLen":   reflectinternal.SliceConverter.Length(map[string]int{"a": 1}),
		"nilLen":   reflectinternal.SliceConverter.Length(nil),
		"intLen":   reflectinternal.SliceConverter.Length(42),
	}
	expected := args.Map{
		"sliceLen": 3,
		"mapLen":   1,
		"nilLen":   0,
		"intLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "SliceConverter_Length returns correct value -- with args", actual)
}

func Test_SliceConverter_ToStrings_FromTypeName(t *testing.T) {
	// Act
	result, err := reflectinternal.SliceConverter.ToStrings([]int{1, 2, 3})
	emptyResult, emptyErr := reflectinternal.SliceConverter.ToStrings([]int{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"emptyLen":  len(emptyResult),
		"emptyNoErr": emptyErr == nil,
	}
	expected := args.Map{
		"resultLen": 3,
		"noError":   true,
		"emptyLen":  0,
		"emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceConverter_ToStrings returns correct value -- with args", actual)
}

func Test_SliceConverter_ToStringsMust_FromTypeName(t *testing.T) {
	// Act
	result := reflectinternal.SliceConverter.ToStringsMust([]int{1, 2})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
	}
	expected := args.Map{
		"resultLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "SliceConverter_ToStringsMust returns correct value -- with args", actual)
}

func Test_SliceConverter_ToAnyItemsAsync_FromTypeName(t *testing.T) {
	// Act
	result := reflectinternal.SliceConverter.ToAnyItemsAsync([]int{1, 2, 3})
	nilResult := reflectinternal.SliceConverter.ToAnyItemsAsync(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"nilLen":    len(nilResult),
	}
	expected := args.Map{
		"resultLen": 3,
		"nilLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "SliceConverter_ToAnyItemsAsync returns correct value -- with args", actual)
}

// ── MapConverter ──

func Test_MapConverter_ToStrings(t *testing.T) {
	// Act
	result, err := reflectinternal.MapConverter.ToStrings(map[string]int{"a": 1, "b": 2})
	nilResult, nilErr := reflectinternal.MapConverter.ToStrings(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilNoErr":  nilErr == nil,
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"nilLen":    0,
		"nilNoErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToStrings returns correct value -- with args", actual)
}

func Test_MapConverter_ToSortedStrings_FromTypeName(t *testing.T) {
	// Act
	result, err := reflectinternal.MapConverter.ToSortedStrings(map[string]int{"b": 2, "a": 1})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"first":     result[0],
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"first":     "a",
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToSortedStrings returns correct value -- with args", actual)
}

func Test_MapConverter_ToKeysValuesAny_FromTypeName(t *testing.T) {
	// Act
	keys, vals, err := reflectinternal.MapConverter.ToKeysValuesAny(map[string]int{"a": 1})
	nilKeys, nilVals, nilErr := reflectinternal.MapConverter.ToKeysValuesAny(nil)

	// Assert
	actual := args.Map{
		"keysLen":    len(keys),
		"valsLen":    len(vals),
		"noError":    err == nil,
		"nilKeysLen": len(nilKeys),
		"nilValsLen": len(nilVals),
		"nilNoErr":   nilErr == nil,
	}
	expected := args.Map{
		"keysLen":    1,
		"valsLen":    1,
		"noError":    true,
		"nilKeysLen": 0,
		"nilValsLen": 0,
		"nilNoErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToKeysValuesAny returns non-empty -- with args", actual)
}

func Test_MapConverter_ToMapStringAny_FromTypeName(t *testing.T) {
	// Act
	result, err := reflectinternal.MapConverter.ToMapStringAny(map[string]int{"a": 1})
	nilResult, nilErr := reflectinternal.MapConverter.ToMapStringAny(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilNoErr":  nilErr == nil,
	}
	expected := args.Map{
		"resultLen": 1,
		"noError":   true,
		"nilLen":    0,
		"nilNoErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToMapStringAny returns correct value -- with args", actual)
}

func Test_MapConverter_Length_FromTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"mapLen": reflectinternal.MapConverter.Length(map[string]int{"a": 1}),
	}
	expected := args.Map{
		"mapLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_Length returns correct value -- with args", actual)
}

// ── ReflectGetter ──

func Test_ReflectGetter_PublicValuesMapStruct_FromTypeName(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}
	s := testStruct{Name: "Alice", Age: 30}

	// Act
	result, err := reflectinternal.ReflectGetter.PublicValuesMapStruct(s)
	nilResult, nilErr := reflectinternal.ReflectGetter.PublicValuesMapStruct(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilHasErr": nilErr != nil,
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"nilLen":    0,
		"nilHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_PublicValuesMapStruct returns non-empty -- with args", actual)
}

func Test_ReflectGetter_FieldNamesMap_FromTypeName(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}

	// Act
	result, err := reflectinternal.ReflectGetter.FieldNamesMap(testStruct{})
	nilResult, nilErr := reflectinternal.ReflectGetter.FieldNamesMap(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilHasErr": nilErr != nil,
	}
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"nilLen":    0,
		"nilHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_FieldNamesMap returns correct value -- with args", actual)
}

func Test_ReflectGetter_StructFieldsMap_FromTypeName(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}

	// Act
	result := reflectinternal.ReflectGetter.StructFieldsMap(testStruct{})
	nilResult := reflectinternal.ReflectGetter.StructFieldsMap(nil)

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"nilLen":    len(nilResult),
	}
	expected := args.Map{
		"resultLen": 2,
		"nilLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_StructFieldsMap returns correct value -- with args", actual)
}
