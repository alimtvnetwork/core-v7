package coreapi

import "gitlab.com/auk-go/core/constants"

func InvalidGenericResponseResult(attribute *ResponseAttribute) *GenericResponseResult {
	if attribute == nil {
		return &GenericResponseResult{
			Attribute: InvalidResponseAttribute(constants.EmptyString),
			Response:  nil,
		}
	}

	return &GenericResponseResult{
		Attribute: attribute,
		Response:  nil,
	}
}
