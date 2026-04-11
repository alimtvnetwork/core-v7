package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AttributesGetters — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Attributes_IsNull_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	a2 := &corepayload.Attributes{}

	// Act
	actual := args.Map{
		"nil": a.IsNull(),
		"notNil": a2.IsNull(),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"notNil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNull returns correct value -- with args", actual)
}

func Test_Attributes_HasSafeItems_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.HasSafeItems()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems returns correct value -- with args", actual)
}

func Test_Attributes_HasStringKey_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")

	// Act
	actual := args.Map{
		"has": a.HasStringKey("k"),
		"no": a.HasStringKey("x"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"no": false,
	}
	expected.ShouldBeEqual(t, 0, "HasStringKey returns correct value -- with args", actual)
}

func Test_Attributes_HasStringKey_NoKV(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"val": a.HasStringKey("k")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasStringKey returns empty -- no kv", actual)
}

func Test_Attributes_HasAnyKey_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.AnyKeyValuePairs.Add("k", 42)

	// Act
	actual := args.Map{
		"has": a.HasAnyKey("k"),
		"no": a.HasAnyKey("x"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"no": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyKey returns correct value -- with args", actual)
}

func Test_Attributes_HasAnyKey_NoAKV(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"val": a.HasAnyKey("k")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasAnyKey returns empty -- no akv", actual)
}

func Test_Attributes_Payloads_Empty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"len": len(a.Payloads())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Payloads returns empty -- empty", actual)
}

func Test_Attributes_Payloads_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"len": len(a.Payloads())}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Payloads returns non-empty -- valid", actual)
}

func Test_Attributes_PayloadsString_Empty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.PayloadsString()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns empty -- empty", actual)
}

func Test_Attributes_PayloadsString_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("hello"))

	// Act
	actual := args.Map{"val": a.PayloadsString()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PayloadsString returns non-empty -- valid", actual)
}

func Test_Attributes_AnyKeyValMap_Empty(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	m := a.AnyKeyValMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyKeyValMap returns empty -- empty", actual)
}

func Test_Attributes_Hashmap_Empty(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	m := a.Hashmap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- empty", actual)
}

func Test_Attributes_CompiledError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"noErr": a.CompiledError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with args", actual)
}

func Test_Attributes_HasIssuesOrEmpty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"val": a.HasIssuesOrEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasIssuesOrEmpty returns empty -- empty", actual)
}

func Test_Attributes_IsSafeValid_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.IsSafeValid()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsSafeValid returns non-empty -- with args", actual)
}

func Test_Attributes_HasAnyItem_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.HasAnyItem()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- with args", actual)
}

func Test_Attributes_Count(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.Count()}

	// Assert
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Count returns correct value -- with args", actual)
}

func Test_Attributes_Capacity(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.Capacity()}

	// Assert
	expected := args.Map{"val": 4}
	expected.ShouldBeEqual(t, 0, "Capacity returns correct value -- with args", actual)
}

func Test_Attributes_Length_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.Length()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Attributes_HasPagingInfo_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{PagingInfo: &corepayload.PagingInfo{TotalPages: 5}}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"has": a.HasPagingInfo(),
		"nil": a2.HasPagingInfo(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasPagingInfo returns correct value -- with args", actual)
}

func Test_Attributes_HasKeyValuePairs_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"has": a.HasKeyValuePairs(),
		"nil": a2.HasKeyValuePairs(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasKeyValuePairs returns correct value -- with args", actual)
}

func Test_Attributes_HasFromTo_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"no": a.HasFromTo(),
		"nil": a2.HasFromTo(),
	}

	// Assert
	expected := args.Map{
		"no": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasFromTo returns correct value -- with args", actual)
}

func Test_Attributes_IsValid_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"valid": a.IsValid(),
		"nil": a2.IsValid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValid returns non-empty -- with args", actual)
}

func Test_Attributes_IsInvalid(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.IsInvalid()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid returns nil -- nil", actual)
}

func Test_Attributes_HasError(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"no": a.HasError(),
		"nil": a2.HasError(),
	}

	// Assert
	expected := args.Map{
		"no": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasError returns error -- with args", actual)
}

func Test_Attributes_Error_Empty(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"noErr": a.Error() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error returns empty -- empty", actual)
}

func Test_Attributes_IsEmptyError_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.IsEmptyError()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyError returns nil -- nil", actual)
}

