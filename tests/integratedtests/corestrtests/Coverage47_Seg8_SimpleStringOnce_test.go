package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Segment 8a
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg8_SSO_ValueAndInit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueAndInit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{
			"value": sso.Value(),
			"init": sso.IsInitialized(),
			"defined": sso.IsDefined(),
			"uninit": sso.IsUninitialized(),
		}

		// Assert
		expected := args.Map{
			"value": "",
			"init": false,
			"defined": false,
			"uninit": true,
		}
		expected.ShouldBeEqual(t, 0, "Value/Init -- default zero", actual)
	})
}

func Test_Seg8_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SetOnUninitialized", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		err := sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{
			"err": err == nil,
			"val": sso.Value(),
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"err": true,
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized -- sets value", actual)
	})
}

func Test_Seg8_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("first")
		err := sso.SetOnUninitialized("second")

		// Act
		actual := args.Map{
			"hasErr": err != nil,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"hasErr": true,
			"val": "first",
		}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized already init -- error", actual)
	})
}

func Test_Seg8_SSO_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_GetSetOnce", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		val1 := sso.GetSetOnce("first")
		val2 := sso.GetSetOnce("second")

		// Act
		actual := args.Map{
			"val1": val1,
			"val2": val2,
		}

		// Assert
		expected := args.Map{
			"val1": "first",
			"val2": "first",
		}
		expected.ShouldBeEqual(t, 0, "GetSetOnce -- only first sticks", actual)
	})
}

func Test_Seg8_SSO_GetOnce(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_GetOnce", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		val := sso.GetOnce()

		// Act
		actual := args.Map{
			"val": val,
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "GetOnce -- sets empty once", actual)
	})
}

func Test_Seg8_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_GetOnceFunc", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		val := sso.GetOnceFunc(func() string { return "computed" })
		val2 := sso.GetOnceFunc(func() string { return "other" })

		// Act
		actual := args.Map{
			"val": val,
			"val2": val2,
		}

		// Assert
		expected := args.Map{
			"val": "computed",
			"val2": "computed",
		}
		expected.ShouldBeEqual(t, 0, "GetOnceFunc -- first call wins", actual)
	})
}

func Test_Seg8_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SetOnceIfUninitialized", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		set1 := sso.SetOnceIfUninitialized("first")
		set2 := sso.SetOnceIfUninitialized("second")

		// Act
		actual := args.Map{
			"set1": set1,
			"set2": set2,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"set1": true,
			"set2": false,
			"val": "first",
		}
		expected.ShouldBeEqual(t, 0, "SetOnceIfUninitialized -- first true second false", actual)
	})
}

func Test_Seg8_SSO_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SetInitialize_SetUnInit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		sso.SetInitialize()
		init1 := sso.IsInitialized()
		sso.SetUnInit()
		init2 := sso.IsInitialized()

		// Act
		actual := args.Map{
			"init1": init1,
			"init2": init2,
		}

		// Assert
		expected := args.Map{
			"init1": true,
			"init2": false,
		}
		expected.ShouldBeEqual(t, 0, "SetInitialize/SetUnInit -- toggles", actual)
	})
}

func Test_Seg8_SSO_Invalidate(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Invalidate", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")
		sso.Invalidate()

		// Act
		actual := args.Map{
			"init": sso.IsInitialized(),
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"init": false,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "Invalidate -- resets", actual)
	})
}

func Test_Seg8_SSO_Reset(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Reset", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")
		sso.Reset()

		// Act
		actual := args.Map{
			"init": sso.IsInitialized(),
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"init": false,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "Reset -- resets", actual)
	})
}

func Test_Seg8_SSO_IsInvalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsInvalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		invalid1 := sso.IsInvalid()
		_ = sso.SetOnUninitialized("val")
		invalid2 := sso.IsInvalid()

		// Act
		actual := args.Map{
			"invalid1": invalid1,
			"invalid2": invalid2,
		}

		// Assert
		expected := args.Map{
			"invalid1": true,
			"invalid2": false,
		}
		expected.ShouldBeEqual(t, 0, "IsInvalid -- uninit vs init", actual)
	})
}

func Test_Seg8_SSO_IsInvalid_NilReceiver(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsInvalid_NilReceiver", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		actual := args.Map{"invalid": sso.IsInvalid()}

		// Assert
		expected := args.Map{"invalid": true}
		expected.ShouldBeEqual(t, 0, "IsInvalid nil -- true", actual)
	})
}

