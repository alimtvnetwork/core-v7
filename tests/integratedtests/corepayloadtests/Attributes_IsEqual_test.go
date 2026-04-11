package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Attributes extended ──

func Test_Attributes_IsEqual(t *testing.T) {
	// Arrange
	a1 := corepayload.New.Attributes.Empty()
	a2 := corepayload.New.Attributes.Empty()
	var nilA *corepayload.Attributes

	// Act
	actual := args.Map{
		"equal":     a1.IsEqual(a2),
		"self":      a1.IsEqual(a1),
		"nilNil":    nilA.IsEqual(nilA),
		"nilNonNil": nilA.IsEqual(a1),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"self": true,
		"nilNil": true,
		"nilNonNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- IsEqual", actual)
}

func Test_Attributes_Clone(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte(`{"a":1}`))
	cloned, err := a.Clone(false)
	deepCloned, deepErr := a.Clone(true)
	var nilA *corepayload.Attributes
	nilCloned, nilErr := nilA.ClonePtr(false)

	// Act
	actual := args.Map{
		"noErr":  err == nil, "len": cloned.Length() > 0,
		"deepNoErr": deepErr == nil, "deepLen": deepCloned.Length() > 0,
		"nilClone": nilCloned == nil, "nilErr": nilErr == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true, "len": true, "deepNoErr": true, "deepLen": true,
		"nilClone": true, "nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- Clone", actual)
}

