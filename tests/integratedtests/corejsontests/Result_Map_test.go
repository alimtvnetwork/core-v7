package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result extended methods ──

func Test_Result_Map_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	m := r.Map()
	var nilR *corejson.Result
	nilM := nilR.Map()

	// Act
	actual := args.Map{
		"gt0": len(m) > 0,
		"nilLen": len(nilM),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Result Map returns populated -- valid", actual)
}

func Test_Result_Map_WithError_FromResultMap(t *testing.T) {
	// Arrange
	r := &corejson.Result{Error: errors.New("err"), TypeName: "test"}
	m := r.Map()

	// Act
	actual := args.Map{"hasErr": len(m) > 0}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result Map with error -- has error key", actual)
}

func Test_Result_BytesTypeName_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"name": r.BytesTypeName() != "",
		"nilName": nilR.BytesTypeName(),
	}

	// Assert
	expected := args.Map{
		"name": true,
		"nilName": "",
	}
	expected.ShouldBeEqual(t, 0, "Result BytesTypeName -- valid and nil", actual)
}

func Test_Result_SafeBytesTypeName_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{"name": r.SafeBytesTypeName() != ""}

	// Assert
	expected := args.Map{"name": true}
	expected.ShouldBeEqual(t, 0, "Result SafeBytesTypeName -- valid", actual)
}

func Test_Result_PrettyJsonString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"notEmpty": r.PrettyJsonString() != "",
		"nilEmpty": nilR.PrettyJsonString(),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nilEmpty": "",
	}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonString -- valid and nil", actual)
}

func Test_Result_PrettyJsonStringOrErrString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	errR := &corejson.Result{Error: errors.New("fail")}
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"valid":  r.PrettyJsonStringOrErrString() != "",
		"errStr": errR.PrettyJsonStringOrErrString() != "",
		"nilStr": nilR.PrettyJsonStringOrErrString() != "",
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"errStr": true,
		"nilStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonStringOrErrString -- all branches", actual)
}

func Test_Result_Length_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"gt0": r.Length() > 0,
		"nilLen": nilR.Length(),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Result Length -- valid and nil", actual)
}

func Test_Result_ErrorString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{
		"empty": r.ErrorString(),
		"hasErr": errR.ErrorString(),
	}

	// Assert
	expected := args.Map{
		"empty": "",
		"hasErr": "fail",
	}
	expected.ShouldBeEqual(t, 0, "Result ErrorString -- no error and error", actual)
}

func Test_Result_IsErrorEqual_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{
		"nilNil":   r.IsErrorEqual(nil),
		"nilErr":   r.IsErrorEqual(errors.New("x")),
		"errMatch": errR.IsErrorEqual(errors.New("fail")),
		"errDiff":  errR.IsErrorEqual(errors.New("other")),
	}

	// Assert
	expected := args.Map{
		"nilNil": true,
		"nilErr": false,
		"errMatch": true,
		"errDiff": false,
	}
	expected.ShouldBeEqual(t, 0, "Result IsErrorEqual -- all branches", actual)
}

func Test_Result_SafeNonIssueBytes_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{
		"gt0": len(r.SafeNonIssueBytes()) > 0,
		"errLen": len(errR.SafeNonIssueBytes()),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"errLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Result SafeNonIssueBytes -- valid and error", actual)
}

func Test_Result_SafeBytes_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"gt0": len(r.SafeBytes()) > 0,
		"nilLen": len(nilR.SafeBytes()),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Result SafeBytes -- valid and nil", actual)
}

func Test_Result_Values_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{"gt0": len(r.Values()) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result Values -- has bytes", actual)
}

func Test_Result_SafeValues_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"gt0": len(r.SafeValues()) > 0,
		"nilLen": len(nilR.SafeValues()),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Result SafeValues -- valid and nil", actual)
}

func Test_Result_SafeValuesPtr_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{"gt0": len(r.SafeValuesPtr()) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result SafeValuesPtr -- has bytes", actual)
}

func Test_Result_Raw_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	bytes, err := r.Raw()
	var nilR *corejson.Result
	_, nilErr := nilR.Raw()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"gt0": len(bytes) > 0,
		"nilErr": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"gt0": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Raw -- valid and nil", actual)
}

func Test_Result_RawString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s, err := r.RawString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Result RawString -- valid", actual)
}

func Test_Result_RawStringMust_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result RawStringMust -- valid", actual)
}

func Test_Result_RawErrString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	bytes, errMsg := r.RawErrString()

	// Act
	actual := args.Map{
		"gt0": len(bytes) > 0,
		"empty": errMsg,
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "Result RawErrString -- valid", actual)
}

func Test_Result_RawPrettyString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]int{"a": 1})
	s, err := r.RawPrettyString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Result RawPrettyString -- valid", actual)
}

