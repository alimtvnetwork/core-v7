package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// TypedPayloadWrapper — Core methods
// ══════════════════════════════════════════════════════════════════════════════

// testUser is declared in TypedCollection_testcases.go

func newTypedWrapper(name, id string, data testUser) *corepayload.TypedPayloadWrapper[testUser] {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[testUser](name, id, "testUser", data)
	return tw
}

func Test_Cov13_TPW_NewTypedPayloadWrapper_Nil(t *testing.T) {
	_, err := corepayload.NewTypedPayloadWrapper[testUser](nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewTPW returns nil -- nil", actual)
}

func Test_Cov13_TPW_NewTypedPayloadWrapper_Valid(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`{"name":"alice","email":"a@b.c","age":0}`)}
	tw, err := corepayload.NewTypedPayloadWrapper[testUser](pw)
	actual := args.Map{"noErr": err == nil, "name": tw.Data().Name}
	expected := args.Map{"noErr": true, "name": "alice"}
	expected.ShouldBeEqual(t, 0, "NewTPW returns non-empty -- valid", actual)
}

func Test_Cov13_TPW_NewTypedPayloadWrapperFrom(t *testing.T) {
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})
	actual := args.Map{"noErr": err == nil, "name": tw.Data().Name, "id": tw.Identifier()}
	expected := args.Map{"noErr": true, "name": "alice", "id": "id"}
	expected.ShouldBeEqual(t, 0, "NewTPWFrom returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Data(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	actual := args.Map{"name": tw.Data().Name, "typed": tw.TypedData().Name}
	expected := args.Map{"name": "alice", "typed": "alice"}
	expected.ShouldBeEqual(t, 0, "Data returns correct value -- with args", actual)
}

func Test_Cov13_TPW_IsParsed(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	var tw2 *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"parsed": tw.IsParsed(), "nil": tw2.IsParsed()}
	expected := args.Map{"parsed": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsParsed returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Name_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.Name()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Name returns nil -- nil", actual)
}

func Test_Cov13_TPW_Identifier_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.Identifier()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Identifier returns nil -- nil", actual)
}

func Test_Cov13_TPW_IdString(t *testing.T) {
	tw := newTypedWrapper("n", "id1", testUser{Name: "a"})
	actual := args.Map{"val": tw.IdString()}
	expected := args.Map{"val": "id1"}
	expected.ShouldBeEqual(t, 0, "IdString returns correct value -- with args", actual)
}

func Test_Cov13_TPW_IdInteger_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.IdInteger()}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "IdInteger returns nil -- nil", actual)
}

func Test_Cov13_TPW_EntityType_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.EntityType()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "EntityType returns nil -- nil", actual)
}

func Test_Cov13_TPW_CategoryName_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.CategoryName()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "CategoryName returns nil -- nil", actual)
}

func Test_Cov13_TPW_TaskTypeName_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.TaskTypeName()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "TaskTypeName returns nil -- nil", actual)
}

func Test_Cov13_TPW_HasManyRecords_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.HasManyRecords()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasManyRecords returns nil -- nil", actual)
}

func Test_Cov13_TPW_HasSingleRecord(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	actual := args.Map{"val": tw.HasSingleRecord()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSingleRecord returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Attributes_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"nil": tw.Attributes() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- nil", actual)
}

func Test_Cov13_TPW_InitializeAttributesOnNull_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"nil": tw.InitializeAttributesOnNull() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "InitializeAttributesOnNull returns nil -- nil", actual)
}

func Test_Cov13_TPW_HasError_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.HasError()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasError returns nil -- nil", actual)
}

func Test_Cov13_TPW_IsEmpty_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.IsEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns nil -- nil", actual)
}

func Test_Cov13_TPW_HasItems(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	actual := args.Map{"val": tw.HasItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasItems returns correct value -- with args", actual)
}

func Test_Cov13_TPW_HasSafeItems(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	actual := args.Map{"val": tw.HasSafeItems()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Error_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"noErr": tw.Error() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- nil", actual)
}

func Test_Cov13_TPW_String_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.String()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
}

func Test_Cov13_TPW_PrettyJsonString_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.PrettyJsonString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns nil -- nil", actual)
}

func Test_Cov13_TPW_JsonString_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.JsonString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns nil -- nil", actual)
}

func Test_Cov13_TPW_Json_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	r := tw.Json()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Json returns nil -- nil", actual)
}

func Test_Cov13_TPW_JsonPtr_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	r := tw.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr returns nil -- nil", actual)
}

