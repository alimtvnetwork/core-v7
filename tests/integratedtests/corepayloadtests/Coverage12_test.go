package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// PayloadWrapper — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_PW_HasSafeItems(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"data"`))
	actual := args.Map{"val": pw.HasSafeItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_Cov12_PW_HasSafeItems_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"val": pw.HasSafeItems()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns empty -- empty", actual)
}

func Test_Cov12_PW_DynamicPayloads_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"len": len(pw.DynamicPayloads())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns nil -- nil", actual)
}

func Test_Cov12_PW_SetDynamicPayloads_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	err := pw.SetDynamicPayloads([]byte("data"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SetDynamicPayloads returns nil -- nil", actual)
}

func Test_Cov12_PW_SetDynamicPayloads_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte("data"))
	actual := args.Map{"noErr": err == nil, "len": len(pw.Payloads)}
	expected := args.Map{"noErr": true, "len": 4}
	expected.ShouldBeEqual(t, 0, "SetDynamicPayloads returns non-empty -- valid", actual)
}

func Test_Cov12_PW_AttrAsBinder(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"notNil": pw.AttrAsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AttrAsBinder returns correct value -- with args", actual)
}

func Test_Cov12_PW_InitializeAttributesOnNull(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	binder := pw.InitializeAttributesOnNull()
	actual := args.Map{"notNil": binder != nil, "hasAttr": pw.Attributes != nil}
	expected := args.Map{"notNil": true, "hasAttr": true}
	expected.ShouldBeEqual(t, 0, "InitializeAttributesOnNull returns correct value -- with args", actual)
}

func Test_Cov12_PW_BasicError_NoError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"nil": pw.BasicError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "BasicError returns empty -- no error", actual)
}

func Test_Cov12_PW_All(t *testing.T) {
	pw := corepayload.PayloadWrapper{Identifier: "id", Name: "n", EntityType: "e", CategoryName: "c", Payloads: []byte("p")}
	id, name, entity, cat, dyn := pw.All()
	actual := args.Map{"id": id, "name": name, "entity": entity, "cat": cat, "dynLen": len(dyn)}
	expected := args.Map{"id": "id", "name": "n", "entity": "e", "cat": "c", "dynLen": 1}
	expected.ShouldBeEqual(t, 0, "All returns correct value -- with args", actual)
}

func Test_Cov12_PW_AllSafe_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	id, name, entity, cat, dyn := pw.AllSafe()
	actual := args.Map{"id": id, "name": name, "entity": entity, "cat": cat, "dynLen": len(dyn)}
	expected := args.Map{"id": "", "name": "", "entity": "", "cat": "", "dynLen": 0}
	expected.ShouldBeEqual(t, 0, "AllSafe returns nil -- nil", actual)
}

func Test_Cov12_PW_PayloadName(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"val": pw.PayloadName()}
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "PayloadName returns correct value -- with args", actual)
}

func Test_Cov12_PW_PayloadCategory(t *testing.T) {
	pw := corepayload.PayloadWrapper{CategoryName: "cat"}
	actual := args.Map{"val": pw.PayloadCategory()}
	expected := args.Map{"val": "cat"}
	expected.ShouldBeEqual(t, 0, "PayloadCategory returns correct value -- with args", actual)
}

func Test_Cov12_PW_PayloadTaskType(t *testing.T) {
	pw := corepayload.PayloadWrapper{TaskTypeName: "task"}
	actual := args.Map{"val": pw.PayloadTaskType()}
	expected := args.Map{"val": "task"}
	expected.ShouldBeEqual(t, 0, "PayloadTaskType returns correct value -- with args", actual)
}

func Test_Cov12_PW_PayloadEntityType(t *testing.T) {
	pw := corepayload.PayloadWrapper{EntityType: "entity"}
	actual := args.Map{"val": pw.PayloadEntityType()}
	expected := args.Map{"val": "entity"}
	expected.ShouldBeEqual(t, 0, "PayloadEntityType returns correct value -- with args", actual)
}

