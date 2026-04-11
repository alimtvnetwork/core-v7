package argstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage12 — coretests/args final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- FuncWrap.VoidCallNoReturn valid path ---

func Test_FuncWrap_VoidCallNoReturn_Valid(t *testing.T) {
	// Arrange
	called := false
	fw := args.NewFuncWrap.Default(func() { called = true })

	// Act
	err := fw.VoidCallNoReturn()

	// Assert
	convey.Convey("VoidCallNoReturn succeeds for valid void func", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(called, convey.ShouldBeTrue)
	})
}

// --- FuncWrap.InvokeFirstAndError success path ---

func Test_FuncWrap_InvokeFirstAndError_Success(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) { return "ok", nil })

	// Act
	first, funcErr, procErr := fw.InvokeFirstAndError()

	// Assert
	convey.Convey("InvokeFirstAndError returns first and nil error", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldBeNil)
		convey.So(first, convey.ShouldEqual, "ok")
	})
}

func Test_FuncWrap_InvokeFirstAndError_WithError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) {
		return "", errors.New("fail")
	})

	// Act
	first, funcErr, procErr := fw.InvokeFirstAndError()

	// Assert
	convey.Convey("InvokeFirstAndError returns func error", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldNotBeNil)
		convey.So(first, convey.ShouldEqual, "")
	})
}

// --- FuncWrap.InvokeAsAnyError branches ---

func Test_FuncWrap_InvokeAsAnyError_SingleReturn(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "solo" })

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	convey.Convey("InvokeAsAnyError with single return", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldBeNil)
		convey.So(result, convey.ShouldEqual, "solo")
	})
}

func Test_FuncWrap_InvokeAsAnyError_WithFuncError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) {
		return "val", errors.New("err")
	})

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	convey.Convey("InvokeAsAnyError returns func error from second return", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldNotBeNil)
		convey.So(result, convey.ShouldEqual, "val")
	})
}

func Test_FuncWrap_InvokeAsAnyError_NilSecondReturn(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) {
		return "val", nil
	})

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	convey.Convey("InvokeAsAnyError with nil second return", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldBeNil)
		convey.So(result, convey.ShouldEqual, "val")
	})
}

func Test_FuncWrap_InvokeAsAnyError_VoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result, funcErr, procErr := fw.InvokeAsAnyError()

	// Assert
	convey.Convey("InvokeAsAnyError with void func returns nils", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeNil)
	})
}

// --- FuncWrap.InvokeAsError with actual error return ---

func Test_FuncWrap_InvokeAsError_ActualError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() error {
		return errors.New("real error")
	})

	// Act
	funcErr, procErr := fw.InvokeAsError()

	// Assert
	convey.Convey("InvokeAsError returns actual error", t, func() {
		convey.So(procErr, convey.ShouldBeNil)
		convey.So(funcErr, convey.ShouldNotBeNil)
		convey.So(funcErr.Error(), convey.ShouldEqual, "real error")
	})
}

// --- FuncWrap.InvokeAsBool/AsString/AsAny void returns ---

func Test_FuncWrap_InvokeAsBool_VoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result, err := fw.InvokeAsBool()

	// Assert
	convey.Convey("InvokeAsBool with void func returns false", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_InvokeAsString_VoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result, err := fw.InvokeAsString()

	// Assert
	convey.Convey("InvokeAsString with void func returns empty", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_FuncWrap_InvokeAsAny_VoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result, err := fw.InvokeAsAny()

	// Assert
	convey.Convey("InvokeAsAny with void func returns nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_FuncWrap_InvokeAsBool_Success(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() bool { return true })

	// Act
	result, err := fw.InvokeAsBool()

	// Assert
	convey.Convey("InvokeAsBool returns true for bool func", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_InvokeAsString_Success(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "hello" })

	// Act
	result, err := fw.InvokeAsString()

	// Assert
	convey.Convey("InvokeAsString returns value for string func", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldEqual, "hello")
	})
}

func Test_FuncWrap_InvokeAsAny_Success(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() int { return 42 })

	// Act
	result, err := fw.InvokeAsAny()

	// Assert
	convey.Convey("InvokeAsAny returns value", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldEqual, 42)
	})
}

// --- FuncWrap.InvokeSkip panic recovery ---

func Test_FuncWrap_InvokeSkip_PanicRecovery(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() { panic("boom") })

	// Act
	results, err := fw.Invoke()

	// Assert
	convey.Convey("Invoke recovers from panic and returns error", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(err.Error(), convey.ShouldContainSubstring, "boom")
		convey.So(results, convey.ShouldNotBeNil)
	})
}

// --- FuncWrap.IsNotEqual and IsEqualValue ---

