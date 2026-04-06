package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Within — StringRange* functions ──

func Test_Cov3_Within_StringRangeInt32_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt32("100")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 100, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt32 returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeInt32_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeInt32("abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt32 returns false -- invalid input", actual)
}

func Test_Cov3_Within_StringRangeInt16_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt16("50")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 50, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt16 returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeInt8_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt8("10")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 10, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInt8 returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeByte_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeByte("200")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 200, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeByte returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeByte_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeByte("abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeByte returns false -- invalid input", actual)
}

func Test_Cov3_Within_StringRangeUint16_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeUint16("1000")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 1000, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeUint16 returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeUint32_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeUint32("5000")
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 5000, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeUint32 returns value -- valid input", actual)
}

func Test_Cov3_Within_StringRangeIntegerDefault_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeIntegerDefault(1, 100, "50")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 50, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeIntegerDefault returns value -- in range", actual)
}

func Test_Cov3_Within_StringRangeIntegerDefault_BelowMin(t *testing.T) {
	val, ok := corerange.Within.StringRangeIntegerDefault(10, 100, "5")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 10, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeIntegerDefault returns min -- below min", actual)
}

func Test_Cov3_Within_StringRangeIntegerDefault_AboveMax(t *testing.T) {
	val, ok := corerange.Within.StringRangeIntegerDefault(10, 100, "200")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 100, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeIntegerDefault returns max -- above max", actual)
}

func Test_Cov3_Within_StringRangeIntegerDefault_ParseErr(t *testing.T) {
	_, ok := corerange.Within.StringRangeIntegerDefault(10, 100, "abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeIntegerDefault returns false -- parse error", actual)
}

func Test_Cov3_Within_StringRangeInteger_NoBoundary(t *testing.T) {
	val, ok := corerange.Within.StringRangeInteger(false, 1, 10, "50")
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 50, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeInteger returns raw -- no boundary", actual)
}

func Test_Cov3_Within_StringRangeFloat_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeFloat(true, 1.0, 100.0, "50.5")
	actual := args.Map{"ok": ok, "inRange": val > 50.0}
	expected := args.Map{"ok": true, "inRange": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat returns value -- valid", actual)
}

func Test_Cov3_Within_StringRangeFloat_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat(true, 1.0, 100.0, "abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat returns false -- invalid input", actual)
}

func Test_Cov3_Within_StringRangeFloat_NoBoundary(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat(false, 1.0, 10.0, "50.5")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat returns zero -- no boundary out of range", actual)
}

func Test_Cov3_Within_StringRangeFloatDefault_Valid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloatDefault("50.5")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloatDefault returns value -- valid", actual)
}

func Test_Cov3_Within_StringRangeFloatDefault_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloatDefault("abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloatDefault returns false -- invalid", actual)
}

func Test_Cov3_Within_StringRangeFloat64_Valid(t *testing.T) {
	val, ok := corerange.Within.StringRangeFloat64(true, 1.0, 100.0, "50.5")
	actual := args.Map{"ok": ok, "inRange": val > 50.0}
	expected := args.Map{"ok": true, "inRange": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat64 returns value -- valid", actual)
}

func Test_Cov3_Within_StringRangeFloat64_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat64(true, 1.0, 100.0, "abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat64 returns false -- invalid", actual)
}

func Test_Cov3_Within_StringRangeFloat64Default_Valid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat64Default("50.5")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat64Default returns value -- valid", actual)
}

func Test_Cov3_Within_StringRangeFloat64Default_Invalid(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat64Default("abc")
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within StringRangeFloat64Default returns false -- invalid", actual)
}

func Test_Cov3_Within_RangeDefaultInteger(t *testing.T) {
	val, ok := corerange.Within.RangeDefaultInteger(1, 10, 5)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 5, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within RangeDefaultInteger returns value -- in range", actual)
}

func Test_Cov3_Within_RangeInteger_NoBoundary_OutOfRange(t *testing.T) {
	val, ok := corerange.Within.RangeInteger(false, 1, 10, 50)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 50, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeInteger returns raw -- no boundary", actual)
}

