package regexnewtests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// Create / CreateLock / CreateLockIf / CreateMust / CreateMustLockIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_Create_Valid(t *testing.T) {
	// Arrange
	re, err := regexnew.Create(`^\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Create returns non-empty -- valid", actual)
}

func Test_Create_Invalid(t *testing.T) {
	// Arrange
	_, err := regexnew.Create(`[invalid`)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid", actual)
}

func Test_Create_Cached(t *testing.T) {
	// Arrange
	re1, _ := regexnew.Create(`^cov6cached\d+$`)
	re2, _ := regexnew.Create(`^cov6cached\d+$`)

	// Act
	actual := args.Map{"same": re1 == re2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Create returns correct value -- cached", actual)
}

func Test_CreateLock_Valid(t *testing.T) {
	// Arrange
	re, err := regexnew.CreateLock(`^cov6lock\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLock returns non-empty -- valid", actual)
}

func Test_CreateLockIf_WithLock(t *testing.T) {
	// Arrange
	re, err := regexnew.CreateLockIf(true, `^cov6lockif\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLockIf returns correct value -- lock", actual)
}

func Test_CreateLockIf_WithoutLock(t *testing.T) {
	// Arrange
	re, err := regexnew.CreateLockIf(false, `^cov6lockifno\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLockIf returns empty -- no lock", actual)
}

func Test_CreateMust_Valid(t *testing.T) {
	// Arrange
	re := regexnew.CreateMust(`^cov6must\d+$`)

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns non-empty -- valid", actual)
}

func Test_CreateMust_Cached(t *testing.T) {
	// Arrange
	re1 := regexnew.CreateMust(`^cov6mustcache\d+$`)
	re2 := regexnew.CreateMust(`^cov6mustcache\d+$`)

	// Act
	actual := args.Map{"same": re1 == re2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns correct value -- cached", actual)
}

func Test_CreateMustLockIf_WithLock_FromCreateValid(t *testing.T) {
	// Arrange
	re := regexnew.CreateMustLockIf(true, `^cov6mustlock\d+$`)

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns correct value -- lock", actual)
}

func Test_CreateMustLockIf_WithoutLock_FromCreateValid(t *testing.T) {
	// Arrange
	re := regexnew.CreateMustLockIf(false, `^cov6mustlockno\d+$`)

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns empty -- no lock", actual)
}

func Test_CreateApplicableLock_Valid_FromCreateValid(t *testing.T) {
	// Arrange
	re, err, ok := regexnew.CreateApplicableLock(`^cov6applock\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
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

func Test_CreateApplicableLock_Invalid_FromCreateValid(t *testing.T) {
	// Arrange
	_, err, ok := regexnew.CreateApplicableLock(`[invalid`)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock returns error -- invalid", actual)
}

func Test_NewMustLock(t *testing.T) {
	// Arrange
	re := regexnew.NewMustLock(`^cov6newmust\d+$`)

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsMatchLock / IsMatchFailed
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsMatchLock_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": regexnew.IsMatchLock(`^\d+$`, "123")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns correct value -- match", actual)
}

func Test_IsMatchLock_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"v": regexnew.IsMatchLock(`^\d+$`, "abc")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns empty -- no match", actual)
}

func Test_IsMatchLock_InvalidRegex(t *testing.T) {
	// Act
	actual := args.Map{"v": regexnew.IsMatchLock(`[invalid`, "abc")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns error -- invalid regex", actual)
}

func Test_IsMatchFailed_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": regexnew.IsMatchFailed(`^\d+$`, "123")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed returns correct value -- match", actual)
}

func Test_IsMatchFailed_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"v": regexnew.IsMatchFailed(`^\d+$`, "abc")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed returns empty -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchError / MatchErrorLock / MatchUsingFuncErrorLock / MatchUsingCustomizeErrorFuncLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_MatchError_Match(t *testing.T) {
	// Act
	actual := args.Map{"noErr": regexnew.MatchError(`^\d+$`, "123") == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- match", actual)
}

func Test_MatchError_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": regexnew.MatchError(`^\d+$`, "abc") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns empty -- no match", actual)
}

func Test_MatchError_InvalidRegex(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": regexnew.MatchError(`[invalid`, "abc") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- invalid", actual)
}

func Test_MatchErrorLock_Match(t *testing.T) {
	// Act
	actual := args.Map{"noErr": regexnew.MatchErrorLock(`^\d+$`, "123") == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock returns error -- match", actual)
}

func Test_MatchErrorLock_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": regexnew.MatchErrorLock(`^\d+$`, "abc") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock returns empty -- no match", actual)
}

func Test_MatchUsingFuncErrorLock_Match(t *testing.T) {
	// Arrange
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"noErr": regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFn) == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns error -- match", actual)
}

func Test_MatchUsingFuncErrorLock_NoMatch(t *testing.T) {
	// Arrange
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"hasErr": regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFn) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns empty -- no match", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_Match(t *testing.T) {
	// Arrange
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"noErr": regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "123", matchFn, nil) == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock returns error -- match", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_NilCustomize(t *testing.T) {
	// Arrange
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"hasErr": regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFn, nil) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock returns nil -- no match nil customize", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch_CustomErr(t *testing.T) {
	// Arrange
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	customErr := func(pattern, term string, err error, re *regexp.Regexp) error {
		return fmt.Errorf("custom error for %s", term)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFn, customErr)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"custom": err.Error() == "custom error for abc",
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"custom": true,
	}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock returns error -- custom err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_LazyRegex_IsNull(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"nil": lr.IsNull()}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsNull", actual)
}

func Test_LazyRegex_IsDefined(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^cov6lazy\d+$`)

	// Act
	actual := args.Map{"defined": lr.IsDefined()}

	// Assert
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsDefined", actual)
}

func Test_LazyRegex_IsUndefined_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.IsUndefined()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns nil -- IsUndefined nil", actual)
}

func Test_LazyRegex_IsApplicable(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6lazyapp\d+$`)

	// Act
	actual := args.Map{"v": lr.IsApplicable()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsApplicable", actual)
}

func Test_LazyRegex_IsApplicable_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.IsApplicable()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns nil -- IsApplicable nil", actual)
}

func Test_LazyRegex_IsApplicable_Cached(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6lazycached\d+$`)
	_ = lr.IsApplicable() // first call compiles

	// Act
	actual := args.Map{"v": lr.IsApplicable()} // second call returns cached

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsApplicable cached", actual)
}

func Test_LazyRegex_Compile_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6comp\d+$`)
	re, err := lr.Compile()

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns non-empty -- Compile valid", actual)
}

func Test_LazyRegex_Compile_Cached(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6compcache\d+$`)
	re1, _ := lr.Compile()
	re2, _ := lr.Compile()

	// Act
	actual := args.Map{"same": re1 == re2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- Compile cached", actual)
}

func Test_LazyRegex_IsCompiled(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6iscomp\d+$`)

	// Act
	actual := args.Map{"before": lr.IsCompiled()}

	// Assert
	expected := args.Map{"before": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsCompiled before", actual)
	lr.Compile()
	actual2 := args.Map{"after": lr.IsCompiled()}
	expected2 := args.Map{"after": true}
	expected2.ShouldBeEqual(t, 0, "LazyRegex returns correct value -- IsCompiled after", actual2)
}

func Test_LazyRegex_IsCompiled_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.IsCompiled()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns nil -- IsCompiled nil", actual)
}

func Test_LazyRegex_OnRequiredCompiled_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex
	err := lr.OnRequiredCompiled()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled returns nil -- nil", actual)
}

func Test_LazyRegex_OnRequiredCompiled_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6reqcomp\d+$`)
	err := lr.OnRequiredCompiled()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled returns non-empty -- valid", actual)
}

func Test_LazyRegex_OnRequiredCompiled_AlreadyCompiled(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6reqcomp2\d+$`)
	lr.Compile()
	err := lr.OnRequiredCompiled()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled returns correct value -- already compiled", actual)
}

func Test_LazyRegex_HasError(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6haserr\d+$`)

	// Act
	actual := args.Map{"v": lr.HasError()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex returns error -- HasError", actual)
}

func Test_LazyRegex_HasAnyIssues_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.HasAnyIssues()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues returns nil -- nil", actual)
}

func Test_LazyRegex_HasAnyIssues_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6anyiss\d+$`)

	// Act
	actual := args.Map{"v": lr.HasAnyIssues()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues returns non-empty -- valid", actual)
}

func Test_LazyRegex_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.IsInvalid()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid returns nil -- nil", actual)
}

func Test_LazyRegex_IsInvalid_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6inv\d+$`)

	// Act
	actual := args.Map{"v": lr.IsInvalid()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsInvalid returns error -- valid", actual)
}

