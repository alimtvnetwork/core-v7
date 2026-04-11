package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── MatchUsingCustomizeErrorFuncLock — custom error ──

func Test_MatchUsingCustomizeErrorFuncLock_CustomError(t *testing.T) {
	// Arrange
	customCalled := false
	customErr := func(pattern, comparing string, err error, r *regexp.Regexp) error {
		customCalled = true
		return err
	}
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	regexnew.MatchUsingCustomizeErrorFuncLock(
		`^\d+$`, "abc", matchFunc, customErr,
	)

	// Assert
	actual := args.Map{"customCalled": customCalled}
	expected := args.Map{"customCalled": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingCustomizeErrorFuncLock returns custom error -- no match with custom func", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_NilCustomizer(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := regexnew.MatchUsingCustomizeErrorFuncLock(
		`^\d+$`, "abc", matchFunc, nil,
	)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingCustomizeErrorFuncLock returns default error -- nil customizer", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_Match_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := regexnew.MatchUsingCustomizeErrorFuncLock(
		`^\d+$`, "123", matchFunc, nil,
	)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "MatchUsingCustomizeErrorFuncLock returns nil -- match", actual)
}

// ── MatchUsingFuncErrorLock ──

func Test_MatchUsingFuncErrorLock_Match_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFunc)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns nil -- match", actual)
}

func Test_MatchUsingFuncErrorLock_NoMatch_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFunc)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns error -- no match", actual)
}

// ── MatchErrorLock ──

func Test_MatchErrorLock_Match_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchErrorLock(`^\d+$`, "123")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock returns nil -- match", actual)
}

func Test_MatchErrorLock_NoMatch_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchErrorLock(`^\d+$`, "abc")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock returns error -- no match", actual)
}

// ── IsMatchFailed ──

func Test_IsMatchFailed_True(t *testing.T) {
	// Arrange & Act
	result := regexnew.IsMatchFailed(`^\d+$`, "abc")

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed returns true -- no match", actual)
}

func Test_IsMatchFailed_False(t *testing.T) {
	// Arrange & Act
	result := regexnew.IsMatchFailed(`^\d+$`, "123")

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": false}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed returns false -- match", actual)
}

// ── newCreator methods ──

func Test_NewCreator_LazyLock_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyLock(`^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":    lr != nil,
		"isDefined": lr.IsDefined(),
	}
	expected := args.Map{
		"notNil":    true,
		"isDefined": true,
	}
	expected.ShouldBeEqual(t, 0, "New.LazyLock returns defined -- valid pattern", actual)
}

func Test_NewCreator_DefaultLockIf_Locked(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.New.DefaultLockIf(true, `^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf returns regex -- locked", actual)
}

func Test_NewCreator_DefaultLockIf_Unlocked(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.New.DefaultLockIf(false, `^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf returns regex -- unlocked", actual)
}

func Test_NewCreator_DefaultApplicableLock_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	r, err, isApplicable := regexnew.New.DefaultApplicableLock(`^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":       r != nil,
		"noError":      err == nil,
		"isApplicable": isApplicable,
	}
	expected := args.Map{
		"notNil":       true,
		"noError":      true,
		"isApplicable": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultApplicableLock returns applicable -- valid", actual)
}

// ── newLazyRegexCreator methods ──

func Test_LazyRegexCreator_TwoLock(t *testing.T) {
	// Arrange & Act
	first, second := regexnew.New.LazyRegex.TwoLock(`^\d+$`, `^[a-z]+$`)

	// Assert
	actual := args.Map{
		"firstNotNil":  first != nil,
		"secondNotNil": second != nil,
	}
	expected := args.Map{
		"firstNotNil":  true,
		"secondNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.TwoLock returns two -- valid patterns", actual)
}

func Test_LazyRegexCreator_ManyUsingLock(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.ManyUsingLock(`^\d+$`, `^[a-z]+$`, `^\w+$`)

	// Assert
	actual := args.Map{
		"length": len(result),
	}
	expected := args.Map{
		"length": 3,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.ManyUsingLock returns map -- 3 patterns", actual)
}

func Test_LazyRegexCreator_ManyUsingLock_Empty(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.ManyUsingLock()

	// Assert
	actual := args.Map{
		"length": len(result),
	}
	expected := args.Map{
		"length": 0,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.ManyUsingLock returns empty -- no patterns", actual)
}

func Test_LazyRegexCreator_AllPatternsMap(t *testing.T) {
	// Arrange
	_ = regexnew.New.Lazy(`test-all-patterns-\d+`)

	// Act
	result := regexnew.New.LazyRegex.AllPatternsMap()

	// Assert
	actual := args.Map{
		"hasItems": len(result) > 0,
	}
	expected := args.Map{
		"hasItems": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.AllPatternsMap returns map -- has items", actual)
}

func Test_LazyRegexCreator_NewLockIf_Locked(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.NewLockIf(true, `^\d+$`)

	// Assert
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewLockIf returns lazy -- locked", actual)
}

func Test_LazyRegexCreator_NewLockIf_Unlocked(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.NewLockIf(false, `^\d+$`)

	// Assert
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewLockIf returns lazy -- unlocked", actual)
}

// ── CreateMustLockIf ──

func Test_CreateMustLockIf_Locked(t *testing.T) {
	// Arrange & Act
	r := regexnew.CreateMustLockIf(true, `^\d+$`)

	// Assert
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns regex -- locked", actual)
}

func Test_CreateMustLockIf_Unlocked(t *testing.T) {
	// Arrange & Act
	r := regexnew.CreateMustLockIf(false, `^\d+$`)

	// Assert
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf returns regex -- unlocked", actual)
}

// ── NewMustLock ──

func Test_NewMustLock_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	r := regexnew.NewMustLock(`^\d+$`)

	// Assert
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock returns regex -- valid", actual)
}

// ── CreateApplicableLock — invalid ──

func Test_CreateApplicableLock_Invalid_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange & Act
	r, err, isApplicable := regexnew.CreateApplicableLock("[invalid")

	// Assert
	actual := args.Map{
		"regexNil":     r == nil,
		"hasError":     err != nil,
		"isApplicable": isApplicable,
	}
	expected := args.Map{
		"regexNil":     true,
		"hasError":     true,
		"isApplicable": false,
	}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock returns not applicable -- invalid", actual)
}

// ── CreateLockIf ──

func Test_CreateLockIf_Locked(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.CreateLockIf(true, `^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLockIf returns regex -- locked", actual)
}

func Test_CreateLockIf_Unlocked(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.CreateLockIf(false, `^\d+$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLockIf returns regex -- unlocked", actual)
}

// ── LazyRegex.MatchError on LazyRegex (vs package-level) ──

func Test_LazyRegex_MatchError_Match_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	err := lr.MatchError("123")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MatchError returns nil -- match", actual)
}

func Test_LazyRegex_MatchError_NoMatch_FromMatchUsingCustomizeE(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	err := lr.MatchError("abc")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MatchError returns error -- no match", actual)
}
