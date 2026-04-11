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

func Test_TPW_NewTypedPayloadWrapper_Nil(t *testing.T) {
	// Arrange
	_, err := corepayload.NewTypedPayloadWrapper[testUser](nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewTPW returns nil -- nil", actual)
}

func Test_TPW_NewTypedPayloadWrapper_Valid(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`{"name":"alice","email":"a@b.c","age":0}`)}
	tw, err := corepayload.NewTypedPayloadWrapper[testUser](pw)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": tw.Data().Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "alice",
	}
	expected.ShouldBeEqual(t, 0, "NewTPW returns non-empty -- valid", actual)
}

func Test_TPW_NewTypedPayloadWrapperFrom(t *testing.T) {
	// Arrange
	tw, err := corepayload.NewTypedPayloadWrapperFrom[testUser]("n", "id", "testUser", testUser{Name: "alice"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": tw.Data().Name,
		"id": tw.Identifier(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "alice",
		"id": "id",
	}
	expected.ShouldBeEqual(t, 0, "NewTPWFrom returns correct value -- with args", actual)
}

func Test_TPW_Data(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})

	// Act
	actual := args.Map{
		"name": tw.Data().Name,
		"typed": tw.TypedData().Name,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"typed": "alice",
	}
	expected.ShouldBeEqual(t, 0, "Data returns correct value -- with args", actual)
}

func Test_TPW_IsParsed(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	var tw2 *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{
		"parsed": tw.IsParsed(),
		"nil": tw2.IsParsed(),
	}

	// Assert
	expected := args.Map{
		"parsed": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsParsed returns correct value -- with args", actual)
}

func Test_TPW_Name_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.Name()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Name returns nil -- nil", actual)
}

func Test_TPW_Identifier_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.Identifier()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Identifier returns nil -- nil", actual)
}

func Test_TPW_IdString(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id1", testUser{Name: "a"})

	// Act
	actual := args.Map{"val": tw.IdString()}

	// Assert
	expected := args.Map{"val": "id1"}
	expected.ShouldBeEqual(t, 0, "IdString returns correct value -- with args", actual)
}

func Test_TPW_IdInteger_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.IdInteger()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "IdInteger returns nil -- nil", actual)
}

func Test_TPW_EntityType_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.EntityType()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "EntityType returns nil -- nil", actual)
}

func Test_TPW_CategoryName_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.CategoryName()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "CategoryName returns nil -- nil", actual)
}

func Test_TPW_TaskTypeName_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.TaskTypeName()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "TaskTypeName returns nil -- nil", actual)
}

func Test_TPW_HasManyRecords_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.HasManyRecords()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasManyRecords returns nil -- nil", actual)
}

func Test_TPW_HasSingleRecord(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})

	// Act
	actual := args.Map{"val": tw.HasSingleRecord()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSingleRecord returns correct value -- with args", actual)
}

func Test_TPW_Attributes_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"nil": tw.Attributes() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- nil", actual)
}

func Test_TPW_InitializeAttributesOnNull_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"nil": tw.InitializeAttributesOnNull() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "InitializeAttributesOnNull returns nil -- nil", actual)
}

func Test_TPW_HasError_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.HasError()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasError returns nil -- nil", actual)
}

func Test_TPW_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.IsEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns nil -- nil", actual)
}

func Test_TPW_HasItems(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})

	// Act
	actual := args.Map{"val": tw.HasItems()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasItems returns correct value -- with args", actual)
}

func Test_TPW_HasSafeItems(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})

	// Act
	actual := args.Map{"val": tw.HasSafeItems()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_TPW_Error_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"noErr": tw.Error() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns nil -- nil", actual)
}

func Test_TPW_String_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.String()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
}

func Test_TPW_PrettyJsonString_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.PrettyJsonString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns nil -- nil", actual)
}

func Test_TPW_JsonString_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.JsonString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns nil -- nil", actual)
}

func Test_TPW_Json_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	r := tw.Json()

	// Act
	actual := args.Map{"empty": r.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Json returns nil -- nil", actual)
}

func Test_TPW_JsonPtr_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	r := tw.JsonPtr()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr returns nil -- nil", actual)
}

func Test_TPW_MarshalJSON_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	_, err := tw.MarshalJSON()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns nil -- nil", actual)
}

func Test_TPW_MarshalJSON_Valid(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.MarshalJSON()

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
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns non-empty -- valid", actual)
}

