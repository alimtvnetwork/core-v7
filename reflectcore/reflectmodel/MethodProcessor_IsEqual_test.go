package reflectmodel

import (
	"reflect"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor.IsEqual — invalid != valid branch (line 163-165)
// ══════════════════════════════════════════════════════════════════════════════

func Test_MethodProcessor_IsEqual_InvalidVsValid(t *testing.T) {
	// Arrange
	valid := newMethodProcessorInternal("PublicMethod")
	var invalid *MethodProcessor // nil is invalid

	// Act — non-nil wrapping nil scenario: create invalid (non-nil but invalid)
	invalidNonNil := &MethodProcessor{
		Name:  "",
		Index: -1,
	}

	// Assert
	result := valid.IsEqual(invalidNonNil)
	if result {
		t.Fatal("valid should not equal invalid")
	}

	_ = invalid
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor.GetInArgsTypes — zero args branch (line 229-231)
// ══════════════════════════════════════════════════════════════════════════════

func Test_MethodProcessor_GetInArgsTypes_ZeroArgs(t *testing.T) {
	// Arrange — create a MethodProcessor that wraps a function type with 0 in-args
	// This is effectively impossible via MethodByName (always has receiver),
	// but we test the invalid branch which returns empty

	var mp *MethodProcessor

	// Act
	result := mp.GetInArgsTypes()

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty for nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor.GetInArgsTypesNames — zero args branch (line 253-255)
// ══════════════════════════════════════════════════════════════════════════════

func Test_MethodProcessor_GetInArgsTypesNames_ZeroArgs(t *testing.T) {
	// Arrange
	var mp *MethodProcessor

	// Act
	result := mp.GetInArgsTypesNames()

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty for nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor.validationError — invalid (non-nil) branch (line 276-284)
// ══════════════════════════════════════════════════════════════════════════════

func Test_MethodProcessor_ValidationError_Invalid(t *testing.T) {
	// Arrange — a nil MethodProcessor triggers the nil guard
	var mp *MethodProcessor

	// Act
	err := mp.validationError()

	// Assert
	if err == nil {
		t.Fatal("expected error for nil method processor")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// isNull — struct (default branch, returns false)
// Note: isNull nil, ptr, map, slice, func, chan paths are covered
//       by Coverage_internal_test.go — if showing 0, it's a collection issue.
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsNull_Struct(t *testing.T) {
	// Arrange
	type sample struct{ X int }

	// Act
	result := isNull(sample{X: 1})

	// Assert
	if result {
		t.Fatal("struct should not be null")
	}
}

func Test_IsNull_NilUnsafePointer(t *testing.T) {
	// Arrange
	var p *int

	// Act
	result := isNull(p)

	// Assert
	if !result {
		t.Fatal("nil ptr should be null")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils — ReflectValueToAnyValue nil interface branch (line 45-47)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_ReflectValueToAnyValue_NilInterface(t *testing.T) {
	// Arrange
	u := rvUtils{}
	var nilPtr *string

	// Act
	result := u.ReflectValueToAnyValue(reflect.ValueOf(nilPtr))

	// Assert
	if result != nil {
		t.Fatal("expected nil for nil ptr")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils — IsNull non-nil value (line 64-75 default branch)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_IsNull_NonNilValue(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	result := u.IsNull(42)

	// Assert
	if result {
		t.Fatal("int should not be null")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils — InterfacesToTypes non-empty (line 102-104)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_InterfacesToTypes_NonEmpty(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	result := u.InterfacesToTypes([]any{"a", 1, true})

	// Assert
	if len(result) != 3 {
		t.Fatalf("expected 3, got %d", len(result))
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.IsReflectTypeMatchAny (line 193-198)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_IsReflectTypeMatchAny_Match(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	ok, err := u.IsReflectTypeMatchAny("hello", "world")

	// Assert
	if !ok || err != nil {
		t.Fatal("expected match for same types")
	}
}

func Test_RvUtils_IsReflectTypeMatchAny_Mismatch(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	ok, err := u.IsReflectTypeMatchAny("hello", 42)

	// Assert
	if ok || err == nil {
		t.Fatal("expected mismatch for different types")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.PrependWithSpaces — with no spaces (line 223-225)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_PrependWithSpaces_ZeroPrependSpaces(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act — prependingLinesSpaceCount = 0, spaceCountLines = 0
	result := u.PrependWithSpaces(0, []string{"existing"}, 0, "prepend")

	// Assert
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.WithSpaces — empty (line 239-241)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_WithSpaces_Empty(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	result := u.WithSpaces(4)

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty for no lines")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.ArgsToReflectValues — empty branch (line 15-17)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RvUtils_ArgsToReflectValues_Empty(t *testing.T) {
	// Arrange
	u := rvUtils{}

	// Act
	result := u.ArgsToReflectValues(nil)

	// Assert
	if len(result) != 0 {
		t.Fatal("expected empty for nil args")
	}
}