func Test_Cov3_Within_RangeInteger_BelowMin(t *testing.T) {
	val, ok := corerange.Within.RangeInteger(true, 10, 100, 5)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 10, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeInteger returns min -- below min", actual)
}

func Test_Cov3_Within_RangeInteger_AboveMax(t *testing.T) {
	val, ok := corerange.Within.RangeInteger(true, 10, 100, 200)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 100, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeInteger returns max -- above max", actual)
}

func Test_Cov3_Within_RangeByteDefault(t *testing.T) {
	val, ok := corerange.Within.RangeByteDefault(100)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 100, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within RangeByteDefault returns value -- in range", actual)
}

func Test_Cov3_Within_RangeByte_NoBoundary(t *testing.T) {
	_, ok := corerange.Within.RangeByte(false, -1)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeByte returns false -- no boundary", actual)
}

func Test_Cov3_Within_RangeByte_BelowZero(t *testing.T) {
	val, ok := corerange.Within.RangeByte(true, -1)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 0, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeByte returns 0 -- below zero", actual)
}

func Test_Cov3_Within_RangeByte_AboveMax(t *testing.T) {
	val, ok := corerange.Within.RangeByte(true, 300)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 255, "ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeByte returns 255 -- above max", actual)
}

func Test_Cov3_Within_RangeUint16Default(t *testing.T) {
	val, ok := corerange.Within.RangeUint16Default(1000)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 1000, "ok": true}
	expected.ShouldBeEqual(t, 0, "Within RangeUint16Default returns value -- in range", actual)
}

func Test_Cov3_Within_RangeUint16_NoBoundary(t *testing.T) {
	_, ok := corerange.Within.RangeUint16(false, -1)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeUint16 returns false -- no boundary", actual)
}

func Test_Cov3_Within_RangeFloat_InRange(t *testing.T) {
	val, ok := corerange.Within.RangeFloat(true, 1.0, 100.0, 50.0)
	actual := args.Map{"ok": ok, "val": val > 49.0}
	expected := args.Map{"ok": true, "val": true}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat returns value -- in range", actual)
}

func Test_Cov3_Within_RangeFloat_NoBoundary(t *testing.T) {
	_, ok := corerange.Within.RangeFloat(false, 1.0, 10.0, 50.0)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat returns raw -- no boundary", actual)
}

func Test_Cov3_Within_RangeFloat_BelowMin(t *testing.T) {
	val, ok := corerange.Within.RangeFloat(true, 10.0, 100.0, 5.0)
	actual := args.Map{"ok": ok, "isMin": val == 10.0}
	expected := args.Map{"ok": false, "isMin": true}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat returns min -- below min", actual)
}

func Test_Cov3_Within_RangeFloat_AboveMax(t *testing.T) {
	val, ok := corerange.Within.RangeFloat(true, 10.0, 100.0, 200.0)
	actual := args.Map{"ok": ok, "isMax": val == 100.0}
	expected := args.Map{"ok": false, "isMax": true}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat returns max -- above max", actual)
}

func Test_Cov3_Within_RangeFloat64_NoBoundary(t *testing.T) {
	_, ok := corerange.Within.RangeFloat64(false, 1.0, 10.0, 50.0)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat64 returns raw -- no boundary", actual)
}

func Test_Cov3_Within_RangeFloat64_BelowMin(t *testing.T) {
	val, ok := corerange.Within.RangeFloat64(true, 10.0, 100.0, 5.0)
	actual := args.Map{"ok": ok, "isMin": val == 10.0}
	expected := args.Map{"ok": false, "isMin": true}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat64 returns min -- below min", actual)
}

func Test_Cov3_Within_RangeFloat64_AboveMax(t *testing.T) {
	val, ok := corerange.Within.RangeFloat64(true, 10.0, 100.0, 200.0)
	actual := args.Map{"ok": ok, "isMax": val == 100.0}
	expected := args.Map{"ok": false, "isMax": true}
	expected.ShouldBeEqual(t, 0, "Within RangeFloat64 returns max -- above max", actual)
}

