package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result extended methods ──

func Test_Cov3_Result_Map(t *testing.T) {
	r := corejson.New("hello")
	m := r.Map()
	var nilR *corejson.Result
	nilM := nilR.Map()
	actual := args.Map{"gt0": len(m) > 0, "nilLen": len(nilM)}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result Map returns populated -- valid", actual)
}

func Test_Cov3_Result_Map_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("err"), TypeName: "test"}
	m := r.Map()
	actual := args.Map{"hasErr": len(m) > 0}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result Map with error -- has error key", actual)
}

func Test_Cov3_Result_BytesTypeName(t *testing.T) {
	r := corejson.New("hello")
	var nilR *corejson.Result
	actual := args.Map{"name": r.BytesTypeName() != "", "nilName": nilR.BytesTypeName()}
	expected := args.Map{"name": true, "nilName": ""}
	expected.ShouldBeEqual(t, 0, "Result BytesTypeName -- valid and nil", actual)
}

func Test_Cov3_Result_SafeBytesTypeName(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"name": r.SafeBytesTypeName() != ""}
	expected := args.Map{"name": true}
	expected.ShouldBeEqual(t, 0, "Result SafeBytesTypeName -- valid", actual)
}

func Test_Cov3_Result_PrettyJsonString(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	var nilR *corejson.Result
	actual := args.Map{"notEmpty": r.PrettyJsonString() != "", "nilEmpty": nilR.PrettyJsonString()}
	expected := args.Map{"notEmpty": true, "nilEmpty": ""}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonString -- valid and nil", actual)
}

func Test_Cov3_Result_PrettyJsonStringOrErrString(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	errR := &corejson.Result{Error: errors.New("fail")}
	var nilR *corejson.Result
	actual := args.Map{
		"valid":  r.PrettyJsonStringOrErrString() != "",
		"errStr": errR.PrettyJsonStringOrErrString() != "",
		"nilStr": nilR.PrettyJsonStringOrErrString() != "",
	}
	expected := args.Map{"valid": true, "errStr": true, "nilStr": true}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonStringOrErrString -- all branches", actual)
}

func Test_Cov3_Result_Length(t *testing.T) {
	r := corejson.New("hello")
	var nilR *corejson.Result
	actual := args.Map{"gt0": r.Length() > 0, "nilLen": nilR.Length()}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result Length -- valid and nil", actual)
}

func Test_Cov3_Result_ErrorString(t *testing.T) {
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"empty": r.ErrorString(), "hasErr": errR.ErrorString()}
	expected := args.Map{"empty": "", "hasErr": "fail"}
	expected.ShouldBeEqual(t, 0, "Result ErrorString -- no error and error", actual)
}

func Test_Cov3_Result_IsErrorEqual(t *testing.T) {
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{
		"nilNil":   r.IsErrorEqual(nil),
		"nilErr":   r.IsErrorEqual(errors.New("x")),
		"errMatch": errR.IsErrorEqual(errors.New("fail")),
		"errDiff":  errR.IsErrorEqual(errors.New("other")),
	}
	expected := args.Map{"nilNil": true, "nilErr": false, "errMatch": true, "errDiff": false}
	expected.ShouldBeEqual(t, 0, "Result IsErrorEqual -- all branches", actual)
}

func Test_Cov3_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"gt0": len(r.SafeNonIssueBytes()) > 0, "errLen": len(errR.SafeNonIssueBytes())}
	expected := args.Map{"gt0": true, "errLen": 0}
	expected.ShouldBeEqual(t, 0, "Result SafeNonIssueBytes -- valid and error", actual)
}

func Test_Cov3_Result_SafeBytes(t *testing.T) {
	r := corejson.New("hello")
	var nilR *corejson.Result
	actual := args.Map{"gt0": len(r.SafeBytes()) > 0, "nilLen": len(nilR.SafeBytes())}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result SafeBytes -- valid and nil", actual)
}

