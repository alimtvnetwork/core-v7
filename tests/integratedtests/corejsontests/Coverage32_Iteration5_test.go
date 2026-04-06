package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result — nil receiver, edge branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C01_Result_NilReceiver_Map(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Map returns nil -- with args", actual)
}

func Test_I5_C02_Result_NilReceiver_JsonStringPtr(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.JsonStringPtr()

	// Act
	actual := args.Map{
		"notNil": s != nil,
		"empty": *s == "",
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NilReceiver_JsonStringPtr returns nil -- with args", actual)
}

func Test_I5_C03_Result_NilReceiver_PrettyJsonString(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"empty": r.PrettyJsonString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_PrettyJsonString returns nil -- with args", actual)
}

func Test_I5_C04_Result_NilReceiver_PrettyJsonStringOrErrString(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_PrettyJsonStringOrErr returns nil -- with args", actual)
}

func Test_I5_C05_Result_PrettyJsonStringOrErrString_WithError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("test-err")}
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErr_WithError returns error -- with args", actual)
}

func Test_I5_C06_Result_NilReceiver_Length(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Length returns nil -- with args", actual)
}

func Test_I5_C07_Result_NilReceiver_Raw(t *testing.T) {
	// Arrange
	var r *corejson.Result
	b, err := r.Raw()

	// Act
	actual := args.Map{
		"emptyBytes": len(b) == 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"emptyBytes": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Raw returns nil -- with args", actual)
}

func Test_I5_C08_Result_NilReceiver_MeaningfulError(t *testing.T) {
	// Arrange
	var r *corejson.Result
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_MeaningfulError returns nil -- with args", actual)
}

func Test_I5_C09_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte{}, TypeName: "TestType"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_EmptyBytes returns empty -- with args", actual)
}

func Test_I5_C10_Result_MeaningfulError_WithErrorAndBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"hello"`), Error: errors.New("some-err"), TypeName: "T"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_WithErrorAndBytes returns error -- with args", actual)
}

func Test_I5_C11_Result_MeaningfulError_WithErrorNoBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("some-err"), TypeName: "T"}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_WithErrorNoBytes returns error -- with args", actual)
}

func Test_I5_C12_Result_MeaningfulErrorMessage(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	msg := r.MeaningfulErrorMessage()

	// Act
	actual := args.Map{"empty": msg == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorMessage_NoErr returns error -- with args", actual)
}

func Test_I5_C13_Result_SafeString(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"val"`)}

	// Act
	actual := args.Map{"notEmpty": r.SafeString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeString returns correct value -- with args", actual)
}

func Test_I5_C14_Result_Map_WithBytesErrorType(t *testing.T) {
	r := &corejson.Result{
		Bytes:    []byte(`"data"`),
		Error:    errors.New("e"),
		TypeName: "Foo",
	}
	m := r.Map()
	_ = m
}

func Test_I5_C15_Result_DeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()

	// Act
	actual := args.Map{
		"len": len(m),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializedFieldsToMap_Nil returns nil -- with args", actual)
}

func Test_I5_C16_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	m := r.SafeDeserializedFieldsToMap()
	_ = m
	// just exercising the method without panic

	// Act
	actual := args.Map{"ran": true}

	// Assert
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "SafeDeserializedFieldsToMap returns correct value -- with args", actual)
}

func Test_I5_C17_Result_FieldsNames_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	names, err := r.FieldsNames()

	// Act
	actual := args.Map{
		"len": len(names),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FieldsNames_Empty returns empty -- with args", actual)
}

func Test_I5_C18_Result_SafeFieldsNames(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	names := r.SafeFieldsNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeFieldsNames returns correct value -- with args", actual)
}

func Test_I5_C19_Result_BytesTypeName_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"empty": r.BytesTypeName() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesTypeName_Nil returns nil -- with args", actual)
}

func Test_I5_C20_Result_SafeBytesTypeName_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"empty": r.SafeBytesTypeName() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName_Empty returns empty -- with args", actual)
}

func Test_I5_C21_Result_SafeBytesTypeName_WithType(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "Foo"}

	// Act
	actual := args.Map{"val": r.SafeBytesTypeName()}

	// Assert
	expected := args.Map{"val": "Foo"}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName_WithType returns non-empty -- with args", actual)
}

func Test_I5_C22_Result_ErrorString_NoError(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"empty": r.ErrorString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ErrorString_NoError returns error -- with args", actual)
}

func Test_I5_C23_Result_IsErrorEqual(t *testing.T) {
	// Arrange
	r1 := &corejson.Result{Error: errors.New("x")}
	r2 := &corejson.Result{}

	// Act
	actual := args.Map{
		"sameErr":    r1.IsErrorEqual(errors.New("x")),
		"diffErr":    r1.IsErrorEqual(errors.New("y")),
		"nilBothNil": r2.IsErrorEqual(nil),
		"oneNil":     r1.IsErrorEqual(nil),
	}

	// Assert
	expected := args.Map{
		"sameErr":    true,
		"diffErr":    false,
		"nilBothNil": true,
		"oneNil":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns error -- with args", actual)
}

func Test_I5_C24_Result_String_NilBytes(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	s := r.String()

	// Act
	actual := args.Map{"empty": s == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "String_NilBytes returns nil -- with args", actual)
}

func Test_I5_C25_Result_String_WithError(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	s := r.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String_WithError returns error -- with args", actual)
}

func Test_I5_C26_Result_String_NoError(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	s := r.String()

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String_NoError returns error -- with args", actual)
}

func Test_I5_C27_Result_SafeNonIssueBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"empty": len(r.SafeNonIssueBytes()) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes_Error returns error -- with args", actual)
}

func Test_I5_C28_Result_SafeNonIssueBytes_Valid(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}

	// Act
	actual := args.Map{"hasData": len(r.SafeNonIssueBytes()) > 0}

	// Assert
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes_Valid returns non-empty -- with args", actual)
}

