package args

import (
	"errors"
	"fmt"
	"testing"
)

// Helper functions for tests
func sampleAdd(a, b int) int          { return a + b }
func sampleGreet(name string) string  { return "Hello " + name }
func sampleNoArgs() string            { return "no-args" }
func sampleError(s string) error      { return errors.New(s) }
func sampleBool(v int) bool           { return v > 0 }
func sampleMultiReturn(a int) (string, error) {
	if a < 0 {
		return "", errors.New("negative")
	}
	return fmt.Sprintf("%d", a), nil
}
func sampleThreeArgs(a, b, c int) int    { return a + b + c }
func sampleFourArgs(a, b, c, d int) int  { return a + b + c + d }
func sampleFiveArgs(a, b, c, d, e int) int { return a + b + c + d + e }
func sampleSixArgs(a, b, c, d, e, f int) int { return a + b + c + d + e + f }

// ══════════════════════════════════════════════════════════════════════════════
// OneFunc — Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs
// Covers OneFunc.go L84-86, L89-91, L103-105
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_OneFunc_Invoke(t *testing.T) {
	of := &OneFunc[string]{
		First:    "World",
		WorkFunc: sampleGreet,
	}

	results, err := of.Invoke("World")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if results[0] != "Hello World" {
		t.Errorf("got %v, want Hello World", results[0])
	}
}

func Test_Cov9_OneFunc_InvokeMust(t *testing.T) {
	of := &OneFunc[string]{
		First:    "Test",
		WorkFunc: sampleGreet,
	}
	results := of.InvokeMust("Test")
	if results[0] != "Hello Test" {
		t.Errorf("got %v, want Hello Test", results[0])
	}
}