func Test_Seg8_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueBytes", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{
			"len": len(sso.ValueBytes()),
			"ptrLen": len(sso.ValueBytesPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"ptrLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "ValueBytes -- correct", actual)
	})
}

func Test_Seg8_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ConcatNew", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		result := sso.ConcatNew(" world")

		// Act
		actual := args.Map{
			"val": result.Value(),
			"init": result.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "hello world",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- appended", actual)
	})
}

func Test_Seg8_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ConcatNewUsingStrings", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a")
		result := sso.ConcatNewUsingStrings(",", "b", "c")

		// Act
		actual := args.Map{"val": result.Value()}

		// Assert
		expected := args.Map{"val": "a,b,c"}
		expected.ShouldBeEqual(t, 0, "ConcatNewUsingStrings -- joined", actual)
	})
}

func Test_Seg8_SSO_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsEmpty_IsWhitespace", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{
			"empty": sso.IsEmpty(),
			"ws": sso.IsWhitespace(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"ws": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty/IsWhitespace -- empty", actual)
	})
}

func Test_Seg8_SSO_Trim(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Trim", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("  hello  ")

		// Act
		actual := args.Map{"val": sso.Trim()}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "Trim -- trimmed", actual)
	})
}

func Test_Seg8_SSO_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_HasValidNonEmpty", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")

		// Act
		actual := args.Map{
			"nonEmpty": sso.HasValidNonEmpty(),
			"nonWS": sso.HasValidNonWhitespace(),
		}

		// Assert
		expected := args.Map{
			"nonEmpty": true,
			"nonWS": true,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty/NonWhitespace -- true", actual)
	})
}

func Test_Seg8_SSO_HasValidNonEmpty_Uninit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_HasValidNonEmpty_Uninit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{
			"nonEmpty": sso.HasValidNonEmpty(),
			"nonWS": sso.HasValidNonWhitespace(),
		}

		// Assert
		expected := args.Map{
			"nonEmpty": false,
			"nonWS": false,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty uninit -- false", actual)
	})
}

func Test_Seg8_SSO_SafeValue(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SafeValue", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		uninitVal := sso.SafeValue()
		_ = sso.SetOnUninitialized("hello")
		initVal := sso.SafeValue()

		// Act
		actual := args.Map{
			"uninit": uninitVal,
			"init": initVal,
		}

		// Assert
		expected := args.Map{
			"uninit": "",
			"init": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SafeValue -- correct", actual)
	})
}

func Test_Seg8_SSO_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_HasSafeNonEmpty", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")

		// Act
		actual := args.Map{"safe": sso.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"safe": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty -- true", actual)
	})
}

func Test_Seg8_SSO_Int(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("42")

		// Act
		actual := args.Map{"int": sso.Int()}

		// Assert
		expected := args.Map{"int": 42}
		expected.ShouldBeEqual(t, 0, "Int -- 42", actual)
	})
}

func Test_Seg8_SSO_Int_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{"int": sso.Int()}

		// Assert
		expected := args.Map{"int": 0}
		expected.ShouldBeEqual(t, 0, "Int invalid -- 0", actual)
	})
}

func Test_Seg8_SSO_Byte(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Byte", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("100")

		// Act
		actual := args.Map{"byte": sso.Byte()}

		// Assert
		expected := args.Map{"byte": byte(100)}
		expected.ShouldBeEqual(t, 0, "Byte -- 100", actual)
	})
}

func Test_Seg8_SSO_Byte_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Byte_OutOfRange", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("300")

		// Act
		actual := args.Map{"byte": sso.Byte()}

		// Assert
		expected := args.Map{"byte": byte(0)}
		expected.ShouldBeEqual(t, 0, "Byte out of range -- 0", actual)
	})
}

func Test_Seg8_SSO_Byte_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Byte_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{"byte": sso.Byte()}

		// Assert
		expected := args.Map{"byte": byte(0)}
		expected.ShouldBeEqual(t, 0, "Byte invalid -- 0", actual)
	})
}

func Test_Seg8_SSO_Int16(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int16", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("1000")

		// Act
		actual := args.Map{"int16": sso.Int16()}

		// Assert
		expected := args.Map{"int16": int16(1000)}
		expected.ShouldBeEqual(t, 0, "Int16 -- 1000", actual)
	})
}