func Test_Attributes_Getters(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	var nilA *corepayload.Attributes

	// Act
	actual := args.Map{
		"isNull":           nilA.IsNull(),
		"hasSafe":          a.HasSafeItems(),
		"hasStringKey":     a.HasStringKey("x"),
		"hasAnyKey":        a.HasAnyKey("x"),
		"payloadsLen":      len(a.Payloads()),
		"payloadsStr":      a.PayloadsString(),
		"anyMap":           len(a.AnyKeyValMap()),
		"hashmap":          len(a.Hashmap()),
		"compiledErr":      a.CompiledError() == nil,
		"hasIssues":        a.HasIssuesOrEmpty(),
		"isSafeValid":      a.IsSafeValid(),
		"hasAnyItem":       a.HasAnyItem(),
		"count":            a.Count(),
		"capacity":         a.Capacity(),
		"length":           a.Length(),
		"nilLength":        nilA.Length(),
		"hasPaging":        a.HasPagingInfo(),
		"hasKV":            a.HasKeyValuePairs(),
		"hasFromTo":        a.HasFromTo(),
		"isValid":          a.IsValid(),
		"isInvalid":        a.IsInvalid(),
		"hasError":         a.HasError(),
		"isEmptyError":     a.IsEmptyError(),
		"dynamicLen":       a.DynamicBytesLength(),
		"stringKVLen":      a.StringKeyValuePairsLength(),
		"anyKVLen":         a.AnyKeyValuePairsLength(),
		"nilDynLen":        nilA.DynamicBytesLength(),
		"nilStrKVLen":      nilA.StringKeyValuePairsLength(),
		"nilAnyKVLen":      nilA.AnyKeyValuePairsLength(),
		"hasStringKVP":     a.HasStringKeyValuePairs(),
		"hasAnyKVP":        a.HasAnyKeyValuePairs(),
		"hasDynamic":       a.HasDynamicPayloads(),
		"isPagingEmpty":    a.IsPagingInfoEmpty(),
		"isKVEmpty":        a.IsKeyValuePairsEmpty(),
		"isAnyKVEmpty":     a.IsAnyKeyValuePairsEmpty(),
		"isUserInfoEmpty":  a.IsUserInfoEmpty(),
		"isAuthInfoEmpty":  a.IsAuthInfoEmpty(),
		"isSessionInfoEmpty": a.IsSessionInfoEmpty(),
		"hasUserInfo":      a.HasUserInfo(),
		"hasAuthInfo":      a.HasAuthInfo(),
		"hasSessionInfo":   a.HasSessionInfo(),
	}

	// Assert
	expected := args.Map{
		"isNull": true, "hasSafe": false, "hasStringKey": false, "hasAnyKey": false,
		"payloadsLen": 0, "payloadsStr": "", "anyMap": 0, "hashmap": 0,
		"compiledErr": true, "hasIssues": true, "isSafeValid": false,
		"hasAnyItem": false, "count": 0, "capacity": 0, "length": 0, "nilLength": 0,
		"hasPaging": false, "hasKV": false, "hasFromTo": false,
		"isValid": true, "isInvalid": true, "hasError": false, "isEmptyError": true,
		"dynamicLen": 0, "stringKVLen": 0, "anyKVLen": 0,
		"nilDynLen": 0, "nilStrKVLen": 0, "nilAnyKVLen": 0,
		"hasStringKVP": false, "hasAnyKVP": false, "hasDynamic": false,
		"isPagingEmpty": true, "isKVEmpty": true, "isAnyKVEmpty": true,
		"isUserInfoEmpty": true, "isAuthInfoEmpty": true, "isSessionInfoEmpty": true,
		"hasUserInfo": false, "hasAuthInfo": false, "hasSessionInfo": false,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- Getters", actual)
}

func Test_Attributes_GetStringKeyValue_FromAttributesIsEqual(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	_, found := a.GetStringKeyValue("x")
	var nilA *corepayload.Attributes
	_, nilFound := nilA.GetStringKeyValue("x")

	// Act
	actual := args.Map{
		"found": found,
		"nilFound": nilFound,
	}

	// Assert
	expected := args.Map{
		"found": false,
		"nilFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- GetStringKeyValue", actual)
}

func Test_Attributes_IsErrorEqual(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"equalNil":     a.IsErrorEqual(nil),
		"differentNil": a.IsErrorDifferent(nil),
	}

	// Assert
	expected := args.Map{
		"equalNil": true,
		"differentNil": false,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns error -- IsErrorEqual", actual)
}

// ── PagingInfo ──

func Test_PagingInfo(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}
	var nilP *corepayload.PagingInfo

	// Act
	actual := args.Map{
		"isEmpty":     p.IsEmpty(),
		"isEqual":     p.IsEqual(p),
		"nilNil":      nilP.IsEqual(nilP),
		"nilNonNil":   nilP.IsEqual(p),
		"hasTotal":    p.HasTotalPages(),
		"hasCurrent":  p.HasCurrentPageIndex(),
		"hasPerPage":  p.HasPerPageItems(),
		"hasTotalI":   p.HasTotalItems(),
		"invalidTotal": p.IsInvalidTotalPages(),
		"invalidCurr":  p.IsInvalidCurrentPageIndex(),
		"invalidPer":   p.IsInvalidPerPageItems(),
		"invalidTotalI": p.IsInvalidTotalItems(),
		"cloneTotalP":  p.Clone().TotalPages,
		"clonePtrNil":  nilP.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "isEqual": true, "nilNil": true, "nilNonNil": false,
		"hasTotal": true, "hasCurrent": true, "hasPerPage": true, "hasTotalI": true,
		"invalidTotal": false, "invalidCurr": false, "invalidPer": false, "invalidTotalI": false,
		"cloneTotalP": 5, "clonePtrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- with args", actual)
}

// ── SessionInfo ──

func Test_SessionInfo_FromAttributesIsEqual(t *testing.T) {
	// Arrange
	u := corepayload.New.User.UsingName("alice")
	s := &corepayload.SessionInfo{Id: "42", User: u, SessionPath: "/sess"}
	var nilS *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"isEmpty":  s.IsEmpty(),
		"isValid":  s.IsValid(),
		"hasUser":  s.HasUser(),
		"isUserE":  s.IsUserEmpty(),
		"isNameE":  s.IsUserNameEmpty(),
		"nameEq":   s.IsUsernameEqual("alice"),
		"idInt":    s.IdentifierInteger(),
		"idUint":   s.IdentifierUnsignedInteger(),
		"cloneId":  s.ClonePtr().Id,
		"nilEmpty": nilS.IsEmpty(),
		"nilClone": nilS.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "isValid": true, "hasUser": true,
		"isUserE": false, "isNameE": false, "nameEq": true,
		"idInt": 42, "idUint": uint(42), "cloneId": "42",
		"nilEmpty": true, "nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- with args", actual)
}

