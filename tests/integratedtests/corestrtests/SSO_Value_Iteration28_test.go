package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Init/Set/Get
// ══════════════════════════════════════════════════════════════════════════════

func Test_SSO_Value_Empty(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Value_Empty", func() {
		// Arrange
		var sso corestr.SimpleStringOnce

		// Act
		actual := args.Map{
			"val": sso.Value(),
			"init": sso.IsInitialized(),
			"defined": sso.IsDefined(),
			"uninit": sso.IsUninitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"init": false,
			"defined": false,
			"uninit": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns empty -- empty", actual)
	})
}

func Test_SSO_SetOnUninitialized_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_SetOnUninitialized", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		err := sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": sso.Value(),
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- SetOnUninitialized", actual)
	})
}

func Test_SSO_SetOnUninitialized_AlreadyInit_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
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
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- SetOnUninitialized already init", actual)
	})
}

func Test_SSO_GetSetOnce_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_GetSetOnce", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		v1 := sso.GetSetOnce("first")
		v2 := sso.GetSetOnce("second")

		// Act
		actual := args.Map{
			"v1": v1,
			"v2": v2,
		}

		// Assert
		expected := args.Map{
			"v1": "first",
			"v2": "first",
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- GetSetOnce", actual)
	})
}

func Test_SSO_GetOnce_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_GetOnce", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		v := sso.GetOnce()

		// Act
		actual := args.Map{
			"val": v,
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- GetOnce", actual)
	})
}

func Test_SSO_GetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_I28_SSO_GetOnce_AlreadyInit", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("hello")
		v := sso.GetOnce()

		// Act
		actual := args.Map{"val": v}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- GetOnce already init", actual)
	})
}

func Test_SSO_GetOnceFunc_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_GetOnceFunc", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		v := sso.GetOnceFunc(func() string { return "computed" })
		v2 := sso.GetOnceFunc(func() string { return "other" })

		// Act
		actual := args.Map{
			"v": v,
			"v2": v2,
		}

		// Assert
		expected := args.Map{
			"v": "computed",
			"v2": "computed",
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- GetOnceFunc", actual)
	})
}

func Test_SSO_SetOnceIfUninitialized_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_SetOnceIfUninitialized", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		ok1 := sso.SetOnceIfUninitialized("hello")
		ok2 := sso.SetOnceIfUninitialized("world")

		// Act
		actual := args.Map{
			"ok1": ok1,
			"ok2": ok2,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"ok1": true,
			"ok2": false,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- SetOnceIfUninitialized", actual)
	})
}

func Test_SSO_Invalidate_Reset(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Invalidate_Reset", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("hello")
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
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Invalidate", actual)

		sso.GetSetOnce("world")
		sso.Reset()
		actual2 := args.Map{"init": sso.IsInitialized()}
		expected2 := args.Map{"init": false}
		expected2.ShouldBeEqual(t, 0, "SSO returns correct value -- Reset", actual2)
	})
}

func Test_SSO_IsInvalid_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsInvalid", func() {
		// Arrange
		var sso corestr.SimpleStringOnce

		// Act
		actual := args.Map{"invalid": sso.IsInvalid()}

		// Assert
		expected := args.Map{"invalid": true}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- IsInvalid uninit", actual)

		sso.GetSetOnce("hello")
		actual2 := args.Map{"invalid": sso.IsInvalid()}
		expected2 := args.Map{"invalid": false}
		expected2.ShouldBeEqual(t, 0, "SSO returns error -- IsInvalid init", actual2)
	})
}

func Test_SSO_IsInvalid_Nil(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsInvalid_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act
		actual := args.Map{"invalid": sso.IsInvalid()}

		// Assert
		expected := args.Map{"invalid": true}
		expected.ShouldBeEqual(t, 0, "SSO returns nil -- IsInvalid nil", actual)
	})
}

