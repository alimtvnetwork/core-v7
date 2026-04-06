package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result — Core Methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_Result_Map_Nil(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Result.Map returns nil -- nil", actual)
}

func Test_Cov8_Result_Map_WithAll(t *testing.T) {
	// IsEmptyJsonBytes returns true when Error is set, so Map() skips Bytes key
	r := &corejson.Result{
		Bytes:    []byte(`{"a":1}`),
		Error:    errors.New("test"),
		TypeName: "TestType",
	}
	m := r.Map()
	actual := args.Map{"hasBytes": m["Bytes"] != "", "hasError": m["Error"] != "", "hasType": m["Type"] != ""}
	expected := args.Map{"hasBytes": false, "hasError": true, "hasType": true}
	expected.ShouldBeEqual(t, 0, "Result.Map returns non-empty -- with all", actual)
}

func Test_Cov8_Result_Map_Empty(t *testing.T) {
	r := &corejson.Result{}
	m := r.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Result.Map returns empty -- empty", actual)
}

func Test_Cov8_Result_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"val": r.BytesTypeName()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "BytesTypeName returns nil -- nil", actual)
}

func Test_Cov8_Result_BytesTypeName_Valid(t *testing.T) {
	r := &corejson.Result{TypeName: "MyType"}
	actual := args.Map{"val": r.BytesTypeName()}
	expected := args.Map{"val": "MyType"}
	expected.ShouldBeEqual(t, 0, "BytesTypeName returns non-empty -- valid", actual)
}

func Test_Cov8_Result_SafeBytesTypeName_Empty(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.SafeBytesTypeName()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName returns empty -- empty", actual)
}

