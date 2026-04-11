package resultstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── safeInterface: invalid reflect.Value ──

type voidMethod struct{}

func (v voidMethod) VoidReturn() {}

func Test_Invoke_VoidMethod_ReturnsNilValue(t *testing.T) {
	// Arrange
	funcRef := (voidMethod).VoidReturn
	receiver := voidMethod{}

	// Act
	r := results.InvokeWithPanicRecovery(funcRef, receiver)

	// Assert
	actual := args.Map{
		"panicked":    r.Panicked,
		"value":       r.Value,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked":    false,
		"value":       nil,
		"returnCount": 0,
	}
	expected.ShouldBeEqual(
		t, 0,
		"InvokeWithPanicRecovery returns nil value -- void method",
		actual,
	)
}

// ── extractErrorFromValue: non-error interface return ──

type multiNonErrorReturn struct{}

func (m multiNonErrorReturn) TwoStrings() (string, string) {
	return "a", "b"
}

func Test_Invoke_NonErrorSecondReturn(t *testing.T) {
	// Arrange
	funcRef := (multiNonErrorReturn).TwoStrings
	receiver := multiNonErrorReturn{}

	// Act
	r := results.InvokeWithPanicRecovery(funcRef, receiver)

	// Assert
	actual := args.Map{
		"panicked":    r.Panicked,
		"error":       r.Error,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked":    false,
		"error":       nil,
		"returnCount": 2,
	}
	expected.ShouldBeEqual(
		t, 0,
		"InvokeWithPanicRecovery returns nil error -- non-error second return",
		actual,
	)
}

// ── extractErrorFromValue: nil pointer return that implements error ──

type nilPtrErrorReturn struct{}

func (n nilPtrErrorReturn) NilPtrErr() (*customErr, string) {
	return nil, "ok"
}

type customErr struct{ msg string }

func (e *customErr) Error() string { return e.msg }

func Test_Invoke_NilPtrErrorReturn(t *testing.T) {
	// Arrange
	funcRef := (nilPtrErrorReturn).NilPtrErr
	receiver := nilPtrErrorReturn{}

	// Act
	r := results.InvokeWithPanicRecovery(funcRef, receiver)

	// Assert
	actual := args.Map{
		"panicked":    r.Panicked,
		"error":       r.Error,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked":    false,
		"error":       nil,
		"returnCount": 2,
	}
	expected.ShouldBeEqual(
		t, 0,
		"InvokeWithPanicRecovery returns nil error -- nil ptr implementing error",
		actual,
	)
}

// ── MethodName: function without dots in runtime name ──

func Test_MethodName_NilFuncRef(t *testing.T) {
	// Arrange / Act
	name := results.MethodName(nil)

	// Assert
	actual := args.Map{
		"name": name,
	}
	expected := args.Map{
		"name": "",
	}
	expected.ShouldBeEqual(
		t, 0,
		"MethodName returns empty -- nil funcRef",
		actual,
	)
}

func Test_MethodName_NonFuncRef(t *testing.T) {
	// Arrange / Act
	name := results.MethodName(42)

	// Assert
	actual := args.Map{
		"name": name,
	}
	expected := args.Map{
		"name": "",
	}
	expected.ShouldBeEqual(
		t, 0,
		"MethodName returns empty -- non-func ref",
		actual,
	)
}

func Test_MethodName_ValidMethod(t *testing.T) {
	// Arrange / Act
	name := results.MethodName((voidMethod).VoidReturn)

	// Assert
	actual := args.Map{
		"name": name,
	}
	expected := args.Map{
		"name": "VoidReturn",
	}
	expected.ShouldBeEqual(
		t, 0,
		"MethodName returns method name -- valid method expression",
		actual,
	)
}

// ── safeInterface: nil pointer in return value ──

type nilPtrReturn struct{}

func (n nilPtrReturn) ReturnNilPtr() *int {
	return nil
}

func Test_Invoke_NilPtrReturnValue(t *testing.T) {
	// Arrange
	funcRef := (nilPtrReturn).ReturnNilPtr
	receiver := nilPtrReturn{}

	// Act
	r := results.InvokeWithPanicRecovery(funcRef, receiver)

	// Assert
	actual := args.Map{
		"panicked":    r.Panicked,
		"value":       r.Value,
		"returnCount": r.ReturnCount,
	}
	expected := args.Map{
		"panicked":    false,
		"value":       nil,
		"returnCount": 1,
	}
	expected.ShouldBeEqual(
		t, 0,
		fmt.Sprintf(
			"InvokeWithPanicRecovery extracts nil -- nil pointer return value"),
		actual,
	)
}

// ── Nil receiver causing panic ──

type ptrReceiver struct{ val int }

func (p *ptrReceiver) GetVal() int { return p.val }

func Test_Invoke_NilReceiver_PanicsRecovery(t *testing.T) {
	// Arrange
	funcRef := (*ptrReceiver).GetVal

	// Act
	r := results.InvokeWithPanicRecovery(funcRef, nil)

	// Assert
	actual := args.Map{
		"panicked": r.Panicked,
	}
	expected := args.Map{
		"panicked": true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"InvokeWithPanicRecovery panics -- nil receiver on value method",
		actual,
	)
}
