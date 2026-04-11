package reflectinternaltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// ── getFunc ──

func Test_GetFunc_FullName(t *testing.T) {
	// Arrange
	fn := func() {}
	fullName := reflectinternal.GetFunc.FullName(fn)
	nilName := reflectinternal.GetFunc.FullName(nil)

	// Act
	actual := args.Map{
		"hasName": fullName != "",
		"nilName": nilName,
	}

	// Assert
	expected := args.Map{
		"hasName": true,
		"nilName": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FullName returns correct value -- with args", actual)
}

func Test_GetFunc_FullNameWithName(t *testing.T) {
	// Arrange
	fn := func() {}
	fullName, name := reflectinternal.GetFunc.FullNameWithName(fn)
	nilFull, nilName := reflectinternal.GetFunc.FullNameWithName(nil)

	// Act
	actual := args.Map{
		"hasFullName": fullName != "",
		"hasName":     name != "",
		"nilFull":     nilFull,
		"nilName":     nilName,
	}

	// Assert
	expected := args.Map{
		"hasFullName": true,
		"hasName":     true,
		"nilFull":     "",
		"nilName":     "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FullNameWithName returns non-empty -- with args", actual)
}

func Test_GetFunc_NameOnly(t *testing.T) {
	// Arrange
	fn := func() {}
	name := reflectinternal.GetFunc.NameOnly(fn)
	nilName := reflectinternal.GetFunc.NameOnly(nil)

	// Act
	actual := args.Map{
		"hasName": name != "",
		"nilName": nilName,
	}

	// Assert
	expected := args.Map{
		"hasName": true,
		"nilName": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_NameOnly returns correct value -- with args", actual)
}

func Test_GetFunc_All(t *testing.T) {
	// Arrange
	full, pkg, method := reflectinternal.GetFunc.All("mypackage.MyFunc")
	emFull, emPkg, emMethod := reflectinternal.GetFunc.All("")

	// Act
	actual := args.Map{
		"full":     full,
		"pkg":      pkg,
		"method":   method,
		"emFull":   emFull,
		"emPkg":    emPkg,
		"emMethod": emMethod,
	}

	// Assert
	expected := args.Map{
		"full":     "mypackage.MyFunc",
		"pkg":      "mypackage",
		"method":   "MyFunc",
		"emFull":   "",
		"emPkg":    "",
		"emMethod": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_All returns correct value -- with args", actual)
}

func Test_GetFunc_FuncDirectInvokeName(t *testing.T) {
	// Arrange
	fn := func() {}
	name := reflectinternal.GetFunc.FuncDirectInvokeName(fn)
	emptyName := reflectinternal.GetFunc.FuncDirectInvokeNameUsingFullName("")

	// Act
	actual := args.Map{
		"hasName":  name != "",
		"emptyRes": emptyName,
	}

	// Assert
	expected := args.Map{
		"hasName":  true,
		"emptyRes": "",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_FuncDirectInvokeName returns correct value -- with args", actual)
}

func Test_GetFunc_PascalFuncName(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":  reflectinternal.GetFunc.PascalFuncName("hello"),
		"single":  reflectinternal.GetFunc.PascalFuncName("h"),
		"empty":   reflectinternal.GetFunc.PascalFuncName(""),
		"already": reflectinternal.GetFunc.PascalFuncName("Hello"),
	}

	// Assert
	expected := args.Map{
		"simple":  "Hello",
		"single":  "H",
		"empty":   "",
		"already": "Hello",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_PascalFuncName returns correct value -- with args", actual)
}

func Test_GetFunc_GetPkgPathFullName(t *testing.T) {
	// Arrange
	result := reflectinternal.GetFunc.GetPkgPathFullName("github.com/org/repo/pkg.Func")
	empty := reflectinternal.GetFunc.GetPkgPathFullName("")
	noslash := reflectinternal.GetFunc.GetPkgPathFullName("pkg.Func")

	// Act
	actual := args.Map{
		"result":  result != "",
		"empty":   empty,
		"noslash": noslash,
	}

	// Assert
	expected := args.Map{
		"result":  true,
		"empty":   "",
		"noslash": "pkg.Func",
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetPkgPathFullName returns correct value -- with args", actual)
}

func Test_GetFunc_RunTime(t *testing.T) {
	// Arrange
	fn := func() {}
	rt := reflectinternal.GetFunc.RunTime(fn)
	nilRt := reflectinternal.GetFunc.RunTime(nil)
	intRt := reflectinternal.GetFunc.RunTime(42)

	// Act
	actual := args.Map{
		"rtNotNil": rt != nil,
		"nilRt":    nilRt == nil,
		"intRt":    intRt == nil,
	}

	// Assert
	expected := args.Map{
		"rtNotNil": true,
		"nilRt":    true,
		"intRt":    true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_RunTime returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethod(t *testing.T) {
	// Arrange
	type myStruct struct{}
	s := myStruct{}

	nilMethod := reflectinternal.GetFunc.GetMethod("", s)
	missingMethod := reflectinternal.GetFunc.GetMethod("NotExist", s)
	nilItem := reflectinternal.GetFunc.GetMethod("Name", nil)

	// Act
	actual := args.Map{
		"nilMethod":     nilMethod == nil,
		"missingMethod": missingMethod == nil,
		"nilItem":       nilItem == nil,
	}

	// Assert
	expected := args.Map{
		"nilMethod":     true,
		"missingMethod": true,
		"nilItem":       true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethod returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethods(t *testing.T) {
	// Arrange
	type myStruct struct{}
	s := myStruct{}
	methods := reflectinternal.GetFunc.GetMethods(s)
	nilMethods := reflectinternal.GetFunc.GetMethods(nil)

	// Act
	actual := args.Map{
		"methodsNotNil":    methods != nil,
		"nilMethodsNotNil": nilMethods != nil,
	}

	// Assert
	expected := args.Map{
		"methodsNotNil":    true,
		"nilMethodsNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethods returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethodsNames(t *testing.T) {
	// Arrange
	type myStruct struct{}
	s := myStruct{}
	names := reflectinternal.GetFunc.GetMethodsNames(s)
	nilNames := reflectinternal.GetFunc.GetMethodsNames(nil)

	// Act
	actual := args.Map{
		"namesNotNil":    names != nil,
		"nilNamesNotNil": nilNames != nil,
	}

	// Assert
	expected := args.Map{
		"namesNotNil":    names != nil,
		"nilNamesNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethodsNames returns correct value -- with args", actual)
}

func Test_GetFunc_GetMethodsMap(t *testing.T) {
	// Arrange
	type myStruct struct{}
	s := myStruct{}
	m := reflectinternal.GetFunc.GetMethodsMap(s)
	nilMap := reflectinternal.GetFunc.GetMethodsMap(nil)

	// Act
	actual := args.Map{
		"mapNotNil":    m != nil,
		"nilMapNotNil": nilMap != nil,
	}

	// Assert
	expected := args.Map{
		"mapNotNil":    true,
		"nilMapNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFunc_GetMethodsMap returns correct value -- with args", actual)
}

// ── reflectUtils ──

func Test_Utils_MaxLimit(t *testing.T) {
	// Act
	actual := args.Map{
		"noMax":     reflectinternal.Utils.MaxLimit(10, -1),
		"underMax":  reflectinternal.Utils.MaxLimit(5, 10),
		"overMax":   reflectinternal.Utils.MaxLimit(15, 10),
		"equalMax":  reflectinternal.Utils.MaxLimit(10, 10),
	}

	// Assert
	expected := args.Map{
		"noMax":     10,
		"underMax":  5,
		"overMax":   10,
		"equalMax":  10,
	}
	expected.ShouldBeEqual(t, 0, "Utils_MaxLimit returns correct value -- with args", actual)
}

func Test_Utils_AppendArgs(t *testing.T) {
	// Arrange
	emptyResult := reflectinternal.Utils.AppendArgs("item", []any{})
	withResult := reflectinternal.Utils.AppendArgs("item", []any{"a", "b"})

	// Act
	actual := args.Map{
		"emptyLen": len(emptyResult),
		"withLen":  len(withResult),
		"first":    emptyResult[0],
	}

	// Assert
	expected := args.Map{
		"emptyLen": 1,
		"withLen":  3,
		"first":    "item",
	}
	expected.ShouldBeEqual(t, 0, "Utils_AppendArgs returns correct value -- with args", actual)
}

func Test_Utils_PkgNameOnly(t *testing.T) {
	// Arrange
	fn := func() {}
	name := reflectinternal.Utils.PkgNameOnly(fn)

	// Act
	actual := args.Map{
		"hasName": name != "",
	}

	// Assert
	expected := args.Map{
		"hasName": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_PkgNameOnly returns correct value -- with args", actual)
}

func Test_Utils_FullNameToPkgName_FromGetFuncFullName(t *testing.T) {
	// Arrange
	name := reflectinternal.Utils.FullNameToPkgName("mypackage.MyFunc")

	// Act
	actual := args.Map{
		"name": name,
	}

	// Assert
	expected := args.Map{
		"name": "mypackage",
	}
	expected.ShouldBeEqual(t, 0, "Utils_FullNameToPkgName returns correct value -- with args", actual)
}

func Test_Utils_IsReflectTypeMatch(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	anyType := reflect.TypeOf((*any)(nil)).Elem()

	ok1, err1 := reflectinternal.Utils.IsReflectTypeMatch(intType, intType)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatch(intType, strType)
	ok3, err3 := reflectinternal.Utils.IsReflectTypeMatch(anyType, strType)

	// Act
	actual := args.Map{
		"sameOk":   ok1,
		"sameErr":  err1 == nil,
		"diffOk":   ok2,
		"diffErr":  err2 != nil,
		"anyOk":    ok3,
		"anyNoErr": err3 == nil,
	}

	// Assert
	expected := args.Map{
		"sameOk":   true,
		"sameErr":  true,
		"diffOk":   false,
		"diffErr":  true,
		"anyOk":    true,
		"anyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_IsReflectTypeMatch returns correct value -- with args", actual)
}

func Test_Utils_IsReflectTypeMatchAny(t *testing.T) {
	// Arrange
	ok1, err1 := reflectinternal.Utils.IsReflectTypeMatchAny(42, 99)
	ok2, err2 := reflectinternal.Utils.IsReflectTypeMatchAny(42, "hello")

	// Act
	actual := args.Map{
		"sameOk":  ok1,
		"sameErr": err1 == nil,
		"diffOk":  ok2,
		"diffErr": err2 != nil,
	}

	// Assert
	expected := args.Map{
		"sameOk":  true,
		"sameErr": true,
		"diffOk":  false,
		"diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_IsReflectTypeMatchAny returns correct value -- with args", actual)
}

func Test_Utils_VerifyReflectTypesAny(t *testing.T) {
	// Arrange
	ok1, err1 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42, "hello"},
		[]any{99, "world"},
	)
	ok2, err2 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42},
		[]any{42, "extra"},
	)
	ok3, err3 := reflectinternal.Utils.VerifyReflectTypesAny(
		[]any{42, "hello"},
		[]any{99, 100},
	)

	// Act
	actual := args.Map{
		"matchOk":     ok1,
		"matchErr":    err1 == nil,
		"lenMismatch": ok2,
		"lenErr":      err2 != nil,
		"typeMismatch": ok3,
		"typeErr":     err3 != nil,
	}

	// Assert
	expected := args.Map{
		"matchOk":     true,
		"matchErr":    true,
		"lenMismatch": false,
		"lenErr":      true,
		"typeMismatch": false,
		"typeErr":     true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_VerifyReflectTypesAny returns correct value -- with args", actual)
}

func Test_Utils_VerifyReflectTypes(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")

	ok1, err1 := reflectinternal.Utils.VerifyReflectTypes(
		"test",
		[]reflect.Type{intType, strType},
		[]reflect.Type{intType, strType},
	)
	ok2, err2 := reflectinternal.Utils.VerifyReflectTypes(
		"test",
		[]reflect.Type{intType},
		[]reflect.Type{intType, strType},
	)

	// Act
	actual := args.Map{
		"matchOk":  ok1,
		"matchErr": err1 == nil,
		"diffOk":   ok2,
		"diffErr":  err2 != nil,
	}

	// Assert
	expected := args.Map{
		"matchOk":  true,
		"matchErr": true,
		"diffOk":   false,
		"diffErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "Utils_VerifyReflectTypes returns correct value -- with args", actual)
}

// ── reflectGetter additional ──

func Test_ReflectGetter_FieldNameWithValuesMap(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		age  int
	}
	s := testStruct{Name: "Alice", age: 30}

	result, err := reflectinternal.ReflectGetter.FieldNameWithValuesMap(s)
	nilResult, nilErr := reflectinternal.ReflectGetter.FieldNameWithValuesMap(nil)

	// Act
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"nilLen":    len(nilResult),
		"nilHasErr": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"nilLen":    0,
		"nilHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_FieldNameWithValuesMap returns non-empty -- with args", actual)
}

func Test_ReflectGetter_NullFieldsMap(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name *string
		Age  int
	}
	s := testStruct{Name: nil, Age: 0}

	result := reflectinternal.ReflectGetter.NullFieldsMap(s)
	nilResult := reflectinternal.ReflectGetter.NullFieldsMap(nil)

	// Act
	actual := args.Map{
		"hasNullName": result["Name"],
		"nilLen":      len(nilResult),
	}

	// Assert
	expected := args.Map{
		"hasNullName": true,
		"nilLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_NullFieldsMap returns correct value -- with args", actual)
}

func Test_ReflectGetter_NullOrZeroFieldsMap(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}
	s := testStruct{Name: "", Age: 0}

	result := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(s)
	nilResult := reflectinternal.ReflectGetter.NullOrZeroFieldsMap(nil)

	// Act
	actual := args.Map{
		"hasZeroName": result["Name"],
		"hasZeroAge":  result["Age"],
		"nilLen":      len(nilResult),
	}

	// Assert
	expected := args.Map{
		"hasZeroName": true,
		"hasZeroAge":  true,
		"nilLen":      0,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetter_NullOrZeroFieldsMap returns correct value -- with args", actual)
}

// ── reflectGetUsingReflectValue additional ──

func Test_ReflectGetRv_FieldNameWithTypeMap(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}
	s := testStruct{Name: "Alice", Age: 30}
	rv := reflect.ValueOf(s)
	result := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(rv)

	nilResult := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithTypeMap(
		reflect.ValueOf(42),
	)

	// Act
	actual := args.Map{
		"resultLen": len(result),
		"nilResult": nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"resultLen": 2,
		"nilResult": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetRv_FieldNameWithTypeMap returns non-empty -- with args", actual)
}

func Test_ReflectGetRv_FieldNameWithValuesMap(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		age  int
	}
	s := testStruct{Name: "Alice", age: 30}
	rv := reflect.ValueOf(s)

	result, err := reflectinternal.ReflectGetterUsingReflectValue.FieldNameWithValuesMap(rv)

	// Act
	actual := args.Map{
		"resultLen": len(result),
		"noError":   err == nil,
		"hasName":   result["Name"] == "Alice",
	}

	// Assert
	expected := args.Map{
		"resultLen": 2,
		"noError":   true,
		"hasName":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectGetRv_FieldNameWithValuesMap returns non-empty -- with args", actual)
}

// ── MapConverter additional ──

func Test_MapConverter_ToKeysStrings_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToKeysStrings(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"len":   len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len":   1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToKeysStrings returns correct value -- with args", actual)
}

func Test_MapConverter_ToValuesAny_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToValuesAny(map[string]int{"a": 1})
	nilResult, nilErr := reflectinternal.MapConverter.ToValuesAny(nil)

	// Act
	actual := args.Map{
		"len":    len(result),
		"noErr":  err == nil,
		"nilLen": len(nilResult),
		"nilErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"len":    1,
		"noErr":  true,
		"nilLen": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToValuesAny returns non-empty -- with args", actual)
}

func Test_MapConverter_ToKeysAny_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result, err := reflectinternal.MapConverter.ToKeysAny(map[string]int{"a": 1})
	nilResult, nilErr := reflectinternal.MapConverter.ToKeysAny(nil)

	// Act
	actual := args.Map{
		"len":    len(result),
		"noErr":  err == nil,
		"nilLen": len(nilResult),
		"nilErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"len":    1,
		"noErr":  true,
		"nilLen": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToKeysAny returns correct value -- with args", actual)
}

func Test_MapConverter_ToStringsMust(t *testing.T) {
	// Arrange
	result := reflectinternal.MapConverter.ToStringsMust(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToStringsMust returns correct value -- with args", actual)
}

func Test_MapConverter_ToSortedStringsMust_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result := reflectinternal.MapConverter.ToSortedStringsMust(map[string]int{"b": 2, "a": 1})
	nilResult := reflectinternal.MapConverter.ToSortedStringsMust(nil)

	// Act
	actual := args.Map{
		"len":    len(result),
		"first":  result[0],
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len":    2,
		"first":  "a",
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToSortedStringsMust returns correct value -- with args", actual)
}

func Test_MapConverter_ToMapStringAnyRv_NonStringKey(t *testing.T) {
	// Arrange
	m := map[int]string{1: "one", 2: "two"}
	rv := reflect.ValueOf(m)
	result, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	// Act
	actual := args.Map{
		"len":   len(result),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len":   2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToMapStringAnyRv_NonStringKey returns correct value -- with args", actual)
}

func Test_MapConverter_ToMapStringAnyRv_NotMap(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToMapStringAnyRv(rv)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToMapStringAnyRv_NotMap returns correct value -- with args", actual)
}

func Test_MapConverter_ToStringsRv_NotMap_FromGetFuncFullName(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	_, err := reflectinternal.MapConverter.ToStringsRv(rv)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter_ToStringsRv_NotMap returns correct value -- with args", actual)
}

// ── reflectPath ──

func Test_Path_RepoDir_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result := reflectinternal.Path.RepoDir()

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Path_RepoDir returns correct value -- with args", actual)
}

func Test_Path_CurDir_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result := reflectinternal.Path.CurDir()

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Path_CurDir returns correct value -- with args", actual)
}

// ── TypeNamesReferenceString ──

func Test_TypeNamesReferenceString(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNamesReferenceString(true, 42, "hello")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeNamesReferenceString returns correct value -- with args", actual)
}

// ── TypeNamesString ──

func Test_TypeNamesString(t *testing.T) {
	// Arrange
	result := reflectinternal.TypeNamesString(true, 42, "hello")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeNamesString returns correct value -- with args", actual)
}

// ── ReflectType converter ──

func Test_ReflectType_SafeName_FromGetFuncFullName(t *testing.T) {
	// Arrange
	result := reflectinternal.ReflectType.SafeName(42)
	nilResult := reflectinternal.ReflectType.SafeName(nil)

	// Act
	actual := args.Map{
		"result":    result,
		"nilResult": nilResult,
	}

	// Assert
	expected := args.Map{
		"result":    "int",
		"nilResult": "",
	}
	expected.ShouldBeEqual(t, 0, "ReflectType_SafeName returns correct value -- with args", actual)
}