func Test_Cov8_Result_SafeBytesTypeName_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`), TypeName: "MyType"}
	actual := args.Map{"val": r.SafeBytesTypeName()}
	expected := args.Map{"val": "MyType"}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName returns non-empty -- valid", actual)
}

func Test_Cov8_Result_JsonStringPtr_Nil(t *testing.T) {
	var r *corejson.Result
	val := r.JsonStringPtr()
	actual := args.Map{"empty": *val == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr returns nil -- nil", actual)
}

func Test_Cov8_Result_JsonStringPtr_WithBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	val := r.JsonStringPtr()
	actual := args.Map{"val": *val}
	expected := args.Map{"val": `"hello"`}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr returns non-empty -- with bytes", actual)
}

func Test_Cov8_Result_JsonStringPtr_NoBytes(t *testing.T) {
	r := &corejson.Result{}
	val := r.JsonStringPtr()
	actual := args.Map{"empty": *val == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr returns empty -- no bytes", actual)
}

func Test_Cov8_Result_JsonStringPtr_Cached(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	p1 := r.JsonStringPtr()
	p2 := r.JsonStringPtr()
	actual := args.Map{"same": p1 == p2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "JsonStringPtr returns correct value -- cached", actual)
}

func Test_Cov8_Result_SafeString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"val": r.SafeString()}
	expected := args.Map{"val": `"hi"`}
	expected.ShouldBeEqual(t, 0, "SafeString returns correct value -- with args", actual)
}

func Test_Cov8_Result_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{"empty": buf.Len() == 0, "noErr": err == nil}
	expected := args.Map{"empty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonBuffer returns empty -- empty", actual)
}

func Test_Cov8_Result_PrettyJsonBuffer_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{"hasData": buf.Len() > 0, "noErr": err == nil}
	expected := args.Map{"hasData": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonBuffer returns non-empty -- valid", actual)
}

func Test_Cov8_Result_PrettyJsonString_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"val": r.PrettyJsonString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns nil -- nil", actual)
}

func Test_Cov8_Result_PrettyJsonString_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	actual := args.Map{"notEmpty": r.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns non-empty -- valid", actual)
}

func Test_Cov8_Result_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"notEmpty": r.PrettyJsonStringOrErrString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString returns nil -- nil", actual)
}

func Test_Cov8_Result_PrettyJsonStringOrErrString_Error(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"notEmpty": r.PrettyJsonStringOrErrString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString returns error -- error", actual)
}

func Test_Cov8_Result_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	actual := args.Map{"notEmpty": r.PrettyJsonStringOrErrString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErrString returns error -- valid", actual)
}

func Test_Cov8_Result_Length_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"val": r.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Cov8_Result_Length_NilBytes(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil bytes", actual)
}

func Test_Cov8_Result_Length_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"val": r.Length()}
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Length returns non-empty -- valid", actual)
}

func Test_Cov8_Result_HasError(t *testing.T) {
	r1 := &corejson.Result{}
	r2 := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"noErr": r1.HasError(), "hasErr": r2.HasError()}
	expected := args.Map{"noErr": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "HasError returns error -- with args", actual)
}

func Test_Cov8_Result_ErrorString_Empty(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.ErrorString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ErrorString returns empty -- empty", actual)
}

func Test_Cov8_Result_ErrorString_WithErr(t *testing.T) {
	r := &corejson.Result{Error: errors.New("boom")}
	actual := args.Map{"val": r.ErrorString()}
	expected := args.Map{"val": "boom"}
	expected.ShouldBeEqual(t, 0, "ErrorString returns error -- with error", actual)
}

func Test_Cov8_Result_IsErrorEqual_BothNil(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.IsErrorEqual(nil)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns nil -- both nil", actual)
}

func Test_Cov8_Result_IsErrorEqual_OneNil(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"val": r.IsErrorEqual(nil)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns nil -- one nil", actual)
}

func Test_Cov8_Result_IsErrorEqual_Same(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"val": r.IsErrorEqual(errors.New("e"))}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns error -- same", actual)
}

func Test_Cov8_Result_IsErrorEqual_Different(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"val": r.IsErrorEqual(errors.New("f"))}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns error -- different", actual)
}

func Test_Cov8_Result_String_Valid(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- valid", actual)
}

func Test_Cov8_Result_String_WithError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), Error: errors.New("e"), TypeName: "T"}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns error -- with error", actual)
}

func Test_Cov8_Result_SafeNonIssueBytes_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"len": len(r.SafeNonIssueBytes())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes returns correct value -- issues", actual)
}

func Test_Cov8_Result_SafeNonIssueBytes_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"len": len(r.SafeNonIssueBytes())}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes returns non-empty -- valid", actual)
}

func Test_Cov8_Result_SafeBytes_AnyNull(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"len": len(r.SafeBytes())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeBytes returns correct value -- null", actual)
}

func Test_Cov8_Result_Values(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"len": len(r.Values())}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Values returns non-empty -- with args", actual)
}

func Test_Cov8_Result_SafeValues_Nil(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"len": len(r.SafeValues())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeValues returns nil -- nil", actual)
}

func Test_Cov8_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"len": len(r.SafeValuesPtr())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeValuesPtr returns non-empty -- issues", actual)
}

func Test_Cov8_Result_Raw_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Raw()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Raw returns nil -- nil", actual)
}

func Test_Cov8_Result_Raw_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	b, err := r.Raw()
	actual := args.Map{"len": len(b), "noErr": err == nil}
	expected := args.Map{"len": 4, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Raw returns non-empty -- valid", actual)
}

func Test_Cov8_Result_RawString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	s, err := r.RawString()
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": `"hi"`, "noErr": true}
	expected.ShouldBeEqual(t, 0, "RawString returns correct value -- with args", actual)
}

func Test_Cov8_Result_RawErrString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, errMsg := r.RawErrString()
	actual := args.Map{"len": len(b), "emptyErr": errMsg == ""}
	expected := args.Map{"len": 3, "emptyErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrString returns error -- with args", actual)
}

func Test_Cov8_Result_RawPrettyString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := r.RawPrettyString()
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "RawPrettyString returns correct value -- with args", actual)
}

func Test_Cov8_Result_MeaningfulError_Nil(t *testing.T) {
	var r *corejson.Result
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil", actual)
}

func Test_Cov8_Result_MeaningfulError_Ok(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	err := r.MeaningfulError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- ok", actual)
}

func Test_Cov8_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{TypeName: "T"}
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns empty -- empty bytes", actual)
}

func Test_Cov8_Result_MeaningfulError_WithError(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("fail"), TypeName: "T"}
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

func Test_Cov8_Result_MeaningfulErrorMessage_NoErr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"val": r.MeaningfulErrorMessage()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorMessage returns empty -- no err", actual)
}

func Test_Cov8_Result_IsEmptyError(t *testing.T) {
	r1 := &corejson.Result{}
	r2 := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"empty": r1.IsEmptyError(), "notEmpty": r2.IsEmptyError()}
	expected := args.Map{"empty": true, "notEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyError returns empty -- with args", actual)
}

func Test_Cov8_Result_HasSafeItems(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"hi"`)}
	r2 := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"safe": r1.HasSafeItems(), "notSafe": r2.HasSafeItems()}
	expected := args.Map{"safe": true, "notSafe": false}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_Cov8_Result_IsAnyNull(t *testing.T) {
	r1 := &corejson.Result{}
	r2 := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"null": r1.IsAnyNull(), "notNull": r2.IsAnyNull()}
	expected := args.Map{"null": true, "notNull": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNull returns correct value -- with args", actual)
}