func Test_Cov13_TPW_MarshalJSON_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	_, err := tw.MarshalJSON()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns nil -- nil", actual)
}

func Test_Cov13_TPW_MarshalJSON_Valid(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns non-empty -- valid", actual)
}

func Test_Cov13_TPW_Serialize_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	_, err := tw.Serialize()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns nil -- nil", actual)
}

func Test_Cov13_TPW_Serialize_Valid(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns non-empty -- valid", actual)
}

func Test_Cov13_TPW_GetAsString(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")
	val, ok := tw.GetAsString()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": "hello", "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsString returns correct value -- with args", actual)
}

func Test_Cov13_TPW_GetAsInt(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)
	val, ok := tw.GetAsInt()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns correct value -- with args", actual)
}

func Test_Cov13_TPW_GetAsBool(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "id", "bool", true)
	val, ok := tw.GetAsBool()
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "GetAsBool returns correct value -- with args", actual)
}

func Test_Cov13_TPW_ValueString(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")
	actual := args.Map{"val": tw.ValueString()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- with args", actual)
}

func Test_Cov13_TPW_ValueString_NonString(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)
	actual := args.Map{"notEmpty": tw.ValueString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- non-string", actual)
}

func Test_Cov13_TPW_ValueInt(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)
	actual := args.Map{"val": tw.ValueInt()}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
}

func Test_Cov13_TPW_ValueInt_NonInt(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")
	actual := args.Map{"val": tw.ValueInt()}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ValueInt returns non-empty -- non-int", actual)
}