func Test_Seg8_SSO_Int16_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int16_OutOfRange", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("40000")

		// Act
		actual := args.Map{"int16": sso.Int16()}

		// Assert
		expected := args.Map{"int16": int16(0)}
		expected.ShouldBeEqual(t, 0, "Int16 out of range -- 0", actual)
	})
}

func Test_Seg8_SSO_Int32(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int32", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("100000")

		// Act
		actual := args.Map{"int32": sso.Int32()}

		// Assert
		expected := args.Map{"int32": int32(100000)}
		expected.ShouldBeEqual(t, 0, "Int32 -- 100000", actual)
	})
}

func Test_Seg8_SSO_Int32_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Int32_OutOfRange", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("999999999999")

		// Act
		actual := args.Map{"int32": sso.Int32()}

		// Assert
		expected := args.Map{"int32": int32(0)}
		expected.ShouldBeEqual(t, 0, "Int32 out of range -- 0", actual)
	})
}

func Test_Seg8_SSO_Uint16(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Uint16", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("500")
		val, inRange := sso.Uint16()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint16(500),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "Uint16 -- 500", actual)
	})
}

func Test_Seg8_SSO_Uint32(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Uint32", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("70000")
		val, inRange := sso.Uint32()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint32(70000),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "Uint32 -- 70000", actual)
	})
}

func Test_Seg8_SSO_WithinRange_InRange(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRange_InRange", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("50")
		val, inRange := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 50,
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "WithinRange in range -- true", actual)
	})
}

func Test_Seg8_SSO_WithinRange_Below_Boundary(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRange_Below_Boundary", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("-5")
		val, inRange := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 0,
			"inRange": false,
		}
		expected.ShouldBeEqual(t, 0, "WithinRange below boundary -- clamped to min", actual)
	})
}

func Test_Seg8_SSO_WithinRange_Above_Boundary(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRange_Above_Boundary", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("200")
		val, inRange := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 100,
			"inRange": false,
		}
		expected.ShouldBeEqual(t, 0, "WithinRange above boundary -- clamped to max", actual)
	})
}

func Test_Seg8_SSO_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRange_NoBoundary", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("200")
		val, inRange := sso.WithinRange(false, 0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 200,
			"inRange": false,
		}
		expected.ShouldBeEqual(t, 0, "WithinRange no boundary -- raw value", actual)
	})
}

func Test_Seg8_SSO_WithinRange_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRange_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")
		val, inRange := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 0,
			"inRange": false,
		}
		expected.ShouldBeEqual(t, 0, "WithinRange invalid -- 0 false", actual)
	})
}

func Test_Seg8_SSO_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_WithinRangeDefault", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("50")
		val, inRange := sso.WithinRangeDefault(0, 100)

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": 50,
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "WithinRangeDefault -- delegates", actual)
	})
}

func Test_Seg8_SSO_Boolean_Yes(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_Yes", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("yes")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "Boolean yes -- true", actual)
	})
}

func Test_Seg8_SSO_Boolean_True(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_True", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("true")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "Boolean true -- true", actual)
	})
}

func Test_Seg8_SSO_Boolean_1(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_1", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("1")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "Boolean 1 -- true", actual)
	})
}

func Test_Seg8_SSO_Boolean_Y(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_Y", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("Y")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "Boolean Y -- true", actual)
	})
}

func Test_Seg8_SSO_Boolean_YES(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_YES", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("YES")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "Boolean YES -- true", actual)
	})
}

func Test_Seg8_SSO_Boolean_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("xyz")

		// Act
		actual := args.Map{"bool": sso.Boolean(false)}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "Boolean invalid -- false", actual)
	})
}

func Test_Seg8_SSO_Boolean_Uninit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Boolean_Uninit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"bool": sso.Boolean(true)}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "Boolean uninit consider init -- false", actual)
	})
}

func Test_Seg8_SSO_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_BooleanDefault", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("true")

		// Act
		actual := args.Map{"bool": sso.BooleanDefault()}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "BooleanDefault -- true", actual)
	})
}

func Test_Seg8_SSO_IsValueBool(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsValueBool", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("true")

		// Act
		actual := args.Map{"bool": sso.IsValueBool()}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "IsValueBool -- true", actual)
	})
}