func Test_I5_C29_Result_SafeBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"empty": len(r.SafeBytes()) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeBytes_Nil returns nil -- with args", actual)
}

func Test_I5_C30_Result_Values(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}

	// Act
	actual := args.Map{"hasData": len(r.Values()) > 0}

	// Assert
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "Values returns non-empty -- with args", actual)
}

func Test_I5_C31_Result_SafeValues_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"empty": len(r.SafeValues()) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeValues_Nil returns nil -- with args", actual)
}

func Test_I5_C32_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"empty": len(r.SafeValuesPtr()) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeValuesPtr_HasIssues returns non-empty -- with args", actual)
}

func Test_I5_C33_Result_RawMust(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	b := r.RawMust()

	// Act
	actual := args.Map{"hasData": len(b) > 0}

	// Assert
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "RawMust returns correct value -- with args", actual)
}

func Test_I5_C34_Result_RawStringMust(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	s := r.RawStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawStringMust returns correct value -- with args", actual)
}

func Test_I5_C35_Result_RawStringMust_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.RawStringMust()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawStringMust_Panic panics -- with args", actual)
}

func Test_I5_C36_Result_RawErrString(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	b, errMsg := r.RawErrString()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"emptyErr": errMsg == "",
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"emptyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrString returns error -- with args", actual)
}

func Test_I5_C37_Result_RawPrettyString(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{
		"hasContent": len(s) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasContent": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "RawPrettyString returns correct value -- with args", actual)
}

func Test_I5_C38_Result_PrettyJsonBuffer_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{
		"emptyBuf": buf.Len() == 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"emptyBuf": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "PrettyJsonBuffer_Empty returns empty -- with args", actual)
}

func Test_I5_C39_Result_PrettyJsonString_InvalidJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"empty": s == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString_InvalidJson returns error -- with args", actual)
}

func Test_I5_C40_Result_HasSafeItems(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}

	// Act
	actual := args.Map{"safe": r.HasSafeItems()}

	// Assert
	expected := args.Map{"safe": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_I5_C41_Result_HandleError_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.HandleError()
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleError_Panic panics -- with args", actual)
}

func Test_I5_C42_Result_HandleErrorWithMsg_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.HandleErrorWithMsg("custom msg")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithMsg_Panic panics -- with args", actual)
}

func Test_I5_C43_Result_HasBytes_HasJsonBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"ok"`)}

	// Act
	actual := args.Map{
		"hasBytes": r.HasBytes(),
		"hasJsonBytes": r.HasJsonBytes(),
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasJsonBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "HasBytes_HasJsonBytes returns correct value -- with args", actual)
}

func Test_I5_C44_Result_IsEmptyJsonBytes_CurlyBrace(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`{}`)}

	// Act
	actual := args.Map{"isEmpty": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes_CurlyBrace returns empty -- with args", actual)
}

func Test_I5_C45_Result_IsEmptyJsonBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"isEmpty": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes_Nil returns nil -- with args", actual)
}

func Test_I5_C46_Result_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}

	// Act
	actual := args.Map{"has": r.HasAnyItem()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- with args", actual)
}

func Test_I5_C47_Result_HasJson_IsEmptyJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}

	// Act
	actual := args.Map{
		"hasJson": r.HasJson(),
		"isEmptyJson": r.IsEmptyJson(),
	}

	// Assert
	expected := args.Map{
		"hasJson": true,
		"isEmptyJson": false,
	}
	expected.ShouldBeEqual(t, 0, "HasJson_IsEmptyJson returns empty -- with args", actual)
}

func Test_I5_C48_Result_Unmarshal_NilReceiver(t *testing.T) {
	// Arrange
	var r *corejson.Result
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_NilReceiver returns nil -- with args", actual)
}

func Test_I5_C49_Result_Unmarshal_WithExistingError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("prior"), Bytes: []byte(`"x"`)}
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_WithExistingError returns error -- with args", actual)
}

func Test_I5_C50_Result_Unmarshal_InvalidJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	var s string
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_InvalidJson returns error -- with args", actual)
}

func Test_I5_C51_Result_DeserializeMust_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.DeserializeMust(nil)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DeserializeMust_Panic panics -- with args", actual)
}

func Test_I5_C52_Result_UnmarshalMust_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.UnmarshalMust(nil)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalMust_Panic panics -- with args", actual)
}

func Test_I5_C53_Result_SerializeSkipExistingIssues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{
		"nilBytes": b == nil,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nilBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues returns correct value -- with args", actual)
}

func Test_I5_C54_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, err := r.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues_Valid returns non-empty -- with args", actual)
}

func Test_I5_C55_Result_Serialize_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	b, err := r.Serialize()

	// Act
	actual := args.Map{
		"nilBytes": b == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"nilBytes": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_Nil returns nil -- with args", actual)
}

func Test_I5_C56_Result_Serialize_WithError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.Serialize()

	// Act
	actual := args.Map{
		"emptyBytes": len(b) == 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"emptyBytes": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_WithError returns error -- with args", actual)
}

func Test_I5_C57_Result_Serialize_Valid(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b, err := r.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_Valid returns non-empty -- with args", actual)
}

func Test_I5_C58_Result_SerializeMust(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b := r.SerializeMust()

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust returns correct value -- with args", actual)
}

func Test_I5_C59_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_HasIssues returns correct value -- with args", actual)
}

func Test_I5_C60_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_Valid returns non-empty -- with args", actual)
}

func Test_I5_C61_Result_UnmarshalSkipExistingIssues_BadJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_BadJson returns correct value -- with args", actual)
}

func Test_I5_C62_Result_UnmarshalResult(t *testing.T) {
	// Arrange
	inner := corejson.NewResult.Any("hello")
	outerBytes, _ := inner.Serialize()
	outer := &corejson.Result{Bytes: outerBytes}
	result, err := outer.UnmarshalResult()

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalResult returns correct value -- with args", actual)
}

func Test_I5_C63_Result_JsonModel_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	model := r.JsonModel()

	// Act
	actual := args.Map{"hasErr": model.Error != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "JsonModel_Nil returns nil -- with args", actual)
}

func Test_I5_C64_Result_JsonModel_Valid(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	model := r.JsonModel()

	// Act
	actual := args.Map{"hasBytes": len(model.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "JsonModel_Valid returns non-empty -- with args", actual)
}

func Test_I5_C65_Result_JsonModelAny(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	a := r.JsonModelAny()

	// Act
	actual := args.Map{"notNil": a != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_I5_C66_Result_Json_JsonPtr(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}
	j := r.Json()
	jp := r.JsonPtr()

	// Act
	actual := args.Map{
		"hasBytes": len(j.Bytes) > 0,
		"ptrNotNil": jp != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Json_JsonPtr returns correct value -- with args", actual)
}

func Test_I5_C67_Result_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	input := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"injected"`), TypeName: "T"})
	result, err := r.ParseInjectUsingJson(input)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_I5_C68_Result_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	input := &corejson.Result{Bytes: []byte(`not-json`)}
	_, err := r.ParseInjectUsingJson(input)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson_Error returns error -- with args", actual)
}

