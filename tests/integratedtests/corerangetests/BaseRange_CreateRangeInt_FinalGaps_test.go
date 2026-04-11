package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage8 — corerange final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- BaseRange methods ---

func Test_BaseRange_CreateRangeInt(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{
		RawInput:  "1-5",
		Separator: "-",
		IsValid:   true,
	}
	minMax := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	result := br.CreateRangeInt(minMax)

	// Assert
	convey.Convey("BaseRange.CreateRangeInt creates RangeInt", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_BaseRange_IsInvalid(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{IsValid: false}

	// Act & Assert
	convey.Convey("BaseRange.IsInvalid returns true", t, func() {
		convey.So(br.IsInvalid(), convey.ShouldBeTrue)
	})
}

func Test_BaseRange_String(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{Separator: "-"}

	// Act
	result := br.String(1, 5)

	// Assert
	convey.Convey("BaseRange.String formats correctly", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "-")
	})
}

func Test_BaseRange_Clone(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{
		RawInput:  "1-5",
		Separator: "-",
		IsValid:   true,
		HasStart:  true,
		HasEnd:    true,
	}

	// Act
	cloned := br.BaseRangeClone()

	// Assert
	convey.Convey("BaseRange.Clone returns copy", t, func() {
		convey.So(cloned.RawInput, convey.ShouldEqual, "1-5")
	})
}

// --- MinMaxByte methods ---

func Test_MinMaxByte_CreateMinMaxInt(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	result := mmb.CreateMinMaxInt()

	// Assert
	convey.Convey("MinMaxByte.CreateMinMaxInt returns MinMaxInt", t, func() {
		convey.So(result.Min, convey.ShouldEqual, 1)
	})
}