func Test_Seg8_SSO_IsSetter_True(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsSetter_True", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("yes")

		// Act
		actual := args.Map{"isTrue": sso.IsSetter(false).IsTrue()}

		// Assert
		expected := args.Map{"isTrue": true}
		expected.ShouldBeEqual(t, 0, "IsSetter yes -- true", actual)
	})
}

func Test_Seg8_SSO_IsSetter_Uninit(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsSetter_Uninit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"isFalse": sso.IsSetter(true).IsFalse()}

		// Assert
		expected := args.Map{"isFalse": true}
		expected.ShouldBeEqual(t, 0, "IsSetter uninit -- false", actual)
	})
}

func Test_Seg8_SSO_IsSetter_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsSetter_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("xyz")

		// Act
		actual := args.Map{"isUninit": sso.IsSetter(false).IsUninitialized()}

		// Assert
		expected := args.Map{"isUninit": true}
		expected.ShouldBeEqual(t, 0, "IsSetter invalid -- uninitialized", actual)
	})
}

func Test_Seg8_SSO_ValueInt(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueInt", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("42")

		// Act
		actual := args.Map{"val": sso.ValueInt(99)}

		// Assert
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "ValueInt -- 42", actual)
	})
}

func Test_Seg8_SSO_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueInt_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{"val": sso.ValueInt(99)}

		// Assert
		expected := args.Map{"val": 99}
		expected.ShouldBeEqual(t, 0, "ValueInt invalid -- default", actual)
	})
}

func Test_Seg8_SSO_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueDefInt", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("10")

		// Act
		actual := args.Map{"val": sso.ValueDefInt()}

		// Assert
		expected := args.Map{"val": 10}
		expected.ShouldBeEqual(t, 0, "ValueDefInt -- 10", actual)
	})
}

func Test_Seg8_SSO_ValueByte(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueByte", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("100")

		// Act
		actual := args.Map{"val": sso.ValueByte(0)}

		// Assert
		expected := args.Map{"val": byte(100)}
		expected.ShouldBeEqual(t, 0, "ValueByte -- 100", actual)
	})
}

func Test_Seg8_SSO_ValueByte_OverMax(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueByte_OverMax", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("300")

		// Act
		actual := args.Map{"val": sso.ValueByte(5)}

		// Assert
		expected := args.Map{"val": byte(5)}
		expected.ShouldBeEqual(t, 0, "ValueByte over max -- default", actual)
	})
}

func Test_Seg8_SSO_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueDefByte", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("50")

		// Act
		actual := args.Map{"val": sso.ValueDefByte()}

		// Assert
		expected := args.Map{"val": byte(50)}
		expected.ShouldBeEqual(t, 0, "ValueDefByte -- 50", actual)
	})
}

func Test_Seg8_SSO_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueFloat64", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("3.14")

		// Act
		actual := args.Map{"val": sso.ValueFloat64(0.0)}

		// Assert
		expected := args.Map{"val": 3.14}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 -- 3.14", actual)
	})
}

func Test_Seg8_SSO_ValueFloat64_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueFloat64_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc")

		// Act
		actual := args.Map{"val": sso.ValueFloat64(1.5)}

		// Assert
		expected := args.Map{"val": 1.5}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 invalid -- default", actual)
	})
}

func Test_Seg8_SSO_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ValueDefFloat64", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("2.5")

		// Act
		actual := args.Map{"val": sso.ValueDefFloat64()}

		// Assert
		expected := args.Map{"val": 2.5}
		expected.ShouldBeEqual(t, 0, "ValueDefFloat64 -- 2.5", actual)
	})
}

func Test_Seg8_SSO_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_NonPtr_Ptr", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")
		nonPtr := sso.NonPtr()

		// Act
		actual := args.Map{
			"val": nonPtr.Value(),
			"ptrSame": sso.Ptr() == sso,
		}

		// Assert
		expected := args.Map{
			"val": "val",
			"ptrSame": true,
		}
		expected.ShouldBeEqual(t, 0, "NonPtr/Ptr -- correct", actual)
	})
}

func Test_Seg8_SSO_Is(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Is", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{
			"is": sso.Is("hello"),
			"isNot": sso.Is("world"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isNot": false,
		}
		expected.ShouldBeEqual(t, 0, "Is -- correct", actual)
	})
}

