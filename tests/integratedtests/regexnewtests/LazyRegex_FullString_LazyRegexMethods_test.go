package regexnewtests

import (
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── LazyRegex.FullString — valid pattern ──

func Test_LazyRegex_FullString_ValidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	result := lr.FullString()

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"hasPattern": strings.Contains(result, "pattern"),
	}
	expected := args.Map{
		"hasContent": true,
		"hasPattern": true,
	}
	expected.ShouldBeEqual(t, 0, "FullString returns json -- valid pattern", actual)
}

func Test_LazyRegex_FullString_NilReceiver(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	result := lr.FullString()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "FullString returns empty -- nil receiver", actual)
}

// ── LazyRegex.FirstMatchLine ──

func Test_LazyRegex_FirstMatchLine_ValidMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`(\d+)`)

	// Act
	match, isInvalid := lr.FirstMatchLine("abc123def")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "123",
		"isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine returns match -- valid content", actual)
}

func Test_LazyRegex_FirstMatchLine_NoMatch_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`(\d+)`)

	// Act
	match, isInvalid := lr.FirstMatchLine("abcdef")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine returns invalid -- no match", actual)
}

func Test_LazyRegex_FirstMatchLine_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	match, isInvalid := lr.FirstMatchLine("test")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine returns invalid -- bad pattern", actual)
}

// ── LazyRegex.CompileMust — panic path ──

func Test_LazyRegex_CompileMust_ValidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	regEx := lr.CompileMust()

	// Assert
	actual := args.Map{"notNil": regEx != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompileMust returns regex -- valid pattern", actual)
}

func Test_LazyRegex_CompileMust_Panics(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.CompileMust()
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "CompileMust panics -- invalid pattern", actual)
}

// ── LazyRegex.OnRequiredCompiledMust ──

func Test_LazyRegex_OnRequiredCompiledMust_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.OnRequiredCompiledMust()
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiledMust returns ok -- valid", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust_NilPanics(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.OnRequiredCompiledMust()
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiledMust panics -- nil receiver", actual)
}

// ── LazyRegex.MustBeSafe ──

func Test_LazyRegex_MustBeSafe_InvalidPanics(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.MustBeSafe()
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe panics -- invalid pattern", actual)
}

func Test_LazyRegex_MustBeSafe_ValidOk(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	panicked := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.MustBeSafe()
	}()

	// Assert
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "MustBeSafe returns ok -- valid pattern", actual)
}

// ── LazyRegex.HasError ──

func Test_LazyRegex_HasError_ValidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	result := lr.HasError()

	// Assert
	actual := args.Map{"hasError": result}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "HasError returns false -- valid pattern", actual)
}

func Test_LazyRegex_HasError_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	result := lr.HasError()

	// Assert
	actual := args.Map{"hasError": result}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "HasError returns true -- invalid pattern", actual)
}

// ── LazyRegex.HasAnyIssues / IsInvalid ──

func Test_LazyRegex_HasAnyIssues_Nil_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	result := lr.HasAnyIssues()

	// Assert
	actual := args.Map{"hasIssues": result}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues returns true -- nil", actual)
}

func Test_LazyRegex_IsInvalid_Nil_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	result := lr.IsInvalid()

	// Assert
	actual := args.Map{"isInvalid": result}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid returns true -- nil", actual)
}

// ── LazyRegex.IsFailedMatch / IsFailedMatchBytes ──

func Test_LazyRegex_IsFailedMatch_ValidNoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	result := lr.IsFailedMatch("abc")

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatch returns true -- no match", actual)
}

func Test_LazyRegex_IsFailedMatch_ValidMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	result := lr.IsFailedMatch("123")

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": false}
	expected.ShouldBeEqual(t, 0, "IsFailedMatch returns false -- match", actual)
}

func Test_LazyRegex_IsFailedMatchBytes_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	result := lr.IsFailedMatchBytes([]byte("test"))

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatchBytes returns true -- invalid pattern", actual)
}

func Test_LazyRegex_IsFailedMatchBytes_ValidMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	result := lr.IsFailedMatchBytes([]byte("123"))

	// Assert
	actual := args.Map{"isFailed": result}
	expected := args.Map{"isFailed": false}
	expected.ShouldBeEqual(t, 0, "IsFailedMatchBytes returns false -- match", actual)
}

// ── LazyRegex.MatchUsingFuncError ──

func Test_LazyRegex_MatchUsingFuncError_Match_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := lr.MatchUsingFuncError("123", matchFunc)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns nil -- match", actual)
}

func Test_LazyRegex_MatchUsingFuncError_NoMatch_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	matchFunc := func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	}

	// Act
	err := lr.MatchUsingFuncError("abc", matchFunc)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns error -- no match", actual)
}

// ── LazyRegex.CompiledError / Error ──

func Test_LazyRegex_CompiledError_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act
	err := lr.CompiledError()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "CompiledError returns nil -- valid", actual)
}

func Test_LazyRegex_Error_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	err := lr.Error()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Error returns error -- invalid", actual)
}

// ── LazyRegex.OnRequiredCompiled — nil path ──

func Test_LazyRegex_OnRequiredCompiled_Nil_FromLazyRegexFullStringL(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	err := lr.OnRequiredCompiled()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled returns error -- nil", actual)
}