func Test_Cov12_PW_PayloadDynamic(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"len": len(pw.PayloadDynamic())}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "PayloadDynamic returns correct value -- with args", actual)
}

func Test_Cov12_PW_Value(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"notNil": pw.Value() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Value returns correct value -- with args", actual)
}

func Test_Cov12_PW_Error_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"noErr": pw.Error() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns empty -- empty", actual)
}

func Test_Cov12_PW_HasError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"no": pw.HasError(), "nil": pw2.HasError()}
	expected := args.Map{"no": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasError returns error -- with args", actual)
}

func Test_Cov12_PW_IsEmptyError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"empty": pw.IsEmptyError(), "nil": pw2.IsEmptyError()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyError returns empty -- with args", actual)
}

func Test_Cov12_PW_HasAttributes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"has": pw.HasAttributes(), "nil": pw2.HasAttributes()}
	expected := args.Map{"has": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasAttributes returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsEmptyAttributes(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"empty": pw.IsEmptyAttributes(), "nil": pw2.IsEmptyAttributes()}
	expected := args.Map{"empty": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyAttributes returns empty -- with args", actual)
}

func Test_Cov12_PW_HasSingleRecord(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"single": pw.HasSingleRecord(), "nil": pw2.HasSingleRecord()}
	expected := args.Map{"single": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasSingleRecord returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsNull(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	pw2 := &corepayload.PayloadWrapper{}
	actual := args.Map{"nil": pw.IsNull(), "notNil": pw2.IsNull()}
	expected := args.Map{"nil": true, "notNil": false}
	expected.ShouldBeEqual(t, 0, "IsNull returns correct value -- with args", actual)
}

func Test_Cov12_PW_HasAnyNil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"val": pw.HasAnyNil()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyNil returns nil -- with args", actual)
}

func Test_Cov12_PW_Count(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"val": pw.Count()}
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Count returns correct value -- with args", actual)
}

func Test_Cov12_PW_Length_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	actual := args.Map{"val": pw.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Cov12_PW_IsEmpty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"val": pw.IsEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_Cov12_PW_HasItems(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("d")}
	actual := args.Map{"val": pw.HasItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasItems returns correct value -- with args", actual)
}

func Test_Cov12_PW_HasAnyItem(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("d")}
	actual := args.Map{"val": pw.HasAnyItem()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- with args", actual)
}

func Test_Cov12_PW_HasIssuesOrEmpty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"val": pw.HasIssuesOrEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty returns empty -- with args", actual)
}

func Test_Cov12_PW_IdentifierInteger_Empty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"val": pw.IdentifierInteger()}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns empty -- empty", actual)
}

func Test_Cov12_PW_IdentifierInteger_Valid(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}
	actual := args.Map{"val": pw.IdentifierInteger()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns non-empty -- valid", actual)
}

func Test_Cov12_PW_IdentifierUnsignedInteger_Negative(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "-1"}
	actual := args.Map{"val": pw.IdentifierUnsignedInteger()}
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "IdentifierUnsignedInteger returns correct value -- negative", actual)
}

func Test_Cov12_PW_IdString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "id1"}
	actual := args.Map{"val": pw.IdString()}
	expected := args.Map{"val": "id1"}
	expected.ShouldBeEqual(t, 0, "IdString returns correct value -- with args", actual)
}

