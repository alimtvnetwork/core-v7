package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Attributes IsEqual ──

func Test_Attributes_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *corepayload.Attributes

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- IsEqual both nil", actual)
}

func Test_Attributes_IsEqual_OneNil(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"equal": a.IsEqual(nil)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- IsEqual one nil", actual)
}

func Test_Attributes_IsEqual_SamePtr(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}

	// Act
	actual := args.Map{"equal": a.IsEqual(a)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- IsEqual same pointer", actual)
}

// ── Attributes Clone ──

func Test_Attributes_Clone_Nil(t *testing.T) {
	// Arrange
	var a *corepayload.Attributes
	cloned, err := a.ClonePtr(false)

	// Act
	actual := args.Map{
		"isNil": cloned == nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns nil -- ClonePtr nil", actual)
}

func Test_Attributes_Clone_Shallow(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	cloned, err := a.ClonePtr(false)

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"noErr":  err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- ClonePtr shallow", actual)
}

func Test_Attributes_Clone_Value(t *testing.T) {
	// Arrange
	a := &corepayload.Attributes{}
	cloned, err := a.Clone(false)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		// check it doesn't panic
		"dynPayloadsNil": cloned.DynamicPayloads == nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"dynPayloadsNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Attributes returns correct value -- Clone value", actual)
}

// ── AuthInfo ──

func Test_AuthInfo_Nil(t *testing.T) {
	// Arrange
	var ai *corepayload.AuthInfo
	cloned := ai.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns nil -- ClonePtr nil", actual)
}

// ── PagingInfo ──

func Test_PagingInfo_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var a, b *corepayload.PagingInfo

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- IsEqual both nil", actual)
}

func Test_PagingInfo_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var pi *corepayload.PagingInfo

	// Act
	actual := args.Map{"isNil": pi.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- ClonePtr nil", actual)
}

// ── SessionInfo ──

func Test_SessionInfo(t *testing.T) {
	// Arrange
	si := corepayload.SessionInfo{Id: "sess-123"}

	// Act
	actual := args.Map{"id": si.Id}

	// Assert
	expected := args.Map{"id": "sess-123"}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- struct", actual)
}

// ── UserInfo ──

func Test_UserInfo(t *testing.T) {
	// Arrange
	ui := corepayload.UserInfo{}

	// Act
	actual := args.Map{
		"isEmpty": ui.IsEmpty(),
		"hasUser": ui.HasUser(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"hasUser": false,
	}
	expected.ShouldBeEqual(t, 0, "UserInfo returns correct value -- struct", actual)
}

// ── User ──

func Test_User_Nil(t *testing.T) {
	// Arrange
	var u *corepayload.User

	// Act
	actual := args.Map{"isNil": u == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "User returns nil -- nil", actual)
}
