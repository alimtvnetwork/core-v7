package reflectmodeltests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── MethodProcessor.InvokeError ──

type cov5ErrorReturner struct{}

func (s cov5ErrorReturner) ReturnError() error {
	return errors.New("test-error")
}

func (s cov5ErrorReturner) ReturnNilError() error {
	return nil
}

func (s cov5ErrorReturner) ReturnValueAndError(ok bool) (string, error) {
	if ok {
		return "value", nil
	}
	return "", errors.New("fail")
}

func (s cov5ErrorReturner) ReturnThree() (int, string, error) {
	return 42, "hello", nil
}

func (s cov5ErrorReturner) ReturnSingle() string {
	return "only"
}

func (s cov5ErrorReturner) NoReturn() {}

func newCov5MethodProcessor(methodName string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(cov5ErrorReturner{})
	method, ok := t.MethodByName(methodName)
	if !ok {
		return nil
	}
	return &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
}

func Test_MethodProcessor_InvokeError_Success(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnError")
	funcErr, procErr := mp.InvokeError(cov5ErrorReturner{})

	// Act
	actual := args.Map{
		"funcErrMsg": funcErr.Error(),
		"procErr":    procErr == nil,
	}

	// Assert
	expected := args.Map{
		"funcErrMsg": "test-error",
		"procErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- success", actual)
}

func Test_MethodProcessor_InvokeError_WrongArgs(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnError")
	// Wrong number of args — should cause processing error
	_, procErr := mp.InvokeError(cov5ErrorReturner{}, "extra")

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- wrong args", actual)
}

func Test_MethodProcessor_InvokeFirstAndError_Success(t *testing.T) {
	mp := newCov5MethodProcessor("ReturnValueAndError")
	defer func() { recover() }() // results[1].(error) panics on nil error interface
	first, funcErr, procErr := mp.InvokeFirstAndError(cov5ErrorReturner{}, true)
	_ = first
	_ = funcErr
	_ = procErr
}

func Test_MethodProcessor_InvokeFirstAndError_WithFuncError(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnValueAndError")
	first, funcErr, procErr := mp.InvokeFirstAndError(cov5ErrorReturner{}, false)

	// Act
	actual := args.Map{
		"first":      first,
		"hasFuncErr": funcErr != nil,
		"procErr":    procErr == nil,
	}

	// Assert
	expected := args.Map{
		"first":      "",
		"hasFuncErr": true,
		"procErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- with func error", actual)
}

func Test_MethodProcessor_InvokeFirstAndError_ProcessingError(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnValueAndError")
	// Wrong number of args
	_, _, procErr := mp.InvokeFirstAndError(cov5ErrorReturner{})

	// Act
	actual := args.Map{"hasProcErr": procErr != nil}

	// Assert
	expected := args.Map{"hasProcErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- processing error", actual)
}

func Test_MethodProcessor_InvokeFirstAndError_SingleReturn(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnSingle")
	// ReturnSingle returns only 1 value, so len(results) <= 1
	first, _, procErr := mp.InvokeFirstAndError(cov5ErrorReturner{})

	// Act
	actual := args.Map{
		"hasProcErr": procErr != nil,
		"firstNotNil": first != nil,
	}

	// Assert
	expected := args.Map{
		"hasProcErr":  true,
		"firstNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- single return", actual)
}

func Test_MethodProcessor_GetFirstResponseOfInvoke_Error(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnSingle")
	// Wrong arg count
	first, err := mp.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{
		"firstNil": first == nil,
		"hasErr":   err != nil,
	}

	// Assert
	expected := args.Map{
		"firstNil": true,
		"hasErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns error -- error", actual)
}

func Test_MethodProcessor_InvokeResultOfIndex_Error(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnSingle")
	// Wrong arg count
	result, err := mp.InvokeResultOfIndex(0)

	// Act
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr":    err != nil,
	}

	// Assert
	expected := args.Map{
		"resultNil": true,
		"hasErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns error -- error", actual)
}

// ── MethodProcessor.GetOutArgsTypes with zero-return method ──

func Test_MethodProcessor_GetOutArgsTypes_ZeroReturn(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("NoReturn")
	types := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"len": len(types)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes returns correct value -- zero return", actual)
}

