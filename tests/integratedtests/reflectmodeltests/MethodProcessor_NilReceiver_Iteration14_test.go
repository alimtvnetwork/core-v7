package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_MethodProcessor_NilReceiver_HasValidFunc(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"hasValid": mp.HasValidFunc()}

	// Assert
	expected := args.Map{"hasValid": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- HasValidFunc", actual)
}

func Test_MethodProcessor_NilReceiver_IsInvalid(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"isInvalid": mp.IsInvalid()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- IsInvalid", actual)
}

func Test_MethodProcessor_NilReceiver_Func(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"funcNil": mp.Func() == nil}

	// Assert
	expected := args.Map{"funcNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Func", actual)
}

func Test_MethodProcessor_NilReceiver_ReturnLength(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"retLen": mp.ReturnLength()}

	// Assert
	expected := args.Map{"retLen": -1}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- ReturnLength", actual)
}

func Test_MethodProcessor_NilReceiver_GetType(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"typeNil": mp.GetType() == nil}

	// Assert
	expected := args.Map{"typeNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetType", actual)
}

func Test_MethodProcessor_NilReceiver_GetOutArgsTypes(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	out := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"len": len(out)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetOutArgsTypes", actual)
}

func Test_MethodProcessor_NilReceiver_GetInArgsTypes(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	in := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"len": len(in)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetInArgsTypes", actual)
}

func Test_MethodProcessor_NilReceiver_GetInArgsTypesNames(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetInArgsTypesNames", actual)
}

func Test_MethodProcessor_NilReceiver_Invoke(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Invoke", actual)
}

func Test_MethodProcessor_NilReceiver_IsPublic(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{
		"isPublic": mp.IsPublicMethod(),
		"isPrivate": mp.IsPrivateMethod(),
	}

	// Assert
	expected := args.Map{
		"isPublic": false,
		"isPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- IsPublicMethod/IsPrivateMethod", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEqual — nil and same-pointer paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsEqual_BothNil_FromMethodProcessorNilRe(t *testing.T) {
	// Arrange
	var a, b *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- both nil", actual)
}

func Test_IsEqual_LeftNil(t *testing.T) {
	// Arrange
	var a *reflectmodel.MethodProcessor
	b := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- left nil", actual)
}

func Test_IsEqual_RightNil(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"eq": a.IsEqual(nil)}

	// Assert
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- right nil", actual)
}

func Test_IsEqual_SamePointer_FromMethodProcessorNilRe(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"eq": a.IsEqual(a)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same pointer", actual)
}

func Test_IsEqual_SameSignature(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"eq": a.IsEqual(b)}

	// Assert
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same signature", actual)
}

func Test_IsNotEqual_DiffSignature(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")

	// Act
	actual := args.Map{"notEq": a.IsNotEqual(b)}

	// Assert
	expected := args.Map{"notEq": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual returns correct value -- different signature", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeFirstAndError — single return (error path)
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeFirstAndError_SingleReturn(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")
	_, _, procErr := mp.InvokeFirstAndError(sampleStruct{})

	// Act
	actual := args.Map{"hasErr": procErr != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- single return", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeError — with non-error return type (should panic or error)
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeError_NoArgsError(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "InvokeError panics -- non-error return panics", actual)
	}()
	mp.InvokeError(sampleStruct{})
}

// ══════════════════════════════════════════════════════════════════════════════
// InvalidReflectValueKindModel — constructor
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvalidReflectValueKindModel(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("test error msg")

	// Act
	actual := args.Map{
		"isValid":   rvk.IsValid,
		"hasErr":    rvk.HasError(),
		"isInvalid": rvk.IsInvalid(),
		"emptyErr":  rvk.IsEmptyError(),
		"typeName":  rvk.TypeName(),
		"pkgPath":   rvk.PkgPath(),
	}

	// Assert
	expected := args.Map{
		"isValid":   false,
		"hasErr":    true,
		"isInvalid": true,
		"emptyErr":  false,
		"typeName":  "",
		"pkgPath":   "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidReflectValueKindModel returns error -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_NilReceiver_IsInvalid(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"hasErr": rvk.HasError(),
		"emptyErr": rvk.IsEmptyError(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true,
		"hasErr": false,
		"emptyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK IsInvalid/HasError/IsEmptyError", actual)
}

func Test_RVK_NilReceiver_ActualInstance(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"actNil": rvk.ActualInstance() == nil}

	// Assert
	expected := args.Map{"actNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK ActualInstance", actual)
}

func Test_RVK_NilReceiver_PkgPath(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"pkgPath": rvk.PkgPath()}

	// Assert
	expected := args.Map{"pkgPath": ""}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PkgPath", actual)
}

func Test_RVK_NilReceiver_PointerRv(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"nil": rvk.PointerRv() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PointerRv", actual)
}

func Test_RVK_NilReceiver_PointerInterface(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"nil": rvk.PointerInterface() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PointerInterface", actual)
}

func Test_RVK_NilReceiver_TypeName(t *testing.T) {
	// Arrange
	var rvk *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{"typeName": rvk.TypeName()}

	// Assert
	expected := args.Map{"typeName": ""}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK TypeName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — invalid (IsValid=false) with non-nil receiver
// ══════════════════════════════════════════════════════════════════════════════

func Test_RVK_InvalidNotNil_PointerRv(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
	}
	ptr := rvk.PointerRv()

	// Act
	actual := args.Map{
		"notNil": ptr != nil,
		"pkgPath": rvk.PkgPath(),
		"typeName": rvk.TypeName(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"pkgPath": "",
		"typeName": "",
	}
	expected.ShouldBeEqual(t, 0, "RVK returns nil -- invalid non-nil PointerRv", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_FieldProcessor_NilReceiver_IsFieldType(t *testing.T) {
	// Arrange
	var fp *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{"isType": fp.IsFieldType(reflect.TypeOf(0))}

	// Assert
	expected := args.Map{"isType": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- FieldProcessor IsFieldType", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// argsCountMismatchErrorMessage — triggered via ValidateMethodArgs
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidateMethodArgs_TooFewArgs(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")
	// PublicMethod expects (sampleStruct, string, int) = 3 args; give 1
	err := mp.ValidateMethodArgs([]any{sampleStruct{}})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- too few args", actual)
}

func Test_ValidateMethodArgs_TooManyArgs(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{sampleStruct{}, "a", 1, "extra", "extra2"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- too many args", actual)
}

func Test_ValidateMethodArgs_EmptyArgs(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns empty -- empty args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MultiReturn method — InvokeFirstAndError exercises multi-return parsing
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeFirstAndError_MultiReturn(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("MultiReturn")
	// InvokeFirstAndError should return processing error when 2nd return is non-error.
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		mp.InvokeFirstAndError(sampleStruct{})
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": false}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError panics -- MultiReturn string not error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFirstResponseOfInvoke — success path
// ══════════════════════════════════════════════════════════════════════════════

func Test_GetFirstResponseOfInvoke_Success(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")
	first, err := mp.GetFirstResponseOfInvoke(sampleStruct{})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": first,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns correct value -- success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeResultOfIndex — success path index 1
// ══════════════════════════════════════════════════════════════════════════════

func Test_InvokeResultOfIndex_SecondResult(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")
	second, err := mp.InvokeResultOfIndex(1, sampleStruct{}, "test", 42)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nilErr": second == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns correct value -- second result", actual)
}