func Test_TPW_Serialize_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	_, err := tw.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns nil -- nil", actual)
}

func Test_TPW_Serialize_Valid(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.Serialize()

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
	expected.ShouldBeEqual(t, 0, "Serialize returns non-empty -- valid", actual)
}

func Test_TPW_GetAsString(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")
	val, ok := tw.GetAsString()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsString returns correct value -- with args", actual)
}

func Test_TPW_GetAsInt(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)
	val, ok := tw.GetAsInt()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns correct value -- with args", actual)
}

func Test_TPW_GetAsBool(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "id", "bool", true)
	val, ok := tw.GetAsBool()

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsBool returns correct value -- with args", actual)
}

func Test_TPW_ValueString(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")

	// Act
	actual := args.Map{"val": tw.ValueString()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- with args", actual)
}

func Test_TPW_ValueString_NonString(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)

	// Act
	actual := args.Map{"notEmpty": tw.ValueString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- non-string", actual)
}

func Test_TPW_ValueInt(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[int]("n", "id", "int", 42)

	// Act
	actual := args.Map{"val": tw.ValueInt()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
}

func Test_TPW_ValueInt_NonInt(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")

	// Act
	actual := args.Map{"val": tw.ValueInt()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "ValueInt returns non-empty -- non-int", actual)
}

func Test_TPW_ValueBool(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[bool]("n", "id", "bool", true)

	// Act
	actual := args.Map{"val": tw.ValueBool()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_TPW_ValueBool_NonBool(t *testing.T) {
	// Arrange
	tw, _ := corepayload.NewTypedPayloadWrapperFrom[string]("n", "id", "string", "hello")

	// Act
	actual := args.Map{"val": tw.ValueBool()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "ValueBool returns non-empty -- non-bool", actual)
}

func Test_TPW_SetName(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetName("new")

	// Act
	actual := args.Map{"val": tw.Name()}

	// Assert
	expected := args.Map{"val": "new"}
	expected.ShouldBeEqual(t, 0, "SetName returns correct value -- with args", actual)
}

func Test_TPW_SetIdentifier(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetIdentifier("new-id")

	// Act
	actual := args.Map{"val": tw.Identifier()}

	// Assert
	expected := args.Map{"val": "new-id"}
	expected.ShouldBeEqual(t, 0, "SetIdentifier returns correct value -- with args", actual)
}

func Test_TPW_SetEntityType(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetEntityType("newEntity")

	// Act
	actual := args.Map{"val": tw.EntityType()}

	// Assert
	expected := args.Map{"val": "newEntity"}
	expected.ShouldBeEqual(t, 0, "SetEntityType returns correct value -- with args", actual)
}

func Test_TPW_SetCategoryName(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tw.SetCategoryName("newCat")

	// Act
	actual := args.Map{"val": tw.CategoryName()}

	// Assert
	expected := args.Map{"val": "newCat"}
	expected.ShouldBeEqual(t, 0, "SetCategoryName returns correct value -- with args", actual)
}

func Test_TPW_SetTypedData_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	err := tw.SetTypedData(testUser{Name: "a"})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SetTypedData returns nil -- nil", actual)
}

func Test_TPW_SetTypedData_Valid(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	err := tw.SetTypedData(testUser{Name: "b"})

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": tw.Data().Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "b",
	}
	expected.ShouldBeEqual(t, 0, "SetTypedData returns non-empty -- valid", actual)
}

func Test_TPW_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	c, err := tw.ClonePtr(true)

	// Act
	actual := args.Map{
		"nil": c == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_TPW_ClonePtr_Valid(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	c, err := tw.ClonePtr(true)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": c.Data().Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "a",
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns non-empty -- valid", actual)
}

func Test_TPW_ToPayloadWrapper(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	pw := tw.ToPayloadWrapper()

	// Act
	actual := args.Map{
		"notNil": pw != nil,
		"name": pw.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"name": "n",
	}
	expected.ShouldBeEqual(t, 0, "ToPayloadWrapper returns correct value -- with args", actual)
}

func Test_TPW_ToPayloadWrapper_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"nil": tw.ToPayloadWrapper() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ToPayloadWrapper returns nil -- nil", actual)
}

func Test_TPW_PayloadWrapperValue(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})

	// Act
	actual := args.Map{"notNil": tw.PayloadWrapperValue() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapperValue returns correct value -- with args", actual)
}

func Test_TPW_Reparse_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	err := tw.Reparse()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Reparse returns nil -- nil", actual)
}

