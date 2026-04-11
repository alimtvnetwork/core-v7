package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── LazyRegex.IsMatchBytes — failure path ──
// Covers LazyRegex.go L244-246

func Test_LazyRegex_IsMatchBytes_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	result := lr.IsMatchBytes([]byte("test"))

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes returns error -- invalid pattern", actual)
}

// ── regExMatchValidationError — err != nil path ──
// Covers regExMatchValidationError.go via MatchError with invalid pattern
// Note: The regEx == nil path (L21-26) is dead code because
// regexp.Compile never returns (nil, nil).

func Test_MatchError_InvalidPattern(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchError("[invalid", "test")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- invalid pattern", actual)
}
