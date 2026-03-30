package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage16 — enumimpl remaining gaps (external tests)
//
// Covers accessible gaps:
// - ConvEnumAnyValToInteger type switch branches
// - DynamicMap.Set/AddNewOnly nil receiver
// - DynamicMap.isEqualSingle isRegardlessType branch
// - DiffLeftRight.JsonString json.Marshal error branch (dead code)
// - BasicByte/Int8/Int16/Int32/UInt16/String GetValueByName wrapped name branch
//
// Unexported function gaps (toHashset, toStringPrintableDynamicMap, etc.)
// require internal tests or are documented as dead code.
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

func Test_Cov16_ConvEnumAnyValToInteger_ValueByter(t *testing.T) {
	// Arrange & Act
	result := enumimpl.ConvEnumAnyValToInteger(mockValueByter{val: 42})

	// Assert
	convey.Convey("ConvEnumAnyValToInteger handles valueByter", t, func() {
		convey.So(result, convey.ShouldEqual, 42)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ExactValueByter(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueByter{val: 7})
	convey.Convey("ConvEnumAnyValToInteger handles exactValueByter", t, func() {
		convey.So(result, convey.ShouldEqual, 7)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ValueInter(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockValueInter{val: 100})
	convey.Convey("ConvEnumAnyValToInteger handles valueInter", t, func() {
		convey.So(result, convey.ShouldEqual, 100)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ExactValueInter(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueInter{val: 200})
	convey.Convey("ConvEnumAnyValToInteger handles exactValueInter", t, func() {
		convey.So(result, convey.ShouldEqual, 200)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ValueInt8er(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockValueInt8er{val: 5})
	convey.Convey("ConvEnumAnyValToInteger handles valueInt8er", t, func() {
		convey.So(result, convey.ShouldEqual, 5)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ExactValueInt8er(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueInt8er{val: 3})
	convey.Convey("ConvEnumAnyValToInteger handles exactValueInt8er", t, func() {
		convey.So(result, convey.ShouldEqual, 3)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ValueUInt16er(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockValueUInt16er{val: 300})
	convey.Convey("ConvEnumAnyValToInteger handles valueUInt16er", t, func() {
		convey.So(result, convey.ShouldEqual, 300)
	})
}

func Test_Cov16_ConvEnumAnyValToInteger_ExactValueUInt16er(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(mockExactValueUInt16er{val: 500})
	convey.Convey("ConvEnumAnyValToInteger handles exactValueUInt16er", t, func() {
		convey.So(result, convey.ShouldEqual, 500)
	})
}

// --- DynamicMap nil receiver ---

func Test_Cov16_DynamicMap_Set_NilReceiver(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap(nil)

	// Act
	isNew := dm.Set("key", "val")

	// Assert
	convey.Convey("DynamicMap.Set on nil initializes and adds", t, func() {
		convey.So(isNew, convey.ShouldBeTrue)
	})
}

func Test_Cov16_DynamicMap_AddNewOnly_NilReceiver(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap(nil)

	// Act
	isAdded := dm.AddNewOnly("key", "val")

	// Assert
	convey.Convey("DynamicMap.AddNewOnly on nil initializes and adds", t, func() {
		convey.So(isAdded, convey.ShouldBeTrue)
	})
}

// --- DynamicMap.IsEqualRegardlessTypes ---

func Test_Cov16_DynamicMap_IsEqualRegardlessTypes(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"key": 42}
	dm2 := enumimpl.DynamicMap{"key": "42"}

	// Act
	result := dm1.IsEqualRegardlessTypes(dm2)

	// Assert
	convey.Convey("DynamicMap.IsEqualRegardlessTypes compares by string representation", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// --- DiffLeftRight.JsonString error branch (dead code) ---

func Test_Cov16_DiffLeftRight_JsonString(t *testing.T) {
	// Arrange
	diff := &enumimpl.DiffLeftRight{}

	// Act
	result := diff.JsonString()

	// Assert
	convey.Convey("DiffLeftRight.JsonString returns valid JSON for empty struct", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}

// Coverage note: Remaining enumimpl gaps:
// - BasicByte/Int8/Int16/Int32/UInt16/String GetValueByName wrapped branch:
//   requires constructing BasicXxx with specific internal hashmaps — only via New.BasicXxx
// - DynamicMap unexported methods (isEqualSingle, etc.) need internal tests
// - toHashset, toStringPrintableDynamicMap, numberEnumBase: unexported, need internal tests
// - newBasicStringCreator: requires specific Stringer types
// - DynamicMap byte/int getter branches: require specific value types in map
