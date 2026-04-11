package coretestsargstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── helper functions for *Func types ──

func sampleAdd(a, b int) int    { return a + b }
func sampleGreet(name string) string { return "hi " + name }
func sampleConcat(a, b, c string) string { return a + b + c }
func sampleSum4(a, b, c, d int) int { return a + b + c + d }
func sampleSum5(a, b, c, d, e int) int { return a + b + c + d + e }
func sampleSum6(a, b, c, d, e, f int) int { return a + b + c + d + e + f }
func sampleRetErr() error { return errors.New("test-err") }
func sampleRetBool() bool { return true }
func sampleRetString() string { return "hello" }
func sampleRetAny() any { return 42 }
func sampleRetAnyErr() (any, error) { return "ok", nil }
func sampleRetAnyErrFail() (any, error) { return nil, errors.New("fail") }

// ── OneFunc: Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs, Slice, String ──

func Test_OneFunc_Invoke(t *testing.T) {
	// Arrange
	of := args.OneFunc[string]{
		First:    "world",
		WorkFunc: sampleGreet,
	}

	// Act
	results, err := of.Invoke("world")

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"result": fmt.Sprintf("%v", results[0]),
	}
	expected := args.Map{
		"noErr":  true,
		"result": "hi world",
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct -- OneFunc with valid func", actual)
}

func Test_OneFunc_InvokeMust(t *testing.T) {
	// Arrange
	of := args.OneFunc[string]{
		First:    "test",
		WorkFunc: sampleGreet,
	}

	// Act
	results := of.InvokeMust("test")

	// Assert
	actual := args.Map{"result": fmt.Sprintf("%v", results[0])}
	expected := args.Map{"result": "hi test"}
	expected.ShouldBeEqual(t, 0, "InvokeMust returns correct -- OneFunc", actual)
}

// ── TwoFunc: Invoke, InvokeMust, InvokeWithValidArgs, InvokeArgs, Slice, GetByIndex, String ──

func Test_TwoFunc_Invoke(t *testing.T) {
	// Arrange
	tf := args.TwoFunc[int, int]{
		First:    3,
		Second:   4,
		WorkFunc: sampleAdd,
	}

	// Act
	results, err := tf.Invoke(3, 4)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 7,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct -- TwoFunc with add", actual)
}

func Test_TwoFunc_InvokeMust(t *testing.T) {
	// Arrange
	tf := args.TwoFunc[int, int]{
		First:    1,
		Second:   2,
		WorkFunc: sampleAdd,
	}

	// Act
	results := tf.InvokeMust(1, 2)

	// Assert
	actual := args.Map{"result": results[0]}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "InvokeMust returns correct -- TwoFunc", actual)
}

func Test_TwoFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	tf := args.TwoFunc[int, int]{
		First:    10,
		Second:   20,
		WorkFunc: sampleAdd,
	}

	// Act
	results, err := tf.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 30,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- TwoFunc", actual)
}

func Test_TwoFunc_InvokeArgs(t *testing.T) {
	// Arrange
	tf := args.TwoFunc[int, int]{
		First:    5,
		Second:   6,
		WorkFunc: sampleAdd,
	}

	// Act
	results, err := tf.InvokeArgs(2)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 11,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- TwoFunc upTo 2", actual)
}

func Test_TwoFunc_Slice_GetByIndex_String(t *testing.T) {
	// Arrange
	tf := args.TwoFunc[int, int]{
		First:  1,
		Second: 2,
		Expect: 99,
	}

	// Act
	slice := tf.Slice()
	byIdx := tf.GetByIndex(0)
	str := tf.String()

	// Assert
	actual := args.Map{
		"sliceLen":  len(slice),
		"firstItem": byIdx,
		"hasStr":    len(str) > 0,
	}
	expected := args.Map{
		"sliceLen":  3,
		"firstItem": 1,
		"hasStr":    true,
	}
	expected.ShouldBeEqual(t, 0, "Slice/GetByIndex/String returns correct -- TwoFunc", actual)
}

// ── ThreeFunc ──

func Test_ThreeFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	tf := args.ThreeFunc[string, string, string]{
		First:    "a",
		Second:   "b",
		Third:    "c",
		WorkFunc: sampleConcat,
	}

	// Act
	results, err := tf.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": "abc",
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- ThreeFunc", actual)
}

