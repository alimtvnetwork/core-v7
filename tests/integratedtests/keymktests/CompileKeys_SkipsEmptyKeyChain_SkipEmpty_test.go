package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage6 — appendStringsWithBaseAnyItems skip-empty branch
//
// Target: keymk/appendStringsWithBaseAnyItems.go:13-14
//   isSkipOnEmpty && item == "" → continue
//
// Exercise via CompileKeys with IsSkipEmptyEntry=true and a sub-key
// that has an empty string in its keyChains.
// ══════════════════════════════════════════════════════════════════════════════

func Test_CompileKeys_SkipsEmptyKeyChain(t *testing.T) {
	// Arrange — sub-key built with IsSkipEmptyEntry=false so "" enters keyChains
	noSkipOption := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
	}
	mainKey := keymk.NewKey.Default("root", "chain1")
	subKey := keymk.NewKey.All(noSkipOption, "sub", "", "val")

	// Act — CompileKeys uses mainKey.option.IsSkipEmptyEntry (true) to filter sub keyChains
	result := mainKey.CompileKeys(subKey)

	// Assert
	convey.Convey("CompileKeys skips empty keyChain entries when IsSkipEmptyEntry is true", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "sub")
		convey.So(result, convey.ShouldContainSubstring, "val")
	})
}
