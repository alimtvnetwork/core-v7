package codefuncstests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

func sampleFunc() {}

// =============================================================================
// GetFuncName — positive
// =============================================================================

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncNameTestCases {
		// Act
		name := corefuncs.GetFuncName(sampleFunc)
		isNotEmpty := fmt.Sprintf("%v", name != "")

		// Assert
		tc.ShouldBeEqual(t, caseIndex, isNotEmpty)
	}
}

// =============================================================================
// GetFuncName — nil input (panic recovery)
// =============================================================================

func Test_GetFuncName_NilInput(t *testing.T) {
	tc := getFuncNameNilTestCase

	// Arrange
	var panicked bool
	var result string

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		result = corefuncs.GetFuncName(nil)
	}()

	actual := args.Map{
		"result":   result,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// GetFuncName — non-function input (panic recovery)
// =============================================================================

func Test_GetFuncName_NonFuncInput(t *testing.T) {
	tc := getFuncNameNonFuncTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	var panicked bool
	var result string

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		result = corefuncs.GetFuncName(inputVal)
	}()

	actual := args.Map{
		"result":   result,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// GetFuncFullName — positive
// =============================================================================

func Test_GetFuncFullName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncFullNameTestCases {
		// Act
		fullName := corefuncs.GetFuncFullName(sampleFunc)
		actual := args.Map{
			"isNotEmpty":      fullName != "",
			"containsPackage": strings.Contains(fullName, "codefuncstests"),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// GetFuncFullName — nil input (panic recovery)
// =============================================================================

func Test_GetFuncFullName_NilInput(t *testing.T) {
	tc := getFuncFullNameNilTestCase

	// Arrange
	var panicked bool
	var result string

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		result = corefuncs.GetFuncFullName(nil)
	}()

	actual := args.Map{
		"result":   result,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// GetFuncFullName — non-function input (panic recovery)
// =============================================================================

func Test_GetFuncFullName_NonFuncInput(t *testing.T) {
	tc := getFuncFullNameNonFuncTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	var panicked bool
	var result string

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		result = corefuncs.GetFuncFullName(inputVal)
	}()

	actual := args.Map{
		"result":   result,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// GetFunc — positive
// =============================================================================

func Test_GetFunc_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncTestCases {
		// Act
		f := corefuncs.GetFunc(sampleFunc)
		actual := args.Map{
			"isNotNil": f != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// GetFunc — nil input (panic recovery)
// =============================================================================

func Test_GetFunc_NilInput(t *testing.T) {
	tc := getFuncNilTestCase

	// Arrange
	var panicked bool
	var isNil bool

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		f := corefuncs.GetFunc(nil)
		isNil = f == nil
	}()

	actual := args.Map{
		"isNil":    isNil,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// GetFunc — non-function input (panic recovery)
// =============================================================================

func Test_GetFunc_NonFuncInput(t *testing.T) {
	tc := getFuncNonFuncTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	var panicked bool
	var isNil bool

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()

		f := corefuncs.GetFunc(inputVal)
		isNil = f == nil
	}()

	actual := args.Map{
		"isNil":    isNil,
		"panicked": panicked,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// =============================================================================
// newCreator — factory methods
// =============================================================================

func Test_NewCreator_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		name, _ := input.GetAsString("name")

		var actual args.Map

		// Act
		switch method {
		case "ActionErr":
			w := corefuncs.New.ActionErr(name, func() error { return nil })
			err := w.Exec()
			actual = args.Map{
				"hasError": err != nil,
			}
		case "IsSuccess":
			w := corefuncs.New.IsSuccess(name, func() bool { return true })
			actual = args.Map{
				"result": w.Exec(),
			}
		case "NamedAction":
			tracker := &namedActionTracker{}
			w := corefuncs.New.NamedAction(name, tracker.Action)
			w.Exec()
			actual = args.Map{
				"calledWith": tracker.CalledWith,
			}
		case "LegacyInOutErr":
			w := corefuncs.New.LegacyInOutErr(name, func(in any) (any, error) {
				return "processed", nil
			})
			output, err := w.Exec("input")
			actual = args.Map{
				"output":   output.(string),
				"hasError": err != nil,
			}
		case "LegacyResultDelegating":
			w := corefuncs.New.LegacyResultDelegating(name, func(_ any) error {
				return nil
			})
			err = w.Exec("target")
			actual = args.Map{
				"hasError": err != nil,
			}
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
