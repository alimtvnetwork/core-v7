package trydotests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/internal/trydo"
)

// ── WrappedErr.String() with both error AND exception (IsBothPresent path) ──

var covWrappedErrStringBothTestCases = []coretestcases.CaseV1{
	{
		Title: "WrappedErr.String returns combined string -- both error and exception present",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Error:     errors.New("test-error"),
				HasError:  true,
				Exception: "test-exception",
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"hasError":    true,
			"hasExc":      true,
			"nonEmpty":    true,
			"isBoth":      true,
			"containsErr": true,
		},
	},
}

func Test_WrappedErr_String_BothPresent(t *testing.T) {
	for caseIndex, testCase := range covWrappedErrStringBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		we := input["wrappedErr"].(*trydo.WrappedErr)

		// Act
		str := we.String()
		errStr := we.ErrorString()
		excStr := we.ExceptionString()

		// Assert
		actual := args.Map{
			"hasError":    errStr != "",
			"hasExc":      excStr != "",
			"nonEmpty":    str != "",
			"isBoth":      we.IsBothPresent(),
			"containsErr": len(str) > len(errStr),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── Throw function ──

func Test_Throw(t *testing.T) {
	// Arrange
	exception := trydo.WrapPanic(func() {
		trydo.Throw("thrown-value")
	})

	// Act
	actual := args.Map{
		"isNil": exception == nil,
	}

	// Assert
	expected := args.Map{
		"isNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Throw panics -- panics and is caught", actual)
}

// ── WrappedErr.HasErrorOrException with exception only ──

func Test_WrappedErr_HasErrorOrException_ExcOnly(t *testing.T) {
	// Arrange
	we := &trydo.WrappedErr{
		HasThrown: true,
		Exception: "exc",
	}

	// Act
	actual := args.Map{"result": we.HasErrorOrException()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasErrorOrException returns error -- exception only", actual)
}

// ── WrappedErr.String() exception only path ──

func Test_WrappedErr_String_ExcOnly(t *testing.T) {
	// Arrange
	we := &trydo.WrappedErr{
		Exception: "exc-val",
		HasThrown: true,
	}
	str := we.String()

	// Act
	actual := args.Map{"nonEmpty": str != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- exception only", actual)
}

// ── WrappedErr.ErrorString with nil error ──

func Test_WrappedErr_ErrorString_NilError(t *testing.T) {
	// Arrange
	we := &trydo.WrappedErr{}

	// Act
	actual := args.Map{"result": we.ErrorString()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorString returns nil -- nil error", actual)
}
