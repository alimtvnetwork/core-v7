package corepayload

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
)

type SessionInfo struct {
	Id          string `json:"Id,omitempty"`
	User        *User  `json:"User,omitempty"`
	SessionPath string `json:"SessionPath,omitempty"`
}

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it *SessionInfo) IdentifierInteger() int {
	if it.Id == "" {
		return constants.InvalidValue
	}

	idInt, _ := converters.StringToIntegerWithDefault(
		it.Id,
		constants.InvalidValue)

	return idInt
}

// IdentifierUnsignedInteger
//
// Invalid value returns constants.Zero
func (it *SessionInfo) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
}

func (it *SessionInfo) IsEmpty() bool {
	return it == nil
}

func (it *SessionInfo) IsValid() bool {
	return !it.IsEmpty() && it.Id != ""
}

func (it *SessionInfo) IsUserNameEmpty() bool {
	return it == nil || it.User.IsNameEmpty()
}

func (it *SessionInfo) IsUserEmpty() bool {
	return it == nil || it.User.IsEmpty()
}

func (it *SessionInfo) HasUser() bool {
	return it != nil && it.User.IsValidUser()
}

func (it *SessionInfo) IsUsernameEqual(
	name string,
) bool {
	return it != nil &&
		it.User.IsNameEqual(name)
}

func (it SessionInfo) Clone() SessionInfo {
	return SessionInfo{
		Id:          it.Id,
		User:        it.User.ClonePtr(),
		SessionPath: it.SessionPath,
	}
}

func (it *SessionInfo) ClonePtr() *SessionInfo {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it SessionInfo) Ptr() *SessionInfo {
	return &it
}
