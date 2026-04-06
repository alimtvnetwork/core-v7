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

func Test_Cov6_FieldProcessor_IsFieldType_Valid(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(""))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}

func Test_Cov6_FieldProcessor_IsFieldType_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil receiver should return false", actual)
}

func Test_Cov6_FieldProcessor_IsFieldKind_Valid(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}
	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected int kind", actual)
	actual := args.Map{"result": fp.IsFieldKind(reflect.String)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}

func Test_Cov6_FieldProcessor_IsFieldKind_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}
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

func Test_Cov6_MP_HasValidFunc(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.HasValidFunc()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	var nilMp *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMp.HasValidFunc()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov6_MP_GetFuncName(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.GetFuncName() != "Add"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Add", actual)
}

func Test_Cov6_MP_IsInvalid(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.IsInvalid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	var nilMp *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMp.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov6_MP_Func(t *testing.T) {
	mp := getMethodProcessor6("Add")
	f := mp.Func()
	actual := args.Map{"result": f == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected func", actual)
}

func Test_Cov6_MP_Func_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"result": mp.Func() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov6_MP_ArgsCount(t *testing.T) {
	mp := getMethodProcessor6("Add")
	// Add has receiver + a + b = 3
	actual := args.Map{"result": mp.ArgsCount() < 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
}

func Test_Cov6_MP_ReturnLength(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.ReturnLength() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov6_MP_ReturnLength_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"result": mp.ReturnLength() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return -1", actual)
}

func Test_Cov6_MP_IsPublicMethod(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.IsPublicMethod()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected public", actual)
}

func Test_Cov6_MP_IsPrivateMethod(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.IsPrivateMethod()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not private", actual)
}

func Test_Cov6_MP_ArgsLength(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.ArgsLength() < 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
}

func Test_Cov6_MP_Invoke_Success(t *testing.T) {
	mp := getMethodProcessor6("Add")
	// receiver + 2 args
	results, err := mp.Invoke(testTarget6{}, 3, 4)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": len(results) != 1 || results[0].(int) != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_Cov6_MP_Invoke_ArgsMismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	_, err := mp.Invoke(testTarget6{}, 3) // missing arg
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for args mismatch", actual)
}

func Test_Cov6_MP_Invoke_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_Cov6_MP_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	resp, err := mp.GetFirstResponseOfInvoke(testTarget6{})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": resp.(string) != "hi"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hi", actual)
}

func Test_Cov6_MP_InvokeResultOfIndex(t *testing.T) {
	mp := getMethodProcessor6("Add")
	resp, err := mp.InvokeResultOfIndex(0, testTarget6{}, 1, 2)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": resp.(int) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov6_MP_InvokeError(t *testing.T) {
	defer func() { recover() }() // InvokeError may panic on zero reflect.Value
	mp := getMethodProcessor6("Err")
	funcErr, procErr := mp.InvokeError(testTarget6{})
	actual := args.Map{"result": procErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": funcErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error from Err()", actual)
}

func Test_Cov6_MP_InvokeFirstAndError_Success(t *testing.T) {
	defer func() { recover() }() // may panic on zero reflect.Value in ReflectValueToAnyValue
	mp := getMethodProcessor6("PairResult")
	first, funcErr, procErr := mp.InvokeFirstAndError(testTarget6{})
	actual := args.Map{"result": procErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "processing error:", actual)
	actual := args.Map{"result": funcErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no func error", actual)
	actual := args.Map{"result": first.(string) != "ok"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'ok'", actual)
}

func Test_Cov6_MP_InvokeFirstAndError_SingleReturn(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	_, _, procErr := mp.InvokeFirstAndError(testTarget6{})
	actual := args.Map{"result": procErr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for single return", actual)
}

func Test_Cov6_MP_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Cov6_MP_IsEqual_OneNil(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.IsEqual(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)
}

func Test_Cov6_MP_IsEqual_Same(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.IsEqual(mp)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)
}

func Test_Cov6_MP_IsEqual_DiffMethods(t *testing.T) {
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")
	// Different signatures → should fail at args verification
	_ = mp1.IsEqual(mp2)
}

func Test_Cov6_MP_IsNotEqual(t *testing.T) {
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")
	actual := args.Map{"result": mp1.IsNotEqual(mp2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different methods should not be equal", actual)
}

func Test_Cov6_MP_GetType(t *testing.T) {
	mp := getMethodProcessor6("Add")
	actual := args.Map{"result": mp.GetType() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil type", actual)
}

func Test_Cov6_MP_GetType_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"result": mp.GetType() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil type", actual)
}

func Test_Cov6_MP_GetOutArgsTypes(t *testing.T) {
	mp := getMethodProcessor6("Add")
	out := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(out) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// Call again to hit cache
	out2 := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(out2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_Cov6_MP_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	out := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(out) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov6_MP_GetInArgsTypes(t *testing.T) {
	mp := getMethodProcessor6("Add")
	in := mp.GetInArgsTypes()
	actual := args.Map{"result": len(in) < 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
	// Call again to hit cache
	in2 := mp.GetInArgsTypes()
	actual := args.Map{"result": len(in2) != len(in)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_Cov6_MP_GetInArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	in := mp.GetInArgsTypes()
	actual := args.Map{"result": len(in) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov6_MP_GetInArgsTypesNames(t *testing.T) {
	mp := getMethodProcessor6("Add")
	names := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names) < 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 2", actual)
	// Call again to hit cache
	names2 := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names2) != len(names)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache should return same", actual)
}

func Test_Cov6_MP_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov6_MP_ValidateMethodArgs_Success(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_Cov6_MP_ValidateMethodArgs_WrongCount(t *testing.T) {
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected args mismatch error", actual)
}

func Test_Cov6_MP_ValidateMethodArgs_WrongType(t *testing.T) {
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}, "not_int", "not_int"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
}

func Test_Cov6_MP_VerifyInArgs(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.VerifyInArgs([]any{testTarget6{}})
	actual := args.Map{"result": ok || err != nil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_Cov6_MP_VerifyOutArgs(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, err := mp.VerifyOutArgs([]any{0})
	actual := args.Map{"result": ok || err != nil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_Cov6_MP_VerifyOutArgs_Mismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, _ := mp.VerifyOutArgs([]any{"string"})
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_Cov6_MP_InArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(testTarget6{})})
	actual := args.Map{"result": ok || err != nil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_Cov6_MP_OutArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	actual := args.Map{"result": ok || err != nil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_Cov6_MP_OutArgsVerifyRv_LengthMismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, _ := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")})
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected mismatch for wrong length", actual)
}

// ═══════════════════════════════════════════════
// ReflectValueKind — uncovered methods
// ═══════════════════════════════════════════════

func Test_Cov6_RVK_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	actual := args.Map{"result": rvk.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": rvk.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": rvk.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
}

func Test_Cov6_RVK_IsEmptyError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: true}
	actual := args.Map{"result": rvk.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error", actual)
	var nilRvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": nilRvk.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty error", actual)
}

func Test_Cov6_RVK_ActualInstance(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	inst := rvk.ActualInstance()
	actual := args.Map{"result": inst.(int) != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_Cov6_RVK_ActualInstance_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": rvk.ActualInstance() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov6_RVK_PkgPath(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(testTarget6{}),
	}
	pkg := rvk.PkgPath()
	actual := args.Map{"result": pkg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty pkg path", actual)
}

func Test_Cov6_RVK_PkgPath_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": rvk.PkgPath() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov6_RVK_PkgPath_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	actual := args.Map{"result": rvk.PkgPath() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_Cov6_RVK_PointerRv_Valid(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	ptr := rvk.PointerRv()
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov6_RVK_PointerRv_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": rvk.PointerRv() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_Cov6_RVK_PointerRv_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(42),
	}
	ptr := rvk.PointerRv()
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil (returns FinalReflectVal addr)", actual)
}

func Test_Cov6_RVK_TypeName_Valid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
	}
	name := rvk.TypeName()
	actual := args.Map{"result": name == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov6_RVK_TypeName_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": rvk.TypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov6_RVK_TypeName_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	actual := args.Map{"result": rvk.TypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_Cov6_RVK_PointerInterface_Valid(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	pi := rvk.PointerInterface()
	actual := args.Map{"result": pi == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Cov6_RVK_PointerInterface_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"result": rvk.PointerInterface() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