// ── MinMaxInt16 — uncovered branches ──

func Test_Cov3_MinMaxInt16_Comparisons(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
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
		"isMinEqual2": true, "isMinAboveEqual2": true, "isMinAbove1": true,
		"isMinLess3": true, "isMinLessEqual2": true, "isMaxEqual8": true,
		"isMaxAboveEqual8": true, "isMaxAbove7": true, "isMaxLess9": true,
		"isMaxLessEqual8": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 Comparisons returns true -- matching values", actual)
}

func Test_Cov3_MinMaxInt16_Ranges(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 3, Max: 7}
	actual := args.Map{
		"rangesLen":    len(mm.Ranges()),
		"rangesIntLen": len(mm.RangesInt()),
		"rangeLenInt":  mm.RangeLengthInt(),
	}
	expected := args.Map{"rangesLen": 5, "rangesIntLen": 5, "rangeLenInt": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 Ranges returns 5 -- range 3 to 7", actual)
}

func Test_Cov3_MinMaxInt16_CreateRanges(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 1, Max: 3}
	extra := corerange.MinMaxInt16{Min: 10, Max: 12}
	actual := args.Map{"combinedLen": len(mm.CreateRanges(extra)), "noExtraLen": len(mm.CreateRanges())}
	expected := args.Map{"combinedLen": 6, "noExtraLen": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 CreateRanges returns combined -- with extra", actual)
}

func Test_Cov3_MinMaxInt16_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 1, Max: 5}
	actual := args.Map{"exceptLen": len(mm.RangesExcept(3))}
	expected := args.Map{"exceptLen": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 RangesExcept returns 4 -- excluding 3", actual)
}

func Test_Cov3_MinMaxInt16_Clone(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 1, Max: 10}
	cloned := mm.Clone()
	clonedPtr := mm.ClonePtr()
	var nilMm *corerange.MinMaxInt16
	actual := args.Map{
		"clonedMin": int(cloned.Min), "clonedMax": int(cloned.Max),
		"clonedPtrNil": clonedPtr == nil, "nilCloneNil": nilMm.ClonePtr() == nil,
	}
	expected := args.Map{"clonedMin": 1, "clonedMax": 10, "clonedPtrNil": false, "nilCloneNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 Clone returns copy -- valid input", actual)
}