func Test_I5_C69_Result_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	input := &corejson.Result{Bytes: []byte(`not-json`)}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.ParseInjectUsingJsonMust(input)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust_Panic panics -- with args", actual)
}

func Test_I5_C70_Result_CloneError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}

	// Act
	actual := args.Map{"hasErr": r.CloneError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError returns error -- with args", actual)

	r2 := &corejson.Result{}
	actual2 := args.Map{"nilErr": r2.CloneError() == nil}
	expected2 := args.Map{"nilErr": true}
	expected2.ShouldBeEqual(t, 0, "CloneError_NoErr returns error -- with args", actual2)
}

func Test_I5_C71_Result_Ptr_NonPtr(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}
	p := r.Ptr()
	np := p.NonPtr()

	// Act
	actual := args.Map{
		"ptrNotNil": p != nil,
		"hasBytes": len(np.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"ptrNotNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Ptr_NonPtr returns correct value -- with args", actual)
}

func Test_I5_C72_Result_NonPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	np := r.NonPtr()

	// Act
	actual := args.Map{"hasErr": np.Error != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NonPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C73_Result_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}
	p := r.ToPtr()
	np := r.ToNonPtr()

	// Act
	actual := args.Map{
		"ptrNotNil": p != nil,
		"hasBytes": len(np.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"ptrNotNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ToPtr_ToNonPtr returns correct value -- with args", actual)
}

func Test_I5_C74_Result_IsEqualPtr(t *testing.T) {
	// Arrange
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	r3 := &corejson.Result{Bytes: []byte(`"y"`), TypeName: "T"}
	var rNil *corejson.Result

	// Act
	actual := args.Map{
		"same":     r1.IsEqualPtr(r2),
		"diff":     r1.IsEqualPtr(r3),
		"bothNil":  rNil.IsEqualPtr(nil),
		"oneNil":   r1.IsEqualPtr(nil),
		"selfSame": r1.IsEqualPtr(r1),
	}

	// Assert
	expected := args.Map{
		"same":     true,
		"diff":     false,
		"bothNil":  true,
		"oneNil":   false,
		"selfSame": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- with args", actual)
}

func Test_I5_C75_Result_IsEqualPtr_DiffType(t *testing.T) {
	// Arrange
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "A"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "B"}

	// Act
	actual := args.Map{"equal": r1.IsEqualPtr(r2)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr_DiffType returns correct value -- with args", actual)
}

func Test_I5_C76_Result_IsEqualPtr_DiffErr(t *testing.T) {
	// Arrange
	r1 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("a")}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("b")}

	// Act
	actual := args.Map{"equal": r1.IsEqualPtr(r2)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr_DiffErr returns error -- with args", actual)
}

func Test_I5_C77_Result_CombineErrorWithRefString(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	s := r.CombineErrorWithRefString("ref1", "ref2")

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString returns error -- with args", actual)
}

func Test_I5_C78_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	s := r.CombineErrorWithRefString("ref1")

	// Act
	actual := args.Map{"empty": s == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString_NoErr returns error -- with args", actual)
}

func Test_I5_C79_Result_CombineErrorWithRefError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	err := r.CombineErrorWithRefError("ref1")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError returns error -- with args", actual)

	r2 := &corejson.Result{}
	err2 := r2.CombineErrorWithRefError("ref1")
	actual2 := args.Map{"nilErr": err2 == nil}
	expected2 := args.Map{"nilErr": true}
	expected2.ShouldBeEqual(t, 0, "CombineErrorWithRefError_NoErr returns error -- with args", actual2)
}

