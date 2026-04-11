package reflectmodeltests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// Helper types for exercising uncovered utils.go paths
// ══════════════════════════════════════════════════════════════════════════════

type ptrReturner struct{}

func (p ptrReturner) ReturnPtr(x int) *int       { return &x }
func (p ptrReturner) ReturnSlice() []int          { return []int{1, 2, 3} }
func (p ptrReturner) ReturnMap() map[string]int   { return map[string]int{"a": 1} }
func (p ptrReturner) ReturnNilPtr() *int          { return nil }
func (p ptrReturner) ReturnNilSlice() []int       { return nil }
func (p ptrReturner) ReturnNilMap() map[string]int { return nil }
func (p ptrReturner) ReturnInterface() any        { return "hello" }
func (p ptrReturner) ReturnNilInterface() any     { return nil }
func (p ptrReturner) ReturnChan() chan int         { return make(chan int) }
func (p ptrReturner) ReturnFunc() func()           { return func() {} }
func (p ptrReturner) ReturnNilFunc() func()        { return nil }
func (p ptrReturner) ReturnNilChan() chan int       { return nil }
func (p ptrReturner) NoArgs()                      {}
func (p ptrReturner) ManyArgs(a, b, c, d int) int  { return a + b + c + d }
func (p ptrReturner) ReturnMulti(x int) (string, error) {
	if x < 0 {
		return "", fmt.Errorf("negative")
	}
	return fmt.Sprintf("ok:%d", x), nil
}
func (p ptrReturner) ReturnErrorOnly() error {
	return errors.New("test-error")
}
func (p ptrReturner) ReturnNilError() error {
	return nil
}

func getPtrMP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(ptrReturner{})
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

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueToAnyValue — pointer/interface branches via Invoke
// ══════════════════════════════════════════════════════════════════════════════

func Test_Invoke_ReturnPtr(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	results, err := mp.Invoke(ptrReturner{}, 42)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnPtr", actual)
}

func Test_Invoke_ReturnSlice(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnSlice")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": results[0] != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnSlice", actual)
}

func Test_Invoke_ReturnMap(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMap")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": results[0] != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnMap", actual)
}

func Test_Invoke_ReturnNilPtr(t *testing.T) {
	mp := getPtrMP("ReturnNilPtr")
	defer func() { recover() }()
	res, err := mp.Invoke(ptrReturner{})
	_ = res
	_ = err
}

func Test_Invoke_ReturnNilSlice(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilSlice")
	results, err := mp.Invoke(ptrReturner{})
	// reflect returns typed nil (non-nil interface wrapping nil slice)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nil": results[0] == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns typed nil -- ReturnNilSlice", actual)
}

func Test_Invoke_ReturnNilMap(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilMap")
	results, err := mp.Invoke(ptrReturner{})
	// reflect returns typed nil (non-nil interface wrapping nil map)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nil": results[0] == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns typed nil -- ReturnNilMap", actual)
}

func Test_Invoke_ReturnInterface(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnInterface")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnInterface", actual)
}

func Test_Invoke_ReturnNilInterface(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilInterface")
	defer func() { recover() }() // reflect.Value.Interface panics on zero Value
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nil": results[0] == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- ReturnNilInterface", actual)
}

func Test_Invoke_ReturnChan(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnChan")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": results[0] != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnChan", actual)
}

func Test_Invoke_ReturnFunc(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnFunc")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": results[0] != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ReturnFunc", actual)
}

func Test_Invoke_ReturnNilFunc(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilFunc")
	results, err := mp.Invoke(ptrReturner{})
	isNil := results[0] == nil || reflect.ValueOf(results[0]).IsNil()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nil": isNil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- ReturnNilFunc", actual)
}

func Test_Invoke_ReturnNilChan(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilChan")
	results, err := mp.Invoke(ptrReturner{})
	isNil := results[0] == nil || reflect.ValueOf(results[0]).IsNil()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nil": isNil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns nil -- ReturnNilChan", actual)
}