func Test_Cov13_TPW_ValueBool(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "id", "bool", true)
	actual := args.Map{"val": tw.ValueBool()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_Cov13_TPW_ValueBool_NonBool(t *testing.T) {
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")
	actual := args.Map{"val": tw.ValueBool()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "ValueBool returns non-empty -- non-bool", actual)
}

func Test_Cov13_TPW_SetName(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetName("new")
	actual := args.Map{"val": tw.Name()}
	expected := args.Map{"val": "new"}
	expected.ShouldBeEqual(t, 0, "SetName returns correct value -- with args", actual)
}

func Test_Cov13_TPW_SetIdentifier(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetIdentifier("new-id")
	actual := args.Map{"val": tw.Identifier()}
	expected := args.Map{"val": "new-id"}
	expected.ShouldBeEqual(t, 0, "SetIdentifier returns correct value -- with args", actual)
}

func Test_Cov13_TPW_SetEntityType(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetEntityType("newEntity")
	actual := args.Map{"val": tw.EntityType()}
	expected := args.Map{"val": "newEntity"}
	expected.ShouldBeEqual(t, 0, "SetEntityType returns correct value -- with args", actual)
}

func Test_Cov13_TPW_SetCategoryName(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetCategoryName("newCat")
	actual := args.Map{"val": tw.CategoryName()}
	expected := args.Map{"val": "newCat"}
	expected.ShouldBeEqual(t, 0, "SetCategoryName returns correct value -- with args", actual)
}

func Test_Cov13_TPW_SetTypedData_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	err := tw.SetTypedData(testUser{Name: "a"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SetTypedData returns nil -- nil", actual)
}

func Test_Cov13_TPW_SetTypedData_Valid(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	err := tw.SetTypedData(testUser{Name: "b"})
	actual := args.Map{"noErr": err == nil, "name": tw.Data().Name}
	expected := args.Map{"noErr": true, "name": "b"}
	expected.ShouldBeEqual(t, 0, "SetTypedData returns non-empty -- valid", actual)
}

func Test_Cov13_TPW_ClonePtr_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	c, err := tw.ClonePtr(true)
	actual := args.Map{"nil": c == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov13_TPW_ClonePtr_Valid(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	c, err := tw.ClonePtr(true)
	actual := args.Map{"noErr": err == nil, "name": c.Data().Name}
	expected := args.Map{"noErr": true, "name": "a"}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns non-empty -- valid", actual)
}

func Test_Cov13_TPW_ToPayloadWrapper(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	pw := tw.ToPayloadWrapper()
	actual := args.Map{"notNil": pw != nil, "name": pw.Name}
	expected := args.Map{"notNil": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "ToPayloadWrapper returns correct value -- with args", actual)
}

func Test_Cov13_TPW_ToPayloadWrapper_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"nil": tw.ToPayloadWrapper() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ToPayloadWrapper returns nil -- nil", actual)
}

func Test_Cov13_TPW_PayloadWrapperValue(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	actual := args.Map{"notNil": tw.PayloadWrapperValue() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapperValue returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Reparse_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	err := tw.Reparse()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Reparse returns nil -- nil", actual)
}

func Test_Cov13_TPW_DynamicPayloads_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"len": len(tw.DynamicPayloads())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns nil -- nil", actual)
}

func Test_Cov13_TPW_PayloadsString_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.PayloadsString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns nil -- nil", actual)
}

func Test_Cov13_TPW_Length_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	actual := args.Map{"val": tw.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Cov13_TPW_IsNull(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw2 := newTypedWrapper("n", "id", testUser{Name: "a"})
	actual := args.Map{"nil": tw.IsNull(), "notNil": tw2.IsNull()}
	expected := args.Map{"nil": true, "notNil": false}
	expected.ShouldBeEqual(t, 0, "IsNull returns correct value -- with args", actual)
}

func Test_Cov13_TPW_Clear_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Cov13_TPW_Dispose_Nil(t *testing.T) {
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Cov13_TPW_TypedDataJson(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	r := tw.TypedDataJson()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDataJson returns correct value -- with args", actual)
}

func Test_Cov13_TPW_TypedDataJsonPtr(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	r := tw.TypedDataJsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDataJsonPtr returns correct value -- with args", actual)
}

func Test_Cov13_TPW_TypedDataJsonBytes(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.TypedDataJsonBytes()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDataJsonBytes returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedPayloadCollection — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_TPC_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"empty": tc.IsEmpty(), "len": tc.Length(), "count": tc.Count()}
	expected := args.Map{"empty": true, "len": 0, "count": 0}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Empty", actual)
}

func Test_Cov13_TPC_NewWithCap(t *testing.T) {
	tc := corepayload.NewTypedPayloadCollection[testUser](10)
	actual := args.Map{"empty": tc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- NewWithCap", actual)
}

func Test_Cov13_TPC_Items_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	actual := args.Map{"nil": tc.Items() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Items nil", actual)
}

func Test_Cov13_TPC_Length_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	actual := args.Map{"val": tc.Length()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Length nil", actual)
}

func Test_Cov13_TPC_IsEmpty_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	actual := args.Map{"val": tc.IsEmpty()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- IsEmpty nil", actual)
}

func Test_Cov13_TPC_HasItems_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	actual := args.Map{"val": tc.HasItems()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- HasItems nil", actual)
}

func Test_Cov13_TPC_Add_And_Access(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	tc.Add(tw)
	actual := args.Map{
		"len":       tc.Length(),
		"hasItems":  tc.HasItems(),
		"hasAny":    tc.HasAnyItem(),
		"lastIdx":   tc.LastIndex(),
		"hasIdx0":   tc.HasIndex(0),
		"hasIdx99":  tc.HasIndex(99),
		"firstName": tc.First().Data().Name,
		"lastName":  tc.Last().Data().Name,
	}
	expected := args.Map{
		"len":       1,
		"hasItems":  true,
		"hasAny":    true,
		"lastIdx":   0,
		"hasIdx0":   true,
		"hasIdx99":  false,
		"firstName": "alice",
		"lastName":  "alice",
	}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- Add and Access", actual)
}

func Test_Cov13_TPC_FirstOrDefault_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"nil": tc.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- FirstOrDefault empty", actual)
}

func Test_Cov13_TPC_LastOrDefault_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"nil": tc.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- LastOrDefault empty", actual)
}

func Test_Cov13_TPC_SafeAt(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	tc.Add(tw)
	actual := args.Map{"valid": tc.SafeAt(0) != nil, "oob": tc.SafeAt(5) == nil}
	expected := args.Map{"valid": true, "oob": true}
	expected.ShouldBeEqual(t, 0, "SafeAt returns correct value -- with args", actual)
}

func Test_Cov13_TPC_RemoveAt(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "a"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "b"}))
	ok := tc.RemoveAt(0)
	notOk := tc.RemoveAt(99)
	actual := args.Map{"ok": ok, "notOk": notOk, "len": tc.Length()}
	expected := args.Map{"ok": true, "notOk": false, "len": 1}
	expected.ShouldBeEqual(t, 0, "RemoveAt returns correct value -- with args", actual)
}

func Test_Cov13_TPC_AllData(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	data := tc.AllData()
	actual := args.Map{"len": len(data), "first": data[0].Name}
	expected := args.Map{"len": 2, "first": "alice"}
	expected.ShouldBeEqual(t, 0, "AllData returns correct value -- with args", actual)
}

