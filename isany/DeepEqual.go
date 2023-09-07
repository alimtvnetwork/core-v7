package isany

import (
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func DeepEqual(
	left, right interface{},
) (isEqual bool) {
	return reflectinternal.
		IsAnyEqual(left, right)
}
