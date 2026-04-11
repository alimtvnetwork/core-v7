package regexnewtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── LazyRegex ──

func Test_LazyRegex_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{
		"isNull":      lr.IsNull(),
		"isDefined":   lr.IsDefined(),
		"isUndefined": lr.IsUndefined(),
		"isApplicable": lr.IsApplicable(),
		"isCompiled":  lr.IsCompiled(),
		"hasError":    lr.HasError(),
		"hasIssues":   lr.HasAnyIssues(),
		"isInvalid":   lr.IsInvalid(),
		"string":      lr.String(),
		"fullString":  lr.FullString(),
		"pattern":     lr.Pattern(),
	}

	// Assert
	expected := args.Map{
		"isNull": true, "isDefined": false, "isUndefined": true,
		"isApplicable": false, "isCompiled": false, "hasError": false,
		"hasIssues": true, "isInvalid": true,
		"string": "", "fullString": "", "pattern": "",
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns nil -- nil", actual)
}

func Test_LazyRegex_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Act
	actual := args.Map{
		"isDefined":    lr.IsDefined(),
		"isUndefined":  lr.IsUndefined(),
		"isApplicable": lr.IsApplicable(),
		"isCompiled":   lr.IsCompiled(),
		"hasError":     lr.HasError(),
		"hasIssues":    lr.HasAnyIssues(),
		"isInvalid":    lr.IsInvalid(),
		"pattern":      lr.Pattern(),
		"string":       lr.String(),
	}

	// Assert
	expected := args.Map{
		"isDefined": true, "isUndefined": false, "isApplicable": true,
		"isCompiled": true, "hasError": false, "hasIssues": false,
		"isInvalid": false, "pattern": `^\d+$`, "string": `^\d+$`,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns non-empty -- valid", actual)
}

func Test_LazyRegex_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)

	// Act
	actual := args.Map{
		"isApplicable": lr.IsApplicable(),
		"hasError":     lr.HasError(),
		"hasIssues":    lr.HasAnyIssues(),
		"isInvalid":    lr.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isApplicable": false, "hasError": true,
		"hasIssues": true, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- invalid", actual)
}

func Test_LazyRegex_Compile_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	r, err := lr.Compile()

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- Compile", actual)
}

func Test_LazyRegex_CompileMust_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	r := lr.CompileMust()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- CompileMust", actual)
}

func Test_LazyRegex_CompileMust_Panic_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex panics -- CompileMust panic", actual)
	}()
	lr.CompileMust()
}

func Test_LazyRegex_OnRequiredCompiled_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.OnRequiredCompiled()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- OnRequiredCompiled", actual)
}

func Test_LazyRegex_OnRequiredCompiled_Nil_FromLazyRegexNil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex
	err := lr.OnRequiredCompiled()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns nil -- OnRequiredCompiled nil", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	lr.OnRequiredCompiledMust() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- OnRequiredCompiledMust", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust_Panic(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex panics -- OnRequiredCompiledMust panic", actual)
	}()
	lr.OnRequiredCompiledMust()
}

func Test_LazyRegex_CompiledError_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.CompiledError()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- CompiledError", actual)
}

func Test_LazyRegex_Error_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)
	err := lr.Error()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- Error", actual)
}

func Test_LazyRegex_MustBeSafe_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	lr.MustBeSafe() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- MustBeSafe", actual)
}

func Test_LazyRegex_MustBeSafe_Panic(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex panics -- MustBeSafe panic", actual)
	}()
	lr.MustBeSafe()
}

func Test_LazyRegex_FullString_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	result := lr.FullString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- FullString", actual)
}

func Test_LazyRegex_MatchError_Success(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchError("123")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- MatchError success", actual)
}

func Test_LazyRegex_MatchError_Fail(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchError("abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- MatchError fail", actual)
}