func Test_Attributes_DynamicBytesLength_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.DynamicBytesLength()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "DynamicBytesLength returns nil -- nil", actual)
}

func Test_Attributes_StringKeyValuePairsLength_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.StringKeyValuePairsLength()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "StringKeyValuePairsLength returns nil -- nil", actual)
}

func Test_Attributes_AnyKeyValuePairsLength_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.AnyKeyValuePairsLength()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "AnyKeyValuePairsLength returns nil -- nil", actual)
}

func Test_Attributes_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.IsEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns nil -- nil", actual)
}

func Test_Attributes_HasItems(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.HasItems()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasItems returns correct value -- with args", actual)
}

func Test_Attributes_IsPagingInfoEmpty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsPagingInfoEmpty(),
		"nil": a2.IsPagingInfoEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsPagingInfoEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsKeyValuePairsEmpty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsKeyValuePairsEmpty(),
		"nil": a2.IsKeyValuePairsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsAnyKeyValuePairsEmpty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsAnyKeyValuePairsEmpty(),
		"nil": a2.IsAnyKeyValuePairsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyKeyValuePairsEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsUserInfoEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsUserInfoEmpty(),
		"nil": a2.IsUserInfoEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsUserInfoEmpty returns empty -- with args", actual)
}

func Test_Attributes_VirtualUser_Empty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"nil": a.VirtualUser() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "VirtualUser returns empty -- empty", actual)
}

func Test_Attributes_VirtualUser_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "u"}}}}

	// Act
	actual := args.Map{"name": a.VirtualUser().Name}

	// Assert
	expected := args.Map{"name": "u"}
	expected.ShouldBeEqual(t, 0, "VirtualUser returns non-empty -- valid", actual)
}

func Test_Attributes_SystemUser_Empty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"nil": a.SystemUser() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SystemUser returns empty -- empty", actual)
}

func Test_Attributes_SystemUser_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{SystemUser: &corepayload.User{Name: "su"}}}}

	// Act
	actual := args.Map{"name": a.SystemUser().Name}

	// Assert
	expected := args.Map{"name": "su"}
	expected.ShouldBeEqual(t, 0, "SystemUser returns non-empty -- valid", actual)
}

func Test_Attributes_SessionUser_Empty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"nil": a.SessionUser() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionUser returns empty -- empty", actual)
}

func Test_Attributes_SessionUser_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1", User: &corepayload.User{Name: "su"}}}}

	// Act
	actual := args.Map{"name": a.SessionUser().Name}

	// Assert
	expected := args.Map{"name": "su"}
	expected.ShouldBeEqual(t, 0, "SessionUser returns non-empty -- valid", actual)
}

func Test_Attributes_IsAuthInfoEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsAuthInfoEmpty(),
		"nil": a2.IsAuthInfoEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsAuthInfoEmpty returns empty -- with args", actual)
}

func Test_Attributes_IsSessionInfoEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	var a2 *corepayload.Attributes

	// Act
	actual := args.Map{
		"empty": a.IsSessionInfoEmpty(),
		"nil": a2.IsSessionInfoEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsSessionInfoEmpty returns empty -- with args", actual)
}

func Test_Attributes_HasUserInfo(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "u"}}}}

	// Act
	actual := args.Map{"val": a.HasUserInfo()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasUserInfo returns correct value -- with args", actual)
}

func Test_Attributes_HasAuthInfo(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ActionType: "login"}}

	// Act
	actual := args.Map{"val": a.HasAuthInfo()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAuthInfo returns correct value -- with args", actual)
}

func Test_Attributes_HasSessionInfo(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1"}}}

	// Act
	actual := args.Map{"val": a.HasSessionInfo()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasSessionInfo returns correct value -- with args", actual)
}

func Test_Attributes_SessionInfo_Empty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"nil": a.SessionInfo() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns empty -- empty", actual)
}

func Test_Attributes_SessionInfo_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1"}}}

	// Act
	actual := args.Map{"id": a.SessionInfo().Id}

	// Assert
	expected := args.Map{"id": "1"}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns non-empty -- valid", actual)
}

func Test_Attributes_AuthType_Empty_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"val": a.AuthType()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "AuthType returns empty -- empty", actual)
}

func Test_Attributes_AuthType_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ActionType: "login"}}

	// Act
	actual := args.Map{"val": a.AuthType()}

	// Assert
	expected := args.Map{"val": "login"}
	expected.ShouldBeEqual(t, 0, "AuthType returns non-empty -- valid", actual)
}

