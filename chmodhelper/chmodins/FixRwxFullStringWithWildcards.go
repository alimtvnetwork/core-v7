package chmodins

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

// FixRwxFullStringWithWildcards can be less than 10 and can be
//
//  - "rwx" will be filled with "-rwx******"
//  - "rw-" will be filled with "-rw-******"
//  - "-rw-" will be filled with "-rw-******"
func FixRwxFullStringWithWildcards(rwxFull string) (fixedRwx string) {
	if len(rwxFull) == RwxFullLength {
		return rwxFull
	}

	if len(rwxFull) == constants.Zero {
		return AllWildCardsRwxFullString
	}

	if rwxFull[0] != constants.HyphenChar {
		rwxFull = constants.Hyphen + rwxFull
	}

	return strutilinternal.MaskLine(
		AllWildCardsRwxFullString,
		rwxFull)
}