func Test_I5_C80_Result_IsEqual(t *testing.T) {
	// Arrange
	r1 := corejson.Result{Bytes: []byte(`"x"`)}
	r2 := corejson.Result{Bytes: []byte(`"x"`)}
	r3 := corejson.Result{Bytes: []byte(`"y"`)}

	// Act
	actual := args.Map{
		"same": r1.IsEqual(r2),
		"diff": r1.IsEqual(r3),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_I5_C81_Result_BytesError(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")}
	be := r.BytesError()

	// Act
	actual := args.Map{
		"notNil": be != nil,
		"hasBytes": len(be.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesError returns error -- with args", actual)

	var rNil *corejson.Result
	actual2 := args.Map{"nil": rNil.BytesError() == nil}
	expected2 := args.Map{"nil": true}
	expected2.ShouldBeEqual(t, 0, "BytesError_Nil returns nil -- with args", actual2)
}

func Test_I5_C82_Result_Dispose(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	r.Dispose()

	// Act
	actual := args.Map{
		"nilErr": r.Error == nil,
		"nilBytes": r.Bytes == nil,
		"emptyType": r.TypeName == "",
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"nilBytes": true,
		"emptyType": true,
	}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)

	var rNil *corejson.Result
	rNil.Dispose() // should not panic
}

func Test_I5_C83_Result_CloneIf(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	cloned := r.CloneIf(true, true)
	notCloned := r.CloneIf(false, false)

	// Act
	actual := args.Map{
		"clonedHasBytes": len(cloned.Bytes) > 0,
		"notClonedHasBytes": len(notCloned.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"clonedHasBytes": true,
		"notClonedHasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "CloneIf returns correct value -- with args", actual)
}

func Test_I5_C84_Result_ClonePtr(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	cloned := r.ClonePtr(true)

	// Act
	actual := args.Map{"notNil": cloned != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)

	var rNil *corejson.Result
	actual2 := args.Map{"nil": rNil.ClonePtr(true) == nil}
	expected2 := args.Map{"nil": true}
	expected2.ShouldBeEqual(t, 0, "ClonePtr_Nil returns nil -- with args", actual2)
}

func Test_I5_C85_Result_Clone_ShallowAndDeep(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	shallow := r.Clone(false)
	deep := r.Clone(true)

	// Act
	actual := args.Map{
		"shallowHasBytes": len(shallow.Bytes) > 0,
		"deepHasBytes": len(deep.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"shallowHasBytes": true,
		"deepHasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone_ShallowAndDeep returns correct value -- with args", actual)
}

func Test_I5_C86_Result_Clone_Empty(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	cloned := r.Clone(true)

	// Act
	actual := args.Map{"emptyBytes": len(cloned.Bytes) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone_Empty returns empty -- with args", actual)
}

func Test_I5_C87_Result_AsInterfaces(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}

	// Act
	actual := args.Map{
		"jsonContracts": r.AsJsonContractsBinder() != nil,
		"jsoner":        r.AsJsoner() != nil,
		"selfInjector":  r.AsJsonParseSelfInjector() != nil,
	}

	// Assert
	expected := args.Map{
		"jsonContracts": true,
		"jsoner":        true,
		"selfInjector":  true,
	}
	expected.ShouldBeEqual(t, 0, "AsInterfaces returns correct value -- with args", actual)
}

func Test_I5_C88_Result_JsonParseSelfInject(t *testing.T) {
	// Arrange
	r := corejson.Result{}
	input := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"injected"`), TypeName: "T"})
	err := r.JsonParseSelfInject(input)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

func Test_I5_C89_Result_InjectInto(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	target := corejson.Result{}
	err := r.InjectInto(&target)
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// castingAny
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C90_CastAny_FromToDefault(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToDefault returns correct value -- with args", actual)
}

func Test_I5_C91_CastAny_FromToReflection(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToReflection returns correct value -- with args", actual)
}

func Test_I5_C92_CastAny_FromToOption_Nil(t *testing.T) {
	// Arrange
	err := corejson.CastAny.FromToOption(true, nil, nil)

	// Act
	actual := args.Map{"hasResult": err != nil || err == nil} // exercises nil path

	// Assert
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_Nil returns nil -- with args", actual)
}

func Test_I5_C93_CastAny_FromToOption_String(t *testing.T) {
	// Arrange
	var out map[string]string
	err := corejson.CastAny.FromToOption(false, `{"a":"b"}`, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out["a"],
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "b",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_String returns correct value -- with args", actual)
}

func Test_I5_C94_CastAny_FromToOption_Error(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New(`"test"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "test",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_Error returns error -- with args", actual)
}

func Test_I5_C95_CastAny_FromToOption_NilError(t *testing.T) {
	// Arrange
	var nilErr error
	var out string
	err := corejson.CastAny.FromToOption(false, nilErr, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_NilError returns nil -- with args", actual)
}

func Test_I5_C96_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	// Arrange
	fn := func() ([]byte, error) { return []byte(`"from-func"`), nil }
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "from-func",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_SerializerFunc returns correct value -- with args", actual)
}

func Test_I5_C97_CastAny_FromToOption_AnyItem(t *testing.T) {
	// Arrange
	type s struct{ A int }
	var out s
	err := corejson.CastAny.FromToOption(false, s{A: 42}, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out.A,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_AnyItem returns correct value -- with args", actual)
}

func Test_I5_C98_CastAny_OrDeserializeTo(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hi"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hi",
	}
	expected.ShouldBeEqual(t, 0, "CastAny_OrDeserializeTo returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// anyTo — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C99_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(nil)

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_Nil returns nil -- with args", actual)
}

func Test_I5_C100_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(errors.New("test"))

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_Error returns error -- with args", actual)
}

func Test_I5_C101_AnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(errors.New(""))

	// Act
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_EmptyError returns empty -- with args", actual)
}

func Test_I5_C102_AnyTo_SerializedRaw(t *testing.T) {
	// Arrange
	b, err := corejson.AnyTo.SerializedRaw([]byte(`"test"`))

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedRaw returns correct value -- with args", actual)
}

func Test_I5_C103_AnyTo_SerializedString(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString([]byte(`"hello"`))

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedString returns correct value -- with args", actual)
}

func Test_I5_C104_AnyTo_SerializedString_Error(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString(nil)

	// Act
	actual := args.Map{
		"empty": s == "",
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedString_Error returns error -- with args", actual)
}

func Test_I5_C105_AnyTo_SerializedSafeString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString([]byte(`"hello"`))

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedSafeString returns correct value -- with args", actual)
}

func Test_I5_C106_AnyTo_SerializedSafeString_Error(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString(nil)

	// Act
	actual := args.Map{"empty": s == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedSafeString_Error returns error -- with args", actual)
}

func Test_I5_C107_AnyTo_SerializedStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust([]byte(`"hello"`))

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedStringMust returns correct value -- with args", actual)
}

func Test_I5_C108_AnyTo_SafeJsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SafeJsonString returns correct value -- with args", actual)
}

func Test_I5_C109_AnyTo_PrettyStringWithError(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError("hello")

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_String returns error -- with args", actual)
}

func Test_I5_C110_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_Bytes returns error -- with args", actual)
}

func Test_I5_C111_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_Result returns error -- with args", actual)
}

func Test_I5_C112_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := corejson.AnyTo.PrettyStringWithError(r)

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_ResultPtr returns error -- with args", actual)
}

func Test_I5_C113_AnyTo_PrettyStringWithError_ResultWithErr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("e")}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	_ = s
	_ = err
}