func Test_FuncWrap_IsNotEqual(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() string { return "" })
	fw2 := args.NewFuncWrap.Default(func() int { return 0 })

	// Act
	result := fw1.IsNotEqual(fw2)

	// Assert
	convey.Convey("IsNotEqual returns true for different funcs", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsEqualValue(t *testing.T) {
	// Arrange
	sampleFunc := func() string { return "" }
	fw1 := args.NewFuncWrap.Default(sampleFunc)
	fw2 := *args.NewFuncWrap.Default(sampleFunc)

	// Act
	result := fw1.IsEqualValue(fw2)

	// Assert
	convey.Convey("IsEqualValue returns true for same func value", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// --- FuncWrap.IsEqual nil branches ---

func Test_FuncWrap_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var fw1 *args.FuncWrapAny
	var fw2 *args.FuncWrapAny

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("IsEqual returns true when both nil", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsEqual_OneNil(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() {})
	var fw2 *args.FuncWrapAny

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("IsEqual returns false when one nil", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_IsEqual_SamePointer(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw1.IsEqual(fw1)

	// Assert
	convey.Convey("IsEqual returns true for same pointer", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsEqual_DiffInArgs(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func(a string) {})
	fw2 := args.NewFuncWrap.Default(func(a int) {})
	// Force same name
	fw2.Name = fw1.Name

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("IsEqual returns false for different in-arg types", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_FuncWrap_IsEqual_DiffOutArgs(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(func() string { return "" })
	fw2 := args.NewFuncWrap.Default(func() int { return 0 })
	fw2.Name = fw1.Name

	// Act
	result := fw1.IsEqual(fw2)

	// Assert
	convey.Convey("IsEqual returns false for different out-arg types", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

// --- FuncWrap typed helper signature checkers ---

func Test_FuncWrap_IsBoolFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() bool { return true })

	// Act & Assert
	convey.Convey("IsBoolFunc returns true for bool-returning func", t, func() {
		convey.So(fw.IsBoolFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsErrorFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() error { return nil })

	// Act & Assert
	convey.Convey("IsErrorFunc returns true for error-returning func", t, func() {
		convey.So(fw.IsErrorFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsStringFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })

	// Act & Assert
	convey.Convey("IsStringFunc returns true for string-returning func", t, func() {
		convey.So(fw.IsStringFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsAnyFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() int { return 0 })

	// Act & Assert
	convey.Convey("IsAnyFunc returns true for single-return func", t, func() {
		convey.So(fw.IsAnyFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsValueErrorFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) { return "", nil })

	// Act & Assert
	convey.Convey("IsValueErrorFunc returns true for (T, error) func", t, func() {
		convey.So(fw.IsValueErrorFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsVoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act & Assert
	convey.Convey("IsVoidFunc returns true for void func", t, func() {
		convey.So(fw.IsVoidFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsAnyErrorFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) { return "", nil })

	// Act & Assert
	convey.Convey("IsAnyErrorFunc returns true (alias for IsValueErrorFunc)", t, func() {
		convey.So(fw.IsAnyErrorFunc(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_SignatureCheckers_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act & Assert
	convey.Convey("Signature checkers return false for invalid wrap", t, func() {
		convey.So(fw.IsBoolFunc(), convey.ShouldBeFalse)
		convey.So(fw.IsErrorFunc(), convey.ShouldBeFalse)
		convey.So(fw.IsStringFunc(), convey.ShouldBeFalse)
		convey.So(fw.IsAnyFunc(), convey.ShouldBeFalse)
		convey.So(fw.IsValueErrorFunc(), convey.ShouldBeFalse)
		convey.So(fw.IsVoidFunc(), convey.ShouldBeFalse)
	})
}

// --- FuncWrap.VoidCall ---

func Test_FuncWrap_VoidCall(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "x" })

	// Act
	results, err := fw.VoidCall()

	// Assert
	convey.Convey("VoidCall invokes with no args", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(results), convey.ShouldEqual, 1)
	})
}

// --- FuncWrap.GetPascalCaseFuncName ---

func Test_FuncWrap_GetPascalCaseFuncName(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw.GetPascalCaseFuncName()

	// Assert
	convey.Convey("GetPascalCaseFuncName returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_GetPascalCaseFuncName_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	result := fw.GetPascalCaseFuncName()

	// Assert
	convey.Convey("GetPascalCaseFuncName nil receiver returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- FuncWrap.PkgPath, PkgNameOnly, FuncDirectInvokeName ---

func Test_FuncWrap_PkgPath(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw.PkgPath()

	// Assert
	convey.Convey("PkgPath returns non-empty for valid func", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_PkgPath_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})
	_ = fw.PkgPath() // first call to populate cache

	// Act
	result := fw.PkgPath() // second call from cache

	// Assert
	convey.Convey("PkgPath returns cached value on second call", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_PkgNameOnly(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw.PkgNameOnly()

	// Assert
	convey.Convey("PkgNameOnly returns package name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_PkgNameOnly_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})
	_ = fw.PkgNameOnly()

	// Act
	result := fw.PkgNameOnly()

	// Assert
	convey.Convey("PkgNameOnly returns cached value", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_FuncDirectInvokeName(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw.FuncDirectInvokeName()

	// Assert
	convey.Convey("FuncDirectInvokeName returns non-empty", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncWrap_FuncDirectInvokeName_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})
	_ = fw.FuncDirectInvokeName()

	// Act
	result := fw.FuncDirectInvokeName()

	// Assert
	convey.Convey("FuncDirectInvokeName returns cached value", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- FuncWrap.GetType, IsPublicMethod, IsPrivateMethod ---

func Test_FuncWrap_GetType(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	result := fw.GetType()

	// Assert
	convey.Convey("GetType returns non-nil for valid func", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_GetType_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	result := fw.GetType()

	// Assert
	convey.Convey("GetType returns nil for invalid wrap", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_FuncWrap_IsPublicMethod(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act & Assert
	convey.Convey("IsPublicMethod returns true for anonymous func", t, func() {
		convey.So(fw.IsPublicMethod(), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsPrivateMethod(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act & Assert
	convey.Convey("IsPrivateMethod returns false for anonymous func", t, func() {
		convey.So(fw.IsPrivateMethod(), convey.ShouldBeFalse)
	})
}

// --- FuncWrap.HasValidFunc ---

func Test_FuncWrap_HasValidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act & Assert
	convey.Convey("HasValidFunc returns true for valid func", t, func() {
		convey.So(fw.HasValidFunc(), convey.ShouldBeTrue)
	})
}

// --- FuncWrap.ValidateMethodArgs mismatch ---

func Test_FuncWrap_ValidateMethodArgs_CountMismatch(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a, b string) {})

	// Act
	err := fw.ValidateMethodArgs([]any{"only-one"})

	// Assert
	convey.Convey("ValidateMethodArgs returns error on count mismatch", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(err.Error(), convey.ShouldContainSubstring, "arguments count doesn't match")
	})
}

// --- FuncWrap.MustBeValid panics ---

func Test_FuncWrap_MustBeValid_NilPanics(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act & Assert
	convey.Convey("MustBeValid panics on nil", t, func() {
		convey.So(func() { fw.MustBeValid() }, convey.ShouldPanic)
	})
}

func Test_FuncWrap_MustBeValid_InvalidPanics(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not-a-func")

	// Act & Assert
	convey.Convey("MustBeValid panics on invalid", t, func() {
		convey.So(func() { fw.MustBeValid() }, convey.ShouldPanic)
	})
}

// --- FuncWrap.ValidationError ---

func Test_FuncWrap_ValidationError_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	err := fw.ValidationError()

	// Assert
	convey.Convey("ValidationError returns error for nil", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_ValidationError_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not-a-func")

	// Act
	err := fw.ValidationError()

	// Assert
	convey.Convey("ValidationError returns error for invalid", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_ValidationError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	err := fw.ValidationError()

	// Assert
	convey.Convey("ValidationError returns nil for valid func", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- FuncWrap.InvalidError ---

func Test_FuncWrap_InvalidError_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	err := fw.InvalidError()

	// Assert
	convey.Convey("InvalidError returns error for nil", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncWrap_InvalidError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	err := fw.InvalidError()

	// Assert
	convey.Convey("InvalidError returns nil for valid func", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// --- FuncWrapArgs: IsInTypeMatches, IsOutTypeMatches ---

func Test_FuncWrap_IsInTypeMatches(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(s string) {})

	// Act & Assert
	convey.Convey("IsInTypeMatches returns true for matching types", t, func() {
		convey.So(fw.IsInTypeMatches("hello"), convey.ShouldBeTrue)
	})
}

func Test_FuncWrap_IsOutTypeMatches(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })

	// Act & Assert
	convey.Convey("IsOutTypeMatches returns true for matching types", t, func() {
		convey.So(fw.IsOutTypeMatches(""), convey.ShouldBeTrue)
	})
}

// --- FuncWrapArgs: InArgNames single arg ---

func Test_FuncWrap_InArgNames_SingleArg(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(s string) {})

	// Act
	names := fw.InArgNames()

	// Assert
	convey.Convey("InArgNames returns single name for single-arg func", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_InArgNames_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(s string) {})
	_ = fw.InArgNames()

	// Act
	names := fw.InArgNames()

	// Assert
	convey.Convey("InArgNames returns cached names", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

// --- FuncWrapArgs: OutArgNames single arg ---

func Test_FuncWrap_OutArgNames_SingleArg(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })

	// Act
	names := fw.OutArgNames()

	// Assert
	convey.Convey("OutArgNames returns single name for single-return func", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_OutArgNames_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })
	_ = fw.OutArgNames()

	// Act
	names := fw.OutArgNames()

	// Assert
	convey.Convey("OutArgNames returns cached names", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_OutArgNames_MultiArg(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() (string, error) { return "", nil })

	// Act
	names := fw.OutArgNames()

	// Assert
	convey.Convey("OutArgNames returns names for multi-return func", t, func() {
		convey.So(len(names), convey.ShouldEqual, 2)
	})
}

// --- FuncWrapArgs: InArgNamesEachLine, OutArgNamesEachLine for single arg ---

func Test_FuncWrap_InArgNamesEachLine_Single(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(s string) {})

	// Act
	result := fw.InArgNamesEachLine()

	// Assert
	convey.Convey("InArgNamesEachLine single arg returns as-is", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_OutArgNamesEachLine_Single(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })

	// Act
	result := fw.OutArgNamesEachLine()

	// Assert
	convey.Convey("OutArgNamesEachLine single return returns as-is", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

// --- FuncWrapArgs: cached GetInArgsTypes, GetOutArgsTypes ---

func Test_FuncWrap_GetInArgsTypes_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a string) {})
	_ = fw.GetInArgsTypes()

	// Act
	types := fw.GetInArgsTypes()

	// Assert
	convey.Convey("GetInArgsTypes returns cached types", t, func() {
		convey.So(len(types), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_GetOutArgsTypes_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })
	_ = fw.GetOutArgsTypes()

	// Act
	types := fw.GetOutArgsTypes()

	// Assert
	convey.Convey("GetOutArgsTypes returns cached types", t, func() {
		convey.So(len(types), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_GetInArgsTypesNames_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a string) {})
	_ = fw.GetInArgsTypesNames()

	// Act
	names := fw.GetInArgsTypesNames()

	// Assert
	convey.Convey("GetInArgsTypesNames returns cached names", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_GetOutArgsTypesNames_Cached(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })
	_ = fw.GetOutArgsTypesNames()

	// Act
	names := fw.GetOutArgsTypesNames()

	// Assert
	convey.Convey("GetOutArgsTypesNames returns cached names", t, func() {
		convey.So(len(names), convey.ShouldEqual, 1)
	})
}

// --- FuncMap delegate methods ---

func Test_FuncMap_PkgPath(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() string { return "" })
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	result := fm.PkgPath(name)

	// Assert
	convey.Convey("FuncMap.PkgPath returns path", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncMap_PkgNameOnly(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() string { return "" })
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	result := fm.PkgNameOnly(name)

	// Assert
	convey.Convey("FuncMap.PkgNameOnly returns name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncMap_FuncDirectInvokeName(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() string { return "" })
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	result := fm.FuncDirectInvokeName(name)

	// Assert
	convey.Convey("FuncMap.FuncDirectInvokeName returns name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncMap_NotFound_Methods(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	convey.Convey("FuncMap methods return defaults for not-found name", t, func() {
		convey.So(fm.PkgPath("missing"), convey.ShouldBeEmpty)
		convey.So(fm.PkgNameOnly("missing"), convey.ShouldBeEmpty)
		convey.So(fm.FuncDirectInvokeName("missing"), convey.ShouldBeEmpty)
		convey.So(fm.ArgsCount("missing"), convey.ShouldEqual, 0)
		convey.So(fm.ReturnLength("missing"), convey.ShouldEqual, 0)
		convey.So(fm.IsPublicMethod("missing"), convey.ShouldBeFalse)
		convey.So(fm.IsPrivateMethod("missing"), convey.ShouldBeFalse)
		convey.So(fm.GetType("missing"), convey.ShouldBeNil)
		convey.So(len(fm.GetOutArgsTypes("missing")), convey.ShouldEqual, 0)
		convey.So(len(fm.GetInArgsTypes("missing")), convey.ShouldEqual, 0)
		convey.So(len(fm.GetInArgsTypesNames("missing")), convey.ShouldEqual, 0)
	})
}

func Test_FuncMap_ValidateMethodArgs_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.ValidateMethodArgs("missing", []any{})

	// Assert
	convey.Convey("FuncMap.ValidateMethodArgs returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvokeFirstAndError_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	_, _, err := fm.InvokeFirstAndError("missing")

	// Assert
	convey.Convey("FuncMap.InvokeFirstAndError returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvalidError_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.InvalidError()

	// Assert
	convey.Convey("FuncMap.InvalidError returns error for empty map", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvalidErrorByName_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})

	// Act
	err := fm.InvalidErrorByName("missing")

	// Assert
	convey.Convey("FuncMap.InvalidErrorByName returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvalidErrorByName_Valid(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	err := fm.InvalidErrorByName(name)

	// Assert
	convey.Convey("FuncMap.InvalidErrorByName returns nil for valid func", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_FuncMap_VoidCallNoReturn_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.VoidCallNoReturn("missing")

	// Assert
	convey.Convey("FuncMap.VoidCallNoReturn returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_MustBeValid_Panics(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	convey.Convey("FuncMap.MustBeValid panics for not-found", t, func() {
		convey.So(func() { fm.MustBeValid("missing") }, convey.ShouldPanic)
	})
}

func Test_FuncMap_ValidationError_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	err := fm.ValidationError("missing")

	// Assert
	convey.Convey("FuncMap.ValidationError returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvokeMust_Panics(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	convey.Convey("FuncMap.InvokeMust panics for not-found", t, func() {
		convey.So(func() { fm.InvokeMust("missing") }, convey.ShouldPanic)
	})
}

func Test_FuncMap_VoidCall(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() string { return "x" })
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	results, err := fm.VoidCall(name)

	// Assert
	convey.Convey("FuncMap.VoidCall invokes with no args", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(results), convey.ShouldEqual, 1)
	})
}

func Test_FuncMap_InvokeResultOfIndex_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	_, err := fm.InvokeResultOfIndex("missing", 0)

	// Assert
	convey.Convey("FuncMap.InvokeResultOfIndex returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InvokeError_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	_, err := fm.InvokeError("missing")

	// Assert
	convey.Convey("FuncMap.InvokeError returns error for not-found", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_VerifyInArgs_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	ok, err := fm.VerifyInArgs("missing", []any{})

	// Assert
	convey.Convey("FuncMap.VerifyInArgs returns error for not-found", t, func() {
		convey.So(ok, convey.ShouldBeFalse)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_VerifyOutArgs_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	ok, err := fm.VerifyOutArgs("missing", []any{})

	// Assert
	convey.Convey("FuncMap.VerifyOutArgs returns error for not-found", t, func() {
		convey.So(ok, convey.ShouldBeFalse)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_InArgsVerifyRv_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	ok, err := fm.InArgsVerifyRv("missing", nil)

	// Assert
	convey.Convey("FuncMap.InArgsVerifyRv returns error for not-found", t, func() {
		convey.So(ok, convey.ShouldBeFalse)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_OutArgsVerifyRv_NotFound(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	ok, err := fm.OutArgsVerifyRv("missing", nil)

	// Assert
	convey.Convey("FuncMap.OutArgsVerifyRv returns error for not-found", t, func() {
		convey.So(ok, convey.ShouldBeFalse)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_FuncMap_GetPascalCaseFuncName(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})

	// Act
	name := ""
	for k := range fm {
		name = k
		break
	}
	result := fm.GetPascalCaseFuncName(name)

	// Assert
	convey.Convey("FuncMap.GetPascalCaseFuncName returns pascal name", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_FuncMap_GetPascalCaseFuncName_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	result := fm.GetPascalCaseFuncName("test")

	// Assert
	convey.Convey("FuncMap.GetPascalCaseFuncName empty returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_FuncMap_IsValidFuncOf(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act & Assert
	convey.Convey("FuncMap.IsValidFuncOf and IsInvalidFunc", t, func() {
		convey.So(fm.IsValidFuncOf(name), convey.ShouldBeTrue)
		convey.So(fm.IsInvalidFunc(name), convey.ShouldBeFalse)
		convey.So(fm.IsValidFuncOf("missing"), convey.ShouldBeFalse)
		convey.So(fm.IsInvalidFunc("missing"), convey.ShouldBeTrue)
	})
}

// --- String type ---

func Test_String_Methods(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act & Assert
	convey.Convey("String type methods work correctly", t, func() {
		convey.So(s.String(), convey.ShouldEqual, "hello world")
		convey.So(s.Length(), convey.ShouldEqual, 11)
		convey.So(s.Count(), convey.ShouldEqual, 11)
		convey.So(s.AscIILength(), convey.ShouldEqual, 11)
		convey.So(s.IsEmpty(), convey.ShouldBeFalse)
		convey.So(s.HasCharacter(), convey.ShouldBeTrue)
		convey.So(s.IsDefined(), convey.ShouldBeTrue)
		convey.So(s.IsEmptyOrWhitespace(), convey.ShouldBeFalse)
		convey.So(len(s.Bytes()), convey.ShouldBeGreaterThan, 0)
		convey.So(len(s.Runes()), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_String_Concat(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	result := s.Concat(" ", "world")

	// Assert
	convey.Convey("String.Concat appends strings", t, func() {
		convey.So(result.String(), convey.ShouldEqual, "hello world")
	})
}

func Test_String_Join(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	result := s.Join("-", "world", "test")

	// Assert
	convey.Convey("String.Join joins with separator", t, func() {
		convey.So(result.String(), convey.ShouldContainSubstring, "-")
	})
}

func Test_String_Split(t *testing.T) {
	// Arrange
	s := args.String("a-b-c")

	// Act
	result := s.Split("-")

	// Assert
	convey.Convey("String.Split splits by separator", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_String_Quote_Methods(t *testing.T) {
	// Arrange
	s := args.String("test")

	// Act & Assert
	convey.Convey("String quote methods return non-empty", t, func() {
		convey.So(s.DoubleQuote().String(), convey.ShouldNotBeEmpty)
		convey.So(s.DoubleQuoteQ().String(), convey.ShouldNotBeEmpty)
		convey.So(s.SingleQuote().String(), convey.ShouldNotBeEmpty)
		convey.So(s.ValueDoubleQuote().String(), convey.ShouldNotBeEmpty)
	})
}

func Test_String_TrimSpace(t *testing.T) {
	// Arrange
	s := args.String("  hello  ")

	// Act
	result := s.TrimSpace()

	// Assert
	convey.Convey("String.TrimSpace trims whitespace", t, func() {
		convey.So(result.String(), convey.ShouldEqual, "hello")
	})
}

func Test_String_ReplaceAll(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act
	result := s.ReplaceAll("world", "go")

	// Assert
	convey.Convey("String.ReplaceAll replaces all occurrences", t, func() {
		convey.So(result.String(), convey.ShouldEqual, "hello go")
	})
}

func Test_String_Substring(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act
	result := s.Substring(0, 5)

	// Assert
	convey.Convey("String.Substring returns substring", t, func() {
		convey.So(result.String(), convey.ShouldEqual, "hello")
	})
}

func Test_String_Empty_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	s := args.String("")

	// Act & Assert
	convey.Convey("Empty String methods", t, func() {
		convey.So(s.IsEmpty(), convey.ShouldBeTrue)
		convey.So(s.HasCharacter(), convey.ShouldBeFalse)
		convey.So(s.IsDefined(), convey.ShouldBeFalse)
		convey.So(s.IsEmptyOrWhitespace(), convey.ShouldBeTrue)
	})
}

func Test_String_Whitespace(t *testing.T) {
	// Arrange
	s := args.String("   ")

	// Act & Assert
	convey.Convey("Whitespace String is empty or whitespace", t, func() {
		convey.So(s.IsEmptyOrWhitespace(), convey.ShouldBeTrue)
	})
}

// --- Dynamic methods ---

func Test_Dynamic_Getters(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
		},
		Expect: "expected",
	}

	// Act & Assert
	convey.Convey("Dynamic getter methods return correct values", t, func() {
		convey.So(d.FirstItem(), convey.ShouldEqual, "a")
		convey.So(d.SecondItem(), convey.ShouldEqual, "b")
		convey.So(d.ThirdItem(), convey.ShouldEqual, "c")
		convey.So(d.FourthItem(), convey.ShouldEqual, "d")
		convey.So(d.FifthItem(), convey.ShouldEqual, "e")
		convey.So(d.SixthItem(), convey.ShouldEqual, "f")
		convey.So(d.Expected(), convey.ShouldEqual, "expected")
		convey.So(d.HasExpect(), convey.ShouldBeTrue)
		convey.So(d.ArgsCount(), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_Dynamic_HasDefined(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	convey.Convey("Dynamic.HasDefined and Has", t, func() {
		convey.So(d.HasDefined("key"), convey.ShouldBeTrue)
		convey.So(d.Has("key"), convey.ShouldBeTrue)
		convey.So(d.HasDefined("missing"), convey.ShouldBeFalse)
		convey.So(d.Has("missing"), convey.ShouldBeFalse)
	})
}

func Test_Dynamic_HasDefinedAll_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"a": "1",
			"b": "2",
		},
	}

	// Act & Assert
	convey.Convey("Dynamic.HasDefinedAll", t, func() {
		convey.So(d.HasDefinedAll("a", "b"), convey.ShouldBeTrue)
		convey.So(d.HasDefinedAll("a", "missing"), convey.ShouldBeFalse)
	})
}

func Test_Dynamic_IsKeyInvalid(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"key": "value",
		},
	}

	// Act & Assert
	convey.Convey("Dynamic.IsKeyInvalid and IsKeyMissing", t, func() {
		convey.So(d.IsKeyInvalid("missing"), convey.ShouldBeTrue)
		convey.So(d.IsKeyMissing("missing"), convey.ShouldBeTrue)
		convey.So(d.IsKeyInvalid("key"), convey.ShouldBeFalse)
	})
}

func Test_Dynamic_GetTyped(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"num":     42,
			"str":     "hello",
			"strings": []string{"a", "b"},
			"items":   []any{1, 2},
		},
	}

	// Act & Assert
	convey.Convey("Dynamic typed getters", t, func() {
		num, ok := d.GetAsInt("num")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(num, convey.ShouldEqual, 42)

		def := d.GetAsIntDefault("missing", 99)
		convey.So(def, convey.ShouldEqual, 99)

		str, ok := d.GetAsString("str")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(str, convey.ShouldEqual, "hello")

		strDef := d.GetAsStringDefault("missing")
		convey.So(strDef, convey.ShouldBeEmpty)

		strs, ok := d.GetAsStrings("strings")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(strs), convey.ShouldEqual, 2)

		items, ok := d.GetAsAnyItems("items")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(items), convey.ShouldEqual, 2)
	})
}

func Test_Dynamic_GetLowerCase_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"actual":  "val",
			"arrange": "setup",
		},
	}

	// Act & Assert
	convey.Convey("Dynamic.GetLowerCase and directLower", t, func() {
		v, ok := d.GetLowerCase("ACTUAL")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(v, convey.ShouldEqual, "val")

		convey.So(d.Actual(), convey.ShouldEqual, "val")
		convey.So(d.Arrange(), convey.ShouldEqual, "setup")
		convey.So(d.GetDirectLower("MISSING"), convey.ShouldBeNil)
	})
}

func Test_Dynamic_Slice_And_String(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"key": "value",
		},
		Expect: "expected",
	}

	// Act
	slice := d.Slice()
	str := d.String()
	_ = d.Slice() // cached

	// Assert
	convey.Convey("Dynamic.Slice and String work", t, func() {
		convey.So(len(slice), convey.ShouldBeGreaterThan, 0)
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_Dynamic_ValidArgs_And_Args(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"first": "a",
		},
	}

	// Act
	valid := d.ValidArgs()
	named := d.Args("first")

	// Assert
	convey.Convey("Dynamic.ValidArgs and Args", t, func() {
		convey.So(len(valid), convey.ShouldBeGreaterThan, 0)
		convey.So(len(named), convey.ShouldEqual, 1)
	})
}

func Test_Dynamic_GetByIndex_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"first": "a",
		},
	}

	// Act
	result := d.GetByIndex(0)

	// Assert
	convey.Convey("Dynamic.GetByIndex returns value", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_Dynamic_HasFunc(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"func": func() {},
		},
	}

	// Act & Assert
	convey.Convey("Dynamic.HasFunc returns true", t, func() {
		convey.So(d.HasFunc(), convey.ShouldBeTrue)
		convey.So(d.GetFuncName(), convey.ShouldNotBeEmpty)
	})
}

func Test_Dynamic_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	d := args.Dynamic[string]{
		Params: args.Map{},
	}

	// Act & Assert
	convey.Convey("Dynamic interface casts", t, func() {
		convey.So(d.AsArgsMapper(), convey.ShouldNotBeNil)
		convey.So(d.AsArgFuncNameContractsBinder(), convey.ShouldNotBeNil)
		convey.So(d.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- DynamicFunc methods ---

func Test_DynamicFunc_Getters(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"first":  "a",
			"second": "b",
			"third":  "c",
			"fourth": "d",
			"fifth":  "e",
			"sixth":  "f",
			"when":   "scenario",
			"title":  "test-title",
		},
		WorkFunc: func() {},
		Expect:   "expected",
	}

	// Act & Assert
	convey.Convey("DynamicFunc getter methods", t, func() {
		convey.So(df.FirstItem(), convey.ShouldEqual, "a")
		convey.So(df.SecondItem(), convey.ShouldEqual, "b")
		convey.So(df.ThirdItem(), convey.ShouldEqual, "c")
		convey.So(df.FourthItem(), convey.ShouldEqual, "d")
		convey.So(df.FifthItem(), convey.ShouldEqual, "e")
		convey.So(df.SixthItem(), convey.ShouldEqual, "f")
		convey.So(df.Expected(), convey.ShouldEqual, "expected")
		convey.So(df.When(), convey.ShouldEqual, "scenario")
		convey.So(df.Title(), convey.ShouldEqual, "test-title")
		convey.So(df.HasExpect(), convey.ShouldBeTrue)
		convey.So(df.HasFunc(), convey.ShouldBeTrue)
		convey.So(df.GetFuncName(), convey.ShouldNotBeEmpty)
		convey.So(df.Length(), convey.ShouldBeGreaterThan, 0)
		convey.So(df.ArgsCount(), convey.ShouldBeGreaterThan, 0)
		convey.So(df.GetWorkFunc(), convey.ShouldNotBeNil)
	})
}

func Test_DynamicFunc_HasFirst_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"first": "a",
		},
	}

	// Act & Assert
	convey.Convey("DynamicFunc.HasFirst", t, func() {
		convey.So(df.HasFirst(), convey.ShouldBeTrue)
	})
}

func Test_DynamicFunc_GetByIndex_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"key": "value",
		},
		WorkFunc: func() {},
	}

	// Act
	result := df.GetByIndex(0)
	outOfRange := df.GetByIndex(999)

	// Assert
	convey.Convey("DynamicFunc.GetByIndex", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(outOfRange, convey.ShouldBeNil)
	})
}

func Test_DynamicFunc_TypedGetters(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"num":     42,
			"str":     "hello",
			"strings": []string{"a"},
			"items":   []any{1},
		},
	}

	// Act & Assert
	convey.Convey("DynamicFunc typed getters", t, func() {
		num, ok := df.GetAsInt("num")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(num, convey.ShouldEqual, 42)

		str, ok := df.GetAsString("str")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(str, convey.ShouldEqual, "hello")

		strs, ok := df.GetAsStrings("strings")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(strs), convey.ShouldEqual, 1)

		items, ok := df.GetAsAnyItems("items")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(items), convey.ShouldEqual, 1)
	})
}

func Test_DynamicFunc_HasDefined(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"key": "val",
		},
	}

	// Act & Assert
	convey.Convey("DynamicFunc.HasDefined, Has, HasDefinedAll, IsKeyInvalid, IsKeyMissing", t, func() {
		convey.So(df.HasDefined("key"), convey.ShouldBeTrue)
		convey.So(df.Has("key"), convey.ShouldBeTrue)
		convey.So(df.HasDefinedAll("key"), convey.ShouldBeTrue)
		convey.So(df.IsKeyInvalid("missing"), convey.ShouldBeTrue)
		convey.So(df.IsKeyMissing("missing"), convey.ShouldBeTrue)
	})
}

func Test_DynamicFunc_GetLowerCase_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	df := args.DynamicFunc[func()]{
		Params: args.Map{
			"actual":  "val",
			"arrange": "setup",
		},
	}

	// Act & Assert
	convey.Convey("DynamicFunc lower-case getters", t, func() {
		v, ok := df.GetLowerCase("ACTUAL")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(v, convey.ShouldEqual, "val")

		convey.So(df.Actual(), convey.ShouldEqual, "val")
		convey.So(df.Arrange(), convey.ShouldEqual, "setup")
		convey.So(df.GetDirectLower("MISSING"), convey.ShouldBeNil)
	})
}

func Test_DynamicFunc_Slice_And_String(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"key": "value",
		},
		WorkFunc: func() {},
		Expect:   "expected",
	}

	// Act
	slice := df.Slice()
	str := df.String()
	_ = df.Slice() // cached

	// Assert
	convey.Convey("DynamicFunc.Slice and String", t, func() {
		convey.So(len(slice), convey.ShouldBeGreaterThan, 0)
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_DynamicFunc_ValidArgs_And_Args(t *testing.T) {
	// Arrange
	df := &args.DynamicFunc[func()]{
		Params: args.Map{
			"first": "a",
		},
	}

	// Act & Assert
	convey.Convey("DynamicFunc.ValidArgs and Args", t, func() {
		convey.So(len(df.ValidArgs()), convey.ShouldBeGreaterThan, 0)
		convey.So(len(df.Args("first")), convey.ShouldEqual, 1)
	})
}

func Test_DynamicFunc_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	df := args.DynamicFunc[func()]{
		Params: args.Map{},
	}

	// Act & Assert
	convey.Convey("DynamicFunc interface casts", t, func() {
		convey.So(df.AsArgsMapper(), convey.ShouldNotBeNil)
		convey.So(df.AsArgFuncNameContractsBinder(), convey.ShouldNotBeNil)
		convey.So(df.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- LeftRight methods ---

func Test_LeftRight_Methods(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{
		Left:   "hello",
		Right:  42,
		Expect: true,
	}

	// Act & Assert
	convey.Convey("LeftRight methods work correctly", t, func() {
		convey.So(lr.ArgsCount(), convey.ShouldEqual, 2)
		convey.So(lr.FirstItem(), convey.ShouldEqual, "hello")
		convey.So(lr.SecondItem(), convey.ShouldEqual, 42)
		convey.So(lr.Expected(), convey.ShouldEqual, true)
		convey.So(lr.HasFirst(), convey.ShouldBeTrue)
		convey.So(lr.HasSecond(), convey.ShouldBeTrue)
		convey.So(lr.HasLeft(), convey.ShouldBeTrue)
		convey.So(lr.HasRight(), convey.ShouldBeTrue)
		convey.So(lr.HasExpect(), convey.ShouldBeTrue)
	})
}

func Test_LeftRight_Clone(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{
		Left:   "a",
		Right:  1,
		Expect: "x",
	}

	// Act
	cloned := lr.Clone()

	// Assert
	convey.Convey("LeftRight.Clone returns independent copy", t, func() {
		convey.So(cloned.Left, convey.ShouldEqual, "a")
		convey.So(cloned.Right, convey.ShouldEqual, 1)
	})
}

func Test_LeftRight_ArgTwo(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{
		Left:  "a",
		Right: 1,
	}

	// Act
	two := lr.ArgTwo()

	// Assert
	convey.Convey("LeftRight.ArgTwo returns TwoFunc", t, func() {
		convey.So(two.First, convey.ShouldEqual, "a")
		convey.So(two.Second, convey.ShouldEqual, 1)
	})
}

func Test_LeftRight_ValidArgs_And_Args(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{
		Left:  "a",
		Right: 1,
	}

	// Act & Assert
	convey.Convey("LeftRight.ValidArgs and Args", t, func() {
		convey.So(len(lr.ValidArgs()), convey.ShouldEqual, 2)
		convey.So(len(lr.Args(1)), convey.ShouldEqual, 1)
		convey.So(len(lr.Args(2)), convey.ShouldEqual, 2)
	})
}

func Test_LeftRight_Slice_GetByIndex_String(t *testing.T) {
	// Arrange
	lr := &args.LeftRight[string, int]{
		Left:  "a",
		Right: 1,
	}

	// Act
	slice := lr.Slice()
	idx0 := lr.GetByIndex(0)
	str := lr.String()
	_ = lr.Slice() // cached

	// Assert
	convey.Convey("LeftRight.Slice, GetByIndex, String", t, func() {
		convey.So(len(slice), convey.ShouldBeGreaterThan, 0)
		convey.So(idx0, convey.ShouldNotBeNil)
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_LeftRight_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	lr := args.LeftRight[string, int]{
		Left:  "a",
		Right: 1,
	}

	// Act & Assert
	convey.Convey("LeftRight interface casts", t, func() {
		convey.So(lr.AsTwoParameter(), convey.ShouldNotBeNil)
		convey.So(lr.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- One methods ---

func Test_One_Methods(t *testing.T) {
	// Arrange
	one := &args.One[string]{
		First:  "hello",
		Expect: 42,
	}

	// Act & Assert
	convey.Convey("One methods", t, func() {
		convey.So(one.FirstItem(), convey.ShouldEqual, "hello")
		convey.So(one.Expected(), convey.ShouldEqual, 42)
		convey.So(one.HasFirst(), convey.ShouldBeTrue)
		convey.So(one.HasExpect(), convey.ShouldBeTrue)
		convey.So(one.ArgsCount(), convey.ShouldEqual, 1)
	})
}

func Test_One_ArgTwo(t *testing.T) {
	// Arrange
	one := &args.One[string]{
		First:  "hello",
		Expect: 42,
	}

	// Act
	result := one.ArgTwo()

	// Assert
	convey.Convey("One.ArgTwo returns One copy", t, func() {
		convey.So(result.First, convey.ShouldEqual, "hello")
	})
}

func Test_One_LeftRight(t *testing.T) {
	// Arrange
	one := &args.One[string]{
		First:  "hello",
		Expect: 42,
	}

	// Act
	lr := one.LeftRight()

	// Assert
	convey.Convey("One.LeftRight returns LeftRight", t, func() {
		convey.So(lr.Left, convey.ShouldEqual, "hello")
	})
}

func Test_One_Slice_GetByIndex_String(t *testing.T) {
	// Arrange
	one := &args.One[string]{
		First: "hello",
	}

	// Act
	slice := one.Slice()
	idx := one.GetByIndex(0)
	str := one.String()
	_ = one.Slice() // cached

	// Assert
	convey.Convey("One.Slice, GetByIndex, String", t, func() {
		convey.So(len(slice), convey.ShouldBeGreaterThan, 0)
		convey.So(idx, convey.ShouldNotBeNil)
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_One_ValidArgs_And_Args(t *testing.T) {
	// Arrange
	one := &args.One[string]{
		First: "hello",
	}

	// Act & Assert
	convey.Convey("One.ValidArgs and Args", t, func() {
		convey.So(len(one.ValidArgs()), convey.ShouldEqual, 1)
		convey.So(len(one.Args(1)), convey.ShouldEqual, 1)
		convey.So(len(one.Args(0)), convey.ShouldEqual, 0)
	})
}

func Test_One_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	one := args.One[string]{
		First: "hello",
	}

	// Act & Assert
	convey.Convey("One interface casts", t, func() {
		convey.So(one.AsOneParameter(), convey.ShouldNotBeNil)
		convey.So(one.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Two methods ---

func Test_Two_Methods(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
		Expect: true,
	}

	// Act & Assert
	convey.Convey("Two methods", t, func() {
		convey.So(two.FirstItem(), convey.ShouldEqual, "a")
		convey.So(two.SecondItem(), convey.ShouldEqual, 1)
		convey.So(two.Expected(), convey.ShouldEqual, true)
		convey.So(two.HasFirst(), convey.ShouldBeTrue)
		convey.So(two.HasSecond(), convey.ShouldBeTrue)
		convey.So(two.HasExpect(), convey.ShouldBeTrue)
		convey.So(two.ArgsCount(), convey.ShouldEqual, 2)
	})
}

func Test_Two_ArgTwo(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act
	result := two.ArgTwo()

	// Assert
	convey.Convey("Two.ArgTwo returns TwoFunc", t, func() {
		convey.So(result.First, convey.ShouldEqual, "a")
	})
}

func Test_Two_LeftRight(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act
	lr := two.LeftRight()

	// Assert
	convey.Convey("Two.LeftRight returns LeftRight", t, func() {
		convey.So(lr.Left, convey.ShouldEqual, "a")
	})
}

func Test_Two_Slice_String(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act
	str := two.String()
	_ = two.Slice() // cached

	// Assert
	convey.Convey("Two.String", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_Two_ValidArgs_And_Args(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act & Assert
	convey.Convey("Two.ValidArgs and Args", t, func() {
		convey.So(len(two.ValidArgs()), convey.ShouldEqual, 2)
		convey.So(len(two.Args(1)), convey.ShouldEqual, 1)
		convey.So(len(two.Args(2)), convey.ShouldEqual, 2)
	})
}

func Test_Two_GetByIndex(t *testing.T) {
	// Arrange
	two := &args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act & Assert
	convey.Convey("Two.GetByIndex", t, func() {
		convey.So(two.GetByIndex(0), convey.ShouldNotBeNil)
		convey.So(two.GetByIndex(99), convey.ShouldBeNil)
	})
}

func Test_Two_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	two := args.Two[string, int]{
		First:  "a",
		Second: 1,
	}

	// Act & Assert
	convey.Convey("Two interface casts", t, func() {
		convey.So(two.AsTwoParameter(), convey.ShouldNotBeNil)
		convey.So(two.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Three methods ---

func Test_Three_Methods(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
		Third:  true,
		Expect: "x",
	}

	// Act & Assert
	convey.Convey("Three methods", t, func() {
		convey.So(three.ArgsCount(), convey.ShouldEqual, 3)
		convey.So(three.FirstItem(), convey.ShouldEqual, "a")
		convey.So(three.SecondItem(), convey.ShouldEqual, 1)
		convey.So(three.ThirdItem(), convey.ShouldEqual, true)
		convey.So(three.Expected(), convey.ShouldEqual, "x")
		convey.So(three.HasFirst(), convey.ShouldBeTrue)
		convey.So(three.HasSecond(), convey.ShouldBeTrue)
		convey.So(three.HasThird(), convey.ShouldBeTrue)
		convey.So(three.HasExpect(), convey.ShouldBeTrue)
	})
}

func Test_Three_ArgTwo_ArgThree(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
		Third:  true,
	}

	// Act & Assert
	convey.Convey("Three.ArgTwo and ArgThree", t, func() {
		two := three.ArgTwo()
		convey.So(two.First, convey.ShouldEqual, "a")

		three2 := three.ArgThree()
		convey.So(three2.Third, convey.ShouldEqual, true)
	})
}

func Test_Three_LeftRight(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
	}

	// Act
	lr := three.LeftRight()

	// Assert
	convey.Convey("Three.LeftRight", t, func() {
		convey.So(lr.Left, convey.ShouldEqual, "a")
	})
}

func Test_Three_Slice_String(t *testing.T) {
	// Arrange
	three := args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
		Third:  true,
	}

	// Act
	str := three.String()
	_ = three.Slice() // cached

	// Assert
	convey.Convey("Three.String works", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_Three_ValidArgs_Args(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
		Third:  true,
	}

	// Act & Assert
	convey.Convey("Three.ValidArgs and Args", t, func() {
		convey.So(len(three.ValidArgs()), convey.ShouldEqual, 3)
		convey.So(len(three.Args(2)), convey.ShouldEqual, 2)
		convey.So(len(three.Args(3)), convey.ShouldEqual, 3)
	})
}

func Test_Three_GetByIndex(t *testing.T) {
	// Arrange
	three := &args.Three[string, int, bool]{
		First:  "a",
		Second: 1,
		Third:  true,
	}

	// Act & Assert
	convey.Convey("Three.GetByIndex", t, func() {
		convey.So(three.GetByIndex(0), convey.ShouldNotBeNil)
		convey.So(three.GetByIndex(99), convey.ShouldBeNil)
	})
}

func Test_Three_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	three := args.Three[string, int, bool]{
		First: "a",
	}

	// Act & Assert
	convey.Convey("Three interface casts", t, func() {
		convey.So(three.AsThreeParameter(), convey.ShouldNotBeNil)
		convey.So(three.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Four methods ---

func Test_Four_Methods(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Expect: "x",
	}

	// Act & Assert
	convey.Convey("Four methods", t, func() {
		convey.So(four.ArgsCount(), convey.ShouldEqual, 4)
		convey.So(four.FirstItem(), convey.ShouldEqual, "a")
		convey.So(four.SecondItem(), convey.ShouldEqual, 1)
		convey.So(four.ThirdItem(), convey.ShouldEqual, true)
		convey.So(four.FourthItem(), convey.ShouldAlmostEqual, 3.14)
		convey.So(four.Expected(), convey.ShouldEqual, "x")
		convey.So(four.HasFirst(), convey.ShouldBeTrue)
		convey.So(four.HasSecond(), convey.ShouldBeTrue)
		convey.So(four.HasThird(), convey.ShouldBeTrue)
		convey.So(four.HasFourth(), convey.ShouldBeTrue)
		convey.So(four.HasExpect(), convey.ShouldBeTrue)
	})
}

func Test_Four_ArgTwo_ArgThree(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
	}

	// Act & Assert
	convey.Convey("Four.ArgTwo and ArgThree", t, func() {
		two := four.ArgTwo()
		convey.So(two.First, convey.ShouldEqual, "a")

		three := four.ArgThree()
		convey.So(three.Third, convey.ShouldEqual, true)
	})
}

func Test_Four_ValidArgs_Args(t *testing.T) {
	// Arrange
	four := &args.Four[string, int, bool, float64]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
	}

	// Act & Assert
	convey.Convey("Four.ValidArgs and Args", t, func() {
		convey.So(len(four.ValidArgs()), convey.ShouldEqual, 4)
		convey.So(len(four.Args(3)), convey.ShouldEqual, 3)
		convey.So(len(four.Args(4)), convey.ShouldEqual, 4)
	})
}

func Test_Four_Slice_String_GetByIndex(t *testing.T) {
	// Arrange
	four := args.Four[string, int, bool, float64]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
	}

	// Act
	str := four.String()
	_ = four.Slice() // cached

	// Assert
	convey.Convey("Four.String and cached Slice", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
		convey.So(four.GetByIndex(0), convey.ShouldNotBeNil)
	})
}

func Test_Four_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	four := args.Four[string, int, bool, float64]{
		First: "a",
	}

	// Act & Assert
	convey.Convey("Four interface casts", t, func() {
		convey.So(four.AsFourParameter(), convey.ShouldNotBeNil)
		convey.So(four.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Five methods ---

func Test_Five_Methods(t *testing.T) {
	// Arrange
	five := &args.Five[string, int, bool, float64, byte]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
		Expect: "x",
	}

	// Act & Assert
	convey.Convey("Five methods", t, func() {
		convey.So(five.ArgsCount(), convey.ShouldEqual, 5)
		convey.So(five.FirstItem(), convey.ShouldEqual, "a")
		convey.So(five.SecondItem(), convey.ShouldEqual, 1)
		convey.So(five.ThirdItem(), convey.ShouldEqual, true)
		convey.So(five.FourthItem(), convey.ShouldAlmostEqual, 3.14)
		convey.So(five.FifthItem(), convey.ShouldEqual, byte(5))
		convey.So(five.Expected(), convey.ShouldEqual, "x")
		convey.So(five.HasFirst(), convey.ShouldBeTrue)
		convey.So(five.HasSecond(), convey.ShouldBeTrue)
		convey.So(five.HasThird(), convey.ShouldBeTrue)
		convey.So(five.HasFourth(), convey.ShouldBeTrue)
		convey.So(five.HasFifth(), convey.ShouldBeTrue)
		convey.So(five.HasExpect(), convey.ShouldBeTrue)
	})
}

func Test_Five_ArgTwo_ArgThree_ArgFour(t *testing.T) {
	// Arrange
	five := &args.Five[string, int, bool, float64, byte]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
	}

	// Act & Assert
	convey.Convey("Five.ArgTwo, ArgThree, ArgFour", t, func() {
		convey.So(five.ArgTwo().First, convey.ShouldEqual, "a")
		convey.So(five.ArgThree().Third, convey.ShouldEqual, true)
		convey.So(five.ArgFour().Fourth, convey.ShouldAlmostEqual, 3.14)
	})
}

func Test_Five_ValidArgs_Args(t *testing.T) {
	// Arrange
	five := &args.Five[string, int, bool, float64, byte]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
	}

	// Act & Assert
	convey.Convey("Five.ValidArgs and Args", t, func() {
		convey.So(len(five.ValidArgs()), convey.ShouldEqual, 5)
		convey.So(len(five.Args(4)), convey.ShouldEqual, 4)
		convey.So(len(five.Args(5)), convey.ShouldEqual, 5)
	})
}

func Test_Five_Slice_String_GetByIndex(t *testing.T) {
	// Arrange
	five := args.Five[string, int, bool, float64, byte]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
	}

	// Act
	str := five.String()
	_ = five.Slice() // cached

	// Assert
	convey.Convey("Five.String and cached Slice", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
		convey.So(five.GetByIndex(0), convey.ShouldNotBeNil)
	})
}

func Test_Five_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	five := args.Five[string, int, bool, float64, byte]{
		First: "a",
	}

	// Act & Assert
	convey.Convey("Five interface casts", t, func() {
		convey.So(five.AsFifthParameter(), convey.ShouldNotBeNil)
		convey.So(five.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Six methods ---

func Test_Six_Methods(t *testing.T) {
	// Arrange
	six := &args.Six[string, int, bool, float64, byte, rune]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
		Sixth:  'x',
		Expect: "e",
	}

	// Act & Assert
	convey.Convey("Six methods", t, func() {
		convey.So(six.ArgsCount(), convey.ShouldEqual, 6)
		convey.So(six.FirstItem(), convey.ShouldEqual, "a")
		convey.So(six.SecondItem(), convey.ShouldEqual, 1)
		convey.So(six.ThirdItem(), convey.ShouldEqual, true)
		convey.So(six.FourthItem(), convey.ShouldAlmostEqual, 3.14)
		convey.So(six.FifthItem(), convey.ShouldEqual, byte(5))
		convey.So(six.SixthItem(), convey.ShouldEqual, 'x')
		convey.So(six.Expected(), convey.ShouldEqual, "e")
		convey.So(six.HasFirst(), convey.ShouldBeTrue)
		convey.So(six.HasSecond(), convey.ShouldBeTrue)
		convey.So(six.HasThird(), convey.ShouldBeTrue)
		convey.So(six.HasFourth(), convey.ShouldBeTrue)
		convey.So(six.HasFifth(), convey.ShouldBeTrue)
		convey.So(six.HasSixth(), convey.ShouldBeTrue)
		convey.So(six.HasExpect(), convey.ShouldBeTrue)
	})
}

func Test_Six_ArgTwo_Three_Four_Five(t *testing.T) {
	// Arrange
	six := &args.Six[string, int, bool, float64, byte, rune]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
		Sixth:  'x',
	}

	// Act & Assert
	convey.Convey("Six.ArgTwo through ArgFive", t, func() {
		convey.So(six.ArgTwo().First, convey.ShouldEqual, "a")
		convey.So(six.ArgThree().Third, convey.ShouldEqual, true)
		convey.So(six.ArgFour().Fourth, convey.ShouldAlmostEqual, 3.14)
		convey.So(six.ArgFive().Fifth, convey.ShouldEqual, byte(5))
	})
}

func Test_Six_ValidArgs_Args(t *testing.T) {
	// Arrange
	six := &args.Six[string, int, bool, float64, byte, rune]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
		Sixth:  'x',
	}

	// Act & Assert
	convey.Convey("Six.ValidArgs and Args", t, func() {
		convey.So(len(six.ValidArgs()), convey.ShouldEqual, 6)
		convey.So(len(six.Args(5)), convey.ShouldEqual, 5)
		convey.So(len(six.Args(6)), convey.ShouldEqual, 6)
	})
}

func Test_Six_Slice_String_GetByIndex(t *testing.T) {
	// Arrange
	six := args.Six[string, int, bool, float64, byte, rune]{
		First:  "a",
		Second: 1,
		Third:  true,
		Fourth: 3.14,
		Fifth:  byte(5),
		Sixth:  'x',
	}

	// Act
	str := six.String()
	_ = six.Slice() // cached

	// Assert
	convey.Convey("Six.String and cached Slice", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
		convey.So(six.GetByIndex(0), convey.ShouldNotBeNil)
	})
}

func Test_Six_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	six := args.Six[string, int, bool, float64, byte, rune]{
		First: "a",
	}

	// Act & Assert
	convey.Convey("Six interface casts", t, func() {
		convey.So(six.AsSixthParameter(), convey.ShouldNotBeNil)
		convey.So(six.AsArgBaseContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Holder methods ---

func Test_Holder_Methods(t *testing.T) {
	// Arrange
	h := &args.Holder[func()]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		Fourth:   "d",
		Fifth:    "e",
		Sixth:    "f",
		WorkFunc: func() {},
		Expect:   "expected",
	}

	// Act & Assert
	convey.Convey("Holder methods", t, func() {
		convey.So(h.ArgsCount(), convey.ShouldEqual, 7)
		convey.So(h.FirstItem(), convey.ShouldEqual, "a")
		convey.So(h.SecondItem(), convey.ShouldEqual, "b")
		convey.So(h.ThirdItem(), convey.ShouldEqual, "c")
		convey.So(h.FourthItem(), convey.ShouldEqual, "d")
		convey.So(h.FifthItem(), convey.ShouldEqual, "e")
		convey.So(h.SixthItem(), convey.ShouldEqual, "f")
		convey.So(h.Expected(), convey.ShouldEqual, "expected")
		convey.So(h.HasFirst(), convey.ShouldBeTrue)
		convey.So(h.HasSecond(), convey.ShouldBeTrue)
		convey.So(h.HasThird(), convey.ShouldBeTrue)
		convey.So(h.HasFourth(), convey.ShouldBeTrue)
		convey.So(h.HasFifth(), convey.ShouldBeTrue)
		convey.So(h.HasSixth(), convey.ShouldBeTrue)
		convey.So(h.HasFunc(), convey.ShouldBeTrue)
		convey.So(h.HasExpect(), convey.ShouldBeTrue)
		convey.So(h.GetFuncName(), convey.ShouldNotBeEmpty)
		convey.So(h.GetWorkFunc(), convey.ShouldNotBeNil)
	})
}

func Test_Holder_ArgTwo_Three_Four_Five(t *testing.T) {
	// Arrange
	h := &args.Holder[func()]{
		First:  "a",
		Second: "b",
		Third:  "c",
		Fourth: "d",
		Fifth:  "e",
	}

	// Act & Assert
	convey.Convey("Holder.ArgTwo through ArgFive", t, func() {
		convey.So(h.ArgTwo().First, convey.ShouldEqual, "a")
		convey.So(h.ArgThree().Third, convey.ShouldEqual, "c")
		convey.So(h.ArgFour().Fourth, convey.ShouldEqual, "d")
		convey.So(h.ArgFive().Fifth, convey.ShouldEqual, "e")
	})
}

func Test_Holder_ValidArgs_Args(t *testing.T) {
	// Arrange
	h := &args.Holder[func()]{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act & Assert
	convey.Convey("Holder.ValidArgs and Args", t, func() {
		convey.So(len(h.ValidArgs()), convey.ShouldEqual, 3)
		convey.So(len(h.Args(2)), convey.ShouldEqual, 2)
		convey.So(len(h.Args(6)), convey.ShouldEqual, 6)
	})
}

func Test_Holder_Slice_GetByIndex_String(t *testing.T) {
	// Arrange
	h := &args.Holder[func()]{
		First: "a",
	}

	// Act
	str := h.String()
	_ = h.Slice() // cached
	idx := h.GetByIndex(0)

	// Assert
	convey.Convey("Holder.String, Slice, GetByIndex", t, func() {
		convey.So(str, convey.ShouldNotBeEmpty)
		convey.So(idx, convey.ShouldNotBeNil)
	})
}

func Test_Holder_AsInterfaces_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	h := args.Holder[func()]{
		First: "a",
	}

	// Act & Assert
	convey.Convey("Holder interface casts", t, func() {
		convey.So(h.AsSixthParameter(), convey.ShouldNotBeNil)
		convey.So(h.AsArgFuncContractsBinder(), convey.ShouldNotBeNil)
	})
}

// --- Map methods ---

func Test_Map_Getters(t *testing.T) {
	// Arrange
	m := args.Map{
		"when":    "scenario",
		"title":   "test-title",
		"actual":  "val",
		"arrange": "setup",
		"seventh": "7th",
	}

	// Act & Assert
	convey.Convey("Map getters", t, func() {
		convey.So(m.When(), convey.ShouldEqual, "scenario")
		convey.So(m.Title(), convey.ShouldEqual, "test-title")
		convey.So(m.Actual(), convey.ShouldEqual, "val")
		convey.So(m.Arrange(), convey.ShouldEqual, "setup")
		convey.So(m.Seventh(), convey.ShouldEqual, "7th")
		convey.So(m.Expect(), convey.ShouldBeNil)
	})
}

func Test_Map_SetActual_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	m.SetActual("new-val")

	// Assert
	convey.Convey("Map.SetActual sets value", t, func() {
		convey.So(m.Actual(), convey.ShouldEqual, "new-val")
	})
}

func Test_Map_GetLowerCase(t *testing.T) {
	// Arrange
	m := args.Map{
		"key": "val",
	}

	// Act
	v, ok := m.GetLowerCase("KEY")

	// Assert
	convey.Convey("Map.GetLowerCase", t, func() {
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(v, convey.ShouldEqual, "val")
	})
}

func Test_Map_GetDirectLower(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act & Assert
	convey.Convey("Map.GetDirectLower missing returns nil", t, func() {
		convey.So(m.GetDirectLower("MISSING"), convey.ShouldBeNil)
	})
}

func Test_Map_TypedGetters(t *testing.T) {
	// Arrange
	m := args.Map{
		"num":     42,
		"str":     "hello",
		"bool":    true,
		"strings": []string{"a"},
		"items":   []any{1},
	}

	// Act & Assert
	convey.Convey("Map typed getters", t, func() {
		num, ok := m.GetAsInt("num")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(num, convey.ShouldEqual, 42)

		convey.So(m.GetAsIntDefault("missing", 99), convey.ShouldEqual, 99)

		b, ok := m.GetAsBool("bool")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(b, convey.ShouldBeTrue)

		convey.So(m.GetAsBoolDefault("missing", true), convey.ShouldBeTrue)

		str, ok := m.GetAsString("str")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(str, convey.ShouldEqual, "hello")

		convey.So(m.GetAsStringDefault("missing"), convey.ShouldBeEmpty)

		strs, ok := m.GetAsStrings("strings")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(strs), convey.ShouldEqual, 1)

		items, ok := m.GetAsAnyItems("items")
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(len(items), convey.ShouldEqual, 1)
	})
}

func Test_Map_GetAsStringSliceFirstOfNames_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"items": []string{"a", "b"},
	}

	// Act
	result := m.GetAsStringSliceFirstOfNames("items")

	// Assert
	convey.Convey("Map.GetAsStringSliceFirstOfNames returns slice", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_Map_GetAsStringSliceFirstOfNames_Nil(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	result := m.GetAsStringSliceFirstOfNames("missing")

	// Assert
	convey.Convey("Map.GetAsStringSliceFirstOfNames nil for missing", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_Map_GetAsStringSliceFirstOfNames_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	result := m.GetAsStringSliceFirstOfNames()

	// Assert
	convey.Convey("Map.GetAsStringSliceFirstOfNames empty args returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_Map_WorkFuncName_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"func": func() {},
	}

	// Act
	name := m.WorkFuncName()

	// Assert
	convey.Convey("Map.WorkFuncName returns name", t, func() {
		convey.So(name, convey.ShouldNotBeEmpty)
	})
}

func Test_Map_GetFirstFuncNameOf_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"func": func() {},
	}

	// Act
	name := m.GetFirstFuncNameOf("func")

	// Assert
	convey.Convey("Map.GetFirstFuncNameOf returns func name", t, func() {
		convey.So(name, convey.ShouldNotBeEmpty)
	})
}

func Test_Map_Slice_String(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
	}

	// Act
	slice := m.Slice()
	str := m.String()

	// Assert
	convey.Convey("Map.Slice and String", t, func() {
		convey.So(len(slice), convey.ShouldEqual, 2)
		convey.So(str, convey.ShouldNotBeEmpty)
	})
}

func Test_Map_CompileToStrings_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
	}

	// Act
	lines := m.CompileToStrings()

	// Assert
	convey.Convey("Map.CompileToStrings returns sorted lines", t, func() {
		convey.So(len(lines), convey.ShouldEqual, 2)
	})
}

func Test_Map_CompileToStrings_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.CompileToStrings()

	// Assert
	convey.Convey("Map.CompileToStrings empty returns empty", t, func() {
		convey.So(len(lines), convey.ShouldEqual, 0)
	})
}

func Test_Map_CompileToString_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	result := m.CompileToString()

	// Assert
	convey.Convey("Map.CompileToString returns multi-line string", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "a : 1")
	})
}

func Test_Map_GoLiteralLines_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"name":  "hello",
		"value": 42,
	}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	convey.Convey("Map.GoLiteralLines returns Go literal format", t, func() {
		convey.So(len(lines), convey.ShouldEqual, 2)
	})
}

func Test_Map_GoLiteralLines_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	convey.Convey("Map.GoLiteralLines empty returns empty", t, func() {
		convey.So(len(lines), convey.ShouldEqual, 0)
	})
}

func Test_Map_GoLiteralString_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	m := args.Map{
		"key": "val",
	}

	// Act
	result := m.GoLiteralString()

	// Assert
	convey.Convey("Map.GoLiteralString returns formatted string", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

func Test_Map_NilReceiver(t *testing.T) {
	// Arrange
	var m args.Map

	// Act & Assert
	convey.Convey("Map nil-receiver methods", t, func() {
		convey.So(m.HasDefined("key"), convey.ShouldBeFalse)
		convey.So(m.Has("key"), convey.ShouldBeFalse)
		convey.So(m.HasDefinedAll("key"), convey.ShouldBeFalse)
		convey.So(m.IsKeyInvalid("key"), convey.ShouldBeFalse)
		convey.So(m.IsKeyMissing("key"), convey.ShouldBeFalse)
		v, ok := m.Get("key")
		convey.So(v, convey.ShouldBeNil)
		convey.So(ok, convey.ShouldBeFalse)
	})
}

// --- emptyCreator ---

func Test_EmptyCreator_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Act & Assert
	convey.Convey("Empty creator methods", t, func() {
		convey.So(args.Empty.Map(), convey.ShouldNotBeNil)
		convey.So(args.Empty.FuncWrap(), convey.ShouldNotBeNil)
		convey.So(args.Empty.FuncMap(), convey.ShouldNotBeNil)
		convey.So(fmt.Sprintf("%v", args.Empty.Holder()), convey.ShouldNotBeEmpty)
	})
}

// --- NewFuncWrap.Many ---

func Test_NewFuncWrap_Many(t *testing.T) {
	// Arrange
	f1 := func() {}
	f2 := func() string { return "" }

	// Act
	result := args.NewFuncWrap.Many(f1, f2)

	// Assert
	convey.Convey("NewFuncWrap.Many returns slice", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_NewFuncWrap_Many_Empty(t *testing.T) {
	// Act
	result := args.NewFuncWrap.Many()

	// Assert
	convey.Convey("NewFuncWrap.Many empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NewFuncWrap.Map ---

func Test_NewFuncWrap_Map_Empty(t *testing.T) {
	// Act
	result := args.NewFuncWrap.Map()

	// Assert
	convey.Convey("NewFuncWrap.Map empty returns empty map", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NewFuncWrap.Single ---

func Test_NewFuncWrap_Single(t *testing.T) {
	// Arrange
	f := func() {}

	// Act
	result := args.NewFuncWrap.Single(f)

	// Assert
	convey.Convey("NewFuncWrap.Single is alias for Default", t, func() {
		convey.So(result.IsValid(), convey.ShouldBeTrue)
	})
}

// --- NewFuncWrap.Invalid ---

func Test_NewFuncWrap_Invalid(t *testing.T) {
	// Act
	result := args.NewFuncWrap.Invalid()

	// Assert
	convey.Convey("NewFuncWrap.Invalid returns invalid wrap", t, func() {
		convey.So(result.IsInvalid(), convey.ShouldBeTrue)
	})
}

// --- FuncMap.Add, Adds ---

func Test_FuncMap_Add(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Add(func() {})

	// Assert
	convey.Convey("FuncMap.Add adds a function", t, func() {
		convey.So(fm.HasAnyItem(), convey.ShouldBeTrue)
	})
}

func Test_FuncMap_Adds_FromFuncWrapVoidCallNoRe(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Adds(func() {}, func() string { return "" })

	// Assert
	convey.Convey("FuncMap.Adds adds multiple functions", t, func() {
		convey.So(fm.Length(), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_FuncMap_Adds_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	fm.Adds()

	// Assert
	convey.Convey("FuncMap.Adds empty does nothing", t, func() {
		convey.So(fm.IsEmpty(), convey.ShouldBeTrue)
	})
}

// --- FuncMap basic accessors ---

func Test_FuncMap_BasicAccessors(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})

	// Act & Assert
	convey.Convey("FuncMap basic accessors", t, func() {
		convey.So(fm.IsEmpty(), convey.ShouldBeFalse)
		convey.So(fm.HasAnyItem(), convey.ShouldBeTrue)
		convey.So(fm.Length(), convey.ShouldEqual, 1)
		convey.So(fm.Count(), convey.ShouldEqual, 1)
	})
}

func Test_FuncMap_Has_IsContains(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act & Assert
	convey.Convey("FuncMap.Has and IsContains", t, func() {
		convey.So(fm.Has(name), convey.ShouldBeTrue)
		convey.So(fm.IsContains(name), convey.ShouldBeTrue)
		convey.So(fm.Has("missing"), convey.ShouldBeFalse)
	})
}

func Test_FuncMap_Get(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func() {})
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act & Assert
	convey.Convey("FuncMap.Get returns pointer or nil", t, func() {
		convey.So(fm.Get(name), convey.ShouldNotBeNil)
		convey.So(fm.Get("missing"), convey.ShouldBeNil)
	})
}

func Test_FuncMap_Has_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	convey.Convey("FuncMap.Has returns false for empty map", t, func() {
		convey.So(fm.Has("anything"), convey.ShouldBeFalse)
	})
}

func Test_FuncMap_Get_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act & Assert
	convey.Convey("FuncMap.Get returns nil for empty map", t, func() {
		convey.So(fm.Get("anything"), convey.ShouldBeNil)
	})
}

// --- FuncMap.ArgsLength (alias) ---

func Test_FuncMap_ArgsLength(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(func(a string) {})
	name := ""
	for k := range fm {
		name = k
		break
	}

	// Act
	result := fm.ArgsLength(name)

	// Assert
	convey.Convey("FuncMap.ArgsLength is alias for ArgsCount", t, func() {
		convey.So(result, convey.ShouldEqual, 1)
	})
}

// --- FuncWrapArgs aliases ---

func Test_FuncWrap_InArgsCount(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a string) {})

	// Act & Assert
	convey.Convey("InArgsCount is alias for ArgsCount", t, func() {
		convey.So(fw.InArgsCount(), convey.ShouldEqual, 1)
	})
}

func Test_FuncWrap_ArgsLength(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func(a string) {})

	// Act & Assert
	convey.Convey("ArgsLength is alias for ArgsCount", t, func() {
		convey.So(fw.ArgsLength(), convey.ShouldEqual, 1)
	})
}

// --- Map.Raw ---

func Test_Map_Raw(t *testing.T) {
	// Arrange
	m := args.Map{
		"key": "val",
	}

	// Act
	raw := m.Raw()

	// Assert
	convey.Convey("Map.Raw returns underlying map", t, func() {
		convey.So(raw["key"], convey.ShouldEqual, "val")
	})
}

// --- NewTypedFuncWrap ---

func Test_NewTypedFuncWrap(t *testing.T) {
	// Arrange
	fn := func(s string) int { return len(s) }

	// Act
	fw := args.NewTypedFuncWrap(fn)

	// Assert
	convey.Convey("NewTypedFuncWrap creates valid typed wrap", t, func() {
		convey.So(fw.IsValid(), convey.ShouldBeTrue)
		convey.So(fw.ArgsCount(), convey.ShouldEqual, 1)
	})
}

func Test_NewTypedFuncWrap_NonFunc(t *testing.T) {
	// Arrange
	notFunc := "not-a-func"

	// Act
	fw := args.NewTypedFuncWrap(notFunc)

	// Assert
	convey.Convey("NewTypedFuncWrap with non-func is invalid", t, func() {
		convey.So(fw.IsInvalid(), convey.ShouldBeTrue)
	})
}

// --- FuncDetector.GetFuncWrap ArgsMapper branch ---

func Test_FuncDetector_GetFuncWrap_ArgsMapper(t *testing.T) {
	// Arrange
	d := &args.Dynamic[string]{
		Params: args.Map{
			"func": func() {},
		},
	}

	// Act
	result := args.FuncDetector.GetFuncWrap(d)

	// Assert
	convey.Convey("FuncDetector.GetFuncWrap handles ArgsMapper", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}