func Test_Cov12_PW_IdInteger(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "42"}
	actual := args.Map{"val": pw.IdInteger()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IdInteger returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsPayloadsEqual(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"match": pw.IsPayloadsEqual([]byte("data")), "no": pw.IsPayloadsEqual([]byte("other"))}
	expected := args.Map{"match": true, "no": false}
	expected.ShouldBeEqual(t, 0, "IsPayloadsEqual returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsName(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	var pw2 *corepayload.PayloadWrapper
	actual := args.Map{"match": pw.IsName("test"), "no": pw.IsName("x"), "nil": pw2.IsName("test")}
	expected := args.Map{"match": true, "no": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsName returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsIdentifier(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Identifier: "id1"}
	actual := args.Map{"match": pw.IsIdentifier("id1"), "no": pw.IsIdentifier("x")}
	expected := args.Map{"match": true, "no": false}
	expected.ShouldBeEqual(t, 0, "IsIdentifier returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsTaskTypeName(t *testing.T) {
	pw := &corepayload.PayloadWrapper{TaskTypeName: "task"}
	actual := args.Map{"match": pw.IsTaskTypeName("task")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsTaskTypeName returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsEntityType(t *testing.T) {
	pw := &corepayload.PayloadWrapper{EntityType: "entity"}
	actual := args.Map{"match": pw.IsEntityType("entity")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsEntityType returns correct value -- with args", actual)
}

func Test_Cov12_PW_IsCategory(t *testing.T) {
	pw := &corepayload.PayloadWrapper{CategoryName: "cat"}
	actual := args.Map{"match": pw.IsCategory("cat")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsCategory returns correct value -- with args", actual)
}

func Test_Cov12_PW_Username_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"val": pw.Username()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Username returns empty -- empty", actual)
}

func Test_Cov12_PW_Username_NoUser(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Attributes.AuthInfo = &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{}}
	actual := args.Map{"val": pw.Username()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Username returns empty -- no user", actual)
}

func Test_Cov12_PW_Username_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.Attributes.AuthInfo = &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "alice"}}}
	actual := args.Map{"val": pw.Username()}
	expected := args.Map{"val": "alice"}
	expected.ShouldBeEqual(t, 0, "Username returns non-empty -- valid", actual)
}

func Test_Cov12_PW_String(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notEmpty": pw.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_Cov12_PW_PrettyJsonString(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notEmpty": pw.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonString(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notEmpty": pw.JsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonStringMust(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notEmpty": pw.JsonStringMust() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- with args", actual)
}

func Test_Cov12_PW_Serialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	b, err := pw.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_Cov12_PW_SerializeMust(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	b := pw.SerializeMust()
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust returns correct value -- with args", actual)
}

func Test_Cov12_PW_Deserialize_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"hello"`))
	var s string
	err := pw.Deserialize(&s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns non-empty -- valid", actual)
}

func Test_Cov12_PW_PayloadDeserialize(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	var s string
	err := pw.PayloadDeserialize(&s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadDeserialize returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonModel(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	model := pw.JsonModel()
	actual := args.Map{"name": model.Name}
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonModelAny(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny returns correct value -- with args", actual)
}

func Test_Cov12_PW_Json(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	r := pw.Json()
	actual := args.Map{"notEmpty": !r.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonPtr(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.JsonPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- with args", actual)
}

func Test_Cov12_PW_ParseInjectUsingJson(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	jr := pw.JsonPtr()
	pw2 := &corepayload.PayloadWrapper{}
	result, err := pw2.ParseInjectUsingJson(jr)
	actual := args.Map{"notNil": result != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- with args", actual)
}

func Test_Cov12_PW_JsonParseSelfInject(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	jr := pw.JsonPtr()
	pw2 := &corepayload.PayloadWrapper{}
	err := pw2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns correct value -- with args", actual)
}

func Test_Cov12_PW_Clear_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	pw.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Cov12_PW_Clear_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	pw.Clear()
	actual := args.Map{"empty": pw.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Clear returns non-empty -- valid", actual)
}

func Test_Cov12_PW_Dispose_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	pw.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov12_PW_Dispose_Valid(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	pw.Dispose()
	actual := args.Map{"nilAttr": pw.Attributes == nil}
	expected := args.Map{"nilAttr": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns non-empty -- valid", actual)
}

func Test_Cov12_PW_Clone_Shallow(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	c, err := pw.Clone(false)
	actual := args.Map{"name": c.Name, "noErr": err == nil}
	expected := args.Map{"name": "n", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- shallow", actual)
}

func Test_Cov12_PW_Clone_Deep(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	c, err := pw.Clone(true)
	actual := args.Map{"name": c.Name, "noErr": err == nil}
	expected := args.Map{"name": "n", "noErr": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep", actual)
}

func Test_Cov12_PW_ClonePtr_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	c, err := pw.ClonePtr(false)
	actual := args.Map{"nil": c == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov12_PW_NonPtr_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	np := pw.NonPtr()
	actual := args.Map{"empty": np.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NonPtr returns nil -- nil", actual)
}

func Test_Cov12_PW_ToPtr(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.ToPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToPtr returns correct value -- with args", actual)
}

func Test_Cov12_PW_AsJsonContractsBinder(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"notNil": pw.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns correct value -- with args", actual)
}

func Test_Cov12_PW_AsPayloadsBinder(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.AsPayloadsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsPayloadsBinder returns correct value -- with args", actual)
}

func Test_Cov12_PW_AsJsonMarshaller(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.AsJsonMarshaller() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsJsonMarshaller returns correct value -- with args", actual)
}

func Test_Cov12_PW_AsStandardTaskEntityDefinerContractsBinder(t *testing.T) {
	pw := corepayload.PayloadWrapper{Name: "test"}
	actual := args.Map{"notNil": pw.AsStandardTaskEntityDefinerContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsStandardTaskEntityDefinerContractsBinder returns correct value -- with args", actual)
}

func Test_Cov12_PW_HandleError_NoError(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	pw.HandleError() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError returns empty -- no error", actual)
}

func Test_Cov12_PW_PayloadsString_Empty(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	actual := args.Map{"val": pw.PayloadsString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns empty -- empty", actual)
}

func Test_Cov12_PW_PayloadsString_Valid(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("hello")}
	actual := args.Map{"val": pw.PayloadsString()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns non-empty -- valid", actual)
}

func Test_Cov12_PW_PayloadsJsonResult_Empty(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	r := pw.PayloadsJsonResult()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns empty -- empty", actual)
}

func Test_Cov12_PW_PayloadsJsonResult_Valid(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	r := pw.PayloadsJsonResult()
	actual := args.Map{"notEmpty": !r.IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsJsonResult returns non-empty -- valid", actual)
}

func Test_Cov12_PW_PayloadsPrettyString_Empty(t *testing.T) {
	pw := corepayload.PayloadWrapper{}
	actual := args.Map{"val": pw.PayloadsPrettyString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsPrettyString returns empty -- empty", actual)
}

func Test_Cov12_PW_PayloadsPrettyString_Valid(t *testing.T) {
	pw := corepayload.PayloadWrapper{Payloads: []byte(`{"a":1}`)}
	actual := args.Map{"notEmpty": pw.PayloadsPrettyString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PayloadsPrettyString returns non-empty -- valid", actual)
}

func Test_Cov12_PW_BytesConverter(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte("data")}
	actual := args.Map{"notNil": pw.BytesConverter() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- with args", actual)
}

func Test_Cov12_PW_PayloadProperties(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	props := pw.PayloadProperties()
	actual := args.Map{"notNil": props != nil, "name": props.Name()}
	expected := args.Map{"notNil": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "PayloadProperties returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PayloadsCollection — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_PC_Length_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"val": pc.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- Length nil", actual)
}

func Test_Cov12_PC_Count(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"val": pc.Count()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- Count", actual)
}

func Test_Cov12_PC_IsEmpty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"val": pc.IsEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "PC returns empty -- IsEmpty", actual)
}

func Test_Cov12_PC_HasAnyItem(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.PayloadWrapper{Name: "test", Payloads: []byte(`"d"`)}
	pc.Add(pw)
	actual := args.Map{"val": pc.HasAnyItem()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- HasAnyItem", actual)
}

func Test_Cov12_PC_FirstOrDefault_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"nil": pc.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PC returns empty -- FirstOrDefault empty", actual)
}

func Test_Cov12_PC_LastOrDefault_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"nil": pc.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PC returns empty -- LastOrDefault empty", actual)
}

func Test_Cov12_PC_Add_And_Access(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.PayloadWrapper{Name: "first", Payloads: []byte(`"d"`)}
	pc.Add(pw)
	pw2 := corepayload.PayloadWrapper{Name: "second", Payloads: []byte(`"e"`)}
	pc.Add(pw2)
	actual := args.Map{
		"len":   pc.Length(),
		"first": pc.First().Name,
		"last":  pc.Last().Name,
	}
	expected := args.Map{
		"len":   2,
		"first": "first",
		"last":  "second",
	}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- Add and Access", actual)
}

func Test_Cov12_PC_Clone(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.PayloadWrapper{Name: "test", Payloads: []byte(`"d"`)}
	pc.Add(pw)
	c := pc.Clone()
	actual := args.Map{"len": c.Length(), "name": c.First().Name}
	expected := args.Map{"len": 1, "name": "test"}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- Clone", actual)
}

func Test_Cov12_PC_ClonePtr_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"nil": pc.ClonePtr() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- ClonePtr nil", actual)
}

func Test_Cov12_PC_Clear_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	result := pc.Clear()
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- Clear nil", actual)
}

func Test_Cov12_PC_Dispose_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	pc.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- Dispose nil", actual)
}

func Test_Cov12_PC_GetPagesSize(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		pc.Add(corepayload.PayloadWrapper{Name: "t", Payloads: []byte(`"d"`)})
	}
	actual := args.Map{"val": pc.GetPagesSize(10), "zero": pc.GetPagesSize(0)}
	expected := args.Map{"val": 3, "zero": 0}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- GetPagesSize", actual)
}

func Test_Cov12_PC_IsEqual_BothNil(t *testing.T) {
	var p1, p2 *corepayload.PayloadsCollection
	actual := args.Map{"val": p1.IsEqual(p2)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- IsEqual both nil", actual)
}

func Test_Cov12_PC_IsEqual_OneNil(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"val": pc.IsEqual(nil)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PC returns nil -- IsEqual one nil", actual)
}

func Test_Cov12_PC_IsEqual_DiffLen(t *testing.T) {
	pc1 := corepayload.New.PayloadsCollection.Empty()
	pc1.Add(corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"d"`)})
	pc2 := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"val": pc1.IsEqual(pc2)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PC returns correct value -- IsEqual diff len", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NewPayloadWrapper Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewPW_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()
	actual := args.Map{"notNil": pw != nil, "hasAttr": pw.Attributes != nil}
	expected := args.Map{"notNil": true, "hasAttr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.Empty returns empty -- with args", actual)
}

