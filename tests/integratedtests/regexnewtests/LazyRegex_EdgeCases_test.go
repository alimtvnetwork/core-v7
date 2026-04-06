package regexnewtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// Nil receiver tests (migrated to CaseNilSafe)
// =============================================================================

func Test_NilLazyRegex_NilSafe(t *testing.T) {
	for caseIndex, tc := range lazyRegexNilReceiverTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Empty pattern edge cases
// =============================================================================

func Test_EmptyPattern_IsUndefined(t *testing.T) {
	tc := emptyPatternEdgeCaseTestCases[0]

	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	actual := args.Map{
		params.isUndefined: lazy.IsUndefined(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyPattern_IsNotApplicable(t *testing.T) {
	tc := emptyPatternEdgeCaseTestCases[1]

	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	actual := args.Map{
		params.isApplicable: lazy.IsApplicable(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyPattern_IsMatch_ReturnsFalse(t *testing.T) {
	tc := emptyPatternEdgeCaseTestCases[2]

	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	actual := args.Map{
		params.isMatch: lazy.IsMatch("anything"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyPattern_IsFailedMatch_ReturnsTrue(t *testing.T) {
	tc := emptyPatternEdgeCaseTestCases[3]

	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	actual := args.Map{
		params.isFailedMatch: lazy.IsFailedMatch("anything"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyPattern_Compile_ReturnsError(t *testing.T) {
	tc := emptyPatternEdgeCaseTestCases[4]

	// Arrange
	lazy := regexnew.New.Lazy("")

	// Act
	regex, err := lazy.Compile()

	actual := args.Map{
		params.hasError:   err != nil,
		params.regexIsNil: regex == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// Concurrency tests (these use channels by nature, kept as-is with assertions)
// =============================================================================

func Test_InvalidPattern_ConcurrentAccess(t *testing.T) {
	// Arrange
	invalidPatterns := []string{"[bad", "(unclosed", "*invalid", "(?P<>bad)"}
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			p := invalidPatterns[index%len(invalidPatterns)]
			lazy := regexnew.New.LazyLock(p)

			if lazy == nil {
				errors <- fmt.Sprintf("goroutine %d: LazyLock returned nil for invalid pattern %s", index, p)
				return
			}

			if lazy.IsApplicable() {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not be applicable", index, p)
			}

			if lazy.IsMatch("test") {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not match", index, p)
			}

			if !lazy.HasAnyIssues() {
				errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should have issues", index, p)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	for errMsg := range errors {

	// Act
		actual := args.Map{"error": errMsg}

	// Assert
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_InvalidPattern_ConcurrentCompileError(t *testing.T) {
	// Arrange
	pattern := "[broken"
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			lazy := regexnew.New.LazyLock(pattern)
			regex, err := lazy.Compile()

			if err == nil {
				errors <- fmt.Sprintf("goroutine %d: expected compile error", index)
			}

			if regex != nil {
				errors <- fmt.Sprintf("goroutine %d: expected nil regex", index)
			}

			if !lazy.HasError() {
				errors <- fmt.Sprintf("goroutine %d: HasError should be true", index)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	for errMsg := range errors {

	// Act
		actual := args.Map{"error": errMsg}

	// Assert
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_MixedValidInvalid_ConcurrentAccess(t *testing.T) {
	// Arrange
	patterns := []string{`\d+`, "[bad", `[a-z]+`, "(unclosed"}
	goroutineCount := 80
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			p := patterns[index%len(patterns)]
			lazy := regexnew.New.LazyLock(p)
			isValid := (index%len(patterns))%2 == 0

			if isValid {
				if !lazy.IsApplicable() {
					errors <- fmt.Sprintf("goroutine %d: valid pattern %s should be applicable", index, p)
				}
			} else {
				if lazy.IsApplicable() {
					errors <- fmt.Sprintf("goroutine %d: invalid pattern %s should not be applicable", index, p)
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	for errMsg := range errors {

	// Act
		actual := args.Map{"error": errMsg}

	// Assert
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}
