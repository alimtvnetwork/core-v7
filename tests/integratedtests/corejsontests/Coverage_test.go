package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── New / NewPtr ──

func Test_New_Simple_Cov(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"hasError": r.Error != nil, "isEmpty": r.IsEmpty(), "hasBytes": len(r.Bytes) > 0, "typeName": r.TypeName != ""}
	expected := args.Map{"hasError": false, "isEmpty": false, "hasBytes": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "New_Simple returns correct value -- with args", actual)
}

func Test_NewPtr_Simple_Cov(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"isNil": r == nil, "hasError": r.Error != nil}
	expected := args.Map{"isNil": false, "hasError": false}
	expected.ShouldBeEqual(t, 0, "NewPtr_Simple returns correct value -- with args", actual)
}

func Test_New_Struct_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	actual := args.Map{"hasError": r.Error != nil, "hasBytes": len(r.Bytes) > 0, "stringNotEmpty": r.String() != "", "jsonStringNotNil": r.JsonString() != ""}
	expected := args.Map{"hasError": false, "hasBytes": true, "stringNotEmpty": true, "jsonStringNotNil": true}
	expected.ShouldBeEqual(t, 0, "New_Struct returns correct value -- with args", actual)
}

func Test_New_Nil_Cov(t *testing.T) {
	actual := args.Map{"hasError": corejson.New(nil).Error != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "New_Nil returns nil -- with args", actual)
}

// ── Result methods ──

func Test_Result_IsEmpty_Cov(t *testing.T) {
	empty := corejson.Result{}
	actual := args.Map{
		"isEmpty": empty.IsEmpty(), "isEmptyError": empty.IsEmptyError(),
		"hasError": empty.HasError(), "hasNoError": !empty.HasError(),
	}
	expected := args.Map{
		"isEmpty": true, "isEmptyError": true,
		"hasError": false, "hasNoError": true,
	}
	expected.ShouldBeEqual(t, 0, "Result_IsEmpty returns empty -- with args", actual)
}

func Test_Result_String_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": corejson.New(42).String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_String returns correct value -- with args", actual)
}

func Test_Result_JsonString_Cov(t *testing.T) {
	actual := args.Map{"result": corejson.NewPtr(42).JsonString()}
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Result_JsonString returns correct value -- with args", actual)
}

func Test_Result_SafeBytes_Cov(t *testing.T) {
	actual := args.Map{"hasBytes": len(corejson.NewPtr(42).SafeBytes()) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_SafeBytes returns correct value -- with args", actual)
}

func Test_Result_SafeBytes_Empty_Cov(t *testing.T) {
	r := corejson.NewPtr(nil)
	actual := args.Map{"notNil": r.SafeBytes() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Result_SafeBytes_Empty returns empty -- with args", actual)
}

func Test_Result_MustBytesJson_Cov(t *testing.T) {
	r := corejson.NewPtr(42)
	actual := args.Map{"hasBytes": len(r.SafeBytes()) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_MustBytes returns correct value -- with args", actual)
}

func Test_Result_PrettyJsonString_Cov(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": r.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_PrettyJsonString returns correct value -- with args", actual)
}

func Test_Result_PrettyJsonBytes_Cov(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": r.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_PrettyJsonBytes returns correct value -- with args", actual)
}

func Test_Result_Clone_Cov(t *testing.T) {
	c := corejson.NewPtr(42).Clone(false)
	actual := args.Map{"hasError": c.Error != nil, "hasBytes": len(c.Bytes) > 0}
	expected := args.Map{"hasError": false, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_Clone returns correct value -- with args", actual)
}

func Test_Result_ClonePtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewPtr(42).ClonePtr(false) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Result_ClonePtr returns correct value -- with args", actual)
}

func Test_Result_IsEqual_Cov(t *testing.T) {
	r1 := corejson.New(42)
	r2 := corejson.New(42)
	r3 := corejson.New(43)
	actual := args.Map{"same": r1.IsEqual(r2), "diff": r1.IsEqual(r3), "nil": r1.IsEqual(corejson.Result{})}
	expected := args.Map{"same": true, "diff": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "Result_IsEqual returns correct value -- with args", actual)
}

func Test_Result_Unmarshal_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := r.Unmarshal(&out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Result_Unmarshal returns correct value -- with args", actual)
}

// ── Empty creator ──

func Test_EmptyResult_Cov(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"isEmpty": r.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult returns empty -- with args", actual)
}

func Test_EmptyResultPtr_Cov(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": r != nil, "isEmpty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResultPtr returns empty -- with args", actual)
}

// ── Serialize / Deserialize ──

func Test_Serialize_Apply_Cov(t *testing.T) {
	actual := args.Map{"hasErr": corejson.Serialize.Apply(42).HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Serialize_Apply returns correct value -- with args", actual)
}

func Test_Serialize_UsingAnyPtr_Cov(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr(42)
	actual := args.Map{"notNil": r != nil, "hasErr": r.HasError()}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "Serialize_UsingAnyPtr returns correct value -- with args", actual)
}

func Test_Deserialize_UsingResult_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := corejson.Deserialize.UsingResult(&r, &out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingResult returns correct value -- with args", actual)
}

func Test_Deserialize_UsingBytes_Cov(t *testing.T) {
	type testS struct{ A int }
	var out testS
	err := corejson.Deserialize.UsingBytes([]byte(`{"A":42}`), &out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytes returns correct value -- with args", actual)
}

// ── CastAny / AnyTo / NewResult ──

func Test_CastAny_FromToDefault_Cov(t *testing.T) {
	type testS struct{ A int }
	from := testS{A: 42}
	var to testS
	err := corejson.CastAny.FromToDefault(from, &to)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToDefault returns correct value -- with args", actual)
}

func Test_AnyTo_SerializedJsonResult_Cov(t *testing.T) {
	actual := args.Map{"hasErr": corejson.AnyTo.SerializedJsonResult(42).HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult returns correct value -- with args", actual)
}

func Test_AnyTo_ResultPtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.AnyTo.SerializedJsonResult(42) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_ResultPtr returns correct value -- with args", actual)
}

func Test_NewResult_UsingBytes_Cov(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr([]byte(`"hello"`))
	actual := args.Map{"notEmpty": r != nil}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytes returns correct value -- with args", actual)
}

func Test_NewResult_UsingBytesPtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResult.UsingBytesPtr([]byte(`"hello"`)) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtr returns correct value -- with args", actual)
}

// ── Collections ──

func Test_NewBytesCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewBytesCollection.UsingCap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesCollection_UsingCap returns correct value -- with args", actual)
}

func Test_NewResultsCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResultsCollection.UsingCap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection_UsingCap returns correct value -- with args", actual)
}

func Test_NewResultsPtrCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResultsPtrCollection.UsingCap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsPtrCollection_UsingCap returns correct value -- with args", actual)
}

// ── Pretty ──

func Test_Pretty_Bytes_Cov(t *testing.T) {
	actual := args.Map{"hasBytes": len(corejson.Pretty.Bytes.SafeDefault([]byte(`{"a":1}`))) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Pretty_Bytes returns correct value -- with args", actual)
}

func Test_Pretty_String_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": corejson.Pretty.String.SafeDefault(`{"a":1}`) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty_String returns correct value -- with args", actual)
}

// ── NewMapResults ──

func Test_NewMapResults_Empty_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewMapResults.Empty() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMapResults_Empty returns empty -- with args", actual)
}
