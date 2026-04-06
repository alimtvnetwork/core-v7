package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// PayloadWrapper — core methods
// ═══════════════════════════════════════════

func Test_Cov6_PayloadWrapper_Basic(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name: "test", Identifier: "id-1", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
		HasManyRecords: false, Payloads: []byte(`"hello"`),
	}
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
	expected := args.Map{
		"name": "test", "entity": "entity", "category": "cat",
		"taskType": "task", "idStr": "id-1", "hasAny": true,
		"payloadDyn": true, "dynPayloads": true, "payloadsStr": true, "value": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_Cov6_PayloadWrapper_IsChecks(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Name: "test", Identifier: "id", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
	}
	actual := args.Map{
		"isName":     pw.IsName("test"),
		"isNotName":  pw.IsName("other"),
		"isId":       pw.IsIdentifier("id"),
		"isTask":     pw.IsTaskTypeName("task"),
		"isEntity":   pw.IsEntityType("entity"),
		"isCat":      pw.IsCategory("cat"),
	}
	expected := args.Map{
		"isName": true, "isNotName": false, "isId": true,
		"isTask": true, "isEntity": true, "isCat": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- IsChecks", actual)
}

func Test_Cov6_PayloadWrapper_IsEqual(t *testing.T) {
	pw1 := &corepayload.PayloadWrapper{Name: "test", Identifier: "id", Payloads: []byte("p")}
	pw2 := &corepayload.PayloadWrapper{Name: "test", Identifier: "id", Payloads: []byte("p")}
	pw3 := &corepayload.PayloadWrapper{Name: "other"}
	var nilPW *corepayload.PayloadWrapper
	actual := args.Map{
		"equal":     pw1.IsEqual(pw2),
		"notEqual":  pw1.IsEqual(pw3),
		"samePtr":   pw1.IsEqual(pw1),
		"nilBoth":   nilPW.IsEqual(nil),
		"nilLeft":   nilPW.IsEqual(pw1),
		"payEq":     pw1.IsPayloadsEqual([]byte("p")),
		"payNotEq":  pw1.IsPayloadsEqual([]byte("x")),
	}
	expected := args.Map{
		"equal": true, "notEqual": false, "samePtr": true,
		"nilBoth": true, "nilLeft": false, "payEq": true, "payNotEq": false,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- IsEqual", actual)
}

func Test_Cov6_PayloadWrapper_JSON(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Name: "test", Payloads: []byte(`"hello"`)}
	jsonStr := pw.JsonString()
	prettyStr := pw.PrettyJsonString()
	str := pw.String()
	b, err := pw.Serialize()
	actual := args.Map{
		"jsonNotEmpty":   jsonStr != "",
		"prettyNotEmpty": prettyStr != "",
		"strNotEmpty":    str != "",
		"bLen":           len(b) > 0,
		"noErr":          err == nil,
	}
	expected := args.Map{
		"jsonNotEmpty": true, "prettyNotEmpty": true, "strNotEmpty": true,
		"bLen": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- JSON", actual)
}

func Test_Cov6_PayloadWrapper_SetDynamicPayloads(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	err := pw.SetDynamicPayloads([]byte("new"))
	var nilPW *corepayload.PayloadWrapper
	nilErr := nilPW.SetDynamicPayloads([]byte("x"))
	actual := args.Map{
		"noErr":   err == nil,
		"payload": string(pw.Payloads),
		"nilErr":  nilErr != nil,
	}
	expected := args.Map{"noErr": true, "payload": "new", "nilErr": true}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- SetDynamicPayloads", actual)
}

func Test_Cov6_PayloadWrapper_All(t *testing.T) {
	pw := corepayload.PayloadWrapper{
		Name: "n", Identifier: "id", EntityType: "e", CategoryName: "c", Payloads: []byte("p"),
	}
	id, name, entity, cat, dynP := pw.All()
	actual := args.Map{"id": id, "name": name, "entity": entity, "cat": cat, "dynP": string(dynP)}
	expected := args.Map{"id": "id", "name": "n", "entity": "e", "cat": "c", "dynP": "p"}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- All", actual)
}

func Test_Cov6_PayloadWrapper_AllSafe_Nil(t *testing.T) {
	var pw *corepayload.PayloadWrapper
	id, name, entity, cat, dynP := pw.AllSafe()
	actual := args.Map{"id": id, "name": name, "entity": entity, "cat": cat, "dynP": string(dynP)}
	expected := args.Map{"id": "", "name": "", "entity": "", "cat": "", "dynP": ""}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns nil -- AllSafe nil", actual)
}

// ═══════════════════════════════════════════
// User — comprehensive
// ═══════════════════════════════════════════

func Test_Cov6_User_Comprehensive(t *testing.T) {
	u := &corepayload.User{
		Name: "Alice", Type: "admin", Identifier: "123",
		AuthToken: "token", PasswordHash: "hash", IsSystemUser: false,
	}
	var nilU *corepayload.User
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

func Test_Cov6_User_Clone(t *testing.T) {
	u := &corepayload.User{Name: "Alice", Type: "admin"}
	cloned := u.Clone()
	clonedPtr := u.ClonePtr()
	var nilU *corepayload.User
	actual := args.Map{
		"cloneName":   cloned.Name,
		"cpName":      clonedPtr.Name,
		"nilClonePtr": nilU.ClonePtr() == nil,
	}
	expected := args.Map{"cloneName": "Alice", "cpName": "Alice", "nilClonePtr": true}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- Clone", actual)
}

func Test_Cov6_User_JSON(t *testing.T) {
	u := &corepayload.User{Name: "Alice"}
	j := u.Json()
	jp := u.JsonPtr()
	b, err := u.Serialize()
	actual := args.Map{
		"jHas": j.HasBytes(), "jpNN": jp != nil,
		"bLen": len(b) > 0, "noErr": err == nil,
	}
	expected := args.Map{"jHas": true, "jpNN": true, "bLen": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "User returns correct value -- JSON", actual)
}

// ═══════════════════════════════════════════
// AuthInfo — comprehensive
// ═══════════════════════════════════════════

func Test_Cov6_AuthInfo_Basic(t *testing.T) {
	ai := &corepayload.AuthInfo{
		Identifier: "123", ActionType: "create", ResourceName: "/api",
	}
	var nilAI *corepayload.AuthInfo
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
	expected := args.Map{
		"isEmpty": false, "hasAny": true, "isValid": true,
		"isActionEmpty": false, "isResEmpty": false,
		"hasAction": true, "hasResource": true,
		"idInt": 123, "idUint": uint(123),
		"strNN": true, "nilEmpty": true, "nilActionEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- basic", actual)
}

func Test_Cov6_AuthInfo_Setters(t *testing.T) {
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("create")
	ai.SetResourceName("/api")
	ai.SetIdentifier("id-1")
	actual := args.Map{
		"action": ai.ActionType, "resource": ai.ResourceName, "id": ai.Identifier,
	}
	expected := args.Map{"action": "create", "resource": "/api", "id": "id-1"}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- setters", actual)
}

func Test_Cov6_AuthInfo_Clone(t *testing.T) {
	ai := &corepayload.AuthInfo{ActionType: "create"}
	cloned := ai.Clone()
	clonedPtr := ai.ClonePtr()
	var nilAI *corepayload.AuthInfo
	actual := args.Map{
		"cloneAction": cloned.ActionType,
		"cpAction":    clonedPtr.ActionType,
		"nilClone":    nilAI.ClonePtr() == nil,
	}
	expected := args.Map{"cloneAction": "create", "cpAction": "create", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- Clone", actual)
}

func Test_Cov6_AuthInfo_JSON(t *testing.T) {
	ai := corepayload.AuthInfo{ActionType: "create"}
	j := ai.Json()
	jp := ai.JsonPtr()
	actual := args.Map{"jHas": j.HasBytes(), "jpNN": jp != nil}
	expected := args.Map{"jHas": true, "jpNN": true}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- JSON", actual)
}

// ═══════════════════════════════════════════
// PayloadsCollection
// ═══════════════════════════════════════════

func Test_Cov6_PayloadsCollection_Add(t *testing.T) {
	pc := &corepayload.PayloadsCollection{}
	pw := corepayload.PayloadWrapper{Name: "test"}
	pc.Add(pw)
	pc.Adds(corepayload.PayloadWrapper{Name: "t2"}, corepayload.PayloadWrapper{Name: "t3"})
	actual := args.Map{"len": len(pc.Items)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- Add", actual)
}

func Test_Cov6_PayloadsCollection_AddsPtr(t *testing.T) {
	pc := &corepayload.PayloadsCollection{}
	pw := &corepayload.PayloadWrapper{Name: "test"}
	pc.AddsPtr(pw)
	pc.AddsPtr()
	actual := args.Map{"len": len(pc.Items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- AddsPtr", actual)
}

// ═══════════════════════════════════════════
// BytesCreateInstruction
// ═══════════════════════════════════════════

func Test_Cov6_BytesCreateInstruction(t *testing.T) {
	bci := corepayload.BytesCreateInstruction{
		Name: "test", Identifier: "id", TaskTypeName: "task",
		EntityType: "entity", CategoryName: "cat",
		HasManyRecords: true, Payloads: []byte("payload"),
	}
	actual := args.Map{
		"name": bci.Name, "id": bci.Identifier, "task": bci.TaskTypeName,
		"entity": bci.EntityType, "cat": bci.CategoryName,
		"hasMany": bci.HasManyRecords, "payLen": len(bci.Payloads),
	}
	expected := args.Map{
		"name": "test", "id": "id", "task": "task",
		"entity": "entity", "cat": "cat", "hasMany": true, "payLen": 7,
	}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstruction returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// SessionInfo
// ═══════════════════════════════════════════

func Test_Cov6_SessionInfo_IsEmpty(t *testing.T) {
	si := &corepayload.SessionInfo{}
	var nilSI *corepayload.SessionInfo
	actual := args.Map{
		"isEmpty":  si.IsEmpty(),
		"nilEmpty": nilSI.IsEmpty(),
	}
	expected := args.Map{"isEmpty": true, "nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns empty -- IsEmpty", actual)
}

func Test_Cov6_SessionInfo_Clone(t *testing.T) {
	si := &corepayload.SessionInfo{Id: "s1"}
	cloned := si.ClonePtr()
	var nilSI *corepayload.SessionInfo
	actual := args.Map{
		"cloneId":  cloned.Id,
		"nilClone": nilSI.ClonePtr() == nil,
	}
	expected := args.Map{"cloneId": "s1", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- Clone", actual)
}
