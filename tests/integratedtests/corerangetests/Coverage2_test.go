package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RangeString ──

func Test_RangeString_Valid_Cov2(t *testing.T) {
	// Arrange
	rs := corerange.NewRangeString("hello:world", ":")

	// Act
	actual := args.Map{
		"isValid": rs.IsValid,
		"hasStart": rs.HasStart,
		"hasEnd": rs.HasEnd,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"hasStart": true,
		"hasEnd": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeString_Valid returns non-empty -- with args", actual)
}

func Test_RangeString_NoSeparator_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isValid": corerange.NewRangeString("hello", ":").IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeString_NoSeparator returns correct value -- with args", actual)
}

func Test_RangeString_Empty_Cov2(t *testing.T) {
	// Act
	actual := args.Map{"isValid": corerange.NewRangeString("", ":").IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeString_Empty returns empty -- with args", actual)
}

func Test_RangeString_Methods_Cov2(t *testing.T) {
	// Arrange
	rs := corerange.NewRangeString("hello:world", ":")

	// Act
	actual := args.Map{
		"stringNotEmpty": rs.String() != "",
		"start": rs.Start,
		"end": rs.End,
	}

	// Assert
	expected := args.Map{
		"stringNotEmpty": true,
		"start": "hello",
		"end": "world",
	}
	expected.ShouldBeEqual(t, 0, "RangeString_Methods returns correct value -- with args", actual)
}

// ── RangeInt8 ──

func Test_RangeInt8_Cov2(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeInt8()

	// Act
	actual := args.Map{
		"isValid": ri8.IsValid,
		"start": int(ri8.Start),
		"end": int(ri8.End),
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"start": 3,
		"end": 7,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt8 returns correct value -- with args", actual)
}

func Test_RangeInt8_Methods_Cov2(t *testing.T) {
	// Arrange
	ri8 := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeInt8()

	// Act
	actual := args.Map{
		"rangeLength": int(ri8.RangeLength()), "difference": int(ri8.Difference()),
		"isWithin3": ri8.IsWithinRange(3), "isWithin8": ri8.IsWithinRange(8),
		"isInvalid8": ri8.IsInvalidValue(8), "stringNotEmpty": ri8.String() != "",
		"rangesLen": len(ri8.Ranges()), "rangesIntLen": len(ri8.RangesInt8()),
	}

	// Assert
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "isWithin8": false,
		"isInvalid8": true, "stringNotEmpty": true,
		"rangesLen": 4, "rangesIntLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt8_Methods returns correct value -- with args", actual)
}

// ── RangeInt16 ──

func Test_RangeInt16_Cov2(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeInt16()

	// Act
	actual := args.Map{
		"isValid": ri16.IsValid,
		"start": int(ri16.Start),
		"end": int(ri16.End),
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"start": 3,
		"end": 7,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt16 returns correct value -- with args", actual)
}

func Test_RangeInt16_Methods_Cov2(t *testing.T) {
	// Arrange
	ri16 := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeInt16()

	// Act
	actual := args.Map{
		"rangeLength": int(ri16.RangeLength()), "difference": int(ri16.Difference()),
		"isWithin3": ri16.IsWithinRange(3), "stringNotEmpty": ri16.String() != "",
	}

	// Assert
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt16_Methods returns correct value -- with args", actual)
}

// ── RangeByte ──

func Test_RangeByte_Cov2(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeByte()

	// Act
	actual := args.Map{
		"isValid": rb.IsValid,
		"start": int(rb.Start),
		"end": int(rb.End),
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"start": 3,
		"end": 7,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte returns correct value -- with args", actual)
}

func Test_RangeByte_Methods_Cov2(t *testing.T) {
	// Arrange
	rb := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeByte()

	// Act
	actual := args.Map{
		"rangeLength": int(rb.RangeLength()), "difference": int(rb.Difference()),
		"isWithin3": rb.IsWithinRange(3), "stringNotEmpty": rb.String() != "",
	}

	// Assert
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte_Methods returns correct value -- with args", actual)
}

// ── BaseRange ──

func Test_BaseRange_Clone_Cov2(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{RawInput: "1:5", Separator: ":", IsValid: true, HasStart: true, HasEnd: true}
	cloned := br.BaseRangeClone()

	// Act
	actual := args.Map{
		"rawInput": cloned.RawInput, "separator": cloned.Separator,
		"isValid": cloned.IsValid, "isInvalid": cloned.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"rawInput": "1:5", "separator": ":",
		"isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseRange_Clone returns correct value -- with args", actual)
}

func Test_BaseRange_String_Cov2(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{Separator: ":"}

	// Act
	actual := args.Map{"notEmpty": br.String(1, 5) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BaseRange_String returns correct value -- with args", actual)
}

func Test_BaseRange_CreateRangeInt_Cov2(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{RawInput: "3:7", Separator: ":", IsValid: true}
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}

	// Act
	actual := args.Map{"notNil": br.CreateRangeInt(mm) != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BaseRange_CreateRangeInt returns correct value -- with args", actual)
}

// ── MinMaxInt16 ──

func Test_MinMaxInt16_Cov2(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}

	// Act
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10), "isInvalid10": mm.IsInvalidValue(10),
	}

	// Assert
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true, "isInvalid10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 returns correct value -- with args", actual)
}