func Test_LazyRegex_MatchUsingFuncError_Success(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchUsingFuncError("123", func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- MatchUsingFuncError success", actual)
}

func Test_LazyRegex_MatchUsingFuncError_Fail(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchUsingFuncError("abc", func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- MatchUsingFuncError fail", actual)
}

func Test_LazyRegex_IsMatch_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Act
	actual := args.Map{
		"match":   lr.IsMatch("123"),
		"noMatch": lr.IsMatch("abc"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsMatch", actual)
}

func Test_LazyRegex_IsMatchBytes_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Act
	actual := args.Map{
		"match":   lr.IsMatchBytes([]byte("123")),
		"noMatch": lr.IsMatchBytes([]byte("abc")),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsMatchBytes", actual)
}

func Test_LazyRegex_IsFailedMatch_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Act
	actual := args.Map{
		"failOnAlpha":  lr.IsFailedMatch("abc"),
		"failOnDigits": lr.IsFailedMatch("123"),
	}

	// Assert
	expected := args.Map{
		"failOnAlpha": true,
		"failOnDigits": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsFailedMatch", actual)
}

func Test_LazyRegex_IsFailedMatchBytes_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Act
	actual := args.Map{
		"failOnAlpha":  lr.IsFailedMatchBytes([]byte("abc")),
		"failOnDigits": lr.IsFailedMatchBytes([]byte("123")),
	}

	// Assert
	expected := args.Map{
		"failOnAlpha": true,
		"failOnDigits": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsFailedMatchBytes", actual)
}

func Test_LazyRegex_FirstMatchLine_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`\d+`)
	match, isInvalid := lr.FirstMatchLine("abc 123 def")

	// Act
	actual := args.Map{
		"match": match,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"match": "123",
		"isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- FirstMatchLine", actual)
}

func Test_LazyRegex_FirstMatchLine_NoMatch_FromLazyRegexNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`\d+`)
	match, isInvalid := lr.FirstMatchLine("abc def")

	// Act
	actual := args.Map{
		"match": match,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"match": "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns empty -- FirstMatchLine no match", actual)
}

func Test_LazyRegex_FirstMatchLine_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`[invalid`)
	match, isInvalid := lr.FirstMatchLine("abc")

	// Act
	actual := args.Map{
		"match": match,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"match": "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- FirstMatchLine invalid regex", actual)
}

// ── MatchUsingCustomizeErrorFuncLock with custom err ──

func Test_MatchCustomErr_WithCustom(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	customErr := func(pattern, match string, compileErr error, r *regexp.Regexp) error {
		return errors.New("custom: " + match)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, customErr)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"isCustom": err.Error() == "custom: abc",
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"isCustom": true,
	}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr returns error -- with custom func", actual)
}

// ── NewMustLock ──

func Test_NewMustLock_FromLazyRegexNil(t *testing.T) {
	// Arrange
	r := regexnew.NewMustLock(`^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock returns correct value -- with args", actual)
}

// ── Create / CreateLock / CreateMust ──

func Test_Create_Valid_FromLazyRegexNil(t *testing.T) {
	// Arrange
	r, err := regexnew.Create(`^[a-z]+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Create returns non-empty -- valid", actual)
}

func Test_Create_Invalid_FromLazyRegexNil(t *testing.T) {
	// Arrange
	_, err := regexnew.Create(`[invalid`)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid", actual)
}

func Test_CreateLock_Valid_FromLazyRegexNil(t *testing.T) {
	// Arrange
	r, err := regexnew.CreateLock(`^[A-Z]+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLock returns non-empty -- valid", actual)
}

func Test_CreateMust(t *testing.T) {
	// Arrange
	r := regexnew.CreateMust(`^test\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns correct value -- with args", actual)
}

func Test_IsMatchLock_Valid_FromLazyRegexNil(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   regexnew.IsMatchLock(`^\d+$`, "456"),
		"noMatch": regexnew.IsMatchLock(`^\d+$`, "abc"),
		"invalid": regexnew.IsMatchLock(`[invalid`, "abc"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns correct value -- with args", actual)
}

// ── New.Lazy creators ──

func Test_NewLazy_NoLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"isDefined": lr.IsDefined(),
		"isApplicable": lr.IsApplicable(),
	}

	// Assert
	expected := args.Map{
		"isDefined": true,
		"isApplicable": true,
	}
	expected.ShouldBeEqual(t, 0, "NewLazy returns correct value -- NoLock", actual)
}

func Test_NewLazy_LockIf_True(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyRegex.NewLockIf(true, `^\d+$`)

	// Act
	actual := args.Map{"isDefined": lr.IsDefined()}

	// Assert
	expected := args.Map{"isDefined": true}
	expected.ShouldBeEqual(t, 0, "NewLazy returns correct value -- LockIf true", actual)
}

func Test_NewLazy_LockIf_False(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyRegex.NewLockIf(false, `^\d+$`)

	// Act
	actual := args.Map{"isDefined": lr.IsDefined()}

	// Assert
	expected := args.Map{"isDefined": true}
	expected.ShouldBeEqual(t, 0, "NewLazy returns correct value -- LockIf false", actual)
}