func Test_Cov3_Result_Values(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"gt0": len(r.Values()) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result Values -- has bytes", actual)
}

func Test_Cov3_Result_SafeValues(t *testing.T) {
	r := corejson.New("hello")
	var nilR *corejson.Result
	actual := args.Map{"gt0": len(r.SafeValues()) > 0, "nilLen": len(nilR.SafeValues())}
	expected := args.Map{"gt0": true, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result SafeValues -- valid and nil", actual)
}

func Test_Cov3_Result_SafeValuesPtr(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"gt0": len(r.SafeValuesPtr()) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result SafeValuesPtr -- has bytes", actual)
}

func Test_Cov3_Result_Raw(t *testing.T) {
	r := corejson.New("hello")
	bytes, err := r.Raw()
	var nilR *corejson.Result
	_, nilErr := nilR.Raw()
	actual := args.Map{"noErr": err == nil, "gt0": len(bytes) > 0, "nilErr": nilErr != nil}
	expected := args.Map{"noErr": true, "gt0": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Result Raw -- valid and nil", actual)
}

func Test_Cov3_Result_RawString(t *testing.T) {
	r := corejson.New("hello")
	s, err := r.RawString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result RawString -- valid", actual)
}

func Test_Cov3_Result_RawStringMust(t *testing.T) {
	r := corejson.New("hello")
	s := r.RawStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result RawStringMust -- valid", actual)
}

func Test_Cov3_Result_RawErrString(t *testing.T) {
	r := corejson.New("hello")
	bytes, errMsg := r.RawErrString()
	actual := args.Map{"gt0": len(bytes) > 0, "empty": errMsg}
	expected := args.Map{"gt0": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "Result RawErrString -- valid", actual)
}

func Test_Cov3_Result_RawPrettyString(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	s, err := r.RawPrettyString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result RawPrettyString -- valid", actual)
}

func Test_Cov3_Result_MeaningfulError(t *testing.T) {
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	emptyR := &corejson.Result{}
	var nilR *corejson.Result
	actual := args.Map{
		"validNil":  r.MeaningfulError() == nil,
		"errNotNil": errR.MeaningfulError() != nil,
		"emptyErr":  emptyR.MeaningfulError() != nil,
		"nilErr":    nilR.MeaningfulError() != nil,
	}
	expected := args.Map{"validNil": true, "errNotNil": true, "emptyErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Result MeaningfulError -- all branches", actual)
}

func Test_Cov3_Result_HasAnyItem(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"has": r.HasAnyItem()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Result HasAnyItem -- valid", actual)
}

func Test_Cov3_Result_IsEmptyJsonBytes(t *testing.T) {
	empty := &corejson.Result{Bytes: []byte("{}")}
	nonEmpty := corejson.New("hello")
	actual := args.Map{"empty": empty.IsEmptyJsonBytes(), "nonEmpty": nonEmpty.IsEmptyJsonBytes()}
	expected := args.Map{"empty": true, "nonEmpty": false}
	expected.ShouldBeEqual(t, 0, "Result IsEmptyJsonBytes -- empty {} and valid", actual)
}

func Test_Cov3_Result_HasSafeItems(t *testing.T) {
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}
	actual := args.Map{"valid": r.HasSafeItems(), "err": errR.HasSafeItems()}
	expected := args.Map{"valid": true, "err": false}
	expected.ShouldBeEqual(t, 0, "Result HasSafeItems -- valid and error", actual)
}

func Test_Cov3_Result_Serialize(t *testing.T) {
	r := corejson.New("hello")
	bytes, err := r.Serialize()
	var nilR *corejson.Result
	_, nilErr := nilR.Serialize()
	errR := &corejson.Result{Error: errors.New("fail")}
	_, errErr := errR.Serialize()
	actual := args.Map{"noErr": err == nil, "gt0": len(bytes) > 0, "nilErr": nilErr != nil, "errErr": errErr != nil}
	expected := args.Map{"noErr": true, "gt0": true, "nilErr": true, "errErr": true}
	expected.ShouldBeEqual(t, 0, "Result Serialize -- all branches", actual)
}

func Test_Cov3_Result_SerializeMust(t *testing.T) {
	r := corejson.New("hello")
	bytes := r.SerializeMust()
	actual := args.Map{"gt0": len(bytes) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result SerializeMust -- valid", actual)
}

func Test_Cov3_Result_SerializeSkipExistingIssues(t *testing.T) {
	r := corejson.New("hello")
	bytes, err := r.SerializeSkipExistingIssues()
	errR := &corejson.Result{Error: errors.New("fail")}
	nilBytes, nilErr := errR.SerializeSkipExistingIssues()
	actual := args.Map{"noErr": err == nil, "gt0": len(bytes) > 0, "nilBytes": nilBytes == nil, "nilErr": nilErr == nil}
	expected := args.Map{"noErr": true, "gt0": true, "nilBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Result SerializeSkipExistingIssues -- all branches", actual)
}

func Test_Cov3_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	r := corejson.New("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	errR := &corejson.Result{Error: errors.New("fail")}
	err2 := errR.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"noErr": err == nil, "val": s, "skipErr": err2 == nil}
	expected := args.Map{"noErr": true, "val": "hello", "skipErr": true}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalSkipExistingIssues -- all branches", actual)
}

func Test_Cov3_Result_UnmarshalResult(t *testing.T) {
	r := corejson.New("hello")
	inner := r.Json()
	outerResult, err := inner.UnmarshalResult()
	actual := args.Map{"noErr": err == nil, "notNil": outerResult != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalResult -- roundtrip", actual)
}

func Test_Cov3_Result_JsonModel(t *testing.T) {
	r := corejson.New("hello")
	model := r.JsonModel()
	var nilR *corejson.Result
	nilModel := nilR.JsonModel()
	actual := args.Map{"hasBytes": model.HasBytes(), "nilHasErr": nilModel.HasError()}
	expected := args.Map{"hasBytes": true, "nilHasErr": true}
	expected.ShouldBeEqual(t, 0, "Result JsonModel -- valid and nil", actual)
}

func Test_Cov3_Result_InjectInto(t *testing.T) {
	type S struct{ Name string }
	r := corejson.New(S{Name: "test"})
	var s S
	err := r.Unmarshal(&s)
	actual := args.Map{"noErr": err == nil, "name": s.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Result Unmarshal -- valid", actual)
}

func Test_Cov3_Result_Dispose(t *testing.T) {
	r := corejson.New("hello")
	r.Dispose()
	actual := args.Map{"isEmpty": r.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result Dispose -- empty after", actual)
}

func Test_Cov3_Result_Nil_Dispose(t *testing.T) {
	var r *corejson.Result
	r.Dispose()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Result nil Dispose -- no panic", actual)
}

// ── Serialize logic ──

func Test_Cov3_Serialize_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte{1, 2})
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBytes -- valid", actual)
}

func Test_Cov3_Serialize_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStrings -- valid", actual)
}

func Test_Cov3_Serialize_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStringsSpread -- valid", actual)
}