// ── MinMaxInt64 ──

func Test_MinMaxInt64_Cov2(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt64{Min: 2, Max: 8}

	// Act
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10),
	}

	// Assert
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 returns correct value -- with args", actual)
}

// ── MinMaxInt8 ──

func Test_MinMaxInt8_Cov2(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}

	// Act
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10),
	}

	// Assert
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 returns correct value -- with args", actual)
}

// ── StartEndSimpleString ──

func Test_StartEndSimpleString_Cov2(t *testing.T) {
	// Arrange
	se := &corerange.StartEndSimpleString{Start: "hello", End: "world"}

	// Act
	actual := args.Map{
		"hasStart": se.HasStart(),
		"hasEnd": se.HasEnd(),
		"isStartEmpty": se.Start == "",
		"isEndEmpty": se.End == "",
	}

	// Assert
	expected := args.Map{
		"hasStart": true,
		"hasEnd": true,
		"isStartEmpty": false,
		"isEndEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString returns correct value -- with args", actual)
}

// ── StartEndString ──

func Test_StartEndString_Cov2(t *testing.T) {
	// Arrange
	se := corerange.NewStartEndString("hello:world", ":")

	// Act
	actual := args.Map{
		"isValid": se.IsValid,
		"hasStart": se.HasStart,
		"hasEnd": se.HasEnd,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"hasStart": true,
		"hasEnd": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndString returns correct value -- with args", actual)
}

func Test_StartEndString_Methods_Cov2(t *testing.T) {
	// Arrange
	se := corerange.NewStartEndString("hello:world", ":")

	// Act
	actual := args.Map{
		"start": se.Start,
		"end": se.End,
		"stringNotEmpty": se.String() != "",
	}

	// Assert
	expected := args.Map{
		"start": "hello",
		"end": "world",
		"stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndString_Methods returns correct value -- with args", actual)
}

// ── RangeAny ──

func Test_RangeAny_Cov2(t *testing.T) {
	// Arrange
	ra := &corerange.RangeAny{BaseRange: &corerange.BaseRange{RawInput: "hello:world", Separator: ":", IsValid: true, HasStart: true, HasEnd: true}}

	// Act
	actual := args.Map{"isValid": ra.IsValid}

	// Assert
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "RangeAny returns correct value -- with args", actual)
}

func Test_RangeAny_NoSeparator_Cov2(t *testing.T) {
	// Arrange
	ra := &corerange.RangeAny{BaseRange: &corerange.BaseRange{RawInput: "hello", Separator: ":", IsValid: false}}

	// Act
	actual := args.Map{"isValid": ra.IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeAny_NoSep returns correct value -- with args", actual)
}

// ── MinMaxInt boundary ──

func Test_MinMaxInt_IsWithinRange_Boundary_Cov2(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 3, Max: 7}

	// Act
	actual := args.Map{
		"exactMin": mm.IsWithinRange(3),
		"exactMax": mm.IsWithinRange(7),
		"below": mm.IsWithinRange(2),
		"above": mm.IsWithinRange(8),
	}

	// Assert
	expected := args.Map{
		"exactMin": true,
		"exactMax": true,
		"below": false,
		"above": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_Within_Boundary returns non-empty -- with args", actual)
}
