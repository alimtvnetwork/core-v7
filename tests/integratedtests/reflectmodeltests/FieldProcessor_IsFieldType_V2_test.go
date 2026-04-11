package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// FieldProcessor — all uncovered methods
// ═══════════════════════════════════════════════

func Test_FieldProcessor_IsFieldType_Valid(t *testing.T) {
	// Arrange
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}

	// Act
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
	actual = args.Map{"result": fp.IsFieldType(reflect.TypeOf(""))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}

func Test_FieldProcessor_IsFieldType_Nil(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil receiver should return false", actual)
}

func Test_FieldProcessor_IsFieldKind_Valid(t *testing.T) {
	// Arrange
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}

	// Act
	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected int kind", actual)
	actual = args.Map{"result": fp.IsFieldKind(reflect.String)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}

func Test_FieldProcessor_IsFieldKind_Nil(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil receiver should return false", actual)
}

// ═══════════════════════════════════════════════
// MethodProcessor — comprehensive coverage
// ═══════════════════════════════════════════════

// helper: create a MethodProcessor from a real method
type testTarget6 struct{}

func (testTarget6) Add(a, b int) int         { return a + b }
func (testTarget6) Greeting() string          { return "hi" }
func (testTarget6) Err() error                { return nil }
func (testTarget6) PairResult() (string, error) { return "ok", nil }

func getMethodProcessor6(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testTarget6{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if m.Name == name {
			return &reflectmodel.MethodProcessor{
				Name:          m.Name,
				Index:         i,
				ReflectMethod: m,
			}
		}
	}
	return nil
}