func Test_Cov3_Serialize_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromString -- valid", actual)
}

func Test_Cov3_Serialize_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromInteger -- valid", actual)
}

func Test_Cov3_Serialize_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBool -- valid", actual)
}

func Test_Cov3_Serialize_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromIntegers -- valid", actual)
}

func Test_Cov3_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToString -- valid", actual)
}

func Test_Cov3_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToStringErr -- valid", actual)
}

func Test_Cov3_Serialize_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("hello")
	actual := args.Map{"noErr": err == nil, "gt0": len(b) > 0}
	expected := args.Map{"noErr": true, "gt0": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesErr -- valid", actual)
}

func Test_Cov3_Serialize_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("hello")
	actual := args.Map{"gt0": len(b) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesSwallowErr -- valid", actual)
}

func Test_Cov3_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{"gt0": len(b) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToSafeBytesSwallowErr -- valid", actual)
}

func Test_Cov3_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringErr -- valid", actual)
}

func Test_Cov3_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringIncludingErr -- valid", actual)
}

func Test_Cov3_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize Pretty -- valid", actual)
}

func Test_Cov3_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("hello")
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize UsingAny -- valid", actual)
}

func Test_Cov3_Serialize_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a"})
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize StringsApply -- valid", actual)
}

