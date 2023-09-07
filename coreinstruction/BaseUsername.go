package coreinstruction

import (
	"regexp"
	"strings"

	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type BaseUsername struct {
	Username string `json:"Username"`
}

func NewUsername(user string) *BaseUsername {
	return &BaseUsername{Username: user}
}

func (it *BaseUsername) UsernameString() string {
	return it.Username
}

func (it *BaseUsername) IsUsernameEmpty() bool {
	return it == nil || it.Username == ""
}

func (it *BaseUsername) IsUsernameWhitespace() bool {
	return it == nil || strutilinternal.IsNullOrEmptyOrWhitespace(&it.Username)
}

func (it *BaseUsername) IsUsername(user string) bool {
	return it.Username == user
}

func (it *BaseUsername) IsUsernameCaseInsensitive(usernameInsensitive string) bool {
	return strings.EqualFold(it.Username, usernameInsensitive)
}

func (it *BaseUsername) IsUsernameContains(usernameContains string) bool {
	return strings.Contains(it.Username, usernameContains)

}

func (it *BaseUsername) IsUsernameRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(it.Username)
}

func (it *BaseUsername) IsEqual(right *BaseUsername) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	return it.Username == right.Username
}

func (it *BaseUsername) IsNotEqual(right *BaseUsername) bool {
	return !it.IsEqual(right)
}

func (it *BaseUsername) ClonePtr() *BaseUsername {
	if it == nil {
		return nil
	}

	return &BaseUsername{
		Username: it.Username,
	}
}

func (it BaseUsername) Clone() BaseUsername {
	return BaseUsername{
		Username: it.Username,
	}
}
