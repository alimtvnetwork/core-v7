package isany

import "gitlab.com/auk-go/core/internal/reflectinternal"

func NotDeepEqual(
	left, right interface{},
) (isNotEqual bool) {
	return !reflectinternal.Is.AnyEqual(left, right)
}
