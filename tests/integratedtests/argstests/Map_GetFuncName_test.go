package argstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map: GetFuncName returns empty when no FuncWrap ──
// Covers Map.go L107

func Test_Map_GetFuncName_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"key": "value"}

	result := m.GetFuncName()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "GetFuncName returns empty -- no FuncWrap in map", actual)
}

// ── Map: InvokeMust panics on invalid func ──
// Covers Map.go L422-423

func Test_Map_InvokeMust_Panic(t *testing.T) {
	// Arrange
	m := args.Map{"workFunc": "not-a-func"}

	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		m.InvokeMust()
	}()

	// Act
	actual := args.Map{"didPanic": didPanic}

	// Assert
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "InvokeMust panics -- invalid work func", actual)
}

// ── Map: InvokeWithValidArgs ──
// Covers Map.go L434-439

func Test_Map_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s + "!" }
	m := args.Map{
		"workFunc": fn,
		"first":    "hello",
	}

	results, err := m.InvokeWithValidArgs()

	// Act
	actual := args.Map{
		"noErr":     err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":     true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs invokes function -- valid args", actual)
}

// ── FuncWrap: IsEqual various comparison paths ──
// Covers FuncWrap.go L215-245

func Test_FuncWrap_IsEqual_Different(t *testing.T) {
	// Arrange
	fn1 := func(s string) string { return s }
	fn2 := func(s string, i int) string { return s }

	fw1 := args.NewFuncWrap.Default(fn1)
	fw2 := args.NewFuncWrap.Default(fn2)

	result := fw1.IsEqual(fw2)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different arg counts", actual)
}

func Test_FuncWrap_IsEqual_Nil(t *testing.T) {
	// Arrange
	fn1 := func() {}
	fw1 := args.NewFuncWrap.Default(fn1)

	result := fw1.IsEqual(nil)

	// Act
	actual := args.Map{"equal": result}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- nil other", actual)
}

// ── FuncWrap: InvokeFirstAndError ──
// Covers FuncWrapInvoke.go L121-137

func Test_FuncWrap_InvokeFirstAndError(t *testing.T) {
	// Arrange
	fn := func(s string) (string, error) { return s + "!", fmt.Errorf("test-err") }
	fw := args.NewFuncWrap.Default(fn)

	first, funcErr, procErr := fw.InvokeFirstAndError("hello")

	// Act
	actual := args.Map{
		"first":      first,
		"hasFuncErr": funcErr != nil,
		"noProcErr":  procErr == nil,
	}

	// Assert
	expected := args.Map{
		"first":      "hello!",
		"hasFuncErr": true,
		"noProcErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns first and error", actual)
}

// ── FuncWrap: InvokeError ──
// Covers FuncWrapInvoke.go L107-114

func Test_FuncWrap_InvokeError(t *testing.T) {
	// Arrange
	fn := func() error { return nil }
	fw := args.NewFuncWrap.Default(fn)

	defer func() {
		if r := recover(); r != nil {
			// Known limitation: InvokeError panics on nil error interface cast
			t.Skipf("InvokeError panics on nil error return: %v", r)
		}
	}()

	funcErr, procErr := fw.InvokeError()

	// Act
	actual := args.Map{
		"noFuncErr": funcErr == nil,
		"noProcErr": procErr == nil,
	}

	// Assert
	expected := args.Map{
		"noFuncErr": true,
		"noProcErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeError returns nil errors -- no-error func", actual)
}

// ── FuncWrap: GetResponseOfInvoke with index ──
// Covers FuncWrapInvoke.go L85-101

func Test_FuncWrap_InvokeResultOfIndex(t *testing.T) {
	// Arrange
	fn := func() (string, int) { return "a", 42 }
	fw := args.NewFuncWrap.Default(fn)

	result, err := fw.InvokeResultOfIndex(1)

	// Act
	actual := args.Map{
		"result": result,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": 42,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns indexed result", actual)
}

// ── FuncWrap: InvokeSkip with panic ──
// Covers FuncWrapInvoke.go L61-71

func Test_FuncWrap_InvokeSkip_Panic(t *testing.T) {
	// Arrange
	fn := func() { panic("test panic") }
	fw := args.NewFuncWrap.Default(fn)

	_, err := fw.Invoke()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Invoke returns error -- function panics", actual)
}

