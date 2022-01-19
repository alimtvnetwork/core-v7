package isany

import "gitlab.com/evatix-go/core/internal/reflectinternal"

func NotDeepEqual(
	left, right interface{},
) (isNotEqual bool) {
	return !reflectinternal.
		IsAnyEqual(left, right)
}