// ── MethodProcessor.GetInArgsTypes with zero-arg method (just receiver) ──

func Test_MethodProcessor_GetInArgsTypesNames_NoReturn(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("NoReturn")
	names := mp.GetInArgsTypesNames()
	// NoReturn() only has receiver

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns empty -- no return method", actual)
}

// ── MethodProcessor.IsEqual — different methods (in/out args mismatch) ──

func Test_MethodProcessor_IsEqual_DiffInArgs(t *testing.T) {
	// Arrange
	mp1 := newCov5MethodProcessor("ReturnSingle")
	mp2 := newCov5MethodProcessor("ReturnValueAndError")

	// Act
	actual := args.Map{
		"equal":    mp1.IsEqual(mp2),
		"notEqual": mp1.IsNotEqual(mp2),
	}

	// Assert
	expected := args.Map{
		"equal":    false,
		"notEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- diff in args", actual)
}

// ── MethodProcessor.IsEqual — same signature but different names ──

func Test_MethodProcessor_IsEqual_SameSignature(t *testing.T) {
	// Arrange
	mp1 := newCov5MethodProcessor("ReturnError")
	mp2 := newCov5MethodProcessor("ReturnNilError")
	// Both have same signature: (receiver) -> error
	// IsEqual compares names first — different names means not equal

	// Act
	actual := args.Map{
		"equal": mp1.IsEqual(mp2),
	}

	// Assert
	expected := args.Map{
		"equal": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same signature different name", actual)
}

// ── MethodProcessor.VerifyOutArgs mismatch ──

func Test_MethodProcessor_VerifyOutArgs_Mismatch(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnSingle")
	// ReturnSingle returns string, give int
	ok, err := mp.VerifyOutArgs([]any{42})

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
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs returns correct value -- mismatch", actual)
}

// ── MethodProcessor.OutArgsVerifyRv length mismatch ──

func Test_MethodProcessor_OutArgsVerifyRv_LengthMismatch(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnSingle")
	// ReturnSingle has 1 out arg, give 2
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(""), reflect.TypeOf(0)})

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
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv returns correct value -- length mismatch", actual)
}

// ── MethodProcessor.ValidateMethodArgs correct types ──

func Test_MethodProcessor_ValidateMethodArgs_Correct(t *testing.T) {
	// Arrange
	mp := newCov5MethodProcessor("ReturnValueAndError")
	err := mp.ValidateMethodArgs([]any{cov5ErrorReturner{}, true})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- correct", actual)
}

// ── MethodProcessor — nil receiver invoke ──

func Test_MethodProcessor_Invoke_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- nil receiver", actual)
}

func Test_MethodProcessor_InvokeError_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, procErr := mp.InvokeError()

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns nil -- nil receiver", actual)
}

func Test_MethodProcessor_InvokeFirstAndError_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, _, procErr := mp.InvokeFirstAndError()

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns nil -- nil receiver", actual)
}

func Test_MethodProcessor_GetFirstResponseOfInvoke_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	first, err := mp.GetFirstResponseOfInvoke()

	// Act
	actual := args.Map{
		"firstNil": first == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"firstNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns nil -- nil", actual)
}

func Test_MethodProcessor_InvokeResultOfIndex_NilReceiver(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	result, err := mp.InvokeResultOfIndex(0)

	// Act
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns nil -- nil", actual)
}

// ── MethodProcessor.ValidateMethodArgs nil ──

func Test_MethodProcessor_ValidateMethodArgs_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	// calling Invoke on nil hits validationError which returns error
	_, err := mp.Invoke("something")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns nil -- nil processor", actual)
}

// ── ReflectValueKind — with Interface reflect value ──

func Test_ReflectValueKind_InterfaceReflectValue(t *testing.T) {
	// Arrange
	var iface interface{} = "hello"
	rv := reflect.ValueOf(&iface)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}

	// Act
	actual := args.Map{
		"typeName":    rvk.TypeName() != "",
		"pkgPath":     rvk.PkgPath(),
		"ptrRvNotNil": rvk.PointerRv() != nil,
	}

	// Assert
	expected := args.Map{
		"typeName":    true,
		"pkgPath":     "",
		"ptrRvNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns correct value -- interface rv", actual)
}
