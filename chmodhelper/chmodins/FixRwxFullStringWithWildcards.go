package chmodins

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

// FixRwxFullStringWithWildcards can be less than 10 and can be
//
//  - "-rwx" will be "-rwx******"
//  - "-rwxr-x" will be "-rwxr-x***"
//  - "-rwxr-x" will be "-rwxr-x***"
func FixRwxFullStringWithWildcards(rwxFull string) (fixedRwx string) {
	length := len(rwxFull)

	if length == RwxFullLength {
		return rwxFull
	} else if length == constants.Zero {
		return AllWildCardsRwxFullString
	}

	return strutilinternal.MaskLine(
		AllWildCardsRwxFullString,
		rwxFull)
}
