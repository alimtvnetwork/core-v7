package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corerange"
)

// ── RangeInt ──

func Test_Cov4_RangeInt_CreateRanges(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 5, true)
	extra := corerange.MinMaxInt{Min: 10, Max: 12}
	ranges := ri.CreateRanges(extra)
	actual := args.Map{
		"len":   len(ranges),
		"first": ranges[0],
		"last":  ranges[len(ranges)-1],
	}
	expected := args.Map{"len": 8, "first": 1, "last": 12}
	expected.ShouldBeEqual(t, 0, "RangeInt returns non-empty -- CreateRanges with extra MinMaxInt", actual)
}

func Test_Cov4_RangeInt_CreateRanges_NoExtra(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 3, true)
	ranges := ri.CreateRanges()
	actual := args.Map{"len": len(ranges)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RangeInt returns empty -- CreateRanges no extra", actual)
}

func Test_Cov4_RangeInt_RangesExcept(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 5, true)
	result := ri.RangesExcept(2, 4)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RangeInt returns correct value -- RangesExcept", actual)
}

func Test_Cov4_RangeInt_Conversions(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 5, true)
	ri8 := ri.CreateRangeInt8()
	riByte := ri.CreateRangeByte()
	ri16 := ri.CreateRangeInt16()
	se := ri.CreateStartEnd()
	actual := args.Map{
		"int8Start":    int(ri8.Start),
		"byteStart":   int(riByte.Start),
		"int16Start":  int(ri16.Start),
		"seStart":     se.Start,
		"shallowI16":  int(ri.ShallowCreateRangeInt16().Start),
		"shallowI8":   int(ri.ShallowCreateRangeInt8().Start),
		"shallowByte": int(ri.ShallowCreateRangeByte().Start),
	}
	expected := args.Map{
		"int8Start": 1, "byteStart": 1, "int16Start": 1,
		"seStart": 1, "shallowI16": 1, "shallowI8": 1, "shallowByte": 1,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt returns correct value -- conversions", actual)
}

func Test_Cov4_RangeInt_IsWithinRange(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 10, true)
	actual := args.Map{
		"within5":          ri.IsWithinRange(5),
		"within0":          ri.IsWithinRange(0),
		"validPlusWithin":  ri.IsValidPlusWithinRange(5),
		"invalidValue":     ri.IsInvalidValue(0),
	}
	expected := args.Map{
		"within5": true, "within0": false,
		"validPlusWithin": true, "invalidValue": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt returns non-empty -- IsWithinRange", actual)
}

