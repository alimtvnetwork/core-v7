package corepayload

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corejson"
)

type User struct {
	Identifier   string `json:"Identifier,omitempty"`
	Name         string `json:"Name,omitempty"`
	Type         string `json:"Type,omitempty"`
	AuthToken    string `json:"AuthToken,omitempty"`
	PasswordHash string `json:"PasswordHash,omitempty"`
	IsSystemUser bool   `json:"IsSystemUser"`
}

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it *User) IdentifierInteger() int {
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
func (it *User) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
}

func (it *User) HasAuthToken() bool {
	return it != nil && it.AuthToken != ""
}

func (it *User) HasPasswordHash() bool {
	return it != nil && it.PasswordHash != ""
}

func (it *User) IsPasswordHashEmpty() bool {
	return it == nil || it.PasswordHash == ""
}

func (it *User) IsAuthTokenEmpty() bool {
	return it == nil || it.AuthToken == ""
}

func (it *User) IsEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *User) IsValidUser() bool {
	return !it.IsEmpty()
}

func (it *User) IsNameEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *User) IsNameEqual(name string) bool {
	return it != nil && it.Name == name
}

func (it *User) IsNotSystemUser() bool {
	return it != nil && !it.IsSystemUser
}

func (it *User) IsVirtualUser() bool {
	return it != nil && !it.IsSystemUser
}

func (it *User) HasType() bool {
	return it != nil && it.Type != ""
}

func (it *User) IsTypeEmpty() bool {
	return it == nil || it.Type == ""
}

func (it User) String() string {
	return it.JsonPtr().JsonString()
}

func (it *User) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it *User) Json() corejson.Result {
	return corejson.New(it)
}

func (it *User) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *User) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *User) Deserialize(rawJsonBytes []byte) error {
	return corejson.
		Deserialize.
		UsingBytes(rawJsonBytes, it)
}

func (it User) Clone() User {
	return User{
		Name:         it.Name,
		Type:         it.Type,
		AuthToken:    it.AuthToken,
		PasswordHash: it.PasswordHash,
		IsSystemUser: it.IsSystemUser,
	}
}

func (it *User) ClonePtr() *User {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it User) Ptr() *User {
	return &it
}
