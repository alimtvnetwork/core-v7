package args

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/internal/convertinternal"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/internal/trydo"
)

// VoidCallNoReturn invokes the function ignoring return values.
func (it *FuncWrap[T]) VoidCallNoReturn(args ...any) (processingErr error) {
	it.MustBeValid()
	_, err := it.Invoke(args...)

	return err
}

// InvokeMust invokes the function, panicking on error.
func (it *FuncWrap[T]) InvokeMust(args ...any) []any {
	results, err := it.Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}

// Invoke dynamically calls the wrapped function with the given arguments.
// Returns the results as []any and any processing error.
func (it *FuncWrap[T]) Invoke(args ...any) (results []any, processingErr error) {
	return it.InvokeSkip(codestack.Skip1, args...)
}

// InvokeSkip invokes the function with a custom stack skip for error reporting.
func (it *FuncWrap[T]) InvokeSkip(
	skipStack int,
	args ...any,
) (results []any, processingErr error) {
	firstErr := it.ValidationError()

	if firstErr != nil {
		return nil, firstErr
	}

	argsValidationErr := it.ValidateMethodArgs(args)

	if argsValidationErr != nil {
		return nil, argsValidationErr
	}

	rvs := argsToRvFunc(args)
	var resultsRawValues []reflect.Value
	exception := trydo.WrapPanic(func() { resultsRawValues = it.rv.Call(rvs) })

	if exception != nil {
		toMsg := convertinternal.AnyTo.SmartString(exception)
		finalError := fmt.Errorf(
			"%s - func invoke failed\nstack-trace:%s\nerr:%s",
			it.GetFuncName(),
			reflectinternal.CodeStack.StacksString(codestack.Skip1+skipStack),
			toMsg,
		)

		return rvToInterfacesFunc(resultsRawValues), finalError
	}

	return rvToInterfacesFunc(resultsRawValues), nil
}

// VoidCall invokes the function with no arguments.
func (it *FuncWrap[T]) VoidCall() ([]any, error) {
	return it.Invoke()
}

// GetFirstResponseOfInvoke invokes the function and returns only the first result.
func (it *FuncWrap[T]) GetFirstResponseOfInvoke(args ...any) (firstResponse any, err error) {
	result, err := it.InvokeResultOfIndex(0, args...)

	if err != nil {
		return nil, err
	}

	return result, err
}

// InvokeResultOfIndex invokes the function and returns the result at the given index.
func (it *FuncWrap[T]) InvokeResultOfIndex(
	index int,
	args ...any,
) (firstResponse any, err error) {
	results, err := it.Invoke(args...)

	if err != nil {
		return nil, err
	}

	return results[index], err
}

// InvokeError invokes the function and returns the first result as an error.
func (it *FuncWrap[T]) InvokeError(args ...any) (funcErr, processingErr error) {
	result, err := it.GetFirstResponseOfInvoke(args...)

	if err != nil {
		return nil, err
	}

	return result.(error), err
}

// InvokeFirstAndError is useful for methods returning (something, error).
// It invokes the function and separates the first result from the error result.
func (it *FuncWrap[T]) InvokeFirstAndError(
	args ...any,
) (firstResponse any, funcErr, processingErr error) {
	results, processingErr := it.Invoke(args...)

	if processingErr != nil {
		return nil, nil, processingErr
	}

	if len(results) <= 1 {
		return results, nil, errors.New(
			it.GetFuncName() + " doesn't return at least 2 return args",
		)
	}

	first := results[0]
	second, _ := results[1].(error)

	return first, second, processingErr
}

// argsCountMismatchErrorMessage builds a detailed error message for argument count mismatches.
func (it *FuncWrap[T]) argsCountMismatchErrorMessage(
	expectedCount int,
	given int,
	args []any,
) string {
	expectedTypes := it.GetInArgsTypesNames()
	expectedToNames := strings.Join(expectedTypes, newLineSpaceIndent)
	actualTypes := reflectinternal.Converter.InterfacesToTypesNamesWithValues(args)
	actualTypesName := strings.Join(actualTypes, newLineSpaceIndent)

	return fmt.Sprintf(
		"%s [Func] =>\n  arguments count doesn't match for - count:\n    expected : %d\n    given    : %d\n  expected types listed :\n    - %s\n  actual given types list :\n    - %s",
		it.Name,
		expectedCount,
		given,
		expectedToNames,
		actualTypesName,
	)
}