func Test_Cov4_RangeInt_String(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(1, 5, true)
	actual := args.Map{"notEmpty": ri.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeInt returns correct value -- String", actual)
}

func Test_Cov4_RangeInt_DifferenceAbsolute_Negative(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(5, 1, false)
	actual := args.Map{"diffAbs": ri.DifferenceAbsolute()}
	expected := args.Map{"diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "RangeInt returns correct value -- DifferenceAbsolute negative", actual)
}

// ── StartEndInt ──

func Test_Cov4_StartEndInt_Methods(t *testing.T) {
	se := &corerange.StartEndInt{Start: 2, End: 10}
	actual := args.Map{
		"invalidStart":   se.IsInvalidStart(),
		"bothDefined":    se.IsStartEndBothDefined(),
		"invalidBoth":    se.IsInvalidStartEndBoth(),
		"invalidAny":     se.IsInvalidAnyStartEnd(),
		"hasStart":       se.HasStart(),
		"hasEnd":         se.HasEnd(),
		"invalidEnd":     se.IsInvalidEnd(),
		"isInvalid":      se.IsInvalid(),
		"startGt1":       se.IsStartGraterThan(1),
		"endGt5":         se.IsEndGraterThan(5),
		"diff":           se.Diff(),
		"diffAbs":        se.DifferenceAbsolute(),
		"rangeLen":       se.RangeLength(),
		"string":         se.String(),
		"stringSpace":    se.StringSpace(),
		"stringHyphen":   se.StringHyphen(),
		"stringColon":    se.StringColon(),
	}
	expected := args.Map{
		"invalidStart": false, "bothDefined": true,
		"invalidBoth": false, "invalidAny": false,
		"hasStart": true, "hasEnd": true,
		"invalidEnd": false, "isInvalid": false,
		"startGt1": true, "endGt5": true,
		"diff": 8, "diffAbs": 8, "rangeLen": 9,
		"string": "2-10", "stringSpace": "2 10",
		"stringHyphen": "2-10", "stringColon": "2:10",
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt returns correct value -- methods", actual)
}

func Test_Cov4_StartEndInt_Nil(t *testing.T) {
	var se *corerange.StartEndInt
	actual := args.Map{
		"invalidStart": se.IsInvalidStart(),
		"hasStart":     se.HasStart(),
		"invalidEnd":   se.IsInvalidEnd(),
		"hasEnd":       se.HasEnd(),
		"isInvalid":    se.IsInvalid(),
		"startGt":      se.IsStartGraterThan(0),
		"endGt":        se.IsEndGraterThan(0),
	}
	expected := args.Map{
		"invalidStart": true, "hasStart": false,
		"invalidEnd": true, "hasEnd": false,
		"isInvalid": true, "startGt": false, "endGt": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt returns nil -- nil receiver", actual)
}

func Test_Cov4_StartEndInt_Ranges(t *testing.T) {
	se := &corerange.StartEndInt{Start: 1, End: 3}
	ranges := se.Ranges()
	actual := args.Map{"len": len(ranges)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "StartEndInt returns correct value -- Ranges", actual)
}

func Test_Cov4_StartEndInt_CreateRanges(t *testing.T) {
	se := &corerange.StartEndInt{Start: 1, End: 3}
	extra := corerange.StartEndInt{Start: 10, End: 12}
	result := se.CreateRanges(extra)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 6}
	expected.ShouldBeEqual(t, 0, "StartEndInt returns correct value -- CreateRanges", actual)
}

func Test_Cov4_StartEndInt_RangesExcept(t *testing.T) {
	se := &corerange.StartEndInt{Start: 1, End: 5}
	result := se.RangesExcept(2, 4)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "StartEndInt returns correct value -- RangesExcept", actual)
}

// ── StartEndSimpleString ──

func Test_Cov4_StartEndSimpleString_Methods(t *testing.T) {
	ss := &corerange.StartEndSimpleString{Start: "abc", End: "xyz"}
	actual := args.Map{
		"invalidStart":  ss.IsInvalidStart(),
		"bothDefined":   ss.IsStartEndBothDefined(),
		"invalidBoth":   ss.IsInvalidStartEndBoth(),
		"invalidAny":    ss.IsInvalidAnyStartEnd(),
		"hasStart":      ss.HasStart(),
		"hasEnd":        ss.HasEnd(),
		"invalidEnd":    ss.IsInvalidEnd(),
		"stringSpace":   ss.StringSpace(),
		"stringHyphen":  ss.StringHyphen(),
		"stringColon":   ss.StringColon(),
		"startVVNotNil": ss.StartValidValue() != nil,
		"endVVNotNil":   ss.EndValidValue() != nil,
	}
	expected := args.Map{
		"invalidStart": false, "bothDefined": true,
		"invalidBoth": false, "invalidAny": false,
		"hasStart": true, "hasEnd": true, "invalidEnd": false,
		"stringSpace": "abc xyz", "stringHyphen": "abc-xyz",
		"stringColon": "abc:xyz",
		"startVVNotNil": true, "endVVNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString returns correct value -- methods", actual)
}

func Test_Cov4_StartEndSimpleString_Nil(t *testing.T) {
	var ss *corerange.StartEndSimpleString
	actual := args.Map{
		"invalidStart": ss.IsInvalidStart(),
		"hasStart":     ss.HasStart(),
		"invalidEnd":   ss.IsInvalidEnd(),
		"hasEnd":       ss.HasEnd(),
		"startVV":      ss.StartValidValue() == nil,
		"endVV":        ss.EndValidValue() == nil,
		"startEnd":     ss.StartEndString() == nil,
	}
	expected := args.Map{
		"invalidStart": true, "hasStart": false,
		"invalidEnd": true, "hasEnd": false,
		"startVV": true, "endVV": true, "startEnd": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString returns nil -- nil", actual)
}

func Test_Cov4_StartEndSimpleString_StartEndString(t *testing.T) {
	ss := &corerange.StartEndSimpleString{Start: "a", End: "b"}
	result := ss.StartEndString()
	actual := args.Map{
		"notNil": result != nil,
		"start":  result.Start,
		"end":    result.End,
	}
	expected := args.Map{"notNil": true, "start": "a", "end": "b"}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString returns correct value -- StartEndString", actual)
}

// ── RangeAny ──

func Test_Cov4_RangeAny_Methods(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{
			RawInput:  "1|5",
			Separator: "|",
			IsValid:   true,
			HasStart:  true,
			HasEnd:    true,
		},
		RawInput: "1|5",
		Start:    1,
		End:      5,
	}
	actual := args.Map{
		"rawInputStr": ra.RawInputString(),
		"startStr":    ra.StartString(),
		"endStr":      ra.EndString(),
		"stringVal":   ra.String() != "",
	}
	expected := args.Map{
		"rawInputStr": "1|5",
		"startStr":    "1",
		"endStr":      "5",
		"stringVal":   true,
	}
	expected.ShouldBeEqual(t, 0, "RangeAny returns correct value -- methods", actual)
}

// ── Within ──

func Test_Cov4_Within_RangeInteger(t *testing.T) {
	val, ok := corerange.Within.RangeInteger(true, 0, 100, 50)
	val2, ok2 := corerange.Within.RangeInteger(true, 0, 100, -5)
	val3, ok3 := corerange.Within.RangeInteger(true, 0, 100, 200)
	val4, ok4 := corerange.Within.RangeInteger(false, 0, 100, -5)
	actual := args.Map{
		"inRange": val, "inRangeOk": ok,
		"belowMin": val2, "belowMinOk": ok2,
		"aboveMax": val3, "aboveMaxOk": ok3,
		"noBoundBelow": val4, "noBoundBelowOk": ok4,
	}
	expected := args.Map{
		"inRange": 50, "inRangeOk": true,
		"belowMin": 0, "belowMinOk": false,
		"aboveMax": 100, "aboveMaxOk": false,
		"noBoundBelow": -5, "noBoundBelowOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeInteger", actual)
}

func Test_Cov4_Within_RangeDefaultInteger(t *testing.T) {
	val, ok := corerange.Within.RangeDefaultInteger(0, 10, 5)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 5, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeDefaultInteger", actual)
}

func Test_Cov4_Within_StringRange_Types(t *testing.T) {
	v8, ok8 := corerange.Within.StringRangeInt8("5")
	v16, ok16 := corerange.Within.StringRangeInt16("100")
	v32, ok32 := corerange.Within.StringRangeInt32("1000")
	vB, okB := corerange.Within.StringRangeByte("200")
	vU16, okU16 := corerange.Within.StringRangeUint16("500")
	actual := args.Map{
		"v8": int(v8), "ok8": ok8,
		"v16": int(v16), "ok16": ok16,
		"v32": int(v32), "ok32": ok32,
		"vB": int(vB), "okB": okB,
		"vU16": int(vU16), "okU16": okU16,
	}
	expected := args.Map{
		"v8": 5, "ok8": true,
		"v16": 100, "ok16": true,
		"v32": 1000, "ok32": true,
		"vB": 200, "okB": true,
		"vU16": 500, "okU16": true,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- StringRange type variants", actual)
}

func Test_Cov4_Within_StringRangeInteger_InvalidInput(t *testing.T) {
	_, ok := corerange.Within.StringRangeInteger(true, 0, 100, "abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within returns error -- StringRangeInteger invalid", actual)
}

func Test_Cov4_Within_RangeByte(t *testing.T) {
	v, ok := corerange.Within.RangeByteDefault(100)
	v2, ok2 := corerange.Within.RangeByte(true, -1)
	v3, ok3 := corerange.Within.RangeByte(true, 300)
	v4, ok4 := corerange.Within.RangeByte(false, -1)
	actual := args.Map{
		"v": int(v), "ok": ok,
		"belowV": int(v2), "belowOk": ok2,
		"aboveV": int(v3), "aboveOk": ok3,
		"noBoundV": int(v4), "noBoundOk": ok4,
	}
	expected := args.Map{
		"v": 100, "ok": true,
		"belowV": 0, "belowOk": false,
		"aboveV": 255, "aboveOk": false,
		"noBoundV": 0, "noBoundOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeByte", actual)
}

func Test_Cov4_Within_RangeUint16(t *testing.T) {
	v, ok := corerange.Within.RangeUint16Default(500)
	v2, ok2 := corerange.Within.RangeUint16(false, -1)
	actual := args.Map{
		"v": int(v), "ok": ok,
		"noBoundV": int(v2), "noBoundOk": ok2,
	}
	expected := args.Map{
		"v": 500, "ok": true,
		"noBoundV": 0, "noBoundOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeUint16", actual)
}

func Test_Cov4_Within_RangeFloat(t *testing.T) {
	v, ok := corerange.Within.RangeFloat(true, 0, 100, 50)
	v2, ok2 := corerange.Within.RangeFloat(true, 0, 100, -5)
	v3, ok3 := corerange.Within.RangeFloat(false, 0, 100, -5)
	actual := args.Map{
		"inV": int(v), "inOk": ok,
		"belowV": int(v2), "belowOk": ok2,
		"noBoundV": int(v3), "noBoundOk": ok3,
	}
	expected := args.Map{
		"inV": 50, "inOk": true,
		"belowV": 0, "belowOk": false,
		"noBoundV": -5, "noBoundOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeFloat", actual)
}

func Test_Cov4_Within_RangeFloat64(t *testing.T) {
	v, ok := corerange.Within.RangeFloat64(true, 0, 100, 50)
	v2, ok2 := corerange.Within.RangeFloat64(true, 0, 100, 200)
	actual := args.Map{
		"v": int(v), "ok": ok,
		"aboveV": int(v2), "aboveOk": ok2,
	}
	expected := args.Map{
		"v": 50, "ok": true,
		"aboveV": 100, "aboveOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- RangeFloat64", actual)
}

func Test_Cov4_Within_StringRangeFloat_InvalidInput(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat(true, 0, 100, "abc")
	_, ok64 := corerange.Within.StringRangeFloat64(true, 0, 100, "abc")
	actual := args.Map{"ok": ok, "ok64": ok64}
	expected := args.Map{"ok": false, "ok64": false}
	expected.ShouldBeEqual(t, 0, "Within returns error -- StringRangeFloat invalid", actual)
}

func Test_Cov4_Within_StringRangeIntegerDefault(t *testing.T) {
	v, ok := corerange.Within.StringRangeIntegerDefault(0, 100, "50")
	v2, ok2 := corerange.Within.StringRangeIntegerDefault(0, 100, "abc")
	v3, ok3 := corerange.Within.StringRangeIntegerDefault(0, 100, "-5")
	v4, ok4 := corerange.Within.StringRangeIntegerDefault(0, 100, "200")
	actual := args.Map{
		"v": v, "ok": ok,
		"errV": v2, "errOk": ok2,
		"belowV": v3, "belowOk": ok3,
		"aboveV": v4, "aboveOk": ok4,
	}
	expected := args.Map{
		"v": 50, "ok": true,
		"errV": 0, "errOk": false,
		"belowV": 0, "belowOk": false,
		"aboveV": 100, "aboveOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Within returns non-empty -- StringRangeIntegerDefault", actual)
}

// ── MinMaxInt ──

func Test_Cov4_MinMaxInt_IsEqual(t *testing.T) {
	a := &corerange.MinMaxInt{Min: 1, Max: 10}
	b := &corerange.MinMaxInt{Min: 1, Max: 10}
	c := &corerange.MinMaxInt{Min: 2, Max: 10}
	var nilMM *corerange.MinMaxInt
	actual := args.Map{
		"equal":    a.IsEqual(b),
		"notEqual": a.IsEqual(c),
		"selfEq":   a.IsEqual(a),
		"nilBoth":  nilMM.IsEqual(nil),
		"nilLeft":  nilMM.IsEqual(a),
	}
	expected := args.Map{
		"equal": true, "notEqual": false,
		"selfEq": true, "nilBoth": true, "nilLeft": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt returns correct value -- IsEqual", actual)
}

func Test_Cov4_MinMaxInt_String(t *testing.T) {
	mm := corerange.MinMaxInt{Min: 1, Max: 10}
	actual := args.Map{"notEmpty": mm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt returns correct value -- String", actual)
}

func Test_Cov4_MinMaxInt_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 1, Max: 5}
	result := mm.RangesExcept(2, 4)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt returns correct value -- RangesExcept", actual)
}

func Test_Cov4_MinMaxInt_IsOutOfRange(t *testing.T) {
	mm := corerange.MinMaxInt{Min: 1, Max: 10}
	actual := args.Map{
		"outOf0":  mm.IsOutOfRange(0),
		"outOf5":  mm.IsOutOfRange(5),
		"outOf11": mm.IsOutOfRange(11),
	}
	expected := args.Map{"outOf0": true, "outOf5": false, "outOf11": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt returns correct value -- IsOutOfRange", actual)
}

// ── MinMaxByte ──

func Test_Cov4_MinMaxByte_Methods(t *testing.T) {
	mb := &corerange.MinMaxByte{Min: 1, Max: 10}
	actual := args.Map{
		"diff":       int(mb.Difference()),
		"rangeLen":   int(mb.RangeLength()),
		"rangeLenI":  mb.RangeLengthInt(),
		"rangesLen":  len(mb.Ranges()),
		"rangesILen": len(mb.RangesInt()),
		"within5":    mb.IsWithinRange(5),
		"within0":    mb.IsWithinRange(0),
		"invalidV":   mb.IsInvalidValue(0),
		"outOfRange": mb.IsOutOfRange(0),
		"minEq1":     mb.IsMinEqual(1),
		"minAboveEq": mb.IsMinAboveEqual(1),
		"minAbove":   mb.IsMinAbove(0),
		"minLess":    mb.IsMinLess(5),
		"minLessEq":  mb.IsMinLessEqual(1),
		"maxEq10":    mb.IsMaxEqual(10),
		"maxAboveEq": mb.IsMaxAboveEqual(10),
		"maxAbove":   mb.IsMaxAbove(5),
		"maxLess":    mb.IsMaxLess(20),
		"maxLessEq":  mb.IsMaxLessEqual(10),
	}
	expected := args.Map{
		"diff": 9, "rangeLen": 10, "rangeLenI": 10,
		"rangesLen": 10, "rangesILen": 10,
		"within5": true, "within0": false,
		"invalidV": true, "outOfRange": true,
		"minEq1": true, "minAboveEq": true, "minAbove": true,
		"minLess": true, "minLessEq": true,
		"maxEq10": true, "maxAboveEq": true, "maxAbove": true,
		"maxLess": true, "maxLessEq": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte returns correct value -- methods", actual)
}

func Test_Cov4_MinMaxByte_Clone(t *testing.T) {
	mb := &corerange.MinMaxByte{Min: 1, Max: 10}
	cloned := mb.ClonePtr()
	clonedV := mb.Clone()
	var nilMB *corerange.MinMaxByte
	actual := args.Map{
		"clonedMin": int(cloned.Min),
		"clonedMax": int(cloned.Max),
		"valueMin":  int(clonedV.Min),
		"nilClone":  nilMB.ClonePtr() == nil,
	}
	expected := args.Map{
		"clonedMin": 1, "clonedMax": 10,
		"valueMin": 1, "nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxByte returns correct value -- Clone", actual)
}

// ── RangeByte ──

func Test_Cov4_RangeByte_Difference(t *testing.T) {
	rb := corerange.NewRangeByteMinMax("1|5", "|", 0, 255)
	actual := args.Map{
		"diffAbs": int(rb.DifferenceAbsolute()),
		"rangeL":  int(rb.RangeLength()),
	}
	expected := args.Map{
		"diffAbs": 4,
		"rangeL":  5,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte returns correct value -- Difference", actual)
}

// ── RangeString ──

func Test_Cov4_RangeString_Methods(t *testing.T) {
	rs := corerange.NewRangeString("hello:world", ":")
	actual := args.Map{
		"start":    rs.Start,
		"end":      rs.End,
		"notEmpty": rs.String() != "",
	}
	expected := args.Map{
		"start": "hello", "end": "world", "notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeString returns correct value -- methods", actual)
}

// ── StartEndString ──

func Test_Cov4_StartEndString_UsingLines(t *testing.T) {
	se := corerange.NewStartEndStringUsingLines([]string{"first", "last"})
	actual := args.Map{
		"start": se.Start, "end": se.End,
		"isValid": se.IsValid,
	}
	expected := args.Map{
		"start": "first", "end": "last", "isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "StartEndString returns correct value -- UsingLines", actual)
}

func Test_Cov4_StartEndString_CreateRangeString(t *testing.T) {
	se := corerange.NewStartEndString("a:b", ":")
	rs := se.CreateRangeString()
	actual := args.Map{"start": rs.Start, "end": rs.End}
	expected := args.Map{"start": "a", "end": "b"}
	expected.ShouldBeEqual(t, 0, "StartEndString returns correct value -- CreateRangeString", actual)
}
