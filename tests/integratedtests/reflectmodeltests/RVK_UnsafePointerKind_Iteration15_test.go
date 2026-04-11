package reflectmodeltests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.IsNull — UnsafePointer kind via internal test exercises
// These test ReflectValueKind with various pointer/interface kinds
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_UnsafePointerKind(t *testing.T) {
	// Arrange
	x := 42
	up := unsafe.Pointer(&x)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(up),
		Kind:            reflect.UnsafePointer,
	}

	// Act
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
		"typeName":  rvk.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"isInvalid": false,
		"actNotNil": true,
		"typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- UnsafePointer kind", actual)
}

func Test_RVK_FuncKind(t *testing.T) {
	// Arrange
	fn := func() string { return "hello" }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(fn),
		Kind:            reflect.Func,
	}

	// Act
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
		"typeName":  rvk.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"isInvalid": false,
		"actNotNil": true,
		"typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Func kind", actual)
}

func Test_RVK_ChanKind(t *testing.T) {
	// Arrange
	ch := make(chan string, 1)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(ch),
		Kind:            reflect.Chan,
	}

	// Act
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
	}

	// Assert
	expected := args.Map{
		"isInvalid": false,
		"actNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Chan kind", actual)
}

func Test_RVK_Float64Kind(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(3.14),
		Kind:            reflect.Float64,
	}

	// Act
	actual := args.Map{
		"actInst":  rvk.ActualInstance(),
		"typeName": rvk.TypeName() != "",
		"pkgPath":  rvk.PkgPath(),
	}

	// Assert
	expected := args.Map{
		"actInst": 3.14,
		"typeName": true,
		"pkgPath": "",
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Float64 kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — PointerRv and PointerInterface with valid struct
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_Valid_PointerRv_Struct(t *testing.T) {
	// Arrange
	type testStruct struct{ X int }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(testStruct{X: 99}),
		Kind:            reflect.Struct,
	}
	ptr := rvk.PointerRv()
	ptrIface := rvk.PointerInterface()

	// Act
	actual := args.Map{
		"ptrNotNil": ptr != nil,
		"ptrIfaceNotNil": ptrIface != nil,
	}

	// Assert
	expected := args.Map{
		"ptrNotNil": true,
		"ptrIfaceNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- PointerRv struct", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — HasError with actual error set
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_HasError_True(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("some error message")

	// Act
	actual := args.Map{
		"hasErr":   rvk.HasError(),
		"emptyErr": rvk.IsEmptyError(),
		"errMsg":   rvk.Error.Error(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"emptyErr": false,
		"errMsg": "some error message",
	}
	expected.ShouldBeEqual(t, 0, "RVK returns error -- HasError with message", actual)
}

func Test_RVK_HasError_False(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
		Error:           nil,
	}

	// Act
	actual := args.Map{
		"hasErr": rvk.HasError(),
		"emptyErr": rvk.IsEmptyError(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"emptyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- no error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InvokeFirstAndError with nil error interface (panic recovery)
// ══════════════════════════════════════════════════════════════════════════════

type helperI15 struct{}

func (h helperI15) ReturnTwoNoError() (string, error) {
	return "ok", nil
}

func (h helperI15) ReturnStringOnly() string {
	return "hello"
}

func (h helperI15) ReturnBoolInt(x int) (bool, int) {
	return x > 0, x * 2
}

func (h helperI15) ReturnThree() (int, string, bool) {
	return 1, "two", true
}

func getI15MP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(helperI15{})
	m, ok := t.MethodByName(name)
	if !ok {
		return nil
	}
	return &reflectmodel.MethodProcessor{
		Name:          m.Name,
		Index:         m.Index,
		ReflectMethod: m,
	}
}

func Test_InvokeFirstAndError_NilErrorInterface(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnTwoNoError")
	// InvokeFirstAndError does results[1].Interface().(error) on a nil interface → panics
	didPanic := false
	func() {
		defer func() {
			recover()
			didPanic = true
		}()
		_, _, _ = mp.InvokeFirstAndError(helperI15{})
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError panics -- nil error interface conversion", actual)
}

func Test_InvokeFirstAndError_NonErrorSecondReturn_Panics(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnBoolInt")
	panicked := false
	_, _, processingErr := func() (any, error, error) {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()

		return mp.InvokeFirstAndError(helperI15{}, 5)
	}()

	// Act
	actual := args.Map{
		"panicked": panicked,
		"hasProcessingErr": processingErr != nil,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"hasProcessingErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError panics -- non-error second return", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — Invoke success with zero-return method
// ══════════════════════════════════════════════════════════════════════════════

func Test_Invoke_ZeroReturn(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")
	results, err := mp.Invoke(sampleStruct{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(results),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- single return", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — VerifyInArgs with wrong arg count
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyInArgs_WrongCount(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnBoolInt")
	ok, err := mp.VerifyInArgs([]any{helperI15{}}) // needs receiver + int

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
	expected.ShouldBeEqual(t, 0, "VerifyInArgs returns error -- wrong count", actual)
}

func Test_VerifyOutArgs_WrongCount(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnBoolInt")
	ok, err := mp.VerifyOutArgs([]any{true}) // needs 2 out

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
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs returns error -- wrong count", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — argsCountMismatchErrorMessage exercises with 0 args
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidateMethodArgs_ZeroGiven(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnBoolInt")
	err := mp.ValidateMethodArgs([]any{})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns error -- zero args given", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — with various types
// ══════════════════════════════════════════════════════════════════════════════

func Test_FieldProcessor_StringField(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Name", 0)

	// Act
	actual := args.Map{
		"isStr":  fp.IsFieldKind(reflect.String),
		"isInt":  fp.IsFieldKind(reflect.Int),
		"typeOk": fp.IsFieldType(reflect.TypeOf("")),
	}

	// Assert
	expected := args.Map{
		"isStr": true,
		"isInt": false,
		"typeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- String field", actual)
}

func Test_FieldProcessor_NilReceiver_IsFieldKind(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{"kind": fp.IsFieldKind(reflect.Int)}

	// Assert
	expected := args.Map{"kind": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns false -- nil IsFieldKind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValue — nil RawData
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectValue_NilRawData(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{
		TypeName:     "NilType",
		FieldsNames:  nil,
		MethodsNames: nil,
		RawData:      nil,
	}

	// Act
	actual := args.Map{
		"typeName":   rv.TypeName,
		"fieldsNil":  rv.FieldsNames == nil,
		"methodsNil": rv.MethodsNames == nil,
		"rawNil":     rv.RawData == nil,
	}

	// Assert
	expected := args.Map{
		"typeName": "NilType",
		"fieldsNil": true,
		"methodsNil": true,
		"rawNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns nil -- nil raw data", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — GetOutArgsTypes with 3 returns
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetOutArgsTypes_ThreeReturns(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnThree")
	out := mp.GetOutArgsTypes()
	out2 := mp.GetOutArgsTypes() // cached

	// Act
	actual := args.Map{
		"len": len(out),
		"cached": len(out2),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"cached": 3,
	}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes returns correct value -- three returns", actual)
}

func Test_GetInArgsTypesNames_ThreeReturns(t *testing.T) {
	// Arrange
	mp := getI15MP("ReturnThree")
	names := mp.GetInArgsTypesNames()
	names2 := mp.GetInArgsTypesNames() // cached

	// Act
	actual := args.Map{
		"len": len(names),
		"cached": len(names2),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"cached": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns correct value -- receiver only", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — IsEqual with different out args
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsEqual_DiffOutArgs(t *testing.T) {
	// Arrange
	mp1 := getI15MP("ReturnStringOnly")
	mp2 := getI15MP("ReturnThree")

	// Act
	actual := args.Map{"eq": mp1.IsEqual(mp2)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- diff out args", actual)
}

func Test_IsEqual_SameSig(t *testing.T) {
	// Arrange
	mp1 := getI15MP("ReturnStringOnly")
	mp2 := getI15MP("ReturnStringOnly")

	// Act
	actual := args.Map{"eq": mp1.IsEqual(mp2)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- same sig", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InvokeError with nil processing error
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeError_NilFunc(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, procErr := mp.InvokeError()

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- nil receiver", actual)
}

func Test_GetFirstResponseOfInvoke_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, err := mp.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns error -- nil receiver", actual)
}

func Test_InvokeResultOfIndex_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, err := mp.InvokeResultOfIndex(0)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns error -- nil receiver", actual)
}

func Test_InvokeFirstAndError_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, _, procErr := mp.InvokeFirstAndError()

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- nil receiver", actual)
}