func Test_Cov12_NewPW_Deserialize(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	b, _ := pw.Serialize()
	pw2, err := corepayload.New.PayloadWrapper.Deserialize(b)
	actual := args.Map{"notNil": pw2 != nil, "noErr": err == nil, "name": pw2.Name}
	expected := args.Map{"notNil": true, "noErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "NewPW.Deserialize returns correct value -- with args", actual)
}

func Test_Cov12_NewPW_Deserialize_Bad(t *testing.T) {
	_, err := corepayload.New.PayloadWrapper.Deserialize([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.Deserialize returns correct value -- bad", actual)
}

func Test_Cov12_NewPW_CastOrDeserializeFrom_Nil(t *testing.T) {
	_, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.CastOrDeserializeFrom returns nil -- nil", actual)
}

func Test_Cov12_NewPW_DeserializeToMany_Bad(t *testing.T) {
	_, err := corepayload.New.PayloadWrapper.DeserializeToMany([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.DeserializeToMany returns correct value -- bad", actual)
}

func Test_Cov12_NewPW_DeserializeUsingJsonResult(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`"d"`))
	jr := pw.JsonPtr()
	pw2, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(jr)
	actual := args.Map{"notNil": pw2 != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.DeserializeUsingJsonResult returns correct value -- with args", actual)
}

func Test_Cov12_NewPW_DeserializeUsingJsonResult_Bad(t *testing.T) {
	jr := corejson.NewResult.UsingBytes([]byte("{bad"))
	_, err := corepayload.New.PayloadWrapper.DeserializeUsingJsonResult(&jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPW.DeserializeUsingJsonResult returns correct value -- bad", actual)
}

func Test_Cov12_NewPW_All(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.All("n", "id", "task", "cat", "entity", true, nil, []byte(`"d"`))
	actual := args.Map{"name": pw.Name, "many": pw.HasManyRecords}
	expected := args.Map{"name": "n", "many": true}
	expected.ShouldBeEqual(t, 0, "NewPW.All returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NewPayloadsCollection Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewPC_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"notNil": pc != nil, "empty": pc.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "NewPC.Empty returns empty -- with args", actual)
}

func Test_Cov12_NewPC_UsingCap(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingCap(10)
	actual := args.Map{"notNil": pc != nil, "empty": pc.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "NewPC.UsingCap returns correct value -- with args", actual)
}

func Test_Cov12_NewPC_UsingWrappers_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.UsingWrappers()
	actual := args.Map{"empty": pc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewPC.UsingWrappers returns empty -- empty", actual)
}

func Test_Cov12_NewPC_Deserialize_Bad(t *testing.T) {
	_, err := corepayload.New.PayloadsCollection.Deserialize([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPC.Deserialize returns correct value -- bad", actual)
}

func Test_Cov12_NewPC_DeserializeToMany_Bad(t *testing.T) {
	_, err := corepayload.New.PayloadsCollection.DeserializeToMany([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeToMany returns correct value -- bad", actual)
}

func Test_Cov12_NewPC_DeserializeUsingJsonResult_Bad(t *testing.T) {
	jr := corejson.NewResult.UsingBytes([]byte("{bad"))
	_, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(&jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeUsingJsonResult returns correct value -- bad", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Generic Helpers
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DeserializePayloadTo_Nil(t *testing.T) {
	_, err := corepayload.DeserializePayloadTo[string](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns nil -- nil", actual)
}

func Test_Cov12_DeserializePayloadTo_Valid(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	s, err := corepayload.DeserializePayloadTo[string](pw)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadTo returns non-empty -- valid", actual)
}

func Test_Cov12_DeserializePayloadToSlice_Nil(t *testing.T) {
	_, err := corepayload.DeserializePayloadToSlice[string](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns nil -- nil", actual)
}

func Test_Cov12_DeserializePayloadToSlice_Valid(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`["a","b"]`)}
	s, err := corepayload.DeserializePayloadToSlice[string](pw)
	actual := args.Map{"len": len(s), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializePayloadToSlice returns non-empty -- valid", actual)
}

func Test_Cov12_DeserializeAttributesPayloadTo_Nil(t *testing.T) {
	_, err := corepayload.DeserializeAttributesPayloadTo[string](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns nil -- nil", actual)
}

func Test_Cov12_DeserializeAttributesPayloadTo_Valid(t *testing.T) {
	a := &corepayload.Attributes{DynamicPayloads: []byte(`"hello"`)}
	s, err := corepayload.DeserializeAttributesPayloadTo[string](a)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadTo returns non-empty -- valid", actual)
}

func Test_Cov12_DeserializeAttributesPayloadToSlice_Nil(t *testing.T) {
	_, err := corepayload.DeserializeAttributesPayloadToSlice[string](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToSlice returns nil -- nil", actual)
}

func Test_Cov12_DeserializeAttributesPayloadToSlice_Valid(t *testing.T) {
	a := &corepayload.Attributes{DynamicPayloads: []byte(`["a","b"]`)}
	s, err := corepayload.DeserializeAttributesPayloadToSlice[string](a)
	actual := args.Map{"len": len(s), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeAttributesPayloadToSlice returns non-empty -- valid", actual)
}
