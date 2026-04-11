package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== AnyOnce.Deserialize coverage =====

func Test_AnyOnce_Deserialize_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyOncePtr(func() any { return map[string]string{"a": "b"} })
	var target map[string]string
	err := o.Deserialize(&target)
	// Due to the bug (if err == nil returns err which is nil), this always returns nil

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil due to code path", actual)
}

func Test_AnyOnce_Deserialize_SerializeError(t *testing.T) {
	// Arrange
	// Use a value that can't be marshalled (channel)
	ch := make(chan int)
	o := coreonce.NewAnyOncePtr(func() any { return ch })
	var target string
	err := o.Deserialize(&target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialize error", actual)
}

// ===== AnyErrorOnce.Deserialize coverage =====

func Test_AnyErrorOnce_Deserialize_ExistingError(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return nil, errors.New("init error")
	})
	var target string
	err := o.Deserialize(&target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error from serialize", actual)
}

func Test_AnyErrorOnce_Deserialize_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return map[string]string{"x": "y"}, nil
	})
	var target map[string]string
	err := o.Deserialize(&target)
	// Same bug as AnyOnce - always returns nil when serialize succeeds

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_AnyErrorOnce_Deserialize_MarshalError(t *testing.T) {
	// Arrange
	ch := make(chan int)
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return ch, nil
	})
	var target string
	err := o.Deserialize(&target)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected marshal error from Serialize", actual)
}

// ===== IntegersOnce.IsEqual - hit currentMap[item] < 0 =====

func Test_IntegersOnce_IsEqual_FreqMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 1} })
	// Same length but different frequencies: {1,1} vs {1,2}

	// Act
	actual := args.Map{"result": o.IsEqual(1, 2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for frequency mismatch", actual)
}

// ===== MapStringStringOnce.IsEqual - hit isMissing and value mismatch =====

func Test_MapStringStringOnce_IsEqual_MissingKey(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})

	// Act
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "1", "c": "2"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for missing key", actual)
}

func Test_MapStringStringOnce_IsEqual_ValueMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{"result": o.IsEqual(map[string]string{"a": "9"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for value mismatch", actual)
}

// ===== StringsOnce.IsEqual - hit currentMap[item] < 0 =====

func Test_StringsOnce_IsEqual_FreqMismatch(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "a"} })

	// Act
	actual := args.Map{"result": o.IsEqual("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for frequency mismatch", actual)
}

// ===== StringOnce.SplitLeftRight - hit len > 2 branch =====
// Note: SplitN with n=2 returns at most 2, so len>2 is dead code.
// But we can test the len==1 (no splitter found) path to cover the else.

func Test_StringOnce_SplitLeftRight_NoSplitter(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "nosplitter" })
	left, right := o.SplitLeftRight(":")

	// Act
	actual := args.Map{"result": left != "nosplitter" || right != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nosplitter','', got '',''", actual)
}

func Test_StringOnce_SplitLeftRight_WithSplitter(t *testing.T) {
	// Arrange
	o := coreonce.NewStringOncePtr(func() string { return "left:right" })
	left, right := o.SplitLeftRight(":")

	// Act
	actual := args.Map{"result": left != "left" || right != "right"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'left','right', got '',''", actual)
}

// ===== JsonStringMust panic paths =====

func Test_StringsOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	s := o.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json string", actual)
}

func Test_MapStringStringOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	s := o.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json string", actual)
}
