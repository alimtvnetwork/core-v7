package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// PayloadWrapper — core methods
// ═══════════════════════════════════════════

func Test_PayloadWrapper_Basic_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name: "test", Identifier: "id-1", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
		HasManyRecords: false, Payloads: []byte(`"hello"`),
	}

	// Act
	actual := args.Map{
		"name":       pw.PayloadName(),
		"entity":     pw.PayloadEntityType(),
		"category":   pw.PayloadCategory(),
		"taskType":   pw.PayloadTaskType(),
		"idStr":      pw.IdString(),
		"hasAny":     pw.HasAnyItem(),
		"payloadDyn": len(pw.PayloadDynamic()) > 0,
		"dynPayloads": len(pw.DynamicPayloads()) > 0,
		"payloadsStr": pw.PayloadsString() != "",
		"value":      pw.Value() != nil,
	}

	// Assert
	expected := args.Map{
		"name": "test", "entity": "entity", "category": "cat",
		"taskType": "task", "idStr": "id-1", "hasAny": true,
		"payloadDyn": true, "dynPayloads": true, "payloadsStr": true, "value": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_PayloadWrapper_IsChecks(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{
		Name: "test", Identifier: "id", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
	}

	// Act
	actual := args.Map{
		"isName":     pw.IsName("test"),
		"isNotName":  pw.IsName("other"),
		"isId":       pw.IsIdentifier("id"),
		"isTask":     pw.IsTaskTypeName("task"),
		"isEntity":   pw.IsEntityType("entity"),
		"isCat":      pw.IsCategory("cat"),
	}

	// Assert
	expected := args.Map{
		"isName": true, "isNotName": false, "isId": true,
		"isTask": true, "isEntity": true, "isCat": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- IsChecks", actual)
}

func Test_PayloadWrapper_IsEqual(t *testing.T) {
	// Arrange
	pw1 := &corepayload.PayloadWrapper{Name: "test", Identifier: "id", Payloads: []byte("p")}
	pw2 := &corepayload.PayloadWrapper{Name: "test", Identifier: "id", Payloads: []byte("p")}
	pw3 := &corepayload.PayloadWrapper{Name: "other"}
	var nilPW *corepayload.PayloadWrapper

	// Act
	actual := args.Map{
		"equal":     pw1.IsEqual(pw2),
		"notEqual":  pw1.IsEqual(pw3),
		"samePtr":   pw1.IsEqual(pw1),
		"nilBoth":   nilPW.IsEqual(nil),
		"nilLeft":   nilPW.IsEqual(pw1),
		"payEq":     pw1.IsPayloadsEqual([]byte("p")),
		"payNotEq":  pw1.IsPayloadsEqual([]byte("x")),
	}

	// Assert
	expected := args.Map{
		"equal": true, "notEqual": false, "samePtr": true,
		"nilBoth": true, "nilLeft": false, "payEq": true, "payNotEq": false,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- IsEqual", actual)
}

func Test_PayloadWrapper_JSON(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{Name: "test", Payloads: []byte(`"hello"`)}
	jsonStr := pw.JsonString()
	prettyStr := pw.PrettyJsonString()
	str := pw.String()
	b, err := pw.Serialize()

	// Act
	actual := args.Map{
		"jsonNotEmpty":   jsonStr != "",
		"prettyNotEmpty": prettyStr != "",
		"strNotEmpty":    str != "",
		"bLen":           len(b) > 0,
		"noErr":          err == nil,
	}

	// Assert
	expected := args.Map{
		"jsonNotEmpty": true, "prettyNotEmpty": true, "strNotEmpty": true,
		"bLen": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- JSON", actual)
}

func Test_PayloadWrapper_SetDynamicPayloads_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	pw := &corepayload.PayloadWrapper{}
	err := pw.SetDynamicPayloads([]byte("new"))
	var nilPW *corepayload.PayloadWrapper
	nilErr := nilPW.SetDynamicPayloads([]byte("x"))

	// Act
	actual := args.Map{
		"noErr":   err == nil,
		"payload": string(pw.Payloads),
		"nilErr":  nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"payload": "new",
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- SetDynamicPayloads", actual)
}

func Test_PayloadWrapper_All(t *testing.T) {
	// Arrange
	pw := corepayload.PayloadWrapper{
		Name: "n", Identifier: "id", EntityType: "e", CategoryName: "c", Payloads: []byte("p"),
	}
	id, name, entity, cat, dynP := pw.All()

	// Act
	actual := args.Map{
		"id": id,
		"name": name,
		"entity": entity,
		"cat": cat,
		"dynP": string(dynP),
	}

	// Assert
	expected := args.Map{
		"id": "id",
		"name": "n",
		"entity": "e",
		"cat": "c",
		"dynP": "p",
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- All", actual)
}

func Test_PayloadWrapper_AllSafe_Nil_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	var pw *corepayload.PayloadWrapper
	id, name, entity, cat, dynP := pw.AllSafe()

	// Act
	actual := args.Map{
		"id": id,
		"name": name,
		"entity": entity,
		"cat": cat,
		"dynP": string(dynP),
	}

	// Assert
	expected := args.Map{
		"id": "",
		"name": "",
		"entity": "",
		"cat": "",
		"dynP": "",
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns nil -- AllSafe nil", actual)
}

// ═══════════════════════════════════════════
// User — comprehensive
// ═══════════════════════════════════════════

func Test_User_Comprehensive(t *testing.T) {
	// Arrange
	u := &corepayload.User{
		Name: "Alice", Type: "admin", Identifier: "123",
		AuthToken: "token", PasswordHash: "hash", IsSystemUser: false,
	}
	var nilU *corepayload.User

	// Act
	actual := args.Map{
		"isEmpty":      u.IsEmpty(),
		"isValid":      u.IsValidUser(),
		"isNameEmpty":  u.IsNameEmpty(),
		"isNameEq":     u.IsNameEqual("Alice"),
		"hasAuth":      u.HasAuthToken(),
		"hasPwHash":    u.HasPasswordHash(),
		"isPwEmpty":    u.IsPasswordHashEmpty(),
		"isAuthEmpty":  u.IsAuthTokenEmpty(),
		"isNotSysUser": u.IsNotSystemUser(),
		"isVirtual":    u.IsVirtualUser(),
		"hasType":      u.HasType(),
		"isTypeEmpty":  u.IsTypeEmpty(),
		"idInt":        u.IdentifierInteger(),
		"idUint":       u.IdentifierUnsignedInteger(),
		"strNN":        u.String() != "",
		"prettyNN":     u.PrettyJsonString() != "",
		"nilEmpty":     nilU.IsEmpty(),
		"nilPwEmpty":   nilU.IsPasswordHashEmpty(),
		"nilAuthEmpty": nilU.IsAuthTokenEmpty(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "isValid": true, "isNameEmpty": false,
		"isNameEq": true, "hasAuth": true, "hasPwHash": true,
		"isPwEmpty": false, "isAuthEmpty": false,
		"isNotSysUser": true, "isVirtual": true,
		"hasType": true, "isTypeEmpty": false,
		"idInt": 123, "idUint": uint(123),
		"strNN": true, "prettyNN": true,
		"nilEmpty": true, "nilPwEmpty": true, "nilAuthEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- comprehensive", actual)
}

func Test_User_Clone_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "Alice", Type: "admin"}
	cloned := u.Clone()
	clonedPtr := u.ClonePtr()
	var nilU *corepayload.User

	// Act
	actual := args.Map{
		"cloneName":   cloned.Name,
		"cpName":      clonedPtr.Name,
		"nilClonePtr": nilU.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"cloneName": "Alice",
		"cpName": "Alice",
		"nilClonePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- Clone", actual)
}

func Test_User_JSON(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "Alice"}
	j := u.Json()
	jp := u.JsonPtr()
	b, err := u.Serialize()

	// Act
	actual := args.Map{
		"jHas": j.HasBytes(), "jpNN": jp != nil,
		"bLen": len(b) > 0, "noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"jHas": true,
		"jpNN": true,
		"bLen": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- JSON", actual)
}

// ═══════════════════════════════════════════
// AuthInfo — comprehensive
// ═══════════════════════════════════════════

func Test_AuthInfo_Basic(t *testing.T) {
	// Arrange
	ai := &corepayload.AuthInfo{
		Identifier: "123", ActionType: "create", ResourceName: "/api",
	}
	var nilAI *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"isEmpty":         ai.IsEmpty(),
		"hasAny":          ai.HasAnyItem(),
		"isValid":         ai.IsValid(),
		"isActionEmpty":   ai.IsActionTypeEmpty(),
		"isResEmpty":      ai.IsResourceNameEmpty(),
		"hasAction":       ai.HasActionType(),
		"hasResource":     ai.HasResourceName(),
		"idInt":           ai.IdentifierInteger(),
		"idUint":          ai.IdentifierUnsignedInteger(),
		"strNN":           ai.String() != "",
		"nilEmpty":        nilAI.IsEmpty(),
		"nilActionEmpty":  nilAI.IsActionTypeEmpty(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "hasAny": true, "isValid": true,
		"isActionEmpty": false, "isResEmpty": false,
		"hasAction": true, "hasResource": true,
		"idInt": 123, "idUint": uint(123),
		"strNN": true, "nilEmpty": true, "nilActionEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- basic", actual)
}

func Test_AuthInfo_Setters_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("create")
	ai.SetResourceName("/api")
	ai.SetIdentifier("id-1")

	// Act
	actual := args.Map{
		"action": ai.ActionType, "resource": ai.ResourceName, "id": ai.Identifier,
	}

	// Assert
	expected := args.Map{
		"action": "create",
		"resource": "/api",
		"id": "id-1",
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- setters", actual)
}

func Test_AuthInfo_Clone_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	ai := &corepayload.AuthInfo{ActionType: "create"}
	cloned := ai.Clone()
	clonedPtr := ai.ClonePtr()
	var nilAI *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"cloneAction": cloned.ActionType,
		"cpAction":    clonedPtr.ActionType,
		"nilClone":    nilAI.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"cloneAction": "create",
		"cpAction": "create",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- Clone", actual)
}

func Test_AuthInfo_JSON(t *testing.T) {
	// Arrange
	ai := corepayload.AuthInfo{ActionType: "create"}
	j := ai.Json()
	jp := ai.JsonPtr()

	// Act
	actual := args.Map{
		"jHas": j.HasBytes(),
		"jpNN": jp != nil,
	}

	// Assert
	expected := args.Map{
		"jHas": true,
		"jpNN": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- JSON", actual)
}

// ═══════════════════════════════════════════
// PayloadsCollection
// ═══════════════════════════════════════════

func Test_PayloadsCollection_Add_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	pc := &corepayload.PayloadsCollection{}
	pw := corepayload.PayloadWrapper{Name: "test"}
	pc.Add(pw)
	pc.Adds(corepayload.PayloadWrapper{Name: "t2"}, corepayload.PayloadWrapper{Name: "t3"})

	// Act
	actual := args.Map{"len": len(pc.Items)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- Add", actual)
}

func Test_PayloadsCollection_AddsPtr_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	pc := &corepayload.PayloadsCollection{}
	pw := &corepayload.PayloadWrapper{Name: "test"}
	pc.AddsPtr(pw)
	pc.AddsPtr()

	// Act
	actual := args.Map{"len": len(pc.Items)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- AddsPtr", actual)
}

// ═══════════════════════════════════════════
// BytesCreateInstruction
// ═══════════════════════════════════════════

func Test_BytesCreateInstruction_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	bci := corepayload.BytesCreateInstruction{
		Name: "test", Identifier: "id", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
		HasManyRecords: true, Payloads: []byte("payload"),
	}

	// Act
	actual := args.Map{
		"name": bci.Name, "id": bci.Identifier, "task": bci.TaskTypeName,
		"entity": bci.EntityType, "cat": bci.CategoryName,
		"hasMany": bci.HasManyRecords, "payLen": len(bci.Payloads),
	}

	// Assert
	expected := args.Map{
		"name": "test", "id": "id", "task": "task",
		"entity": "entity", "cat": "cat", "hasMany": true, "payLen": 7,
	}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstruction returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// SessionInfo
// ═══════════════════════════════════════════

func Test_SessionInfo_IsEmpty(t *testing.T) {
	// Arrange
	si := &corepayload.SessionInfo{}
	var nilSI *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"isEmpty":  si.IsEmpty(),
		"nilEmpty": nilSI.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns empty -- IsEmpty", actual)
}

func Test_SessionInfo_Clone_FromPayloadWrapperBasic(t *testing.T) {
	// Arrange
	si := &corepayload.SessionInfo{Id: "s1"}
	cloned := si.ClonePtr()
	var nilSI *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"cloneId":  cloned.Id,
		"nilClone": nilSI.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"cloneId": "s1",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- Clone", actual)
}
