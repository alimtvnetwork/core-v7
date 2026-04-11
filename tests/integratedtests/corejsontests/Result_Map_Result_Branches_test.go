package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Result — Map
// =============================================================================

func Test_Result_Map_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Result Map nil", actual)
}

func Test_Result_Map_WithBytesAndError(t *testing.T) {
	// Arrange
	r := &corejson.Result{
		Bytes:    []byte(`"hello"`),
		Error:    errors.New("fail"),
		TypeName: "string",
	}
	m := r.Map()
	// When Result has Error, IsEmptyJsonBytes returns true (HasError check),
	// so JsonString() returns "" → m["Bytes"] is ""

	// Act
	actual := args.Map{
		"hasBytes": m["Bytes"] != "",
		"hasError": m["Error"] != "",
		"hasType": m["Type"] != "",
	}

	// Assert
	expected := args.Map{
		"hasBytes": false,
		"hasError": true,
		"hasType": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Map with all fields", actual)
}

func Test_Result_Map_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	m := r.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Result Map empty", actual)
}

// =============================================================================
// Result — DeserializedFieldsToMap / SafeDeserializedFieldsToMap
// =============================================================================

func Test_Result_DeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()

	// Act
	actual := args.Map{
		"len": len(m),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializedFieldsToMap nil", actual)
}

func Test_Result_SafeDeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.SafeDeserializedFieldsToMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeDeserializedFieldsToMap nil", actual)
}

// =============================================================================
// Result — FieldsNames / SafeFieldsNames
// =============================================================================

func Test_Result_FieldsNames_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	names, err := r.FieldsNames()

	// Act
	actual := args.Map{
		"len": len(names),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldsNames nil", actual)
}

func Test_Result_SafeFieldsNames_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	names := r.SafeFieldsNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeFieldsNames nil", actual)
}

// =============================================================================
// Result — BytesTypeName / SafeBytesTypeName
// =============================================================================

func Test_Result_BytesTypeName_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"r": r.BytesTypeName()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesTypeName nil", actual)
}

func Test_Result_BytesTypeName_Set(t *testing.T) {
	// Arrange
	r := &corejson.Result{TypeName: "int"}

	// Act
	actual := args.Map{"r": r.BytesTypeName()}

	// Assert
	expected := args.Map{"r": "int"}
	expected.ShouldBeEqual(t, 0, "BytesTypeName set", actual)
}

func Test_Result_SafeBytesTypeName_Empty_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"r": r.SafeBytesTypeName()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName empty", actual)
}

func Test_Result_SafeBytesTypeName_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"hasName": r.SafeBytesTypeName() != ""}

	// Assert
	expected := args.Map{"hasName": true}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName valid", actual)
}

// =============================================================================
// Result — JsonStringPtr caching
// =============================================================================

func Test_Result_JsonStringPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"r": *r.JsonStringPtr()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr nil", actual)
}

func Test_Result_JsonStringPtr_NoBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	s := r.JsonStringPtr()

	// Act
	actual := args.Map{"r": *s}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr no bytes", actual)
}

func Test_Result_JsonStringPtr_WithBytes(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s1 := r.JsonStringPtr()
	s2 := r.JsonStringPtr() // cached

	// Act
	actual := args.Map{
		"same": s1 == s2,
		"hasContent": len(*s1) > 0,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"hasContent": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr cached", actual)
}

// =============================================================================
// Result — PrettyJsonString / PrettyJsonStringOrErrString / PrettyJsonBuffer
// =============================================================================

func Test_Result_PrettyJsonString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"r": r.PrettyJsonString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString nil", actual)
}

func Test_Result_PrettyJsonString_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(map[string]int{"a": 1})
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString valid", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString nil", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Error(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString error", actual)
}

func Test_Result_PrettyJsonStringOrErrString_Valid_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString valid", actual)
}

func Test_Result_PrettyJsonBuffer_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"empty": buf.Len() == 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "PrettyJsonBuffer empty", actual)
}

// =============================================================================
// Result — Length, HasError, ErrorString, IsErrorEqual
// =============================================================================

func Test_Result_Length_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"r": r.Length()}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "Result Length nil", actual)
}

func Test_Result_ErrorString_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": r.ErrorString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "ErrorString no error", actual)
}

func Test_Result_ErrorString_HasError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{"hasContent": len(r.ErrorString()) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "ErrorString has error", actual)
}

func Test_Result_IsErrorEqual_BothNil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"r": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual both nil", actual)
}