func Test_TPW_DynamicPayloads_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"len": len(tw.DynamicPayloads())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicPayloads returns nil -- nil", actual)
}

func Test_TPW_PayloadsString_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.PayloadsString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns nil -- nil", actual)
}

func Test_TPW_Length_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]

	// Act
	actual := args.Map{"val": tw.Length()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_TPW_IsNull(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw2 := newTypedWrapper("n", "id", testUser{Name: "a"})

	// Act
	actual := args.Map{
		"nil": tw.IsNull(),
		"notNil": tw2.IsNull(),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"notNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNull returns correct value -- with args", actual)
}

func Test_TPW_Clear_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_TPW_Dispose_Nil(t *testing.T) {
	// Arrange
	var tw *corepayload.TypedPayloadWrapper[testUser]
	tw.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_TPW_TypedDataJson(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	r := tw.TypedDataJson()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDataJson returns correct value -- with args", actual)
}

func Test_TPW_TypedDataJsonPtr(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	r := tw.TypedDataJsonPtr()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDataJsonPtr returns correct value -- with args", actual)
}

func Test_TPW_TypedDataJsonBytes(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	b, err := tw.TypedDataJsonBytes()

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
	expected.ShouldBeEqual(t, 0, "TypedDataJsonBytes returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedPayloadCollection — Core methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_TPC_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{
		"empty": tc.IsEmpty(),
		"len": tc.Length(),
		"count": tc.Count(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"len": 0,
		"count": 0,
	}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Empty", actual)
}

func Test_TPC_NewWithCap(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollection[testUser](10)

	// Act
	actual := args.Map{"empty": tc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- NewWithCap", actual)
}

func Test_TPC_Items_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]

	// Act
	actual := args.Map{"nil": tc.Items() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Items nil", actual)
}

func Test_TPC_Length_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]

	// Act
	actual := args.Map{"val": tc.Length()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Length nil", actual)
}

func Test_TPC_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]

	// Act
	actual := args.Map{"val": tc.IsEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- IsEmpty nil", actual)
}

func Test_TPC_HasItems_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]

	// Act
	actual := args.Map{"val": tc.HasItems()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- HasItems nil", actual)
}

func Test_TPC_Add_And_Access(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	tc.Add(tw)

	// Act
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

	// Assert
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

func Test_TPC_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{"nil": tc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- FirstOrDefault empty", actual)
}

func Test_TPC_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{"nil": tc.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- LastOrDefault empty", actual)
}

func Test_TPC_SafeAt(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tw := newTypedWrapper("n", "id", testUser{Name: "alice"})
	tc.Add(tw)

	// Act
	actual := args.Map{
		"valid": tc.SafeAt(0) != nil,
		"oob": tc.SafeAt(5) == nil,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"oob": true,
	}
	expected.ShouldBeEqual(t, 0, "SafeAt returns correct value -- with args", actual)
}

func Test_TPC_RemoveAt(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "a"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "b"}))
	ok := tc.RemoveAt(0)
	notOk := tc.RemoveAt(99)

	// Act
	actual := args.Map{
		"ok": ok,
		"notOk": notOk,
		"len": tc.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"notOk": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAt returns correct value -- with args", actual)
}

func Test_TPC_AllData(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	data := tc.AllData()

	// Act
	actual := args.Map{
		"len": len(data),
		"first": data[0].Name,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "alice",
	}
	expected.ShouldBeEqual(t, 0, "AllData returns correct value -- with args", actual)
}

func Test_TPC_AllData_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	data := tc.AllData()

	// Act
	actual := args.Map{"len": len(data)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllData returns empty -- empty", actual)
}

func Test_TPC_AllNames(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("myName", "1", testUser{Name: "a"}))
	names := tc.AllNames()

	// Act
	actual := args.Map{
		"len": len(names),
		"first": names[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "myName",
	}
	expected.ShouldBeEqual(t, 0, "AllNames returns correct value -- with args", actual)
}