func Test_Result_MeaningfulError_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail"), Bytes: []byte("x")}
	emptyR := &corejson.Result{}
	var nilR *corejson.Result

	// Act
	actual := args.Map{
		"validNil":  r.MeaningfulError() == nil,
		"errNotNil": errR.MeaningfulError() != nil,
		"emptyErr":  emptyR.MeaningfulError() != nil,
		"nilErr":    nilR.MeaningfulError() != nil,
	}

	// Assert
	expected := args.Map{
		"validNil": true,
		"errNotNil": true,
		"emptyErr": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result MeaningfulError -- all branches", actual)
}

func Test_Result_HasAnyItem_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{"has": r.HasAnyItem()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Result HasAnyItem -- valid", actual)
}

func Test_Result_IsEmptyJsonBytes(t *testing.T) {
	// Arrange
	empty := &corejson.Result{Bytes: []byte("{}")}
	nonEmpty := corejson.New("hello")

	// Act
	actual := args.Map{
		"empty": empty.IsEmptyJsonBytes(),
		"nonEmpty": nonEmpty.IsEmptyJsonBytes(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "Result IsEmptyJsonBytes -- empty {} and valid", actual)
}

func Test_Result_HasSafeItems_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	errR := &corejson.Result{Error: errors.New("fail")}

	// Act
	actual := args.Map{
		"valid": r.HasSafeItems(),
		"err": errR.HasSafeItems(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"err": false,
	}
	expected.ShouldBeEqual(t, 0, "Result HasSafeItems -- valid and error", actual)
}

func Test_Result_Serialize_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	bytes, err := r.Serialize()
	var nilR *corejson.Result
	_, nilErr := nilR.Serialize()
	errR := &corejson.Result{Error: errors.New("fail")}
	_, errErr := errR.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"gt0": len(bytes) > 0,
		"nilErr": nilErr != nil,
		"errErr": errErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"gt0": true,
		"nilErr": true,
		"errErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Serialize -- all branches", actual)
}

func Test_Result_SerializeMust_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	bytes := r.SerializeMust()

	// Act
	actual := args.Map{"gt0": len(bytes) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Result SerializeMust -- valid", actual)
}

func Test_Result_SerializeSkipExistingIssues_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	bytes, err := r.SerializeSkipExistingIssues()
	errR := &corejson.Result{Error: errors.New("fail")}
	nilBytes, nilErr := errR.SerializeSkipExistingIssues()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"gt0": len(bytes) > 0,
		"nilBytes": nilBytes == nil,
		"nilErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"gt0": true,
		"nilBytes": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result SerializeSkipExistingIssues -- all branches", actual)
}

func Test_Result_UnmarshalSkipExistingIssues_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	errR := &corejson.Result{Error: errors.New("fail")}
	err2 := errR.UnmarshalSkipExistingIssues(&s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
		"skipErr": err2 == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
		"skipErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalSkipExistingIssues -- all branches", actual)
}

func Test_Result_UnmarshalResult_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	inner := r.Json()
	outerResult, err := inner.UnmarshalResult()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": outerResult != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalResult -- roundtrip", actual)
}

func Test_Result_JsonModel_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	model := r.JsonModel()
	var nilR *corejson.Result
	nilModel := nilR.JsonModel()

	// Act
	actual := args.Map{
		"hasBytes": model.HasBytes(),
		"nilHasErr": nilModel.HasError(),
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"nilHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result JsonModel -- valid and nil", actual)
}

func Test_Result_InjectInto_FromResultMap(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	r := corejson.New(S{Name: "test"})
	var s S
	err := r.Unmarshal(&s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": s.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "Result Unmarshal -- valid", actual)
}

func Test_Result_Dispose_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	r.Dispose()

	// Act
	actual := args.Map{"isEmpty": r.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result Dispose -- empty after", actual)
}

func Test_Result_Nil_Dispose(t *testing.T) {
	// Arrange
	var r *corejson.Result
	r.Dispose()

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Result nil Dispose -- no panic", actual)
}

// ── Serialize logic ──

func Test_Serialize_FromBytes(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBytes([]byte{1, 2})

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBytes -- valid", actual)
}

func Test_Serialize_FromStrings(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStrings([]string{"a"})

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStrings -- valid", actual)
}

func Test_Serialize_FromStringsSpread(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromStringsSpread("a", "b")

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromStringsSpread -- valid", actual)
}

func Test_Serialize_FromString_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromString("hello")

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromString -- valid", actual)
}

func Test_Serialize_FromInteger_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromInteger(42)

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromInteger -- valid", actual)
}

func Test_Serialize_FromBool_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromBool(true)

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromBool -- valid", actual)
}

