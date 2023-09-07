package stringslice

import "gitlab.com/auk-go/core/constants"

func MakeDefaultPtr(capacity int) *[]string {
	slice := make([]string, constants.Zero, capacity)

	return &slice
}
