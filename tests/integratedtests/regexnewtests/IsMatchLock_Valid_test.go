package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// IsMatchLock — additional branches
// =============================================================================

func Test_IsMatchLock_Valid(t *testing.T) {
	// Arrange
	isMatch := regexnew.IsMatchLock(`^\d+$`, "123")

	// Act
	actual := args.Map{"isMatch": isMatch}

	// Assert
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns true -- valid match", actual)
}

func Test_IsMatchLock_NoMatch_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	isMatch := regexnew.IsMatchLock(`^\d+$`, "abc")

	// Act
	actual := args.Map{"isMatch": isMatch}

	// Assert
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- no match", actual)
}

func Test_IsMatchLock_InvalidPattern(t *testing.T) {
	// Arrange
	isMatch := regexnew.IsMatchLock(`[invalid`, "abc")

	// Act
	actual := args.Map{"isMatch": isMatch}

	// Assert
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- invalid pattern", actual)
}

// =============================================================================
// Create — additional branches
// =============================================================================

func Test_Create_Valid_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	r, err := regexnew.Create(`^\d+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Create returns regex -- valid pattern", actual)
}

func Test_Create_Invalid_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	_, err := regexnew.Create(`[invalid`)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid pattern", actual)
}

func Test_CreateLock_Valid_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	r, err := regexnew.CreateLock(`^\d+$`)

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "CreateLock returns regex -- valid", actual)
}

func Test_CreateMust_Valid_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	r := regexnew.CreateMust(`^\d+$`)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns regex -- valid", actual)
}

// =============================================================================
// LazyRegex — additional methods
// =============================================================================

func Test_LazyRegex_Compile(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)
	compiled, err := lazy.Compile()

	// Act
	actual := args.Map{
		"notNil": compiled != nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compile returns regex -- valid", actual)
}

func Test_LazyRegex_IsMatch_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"matchDigits": lazy.IsMatch("123"),
		"failAlpha":   lazy.IsMatch("abc"),
	}

	// Assert
	expected := args.Map{
		"matchDigits": true,
		"failAlpha": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatch returns expected -- digits vs alpha", actual)
}

func Test_LazyRegex_IsMatchBytes_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"matchDigits": lazy.IsMatchBytes([]byte("123")),
		"failAlpha":   lazy.IsMatchBytes([]byte("abc")),
	}

	// Assert
	expected := args.Map{
		"matchDigits": true,
		"failAlpha": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatchBytes returns expected -- digits vs alpha", actual)
}

func Test_LazyRegex_IsFailedMatch_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"failAlpha":   lazy.IsFailedMatch("abc"),
		"failDigits":  lazy.IsFailedMatch("123"),
	}

	// Assert
	expected := args.Map{
		"failAlpha": true,
		"failDigits": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatch returns expected -- alpha fails", actual)
}

func Test_LazyRegex_IsFailedMatchBytes_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"failAlpha": lazy.IsFailedMatchBytes([]byte("abc")),
	}

	// Assert
	expected := args.Map{"failAlpha": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatchBytes returns true -- alpha", actual)
}

func Test_LazyRegex_OnRequiredCompiled(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)
	err := lazy.OnRequiredCompiled()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiled returns nil -- valid", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)
	lazy.OnRequiredCompiledMust() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiledMust no panic -- valid", actual)
}

func Test_LazyRegex_StringAndPattern(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"string":  lazy.String(),
		"pattern": lazy.Pattern(),
		"isNull":  lazy.IsNull(),
		"isDef":   lazy.IsDefined(),
		"isUndef": lazy.IsUndefined(),
		"isComp":  lazy.IsCompiled(),
		"isAppl":  lazy.IsApplicable(),
	}

	// Assert
	expected := args.Map{
		"string": `^\d+$`, "pattern": `^\d+$`,
		"isNull": false, "isDef": true, "isUndef": false,
		"isComp": true, "isAppl": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex String/Pattern returns expected -- valid", actual)
}

func Test_LazyRegex_MatchError_Valid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"matchErr":   lazy.MatchError("123") == nil,
		"noMatchErr": lazy.MatchError("abc") != nil,
	}

	// Assert
	expected := args.Map{
		"matchErr": true,
		"noMatchErr": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchError returns expected -- match vs no match", actual)
}

func Test_LazyRegex_FirstMatchLine_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`\d+`)
	match, isInvalid := lazy.FirstMatchLine("abc123def")

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
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine returns 123 -- valid", actual)
}

func Test_LazyRegex_FirstMatchLine_NoMatch_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`\d+`)
	match, isInvalid := lazy.FirstMatchLine("abc")

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
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine returns empty -- no match", actual)
}

func Test_LazyRegex_HasError_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"hasError": lazy.HasError(),
		"hasIssues": lazy.HasAnyIssues(),
		"isInvalid": lazy.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"hasError": false,
		"hasIssues": false,
		"isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasError returns false -- valid pattern", actual)
}

func Test_LazyRegex_CompiledError_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{
		"hasErr": lazy.CompiledError() != nil,
		"errNil": lazy.Error() == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompiledError returns nil -- valid", actual)
}

func Test_LazyRegex_MustBeSafe_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)
	lazy.MustBeSafe() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MustBeSafe no panic -- valid", actual)
}

func Test_LazyRegex_FullString_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)

	// Act
	actual := args.Map{"notEmpty": lazy.FullString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FullString returns non-empty -- valid", actual)
}

func Test_LazyRegex_CompileMust_FromIsMatchLockValid(t *testing.T) {
	// Arrange
	lazy := regexnew.New.Lazy(`^\d+$`)
	r := lazy.CompileMust()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompileMust returns regex -- valid", actual)
}