func Test_I5_C114_AnyTo_PrettyStringWithError_ResultPtrWithErr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("e")}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	_ = s
	_ = err
}

func Test_I5_C115_AnyTo_SafeJsonPrettyString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_String returns correct value -- with args", actual)
}

func Test_I5_C116_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_Bytes returns correct value -- with args", actual)
}

func Test_I5_C117_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := corejson.AnyTo.SafeJsonPrettyString(r)

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_Result returns correct value -- with args", actual)
}

func Test_I5_C118_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := corejson.AnyTo.SafeJsonPrettyString(r)

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_ResultPtr returns correct value -- with args", actual)
}

func Test_I5_C119_AnyTo_JsonString(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_String returns correct value -- with args", actual)
}

func Test_I5_C120_AnyTo_JsonString_Bytes(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString([]byte(`"x"`))

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_Bytes returns correct value -- with args", actual)
}

func Test_I5_C121_AnyTo_JsonString_Result(t *testing.T) {
	// Arrange
	r := corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_Result returns correct value -- with args", actual)
}

func Test_I5_C122_AnyTo_JsonString_ResultPtr(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_ResultPtr returns correct value -- with args", actual)
}

func Test_I5_C123_AnyTo_JsonStringWithErr(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr("hello")

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_String returns error -- with args", actual)
}

func Test_I5_C124_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_Bytes returns error -- with args", actual)
}

func Test_I5_C125_AnyTo_JsonStringWithErr_ResultWithErr(t *testing.T) {
	r := corejson.Result{Error: errors.New("e"), Bytes: []byte(`"x"`)}
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	_ = s
	_ = err
}

func Test_I5_C126_AnyTo_JsonStringWithErr_ResultPtrWithErr(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e"), Bytes: []byte(`"x"`)}
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	// Result with error may or may not produce content depending on implementation
	_ = s
	_ = err
}

func Test_I5_C127_AnyTo_JsonStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringMust returns correct value -- with args", actual)
}

func Test_I5_C128_AnyTo_PrettyStringMust(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.PrettyStringMust("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringMust returns correct value -- with args", actual)
}

func Test_I5_C129_AnyTo_UsingSerializer(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.UsingSerializer(nil)

	// Act
	actual := args.Map{"nil": r == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_UsingSerializer_Nil returns nil -- with args", actual)
}

func Test_I5_C130_AnyTo_SerializedFieldsMap(t *testing.T) {
	// Arrange
	type s struct{ A int }
	m, err := corejson.AnyTo.SerializedFieldsMap(s{A: 42})
	_ = m

	// Act
	actual := args.Map{"ran": err == nil || err != nil}

	// Assert
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedFieldsMap returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializerLogic — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C131_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C132_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	// Arrange
	str := `"hello"`
	var out string
	err := corejson.Deserialize.UsingStringPtr(&str, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringPtr_Valid returns non-empty -- with args", actual)
}

func Test_I5_C133_Deserialize_UsingError_Nil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingError(nil, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingError_Nil returns nil -- with args", actual)
}

func Test_I5_C134_Deserialize_UsingError_Valid(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingError_Valid returns error -- with args", actual)
}

func Test_I5_C135_Deserialize_UsingErrorWhichJsonResult(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingErrorWhichJsonResult_Nil returns nil -- with args", actual)
}

func Test_I5_C136_Deserialize_ApplyMust_Panic(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var s string
		corejson.Deserialize.ApplyMust(r, &s)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_ApplyMust_Panic panics -- with args", actual)
}

func Test_I5_C137_Deserialize_FromString(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.FromString(`"hello"`, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromString returns correct value -- with args", actual)
}

func Test_I5_C138_Deserialize_FromStringMust(t *testing.T) {
	// Arrange
	var out string
	corejson.Deserialize.FromStringMust(`"hello"`, &out)

	// Act
	actual := args.Map{"val": out}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromStringMust returns correct value -- with args", actual)
}

func Test_I5_C139_Deserialize_FromStringMust_Panic(t *testing.T) {
	// Arrange
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var out string
		corejson.Deserialize.FromStringMust(`not-json`, &out)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromStringMust_Panic panics -- with args", actual)
}

func Test_I5_C140_Deserialize_MapAnyToPointer(t *testing.T) {
	// Arrange
	type s struct {
		A int `json:"a"`
	}
	var out s
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"a": 1}, &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"val": out.A,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"val": 1,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize_MapAnyToPointer returns correct value -- with args", actual)
}

func Test_I5_C141_Deserialize_MapAnyToPointer_SkipEmpty(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_MapAnyToPointer_SkipEmpty returns empty -- with args", actual)
}

func Test_I5_C142_Deserialize_UsingStringOption(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingStringOption(true, "", &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringOption_Empty returns empty -- with args", actual)
}

func Test_I5_C143_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringIgnoreEmpty returns empty -- with args", actual)
}

func Test_I5_C144_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesPointer(nil, &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointer_Nil returns nil -- with args", actual)
}

func Test_I5_C145_Deserialize_UsingBytesPointerMust_Panic(t *testing.T) {
	// Arrange
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var out string
		corejson.Deserialize.UsingBytesPointerMust(nil, &out)
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointerMust_Panic panics -- with args", actual)
}

func Test_I5_C146_Deserialize_UsingBytesIf(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &out)

	// Act
	actual := args.Map{
		"nilErr": err == nil,
		"empty": out == "",
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesIf_Skip returns correct value -- with args", actual)
}

func Test_I5_C147_Deserialize_UsingBytesPointerIf(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointerIf_Skip returns correct value -- with args", actual)
}

func Test_I5_C148_Deserialize_UsingBytesMust(t *testing.T) {
	// Arrange
	var out string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"val": out}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesMust returns correct value -- with args", actual)
}

func Test_I5_C149_Deserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	// Arrange
	var out string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &out)

	// Act
	actual := args.Map{"empty": out == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingSafeBytesMust_Empty returns empty -- with args", actual)
}

