package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Attributes IsEqual ──

func Test_Cov5_Attributes_IsEqual_BothNil(t *testing.T) {
	var a, b *corepayload.Attributes
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- IsEqual both nil", actual)
}

func Test_Cov5_Attributes_IsEqual_OneNil(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"equal": a.IsEqual(nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- IsEqual one nil", actual)
}

func Test_Cov5_Attributes_IsEqual_SamePtr(t *testing.T) {
	a := &corepayload.Attributes{}
	actual := args.Map{"equal": a.IsEqual(a)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- IsEqual same pointer", actual)
}
	a := &corepayload.Attributes{}
	cloned, err := a.Clone(false)
	actual := args.Map{
		"noErr": err == nil,
		// check it doesn't panic
		"dynPayloadsNil": cloned.DynamicPayloads == nil,
	}
	expected := args.Map{"noErr": true, "dynPayloadsNil": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- Clone value", actual)
}
func Test_Cov5_PagingInfo_IsEqual_BothNil(t *testing.T) {
	var a, b *corepayload.PagingInfo
	actual := args.Map{"equal": a.IsEqual(b)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- IsEqual both nil", actual)
}
func Test_Cov5_SessionInfo(t *testing.T) {
	si := corepayload.SessionInfo{Id: "sess-123"}
	actual := args.Map{"id": si.Id}
	expected := args.Map{"id": "sess-123"}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- struct", actual)
}

// ── UserInfo ──

func Test_Cov5_UserInfo(t *testing.T) {
	ui := corepayload.UserInfo{}
	actual := args.Map{"isEmpty": ui.IsEmpty(), "hasUser": ui.HasUser()}
	expected := args.Map{"isEmpty": true, "hasUser": false}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- struct", actual)
}

// ── User ──

func Test_Cov5_User_Nil(t *testing.T) {
	var u *corepayload.User
	actual := args.Map{"isNil": u == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "User returns nil -- nil", actual)
}
