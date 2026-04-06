package coretesttests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DraftType ──

func Test_Cov_DraftType_Methods(t *testing.T) {
	d := &coretests.DraftType{SampleString1: "a", SampleString2: "b", SampleInteger: 1, Lines: []string{"l"}, RawBytes: []byte{1}}
	d.SetF2Integer(5)
	actual := args.Map{
		"f1":       d.F1String(),
		"f2":       d.F2Integer(),
		"nonPtr":   d.NonPtr().SampleString1,
		"linesLen": d.LinesLength(),
		"bytesLen": d.RawBytesLength(),
		"json":     d.JsonString() != "",
		"jsonByte": len(d.JsonBytes()) > 0,
		"jsonPtr":  len(d.JsonBytesPtr()) > 0,
	}
	expected := args.Map{
		"f1": "", "f2": 5, "nonPtr": "a",
		"linesLen": 1, "bytesLen": 1,
		"json": true, "jsonByte": true, "jsonPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType returns correct value -- Methods", actual)
}

func Test_Cov_DraftType_PtrOrNonPtr(t *testing.T) {
	d := &coretests.DraftType{SampleString1: "a"}
	actual := args.Map{
		"ptrNotNil":    d.PtrOrNonPtr(true) != nil,
		"nonPtrNotNil": d.PtrOrNonPtr(false) != nil,
	}
	var nilD *coretests.DraftType
	actual["nilResult"] = nilD.PtrOrNonPtr(true) == nil
	expected := args.Map{"ptrNotNil": true, "nonPtrNotNil": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "DraftType returns correct value -- PtrOrNonPtr", actual)
}

func Test_Cov_DraftType_Clone(t *testing.T) {
	d := &coretests.DraftType{SampleString1: "a", Lines: []string{"l"}, RawBytes: []byte{1}}
	c := d.Clone()
	cp := d.ClonePtr()
	var nilD *coretests.DraftType
	nilCp := nilD.ClonePtr()
	actual := args.Map{
		"name": c.SampleString1, "cpName": cp.SampleString1,
		"nilCp": nilCp == nil,
	}
	expected := args.Map{"name": "a", "cpName": "a", "nilCp": true}
	expected.ShouldBeEqual(t, 0, "DraftType returns correct value -- Clone", actual)
}