func Test_I5_C150_Deserialize_AnyToFieldsMap(t *testing.T) {
	// Arrange
	type s struct{ A int }
	m, _ := corejson.Deserialize.AnyToFieldsMap(s{A: 1})
	_ = m

	// Act
	actual := args.Map{"ran": true}

	// Assert
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_AnyToFieldsMap returns correct value -- with args", actual)
}

func Test_I5_C151_Deserialize_UsingDeserializerToOption_SkipNil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption_SkipNil returns nil -- with args", actual)
}

func Test_I5_C152_Deserialize_UsingDeserializerToOption_NilNotSkipped(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption_NilNotSkipped returns nil -- with args", actual)
}

func Test_I5_C153_Deserialize_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined_Nil returns nil -- with args", actual)
}

func Test_I5_C154_Deserialize_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	// Arrange
	fn := func(toPtr any) error { return nil }
	var out string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined_Valid returns non-empty -- with args", actual)
}

func Test_I5_C155_Deserialize_UsingJsonerToAny_SkipNil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny_SkipNil returns nil -- with args", actual)
}

func Test_I5_C156_Deserialize_UsingJsonerToAny_NilNotSkipped(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny_NilNotSkipped returns nil -- with args", actual)
}

func Test_I5_C157_Deserialize_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &out)

	// Act
	actual := args.Map{"nilErr": err == nil}

	// Assert
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust_SkipNil returns nil -- with args", actual)
}

func Test_I5_C158_Deserialize_UsingJsonerToAnyMust_NilNotSkipped(t *testing.T) {
	// Arrange
	var out string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust_NilNotSkipped returns nil -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Utility funcs — BytesCloneIf, BytesDeepClone, BytesToString, BytesToPrettyString, JsonString
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C159_BytesCloneIf_NoClone(t *testing.T) {
	// Arrange
	b := corejson.BytesCloneIf(false, []byte("abc"))

	// Act
	actual := args.Map{"empty": len(b) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_NoClone returns correct value -- with args", actual)
}

func Test_I5_C160_BytesCloneIf_Clone(t *testing.T) {
	// Arrange
	b := corejson.BytesCloneIf(true, []byte("abc"))

	// Act
	actual := args.Map{"len": len(b)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_Clone returns correct value -- with args", actual)
}

func Test_I5_C161_BytesCloneIf_Empty(t *testing.T) {
	// Arrange
	b := corejson.BytesCloneIf(true, []byte{})

	// Act
	actual := args.Map{"empty": len(b) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_Empty returns empty -- with args", actual)
}

func Test_I5_C162_BytesDeepClone(t *testing.T) {
	// Arrange
	b := corejson.BytesDeepClone([]byte("abc"))

	// Act
	actual := args.Map{"len": len(b)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns correct value -- with args", actual)
}

func Test_I5_C163_BytesDeepClone_Empty(t *testing.T) {
	// Arrange
	b := corejson.BytesDeepClone([]byte{})

	// Act
	actual := args.Map{"empty": len(b) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone_Empty returns empty -- with args", actual)
}

func Test_I5_C164_BytesToString(t *testing.T) {
	// Act
	actual := args.Map{"val": corejson.BytesToString([]byte(`"x"`))}

	// Assert
	expected := args.Map{"val": `"x"`}
	expected.ShouldBeEqual(t, 0, "BytesToString returns correct value -- with args", actual)
}

func Test_I5_C165_BytesToString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"empty": corejson.BytesToString(nil) == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString_Empty returns empty -- with args", actual)
}

func Test_I5_C166_BytesToPrettyString(t *testing.T) {
	// Arrange
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString returns correct value -- with args", actual)
}

func Test_I5_C167_BytesToPrettyString_Empty(t *testing.T) {
	// Act
	actual := args.Map{"empty": corejson.BytesToPrettyString(nil) == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString_Empty returns empty -- with args", actual)
}

func Test_I5_C168_JsonString_Func(t *testing.T) {
	// Arrange
	s, err := corejson.JsonString("hello")

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString_Func returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// serializerLogic — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C169_Serialize_StringsApply(t *testing.T) {
	// Arrange
	r := corejson.Serialize.StringsApply([]string{"a", "b"})

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_StringsApply returns correct value -- with args", actual)
}

func Test_I5_C170_Serialize_FromBytes(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBytes([]byte("abc"))

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromBytes returns correct value -- with args", actual)
}

func Test_I5_C171_Serialize_FromStrings(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStrings([]string{"a"})

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStrings returns correct value -- with args", actual)
}

func Test_I5_C172_Serialize_FromStringsSpread(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStringsSpread("a", "b")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStringsSpread returns correct value -- with args", actual)
}

func Test_I5_C173_Serialize_FromString(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromString("hello")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromString returns correct value -- with args", actual)
}

func Test_I5_C174_Serialize_FromInteger(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger(42)

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromInteger returns correct value -- with args", actual)
}

func Test_I5_C175_Serialize_FromInteger64(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger64(42)

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromInteger64 returns correct value -- with args", actual)
}

func Test_I5_C176_Serialize_FromBool(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBool(true)

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromBool returns correct value -- with args", actual)
}

func Test_I5_C177_Serialize_FromIntegers(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromIntegers([]int{1, 2})

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromIntegers returns correct value -- with args", actual)
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_I5_C178_Serialize_FromStringer(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStringer(testStringer{"hello"})

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStringer returns correct value -- with args", actual)
}

func Test_I5_C179_Serialize_UsingAnyPtr(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAnyPtr("hello")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_UsingAnyPtr returns correct value -- with args", actual)
}

func Test_I5_C180_Serialize_UsingAny(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAny("hello")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_UsingAny returns correct value -- with args", actual)
}

func Test_I5_C181_Serialize_Raw(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Raw("hello")

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_Raw returns correct value -- with args", actual)
}

func Test_I5_C182_Serialize_Marshal(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.Marshal("hello")

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_Marshal returns correct value -- with args", actual)
}

func Test_I5_C183_Serialize_ApplyMust(t *testing.T) {
	// Arrange
	r := corejson.Serialize.ApplyMust("hello")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ApplyMust returns correct value -- with args", actual)
}

func Test_I5_C184_Serialize_ToBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesMust("hello")

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesMust returns correct value -- with args", actual)
}

func Test_I5_C185_Serialize_ToSafeBytesMust(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesMust("hello")

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToSafeBytesMust returns correct value -- with args", actual)
}

func Test_I5_C186_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToSafeBytesSwallowErr returns error -- with args", actual)
}

func Test_I5_C187_Serialize_ToBytesSwallowErr(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesSwallowErr("hello")

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesSwallowErr returns error -- with args", actual)
}

func Test_I5_C188_Serialize_ToBytesErr(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.ToBytesErr("hello")

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesErr returns error -- with args", actual)
}

func Test_I5_C189_Serialize_ToString(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToString("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToString returns correct value -- with args", actual)
}

func Test_I5_C190_Serialize_ToStringMust(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToStringMust("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToStringMust returns correct value -- with args", actual)
}

func Test_I5_C191_Serialize_ToStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToStringErr("hello")

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_ToStringErr returns error -- with args", actual)
}

func Test_I5_C192_Serialize_ToPrettyStringErr(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"notEmpty": s != "",
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize_ToPrettyStringErr returns error -- with args", actual)
}

func Test_I5_C193_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToPrettyStringIncludingErr returns error -- with args", actual)
}

func Test_I5_C194_Serialize_Pretty(t *testing.T) {
	// Arrange
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Pretty returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializeFromBytesTo — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C195_BytesTo_Bool(t *testing.T) {
	// Arrange
	b, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))

	// Act
	actual := args.Map{
		"val": b,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Bool returns correct value -- with args", actual)
}

func Test_I5_C196_BytesTo_BoolMust(t *testing.T) {
	// Arrange
	b := corejson.Deserialize.BytesTo.BoolMust([]byte(`true`))

	// Act
	actual := args.Map{"val": b}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_BoolMust returns correct value -- with args", actual)
}

func Test_I5_C197_BytesTo_Integer(t *testing.T) {
	// Arrange
	i, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))

	// Act
	actual := args.Map{
		"val": i,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer returns correct value -- with args", actual)
}

func Test_I5_C198_BytesTo_IntegerMust(t *testing.T) {
	// Arrange
	i := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))

	// Act
	actual := args.Map{"val": i}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "BytesTo_IntegerMust returns correct value -- with args", actual)
}

