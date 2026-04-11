package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// User — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_User_IdentifierInteger_Empty(t *testing.T) {
	// Arrange
	u := corepayload.User{}

	// Act
	actual := args.Map{"val": u.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns empty -- empty", actual)
}

func Test_User_IdentifierInteger_Valid(t *testing.T) {
	// Arrange
	u := corepayload.User{Identifier: "42"}

	// Act
	actual := args.Map{"val": u.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns non-empty -- valid", actual)
}

func Test_User_IdentifierUnsignedInteger_Negative(t *testing.T) {
	// Arrange
	u := corepayload.User{Identifier: "-1"}

	// Act
	actual := args.Map{"val": u.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "IdentifierUnsignedInteger returns correct value -- negative", actual)
}

func Test_User_IdentifierUnsignedInteger_Valid(t *testing.T) {
	// Arrange
	u := corepayload.User{Identifier: "42"}

	// Act
	actual := args.Map{"val": u.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"val": uint(42)}
	expected.ShouldBeEqual(t, 0, "IdentifierUnsignedInteger returns non-empty -- valid", actual)
}

func Test_User_HasAuthToken_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u1 := &corepayload.User{AuthToken: "tok"}
	var u2 *corepayload.User

	// Act
	actual := args.Map{
		"has": u1.HasAuthToken(),
		"nil": u2.HasAuthToken(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAuthToken returns correct value -- with args", actual)
}

func Test_User_HasPasswordHash_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := &corepayload.User{PasswordHash: "hash"}

	// Act
	actual := args.Map{"has": u.HasPasswordHash()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasPasswordHash returns correct value -- with args", actual)
}

func Test_User_IsPasswordHashEmpty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := &corepayload.User{}
	var u2 *corepayload.User

	// Act
	actual := args.Map{
		"empty": u.IsPasswordHashEmpty(),
		"nil": u2.IsPasswordHashEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsPasswordHashEmpty returns empty -- with args", actual)
}

func Test_User_IsAuthTokenEmpty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := &corepayload.User{}

	// Act
	actual := args.Map{"empty": u.IsAuthTokenEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "IsAuthTokenEmpty returns empty -- with args", actual)
}

func Test_User_IsEmpty(t *testing.T) {
	// Arrange
	u1 := &corepayload.User{}
	u2 := &corepayload.User{Name: "test"}
	var u3 *corepayload.User

	// Act
	actual := args.Map{
		"empty": u1.IsEmpty(),
		"notEmpty": u2.IsEmpty(),
		"nil": u3.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_User_IsValidUser(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "test"}

	// Act
	actual := args.Map{"val": u.IsValidUser()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsValidUser returns non-empty -- with args", actual)
}

func Test_User_IsNameEmpty(t *testing.T) {
	// Arrange
	u := &corepayload.User{}

	// Act
	actual := args.Map{"val": u.IsNameEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsNameEmpty returns empty -- with args", actual)
}

func Test_User_IsNameEqual(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	var u2 *corepayload.User

	// Act
	actual := args.Map{
		"match": u.IsNameEqual("alice"),
		"no": u.IsNameEqual("bob"),
		"nil": u2.IsNameEqual("alice"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"no": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNameEqual returns correct value -- with args", actual)
}

func Test_User_IsNotSystemUser(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "u"}

	// Act
	actual := args.Map{"val": u.IsNotSystemUser()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsNotSystemUser returns correct value -- with args", actual)
}

func Test_User_IsVirtualUser(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "u"}

	// Act
	actual := args.Map{"val": u.IsVirtualUser()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsVirtualUser returns correct value -- with args", actual)
}

func Test_User_HasType(t *testing.T) {
	// Arrange
	u := &corepayload.User{Type: "admin"}

	// Act
	actual := args.Map{"has": u.HasType()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasType returns correct value -- with args", actual)
}

func Test_User_IsTypeEmpty(t *testing.T) {
	// Arrange
	u := &corepayload.User{}
	var u2 *corepayload.User

	// Act
	actual := args.Map{
		"empty": u.IsTypeEmpty(),
		"nil": u2.IsTypeEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsTypeEmpty returns empty -- with args", actual)
}

func Test_User_String(t *testing.T) {
	// Arrange
	u := corepayload.User{Name: "alice"}

	// Act
	actual := args.Map{"notEmpty": u.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_User_PrettyJsonString_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}

	// Act
	actual := args.Map{"notEmpty": u.PrettyJsonString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns correct value -- with args", actual)
}

func Test_User_Json(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	r := u.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- with args", actual)
}

func Test_User_JsonPtr(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	r := u.JsonPtr()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- with args", actual)
}

func Test_User_Serialize(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	b, err := u.Serialize()

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
	expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- with args", actual)
}

func Test_User_Deserialize(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	b, _ := u.Serialize()
	u2 := &corepayload.User{}
	err := u2.Deserialize(b)

	// Act
	actual := args.Map{
		"name": u2.Name,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_User_Clone_FromUserIdentifierIntege_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.User{Name: "alice", Identifier: "id1", Type: "admin", AuthToken: "tok", PasswordHash: "hash", IsSystemUser: true}
	c := u.Clone()

	// Act
	actual := args.Map{
		"name": c.Name,
		"type": c.Type,
		"token": c.AuthToken,
		"hash": c.PasswordHash,
		"sys": c.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"type": "admin",
		"token": "tok",
		"hash": "hash",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_User_ClonePtr_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var u *corepayload.User

	// Act
	actual := args.Map{"nil": u.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_User_ClonePtr_Valid(t *testing.T) {
	// Arrange
	u := &corepayload.User{Name: "alice"}
	c := u.ClonePtr()

	// Act
	actual := args.Map{
		"name": c.Name,
		"diff": c != u,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns non-empty -- valid", actual)
}

func Test_User_Ptr_FromUserIdentifierIntege_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.User{Name: "alice"}
	p := u.Ptr()

	// Act
	actual := args.Map{
		"notNil": p != nil,
		"name": p.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"name": "alice",
	}
	expected.ShouldBeEqual(t, 0, "Ptr returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// UserInfo — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_UserInfo_SetUserSystemUser_Nil(t *testing.T) {
	// Arrange
	var ui *corepayload.UserInfo
	u := &corepayload.User{Name: "a"}
	su := &corepayload.User{Name: "b"}
	result := ui.SetUserSystemUser(u, su)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"user": result.User.Name,
		"sys": result.SystemUser.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"user": "a",
		"sys": "b",
	}
	expected.ShouldBeEqual(t, 0, "SetUserSystemUser returns nil -- nil", actual)
}

func Test_UserInfo_SetUserSystemUser_NonNil(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{}
	u := &corepayload.User{Name: "a"}
	su := &corepayload.User{Name: "b"}
	result := ui.SetUserSystemUser(u, su)

	// Act
	actual := args.Map{
		"same": result == ui,
		"user": result.User.Name,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"user": "a",
	}
	expected.ShouldBeEqual(t, 0, "SetUserSystemUser returns nil -- non-nil", actual)
}

func Test_UserInfo_SetUser_Nil(t *testing.T) {
	// Arrange
	var ui *corepayload.UserInfo
	result := ui.SetUser(&corepayload.User{Name: "a"})

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"user": result.User.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"user": "a",
	}
	expected.ShouldBeEqual(t, 0, "SetUser returns nil -- nil", actual)
}

func Test_UserInfo_SetUser_NonNil(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{}
	result := ui.SetUser(&corepayload.User{Name: "a"})

	// Act
	actual := args.Map{"same": result == ui}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SetUser returns nil -- non-nil", actual)
}

func Test_UserInfo_SetSystemUser_Nil(t *testing.T) {
	// Arrange
	var ui *corepayload.UserInfo
	result := ui.SetSystemUser(&corepayload.User{Name: "sys"})

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"name": result.SystemUser.Name,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"name": "sys",
	}
	expected.ShouldBeEqual(t, 0, "SetSystemUser returns nil -- nil", actual)
}

func Test_UserInfo_HasUser(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{User: &corepayload.User{Name: "a"}}
	var ui2 *corepayload.UserInfo

	// Act
	actual := args.Map{
		"has": ui.HasUser(),
		"nil": ui2.HasUser(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasUser returns correct value -- with args", actual)
}

func Test_UserInfo_HasSystemUser(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{SystemUser: &corepayload.User{Name: "sys"}}

	// Act
	actual := args.Map{"has": ui.HasSystemUser()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasSystemUser returns correct value -- with args", actual)
}

func Test_UserInfo_IsEmpty(t *testing.T) {
	// Arrange
	ui1 := &corepayload.UserInfo{}
	ui2 := &corepayload.UserInfo{User: &corepayload.User{Name: "a"}}
	var ui3 *corepayload.UserInfo

	// Act
	actual := args.Map{
		"empty": ui1.IsEmpty(),
		"notEmpty": ui2.IsEmpty(),
		"nil": ui3.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_UserInfo_IsUserEmpty(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{}

	// Act
	actual := args.Map{"val": ui.IsUserEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsUserEmpty returns empty -- with args", actual)
}

func Test_UserInfo_IsSystemUserEmpty(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{}

	// Act
	actual := args.Map{"val": ui.IsSystemUserEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsSystemUserEmpty returns empty -- with args", actual)
}

func Test_UserInfo_Clone_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{
		User:       &corepayload.User{Name: "a"},
		SystemUser: &corepayload.User{Name: "b"},
	}
	c := ui.Clone()

	// Act
	actual := args.Map{
		"user": c.User.Name,
		"sys": c.SystemUser.Name,
	}

	// Assert
	expected := args.Map{
		"user": "a",
		"sys": "b",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_UserInfo_ClonePtr_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var ui *corepayload.UserInfo

	// Act
	actual := args.Map{"nil": ui.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_UserInfo_Ptr(t *testing.T) {
	// Arrange
	ui := corepayload.UserInfo{User: &corepayload.User{Name: "a"}}
	p := ui.Ptr()

	// Act
	actual := args.Map{"notNil": p != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Ptr returns correct value -- with args", actual)
}

func Test_UserInfo_ToNonPtr_Nil(t *testing.T) {
	// Arrange
	var ui *corepayload.UserInfo
	np := ui.ToNonPtr()

	// Act
	actual := args.Map{"empty": np.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ToNonPtr returns nil -- nil", actual)
}

func Test_UserInfo_ToNonPtr_Valid(t *testing.T) {
	// Arrange
	ui := &corepayload.UserInfo{User: &corepayload.User{Name: "a"}}
	np := ui.ToNonPtr()

	// Act
	actual := args.Map{"name": np.User.Name}

	// Assert
	expected := args.Map{"name": "a"}
	expected.ShouldBeEqual(t, 0, "ToNonPtr returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SessionInfo — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SessionInfo_IdentifierInteger_Empty(t *testing.T) {
	// Arrange
	si := corepayload.SessionInfo{}

	// Act
	actual := args.Map{"val": si.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "SI returns empty -- IdentifierInteger empty", actual)
}

func Test_SessionInfo_IdentifierInteger_Valid(t *testing.T) {
	// Arrange
	si := corepayload.SessionInfo{Id: "42"}

	// Act
	actual := args.Map{"val": si.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "SI returns non-empty -- IdentifierInteger valid", actual)
}

func Test_SessionInfo_IdentifierUnsignedInteger_Negative(t *testing.T) {
	// Arrange
	si := corepayload.SessionInfo{Id: "-1"}

	// Act
	actual := args.Map{"val": si.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "SI returns correct value -- IdentifierUnsignedInteger negative", actual)
}

func Test_SessionInfo_IsEmpty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	s1 := &corepayload.SessionInfo{}
	s2 := &corepayload.SessionInfo{Id: "1"}
	var s3 *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"empty": s1.IsEmpty(),
		"notEmpty": s2.IsEmpty(),
		"nil": s3.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "SI returns empty -- IsEmpty", actual)
}

func Test_SessionInfo_IsValid(t *testing.T) {
	// Arrange
	s1 := &corepayload.SessionInfo{Id: "1", User: &corepayload.User{Name: "a"}}
	s2 := &corepayload.SessionInfo{}

	// Act
	actual := args.Map{
		"valid": s1.IsValid(),
		"invalid": s2.IsValid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "SI returns non-empty -- IsValid", actual)
}

func Test_SessionInfo_IsUserNameEmpty(t *testing.T) {
	// Arrange
	s := &corepayload.SessionInfo{}
	var s2 *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"empty": s.IsUserNameEmpty(),
		"nil": s2.IsUserNameEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsUserNameEmpty returns empty -- with args", actual)
}

func Test_SessionInfo_IsUserEmpty(t *testing.T) {
	// Arrange
	s := &corepayload.SessionInfo{}

	// Act
	actual := args.Map{"val": s.IsUserEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsUserEmpty returns empty -- with args", actual)
}

func Test_SessionInfo_HasUser(t *testing.T) {
	// Arrange
	s := &corepayload.SessionInfo{User: &corepayload.User{Name: "a"}}

	// Act
	actual := args.Map{"val": s.HasUser()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "SI returns correct value -- HasUser", actual)
}

func Test_SessionInfo_IsUsernameEqual_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	s := &corepayload.SessionInfo{User: &corepayload.User{Name: "alice"}}
	var s2 *corepayload.SessionInfo

	// Act
	actual := args.Map{
		"match": s.IsUsernameEqual("alice"),
		"no": s.IsUsernameEqual("bob"),
		"nil": s2.IsUsernameEqual("alice"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"no": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "IsUsernameEqual returns correct value -- with args", actual)
}

func Test_SessionInfo_Clone_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	s := corepayload.SessionInfo{Id: "1", User: &corepayload.User{Name: "a"}, SessionPath: "/p"}
	c := s.Clone()

	// Act
	actual := args.Map{
		"id": c.Id,
		"user": c.User.Name,
		"path": c.SessionPath,
	}

	// Assert
	expected := args.Map{
		"id": "1",
		"user": "a",
		"path": "/p",
	}
	expected.ShouldBeEqual(t, 0, "SI returns correct value -- Clone", actual)
}

func Test_SessionInfo_ClonePtr_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var s *corepayload.SessionInfo

	// Act
	actual := args.Map{"nil": s.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SI returns nil -- ClonePtr nil", actual)
}

func Test_SessionInfo_Ptr(t *testing.T) {
	// Arrange
	s := corepayload.SessionInfo{Id: "1"}
	p := s.Ptr()

	// Act
	actual := args.Map{"notNil": p != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SI returns correct value -- Ptr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AuthInfo — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_AuthInfo_IdentifierInteger_Empty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	ai := corepayload.AuthInfo{}

	// Act
	actual := args.Map{"val": ai.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "AI returns empty -- IdentifierInteger empty", actual)
}

func Test_AuthInfo_IdentifierInteger_Valid(t *testing.T) {
	// Arrange
	ai := corepayload.AuthInfo{Identifier: "10"}

	// Act
	actual := args.Map{"val": ai.IdentifierInteger()}

	// Assert
	expected := args.Map{"val": 10}
	expected.ShouldBeEqual(t, 0, "AI returns non-empty -- IdentifierInteger valid", actual)
}

func Test_AuthInfo_IdentifierUnsignedInteger_Negative_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	ai := corepayload.AuthInfo{Identifier: "-5"}

	// Act
	actual := args.Map{"val": ai.IdentifierUnsignedInteger()}

	// Assert
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- IdentifierUnsignedInteger negative", actual)
}

func Test_AuthInfo_IsEmpty(t *testing.T) {
	// Arrange
	a1 := &corepayload.AuthInfo{}
	a2 := &corepayload.AuthInfo{ActionType: "login"}
	var a3 *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"empty": a1.IsEmpty(),
		"notEmpty": a2.IsEmpty(),
		"nil": a3.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "AI returns empty -- IsEmpty", actual)
}

func Test_AuthInfo_HasAnyItem(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ActionType: "x"}

	// Act
	actual := args.Map{"val": a.HasAnyItem()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- HasAnyItem", actual)
}

func Test_AuthInfo_IsActionTypeEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	var a2 *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"empty": a.IsActionTypeEmpty(),
		"nil": a2.IsActionTypeEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsActionTypeEmpty returns empty -- with args", actual)
}

func Test_AuthInfo_IsResourceNameEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}

	// Act
	actual := args.Map{"val": a.IsResourceNameEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsResourceNameEmpty returns empty -- with args", actual)
}

func Test_AuthInfo_IsValid(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ActionType: "x"}

	// Act
	actual := args.Map{"val": a.IsValid()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AI returns non-empty -- IsValid", actual)
}

func Test_AuthInfo_HasActionType(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ActionType: "x"}
	var a2 *corepayload.AuthInfo

	// Act
	actual := args.Map{
		"has": a.HasActionType(),
		"nil": a2.HasActionType(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasActionType returns correct value -- with args", actual)
}

func Test_AuthInfo_HasResourceName(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{ResourceName: "r"}

	// Act
	actual := args.Map{"has": a.HasResourceName()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasResourceName returns correct value -- with args", actual)
}

func Test_AuthInfo_IsUserInfoEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	var a2 *corepayload.AuthInfo

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
	expected.ShouldBeEqual(t, 0, "AI returns empty -- IsUserInfoEmpty", actual)
}

func Test_AuthInfo_IsSessionInfoEmpty(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}

	// Act
	actual := args.Map{"val": a.IsSessionInfoEmpty()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AI returns empty -- IsSessionInfoEmpty", actual)
}

func Test_AuthInfo_HasUserInfo(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{User: &corepayload.User{Name: "u"}}}

	// Act
	actual := args.Map{"val": a.HasUserInfo()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- HasUserInfo", actual)
}

func Test_AuthInfo_HasSessionInfo(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{SessionInfo: &corepayload.SessionInfo{Id: "1"}}

	// Act
	actual := args.Map{"val": a.HasSessionInfo()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- HasSessionInfo", actual)
}

func Test_AuthInfo_String(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{ActionType: "login"}

	// Act
	actual := args.Map{"notEmpty": a.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- String", actual)
}

func Test_AuthInfo_PrettyJsonString(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{ActionType: "x"}

	// Act
	actual := args.Map{"notEmpty": a.PrettyJsonString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- PrettyJsonString", actual)
}

func Test_AuthInfo_Json_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{ActionType: "x"}
	r := a.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- Json", actual)
}

func Test_AuthInfo_JsonPtr(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{ActionType: "x"}

	// Act
	actual := args.Map{"notNil": a.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- JsonPtr", actual)
}

func Test_AuthInfo_Clone_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{
		Identifier:   "id1",
		ActionType:   "login",
		ResourceName: "res",
		SessionInfo:  &corepayload.SessionInfo{Id: "s1"},
		UserInfo:     &corepayload.UserInfo{User: &corepayload.User{Name: "u"}},
	}
	c := a.Clone()

	// Act
	actual := args.Map{
		"id": c.Identifier,
		"action": c.ActionType,
		"res": c.ResourceName,
		"hasSession": c.SessionInfo != nil,
		"hasUser": c.UserInfo != nil,
	}

	// Assert
	expected := args.Map{
		"id": "id1",
		"action": "login",
		"res": "res",
		"hasSession": true,
		"hasUser": true,
	}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- Clone", actual)
}

func Test_AuthInfo_ClonePtr_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo

	// Act
	actual := args.Map{"nil": a.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- ClonePtr nil", actual)
}

func Test_AuthInfo_Ptr(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{ActionType: "x"}

	// Act
	actual := args.Map{"notNil": a.Ptr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns correct value -- Ptr", actual)
}

// AuthInfo Set* methods
func Test_AuthInfo_SetUserInfo_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetUserInfo(&corepayload.UserInfo{User: &corepayload.User{Name: "u"}})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetUserInfo nil", actual)
}

func Test_AuthInfo_SetUserInfo_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	result := a.SetUserInfo(&corepayload.UserInfo{})

	// Act
	actual := args.Map{"same": result == a}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetUserInfo non-nil", actual)
}

func Test_AuthInfo_SetActionType_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetActionType("login")

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"action": result.ActionType,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"action": "login",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetActionType nil", actual)
}

func Test_AuthInfo_SetActionType_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	result := a.SetActionType("login")

	// Act
	actual := args.Map{
		"same": result == a,
		"action": result.ActionType,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"action": "login",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetActionType non-nil", actual)
}

func Test_AuthInfo_SetResourceName_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetResourceName("res")

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"res": result.ResourceName,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"res": "res",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetResourceName nil", actual)
}

func Test_AuthInfo_SetResourceName_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	a.SetResourceName("res")

	// Act
	actual := args.Map{"res": a.ResourceName}

	// Assert
	expected := args.Map{"res": "res"}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetResourceName non-nil", actual)
}

func Test_AuthInfo_SetIdentifier_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetIdentifier("id1")

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"id": result.Identifier,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"id": "id1",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetIdentifier nil", actual)
}

func Test_AuthInfo_SetIdentifier_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{}
	a.SetIdentifier("id1")

	// Act
	actual := args.Map{"id": a.Identifier}

	// Assert
	expected := args.Map{"id": "id1"}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetIdentifier non-nil", actual)
}

func Test_AuthInfo_SetSessionInfo_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetSessionInfo(&corepayload.SessionInfo{Id: "s1"})

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"id": result.SessionInfo.Id,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"id": "s1",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetSessionInfo nil", actual)
}

func Test_AuthInfo_SetUserSystemUser_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	u := &corepayload.User{Name: "u"}
	su := &corepayload.User{Name: "su"}
	result := a.SetUserSystemUser(u, su)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetUserSystemUser nil", actual)
}

func Test_AuthInfo_SetUserSystemUser_NonNil(t *testing.T) {
	// Arrange
	a := &corepayload.AuthInfo{UserInfo: &corepayload.UserInfo{}}
	a.SetUserSystemUser(&corepayload.User{Name: "u"}, &corepayload.User{Name: "su"})

	// Act
	actual := args.Map{
		"user": a.UserInfo.User.Name,
		"sys": a.UserInfo.SystemUser.Name,
	}

	// Assert
	expected := args.Map{
		"user": "u",
		"sys": "su",
	}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetUserSystemUser non-nil", actual)
}

func Test_AuthInfo_SetUser_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetUser(&corepayload.User{Name: "u"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetUser nil", actual)
}

func Test_AuthInfo_SetSystemUser_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.AuthInfo
	result := a.SetSystemUser(&corepayload.User{Name: "su"})

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AI returns nil -- SetSystemUser nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PagingInfo — All methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_PagingInfo_IsEmpty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{}
	p2 := &corepayload.PagingInfo{TotalPages: 5, TotalItems: 50}
	var p3 *corepayload.PagingInfo

	// Act
	actual := args.Map{
		"empty": p1.IsEmpty(),
		"notEmpty": p2.IsEmpty(),
		"nil": p3.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "PI returns empty -- IsEmpty", actual)
}

func Test_PagingInfo_IsEqual_BothNil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var p1, p2 *corepayload.PagingInfo

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "PI returns nil -- IsEqual both nil", actual)
}

func Test_PagingInfo_IsEqual_OneNil(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{TotalPages: 1}

	// Act
	actual := args.Map{"val": p.IsEqual(nil)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PI returns nil -- IsEqual one nil", actual)
}

func Test_PagingInfo_IsEqual_DiffTotalPages(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{TotalPages: 1}
	p2 := &corepayload.PagingInfo{TotalPages: 2}

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- IsEqual diff total pages", actual)
}

func Test_PagingInfo_IsEqual_DiffCurrentPage(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1}
	p2 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 2}

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- IsEqual diff current page", actual)
}

func Test_PagingInfo_IsEqual_DiffPerPage(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10}
	p2 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 20}

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- IsEqual diff per page", actual)
}

func Test_PagingInfo_IsEqual_DiffTotal(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}
	p2 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 20}

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- IsEqual diff total items", actual)
}

func Test_PagingInfo_IsEqual_Same(t *testing.T) {
	// Arrange
	p1 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}
	p2 := &corepayload.PagingInfo{TotalPages: 1, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 10}

	// Act
	actual := args.Map{"val": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- IsEqual same", actual)
}

func Test_PagingInfo_HasMethods_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 50}
	var pNil *corepayload.PagingInfo

	// Act
	actual := args.Map{
		"hasTotalPages": p.HasTotalPages(), "hasCurrentPage": p.HasCurrentPageIndex(),
		"hasPerPage": p.HasPerPageItems(), "hasTotal": p.HasTotalItems(),
		"nilTotalPages": pNil.HasTotalPages(), "nilCurrentPage": pNil.HasCurrentPageIndex(),
	}

	// Assert
	expected := args.Map{
		"hasTotalPages": true, "hasCurrentPage": true,
		"hasPerPage": true, "hasTotal": true,
		"nilTotalPages": false, "nilCurrentPage": false,
	}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- Has methods", actual)
}

func Test_PagingInfo_IsInvalidMethods_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{}
	var pNil *corepayload.PagingInfo

	// Act
	actual := args.Map{
		"invTotalPages": p.IsInvalidTotalPages(), "invCurrentPage": p.IsInvalidCurrentPageIndex(),
		"invPerPage": p.IsInvalidPerPageItems(), "invTotal": p.IsInvalidTotalItems(),
		"nilInvTotal": pNil.IsInvalidTotalPages(),
	}

	// Assert
	expected := args.Map{
		"invTotalPages": true, "invCurrentPage": true,
		"invPerPage": true, "invTotal": true,
		"nilInvTotal": true,
	}
	expected.ShouldBeEqual(t, 0, "PI returns error -- IsInvalid methods", actual)
}

func Test_PagingInfo_Clone_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	p := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 50}
	c := p.Clone()

	// Act
	actual := args.Map{
		"total": c.TotalPages,
		"current": c.CurrentPageIndex,
		"perPage": c.PerPageItems,
		"items": c.TotalItems,
	}

	// Assert
	expected := args.Map{
		"total": 5,
		"current": 2,
		"perPage": 10,
		"items": 50,
	}
	expected.ShouldBeEqual(t, 0, "PI returns correct value -- Clone", actual)
}

func Test_PagingInfo_ClonePtr_Nil_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	var p *corepayload.PagingInfo

	// Act
	actual := args.Map{"nil": p.ClonePtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "PI returns nil -- ClonePtr nil", actual)
}

