package coreinstructiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestStringSearch_IsMatch verifies match logic.
func TestStringSearch_IsMatch(t *testing.T) {
	for _, tc := range stringSearchIsMatchCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.search.IsMatch(tc.content)

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestStringSearch_NilIsMatch verifies nil receiver returns true.
func TestStringSearch_NilIsMatch(t *testing.T) {
	// Arrange
	var s *coreinstruction.StringSearch

	// Act & Assert
	actual := args.Map{"result": s.IsMatch("anything")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil StringSearch.IsMatch should return true", actual)
}

// TestStringSearch_IsEmpty verifies nil check.
func TestStringSearch_IsEmpty(t *testing.T) {
	var s *coreinstruction.StringSearch
	actual := args.Map{"result": s.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
	actual := args.Map{"result": s.IsExist()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not exist", actual)
	actual := args.Map{"result": s.Has()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have", actual)
}

// TestStringSearch_IsAllMatch verifies all-match logic.
func TestStringSearch_IsAllMatch(t *testing.T) {
	// Arrange
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act & Assert
	actual := args.Map{"result": s.IsAllMatch("hello", "hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all 'hello' should match", actual)
	actual := args.Map{"result": s.IsAllMatch("hello", "world")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mixed should fail", actual)
	actual := args.Map{"result": s.IsAllMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty contents should return true", actual)
}

// TestStringSearch_IsAnyMatchFailed verifies any-fail logic.
func TestStringSearch_IsAnyMatchFailed(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	actual := args.Map{"result": s.IsAnyMatchFailed("hello", "world")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected any match failed", actual)
}

// TestStringSearch_IsMatchFailed verifies match failure.
func TestStringSearch_IsMatchFailed(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	actual := args.Map{"result": s.IsMatchFailed("world")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match failed", actual)
}

// TestStringSearch_VerifyError verifies error on mismatch.
func TestStringSearch_VerifyError(t *testing.T) {
	// Nil returns nil
	var nilS *coreinstruction.StringSearch
	actual := args.Map{"result": nilS.VerifyError("x") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil error", actual)

	// Equal match → nil error
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	actual := args.Map{"result": s.VerifyError("hello") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching content should return nil error", actual)
	actual := args.Map{"result": s.VerifyError("world") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatched content should return error", actual)
}

// TestStringCompare_IsMatch verifies compare match.
func TestStringCompare_IsMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		false,
		"test",
		"test",
	)

	// Act & Assert
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "equal content should match", actual)
	actual := args.Map{"result": sc.IsMatchFailed()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be match failed", actual)
}

// TestStringCompare_Nil verifies nil receiver.
func TestStringCompare_Nil(t *testing.T) {
	var sc *coreinstruction.StringCompare
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should match", actual)
	actual := args.Map{"result": sc.IsDefined()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not be defined", actual)
	actual := args.Map{"result": sc.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
	actual := args.Map{"result": sc.VerifyError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil error", actual)
}

// TestStringCompare_VerifyError verifies verify error.
func TestStringCompare_VerifyError(t *testing.T) {
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		false,
		"expected",
		"actual",
	)
	actual := args.Map{"result": sc.VerifyError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatched should return error", actual)
}

// TestNewStringCompareEqual verifies constructor.
func TestNewStringCompareEqual(t *testing.T) {
	sc := coreinstruction.NewStringCompareEqual("a", "a")
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "equal should match", actual)
}

// TestNewStringCompareStartsWith verifies constructor.
func TestNewStringCompareStartsWith(t *testing.T) {
	sc := coreinstruction.NewStringCompareStartsWith(false, "hel", "hello")
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should start with 'hel'", actual)
}

// TestNewStringCompareEndsWith verifies constructor.
func TestNewStringCompareEndsWith(t *testing.T) {
	sc := coreinstruction.NewStringCompareEndsWith(false, "llo", "hello")
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should end with 'llo'", actual)
}

// TestNewStringCompareContains verifies constructor.
func TestNewStringCompareContains(t *testing.T) {
	sc := coreinstruction.NewStringCompareContains(false, "ell", "hello")
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain 'ell'", actual)
}

// TestNewStringCompareRegex verifies regex constructor.
func TestNewStringCompareRegex(t *testing.T) {
	sc := coreinstruction.NewStringCompareRegex("^he.*o$", "hello")
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "regex should match", actual)
}

// TestNewStringCompare_IgnoreCase verifies case-insensitive.
func TestNewStringCompare_IgnoreCase(t *testing.T) {
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		true,
		"HELLO",
		"hello",
	)
	actual := args.Map{"result": sc.IsMatch()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ignore case equal should match", actual)
}

// TestStringCompare_VerifyError_Regex verifies regex verify error.
func TestStringCompare_VerifyError_Regex(t *testing.T) {
	sc := coreinstruction.NewStringCompareRegex("^abc$", "abc")
	actual := args.Map{"result": sc.VerifyError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching regex should return nil error", actual)
	sc2 := coreinstruction.NewStringCompareRegex("^abc$", "xyz")
	actual := args.Map{"result": sc2.VerifyError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-matching regex should return error", actual)
}

// TestStringSearch_VerifyError_Regex verifies regex verify on StringSearch.
func TestStringSearch_VerifyError_Regex(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Regex,
		Search:        "^test$",
		BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
			IsIgnoreCase: false,
		},
	}
	actual := args.Map{"result": s.VerifyError("test") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching regex should return nil error", actual)
	actual := args.Map{"result": s.VerifyError("nope") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-matching regex should return error", actual)
}
