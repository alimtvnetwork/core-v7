package regexnewtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_LazyLock_ConcurrentAccess(t *testing.T) {
	// Arrange
	pattern := `\d+`
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			lazy := regexnew.New.LazyLock(pattern)
			if lazy == nil {
				errors <- fmt.Sprintf("goroutine %d: LazyLock returned nil", index)
				return
			}

			if !lazy.IsApplicable() {
				errors <- fmt.Sprintf("goroutine %d: not applicable", index)
				return
			}

			if !lazy.IsMatch("123") {
				errors <- fmt.Sprintf("goroutine %d: failed match", index)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		actual := args.Map{"error": errMsg}
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_LazyLock_ConcurrentDifferentPatterns(t *testing.T) {
	// Arrange
	patterns := []string{
		`\d+`, `[a-z]+`, `\w+`, `^hello$`, `\s+`,
		`[A-Z]{2,}`, `\d{3}-\d{4}`, `[0-9]+`, `^test`, `end$`,
	}
	goroutineCount := 100
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			p := patterns[index%len(patterns)]
			lazy := regexnew.New.LazyLock(p)

			if lazy == nil {
				errors <- fmt.Sprintf("goroutine %d: nil for pattern %s", index, p)
				return
			}

			if lazy.Pattern() != p {
				errors <- fmt.Sprintf("goroutine %d: pattern mismatch got %s want %s", index, lazy.Pattern(), p)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		actual := args.Map{"error": errMsg}
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_TwoLock_ConcurrentAccess(t *testing.T) {
	// Arrange
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount*2)

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			first, second := regexnew.New.LazyRegex.TwoLock(`\d+`, `[a-z]+`)

			if first == nil || second == nil {
				errors <- fmt.Sprintf("goroutine %d: TwoLock returned nil", index)
				return
			}

			if !first.IsMatch("42") {
				errors <- fmt.Sprintf("goroutine %d: first failed match", index)
			}

			if !second.IsMatch("abc") {
				errors <- fmt.Sprintf("goroutine %d: second failed match", index)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		actual := args.Map{"error": errMsg}
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_ManyUsingLock_ConcurrentAccess(t *testing.T) {
	// Arrange
	patterns := []string{`\d+`, `[a-z]+`, `\w+`, `^test$`}
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount*len(patterns))

	// Act
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			result := regexnew.New.LazyRegex.ManyUsingLock(patterns...)

			if len(result) != len(patterns) {
				errors <- fmt.Sprintf("goroutine %d: expected %d patterns got %d", index, len(patterns), len(result))
				return
			}

			for _, p := range patterns {
				lazy, exists := result[p]
				if !exists || lazy == nil {
					errors <- fmt.Sprintf("goroutine %d: missing pattern %s", index, p)
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		actual := args.Map{"error": errMsg}
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}

func Test_LazyLock_ConcurrentCompileAndMatch(t *testing.T) {
	// Arrange
	pattern := `^[a-z]+\d+$`
	goroutineCount := 50
	wg := sync.WaitGroup{}
	wg.Add(goroutineCount)
	errors := make(chan string, goroutineCount)

	// Act - half compile, half match simultaneously
	for i := 0; i < goroutineCount; i++ {
		go func(index int) {
			defer wg.Done()

			lazy := regexnew.New.LazyLock(pattern)

			if index%2 == 0 {
				regex, err := lazy.Compile()
				if err != nil || regex == nil {
					errors <- fmt.Sprintf("goroutine %d: compile failed", index)
				}
			} else {
				if !lazy.IsMatch("abc123") {
					errors <- fmt.Sprintf("goroutine %d: match failed", index)
				}
				if lazy.IsMatch("123abc") {
					errors <- fmt.Sprintf("goroutine %d: should not match", index)
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Assert
	for errMsg := range errors {
		actual := args.Map{"error": errMsg}
		expected := args.Map{"error": ""}
		expected.ShouldBeEqual(t, 0, "concurrent operation should not error", actual)
	}
}