func Test_LazyRegex_CompiledError(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6comperr\d+$`)

	// Act
	actual := args.Map{"noErr": lr.CompiledError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with args", actual)
}

func Test_LazyRegex_Error(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6err\d+$`)

	// Act
	actual := args.Map{"noErr": lr.Error() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns error -- with args", actual)
}

func Test_LazyRegex_MustBeSafe(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6safe\d+$`)
	// should not panic
	lr.MustBeSafe()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe returns correct value -- with args", actual)
}

func Test_LazyRegex_String_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.String()}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
}

func Test_LazyRegex_String_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6str\d+$`)

	// Act
	actual := args.Map{"v": lr.String()}

	// Assert
	expected := args.Map{"v": `^cov6str\d+$`}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- valid", actual)
}

func Test_LazyRegex_FullString_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.FullString()}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "FullString returns nil -- nil", actual)
}

func Test_LazyRegex_FullString_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6full\d+$`)

	// Act
	actual := args.Map{"notEmpty": lr.FullString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns non-empty -- valid", actual)
}

func Test_LazyRegex_Pattern_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	actual := args.Map{"v": lr.Pattern()}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Pattern returns nil -- nil", actual)
}

func Test_LazyRegex_Pattern_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6pat\d+$`)

	// Act
	actual := args.Map{"v": lr.Pattern()}

	// Assert
	expected := args.Map{"v": `^cov6pat\d+$`}
	expected.ShouldBeEqual(t, 0, "Pattern returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — Match methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_LazyRegex_MatchError_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6me\d+$`)

	// Act
	actual := args.Map{"noErr": lr.MatchError("cov6me123") == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- match", actual)
}

func Test_LazyRegex_MatchError_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6me2\d+$`)

	// Act
	actual := args.Map{"hasErr": lr.MatchError("abc") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns empty -- no match", actual)
}

func Test_LazyRegex_MatchUsingFuncError_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6mfe\d+$`)
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"noErr": lr.MatchUsingFuncError("cov6mfe123", matchFn) == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns error -- match", actual)
}