func Test_Result_IsErrorEqual_OneNil(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("x")}

	// Act
	actual := args.Map{"r": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual one nil", actual)
}

func Test_Result_IsErrorEqual_SameMsg(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("x")}

	// Act
	actual := args.Map{"r": r.IsErrorEqual(errors.New("x"))}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual same msg", actual)
}

func Test_Result_IsErrorEqual_DiffMsg(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("x")}

	// Act
	actual := args.Map{"r": r.IsErrorEqual(errors.New("y"))}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual diff msg", actual)
}

// =============================================================================
// Result — String
// =============================================================================

func Test_Result_String_Empty(t *testing.T) {
	// Arrange
	r := corejson.Result{}

	// Act
	actual := args.Map{"r": r.String()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "Result String empty", actual)
}

func Test_Result_String_WithError(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail"), TypeName: "string"}
	s := r.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Result String with error", actual)
}

func Test_Result_String_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s := r.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Result String no error", actual)
}

// =============================================================================
// Result — SafeNonIssueBytes, SafeBytes, Values, SafeValues, SafeValuesPtr
// =============================================================================

func Test_Result_SafeNonIssueBytes_HasIssues_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("x")}

	// Act
	actual := args.Map{"len": len(r.SafeNonIssueBytes())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes has issues", actual)
}

func Test_Result_SafeNonIssueBytes_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"hasBytes": len(r.SafeNonIssueBytes()) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes valid", actual)
}

func Test_Result_SafeBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"len": len(r.SafeBytes())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeBytes nil", actual)
}

func Test_Result_SafeValues_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"len": len(r.SafeValues())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeValues nil", actual)
}

func Test_Result_SafeValuesPtr_HasIssues_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"len": len(r.SafeValuesPtr())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeValuesPtr has issues", actual)
}

// =============================================================================
// Result — Raw, RawMust, RawString, RawStringMust, RawErrString, RawPrettyString
// =============================================================================

func Test_Result_Raw_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	b, err := r.Raw()

	// Act
	actual := args.Map{
		"len": len(b),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Raw nil", actual)
}

func Test_Result_Raw_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	b, err := r.Raw()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Raw valid", actual)
}

func Test_Result_RawString_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s, err := r.RawString()

	// Act
	actual := args.Map{
		"hasContent": len(s) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawString valid", actual)
}

func Test_Result_RawStringMust_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "RawStringMust valid", actual)
}

func Test_Result_RawStringMust_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RawStringMust panics", actual)
	}()
	r := &corejson.Result{Error: errors.New("fail")}
	r.RawStringMust()
}

func Test_Result_RawErrString_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail")}
	b, msg := r.RawErrString()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"hasMsg": len(msg) > 0,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasMsg": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrString", actual)
}

func Test_Result_RawPrettyString_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{
		"hasContent": len(s) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawPrettyString", actual)
}

// =============================================================================
// Result — MeaningfulError / MeaningfulErrorMessage
// =============================================================================

func Test_Result_MeaningfulError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError nil", actual)
}

func Test_Result_MeaningfulError_OK(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"noErr": r.MeaningfulError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError OK", actual)
}

func Test_Result_MeaningfulError_EmptyBytes_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{TypeName: "int"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError empty bytes", actual)
}

func Test_Result_MeaningfulError_WithBothErrorAndBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail"), TypeName: "string"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError with error and bytes", actual)
}

func Test_Result_MeaningfulErrorMessage_NoError(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": r.MeaningfulErrorMessage()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorMessage no error", actual)
}

// =============================================================================
// Result — IsEmptyJsonBytes branches
// =============================================================================

func Test_Result_IsEmptyJsonBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"r": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes nil", actual)
}

func Test_Result_IsEmptyJsonBytes_EmptyJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte("{}")}

	// Act
	actual := args.Map{"r": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes {}", actual)
}

func Test_Result_IsEmptyJsonBytes_ZeroLen(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte{}}

	// Act
	actual := args.Map{"r": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes zero len", actual)
}

func Test_Result_IsEmptyJsonBytes_HasContent(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes has content", actual)
}

// =============================================================================
// Result — Unmarshal / UnmarshalMust / DeserializeMust / UnmarshalSkipExistingIssues
// =============================================================================

func Test_Result_Unmarshal_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal nil", actual)
}

func Test_Result_Unmarshal_HasError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail"), TypeName: "x"}
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal has error", actual)
}

func Test_Result_Unmarshal_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Unmarshal valid", actual)
}