func Test_Cov13_TPC_AllData_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	data := tc.AllData()
	actual := args.Map{"len": len(data)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllData returns empty -- empty", actual)
}

func Test_Cov13_TPC_AllNames(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("myName", "1", testUser{Name: "a"}))
	names := tc.AllNames()
	actual := args.Map{"len": len(names), "first": names[0]}
	expected := args.Map{"len": 1, "first": "myName"}
	expected.ShouldBeEqual(t, 0, "AllNames returns correct value -- with args", actual)
}

func Test_Cov13_TPC_AllIdentifiers(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id1", testUser{Name: "a"}))
	ids := tc.AllIdentifiers()
	actual := args.Map{"len": len(ids), "first": ids[0]}
	expected := args.Map{"len": 1, "first": "id1"}
	expected.ShouldBeEqual(t, 0, "AllIdentifiers returns correct value -- with args", actual)
}

func Test_Cov13_TPC_ToPayloadsCollection(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	pc := tc.ToPayloadsCollection()
	actual := args.Map{"len": pc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns correct value -- with args", actual)
}

func Test_Cov13_TPC_ToPayloadsCollection_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	pc := tc.ToPayloadsCollection()
	actual := args.Map{"empty": pc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns empty -- empty", actual)
}

func Test_Cov13_TPC_Clone_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	c, err := tc.Clone()
	actual := args.Map{"empty": c.IsEmpty(), "noErr": err == nil}
	expected := args.Map{"empty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Clone empty", actual)
}

func Test_Cov13_TPC_Clone_Valid(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	c, err := tc.Clone()
	actual := args.Map{"len": c.Length(), "noErr": err == nil}
	expected := args.Map{"len": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- Clone valid", actual)
}

func Test_Cov13_TPC_Clear_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	tc.Clear() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Clear nil", actual)
}

func Test_Cov13_TPC_Dispose_Nil(t *testing.T) {
	var tc *corepayload.TypedPayloadCollection[testUser]
	tc.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Dispose nil", actual)
}

func Test_Cov13_TPC_IsValid(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	actual := args.Map{"val": tc.IsValid()}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- IsValid", actual)
}

func Test_Cov13_TPC_HasErrors(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	actual := args.Map{"val": tc.HasErrors()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- HasErrors", actual)
}

func Test_Cov13_TPC_FirstError(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	actual := args.Map{"nil": tc.FirstError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- FirstError", actual)
}

func Test_Cov13_TPC_Errors_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"nil": tc.Errors() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Errors empty", actual)
}

func Test_Cov13_TPC_MergedError_None(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	actual := args.Map{"nil": tc.MergedError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- MergedError none", actual)
}

func Test_Cov13_TPC_GetPagesSize(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	for i := 0; i < 25; i++ {
		tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	}
	actual := args.Map{"val": tc.GetPagesSize(10), "zero": tc.GetPagesSize(0)}
	expected := args.Map{"val": 3, "zero": 0}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- GetPagesSize", actual)
}

func Test_Cov13_TPC_Filter(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	filtered := tc.Filter(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- Filter", actual)
}

func Test_Cov13_TPC_FilterByData(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	filtered := tc.FilterByData(func(data testUser) bool {
		return data.Name == "bob"
	})
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FilterByData", actual)
}

func Test_Cov13_TPC_FirstByName(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("myName", "1", testUser{Name: "alice"}))
	item := tc.FirstByName("myName")
	noItem := tc.FirstByName("unknown")
	actual := args.Map{"found": item != nil, "notFound": noItem == nil}
	expected := args.Map{"found": true, "notFound": true}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FirstByName", actual)
}

func Test_Cov13_TPC_FirstById(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id1", testUser{Name: "alice"}))
	item := tc.FirstById("id1")
	actual := args.Map{"found": item != nil}
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FirstById", actual)
}

func Test_Cov13_TPC_CountFunc(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	count := tc.CountFunc(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return true
	})
	actual := args.Map{"val": count}
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- CountFunc", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Typed Collection Funcs
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov13_MapTypedPayloads(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "alice"}))
	names := corepayload.MapTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Data().Name
	})
	actual := args.Map{"len": len(names), "first": names[0]}
	expected := args.Map{"len": 1, "first": "alice"}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns correct value -- with args", actual)
}

func Test_Cov13_MapTypedPayloads_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	names := corepayload.MapTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return ""
	})
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns empty -- empty", actual)
}