func Test_SSO_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_I28_SSO_SetInitialize_SetUnInit", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.SetInitialize()

		// Act
		actual := args.Map{"init": sso.IsInitialized()}

		// Assert
		expected := args.Map{"init": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- SetInitialize", actual)

		sso.SetUnInit()
		actual2 := args.Map{"init": sso.IsInitialized()}
		expected2 := args.Map{"init": false}
		expected2.ShouldBeEqual(t, 0, "SSO returns correct value -- SetUnInit", actual2)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Bytes, Checks
// ══════════════════════════════════════════════════════════════════════════════

func Test_SSO_ValueBytes_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ValueBytes", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{
			"len": len(sso.ValueBytes()),
			"lenPtr": len(sso.ValueBytesPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"lenPtr": 3,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ValueBytes", actual)
	})
}

func Test_SSO_IsEmpty_IsWhitespace_Trim(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsEmpty_IsWhitespace_Trim", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("  hi  ")

		// Act
		actual := args.Map{
			"empty": sso.IsEmpty(),
			"ws": sso.IsWhitespace(),
			"trim": sso.Trim(),
		}

		// Assert
		expected := args.Map{
			"empty": false,
			"ws": false,
			"trim": "hi",
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- checks", actual)
	})
}

func Test_SSO_HasValidNonEmpty_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I28_SSO_HasValidNonEmpty_HasValidNonWhitespace", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("x")

		// Act
		actual := args.Map{
			"hv": sso.HasValidNonEmpty(),
			"hvw": sso.HasValidNonWhitespace(),
			"safe": sso.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"hv": true,
			"hvw": true,
			"safe": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- HasValid", actual)
	})
}

func Test_SSO_SafeValue_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_SafeValue", func() {
		// Arrange
		var sso corestr.SimpleStringOnce

		// Act
		actual := args.Map{"uninit": sso.SafeValue()}

		// Assert
		expected := args.Map{"uninit": ""}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- SafeValue uninit", actual)

		sso.GetSetOnce("hello")
		actual2 := args.Map{"init": sso.SafeValue()}
		expected2 := args.Map{"init": "hello"}
		expected2.ShouldBeEqual(t, 0, "SSO returns correct value -- SafeValue init", actual2)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Numeric conversions
// ══════════════════════════════════════════════════════════════════════════════

func Test_SSO_Int_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("42")

		// Act
		actual := args.Map{"val": sso.Int()}

		// Assert
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Int", actual)
	})
}

func Test_SSO_Int_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{"val": sso.Int()}

		// Assert
		expected := args.Map{"val": 0}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Int err", actual)
	})
}

func Test_SSO_Byte_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Byte", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("100")

		// Act
		actual := args.Map{"val": sso.Byte()}

		// Assert
		expected := args.Map{"val": byte(100)}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Byte", actual)
	})
}

func Test_SSO_Byte_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Byte_OutOfRange", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("300")

		// Act
		actual := args.Map{"val": sso.Byte()}

		// Assert
		expected := args.Map{"val": byte(0)}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Byte out of range", actual)
	})
}

func Test_SSO_Byte_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Byte_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{"val": sso.Byte()}

		// Assert
		expected := args.Map{"val": byte(0)}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Byte err", actual)
	})
}

func Test_SSO_Int16_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int16", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("100")

		// Act
		actual := args.Map{"val": sso.Int16()}

		// Assert
		expected := args.Map{"val": int16(100)}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Int16", actual)
	})
}

func Test_SSO_Int16_OutOfRange_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int16_OutOfRange", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("99999")

		// Act
		actual := args.Map{"val": sso.Int16()}

		// Assert
		expected := args.Map{"val": int16(0)}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Int16 out of range", actual)
	})
}

func Test_SSO_Int16_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int16_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{"val": sso.Int16()}

		// Assert
		expected := args.Map{"val": int16(0)}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Int16 err", actual)
	})
}

