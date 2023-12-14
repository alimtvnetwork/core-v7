package corepayload

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corejson"
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

	idInt, _ := converters.StringTo.IntegerWithDefault(
		it.Identifier,
		constants.InvalidValue,
	)

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

func (it *AuthInfo) String() string {
	return it.JsonPtr().JsonString()
}

// SetUserInfo
//
// on null creates new
func (it *AuthInfo) SetUserInfo(
	userInfo *UserInfo,
) *AuthInfo {
	if it == nil {
		return &AuthInfo{
			UserInfo: userInfo,
		}
	}

	it.UserInfo = userInfo

	return it
}

func (it *AuthInfo) SetActionType(
	actionType string,
) *AuthInfo {
	if it == nil {
		return &AuthInfo{
			ActionType: actionType,
		}
	}

	it.ActionType = actionType

	return it
}

func (it *AuthInfo) SetResourceName(
	resourceName string,
) *AuthInfo {
	if it == nil {
		return &AuthInfo{
			ResourceName: resourceName,
		}
	}

	it.ResourceName = resourceName

	return it
}

func (it *AuthInfo) SetIdentifier(
	identifier string,
) *AuthInfo {
	if it == nil {
		return &AuthInfo{
			Identifier: identifier,
		}
	}

	it.Identifier = identifier

	return it
}

func (it *AuthInfo) SetSessionInfo(
	sessionInfo *SessionInfo,
) *AuthInfo {
	if it == nil {
		return &AuthInfo{
			SessionInfo: sessionInfo,
		}
	}

	it.SessionInfo = sessionInfo

	return it
}

func (it *AuthInfo) SetUserSystemUser(
	user, systemUser *User,
) *AuthInfo {
	if it == nil {
		empty := &AuthInfo{}
		empty.UserInfo = &UserInfo{}
		empty.UserInfo.SetUserSystemUser(user, systemUser)

		return empty
	}

	it.UserInfo.SetUserSystemUser(user, systemUser)

	return it
}

// SetUser
//
// on null creates new
func (it *AuthInfo) SetUser(
	user *User,
) *AuthInfo {
	if it == nil {
		empty := &AuthInfo{}
		empty.UserInfo = &UserInfo{}
		empty.UserInfo.SetUser(user)

		return empty
	}

	it.UserInfo.SetUser(user)

	return it
}

// SetSystemUser
//
// on null creates new
func (it *AuthInfo) SetSystemUser(
	systemUser *User,
) *AuthInfo {
	if it == nil {
		empty := &AuthInfo{}
		empty.UserInfo = &UserInfo{}
		empty.UserInfo.SetSystemUser(systemUser)

		return empty
	}

	it.UserInfo.SetSystemUser(systemUser)

	return it
}

func (it *AuthInfo) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it *AuthInfo) Json() corejson.Result {
	return corejson.New(it)
}

func (it *AuthInfo) JsonPtr() *corejson.Result {
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