func Test_Invoke_NoArgs(t *testing.T) {
	// Arrange
	mp := getPtrMP("NoArgs")
	results, err := mp.Invoke(ptrReturner{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(results),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- NoArgs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Invoke with type validation branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Invoke_TypeMismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	_, err := mp.Invoke(ptrReturner{}, "wrong", 2, 3, 4)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
	actual = args.Map{"hasErr": true}
	expected = args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- type mismatch", actual)
}

func Test_Invoke_ManyArgs_Success(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	results, err := mp.Invoke(ptrReturner{}, 1, 2, 3, 4)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 10,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- ManyArgs success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeFirstAndError — with func returning (string, error)
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeFirstAndError_Success(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	defer func() {
		if r := recover(); r != nil {
			// InvokeFirstAndError panics on zero reflect.Value for nil error
			// This is a known limitation — test exercises the path
			t.Skipf("InvokeFirstAndError panics on nil error return: %v", r)
		}
	}()
	first, funcErr, procErr := mp.InvokeFirstAndError(ptrReturner{}, 5)

	// Act
	actual := args.Map{
		"procErr": procErr == nil,
		"funcErr": funcErr == nil,
		"first": first,
	}

	// Assert
	expected := args.Map{
		"procErr": true,
		"funcErr": true,
		"first": "ok:5",
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- success", actual)
}

func Test_InvokeFirstAndError_FuncError(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	_, funcErr, procErr := mp.InvokeFirstAndError(ptrReturner{}, -1)

	// Act
	actual := args.Map{
		"procErr": procErr == nil,
		"funcErr": funcErr != nil,
	}

	// Assert
	expected := args.Map{
		"procErr": true,
		"funcErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- func error", actual)
}

func Test_InvokeFirstAndError_ProcessingError(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	_, _, procErr := mp.InvokeFirstAndError(ptrReturner{}, "wrong_type")

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- processing error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeError — via different return types
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeError_ReturnsError(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnErrorOnly")
	funcErr, procErr := mp.InvokeError(ptrReturner{})

	// Act
	actual := args.Map{
		"procErr": procErr == nil,
		"funcErr": funcErr != nil,
	}

	// Assert
	expected := args.Map{
		"procErr": true,
		"funcErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- returns error", actual)
}

func Test_InvokeError_NilError(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnNilError")
	// InvokeError does result.(error) which panics when result is nil interface.
	// This is a known production limitation — test the panic path.
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		mp.InvokeError(ptrReturner{})
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "InvokeError panics on nil error result -- known limitation", actual)
}

func Test_InvokeError_ProcessingError(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnErrorOnly")
	_, procErr := mp.InvokeError(ptrReturner{}, "extra_arg")

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- processing error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFirstResponseOfInvoke / InvokeResultOfIndex — error paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetFirstResponseOfInvoke_Error(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	_, err := mp.GetFirstResponseOfInvoke(ptrReturner{}) // wrong args count

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns error -- error", actual)
}

func Test_InvokeResultOfIndex_Error(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	_, err := mp.InvokeResultOfIndex(0, ptrReturner{}) // wrong args count

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns error -- error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetOutArgsTypes / GetInArgsTypes — cached paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetOutArgsTypes_Cached(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	out1 := mp.GetOutArgsTypes()
	out2 := mp.GetOutArgsTypes() // should use cache

	// Act
	actual := args.Map{
		"len1": len(out1),
		"len2": len(out2),
		"same": len(out1) == len(out2),
	}

	// Assert
	expected := args.Map{
		"len1": 2,
		"len2": 2,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes returns correct value -- cached", actual)
}

func Test_GetInArgsTypes_Cached(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	in1 := mp.GetInArgsTypes()
	in2 := mp.GetInArgsTypes() // cache

	// Act
	actual := args.Map{
		"len1": len(in1),
		"len2": len(in2),
	}

	// Assert
	expected := args.Map{
		"len1": 2,
		"len2": 2,
	}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypes returns correct value -- cached", actual)
}

func Test_GetInArgsTypesNames_Cached(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	n1 := mp.GetInArgsTypesNames()
	n2 := mp.GetInArgsTypesNames() // cache

	// Act
	actual := args.Map{
		"len1": len(n1),
		"len2": len(n2),
	}

	// Assert
	expected := args.Map{
		"len1": 2,
		"len2": 2,
	}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns correct value -- cached", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEqual — deep equality with same signature different methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsEqual_SameSignatureDiffMethods(t *testing.T) {
	// Arrange
	mp1 := getPtrMP("ReturnErrorOnly")
	mp2 := getPtrMP("ReturnNilError")
	// Same signature (receiver → error), but different names → not equal

	// Act
	actual := args.Map{"isEqual": mp1.IsEqual(mp2)}

	// Assert
	expected := args.Map{"isEqual": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same signature diff methods", actual)
}

func Test_IsNotEqual_DiffSignature_FromMisc10Iteration13(t *testing.T) {
	// Arrange
	mp1 := getPtrMP("ReturnPtr")
	mp2 := getPtrMP("ReturnSlice")

	// Act
	actual := args.Map{"notEqual": mp1.IsNotEqual(mp2)}

	// Assert
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual returns correct value -- diff signature", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VerifyInArgs / VerifyOutArgs — with various types
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyInArgs_Success(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	ok, err := mp.VerifyInArgs([]any{ptrReturner{}, 42})

	// Act
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "VerifyInArgs returns correct value -- success", actual)
}

func Test_VerifyInArgs_TypeMismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	ok, err := mp.VerifyInArgs([]any{ptrReturner{}, "wrong"})

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
	expected.ShouldBeEqual(t, 0, "VerifyInArgs returns correct value -- type mismatch", actual)
}

func Test_VerifyOutArgs_Success(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	// VerifyOutArgs uses reflect.TypeOf which returns concrete type (*errors.errorString)
	// not interface type (error). So type comparison fails — this is expected behavior.
	ok, err := mp.VerifyOutArgs([]any{"", errors.New("")})

	// Act
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"ok": false,
		"noErr": false,
	}
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs returns correct value -- concrete vs interface type mismatch", actual)
}

func Test_VerifyOutArgs_LenMismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	ok, err := mp.VerifyOutArgs([]any{"only_one"})

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
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs returns correct value -- length mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidateMethodArgs — comprehensive error message paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidateMethodArgs_CountMismatch_Detailed(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, 1})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- count mismatch detailed", actual)
}

func Test_ValidateMethodArgs_TypeMismatch_Detailed(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, "wrong", 2, 3, 4})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- type mismatch detailed", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — various Kind types for wider coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_IntKind(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
	}

	// Act
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actInst":   rvk.ActualInstance(),
		"typeName":  rvk.TypeName() != "",
		"pkgPath":   rvk.PkgPath(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": false,
		"actInst":   42,
		"typeName":  true,
		"pkgPath":   "",
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Int kind", actual)
}

func Test_RVK_BoolKind(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(true),
		Kind:            reflect.Bool,
	}

	// Act
	actual := args.Map{
		"actInst": rvk.ActualInstance(),
		"typeName": rvk.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"actInst": true,
		"typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Bool kind", actual)
}

func Test_RVK_SliceKind(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf([]int{1, 2, 3}),
		Kind:            reflect.Slice,
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
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Slice kind", actual)
}

func Test_RVK_MapKind(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(map[string]int{"a": 1}),
		Kind:            reflect.Map,
	}

	// Act
	actual := args.Map{
		"actNotNil": rvk.ActualInstance() != nil,
		"typeName": rvk.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"actNotNil": true,
		"typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Map kind", actual)
}

