package integratedtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── AnyToBytes ──

func Test_Cov_AnyToBytes_FromBytes(t *testing.T) {
	result := coretests.AnyToBytes([]byte("hello"))
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyToBytes from bytes", actual)
}

func Test_Cov_AnyToBytes_FromNilBytes(t *testing.T) {
	result := coretests.AnyToBytes([]byte(nil))
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyToBytes from nil bytes", actual)
}

func Test_Cov_AnyToBytes_FromString(t *testing.T) {
	result := coretests.AnyToBytes("hello")
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyToBytes from string", actual)
}

func Test_Cov_AnyToBytes_FromStruct(t *testing.T) {
	result := coretests.AnyToBytes(struct{ N int }{42})
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToBytes from struct", actual)
}

// ── AnyToBytesPtr (deprecated alias) ──

func Test_Cov_AnyToBytesPtr(t *testing.T) {
	result := coretests.AnyToBytesPtr("test")
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "AnyToBytesPtr", actual)
}

// ── AnyToDraftType ──

func Test_Cov_AnyToDraftType_FromValue(t *testing.T) {
	dt := coretests.DraftType{SampleString1: "hello"}
	result := coretests.AnyToDraftType(dt)
	actual := args.Map{"notNil": result != nil, "val": result.SampleString1}
	expected := args.Map{"notNil": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyToDraftType from value", actual)
}

func Test_Cov_AnyToDraftType_FromPtr(t *testing.T) {
	dt := &coretests.DraftType{SampleString1: "hello"}
	result := coretests.AnyToDraftType(dt)
	actual := args.Map{"notNil": result != nil, "val": result.SampleString1}
	expected := args.Map{"notNil": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyToDraftType from ptr", actual)
}

func Test_Cov_AnyToDraftType_FromOther(t *testing.T) {
	result := coretests.AnyToDraftType("not a draft type")
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyToDraftType from other", actual)
}

// ── DraftType ──

func Test_Cov_DraftType_Getters(t *testing.T) {
	dt := coretests.DraftType{
		SampleString1: "s1",
		SampleString2: "s2",
		SampleInteger: 42,
		Lines:         []string{"a"},
		RawBytes:      []byte("b"),
	}
	dt.SetF2Integer(99)

	actual := args.Map{
		"f1":       dt.F1String(),
		"f2":       dt.F2Integer(),
		"linesLen": dt.LinesLength(),
		"bytesLen": dt.RawBytesLength(),
	}
	expected := args.Map{
		"f1":       "",
		"f2":       99,
		"linesLen": 1,
		"bytesLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "DraftType getters", actual)
}

func Test_Cov_DraftType_JsonString(t *testing.T) {
	dt := coretests.DraftType{SampleString1: "test"}
	actual := args.Map{
		"jsonNotEmpty":  dt.JsonString() != "",
		"bytesNotEmpty": len(dt.JsonBytes()) > 0,
		"ptrNotEmpty":   len(dt.JsonBytesPtr()) > 0,
	}
	expected := args.Map{
		"jsonNotEmpty":  true,
		"bytesNotEmpty": true,
		"ptrNotEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType json", actual)
}

func Test_Cov_DraftType_NonPtr(t *testing.T) {
	dt := coretests.DraftType{SampleString1: "test"}
	nonPtr := dt.NonPtr()
	actual := args.Map{"val": nonPtr.SampleString1}
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "DraftType NonPtr", actual)
}

func Test_Cov_DraftType_PtrOrNonPtr(t *testing.T) {
	dt := &coretests.DraftType{SampleString1: "test"}
	ptrResult := dt.PtrOrNonPtr(true)
	nonPtrResult := dt.PtrOrNonPtr(false)
	var nilDt *coretests.DraftType
	nilResult := nilDt.PtrOrNonPtr(true)
	actual := args.Map{
		"ptrNotNil":    ptrResult != nil,
		"nonPtrNotNil": nonPtrResult != nil,
		"nilIsNil":     nilResult == nil,
	}
	expected := args.Map{
		"ptrNotNil":    true,
		"nonPtrNotNil": true,
		"nilIsNil":     true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType PtrOrNonPtr", actual)
}

func Test_Cov_DraftType_Clone(t *testing.T) {
	dt := coretests.DraftType{SampleString1: "test", Lines: []string{"a"}, RawBytes: []byte("b")}
	clone := dt.Clone()
	actual := args.Map{"val": clone.SampleString1, "linesLen": len(clone.Lines)}
	expected := args.Map{"val": "test", "linesLen": 1}
	expected.ShouldBeEqual(t, 0, "DraftType Clone", actual)
}

func Test_Cov_DraftType_ClonePtr(t *testing.T) {
	dt := &coretests.DraftType{SampleString1: "test"}
	clone := dt.ClonePtr()
	var nilDt *coretests.DraftType
	nilClone := nilDt.ClonePtr()
	actual := args.Map{"notNil": clone != nil, "nilIsNil": nilClone == nil}
	expected := args.Map{"notNil": true, "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "DraftType ClonePtr", actual)
}

func Test_Cov_DraftType_IsEqual(t *testing.T) {
	dt1 := &coretests.DraftType{SampleString1: "a"}
	dt2 := &coretests.DraftType{SampleString1: "a"}
	dt3 := &coretests.DraftType{SampleString1: "b"}
	actual := args.Map{
		"same":       dt1.IsEqual(true, dt2),
		"diff":       dt1.IsEqual(true, dt3),
		"bothNil":    (*coretests.DraftType)(nil).IsEqual(true, nil),
		"leftNil":    (*coretests.DraftType)(nil).IsEqual(true, dt1),
		"selfEqual":  dt1.IsEqual(true, dt1),
		"isEqualAll": dt1.IsEqualAll(dt2),
	}
	expected := args.Map{
		"same":       true,
		"diff":       false,
		"bothNil":    true,
		"leftNil":    false,
		"selfEqual":  true,
		"isEqualAll": true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType IsEqual", actual)
}

func Test_Cov_DraftType_Verify(t *testing.T) {
	dt1 := &coretests.DraftType{SampleString1: "a"}
	dt2 := &coretests.DraftType{SampleString1: "b"}
	dt3 := &coretests.DraftType{SampleString1: "a"}
	actual := args.Map{
		"diffMsg":        dt1.VerifyNotEqualMessage(false, dt2) != "",
		"sameMsg":        dt1.VerifyNotEqualMessage(false, dt3),
		"allMsg":         dt1.VerifyAllNotEqualMessage(dt2) != "",
		"diffErr":        dt1.VerifyNotEqualErr(false, dt2) != nil,
		"sameErr":        dt1.VerifyNotEqualErr(false, dt3) == nil,
		"allErr":         dt1.VerifyAllNotEqualErr(dt2) != nil,
		"exclInnerErr":   dt1.VerifyNotEqualExcludingInnerFieldsErr(dt2) != nil,
	}
	expected := args.Map{
		"diffMsg":        true,
		"sameMsg":        "",
		"allMsg":         true,
		"diffErr":        true,
		"sameErr":        true,
		"allErr":         true,
		"exclInnerErr":   true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType Verify", actual)
}

// ── TestFuncName ──

func Test_Cov_TestFuncName(t *testing.T) {
	fn := coretests.TestFuncName("myFunc")
	actual := args.Map{"val": fn.Value()}
	expected := args.Map{"val": "myFunc"}
	expected.ShouldBeEqual(t, 0, "TestFuncName", actual)
}

// ── SomeString ──

func Test_Cov_SomeString(t *testing.T) {
	s := coretests.SomeString{Value: "hello"}
	actual := args.Map{
		"str":       s.String(),
		"stringer":  s.AsStringer() != nil,
	}
	expected := args.Map{
		"str":       "hello",
		"stringer":  true,
	}
	expected.ShouldBeEqual(t, 0, "SomeString", actual)
}

// ── VerifyTypeOf ──

func Test_Cov_VerifyTypeOf(t *testing.T) {
	vt := coretests.NewVerifyTypeOf("hello")
	actual := args.Map{
		"isDefined":      vt.IsDefined(),
		"isInvalid":      vt.IsInvalid(),
		"isSkipVerify":   vt.IsInvalidOrSkipVerify(),
	}
	expected := args.Map{
		"isDefined":      true,
		"isInvalid":      true,
		"isSkipVerify":   false,
	}
	expected.ShouldBeEqual(t, 0, "VerifyTypeOf", actual)
}

func Test_Cov_VerifyTypeOf_Nil(t *testing.T) {
	var vt *coretests.VerifyTypeOf
	actual := args.Map{
		"isDefined":    vt.IsDefined(),
		"isSkipVerify": vt.IsInvalidOrSkipVerify(),
	}
	expected := args.Map{
		"isDefined":    false,
		"isSkipVerify": true,
	}
	expected.ShouldBeEqual(t, 0, "VerifyTypeOf nil", actual)
}

// ── LogOnFail ──

func Test_Cov_LogOnFail_Pass(t *testing.T) {
	// Should not panic
	coretests.LogOnFail(true, "expected", "actual")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogOnFail pass", actual)
}

func Test_Cov_LogOnFail_Fail(t *testing.T) {
	// Should log but not panic
	coretests.LogOnFail(false, "expected", "actual")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogOnFail fail", actual)
}

// ── ToStringValues / ToStringNameValues ──

func Test_Cov_ToStringValues(t *testing.T) {
	actual := args.Map{
		"val": coretests.ToStringValues(42) != "",
		"nil": coretests.ToStringValues(nil) != "",
	}
	expected := args.Map{
		"val": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringValues", actual)
}

func Test_Cov_ToStringNameValues(t *testing.T) {
	actual := args.Map{
		"val": coretests.ToStringNameValues(42) != "",
		"nil": coretests.ToStringNameValues(nil) != "",
	}
	expected := args.Map{
		"val": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "ToStringNameValues", actual)
}

// ── SimpleGherkins ──

func Test_Cov_SimpleGherkins_ToString(t *testing.T) {
	g := &coretests.SimpleGherkins{
		Feature: "Login",
		Given:   "user exists",
		When:    "user logs in",
		Then:    "redirect to home",
		Expect:  "home page",
		Actual:  "home page",
	}

	actual := args.Map{
		"toString":      g.ToString(0) != "",
		"string":        g.String() != "",
		"withExpect":    g.GetWithExpectation(0) != "",
		"condTrue":      g.GetMessageConditional(true, 0) != "",
		"condFalse":     g.GetMessageConditional(false, 0) != "",
	}
	expected := args.Map{
		"toString":      true,
		"string":        true,
		"withExpect":    true,
		"condTrue":      true,
		"condFalse":     true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleGherkins", actual)
}

// ── BaseTestCase ──

func Test_Cov_BaseTestCase_Getters(t *testing.T) {
	bt := &coretests.BaseTestCase{
		Title:         "test",
		ArrangeInput:  "input",
		ExpectedInput: "expected",
	}

	actual := args.Map{
		"title":          bt.CaseTitle(),
		"typeName":       bt.ArrangeTypeName(),
		"arrangeStr":     bt.ArrangeString() != "",
		"isTypeInvalid":  bt.IsTypeInvalidOrSkipVerify(),
		"hasParams":      bt.HasParameters(),
		"invalidParams":  bt.IsInvalidParameters(),
		"firstParam":     bt.FirstParam() == nil,
		"secondParam":    bt.SecondParam() == nil,
		"thirdParam":     bt.ThirdParam() == nil,
		"fourthParam":    bt.FourthParam() == nil,
		"fifthParam":     bt.FifthParam() == nil,
		"isVerify":       bt.IsVerifyType(),
		"wrapperNotNil":  bt.AsSimpleTestCaseWrapper() != nil,
		"baseWrapNotNil": bt.AsBaseTestCaseWrapper() != nil,
	}
	expected := args.Map{
		"title":          "test",
		"typeName":       "string",
		"arrangeStr":     true,
		"isTypeInvalid":  true,
		"hasParams":      false,
		"invalidParams":  true,
		"firstParam":     true,
		"secondParam":    true,
		"thirdParam":     true,
		"fourthParam":    true,
		"fifthParam":     true,
		"isVerify":       false,
		"wrapperNotNil":  true,
		"baseWrapNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase getters", actual)
}

func Test_Cov_BaseTestCase_WithParams(t *testing.T) {
	bt := &coretests.BaseTestCase{
		Title: "test",
		Parameters: &args.HolderAny{
			First:   "f1",
			Second:  "f2",
			Third:   "f3",
			Fourth:  "f4",
			Fifth:   "f5",
			Hashmap: map[string]any{"key": "val"},
		},
	}

	hasMap, hashMap := bt.HashmapParam()
	actual := args.Map{
		"hasParams":     bt.HasParameters(),
		"firstParam":    bt.FirstParam(),
		"secondParam":   bt.SecondParam(),
		"thirdParam":    bt.ThirdParam(),
		"fourthParam":   bt.FourthParam(),
		"fifthParam":    bt.FifthParam(),
		"hasMap":        hasMap,
		"hasValidMap":   bt.HasValidHashmapParam(),
		"hashMapLen":    len(hashMap),
	}
	expected := args.Map{
		"hasParams":     true,
		"firstParam":    "f1",
		"secondParam":   "f2",
		"thirdParam":    "f3",
		"fourthParam":   "f4",
		"fifthParam":    "f5",
		"hasMap":        true,
		"hasValidMap":   true,
		"hashMapLen":    1,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase with params", actual)
}

func Test_Cov_BaseTestCase_Nil(t *testing.T) {
	var bt *coretests.BaseTestCase
	actual := args.Map{
		"isTypeInvalid": bt.IsTypeInvalidOrSkipVerify(),
		"invalidParams": bt.IsInvalidParameters(),
	}
	expected := args.Map{
		"isTypeInvalid": true,
		"invalidParams": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase nil", actual)
}

// ── CaseIndexPlusIsPrint ──

func Test_Cov_CaseIndexPlusIsPrint(t *testing.T) {
	c := &coretests.CaseIndexPlusIsPrint{
		IsPrint:   true,
		CaseIndex: 5,
	}
	actual := args.Map{"isPrint": c.IsPrint, "index": c.CaseIndex}
	expected := args.Map{"isPrint": true, "index": 5}
	expected.ShouldBeEqual(t, 0, "CaseIndexPlusIsPrint", actual)
}

// ── SimpleTestCase ──

func Test_Cov_SimpleTestCase_Getters(t *testing.T) {
	tc := coretests.SimpleTestCase{
		Title:         "test",
		ArrangeInput:  "arrange",
		ExpectedInput: "expected",
	}
	tc.SetActual("actual")

	actual := args.Map{
		"title":       tc.CaseTitle(),
		"input":       tc.Input(),
		"expected":    tc.Expected(),
		"arrangeStr":  tc.ArrangeString() != "",
		"expectedStr": tc.ExpectedString() != "",
		"formTitle":   tc.FormTitle(0) != "",
		"customTitle": tc.CustomTitle(0, "custom") != "",
		"wrapperOk":   tc.AsSimpleTestCaseWrapper() != nil,
	}
	expected := args.Map{
		"title":       "test",
		"input":       "arrange",
		"expected":    "expected",
		"arrangeStr":  true,
		"expectedStr": true,
		"formTitle":   true,
		"customTitle": true,
		"wrapperOk":   true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleTestCase getters", actual)
}

// ── GetAssert ──

func Test_Cov_GetAssert_ToString(t *testing.T) {
	result := coretests.GetAssert.ToString("hello")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "GetAssert.ToString", actual)
}

func Test_Cov_GetAssert_SortedMessage(t *testing.T) {
	result := coretests.GetAssert.SortedMessage(false, "c b a", " ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetAssert.SortedMessage", actual)
}

func Test_Cov_GetAssert_SortedArrayNoPrint(t *testing.T) {
	result := coretests.GetAssert.SortedArrayNoPrint("c b a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "GetAssert.SortedArrayNoPrint", actual)
}

func Test_Cov_GetAssert_ToStrings(t *testing.T) {
	result := coretests.GetAssert.ToStrings("hello")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetAssert.ToStrings", actual)
}

func Test_Cov_GetAssert_ToStringsWithSpace(t *testing.T) {
	result := coretests.GetAssert.ToStringsWithSpace(2, "hello")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetAssert.ToStringsWithSpace", actual)
}

func Test_Cov_GetAssert_ErrorToLinesWithSpaces(t *testing.T) {
	nilResult := coretests.GetAssert.ErrorToLinesWithSpaces(2, nil)
	actual := args.Map{"nilLen": len(nilResult)}
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "GetAssert.ErrorToLinesWithSpaces nil", actual)
}

func Test_Cov_GetAssert_ErrorToLinesWithSpacesDefault(t *testing.T) {
	result := coretests.GetAssert.ErrorToLinesWithSpacesDefault(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetAssert.ErrorToLinesWithSpacesDefault", actual)
}

func Test_Cov_GetAssert_AnyToDoubleQuoteLines(t *testing.T) {
	result := coretests.GetAssert.AnyToDoubleQuoteLines(2, "hello")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetAssert.AnyToDoubleQuoteLines", actual)
}

func Test_Cov_GetAssert_ConvertLinesToDoubleQuoteThenString(t *testing.T) {
	result := coretests.GetAssert.ConvertLinesToDoubleQuoteThenString(2, []string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetAssert.ConvertLinesToDoubleQuoteThenString", actual)
}

func Test_Cov_GetAssert_AnyToStringDoubleQuoteLine(t *testing.T) {
	result := coretests.GetAssert.AnyToStringDoubleQuoteLine(2, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetAssert.AnyToStringDoubleQuoteLine", actual)
}

// ── Compare ──

func Test_Cov_Compare_SortedStrings(t *testing.T) {
	c := &coretests.Compare{StringContains: "c b a"}
	ss := c.SortedStrings()
	ss2 := c.SortedStrings() // cached

	actual := args.Map{
		"len":      len(ss),
		"cachedEq": len(ss) == len(ss2),
	}
	expected := args.Map{
		"len":      3,
		"cachedEq": true,
	}
	expected.ShouldBeEqual(t, 0, "Compare.SortedStrings", actual)
}

func Test_Cov_Compare_SortedString(t *testing.T) {
	c := &coretests.Compare{StringContains: "b a"}
	s := c.SortedString()
	s2 := c.SortedString() // cached

	actual := args.Map{
		"notEmpty": s != "",
		"cached":   s == s2,
	}
	expected := args.Map{
		"notEmpty": true,
		"cached":   true,
	}
	expected.ShouldBeEqual(t, 0, "Compare.SortedString", actual)
}

func Test_Cov_Compare_GetPrintMessage(t *testing.T) {
	c := &coretests.Compare{StringContains: "hello"}
	msg := c.GetPrintMessage(0)
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare.GetPrintMessage", actual)
}
