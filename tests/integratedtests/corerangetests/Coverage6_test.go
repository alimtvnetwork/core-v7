package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── NewStartEndStringUsingLines ──

func Test_Cov6_NewStartEndStringUsingLines_Valid(t *testing.T) {
	ses := corerange.NewStartEndStringUsingLines([]string{"hello", "world"})
	actual := args.Map{
		"isValid":  ses.IsValid,
		"hasStart": ses.HasStart,
		"hasEnd":   ses.HasEnd,
		"start":    ses.Start,
		"end":      ses.End,
	}
	expected := args.Map{"isValid": true, "hasStart": true, "hasEnd": true, "start": "hello", "end": "world"}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns valid -- two lines", actual)
}

func Test_Cov6_NewStartEndStringUsingLines_Single(t *testing.T) {
	ses := corerange.NewStartEndStringUsingLines([]string{"only"})
	actual := args.Map{
		"isValid":  ses.IsValid,
		"hasStart": ses.HasStart,
		"hasEnd":   ses.HasEnd,
		"start":    ses.Start,
	}
	expected := args.Map{"isValid": false, "hasStart": true, "hasEnd": false, "start": "only"}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns invalid -- single line", actual)
}

func Test_Cov6_NewStartEndStringUsingLines_Empty(t *testing.T) {
	ses := corerange.NewStartEndStringUsingLines([]string{})
	actual := args.Map{"isValid": ses.IsValid, "hasStart": ses.HasStart, "hasEnd": ses.HasEnd}
	expected := args.Map{"isValid": false, "hasStart": false, "hasEnd": false}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns invalid -- empty", actual)
}

func Test_Cov6_NewStartEndStringUsingLines_ThreePlus(t *testing.T) {
	ses := corerange.NewStartEndStringUsingLines([]string{"a", "b", "c"})
	actual := args.Map{
		"isValid": ses.IsValid,
		"start":   ses.Start,
		"end":     ses.End,
	}
	expected := args.Map{"isValid": false, "start": "a", "end": "c"}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns invalid -- three lines", actual)
}

// ── StartEndString.CreateRangeString ──

func Test_Cov6_StartEndString_CreateRangeString(t *testing.T) {
	ses := corerange.NewStartEndString("hello:world", ":")
	rs := ses.CreateRangeString()
	actual := args.Map{"notNil": rs != nil, "start": rs.Start, "end": rs.End}
	expected := args.Map{"notNil": true, "start": "hello", "end": "world"}
	expected.ShouldBeEqual(t, 0, "StartEndString CreateRangeString returns valid -- delegates", actual)
}

// ── RangeAny — uncovered methods ──

func Test_Cov6_RangeAny_CreateRangeInt(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{RawInput: "3|7", Separator: "|", IsValid: true, HasStart: true, HasEnd: true},
		RawInput:  "3|7",
		Start:     3,
		End:       7,
	}
	ri := ra.CreateRangeInt()
	actual := args.Map{"notNil": ri != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeAny CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov6_RangeAny_CreateRangeIntMinMax(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{RawInput: "3|7", Separator: "|", IsValid: true, HasStart: true, HasEnd: true},
		RawInput:  "3|7",
		Start:     3,
		End:       7,
	}
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}
	ri := ra.CreateRangeIntMinMax(mm)
	actual := args.Map{"notNil": ri != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeAny CreateRangeIntMinMax returns non-nil -- with minMax", actual)
}

func Test_Cov6_RangeAny_CreateRangeString(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{RawInput: "hello:world", Separator: ":", IsValid: true, HasStart: true, HasEnd: true},
		RawInput:  "hello:world",
		Start:     "hello",
		End:       "world",
	}
	rs := ra.CreateRangeString()
	actual := args.Map{"notNil": rs != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeAny CreateRangeString returns non-nil -- valid", actual)
}