func Test_LazyRegex_MatchUsingFuncError_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6mfe2\d+$`)
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }

	// Act
	actual := args.Map{"hasErr": lr.MatchUsingFuncError("abc", matchFn) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns empty -- no match", actual)
}

func Test_LazyRegex_IsMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6im\d+$`)

	// Act
	actual := args.Map{
		"match": lr.IsMatch("cov6im123"),
		"noMatch": lr.IsMatch("abc"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct value -- with args", actual)
}

func Test_LazyRegex_IsMatchBytes(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6imb\d+$`)

	// Act
	actual := args.Map{
		"match": lr.IsMatchBytes([]byte("cov6imb123")),
		"noMatch": lr.IsMatchBytes([]byte("abc")),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes returns correct value -- with args", actual)
}

func Test_LazyRegex_IsFailedMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6ifm\d+$`)

	// Act
	actual := args.Map{
		"match": lr.IsFailedMatch("cov6ifm123"),
		"noMatch": lr.IsFailedMatch("abc"),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"noMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "IsFailedMatch returns correct value -- with args", actual)
}

func Test_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6ifmb\d+$`)

	// Act
	actual := args.Map{
		"match": lr.IsFailedMatchBytes([]byte("cov6ifmb123")),
		"noMatch": lr.IsFailedMatchBytes([]byte("abc")),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"noMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "IsFailedMatchBytes returns correct value -- with args", actual)
}

func Test_LazyRegex_FirstMatchLine_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`cov6fml(\d+)`)
	line, invalid := lr.FirstMatchLine("cov6fml123")

	// Act
	actual := args.Map{
		"line": line,
		"invalid": invalid,
	}

	// Assert
	expected := args.Map{
		"line": "cov6fml123",
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine returns correct value -- match", actual)
}

func Test_LazyRegex_FirstMatchLine_NoMatch_FromCreateValid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6fml2\d+$`)
	line, invalid := lr.FirstMatchLine("abc")

	// Act
	actual := args.Map{
		"line": line,
		"invalid": invalid,
	}

	// Assert
	expected := args.Map{
		"line": "",
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine returns empty -- no match", actual)
}

func Test_LazyRegex_CompileMust(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6cm\d+$`)
	re := lr.CompileMust()

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompileMust returns correct value -- with args", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6orcm\d+$`)
	lr.OnRequiredCompiledMust() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiledMust returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// lazyRegexMap — via New.LazyRegex methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewCreator_Lazy(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^cov6nclazy\d+$`)

	// Act
	actual := args.Map{
		"notNil": lr != nil,
		"defined": lr.IsDefined(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"defined": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Lazy returns correct value -- with args", actual)
}

func Test_NewCreator_LazyLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6nclazylock\d+$`)

	// Act
	actual := args.Map{"notNil": lr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.LazyLock returns correct value -- with args", actual)
}

func Test_NewCreator_Default(t *testing.T) {
	// Arrange
	re, err := regexnew.New.Default(`^cov6ncdef\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Default returns correct value -- with args", actual)
}

