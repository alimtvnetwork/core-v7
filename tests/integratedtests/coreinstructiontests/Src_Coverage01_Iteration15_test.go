package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coreinstruction"
	"regexp"
	"testing"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsContinueOnError — IsExitOnError nil receiver branch (line 7)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_BaseIsContinueOnError_IsExitOnError_Nil(t *testing.T) {
	var b *coreinstruction.BaseIsContinueOnError
	actual := args.Map{"result": b.IsExitOnError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil receiver", actual)
}

func Test_I15_BaseIsContinueOnError_IsExitOnError_ContinueTrue(t *testing.T) {
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: true}
	actual := args.Map{"result": b.IsExitOnError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false when IsContinueOnError is true", actual)
}

func Test_I15_BaseIsContinueOnError_IsExitOnError_ContinueFalse(t *testing.T) {
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: false}
	actual := args.Map{"result": b.IsExitOnError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true when IsContinueOnError is false", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsSecure — NewSecure, NewPlain, IsPlainText, IsIncludePayload (lines 7-25)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_BaseIsSecure_NewSecure(t *testing.T) {
	s := coreinstruction.NewSecure()
	actual := args.Map{"result": s.IsSecure}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected secure", actual)
	actual := args.Map{"result": s.IsPlainText()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not plain text", actual)
	actual := args.Map{"result": s.IsIncludePayload()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not include payload", actual)
}

func Test_I15_BaseIsSecure_NewPlain(t *testing.T) {
	s := coreinstruction.NewPlain()
	actual := args.Map{"result": s.IsSecure}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not secure", actual)
	actual := args.Map{"result": s.IsPlainText()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected plain text", actual)
	actual := args.Map{"result": s.IsIncludePayload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected include payload", actual)
}

func Test_I15_BaseIsSecure_NilReceiver(t *testing.T) {
	var s *coreinstruction.BaseIsSecure
	actual := args.Map{"result": s.IsPlainText()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected plain text for nil", actual)
	actual := args.Map{"result": s.IsIncludePayload()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected include payload for nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseTags — NewTagsPtr, TagsLength nil, TagsHashset, IsAnyTagMatchesRegex
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_BaseTags_NewTagsPtr_Empty(t *testing.T) {
	bt := coreinstruction.NewTagsPtr([]string{})
	actual := args.Map{"result": bt == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": bt.TagsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 tags", actual)
}

func Test_I15_BaseTags_NewTagsPtr_NonEmpty(t *testing.T) {
	bt := coreinstruction.NewTagsPtr([]string{"a", "b"})
	actual := args.Map{"result": bt.TagsLength() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 tags", actual)
}

func Test_I15_BaseTags_TagsLength_NilTags(t *testing.T) {
	bt := coreinstruction.BaseTags{Tags: nil}
	actual := args.Map{"result": bt.TagsLength() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_I15_BaseTags_TagsHashset_Cached(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"x", "y"})
	h1 := bt.TagsHashset()
	h2 := bt.TagsHashset()
	actual := args.Map{"result": h1 != h2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same pointer (cached)", actual)
}

func Test_I15_BaseTags_HasAllTags(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"a", "b", "c"})
	actual := args.Map{"result": bt.HasAllTags("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": bt.HasAllTags("a", "z")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// Empty tags => true
	actual := args.Map{"result": bt.HasAllTags()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
}

func Test_I15_BaseTags_HasAnyTags(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"a", "b"})
	actual := args.Map{"result": bt.HasAnyTags("a", "z")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": bt.HasAnyTags("x", "z")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_Empty(t *testing.T) {
	bt := coreinstruction.NewTags([]string{})
	r := regexp.MustCompile(`.*`)
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty tags", actual)
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_Match(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"hello-world", "foo"})
	r := regexp.MustCompile(`^hello`)
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_NoMatch(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"foo", "bar"})
	r := regexp.MustCompile(`^hello`)
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}