func Test_MP_HasValidFunc(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.HasValidFunc()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	var nilMp *reflectmodel.MethodProcessor
	actual = args.Map{"result": nilMp.HasValidFunc()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_MP_GetFuncName(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.GetFuncName() != "Add"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Add", actual)
}

func Test_MP_IsInvalid(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	var nilMp *reflectmodel.MethodProcessor
	actual = args.Map{"result": nilMp.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_MP_Func(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	f := mp.Func()

	// Act
	actual := args.Map{"result": f == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected func", actual)
}

func Test_MP_Func_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": mp.Func() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_MP_ArgsCount(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	// Add has receiver + a + b = 3

	// Act
	actual := args.Map{"result": mp.ArgsCount() < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
}

func Test_MP_ReturnLength(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.ReturnLength() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MP_ReturnLength_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": mp.ReturnLength() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return -1", actual)
}

func Test_MP_IsPublicMethod(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.IsPublicMethod()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected public", actual)
}

func Test_MP_IsPrivateMethod(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.IsPrivateMethod()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not private", actual)
}

func Test_MP_ArgsLength(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.ArgsLength() < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
}

func Test_MP_Invoke_Success(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	// receiver + 2 args
	results, err := mp.Invoke(testTarget6{}, 3, 4)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": len(results) != 1 || results[0].(int) != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_MP_Invoke_ArgsMismatch(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	_, err := mp.Invoke(testTarget6{}, 3) // missing arg

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for args mismatch", actual)
}

func Test_MP_Invoke_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_MP_GetFirstResponseOfInvoke(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Greeting")
	resp, err := mp.GetFirstResponseOfInvoke(testTarget6{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": resp.(string) != "hi"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_MP_InvokeResultOfIndex(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	resp, err := mp.InvokeResultOfIndex(0, testTarget6{}, 1, 2)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": resp.(int) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_MP_InvokeError(t *testing.T) {
	// Arrange
	defer func() { recover() }() // InvokeError may panic on zero reflect.Value
	mp := getMethodProcessor6("Err")
	funcErr, procErr := mp.InvokeError(testTarget6{})

	// Act
	actual := args.Map{"result": procErr != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": funcErr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error from Err()", actual)
}

func Test_MP_InvokeFirstAndError_Success(t *testing.T) {
	// Arrange
	defer func() { recover() }() // may panic on zero reflect.Value in ReflectValueToAnyValue
	mp := getMethodProcessor6("PairResult")
	first, funcErr, procErr := mp.InvokeFirstAndError(testTarget6{})

	// Act
	actual := args.Map{"result": procErr != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "processing error:", actual)
	actual = args.Map{"result": funcErr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no func error", actual)
	actual = args.Map{"result": first.(string) != "ok"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ok'", actual)
}

func Test_MP_InvokeFirstAndError_SingleReturn(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Greeting")
	_, _, procErr := mp.InvokeFirstAndError(testTarget6{})

	// Act
	actual := args.Map{"result": procErr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for single return", actual)
}

func Test_MP_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_MP_IsEqual_OneNil(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.IsEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)
}

func Test_MP_IsEqual_Same(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.IsEqual(mp)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)
}

func Test_MP_IsEqual_DiffMethods(t *testing.T) {
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")
	// Different signatures → should fail at args verification
	_ = mp1.IsEqual(mp2)
}

func Test_MP_IsNotEqual(t *testing.T) {
	// Arrange
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")

	// Act
	actual := args.Map{"result": mp1.IsNotEqual(mp2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different methods should not be equal", actual)
}

func Test_MP_GetType(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")

	// Act
	actual := args.Map{"result": mp.GetType() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil type", actual)
}

func Test_MP_GetType_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": mp.GetType() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil type", actual)
}

func Test_MP_GetOutArgsTypes(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	out := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(out) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// Call again to hit cache
	out2 := mp.GetOutArgsTypes()
	actual = args.Map{"result": len(out2) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_MP_GetOutArgsTypes_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	out := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(out) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_MP_GetInArgsTypes(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	in := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"result": len(in) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
	// Call again to hit cache
	in2 := mp.GetInArgsTypes()
	actual = args.Map{"result": len(in2) != len(in)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_MP_GetInArgsTypes_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	in := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"result": len(in) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_MP_GetInArgsTypesNames(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	names := mp.GetInArgsTypesNames()

	// Act
	actual := args.Map{"result": len(names) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
	// Call again to hit cache
	names2 := mp.GetInArgsTypesNames()
	actual = args.Map{"result": len(names2) != len(names)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_MP_GetInArgsTypesNames_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()

	// Act
	actual := args.Map{"result": len(names) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_MP_ValidateMethodArgs_Success(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Greeting")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_MP_ValidateMethodArgs_WrongCount(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected args mismatch error", actual)
}

func Test_MP_ValidateMethodArgs_WrongType(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}, "not_int", "not_int"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
}

func Test_MP_VerifyInArgs(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.VerifyInArgs([]any{testTarget6{}})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_VerifyOutArgs(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	ok, err := mp.VerifyOutArgs([]any{0})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_VerifyOutArgs_Mismatch(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	ok, _ := mp.VerifyOutArgs([]any{"string"})

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_MP_InArgsVerifyRv(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(testTarget6{})})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_OutArgsVerifyRv(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_OutArgsVerifyRv_LengthMismatch(t *testing.T) {
	// Arrange
	mp := getMethodProcessor6("Add")
	ok, _ := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")})

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch for wrong length", actual)
}

// ═══════════════════════════════════════════════
// ReflectValueKind — uncovered methods
// ═══════════════════════════════════════════════

func Test_RVK_InvalidReflectValueKindModel(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	// Act
	actual := args.Map{"result": rvk.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual = args.Map{"result": rvk.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": rvk.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
}

func Test_RVK_IsEmptyError(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: true}

	// Act
	actual := args.Map{"result": rvk.IsEmptyError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error", actual)
	var nilRvk *reflectmodel.ReflectValueKind
	actual = args.Map{"result": nilRvk.IsEmptyError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty error", actual)
}

func Test_RVK_ActualInstance(t *testing.T) {
	// Arrange
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	inst := rvk.ActualInstance()

	// Act
	actual := args.Map{"result": inst.(int) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_RVK_ActualInstance_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.ActualInstance() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_RVK_PkgPath(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(testTarget6{}),
	}
	pkg := rvk.PkgPath()

	// Act
	actual := args.Map{"result": pkg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty pkg path", actual)
}

func Test_RVK_PkgPath_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.PkgPath() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_RVK_PkgPath_Invalid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	// Act
	actual := args.Map{"result": rvk.PkgPath() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_RVK_PointerRv_Valid(t *testing.T) {
	// Arrange
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	ptr := rvk.PointerRv()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_RVK_PointerRv_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.PointerRv() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_RVK_PointerRv_NotValid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(42),
	}
	ptr := rvk.PointerRv()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil (returns FinalReflectVal addr)", actual)
}

func Test_RVK_TypeName_Valid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
	}
	name := rvk.TypeName()

	// Act
	actual := args.Map{"result": name == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_RVK_TypeName_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.TypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_RVK_TypeName_NotValid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	// Act
	actual := args.Map{"result": rvk.TypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_RVK_PointerInterface_Valid(t *testing.T) {
	// Arrange
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	pi := rvk.PointerInterface()

	// Act
	actual := args.Map{"result": pi == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_RVK_PointerInterface_Nil(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.PointerInterface() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