// ── AuthInfo ──

func Test_AuthInfo(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ActionType: "login", ResourceName: "/api", Identifier: "42"}
	var nilA *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"isEmpty":      a.IsEmpty(),
		"hasAnyItem":   a.HasAnyItem(),
		"isValid":      a.IsValid(),
		"hasAction":    a.HasActionType(),
		"hasResource":  a.HasResourceName(),
		"isActionE":    a.IsActionTypeEmpty(),
		"isResourceE":  a.IsResourceNameEmpty(),
		"isUserInfoE":  a.IsUserInfoEmpty(),
		"isSessionE":   a.IsSessionInfoEmpty(),
		"hasUserInfo":  a.HasUserInfo(),
		"hasSession":   a.HasSessionInfo(),
		"idInt":        a.IdentifierInteger(),
		"idUint":       a.IdentifierUnsignedInteger(),
		"str":          a.String() != "",
		"pretty":       a.PrettyJsonString() != "",
		"nilEmpty":     nilA.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": false, "hasAnyItem": true, "isValid": true,
		"hasAction": true, "hasResource": true,
		"isActionE": false, "isResourceE": false,
		"isUserInfoE": true, "isSessionE": true,
		"hasUserInfo": false, "hasSession": false,
		"idInt": 42, "idUint": uint(42),
		"str": true, "pretty": true, "nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- with args", actual)
}

func Test_AuthInfo_Setters(t *testing.T) {
	// Arrange
	var nilA *corepayload.AuthInfo
	fromNilUser := nilA.SetUserInfo(&corepayload.UserInfo{})
	fromNilAction := nilA.SetActionType("login")
	fromNilResource := nilA.SetResourceName("/api")
	fromNilId := nilA.SetIdentifier("42")
	fromNilSession := nilA.SetSessionInfo(&corepayload.SessionInfo{})
	fromNilSysUser := nilA.SetUserSystemUser(nil, nil)
	fromNilSetUser := nilA.SetUser(corepayload.New.User.UsingName("a"))
	fromNilSetSysUser := nilA.SetSystemUser(corepayload.New.User.UsingName("s"))

	// Act
	actual := args.Map{
		"userNotNil":    fromNilUser != nil,
		"actionNotNil":  fromNilAction != nil,
		"resourceNotNil": fromNilResource != nil,
		"idNotNil":      fromNilId != nil,
		"sessionNotNil": fromNilSession != nil,
		"sysUserNotNil": fromNilSysUser != nil,
		"setUserNotNil": fromNilSetUser != nil,
		"setSysNotNil":  fromNilSetSysUser != nil,
	}

	// Assert
	expected := args.Map{
		"userNotNil": true, "actionNotNil": true, "resourceNotNil": true,
		"idNotNil": true, "sessionNotNil": true, "sysUserNotNil": true,
		"setUserNotNil": true, "setSysNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- Setters", actual)
}

func Test_AuthInfo_Clone(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ActionType: "login"}
	cloned := a.ClonePtr()
	var nilA *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"action":   cloned.ActionType,
		"nilClone": nilA.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"action": "login",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- Clone", actual)
}

// ── TypedPayloadWrapper ──

func Test_TypedPayloadWrapper(t *testing.T) {
	// Arrange
	type Data struct{ Name string }
	typed, err := corepayload.NewTypedPayloadWrapperFrom[Data](
		"test", "id1", "entity", Data{Name: "alice"},
	)

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"name":     typed.Name(),
		"id":       typed.Identifier(),
		"idStr":    typed.IdString(),
		"entity":   typed.EntityType(),
		"category": typed.CategoryName(),
		"task":     typed.TaskTypeName(),
		"hasMany":  typed.HasManyRecords(),
		"hasSingle": typed.HasSingleRecord(),
		"isParsed": typed.IsParsed(),
		"data":     typed.TypedData().Name,
		"dataAlias": typed.Data().Name,
		"hasItems": typed.HasItems(),
		"isEmpty":  typed.IsEmpty(),
		"hasSafe":  typed.HasSafeItems(),
		"hasError": typed.HasError(),
		"err":      typed.Error() == nil,
		"str":      typed.String() != "",
		"pretty":   typed.PrettyJsonString() != "",
		"jsonStr":  typed.JsonString() != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true, "name": "test", "id": "id1", "idStr": "id1",
		"entity": "entity", "category": "", "task": "",
		"hasMany": false, "hasSingle": true, "isParsed": true,
		"data": "alice", "dataAlias": "alice",
		"hasItems": true, "isEmpty": false, "hasSafe": true,
		"hasError": false, "err": true, "str": true, "pretty": true, "jsonStr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- with args", actual)
}