func Test_Attributes_ResourceName_Empty(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"val": a.ResourceName()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "ResourceName returns empty -- empty", actual)
}

func Test_Attributes_ResourceName_Valid(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{ResourceName: "res"}}

	// Act
	actual := args.Map{"val": a.ResourceName()}

	// Assert
	expected := args.Map{"val": "res"}
	expected.ShouldBeEqual(t, 0, "ResourceName returns non-empty -- valid", actual)
}

func Test_Attributes_HasStringKeyValuePairs_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")

	// Act
	actual := args.Map{"val": a.HasStringKeyValuePairs()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasStringKeyValuePairs returns correct value -- with args", actual)
}

func Test_Attributes_HasAnyKeyValuePairs_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.AnyKeyValuePairs.Add("k", 42)

	// Act
	actual := args.Map{"val": a.HasAnyKeyValuePairs()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyKeyValuePairs returns correct value -- with args", actual)
}

func Test_Attributes_HasDynamicPayloads_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"val": a.HasDynamicPayloads()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasDynamicPayloads returns correct value -- with args", actual)
}

func Test_Attributes_GetStringKeyValue_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	v, found := a.GetStringKeyValue("k")

	// Act
	actual := args.Map{
		"val": v,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": "",
		"found": false,
	}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns nil -- nil", actual)
}

func Test_Attributes_GetStringKeyValue_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.KeyValuePairs.AddOrUpdate("k", "v")
	v, found := a.GetStringKeyValue("k")

	// Act
	actual := args.Map{
		"val": v,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"found": true,
	}
	expected.ShouldBeEqual(t, 0, "GetStringKeyValue returns non-empty -- valid", actual)
}

func Test_Attributes_GetAnyKeyValue_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	_, found := a.GetAnyKeyValue("k")

	// Act
	actual := args.Map{"found": found}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "GetAnyKeyValue returns nil -- nil", actual)
}

func Test_Attributes_IsErrorDifferent(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"val": a.IsErrorDifferent(nil)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsErrorDifferent returns error -- with args", actual)
}

func Test_Attributes_IsErrorEqual_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"val": a.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual returns nil -- nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AttributesSetters — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Attributes_SetAuthInfo_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	result := a.SetAuthInfo(&corepayload.AuthInfo{ActionType: "x"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo returns nil -- nil", actual)
}

func Test_Attributes_SetAuthInfo_NonNil(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	result := a.SetAuthInfo(&corepayload.AuthInfo{ActionType: "x"})

	// Act
	actual := args.Map{"same": result == a}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetAuthInfo returns nil -- non-nil", actual)
}

func Test_Attributes_SetUserInfo_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	result := a.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo returns nil -- nil", actual)
}

func Test_Attributes_SetUserInfo_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{AuthInfo: &corepayload.AuthInfo{}}
	result := a.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{"same": result == a}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetUserInfo returns nil -- non-nil", actual)
}

func Test_Attributes_AddNewStringKeyValueOnly_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.AddNewStringKeyValueOnly("k", "v")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns nil -- nil", actual)
}

func Test_Attributes_AddNewStringKeyValueOnly_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"val": a.AddNewStringKeyValueOnly("k", "v")}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AddNewStringKeyValueOnly returns non-empty -- valid", actual)
}

func Test_Attributes_AddNewAnyKeyValueOnly_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.AddNewAnyKeyValueOnly("k", 42)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns nil -- nil", actual)
}

func Test_Attributes_AddNewAnyKeyValueOnly_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{"val": a.AddNewAnyKeyValueOnly("k", 42)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AddNewAnyKeyValueOnly returns non-empty -- valid", actual)
}

func Test_Attributes_AddOrUpdateString_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.AddOrUpdateString("k", "v")}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateString returns nil -- nil", actual)
}

func Test_Attributes_AddOrUpdateAnyItem_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes

	// Act
	actual := args.Map{"val": a.AddOrUpdateAnyItem("k", 42)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "AddOrUpdateAnyItem returns nil -- nil", actual)
}

func Test_Attributes_SetBasicErr_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	result := a.SetBasicErr(nil)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SetBasicErr returns nil -- nil", actual)
}

func Test_Attributes_Clear_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	a.Clear() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Clear returns nil -- nil", actual)
}

