package reflectmodel

import (
	"reflect"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage4 — reflectcore/reflectmodel remaining gaps (Iteration 16)
//
// Targets:
//   - MethodProcessor.IsEqual: IsPublicMethod differs (line 169-171)
//   - MethodProcessor.IsEqual: ArgsCount differs (line 172-174)
//   - MethodProcessor.IsEqual: ReturnLength differs (line 175-177)
//   - rvUtils.PrependWithSpaces: prependingLinesSpaceCount > 0 (line 223-225)
//   - isNull: nil input (line 5-6), non-nil nilable kinds (line 10-16)
// ══════════════════════════════════════════════════════════════════════════════

// ---------- MethodProcessor.IsEqual — IsPublicMethod branch ----------

func Test_IsEqual_DifferentIsPublicMethod(t *testing.T) {
	// Arrange — both have same Name and same IsInvalid (false)
	// but different PkgPath (one public, one private)
	a := newMethodProcessorInternal("PublicMethod")
	b := newMethodProcessorInternal("PublicMethod")
	b.ReflectMethod.PkgPath = "some/pkg" // makes it private

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when IsPublicMethod differs")
	}
}

// ---------- MethodProcessor.IsEqual — ArgsCount branch ----------

func Test_IsEqual_DifferentArgsCount(t *testing.T) {
	// Arrange — same name, same IsPublicMethod, but different ArgsCount
	a := newMethodProcessorInternal("PublicMethod") // NumIn=2 (receiver+string)
	b := newMethodProcessorInternal("NoArgs")       // NumIn=1 (receiver)
	b.Name = a.Name                                 // force same name

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when ArgsCount differs")
	}
}

// ---------- MethodProcessor.IsEqual — ReturnLength branch ----------

type returnHost struct{}

func (r returnHost) TwoReturns(s string) (string, error) { return s, nil }
func (r returnHost) OneReturn(s string) string            { return s }

func Test_IsEqual_DifferentReturnLength(t *testing.T) {
	// Arrange — same name, same IsPublicMethod, same ArgsCount, different ReturnLength
	rt := reflect.TypeOf(returnHost{})

	m1, _ := rt.MethodByName("TwoReturns") // returns 2
	m2, _ := rt.MethodByName("OneReturn")  // returns 1

	a := &MethodProcessor{Name: "TestFunc", Index: 0, ReflectMethod: m1}
	b := &MethodProcessor{Name: "TestFunc", Index: 0, ReflectMethod: m2}

	// Act
	result := a.IsEqual(b)

	// Assert
	if result {
		t.Fatal("expected false when ReturnLength differs")
	}
}

// ---------- rvUtils.PrependWithSpaces — prependingLinesSpaceCount > 0 ----------

func Test_PrependWithSpaces_WithPrependingSpaces(t *testing.T) {
	// Arrange
	u := rvUtils{}
	existing := []string{"line1", "line2"}

	// Act — prependingLinesSpaceCount = 4 triggers WithSpaces on prepending lines
	result := u.PrependWithSpaces(2, existing, 4, "header")

	// Assert
	if len(result) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(result))
	}
	if result[0] != "    header" {
		t.Fatalf("expected 4-space prepend, got %q", result[0])
	}
}

// ---------- isNull — nil input and nilable kinds ----------

func Test_IsNull_NilInput(t *testing.T) {
	// Arrange & Act
	result := isNull(nil)

	// Assert
	if !result {
		t.Fatal("expected true for nil")
	}
}

func Test_IsNull_NilMap(t *testing.T) {
	// Arrange
	var m map[string]int

	// Act
	result := isNull(m)

	// Assert
	if !result {
		t.Fatal("expected true for nil map")
	}
}

func Test_IsNull_NilSlice(t *testing.T) {
	// Arrange
	var s []int

	// Act
	result := isNull(s)

	// Assert
	if !result {
		t.Fatal("expected true for nil slice")
	}
}

func Test_IsNull_NilChan(t *testing.T) {
	// Arrange
	var ch chan int

	// Act
	result := isNull(ch)

	// Assert
	if !result {
		t.Fatal("expected true for nil chan")
	}
}

func Test_IsNull_NilFunc(t *testing.T) {
	// Arrange
	var fn func()

	// Act
	result := isNull(fn)

	// Assert
	if !result {
		t.Fatal("expected true for nil func")
	}
}

func Test_IsNull_NonNilInt(t *testing.T) {
	// Arrange & Act
	result := isNull(42)

	// Assert
	if result {
		t.Fatal("expected false for int")
	}
}
