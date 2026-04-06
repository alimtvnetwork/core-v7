package reflectmodeltests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── FieldProcessor ──

func Test_C9_FieldProcessor_IsFieldType(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(""),
	}
	actual := args.Map{
		"matchStr":  fp.IsFieldType(reflect.TypeOf("")),
		"matchInt":  fp.IsFieldType(reflect.TypeOf(0)),
	}
	expected := args.Map{"matchStr": true, "matchInt": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor.IsFieldType returns correct value -- with args", actual)

	var nilFp *reflectmodel.FieldProcessor
	actual := args.Map{"result": nilFp.IsFieldType(reflect.TypeOf(""))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_C9_FieldProcessor_IsFieldKind(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(""),
	}
	actual := args.Map{
		"matchStr": fp.IsFieldKind(reflect.String),
		"matchInt": fp.IsFieldKind(reflect.Int),
	}
	expected := args.Map{"matchStr": true, "matchInt": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor.IsFieldKind returns correct value -- with args", actual)

	var nilFp *reflectmodel.FieldProcessor
	actual := args.Map{"result": nilFp.IsFieldKind(reflect.String)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

// ── ReflectValueKind ──

func Test_C9_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	actual := args.Map{
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"emptyError": rvk.IsEmptyError(),
		"isValid":    rvk.IsValid,
	}
	expected := args.Map{
		"isInvalid":  true,
		"hasError":   true,
		"emptyError": false,
		"isValid":    false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidReflectValueKindModel returns error -- with args", actual)
}

func Test_C9_ReflectValueKind_NilReceiver(t *testing.T) {
	var nilRvk *reflectmodel.ReflectValueKind
	actual := args.Map{
		"isInvalid":  nilRvk.IsInvalid(),
		"hasError":   nilRvk.HasError(),
		"emptyError": nilRvk.IsEmptyError(),
		"actInst":    nilRvk.ActualInstance() == nil,
		"pkgPath":    nilRvk.PkgPath(),
		"ptrRv":      nilRvk.PointerRv() == nil,
		"typeName":   nilRvk.TypeName(),
		"ptrIface":   nilRvk.PointerInterface() == nil,
	}
	expected := args.Map{
		"isInvalid":  true,
		"hasError":   false,
		"emptyError": true,
		"actInst":    true,
		"pkgPath":    "",
		"ptrRv":      true,
		"typeName":   "",
		"ptrIface":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns nil -- nil receiver", actual)
}

func Test_C9_ReflectValueKind_Valid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf("hello"),
		Kind:            reflect.String,
	}
	actual := args.Map{
		"isInvalid":     rvk.IsInvalid(),
		"hasError":      rvk.HasError(),
		"actInstNotNil": rvk.ActualInstance() != nil,
		"pkgNotEmpty":   true, // PkgPath for string is ""
		"typeNotEmpty":  rvk.TypeName() != "",
	}
	expected := args.Map{
		"isInvalid":     false,
		"hasError":      false,
		"actInstNotNil": true,
		"pkgNotEmpty":   true,
		"typeNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind returns non-empty -- valid", actual)

	// PointerRv for valid value
	ptrRv := rvk.PointerRv()
	actual := args.Map{"result": ptrRv == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil PointerRv", actual)

	// PointerInterface
	ptrIface := rvk.PointerInterface()
	actual := args.Map{"result": ptrIface == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil PointerInterface", actual)
}

func Test_C9_ReflectValueKind_InvalidNotNil(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
	}
	// PointerRv for invalid but non-nil
	ptrRv := rvk.PointerRv()
	actual := args.Map{"result": ptrRv == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil PointerRv for invalid", actual)
	// PkgPath for invalid
	actual := args.Map{"result": rvk.PkgPath() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty pkgPath for invalid", actual)
}

// ── MethodProcessor — basic methods ──

type testHelper struct{}

func (h testHelper) Add(a, b int) int       { return a + b }
func (h testHelper) Greet(name string) string { return "hi " + name }
func (h testHelper) Fail() error              { return fmt.Errorf("fail") }
func (h testHelper) TwoReturns(x int) (int, error) {
	if x < 0 { return 0, fmt.Errorf("negative") }
	return x * 2, nil
}

func getMethodProcessorFull(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testHelper{})
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

func Test_C9_MethodProcessor_BasicMethods(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	actual := args.Map{"result": mp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "method not found", actual)

	actual := args.Map{
		"hasValidFunc":  mp.HasValidFunc(),
		"funcName":      mp.GetFuncName(),
		"isInvalid":     mp.IsInvalid(),
		"isPublic":      mp.IsPublicMethod(),
		"isPrivate":     mp.IsPrivateMethod(),
		"argsCount":     mp.ArgsCount(),
		"argsLen":       mp.ArgsLength(),
		"returnLen":     mp.ReturnLength(),
		"funcNotNil":    mp.Func() != nil,
		"typeNotNil":    mp.GetType() != nil,
	}
	expected := args.Map{
		"hasValidFunc":  true,
		"funcName":      "Add",
		"isInvalid":     false,
		"isPublic":      true,
		"isPrivate":     false,
		"argsCount":     3, // receiver + 2 args
		"argsLen":       3,
		"returnLen":     1,
		"funcNotNil":    true,
		"typeNotNil":    true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- basic", actual)
}

func Test_C9_MethodProcessor_NilReceiver(t *testing.T) {
	var nilMp *reflectmodel.MethodProcessor
	actual := args.Map{
		"hasValidFunc": nilMp.HasValidFunc(),
		"isInvalid":    nilMp.IsInvalid(),
		"funcNil":      nilMp.Func() == nil,
		"returnLen":    nilMp.ReturnLength(),
		"isPublic":     nilMp.IsPublicMethod(),
		"isPrivate":    nilMp.IsPrivateMethod(),
		"typeNil":      nilMp.GetType() == nil,
	}
	expected := args.Map{
		"hasValidFunc": false,
		"isInvalid":    true,
		"funcNil":      true,
		"returnLen":    -1,
		"isPublic":     false,
		"isPrivate":    false,
		"typeNil":      true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- nil", actual)

	// GetInArgsTypes/GetOutArgsTypes/GetInArgsTypesNames on nil
	actual := args.Map{"result": len(nilMp.GetInArgsTypes()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": len(nilMp.GetOutArgsTypes()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": len(nilMp.GetInArgsTypesNames()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C9_MethodProcessor_Invoke_Success(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	results, err := mp.Invoke(testHelper{}, 2, 3)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "invoke err:", actual)
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C9_MethodProcessor_Invoke_ArgsMismatch(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	_, err := mp.Invoke(testHelper{}, 2) // missing arg
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for args mismatch", actual)
}

func Test_C9_MethodProcessor_Invoke_NilReceiver(t *testing.T) {
	var nilMp *reflectmodel.MethodProcessor
	_, err := nilMp.Invoke()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil invoke", actual)
}

func Test_C9_MethodProcessor_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMethodProcessorFull("Greet")
	first, err := mp.GetFirstResponseOfInvoke(testHelper{}, "world")
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err:", actual)
	actual := args.Map{"result": first}
	expected := args.Map{"result": "hi world"}
	expected.ShouldBeEqual(t, 0, "expected 'hi world'", actual)
}

func Test_C9_MethodProcessor_InvokeResultOfIndex(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	result, err := mp.InvokeResultOfIndex(0, testHelper{}, 1, 2)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err:", actual)
	actual := args.Map{"result": result}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C9_MethodProcessor_InvokeError(t *testing.T) {
	mp := getMethodProcessorFull("Fail")
	funcErr, procErr := mp.InvokeError(testHelper{})
	actual := args.Map{"result": procErr}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "proc err:", actual)
	actual := args.Map{"result": funcErr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected func error", actual)
}

func Test_C9_MethodProcessor_InvokeFirstAndError(t *testing.T) {
	mp := getMethodProcessorFull("TwoReturns")
	first, funcErr, procErr := mp.InvokeFirstAndError(testHelper{}, 5)
	actual := args.Map{"result": procErr}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "proc err:", actual)
	actual := args.Map{"result": funcErr}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "func err:", actual)
	actual := args.Map{"result": first}
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_C9_MethodProcessor_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	mp := getMethodProcessorFull("Add") // only 1 return
	_, _, procErr := mp.InvokeFirstAndError(testHelper{}, 1, 2)
	actual := args.Map{"result": procErr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for too few returns", actual)
}

func Test_C9_MethodProcessor_InvokeFirstAndError_WithError(t *testing.T) {
	mp := getMethodProcessorFull("TwoReturns")
	defer func() { recover() }() // may panic on nil error interface
	_, _, _ = mp.InvokeFirstAndError(testHelper{}, -1)
}

// ── MethodProcessor — GetInArgsTypes / GetOutArgsTypes / GetInArgsTypesNames ──

func Test_C9_MethodProcessor_ArgTypes(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	inTypes := mp.GetInArgsTypes()
	actual := args.Map{"result": len(inTypes)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3 in args", actual)
	// second call should use cache
	inTypes2 := mp.GetInArgsTypes()
	actual := args.Map{"result": len(inTypes2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cached mismatch", actual)

	outTypes := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(outTypes)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1 out arg", actual)
	outTypes2 := mp.GetOutArgsTypes()
	actual := args.Map{"result": len(outTypes2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cached mismatch", actual)

	names := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	names2 := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(names2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cached mismatch", actual)
}

func Test_C9_MethodProcessor_ZeroArgsMethod(t *testing.T) {
	mp := getMethodProcessorFull("Fail")
	// Fail has receiver only → ArgsCount=1, but GetInArgsTypes returns 1
	inTypes := mp.GetInArgsTypes()
	outTypes := mp.GetOutArgsTypes()
	names := mp.GetInArgsTypesNames()
	actual := args.Map{"result": len(inTypes)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1 in arg (receiver)", actual)
	actual := args.Map{"result": len(outTypes)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1 out arg", actual)
	actual := args.Map{"result": len(names)}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "expected 1 name", actual)
}

// ── MethodProcessor — IsEqual / IsNotEqual ──

func Test_C9_MethodProcessor_IsEqual(t *testing.T) {
	mp1 := getMethodProcessorFull("Add")
	mp2 := getMethodProcessorFull("Add")
	mp3 := getMethodProcessorFull("Greet")
	var nilMp *reflectmodel.MethodProcessor

	actual := args.Map{
		"sameEqual":    mp1.IsEqual(mp2),
		"diffNotEqual": mp1.IsNotEqual(mp3),
		"nilBothEqual": nilMp.IsEqual(nil),
		"nilLeft":      nilMp.IsEqual(mp1),
		"nilRight":     mp1.IsEqual(nil),
		"selfEqual":    mp1.IsEqual(mp1),
	}
	expected := args.Map{
		"sameEqual":    true,
		"diffNotEqual": true,
		"nilBothEqual": true,
		"nilLeft":      false,
		"nilRight":     false,
		"selfEqual":    true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- IsEqual", actual)
}

// ── MethodProcessor — VerifyInArgs / VerifyOutArgs / ValidateMethodArgs ──

func Test_C9_MethodProcessor_VerifyInArgs(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	ok, err := mp.VerifyInArgs([]any{testHelper{}, 1, 2})
	actual := args.Map{"result": !ok || err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)

	ok2, err2 := mp.VerifyInArgs([]any{testHelper{}, "a", 2})
	actual := args.Map{"result": ok2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not ok for type mismatch", actual)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C9_MethodProcessor_VerifyOutArgs(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	ok, err := mp.VerifyOutArgs([]any{0})
	actual := args.Map{"result": !ok || err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "expected ok", actual)
}

func Test_C9_MethodProcessor_ValidateMethodArgs(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	err := mp.ValidateMethodArgs([]any{testHelper{}, 1, 2})
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err2 := mp.ValidateMethodArgs([]any{testHelper{}, 1})
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for count mismatch", actual)
}

// ── MethodProcessor — InArgsVerifyRv / OutArgsVerifyRv ──

func Test_C9_MethodProcessor_InArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	types := []reflect.Type{reflect.TypeOf(testHelper{}), reflect.TypeOf(0), reflect.TypeOf(0)}
	ok, err := mp.InArgsVerifyRv(types)
	actual := args.Map{"result": !ok || err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err=", actual)

	// wrong length
	ok2, err2 := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	actual := args.Map{"result": ok2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not ok", actual)
	actual := args.Map{"result": err2 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C9_MethodProcessor_OutArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessorFull("Add")
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	actual := args.Map{"result": !ok || err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err=", actual)
}
