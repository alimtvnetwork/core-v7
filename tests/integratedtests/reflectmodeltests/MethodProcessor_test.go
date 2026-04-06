package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== MethodProcessor Tests =====

// --- Validity & Identity ---

func Test_MethodProcessor_HasValidFunc_True(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create MethodProcessor for PublicMethod", actual)

	actual := args.Map{"result": mp.HasValidFunc()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasValidFunc() = true", actual)
}

// Note: HasValidFunc nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// Note: IsInvalid nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_IsInvalid_Valid(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp.IsInvalid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid() = false for valid method", actual)
}

func Test_MethodProcessor_GetFuncName(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp.GetFuncName() != "PublicMethod"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetFuncName() =, want", mp.GetFuncName(), "PublicMethod", actual)
}

// --- Func ---

func Test_MethodProcessor_Func_Valid(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	fn := mp.Func()

	// Act
	actual := args.Map{"result": fn == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Func() to return non-nil for valid method", actual)
}

// Note: Func nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- Args & Return Counts ---

func Test_MethodProcessor_ArgsCount(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// reflect.Method includes receiver as first arg
	// sampleStruct.PublicMethod(a string, b int) => 3 args (receiver + 2)
	got := mp.ArgsCount()

	// Act
	actual := args.Map{"result": got != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgsCount() =, want 3 (receiver + 2 params)", actual)
}

func Test_MethodProcessor_ArgsLength(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	// NoArgsMethod() => 1 arg (receiver only)
	got := mp.ArgsLength()

	// Act
	actual := args.Map{"result": got != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ArgsLength() =, want 1 (receiver only)", actual)
}

func Test_MethodProcessor_ReturnLength(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	got := mp.ReturnLength()

	// Act
	actual := args.Map{"result": got != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReturnLength() =, want 2 (string, error)", actual)
}

// Note: ReturnLength nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_ReturnLength_MultiReturn(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("MultiReturn")

	got := mp.ReturnLength()

	// Act
	actual := args.Map{"result": got != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReturnLength() =, want 3 (int, string, error)", actual)
}

// --- Public/Private ---

func Test_MethodProcessor_IsPublicMethod(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp.IsPublicMethod()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsPublicMethod() = true for PublicMethod", actual)
}

// Note: IsPublicMethod nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_IsPrivateMethod_False(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp.IsPrivateMethod()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsPrivateMethod() = false for PublicMethod", actual)
}

// --- GetType ---

func Test_MethodProcessor_GetType_Valid(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	got := mp.GetType()

	// Act
	actual := args.Map{"result": got == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected GetType() to return non-nil for valid method", actual)
}

// Note: GetType nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- InArgs & OutArgs Types ---

func Test_MethodProcessor_GetInArgsTypes(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	types := mp.GetInArgsTypes()
	// receiver + 2 params = 3

	// Act
	actual := args.Map{"result": len(types) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypes() len =, want 3", actual)
}

// Note: GetInArgsTypes nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_GetInArgsTypes_Cached(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	first := mp.GetInArgsTypes()
	second := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"result": len(first) != len(second)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached GetInArgsTypes to return same length", actual)
}

func Test_MethodProcessor_GetOutArgsTypes(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	types := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(types) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes() len =, want 2", actual)
}

// Note: GetOutArgsTypes nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_GetOutArgsTypes_NoArgs(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	// NoArgsMethod returns string => 1 out type
	types := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"result": len(types) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes() len =, want 1", actual)
}

func Test_MethodProcessor_GetInArgsTypesNames(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	names := mp.GetInArgsTypesNames()
	// receiver type + string + int = 3

	// Act
	actual := args.Map{"result": len(names) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames() len =, want 3", actual)
}

// Note: GetInArgsTypesNames nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- IsEqual / IsNotEqual ---

func Test_MethodProcessor_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEqual(nil, nil) = true", actual)
}

func Test_MethodProcessor_IsEqual_OneNil(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")
	var nilMp *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"result": mp.IsEqual(nilMp)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEqual(valid, nil) = false", actual)
}

func Test_MethodProcessor_IsEqual_Same(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": mp.IsEqual(mp)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEqual with itself = true", actual)
}

func Test_MethodProcessor_IsEqual_SameMethod(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEqual for same method = true", actual)
}

func Test_MethodProcessor_IsNotEqual(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")

	// Act
	actual := args.Map{"result": a.IsNotEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsNotEqual for different methods = true", actual)
}

// --- ValidateMethodArgs ---

func Test_MethodProcessor_ValidateMethodArgs_WrongCount(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// PublicMethod expects receiver + string + int = 3 args
	err := mp.ValidateMethodArgs([]any{"a"})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong arg count", actual)
}

func Test_MethodProcessor_ValidateMethodArgs_WrongType(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	// receiver + string + int, but we give receiver + int + int
	err := mp.ValidateMethodArgs([]any{sampleStruct{}, 42, 42})

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for wrong arg type", actual)
}

func Test_MethodProcessor_ValidateMethodArgs_Correct(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	err := mp.ValidateMethodArgs([]any{sampleStruct{}, "hello", 42})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for correct args, got:", actual)
}

// --- VerifyInArgs / VerifyOutArgs ---

func Test_MethodProcessor_VerifyInArgs_Match(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	ok, err := mp.VerifyInArgs([]any{sampleStruct{}, "s", 1})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected VerifyInArgs match, got ok= err=", actual)
}

func Test_MethodProcessor_VerifyOutArgs_Match(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	ok, err := mp.VerifyOutArgs([]any{""})

	// Act
	actual := args.Map{"result": ok || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected VerifyOutArgs match, got ok= err=", actual)
}

func Test_MethodProcessor_InArgsVerifyRv_LengthMismatch(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf("")})

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected InArgsVerifyRv = false for length mismatch", actual)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for length mismatch", actual)
}

// --- Invoke ---

func Test_MethodProcessor_Invoke_Success(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	results, err := mp.Invoke(sampleStruct{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invoke error:", actual)

	actual := args.Map{"result": len(results) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invoke results len =, want 1", actual)

	actual := args.Map{"result": results[0] != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invoke result =, want", results[0], "hello", actual)
}

// Note: Invoke nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_NilReceiver(t *testing.T) {
	for caseIndex, tc := range methodProcessorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_MethodProcessor_Invoke_ArgsMismatch(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("PublicMethod")

	_, err := mp.Invoke(sampleStruct{}, "only one arg")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for args count mismatch", actual)
}

func Test_MethodProcessor_GetFirstResponseOfInvoke(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	first, err := mp.GetFirstResponseOfInvoke(sampleStruct{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke error:", actual)

	actual := args.Map{"result": first != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first response =, want", first, "hello", actual)
}

func Test_MethodProcessor_InvokeResultOfIndex(t *testing.T) {
	// Arrange
	mp := newMethodProcessor("NoArgsMethod")

	result, err := mp.InvokeResultOfIndex(0, sampleStruct{})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex error:", actual)

	actual := args.Map{"result": result != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "result =, want", result, "hello", actual)
}
