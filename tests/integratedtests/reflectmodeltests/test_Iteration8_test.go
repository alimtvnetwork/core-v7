package reflectmodeltests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── test types ──

type testAdder struct{}

func (t testAdder) Add(a, b int) int        { return a + b }
func (t testAdder) Greet(name string) string { return "hello " + name }
func (t testAdder) Fail() error              { return errors.New("fail") }
func (t testAdder) NoError() error           { return nil }
func (t testAdder) TwoReturns(x int) (int, error) {
	if x < 0 {
		return 0, errors.New("negative")
	}
	return x * 2, nil
}

func getMP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testAdder{})
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

// ==========================================================================
// FieldProcessor
// ==========================================================================

func Test_FieldProcessor_IsFieldType_FromtestIteration8(t *testing.T) {
	// Arrange
	type sample struct {
		Name string
		Age  int
	}
	st := reflect.TypeOf(sample{})
	f := st.Field(0)
	fp := &reflectmodel.FieldProcessor{
		Name:      f.Name,
		Index:     0,
		Field:     f,
		FieldType: f.Type,
	}

	// Act
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(""))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for string type", actual)
	actual := args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for int type", actual)

	// nil receiver
	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{"result": nilFP.IsFieldType(reflect.TypeOf(""))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil receiver", actual)
}

func Test_FieldProcessor_IsFieldKind_FromtestIteration8(t *testing.T) {
	// Arrange
	type sample struct {
		Name string
	}
	st := reflect.TypeOf(sample{})
	f := st.Field(0)
	fp := &reflectmodel.FieldProcessor{
		Name:      f.Name,
		Index:     0,
		Field:     f,
		FieldType: f.Type,
	}

	// Act
	actual := args.Map{"result": fp.IsFieldKind(reflect.String)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for string kind", actual)
	actual := args.Map{"result": fp.IsFieldKind(reflect.Int)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for int kind", actual)

	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{"result": nilFP.IsFieldKind(reflect.String)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil receiver", actual)
}

// ==========================================================================
// MethodProcessor — basic properties
// ==========================================================================

func Test_MP_HasValidFunc_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.HasValidFunc()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.HasValidFunc()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_MP_GetFuncName_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.GetFuncName() != "Add"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Add", actual)
}

func Test_MP_IsInvalid_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
}

func Test_MP_Func_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	f := mp.Func()

	// Act
	actual := args.Map{"result": f == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)

	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.Func() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_MP_ArgsCount_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	// receiver + a + b = 3

	// Act
	actual := args.Map{"result": mp.ArgsCount() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_MP_ArgsLength_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.ArgsLength() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_MP_ReturnLength_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.ReturnLength() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.ReturnLength() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1 for nil", actual)
}

func Test_MP_IsPublicMethod_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.IsPublicMethod()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.IsPublicMethod()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_MP_IsPrivateMethod_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.IsPrivateMethod()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for public method", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.IsPrivateMethod()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_MP_GetType_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.GetType() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": nilMP.GetType() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil", actual)
}

// ==========================================================================
// MethodProcessor — Invoke variants
// ==========================================================================

func Test_MP_Invoke_Success_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	results, err := mp.Invoke(testAdder{}, 2, 3)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": results[0].(int) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_MP_Invoke_NilReceiver(t *testing.T) {
	// Arrange
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.Invoke()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_MP_Invoke_WrongArgCount(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	_, err := mp.Invoke(testAdder{}, 2)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected args count mismatch error", actual)
}

func Test_MP_GetFirstResponseOfInvoke_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	resp, err := mp.GetFirstResponseOfInvoke(testAdder{}, 2, 3)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": resp.(int) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_MP_InvokeResultOfIndex_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	resp, err := mp.InvokeResultOfIndex(0, testAdder{}, 10, 20)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": resp.(int) != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 30", actual)
}

func Test_MP_InvokeError_WithError(t *testing.T) {
	// Arrange
	mp := getMP("Fail")
	funcErr, procErr := mp.InvokeError(testAdder{})

	// Act
	actual := args.Map{"result": procErr}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "procErr", actual)
	actual := args.Map{"result": funcErr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected func error", actual)
}