func Test_PagingInfo_ClonePtr_Valid(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{TotalPages: 5}
	c := p.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"total": c.TotalPages,
		"diff": c != p,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"total": 5,
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "PI returns non-empty -- ClonePtr valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NewUser Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewUser_Empty_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.Empty()

	// Act
	actual := args.Map{
		"notNil": u != nil,
		"empty": u.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.Empty returns empty -- with args", actual)
}

func Test_NewUser_Create_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.Create(true, "alice", "admin")

	// Act
	actual := args.Map{
		"name": u.Name,
		"type": u.Type,
		"sys": u.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"type": "admin",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.Create returns correct value -- with args", actual)
}

func Test_NewUser_NonSysCreate(t *testing.T) {
	// Arrange
	u := corepayload.New.User.NonSysCreate("alice", "admin")

	// Act
	actual := args.Map{
		"name": u.Name,
		"sys": u.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"sys": false,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.NonSysCreate returns correct value -- with args", actual)
}

func Test_NewUser_NonSysCreateId(t *testing.T) {
	// Arrange
	u := corepayload.New.User.NonSysCreateId("id1", "alice", "admin")

	// Act
	actual := args.Map{
		"id": u.Identifier,
		"name": u.Name,
	}

	// Assert
	expected := args.Map{
		"id": "id1",
		"name": "alice",
	}
	expected.ShouldBeEqual(t, 0, "NewUser.NonSysCreateId returns correct value -- with args", actual)
}

func Test_NewUser_System_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.System("sys", "os")

	// Act
	actual := args.Map{
		"name": u.Name,
		"sys": u.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"name": "sys",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.System returns correct value -- with args", actual)
}

func Test_NewUser_SystemId_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.SystemId("id1", "sys", "os")

	// Act
	actual := args.Map{
		"id": u.Identifier,
		"sys": u.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"id": "id1",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.SystemId returns correct value -- with args", actual)
}

func Test_NewUser_UsingName_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.UsingName("alice")

	// Act
	actual := args.Map{"name": u.Name}

	// Assert
	expected := args.Map{"name": "alice"}
	expected.ShouldBeEqual(t, 0, "NewUser.UsingName returns correct value -- with args", actual)
}

func Test_NewUser_All_FromUserIdentifierIntege_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.All(true, "id1", "alice", "admin", "tok", "hash")

	// Act
	actual := args.Map{
		"id": u.Identifier,
		"name": u.Name,
		"type": u.Type,
		"token": u.AuthToken,
		"hash": u.PasswordHash,
		"sys": u.IsSystemUser,
	}

	// Assert
	expected := args.Map{
		"id": "id1",
		"name": "alice",
		"type": "admin",
		"token": "tok",
		"hash": "hash",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.All returns correct value -- with args", actual)
}

func Test_NewUser_Deserialize_FromUserIdentifierIntege_FromUserIdentifierIntege(t *testing.T) {
	// Arrange
	u := corepayload.New.User.UsingName("alice")
	b, _ := u.Serialize()
	u2, err := corepayload.New.User.Deserialize(b)

	// Act
	actual := args.Map{
		"name": u2.Name,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"name": "alice",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.Deserialize returns correct value -- with args", actual)
}

func Test_NewUser_Deserialize_Bad(t *testing.T) {
	// Arrange
	_, err := corepayload.New.User.Deserialize([]byte("{bad"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewUser.Deserialize returns correct value -- bad", actual)
}

func Test_NewUser_CastOrDeserializeFrom_Nil(t *testing.T) {
	_, err := corepayload.New.User.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewUser.CastOrDeserializeFrom returns nil -- nil", actual)
}