func Test_Cov8_Result_HasIssuesOrEmpty(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"hi"`)}
	r2 := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"ok": r1.HasIssuesOrEmpty(), "err": r2.HasIssuesOrEmpty()}
	expected := args.Map{"ok": false, "err": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty returns empty -- with args", actual)
}

func Test_Cov8_Result_HasBytes(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"hi"`)}
	r2 := &corejson.Result{}
	actual := args.Map{"has": r1.HasBytes(), "notHas": r2.HasBytes()}
	expected := args.Map{"has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "HasBytes returns correct value -- with args", actual)
}

func Test_Cov8_Result_IsEmptyJsonBytes_EmptyJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("{}")}
	actual := args.Map{"empty": r.IsEmptyJsonBytes()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes returns empty -- {}", actual)
}

func Test_Cov8_Result_IsEmpty(t *testing.T) {
	r1 := &corejson.Result{}
	r2 := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"empty": r1.IsEmpty(), "notEmpty": r2.IsEmpty()}
	expected := args.Map{"empty": true, "notEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_Cov8_Result_HasAnyItem(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"has": r.HasAnyItem()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- with args", actual)
}

func Test_Cov8_Result_IsEmptyJson(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.IsEmptyJson()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJson returns empty -- with args", actual)
}

func Test_Cov8_Result_HasJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"val": r.HasJson()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasJson returns correct value -- with args", actual)
}

func Test_Cov8_Result_HasJsonBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"val": r.HasJsonBytes()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasJsonBytes returns correct value -- with args", actual)
}

func Test_Cov8_Result_Deserialize_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.Deserialize(&s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns non-empty -- valid", actual)
}

func Test_Cov8_Result_Unmarshal_Nil(t *testing.T) {
	var r *corejson.Result
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns nil -- nil", actual)
}

func Test_Cov8_Result_Unmarshal_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns error -- with error", actual)
}

func Test_Cov8_Result_Unmarshal_BadJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{bad`)}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns correct value -- bad json", actual)
}

func Test_Cov8_Result_SerializeSkipExistingIssues_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"nilBytes": b == nil, "noErr": err == nil}
	expected := args.Map{"nilBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues returns correct value -- issues", actual)
}

func Test_Cov8_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues returns non-empty -- valid", actual)
}

func Test_Cov8_Result_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	_, err := r.Serialize()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns nil -- nil", actual)
}

func Test_Cov8_Result_Serialize_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	_, err := r.Serialize()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns error -- with error", actual)
}

func Test_Cov8_Result_Serialize_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	b, err := r.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns non-empty -- valid", actual)
}

func Test_Cov8_Result_UnmarshalSkipExistingIssues_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues returns correct value -- issues", actual)
}

func Test_Cov8_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hi"}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues returns non-empty -- valid", actual)
}

func Test_Cov8_Result_UnmarshalSkipExistingIssues_BadJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{bad`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues returns correct value -- bad json", actual)
}

func Test_Cov8_Result_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	model := r.JsonModel()
	actual := args.Map{"hasErr": model.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "JsonModel returns nil -- nil", actual)
}