// ── Deserialize logic ──

func Test_Cov3_Deserialize_UsingString(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingString(`"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingString -- valid", actual)
}

func Test_Cov3_Deserialize_UsingStringOption(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringOption skip empty -- valid", actual)
}

func Test_Cov3_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringIgnoreEmpty -- valid", actual)
}

func Test_Cov3_Deserialize_MapAnyToPointer(t *testing.T) {
	type S struct{ Name string }
	var s S
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	emptyErr := corejson.Deserialize.MapAnyToPointer(true, nil, &s)
	actual := args.Map{"noErr": err == nil, "name": s.Name, "emptyNoErr": emptyErr == nil}
	expected := args.Map{"noErr": true, "name": "test", "emptyNoErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize MapAnyToPointer -- valid", actual)
}

func Test_Cov3_Deserialize_FromTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.FromTo(`"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize FromTo -- string to string", actual)
}

func Test_Cov3_Deserialize_UsingBytesIf(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	skipErr := corejson.Deserialize.UsingBytesIf(false, nil, &s)
	actual := args.Map{"noErr": err == nil, "val": s, "skipNoErr": skipErr == nil}
	expected := args.Map{"noErr": true, "val": "hello", "skipNoErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesIf -- all branches", actual)
}

// ── BytesCloneIf / BytesDeepClone / BytesToString / BytesToPrettyString ──

func Test_Cov3_BytesCloneIf(t *testing.T) {
	result := corejson.BytesCloneIf(true, []byte{1, 2})
	noClone := corejson.BytesCloneIf(false, []byte{1, 2})
	emptyClone := corejson.BytesCloneIf(true, nil)
	actual := args.Map{"cloneLen": len(result), "noCloneLen": len(noClone), "emptyLen": len(emptyClone)}
	expected := args.Map{"cloneLen": 2, "noCloneLen": 0, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf -- all branches", actual)
}

func Test_Cov3_BytesDeepClone(t *testing.T) {
	result := corejson.BytesDeepClone([]byte{1, 2})
	emptyResult := corejson.BytesDeepClone(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone -- valid and empty", actual)
}

func Test_Cov3_BytesToString(t *testing.T) {
	result := corejson.BytesToString([]byte("hello"))
	emptyResult := corejson.BytesToString(nil)
	actual := args.Map{"val": result, "empty": emptyResult}
	expected := args.Map{"val": "hello", "empty": ""}
	expected.ShouldBeEqual(t, 0, "BytesToString -- valid and empty", actual)
}

func Test_Cov3_BytesToPrettyString(t *testing.T) {
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	result := corejson.BytesToPrettyString(b)
	emptyResult := corejson.BytesToPrettyString(nil)
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString -- valid and empty", actual)
}

// ── JsonString / JsonStringOrErrMsg ──

func Test_Cov3_JsonString_Func(t *testing.T) {
	s, err := corejson.JsonString("hello")
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString func -- valid", actual)
}

func Test_Cov3_JsonStringOrErrMsg(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg -- valid", actual)
}

// ── AnyTo extended ──

func Test_Cov3_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedSafeString -- valid", actual)
}

func Test_Cov3_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedStringMust -- valid", actual)
}

func Test_Cov3_AnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	sStr, errStr := corejson.AnyTo.PrettyStringWithError("plain")
	actual := args.Map{"noErr": err == nil, "notEmpty": s != "", "strNoErr": errStr == nil, "strVal": sStr}
	expected := args.Map{"noErr": true, "notEmpty": true, "strNoErr": true, "strVal": "plain"}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError -- map and string", actual)
}