func Test_Cov9_OneFunc_InvokeArgs(t *testing.T) {
	of := &OneFunc[string]{
		First:    "Args",
		WorkFunc: sampleGreet,
	}
	results, err := of.InvokeArgs(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if results[0] != "Hello Args" {
		t.Errorf("got %v, want Hello Args", results[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// TwoFunc — Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs, String
// Covers TwoFunc.go L98-100, L103-105, L117-119, L170-172
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_TwoFunc_Invoke(t *testing.T) {
	tf := &TwoFunc[int, int]{
		First:    1,
		Second:   2,
		WorkFunc: sampleAdd,
	}
	results, err := tf.Invoke(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if results[0] != 3 {
		t.Errorf("got %v, want 3", results[0])
	}
}

func Test_Cov9_TwoFunc_InvokeMust(t *testing.T) {
	tf := &TwoFunc[int, int]{First: 1, Second: 2, WorkFunc: sampleAdd}
	results := tf.InvokeMust(1, 2)
	if results[0] != 3 {
		t.Errorf("got %v, want 3", results[0])
	}
}

func Test_Cov9_TwoFunc_InvokeArgs(t *testing.T) {
	tf := &TwoFunc[int, int]{First: 1, Second: 2, WorkFunc: sampleAdd}
	results, err := tf.InvokeArgs(2)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 3 {
		t.Errorf("got %v, want 3", results[0])
	}
}

func Test_Cov9_TwoFunc_String(t *testing.T) {
	tf := &TwoFunc[int, int]{First: 10, Second: 20, WorkFunc: sampleAdd}
	s := tf.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ThreeFunc — same pattern
// Covers ThreeFunc.go L119-121, L124-126, L138-140, L174-176, L197-199
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_ThreeFunc_Invoke(t *testing.T) {
	tf := &ThreeFunc[int, int, int]{
		First: 1, Second: 2, Third: 3,
		WorkFunc: sampleThreeArgs,
	}
	results, err := tf.Invoke(1, 2, 3)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 6 {
		t.Errorf("got %v, want 6", results[0])
	}
}

func Test_Cov9_ThreeFunc_InvokeMust(t *testing.T) {
	tf := &ThreeFunc[int, int, int]{First: 1, Second: 2, Third: 3, WorkFunc: sampleThreeArgs}
	results := tf.InvokeMust(1, 2, 3)
	if results[0] != 6 {
		t.Errorf("got %v, want 6", results[0])
	}
}

func Test_Cov9_ThreeFunc_InvokeArgs(t *testing.T) {
	tf := &ThreeFunc[int, int, int]{First: 1, Second: 2, Third: 3, WorkFunc: sampleThreeArgs}
	results, err := tf.InvokeArgs(3)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 6 {
		t.Errorf("got %v, want 6", results[0])
	}
}

func Test_Cov9_ThreeFunc_Slice(t *testing.T) {
	tf := &ThreeFunc[int, int, int]{First: 1, Second: 2, Third: 3, WorkFunc: sampleThreeArgs}
	s := tf.Slice()
	if len(s) == 0 {
		t.Error("expected non-empty slice")
	}
}

func Test_Cov9_ThreeFunc_GetByIndex(t *testing.T) {
	tf := &ThreeFunc[int, int, int]{First: 10, Second: 20, Third: 30}
	slice := tf.Slice()
	item := tf.GetByIndex(0)
	if item != slice[0] {
		t.Errorf("got %v, want %v", item, slice[0])
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FourFunc — same pattern
// Covers FourFunc.go L123-125, L130-132, L135-137, L142-144, L149-151, L190-192, L214-216
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FourFunc_Invoke(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{
		First: 1, Second: 2, Third: 3, Fourth: 4,
		WorkFunc: sampleFourArgs,
	}
	results, err := ff.Invoke(1, 2, 3, 4)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 10 {
		t.Errorf("got %v, want 10", results[0])
	}
}

func Test_Cov9_FourFunc_InvokeMust(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, WorkFunc: sampleFourArgs}
	results := ff.InvokeMust(1, 2, 3, 4)
	if results[0] != 10 {
		t.Errorf("got %v, want 10", results[0])
	}
}

func Test_Cov9_FourFunc_InvokeWithValidArgs(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, WorkFunc: sampleFourArgs}
	results, err := ff.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 10 {
		t.Errorf("got %v, want 10", results[0])
	}
}

func Test_Cov9_FourFunc_InvokeArgs(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, WorkFunc: sampleFourArgs}
	results, err := ff.InvokeArgs(4)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 10 {
		t.Errorf("got %v, want 10", results[0])
	}
}

func Test_Cov9_FourFunc_Slice(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, WorkFunc: sampleFourArgs}
	s := ff.Slice()
	if len(s) == 0 {
		t.Error("expected non-empty slice")
	}
}

func Test_Cov9_FourFunc_String(t *testing.T) {
	ff := &FourFunc[int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4}
	s := ff.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// FiveFunc — same pattern
// Covers FiveFunc.go L134-136, L141-143, L146-148, L153-155, L160-162, L206-208, L231-233
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_FiveFunc_Invoke(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{
		First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5,
		WorkFunc: sampleFiveArgs,
	}
	results, err := ff.Invoke(1, 2, 3, 4, 5)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 15 {
		t.Errorf("got %v, want 15", results[0])
	}
}

func Test_Cov9_FiveFunc_InvokeMust(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, WorkFunc: sampleFiveArgs}
	results := ff.InvokeMust(1, 2, 3, 4, 5)
	if results[0] != 15 {
		t.Errorf("got %v, want 15", results[0])
	}
}

func Test_Cov9_FiveFunc_InvokeWithValidArgs(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, WorkFunc: sampleFiveArgs}
	results, err := ff.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 15 {
		t.Errorf("got %v, want 15", results[0])
	}
}

func Test_Cov9_FiveFunc_InvokeArgs(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, WorkFunc: sampleFiveArgs}
	results, err := ff.InvokeArgs(5)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 15 {
		t.Errorf("got %v, want 15", results[0])
	}
}

func Test_Cov9_FiveFunc_Slice(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, WorkFunc: sampleFiveArgs}
	s := ff.Slice()
	if len(s) == 0 {
		t.Error("expected non-empty slice")
	}
}

func Test_Cov9_FiveFunc_String(t *testing.T) {
	ff := &FiveFunc[int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5}
	s := ff.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SixFunc — same pattern
// Covers SixFunc.go L156-158, L163-165, L168-170, L175-177, L182-184, L233-235, L259-261
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_SixFunc_Invoke(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{
		First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6,
		WorkFunc: sampleSixArgs,
	}
	results, err := sf.Invoke(1, 2, 3, 4, 5, 6)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 21 {
		t.Errorf("got %v, want 21", results[0])
	}
}

func Test_Cov9_SixFunc_InvokeMust(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6, WorkFunc: sampleSixArgs}
	results := sf.InvokeMust(1, 2, 3, 4, 5, 6)
	if results[0] != 21 {
		t.Errorf("got %v, want 21", results[0])
	}
}

func Test_Cov9_SixFunc_InvokeWithValidArgs(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6, WorkFunc: sampleSixArgs}
	results, err := sf.InvokeWithValidArgs()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 21 {
		t.Errorf("got %v, want 21", results[0])
	}
}

func Test_Cov9_SixFunc_InvokeArgs(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6, WorkFunc: sampleSixArgs}
	results, err := sf.InvokeArgs(6)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if results[0] != 21 {
		t.Errorf("got %v, want 21", results[0])
	}
}

func Test_Cov9_SixFunc_Slice(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6, WorkFunc: sampleSixArgs}
	s := sf.Slice()
	if len(s) == 0 {
		t.Error("expected non-empty slice")
	}
}

func Test_Cov9_SixFunc_String(t *testing.T) {
	sf := &SixFunc[int, int, int, int, int, int]{First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5, Sixth: 6}
	s := sf.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}