func Test_TypedPayloadWrapper_Nil(t *testing.T) {
	// Arrange
	var typed *corepayload.TypedPayloadWrapper[string]

	// Act
	actual := args.Map{
		"name":     typed.Name(),
		"id":       typed.Identifier(),
		"idStr":    typed.IdString(),
		"entity":   typed.EntityType(),
		"category": typed.CategoryName(),
		"task":     typed.TaskTypeName(),
		"hasMany":  typed.HasManyRecords(),
		"isEmpty":  typed.IsEmpty(),
		"hasError": typed.HasError(),
		"err":      typed.Error() == nil,
		"str":      typed.String(),
		"pretty":   typed.PrettyJsonString(),
		"jsonStr":  typed.JsonString(),
		"isParsed": typed.IsParsed(),
		"attrs":    typed.Attributes() == nil,
	}

	// Assert
	expected := args.Map{
		"name": "", "id": "", "idStr": "", "entity": "", "category": "", "task": "",
		"hasMany": false, "isEmpty": true, "hasError": false, "err": true,
		"str": "", "pretty": "", "jsonStr": "", "isParsed": false, "attrs": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns nil -- Nil", actual)
}

func Test_TypedPayloadWrapper_JSON(t *testing.T) {
	// Arrange
	type Data struct{ Name string }
	typed, _ := corepayload.NewTypedPayloadWrapperFrom[Data](
		"test", "id1", "entity", Data{Name: "alice"},
	)
	bytes, err := typed.MarshalJSON()
	var typed2 corepayload.TypedPayloadWrapper[Data]
	err2 := typed2.UnmarshalJSON(bytes)

	// Act
	actual := args.Map{
		"noErr": err == nil, "noErr2": err2 == nil,
		"name": typed2.Data().Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"noErr2": true,
		"name": "alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- JSON", actual)
}

func Test_TypedPayloadWrapper_Serialize(t *testing.T) {
	// Arrange
	type Data struct{ Name string }
	typed, _ := corepayload.NewTypedPayloadWrapperFrom[Data](
		"test", "id1", "entity", Data{Name: "alice"},
	)
	bytes, err := typed.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapper returns correct value -- Serialize", actual)
}

// ── PayloadCreateInstructionTypeStringer ──

func Test_PayloadCreateInstructionTypeStringer(t *testing.T) {
	// Arrange
	type myStringer struct{ val string }
	ms := corepayload.PayloadCreateInstructionTypeStringer{
		Name:                 "test",
		Identifier:           "id1",
		TaskTypeNameStringer: corepayload.New.User.UsingName("task"),
		CategoryNameStringer: corepayload.New.User.UsingName("cat"),
	}
	result := ms.PayloadCreateInstruction()

	// Act
	actual := args.Map{"name": result.Name}

	// Assert
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ── Empty creator ──

func Test_EmptyCreator(t *testing.T) {
	// Arrange
	attrs := corepayload.Empty.Attributes()
	attrsDef := corepayload.Empty.AttributesDefaults()
	pw := corepayload.Empty.PayloadWrapper()
	pc := corepayload.Empty.PayloadsCollection()

	// Act
	actual := args.Map{
		"attrsNotNil":  attrs != nil,
		"attrsDefHasKV": attrsDef.HasKeyValuePairs(),
		"pwNotNil":     pw != nil,
		"pcNotNil":     pc != nil,
	}

	// Assert
	expected := args.Map{
		"attrsNotNil": true, "attrsDefHasKV": false,
		"pwNotNil": true, "pcNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "EmptyCreator returns empty -- with args", actual)
}