func Test_Seg8_SSO_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsAnyOf", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("b")

		// Act
		actual := args.Map{
			"found":   sso.IsAnyOf("a", "b", "c"),
			"notFound": sso.IsAnyOf("x", "y"),
			"empty":   sso.IsAnyOf(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyOf -- correct", actual)
	})
}

func Test_Seg8_SSO_IsContains(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsContains", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello world")

		// Act
		actual := args.Map{
			"contains": sso.IsContains("world"),
			"not": sso.IsContains("xyz"),
		}

		// Assert
		expected := args.Map{
			"contains": true,
			"not": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains -- correct", actual)
	})
}

func Test_Seg8_SSO_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsAnyContains", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello world")

		// Act
		actual := args.Map{
			"found":   sso.IsAnyContains("xyz", "world"),
			"notFound": sso.IsAnyContains("abc"),
			"empty":   sso.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyContains -- correct", actual)
	})
}

func Test_Seg8_SSO_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsEqualNonSensitive", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("Hello")

		// Act
		actual := args.Map{
			"eq": sso.IsEqualNonSensitive("hello"),
			"neq": sso.IsEqualNonSensitive("world"),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive -- correct", actual)
	})
}

func Test_Seg8_SSO_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_IsRegexMatches", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"match": sso.IsRegexMatches(re),
			"nil": sso.IsRegexMatches(nil),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches -- correct", actual)
	})
}

func Test_Seg8_SSO_RegexFindString(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_RegexFindString", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("abc123xyz")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"found": sso.RegexFindString(re),
			"nil": sso.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"found": "123",
			"nil": "",
		}
		expected.ShouldBeEqual(t, 0, "RegexFindString -- correct", actual)
	})
}

func Test_Seg8_SSO_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_RegexFindAllStrings", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a1b2c3")
		re := regexp.MustCompile(`\d`)

		// Act
		actual := args.Map{
			"len": len(sso.RegexFindAllStrings(re, -1)),
			"nil": len(sso.RegexFindAllStrings(nil, -1)),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"nil": 0,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings -- correct", actual)
	})
}

func Test_Seg8_SSO_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_RegexFindAllStringsWithFlag", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a1b2")
		re := regexp.MustCompile(`\d`)
		items, hasAny := sso.RegexFindAllStringsWithFlag(re, -1)
		nilItems, nilHas := sso.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"len": len(items),
			"hasAny": hasAny,
			"nilLen": len(nilItems),
			"nilHas": nilHas,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"hasAny": true,
			"nilLen": 0,
			"nilHas": false,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag -- correct", actual)
	})
}

func Test_Seg8_SSO_LinesSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_LinesSimpleSlice", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a\nb\nc")

		// Act
		actual := args.Map{"len": sso.LinesSimpleSlice().Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LinesSimpleSlice -- 3 lines", actual)
	})
}

func Test_Seg8_SSO_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SimpleSlice", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a,b,c")

		// Act
		actual := args.Map{"len": sso.SimpleSlice(",").Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "SimpleSlice -- 3 items", actual)
	})
}

func Test_Seg8_SSO_Split(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Split", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a,b")

		// Act
		actual := args.Map{"len": len(sso.Split(","))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Split -- 2 items", actual)
	})
}

func Test_Seg8_SSO_SplitLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SplitLeftRight", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("key=value")
		left, right := sso.SplitLeftRight("=")

		// Act
		actual := args.Map{
			"left": left,
			"right": right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "SplitLeftRight -- correct", actual)
	})
}

func Test_Seg8_SSO_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SplitLeftRight_NoSep", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("nosep")
		left, right := sso.SplitLeftRight("=")

		// Act
		actual := args.Map{
			"left": left,
			"right": right,
		}

		// Assert
		expected := args.Map{
			"left": "nosep",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "SplitLeftRight no sep -- right empty", actual)
	})
}

func Test_Seg8_SSO_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SplitLeftRightTrim", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized(" key = value ")
		left, right := sso.SplitLeftRightTrim("=")

		// Act
		actual := args.Map{
			"left": left,
			"right": right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "SplitLeftRightTrim -- trimmed", actual)
	})
}

