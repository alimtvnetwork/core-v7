package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── CreateMustLockIf ──

func Test_CreateMustLockIf_WithLock(t *testing.T) {
	// Arrange
	r := regexnew.CreateMustLockIf(true, `^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns non-empty -- with lock", actual)
}

func Test_CreateMustLockIf_WithoutLock(t *testing.T) {
	// Arrange
	r := regexnew.CreateMustLockIf(false, `^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns non-empty -- without lock", actual)
}

// ── CreateApplicableLock ──

func Test_CreateApplicableLock_Valid(t *testing.T) {
	// Arrange
	r, err, ok := regexnew.CreateApplicableLock(`^\d+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": err == nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock returns non-empty -- valid", actual)
}

func Test_CreateApplicableLock_Invalid(t *testing.T) {
	// Arrange
	r, err, ok := regexnew.CreateApplicableLock(`[invalid`)

	// Act
	actual := args.Map{
		"isNil": r == nil,
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"hasErr": true,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock returns error -- invalid", actual)
}

// ── IsMatchFailed ──

func Test_IsMatchFailed(t *testing.T) {
	// Act
	actual := args.Map{
		"fail":    regexnew.IsMatchFailed(`^\d+$`, "abc"),
		"noFail":  regexnew.IsMatchFailed(`^\d+$`, "123"),
		"invalid": regexnew.IsMatchFailed(`[invalid`, "abc"),
	}

	// Assert
	expected := args.Map{
		"fail": true,
		"noFail": false,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed returns correct value -- with args", actual)
}

// ── MatchError ──

func Test_MatchError_Success(t *testing.T) {
	// Arrange
	err := regexnew.MatchError(`^\d+$`, "123")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- success", actual)
}

func Test_MatchError_Fail(t *testing.T) {
	// Arrange
	err := regexnew.MatchError(`^\d+$`, "abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- fail", actual)
}

// ── MatchErrorLock ──

func Test_MatchErrorLock(t *testing.T) {
	// Arrange
	err := regexnew.MatchErrorLock(`^\d+$`, "123")
	errFail := regexnew.MatchErrorLock(`^\d+$`, "abc")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasErr": errFail != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock returns error -- with args", actual)
}

// ── MatchUsingFuncErrorLock ──

func Test_MatchUsingFuncErrorLock(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFunc)
	errFail := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFunc)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasErr": errFail != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns error -- with args", actual)
}

// ── MatchUsingCustomizeErrorFuncLock — nil custom error func ──

func Test_MatchCustomErr_NilCustom(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr returns nil -- nil custom func", actual)
}

func Test_MatchCustomErr_InvalidRegex(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool { return false }
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`[invalid`, "abc", matchFunc, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr returns error -- invalid regex", actual)
}

// ── PrettyJson ──

func Test_LazyRegex_FullString(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^\d+$`)
	result := lr.FullString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- FullString", actual)
}

// ── newCreator — All creator methods ──

func Test_New_Must(t *testing.T) {
	// Arrange
	r := regexnew.CreateMust(`^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Must returns correct value -- with args", actual)
}

func Test_New_MustLock(t *testing.T) {
	// Arrange
	r := regexnew.NewMustLock(`^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.MustLock returns correct value -- with args", actual)
}

func Test_New_Create(t *testing.T) {
	// Arrange
	r, err := regexnew.New.Default(`^\d+$`)

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
	expected.ShouldBeEqual(t, 0, "New.Create returns correct value -- with args", actual)
}

func Test_New_CreateLock(t *testing.T) {
	// Arrange
	r, err := regexnew.New.DefaultLock(`^\d+$`)

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
	expected.ShouldBeEqual(t, 0, "New.CreateLock returns correct value -- with args", actual)
}

func Test_New_DefaultLockIf(t *testing.T) {
	// Arrange
	r, err := regexnew.New.DefaultLockIf(true, `^\d+$`)

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
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf returns correct value -- with args", actual)
}

func Test_New_DefaultApplicableLock(t *testing.T) {
	// Arrange
	r, err, ok := regexnew.New.DefaultApplicableLock(`^\d+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"noErr": err == nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultApplicableLock returns correct value -- with args", actual)
}

// ── regExMatchValidationError ──

func Test_RegExMatchValidationError(t *testing.T) {
	// Arrange
	// regExMatchValidationError is unexported; test via MatchError instead
	err := regexnew.MatchError(`^\d+$`, "abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RegExMatchValidationError returns error -- via MatchError", actual)
}

// ── LazyRegex — FindStringSubmatch / FindAllString ──

func Test_LazyRegex_FirstMatchLine(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`(\d+)-(\d+)`)
	result, isInvalid := lr.FirstMatchLine("abc 123-456 def")

	// Act
	actual := args.Map{
		"match": result,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"match": "123-456",
		"isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- FirstMatchLine", actual)
}

func Test_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`(\d+)-(\d+)`)
	result, isInvalid := lr.FirstMatchLine("abc def")

	// Act
	actual := args.Map{
		"match": result,
		"isInvalid": isInvalid,
	}

	// Assert
	expected := args.Map{
		"match": "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns empty -- FirstMatchLine no match", actual)
}