func Test_Cov3_MinMaxInt16_IsEqual(t *testing.T) {
	mm1 := &corerange.MinMaxInt16{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt16{Min: 1, Max: 10}
	mm3 := &corerange.MinMaxInt16{Min: 2, Max: 10}
	actual := args.Map{
		"sameValues": mm1.IsEqual(mm2), "diffValues": mm1.IsEqual(mm3),
		"samePtr": mm1.IsEqual(mm1), "bothNil": (*corerange.MinMaxInt16)(nil).IsEqual(nil),
		"leftNilOnly": (*corerange.MinMaxInt16)(nil).IsEqual(mm1),
	}
	expected := args.Map{
		"sameValues": true, "diffValues": false, "samePtr": true,
		"bothNil": true, "leftNilOnly": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 IsEqual returns expected -- various combos", actual)
}

func Test_Cov3_MinMaxInt16_String(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	actual := args.Map{"notEmpty": mm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 String returns non-empty -- valid input", actual)
}

func Test_Cov3_MinMaxInt16_CreateMinMaxInt(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	mmi := mm.CreateMinMaxInt()
	actual := args.Map{"notNil": mmi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 CreateMinMaxInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt16_CreateRangeInt(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt16_CreateRangeInt8(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt16_CreateRangeInt16(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16 CreateRangeInt16 returns non-nil -- valid", actual)
}

// ── MinMaxInt8 — uncovered branches ──

func Test_Cov3_MinMaxInt8_Comparisons(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
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
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 Comparisons returns true -- matching values", actual)
}

func Test_Cov3_MinMaxInt8_Ranges(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 3, Max: 7}
	actual := args.Map{
		"rangesLen": len(mm.Ranges()), "rangesIntLen": len(mm.RangesInt()),
		"rangeLenInt": mm.RangeLengthInt(),
	}
	expected := args.Map{"rangesLen": 5, "rangesIntLen": 5, "rangeLenInt": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 Ranges returns 5 -- range 3 to 7", actual)
}

func Test_Cov3_MinMaxInt8_CreateRanges(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 3}
	extra := corerange.MinMaxInt8{Min: 10, Max: 12}
	actual := args.Map{"combinedLen": len(mm.CreateRanges(extra)), "noExtraLen": len(mm.CreateRanges())}
	expected := args.Map{"combinedLen": 6, "noExtraLen": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRanges returns combined -- with extra", actual)
}

func Test_Cov3_MinMaxInt8_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 5}
	actual := args.Map{"exceptLen": len(mm.RangesExcept(3))}
	expected := args.Map{"exceptLen": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 RangesExcept returns 4 -- excluding 3", actual)
}

func Test_Cov3_MinMaxInt8_Clone(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 1, Max: 10}
	cloned := mm.Clone()
	var nilMm *corerange.MinMaxInt8
	actual := args.Map{
		"clonedMin": int(cloned.Min), "clonedMax": int(cloned.Max),
		"clonedPtrNil": mm.ClonePtr() == nil, "nilCloneNil": nilMm.ClonePtr() == nil,
	}
	expected := args.Map{"clonedMin": 1, "clonedMax": 10, "clonedPtrNil": false, "nilCloneNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 Clone returns copy -- valid input", actual)
}

func Test_Cov3_MinMaxInt8_IsEqual(t *testing.T) {
	mm1 := &corerange.MinMaxInt8{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt8{Min: 1, Max: 10}
	actual := args.Map{"same": mm1.IsEqual(mm2), "samePtr": mm1.IsEqual(mm1),
		"bothNil": (*corerange.MinMaxInt8)(nil).IsEqual(nil)}
	expected := args.Map{"same": true, "samePtr": true, "bothNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 IsEqual returns true -- same values", actual)
}

func Test_Cov3_MinMaxInt8_String(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notEmpty": mm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 String returns non-empty -- valid input", actual)
}

func Test_Cov3_MinMaxInt8_CreateMinMaxInt(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateMinMaxInt() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateMinMaxInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt8_CreateRangeInt(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt8_CreateRangeInt8(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt8_CreateRangeInt16(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8 CreateRangeInt16 returns non-nil -- valid", actual)
}

// ── MinMaxInt64 — uncovered branches ──

func Test_Cov3_MinMaxInt64_Comparisons(t *testing.T) {
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

func Test_Cov3_MinMaxInt64_Ranges(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 3, Max: 7}
	actual := args.Map{
		"rangesLen": len(mm.Ranges()), "rangesIntLen": len(mm.RangesInt()),
		"rangeLenInt": mm.RangeLengthInt(),
	}
	expected := args.Map{"rangesLen": 5, "rangesIntLen": 5, "rangeLenInt": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 Ranges returns 5 -- range 3 to 7", actual)
}

func Test_Cov3_MinMaxInt64_CreateRanges(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 3}
	extra := corerange.MinMaxInt64{Min: 10, Max: 12}
	actual := args.Map{"combinedLen": len(mm.CreateRanges(extra)), "noExtraLen": len(mm.CreateRanges())}
	expected := args.Map{"combinedLen": 6, "noExtraLen": 3}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRanges returns combined -- with extra", actual)
}

func Test_Cov3_MinMaxInt64_RangesExcept(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 5}
	actual := args.Map{"exceptLen": len(mm.RangesExcept(3))}
	expected := args.Map{"exceptLen": 4}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 RangesExcept returns 4 -- excluding 3", actual)
}

func Test_Cov3_MinMaxInt64_Clone(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 1, Max: 10}
	cloned := mm.Clone()
	var nilMm *corerange.MinMaxInt64
	actual := args.Map{
		"clonedMin": int(cloned.Min), "clonedMax": int(cloned.Max),
		"clonedPtrNil": mm.ClonePtr() == nil, "nilCloneNil": nilMm.ClonePtr() == nil,
	}
	expected := args.Map{"clonedMin": 1, "clonedMax": 10, "clonedPtrNil": false, "nilCloneNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 Clone returns copy -- valid input", actual)
}

func Test_Cov3_MinMaxInt64_IsEqual(t *testing.T) {
	mm1 := &corerange.MinMaxInt64{Min: 1, Max: 10}
	mm2 := &corerange.MinMaxInt64{Min: 1, Max: 10}
	actual := args.Map{"same": mm1.IsEqual(mm2), "samePtr": mm1.IsEqual(mm1),
		"bothNil": (*corerange.MinMaxInt64)(nil).IsEqual(nil)}
	expected := args.Map{"same": true, "samePtr": true, "bothNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 IsEqual returns true -- same values", actual)
}

func Test_Cov3_MinMaxInt64_String(t *testing.T) {
	actual := args.Map{"notEmpty": (&corerange.MinMaxInt64{Min: 2, Max: 8}).String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 String returns non-empty -- valid", actual)
}

func Test_Cov3_MinMaxInt64_CreateMinMaxInt(t *testing.T) {
	actual := args.Map{"notNil": (&corerange.MinMaxInt64{Min: 2, Max: 8}).CreateMinMaxInt() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateMinMaxInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt64_CreateRangeInt(t *testing.T) {
	actual := args.Map{"notNil": (&corerange.MinMaxInt64{Min: 2, Max: 8}).CreateRangeInt("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt64_CreateRangeInt8(t *testing.T) {
	actual := args.Map{"notNil": (&corerange.MinMaxInt64{Min: 2, Max: 8}).CreateRangeInt8("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt64_CreateRangeInt16(t *testing.T) {
	actual := args.Map{"notNil": (&corerange.MinMaxInt64{Min: 2, Max: 8}).CreateRangeInt16("3:5", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64 CreateRangeInt16 returns non-nil -- valid", actual)
}

// ── StartEndSimpleString — uncovered branches ──

func Test_Cov3_StartEndSimpleString_NilReceiver(t *testing.T) {
	var se *corerange.StartEndSimpleString
	actual := args.Map{
		"isInvalidStart":    se.IsInvalidStart(),
		"isInvalidEnd":      se.IsInvalidEnd(),
		"hasStart":          se.HasStart(),
		"hasEnd":            se.HasEnd(),
		"isInvalidBoth":     se.IsInvalidStartEndBoth(),
		"isInvalidAny":      se.IsInvalidAnyStartEnd(),
		"startValidVal":     se.StartValidValue() == nil,
		"endValidVal":       se.EndValidValue() == nil,
		"startEndStr":       se.StartEndString() == nil,
		"isStartEndBoth":    se.IsStartEndBothDefined(),
	}
	expected := args.Map{
		"isInvalidStart": true, "isInvalidEnd": true,
		"hasStart": false, "hasEnd": false,
		"isInvalidBoth": true, "isInvalidAny": true,
		"startValidVal": true, "endValidVal": true,
		"startEndStr": true, "isStartEndBoth": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString nil returns safe defaults -- nil receiver", actual)
}

func Test_Cov3_StartEndSimpleString_AllMethods(t *testing.T) {
	se := &corerange.StartEndSimpleString{Start: "hello", End: "world"}
	actual := args.Map{
		"isInvalidStart": se.IsInvalidStart(), "isInvalidEnd": se.IsInvalidEnd(),
		"isStartEndBoth": se.IsStartEndBothDefined(), "isInvalidBoth": se.IsInvalidStartEndBoth(),
		"isInvalidAny":   se.IsInvalidAnyStartEnd(),
		"startValid":     se.StartValidValue() != nil, "endValid": se.EndValidValue() != nil,
		"startEndStr":    se.StartEndString() != nil,
		"stringSpace":    se.StringSpace(), "stringHyphen": se.StringHyphen(),
		"stringColon":    se.StringColon(),
	}
	expected := args.Map{
		"isInvalidStart": false, "isInvalidEnd": false,
		"isStartEndBoth": true, "isInvalidBoth": false,
		"isInvalidAny":   false,
		"startValid": true, "endValid": true,
		"startEndStr": true,
		"stringSpace": "hello world", "stringHyphen": "hello-world",
		"stringColon": "hello:world",
	}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString methods return expected -- valid input", actual)
}

func Test_Cov3_StartEndSimpleString_StringUsingFormat(t *testing.T) {
	se := &corerange.StartEndSimpleString{Start: "a", End: "b"}
	actual := args.Map{"result": se.StringUsingFormat("%s-%s")}
	expected := args.Map{"result": "a-b"}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString StringUsingFormat returns formatted -- custom format", actual)
}

func Test_Cov3_StartEndSimpleString_RangeInt(t *testing.T) {
	se := &corerange.StartEndSimpleString{Start: "3", End: "7"}
	actual := args.Map{"notNil": se.RangeInt(nil) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString RangeInt returns non-nil -- nil minMax", actual)
}

// ── StartEndInt — uncovered branches ──

func Test_Cov3_StartEndInt_NilReceiver(t *testing.T) {
	var se *corerange.StartEndInt
	actual := args.Map{
		"isInvalid":       se.IsInvalid(),
		"isInvalidStart":  se.IsInvalidStart(),
		"isInvalidEnd":    se.IsInvalidEnd(),
		"hasStart":        se.HasStart(),
		"hasEnd":          se.HasEnd(),
		"isStartEndBoth":  se.IsStartEndBothDefined(),
		"isStartGrater":   se.IsStartGraterThan(5),
		"isEndGrater":     se.IsEndGraterThan(5),
	}
	expected := args.Map{
		"isInvalid": true, "isInvalidStart": true, "isInvalidEnd": true,
		"hasStart": false, "hasEnd": false, "isStartEndBoth": false,
		"isStartGrater": false, "isEndGrater": false,
	}
	expected.ShouldBeEqual(t, 0, "StartEndInt nil returns safe defaults -- nil receiver", actual)
}

func Test_Cov3_StartEndInt_StringUsingFormat(t *testing.T) {
	se := &corerange.StartEndInt{Start: 3, End: 7}
	actual := args.Map{"result": se.StringUsingFormat("%d-%d")}
	expected := args.Map{"result": "3-7"}
	expected.ShouldBeEqual(t, 0, "StartEndInt StringUsingFormat returns formatted -- custom", actual)
}

func Test_Cov3_StartEndInt_RangeInt(t *testing.T) {
	se := &corerange.StartEndInt{Start: 3, End: 7}
	actual := args.Map{"notNil": se.RangeInt(nil) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndInt RangeInt returns non-nil -- nil minMax", actual)
}

func Test_Cov3_StartEndInt_DiffNegative(t *testing.T) {
	se := &corerange.StartEndInt{Start: 7, End: 3}
	actual := args.Map{"diff": se.Diff(), "diffAbs": se.DifferenceAbsolute()}
	expected := args.Map{"diff": -4, "diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "StartEndInt Diff returns negative -- reversed range", actual)
}

// ── RangeInt — uncovered constructor branches ──

func Test_Cov3_RangeInt_NewRangeInt8MinMax(t *testing.T) {
	ri8 := corerange.NewRangeInt8MinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": ri8 != nil, "start": int(ri8.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt8MinMax returns valid -- valid input", actual)
}

func Test_Cov3_RangeInt_NewRangeInt16MinMax(t *testing.T) {
	ri16 := corerange.NewRangeInt16MinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": ri16 != nil, "start": int(ri16.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt16MinMax returns valid -- valid input", actual)
}

func Test_Cov3_RangeInt_NewRangeByteMinMax(t *testing.T) {
	rb := corerange.NewRangeByteMinMax("3:7", ":", 0, 10)
	actual := args.Map{"notNil": rb != nil, "start": int(rb.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeByteMinMax returns valid -- valid input", actual)
}

func Test_Cov3_RangeInt_NewRangeByte_NilMinMax(t *testing.T) {
	rb := corerange.NewRangeByte("3:7", ":", nil)
	actual := args.Map{"notNil": rb != nil, "start": int(rb.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeByte returns valid -- nil minMax", actual)
}

func Test_Cov3_RangeInt_NewRangeInt8_NilMinMax(t *testing.T) {
	ri8 := corerange.NewRangeInt8("3:7", ":", nil)
	actual := args.Map{"notNil": ri8 != nil, "start": int(ri8.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt8 returns valid -- nil minMax", actual)
}

func Test_Cov3_RangeInt_NewRangeInt16_NilMinMax(t *testing.T) {
	ri16 := corerange.NewRangeInt16("3:7", ":", nil)
	actual := args.Map{"notNil": ri16 != nil, "start": int(ri16.Start)}
	expected := args.Map{"notNil": true, "start": 3}
	expected.ShouldBeEqual(t, 0, "NewRangeInt16 returns valid -- nil minMax", actual)
}

// ── RangeInt8 — uncovered branches ──

func Test_Cov3_RangeInt8_ValidPlusWithinRange(t *testing.T) {
	ri8 := corerange.NewRangeInt8MinMax("3:7", ":", 0, 10)
	actual := args.Map{
		"validWithin":   ri8.IsValidPlusWithinRange(5),
		"validOutside":  ri8.IsValidPlusWithinRange(10),
		"rangesInt8Len": len(ri8.RangesInt8()),
	}
	expected := args.Map{"validWithin": true, "validOutside": false, "rangesInt8Len": 5}
	expected.ShouldBeEqual(t, 0, "RangeInt8 IsValidPlusWithinRange returns expected -- valid range", actual)
}

// ── RangeInt16 — uncovered branches ──

func Test_Cov3_RangeInt16_ValidPlusWithinRange(t *testing.T) {
	ri16 := corerange.NewRangeInt16MinMax("3:7", ":", 0, 10)
	actual := args.Map{
		"validWithin":    ri16.IsValidPlusWithinRange(5),
		"validOutside":   ri16.IsValidPlusWithinRange(10),
		"rangesInt16Len": len(ri16.RangesInt16()),
	}
	expected := args.Map{"validWithin": true, "validOutside": false, "rangesInt16Len": 5}
	expected.ShouldBeEqual(t, 0, "RangeInt16 IsValidPlusWithinRange returns expected -- valid range", actual)
}

// ── RangeByte — uncovered branches ──

func Test_Cov3_RangeByte_ValidPlusWithinRange(t *testing.T) {
	rb := corerange.NewRangeByteMinMax("3:7", ":", 0, 10)
	actual := args.Map{
		"validWithin":  rb.IsValidPlusWithinRange(5),
		"validOutside": rb.IsValidPlusWithinRange(10),
		"rangesIntLen": len(rb.RangesInt()),
	}
	expected := args.Map{"validWithin": true, "validOutside": false, "rangesIntLen": 5}
	expected.ShouldBeEqual(t, 0, "RangeByte IsValidPlusWithinRange returns expected -- valid range", actual)
}

func Test_Cov3_RangeByte_DifferenceStartGtEnd(t *testing.T) {
	rb := &corerange.RangeByte{
		BaseRange: &corerange.BaseRange{IsValid: true},
		Start:     7,
		End:       3,
	}
	actual := args.Map{"diff": int(rb.Difference())}
	expected := args.Map{"diff": 4}
	expected.ShouldBeEqual(t, 0, "RangeByte Difference returns positive -- start > end", actual)
}

// ── RangeAny — uncovered branches ──

func Test_Cov3_RangeAny_Methods(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{
			RawInput:  "hello:world",
			Separator: ":",
			IsValid:   true,
			HasStart:  true,
			HasEnd:    true,
		},
		RawInput: "hello:world",
		Start:    "hello",
		End:      "world",
	}
	actual := args.Map{
		"rawInput":       ra.RawInputString(),
		"start":          ra.StartString(),
		"end":            ra.EndString(),
		"rangeIntNotNil": ra.CreateRangeInt() != nil,
		"rangeStrNotNil": ra.CreateRangeString() != nil,
		"startEndNotNil": ra.CreateStartEndString() != nil,
		"stringNotEmpty": ra.String() != "",
	}
	expected := args.Map{
		"rawInput": "hello:world", "start": "hello", "end": "world",
		"rangeIntNotNil": true, "rangeStrNotNil": true, "startEndNotNil": true,
		"stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeAny methods return expected -- valid input", actual)
}

func Test_Cov3_RangeAny_CreateRangeIntMinMax(t *testing.T) {
	ra := &corerange.RangeAny{
		BaseRange: &corerange.BaseRange{
			RawInput:  "3:7",
			Separator: ":",
			IsValid:   true,
			HasStart:  true,
			HasEnd:    true,
		},
		RawInput: "3:7",
		Start:    "3",
		End:      "7",
	}
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}
	actual := args.Map{"notNil": ra.CreateRangeIntMinMax(mm) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RangeAny CreateRangeIntMinMax returns non-nil -- valid", actual)
}

// ── StartEndString — uncovered branches ──

func Test_Cov3_StartEndString_UsingLines(t *testing.T) {
	se := corerange.NewStartEndStringUsingLines([]string{"first", "middle", "last"})
	actual := args.Map{"start": se.Start, "end": se.End, "hasStart": se.HasStart, "hasEnd": se.HasEnd}
	expected := args.Map{"start": "first", "end": "last", "hasStart": true, "hasEnd": true}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns first/last -- 3 lines", actual)
}

func Test_Cov3_StartEndString_UsingLines_SingleElement(t *testing.T) {
	se := corerange.NewStartEndStringUsingLines([]string{"only"})
	actual := args.Map{"start": se.Start, "hasStart": se.HasStart, "hasEnd": se.HasEnd}
	expected := args.Map{"start": "only", "hasStart": true, "hasEnd": false}
	expected.ShouldBeEqual(t, 0, "NewStartEndStringUsingLines returns first -- single element", actual)
}

func Test_Cov3_StartEndString_CreateRangeString(t *testing.T) {
	se := corerange.NewStartEndString("a:b", ":")
	actual := args.Map{"notNil": se.CreateRangeString() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StartEndString CreateRangeString returns non-nil -- valid", actual)
}

// ── MinMaxInt — uncovered branches ──

func Test_Cov3_MinMaxInt_DifferenceAbsolute_Negative(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 8, Max: 3}
	actual := args.Map{"diffAbs": mm.DifferenceAbsolute()}
	expected := args.Map{"diffAbs": 5}
	expected.ShouldBeEqual(t, 0, "MinMaxInt DifferenceAbsolute returns positive -- negative diff", actual)
}

func Test_Cov3_MinMaxInt_CreateRangeInt8(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt8("3:7", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt CreateRangeInt8 returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt_CreateRangeInt16(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}
	actual := args.Map{"notNil": mm.CreateRangeInt16("3:7", ":") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MinMaxInt CreateRangeInt16 returns non-nil -- valid", actual)
}

func Test_Cov3_MinMaxInt_IsInvalidValue(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 3, Max: 7}
	actual := args.Map{"invalid2": mm.IsInvalidValue(2), "valid5": mm.IsInvalidValue(5)}
	expected := args.Map{"invalid2": true, "valid5": false}
	expected.ShouldBeEqual(t, 0, "MinMaxInt IsInvalidValue returns expected -- boundary check", actual)
}

// ── RangeInt — DifferenceAbsolute ──

func Test_Cov3_RangeInt_DifferenceAbsolute(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(3, 7, true)
	actual := args.Map{"diffAbs": ri.DifferenceAbsolute()}
	expected := args.Map{"diffAbs": 4}
	expected.ShouldBeEqual(t, 0, "RangeInt DifferenceAbsolute returns 4 -- 3 to 7", actual)
}

// ── RangeInt — Invalid Ranges returns empty ──

func Test_Cov3_RangeInt_InvalidRanges(t *testing.T) {
	ri := corerange.NewRangeIntUsingValues(7, 3, false)
	actual := args.Map{"rangesLen": len(ri.Ranges())}
	expected := args.Map{"rangesLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeInt Ranges returns empty -- invalid range", actual)
}