func Test_Serialize_FromIntegers(t *testing.T) {
	// Arrange
	r := corejson.Serialize.FromIntegers([]int{1, 2})

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize FromIntegers -- valid", actual)
}

func Test_Serialize_ToString_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToString("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToString -- valid", actual)
}

func Test_Serialize_ToStringErr_FromResultMap(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToStringErr("hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize ToStringErr -- valid", actual)
}

func Test_Serialize_ToBytesErr_FromResultMap(t *testing.T) {
	// Arrange
	b, err := corejson.Serialize.ToBytesErr("hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"gt0": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"gt0": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesErr -- valid", actual)
}

func Test_Serialize_ToBytesSwallowErr_FromResultMap(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToBytesSwallowErr("hello")

	// Act
	actual := args.Map{"gt0": len(b) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToBytesSwallowErr -- valid", actual)
}

func Test_Serialize_ToSafeBytesSwallowErr_FromResultMap(t *testing.T) {
	// Arrange
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")

	// Act
	actual := args.Map{"gt0": len(b) > 0}

	// Assert
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToSafeBytesSwallowErr -- valid", actual)
}

func Test_Serialize_ToPrettyStringErr_FromResultMap(t *testing.T) {
	// Arrange
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringErr -- valid", actual)
}

func Test_Serialize_ToPrettyStringIncludingErr_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize ToPrettyStringIncludingErr -- valid", actual)
}

func Test_Serialize_Pretty_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize Pretty -- valid", actual)
}

func Test_Serialize_UsingAny_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Serialize.UsingAny("hello")

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize UsingAny -- valid", actual)
}

func Test_Serialize_StringsApply_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Serialize.StringsApply([]string{"a"})

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize StringsApply -- valid", actual)
}

// ── Deserialize logic ──

func Test_Deserialize_UsingString_FromResultMap(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingString(`"hello"`, &s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingString -- valid", actual)
}

func Test_Deserialize_UsingStringOption(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringOption skip empty -- valid", actual)
}

func Test_Deserialize_UsingStringIgnoreEmpty_FromResultMap(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingStringIgnoreEmpty -- valid", actual)
}

func Test_Deserialize_MapAnyToPointer(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	var s S
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	emptyErr := corejson.Deserialize.MapAnyToPointer(true, nil, &s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": s.Name,
		"emptyNoErr": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
		"emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize MapAnyToPointer -- valid", actual)
}

func Test_Deserialize_FromTo_FromResultMap(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.FromTo(`"hello"`, &s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Deserialize FromTo -- string to string", actual)
}

func Test_Deserialize_UsingBytesIf(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	skipErr := corejson.Deserialize.UsingBytesIf(false, nil, &s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
		"skipNoErr": skipErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
		"skipNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize UsingBytesIf -- all branches", actual)
}

// ── BytesCloneIf / BytesDeepClone / BytesToString / BytesToPrettyString ──

func Test_BytesCloneIf(t *testing.T) {
	// Arrange
	result := corejson.BytesCloneIf(true, []byte{1, 2})
	noClone := corejson.BytesCloneIf(false, []byte{1, 2})
	emptyClone := corejson.BytesCloneIf(true, nil)

	// Act
	actual := args.Map{
		"cloneLen": len(result),
		"noCloneLen": len(noClone),
		"emptyLen": len(emptyClone),
	}

	// Assert
	expected := args.Map{
		"cloneLen": 2,
		"noCloneLen": 0,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf -- all branches", actual)
}

func Test_BytesDeepClone_FromResultMap(t *testing.T) {
	// Arrange
	result := corejson.BytesDeepClone([]byte{1, 2})
	emptyResult := corejson.BytesDeepClone(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone -- valid and empty", actual)
}

func Test_BytesToString_FromResultMap(t *testing.T) {
	// Arrange
	result := corejson.BytesToString([]byte("hello"))
	emptyResult := corejson.BytesToString(nil)

	// Act
	actual := args.Map{
		"val": result,
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesToString -- valid and empty", actual)
}

func Test_BytesToPrettyString_FromResultMap(t *testing.T) {
	// Arrange
	b, _ := corejson.Serialize.Raw(map[string]int{"a": 1})
	result := corejson.BytesToPrettyString(b)
	emptyResult := corejson.BytesToPrettyString(nil)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString -- valid and empty", actual)
}

// ── JsonString / JsonStringOrErrMsg ──

func Test_JsonString_Func_FromResultMap(t *testing.T) {
	// Arrange
	s, err := corejson.JsonString("hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString func -- valid", actual)
}

func Test_JsonStringOrErrMsg_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg("hello")

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg -- valid", actual)
}

// ── AnyTo extended ──

func Test_AnyTo_SerializedSafeString_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedSafeString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedSafeString -- valid", actual)
}

func Test_AnyTo_SerializedStringMust_FromResultMap(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SerializedStringMust(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedStringMust -- valid", actual)
}

func Test_AnyTo_PrettyStringWithError_FromResultMap(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.PrettyStringWithError(map[string]int{"a": 1})
	sStr, errStr := corejson.AnyTo.PrettyStringWithError("plain")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
		"strNoErr": errStr == nil,
		"strVal": sStr,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
		"strNoErr": true,
		"strVal": "plain",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError -- map and string", actual)
}

func Test_AnyTo_JsonStringWithErr_FromResultMap(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.JsonStringWithErr(map[string]int{"a": 1})
	sStr, errStr := corejson.AnyTo.JsonStringWithErr("plain")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
		"strNoErr": errStr == nil,
		"strVal": sStr,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
		"strNoErr": true,
		"strVal": "plain",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr -- map and string", actual)
}

func Test_AnyTo_SerializedJsonResult_Nil_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(nil)

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult nil -- has error", actual)
}

func Test_AnyTo_SerializedJsonResult_Bytes_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult bytes -- valid", actual)
}

func Test_AnyTo_SerializedJsonResult_String_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(`"hello"`)

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult string -- valid", actual)
}

func Test_AnyTo_SerializedJsonResult_Error_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.SerializedJsonResult(errors.New("test error"))

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult error -- valid", actual)
}

