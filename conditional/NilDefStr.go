package conditional

import "gitlab.com/evatix-go/core/constants"

func NilDefStr(
	strPtr *string,
) string {
	if strPtr == nil {
		return constants.EmptyString
	}

	return *strPtr
}

func NilDefStrPtr(
	strPtr *string,
) *string {
	if strPtr == nil {
		return constants.EmptyStringPtr
	}

	return strPtr
}

func NilStr(
	strPtr *string,
	onNil string,
	onNonNil string,
) string {
	if strPtr == nil {
		return onNil
	}

	return onNonNil
}

func NilOrEmptyStr(
	strPtr *string,
	onNilOrEmpty string,
	onNonNilOrNonEmpty string,
) string {
	if strPtr == nil || *strPtr == "" {
		return onNilOrEmpty
	}

	return onNonNilOrNonEmpty
}

func NilOrEmptyStrPtr(
	strPtr *string,
	onNilOrEmpty string,
	onNonNilOrNonEmpty string,
) *string {
	if strPtr == nil || *strPtr == "" {
		return &onNilOrEmpty
	}

	return &onNonNilOrNonEmpty
}
