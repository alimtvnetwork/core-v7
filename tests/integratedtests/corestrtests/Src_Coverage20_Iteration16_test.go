package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC20_ValidValue_Constructors_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_Constructors_Verification", func() {
		// Arrange
		tc := srcC20ValidValueConstructorsTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			vv := corestr.NewValidValueUsingAny(false, true, "hello")
			_ = vv.IsValid
			_ = corestr.NewValidValueUsingAnyAutoValid(false, "hello")
			_ = corestr.NewValidValueEmpty()
			_ = corestr.InvalidValidValueNoMessage()
			_ = corestr.InvalidValidValue("msg")
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC20_ValidValue_StringChecks_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_StringChecks_Verification", func() {
		// Arrange
		tc := srcC20ValidValueStringChecksTestCase
		vv := corestr.NewValidValue("  hello  ")

		// Act
		actual := args.Map{
			"isEmpty":   vv.IsEmpty(),
			"isWhite":   vv.IsWhitespace(),
			"trim":      vv.Trim(),
			"validNonE": vv.HasValidNonEmpty(),
			"validNonW": vv.HasValidNonWhitespace(),
			"safeNonE":  vv.HasSafeNonEmpty(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValue_Conversions_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_Conversions_Verification", func() {
		// Arrange
		tc := srcC20ValidValueConversionsTestCase

		// Act
		vvBool := corestr.NewValidValue("true")
		vvNum := corestr.NewValidValue("42")
		vvFloat := corestr.NewValidValue("3.14")
		vvInvalid := corestr.NewValidValue("notnum")
		vvEmpty := corestr.NewValidValue("")
		actual := args.Map{
			"boolTrue":    vvBool.ValueBool(),
			"int42":       vvNum.ValueInt(0),
			"defInt42":    vvNum.ValueDefInt(),
			"byte42":      int(vvNum.ValueByte(0)),
			"defByte42":   int(vvNum.ValueDefByte()),
			"float314":    vvFloat.ValueFloat64(0),
			"defFloat314": vvFloat.ValueDefFloat64(),
			"boolInvalid": vvInvalid.ValueBool(),
			"intDefault":  vvInvalid.ValueInt(99),
			"boolEmpty":   vvEmpty.ValueBool(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValue_Comparisons_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_Comparisons_Verification", func() {
		// Arrange
		tc := srcC20ValidValueComparisonsTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			vv := corestr.NewValidValue("hello")
			_ = vv.Is("hello")
			_ = vv.IsAnyOf("world", "hello")
			_ = vv.IsAnyOf("world", "foo")
			_ = vv.IsAnyOf()
			_ = vv.IsContains("ell")
			_ = vv.IsAnyContains("xyz", "ell")
			_ = vv.IsAnyContains("xyz", "abc")
			_ = vv.IsAnyContains()
			_ = vv.IsEqualNonSensitive("HELLO")
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValue_CloneDispose_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_CloneDispose_Verification", func() {
		// Arrange
		tc := srcC20ValidValueCloneDisposeTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			var vvNil *corestr.ValidValue
			_ = vvNil.Clone()
			_ = vvNil.String()
			_ = vvNil.FullString()
			vvNil.Clear()
			vvNil.Dispose()
			vv := corestr.NewValidValue("test")
			c := vv.Clone()
			_ = c.Value
			vv.Clear()
			vv2 := corestr.NewValidValue("x")
			vv2.Dispose()
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValue_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_Json_Verification", func() {
		// Arrange
		tc := srcC20ValidValueJsonTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			vv := corestr.NewValidValue("test")
			_ = vv.Json()
			jp := vv.JsonPtr()
			_ = jp
			b, _ := vv.Serialize()
			_ = len(b)
			jr := vv.JsonPtr()
			vv2 := &corestr.ValidValue{}
			_, _ = vv2.ParseInjectUsingJson(jr)
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValue_Split_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValue_Split_Verification", func() {
		// Arrange
		tc := srcC20ValidValueSplitTestCase
		vv := corestr.NewValidValue("a,b,c")

		// Act
		actual := args.Map{
			"splitLen":    len(vv.Split(",")),
			"nonEmptyGe1": len(vv.SplitNonEmpty(",")) > 0,
			"trimGe1":     len(vv.SplitTrimNonWhitespace(",")) > 0,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValidValues_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValidValues_Verification", func() {
		// Arrange
		tc := srcC20ValidValuesTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			vvs := corestr.EmptyValidValues()
			_ = vvs.HasAnyItem()
			vvs.Add("a").Add("b")
			_ = vvs.Length()
			_ = vvs.Count()
			_ = vvs.HasAnyItem()
			_ = vvs.HasIndex(1)
			_ = vvs.SafeValueAt(0)
			_ = vvs.SafeValueAt(99)
			// UsingValues
			v1 := corestr.ValidValue{Value: "x", IsValid: true}
			_ = corestr.NewValidValuesUsingValues(v1)
			_ = corestr.NewValidValuesUsingValues()
			// Strings, FullStrings, String
			_ = corestr.EmptyValidValues().Strings()
			vvs2 := corestr.EmptyValidValues()
			vvs2.Add("x")
			_ = vvs2.Strings()
			_ = vvs2.FullStrings()
			_ = vvs2.String()
			// Find
			_ = corestr.EmptyValidValues().Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
				return v, true, false
			})
			vvs3 := corestr.EmptyValidValues()
			vvs3.Add("a").Add("b").Add("c")
			_ = vvs3.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
				if v.Value == "b" {
					return v, true, true
				}
				return nil, false, false
			})
			// SafeValidValueAt
			vvs4 := corestr.EmptyValidValues()
			vvs4.AddFull(true, "valid", "")
			vvs4.AddFull(false, "invalid", "err")
			_ = vvs4.SafeValidValueAt(0)
			_ = vvs4.SafeValidValueAt(1)
			_ = vvs4.SafeValidValueAt(99)
			// SafeIndexes
			vvs5 := corestr.EmptyValidValues()
			vvs5.Add("a").Add("b")
			_ = vvs5.SafeValuesAtIndexes(0, 1)
			_ = vvs5.SafeValidValuesAtIndexes(0, 1)
			_ = vvs5.SafeValuesAtIndexes()
			// ConcatNew
			vvs6 := corestr.EmptyValidValues()
			vvs6.Add("a")
			_ = vvs6.ConcatNew(true)
			_ = vvs6.ConcatNew(false)
			other := corestr.EmptyValidValues()
			other.Add("b")
			_ = vvs6.ConcatNew(false, other)
			// AddValidValues
			vvs6.AddValidValues(nil)
			vvs6.AddValidValues(corestr.EmptyValidValues())
			vvs6.Adds()
			vvs6.AddsPtr()
			// Hashmap, Map
			_ = corestr.EmptyValidValues().Hashmap()
			_ = corestr.EmptyValidValues().Map()
			vvs7 := corestr.EmptyValidValues()
			vvs7.Add("k")
			_ = vvs7.Hashmap()
			// AddHashset
			vvs8 := corestr.EmptyValidValues()
			vvs8.AddHashsetMap(nil)
			vvs8.AddHashset(nil)
			vvs8.AddHashsetMap(map[string]bool{"a": true})
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_ValueStatus_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_ValueStatus_Verification", func() {
		// Arrange
		tc := srcC20ValueStatusTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			_ = corestr.InvalidValueStatusNoMessage()
			vs := corestr.InvalidValueStatus("err")
			_ = vs.Clone()
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_TextWithLineNumber_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_TextWithLineNumber_Verification", func() {
		// Arrange
		tc := srcC20TextWithLineNumberTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			var tln *corestr.TextWithLineNumber
			_ = tln.HasLineNumber()
			_ = tln.IsInvalidLineNumber()
			_ = tln.Length()
			_ = tln.IsEmpty()
			_ = tln.IsEmptyText()
			tln2 := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
			_ = tln2.HasLineNumber()
			_ = tln2.IsInvalidLineNumber()
			_ = tln2.Length()
			_ = tln2.IsEmpty()
			_ = tln2.IsEmptyText()
			_ = tln2.IsEmptyTextLineBoth()
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_LeftRight_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_LeftRight_Verification", func() {
		// Arrange
		tc := srcC20LeftRightTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			_ = corestr.InvalidLeftRightNoMessage()
			_ = corestr.InvalidLeftRight("msg")
			lr := corestr.NewLeftRight("a", "b")
			_ = lr.IsValid
			// UsingSlice
			_ = corestr.LeftRightUsingSlice(nil)
			_ = corestr.LeftRightUsingSlice([]string{"a"})
			_ = corestr.LeftRightUsingSlice([]string{"a", "b"})
			// UsingSlicePtr
			_ = corestr.LeftRightUsingSlicePtr([]string{})
			_ = corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
			// TrimmedUsingSlice
			_ = corestr.LeftRightTrimmedUsingSlice(nil)
			_ = corestr.LeftRightTrimmedUsingSlice([]string{})
			_ = corestr.LeftRightTrimmedUsingSlice([]string{" a "})
			_ = corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
			// Methods
			lr2 := corestr.NewLeftRight("hello", "world")
			_ = lr2.LeftBytes()
			_ = lr2.RightBytes()
			_ = lr2.LeftTrim()
			_ = lr2.IsLeftEmpty()
			_ = lr2.IsRightEmpty()
			_ = lr2.IsLeftWhitespace()
			_ = lr2.IsRightWhitespace()
			_ = lr2.HasValidNonEmptyLeft()
			_ = lr2.HasValidNonEmptyRight()
			_ = lr2.HasSafeNonEmpty()
			_ = lr2.NonPtr()
			_ = lr2.Ptr()
			_ = lr2.IsLeft("hello")
			_ = lr2.IsRight("world")
			_ = lr2.Is("hello", "world")
			// IsEqual
			var lrNil *corestr.LeftRight
			_ = lrNil.IsEqual(nil)
			_ = lrNil.IsEqual(lr2)
			_ = lr2.IsEqual(lrNil)
			lr3 := corestr.NewLeftRight("hello", "world")
			_ = lr2.IsEqual(lr3)
			// Clone, Clear, Dispose
			_ = lr2.Clone()
			lr2.Clear()
			lr2.Dispose()
			lrNil.Clear()
			lrNil.Dispose()
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC20_LeftMiddleRight_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC20_LeftMiddleRight_Verification", func() {
		// Arrange
		tc := srcC20LeftMiddleRightTestCase

		// Act
		noPanic := !callPanicsSrcC20(func() {
			_ = corestr.InvalidLeftMiddleRightNoMessage()
			_ = corestr.InvalidLeftMiddleRight("msg")
			lmr := corestr.NewLeftMiddleRight("a", "m", "b")
			_ = lmr.IsValid
			_ = lmr.LeftBytes()
			_ = lmr.RightBytes()
			_ = lmr.MiddleBytes()
			_ = lmr.LeftTrim()
			_ = lmr.RightTrim()
			_ = lmr.MiddleTrim()
			_ = lmr.IsLeftEmpty()
			_ = lmr.IsRightEmpty()
			_ = lmr.IsMiddleEmpty()
			_ = lmr.HasValidNonEmptyLeft()
			_ = lmr.HasValidNonEmptyRight()
			_ = lmr.HasValidNonEmptyMiddle()
			_ = lmr.HasSafeNonEmpty()
			_ = lmr.IsAll("a", "m", "b")
			_ = lmr.Is("a", "b")
			_ = lmr.ToLeftRight()
			_ = lmr.Clone()
			lmr.Clear()
			lmr.Dispose()
			var lmrNil *corestr.LeftMiddleRight
			lmrNil.Clear()
			lmrNil.Dispose()
		})
		actual := args.Map{"noPanic": noPanic}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC20(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
}