func Test_ThreeFunc_InvokeArgs(t *testing.T) {
	// Arrange
	tf := args.ThreeFunc[string, string, string]{
		First:    "x",
		Second:   "y",
		Third:    "z",
		WorkFunc: sampleConcat,
	}

	// Act
	results, err := tf.InvokeArgs(3)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": "xyz",
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- ThreeFunc upTo 3", actual)
}

func Test_ThreeFunc_Slice_String(t *testing.T) {
	// Arrange
	tf := args.ThreeFunc[string, string, string]{
		First:  "a",
		Second: "b",
		Third:  "c",
	}

	// Act
	slice := tf.Slice()
	str := tf.String()

	// Assert
	actual := args.Map{
		"sliceLen": len(slice),
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"sliceLen": 3,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Slice/String returns correct -- ThreeFunc", actual)
}

// ── FourFunc ──

func Test_FourFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	ff := args.FourFunc[int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		WorkFunc: sampleSum4,
	}

	// Act
	results, err := ff.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 10,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- FourFunc", actual)
}

func Test_FourFunc_InvokeArgs(t *testing.T) {
	// Arrange
	ff := args.FourFunc[int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		WorkFunc: sampleSum4,
	}

	// Act
	results, err := ff.InvokeArgs(4)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 10,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- FourFunc upTo 4", actual)
}

func Test_FourFunc_Slice_String(t *testing.T) {
	// Arrange
	ff := args.FourFunc[int, int, int, int]{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
	}

	// Act
	slice := ff.Slice()
	str := ff.String()

	// Assert
	actual := args.Map{
		"sliceLen": len(slice),
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"sliceLen": 4,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Slice/String returns correct -- FourFunc", actual)
}

// ── FiveFunc ──

func Test_FiveFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	ff := args.FiveFunc[int, int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		Fifth:    5,
		WorkFunc: sampleSum5,
	}

	// Act
	results, err := ff.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 15,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- FiveFunc", actual)
}

func Test_FiveFunc_InvokeArgs(t *testing.T) {
	// Arrange
	ff := args.FiveFunc[int, int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		Fifth:    5,
		WorkFunc: sampleSum5,
	}

	// Act
	results, err := ff.InvokeArgs(5)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 15,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- FiveFunc upTo 5", actual)
}

func Test_FiveFunc_Slice_String(t *testing.T) {
	// Arrange
	ff := args.FiveFunc[int, int, int, int, int]{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
		Fifth:  5,
	}

	// Act
	slice := ff.Slice()
	str := ff.String()

	// Assert
	actual := args.Map{
		"sliceLen": len(slice),
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"sliceLen": 5,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Slice/String returns correct -- FiveFunc", actual)
}

// ── SixFunc ──

func Test_SixFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	sf := args.SixFunc[int, int, int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		Fifth:    5,
		Sixth:    6,
		WorkFunc: sampleSum6,
	}

	// Act
	results, err := sf.InvokeWithValidArgs()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 21,
	}
	expected.ShouldBeEqual(t, 0, "InvokeWithValidArgs returns correct -- SixFunc", actual)
}

func Test_SixFunc_InvokeArgs(t *testing.T) {
	// Arrange
	sf := args.SixFunc[int, int, int, int, int, int]{
		First:    1,
		Second:   2,
		Third:    3,
		Fourth:   4,
		Fifth:    5,
		Sixth:    6,
		WorkFunc: sampleSum6,
	}

	// Act
	results, err := sf.InvokeArgs(6)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}
	expected := args.Map{
		"noErr": true,
		"result": 21,
	}
	expected.ShouldBeEqual(t, 0, "InvokeArgs returns correct -- SixFunc upTo 6", actual)
}

func Test_SixFunc_Slice_String(t *testing.T) {
	// Arrange
	sf := args.SixFunc[int, int, int, int, int, int]{
		First:  1,
		Second: 2,
		Third:  3,
		Fourth: 4,
		Fifth:  5,
		Sixth:  6,
	}

	// Act
	slice := sf.Slice()
	str := sf.String()

	// Assert
	actual := args.Map{
		"sliceLen": len(slice),
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"sliceLen": 6,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Slice/String returns correct -- SixFunc", actual)
}