func Test_Cov8_Result_JsonModel_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	model := r.JsonModel()
	actual := args.Map{"len": model.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "JsonModel returns non-empty -- valid", actual)
}

func Test_Cov8_Result_JsonModelAny(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	actual := args.Map{"notNil": r.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_Cov8_Result_Ptr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`)}
	p := r.Ptr()
	actual := args.Map{"notNil": p != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Ptr returns correct value -- with args", actual)
}

func Test_Cov8_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	np := r.NonPtr()
	actual := args.Map{"hasErr": np.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns nil -- nil", actual)
}

func Test_Cov8_Result_NonPtr_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`)}
	np := r.NonPtr()
	actual := args.Map{"len": np.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "NonPtr returns non-empty -- valid", actual)
}

func Test_Cov8_Result_ToPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"notNil": r.ToPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToPtr returns correct value -- with args", actual)
}

func Test_Cov8_Result_ToNonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	np := r.ToNonPtr()
	actual := args.Map{"len": np.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ToNonPtr returns correct value -- with args", actual)
}

func Test_Cov8_Result_IsEqualPtr_BothNil(t *testing.T) {
	var r1 *corejson.Result
	actual := args.Map{"val": r1.IsEqualPtr(nil)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns nil -- both nil", actual)
}

func Test_Cov8_Result_IsEqualPtr_OneNil(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"val": r.IsEqualPtr(nil)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns nil -- one nil", actual)
}

func Test_Cov8_Result_IsEqualPtr_Same(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"val": r.IsEqualPtr(r)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- same", actual)
}

func Test_Cov8_Result_IsEqualPtr_DiffLength(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`)}
	r2 := &corejson.Result{Bytes: []byte(`"xy"`)}
	actual := args.Map{"val": r1.IsEqualPtr(r2)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- diff length", actual)
}

func Test_Cov8_Result_IsEqualPtr_DiffError(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("a")}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("b")}
	actual := args.Map{"val": r1.IsEqualPtr(r2)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns error -- diff error", actual)
}

func Test_Cov8_Result_IsEqualPtr_DiffTypeName(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "A"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "B"}
	actual := args.Map{"val": r1.IsEqualPtr(r2)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- diff type", actual)
}

func Test_Cov8_Result_IsEqualPtr_EqualBytes(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	actual := args.Map{"val": r1.IsEqualPtr(r2)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr returns correct value -- equal bytes", actual)
}

func Test_Cov8_Result_IsEqual(t *testing.T) {
	r1 := corejson.Result{Bytes: []byte(`"x"`)}
	r2 := corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"val": r1.IsEqual(r2)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- with args", actual)
}

func Test_Cov8_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"val": r.CombineErrorWithRefString("ref")}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString returns empty -- no error", actual)
}

func Test_Cov8_Result_CombineErrorWithRefString_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("boom")}
	result := r.CombineErrorWithRefString("ref1")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString returns error -- with error", actual)
}

func Test_Cov8_Result_CombineErrorWithRefError_NoError(t *testing.T) {
	r := &corejson.Result{}
	err := r.CombineErrorWithRefError("ref")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError returns empty -- no error", actual)
}

func Test_Cov8_Result_CombineErrorWithRefError_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("boom")}
	err := r.CombineErrorWithRefError("ref1")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError returns error -- with error", actual)
}

func Test_Cov8_Result_BytesError_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"nil": r.BytesError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "BytesError returns nil -- nil", actual)
}

func Test_Cov8_Result_BytesError_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	be := r.BytesError()
	actual := args.Map{"notNil": be != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesError returns error -- valid", actual)
}

func Test_Cov8_Result_Dispose(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	r.Dispose()
	actual := args.Map{"nilBytes": r.Bytes == nil, "nilErr": r.Error == nil, "emptyType": r.TypeName == ""}
	expected := args.Map{"nilBytes": true, "nilErr": true, "emptyType": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
}

func Test_Cov8_Result_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov8_Result_CloneIf_Clone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	cloned := r.CloneIf(true, false)
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "CloneIf returns correct value -- clone", actual)
}

func Test_Cov8_Result_CloneIf_NoClone(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	cloned := r.CloneIf(false, false)
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "CloneIf returns empty -- no clone", actual)
}

func Test_Cov8_Result_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"nil": r.ClonePtr(false) == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov8_Result_ClonePtr_DeepClone(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	cloned := r.ClonePtr(true)
	actual := args.Map{"notNil": cloned != nil, "len": cloned.Length()}
	expected := args.Map{"notNil": true, "len": 4}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- deep", actual)
}

func Test_Cov8_Result_Clone_EmptyLength(t *testing.T) {
	r := corejson.Result{TypeName: "T"}
	cloned := r.Clone(false)
	actual := args.Map{"len": cloned.Length(), "type": cloned.TypeName}
	expected := args.Map{"len": 0, "type": "T"}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- empty", actual)
}

func Test_Cov8_Result_Clone_ShallowCopy(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	cloned := r.Clone(false)
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_Cov8_Result_Clone_DeepCopy(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"hi"`), TypeName: "T"}
	cloned := r.Clone(true)
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_Cov8_Result_CloneError_WithErr(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	err := r.CloneError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError returns error -- with", actual)
}