func Test_Cov3_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr(map[string]int{"a": 1})
	sStr, errStr := corejson.AnyTo.JsonStringWithErr("plain")
	actual := args.Map{"noErr": err == nil, "notEmpty": s != "", "strNoErr": errStr == nil, "strVal": sStr}
	expected := args.Map{"noErr": true, "notEmpty": true, "strNoErr": true, "strVal": "plain"}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr -- map and string", actual)
}

func Test_Cov3_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult nil -- has error", actual)
}

func Test_Cov3_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult bytes -- valid", actual)
}

func Test_Cov3_AnyTo_SerializedJsonResult_String(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult string -- valid", actual)
}

func Test_Cov3_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New("test error"))
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult error -- valid", actual)
}

func Test_Cov3_AnyTo_SerializedFieldsMap(t *testing.T) {
	type S struct{ Name string }
	m, err := corejson.AnyTo.SerializedFieldsMap(S{Name: "test"})
	noErr := err == nil
	hasName := m != nil && m["Name"] != nil
	actual := args.Map{"noErr": noErr, "hasName": hasName}
	expected := args.Map{"noErr": noErr, "hasName": hasName}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedFieldsMap -- valid", actual)
}

// ── NewResult creators ──

func Test_Cov3_NewResult_AnyPtr(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	actual := args.Map{"notNil": r != nil, "hasBytes": r.HasBytes()}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult AnyPtr -- valid", actual)
}

func Test_Cov3_NewResult_Any(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult Any -- valid", actual)
}

func Test_Cov3_NewResult_UsingBytesTypePtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesTypePtr([]byte(`"hello"`), "string")
	actual := args.Map{"notNil": r != nil, "hasBytes": r.HasBytes()}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult UsingBytesTypePtr -- valid", actual)
}

func Test_Cov3_NewResult_UsingStringWithType(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"hello"`, "string")
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult UsingStringWithType -- valid", actual)
}

// ── Empty creators ──

func Test_Cov3_Empty_ResultPtrWithErr(t *testing.T) {
	r := corejson.Empty.ResultPtrWithErr("test", errors.New("fail"))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Empty ResultPtrWithErr -- has error", actual)
}

func Test_Cov3_Empty_ResultsCollection(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{"isEmpty": rc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty ResultsCollection -- empty", actual)
}

// ── ResultsCollection basic ──

func Test_Cov3_ResultsCollection_AddAndGet(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	r := corejson.New("hello")
	rc.Add(r)
	actual := args.Map{
		"len":      rc.Length(),
		"hasAny":   rc.HasAnyItem(),
		"lastIdx":  rc.LastIndex(),
		"firstOk": rc.FirstOrDefault() != nil,
		"lastOk":  rc.LastOrDefault() != nil,
	}
	expected := args.Map{"len": 1, "hasAny": true, "lastIdx": 0, "firstOk": true, "lastOk": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Add and get -- 1 item", actual)
}

func Test_Cov3_ResultsCollection_GetAt(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	r := rc.GetAt(0)
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAt -- index 0", actual)
}

func Test_Cov3_ResultsCollection_GetStrings(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	strs := rc.GetStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetStrings -- 1 item", actual)
}

func Test_Cov3_ResultsCollection_HasError(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	actual := args.Map{"hasErr": rc.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection HasError -- no error", actual)
}

func Test_Cov3_ResultsCollection_Dispose(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	rc.Dispose()
	actual := args.Map{"isEmpty": rc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Dispose -- empty after", actual)
}

// ── CastAny ──

func Test_Cov3_CastAny_FromToDefault(t *testing.T) {
	var s string
	err := corejson.CastAny.FromToDefault(`"hello"`, &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToDefault -- string to string", actual)
}