func Test_SSO_Int32_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int32", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("1000")

		// Act
		actual := args.Map{"val": sso.Int32()}

		// Assert
		expected := args.Map{"val": int32(1000)}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Int32", actual)
	})
}

func Test_SSO_Int32_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Int32_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{"val": sso.Int32()}

		// Assert
		expected := args.Map{"val": int32(0)}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Int32 err", actual)
	})
}

func Test_SSO_Uint16(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Uint16", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("100")
		val, inRange := sso.Uint16()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint16(100),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Uint16", actual)
	})
}

func Test_SSO_Uint32(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Uint32", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("1000")
		val, inRange := sso.Uint32()

		// Act
		actual := args.Map{
			"val": val,
			"inRange": inRange,
		}

		// Assert
		expected := args.Map{
			"val": uint32(1000),
			"inRange": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Uint32", actual)
	})
}

func Test_SSO_WithinRange_InRange_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRange_InRange", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("50")
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
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- WithinRange in range", actual)
	})
}

func Test_SSO_WithinRange_Below(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRange_Below", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("-5")
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
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- WithinRange below", actual)
	})
}

func Test_SSO_WithinRange_Above(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRange_Above", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("200")
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
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- WithinRange above", actual)
	})
}

func Test_SSO_WithinRange_NoBoundary_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRange_NoBoundary", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("200")
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
		expected.ShouldBeEqual(t, 0, "SSO returns empty -- WithinRange no boundary", actual)
	})
}

func Test_SSO_WithinRange_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRange_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")
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
		expected.ShouldBeEqual(t, 0, "SSO returns error -- WithinRange err", actual)
	})
}

func Test_SSO_WithinRangeDefault_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_WithinRangeDefault", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("50")
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
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- WithinRangeDefault", actual)
	})
}

func Test_SSO_Boolean_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Boolean", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("yes")

		// Act
		actual := args.Map{"val": sso.Boolean(false)}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Boolean yes", actual)
	})
}

func Test_SSO_Boolean_True(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Boolean_True", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("true")

		// Act
		actual := args.Map{"val": sso.Boolean(false)}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "SSO returns non-empty -- Boolean true", actual)
	})
}

func Test_SSO_Boolean_ConsiderInit_Uninit_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Boolean_ConsiderInit_Uninit", func() {
		// Arrange
		var sso corestr.SimpleStringOnce

		// Act
		actual := args.Map{"val": sso.Boolean(true)}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- Boolean consider init uninit", actual)
	})
}

func Test_SSO_Boolean_ParseErr_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_Boolean_ParseErr", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{"val": sso.Boolean(false)}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- Boolean parse err", actual)
	})
}

func Test_SSO_BooleanDefault_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_BooleanDefault", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("y")

		// Act
		actual := args.Map{"val": sso.BooleanDefault()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- BooleanDefault", actual)
	})
}

func Test_SSO_IsValueBool_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsValueBool", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("1")

		// Act
		actual := args.Map{"val": sso.IsValueBool()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- IsValueBool", actual)
	})
}

func Test_SSO_IsSetter(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsSetter", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("yes")
		is := sso.IsSetter(false)

		// Act
		actual := args.Map{"true": is.IsTrue()}

		// Assert
		expected := args.Map{"true": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- IsSetter yes", actual)
	})
}

func Test_SSO_IsSetter_ConsiderInit_Uninit(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsSetter_ConsiderInit_Uninit", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		is := sso.IsSetter(true)

		// Act
		actual := args.Map{"false": is.IsFalse()}

		// Assert
		expected := args.Map{"false": true}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- IsSetter uninit", actual)
	})
}

func Test_SSO_IsSetter_ParseErr_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_IsSetter_ParseErr", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")
		is := sso.IsSetter(false)

		// Act
		actual := args.Map{"uninit": is.IsUninitialized()}

		// Assert
		expected := args.Map{"uninit": true}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- IsSetter parse err", actual)
	})
}