func Test_Cov8_Result_CloneError_NoErr(t *testing.T) {
	r := &corejson.Result{}
	err := r.CloneError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError returns error -- none", actual)
}

func Test_Cov8_Result_AsJsonContractsBinder(t *testing.T) {
	r := corejson.Result{}
	actual := args.Map{"notNil": r.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov8_Result_AsJsoner(t *testing.T) {
	r := corejson.Result{}
	actual := args.Map{"notNil": r.AsJsoner() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsoner returns correct value -- with args", actual)
}

func Test_Cov8_Result_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.Result{}
	actual := args.Map{"notNil": r.AsJsonParseSelfInjector() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonParseSelfInjector returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Funcs — BytesCloneIf, BytesDeepClone, BytesToString, BytesToPrettyString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_BytesCloneIf_NoClone(t *testing.T) {
	result := corejson.BytesCloneIf(false, []byte("hi"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- no clone", actual)
}

func Test_Cov8_BytesCloneIf_Empty(t *testing.T) {
	result := corejson.BytesCloneIf(true, []byte{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- empty", actual)
}

func Test_Cov8_BytesDeepClone_Empty(t *testing.T) {
	result := corejson.BytesDeepClone(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns empty -- empty", actual)
}

func Test_Cov8_BytesDeepClone_Valid(t *testing.T) {
	result := corejson.BytesDeepClone([]byte("hi"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns non-empty -- valid", actual)
}

func Test_Cov8_BytesToString_Empty(t *testing.T) {
	actual := args.Map{"val": corejson.BytesToString(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "BytesToString returns empty -- empty", actual)
}

func Test_Cov8_BytesToString_Valid(t *testing.T) {
	actual := args.Map{"val": corejson.BytesToString([]byte("hi"))}
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesToString returns non-empty -- valid", actual)
}

func Test_Cov8_BytesToPrettyString_Empty(t *testing.T) {
	actual := args.Map{"val": corejson.BytesToPrettyString(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString returns empty -- empty", actual)
}

func Test_Cov8_BytesToPrettyString_Valid(t *testing.T) {
	result := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString returns non-empty -- valid", actual)
}

func Test_Cov8_JsonString_Func(t *testing.T) {
	s, err := corejson.JsonString("hello")
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- func", actual)
}

func Test_Cov8_JsonStringOrErrMsg_Valid(t *testing.T) {
	result := corejson.JsonStringOrErrMsg("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg returns error -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Serialize / Deserialize Logic
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_Serialize_Apply(t *testing.T) {
	r := corejson.Serialize.Apply("hello")
	actual := args.Map{"notNil": r != nil, "noErr": !r.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Apply returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize.StringsApply returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize.FromString returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize.FromInteger returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize.FromBool returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("hello")
	actual := args.Map{"notNil": r != nil, "noErr": !r.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.UsingAnyPtr returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.UsingAny returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToString returns correct value -- with args", actual)
}

func Test_Cov8_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToStringErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToBytesErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToSafeBytesSwallowErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToBytesSwallowErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToPrettyStringErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize.ToPrettyStringIncludingErr returns error -- with args", actual)
}

func Test_Cov8_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Pretty returns correct value -- with args", actual)
}

func Test_Cov8_Deserialize_UsingString(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingString(`"hello"`, &s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingString returns correct value -- with args", actual)
}

func Test_Cov8_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingStringPtr returns nil -- nil", actual)
}

func Test_Cov8_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	str := `"hello"`
	var s string
	err := corejson.Deserialize.UsingStringPtr(&str, &s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingStringPtr returns non-empty -- valid", actual)
}

func Test_Cov8_Deserialize_UsingError_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingError returns nil -- nil", actual)
}

func Test_Cov8_Deserialize_UsingBytes(t *testing.T) {
	var i int
	err := corejson.Deserialize.UsingBytes([]byte("42"), &i)
	actual := args.Map{"val": i, "noErr": err == nil}
	expected := args.Map{"val": 42, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns correct value -- with args", actual)
}

func Test_Cov8_Deserialize_UsingBytes_BadJson(t *testing.T) {
	var i int
	err := corejson.Deserialize.UsingBytes([]byte("{bad"), &i)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns correct value -- bad", actual)
}

func Test_Cov8_Deserialize_UsingBytesIf_Skip(t *testing.T) {
	var i int
	err := corejson.Deserialize.UsingBytesIf(false, []byte("42"), &i)
	actual := args.Map{"noErr": err == nil, "val": i}
	expected := args.Map{"noErr": true, "val": 0}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytesIf returns correct value -- skip", actual)
}

func Test_Cov8_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	var i int
	err := corejson.Deserialize.UsingBytesPointer(nil, &i)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytesPointer returns nil -- nil", actual)
}

func Test_Cov8_Deserialize_UsingBytesPointerIf_Skip(t *testing.T) {
	var i int
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte("42"), &i)
	actual := args.Map{"noErr": err == nil, "val": i}
	expected := args.Map{"noErr": true, "val": 0}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytesPointerIf returns correct value -- skip", actual)
}

func Test_Cov8_Deserialize_UsingStringOption_IgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingStringOption returns empty -- ignore empty", actual)
}

func Test_Cov8_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingStringIgnoreEmpty returns empty -- with args", actual)
}

func Test_Cov8_Deserialize_MapAnyToPointer_Empty(t *testing.T) {
	var s map[string]any
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.MapAnyToPointer returns empty -- empty skip", actual)
}

func Test_Cov8_Deserialize_AnyToFieldsMap(t *testing.T) {
	// AnyToFieldsMap should deserialize valid map input.
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	actual := args.Map{"isNil": m == nil, "hasErr": err != nil}
	expected := args.Map{"isNil": false, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "Deserialize.AnyToFieldsMap returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_Empty_Result(t *testing.T) {
	r := corejson.Empty.Result()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- with args", actual)
}

func Test_Cov8_Empty_ResultPtr(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr returns empty -- with args", actual)
}

func Test_Cov8_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))
	actual := args.Map{"hasErr": r.HasError(), "type": r.TypeName}
	expected := args.Map{"hasErr": true, "type": "T"}
	expected.ShouldBeEqual(t, 0, "Empty.ResultWithErr returns empty -- with args", actual)
}

func Test_Cov8_Empty_ResultPtrWithErr(t *testing.T) {
	r := corejson.Empty.ResultPtrWithErr("T", errors.New("e"))
	actual := args.Map{"notNil": r != nil, "hasErr": r.HasError()}
	expected := args.Map{"notNil": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtrWithErr returns empty -- with args", actual)
}

func Test_Cov8_Empty_BytesCollection(t *testing.T) {
	bc := corejson.Empty.BytesCollection()
	actual := args.Map{"empty": bc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollection returns empty -- with args", actual)
}

func Test_Cov8_Empty_BytesCollectionPtr(t *testing.T) {
	bc := corejson.Empty.BytesCollectionPtr()
	actual := args.Map{"notNil": bc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollectionPtr returns empty -- with args", actual)
}

func Test_Cov8_Empty_ResultsCollection(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsCollection returns empty -- with args", actual)
}

func Test_Cov8_Empty_ResultsPtrCollection(t *testing.T) {
	rc := corejson.Empty.ResultsPtrCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsPtrCollection returns empty -- with args", actual)
}

func Test_Cov8_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	actual := args.Map{"empty": mr.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.MapResults returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// New / NewPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_New_Valid(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"noErr": !r.HasError(), "notEmpty": !r.IsEmpty()}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "New returns non-empty -- valid", actual)
}

func Test_Cov8_NewPtr_Valid(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"notNil": r != nil, "noErr": !r.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPtr returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyTo
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonString returns correct value -- with args", actual)
}

func Test_Cov8_AnyTo_SafeJsonPrettyString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonPrettyString returns correct value -- with args", actual)
}

func Test_Cov8_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString("hi")
	actual := args.Map{"val": s}
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonPrettyString returns correct value -- string", actual)
}

func Test_Cov8_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonPrettyString returns correct value -- bytes", actual)
}

func Test_Cov8_AnyTo_JsonString(t *testing.T) {
	s := corejson.AnyTo.JsonString("hello")
	actual := args.Map{"val": s}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns correct value -- string", actual)
}

func Test_Cov8_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`"x"`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns correct value -- bytes", actual)
}

func Test_Cov8_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns correct value -- Result", actual)
}

func Test_Cov8_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns correct value -- *Result", actual)
}

func Test_Cov8_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonStringWithErr returns error -- string", actual)
}

func Test_Cov8_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonStringWithErr returns error -- bytes", actual)
}

func Test_Cov8_AnyTo_PrettyStringWithError_String(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringWithError returns error -- string", actual)
}

func Test_Cov8_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringWithError returns error -- bytes", actual)
}

func Test_Cov8_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns nil -- nil", actual)
}

func Test_Cov8_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult([]byte(`"x"`))
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- bytes", actual)
}

func Test_Cov8_AnyTo_SerializedJsonResult_String(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- string", actual)
}

func Test_Cov8_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedRaw returns correct value -- with args", actual)
}

func Test_Cov8_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedString returns correct value -- with args", actual)
}