func Test_TPC_AllIdentifiers(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id1", testUser{Name: "a"}))
	ids := tc.AllIdentifiers()

	// Act
	actual := args.Map{
		"len": len(ids),
		"first": ids[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "id1",
	}
	expected.ShouldBeEqual(t, 0, "AllIdentifiers returns correct value -- with args", actual)
}

func Test_TPC_ToPayloadsCollection(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	pc := tc.ToPayloadsCollection()

	// Act
	actual := args.Map{"len": pc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns correct value -- with args", actual)
}

func Test_TPC_ToPayloadsCollection_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	pc := tc.ToPayloadsCollection()

	// Act
	actual := args.Map{"empty": pc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns empty -- empty", actual)
}

func Test_TPC_Clone_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	c, err := tc.Clone()

	// Act
	actual := args.Map{
		"empty": c.IsEmpty(),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Clone empty", actual)
}

func Test_TPC_Clone_Valid(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	c, err := tc.Clone()

	// Act
	actual := args.Map{
		"len": c.Length(),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- Clone valid", actual)
}

func Test_TPC_Clear_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]
	tc.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Clear nil", actual)
}

func Test_TPC_Dispose_Nil(t *testing.T) {
	// Arrange
	var tc *corepayload.TypedPayloadCollection[testUser]
	tc.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TPC returns nil -- Dispose nil", actual)
}

func Test_TPC_IsValid(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))

	// Act
	actual := args.Map{"val": tc.IsValid()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "TPC returns non-empty -- IsValid", actual)
}

func Test_TPC_HasErrors(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))

	// Act
	actual := args.Map{"val": tc.HasErrors()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- HasErrors", actual)
}

func Test_TPC_FirstError(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))

	// Act
	actual := args.Map{"nil": tc.FirstError() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- FirstError", actual)
}

func Test_TPC_Errors_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{"nil": tc.Errors() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns empty -- Errors empty", actual)
}

func Test_TPC_MergedError_None(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))

	// Act
	actual := args.Map{"nil": tc.MergedError() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TPC returns error -- MergedError none", actual)
}

func Test_TPC_GetPagesSize(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	for i := 0; i < 25; i++ {
		tc.Add(newTypedWrapper("n", "id", testUser{Name: "a"}))
	}

	// Act
	actual := args.Map{
		"val": tc.GetPagesSize(10),
		"zero": tc.GetPagesSize(0),
	}

	// Assert
	expected := args.Map{
		"val": 3,
		"zero": 0,
	}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- GetPagesSize", actual)
}

func Test_TPC_Filter(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	filtered := tc.Filter(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- Filter", actual)
}

func Test_TPC_FilterByData(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	filtered := tc.FilterByData(func(data testUser) bool {
		return data.Name == "bob"
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FilterByData", actual)
}

func Test_TPC_FirstByName(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("myName", "1", testUser{Name: "alice"}))
	item := tc.FirstByName("myName")
	noItem := tc.FirstByName("unknown")

	// Act
	actual := args.Map{
		"found": item != nil,
		"notFound": noItem == nil,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": true,
	}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FirstByName", actual)
}

func Test_TPC_FirstById(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id1", testUser{Name: "alice"}))
	item := tc.FirstById("id1")

	// Act
	actual := args.Map{"found": item != nil}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- FirstById", actual)
}

func Test_TPC_CountFunc(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	count := tc.CountFunc(func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return true
	})

	// Act
	actual := args.Map{"val": count}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "TPC returns correct value -- CountFunc", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Typed Collection Funcs
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapTypedPayloads_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "alice"}))
	names := corepayload.MapTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Data().Name
	})

	// Act
	actual := args.Map{
		"len": len(names),
		"first": names[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "alice",
	}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns correct value -- with args", actual)
}

func Test_MapTypedPayloads_Empty_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	names := corepayload.MapTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns empty -- empty", actual)
}

func Test_MapTypedPayloadData_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "id", testUser{Name: "alice"}))
	names := corepayload.MapTypedPayloadData[testUser, string](tc, func(data testUser) string {
		return data.Name
	})

	// Act
	actual := args.Map{
		"len": len(names),
		"first": names[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "alice",
	}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloadData returns correct value -- with args", actual)
}

func Test_ReduceTypedPayloads_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	count := corepayload.ReduceTypedPayloads[testUser, int](tc, 0, func(acc int, item *corepayload.TypedPayloadWrapper[testUser]) int {
		return acc + 1
	})

	// Act
	actual := args.Map{"val": count}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns correct value -- with args", actual)
}

func Test_ReduceTypedPayloads_Empty_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	count := corepayload.ReduceTypedPayloads[testUser, int](tc, 0, func(acc int, _ *corepayload.TypedPayloadWrapper[testUser]) int {
		return acc + 1
	})

	// Act
	actual := args.Map{"val": count}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns empty -- empty", actual)
}