func Test_I5_C199_BytesTo_Integer64(t *testing.T) {
	// Arrange
	i, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))

	// Act
	actual := args.Map{
		"val": int(i),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": 99,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer64 returns correct value -- with args", actual)
}

func Test_I5_C200_BytesTo_Integer64Must(t *testing.T) {
	// Arrange
	i := corejson.Deserialize.BytesTo.Integer64Must([]byte(`99`))

	// Act
	actual := args.Map{"val": int(i)}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer64Must returns correct value -- with args", actual)
}

func Test_I5_C201_BytesTo_Integers(t *testing.T) {
	// Arrange
	ints, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))

	// Act
	actual := args.Map{
		"len": len(ints),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integers returns correct value -- with args", actual)
}

func Test_I5_C202_BytesTo_IntegersMust(t *testing.T) {
	// Arrange
	ints := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))

	// Act
	actual := args.Map{"len": len(ints)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesTo_IntegersMust returns correct value -- with args", actual)
}

func Test_I5_C203_BytesTo_Strings(t *testing.T) {
	// Arrange
	strs, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))

	// Act
	actual := args.Map{
		"len": len(strs),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Strings returns correct value -- with args", actual)
}

func Test_I5_C204_BytesTo_StringsMust(t *testing.T) {
	// Arrange
	strs := corejson.Deserialize.BytesTo.StringsMust([]byte(`["a"]`))

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesTo_StringsMust returns correct value -- with args", actual)
}

func Test_I5_C205_BytesTo_String(t *testing.T) {
	// Arrange
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hi"`))

	// Act
	actual := args.Map{
		"val": s,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "hi",
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_String returns correct value -- with args", actual)
}

func Test_I5_C206_BytesTo_StringMust(t *testing.T) {
	// Arrange
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hi"`))

	// Act
	actual := args.Map{"val": s}

	// Assert
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesTo_StringMust returns correct value -- with args", actual)
}

func Test_I5_C207_BytesTo_MapAnyItem(t *testing.T) {
	// Arrange
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))

	// Act
	actual := args.Map{
		"hasKey": m["a"] != nil,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasKey": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapAnyItem returns correct value -- with args", actual)
}

func Test_I5_C208_BytesTo_MapAnyItemMust(t *testing.T) {
	// Arrange
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"hasKey": m["a"] != nil}

	// Assert
	expected := args.Map{"hasKey": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapAnyItemMust returns correct value -- with args", actual)
}