func Test_Cov8_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedSafeString returns correct value -- with args", actual)
}

func Test_Cov8_AnyTo_SerializedFieldsMap(t *testing.T) {
	// SerializedFieldsMap should deserialize valid map input.
	m, err := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	actual := args.Map{"isNil": m == nil, "hasErr": err != nil}
	expected := args.Map{"isNil": false, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedFieldsMap returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NewResult Creator — Key Methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_NewResult_UsingBytes(t *testing.T) {
	r := corejson.NewResult.UsingBytes([]byte(`"hi"`))
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingBytes returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_UsingBytesType(t *testing.T) {
	r := corejson.NewResult.UsingBytesType([]byte(`"hi"`), "T")
	actual := args.Map{"type": r.TypeName}
	expected := args.Map{"type": "T"}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingBytesType returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_UsingBytesPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr(nil)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingBytesPtr returns nil -- nil", actual)
}

func Test_Cov8_NewResult_UsingString(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hi"`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingString returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_UsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingStringPtr returns nil -- nil", actual)
}

func Test_Cov8_NewResult_Error(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("e"))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.Error returns error -- with args", actual)
}

func Test_Cov8_NewResult_ErrorPtr(t *testing.T) {
	r := corejson.NewResult.ErrorPtr(errors.New("e"))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.ErrorPtr returns error -- with args", actual)
}

func Test_Cov8_NewResult_Empty(t *testing.T) {
	r := corejson.NewResult.Empty()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResult.Empty returns empty -- with args", actual)
}

func Test_Cov8_NewResult_EmptyPtr(t *testing.T) {
	r := corejson.NewResult.EmptyPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.EmptyPtr returns empty -- with args", actual)
}

func Test_Cov8_NewResult_Any(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.Any returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_AnyPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	actual := args.Map{"notNil": r != nil, "noErr": !r.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.AnyPtr returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b")
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.Many returns correct value -- with args", actual)
}

func Test_Cov8_NewResult_UsingSerializer_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializer(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingSerializer returns nil -- nil", actual)
}

func Test_Cov8_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingSerializerFunc returns nil -- nil", actual)
}

func Test_Cov8_NewResult_UsingJsoner_Nil(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingJsoner returns nil -- nil", actual)
}

func Test_Cov8_NewResult_PtrUsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.PtrUsingStringPtr(nil, "T")
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.PtrUsingStringPtr returns nil -- nil", actual)
}

