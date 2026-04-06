package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corepayload"
)

// ── User methods ──

func Test_Cov_User_IdentifierInteger(t *testing.T) {
	u := corepayload.User{Identifier: "42"}
	empty := corepayload.User{}
	invalid := corepayload.User{Identifier: "abc"}

	actual := args.Map{
		"valid":   u.IdentifierInteger(),
		"empty":   empty.IdentifierInteger(),
		"invalid": invalid.IdentifierInteger(),
	}
	expected := args.Map{
		"valid":   42,
		"empty":   -1,
		"invalid": -1,
	}
	expected.ShouldBeEqual(t, 0, "IdentifierInteger returns correct value -- with args", actual)
}

func Test_Cov_User_IdentifierUnsignedInteger(t *testing.T) {
	u := corepayload.User{Identifier: "5"}
	neg := corepayload.User{Identifier: "abc"}

	actual := args.Map{
		"valid":    u.IdentifierUnsignedInteger(),
		"negative": neg.IdentifierUnsignedInteger(),
	}
	expected := args.Map{
		"valid":    uint(5),
		"negative": uint(0),
	}
	expected.ShouldBeEqual(t, 0, "IdentifierUnsignedInteger returns correct value -- with args", actual)
}

func Test_Cov_User_HasAuthToken(t *testing.T) {
	u := &corepayload.User{AuthToken: "tok"}
	empty := &corepayload.User{}
	var nilUser *corepayload.User

	actual := args.Map{
		"has":   u.HasAuthToken(),
		"empty": empty.HasAuthToken(),
		"nil":   nilUser.HasAuthToken(),
	}
	expected := args.Map{
		"has":   true,
		"empty": false,
		"nil":   false,
	}
	expected.ShouldBeEqual(t, 0, "HasAuthToken returns correct value -- with args", actual)
}

func Test_Cov_User_HasPasswordHash(t *testing.T) {
	u := &corepayload.User{PasswordHash: "hash"}
	empty := &corepayload.User{}
	var nilUser *corepayload.User

	actual := args.Map{
		"has":   u.HasPasswordHash(),
		"empty": empty.HasPasswordHash(),
		"nil":   nilUser.HasPasswordHash(),
	}
	expected := args.Map{
		"has":   true,
		"empty": false,
		"nil":   false,
	}
	expected.ShouldBeEqual(t, 0, "HasPasswordHash returns correct value -- with args", actual)
}

func Test_Cov_User_IsPasswordHashEmpty(t *testing.T) {
	u := &corepayload.User{PasswordHash: "h"}
	var nilUser *corepayload.User

	actual := args.Map{
		"nonEmpty": u.IsPasswordHashEmpty(),
		"nil":      nilUser.IsPasswordHashEmpty(),
	}
	expected := args.Map{
		"nonEmpty": false,
		"nil":      true,
	}
	expected.ShouldBeEqual(t, 0, "IsPasswordHashEmpty returns empty -- with args", actual)
}

func Test_Cov_User_IsAuthTokenEmpty(t *testing.T) {
	u := &corepayload.User{AuthToken: "t"}
	var nilUser *corepayload.User

	actual := args.Map{
		"nonEmpty": u.IsAuthTokenEmpty(),
		"nil":      nilUser.IsAuthTokenEmpty(),
	}
	expected := args.Map{
		"nonEmpty": false,
		"nil":      true,
	}
	expected.ShouldBeEqual(t, 0, "IsAuthTokenEmpty returns empty -- with args", actual)
}