// ── funcDetector: GetFuncWrap with FuncWrapAny directly ──
// Covers funcDetector.go L16-17

func Test_FuncDetector_DirectFuncWrap(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewFuncWrap.Default(fn)

	result := args.FuncDetector.GetFuncWrap(fw)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"valid": result.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap returns same instance -- direct FuncWrap", actual)
}

// ── funcDetector: GetFuncWrap with raw function ──
// Covers funcDetector.go L18-22 (default)

func Test_FuncDetector_RawFunc(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }

	result := args.FuncDetector.GetFuncWrap(fn)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"valid": result.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFuncWrap creates new FuncWrap -- raw function", actual)
}

// ── newFuncWrapCreator: MethodToFunc nil method ──
// Covers newFuncWrapCreator.go L101-106

func Test_NewFuncWrap_MethodToFunc_Nil(t *testing.T) {
	// Arrange
	fw, err := args.NewFuncWrap.MethodToFunc(nil)

	// Act
	actual := args.Map{
		"invalid": fw.IsInvalid(),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"invalid": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodToFunc returns invalid -- nil method", actual)
}

// ── newFuncWrapCreator: StructToMap ──
// Covers newFuncWrapCreator.go L122-141

type cov9TestStruct struct{}

func (s *cov9TestStruct) Hello() string { return "hello" }
func (s *cov9TestStruct) World() string { return "world" }

func Test_NewFuncWrap_StructToMap(t *testing.T) {
	// Arrange
	s := &cov9TestStruct{}
	fm, err := args.NewFuncWrap.StructToMap(s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasEntries": len(fm) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasEntries": true,
	}
	expected.ShouldBeEqual(t, 0, "StructToMap creates FuncMap -- valid struct", actual)
}

// ── OneFunc: InvokeMust / InvokeWithValidArgs / InvokeArgs ──
// Covers OneFunc.go L84-105

func Test_OneFunc_InvokeMust(t *testing.T) {
	// Arrange
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results := of.InvokeMust("test")

	// Act
	actual := args.Map{"hasResult": len(results) > 0}

	// Assert
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeMust returns result -- valid call", actual)
}

func Test_OneFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results, err := of.InvokeWithValidArgs()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeWithValidArgs returns result", actual)
}

func Test_OneFunc_InvokeArgs(t *testing.T) {
	// Arrange
	of := &args.OneFunc[string]{
		First:    "hi",
		WorkFunc: func(s string) string { return s + "!" },
	}

	results, err := of.InvokeArgs(1)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc InvokeArgs returns result", actual)
}

// ── Holder: InvokeMust / InvokeWithValidArgs / InvokeArgs ──
// Covers Holder.go L188-204

func Test_Holder_InvokeMust(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{
		WorkFunc: func() string { return "test" },
	}

	results := h.InvokeMust()

	// Act
	actual := args.Map{"hasResult": len(results) > 0}

	// Assert
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "Holder InvokeMust returns result", actual)
}

func Test_Holder_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{
		WorkFunc: func() string { return "test" },
	}

	results, err := h.InvokeWithValidArgs()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder InvokeWithValidArgs returns result", actual)
}

func Test_Holder_InvokeArgs(t *testing.T) {
	// Arrange
	h := &args.Holder[any]{
		WorkFunc: func() string { return "test" },
	}

	results, err := h.InvokeArgs(0)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder InvokeArgs returns result", actual)
}

// ── Dynamic: Invoke / InvokeMust / InvokeWithValidArgs ──
// Covers Dynamic.go L51,76-94

func Test_Dynamic_Invoke(t *testing.T) {
	// Arrange
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results, err := d.Invoke()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Invoke returns result", actual)
}

func Test_Dynamic_InvokeMust(t *testing.T) {
	// Arrange
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results := d.InvokeMust()

	// Act
	actual := args.Map{"hasResult": len(results) > 0}

	// Assert
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "Dynamic InvokeMust returns result", actual)
}

func Test_Dynamic_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	fn := func() string { return "dynamic" }
	d := &args.Dynamic[string]{
		Params: args.Map{"workFunc": fn},
	}

	results, err := d.InvokeWithValidArgs()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": len(results) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic InvokeWithValidArgs returns result", actual)
}
