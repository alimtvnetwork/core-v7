package pagingutil

import (
	"math"

	"gitlab.com/evatix-go/core/coreinterface"
)

func GetPagesSizeUsingLengthGetter(
	totalLengthGetter coreinterface.LengthGetter,
	eachPageSize int,
) int {
	pagesPossibleFloat := float64(totalLengthGetter.Length()) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}