func Test_Cov_User_IsEmpty_IsValidUser(t *testing.T) {
	u := &corepayload.User{Name: "test"}
	empty := &corepayload.User{}
	var nilUser *corepayload.User

	actual := args.Map{
		"isEmpty":     u.IsEmpty(),
		"emptyEmpty":  empty.IsEmpty(),
		"nilEmpty":    nilUser.IsEmpty(),
		"isValid":     u.IsValidUser(),
		"emptyValid":  empty.IsValidUser(),
	}
	expected := args.Map{
		"isEmpty":     false,
		"emptyEmpty":  true,
		"nilEmpty":    true,
		"isValid":     true,
		"emptyValid":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty/IsValidUser returns empty -- with args", actual)
}

func Test_Cov_User_IsNameEmpty_IsNameEqual(t *testing.T) {
	u := &corepayload.User{Name: "john"}
	var nilUser *corepayload.User

	actual := args.Map{
		"nameEmpty":  u.IsNameEmpty(),
		"nilEmpty":   nilUser.IsNameEmpty(),
		"nameEqual":  u.IsNameEqual("john"),
		"nameNotEq":  u.IsNameEqual("bob"),
		"nilEqual":   nilUser.IsNameEqual("john"),
	}
	expected := args.Map{
		"nameEmpty":  false,
		"nilEmpty":   true,
		"nameEqual":  true,
		"nameNotEq":  false,
		"nilEqual":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsNameEmpty/IsNameEqual returns empty -- with args", actual)
}

func Test_Cov_User_IsNotSystemUser_IsVirtualUser(t *testing.T) {
	u := &corepayload.User{IsSystemUser: false}
	sys := &corepayload.User{IsSystemUser: true}
	var nilUser *corepayload.User

	actual := args.Map{
		"notSys":     u.IsNotSystemUser(),
		"isSys":      sys.IsNotSystemUser(),
		"nilNotSys":  nilUser.IsNotSystemUser(),
		"virtual":    u.IsVirtualUser(),
		"sysVirtual": sys.IsVirtualUser(),
	}
	expected := args.Map{
		"notSys":     true,
		"isSys":      false,
		"nilNotSys":  false,
		"virtual":    true,
		"sysVirtual": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNotSystemUser/IsVirtualUser returns correct value -- with args", actual)
}

func Test_Cov_User_HasType_IsTypeEmpty(t *testing.T) {
	u := &corepayload.User{Type: "admin"}
	empty := &corepayload.User{}
	var nilUser *corepayload.User

	actual := args.Map{
		"hasType":     u.HasType(),
		"emptyType":   empty.HasType(),
		"nilType":     nilUser.HasType(),
		"typeEmpty":   u.IsTypeEmpty(),
		"nilTypeE":    nilUser.IsTypeEmpty(),
	}
	expected := args.Map{
		"hasType":     true,
		"emptyType":   false,
		"nilType":     false,
		"typeEmpty":   false,
		"nilTypeE":    true,
	}
	expected.ShouldBeEqual(t, 0, "HasType/IsTypeEmpty returns empty -- with args", actual)
}

func Test_Cov_User_String_Json(t *testing.T) {
	u := corepayload.User{Name: "test"}
	actual := args.Map{
		"strNotEmpty":  u.String() != "",
		"jsonNotEmpty": u.JsonPtr().JsonString() != "",
	}
	expected := args.Map{
		"strNotEmpty":  true,
		"jsonNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "String/Json returns correct value -- with args", actual)
}

func Test_Cov_User_PrettyJsonString(t *testing.T) {
	u := &corepayload.User{Name: "test"}
	actual := args.Map{"notEmpty": u.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString returns correct value -- with args", actual)
}

func Test_Cov_User_Serialize_Deserialize(t *testing.T) {
	u := &corepayload.User{Name: "test", Type: "admin"}
	bytes, err := u.Serialize()

	u2 := &corepayload.User{}
	err2 := u2.Deserialize(bytes)

	actual := args.Map{
		"serErr":  err == nil,
		"desErr":  err2 == nil,
		"name":    u2.Name,
		"type":    u2.Type,
	}
	expected := args.Map{
		"serErr":  true,
		"desErr":  true,
		"name":    "test",
		"type":    "admin",
	}
	expected.ShouldBeEqual(t, 0, "Serialize/Deserialize returns correct value -- with args", actual)
}

func Test_Cov_User_Clone(t *testing.T) {
	u := corepayload.User{Name: "test", Type: "admin", AuthToken: "t", PasswordHash: "h", IsSystemUser: true}
	c := u.Clone()
	actual := args.Map{
		"name": c.Name, "type": c.Type,
		"auth": c.AuthToken, "hash": c.PasswordHash,
		"sys": c.IsSystemUser,
	}
	expected := args.Map{
		"name": "test", "type": "admin",
		"auth": "t", "hash": "h",
		"sys": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_Cov_User_ClonePtr(t *testing.T) {
	u := &corepayload.User{Name: "test"}
	c := u.ClonePtr()
	var nilUser *corepayload.User
	nilClone := nilUser.ClonePtr()

	actual := args.Map{
		"name":     c.Name,
		"nilClone": nilClone == nil,
	}
	expected := args.Map{
		"name":     "test",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)
}

func Test_Cov_User_Ptr(t *testing.T) {
	u := corepayload.User{Name: "test"}
	p := u.Ptr()
	actual := args.Map{"notNil": p != nil, "name": p.Name}
	expected := args.Map{"notNil": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Ptr returns correct value -- with args", actual)
}

// ── newUserCreator ──

func Test_Cov_NewUser_All(t *testing.T) {
	u := corepayload.New.User.All(true, "1", "test", "admin", "tok", "hash")
	actual := args.Map{
		"id": u.Identifier, "name": u.Name, "type": u.Type,
		"auth": u.AuthToken, "hash": u.PasswordHash, "sys": u.IsSystemUser,
	}
	expected := args.Map{
		"id": "1", "name": "test", "type": "admin",
		"auth": "tok", "hash": "hash", "sys": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUser.All returns correct value -- with args", actual)
}

func Test_Cov_NewUser_Create(t *testing.T) {
	u := corepayload.New.User.Create(true, "test", "admin")
	actual := args.Map{"name": u.Name, "sys": u.IsSystemUser}
	expected := args.Map{"name": "test", "sys": true}
	expected.ShouldBeEqual(t, 0, "NewUser.Create returns correct value -- with args", actual)
}

func Test_Cov_NewUser_NonSys(t *testing.T) {
	u := corepayload.New.User.NonSysCreate("test", "admin")
	actual := args.Map{"name": u.Name, "sys": u.IsSystemUser}
	expected := args.Map{"name": "test", "sys": false}
	expected.ShouldBeEqual(t, 0, "NewUser.NonSysCreate returns correct value -- with args", actual)
}

func Test_Cov_NewUser_NonSysId(t *testing.T) {
	u := corepayload.New.User.NonSysCreateId("1", "test", "admin")
	actual := args.Map{"id": u.Identifier, "name": u.Name}
	expected := args.Map{"id": "1", "name": "test"}
	expected.ShouldBeEqual(t, 0, "NewUser.NonSysCreateId returns correct value -- with args", actual)
}

func Test_Cov_NewUser_System(t *testing.T) {
	u := corepayload.New.User.System("test", "admin")
	actual := args.Map{"sys": u.IsSystemUser}
	expected := args.Map{"sys": true}
	expected.ShouldBeEqual(t, 0, "NewUser.System returns correct value -- with args", actual)
}

func Test_Cov_NewUser_SystemId(t *testing.T) {
	u := corepayload.New.User.SystemId("1", "test", "admin")
	actual := args.Map{"id": u.Identifier, "sys": u.IsSystemUser}
	expected := args.Map{"id": "1", "sys": true}
	expected.ShouldBeEqual(t, 0, "NewUser.SystemId returns correct value -- with args", actual)
}

func Test_Cov_NewUser_UsingName(t *testing.T) {
	u := corepayload.New.User.UsingName("test")
	actual := args.Map{"name": u.Name}
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "NewUser.UsingName returns correct value -- with args", actual)
}

func Test_Cov_NewUser_Empty(t *testing.T) {
	u := corepayload.New.User.Empty()
	actual := args.Map{"notNil": u != nil, "nameEmpty": u.Name == ""}
	expected := args.Map{"notNil": true, "nameEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewUser.Empty returns empty -- with args", actual)
}

func Test_Cov_NewUser_Deserialize(t *testing.T) {
	original := &corepayload.User{Name: "test", Type: "admin"}
	bytes, _ := original.Serialize()
	u, err := corepayload.New.User.Deserialize(bytes)
	actual := args.Map{"err": err == nil, "name": u.Name}
	expected := args.Map{"err": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "NewUser.Deserialize returns correct value -- with args", actual)
}

func Test_Cov_NewUser_Deserialize_Error(t *testing.T) {
	_, err := corepayload.New.User.Deserialize([]byte("invalid"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewUser.Deserialize returns error -- error", actual)
}

func Test_Cov_NewUser_CastOrDeserializeFrom(t *testing.T) {
	_, err := corepayload.New.User.CastOrDeserializeFrom(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns nil -- nil", actual)
}

func Test_Cov_NewUser_CastOrDeserializeFrom_Valid(t *testing.T) {
	src := map[string]any{"Name": "test"}
	u, err := corepayload.New.User.CastOrDeserializeFrom(src)
	actual := args.Map{"hasErr": err != nil, "notNil": u != nil}
	expected := args.Map{"hasErr": false, "notNil": true}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns non-empty -- valid", actual)
}