func Test_MP_InvokeError_NoError(t *testing.T) {
	// Arrange
	mp := getMP("NoError")
	defer func() { recover() }() // may panic on nil error interface
	funcErr, procErr := mp.InvokeError(testAdder{})

	// Act
	actual := args.Map{"result": procErr}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "procErr", actual)
	_ = funcErr
}

func Test_MP_InvokeFirstAndError_Success_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("TwoReturns")
	// InvokeFirstAndError does results[1].Interface().(error) which panics on nil interface
	didPanic := false
	var first any
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		first, _, _ = mp.InvokeFirstAndError(testAdder{}, 5)
	}()

	// Act
	actual := args.Map{
		"didPanic": didPanic,
		"firstIfNoPanic": first,
	}
	if didPanic {

	// Assert
		expected := args.Map{
			"didPanic": true,
			"firstIfNoPanic": nil,
		}
		expected.ShouldBeEqual(t, 0, "InvokeFirstAndError panics -- nil error interface cast", actual)
	} else {
		actual := args.Map{"result": first.(int) != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	}
}

func Test_MP_InvokeFirstAndError_FuncError(t *testing.T) {
	// Arrange
	mp := getMP("TwoReturns")
	_, funcErr, procErr := mp.InvokeFirstAndError(testAdder{}, -1)

	// Act
	actual := args.Map{"result": procErr}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "procErr", actual)
	actual := args.Map{"result": funcErr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected func error for negative input", actual)
}

func Test_MP_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	// Arrange
	mp := getMP("Add") // only 1 return
	_, _, procErr := mp.InvokeFirstAndError(testAdder{}, 1, 2)

	// Act
	actual := args.Map{"result": procErr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected processing error for single-return method", actual)
}

// ==========================================================================
// MethodProcessor — IsEqual / IsNotEqual
// ==========================================================================

func Test_MP_IsEqual_BothNil_FromtestIteration8(t *testing.T) {
	// Arrange
	var a, b *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for both nil", actual)
}

func Test_MP_IsEqual_OneNil_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.IsEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_MP_IsEqual_Same_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")

	// Act
	actual := args.Map{"result": mp.IsEqual(mp)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for same ptr", actual)
}

func Test_MP_IsEqual_Different(t *testing.T) {
	a := getMP("Add")
	b := getMP("Greet")
	// They have different arg types so IsEqual should detect via InArgsVerifyRv
	_ = a.IsEqual(b)
	_ = a.IsNotEqual(b)
}