func Test_Cov6_RangeAny_CreateStartEndString(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{RawInput: "a:b", Separator: ":", IsValid: true, HasStart: true, HasEnd: true},
		RawInput:  "a:b",
		Start:     "a",
		End:       "b",
	}
	ses := ra.CreateStartEndString()
	actual := args.Map{"notNil": ses != nil, "start": ses.Start, "end": ses.End}
	expected := args.Map{"notNil": true, "start": "a", "end": "b"}
	expected.ShouldBeEqual(t, 0, "RangeAny CreateStartEndString returns valid -- cloned base", actual)
}

// ── MinMaxInt64 — uncovered methods ──

func Test_Cov6_MinMaxInt64_Ranges(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 3, Max: 7}
	ranges := mm.Ranges()
	rangesInt := mm.RangesInt()
	actual := args.Map{"rangesLen": len(ranges), "rangesIntLen": len(rangesInt), "rangeLenInt": mm.RangeLengthInt()}
	expected := args.Map{"rangesLen": 5, "rangesIntLen": 5, "rangeLenInt": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 Ranges returns 5 -- range 3 to 7", actual)
}

func Test_Cov6_MinMaxInt64_CreateRanges(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 3}
	extra := corerange.MinMaxInt64{Min: 10, Max: 12}
	combined := mm.CreateRanges(extra)
	noExtra := mm.CreateRanges()
	actual := args.Map{"combinedLen": len(combined), "noExtraLen": len(noExtra)}
	expected := args.Map{"combinedLen": 6, "noExtraLen": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRanges returns combined -- with extra", actual)
}

func Test_Cov6_MinMaxInt64_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 5}
	result := mm.RangesExcept(3)
	actual := args.Map{"exceptLen": len(result)}
	expected := args.Map{"exceptLen": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 RangesExcept returns 4 -- excluding 3", actual)
}

func Test_Cov6_MinMaxInt64_Clone(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 10}
	cloned := mm.Clone()
	clonedPtr := mm.ClonePtr()
	var nilMm *corerange.MinMaxInt64
	actual := args.Map{
		"clonedMin": int(cloned.Min), "clonedMax": int(cloned.Max),
		"clonedPtrNil": clonedPtr == nil, "nilCloneNil": nilMm.ClonePtr() == nil,
	}
	expected := args.Map{"clonedMin": 1, "clonedMax": 10, "clonedPtrNil": false, "nilCloneNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 Clone returns copy -- valid", actual)
}