func Test_Cov8_NewResult_PtrUsingStringPtr_Valid(t *testing.T) {
	s := `"hi"`
	r := corejson.NewResult.PtrUsingStringPtr(&s, "T")
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.PtrUsingStringPtr returns non-empty -- valid", actual)
}

func Test_Cov8_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingTypePlusStringPtr returns nil -- nil", actual)
}

func Test_Cov8_NewResult_UsingBytesPtrErrPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingBytesPtrErrPtr returns nil -- nil", actual)
}

func Test_Cov8_NewResult_UsingBytesErrPtr_Empty(t *testing.T) {
	r := corejson.NewResult.UsingBytesErrPtr([]byte{}, nil, "T")
	actual := args.Map{"len": len(r.Bytes)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingBytesErrPtr returns empty -- empty", actual)
}

func Test_Cov8_NewResult_PtrUsingBytesPtr_WithErr(t *testing.T) {
	r := corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T")
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.PtrUsingBytesPtr returns error -- with err", actual)
}

func Test_Cov8_NewResult_PtrUsingBytesPtr_NilBytes(t *testing.T) {
	r := corejson.NewResult.PtrUsingBytesPtr(nil, nil, "T")
	actual := args.Map{"notNil": r != nil, "len": len(r.Bytes)}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "NewResult.PtrUsingBytesPtr returns nil -- nil bytes", actual)
}