func Test_SSO_ValueInt_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ValueInt", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("42")

		// Act
		actual := args.Map{
			"val": sso.ValueInt(0),
			"defInt": sso.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"defInt": 42,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ValueInt", actual)
	})
}

func Test_SSO_ValueInt_Err(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ValueInt_Err", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("abc")

		// Act
		actual := args.Map{
			"val": sso.ValueInt(99),
			"defInt": sso.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 99,
			"defInt": 0,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns error -- ValueInt err", actual)
	})
}

func Test_SSO_ValueByte(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ValueByte", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("100")

		// Act
		actual := args.Map{
			"val": sso.ValueByte(0),
			"def": sso.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(100),
			"def": byte(100),
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ValueByte", actual)
	})
}

func Test_SSO_ValueFloat64_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ValueFloat64", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("3.14")

		// Act
		actual := args.Map{
			"close": sso.ValueFloat64(0) > 3.1,
			"def": sso.ValueDefFloat64() > 3.1,
		}

		// Assert
		expected := args.Map{
			"close": true,
			"def": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ValueFloat64", actual)
	})
}

func Test_SSO_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_I28_SSO_NonPtr_Ptr", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("hello")
		np := sso.NonPtr()
		p := sso.Ptr()

		// Act
		actual := args.Map{
			"npVal": np.Value(),
			"pSame": p == &sso,
		}

		// Assert
		expected := args.Map{
			"npVal": "hello",
			"pSame": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- NonPtr/Ptr", actual)
	})
}

func Test_SSO_ConcatNew_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ConcatNew", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("hello")
		newSSO := sso.ConcatNew(" world")

		// Act
		actual := args.Map{"val": newSSO.Value()}

		// Assert
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ConcatNew", actual)
	})
}

func Test_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_I28_SSO_ConcatNewUsingStrings", func() {
		// Arrange
		var sso corestr.SimpleStringOnce
		sso.GetSetOnce("a")
		newSSO := sso.ConcatNewUsingStrings("-", "b", "c")

		// Act
		actual := args.Map{"val": newSSO.Value()}

		// Assert
		expected := args.Map{"val": "a-b-c"}
		expected.ShouldBeEqual(t, 0, "SSO returns correct value -- ConcatNewUsingStrings", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashmapDiff_Length_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Length", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}

		// Act
		actual := args.Map{
			"len": hd.Length(),
			"empty": hd.IsEmpty(),
			"hasAny": hd.HasAnyItem(),
			"lastIdx": hd.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"empty": false,
			"hasAny": true,
			"lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- basics", actual)
	})
}

func Test_HashmapDiff_Nil_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": hd.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns nil -- nil length", actual)
	})
}

func Test_HashmapDiff_AllKeysSorted_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_AllKeysSorted", func() {
		// Arrange
		hd := corestr.HashmapDiff{"b": "2", "a": "1"}
		keys := hd.AllKeysSorted()

		// Act
		actual := args.Map{
			"first": keys[0],
			"second": keys[1],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"second": "b",
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- AllKeysSorted", actual)
	})
}

func Test_HashmapDiff_MapAnyItems_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_MapAnyItems", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		mai := hd.MapAnyItems()

		// Act
		actual := args.Map{"len": len(mai)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- MapAnyItems", actual)
	})
}

func Test_HashmapDiff_MapAnyItems_Nil_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_MapAnyItems_Nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff
		mai := hd.MapAnyItems()

		// Act
		actual := args.Map{"len": len(mai)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns nil -- MapAnyItems nil", actual)
	})
}

func Test_HashmapDiff_Raw_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Raw", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		raw := hd.Raw()

		// Act
		actual := args.Map{"len": len(raw)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- Raw", actual)
	})
}

func Test_HashmapDiff_Raw_Nil_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Raw_Nil", func() {
		// Arrange
		var hd *corestr.HashmapDiff
		raw := hd.Raw()

		// Act
		actual := args.Map{"len": len(raw)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns nil -- Raw nil", actual)
	})
}

