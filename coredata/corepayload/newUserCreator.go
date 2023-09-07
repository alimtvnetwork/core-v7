package corepayload

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/isany"
)

type newUserCreator struct{}

func (it newUserCreator) Deserialize(
	rawJsonBytes []byte,
) (*User, error) {
	user := &User{}

	err := corejson.
		Deserialize.
		UsingBytes(
			rawJsonBytes, user)

	if err == nil {
		return user, nil
	}

	// has error
	return nil, err
}

func (it newUserCreator) CastOrDeserializeFrom(
	anyItem interface{},
) (*User, error) {
	if isany.Null(anyItem) {
		return nil, errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs(
				"given any item is nil failed to convert to user")
	}

	toUser := &User{}
	err := corejson.CastAny.FromToDefault(
		anyItem,
		toUser)

	return toUser, err
}

func (it newUserCreator) Empty() *User {
	return &User{}
}

func (it newUserCreator) Create(
	isSystemUser bool,
	name, userType string,
) *User {
	return &User{
		Name:         name,
		Type:         userType,
		IsSystemUser: isSystemUser,
	}
}

func (it newUserCreator) NonSysCreate(
	name, userType string,
) *User {
	return &User{
		Name: name,
		Type: userType,
	}
}

func (it newUserCreator) NonSysCreateId(
	id, name, userType string,
) *User {
	return &User{
		Identifier: id,
		Name:       name,
		Type:       userType,
	}
}

func (it newUserCreator) System(
	name, userType string,
) *User {
	return &User{
		Name:         name,
		Type:         userType,
		IsSystemUser: true,
	}
}

func (it newUserCreator) SystemId(
	id, name, userType string,
) *User {
	return &User{
		Identifier:   id,
		Name:         name,
		Type:         userType,
		IsSystemUser: true,
	}
}

func (it newUserCreator) UsingNameTypeStringer(
	name string,
	userTypeStringer fmt.Stringer,
) *User {
	return &User{
		Name: name,
		Type: userTypeStringer.String(),
	}
}

func (it newUserCreator) SysUsingNameTypeStringer(
	name string,
	userTypeStringer fmt.Stringer,
) *User {
	return &User{
		Name:         name,
		Type:         userTypeStringer.String(),
		IsSystemUser: true,
	}
}

func (it newUserCreator) UsingName(
	name string,
) *User {
	return &User{
		Name: name,
	}
}

func (it newUserCreator) All(
	isSystemUser bool,
	id, name, userType, authToken, passHash string,
) *User {
	return &User{
		Identifier:   id,
		Name:         name,
		Type:         userType,
		AuthToken:    authToken,
		PasswordHash: passHash,
		IsSystemUser: isSystemUser,
	}
}

func (it newUserCreator) AllTypeStringer(
	isSystemUser bool,
	id, name string,
	userType fmt.Stringer,
	authToken, passHash string,
) *User {
	return &User{
		Identifier:   id,
		Name:         name,
		Type:         userType.String(),
		AuthToken:    authToken,
		PasswordHash: passHash,
		IsSystemUser: isSystemUser,
	}
}

func (it newUserCreator) AllUsingStringer(
	isSystemUser bool,
	id, name string,
	typeStringer fmt.Stringer,
	authToken, passHash string,
) *User {
	return &User{
		Identifier:   id,
		Name:         name,
		Type:         typeStringer.String(),
		AuthToken:    authToken,
		PasswordHash: passHash,
		IsSystemUser: isSystemUser,
	}
}