func Test_Attributes_Clear_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))
	a.Clear()

	// Act
	actual := args.Map{"len": len(a.DynamicPayloads)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns non-empty -- valid", actual)
}

func Test_Attributes_Dispose_Nil_FromAttributesIsNull(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	a.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose returns nil -- nil", actual)
}

func Test_Attributes_HandleErr_NoError_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.HandleErr() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr returns empty -- no error", actual)
}

func Test_Attributes_HandleError_NoError_FromAttributesIsNull(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.HandleError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError returns empty -- no error", actual)
}

func Test_Attributes_MustBeEmptyError_Empty(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	a.MustBeEmptyError() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Attributes — NewAttributesCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewAttributes_Empty(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()

	// Act
	actual := args.Map{
		"notNil": a != nil,
		"valid": a.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Empty returns empty -- with args", actual)
}

func Test_NewAttributes_UsingDynamicPayloadBytes(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingDynamicPayloadBytes([]byte("data"))

	// Act
	actual := args.Map{"len": len(a.DynamicPayloads)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingDynamicPayloadBytes returns correct value -- with args", actual)
}

func Test_NewAttributes_UsingAuthInfo(t *testing.T) {
	// Arrange
	ai := &corepayload.AuthInfo{ActionType: "login"}
	a := corepayload.New.Attributes.UsingAuthInfo(ai)

	// Act
	actual := args.Map{"action": a.AuthInfo.ActionType}

	// Assert
	expected := args.Map{"action": "login"}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingAuthInfo returns correct value -- with args", actual)
}

func Test_NewAttributes_UsingAuthInfoDynamicBytes(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.UsingAuthInfoDynamicBytes(&corepayload.AuthInfo{ActionType: "x"}, []byte("data"))

	// Act
	actual := args.Map{
		"action": a.AuthInfo.ActionType,
		"len": len(a.DynamicPayloads),
	}

	// Assert
	expected := args.Map{
		"action": "x",
		"len": 4,
	}
	expected.ShouldBeEqual(t, 0, "NewAttributes.UsingAuthInfoDynamicBytes returns correct value -- with args", actual)
}

func Test_NewAttributes_Deserialize_Valid(t *testing.T) {
	// Arrange
	a := corepayload.New.Attributes.Empty()
	b := a.JsonPtr().Bytes
	a2, err := corepayload.New.Attributes.Deserialize(b)

	// Act
	actual := args.Map{
		"notNil": a2 != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Deserialize returns correct value -- with args", actual)
}

func Test_NewAttributes_Deserialize_Bad(t *testing.T) {
	// Arrange
	_, err := corepayload.New.Attributes.Deserialize([]byte("{bad"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.Deserialize returns correct value -- bad", actual)
}

func Test_NewAttributes_DeserializeMany_Bad(t *testing.T) {
	_, err := corepayload.New.Attributes.DeserializeMany([]byte("{bad"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.DeserializeMany returns correct value -- bad", actual)
}

func Test_NewAttributes_CastOrDeserializeFrom_Nil_FromAttributesIsNull(t *testing.T) {
	_, err := corepayload.New.Attributes.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewAttributes.CastOrDeserializeFrom returns nil -- nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// EmptyCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Empty_Attributes(t *testing.T) {
	a := corepayload.Empty.Attributes()
	actual := args.Map{"notNil": a != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.Attributes returns empty -- with args", actual)
}

func Test_Empty_AttributesDefaults(t *testing.T) {
	a := corepayload.Empty.AttributesDefaults()
	actual := args.Map{
		"notNil": a != nil,
		"hasKV": a.KeyValuePairs != nil,
		"hasAKV": a.AnyKeyValuePairs != nil,
	}
	expected := args.Map{
		"notNil": true,
		"hasKV": true,
		"hasAKV": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty.AttributesDefaults returns empty -- with args", actual)
}

func Test_Empty_PayloadWrapper(t *testing.T) {
	pw := corepayload.Empty.PayloadWrapper()
	actual := args.Map{"notNil": pw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadWrapper returns empty -- with args", actual)
}

func Test_Empty_PayloadsCollection(t *testing.T) {
	pc := corepayload.Empty.PayloadsCollection()
	actual := args.Map{
		"notNil": pc != nil,
		"empty": pc.IsEmpty(),
	}
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadsCollection returns empty -- with args", actual)
}
