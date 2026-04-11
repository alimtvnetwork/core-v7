package casenilsafetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Test: Nil-safe pointer receiver methods
// =============================================================================

func Test_CaseNilSafe_PointerReceiverMethods(t *testing.T) {
	for caseIndex, tc := range nilSafePointerReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Void methods (no return values)
// =============================================================================

func Test_CaseNilSafe_VoidMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeVoidTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Multi-return methods
// =============================================================================

func Test_CaseNilSafe_MultiReturnMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeMultiReturnTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Unsafe methods (expected panics)
// =============================================================================

func Test_CaseNilSafe_UnsafeMethods(t *testing.T) {
	for caseIndex, tc := range nilUnsafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: MethodName extraction
// =============================================================================

func Test_CaseNilSafe_MethodName(t *testing.T) {
	for caseIndex, tc := range methodNameTestCases {
		// Arrange
		name := tc.MethodName()

		// Act
		actual := args.Map{
			"methodName": name,
		}

		// Assert — MethodName is not an invocation test,
		// so we compare directly via args.Map
		expected := args.Map{
			"methodName": tc.Expected.Value,
		}

		actLines := actual.CompileToStrings()
		expLines := expected.CompileToStrings()

		lineCountActual := args.Map{"lineCount": len(actLines)}
		lineCountExpected := args.Map{"lineCount": len(expLines)}
		lineCountExpected.ShouldBeEqual(t, caseIndex, "line count matches", lineCountActual)
		if len(actLines) != len(expLines) {
			continue
		}

		for i, line := range actLines {
			actual = args.Map{"result": line != expLines[i]}
			expected = args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "Case: got, want", actual)
		}
	}
}

// =============================================================================
// Test: CaseTitle fallback
// =============================================================================

func Test_CaseNilSafe_CaseTitleFallback(t *testing.T) {
	// Arrange
	tc := nilSafePointerReceiverTestCases[0]
	tcNoTitle := tc
	tcNoTitle.Title = ""

	// Act
	titleWithExplicit := tc.CaseTitle()
	titleFromMethod := tcNoTitle.CaseTitle()

	// Assert
	actual := args.Map{"result": titleWithExplicit != "IsValid on nil returns false"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected explicit title", actual)

	actual = args.Map{"result": titleFromMethod != "IsValid"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected method name fallback", actual)
}

// =============================================================================
// Test: Invoke with non-nil receiver
// =============================================================================

func Test_CaseNilSafe_InvokeWithReceiver(t *testing.T) {
	// Arrange
	tc := nilSafePointerReceiverTestCases[0] // IsValid
	receiver := &sampleStruct{Name: "hello", Value: 42}

	// Act
	result := tc.Invoke(receiver)

	// Assert
	actual := args.Map{"result": result.HasPanicked()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not panic with valid receiver", actual)

	actual = args.Map{"result": result.ValueString() != "true"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}