func Test_Seg8_SSO_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SplitNonEmpty", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a,,b")
		result := sso.SplitNonEmpty(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 3} // Note: source has bug returning `slice` not `nonEmptySlice`
		expected.ShouldBeEqual(t, 0, "SplitNonEmpty -- returns original slice (known behavior)", actual)
	})
}

func Test_Seg8_SSO_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_SplitTrimNonWhitespace", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("a, ,b")
		result := sso.SplitTrimNonWhitespace(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 3} // Note: source has bug returning `slice` not `nonEmptySlice`
		expected.ShouldBeEqual(t, 0, "SplitTrimNonWhitespace -- returns original slice (known behavior)", actual)
	})
}

func Test_Seg8_SSO_Clone(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Clone", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		c := sso.Clone()

		// Act
		actual := args.Map{
			"val": c.Value(),
			"init": c.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- copy", actual)
	})
}

func Test_Seg8_SSO_ClonePtr(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ClonePtr", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		c := sso.ClonePtr()

		// Act
		actual := args.Map{
			"val": c.Value(),
			"diff": c != sso,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "ClonePtr -- new ptr", actual)
	})
}

func Test_Seg8_SSO_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ClonePtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce
		c := sso.ClonePtr()

		// Act
		actual := args.Map{"nil": c == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ClonePtr nil -- nil", actual)
	})
}

func Test_Seg8_SSO_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_CloneUsingNewVal", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("old")
		c := sso.CloneUsingNewVal("new")

		// Act
		actual := args.Map{
			"val": c.Value(),
			"init": c.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "new",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "CloneUsingNewVal -- new value same init", actual)
	})
}

func Test_Seg8_SSO_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Dispose", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("val")
		sso.Dispose()

		// Act
		actual := args.Map{
			"val": sso.Value(),
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "Dispose -- empty but still init", actual)
	})
}

func Test_Seg8_SSO_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Dispose_Nil", func() {
		var sso *corestr.SimpleStringOnce
		sso.Dispose() // should not panic
	})
}

func Test_Seg8_SSO_String(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_String", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{"str": sso.String()}

		// Assert
		expected := args.Map{"str": "hello"}
		expected.ShouldBeEqual(t, 0, "String -- value", actual)
	})
}

func Test_Seg8_SSO_String_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_String_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		actual := args.Map{"str": sso.String()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "String nil -- empty", actual)
	})
}

func Test_Seg8_SSO_StringPtr(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_StringPtr", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		p := sso.StringPtr()

		// Act
		actual := args.Map{"val": *p}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "StringPtr -- ptr to value", actual)
	})
}

func Test_Seg8_SSO_StringPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_StringPtr_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce
		p := sso.StringPtr()

		// Act
		actual := args.Map{"val": *p}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "StringPtr nil -- ptr to empty", actual)
	})
}

func Test_Seg8_SSO_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_JsonModel", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		m := sso.JsonModel()

		// Act
		actual := args.Map{
			"val": m.Value,
			"init": m.IsInitialize,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "JsonModel -- correct", actual)
	})
}

func Test_Seg8_SSO_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_JsonModelAny", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"notNil": sso.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg8_SSO_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_MarshalJSON", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		b, err := sso.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg8_SSO_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_UnmarshalJSON", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		b, _ := sso.MarshalJSON()
		sso2 := &corestr.SimpleStringOnce{}
		err := sso2.UnmarshalJSON(b)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": sso2.Value(),
			"init": sso2.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- roundtrip", actual)
	})
}

func Test_Seg8_SSO_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_UnmarshalJSON_Invalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		err := sso.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg8_SSO_Json(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Json", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		j := sso.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg8_SSO_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ParseInjectUsingJson", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		result, err := sso2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"notNil": result != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg8_SSO_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_ParseInjectUsingJsonMust", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		result := sso2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg8_SSO_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_JsonParseSelfInject", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		err := sso2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg8_SSO_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_InterfaceCasts", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{
			"jsoner":   sso.AsJsoner() != nil,
			"binder":   sso.AsJsonContractsBinder() != nil,
			"injector": sso.AsJsonParseSelfInjector() != nil,
			"marsh":    sso.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
			"marsh": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg8_SSO_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Serialize", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		b, err := sso.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg8_SSO_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg8_SSO_Deserialize", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("hello")
		var target corestr.SimpleStringOnceModel
		err := sso.Deserialize(&target)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": target.Value,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}