func Test_NewCreator_DefaultLock(t *testing.T) {
	// Arrange
	re, err := regexnew.New.DefaultLock(`^cov6ncdefl\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultLock returns correct value -- with args", actual)
}

func Test_NewCreator_DefaultLockIf(t *testing.T) {
	// Arrange
	re, err := regexnew.New.DefaultLockIf(true, `^cov6ncdefli\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf returns correct value -- with args", actual)
}

func Test_NewCreator_DefaultApplicableLock(t *testing.T) {
	// Arrange
	re, err, ok := regexnew.New.DefaultApplicableLock(`^cov6ncal\d+$`)

	// Act
	actual := args.Map{
		"notNil": re != nil,
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

func Test_NewLazyRegexCreator_TwoLock(t *testing.T) {
	// Arrange
	first, second := regexnew.New.LazyRegex.TwoLock(`^cov6two1\d+$`, `^cov6two2\d+$`)

	// Act
	actual := args.Map{
		"f": first != nil,
		"s": second != nil,
	}

	// Assert
	expected := args.Map{
		"f": true,
		"s": true,
	}
	expected.ShouldBeEqual(t, 0, "TwoLock returns correct value -- with args", actual)
}

func Test_NewLazyRegexCreator_ManyUsingLock_Empty(t *testing.T) {
	// Arrange
	result := regexnew.New.LazyRegex.ManyUsingLock()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ManyUsingLock returns empty -- empty", actual)
}

func Test_NewLazyRegexCreator_ManyUsingLock(t *testing.T) {
	// Arrange
	result := regexnew.New.LazyRegex.ManyUsingLock(`^cov6many1\d+$`, `^cov6many2\d+$`)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ManyUsingLock returns correct value -- with args", actual)
}

func Test_NewLazyRegexCreator_AllPatternsMap(t *testing.T) {
	// Arrange
	result := regexnew.New.LazyRegex.AllPatternsMap()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AllPatternsMap returns correct value -- with args", actual)
}

func Test_NewLazyRegexCreator_NewLockIf_Lock(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyRegex.NewLockIf(true, `^cov6nli1\d+$`)

	// Act
	actual := args.Map{"notNil": lr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewLockIf returns correct value -- lock", actual)
}

func Test_NewLazyRegexCreator_NewLockIf_NoLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyRegex.NewLockIf(false, `^cov6nli2\d+$`)

	// Act
	actual := args.Map{"notNil": lr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewLockIf returns empty -- no lock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// regExMatchValidationError — all 3 branches (via MatchError)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RegExMatchValidationError_CompileError(t *testing.T) {
	// Arrange
	err := regexnew.MatchError(`[invalid`, "abc")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"hasCompile": err.Error() != "",
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"hasCompile": true,
	}
	expected.ShouldBeEqual(t, 0, "regExMatchValidationError returns error -- compile err", actual)
}

func Test_RegExMatchValidationError_NoMatch(t *testing.T) {
	// Arrange
	err := regexnew.MatchError(`^\d+$`, "abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "regExMatchValidationError returns empty -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// prettyJson — via FullString
// ══════════════════════════════════════════════════════════════════════════════

func Test_PrettyJson_ViaFullString(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov6pj\d+$`)
	s := lr.FullString()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "prettyJson returns correct value -- via FullString", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// regexes-compiled.go — pre-compiled vars
// ══════════════════════════════════════════════════════════════════════════════

func Test_PrecompiledRegexes(t *testing.T) {
	// Act
	actual := args.Map{
		"ws":   regexnew.WhitespaceFinderRegex != nil,
		"hash": regexnew.HashCommentWithSpaceOptionalRegex != nil,
		"pipe": regexnew.WhitespaceOrPipeFinderRegex != nil,
		"dol":  regexnew.DollarIdentifierRegex != nil,
		"pct":  regexnew.PercentIdentifierRegex != nil,
		"pn":   regexnew.PrettyNameRegex != nil,
		"id":   regexnew.ExactIdFieldMatchingRegex != nil,
		"vid":  regexnew.ExactVersionIdFieldMatchingRegex != nil,
		"ubu":  regexnew.UbuntuNameCheckerRegex != nil,
		"cent": regexnew.CentOsNameCheckerRegex != nil,
		"rh":   regexnew.RedHatNameCheckerRegex != nil,
		"num":  regexnew.FirstNumberAnyWhereCheckerRegex != nil,
		"win":  regexnew.WindowsVersionNumberCheckerRegex != nil,
	}

	// Assert
	expected := args.Map{
		"ws": true, "hash": true, "pipe": true, "dol": true, "pct": true,
		"pn": true, "id": true, "vid": true, "ubu": true, "cent": true,
		"rh": true, "num": true, "win": true,
	}
	expected.ShouldBeEqual(t, 0, "pre-compiled returns correct value -- regexes", actual)
}