func Test_ReduceTypedPayloadData_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	result := corepayload.ReduceTypedPayloadData[testUser, string](tc, "", func(acc string, data testUser) string {
		return acc + data.Name
	})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "alice"}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloadData returns correct value -- with args", actual)
}

func Test_AnyTypedPayload(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("n", "1", testUser{Name: "alice"}))
	found := corepayload.AnyTypedPayload[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})
	notFound := corepayload.AnyTypedPayload[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "bob"
	})

	// Act
	actual := args.Map{
		"found": found,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns correct value -- with args", actual)
}

func Test_AnyTypedPayload_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{"val": corepayload.AnyTypedPayload[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return true })}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns empty -- empty", actual)
}

func Test_AllTypedPayloads(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	all := corepayload.AllTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name != ""
	})
	notAll := corepayload.AllTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "bob"
	})

	// Act
	actual := args.Map{
		"all": all,
		"notAll": notAll,
	}

	// Assert
	expected := args.Map{
		"all": true,
		"notAll": false,
	}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns correct value -- with args", actual)
}

func Test_AllTypedPayloads_Empty(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()

	// Act
	actual := args.Map{"val": corepayload.AllTypedPayloads[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return false })}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns empty -- empty", actual)
}

func Test_PartitionTypedPayloads_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	matching, notMatching := corepayload.PartitionTypedPayloads[testUser](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
		return item.Data().Name == "alice"
	})

	// Act
	actual := args.Map{
		"match": matching.Length(),
		"noMatch": notMatching.Length(),
	}

	// Assert
	expected := args.Map{
		"match": 1,
		"noMatch": 1,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns correct value -- with args", actual)
}

func Test_PartitionTypedPayloads_Empty_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	m, nm := corepayload.PartitionTypedPayloads[testUser](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) bool { return true })

	// Act
	actual := args.Map{
		"m": m.Length(),
		"nm": nm.Length(),
	}

	// Assert
	expected := args.Map{
		"m": 0,
		"nm": 0,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns empty -- empty", actual)
}

func Test_GroupTypedPayloads_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	tc.Add(newTypedWrapper("a", "1", testUser{Name: "alice"}))
	tc.Add(newTypedWrapper("b", "2", testUser{Name: "bob"}))
	groups := corepayload.GroupTypedPayloads[testUser, string](tc, func(item *corepayload.TypedPayloadWrapper[testUser]) string {
		return item.Data().Name
	})

	// Act
	actual := args.Map{"len": len(groups)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns correct value -- with args", actual)
}

func Test_GroupTypedPayloads_Empty_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.EmptyTypedPayloadCollection[testUser]()
	groups := corepayload.GroupTypedPayloads[testUser, string](tc, func(_ *corepayload.TypedPayloadWrapper[testUser]) string { return "" })

	// Act
	actual := args.Map{"len": len(groups)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns empty -- empty", actual)
}

func Test_NewTypedPayloadCollectionSingle_Nil_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc := corepayload.NewTypedPayloadCollectionSingle[testUser](nil)

	// Act
	actual := args.Map{"empty": tc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewTPC returns nil -- Single nil", actual)
}

func Test_NewTypedPayloadCollectionSingle_Valid(t *testing.T) {
	// Arrange
	tw := newTypedWrapper("n", "id", testUser{Name: "a"})
	tc := corepayload.NewTypedPayloadCollectionSingle[testUser](tw)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewTPC returns non-empty -- Single valid", actual)
}

func Test_NewTypedPayloadCollectionFromData_Empty_FromTPWNewTypedPayloadWr(t *testing.T) {
	// Arrange
	tc, err := corepayload.NewTypedPayloadCollectionFromData[testUser]("n", nil)

	// Act
	actual := args.Map{
		"empty": tc.IsEmpty(),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTPC returns empty -- FromData empty", actual)
}

func Test_NewTypedPayloadCollectionFromData_Valid(t *testing.T) {
	// Arrange
	tc, err := corepayload.NewTypedPayloadCollectionFromData[testUser]("n", []testUser{{Name: "a"}, {Name: "b"}})

	// Act
	actual := args.Map{
		"len": tc.Length(),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTPC returns non-empty -- FromData valid", actual)
}
