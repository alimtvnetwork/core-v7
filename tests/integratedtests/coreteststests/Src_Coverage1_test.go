package coreteststests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_Src_AnyToBytes_Verification(t *testing.T) {
	for caseIndex, tc := range srcAnyToBytesTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal := input["input"]
		inputType := input["type"].(string)

		// Act
		var result []byte
		switch inputType {
		case "bytes":
			result = coretests.AnyToBytes(inputVal.([]byte))
		case "nilBytes":
			var nilBytes []byte
			result = coretests.AnyToBytes(nilBytes)
		case "string":
			result = coretests.AnyToBytes(inputVal.(string))
		case "other":
			result = coretests.AnyToBytes(inputVal)
		}

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["result"]; has {
			actual["result"] = string(result)
		}
		if _, has := expected["isNil"]; has {
			actual["isNil"] = result == nil
		}
		if _, has := expected["nonEmpty"]; has {
			actual["nonEmpty"] = len(result) > 0
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_DraftType_PtrOrNonPtr_Verification(t *testing.T) {
	for caseIndex, tc := range srcDraftTypePtrOrNonPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		d := &coretests.DraftType{
			SampleString1: input["string1"].(string),
			SampleInteger: input["integer"].(int),
		}
		asPtr := input["asPtr"].(bool)

		// Act
		result := d.PtrOrNonPtr(asPtr)

		// Assert
		expected := tc.ExpectedInput.(args.Map)
		actual := args.Map{}
		if _, has := expected["isNotNil"]; has {
			actual["isNotNil"] = result != nil
		}
		if _, has := expected["isDraftType"]; has {
			_, ok := result.(coretests.DraftType)
			actual["isDraftType"] = ok
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Src_DraftType_PtrOrNonPtr_NilReceiver(t *testing.T) {
	// Arrange
	var nilD *coretests.DraftType

	// Act
	result := nilD.PtrOrNonPtr(true)

	// Assert
	convey.Convey("PtrOrNonPtr returns nil -- nil receiver", t, func() {
		convey.So(result, should.BeNil)
	})
}
	d := coretests.DraftType{SampleString1: "x"}

	// Act
	s := d.JsonString()
	b := d.JsonBytes()
	b2 := d.JsonBytesPtr()

	// Assert
	convey.Convey("JsonString returns non-empty -- DraftType", t, func() {
		convey.So(s, should.NotBeEmpty)
	})
	convey.Convey("JsonBytes returns non-empty -- DraftType", t, func() {
		convey.So(len(b), should.BeGreaterThan, 0)
	})
	convey.Convey("JsonBytesPtr returns non-empty -- DraftType", t, func() {
		convey.So(len(b2), should.BeGreaterThan, 0)
	})

	// Arrange + Act (setters)
	d.SetF2Integer(42)

	// Assert
	convey.Convey("SetF2Integer sets value -- DraftType", t, func() {
		convey.So(d.F2Integer(), should.Equal, 42)
	})
	convey.Convey("F1String returns empty -- DraftType default", t, func() {
		convey.So(d.F1String(), should.BeEmpty)
	})
	_ = d.NonPtr()
}

func Test_Src_SimpleTestCase_Titles_Verification(t *testing.T) {
	// Arrange
	tc := srcSimpleTestCaseTitlesTestCase
	input := tc.ArrangeInput.(args.Map)
	title := input["title"].(string)

	stc := coretests.SimpleTestCase{Title: title}

	// Act
	caseTitle := stc.CaseTitle()
	formTitle := stc.FormTitle(0)
	customTitle := stc.CustomTitle(0, "custom")

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"caseTitle":           caseTitle,
		"formTitleNotEmpty":   formTitle != "",
		"customTitleNotEmpty": customTitle != "",
	})
}

func Test_Src_SimpleTestCase_ArrangeAndExpected(t *testing.T) {
	// Arrange
	stc := coretests.SimpleTestCase{
		Title:         "tc",
		ArrangeInput:  "arrange-val",
		ExpectedInput: "expected-val",
	}

	// Act
	arrangeStr := stc.ArrangeString()
	inputVal := stc.Input()
	expectedVal := stc.Expected()
	expectedStr := stc.ExpectedString()

	// Assert
	convey.Convey("ArrangeString returns non-empty -- SimpleTestCase", t, func() {
		convey.So(arrangeStr, should.NotBeEmpty)
	})
	convey.Convey("Input returns arrange-val -- SimpleTestCase", t, func() {
		convey.So(inputVal, should.Equal, "arrange-val")
	})
	convey.Convey("Expected returns expected-val -- SimpleTestCase", t, func() {
		convey.So(expectedVal, should.Equal, "expected-val")
	})
	convey.Convey("ExpectedString returns non-empty -- SimpleTestCase", t, func() {
		convey.So(expectedStr, should.NotBeEmpty)
	})

	// Act (setters) — SetActual is value receiver (no-op), assign directly
	stc.ActualInput = "actual-val"
	actualStr := stc.ActualString()
	str := stc.String(0)
	linesStr := stc.LinesString(0)

	// Assert
	convey.Convey("ActualString returns non-empty -- after SetActual", t, func() {
		convey.So(actualStr, should.NotBeEmpty)
	})
	convey.Convey("String returns non-empty -- SimpleTestCase", t, func() {
		convey.So(str, should.NotBeEmpty)
	})
	convey.Convey("LinesString returns non-empty -- SimpleTestCase", t, func() {
		convey.So(linesStr, should.NotBeEmpty)
	})
	_ = stc.AsSimpleTestCaseWrapper()
}