func Test_AnyTo_SerializedFieldsMap_FromResultMap(t *testing.T) {
	// Arrange
	type S struct{ Name string }
	m, err := corejson.AnyTo.SerializedFieldsMap(S{Name: "test"})
	noErr := err == nil
	hasName := m != nil && m["Name"] != nil

	// Act
	actual := args.Map{
		"noErr": noErr,
		"hasName": hasName,
	}

	// Assert
	expected := args.Map{
		"noErr": noErr,
		"hasName": hasName,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedFieldsMap -- valid", actual)
}

// ── NewResult creators ──

func Test_NewResult_AnyPtr_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasBytes": r.HasBytes(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult AnyPtr -- valid", actual)
}

func Test_NewResult_Any_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Any("hello")

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult Any -- valid", actual)
}

func Test_NewResult_UsingBytesTypePtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingBytesTypePtr([]byte(`"hello"`), "string")

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"hasBytes": r.HasBytes(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult UsingBytesTypePtr -- valid", actual)
}

func Test_NewResult_UsingStringWithType(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingStringWithType(`"hello"`, "string")

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult UsingStringWithType -- valid", actual)
}

// ── Empty creators ──

func Test_Empty_ResultPtrWithErr_FromResultMap(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtrWithErr("test", errors.New("fail"))

	// Act
	actual := args.Map{"hasErr": r.HasError()}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Empty ResultPtrWithErr -- has error", actual)
}

func Test_Empty_ResultsCollection_FromResultMap(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()

	// Act
	actual := args.Map{"isEmpty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty ResultsCollection -- empty", actual)
}

// ── ResultsCollection basic ──

func Test_ResultsCollection_AddAndGet(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()
	r := corejson.New("hello")
	rc.Add(r)

	// Act
	actual := args.Map{
		"len":      rc.Length(),
		"hasAny":   rc.HasAnyItem(),
		"lastIdx":  rc.LastIndex(),
		"firstOk": rc.FirstOrDefault() != nil,
		"lastOk":  rc.LastOrDefault() != nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasAny": true,
		"lastIdx": 0,
		"firstOk": true,
		"lastOk": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Add and get -- 1 item", actual)
}

func Test_ResultsCollection_GetAt(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	r := rc.GetAt(0)

	// Act
	actual := args.Map{"hasBytes": r.HasBytes()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAt -- index 0", actual)
}

func Test_ResultsCollection_GetStrings_FromResultMap(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	strs := rc.GetStrings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetStrings -- 1 item", actual)
}

func Test_ResultsCollection_HasError(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))

	// Act
	actual := args.Map{"hasErr": rc.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection HasError -- no error", actual)
}

func Test_ResultsCollection_Dispose_FromResultMap(t *testing.T) {
	// Arrange
	rc := corejson.Empty.ResultsCollection()
	rc.Add(corejson.New("hello"))
	rc.Dispose()

	// Act
	actual := args.Map{"isEmpty": rc.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Dispose -- empty after", actual)
}

// ── CastAny ──

func Test_CastAny_FromToDefault_FromResultMap(t *testing.T) {
	// Arrange
	var s string
	err := corejson.CastAny.FromToDefault(`"hello"`, &s)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "CastAny FromToDefault -- string to string", actual)
}
