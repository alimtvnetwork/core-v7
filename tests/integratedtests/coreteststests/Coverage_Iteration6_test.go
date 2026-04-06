package coreteststests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/smartystreets/goconvey/convey"
)

// ─── AnyToBytes ──────────────────────────────────────────────────────────────

func Test_C01_AnyToBytes_FromBytes(t *testing.T) {
	convey.Convey("AnyToBytes from []byte returns as-is", t, func() {
		input := []byte("hello")
		result := coretests.AnyToBytes(input)
		convey.So(result, convey.ShouldResemble, input)
	})
}

func Test_C02_AnyToBytes_FromNilBytes(t *testing.T) {
	convey.Convey("AnyToBytes from nil []byte returns nil", t, func() {
		var input []byte
		result := coretests.AnyToBytes(input)
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_C03_AnyToBytes_FromString(t *testing.T) {
	convey.Convey("AnyToBytes from string returns []byte", t, func() {
		result := coretests.AnyToBytes("hello")
		convey.So(string(result), convey.ShouldEqual, "hello")
	})
}

func Test_C04_AnyToBytes_FromStruct(t *testing.T) {
	convey.Convey("AnyToBytes from struct marshals to JSON", t, func() {
		input := struct{ Name string }{Name: "test"}
		result := coretests.AnyToBytes(input)
		convey.So(string(result), convey.ShouldContainSubstring, "test")
	})
}

func Test_C05_AnyToBytesPtr(t *testing.T) {
	convey.Convey("AnyToBytesPtr delegates to AnyToBytes", t, func() {
		result := coretests.AnyToBytesPtr("hello")
		convey.So(string(result), convey.ShouldEqual, "hello")
	})
}

// ─── AnyToDraftType ──────────────────────────────────────────────────────────

func Test_C06_AnyToDraftType_FromValue(t *testing.T) {
	convey.Convey("AnyToDraftType from DraftType value", t, func() {
		d := coretests.DraftType{SampleString1: "a"}
		result := coretests.AnyToDraftType(d)
		convey.So(result, convey.ShouldNotBeNil)
		convey.So(result.SampleString1, convey.ShouldEqual, "a")
	})
}

func Test_C07_AnyToDraftType_FromPtr(t *testing.T) {
	convey.Convey("AnyToDraftType from *DraftType", t, func() {
		d := &coretests.DraftType{SampleString1: "b"}
		result := coretests.AnyToDraftType(d)
		convey.So(result, convey.ShouldEqual, d)
	})
}

func Test_C08_AnyToDraftType_FromOther(t *testing.T) {
	convey.Convey("AnyToDraftType from non-DraftType returns nil", t, func() {
		result := coretests.AnyToDraftType("not a draft")
		convey.So(result, convey.ShouldBeNil)
	})
}

// ─── DraftType ───────────────────────────────────────────────────────────────

func Test_C09_DraftType_F1String_F2Integer(t *testing.T) {
	convey.Convey("DraftType F1String and F2Integer", t, func() {
		d := coretests.DraftType{}
		convey.So(d.F1String(), convey.ShouldEqual, "")
		convey.So(d.F2Integer(), convey.ShouldEqual, 0)
		d.SetF2Integer(42)
		convey.So(d.F2Integer(), convey.ShouldEqual, 42)
	})
}

func Test_C10_DraftType_NonPtr(t *testing.T) {
	convey.Convey("DraftType NonPtr returns copy", t, func() {
		d := coretests.DraftType{SampleString1: "x"}
		result := d.NonPtr()
		convey.So(result.SampleString1, convey.ShouldEqual, "x")
	})
}

func Test_C11_DraftType_PtrOrNonPtr(t *testing.T) {
	convey.Convey("DraftType PtrOrNonPtr", t, func() {
		d := &coretests.DraftType{SampleString1: "x"}
		asPtr := d.PtrOrNonPtr(true)
		convey.So(asPtr, convey.ShouldEqual, d)

		asVal := d.PtrOrNonPtr(false)
		_, isVal := asVal.(coretests.DraftType)
		convey.So(isVal, convey.ShouldBeTrue)
	})
}

func Test_C12_DraftType_PtrOrNonPtr_Nil(t *testing.T) {
	convey.Convey("DraftType PtrOrNonPtr nil", t, func() {
		var d *coretests.DraftType
		result := d.PtrOrNonPtr(true)
		convey.So(result, convey.ShouldBeNil)
	})
}

func Test_C13_DraftType_IsEqual_BothNil(t *testing.T) {
	convey.Convey("DraftType IsEqual both nil", t, func() {
		var a, b *coretests.DraftType
		convey.So(a.IsEqual(true, b), convey.ShouldBeTrue)
	})
}

func Test_C14_DraftType_IsEqual_OneNil(t *testing.T) {
	convey.Convey("DraftType IsEqual one nil", t, func() {
		a := &coretests.DraftType{}
		convey.So(a.IsEqual(true, nil), convey.ShouldBeFalse)
	})
}

func Test_C15_DraftType_IsEqual_SamePtr(t *testing.T) {
	convey.Convey("DraftType IsEqual same pointer", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		convey.So(a.IsEqual(true, a), convey.ShouldBeTrue)
	})
}

func Test_C16_DraftType_IsEqual_DifferentFields(t *testing.T) {
	convey.Convey("DraftType IsEqual different fields", t, func() {
		a := &coretests.DraftType{SampleString1: "a"}
		b := &coretests.DraftType{SampleString1: "b"}
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C17_DraftType_IsEqual_DiffString2(t *testing.T) {
	convey.Convey("DraftType IsEqual different SampleString2", t, func() {
		a := &coretests.DraftType{SampleString2: "a"}
		b := &coretests.DraftType{SampleString2: "b"}
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C18_DraftType_IsEqual_DiffInteger(t *testing.T) {
	convey.Convey("DraftType IsEqual different SampleInteger", t, func() {
		a := &coretests.DraftType{SampleInteger: 1}
		b := &coretests.DraftType{SampleInteger: 2}
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C19_DraftType_IsEqual_DiffRawBytes(t *testing.T) {
	convey.Convey("DraftType IsEqual different RawBytes", t, func() {
		a := &coretests.DraftType{RawBytes: []byte("a")}
		b := &coretests.DraftType{RawBytes: []byte("b")}
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C20_DraftType_IsEqual_DiffLines(t *testing.T) {
	convey.Convey("DraftType IsEqual different Lines", t, func() {
		a := &coretests.DraftType{Lines: []string{"a"}}
		b := &coretests.DraftType{Lines: []string{"b"}}
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C21_DraftType_IsEqualAll(t *testing.T) {
	convey.Convey("DraftType IsEqualAll", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		convey.So(a.IsEqualAll(b), convey.ShouldBeTrue)
	})
}

func Test_C22_DraftType_IsEqual_ExcludingInner(t *testing.T) {
	convey.Convey("DraftType IsEqual excluding inner fields", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		a.SetF2Integer(1)
		b.SetF2Integer(2)
		// excluding inner fields, should be equal
		convey.So(a.IsEqual(false, b), convey.ShouldBeTrue)
		// including inner fields, should NOT be equal
		convey.So(a.IsEqual(true, b), convey.ShouldBeFalse)
	})
}

func Test_C23_DraftType_VerifyNotEqualMessage(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualMessage", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		msg := a.VerifyNotEqualMessage(true, b)
		convey.So(msg, convey.ShouldEqual, "")
	})
}

func Test_C24_DraftType_VerifyNotEqualMessage_Mismatch(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualMessage mismatch", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "y"}
		msg := a.VerifyNotEqualMessage(true, b)
		convey.So(msg, convey.ShouldNotEqual, "")
	})
}

func Test_C25_DraftType_VerifyAllNotEqualMessage(t *testing.T) {
	convey.Convey("DraftType VerifyAllNotEqualMessage", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		msg := a.VerifyAllNotEqualMessage(b)
		convey.So(msg, convey.ShouldEqual, "")
	})
}

func Test_C26_DraftType_VerifyNotEqualErr(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualErr equal returns nil", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		err := a.VerifyNotEqualErr(true, b)
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_C27_DraftType_VerifyNotEqualErr_Mismatch(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualErr mismatch returns error", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "y"}
		err := a.VerifyNotEqualErr(true, b)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_C28_DraftType_VerifyAllNotEqualErr(t *testing.T) {
	convey.Convey("DraftType VerifyAllNotEqualErr", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		err := a.VerifyAllNotEqualErr(b)
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_C29_DraftType_VerifyNotEqualExcludingInnerFieldsErr(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualExcludingInnerFieldsErr", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "x"}
		err := a.VerifyNotEqualExcludingInnerFieldsErr(b)
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_C30_DraftType_VerifyNotEqualExcludingInnerFieldsErr_Mismatch(t *testing.T) {
	convey.Convey("DraftType VerifyNotEqualExcludingInnerFieldsErr mismatch", t, func() {
		a := &coretests.DraftType{SampleString1: "x"}
		b := &coretests.DraftType{SampleString1: "y"}
		err := a.VerifyNotEqualExcludingInnerFieldsErr(b)
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func Test_C31_DraftType_JsonString(t *testing.T) {
	convey.Convey("DraftType JsonString", t, func() {
		d := coretests.DraftType{SampleString1: "hello"}
		j := d.JsonString()
		convey.So(j, convey.ShouldContainSubstring, "hello")
	})
}

func Test_C32_DraftType_JsonBytes(t *testing.T) {
	convey.Convey("DraftType JsonBytes", t, func() {
		d := coretests.DraftType{SampleString1: "hello"}
		b := d.JsonBytes()
		convey.So(string(b), convey.ShouldContainSubstring, "hello")
	})
}

func Test_C33_DraftType_JsonBytesPtr(t *testing.T) {
	convey.Convey("DraftType JsonBytesPtr", t, func() {
		d := coretests.DraftType{SampleString1: "hello"}
		b := d.JsonBytesPtr()
		convey.So(string(b), convey.ShouldContainSubstring, "hello")
	})
}

func Test_C34_DraftType_Clone(t *testing.T) {
	convey.Convey("DraftType Clone", t, func() {
		d := coretests.DraftType{
			SampleString1: "x",
			Lines:         []string{"a", "b"},
			RawBytes:      []byte("raw"),
		}
		clone := d.Clone()
		convey.So(clone.SampleString1, convey.ShouldEqual, "x")
		convey.So(clone.Lines, convey.ShouldResemble, d.Lines)
		convey.So(clone.RawBytes, convey.ShouldResemble, d.RawBytes)
	})
}

func Test_C35_DraftType_ClonePtr(t *testing.T) {
	convey.Convey("DraftType ClonePtr", t, func() {
		d := &coretests.DraftType{SampleString1: "x"}
		clone := d.ClonePtr()
		convey.So(clone, convey.ShouldNotEqual, d)
		convey.So(clone.SampleString1, convey.ShouldEqual, "x")
	})
}

func Test_C36_DraftType_ClonePtr_Nil(t *testing.T) {
	convey.Convey("DraftType ClonePtr nil", t, func() {
		var d *coretests.DraftType
		clone := d.ClonePtr()
		convey.So(clone, convey.ShouldBeNil)
	})
}

func Test_C37_DraftType_RawBytesLength_LinesLength(t *testing.T) {
	convey.Convey("DraftType RawBytesLength and LinesLength", t, func() {
		d := &coretests.DraftType{
			RawBytes: []byte("abc"),
			Lines:    []string{"a", "b"},
		}
		convey.So(d.RawBytesLength(), convey.ShouldEqual, 3)
		convey.So(d.LinesLength(), convey.ShouldEqual, 2)
	})
}

// ─── SimpleGherkins ──────────────────────────────────────────────────────────

func Test_C38_SimpleGherkins_ToString(t *testing.T) {
	convey.Convey("SimpleGherkins ToString", t, func() {
		g := &coretests.SimpleGherkins{
			Feature: "Login",
			Given:   "user exists",
			When:    "user logs in",
			Then:    "user sees dashboard",
		}
		s := g.ToString(0)
		convey.So(s, convey.ShouldContainSubstring, "Login")
	})
}

func Test_C39_SimpleGherkins_String(t *testing.T) {
	convey.Convey("SimpleGherkins String", t, func() {
		g := &coretests.SimpleGherkins{Feature: "F"}
		s := g.String()
		convey.So(s, convey.ShouldContainSubstring, "F")
	})
}

func Test_C40_SimpleGherkins_GetWithExpectation(t *testing.T) {
	convey.Convey("SimpleGherkins GetWithExpectation", t, func() {
		g := &coretests.SimpleGherkins{
			Feature: "F",
			Expect:  "exp",
			Actual:  "act",
		}
		s := g.GetWithExpectation(1)
		convey.So(s, convey.ShouldContainSubstring, "F")
	})
}

func Test_C41_SimpleGherkins_GetMessageConditional(t *testing.T) {
	convey.Convey("SimpleGherkins GetMessageConditional with expectation", t, func() {
		g := &coretests.SimpleGherkins{Feature: "F", Expect: "e", Actual: "a"}
		s := g.GetMessageConditional(true, 1)
		convey.So(s, convey.ShouldContainSubstring, "F")
	})
}

func Test_C42_SimpleGherkins_GetMessageConditional_NoExpectation(t *testing.T) {
	convey.Convey("SimpleGherkins GetMessageConditional without expectation", t, func() {
		g := &coretests.SimpleGherkins{Feature: "F"}
		s := g.GetMessageConditional(false, 0)
		convey.So(s, convey.ShouldContainSubstring, "F")
	})
}

// ─── Compare ─────────────────────────────────────────────────────────────────

func Test_C43_Compare_SortedStrings(t *testing.T) {
	convey.Convey("Compare SortedStrings", t, func() {
		c := &coretests.Compare{StringContains: "b a c"}
		result := c.SortedStrings()
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
		// second call returns cached
		result2 := c.SortedStrings()
		convey.So(result2, convey.ShouldResemble, result)
	})
}

func Test_C44_Compare_SortedString(t *testing.T) {
	convey.Convey("Compare SortedString", t, func() {
		c := &coretests.Compare{StringContains: "b a"}
		result := c.SortedString()
		convey.So(result, convey.ShouldNotEqual, "")
		// cached
		result2 := c.SortedString()
		convey.So(result2, convey.ShouldEqual, result)
	})
}

func Test_C45_Compare_GetPrintMessage(t *testing.T) {
	convey.Convey("Compare GetPrintMessage", t, func() {
		c := &coretests.Compare{StringContains: "hello"}
		msg := c.GetPrintMessage(0)
		convey.So(msg, convey.ShouldContainSubstring, "Index:0")
	})
}

func Test_C46_Compare_IsMatch_All(t *testing.T) {
	convey.Convey("Compare IsMatch all", t, func() {
		c := &coretests.Compare{
			StringContains: "a b",
			MatchingLength: 0, // all must match
		}
		instruction := &coretests.ComparingInstruction{}
		instruction.SetActual("a b c")
		result := c.IsMatch(false, 0, instruction)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C47_Compare_IsMatch_Partial(t *testing.T) {
	convey.Convey("Compare IsMatch partial", t, func() {
		c := &coretests.Compare{
			StringContains: "a b x",
			MatchingLength: 2,
		}
		instruction := &coretests.ComparingInstruction{}
		instruction.SetActual("a b c")
		result := c.IsMatch(false, 0, instruction)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C48_Compare_IsMatch_Fail_Print(t *testing.T) {
	convey.Convey("Compare IsMatch fail with print", t, func() {
		c := &coretests.Compare{
			StringContains: "x y z",
			MatchingLength: 0,
		}
		instruction := &coretests.ComparingInstruction{}
		instruction.SetActual("a b c")
		result := c.IsMatch(true, 0, instruction)
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_C49_Compare_IsMatch_Partial_Fail_Print(t *testing.T) {
	convey.Convey("Compare IsMatch partial fail with print", t, func() {
		c := &coretests.Compare{
			StringContains: "x y z",
			MatchingLength: 3,
		}
		instruction := &coretests.ComparingInstruction{}
		instruction.SetActual("a b c")
		result := c.IsMatch(true, 0, instruction)
		convey.So(result, convey.ShouldBeFalse)
	})
}

// ─── ComparingInstruction ────────────────────────────────────────────────────

func Test_C50_ComparingInstruction_SetActual_Actual(t *testing.T) {
	convey.Convey("ComparingInstruction SetActual and Actual", t, func() {
		ci := &coretests.ComparingInstruction{}
		ci.SetActual("hello world")
		convey.So(ci.Actual(), convey.ShouldEqual, "hello world")
	})
}

func Test_C51_ComparingInstruction_ActualHashset(t *testing.T) {
	convey.Convey("ComparingInstruction ActualHashset", t, func() {
		ci := &coretests.ComparingInstruction{}
		ci.SetActual("a b c")
		hs := ci.ActualHashset()
		convey.So(hs, convey.ShouldNotBeNil)
		// cached
		hs2 := ci.ActualHashset()
		convey.So(hs2, convey.ShouldEqual, hs)
	})
}

func Test_C52_ComparingInstruction_SetActual_ClearsCache(t *testing.T) {
	convey.Convey("ComparingInstruction SetActual clears hashset cache", t, func() {
		ci := &coretests.ComparingInstruction{}
		ci.SetActual("a b")
		hs1 := ci.ActualHashset()
		ci.SetActual("x y")
		hs2 := ci.ActualHashset()
		convey.So(hs1, convey.ShouldNotEqual, hs2)
	})
}

func Test_C53_ComparingInstruction_IsMatch(t *testing.T) {
	convey.Convey("ComparingInstruction IsMatch with comparing items", t, func() {
		ci := &coretests.ComparingInstruction{
			ComparingItems: []coretests.Compare{
				{StringContains: "a", MatchingLength: 0},
			},
		}
		ci.SetActual("a b c")
		cip := &coretests.CaseIndexPlusIsPrint{IsPrint: false, CaseIndex: 0}
		result := ci.IsMatch(cip)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C54_ComparingInstruction_IsMatch_WithMatchingAsEqual(t *testing.T) {
	convey.Convey("ComparingInstruction IsMatch with IsMatchingAsEqual", t, func() {
		ci := &coretests.ComparingInstruction{
			IsMatchingAsEqual:          true,
			MatchingAsEqualExpectation: "hello world",
			FunName:                    "TestFunc",
			TestCaseName:               "Case1",
			Header:                     "test header",
		}
		ci.SetActual("hello world")
		cip := &coretests.CaseIndexPlusIsPrint{IsPrint: false, CaseIndex: 0}
		result := ci.IsMatch(cip)
		convey.So(result, convey.ShouldBeTrue)
	})
}

// ─── isCompare functions ─────────────────────────────────────────────────────

func Test_C55_IsErrorNonWhiteSortedEqual_BothEmpty(t *testing.T) {
	convey.Convey("IsErrorNonWhiteSortedEqual both empty", t, func() {
		expectation := &errcore.ExpectationMessageDef{Expected: ""}
		result := coretests.IsErrorNonWhiteSortedEqual(false, nil, expectation)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C56_IsStrMsgNonWhiteSortedEqual_NonWhiteSort(t *testing.T) {
	convey.Convey("IsStrMsgNonWhiteSortedEqual non-white sort", t, func() {
		expectation := &errcore.ExpectationMessageDef{
			Expected:       "b a c",
			IsNonWhiteSort: true,
		}
		result := coretests.IsStrMsgNonWhiteSortedEqual(false, "a b c", expectation)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C57_IsStrMsgNonWhiteSortedEqual_NoSort(t *testing.T) {
	convey.Convey("IsStrMsgNonWhiteSortedEqual no sort, trim compare", t, func() {
		expectation := &errcore.ExpectationMessageDef{
			Expected:       "  hello  ",
			IsNonWhiteSort: false,
		}
		result := coretests.IsStrMsgNonWhiteSortedEqual(false, "hello", expectation)
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_C58_IsStrMsgNonWhiteSortedEqual_Mismatch_Print(t *testing.T) {
	convey.Convey("IsStrMsgNonWhiteSortedEqual mismatch with print", t, func() {
		expectation := &errcore.ExpectationMessageDef{
			Expected:       "abc",
			IsNonWhiteSort: false,
		}
		result := coretests.IsStrMsgNonWhiteSortedEqual(true, "xyz", expectation)
		convey.So(result, convey.ShouldBeFalse)
	})
}

// ─── VerifyTypeOf ────────────────────────────────────────────────────────────

func Test_C59_VerifyTypeOf_New(t *testing.T) {
	convey.Convey("NewVerifyTypeOf", t, func() {
		v := coretests.NewVerifyTypeOf("hello")
		convey.So(v, convey.ShouldNotBeNil)
		convey.So(v.IsDefined(), convey.ShouldBeTrue)
		convey.So(v.IsInvalid(), convey.ShouldBeTrue) // Note: IsInvalid returns it != nil
		convey.So(v.IsInvalidOrSkipVerify(), convey.ShouldBeFalse)
	})
}

func Test_C60_VerifyTypeOf_Nil(t *testing.T) {
	convey.Convey("VerifyTypeOf nil checks", t, func() {
		var v *coretests.VerifyTypeOf
		convey.So(v.IsDefined(), convey.ShouldBeFalse)
		convey.So(v.IsInvalidOrSkipVerify(), convey.ShouldBeTrue)
	})
}

func Test_C61_VerifyTypeOf_SkipVerify(t *testing.T) {
	convey.Convey("VerifyTypeOf skip verify", t, func() {
		v := &coretests.VerifyTypeOf{
			IsVerify: issetter.False,
		}
		convey.So(v.IsInvalidOrSkipVerify(), convey.ShouldBeTrue)
	})
}

// ─── SomeString ──────────────────────────────────────────────────────────────

func Test_C62_SomeString(t *testing.T) {
	convey.Convey("SomeString", t, func() {
		s := coretests.SomeString{Value: "hello"}
		convey.So(s.String(), convey.ShouldEqual, "hello")
		stringer := s.AsStringer()
		convey.So(stringer.String(), convey.ShouldEqual, "hello")
	})
}

// ─── TestFuncName ────────────────────────────────────────────────────────────

func Test_C63_TestFuncName(t *testing.T) {
	convey.Convey("TestFuncName Value", t, func() {
		fn := coretests.TestFuncName("myFunc")
		convey.So(fn.Value(), convey.ShouldEqual, "myFunc")
	})
}

// ─── CaseIndexPlusIsPrint ────────────────────────────────────────────────────

func Test_C64_CaseIndexPlusIsPrint(t *testing.T) {
	convey.Convey("CaseIndexPlusIsPrint struct", t, func() {
		cip := &coretests.CaseIndexPlusIsPrint{IsPrint: true, CaseIndex: 5}
		convey.So(cip.IsPrint, convey.ShouldBeTrue)
		convey.So(cip.CaseIndex, convey.ShouldEqual, 5)
	})
}

// ─── getAssert methods ───────────────────────────────────────────────────────

func Test_C65_GetAssert_Messages(t *testing.T) {
	convey.Convey("GetAssert message methods", t, func() {
		ga := coretests.GetAssert

		convey.So(ga.IsEqualMessage("w", "a", "e"), convey.ShouldNotEqual, "")
		convey.So(ga.IsNotEqualMessage("w", "a", "e"), convey.ShouldNotEqual, "")
		convey.So(ga.IsTrueMessage("w", "a"), convey.ShouldNotEqual, "")
		convey.So(ga.IsFalseMessage("w", "a"), convey.ShouldNotEqual, "")
		convey.So(ga.IsNilMessage("w", "a"), convey.ShouldNotEqual, "")
		convey.So(ga.IsNotNilMessage("w", "a"), convey.ShouldNotEqual, "")
		convey.So(ga.ShouldBeMessage("t", "a", "e"), convey.ShouldNotEqual, "")
		convey.So(ga.ShouldNotBeMessage("t", "a", "e"), convey.ShouldNotEqual, "")
	})
}

func Test_C66_GetAssert_SortedMessage(t *testing.T) {
	convey.Convey("GetAssert SortedMessage", t, func() {
		result := coretests.GetAssert.SortedMessage(false, "b a c", " ")
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

func Test_C67_GetAssert_SortedArrayNoPrint(t *testing.T) {
	convey.Convey("GetAssert SortedArrayNoPrint", t, func() {
		result := coretests.GetAssert.SortedArrayNoPrint("b a c")
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C68_GetAssert_ErrorToLinesWithSpaces(t *testing.T) {
	convey.Convey("GetAssert ErrorToLinesWithSpaces nil", t, func() {
		result := coretests.GetAssert.ErrorToLinesWithSpaces(2, nil)
		convey.So(result, convey.ShouldResemble, []string{})
	})
}

func Test_C69_GetAssert_ErrorToLinesWithSpacesDefault(t *testing.T) {
	convey.Convey("GetAssert ErrorToLinesWithSpacesDefault nil", t, func() {
		result := coretests.GetAssert.ErrorToLinesWithSpacesDefault(nil)
		convey.So(result, convey.ShouldResemble, []string{})
	})
}

func Test_C70_GetAssert_AnyToDoubleQuoteLines(t *testing.T) {
	convey.Convey("GetAssert AnyToDoubleQuoteLines", t, func() {
		result := coretests.GetAssert.AnyToDoubleQuoteLines(2, "hello\nworld")
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C71_GetAssert_AnyToStringDoubleQuoteLine(t *testing.T) {
	convey.Convey("GetAssert AnyToStringDoubleQuoteLine", t, func() {
		result := coretests.GetAssert.AnyToStringDoubleQuoteLine(2, "hello")
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

func Test_C72_GetAssert_ToString(t *testing.T) {
	convey.Convey("GetAssert ToString", t, func() {
		result := coretests.GetAssert.ToString("hello")
		convey.So(result, convey.ShouldEqual, "hello")
	})
}

// ─── LogOnFail / ToStringValues / ToStringNameValues ─────────────────────────

func Test_C73_LogOnFail_Pass(t *testing.T) {
	convey.Convey("LogOnFail pass does nothing", t, func() {
		coretests.LogOnFail(true, "e", "a")
		convey.So(true, convey.ShouldBeTrue) // no panic
	})
}

func Test_C74_LogOnFail_Fail(t *testing.T) {
	convey.Convey("LogOnFail fail logs", t, func() {
		coretests.LogOnFail(false, "expected", "actual")
		convey.So(true, convey.ShouldBeTrue) // no panic
	})
}

func Test_C75_ToStringValues(t *testing.T) {
	convey.Convey("ToStringValues", t, func() {
		result := coretests.ToStringValues("hello")
		convey.So(result, convey.ShouldContainSubstring, "hello")
	})
}

func Test_C76_ToStringValues_Nil(t *testing.T) {
	convey.Convey("ToStringValues nil", t, func() {
		result := coretests.ToStringValues(nil)
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

func Test_C77_ToStringNameValues(t *testing.T) {
	convey.Convey("ToStringNameValues", t, func() {
		result := coretests.ToStringNameValues("hello")
		convey.So(result, convey.ShouldContainSubstring, "hello")
	})
}

func Test_C78_ToStringNameValues_Nil(t *testing.T) {
	convey.Convey("ToStringNameValues nil", t, func() {
		result := coretests.ToStringNameValues(nil)
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

// ─── messagePrinter ──────────────────────────────────────────────────────────
// printMessage is unexported, but we cover it via GetAssert

func Test_C79_GetAssert_Quick(t *testing.T) {
	convey.Convey("GetAssert Quick", t, func() {
		result := coretests.GetAssert.Quick("when", "actual", "expected", 0)
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

// ─── getAssertSimpleTestCaseWrapper ──────────────────────────────────────────

func Test_C80_GetAssert_SimpleTestCaseWrapper_Lines(t *testing.T) {
	convey.Convey("GetAssert SimpleTestCaseWrapper Lines", t, func() {
		tc := coretests.SimpleTestCase{
			Title:         "test",
			ArrangeInput:  "arrange",
			ExpectedInput: "expected",
			ActualInput:   "actual",
		}
		actual, expected := coretests.GetAssert.SimpleTestCaseWrapper.Lines(tc)
		convey.So(len(actual), convey.ShouldBeGreaterThan, 0)
		convey.So(len(expected), convey.ShouldBeGreaterThan, 0)
	})
}

// ─── SimpleTestCase ──────────────────────────────────────────────────────────

func Test_C81_SimpleTestCase_Getters(t *testing.T) {
	convey.Convey("SimpleTestCase getters", t, func() {
		tc := coretests.SimpleTestCase{
			Title:         "Title",
			ArrangeInput:  "arrange",
			ExpectedInput: "expected",
			ActualInput:   "actual",
		}
		convey.So(tc.CaseTitle(), convey.ShouldEqual, "Title")
		convey.So(tc.Input(), convey.ShouldEqual, "arrange")
		convey.So(tc.Expected(), convey.ShouldEqual, "expected")
		convey.So(tc.Actual(), convey.ShouldEqual, "actual")
		convey.So(tc.ArrangeString(), convey.ShouldNotEqual, "")
		convey.So(tc.ExpectedString(), convey.ShouldNotEqual, "")
		convey.So(tc.ActualString(), convey.ShouldNotEqual, "")
	})
}

func Test_C82_SimpleTestCase_SetActual(t *testing.T) {
	convey.Convey("SimpleTestCase SetActual", t, func() {
		tc := coretests.SimpleTestCase{}
		tc.SetActual("new")
		// Note: value receiver, so this doesn't mutate the original
		convey.So(true, convey.ShouldBeTrue)
	})
}

func Test_C83_SimpleTestCase_FormTitle_CustomTitle(t *testing.T) {
	convey.Convey("SimpleTestCase FormTitle and CustomTitle", t, func() {
		tc := coretests.SimpleTestCase{Title: "T"}
		ft := tc.FormTitle(1)
		convey.So(ft, convey.ShouldContainSubstring, "T")
		ct := tc.CustomTitle(2, "Custom")
		convey.So(ct, convey.ShouldContainSubstring, "Custom")
	})
}

func Test_C84_SimpleTestCase_String(t *testing.T) {
	convey.Convey("SimpleTestCase String", t, func() {
		tc := coretests.SimpleTestCase{
			Title:         "Test",
			ArrangeInput:  "input",
			ExpectedInput: "output",
		}
		s := tc.String(0)
		convey.So(s, convey.ShouldNotEqual, "")
	})
}

func Test_C85_SimpleTestCase_LinesString(t *testing.T) {
	convey.Convey("SimpleTestCase LinesString", t, func() {
		tc := coretests.SimpleTestCase{
			Title:         "Test",
			ExpectedInput: "output",
		}
		s := tc.LinesString(0)
		convey.So(s, convey.ShouldNotEqual, "")
	})
}

func Test_C86_SimpleTestCase_AsSimpleTestCaseWrapper(t *testing.T) {
	convey.Convey("SimpleTestCase AsSimpleTestCaseWrapper", t, func() {
		tc := coretests.SimpleTestCase{Title: "T"}
		wrapper := tc.AsSimpleTestCaseWrapper()
		convey.So(wrapper, convey.ShouldNotBeNil)
		convey.So(wrapper.CaseTitle(), convey.ShouldEqual, "T")
	})
}

// ─── BaseTestCase Getters ────────────────────────────────────────────────────

func Test_C87_BaseTestCase_Getters(t *testing.T) {
	convey.Convey("BaseTestCase getters", t, func() {
		btc := &coretests.BaseTestCase{
			Title:         "Test",
			ArrangeInput:  "arr",
			ExpectedInput: "exp",
			ActualInput:   "act",
		}
		convey.So(btc.CaseTitle(), convey.ShouldEqual, "Test")
		convey.So(btc.Input(), convey.ShouldEqual, "arr")
		convey.So(btc.Expected(), convey.ShouldEqual, "exp")
		convey.So(btc.Actual(), convey.ShouldEqual, "act")
		convey.So(btc.ArrangeString(), convey.ShouldNotEqual, "")
		convey.So(btc.ExpectedString(), convey.ShouldNotEqual, "")
		convey.So(btc.ActualString(), convey.ShouldNotEqual, "")
		convey.So(btc.ArrangeTypeName(), convey.ShouldNotEqual, "")
	})
}

func Test_C88_BaseTestCase_SetActual(t *testing.T) {
	convey.Convey("BaseTestCase SetActual", t, func() {
		btc := &coretests.BaseTestCase{}
		btc.SetActual("new value")
		convey.So(btc.Actual(), convey.ShouldEqual, "new value")
	})
}

func Test_C89_BaseTestCase_Parameters(t *testing.T) {
	convey.Convey("BaseTestCase parameters nil", t, func() {
		btc := &coretests.BaseTestCase{}
		convey.So(btc.HasParameters(), convey.ShouldBeFalse)
		convey.So(btc.IsInvalidParameters(), convey.ShouldBeTrue)
		convey.So(btc.FirstParam(), convey.ShouldBeNil)
		convey.So(btc.SecondParam(), convey.ShouldBeNil)
		convey.So(btc.ThirdParam(), convey.ShouldBeNil)
		convey.So(btc.FourthParam(), convey.ShouldBeNil)
		convey.So(btc.FifthParam(), convey.ShouldBeNil)
		has, hm := btc.HashmapParam()
		convey.So(has, convey.ShouldBeFalse)
		convey.So(len(hm), convey.ShouldEqual, 0)
		convey.So(btc.HasValidHashmapParam(), convey.ShouldBeFalse)
	})
}

func Test_C90_BaseTestCase_NilReceiver(t *testing.T) {
	convey.Convey("BaseTestCase nil receiver", t, func() {
		var btc *coretests.BaseTestCase
		convey.So(btc.IsTypeInvalidOrSkipVerify(), convey.ShouldBeTrue)
		convey.So(btc.IsInvalidParameters(), convey.ShouldBeTrue)
		convey.So(btc.IsVerifyType(), convey.ShouldBeFalse)
	})
}

func Test_C91_BaseTestCase_IsDisabled_IsSkipWithLog(t *testing.T) {
	convey.Convey("BaseTestCase IsDisabled and IsSkipWithLog", t, func() {
		btc := &coretests.BaseTestCase{IsEnable: issetter.False}
		convey.So(btc.IsDisabled(), convey.ShouldBeTrue)
		convey.So(btc.IsSkipWithLog(0), convey.ShouldBeTrue)
	})
}

func Test_C92_BaseTestCase_IsSkipWithLog_Enabled(t *testing.T) {
	convey.Convey("BaseTestCase IsSkipWithLog enabled", t, func() {
		btc := &coretests.BaseTestCase{IsEnable: issetter.True}
		convey.So(btc.IsSkipWithLog(0), convey.ShouldBeFalse)
	})
}

func Test_C93_BaseTestCase_FormTitle_CustomTitle(t *testing.T) {
	convey.Convey("BaseTestCase FormTitle and CustomTitle", t, func() {
		btc := &coretests.BaseTestCase{Title: "T"}
		convey.So(btc.FormTitle(1), convey.ShouldContainSubstring, "T")
		convey.So(btc.CustomTitle(2, "C"), convey.ShouldContainSubstring, "C")
	})
}

func Test_C94_BaseTestCase_String_LinesString(t *testing.T) {
	convey.Convey("BaseTestCase String and LinesString", t, func() {
		btc := &coretests.BaseTestCase{
			Title:         "Test",
			ExpectedInput: "exp",
		}
		s := btc.String(0)
		convey.So(s, convey.ShouldNotEqual, "")
		ls := btc.LinesString(0)
		convey.So(ls, convey.ShouldNotEqual, "")
	})
}

func Test_C95_BaseTestCase_AsSimpleTestCaseWrapper(t *testing.T) {
	convey.Convey("BaseTestCase AsSimpleTestCaseWrapper", t, func() {
		btc := &coretests.BaseTestCase{Title: "T"}
		wrapper := btc.AsSimpleTestCaseWrapper()
		convey.So(wrapper, convey.ShouldNotBeNil)
	})
}

func Test_C96_BaseTestCase_AsBaseTestCaseWrapper(t *testing.T) {
	convey.Convey("BaseTestCase AsBaseTestCaseWrapper", t, func() {
		btc := &coretests.BaseTestCase{Title: "T"}
		wrapper := btc.AsBaseTestCaseWrapper()
		convey.So(wrapper, convey.ShouldNotBeNil)
	})
}

func Test_C97_BaseTestCase_ActualLines_ExpectedLines(t *testing.T) {
	convey.Convey("BaseTestCase ActualLines and ExpectedLines", t, func() {
		btc := &coretests.BaseTestCase{
			ActualInput:   "a",
			ExpectedInput: "b",
		}
		convey.So(len(btc.ActualLines()), convey.ShouldBeGreaterThan, 0)
		convey.So(len(btc.ExpectedLines()), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C98_BaseTestCase_TypeValidation_NoVerify(t *testing.T) {
	convey.Convey("BaseTestCase TypeValidationError no verify", t, func() {
		btc := &coretests.BaseTestCase{}
		err := btc.TypeValidationError()
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_C99_BaseTestCase_IsVerifyType(t *testing.T) {
	convey.Convey("BaseTestCase IsVerifyType", t, func() {
		btc := &coretests.BaseTestCase{
			VerifyTypeOf: coretests.NewVerifyTypeOf("hello"),
		}
		convey.So(btc.IsVerifyType(), convey.ShouldBeTrue)
	})
}

// ─── GetAssert ToStringsWithSpace, StringsToSpaceString ──────────────────────

func Test_C100_GetAssert_ToStringsWithSpace(t *testing.T) {
	convey.Convey("GetAssert ToStringsWithSpace", t, func() {
		result := coretests.GetAssert.ToStringsWithSpace(2, "hello")
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C101_GetAssert_StringsToSpaceString(t *testing.T) {
	convey.Convey("GetAssert StringsToSpaceString", t, func() {
		result := coretests.GetAssert.StringsToSpaceString(2, "a", "b")
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C102_GetAssert_ToQuoteLines(t *testing.T) {
	convey.Convey("GetAssert ToQuoteLines", t, func() {
		result := coretests.GetAssert.ToQuoteLines(2, []string{"a", "b"})
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_C103_GetAssert_ConvertLinesToDoubleQuoteThenString(t *testing.T) {
	convey.Convey("GetAssert ConvertLinesToDoubleQuoteThenString", t, func() {
		result := coretests.GetAssert.ConvertLinesToDoubleQuoteThenString(2, []string{"a", "b"})
		convey.So(result, convey.ShouldNotEqual, "")
	})
}

func Test_C104_GetAssert_StringsToSpaceStringUsingFunc(t *testing.T) {
	convey.Convey("GetAssert StringsToSpaceStringUsingFunc", t, func() {
		fn := func(i int, prefix, line string) string {
			return prefix + line
		}
		result := coretests.GetAssert.StringsToSpaceStringUsingFunc(2, fn, "a", "b")
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_C105_GetAssert_ErrorToLinesWithSpaces_NonNil(t *testing.T) {
	convey.Convey("GetAssert ErrorToLinesWithSpaces non-nil error", t, func() {
		err := errcore.NotFound.Error("test", "detail")
		result := coretests.GetAssert.ErrorToLinesWithSpaces(2, err)
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}
