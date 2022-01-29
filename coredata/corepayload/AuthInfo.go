package corepayload

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type AuthInfo struct {
	Identifier   string       `json:"Identifier,omitempty"`
	ActionType   string       `json:"ActionType,omitempty"`
	ResourceName string       `json:"ResourceName,omitempty"` // can be url or any name
	SessionInfo  *SessionInfo `json:"SessionInfo,omitempty"`
	UserInfo     *UserInfo    `json:"UserInfo,omitempty"`
}

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it *AuthInfo) IdentifierInteger() int {
	if it.Identifier == "" {
		return constants.InvalidValue
	}

	idInt, _ := converters.StringToIntegerWithDefault(
		it.Identifier,
		constants.InvalidValue)

	return idInt
}

// IdentifierUnsignedInteger
//
// Invalid value returns constants.Zero
func (it *AuthInfo) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
}

func (it *AuthInfo) IsEmpty() bool {
	return it == nil ||
		it.ActionType == "" &&
			it.ResourceName == "" &&
			it.SessionInfo.IsEmpty() &&
			it.UserInfo.IsEmpty()
}

func (it *AuthInfo) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *AuthInfo) IsActionTypeEmpty() bool {
	return it == nil ||
		it.ActionType == ""
}

func (it *AuthInfo) IsResourceNameEmpty() bool {
	return it == nil ||
		it.ResourceName == ""
}

func (it *AuthInfo) IsValid() bool {
	return !it.IsEmpty()
}

func (it *AuthInfo) HasActionType() bool {
	return it != nil &&
		it.ActionType != ""
}

func (it *AuthInfo) HasResourceName() bool {
	return it != nil &&
		it.ResourceName != ""
}

func (it *AuthInfo) IsUserInfoEmpty() bool {
	return it == nil ||
		it.UserInfo.IsEmpty()
}

func (it *AuthInfo) IsSessionInfoEmpty() bool {
	return it == nil ||
		it.SessionInfo.IsEmpty()
}

func (it *AuthInfo) HasUserInfo() bool {
	return !it.IsUserInfoEmpty()
}

func (it *AuthInfo) HasSessionInfo() bool {
	return !it.IsSessionInfoEmpty()
}

func (it AuthInfo) String() string {
	return it.Json().JsonString()
}

func (it AuthInfo) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it AuthInfo) Json() corejson.Result {
	return corejson.New(it)
}

func (it AuthInfo) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *AuthInfo) Clone() AuthInfo {
	return AuthInfo{
		ActionType:   it.ActionType,
		ResourceName: it.ResourceName,
		SessionInfo:  it.SessionInfo.ClonePtr(),
		UserInfo:     it.UserInfo.ClonePtr(),
	}
}

func (it *AuthInfo) ClonePtr() *AuthInfo {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it AuthInfo) Ptr() *AuthInfo {
	return &it
}