func Test_Result_Unmarshal_BadPayload(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`bad`), TypeName: "x"}
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal bad payload", actual)
}

func Test_Result_DeserializeMust_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "DeserializeMust panics", actual)
	}()
	r := &corejson.Result{Error: errors.New("fail")}
	var s string
	r.DeserializeMust(&s)
}

func Test_Result_UnmarshalSkipExistingIssues_HasIssues_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues has issues", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_Valid_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues valid", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_BadPayload(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`bad`), TypeName: "x"}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues bad payload", actual)
}

func Test_Result_UnmarshalResult_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(corejson.Result{Bytes: []byte(`"x"`), TypeName: "string"})
	inner, err := r.UnmarshalResult()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": inner != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalResult", actual)
}

// =============================================================================
// Result — Serialize / SerializeMust / SerializeSkipExistingIssues
// =============================================================================

func Test_Result_Serialize_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize nil", actual)
}

func Test_Result_Serialize_HasError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	_, err := r.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize has error", actual)
}

func Test_Result_Serialize_Valid_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	b, err := r.Serialize()

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
	expected.ShouldBeEqual(t, 0, "Serialize valid", actual)
}

func Test_Result_SerializeSkipExistingIssues_HasIssues_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{
		"nilBytes": b == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nilBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues has issues", actual)
}

func Test_Result_SerializeSkipExistingIssues_Valid_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	b, err := r.SerializeSkipExistingIssues()

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
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues valid", actual)
}

// =============================================================================
// Result — HandleError / MustBeSafe / HandleErrorWithMsg (panic paths)
// =============================================================================

func Test_Result_HandleError_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "HandleError panics", actual)
	}()
	r := &corejson.Result{Error: errors.New("fail")}
	r.HandleError()
}

func Test_Result_HandleError_Safe(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	r.HandleError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError safe", actual)
}

func Test_Result_MustBeSafe_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeSafe panics", actual)
	}()
	r := &corejson.Result{}
	r.MustBeSafe()
}

func Test_Result_HandleErrorWithMsg_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "HandleErrorWithMsg panics", actual)
	}()
	r := &corejson.Result{}
	r.HandleErrorWithMsg("context")
}

// =============================================================================
// Result — JsonModel / JsonModelAny / Json / JsonPtr
// =============================================================================

func Test_Result_JsonModel_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.JsonModel()

	// Act
	actual := args.Map{"hasErr": m.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "JsonModel nil", actual)
}

func Test_Result_JsonModel_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	m := r.JsonModel()

	// Act
	actual := args.Map{"noErr": !m.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonModel valid", actual)
}

func Test_Result_JsonModelAny_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	a := r.JsonModelAny()

	// Act
	actual := args.Map{"notNil": a != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
}

// =============================================================================
// Result — ParseInjectUsingJson / ParseInjectUsingJsonMust
// =============================================================================

func Test_Result_ParseInjectUsingJson_Valid(t *testing.T) {
	// Arrange
	original := corejson.NewPtr("hello")
	serialized := corejson.NewPtr(*original)
	target := corejson.Empty.ResultPtr()
	_, err := target.ParseInjectUsingJson(serialized)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson valid", actual)
}

func Test_Result_ParseInjectUsingJson_Fail(t *testing.T) {
	// Arrange
	bad := &corejson.Result{Error: errors.New("fail")}
	target := corejson.Empty.ResultPtr()
	_, err := target.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson fail", actual)
}

func Test_Result_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics", actual)
	}()
	bad := &corejson.Result{Error: errors.New("fail")}
	target := corejson.Empty.ResultPtr()
	target.ParseInjectUsingJsonMust(bad)
}

// =============================================================================
// Result — Clone / ClonePtr / CloneIf / CloneError
// =============================================================================

func Test_Result_Clone_Empty_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	c := r.Clone(true)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone empty", actual)
}

func Test_Result_Clone_ShallowCopy(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	c := r.Clone(false)

	// Act
	actual := args.Map{"hasBytes": c.Length() > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone shallow", actual)
}

func Test_Result_Clone_DeepCopy(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	c := r.Clone(true)

	// Act
	actual := args.Map{"hasBytes": c.Length() > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone deep", actual)
}

func Test_Result_ClonePtr_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"isNil": r.ClonePtr(true) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr nil", actual)
}

func Test_Result_CloneIf_True_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	c := r.CloneIf(true, true)

	// Act
	actual := args.Map{"hasBytes": c.Length() > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "CloneIf true", actual)
}