func Test_MinMaxByte_CreateRangeInt(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	result := mmb.CreateRangeInt("1-5", "-")

	// Assert
	convey.Convey("MinMaxByte.CreateRangeInt returns RangeInt", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_MinMaxByte_CreateRangeInt8(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	result := mmb.CreateRangeInt8("1-5", "-")

	// Assert
	convey.Convey("MinMaxByte.CreateRangeInt8 returns RangeInt8", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_MinMaxByte_CreateRangeInt16(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	result := mmb.CreateRangeInt16("1-5", "-")

	// Assert
	convey.Convey("MinMaxByte.CreateRangeInt16 returns RangeInt16", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_MinMaxByte_Comparison(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 5, Max: 10}

	// Act & Assert
	convey.Convey("MinMaxByte comparison methods", t, func() {
		convey.So(mmb.IsMinEqual(5), convey.ShouldBeTrue)
		convey.So(mmb.IsMinAboveEqual(5), convey.ShouldBeTrue)
		convey.So(mmb.IsMinAbove(4), convey.ShouldBeTrue)
		convey.So(mmb.IsMinLess(6), convey.ShouldBeTrue)
		convey.So(mmb.IsMinLessEqual(5), convey.ShouldBeTrue)
		convey.So(mmb.IsMaxEqual(10), convey.ShouldBeTrue)
		convey.So(mmb.IsMaxAboveEqual(10), convey.ShouldBeTrue)
		convey.So(mmb.IsMaxAbove(9), convey.ShouldBeTrue)
		convey.So(mmb.IsMaxLess(11), convey.ShouldBeTrue)
		convey.So(mmb.IsMaxLessEqual(10), convey.ShouldBeTrue)
	})
}

func Test_MinMaxByte_RangesInt(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 3, Max: 5}

	// Act
	result := mmb.RangesInt()

	// Assert
	convey.Convey("MinMaxByte.RangesInt returns int ranges", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_MinMaxByte_IsWithinRange(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 3, Max: 5}

	// Act & Assert
	convey.Convey("MinMaxByte range checks", t, func() {
		convey.So(mmb.IsWithinRange(4), convey.ShouldBeTrue)
		convey.So(mmb.IsInvalidValue(2), convey.ShouldBeTrue)
		convey.So(mmb.IsOutOfRange(6), convey.ShouldBeTrue)
	})
}

func Test_MinMaxByte_Clone(t *testing.T) {
	// Arrange
	mmb := &corerange.MinMaxByte{Min: 3, Max: 5}

	// Act
	cloned := mmb.Clone()
	clonedPtr := mmb.ClonePtr()

	// Assert
	convey.Convey("MinMaxByte.Clone and ClonePtr", t, func() {
		convey.So(cloned.Min, convey.ShouldEqual, 3)
		convey.So(clonedPtr.Min, convey.ShouldEqual, 3)
	})
}

func Test_MinMaxByte_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var mmb *corerange.MinMaxByte

	// Act
	result := mmb.ClonePtr()

	// Assert
	convey.Convey("MinMaxByte.ClonePtr nil returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

// --- MinMaxInt methods ---

func Test_MinMaxInt_Comparison(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt{Min: 5, Max: 10}

	// Act & Assert
	convey.Convey("MinMaxInt comparison methods", t, func() {
		convey.So(mmi.IsMinEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMinAboveEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMinAbove(4), convey.ShouldBeTrue)
		convey.So(mmi.IsMinLess(6), convey.ShouldBeTrue)
		convey.So(mmi.IsMinLessEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxEqual(10), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxAboveEqual(10), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxAbove(9), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxLess(11), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxLessEqual(10), convey.ShouldBeTrue)
	})
}

func Test_MinMaxInt_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange — Min > Max (unusual but tests the branch)
	mmi := &corerange.MinMaxInt{Min: 10, Max: 5}

	// Act
	result := mmi.DifferenceAbsolute()

	// Assert
	convey.Convey("DifferenceAbsolute handles negative diff", t, func() {
		convey.So(result, convey.ShouldEqual, 5)
	})
}

func Test_MinMaxInt_CreateRanges(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt{Min: 1, Max: 3}
	extra := corerange.MinMaxInt{Min: 5, Max: 6}

	// Act
	result := mmi.CreateRanges(extra)

	// Assert
	convey.Convey("CreateRanges combines ranges", t, func() {
		convey.So(len(result), convey.ShouldEqual, 5) // 1,2,3 + 5,6
	})
}

func Test_MinMaxInt_CreateRanges_NoExtra(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt{Min: 1, Max: 3}

	// Act
	result := mmi.CreateRanges()

	// Assert
	convey.Convey("CreateRanges no extra returns own ranges", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_MinMaxInt_RangesExcept(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt{Min: 1, Max: 5}

	// Act
	result := mmi.RangesExcept(2, 4)

	// Assert
	convey.Convey("RangesExcept excludes items", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_MinMaxInt_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a *corerange.MinMaxInt
	var b *corerange.MinMaxInt

	// Act & Assert
	convey.Convey("IsEqual both nil returns true", t, func() {
		convey.So(a.IsEqual(b), convey.ShouldBeTrue)
	})
}

func Test_MinMaxInt_IsEqual_OneNil(t *testing.T) {
	// Arrange
	a := &corerange.MinMaxInt{Min: 1, Max: 5}

	// Act & Assert
	convey.Convey("IsEqual one nil returns true (current impl)", t, func() {
		convey.So(a.IsEqual(nil), convey.ShouldBeTrue) // matches impl
	})
}

func Test_MinMaxInt_IsEqual_SamePointer(t *testing.T) {
	// Arrange
	a := &corerange.MinMaxInt{Min: 1, Max: 5}

	// Act & Assert
	convey.Convey("IsEqual same pointer returns true", t, func() {
		convey.So(a.IsEqual(a), convey.ShouldBeTrue)
	})
}

func Test_MinMaxInt_String(t *testing.T) {
	// Arrange
	mmi := corerange.MinMaxInt{Min: 1, Max: 5}

	// Act
	result := mmi.String()

	// Assert
	convey.Convey("MinMaxInt.String formats correctly", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// --- MinMaxInt16 methods ---

func Test_MinMaxInt16_Comparison(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt16{Min: 5, Max: 10}

	// Act & Assert
	convey.Convey("MinMaxInt16 comparison methods", t, func() {
		convey.So(mmi.IsMinEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMinAboveEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMinAbove(4), convey.ShouldBeTrue)
		convey.So(mmi.IsMinLess(6), convey.ShouldBeTrue)
		convey.So(mmi.IsMinLessEqual(5), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxEqual(10), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxAboveEqual(10), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxAbove(9), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxLess(11), convey.ShouldBeTrue)
		convey.So(mmi.IsMaxLessEqual(10), convey.ShouldBeTrue)
	})
}

func Test_MinMaxInt16_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt16{Min: 10, Max: 5}

	// Act
	result := mmi.DifferenceAbsolute()

	// Assert
	convey.Convey("MinMaxInt16.DifferenceAbsolute handles negative", t, func() {
		convey.So(result, convey.ShouldEqual, int16(5))
	})
}

func Test_MinMaxInt16_Clone(t *testing.T) {
	// Arrange
	mmi := &corerange.MinMaxInt16{Min: 3, Max: 8}

	// Act
	cloned := mmi.Clone()
	clonedPtr := mmi.ClonePtr()

	// Assert
	convey.Convey("MinMaxInt16.Clone and ClonePtr", t, func() {
		convey.So(cloned.Min, convey.ShouldEqual, int16(3))
		convey.So(clonedPtr.Min, convey.ShouldEqual, int16(3))
	})
}

func Test_MinMaxInt16_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var mmi *corerange.MinMaxInt16

	// Act
	result := mmi.ClonePtr()

	// Assert
	convey.Convey("MinMaxInt16.ClonePtr nil returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

// --- RangeInt methods ---

func Test_RangeInt_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntUsingValues(10, 5, true)

	// Act
	result := ri.DifferenceAbsolute()

	// Assert
	convey.Convey("RangeInt.DifferenceAbsolute handles negative", t, func() {
		convey.So(result, convey.ShouldEqual, 5)
	})
}

func Test_RangeInt_CreateStartEnd(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntUsingValues(1, 5, true)

	// Act
	result := ri.CreateStartEnd()

	// Assert
	convey.Convey("CreateStartEnd returns StartEndInt", t, func() {
		convey.So(result.Start, convey.ShouldEqual, 1)
		convey.So(result.End, convey.ShouldEqual, 5)
	})
}

func Test_RangeInt_ShallowCreate(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntUsingValues(1, 5, true)

	// Act
	ri16 := ri.ShallowCreateRangeInt16()
	ri8 := ri.ShallowCreateRangeInt8()
	riByte := ri.ShallowCreateRangeByte()

	// Assert
	convey.Convey("ShallowCreate methods return correct types", t, func() {
		convey.So(ri16.Start, convey.ShouldEqual, int16(1))
		convey.So(ri8.Start, convey.ShouldEqual, int8(1))
		convey.So(riByte.Start, convey.ShouldEqual, byte(1))
	})
}

func Test_RangeInt_CreateRanges(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("1-3", "-", 0, 10)

	// Act
	result := ri.CreateRanges(corerange.MinMaxInt{Min: 5, Max: 6})

	// Assert
	convey.Convey("RangeInt.CreateRanges combines", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_RangeInt_RangesExcept(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("1-5", "-", 0, 10)

	// Act
	result := ri.RangesExcept(2, 4)

	// Assert
	convey.Convey("RangeInt.RangesExcept excludes items", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// --- RangeByte methods ---

func Test_RangeByte_NewRangeByteMinMax(t *testing.T) {
	// Act
	result := corerange.NewRangeByteMinMax("1-5", "-", 0, 10)

	// Assert
	convey.Convey("NewRangeByteMinMax creates RangeByte", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_RangeByte_NewRangeByte_NilMinMax(t *testing.T) {
	// Act
	result := corerange.NewRangeByte("1-5", "-", nil)

	// Assert
	convey.Convey("NewRangeByte nil minMax uses default", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_RangeByte_Difference(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeByteMinMax("1-5", "-", 0, 10)

	// Act & Assert
	convey.Convey("RangeByte.Difference and DifferenceAbsolute", t, func() {
		convey.So(rb.Difference(), convey.ShouldBeGreaterThanOrEqualTo, 0)
		convey.So(rb.DifferenceAbsolute(), convey.ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func Test_RangeByte_Ranges_Invalid(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeByteMinMax("invalid", "-", 0, 10)

	// Act
	result := rb.Ranges()

	// Assert
	convey.Convey("RangeByte.Ranges invalid returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- RangeAny methods ---

func Test_RangeAny_Methods(t *testing.T) {
	// Arrange
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{
			RawInput:  "1-5",
			Separator: "-",
			IsValid:   true,
		},
		RawInput: "1-5",
		Start:    1,
		End:      5,
	}

	// Act & Assert
	convey.Convey("RangeAny methods", t, func() {
		convey.So(ra.RawInputString(), convey.ShouldEqual, "1-5")
		convey.So(ra.StartString(), convey.ShouldEqual, "1")
		convey.So(ra.EndString(), convey.ShouldEqual, "5")
		convey.So(ra.CreateRangeInt(), convey.ShouldNotBeNil)
		convey.So(ra.CreateRangeString(), convey.ShouldNotBeNil)
		convey.So(ra.CreateStartEndString(), convey.ShouldNotBeNil)
		convey.So(ra.String(), convey.ShouldNotBeEmpty)
	})
}

func Test_RangeAny_CreateRangeIntMinMax(t *testing.T) {
	// Arrange
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{
			RawInput:  "1-5",
			Separator: "-",
			IsValid:   true,
		},
		RawInput: "1-5",
		Start:    1,
		End:      5,
	}
	minMax := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	result := ra.CreateRangeIntMinMax(minMax)

	// Assert
	convey.Convey("RangeAny.CreateRangeIntMinMax returns RangeInt", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

// --- RangeInt16 methods ---

func Test_RangeInt16_NewRangeInt16_NilMinMax(t *testing.T) {
	// Act
	result := corerange.NewRangeInt16("1-5", "-", nil)

	// Assert
	convey.Convey("NewRangeInt16 nil minMax uses default", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
	})
}

func Test_RangeInt16_DifferenceAbsolute_Negative(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeInt16MinMax("5-1", "-", 0, 10)

	// Act
	result := ri.DifferenceAbsolute()

	// Assert
	convey.Convey("RangeInt16.DifferenceAbsolute handles negative", t, func() {
		convey.So(result, convey.ShouldBeGreaterThanOrEqualTo, 0)
	})
}

// --- StartEndInt methods ---

func Test_StartEndInt(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 1, End: 5}

	// Act & Assert
	convey.Convey("StartEndInt methods", t, func() {
		convey.So(se.RangeLength(), convey.ShouldEqual, 5)
		convey.So(se.Ranges(), convey.ShouldNotBeEmpty)
		convey.So(se.String(), convey.ShouldNotBeEmpty)
		convey.So(se.IsInvalid(), convey.ShouldBeFalse)
	})
}

func Test_StartEndInt_Diff(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 1, End: 5}

	// Act
	diff := se.Diff()
	diffAbs := se.DifferenceAbsolute()

	// Assert
	convey.Convey("StartEndInt.Diff and DifferenceAbsolute", t, func() {
		convey.So(diff, convey.ShouldEqual, 4)
		convey.So(diffAbs, convey.ShouldEqual, 4)
	})
}

// StartEndInt.ClonePtr, StartEndString.IsEmpty, StartEndSimpleString.Length/IsEmpty/String/Clone/ClonePtr,
// and Within.StringRangeInt do not exist — removed.

func Test_StartEndString_String(t *testing.T) {
	// Arrange
	ses := &corerange.StartEndString{
		BaseRange: &corerange.BaseRange{
			Separator: "-",
		},
		Start: "hello",
		End:   "world",
	}

	// Act & Assert
	convey.Convey("StartEndString.String", t, func() {
		convey.So(ses.String(), convey.ShouldNotBeEmpty)
	})
}

func Test_StartEndSimpleString_StringHyphen(t *testing.T) {
	// Arrange
	sess := &corerange.StartEndSimpleString{
		Start: "a",
		End:   "z",
	}

	// Act & Assert
	convey.Convey("StartEndSimpleString.StringHyphen", t, func() {
		convey.So(sess.StringHyphen(), convey.ShouldNotBeEmpty)
		convey.So(sess.StringColon(), convey.ShouldNotBeEmpty)
		convey.So(sess.StringSpace(), convey.ShouldNotBeEmpty)
	})
}

func Test_Within_StringRangeInt8(t *testing.T) {
	// Act
	val, ok := corerange.Within.StringRangeInt8("42")

	// Assert
	convey.Convey("Within.StringRangeInt8 parses valid", t, func() {
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(val, convey.ShouldEqual, int8(42))
	})
}

func Test_Within_StringRangeInt16(t *testing.T) {
	// Act
	val, ok := corerange.Within.StringRangeInt16("42")

	// Assert
	convey.Convey("Within.StringRangeInt16 parses valid", t, func() {
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(val, convey.ShouldEqual, int16(42))
	})
}

func Test_Within_StringRangeInt32(t *testing.T) {
	// Act
	val, ok := corerange.Within.StringRangeInt32("42")

	// Assert
	convey.Convey("Within.StringRangeInt32 parses valid", t, func() {
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(val, convey.ShouldEqual, int32(42))
	})
}

func Test_Within_StringRangeByte(t *testing.T) {
	// Act
	val, ok := corerange.Within.StringRangeByte("42")

	// Assert
	convey.Convey("Within.StringRangeByte parses valid", t, func() {
		convey.So(ok, convey.ShouldBeTrue)
		convey.So(val, convey.ShouldEqual, byte(42))
	})
}

func Test_Within_StringRangeByte_OutOfRange(t *testing.T) {
	// Act
	_, ok := corerange.Within.StringRangeByte("300")

	// Assert
	convey.Convey("Within.StringRangeByte out of range", t, func() {
		convey.So(ok, convey.ShouldBeFalse)
	})
}

// --- NewRangeInt edge cases ---

func Test_NewRangeInt_NilMinMax(t *testing.T) {
	// Act
	result := corerange.NewRangeInt("1-5", "-", nil)

	// Assert
	convey.Convey("NewRangeInt nil minMax", t, func() {
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(result.IsValid, convey.ShouldBeTrue)
	})
}

func Test_NewRangeInt_Invalid(t *testing.T) {
	// Act
	result := corerange.NewRangeInt("abc", "-", nil)

	// Assert
	convey.Convey("NewRangeInt invalid input", t, func() {
		convey.So(result.IsInvalid(), convey.ShouldBeTrue)
	})
}

func Test_NewRangeIntUsingValues(t *testing.T) {
	// Act
	result := corerange.NewRangeIntUsingValues(1, 5, true)

	// Assert
	convey.Convey("NewRangeIntUsingValues creates valid", t, func() {
		convey.So(result.IsValid, convey.ShouldBeTrue)
		convey.So(result.Start, convey.ShouldEqual, 1)
		convey.So(result.End, convey.ShouldEqual, 5)
	})
}