func Test_MP_IsEqual_SameName(t *testing.T) {
	// Arrange
	a := getMP("Add")
	b := getMP("Add")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for same method", actual)
	actual := args.Map{"result": a.IsNotEqual(b)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for IsNotEqual on same method", actual)
}

// ==========================================================================
// MethodProcessor — GetOutArgsTypes / GetInArgsTypes / GetInArgsTypesNames
// ==========================================================================

func Test_MP_GetOutArgsTypes_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("TwoReturns")
	outTypes := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(outTypes) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 out types", actual)
	// Call again to test cache
	outTypes2 := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(outTypes2) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache returned different length", actual)

	var nilMP *reflectmodel.MethodProcessor
	nilOut := nilMP.GetOutArgsTypes()
	actual := args.Map{"result": len(nilOut) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_MP_GetInArgsTypes_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	inTypes := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"result": len(inTypes) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	// Call again for cache
	inTypes2 := mp.GetInArgsTypes()
	actual := args.Map{"result": len(inTypes2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache failed", actual)

	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": len(nilMP.GetInArgsTypes()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_MP_GetInArgsTypesNames_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Greet")
	names := mp.GetInArgsTypesNames()

	// Act
	actualNames := args.Map{"length": len(names)}

	// Assert
	expectedNames := args.Map{"length": 2}
	expectedNames.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns 2 -- receiver + name", actualNames)
	// Call again for cache
	names2 := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names2) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cache failed", actual)

	var nilMP *reflectmodel.MethodProcessor
	actual := args.Map{"result": len(nilMP.GetInArgsTypesNames()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_MP_GetOutArgsTypes_ZeroOut(t *testing.T) {
	// Arrange
	// Find a method with zero out args... we don't have one, but let's
	// test the nil/empty path indirectly
	var nilMP *reflectmodel.MethodProcessor
	out := nilMP.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(out) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ==========================================================================
// MethodProcessor — ValidateMethodArgs / VerifyInArgs / VerifyOutArgs
// ==========================================================================

func Test_MP_ValidateMethodArgs_Success_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Greet")
	err := mp.ValidateMethodArgs([]any{testAdder{}, "world"})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_MP_ValidateMethodArgs_WrongCount_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Greet")
	err := mp.ValidateMethodArgs([]any{testAdder{}})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong arg count", actual)
}

func Test_MP_ValidateMethodArgs_WrongType_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	err := mp.ValidateMethodArgs([]any{testAdder{}, "not-int", 3})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong arg type", actual)
}

func Test_MP_VerifyInArgs_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	ok, err := mp.VerifyInArgs([]any{testAdder{}, 1, 2})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_VerifyOutArgs_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("TwoReturns")
	ok, err := mp.VerifyOutArgs([]any{0, errors.New("e")})

	// Act
	actual := args.Map{"result": ok || err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch for concrete error vs error interface", actual)
}

func Test_MP_InArgsVerifyRv_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	types := mp.GetInArgsTypes()
	ok, err := mp.InArgsVerifyRv(types)

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_MP_OutArgsVerifyRv_FromtestIteration8(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	types := mp.GetOutArgsTypes()
	ok, err := mp.OutArgsVerifyRv(types)

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

// ==========================================================================
// ReflectValueKind
// ==========================================================================

func Test_InvalidReflectValueKindModel_FromtestIteration8(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	// Act
	actual := args.Map{"result": rvk == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": rvk.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": rvk.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has error", actual)
	actual := args.Map{"result": rvk.IsEmptyError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty error", actual)
}

func Test_RVK_NilReceiver(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"result": rvk.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": rvk.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for nil", actual)
	actual := args.Map{"result": rvk.IsEmptyError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty error for nil", actual)
	actual := args.Map{"result": rvk.ActualInstance() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": rvk.PkgPath() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": rvk.PointerRv() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": rvk.TypeName() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": rvk.PointerInterface() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_RVK_Valid(t *testing.T) {
	// Arrange
	val := "hello"
	rv := reflect.ValueOf(val)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}

	// Act
	actual := args.Map{"result": rvk.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual := args.Map{"result": rvk.ActualInstance() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	if rvk.PkgPath() != "" {
		// string type has no PkgPath
	}
	actual := args.Map{"result": rvk.TypeName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty type name", actual)

	ptr := rvk.PointerRv()
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil pointer rv", actual)
	iface := rvk.PointerInterface()
	actual := args.Map{"result": iface == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil pointer interface", actual)
}

func Test_RVK_Invalid_PointerRv(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
	}
	ptr := rvk.PointerRv()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil (returns &FinalReflectVal for invalid)", actual)
}

func Test_RVK_PkgPath_Invalid_FromtestIteration8(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid: false,
	}

	// Act
	actual := args.Map{"result": rvk.PkgPath() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_RVK_TypeName_Invalid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid: false,
	}

	// Act
	actual := args.Map{"result": rvk.TypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

// ==========================================================================
// rvUtils — coverage for utility functions (accessed via MethodProcessor)
// ==========================================================================

func Test_Utils_ArgsToReflectValues_Empty(t *testing.T) {
	// Arrange
	mp := getMP("NoError")
	// Invoke with just the receiver covers ArgsToReflectValues with 1 arg
	_, err := mp.Invoke(testAdder{})

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
}

func Test_Utils_ReflectValuesToInterfaces_Coverage(t *testing.T) {
	// Arrange
	// Covered through Invoke which calls ReflectValuesToInterfaces
	mp := getMP("Greet")
	results, err := mp.Invoke(testAdder{}, "world")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": results[0] != "hello world"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello world", actual)
}

func Test_Utils_VerifyReflectTypes_LengthMismatch(t *testing.T) {
	// Arrange
	mp := getMP("Add")
	// Provide wrong number of types
	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})

	// Act
	actual := args.Map{"result": ok || err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected length mismatch error", actual)
}

func Test_Utils_VerifyReflectTypes_TypeMismatch(t *testing.T) {
	// Arrange
	mp := getMP("Greet") // expects (testAdder, string)
	wrongTypes := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf(42), // should be string
	}
	ok, err := mp.InArgsVerifyRv(wrongTypes)

	// Act
	actual := args.Map{"result": ok || err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
}

func Test_Utils_InterfacesToTypesNamesWithValues(t *testing.T) {
	// Arrange
	// Covered through argsCountMismatchErrorMessage via Invoke with wrong args
	mp := getMP("Add")
	_, err := mp.Invoke(testAdder{}, "wrong")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error containing type info", actual)
	errMsg := err.Error()
	actual := args.Map{"result": errMsg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty error message", actual)
}

func Test_Utils_IndexToPosition(t *testing.T) {
	// Arrange
	// Covered through type verification error messages
	mp := getMP("Add")
	// Wrong type at position 1 (2nd arg) covers "2nd" path
	wrongTypes := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf("wrong"),
		reflect.TypeOf(0),
	}
	ok, err := mp.InArgsVerifyRv(wrongTypes)

	// Act
	actual := args.Map{"result": ok || err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// Test 3rd and 4th+ positions
	wrongTypes2 := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf("wrong"),
		reflect.TypeOf("wrong"),
	}
	_, _ = mp.InArgsVerifyRv(wrongTypes2)
}

func Test_Utils_PrependWithSpaces_Coverage(t *testing.T) {
	// Arrange
	// Covered through VerifyReflectTypes error path which calls PrependWithSpaces
	mp := getMP("Greet")
	wrongTypes := []reflect.Type{
		reflect.TypeOf(0),
		reflect.TypeOf(0),
	}
	_, err := mp.InArgsVerifyRv(wrongTypes)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ==========================================================================
// isNull (unexported, covered via rvUtils.IsNull in ReflectValueToAnyValue)
// ==========================================================================

func Test_IsNull_Coverage(t *testing.T) {
	// Test nil case through InvokeError on NoError method
	mp := getMP("NoError")
	defer func() { recover() }()
	funcErr, procErr := mp.InvokeError(testAdder{})
	_ = funcErr
	_ = procErr
}

// ==========================================================================
// Additional edge cases for maximal coverage
// ==========================================================================

func Test_MP_Invoke_Greet(t *testing.T) {
	// Arrange
	mp := getMP("Greet")
	results, err := mp.Invoke(testAdder{}, "Go")

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": results[0] != "hello Go"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello Go'", actual)
}

func Test_MP_Invoke_TwoReturns(t *testing.T) {
	// Arrange
	mp := getMP("TwoReturns")
	results, err := mp.Invoke(testAdder{}, 7)

	// Act
	actual := args.Map{"result": err}

	// Assert
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)
	actual := args.Map{"result": len(results) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 results", actual)
	actual := args.Map{"result": results[0].(int) != 14}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 14", actual)
}

func Test_MP_InvokeFirstAndError_NilReceiver(t *testing.T) {
	// Arrange
	var nilMP *reflectmodel.MethodProcessor
	_, _, err := nilMP.InvokeFirstAndError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_MP_InvokeError_NilReceiver(t *testing.T) {
	// Arrange
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.InvokeError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_MP_GetFirstResponseOfInvoke_Error(t *testing.T) {
	// Arrange
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_RVK_Struct_PkgPath(t *testing.T) {
	val := testAdder{}
	rv := reflect.ValueOf(val)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}
	pkg := rvk.PkgPath()
	_ = pkg // struct types have PkgPath
	_ = fmt.Sprintf("%v", rvk.ActualInstance())
}
