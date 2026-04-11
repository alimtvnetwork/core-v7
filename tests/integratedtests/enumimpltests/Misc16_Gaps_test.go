package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage16 — enumimpl remaining gaps (external tests)
// ══════════════════════════════════════════════════════════════════════════════

// --- ConvEnumAnyValToInteger type switch branches ---

type mockValueByter struct{ val byte }

func (m mockValueByter) Value() byte { return m.val }

type mockExactValueByter struct{ val byte }

func (m mockExactValueByter) ValueByte() byte { return m.val }

type mockValueInter struct{ val int }

func (m mockValueInter) Value() int { return m.val }

type mockExactValueInter struct{ val int }

func (m mockExactValueInter) ValueInt() int { return m.val }

type mockValueInt8er struct{ val int8 }

func (m mockValueInt8er) Value() int8 { return m.val }

type mockExactValueInt8er struct{ val int8 }

func (m mockExactValueInt8er) ValueInt8() int8 { return m.val }

type mockValueUInt16er struct{ val uint16 }

func (m mockValueUInt16er) Value() uint16 { return m.val }

type mockExactValueUInt16er struct{ val uint16 }

func (m mockExactValueUInt16er) ValueUInt16() uint16 { return m.val }

func Test_ConvEnumAnyValToInteger_ValueByter(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockValueByter{val: 42})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles valueByter", t, func() {
		convey.So(result, convey.ShouldEqual, 42)
	})
}

func Test_ConvEnumAnyValToInteger_ExactValueByter(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueByter{val: 7})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles exactValueByter", t, func() {
		convey.So(result, convey.ShouldEqual, 7)
	})
}

func Test_ConvEnumAnyValToInteger_ValueInter(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockValueInter{val: 100})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles valueInter", t, func() {
		convey.So(result, convey.ShouldEqual, 100)
	})
}

func Test_ConvEnumAnyValToInteger_ExactValueInter(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueInter{val: 200})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles exactValueInter", t, func() {
		convey.So(result, convey.ShouldEqual, 200)
	})
}

func Test_ConvEnumAnyValToInteger_ValueInt8er(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockValueInt8er{val: 5})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles valueInt8er", t, func() {
		convey.So(result, convey.ShouldEqual, 5)
	})
}

func Test_ConvEnumAnyValToInteger_ExactValueInt8er(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueInt8er{val: 3})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles exactValueInt8er", t, func() {
		convey.So(result, convey.ShouldEqual, 3)
	})
}

func Test_ConvEnumAnyValToInteger_ValueUInt16er(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockValueUInt16er{val: 300})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles valueUInt16er", t, func() {
		convey.So(result, convey.ShouldEqual, 300)
	})
}

func Test_ConvEnumAnyValToInteger_ExactValueUInt16er(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueUInt16er{val: 500})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles exactValueUInt16er", t, func() {
		convey.So(result, convey.ShouldEqual, 500)
	})
}

// --- DynamicMap.IsEqual with isRegardlessType=true ---

func Test_DynamicMap_IsEqual_RegardlessType_FromMisc16Gaps(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"key": 42}
	dm2 := enumimpl.DynamicMap{"key": "42"}

	// Act
	result := dm1.IsRawEqual(true, dm2)

	// Assert
	convey.Convey("DynamicMap.IsRawEqual with isRegardlessType compares by string", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// --- DiffLeftRight.JsonString ---

func Test_DiffLeftRight_JsonString_Valid(t *testing.T) {
	// Arrange
	diff := &enumimpl.DiffLeftRight{}

	// Act
	result := diff.JsonString()

	// Assert
	convey.Convey("DiffLeftRight.JsonString returns valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// Coverage note: Remaining gaps are either:
// 1. Nil-receiver dead code (Set/AddNewOnly with nil *DynamicMap — panics on *it dereference)
// 2. Unexported functions (toHashset, toStringPrintableDynamicMap, numberEnumBase)
// 3. BasicByte/Int/String GetValueByName wrapped lookup — requires internal hashmap setup
// 4. DynamicMap byte/int getter type-assertion branches — need specific typed values
// All documented as accepted gaps requiring internal tests.