func Test_HashmapDiff_IsRawEqual_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_IsRawEqual", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}

		// Act
		actual := args.Map{
			"eq": hd.IsRawEqual(map[string]string{"a": "1"}),
			"neq": hd.IsRawEqual(map[string]string{"a": "2"}),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- IsRawEqual", actual)
	})
}

func Test_HashmapDiff_HasAnyChanges_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_HasAnyChanges", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}

		// Act
		actual := args.Map{
			"changes": hd.HasAnyChanges(map[string]string{"a": "2"}),
			"noChanges": hd.HasAnyChanges(map[string]string{"a": "1"}),
		}

		// Assert
		expected := args.Map{
			"changes": true,
			"noChanges": false,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- HasAnyChanges", actual)
	})
}

func Test_HashmapDiff_DiffRaw_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_DiffRaw", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}
		diff := hd.DiffRaw(map[string]string{"a": "1", "b": "99"})

		// Act
		actual := args.Map{"hasDiff": len(diff) > 0}

		// Assert
		expected := args.Map{"hasDiff": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- DiffRaw", actual)
	})
}

func Test_HashmapDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_HashmapDiffUsingRaw_NoDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		diff := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"empty": diff.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns empty -- HashmapDiffUsingRaw no diff", actual)
	})
}

func Test_HashmapDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_HashmapDiffUsingRaw_HasDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		diff := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasDiff": diff.HasAnyItem()}

		// Assert
		expected := args.Map{"hasDiff": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- HashmapDiffUsingRaw has diff", actual)
	})
}

func Test_HashmapDiff_DiffJsonMessage_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_DiffJsonMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		msg := hd.DiffJsonMessage(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"notEmpty": msg != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- DiffJsonMessage", actual)
	})
}

func Test_HashmapDiff_ShouldDiffMessage_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_ShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"notEmpty": msg != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- ShouldDiffMessage", actual)
	})
}

func Test_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_LogShouldDiffMessage", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"notEmpty": msg != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- LogShouldDiffMessage", actual)
	})
}

func Test_HashmapDiff_ToStringsSliceOfDiffMap_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_ToStringsSliceOfDiffMap", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		diff := hd.DiffRaw(map[string]string{"a": "2"})
		strs := hd.ToStringsSliceOfDiffMap(diff)

		// Act
		actual := args.Map{"hasItems": len(strs) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- ToStringsSliceOfDiffMap", actual)
	})
}

func Test_HashmapDiff_RawMapStringAnyDiff_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_RawMapStringAnyDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		d := hd.RawMapStringAnyDiff()

		// Act
		actual := args.Map{"notNil": d != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- RawMapStringAnyDiff", actual)
	})
}

func Test_HashmapDiff_Serialize_FromSSOValueIteration28(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Serialize", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		b, err := hd.Serialize()

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
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- Serialize", actual)
	})
}

func Test_HashmapDiff_Deserialize(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDiff_Deserialize", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1"}
		target := map[string]string{}
		err := hd.Deserialize(&target)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff returns correct value -- Deserialize", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashmapDataModel_NewUsing(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDataModel_NewUsing", func() {
		// Arrange
		dm := &corestr.HashmapDataModel{Items: map[string]string{"a": "1"}}
		hm := corestr.NewHashmapUsingDataModel(dm)

		// Act
		actual := args.Map{
			"notNil": hm != nil,
			"has": hm.Has("a"),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDataModel returns correct value -- NewUsing", actual)
	})
}

func Test_HashmapDataModel_NewFromCollection(t *testing.T) {
	safeTest(t, "Test_I28_HashmapDataModel_NewFromCollection", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)

		// Act
		actual := args.Map{
			"notNil": dm != nil,
			"len": len(dm.Items),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDataModel returns correct value -- NewFromCollection", actual)
	})
}
