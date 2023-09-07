package stringslice

import "gitlab.com/auk-go/core/constants"

type FirstLastStatus struct {
	First, Last       string
	IsValid           bool
	HasFirst, HasLast bool
}

func InvalidFirstLastStatus() *FirstLastStatus {
	return &FirstLastStatus{
		First:    constants.EmptyString,
		Last:     constants.EmptyString,
		IsValid:  false,
		HasFirst: false,
		HasLast:  false,
	}
}

func InvalidFirstLastStatusForLast(first string) *FirstLastStatus {
	return &FirstLastStatus{
		First:    first,
		Last:     constants.EmptyString,
		IsValid:  false,
		HasFirst: true,
		HasLast:  false,
	}
}
