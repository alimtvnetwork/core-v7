package coretestsresultstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── filterByFields: missing key triggers "<missing key: ...>" path ──
// Covers ResultAssert.go L92-94

func Test_ShouldMatchResult_MissingFieldKey(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 42, Panicked: false}
	exp := results.ResultAny{Panicked: false}

	// Act — pass a field name that doesn't exist in ToMap()
	// This triggers filterByFields' else branch (L92-94)
	r.ShouldMatchResult(t, 0, "missing field key", exp, "panicked", "nonexistent")
}

// ── MethodName: non-func input ──
// Covers MethodName.go L22-24

func Test_MethodName_NonFunc(t *testing.T) {
	// Arrange & Act
	name := results.MethodName("not-a-func")

	// Assert
	actual := args.Map{"name": name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "MethodName returns empty -- non-func input", actual)
}

// ── MethodName: nil input ──
// Covers MethodName.go L16-18

func Test_MethodName_Nil_FromShouldMatchResultMis(t *testing.T) {
	// Arrange & Act
	name := results.MethodName(nil)

	// Assert
	actual := args.Map{"name": name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "MethodName returns empty -- nil input", actual)
}

// ── MethodName: valid function ──

type cov4Struct struct{}

func (s *cov4Struct) Hello() string { return "hi" }

func Test_MethodName_ValidFunc(t *testing.T) {
	// Arrange & Act
	name := results.MethodName((*cov4Struct).Hello)

	// Assert
	actual := args.Map{"name": name}
	expected := args.Map{"name": "Hello"}
	expected.ShouldBeEqual(t, 0, "MethodName returns method name -- valid func", actual)
}

// ── InvokeWithPanicRecovery: non-func input ──
// Covers Invoke.go L48-53

func Test_Invoke_NonFunc(t *testing.T) {
	// Arrange & Act
	r := results.InvokeWithPanicRecovery("not-a-func", nil)

	// Assert
	actual := args.Map{
		"panicked":   r.Panicked,
		"panicValue": fmt.Sprintf("%v", r.PanicValue),
	}
	expected := args.Map{
		"panicked":   true,
		"panicValue": "funcRef is not a function: string",
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithPanicRecovery panics -- non-func input", actual)
}
