package coreinstructiontests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BaseIsContinueOnError ──

func Test_IsExitOnError_True(t *testing.T) {
	// Arrange
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: false}

	// Act
	result := b.IsExitOnError()

	// Assert
	actual := args.Map{"isExit": result}
	expected := args.Map{"isExit": true}
	expected.ShouldBeEqual(t, 0, "IsExitOnError returns error -- when IsContinueOnError=false", actual)
}

// ── BaseIsSecure ──

func Test_NewSecure(t *testing.T) {
	// Arrange & Act
	s := coreinstruction.NewSecure()

	// Assert
	actual := args.Map{"isSecure": s.IsSecure}
	expected := args.Map{"isSecure": true}
	expected.ShouldBeEqual(t, 0, "NewSecure returns correct value -- with args", actual)
}

func Test_NewPlain(t *testing.T) {
	// Arrange & Act
	s := coreinstruction.NewPlain()

	// Assert
	actual := args.Map{"isSecure": s.IsSecure}
	expected := args.Map{"isSecure": false}
	expected.ShouldBeEqual(t, 0, "NewPlain returns correct value -- with args", actual)
}

func Test_IsPlainText(t *testing.T) {
	// Arrange
	s := &coreinstruction.BaseIsSecure{IsSecure: false}

	// Act
	result := s.IsPlainText()

	// Assert
	actual := args.Map{"isPlain": result}
	expected := args.Map{"isPlain": true}
	expected.ShouldBeEqual(t, 0, "IsPlainText returns correct value -- when not secure", actual)
}

func Test_IsIncludePayload(t *testing.T) {
	// Arrange
	s := &coreinstruction.BaseIsSecure{IsSecure: false}

	// Act
	result := s.IsIncludePayload()

	// Assert
	actual := args.Map{"includePayload": result}
	expected := args.Map{"includePayload": true}
	expected.ShouldBeEqual(t, 0, "IsIncludePayload returns correct value -- when not secure", actual)
}

// ── BaseTags ──

func Test_NewTagsPtr_NonEmpty(t *testing.T) {
	// Arrange & Act
	tags := coreinstruction.NewTagsPtr([]string{"alpha", "beta"})

	// Assert
	actual := args.Map{"length": tags.TagsLength()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "NewTagsPtr returns empty -- non-empty", actual)
}

func Test_NewTagsPtr_Empty(t *testing.T) {
	// Arrange & Act
	tags := coreinstruction.NewTagsPtr([]string{})

	// Assert
	actual := args.Map{"length": tags.TagsLength()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "NewTagsPtr returns empty -- empty", actual)
}

func Test_TagsLength_NilTags(t *testing.T) {
	// Arrange
	bt := coreinstruction.BaseTags{}

	// Act
	result := bt.TagsLength()

	// Assert
	actual := args.Map{"length": result}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "TagsLength returns nil -- nil tags", actual)
}

func Test_TagsHashset_Cached(t *testing.T) {
	// Arrange
	tags := coreinstruction.NewTags([]string{"x", "y"})

	// Act — call twice, second should use cache
	h1 := tags.TagsHashset()
	h2 := tags.TagsHashset()

	// Assert
	actual := args.Map{"samePtr": h1 == h2}
	expected := args.Map{"samePtr": true}
	expected.ShouldBeEqual(t, 0, "TagsHashset returns correct value -- cached", actual)
}

func Test_IsAnyTagMatchesRegex_Match(t *testing.T) {
	// Arrange
	tags := coreinstruction.NewTags([]string{"hello-world", "foo-bar"})
	rx := regexp.MustCompile(`^foo-`)

	// Act
	result := tags.IsAnyTagMatchesRegex(rx)

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "IsAnyTagMatchesRegex returns non-empty -- with match", actual)
}

func Test_IsAnyTagMatchesRegex_NoMatch(t *testing.T) {
	// Arrange
	tags := coreinstruction.NewTags([]string{"hello", "world"})
	rx := regexp.MustCompile(`^zzz`)

	// Act
	result := tags.IsAnyTagMatchesRegex(rx)

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTagMatchesRegex returns empty -- no match", actual)
}

func Test_IsAnyTagMatchesRegex_EmptyTags(t *testing.T) {
	// Arrange
	tags := coreinstruction.NewTags(nil)
	rx := regexp.MustCompile(`.*`)

	// Act
	result := tags.IsAnyTagMatchesRegex(rx)

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTagMatchesRegex returns empty -- empty tags", actual)
}
