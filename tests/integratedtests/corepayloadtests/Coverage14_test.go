package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov14_Attributes_Empty(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	actual := args.Map{"empty": a.IsEmpty(), "len": a.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Attributes returns empty -- empty", actual)
}

func Test_Cov14_Attributes_KeyValues(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("key", "val")
	v, ok := a.KeyValuePairs.Get("key")
	actual := args.Map{"val": v, "ok": ok, "len": a.StringKeyValuePairsLength()}
	expected := args.Map{"val": "val", "ok": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Attributes returns non-empty -- KeyValues", actual)
}

func Test_Cov14_Attributes_GetStringKeyValue(t *testing.T) {
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("key", "val")
	v, found := a.GetStringKeyValue("key")
	_, notFound := a.GetStringKeyValue("miss")
	actual := args.Map{"val": v, "found": found, "notFound": notFound}
	expected := args.Map{"val": "val", "found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns correct value -- with args", actual)
}

func Test_Cov14_PagingInfo_Empty(t *testing.T) {
	p := corepayload.PagingInfo{}
	actual := args.Map{"page": p.CurrentPageIndex, "size": p.PerPageItems}
	expected := args.Map{"page": 0, "size": 0}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns empty -- empty", actual)
}

func Test_Cov14_PayloadWrapper_UsingBytes(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "42", "task", "cat", "entity", []byte(`{}`))
	actual := args.Map{
		"name":     pw.Name,
		"id":       pw.Identifier,
		"notEmpty": !pw.IsEmpty(),
		"idInt":    pw.IdInteger(),
	}
	expected := args.Map{"name": "n", "id": "42", "notEmpty": true, "idInt": 42}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- UsingBytes", actual)
}

func Test_Cov14_PayloadWrapper_Clone(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`{}`))
	cloned, err := pw.Clone(true)
	actual := args.Map{"noErr": err == nil, "name": cloned.Name}
	expected := args.Map{"noErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- Clone", actual)
}

func Test_Cov14_PayloadWrapper_Clone_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	_, err := pw.Clone(true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns nil -- Clone nil", actual)
}

func Test_Cov14_PayloadWrapper_Dispose(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`{}`))
	pw.Dispose()
	actual := args.Map{"empty": pw.IsEmpty(), "attrNil": pw.Attributes == nil}
	expected := args.Map{"empty": true, "attrNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- Dispose", actual)
}

func Test_Cov14_PayloadsCollection_Empty(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	actual := args.Map{"empty": pc.IsEmpty(), "len": pc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns empty -- empty", actual)
}

func Test_Cov14_PayloadsCollection_Add(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	pw := corepayload.New.PayloadWrapper.UsingBytes("n", "id", "task", "cat", "entity", []byte(`{}`))
	pc.Add(*pw)
	actual := args.Map{"len": pc.Length(), "hasAny": pc.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- add", actual)
}

func Test_Cov14_BytesCreateInstruction(t *testing.T) {
	bci := corepayload.BytesCreateInstruction{Name: "n", Identifier: "id", EntityType: "e", Payloads: []byte(`{}`)}
	actual := args.Map{"name": bci.Name, "id": bci.Identifier}
	expected := args.Map{"name": "n", "id": "id"}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstruction returns correct value -- with args", actual)
}

func Test_Cov14_PayloadCreateInstruction(t *testing.T) {
	pci := corepayload.PayloadCreateInstruction{Name: "n", Identifier: "id", EntityType: "e"}
	actual := args.Map{"name": pci.Name, "id": pci.Identifier}
	expected := args.Map{"name": "n", "id": "id"}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstruction returns correct value -- with args", actual)
}

func Test_Cov14_User_NonSysCreate(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	actual := args.Map{"name": u.Name, "type": u.Type, "sys": u.IsSystemUser}
	expected := args.Map{"name": "alice", "type": "admin", "sys": false}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- NonSysCreate", actual)
}

func Test_Cov14_User_ClonePtr(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("alice", "admin")
	c := u.ClonePtr()
	actual := args.Map{"notNil": c != nil, "name": c.Name}
	expected := args.Map{"notNil": true, "name": "alice"}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- ClonePtr", actual)
}

func Test_Cov14_TypedPayloadCollection_ToPayloadsCollectionJson(t *testing.T) {
	tc := corepayload.NewTypedPayloadCollection[testUser](5)
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})
	if err == nil {
		tc.Add(tw)
	}
	pc := tc.ToPayloadsCollection()
	jsonResult := pc.Json()
	actual := args.Map{"twNoErr": err == nil, "jsonNoErr": !jsonResult.HasError(), "len": pc.Length()}
	expected := args.Map{"twNoErr": true, "jsonNoErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollection returns correct value -- ToPayloadsCollection Json", actual)
}

func Test_Cov14_NewPayloadWrapper_Create_Deserialize(t *testing.T) {
	pw, createErr := corepayload.New.PayloadWrapper.Create("n", "id", "task", "cat", map[string]string{"k": "v"})
	if createErr != nil {
		actual := args.Map{"createErr": true}
		expected := args.Map{"createErr": false}
		expected.ShouldBeEqual(t, 0, "NewPayloadWrapper returns correct value -- Create", actual)
		return
	}
	bytes, serErr := pw.Serialize()
	pw2, deErr := corepayload.New.PayloadWrapper.Deserialize(bytes)
	actual := args.Map{"createNoErr": createErr == nil, "serNoErr": serErr == nil, "deNoErr": deErr == nil, "name": pw2.Name}
	expected := args.Map{"createNoErr": true, "serNoErr": true, "deNoErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "NewPayloadWrapper returns correct value -- Create+Deserialize", actual)
}

func Test_Cov14_Attributes_UsingDynamicPayloadBytes_Deserialize(t *testing.T) {
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"k":"v"}`))
	bytes, serErr := corejson.Serialize.Raw(a)
	a2, deErr := corepayload.New.Attributes.Deserialize(bytes)
	actual := args.Map{"serNoErr": serErr == nil, "deNoErr": deErr == nil, "dynLen": a2.DynamicBytesLength()}
	expected := args.Map{"serNoErr": true, "deNoErr": true, "dynLen": len([]byte(`{"k":"v"}`))}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- UsingDynamicPayloadBytes+Deserialize", actual)
}