func Test_Cov13_MapTypedPayloadData(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "alice"}))
	names := corepayload.MapTypedPayloadData[testUser, string](tc, func(data testUser) string {
		return data.Name
	})
	actual := args.Map{"len": len(names), "first": names[0]}
	expected := args.Map{"len": 1, "first": "alice"}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloadData returns correct value -- with args", actual)
}

func Test_Cov13_ReduceTypedPayloads(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	count := corepayload.ReduceTypedPayloads[testUser, int](tc, 0, func(acc int, item *corepayload.TypedPayloadWrapper[testUser]) int {
		return acc + 1
	})
	actual := args.Map{"val": count}
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns correct value -- with args", actual)
}

func Test_Cov13_ReduceTypedPayloads_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	count := corepayload.ReduceTypedPayloads[testUser, int](tc, 0, func(acc int, _ *corepayload.TypedPayloadWrapper[testUser]) int {
		return acc + 1
	})
	actual := args.Map{"val": count}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns empty -- empty", actual)
}

func Test_Cov13_ReduceTypedPayloadData(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	result := corepayload.ReduceTypedPayloadData[testUser, string](tc, "", func(acc string, data testUser) string {
		return acc + data.Name
	})
	actual := args.Map{"val": result}
	expected := args.Map{"val": "alice"}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloadData returns correct value -- with args", actual)
}

func Test_Cov13_AnyTypedPayload(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "1", testUser{Name: "alice"}))
	found := corepayload.AnyTypedPayload[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})
	notFound := corepayload.AnyTypedPayload[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "bob"
	})
	actual := args.Map{"found": found, "notFound": notFound}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns correct value -- with args", actual)
}

func Test_Cov13_AnyTypedPayload_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"val": corepayload.AnyTypedPayload[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return true })}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns empty -- empty", actual)
}

func Test_Cov13_AllTypedPayloads(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	all := corepayload.AllTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name != ""
	})
	notAll := corepayload.AllTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "bob"
	})
	actual := args.Map{"all": all, "notAll": notAll}
	expected := args.Map{"all": true, "notAll": false}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns correct value -- with args", actual)
}

func Test_Cov13_AllTypedPayloads_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	actual := args.Map{"val": corepayload.AllTypedPayloads[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return false })}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns empty -- empty", actual)
}

func Test_Cov13_PartitionTypedPayloads(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	matching, notMatching := corepayload.PartitionTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})
	actual := args.Map{"match": matching.Length(), "noMatch": notMatching.Length()}
	expected := args.Map{"match": 1, "noMatch": 1}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns correct value -- with args", actual)
}

func Test_Cov13_PartitionTypedPayloads_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	m, nm := corepayload.PartitionTypedPayloads[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return true })
	actual := args.Map{"m": m.Length(), "nm": nm.Length()}
	expected := args.Map{"m": 0, "nm": 0}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns empty -- empty", actual)
}

func Test_Cov13_GroupTypedPayloads(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	groups := corepayload.GroupTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Data().Name
	})
	actual := args.Map{"len": len(groups)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns correct value -- with args", actual)
}

func Test_Cov13_GroupTypedPayloads_Empty(t *testing.T) {
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	groups := corepayload.GroupTypedPayloads[testUser, string](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) string { return "" })
	actual := args.Map{"len": len(groups)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns empty -- empty", actual)
}

func Test_Cov13_NewTypedPayloadCollectionSingle_Nil(t *testing.T) {
	tc := corepayload.NewTypedPayloadCollectionSingle[testUser](nil)
	actual := args.Map{"empty": tc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewTPC returns nil -- Single nil", actual)
}

func Test_Cov13_NewTypedPayloadCollectionSingle_Valid(t *testing.T) {
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tc := corepayload.NewTypedPayloadCollectionSingle[testUser](tw)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewTPC returns non-empty -- Single valid", actual)
}

func Test_Cov13_NewTypedPayloadCollectionFromData_Empty(t *testing.T) {
	tc, err := corepayload.NewTypedPayloadCollectionFromData[testUser]("n", nil)
	actual := args.Map{"empty": tc.IsEmpty(), "noErr": err == nil}
	expected := args.Map{"empty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewTPC returns empty -- FromData empty", actual)
}

func Test_Cov13_NewTypedPayloadCollectionFromData_Valid(t *testing.T) {
	tc, err := corepayload.NewTypedPayloadCollectionFromData[testUser]("n", []testUser{{Name: "a"}, {Name: "b"}})
	actual := args.Map{"len": tc.Length(), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewTPC returns non-empty -- FromData valid", actual)
}
