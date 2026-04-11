package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── VersionSliceByte ──

func Test_VersionSliceByte_BothNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns Equal -- both nil", actual)
}

func Test_VersionSliceByte_LeftNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte(nil, []byte{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns NotEqual -- left nil", actual)
}

func Test_VersionSliceByte_RightNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1}, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns NotEqual -- right nil", actual)
}

func Test_VersionSliceByte_Equal_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns Equal -- same values", actual)
}

func Test_VersionSliceByte_LeftLess_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns LeftLess -- left smaller at index", actual)
}

func Test_VersionSliceByte_LeftGreater_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 4}, []byte{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns LeftGreater -- left bigger at index", actual)
}

func Test_VersionSliceByte_LeftShorter(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns LeftLess -- left shorter", actual)
}

func Test_VersionSliceByte_LeftLonger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceByte([]byte{1, 2, 3}, []byte{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceByte returns LeftGreater -- left longer", actual)
}

// ── VersionSliceInteger ──

func Test_VersionSliceInteger_BothNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns Equal -- both nil", actual)
}

func Test_VersionSliceInteger_LeftNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger(nil, []int{1})}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns NotEqual -- left nil", actual)
}

func Test_VersionSliceInteger_Equal_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns Equal -- same values", actual)
}

func Test_VersionSliceInteger_LeftLess_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns LeftLess -- left smaller", actual)
}

func Test_VersionSliceInteger_LeftGreater_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 4}, []int{1, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns LeftGreater -- left bigger", actual)
}

func Test_VersionSliceInteger_LeftShorter(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2, 3})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns LeftLess -- left shorter", actual)
}

func Test_VersionSliceInteger_LeftLonger(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.VersionSliceInteger([]int{1, 2, 3}, []int{1, 2})}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "VersionSliceInteger returns LeftGreater -- left longer", actual)
}

// ── IsStringsEqualWithoutOrder ──

func Test_IsStringsEqualWithoutOrder_BothNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns true -- both nil", actual)
}

func Test_IsStringsEqualWithoutOrder_LeftNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder(nil, []string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns false -- left nil", actual)
}

func Test_IsStringsEqualWithoutOrder_DiffLen_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns false -- diff length", actual)
}

func Test_IsStringsEqualWithoutOrder_SameUnordered(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"b", "a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns true -- same unordered", actual)
}

func Test_IsStringsEqualWithoutOrder_Different(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualWithoutOrder([]string{"a", "b"}, []string{"a", "c"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualWithoutOrder returns false -- different", actual)
}

// ── IsStringsEqual extended ──

func Test_IsStringsEqual_DiffLen_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqual([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqual returns false -- diff length", actual)
}

// ── IsStringsEqualPtr extended ──

func Test_IsStringsEqualPtr_DiffLen_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.IsStringsEqualPtr([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsStringsEqualPtr returns false -- diff length", actual)
}

// ── Time extended ──

func Test_Time_Equal_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	now := time.Now()

	// Act
	actual := args.Map{"result": corecmp.Time(now, now)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "Time returns Equal -- same time", actual)
}

func Test_Time_Before_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	t1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftLess}
	expected.ShouldBeEqual(t, 0, "Time returns LeftLess -- before", actual)
}

func Test_Time_After_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.Time(t1, t2)}

	// Assert
	expected := args.Map{"result": corecomparator.LeftGreater}
	expected.ShouldBeEqual(t, 0, "Time returns LeftGreater -- after", actual)
}

// ── TimePtr ──

func Test_TimePtr_BothNil_FromVersionSliceByteBoth(t *testing.T) {
	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns Equal -- both nil", actual)
}

func Test_TimePtr_LeftNil_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	t2 := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(nil, &t2)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns NotEqual -- left nil", actual)
}

func Test_TimePtr_RightNil_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	t1 := time.Now()

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&t1, nil)}

	// Assert
	expected := args.Map{"result": corecomparator.NotEqual}
	expected.ShouldBeEqual(t, 0, "TimePtr returns NotEqual -- right nil", actual)
}

func Test_TimePtr_Equal_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	t1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	actual := args.Map{"result": corecmp.TimePtr(&t1, &t2)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "TimePtr returns Equal -- same time", actual)
}

// ── BytePtr equal ──

func Test_BytePtr_Equal_FromVersionSliceByteBoth(t *testing.T) {
	// Arrange
	l, r := byte(5), byte(5)

	// Act
	actual := args.Map{"result": corecmp.BytePtr(&l, &r)}

	// Assert
	expected := args.Map{"result": corecomparator.Equal}
	expected.ShouldBeEqual(t, 0, "BytePtr returns Equal -- same values", actual)
}
