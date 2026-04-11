package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_AnyFunctionsExecuteResults_False(t *testing.T) {
	// Arrange
	result := conditional.AnyFunctionsExecuteResults(
		false,
		nil,
		[]func() (any, bool, bool){
			func() (any, bool, bool) { return "b", true, false },
		},
	)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_TypedErrorFunctionsExecuteResults_False(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		false,
		nil,
		[]func() (int, error){
			func() (int, error) { return 42, nil },
		},
	)

	// Act
	actual := args.Map{"result": err != nil || len(results) != 1 || results[0] != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_TypedErrorFunctionsExecuteResults_WithError(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		[]func() (int, error){
			func() (int, error) { return 0, errors.New("fail") },
			nil,
			func() (int, error) { return 1, nil },
		},
		nil,
	)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual = args.Map{"result": len(results) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_TypedErrorFunctionsExecuteResults_Empty(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](
		true,
		nil,
		nil,
	)

	// Act
	actual := args.Map{"result": err != nil || results != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}
