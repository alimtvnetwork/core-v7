package coreinstruction

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type BaseTypeDotFilter struct {
	splitDotFilters []string
	TypeDotFilter   string `json:"TypeDotFilter"`
}

func (receiver BaseTypeDotFilter) GetDotSplitTypes() []string {
	if receiver.splitDotFilters != nil {
		return receiver.splitDotFilters
	}

	receiver.splitDotFilters = strings.Split(
		receiver.TypeDotFilter,
		constants.Dot)

	return receiver.splitDotFilters
}