func Test_Result_CloneIf_False_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	c := r.CloneIf(false, true)

	// Act
	actual := args.Map{"hasBytes": c.Length() > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "CloneIf false", actual)
}

func Test_Result_CloneError_HasError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{"hasErr": r.CloneError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError has error", actual)
}

func Test_Result_CloneError_NoError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"noErr": r.CloneError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError no error", actual)
}

// =============================================================================
// Result — NonPtr / Ptr / ToPtr / ToNonPtr / IsEqualPtr / IsEqual
// =============================================================================

func Test_Result_NonPtr_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	nr := r.NonPtr()

	// Act
	actual := args.Map{"hasErr": nr.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NonPtr nil", actual)
}

func Test_Result_IsEqualPtr_BothNil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var a, b *corejson.Result

	// Act
	actual := args.Map{"r": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr both nil", actual)
}

func Test_Result_IsEqualPtr_OneNil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	a := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": a.IsEqualPtr(nil)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr one nil", actual)
}

func Test_Result_IsEqualPtr_Same_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	a := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": a.IsEqualPtr(a)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr same", actual)
}

func Test_Result_IsEqualPtr_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.NewPtr("hello")
	b := corejson.NewPtr("hi")

	// Act
	actual := args.Map{"r": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len", actual)
}

func Test_Result_IsEqualPtr_DiffError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	a := &corejson.Result{Bytes: []byte("x"), Error: errors.New("a")}
	b := &corejson.Result{Bytes: []byte("x"), Error: errors.New("b")}

	// Act
	actual := args.Map{"r": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr diff error", actual)
}

func Test_Result_IsEqualPtr_DiffTypeName(t *testing.T) {
	// Arrange
	a := &corejson.Result{Bytes: []byte("x"), TypeName: "a"}
	b := &corejson.Result{Bytes: []byte("x"), TypeName: "b"}

	// Act
	actual := args.Map{"r": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr diff type name", actual)
}

func Test_Result_IsEqualPtr_SameContent(t *testing.T) {
	// Arrange
	a := corejson.NewPtr("hello")
	b := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr same content", actual)
}

func Test_Result_IsEqual_Same(t *testing.T) {
	// Arrange
	a := corejson.New("hello")
	b := corejson.New("hello")

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEqual same", actual)
}

func Test_Result_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.New("hello")
	b := corejson.New("hi")

	// Act
	actual := args.Map{"r": a.IsEqual(b)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsEqual diff len", actual)
}

// =============================================================================
// Result — CombineErrorWithRefString / CombineErrorWithRefError
// =============================================================================

func Test_Result_CombineErrorWithRefString_NoError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"r": r.CombineErrorWithRefString("ref1")}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString no error", actual)
}

func Test_Result_CombineErrorWithRefString_HasError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	s := r.CombineErrorWithRefString("ref1", "ref2")

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString has error", actual)
}

func Test_Result_CombineErrorWithRefError_NoError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{"noErr": r.CombineErrorWithRefError("ref1") == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError no error", actual)
}

func Test_Result_CombineErrorWithRefError_HasError_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("fail")}
	err := r.CombineErrorWithRefError("ref1")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError has error", actual)
}

// =============================================================================
// Result — BytesError / Dispose / AsJsonContractsBinder / AsJsoner / AsJsonParseSelfInjector
// =============================================================================

func Test_Result_BytesError_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"isNil": r.BytesError() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BytesError nil", actual)
}

func Test_Result_BytesError_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	be := r.BytesError()

	// Act
	actual := args.Map{"notNil": be != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesError valid", actual)
}

func Test_Result_Dispose_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")
	r.Dispose()

	// Act
	actual := args.Map{"empty": r.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)
}

func Test_Result_Dispose_Nil_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	var r *corejson.Result
	r.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose nil", actual)
}

func Test_Result_AsJsonContractsBinder_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	binder := r.AsJsonContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder", actual)
}

func Test_Result_AsJsoner_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	jsoner := r.AsJsoner()

	// Act
	actual := args.Map{"notNil": jsoner != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsoner", actual)
}

func Test_Result_AsJsonParseSelfInjector_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	inj := r.AsJsonParseSelfInjector()

	// Act
	actual := args.Map{"notNil": inj != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector", actual)
}

func Test_Result_JsonParseSelfInject_FromResultMapResultBranc(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	source := corejson.NewPtr(r)
	err := r.JsonParseSelfInject(source)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
}
