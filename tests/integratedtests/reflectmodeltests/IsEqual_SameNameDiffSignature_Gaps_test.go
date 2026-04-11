package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── IsEqual: IsInvalid mismatch (line 163-165) ──
// This requires one valid, one with IsInvalid() == true.
// IsInvalid() returns true when receiver is nil.
// But we already test nil vs valid above...
// The tricky part: line 163 is after the nil checks (154-159)
// and same-pointer check (160-162). So both must be non-nil
// and different pointers, but one IsInvalid (which means nil).
// This is logically unreachable since nil is caught at line 157-158.

// ── IsEqual: InArgs mismatch (line 180-182) ──

func Test_IsEqual_SameNameDiffSignature(t *testing.T) {
	// Arrange
	pub := newMethodProcessor("PublicMethod")

	// Create a fake processor with same name but different method
	noArgs := newMethodProcessor("NoArgsMethod")
	fake := &reflectmodel.MethodProcessor{
		Name:          pub.Name,
		Index:         pub.Index,
		ReflectMethod: noArgs.ReflectMethod,
	}

	// Act
	result := pub.IsEqual(fake)

	// Assert
	actual := args.Map{
		"equal": result,
	}
	expected := args.Map{
		"equal": false,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsEqual returns false -- same name different signature",
		actual,
	)
}

// ── IsEqual: OutArgs mismatch (line 185-187) ──

type outArgsDiffStruct struct{}

func (s outArgsDiffStruct) MethodA(a string, b int) (string, error) { return "", nil }
func (s outArgsDiffStruct) MethodB(a string, b int) (int, error)    { return 0, nil }

func Test_IsEqual_SameInArgsDiffOutArgs(t *testing.T) {
	// Arrange
	typeA := reflect.TypeOf(outArgsDiffStruct{})
	methodA, _ := typeA.MethodByName("MethodA")
	methodB, _ := typeA.MethodByName("MethodB")

	procA := &reflectmodel.MethodProcessor{
		Name:          "SameMethod",
		Index:         methodA.Index,
		ReflectMethod: methodA,
	}
	procB := &reflectmodel.MethodProcessor{
		Name:          "SameMethod",
		Index:         methodB.Index,
		ReflectMethod: methodB,
	}

	// Act
	result := procA.IsEqual(procB)

	// Assert
	actual := args.Map{
		"equal": result,
	}
	expected := args.Map{
		"equal": false,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsEqual returns false -- same in-args different out-args",
		actual,
	)
}

// ── GetInArgsTypes: argsCount == 0 (line 229-231) ──
// In Go reflect, methods on concrete types always include the receiver.
// So NumIn() is always >= 1 for concrete type methods.
// This path is only reachable for MethodProcessor wrapping a non-method func type
// or through an invalid state. Document as accepted gap.

// ── GetInArgsTypesNames: argsCount == 0 (line 253-255) ──
// Same reasoning as above — accepted gap.

// ── validationError: IsInvalid branch (line 276-284) ──
// validationError is unexported. It's called by Invoke.
// IsInvalid() returns (it == nil), but if it's nil, line 272 catches it first.
// So line 276 is unreachable dead code. Accepted gap.

// ── ReflectValueToAnyValue: IsNull path (line 45-47) ──

func Test_Invoke_ReturnsNilInterface(t *testing.T) {
	// Arrange — invoke a method that returns nil values
	// Use Invoke which internally calls ReflectValuesToInterfaces
	proc := newMethodProcessor("PublicMethod")

	// Act — PublicMethod(a string, b int) (string, error)
	// receiver is sampleStruct
	responses, err := proc.Invoke(
		sampleStruct{},
		"test",
		42,
	)

	// Assert
	actual := args.Map{
		"hasResponses": len(responses) > 0,
		"err":          err,
	}
	expected := args.Map{
		"hasResponses": true,
		"err":          nil,
	}
	expected.ShouldBeEqual(
		t, 0,
		"Invoke returns responses -- valid method call",
		actual,
	)
}

// ── IsEqual: same pointer (line 160-161) ──

func Test_IsEqual_SamePointer(t *testing.T) {
	// Arrange
	proc := newMethodProcessor("PublicMethod")

	// Act
	result := proc.IsEqual(proc)

	// Assert
	actual := args.Map{
		"equal": result,
	}
	expected := args.Map{
		"equal": true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsEqual returns true -- same pointer",
		actual,
	)
}

// ── IsEqual: both nil ──

func Test_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *reflectmodel.MethodProcessor

	// Act
	result := a.IsEqual(b)

	// Assert
	actual := args.Map{
		"equal": result,
	}
	expected := args.Map{
		"equal": true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsEqual returns true -- both nil",
		actual,
	)
}

// ── IsNotEqual ──

func Test_IsNotEqual_DifferentProcessors(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")

	// Act
	result := a.IsNotEqual(b)

	// Assert
	actual := args.Map{
		"notEqual": result,
	}
	expected := args.Map{
		"notEqual": true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsNotEqual returns true -- different methods",
		actual,
	)
}

// ── IsEqual: identical methods ──

func Test_IsEqual_IdenticalMethods(t *testing.T) {
	// Arrange
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")

	// Act
	result := a.IsEqual(b)

	// Assert
	actual := args.Map{
		"equal": result,
	}
	expected := args.Map{
		"equal": true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"IsEqual returns true -- identical methods from same type",
		actual,
	)
}

// ── Invoke on nil processor triggers validationError nil check (line 272) ──

func Test_Invoke_NilProcessor(t *testing.T) {
	// Arrange
	var proc *reflectmodel.MethodProcessor

	// Act
	responses, err := proc.Invoke()

	// Assert
	actual := args.Map{
		"responses": responses == nil,
		"hasError":  err != nil,
	}
	expected := args.Map{
		"responses": true,
		"hasError":  true,
	}
	expected.ShouldBeEqual(
		t, 0,
		"Invoke returns error -- nil processor",
		actual,
	)
}
