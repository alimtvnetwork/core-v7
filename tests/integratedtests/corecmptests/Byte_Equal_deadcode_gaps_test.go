package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// DEAD CODE DOCUMENTATION
//
// The following functions all have a final `return corecomparator.NotEqual`
// that is UNREACHABLE because the preceding if-else chain exhausts all
// possibilities (==, <, >):
//
//   - Byte (line 14)
//   - Integer (line 14)
//   - Integer8 (line 14)
//   - Integer16 (line 14)
//   - Integer32 (line 14)
//   - Integer64 (line 14)
//   - Time (line 18)
//   - VersionSliceByte (line 41)
//   - VersionSliceInteger (line 41)
//
// These are defensive fallbacks that can never execute. All reachable
// branches are tested below.
// ══════════════════════════════════════════════════════════════════════════════

// ---------- Byte ----------

func Test_Byte_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	var a, b byte = 5, 5

	// Act
	result := corecmp.Byte(a, b)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Byte_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	var a, b byte = 3, 7

	// Act
	result := corecmp.Byte(a, b)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Byte_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	var a, b byte = 9, 2

	// Act
	result := corecmp.Byte(a, b)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Integer ----------

func Test_Integer_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer(10, 10)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Integer_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer(1, 10)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Integer_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer(10, 1)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Integer8 ----------

func Test_Integer8_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer8(5, 5)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Integer8_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer8(1, 5)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Integer8_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer8(5, 1)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Integer16 ----------

func Test_Integer16_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer16(100, 100)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Integer16_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer16(10, 100)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Integer16_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer16(100, 10)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Integer32 ----------

func Test_Integer32_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer32(1000, 1000)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Integer32_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer32(100, 1000)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Integer32_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer32(1000, 100)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Integer64 ----------

func Test_Integer64_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer64(10000, 10000)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Integer64_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer64(1000, 10000)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Integer64_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.Integer64(10000, 1000)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- Time ----------

func Test_Time_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	result := corecmp.Time(now, now)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Time_LeftLess_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	earlier := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	later := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	result := corecmp.Time(earlier, later)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_Time_LeftGreater_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange
	earlier := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	later := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	result := corecmp.Time(later, earlier)

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- VersionSliceByte ----------

func Test_VersionSliceByte_BothNil_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte(nil, nil)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_VersionSliceByte_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_VersionSliceByte_LeftLess_ByElement(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte([]byte{1, 1}, []byte{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_VersionSliceByte_LeftGreater_ByElement(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

func Test_VersionSliceByte_LeftLess_ByLength(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2, 3})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_VersionSliceByte_LeftGreater_ByLength(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

// ---------- VersionSliceInteger ----------

func Test_VersionSliceInteger_BothNil_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger(nil, nil)

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_VersionSliceInteger_Equal_FromByteEqualdeadcodegap(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})

	// Assert
	actual := args.Map{"result": result != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_VersionSliceInteger_LeftLess_ByElement(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger([]int{1, 1}, []int{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_VersionSliceInteger_LeftGreater_ByElement(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}

func Test_VersionSliceInteger_LeftLess_ByLength(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2, 3})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
}

func Test_VersionSliceInteger_LeftGreater_ByLength(t *testing.T) {
	// Arrange & Act
	result := corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2})

	// Assert
	actual := args.Map{"result": result != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
}
