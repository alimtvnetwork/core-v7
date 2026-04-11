package corerangetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MinMaxByte ──

func Test_MinMaxByte_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxByteCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		minVal, _ := input.GetAsInt("min")
		maxVal, _ := input.GetAsInt("max")
		mm := &corerange.MinMaxByte{Min: byte(minVal), Max: byte(maxVal)}

		// Act
		actual := args.Map{
			"difference":     int(mm.Difference()),
			"rangeLength":    int(mm.RangeLength()),
			"rangeLengthInt": mm.RangeLengthInt(),
			"isWithin5":      mm.IsWithinRange(5),
			"isWithin10":     mm.IsWithinRange(10),
			"isInvalid10":    mm.IsInvalidValue(10),
			"isOutOfRange10": mm.IsOutOfRange(10),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxByte_Comparisons(t *testing.T) {
	for caseIndex, testCase := range minMaxByteComparisonCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		minVal, _ := input.GetAsInt("min")
		maxVal, _ := input.GetAsInt("max")
		mm := &corerange.MinMaxByte{Min: byte(minVal), Max: byte(maxVal)}

		// Act
		actual := args.Map{
			"isMinEqual2":      mm.IsMinEqual(2),
			"isMinAboveEqual2": mm.IsMinAboveEqual(2),
			"isMinAbove1":      mm.IsMinAbove(1),
			"isMinLess3":       mm.IsMinLess(3),
			"isMinLessEqual2":  mm.IsMinLessEqual(2),
			"isMaxEqual8":      mm.IsMaxEqual(8),
			"isMaxAboveEqual8": mm.IsMaxAboveEqual(8),
			"isMaxAbove7":      mm.IsMaxAbove(7),
			"isMaxLess9":       mm.IsMaxLess(9),
			"isMaxLessEqual8":  mm.IsMaxLessEqual(8),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxByte_Ranges(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxByte{Min: 2, Max: 5}

	// Act
	ranges := mm.Ranges()
	rangesInt := mm.RangesInt()

	// Assert
	actual := args.Map{
		"rangesLen":    len(ranges),
		"rangesIntLen": len(rangesInt),
	}
	expected := args.Map{
		"rangesLen":    4,
		"rangesIntLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte_Ranges returns correct value -- with args", actual)
}

func Test_MinMaxByte_Clone(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxByte{Min: 1, Max: 10}

	// Act
	cloned := mm.Clone()
	clonedPtr := mm.ClonePtr()

	// Assert
	actual := args.Map{
		"clonedMin":    int(cloned.Min),
		"clonedMax":    int(cloned.Max),
		"clonedPtrNil": clonedPtr == nil,
	}
	expected := args.Map{
		"clonedMin":    1,
		"clonedMax":    10,
		"clonedPtrNil": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte_Clone returns correct value -- with args", actual)
}

func Test_MinMaxByte_CreateMinMaxInt(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxByte{Min: 3, Max: 7}

	// Act
	mmi := mm.CreateMinMaxInt()

	// Assert
	actual := args.Map{
		"isNil": mmi == nil,
	}
	expected := args.Map{
		"isNil": false,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte_CreateMinMaxInt returns correct value -- with args", actual)
}

// ── MinMaxInt ──

func Test_MinMaxInt_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxIntCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		minVal, _ := input.GetAsInt("min")
		maxVal, _ := input.GetAsInt("max")
		mm := &corerange.MinMaxInt{Min: minVal, Max: maxVal}

		// Act
		actual := args.Map{
			"difference":    mm.Difference(),
			"diffAbsolute":  mm.DifferenceAbsolute(),
			"rangeLength":   mm.RangeLength(),
			"isWithin5":     mm.IsWithinRange(5),
			"isWithin10":    mm.IsWithinRange(10),
			"isOutOfRange2": mm.IsOutOfRange(2),
			"stringVal":     mm.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxInt_Comparisons(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 2, Max: 8}

	// Act
	actual := args.Map{
		"isMinEqual2":      mm.IsMinEqual(2),
		"isMinAboveEqual2": mm.IsMinAboveEqual(2),
		"isMinAbove1":      mm.IsMinAbove(1),
		"isMinLess3":       mm.IsMinLess(3),
		"isMinLessEqual2":  mm.IsMinLessEqual(2),
		"isMaxEqual8":      mm.IsMaxEqual(8),
		"isMaxAboveEqual8": mm.IsMaxAboveEqual(8),
		"isMaxAbove7":      mm.IsMaxAbove(7),
		"isMaxLess9":       mm.IsMaxLess(9),
		"isMaxLessEqual8":  mm.IsMaxLessEqual(8),
	}
	expected := args.Map{
		"isMinEqual2":      true,
		"isMinAboveEqual2": true,
		"isMinAbove1":      true,
		"isMinLess3":       true,
		"isMinLessEqual2":  true,
		"isMaxEqual8":      true,
		"isMaxAboveEqual8": true,
		"isMaxAbove7":      true,
		"isMaxLess9":       true,
		"isMaxLessEqual8":  true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_Comparisons returns correct value -- with args", actual)
}

func Test_MinMaxInt_Ranges(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 3, Max: 7}

	// Act
	ranges := mm.Ranges()
	rangesInt := mm.RangesInt()

	// Assert
	actual := args.Map{
		"rangesLen":    len(ranges),
		"rangesIntLen": len(rangesInt),
		"firstRange":   ranges[0],
		"lastRange":    ranges[len(ranges)-1],
	}
	expected := args.Map{
		"rangesLen":    5,
		"rangesIntLen": 5,
		"firstRange":   3,
		"lastRange":    7,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_Ranges returns correct value -- with args", actual)
}

func Test_MinMaxInt_CreateRanges(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 1, Max: 3}
	extra := corerange.MinMaxInt{Min: 10, Max: 12}

	// Act
	combined := mm.CreateRanges(extra)
	noExtra := mm.CreateRanges()

	// Assert
	actual := args.Map{
		"combinedLen": len(combined),
		"noExtraLen":  len(noExtra),
	}
	expected := args.Map{
		"combinedLen": 6,
		"noExtraLen":  3,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_CreateRanges returns correct value -- with args", actual)
}

func Test_MinMaxInt_RangesExcept(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 1, Max: 5}

	// Act
	except := mm.RangesExcept(3)

	// Assert
	actual := args.Map{
		"exceptLen": len(except),
	}
	expected := args.Map{
		"exceptLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_RangesExcept returns correct value -- with args", actual)
}

func Test_MinMaxInt_Clone(t *testing.T) {
	// Arrange
	mm := &corerange.MinMaxInt{Min: 1, Max: 10}

	// Act
	cloned := mm.Clone()
	clonedPtr := mm.ClonePtr()
	var nilMm *corerange.MinMaxInt
	nilClone := nilMm.ClonePtr()

	// Assert
	actual := args.Map{
		"clonedMin":    cloned.Min,
		"clonedMax":    cloned.Max,
		"clonedPtrNil": clonedPtr == nil,
		"nilCloneNil":  nilClone == nil,
	}
	expected := args.Map{
		"clonedMin":    1,
		"clonedMax":    10,
		"clonedPtrNil": false,
		"nilCloneNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_Clone returns correct value -- with args", actual)
}

func Test_MinMaxInt_IsEqual(t *testing.T) {
	// Arrange
	mm1 := &corerange.MinMaxInt{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt{Min: 1, Max: 10}
	mm3 := &corerange.MinMaxInt{Min: 2, Max: 10}

	// Act
	actual := args.Map{
		"sameValues":  mm1.IsEqual(mm2),
		"diffValues":  mm1.IsEqual(mm3),
		"samePtr":     mm1.IsEqual(mm1),
		"bothNil":     (*corerange.MinMaxInt)(nil).IsEqual(nil),
		"leftNilOnly": (*corerange.MinMaxInt)(nil).IsEqual(mm1),
	}
	expected := args.Map{
		"sameValues":  true,
		"diffValues":  false,
		"samePtr":     true,
		"bothNil":     true,
		"leftNilOnly": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_IsEqual returns correct value -- with args", actual)
}

// ── RangeInt ──

func Test_RangeInt_Verification(t *testing.T) {
	for caseIndex, testCase := range rangeIntCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw, _ := input.GetAsString("raw")
		sep, _ := input.GetAsString("separator")
		minVal, _ := input.GetAsInt("min")
		maxVal, _ := input.GetAsInt("max")

		// Act
		ri := corerange.NewRangeIntMinMax(raw, sep, minVal, maxVal)
		actual := args.Map{
			"isValid": ri.IsValid,
		}

		if ri.IsValid || (ri.Start != 0 && ri.End != 0) {
			actual["start"] = ri.Start
			actual["end"] = ri.End
			actual["rangeLength"] = ri.RangeLength()
			actual["difference"] = ri.Difference()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RangeInt_Methods(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("3:7", ":", 0, 10)

	// Act
	actual := args.Map{
		"isWithin5":          ri.IsWithinRange(5),
		"isWithin11":         ri.IsWithinRange(11),
		"isValidPlusWithin5": ri.IsValidPlusWithinRange(5),
		"isInvalidVal11":     ri.IsInvalidValue(11),
		"stringNotEmpty":     ri.String() != "",
		"rangesLen":          len(ri.Ranges()),
		"rangesIntLen":       len(ri.RangesInt()),
	}
	expected := args.Map{
		"isWithin5":          true,
		"isWithin11":         false,
		"isValidPlusWithin5": true,
		"isInvalidVal11":     true,
		"stringNotEmpty":     true,
		"rangesLen":          5,
		"rangesIntLen":       5,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_Methods returns correct value -- with args", actual)
}

func Test_RangeInt_Conversions(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("3:7", ":", 0, 10)

	// Act
	startEnd := ri.CreateStartEnd()
	rangeInt8 := ri.CreateRangeInt8()
	rangeByte := ri.CreateRangeByte()
	rangeInt16 := ri.CreateRangeInt16()
	shallowInt16 := ri.ShallowCreateRangeInt16()
	shallowInt8 := ri.ShallowCreateRangeInt8()
	shallowByte := ri.ShallowCreateRangeByte()

	// Assert
	actual := args.Map{
		"startEndStart":    startEnd.Start,
		"rangeInt8Start":   int(rangeInt8.Start),
		"rangeByteStart":   int(rangeByte.Start),
		"rangeInt16Start":  int(rangeInt16.Start),
		"shallowInt16Nil":  shallowInt16 == nil,
		"shallowInt8Nil":   shallowInt8 == nil,
		"shallowByteNil":   shallowByte == nil,
	}
	expected := args.Map{
		"startEndStart":    3,
		"rangeInt8Start":   3,
		"rangeByteStart":   3,
		"rangeInt16Start":  3,
		"shallowInt16Nil":  false,
		"shallowInt8Nil":   false,
		"shallowByteNil":   false,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_Conversions returns correct value -- with args", actual)
}

func Test_RangeInt_RangesExcept(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("1:5", ":", 0, 10)

	// Act
	except := ri.RangesExcept(3)

	// Assert
	actual := args.Map{
		"exceptLen": len(except),
	}
	expected := args.Map{
		"exceptLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_RangesExcept returns correct value -- with args", actual)
}

func Test_RangeInt_CreateRanges(t *testing.T) {
	// Arrange
	ri := corerange.NewRangeIntMinMax("1:3", ":", 0, 20)
	extra := corerange.MinMaxInt{Min: 10, Max: 12}

	// Act
	combined := ri.CreateRanges(extra)
	noExtra := ri.CreateRanges()

	// Assert
	actual := args.Map{
		"combinedLen": len(combined),
		"noExtraLen":  len(noExtra),
	}
	expected := args.Map{
		"combinedLen": 6,
		"noExtraLen":  3,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_CreateRanges returns correct value -- with args", actual)
}

func Test_RangeInt_NoMinMax(t *testing.T) {
	// Arrange & Act
	ri := corerange.NewRangeInt("3:7", ":", nil)

	// Assert
	actual := args.Map{
		"isValid": ri.IsValid,
		"start":   ri.Start,
		"end":     ri.End,
	}
	expected := args.Map{
		"isValid": true,
		"start":   3,
		"end":     7,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_NoMinMax returns correct value -- with args", actual)
}

func Test_RangeInt_UsingValues(t *testing.T) {
	// Arrange & Act
	ri := corerange.NewRangeIntUsingValues(3, 7, true)

	// Assert
	actual := args.Map{
		"isValid": ri.IsValid,
		"start":   ri.Start,
		"end":     ri.End,
	}
	expected := args.Map{
		"isValid": true,
		"start":   3,
		"end":     7,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt_UsingValues returns non-empty -- with args", actual)
}

// ── StartEndInt ──

func Test_StartEndInt_Verification(t *testing.T) {
	for caseIndex, testCase := range startEndIntCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		start, _ := input.GetAsInt("start")
		end, _ := input.GetAsInt("end")
		se := &corerange.StartEndInt{Start: start, End: end}

		// Act
		actual := args.Map{
			"hasStart":           se.HasStart(),
			"hasEnd":             se.HasEnd(),
			"isInvalidStart":     se.IsInvalidStart(),
			"isInvalidEnd":       se.IsInvalidEnd(),
			"isStartEndBoth":     se.IsStartEndBothDefined(),
			"isInvalidBoth":      se.IsInvalidStartEndBoth(),
			"isInvalidAny":       se.IsInvalidAnyStartEnd(),
			"diff":               se.Diff(),
			"diffAbs":            se.DifferenceAbsolute(),
			"rangeLength":        se.RangeLength(),
			"stringVal":          se.String(),
			"stringSpace":        se.StringSpace(),
			"stringHyphen":       se.StringHyphen(),
			"stringColon":        se.StringColon(),
			"isStartGraterThan2": se.IsStartGraterThan(2),
			"isEndGraterThan5":   se.IsEndGraterThan(5),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StartEndInt_Ranges(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 2, End: 5}

	// Act
	ranges := se.Ranges()
	except := se.RangesExcept(3)
	combined := se.CreateRanges(corerange.StartEndInt{Start: 10, End: 12})
	noExtra := se.CreateRanges()

	// Assert
	actual := args.Map{
		"rangesLen":   len(ranges),
		"exceptLen":   len(except),
		"combinedLen": len(combined),
		"noExtraLen":  len(noExtra),
	}
	expected := args.Map{
		"rangesLen":   4,
		"exceptLen":   3,
		"combinedLen": 7,
		"noExtraLen":  4,
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt_Ranges returns correct value -- with args", actual)
}

func Test_StartEndInt_NilIsInvalid(t *testing.T) {
	// Arrange
	var se *corerange.StartEndInt

	// Act
	actual := args.Map{
		"isInvalid": se.IsInvalid(),
	}
	expected := args.Map{
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt_NilIsInvalid returns nil -- with args", actual)
}

func Test_StartEndInt_StringFormat(t *testing.T) {
	// Arrange
	se := &corerange.StartEndInt{Start: 5, End: 10}

	// Act
	formatted := se.StringUsingFormat("%d to %d")

	// Assert
	actual := args.Map{
		"formatted": formatted,
	}
	expected := args.Map{
		"formatted": "5 to 10",
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt_StringFormat returns correct value -- with args", actual)
}

// ── BaseRange ──

func Test_BaseRange_Verification(t *testing.T) {
	// Arrange
	br := &corerange.BaseRange{
		RawInput:  "3:7",
		Separator: ":",
		IsValid:   true,
		HasStart:  true,
		HasEnd:    true,
	}

	// Act
	cloned := br.BaseRangeClone()
	strVal := br.String(3, 7)

	// Assert
	actual := args.Map{
		"isInvalid":      br.IsInvalid(),
		"clonedNotNil":   cloned != nil,
		"clonedRawInput": cloned.RawInput,
		"stringNotEmpty": strVal != "",
	}
	expected := args.Map{
		"isInvalid":      false,
		"clonedNotNil":   true,
		"clonedRawInput": "3:7",
		"stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseRange_Verification returns correct value -- with args", actual)
}

// ── Within ──

func Test_Within_RangeInteger(t *testing.T) {
	for caseIndex, testCase := range withinIntegerCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsInt("input")
		minVal, _ := input.GetAsInt("min")
		maxVal, _ := input.GetAsInt("max")

		// Act
		val, isInRange := corerange.Within.RangeInteger(true, minVal, maxVal, inputVal)
		actual := args.Map{
			"value":     val,
			"isInRange": isInRange,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Within_RangeInteger_NoBoundary(t *testing.T) {
	// Arrange & Act
	val, isInRange := corerange.Within.RangeInteger(false, 1, 10, 15)

	// Assert
	actual := args.Map{
		"value":     val,
		"isInRange": isInRange,
	}
	expected := args.Map{
		"value":     15,
		"isInRange": false,
	}
	expected.ShouldBeEqual(t, 0, "Within_NoBoundary returns non-empty -- with args", actual)
}

func Test_Within_StringRangeInteger(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.StringRangeInteger(true, 0, 100, "50")
	valBad, okBad := corerange.Within.StringRangeInteger(true, 0, 100, "abc")

	// Assert
	actual := args.Map{
		"val":    val,
		"ok":     ok,
		"valBad": valBad,
		"okBad":  okBad,
	}
	expected := args.Map{
		"val":    50,
		"ok":     true,
		"valBad": 0,
		"okBad":  false,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeInteger returns non-empty -- with args", actual)
}

func Test_Within_StringRangeIntegerDefault(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.StringRangeIntegerDefault(1, 100, "50")
	valLow, okLow := corerange.Within.StringRangeIntegerDefault(1, 100, "0")
	valHigh, okHigh := corerange.Within.StringRangeIntegerDefault(1, 100, "200")

	// Assert
	actual := args.Map{
		"val":     val,
		"ok":      ok,
		"valLow":  valLow,
		"okLow":   okLow,
		"valHigh": valHigh,
		"okHigh":  okHigh,
	}
	expected := args.Map{
		"val":     50,
		"ok":      true,
		"valLow":  1,
		"okLow":   false,
		"valHigh": 100,
		"okHigh":  false,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeIntegerDefault returns non-empty -- with args", actual)
}

func Test_Within_StringRangeTyped(t *testing.T) {
	// Arrange & Act
	valInt32, okInt32 := corerange.Within.StringRangeInt32("50")
	valInt16, okInt16 := corerange.Within.StringRangeInt16("50")
	valInt8, okInt8 := corerange.Within.StringRangeInt8("50")
	valByte, okByte := corerange.Within.StringRangeByte("50")
	valUint16, okUint16 := corerange.Within.StringRangeUint16("50")

	// Assert
	actual := args.Map{
		"valInt32":  fmt.Sprintf("%d", valInt32),
		"okInt32":   okInt32,
		"valInt16":  fmt.Sprintf("%d", valInt16),
		"okInt16":   okInt16,
		"valInt8":   fmt.Sprintf("%d", valInt8),
		"okInt8":    okInt8,
		"valByte":   fmt.Sprintf("%d", valByte),
		"okByte":    okByte,
		"valUint16": fmt.Sprintf("%d", valUint16),
		"okUint16":  okUint16,
	}
	expected := args.Map{
		"valInt32":  "50",
		"okInt32":   true,
		"valInt16":  "50",
		"okInt16":   true,
		"valInt8":   "50",
		"okInt8":    true,
		"valByte":   "50",
		"okByte":    true,
		"valUint16": "50",
		"okUint16":  true,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeTyped returns non-empty -- with args", actual)
}

func Test_Within_RangeByte(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.RangeByteDefault(50)
	valNeg, okNeg := corerange.Within.RangeByte(true, -5)
	valOver, okOver := corerange.Within.RangeByte(true, 300)
	valNoBound, okNoBound := corerange.Within.RangeByte(false, -5)

	// Assert
	actual := args.Map{
		"val":        int(val),
		"ok":         ok,
		"valNeg":     int(valNeg),
		"okNeg":      okNeg,
		"valOver":    int(valOver),
		"okOver":     okOver,
		"valNoBound": int(valNoBound),
		"okNoBound":  okNoBound,
	}
	expected := args.Map{
		"val":        50,
		"ok":         true,
		"valNeg":     0,
		"okNeg":      false,
		"valOver":    255,
		"okOver":     false,
		"valNoBound": 0,
		"okNoBound":  false,
	}
	expected.ShouldBeEqual(t, 0, "Within_RangeByte returns non-empty -- with args", actual)
}

func Test_Within_RangeUint16(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.RangeUint16Default(100)
	valNoBound, okNoBound := corerange.Within.RangeUint16(false, -5)

	// Assert
	actual := args.Map{
		"val":        int(val),
		"ok":         ok,
		"valNoBound": int(valNoBound),
		"okNoBound":  okNoBound,
	}
	expected := args.Map{
		"val":        100,
		"ok":         true,
		"valNoBound": 0,
		"okNoBound":  false,
	}
	expected.ShouldBeEqual(t, 0, "Within_RangeUint16 returns non-empty -- with args", actual)
}

func Test_Within_RangeFloat(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.RangeFloat(true, 1.0, 10.0, 5.0)
	valBelow, okBelow := corerange.Within.RangeFloat(true, 1.0, 10.0, 0.5)
	valAbove, okAbove := corerange.Within.RangeFloat(true, 1.0, 10.0, 15.0)
	valNoBound, okNoBound := corerange.Within.RangeFloat(false, 1.0, 10.0, 15.0)

	// Assert
	actual := args.Map{
		"ok":         ok,
		"okBelow":    okBelow,
		"okAbove":    okAbove,
		"okNoBound":  okNoBound,
		"inRange":    fmt.Sprintf("%.1f", val),
		"below":      fmt.Sprintf("%.1f", valBelow),
		"above":      fmt.Sprintf("%.1f", valAbove),
		"noBound":    fmt.Sprintf("%.1f", valNoBound),
	}
	expected := args.Map{
		"ok":         true,
		"okBelow":    false,
		"okAbove":    false,
		"okNoBound":  false,
		"inRange":    "5.0",
		"below":      "1.0",
		"above":      "10.0",
		"noBound":    "15.0",
	}
	expected.ShouldBeEqual(t, 0, "Within_RangeFloat returns non-empty -- with args", actual)
}

func Test_Within_RangeFloat64(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.RangeFloat64(true, 1.0, 10.0, 5.0)
	valBelow, okBelow := corerange.Within.RangeFloat64(true, 1.0, 10.0, 0.5)

	// Assert
	actual := args.Map{
		"ok":      ok,
		"okBelow": okBelow,
		"inRange": fmt.Sprintf("%.1f", val),
		"below":   fmt.Sprintf("%.1f", valBelow),
	}
	expected := args.Map{
		"ok":      true,
		"okBelow": false,
		"inRange": "5.0",
		"below":   "1.0",
	}
	expected.ShouldBeEqual(t, 0, "Within_RangeFloat64 returns non-empty -- with args", actual)
}

func Test_Within_StringRangeFloat(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.StringRangeFloat(true, 0.0, 100.0, "50.5")
	valBad, okBad := corerange.Within.StringRangeFloat(true, 0.0, 100.0, "abc")

	// Assert
	actual := args.Map{
		"ok":    ok,
		"okBad": okBad,
		"valid": val > 0,
		"bad":   fmt.Sprintf("%.1f", valBad),
	}
	expected := args.Map{
		"ok":    true,
		"okBad": false,
		"valid": true,
		"bad":   "0.0",
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeFloat returns non-empty -- with args", actual)
}

func Test_Within_StringRangeFloat64(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.StringRangeFloat64(true, 0.0, 100.0, "50.5")
	valBad, okBad := corerange.Within.StringRangeFloat64(true, 0.0, 100.0, "abc")

	// Assert
	actual := args.Map{
		"ok":    ok,
		"okBad": okBad,
		"valid": val > 0,
		"bad":   fmt.Sprintf("%.1f", valBad),
	}
	expected := args.Map{
		"ok":    true,
		"okBad": false,
		"valid": true,
		"bad":   "0.0",
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeFloat64 returns non-empty -- with args", actual)
}

func Test_Within_StringRangeFloatDefault(t *testing.T) {
	// Arrange & Act
	_, ok := corerange.Within.StringRangeFloatDefault("50.0")
	_, okBad := corerange.Within.StringRangeFloatDefault("abc")

	// Assert
	actual := args.Map{
		"ok":    ok,
		"okBad": okBad,
	}
	expected := args.Map{
		"ok":    true,
		"okBad": false,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeFloatDefault returns non-empty -- with args", actual)
}

func Test_Within_StringRangeFloat64Default(t *testing.T) {
	// Arrange & Act
	_, ok := corerange.Within.StringRangeFloat64Default("50.0")
	_, okBad := corerange.Within.StringRangeFloat64Default("abc")

	// Assert
	actual := args.Map{
		"ok":    ok,
		"okBad": okBad,
	}
	expected := args.Map{
		"ok":    true,
		"okBad": false,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeFloat64Default returns non-empty -- with args", actual)
}

func Test_Within_StringRangeUint32(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.StringRangeUint32("100")
	_, okBad := corerange.Within.StringRangeUint32("abc")

	// Assert
	actual := args.Map{
		"val":   int(val),
		"ok":    ok,
		"okBad": okBad,
	}
	expected := args.Map{
		"val":   100,
		"ok":    true,
		"okBad": false,
	}
	expected.ShouldBeEqual(t, 0, "Within_StringRangeUint32 returns non-empty -- with args", actual)
}

func Test_Within_RangeDefaultInteger(t *testing.T) {
	// Arrange & Act
	val, ok := corerange.Within.RangeDefaultInteger(1, 10, 5)

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": 5,
		"ok":  true,
	}
	expected.ShouldBeEqual(t, 0, "Within_RangeDefaultInteger returns non-empty -- with args", actual)
}

// ── StartEndString ──

func Test_StartEndString_Verification(t *testing.T) {
	// Arrange
	ses := corerange.NewStartEndString("hello:world", ":")

	// Act
	actual := args.Map{
		"start":     ses.Start,
		"end":       ses.End,
		"hasStart":  ses.HasStart,
		"hasEnd":    ses.HasEnd,
		"strNotEmpty": ses.String() != "",
	}
	expected := args.Map{
		"start":     "hello",
		"end":       "world",
		"hasStart":  true,
		"hasEnd":    true,
		"strNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndString returns correct value -- with args", actual)
}

func Test_RangeString_Verification(t *testing.T) {
	// Arrange
	rs := corerange.NewRangeString("a:b", ":")

	// Act
	actual := args.Map{
		"start":       rs.Start,
		"end":         rs.End,
		"strNotEmpty": rs.String() != "",
	}
	expected := args.Map{
		"start":       "a",
		"end":         "b",
		"strNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeString returns correct value -- with args", actual)
}

// ── StartEndSimpleString ──

func Test_StartEndSimpleString_Verification(t *testing.T) {
	// Arrange
	sess := &corerange.StartEndSimpleString{Start: "hello", End: "world"}

	// Act
	actual := args.Map{
		"hasStart":       sess.HasStart(),
		"hasEnd":         sess.HasEnd(),
		"isInvalidStart": sess.IsInvalidStart(),
		"isInvalidEnd":   sess.IsInvalidEnd(),
		"isBothDefined":  sess.IsStartEndBothDefined(),
		"isInvalidBoth":  sess.IsInvalidStartEndBoth(),
		"isInvalidAny":   sess.IsInvalidAnyStartEnd(),
		"stringSpace":    sess.StringSpace(),
		"stringHyphen":   sess.StringHyphen(),
		"stringColon":    sess.StringColon(),
		"startVVNil":     sess.StartValidValue() == nil,
		"endVVNil":       sess.EndValidValue() == nil,
		"sesNotNil":      sess.StartEndString() != nil,
	}
	expected := args.Map{
		"hasStart":       true,
		"hasEnd":         true,
		"isInvalidStart": false,
		"isInvalidEnd":   false,
		"isBothDefined":  true,
		"isInvalidBoth":  false,
		"isInvalidAny":   false,
		"stringSpace":    "hello world",
		"stringHyphen":   "hello-world",
		"stringColon":    "hello:world",
		"startVVNil":     false,
		"endVVNil":       false,
		"sesNotNil":      true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString returns correct value -- with args", actual)
}

func Test_StartEndSimpleString_Empty(t *testing.T) {
	// Arrange
	sess := &corerange.StartEndSimpleString{}

	// Act
	actual := args.Map{
		"hasStart":       sess.HasStart(),
		"hasEnd":         sess.HasEnd(),
		"isInvalidStart": sess.IsInvalidStart(),
		"isInvalidEnd":   sess.IsInvalidEnd(),
	}
	expected := args.Map{
		"hasStart":       false,
		"hasEnd":         false,
		"isInvalidStart": true,
		"isInvalidEnd":   true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString_Empty returns empty -- with args", actual)
}

func Test_StartEndSimpleString_Nil(t *testing.T) {
	// Arrange
	var sess *corerange.StartEndSimpleString

	// Act
	actual := args.Map{
		"isInvalidStart": sess.IsInvalidStart(),
		"startVVNil":     sess.StartValidValue() == nil,
		"endVVNil":       sess.EndValidValue() == nil,
		"sesNil":         sess.StartEndString() == nil,
	}
	expected := args.Map{
		"isInvalidStart": true,
		"startVVNil":     true,
		"endVVNil":       true,
		"sesNil":         true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString_Nil returns nil -- with args", actual)
}