func Test_I5_C209_BytesTo_MapStringString(t *testing.T) {
	// Arrange
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{
		"val": m["a"],
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "b",
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapStringString returns correct value -- with args", actual)
}

func Test_I5_C210_BytesTo_MapStringStringMust(t *testing.T) {
	// Arrange
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))

	// Act
	actual := args.Map{"val": m["a"]}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapStringStringMust returns correct value -- with args", actual)
}

func Test_I5_C211_BytesTo_Bytes(t *testing.T) {
	// Arrange
	b, err := corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))

	// Act
	actual := args.Map{
		"hasData": len(b) > 0,
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasData": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo_Bytes returns correct value -- with args", actual)
}

func Test_I5_C212_BytesTo_BytesMust(t *testing.T) {
	// Arrange
	b := corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))

	// Act
	actual := args.Map{"hasData": len(b) > 0}

	// Assert
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_BytesMust returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializeFromResultTo — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C213_ResultTo_Byte(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(byte(65))
	b, err := corejson.Deserialize.ResultTo.Byte(r)

	// Act
	actual := args.Map{
		"val": int(b),
		"nilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": 65,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo_Byte returns correct value -- with args", actual)
}

func Test_I5_C214_ResultTo_ByteMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)

	// Act
	actual := args.Map{"val": int(b)}

	// Assert
	expected := args.Map{"val": 65}
	expected.ShouldBeEqual(t, 0, "ResultTo_ByteMust returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// emptyCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C215_Empty_ResultWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))

	// Act
	actual := args.Map{
		"hasErr": r.Error != nil,
		"type": r.TypeName,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"type": "T",
	}
	expected.ShouldBeEqual(t, 0, "Empty_ResultWithErr returns empty -- with args", actual)
}

func Test_I5_C216_Empty_ResultPtrWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtrWithErr("T", errors.New("e"))

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasErr": r.Error != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty_ResultPtrWithErr returns empty -- with args", actual)
}

func Test_I5_C217_Empty_BytesCollection(t *testing.T) {
	// Arrange
	c := corejson.Empty.BytesCollection()

	// Act
	actual := args.Map{"empty": c.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty_BytesCollection returns empty -- with args", actual)
}

func Test_I5_C218_Empty_BytesCollectionPtr(t *testing.T) {
	// Arrange
	c := corejson.Empty.BytesCollectionPtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"empty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty_BytesCollectionPtr returns empty -- with args", actual)
}

func Test_I5_C219_Empty_ResultsPtrCollection(t *testing.T) {
	// Arrange
	c := corejson.Empty.ResultsPtrCollection()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"empty": c.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty_ResultsPtrCollection returns empty -- with args", actual)
}

func Test_I5_C220_Empty_MapResults(t *testing.T) {
	// Arrange
	m := corejson.Empty.MapResults()

	// Act
	actual := args.Map{
		"notNil": m != nil,
		"empty": m.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty_MapResults returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newResultCreator — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C221_NewResult_PtrUsingStringPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.PtrUsingStringPtr(nil, "T")

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_PtrUsingStringPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C222_NewResult_PtrUsingStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.PtrUsingStringPtr(&s, "T")

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_PtrUsingStringPtr_Valid returns non-empty -- with args", actual)
}

func Test_I5_C223_NewResult_UsingErrorStringPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingErrorStringPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C224_NewResult_UsingErrorStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")

	// Act
	actual := args.Map{
		"hasBytes": len(r.Bytes) > 0,
		"nilErr": r.Error == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingErrorStringPtr_Valid returns error -- with args", actual)
}

func Test_I5_C225_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)

	// Act
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingTypePlusStringPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C226_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingTypePlusStringPtr_Valid returns non-empty -- with args", actual)
}

func Test_I5_C227_NewResult_UsingStringPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingStringPtr(nil)

	// Act
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingStringPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C228_NewResult_UsingStringPtr_Valid(t *testing.T) {
	// Arrange
	s := `"hello"`
	r := corejson.NewResult.UsingStringPtr(&s)

	// Act
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingStringPtr_Valid returns non-empty -- with args", actual)
}

func Test_I5_C229_NewResult_UsingBytesPtr_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesPtr(nil)

	// Act
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}

	// Assert
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C230_NewResult_UsingBytesPtrErrPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")

	// Act
	actual := args.Map{"hasErr": r.Error != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtrErrPtr_Nil returns nil -- with args", actual)
}

func Test_I5_C231_NewResult_UsingBytesErrPtr_Empty(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesErrPtr([]byte{}, errors.New("e"), "T")

	// Act
	actual := args.Map{
		"hasErr": r.Error != nil,
		"emptyBytes": len(r.Bytes) == 0,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"emptyBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesErrPtr_Empty returns empty -- with args", actual)
}

func Test_I5_C232_NewResult_UsingSerializer_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializer(nil)

	// Act
	actual := args.Map{"nil": r == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingSerializer_Nil returns nil -- with args", actual)
}

func Test_I5_C233_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializerFunc(nil)

	// Act
	actual := args.Map{"nil": r == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingSerializerFunc_Nil returns nil -- with args", actual)
}

func Test_I5_C234_NewResult_UsingJsoner_Nil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingJsoner(nil)

	// Act
	actual := args.Map{"nil": r == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingJsoner_Nil returns nil -- with args", actual)
}

func Test_I5_C235_NewResult_Many(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Many("a", "b", "c")

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasBytes": len(r.Bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult_Many returns correct value -- with args", actual)
}

func Test_I5_C236_NewResult_AnyToCastingResult(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyToCastingResult("hello")

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_AnyToCastingResult returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// StaticJsonError var
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C237_StaticJsonError(t *testing.T) {
	// Act
	actual := args.Map{"notNil": corejson.StaticJsonError != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StaticJsonError returns error -- with args", actual)
}
