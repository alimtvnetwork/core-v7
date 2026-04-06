package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// SessionInfo
// ═══════════════════════════════════════════

func Test_Cov7_SessionInfo_Empty(t *testing.T) {
	si := corepayload.SessionInfo{}
	var nilSI *corepayload.SessionInfo
	actual := args.Map{
		"isEmpty":     si.IsEmpty(),
		"isValid":     si.IsValid(),
		"hasUser":     si.HasUser(),
		"isUserEmpty": si.IsUserEmpty(),
		"isNameEmpty": si.IsUserNameEmpty(),
		"idInt":       si.IdentifierInteger(),
		"idUint":      si.IdentifierUnsignedInteger(),
		"nilEmpty":    nilSI.IsEmpty(),
	}
	expected := args.Map{
		"isEmpty": true, "isValid": false,
		"hasUser": false, "isUserEmpty": true,
		"isNameEmpty": true, "idInt": -1, "idUint": uint(0),
		"nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns empty -- empty", actual)
}

func Test_Cov7_SessionInfo_Valid(t *testing.T) {
	u := &corepayload.User{Name: "admin"}
	si := corepayload.SessionInfo{
		Id: "42", User: u, SessionPath: "/path",
	}
	actual := args.Map{
		"isEmpty":    si.IsEmpty(),
		"isValid":    si.IsValid(),
		"hasUser":    si.HasUser(),
		"isNameEq":   si.IsUsernameEqual("admin"),
		"isNameNeq":  si.IsUsernameEqual("other"),
		"idInt":      si.IdentifierInteger(),
		"idUint":     si.IdentifierUnsignedInteger(),
	}
	expected := args.Map{
		"isEmpty": false, "isValid": true,
		"hasUser": true, "isNameEq": true, "isNameNeq": false,
		"idInt": 42, "idUint": uint(42),
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns non-empty -- valid", actual)
}

func Test_Cov7_SessionInfo_Clone(t *testing.T) {
	si := corepayload.SessionInfo{
		Id: "1", User: &corepayload.User{Name: "u"}, SessionPath: "/p",
	}
	cloned := si.Clone()
	ptr := si.Ptr()
	clonedPtr := si.ClonePtr()
	var nilSI *corepayload.SessionInfo
	nilClone := nilSI.ClonePtr()
	actual := args.Map{
		"clonedId":     cloned.Id,
		"ptrNotNil":    ptr != nil,
		"clonePtrNN":   clonedPtr != nil,
		"nilCloneNil":  nilClone == nil,
	}
	expected := args.Map{
		"clonedId": "1", "ptrNotNil": true,
		"clonePtrNN": true, "nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- clone", actual)
}

// ═══════════════════════════════════════════
// AuthInfo
// ═══════════════════════════════════════════

func Test_Cov7_AuthInfo_Empty(t *testing.T) {
	ai := corepayload.AuthInfo{}
	var nilAI *corepayload.AuthInfo
	actual := args.Map{
		"isEmpty":      ai.IsEmpty(),
		"hasAny":       ai.HasAnyItem(),
		"isValid":      ai.IsValid(),
		"isActionE":    ai.IsActionTypeEmpty(),
		"isResourceE":  ai.IsResourceNameEmpty(),
		"hasAction":    ai.HasActionType(),
		"hasResource":  ai.HasResourceName(),
		"isUserInfoE":  ai.IsUserInfoEmpty(),
		"isSessionE":   ai.IsSessionInfoEmpty(),
		"hasUserInfo":  ai.HasUserInfo(),
		"hasSession":   ai.HasSessionInfo(),
		"idInt":        ai.IdentifierInteger(),
		"idUint":       ai.IdentifierUnsignedInteger(),
		"nilEmpty":     nilAI.IsEmpty(),
	}
	expected := args.Map{
		"isEmpty": true, "hasAny": false, "isValid": false,
		"isActionE": true, "isResourceE": true,
		"hasAction": false, "hasResource": false,
		"isUserInfoE": true, "isSessionE": true,
		"hasUserInfo": false, "hasSession": false,
		"idInt": -1, "idUint": uint(0), "nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns empty -- empty", actual)
}

func Test_Cov7_AuthInfo_Setters(t *testing.T) {
	ai := &corepayload.AuthInfo{}
	ai.SetActionType("action").
		SetResourceName("resource").
		SetIdentifier("42")
	ui := &corepayload.UserInfo{}
	ai.SetUserInfo(ui)
	si := &corepayload.SessionInfo{Id: "1"}
	ai.SetSessionInfo(si)
	actual := args.Map{
		"action":   ai.ActionType,
		"resource": ai.ResourceName,
		"id":       ai.Identifier,
		"hasUI":    ai.HasUserInfo(),
		"hasSI":    ai.HasSessionInfo(),
		"idInt":    ai.IdentifierInteger(),
	}
	expected := args.Map{
		"action": "action", "resource": "resource", "id": "42",
		"hasUI": false, "hasSI": true, "idInt": 42,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- setters", actual)
}

func Test_Cov7_AuthInfo_NilSetters(t *testing.T) {
	var nilAI *corepayload.AuthInfo
	r1 := nilAI.SetActionType("a")
	r2 := nilAI.SetResourceName("r")
	r3 := nilAI.SetIdentifier("id")
	r4 := nilAI.SetUserInfo(nil)
	r5 := nilAI.SetSessionInfo(nil)
	r6 := nilAI.SetUser(&corepayload.User{Name: "u"})
	r7 := nilAI.SetSystemUser(&corepayload.User{Name: "s"})
	r8 := nilAI.SetUserSystemUser(&corepayload.User{Name: "u"}, &corepayload.User{Name: "s"})
	actual := args.Map{
		"r1NN": r1 != nil, "r2NN": r2 != nil, "r3NN": r3 != nil,
		"r4NN": r4 != nil, "r5NN": r5 != nil, "r6NN": r6 != nil,
		"r7NN": r7 != nil, "r8NN": r8 != nil,
	}
	expected := args.Map{
		"r1NN": true, "r2NN": true, "r3NN": true,
		"r4NN": true, "r5NN": true, "r6NN": true,
		"r7NN": true, "r8NN": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns nil -- nil setters", actual)
}

func Test_Cov7_AuthInfo_Clone(t *testing.T) {
	ai := corepayload.AuthInfo{
		Identifier: "1", ActionType: "a", ResourceName: "r",
	}
	cloned := ai.Clone()
	ptr := ai.Ptr()
	clonePtr := ai.ClonePtr()
	var nilAI *corepayload.AuthInfo
	nilClone := nilAI.ClonePtr()
	actual := args.Map{
		"clonedId":    cloned.Identifier,
		"ptrNN":       ptr != nil,
		"clonePtrNN":  clonePtr != nil,
		"nilCloneNil": nilClone == nil,
	}
	expected := args.Map{
		"clonedId": "1", "ptrNN": true,
		"clonePtrNN": true, "nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- clone", actual)
}

func Test_Cov7_AuthInfo_Json(t *testing.T) {
	ai := corepayload.AuthInfo{ActionType: "test"}
	j := ai.Json()
	jp := ai.JsonPtr()
	str := ai.String()
	pretty := ai.PrettyJsonString()
	actual := args.Map{
		"jsonLen":  j.Length() > 0,
		"jpNN":     jp != nil,
		"strNE":    str != "",
		"prettyNE": pretty != "",
	}
	expected := args.Map{
		"jsonLen": true, "jpNN": true, "strNE": true, "prettyNE": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- json", actual)
}

// ═══════════════════════════════════════════
// UserInfo
// ═══════════════════════════════════════════

func Test_Cov7_UserInfo_Empty(t *testing.T) {
	ui := corepayload.UserInfo{}
	var nilUI *corepayload.UserInfo
	actual := args.Map{
		"isEmpty":    ui.IsEmpty(),
		"hasUser":    ui.HasUser(),
		"hasSysUser": ui.HasSystemUser(),
		"isUserE":    ui.IsUserEmpty(),
		"isSysE":     ui.IsSystemUserEmpty(),
		"nilEmpty":   nilUI.IsEmpty(),
		"nilUserE":   nilUI.IsUserEmpty(),
		"nilSysE":    nilUI.IsSystemUserEmpty(),
	}
	expected := args.Map{
		"isEmpty": true, "hasUser": false, "hasSysUser": false,
		"isUserE": true, "isSysE": true,
		"nilEmpty": true, "nilUserE": true, "nilSysE": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns empty -- empty", actual)
}

func Test_Cov7_UserInfo_Setters(t *testing.T) {
	ui := &corepayload.UserInfo{}
	u := &corepayload.User{Name: "admin"}
	su := &corepayload.User{Name: "system"}
	ui.SetUser(u).SetSystemUser(su)
	actual := args.Map{
		"hasUser":    ui.HasUser(),
		"hasSysUser": ui.HasSystemUser(),
	}
	expected := args.Map{"hasUser": true, "hasSysUser": true}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- setters", actual)
}

func Test_Cov7_UserInfo_NilSetters(t *testing.T) {
	var nilUI *corepayload.UserInfo
	r1 := nilUI.SetUser(&corepayload.User{Name: "u"})
	r2 := nilUI.SetSystemUser(&corepayload.User{Name: "s"})
	r3 := nilUI.SetUserSystemUser(&corepayload.User{Name: "u"}, &corepayload.User{Name: "s"})
	actual := args.Map{
		"r1NN": r1 != nil, "r2NN": r2 != nil, "r3NN": r3 != nil,
	}
	expected := args.Map{"r1NN": true, "r2NN": true, "r3NN": true}
	expected.ShouldBeEqual(t, 0, "UserInfo returns nil -- nil setters", actual)
}

func Test_Cov7_UserInfo_Clone(t *testing.T) {
	ui := corepayload.UserInfo{
		User:       &corepayload.User{Name: "u"},
		SystemUser: &corepayload.User{Name: "s"},
	}
	cloned := ui.Clone()
	ptr := ui.Ptr()
	clonePtr := ui.ClonePtr()
	nonPtr := ui.ToNonPtr()
	var nilUI *corepayload.UserInfo
	nilClone := nilUI.ClonePtr()
	nilNonPtr := nilUI.ToNonPtr()
	actual := args.Map{
		"clonedHasU":  cloned.HasUser(),
		"ptrNN":       ptr != nil,
		"clonePtrNN":  clonePtr != nil,
		"nonPtrHasU":  nonPtr.HasUser(),
		"nilCloneNil": nilClone == nil,
		"nilNonPtrE":  nilNonPtr.IsEmpty(),
	}
	expected := args.Map{
		"clonedHasU": true, "ptrNN": true,
		"clonePtrNN": true, "nonPtrHasU": true,
		"nilCloneNil": true, "nilNonPtrE": true,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- clone", actual)
}

// ═══════════════════════════════════════════
// payloadProperties via PayloadWrapper
// ═══════════════════════════════════════════

func Test_Cov7_PayloadProperties_Basic(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Identifier: "42", Name: "test",
		EntityType: "ent", CategoryName: "cat",
		HasManyRecords: false,
		Payloads:       []byte(`"data"`),
	}
	props := pw.PayloadProperties()
	actual := args.Map{
		"propsNN":  props != nil,
		"name":     props.Name(),
		"idStr":    props.IdString(),
		"idInt":    props.IdInteger(),
		"idUint":   props.IdUnsignedInteger(),
		"entity":   props.EntityType(),
		"category": props.Category(),
		"hasManyF": props.HasManyRecord(),
		"single":   props.HasSingleRecordOnly(),
	}
	expected := args.Map{
		"propsNN": true, "name": "test",
		"idStr": "42", "idInt": 42, "idUint": uint(42),
		"entity": "ent", "category": "cat",
		"hasManyF": false, "single": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- basic", actual)
}

func Test_Cov7_PayloadProperties_Setters(t *testing.T) {
	pw := &corepayload.PayloadWrapper{}
	props := pw.PayloadProperties()
	props.SetName("n")
	props.SetNameMust("nm")
	props.SetIdString("id1")
	props.SetIdStringMust("id2")
	props.SetCategory("c1")
	props.SetCategoryMust("c2")
	props.SetEntityType("e1")
	props.SetEntityTypeMust("e2")
	props.SetManyRecordFlag()
	actual := args.Map{
		"name":     pw.Name,
		"id":       pw.Identifier,
		"category": pw.CategoryName,
		"entity":   pw.EntityType,
		"hasMany":  pw.HasManyRecords,
	}
	expected := args.Map{
		"name": "nm", "id": "id2",
		"category": "c2", "entity": "e2", "hasMany": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- setters", actual)
}

func Test_Cov7_PayloadProperties_SingleRecord(t *testing.T) {
	pw := &corepayload.PayloadWrapper{HasManyRecords: true}
	props := pw.PayloadProperties()
	props.SetSingleRecordFlag()
	actual := args.Map{"hasMany": pw.HasManyRecords}
	expected := args.Map{"hasMany": false}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- single record", actual)
}

func Test_Cov7_PayloadProperties_DynPayloads(t *testing.T) {
	pw := &corepayload.PayloadWrapper{Payloads: []byte(`"hello"`)}
	props := pw.PayloadProperties()
	dyn := props.DynamicPayloads()
	err := props.SetDynamicPayloads([]byte(`"world"`))
	actual := args.Map{
		"dynLen": len(dyn) > 0,
		"errNil": err == nil,
		"newDyn": len(props.DynamicPayloads()) > 0,
	}
	expected := args.Map{"dynLen": true, "errNil": true, "newDyn": true}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- dyn payloads", actual)
}

func Test_Cov7_PayloadProperties_All(t *testing.T) {
	pw := &corepayload.PayloadWrapper{
		Identifier: "1", Name: "n",
		EntityType: "e", CategoryName: "c",
		Payloads: []byte(`"d"`),
	}
	props := pw.PayloadProperties()
	id, name, entity, category, dynPayloads := props.All()
	id2, name2, entity2, category2, dynPayloads2 := props.AllSafe()
	actual := args.Map{
		"id": id, "name": name, "entity": entity, "category": category,
		"dynLen": len(dynPayloads) > 0,
		"id2": id2, "name2": name2, "entity2": entity2, "category2": category2,
		"dynLen2": len(dynPayloads2) > 0,
	}
	expected := args.Map{
		"id": "1", "name": "n", "entity": "e", "category": "c",
		"dynLen": true,
		"id2": "1", "name2": "n", "entity2": "e", "category2": "c",
		"dynLen2": true,
	}
	expected.ShouldBeEqual(t, 0, "payloadProperties returns correct value -- all", actual)
}

// ═══════════════════════════════════════════
// PayloadsCollection — add variations
// ═══════════════════════════════════════════

func Test_Cov7_PayloadsCollection_Adds(t *testing.T) {
	pc := &corepayload.PayloadsCollection{}
	pw1 := corepayload.PayloadWrapper{Name: "a"}
	pw2 := corepayload.PayloadWrapper{Name: "b"}
	pc.Add(pw1)
	pc.Adds(pw2)
	pc.Adds()
	pc.AddsPtr(&pw1)
	pc.AddsPtr()
	actual := args.Map{"len": len(pc.Items)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- adds", actual)
}

func Test_Cov7_PayloadsCollection_AddsOptions(t *testing.T) {
	pc := &corepayload.PayloadsCollection{}
	pw1 := corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)}
	pw2 := corepayload.PayloadWrapper{Name: ""} // empty = has issues
	pc.AddsOptions(true, pw1, pw2)
	pc.AddsOptions(false, pw1, pw2)
	actual := args.Map{"len": len(pc.Items)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- adds options", actual)
}

func Test_Cov7_PayloadsCollection_AddsPtrOptions(t *testing.T) {
	pc := &corepayload.PayloadsCollection{}
	pw1 := &corepayload.PayloadWrapper{Name: "a", Payloads: []byte(`"x"`)}
	pw2 := &corepayload.PayloadWrapper{Name: ""}
	pc.AddsPtrOptions(true, pw1, pw2)
	pc.AddsPtrOptions(false, pw1, pw2)
	actual := args.Map{"len": len(pc.Items)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "PayloadsCollection returns correct value -- adds ptr options", actual)
}