func Test_Cov8_NewResult_PtrUsingBytesPtr_Valid(t *testing.T) {
	r := corejson.NewResult.PtrUsingBytesPtr([]byte(`"hi"`), nil, "T")
	actual := args.Map{"noErr": !r.HasError(), "len": r.Length()}
	expected := args.Map{"noErr": true, "len": 4}
	expected.ShouldBeEqual(t, 0, "NewResult.PtrUsingBytesPtr returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsCollection — Basic
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov8_ResultsCollection_NilLength(t *testing.T) {
	var rc *corejson.ResultsCollection
	actual := args.Map{"len": rc.Length(), "empty": rc.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection returns nil -- nil", actual)
}

func Test_Cov8_ResultsCollection_Add(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	actual := args.Map{"len": rc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.Add returns correct value -- with args", actual)
}

func Test_Cov8_ResultsCollection_AddAny(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.AddAny("hello")
	actual := args.Map{"len": rc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.AddAny returns correct value -- with args", actual)
}

func Test_Cov8_ResultsCollection_FirstOrDefault_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"nil": rc.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.FirstOrDefault returns empty -- empty", actual)
}

func Test_Cov8_ResultsCollection_FirstOrDefault_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	actual := args.Map{"notNil": rc.FirstOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.FirstOrDefault returns non-empty -- valid", actual)
}

func Test_Cov8_ResultsCollection_LastOrDefault_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"nil": rc.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.LastOrDefault returns empty -- empty", actual)
}

func Test_Cov8_ResultsCollection_HasError_NoError(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	actual := args.Map{"hasErr": rc.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.HasError returns error -- no", actual)
}

func Test_Cov8_ResultsCollection_GetStrings(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	strs := rc.GetStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.GetStrings returns correct value -- with args", actual)
}

func Test_Cov8_ResultsCollection_GetPagesSize(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"val": rc.GetPagesSize(0)}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.GetPagesSize returns correct value -- zero", actual)
}

func Test_Cov8_ResultsCollection_Dispose(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.New("hello"))
	rc.Dispose()
	actual := args.Map{"nil": rc.Items == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.Dispose returns correct value -- with args", actual)
}

func Test_Cov8_ResultsCollection_Dispose_Nil(t *testing.T) {
	var rc *corejson.ResultsCollection
	rc.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection.Dispose returns nil -- nil", actual)
}