func Test_Cov6_MinMaxInt64_IsEqual(t *testing.T) {
	mm1 := &corerange.MinMaxInt64{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt64{Min: 1, Max: 10}
	mm3 := &corerange.MinMaxInt64{Min: 2, Max: 10}
	var nilMm *corerange.MinMaxInt64
	actual := args.Map{
		"sameValues": mm1.IsEqual(mm2), "diffValues": mm1.IsEqual(mm3),
		"samePtr": mm1.IsEqual(mm1), "bothNil": nilMm.IsEqual(nil),
		"leftNilOnly": nilMm.IsEqual(mm1),
	}
	expected := args.Map{
		"sameValues": true, "diffValues": false, "samePtr": true,
		"bothNil": true, "leftNilOnly": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 IsEqual returns expected -- various combos", actual)
}

func Test_Cov6_MinMaxInt64_String(t *testing.T) {
	mm := corerange.MinMaxInt64{Min: 2, Max: 8}
	actual := args.Map{"notEmpty": mm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 String returns non-empty -- valid", actual)
}

func Test_Cov6_MinMaxInt64_CreateMinMaxInt(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 2, Max: 8}
	mmi := mm.CreateMinMaxInt()
	actual := args.Map{"notNil": mmi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateMinMaxInt returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt64_CreateRangeInt(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt("3:7", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt64_CreateRangeInt8(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:7", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt64_CreateRangeInt16(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:7", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt16 returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt64_Comparisons(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 2, Max: 8}
	actual := args.Map{
		"isMinEqual2": mm.IsMinEqual(2), "isMinAboveEqual2": mm.IsMinAboveEqual(2),
		"isMinAbove1": mm.IsMinAbove(1), "isMinLess3": mm.IsMinLess(3),
		"isMinLessEqual2": mm.IsMinLessEqual(2), "isMaxEqual8": mm.IsMaxEqual(8),
		"isMaxAboveEqual8": mm.IsMaxAboveEqual(8), "isMaxAbove7": mm.IsMaxAbove(7),
		"isMaxLess9": mm.IsMaxLess(9), "isMaxLessEqual8": mm.IsMaxLessEqual(8),
	}
	expected := args.Map{
		"isMinEqual2": true, "isMinAboveEqual2": true, "isMinAbove1": true,
		"isMinLess3": true, "isMinLessEqual2": true, "isMaxEqual8": true,
		"isMaxAboveEqual8": true, "isMaxAbove7": true, "isMaxLess9": true,
		"isMaxLessEqual8": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 Comparisons returns true -- matching values", actual)
}

// ── MinMaxInt8 — uncovered factory methods ──

func Test_Cov6_MinMaxInt8_CreateMinMaxInt(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	mmi := mm.CreateMinMaxInt()
	actual := args.Map{"notNil": mmi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateMinMaxInt returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt8_CreateRangeInt(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt8_CreateRangeInt8(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt8_CreateRangeInt16(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt16 returns non-nil -- valid", actual)
}

func Test_Cov6_MinMaxInt8_String(t *testing.T) {
	mm := corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notEmpty": mm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 String returns non-empty -- valid", actual)
}

// ── NewRangeByte with nil minMax ──

func Test_Cov6_NewRangeByte_NilMinMax(t *testing.T) {
	rb := corerange.NewRangeByte("3:7", ":", nil)
	actual := args.Map{"notNil": rb != nil, "start": int(rb.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeByte returns valid -- nil minMax uses defaults", actual)
}

// ── NewRangeInt8 with nil minMaxInt8 ──

func Test_Cov6_NewRangeInt8_NilMinMax(t *testing.T) {
	ri8 := corerange.NewRangeInt8("3:7", ":", nil)
	actual := args.Map{"notNil": ri8 != nil, "start": int(ri8.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt8 returns valid -- nil minMax uses defaults", actual)
}

// ── NewRangeInt16 with nil minMaxInt16 ──

func Test_Cov6_NewRangeInt16_NilMinMax(t *testing.T) {
	ri16 := corerange.NewRangeInt16("3:7", ":", nil)
	actual := args.Map{"notNil": ri16 != nil, "start": int(ri16.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt16 returns valid -- nil minMax uses defaults", actual)
}

// ── StartEndSimpleString.StringUsingFormat ──

func Test_Cov6_StartEndSimpleString_StringUsingFormat(t *testing.T) {
	ss := &corerange.StartEndSimpleString{Start: "abc", End: "xyz"}
	result := ss.StringUsingFormat("%s to %s")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "abc to xyz"}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString StringUsingFormat returns formatted -- custom format", actual)
}

// ── StartEndInt.DifferenceAbsolute negative ──

func Test_Cov6_StartEndInt_DifferenceAbsolute_Negative(t *testing.T) {
	se := &corerange.StartEndInt{Start: 10, End: 3}
	actual := args.Map{"diffAbs": se.DifferenceAbsolute(), "diff": se.Diff()}
	expected := args.Map{"diffAbs": 7, "diff": -7}
	expected.ShouldBeEqual(t, 0, "StartEndInt DifferenceAbsolute returns positive -- reversed range", actual)
}

// ── RangeByte.Difference where Start > End ──

func Test_Cov6_RangeByte_Difference_StartGreaterThanEnd(t *testing.T) {
	rb := &corerange.RangeByte{
		BaseRange: &corerange.BaseRange{IsValid: false},
		Start:     7,
		End:       3,
	}
	actual := args.Map{"diff": int(rb.Difference())}
	expected := args.Map{"diff": 4}
	expected.ShouldBeEqual(t, 0, "RangeByte Difference returns positive -- Start > End branch", actual)
}

// ── MinMaxInt8 — CreateRanges with extra ──

func Test_Cov6_MinMaxInt8_CreateRanges(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 3}
	extra := corerange.MinMaxInt8{Min: 10, Max: 12}
	combined := mm.CreateRanges(extra)
	noExtra := mm.CreateRanges()
	actual := args.Map{"combinedLen": len(combined), "noExtraLen": len(noExtra)}
	expected := args.Map{"combinedLen": 6, "noExtraLen": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRanges returns combined -- with extra", actual)
}

func Test_Cov6_MinMaxInt8_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 5}
	result := mm.RangesExcept(3)
	actual := args.Map{"exceptLen": len(result)}
	expected := args.Map{"exceptLen": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 RangesExcept returns 4 -- excluding 3", actual)
}

func Test_Cov6_MinMaxInt8_IsEqual(t *testing.T) {
	mm1 := &corerange.MinMaxInt8{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt8{Min: 1, Max: 10}
	mm3 := &corerange.MinMaxInt8{Min: 2, Max: 10}
	var nilMm *corerange.MinMaxInt8
	actual := args.Map{
		"sameValues": mm1.IsEqual(mm2), "diffValues": mm1.IsEqual(mm3),
		"samePtr": mm1.IsEqual(mm1), "bothNil": nilMm.IsEqual(nil),
		"leftNilOnly": nilMm.IsEqual(mm1),
	}
	expected := args.Map{
		"sameValues": true, "diffValues": false, "samePtr": true,
		"bothNil": true, "leftNilOnly": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 IsEqual returns expected -- various combos", actual)
}

func Test_Cov6_MinMaxInt8_Clone(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 10}
	cloned := mm.Clone()
	clonedPtr := mm.ClonePtr()
	var nilMm *corerange.MinMaxInt8
	actual := args.Map{
		"clonedMin": int(cloned.Min), "clonedMax": int(cloned.Max),
		"clonedPtrNil": clonedPtr == nil, "nilCloneNil": nilMm.ClonePtr() == nil,
	}
	expected := args.Map{"clonedMin": 1, "clonedMax": 10, "clonedPtrNil": false, "nilCloneNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 Clone returns copy -- valid", actual)
}

// ── MinMaxByte — missing IsEqual ──

func Test_Cov6_MinMaxByte_CreateRangeRanges(t *testing.T) {
	mb := &corerange.MinMaxByte{Min: 1, Max: 5}
	actual := args.Map{
		"diffAbs": int(mb.DifferenceAbsolute()),
	}
	expected := args.Map{"diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxByte DifferenceAbsolute returns correct -- valid", actual)
}

// ── NewRangeByteMinMax ──

func Test_Cov6_NewRangeByteMinMax(t *testing.T) {
	rb := corerange.NewRangeByteMinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": rb != nil, "isValid": rb.IsValid, "start": int(rb.Start)}
	expected := args.Map{"notNil": true, "isValid": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeByteMinMax returns valid -- in bounds", actual)
}

// ── NewRangeInt8MinMax ──

func Test_Cov6_NewRangeInt8MinMax(t *testing.T) {
	ri8 := corerange.NewRangeInt8MinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": ri8 != nil, "isValid": ri8.IsValid, "start": int(ri8.Start)}
	expected := args.Map{"notNil": true, "isValid": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt8MinMax returns valid -- in bounds", actual)
}

// ── NewRangeInt16MinMax ──

func Test_Cov6_NewRangeInt16MinMax(t *testing.T) {
	ri16 := corerange.NewRangeInt16MinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": ri16 != nil, "isValid": ri16.IsValid, "start": int(ri16.Start)}
	expected := args.Map{"notNil": true, "isValid": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt16MinMax returns valid -- in bounds", actual)
}
