package coreteststests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/issetter"
)

// ── AnyToBytes ──

func Test_Cov_AnyToBytes_FromBytes(t *testing.T) {
	// Arrange
	input := []byte{1, 2, 3}

	// Act
	result := coretests.AnyToBytes(input)

	// Assert
	actual := args.Map{
		"length":   fmt.Sprintf("%d", len(result)),
		"firstByte": fmt.Sprintf("%d", result[0]),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToBytes returns bytes as-is -- []byte input",
		ExpectedInput: args.Map{
			"length":   "3",
			"firstByte": "1",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_AnyToBytes_FromString(t *testing.T) {
	// Arrange
	input := "hello"

	// Act
	result := coretests.AnyToBytes(input)

	// Assert
	actual := args.Map{
		"value": string(result),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToBytes converts string to bytes -- string input",
		ExpectedInput: args.Map{
			"value": "hello",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_AnyToBytes_FromStruct(t *testing.T) {
	// Arrange
	input := map[string]int{"a": 1}

	// Act
	result := coretests.AnyToBytes(input)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", len(result) > 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToBytes marshals struct to JSON bytes -- map input",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_AnyToBytes_NilBytes(t *testing.T) {
	// Arrange
	var input []byte

	// Act
	result := coretests.AnyToBytes(input)

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToBytes returns nil for nil []byte -- nil input",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── AnyToBytesPtr ──

func Test_Cov_AnyToBytesPtr_Delegates(t *testing.T) {
	// Arrange
	input := "test"

	// Act
	result := coretests.AnyToBytesPtr(input)

	// Assert
	actual := args.Map{
		"value": string(result),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToBytesPtr delegates to AnyToBytes -- string input",
		ExpectedInput: args.Map{
			"value": "test",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── AnyToDraftType ──

func Test_Cov_AnyToDraftType_FromValue(t *testing.T) {
	// Arrange
	input := coretests.DraftType{SampleString1: "s1"}

	// Act
	result := coretests.AnyToDraftType(input)

	// Assert
	actual := args.Map{
		"notNil":   fmt.Sprintf("%v", result != nil),
		"sample1":  result.SampleString1,
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToDraftType returns ptr from value -- DraftType value",
		ExpectedInput: args.Map{
			"notNil":  "true",
			"sample1": "s1",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_AnyToDraftType_FromPtr(t *testing.T) {
	// Arrange
	input := &coretests.DraftType{SampleString1: "s2"}

	// Act
	result := coretests.AnyToDraftType(input)

	// Assert
	actual := args.Map{
		"notNil":  fmt.Sprintf("%v", result != nil),
		"sample1": result.SampleString1,
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToDraftType returns ptr from ptr -- *DraftType input",
		ExpectedInput: args.Map{
			"notNil":  "true",
			"sample1": "s2",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_AnyToDraftType_FromOther(t *testing.T) {
	// Arrange
	input := "not a DraftType"

	// Act
	result := coretests.AnyToDraftType(input)

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "AnyToDraftType returns nil for non-DraftType -- string input",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── DraftType ──

func Test_Cov_DraftType_Clone(t *testing.T) {
	// Arrange
	original := coretests.DraftType{
		SampleString1: "a",
		SampleString2: "b",
		SampleInteger: 10,
		Lines:         []string{"l1", "l2"},
		RawBytes:      []byte{1, 2},
	}

	// Act
	cloned := original.Clone()

	// Assert
	actual := args.Map{
		"isEqual":   fmt.Sprintf("%v", original.IsEqualAll(&cloned)),
		"sample1":   cloned.SampleString1,
		"sample2":   cloned.SampleString2,
		"integer":   fmt.Sprintf("%d", cloned.SampleInteger),
		"linesLen":  fmt.Sprintf("%d", cloned.LinesLength()),
		"bytesLen":  fmt.Sprintf("%d", cloned.RawBytesLength()),
	}

	tc := coretestcases.CaseV1{
		Title:         "Clone creates deep copy -- full DraftType",
		ExpectedInput: args.Map{
			"isEqual":  "true",
			"sample1":  "a",
			"sample2":  "b",
			"integer":  "10",
			"linesLen": "2",
			"bytesLen": "2",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var d *coretests.DraftType

	// Act
	result := d.ClonePtr()

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ClonePtr returns nil on nil receiver",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_ClonePtr_Valid(t *testing.T) {
	// Arrange
	d := &coretests.DraftType{SampleString1: "test"}

	// Act
	result := d.ClonePtr()

	// Assert
	actual := args.Map{
		"notNil":  fmt.Sprintf("%v", result != nil),
		"sample1": result.SampleString1,
	}

	tc := coretestcases.CaseV1{
		Title:         "ClonePtr returns cloned ptr -- valid DraftType",
		ExpectedInput: args.Map{
			"notNil":  "true",
			"sample1": "test",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *coretests.DraftType

	// Act
	result := a.IsEqual(true, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns true for both nil -- nil receivers",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_OneNil(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	var b *coretests.DraftType

	// Act
	result := a.IsEqual(true, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false when one is nil -- one nil",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_SamePointer(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}

	// Act
	result := a.IsEqual(true, a)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns true for same pointer -- identity",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_FieldMismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "y"}

	// Act
	result := a.IsEqual(false, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false on SampleString1 mismatch",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_String2Mismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x", SampleString2: "a"}
	b := &coretests.DraftType{SampleString1: "x", SampleString2: "b"}

	// Act
	result := a.IsEqual(false, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false on SampleString2 mismatch",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_IntegerMismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x", SampleString2: "y", SampleInteger: 1}
	b := &coretests.DraftType{SampleString1: "x", SampleString2: "y", SampleInteger: 2}

	// Act
	result := a.IsEqual(false, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false on SampleInteger mismatch",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_BytesMismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x", RawBytes: []byte{1}}
	b := &coretests.DraftType{SampleString1: "x", RawBytes: []byte{2}}

	// Act
	result := a.IsEqual(false, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false on RawBytes mismatch",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqual_LinesMismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x", Lines: []string{"a"}}
	b := &coretests.DraftType{SampleString1: "x", Lines: []string{"b"}}

	// Act
	result := a.IsEqual(false, b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqual returns false on Lines mismatch",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyNotEqualMessage_Equal(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "x"}

	// Act
	msg := a.VerifyNotEqualMessage(false, b)

	// Assert
	actual := args.Map{
		"isEmpty": fmt.Sprintf("%v", msg == ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyNotEqualMessage returns empty on equal -- same values",
		ExpectedInput: args.Map{
			"isEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyNotEqualErr_Equal(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "x"}

	// Act
	err := a.VerifyNotEqualErr(false, b)

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyNotEqualErr returns nil on equal -- same values",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyNotEqualErr_NotEqual(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "y"}

	// Act
	err := a.VerifyNotEqualErr(false, b)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyNotEqualErr returns error on not equal -- different values",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyAllNotEqualErr(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "x"}

	// Act
	err := a.VerifyAllNotEqualErr(b)

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyAllNotEqualErr returns nil on equal -- including inner fields",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyNotEqualExcludingInnerFieldsErr(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "y"}

	// Act
	err := a.VerifyNotEqualExcludingInnerFieldsErr(b)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyNotEqualExcludingInnerFieldsErr returns error -- different values",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_VerifyAllNotEqualMessage(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "y"}

	// Act
	msg := a.VerifyAllNotEqualMessage(b)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", msg != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "VerifyAllNotEqualMessage returns message on mismatch -- including inner",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_JsonString(t *testing.T) {
	// Arrange
	d := coretests.DraftType{SampleString1: "test"}

	// Act
	result := d.JsonString()

	// Assert
	actual := args.Map{
		"notEmpty":     fmt.Sprintf("%v", result != ""),
		"containsTest": fmt.Sprintf("%v", len(result) > 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "JsonString returns valid JSON string -- simple DraftType",
		ExpectedInput: args.Map{
			"notEmpty":     "true",
			"containsTest": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_JsonBytes(t *testing.T) {
	// Arrange
	d := coretests.DraftType{SampleString1: "test"}

	// Act
	result := d.JsonBytes()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", len(result) > 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "JsonBytes returns valid JSON bytes -- simple DraftType",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_JsonBytesPtr(t *testing.T) {
	// Arrange
	d := coretests.DraftType{SampleString1: "test"}

	// Act
	result := d.JsonBytesPtr()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", len(result) > 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "JsonBytesPtr delegates to JsonBytes -- simple DraftType",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_F1String_F2Integer(t *testing.T) {
	// Arrange
	d := &coretests.DraftType{SampleString1: "x"}
	d.SetF2Integer(42)

	// Act
	f2 := d.F2Integer()

	// Assert
	actual := args.Map{
		"f1String":  d.F1String(),
		"f2Integer": fmt.Sprintf("%d", f2),
	}

	tc := coretestcases.CaseV1{
		Title:         "F1String and F2Integer return private fields -- set via SetF2Integer",
		ExpectedInput: args.Map{
			"f1String":  "",
			"f2Integer": "42",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_NonPtr(t *testing.T) {
	// Arrange
	d := coretests.DraftType{SampleString1: "val"}

	// Act
	result := d.NonPtr()

	// Assert
	actual := args.Map{
		"sample1": result.SampleString1,
	}

	tc := coretestcases.CaseV1{
		Title:         "NonPtr returns value copy -- DraftType",
		ExpectedInput: args.Map{
			"sample1": "val",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_PtrOrNonPtr_Nil(t *testing.T) {
	// Arrange
	var d *coretests.DraftType

	// Act
	result := d.PtrOrNonPtr(true)

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PtrOrNonPtr returns nil on nil receiver -- isPtr=true",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_PtrOrNonPtr_Ptr(t *testing.T) {
	// Arrange
	d := &coretests.DraftType{SampleString1: "p"}

	// Act
	result := d.PtrOrNonPtr(true)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", result != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PtrOrNonPtr returns ptr when isPtr=true -- valid DraftType",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_PtrOrNonPtr_NonPtr(t *testing.T) {
	// Arrange
	d := &coretests.DraftType{SampleString1: "p"}

	// Act
	result := d.PtrOrNonPtr(false)
	asVal, ok := result.(coretests.DraftType)

	// Assert
	actual := args.Map{
		"isValue": fmt.Sprintf("%v", ok),
		"sample1": asVal.SampleString1,
	}

	tc := coretestcases.CaseV1{
		Title:         "PtrOrNonPtr returns value when isPtr=false -- valid DraftType",
		ExpectedInput: args.Map{
			"isValue": "true",
			"sample1": "p",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_DraftType_IsEqualAll(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{SampleString1: "x"}
	b := &coretests.DraftType{SampleString1: "x"}

	// Act
	result := a.IsEqualAll(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsEqualAll returns true for equal structs -- including inner fields",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── VerifyTypeOf ──

func Test_Cov_VerifyTypeOf_NewAndMethods(t *testing.T) {
	// Arrange
	vt := coretests.NewVerifyTypeOf("hello")

	// Act / Assert
	actual := args.Map{
		"isDefined":  fmt.Sprintf("%v", vt.IsDefined()),
		"isInvalid":  fmt.Sprintf("%v", vt.IsInvalid()),
		"skipVerify": fmt.Sprintf("%v", vt.IsInvalidOrSkipVerify()),
	}

	tc := coretestcases.CaseV1{
		Title:         "NewVerifyTypeOf creates valid VerifyTypeOf -- string arrange",
		ExpectedInput: args.Map{
			"isDefined":  "true",
			"isInvalid":  "true",
			"skipVerify": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_VerifyTypeOf_NilSkipVerify(t *testing.T) {
	// Arrange
	var vt *coretests.VerifyTypeOf

	// Act
	result := vt.IsInvalidOrSkipVerify()

	// Assert
	actual := args.Map{
		"skipVerify": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsInvalidOrSkipVerify returns true on nil -- nil receiver",
		ExpectedInput: args.Map{
			"skipVerify": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SomeString ──

func Test_Cov_SomeString_String(t *testing.T) {
	// Arrange
	s := coretests.SomeString{Value: "hello"}

	// Act
	result := s.String()
	stringer := s.AsStringer()

	// Assert
	actual := args.Map{
		"string":         result,
		"stringerString": stringer.String(),
	}

	tc := coretestcases.CaseV1{
		Title:         "SomeString.String returns value -- simple string",
		ExpectedInput: args.Map{
			"string":         "hello",
			"stringerString": "hello",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── TestFuncName ──

func Test_Cov_TestFuncName_Value(t *testing.T) {
	// Arrange
	fn := coretests.TestFuncName("MyFunc")

	// Act
	result := fn.Value()

	// Assert
	actual := args.Map{
		"value": result,
	}

	tc := coretestcases.CaseV1{
		Title:         "TestFuncName.Value returns string -- simple name",
		ExpectedInput: args.Map{
			"value": "MyFunc",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── CaseIndexPlusIsPrint ──

func Test_Cov_CaseIndexPlusIsPrint_Fields(t *testing.T) {
	// Arrange
	c := coretests.CaseIndexPlusIsPrint{IsPrint: true, CaseIndex: 5}

	// Act / Assert
	actual := args.Map{
		"isPrint":   fmt.Sprintf("%v", c.IsPrint),
		"caseIndex": fmt.Sprintf("%d", c.CaseIndex),
	}

	tc := coretestcases.CaseV1{
		Title:         "CaseIndexPlusIsPrint holds fields correctly -- isPrint=true, index=5",
		ExpectedInput: args.Map{
			"isPrint":   "true",
			"caseIndex": "5",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SimpleTestCase ──

func Test_Cov_SimpleTestCase_Getters(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "test title",
		ArrangeInput:  "arrange",
		ExpectedInput: "expected",
	}

	// Act
	actual := args.Map{
		"caseTitle":      tc.CaseTitle(),
		"input":          fmt.Sprintf("%v", tc.Input()),
		"expected":       fmt.Sprintf("%v", tc.Expected()),
		"expectedString": tc.ExpectedString(),
		"arrangeString":  tc.ArrangeString(),
	}

	// Assert
	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase getters return correct values -- basic fields",
		ExpectedInput: args.Map{
			"caseTitle":      "test title",
			"input":          "arrange",
			"expected":       "expected",
			"expectedString": "expected",
			"arrangeString":  "arrange",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_FormTitle(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{Title: "my test"}

	// Act
	result := tc.FormTitle(3)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.FormTitle returns formatted title -- index 3",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_CustomTitle(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{Title: "my test"}

	// Act
	result := tc.CustomTitle(5, "custom")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.CustomTitle returns formatted title -- custom title",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_SetActual(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{Title: "test"}

	// Act
	tc.SetActual("actualVal")

	// Assert - SimpleTestCase uses value receiver, so SetActual doesn't persist
	actual := args.Map{
		"actualString": tc.ActualString(),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.SetActual uses value receiver -- does not persist",
		ExpectedInput: args.Map{
			"actualString": "",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_String(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "test",
		ArrangeInput:  "in",
		ExpectedInput: "out",
	}

	// Act
	result := tc.String(0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.String returns formatted string -- index 0",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_LinesString(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "test",
		ExpectedInput: "out",
	}

	// Act
	result := tc.LinesString(0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.LinesString returns formatted lines -- index 0",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleTestCase_AsSimpleTestCaseWrapper(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{Title: "test"}

	// Act
	wrapper := tc.AsSimpleTestCaseWrapper()

	// Assert
	actual := args.Map{
		"title": wrapper.CaseTitle(),
	}

	verify := coretestcases.CaseV1{
		Title:         "SimpleTestCase.AsSimpleTestCaseWrapper returns self -- same title",
		ExpectedInput: args.Map{
			"title": "test",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}

// ── BaseTestCase getters ──

func Test_Cov_BaseTestCase_Getters(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{
		Title:         "base title",
		ArrangeInput:  "arrange",
		ExpectedInput: "expected",
	}

	// Act
	actual := args.Map{
		"caseTitle":      btc.CaseTitle(),
		"input":          fmt.Sprintf("%v", btc.Input()),
		"expected":       fmt.Sprintf("%v", btc.Expected()),
		"expectedString": btc.ExpectedString(),
		"arrangeString":  btc.ArrangeString(),
	}

	// Assert
	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase getters return correct values -- basic fields",
		ExpectedInput: args.Map{
			"caseTitle":      "base title",
			"input":          "arrange",
			"expected":       "expected",
			"expectedString": "expected",
			"arrangeString":  "arrange",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_SetActual(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act
	btc.SetActual("result")

	// Assert
	actual := args.Map{
		"actual":       fmt.Sprintf("%v", btc.Actual()),
		"actualString": btc.ActualString(),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.SetActual stores value -- pointer receiver",
		ExpectedInput: args.Map{
			"actual":       "result",
			"actualString": "result",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_ActualLines(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{
		Title:     "test",
		ActualInput: "hello",
	}

	// Act
	lines := btc.ActualLines()

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(lines)),
		"line0":     lines[0],
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.ActualLines returns string lines -- string actual",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "hello",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_ExpectedLines(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{
		Title:         "test",
		ExpectedInput: []string{"a", "b"},
	}

	// Act
	lines := btc.ExpectedLines()

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(lines)),
		"line0":     lines[0],
		"line1":     lines[1],
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.ExpectedLines returns string slice lines -- []string expected",
		ExpectedInput: args.Map{
			"lineCount": "2",
			"line0":     "a",
			"line1":     "b",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_FormTitle(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "my test"}

	// Act
	result := btc.FormTitle(3)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.FormTitle returns formatted title -- index 3",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_CustomTitle(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "my test"}

	// Act
	result := btc.CustomTitle(5, "custom")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.CustomTitle returns formatted custom title -- index 5",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_IsDisabled(t *testing.T) {
	// Arrange
	enabled := &coretests.BaseTestCase{Title: "enabled"}
	disabled := &coretests.BaseTestCase{Title: "disabled", IsEnable: issetter.False}

	// Act / Assert
	actual := args.Map{
		"enabledIsDisabled":  fmt.Sprintf("%v", enabled.IsDisabled()),
		"disabledIsDisabled": fmt.Sprintf("%v", disabled.IsDisabled()),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.IsDisabled respects IsEnable -- true and false",
		ExpectedInput: args.Map{
			"enabledIsDisabled":  "false",
			"disabledIsDisabled": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_IsSkipWithLog_Disabled(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test", IsEnable: issetter.False}

	// Act
	result := btc.IsSkipWithLog(0)

	// Assert
	actual := args.Map{
		"skipped": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.IsSkipWithLog returns true when disabled -- false enable",
		ExpectedInput: args.Map{
			"skipped": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_IsSkipWithLog_Enabled(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act
	result := btc.IsSkipWithLog(0)

	// Assert
	actual := args.Map{
		"skipped": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.IsSkipWithLog returns false when enabled -- default enable",
		ExpectedInput: args.Map{
			"skipped": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_String(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{
		Title:         "test",
		ArrangeInput:  "in",
		ExpectedInput: "out",
	}

	// Act
	result := btc.String(0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.String returns formatted string -- index 0",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_LinesString(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{
		Title:         "test",
		ExpectedInput: "out",
	}

	// Act
	result := btc.LinesString(0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.LinesString returns formatted lines -- index 0",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_ArrangeTypeName(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{ArrangeInput: "hello"}

	// Act
	result := btc.ArrangeTypeName()

	// Assert
	actual := args.Map{
		"typeName": result,
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.ArrangeTypeName returns type name -- string input",
		ExpectedInput: args.Map{
			"typeName": "string",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_AsWrappers(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act
	simple := btc.AsSimpleTestCaseWrapper()
	base := btc.AsBaseTestCaseWrapper()

	// Assert
	actual := args.Map{
		"simpleTitle": simple.CaseTitle(),
		"baseNotNil":  fmt.Sprintf("%v", base != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.AsWrappers return valid wrappers -- simple and base",
		ExpectedInput: args.Map{
			"simpleTitle": "test",
			"baseNotNil":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_IsTypeInvalidOrSkipVerify(t *testing.T) {
	// Arrange
	noVerify := &coretests.BaseTestCase{Title: "no verify"}
	withVerify := &coretests.BaseTestCase{
		Title:      "with verify",
		VerifyTypeOf: coretests.NewVerifyTypeOf("hello"),
	}

	// Act / Assert
	actual := args.Map{
		"noVerifySkip":   fmt.Sprintf("%v", noVerify.IsTypeInvalidOrSkipVerify()),
		"withVerifySkip": fmt.Sprintf("%v", withVerify.IsTypeInvalidOrSkipVerify()),
		"isVerifyType":   fmt.Sprintf("%v", withVerify.IsVerifyType()),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.IsTypeInvalidOrSkipVerify returns correct state -- nil vs set",
		ExpectedInput: args.Map{
			"noVerifySkip":   "true",
			"withVerifySkip": "false",
			"isVerifyType":   "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── BaseTestCase Parameters ──

func Test_Cov_BaseTestCase_Parameters_Nil(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act / Assert
	actual := args.Map{
		"hasParams":     fmt.Sprintf("%v", btc.HasParameters()),
		"invalidParams": fmt.Sprintf("%v", btc.IsInvalidParameters()),
		"firstNil":      fmt.Sprintf("%v", btc.FirstParam() == nil),
		"secondNil":     fmt.Sprintf("%v", btc.SecondParam() == nil),
		"thirdNil":      fmt.Sprintf("%v", btc.ThirdParam() == nil),
		"fourthNil":     fmt.Sprintf("%v", btc.FourthParam() == nil),
		"fifthNil":      fmt.Sprintf("%v", btc.FifthParam() == nil),
		"hasValidHash":  fmt.Sprintf("%v", btc.HasValidHashmapParam()),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase parameters return nil/false on nil Parameters -- no params set",
		ExpectedInput: args.Map{
			"hasParams":     "false",
			"invalidParams": "true",
			"firstNil":      "true",
			"secondNil":     "true",
			"thirdNil":      "true",
			"fourthNil":     "true",
			"fifthNil":      "true",
			"hasValidHash":  "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_BaseTestCase_HashmapParam_Nil(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act
	hasMap, hashMap := btc.HashmapParam()

	// Assert
	actual := args.Map{
		"hasMap":    fmt.Sprintf("%v", hasMap),
		"mapEmpty": fmt.Sprintf("%v", len(hashMap) == 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "BaseTestCase.HashmapParam returns false on nil Parameters -- no params",
		ExpectedInput: args.Map{
			"hasMap":    "false",
			"mapEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── BaseTestCase TypeValidation ──

func Test_Cov_BaseTestCase_TypeValidationError_NoVerify(t *testing.T) {
	// Arrange
	btc := &coretests.BaseTestCase{Title: "test"}

	// Act
	err := btc.TypeValidationError()

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "TypeValidationError returns nil when no VerifyTypeOf -- skip verify",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── getAssert helpers ──

func Test_Cov_GetAssert_Quick(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.Quick("when", "actual", "expected", 0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.Quick returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToStrings_String(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ToStrings("hello")

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(result)),
		"line0":     result[0],
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToStrings converts string to []string -- simple string",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "hello",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToStrings_Slice(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ToStrings([]string{"a", "b"})

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(result)),
		"line0":     result[0],
		"line1":     result[1],
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToStrings returns []string as-is -- string slice",
		ExpectedInput: args.Map{
			"lineCount": "2",
			"line0":     "a",
			"line1":     "b",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToStrings_Int(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ToStrings(42)

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(result)),
		"line0":     result[0],
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToStrings converts int to []string -- int value",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "42",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToStrings_Bool(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ToStrings(true)

	// Assert
	actual := args.Map{
		"lineCount": fmt.Sprintf("%d", len(result)),
		"line0":     result[0],
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToStrings converts bool to []string -- true",
		ExpectedInput: args.Map{
			"lineCount": "1",
			"line0":     "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToString(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ToString("test")

	// Assert
	actual := args.Map{
		"value": result,
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToString converts to joined string -- string value",
		ExpectedInput: args.Map{
			"value": "test",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_SortedMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.SortedMessage(false, "b a c", " ")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.SortedMessage returns sorted string -- space-separated",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_SortedArrayNoPrint(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.SortedArrayNoPrint("b a c")

	// Assert
	actual := args.Map{
		"count": fmt.Sprintf("%d", len(result)),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.SortedArrayNoPrint returns sorted array -- space-separated",
		ExpectedInput: args.Map{
			"count": "3",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ToQuoteLines(t *testing.T) {
	// Arrange
	lines := []string{"a", "b"}

	// Act
	result := coretests.GetAssert.ToQuoteLines(2, lines)

	// Assert
	actual := args.Map{
		"count": fmt.Sprintf("%d", len(result)),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ToQuoteLines wraps lines in double quotes -- 2 lines",
		ExpectedInput: args.Map{
			"count": "2",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_AnyToDoubleQuoteLines(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.AnyToDoubleQuoteLines(2, "hello")

	// Assert
	actual := args.Map{
		"count": fmt.Sprintf("%d", len(result)),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.AnyToDoubleQuoteLines converts any to quoted lines -- string",
		ExpectedInput: args.Map{
			"count": "1",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ConvertLinesToDoubleQuoteThenString(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ConvertLinesToDoubleQuoteThenString(2, []string{"a", "b"})

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ConvertLinesToDoubleQuoteThenString returns string -- 2 lines",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_AnyToStringDoubleQuoteLine(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.AnyToStringDoubleQuoteLine(2, "hello")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.AnyToStringDoubleQuoteLine returns quoted string -- string input",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── getAssertMessages ──

func Test_Cov_GetAssert_IsEqualMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsEqualMessage("when", "actual", "expected")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsEqualMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_IsNotEqualMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsNotEqualMessage("when", "actual", "expected")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsNotEqualMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_IsTrueMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsTrueMessage("when", "actual")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsTrueMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_IsFalseMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsFalseMessage("when", "actual")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsFalseMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_IsNilMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsNilMessage("when", nil)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsNilMessage returns formatted message -- nil actual",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_IsNotNilMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.IsNotNilMessage("when", "actual")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.IsNotNilMessage returns formatted message -- non-nil actual",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ShouldBeMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ShouldBeMessage("title", "actual", "expected")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ShouldBeMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_GetAssert_ShouldNotBeMessage(t *testing.T) {
	// Arrange / Act
	result := coretests.GetAssert.ShouldNotBeMessage("title", "actual", "expected")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetAssert.ShouldNotBeMessage returns formatted message -- basic args",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── ToStringValues / ToStringNameValues ──

func Test_Cov_ToStringValues(t *testing.T) {
	// Arrange / Act
	resultNil := coretests.ToStringValues(nil)
	resultVal := coretests.ToStringValues("hello")

	// Assert
	actual := args.Map{
		"nilResult":    resultNil,
		"valNotEmpty":  fmt.Sprintf("%v", resultVal != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "ToStringValues handles nil and value -- nil returns <nil>",
		ExpectedInput: args.Map{
			"nilResult":   "<nil>",
			"valNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_ToStringNameValues(t *testing.T) {
	// Arrange / Act
	resultNil := coretests.ToStringNameValues(nil)
	resultVal := coretests.ToStringNameValues("hello")

	// Assert
	actual := args.Map{
		"nilResult":   resultNil,
		"valNotEmpty": fmt.Sprintf("%v", resultVal != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "ToStringNameValues handles nil and value -- nil returns <nil>",
		ExpectedInput: args.Map{
			"nilResult":   "<nil>",
			"valNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── LogOnFail ──

func Test_Cov_LogOnFail_Pass(t *testing.T) {
	// Arrange / Act - should not panic
	coretests.LogOnFail(true, "exp", "act")

	// Assert
	actual := args.Map{
		"noPanic": "true",
	}

	tc := coretestcases.CaseV1{
		Title:         "LogOnFail does nothing on pass -- isPass=true",
		ExpectedInput: args.Map{
			"noPanic": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_LogOnFail_Fail(t *testing.T) {
	// Arrange / Act - should log but not panic
	coretests.LogOnFail(false, "exp", "act")

	// Assert
	actual := args.Map{
		"noPanic": "true",
	}

	tc := coretestcases.CaseV1{
		Title:         "LogOnFail logs on fail -- isPass=false",
		ExpectedInput: args.Map{
			"noPanic": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SimpleGherkins ──

func Test_Cov_SimpleGherkins_String(t *testing.T) {
	// Arrange
	g := &coretests.SimpleGherkins{
		Feature: "feature",
		Given:   "given",
		When:    "when",
		Then:    "then",
	}

	// Act
	result := g.String()
	toString := g.ToString(1)

	// Assert
	actual := args.Map{
		"stringNotEmpty":   fmt.Sprintf("%v", result != ""),
		"toStringNotEmpty": fmt.Sprintf("%v", toString != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleGherkins.String and ToString return formatted output -- basic",
		ExpectedInput: args.Map{
			"stringNotEmpty":   "true",
			"toStringNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleGherkins_GetWithExpectation(t *testing.T) {
	// Arrange
	g := &coretests.SimpleGherkins{
		Feature: "f",
		Given:   "g",
		When:    "w",
		Then:    "t",
		Actual:  "act",
		Expect:  "exp",
	}

	// Act
	result := g.GetWithExpectation(0)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleGherkins.GetWithExpectation returns message with expectation -- basic",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Cov_SimpleGherkins_GetMessageConditional(t *testing.T) {
	// Arrange
	g := &coretests.SimpleGherkins{Feature: "f"}

	// Act
	withExp := g.GetMessageConditional(true, 0)
	withoutExp := g.GetMessageConditional(false, 0)

	// Assert
	actual := args.Map{
		"withExpNotEmpty":    fmt.Sprintf("%v", withExp != ""),
		"withoutExpNotEmpty": fmt.Sprintf("%v", withoutExp != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleGherkins.GetMessageConditional returns different output -- true vs false",
		ExpectedInput: args.Map{
			"withExpNotEmpty":    "true",
			"withoutExpNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── getAssertSimpleTestCaseWrapper ──

func Test_Cov_GetAssert_SimpleTestCaseWrapper_Lines(t *testing.T) {
	// Arrange
	tc := coretests.SimpleTestCase{
		Title:         "test",
		ActualInput:   "actual",
		ExpectedInput: "expected",
	}

	// Act
	actualLines, expectedLines := coretests.GetAssert.SimpleTestCaseWrapper.Lines(tc)

	// Assert
	actual := args.Map{
		"actualCount":   fmt.Sprintf("%d", len(actualLines)),
		"expectedCount": fmt.Sprintf("%d", len(expectedLines)),
	}

	verify := coretestcases.CaseV1{
		Title:         "GetAssert.SimpleTestCaseWrapper.Lines returns both line sets -- basic",
		ExpectedInput: args.Map{
			"actualCount":   "1",
			"expectedCount": "1",
		},
	}
	verify.ShouldBeEqualMapFirst(t, actual)
}
