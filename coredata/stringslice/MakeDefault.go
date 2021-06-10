package stringslice

import "gitlab.com/evatix-go/core/constants"

func MakeDefault(capacity int) []string {
	return make([]string, constants.Zero, capacity)
}
