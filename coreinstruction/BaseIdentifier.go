package coreinstruction

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

type BaseIdentifier struct {
	Id string `json:"Id"`
}

func (identifier BaseIdentifier) IdString() string {
	return identifier.Id
}

func (identifier BaseIdentifier) IsIdEmpty() bool {
	return identifier.Id == ""
}

func (identifier BaseIdentifier) IsIdWhitespace() bool {
	return strutilinternal.IsNullOrEmptyOrWhitespace(&identifier.Id)
}

func (identifier BaseIdentifier) IsId(id string) bool {
	return identifier.Id == id
}

func (identifier BaseIdentifier) IsIdCaseInsensitive(idInsensitive string) bool {
	return strings.EqualFold(identifier.Id, idInsensitive)
}

func (identifier BaseIdentifier) IsIdContains(idContains string) bool {
	return strings.Contains(identifier.Id, idContains)

}

func (identifier BaseIdentifier) IsIdRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(identifier.Id)
}
