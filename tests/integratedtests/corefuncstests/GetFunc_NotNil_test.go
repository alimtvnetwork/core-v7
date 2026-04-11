package corefuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── GetFunc / GetFuncFullName / GetFuncName ──

func Test_GetFunc_NotNil(t *testing.T) {
	// Arrange
	f := corefuncs.GetFunc(corefuncs.GetFuncName)

	// Act
	actual := args.Map{"notNil": f != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetFunc returns nil -- returns non-nil for valid func", actual)
}

func Test_GetFuncFullName_HasDot(t *testing.T) {
	// Arrange
	name := corefuncs.GetFuncFullName(corefuncs.GetFuncName)

	// Act
	actual := args.Map{"notEmpty": name != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetFuncFullName returns empty -- returns non-empty", actual)
}

func Test_GetFuncName_Short(t *testing.T) {
	// Arrange
	name := corefuncs.GetFuncName(corefuncs.GetFuncName)

	// Act
	actual := args.Map{"result": name}

	// Assert
	expected := args.Map{"result": "GetFuncName"}
	expected.ShouldBeEqual(t, 0, "GetFuncName returns correct value -- returns short name", actual)
}

// ── NamedActionFuncWrapper.Next ──

func Test_NamedAction_Next(t *testing.T) {
	// Arrange
	var firstCalled, secondCalled bool
	w1 := corefuncs.New.NamedAction("first", func(name string) {
		firstCalled = true
	})
	w2 := corefuncs.New.NamedAction("second", func(name string) {
		secondCalled = true
	})
	w1.Next(&w2)

	// Act
	actual := args.Map{
		"first": firstCalled,
		"second": secondCalled,
	}

	// Assert
	expected := args.Map{
		"first": true,
		"second": true,
	}
	expected.ShouldBeEqual(t, 0, "NamedAction Next -- both called", actual)
}

// ── NamedActionFuncWrapper.Exec / AsActionReturnsErrorFunc ──

func Test_NamedAction_Exec(t *testing.T) {
	// Arrange
	var called bool
	w := corefuncs.New.NamedAction("test", func(name string) {
		called = true
	})
	w.Exec()

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "NamedAction Exec -- called", actual)
}

func Test_NamedAction_AsActionReturnsErrorFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.NamedAction("test", func(name string) {})
	errFunc := w.AsActionReturnsErrorFunc()
	err := errFunc()

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "NamedAction AsActionReturnsErrorFunc -- nil error", actual)
}

// ── IsSuccessFuncWrapper ──

func Test_IsSuccess_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.New.IsSuccess("check", func() bool { return true })

	// Act
	actual := args.Map{"result": w.Exec()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsSuccess Exec -- true", actual)
}

func Test_IsSuccess_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.IsSuccess("check", func() bool { return true })
	w.AsActionFunc()()

	// Act
	actual := args.Map{"called": true}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "IsSuccess AsActionFunc -- no panic", actual)
}

func Test_IsSuccess_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.IsSuccess("check", func() bool { return false })
	err := w.AsActionReturnsErrorFunc()()

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "IsSuccess failure -- returns error", actual)
}

func Test_IsSuccess_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.IsSuccess("check", func() bool { return true })
	err := w.AsActionReturnsErrorFunc()()

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "IsSuccess success -- nil error", actual)
}

// ── ActionReturnsErrorFuncWrapper.Exec ──

func Test_ActionErr_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return nil })
	err := w.Exec()

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ActionErr Exec -- nil error", actual)
}

func Test_ActionErr_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return errors.New("fail") })
	err := w.AsActionReturnsErrorFunc()()

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "ActionErr failure -- wraps error", actual)
}

// ── ToLegacy conversions ──

func Test_InOutErrWrapperOf_ToLegacy(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int](
		"test", func(s string) (int, error) { return len(s), nil },
	)
	legacy := w.ToLegacy()
	out, err := legacy.Exec("hi")

	// Act
	actual := args.Map{
		"output": out,
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"output": 2,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf returns error -- ToLegacy", actual)
}

func Test_InActionErrWrapperOf_ToLegacy(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string](
		"test", func(s string) error { return nil },
	)
	legacy := w.ToLegacy()
	_, err := legacy.Exec("hi")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf returns error -- ToLegacy", actual)
}

func Test_InOutFuncWrapperOf_ToLegacy(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int](
		"test", func(s string) int { return len(s) },
	)
	legacy := w.ToLegacy()
	out, err := legacy.Exec("hi")

	// Act
	actual := args.Map{
		"output": out,
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"output": 2,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf returns correct value -- ToLegacy", actual)
}

func Test_ResultDelegatingWrapperOf_ToLegacy(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string](
		"test", func(t *string) error { return nil },
	)
	legacy := w.ToLegacy()
	var s string
	err := legacy.Exec(&s)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf returns correct value -- ToLegacy", actual)
}

// ── InOutErrFuncWrapper.Exec ──

func Test_LegacyInOutErr_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return "ok", nil
	})
	out, err := w.Exec("input")

	// Act
	actual := args.Map{
		"output": out,
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"output": "ok",
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr returns error -- Exec", actual)
}

// ── ResultDelegatingFuncWrapper.Exec ──

func Test_LegacyResultDelegating_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return nil
	})
	err := w.Exec("target")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating returns correct value -- Exec", actual)
}

// ── SerializeOutputFuncWrapperOf.Exec ──

func Test_SerializeWrapper_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.NewSerializeWrapper[string](
		"json", func(s string) ([]byte, error) { return []byte(s), nil },
	)
	bytes, err := w.Exec("hello")

	// Act
	actual := args.Map{
		"result": string(bytes),
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": "hello",
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeWrapper returns correct value -- Exec", actual)
}

// ── InActionReturnsErrFuncWrapperOf.Exec / AsActionFunc ──

func Test_InActionErrWrapper_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string](
		"validate", func(s string) error { return nil },
	)
	err := w.Exec("test")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapper returns error -- Exec", actual)
}

// ── InOutErrFuncWrapperOf.Exec ──

func Test_InOutErrWrapperOf_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int](
		"parse", func(s string) (int, error) { return len(s), nil },
	)
	out, err := w.Exec("abc")

	// Act
	actual := args.Map{
		"output": out,
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"output": 3,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf returns error -- Exec", actual)
}

// ── InOutFuncWrapperOf.Exec ──

func Test_InOutFuncWrapperOf_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int](
		"strlen", func(s string) int { return len(s) },
	)
	out := w.Exec("test")

	// Act
	actual := args.Map{"output": out}

	// Assert
	expected := args.Map{"output": 4}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf returns correct value -- Exec", actual)
}

// ── ResultDelegatingFuncWrapperOf.Exec ──

func Test_ResultDelegatingWrapperOf_Exec(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string](
		"bind", func(t *string) error { *t = "bound"; return nil },
	)
	var s string
	err := w.Exec(&s)

	// Act
	actual := args.Map{
		"value": s,
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"value": "bound",
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf returns correct value -- Exec", actual)
}
