package pagingutil

import (
	"math"
)

// GetPagesSize returns ceiling int for pages possible
func GetPagesSize(
	eachPageSize,
	totalLength int,
) int {
	pagesPossibleFloat := float64(totalLength) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}