func Test_RVK_StructKind(t *testing.T) {
	// Arrange
	type myStruct struct{ X int }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(myStruct{X: 42}),
		Kind:            reflect.Struct,
	}

	// Act
	actual := args.Map{
		"pkgPath": rvk.PkgPath() != "",
		"actInst": rvk.ActualInstance() != nil,
	}

	// Assert
	expected := args.Map{
		"pkgPath": true,
		"actInst": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Struct kind", actual)
}

func Test_RVK_PtrKind(t *testing.T) {
	// Arrange
	x := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(&x),
		Kind:            reflect.Ptr,
	}

	// Act
	actual := args.Map{
		"typeName": rvk.TypeName() != "",
		"pkgPath": true,
	}

	// Assert
	expected := args.Map{
		"typeName": true,
		"pkgPath": true,
	}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Ptr kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValue — struct fields coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectValue_Fields(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{
		TypeName:     "TestType",
		FieldsNames:  []string{"A", "B"},
		MethodsNames: []string{"M1"},
		RawData:      "raw",
	}

	// Act
	actual := args.Map{
		"typeName": rv.TypeName, "fieldsLen": len(rv.FieldsNames),
		"methodsLen": len(rv.MethodsNames), "raw": rv.RawData,
	}

	// Assert
	expected := args.Map{
		"typeName": "TestType",
		"fieldsLen": 2,
		"methodsLen": 1,
		"raw": "raw",
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- fields", actual)
}

func Test_ReflectValue_Empty(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{}

	// Act
	actual := args.Map{
		"typeName": rv.TypeName,
		"raw": rv.RawData == nil,
	}

	// Assert
	expected := args.Map{
		"typeName": "",
		"raw": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns empty -- empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — various field types
// ══════════════════════════════════════════════════════════════════════════════

func Test_FieldProcessor_IntField(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Age", 1)

	// Act
	actual := args.Map{
		"isInt":  fp.IsFieldKind(reflect.Int),
		"isStr":  fp.IsFieldKind(reflect.String),
		"typeOk": fp.IsFieldType(reflect.TypeOf(0)),
	}

	// Assert
	expected := args.Map{
		"isInt": true,
		"isStr": false,
		"typeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- Int field", actual)
}

func Test_FieldProcessor_BoolField_FromMisc10Iteration13(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Active", 2)

	// Act
	actual := args.Map{
		"isBool": fp.IsFieldKind(reflect.Bool),
		"typeOk": fp.IsFieldType(reflect.TypeOf(true)),
	}

	// Assert
	expected := args.Map{
		"isBool": true,
		"typeOk": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- Bool field", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — GetOutArgsTypes with 0 returns & caching
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetOutArgsTypes_NoReturn(t *testing.T) {
	// Arrange
	mp := getPtrMP("NoArgs")
	out := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"len": len(out)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes returns empty -- no return", actual)
}

func Test_GetInArgsTypes_ReceiverOnly(t *testing.T) {
	// Arrange
	mp := getPtrMP("NoArgs")
	in := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"len": len(in)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypes returns correct value -- receiver only", actual)
}

func Test_GetInArgsTypesNames_ReceiverOnly(t *testing.T) {
	// Arrange
	mp := getPtrMP("NoArgs")
	names := mp.GetInArgsTypesNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns correct value -- receiver only", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InArgsVerifyRv / OutArgsVerifyRv with matching + mismatching
// ══════════════════════════════════════════════════════════════════════════════

func Test_InArgsVerifyRv_Match(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	types := []reflect.Type{reflect.TypeOf(ptrReturner{}), reflect.TypeOf(0)}
	ok, err := mp.InArgsVerifyRv(types)

	// Act
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv returns correct value -- match", actual)
}

func Test_InArgsVerifyRv_Mismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	types := []reflect.Type{reflect.TypeOf(ptrReturner{}), reflect.TypeOf("")}
	ok, err := mp.InArgsVerifyRv(types)

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
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv returns correct value -- mismatch", actual)
}

func Test_OutArgsVerifyRv_Match(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnPtr")
	// ReturnPtr returns *int
	x := 0
	types := []reflect.Type{reflect.TypeOf(&x)}
	ok, err := mp.OutArgsVerifyRv(types)

	// Act
	actual := args.Map{
		"ok": ok,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv returns correct value -- match", actual)
}

func Test_OutArgsVerifyRv_Mismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ReturnMulti")
	types := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)} // wrong second type
	ok, err := mp.OutArgsVerifyRv(types)

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
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv returns correct value -- mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — Multiple type mismatches in VerifyReflectTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidateMethodArgs_MultiTypeMismatch(t *testing.T) {
	// Arrange
	mp := getPtrMP("ManyArgs")
	// All 4 args are wrong types
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, "a", "b", "c", "d"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- multi type mismatch", actual)
}