func Test_Cov_DraftType_IsEqual(t *testing.T) {
	d1 := &coretests.DraftType{SampleString1: "a", SampleInteger: 1}
	d2 := &coretests.DraftType{SampleString1: "a", SampleInteger: 1}
	d3 := &coretests.DraftType{SampleString1: "b"}
	var nilD *coretests.DraftType
	actual := args.Map{
		"equal":        d1.IsEqual(true, d2),
		"equalAll":     d1.IsEqualAll(d2),
		"notEqual":     d1.IsEqual(true, d3),
		"nilNil":       nilD.IsEqual(true, nilD),
		"nilNonNil":    nilD.IsEqual(true, d1),
		"self":         d1.IsEqual(true, d1),
		"excludeInner": d1.IsEqual(false, d2),
	}
	expected := args.Map{
		"equal": true, "equalAll": true, "notEqual": false,
		"nilNil": true, "nilNonNil": false, "self": true,
		"excludeInner": true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType returns correct value -- IsEqual", actual)
}

func Test_Cov_DraftType_VerifyNotEqual(t *testing.T) {
	d1 := &coretests.DraftType{SampleString1: "a"}
	d2 := &coretests.DraftType{SampleString1: "a"}
	d3 := &coretests.DraftType{SampleString1: "b"}
	actual := args.Map{
		"equalMsg":    d1.VerifyAllNotEqualMessage(d2),
		"notEqualMsg": d1.VerifyAllNotEqualMessage(d3) != "",
		"equalErr":    d1.VerifyAllNotEqualErr(d2) == nil,
		"notEqualErr": d1.VerifyAllNotEqualErr(d3) != nil,
		"excludeErr":  d1.VerifyNotEqualExcludingInnerFieldsErr(d2) == nil,
	}
	expected := args.Map{
		"equalMsg": "", "notEqualMsg": true,
		"equalErr": true, "notEqualErr": true,
		"excludeErr": true,
	}
	expected.ShouldBeEqual(t, 0, "DraftType returns correct value -- VerifyNotEqual", actual)
}

// ── AnyToBytes / AnyToBytesPtr / AnyToDraftType ──

func Test_Cov_AnyToBytes(t *testing.T) {
	bytesResult := coretests.AnyToBytes([]byte{1, 2})
	strResult := coretests.AnyToBytes("hello")
	intResult := coretests.AnyToBytes(42)
	nilResult := coretests.AnyToBytes([]byte(nil))
	actual := args.Map{
		"bytesLen": len(bytesResult), "strLen": len(strResult) > 0,
		"intLen": len(intResult) > 0, "nilIsNil": nilResult == nil,
	}
	expected := args.Map{
		"bytesLen": 2, "strLen": true, "intLen": true, "nilIsNil": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyToBytes returns correct value -- with args", actual)
}

func Test_Cov_AnyToBytesPtr(t *testing.T) {
	result := coretests.AnyToBytesPtr("hello")
	actual := args.Map{"len": len(result) > 0}
	expected := args.Map{"len": true}
	expected.ShouldBeEqual(t, 0, "AnyToBytesPtr returns correct value -- with args", actual)
}

func Test_Cov_AnyToDraftType(t *testing.T) {
	d := coretests.DraftType{SampleString1: "a"}
	dp := &coretests.DraftType{SampleString1: "b"}
	actual := args.Map{
		"fromVal":  coretests.AnyToDraftType(d).SampleString1,
		"fromPtr":  coretests.AnyToDraftType(dp).SampleString1,
		"fromNil":  coretests.AnyToDraftType("not draft") == nil,
	}
	expected := args.Map{"fromVal": "a", "fromPtr": "b", "fromNil": true}
	expected.ShouldBeEqual(t, 0, "AnyToDraftType returns correct value -- with args", actual)
}

// ── CaseIndexPlusIsPrint ──

func Test_Cov_CaseIndexPlusIsPrint(t *testing.T) {
	c := &coretests.CaseIndexPlusIsPrint{IsPrint: true, CaseIndex: 5}
	actual := args.Map{"isPrint": c.IsPrint, "index": c.CaseIndex}
	expected := args.Map{"isPrint": true, "index": 5}
	expected.ShouldBeEqual(t, 0, "CaseIndexPlusIsPrint returns correct value -- with args", actual)
}

// ── SomeString ──

func Test_Cov_SomeString(t *testing.T) {
	s := coretests.SomeString{Value: "hello"}
	actual := args.Map{
		"str":      s.String(),
		"stringer": s.AsStringer().String(),
	}
	expected := args.Map{"str": "hello", "stringer": "hello"}
	expected.ShouldBeEqual(t, 0, "SomeString returns correct value -- with args", actual)
}

// ── TestFuncName ──

func Test_Cov_TestFuncName(t *testing.T) {
	fn := coretests.TestFuncName("myFunc")
	actual := args.Map{"val": fn.Value()}
	expected := args.Map{"val": "myFunc"}
	expected.ShouldBeEqual(t, 0, "TestFuncName returns correct value -- with args", actual)
}

// ── Compare ──

func Test_Cov_Compare_SortedStrings(t *testing.T) {
	c := &coretests.Compare{StringContains: "hello world"}
	sorted := c.SortedStrings()
	sortedStr := c.SortedString()
	actual := args.Map{"len": len(sorted) > 0, "strNotEmpty": sortedStr != ""}
	expected := args.Map{"len": true, "strNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- SortedStrings", actual)
}

func Test_Cov_Compare_GetPrintMessage(t *testing.T) {
	c := &coretests.Compare{StringContains: "test"}
	msg := c.GetPrintMessage(0)
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- GetPrintMessage", actual)
}

// ── ComparingInstruction ──

func Test_Cov_ComparingInstruction(t *testing.T) {
	ci := &coretests.ComparingInstruction{
		FunName:      "fn",
		Header:       "header",
		TestCaseName: "test",
	}
	ci.SetActual("hello world")
	actual := args.Map{
		"actual":       ci.Actual(),
		"hashsetNotNil": ci.ActualHashset() != nil,
	}
	expected := args.Map{"actual": "hello world", "hashsetNotNil": true}
	expected.ShouldBeEqual(t, 0, "ComparingInstruction returns correct value -- with args", actual)
}

// ── BaseTestCase getters ──

func Test_Cov_BaseTestCase_Getters(t *testing.T) {
	bt := &coretests.BaseTestCase{
		Title:         "test",
		ArrangeInput:  "arrange",
		ActualInput:   "actual",
		ExpectedInput: "expected",
	}
	actual := args.Map{
		"title":        bt.CaseTitle(),
		"input":        bt.Input(),
		"expected":     bt.Expected(),
		"actual":       bt.Actual(),
		"arrangeStr":   bt.ArrangeString() != "",
		"actualStr":    bt.ActualString() != "",
		"expectedStr":  bt.ExpectedString() != "",
		"arrangeType":  bt.ArrangeTypeName(),
		"skipVerify":   bt.IsTypeInvalidOrSkipVerify(),
		"isVerify":     bt.IsVerifyType(),
		"hasPar":       bt.HasParameters(),
		"invalidPar":   bt.IsInvalidParameters(),
	}
	expected := args.Map{
		"title": "test", "input": "arrange",
		"expected": "expected", "actual": "actual",
		"arrangeStr": true, "actualStr": true, "expectedStr": true,
		"arrangeType": "string", "skipVerify": true, "isVerify": false,
		"hasPar": false, "invalidPar": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- Getters", actual)
}

func Test_Cov_BaseTestCase_Params(t *testing.T) {
	bt := &coretests.BaseTestCase{
		Parameters: &args.HolderAny{
			First: 1, Second: 2, Third: 3, Fourth: 4, Fifth: 5,
			Hashmap: map[string]any{"k": "v"},
		},
	}
	actual := args.Map{
		"first": bt.FirstParam(), "second": bt.SecondParam(),
		"third": bt.ThirdParam(), "fourth": bt.FourthParam(),
		"fifth": bt.FifthParam(),
		"hasHashmap": bt.HasValidHashmapParam(),
	}
	expected := args.Map{
		"first": 1, "second": 2, "third": 3, "fourth": 4, "fifth": 5,
		"hasHashmap": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- Params", actual)
}

func Test_Cov_BaseTestCase_NilParams(t *testing.T) {
	bt := &coretests.BaseTestCase{}
	actual := args.Map{
		"first": bt.FirstParam() == nil, "second": bt.SecondParam() == nil,
		"third": bt.ThirdParam() == nil, "fourth": bt.FourthParam() == nil,
		"fifth": bt.FifthParam() == nil,
		"hasHashmap": bt.HasValidHashmapParam(),
	}
	expected := args.Map{
		"first": true, "second": true, "third": true, "fourth": true, "fifth": true,
		"hasHashmap": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns nil -- NilParams", actual)
}

func Test_Cov_BaseTestCase_HashmapParam(t *testing.T) {
	bt := &coretests.BaseTestCase{
		Parameters: &args.HolderAny{Hashmap: map[string]any{"k": "v"}},
	}
	hasItem, hm := bt.HashmapParam()
	bt2 := &coretests.BaseTestCase{}
	hasItem2, _ := bt2.HashmapParam()
	actual := args.Map{"hasItem": hasItem, "len": len(hm), "noParam": hasItem2}
	expected := args.Map{"hasItem": true, "len": 1, "noParam": false}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- HashmapParam", actual)
}

func Test_Cov_BaseTestCase_SetActual(t *testing.T) {
	bt := &coretests.BaseTestCase{}
	bt.SetActual("test")
	actual := args.Map{"actual": bt.Actual()}
	expected := args.Map{"actual": "test"}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- SetActual", actual)
}

func Test_Cov_BaseTestCase_FormTitle(t *testing.T) {
	bt := &coretests.BaseTestCase{Title: "test"}
	actual := args.Map{
		"formTitle":   bt.FormTitle(0) != "",
		"customTitle": bt.CustomTitle(0, "custom") != "",
	}
	expected := args.Map{"formTitle": true, "customTitle": true}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- FormTitle", actual)
}

func Test_Cov_BaseTestCase_Contracts(t *testing.T) {
	bt := &coretests.BaseTestCase{}
	actual := args.Map{
		"simpleWrapper": bt.AsSimpleTestCaseWrapper() != nil,
		"baseWrapper":   bt.AsBaseTestCaseWrapper() != nil,
	}
	expected := args.Map{"simpleWrapper": true, "baseWrapper": true}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- Contracts", actual)
}

func Test_Cov_BaseTestCase_ActualExpectedLines(t *testing.T) {
	bt := &coretests.BaseTestCase{ActualInput: "line1\nline2", ExpectedInput: "e1\ne2"}
	actual := args.Map{
		"actualLen":   len(bt.ActualLines()),
		"expectedLen": len(bt.ExpectedLines()),
	}
	expected := args.Map{"actualLen": 2, "expectedLen": 2}
	expected.ShouldBeEqual(t, 0, "BaseTestCase returns correct value -- ActualExpectedLines", actual)
}
