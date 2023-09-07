package stringslice

import "gitlab.com/auk-go/core/constants"

func Last(slice []string) string {
	return (slice)[len(slice)-constants.One]
}
