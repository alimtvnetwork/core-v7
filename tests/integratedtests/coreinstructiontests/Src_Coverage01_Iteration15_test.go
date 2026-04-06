package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coreinstruction"
	"regexp"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsContinueOnError — IsExitOnError nil receiver branch (line 7)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_BaseIsContinueOnError_IsExitOnError_Nil(t *testing.T) {
	var b *coreinstruction.BaseIsContinueOnError
	if b.IsExitOnError() {
		t.Fatal("expected false for nil receiver")
	}
}

func Test_I15_BaseIsContinueOnError_IsExitOnError_ContinueTrue(t *testing.T) {
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: true}
	if b.IsExitOnError() {
		t.Fatal("expected false when IsContinueOnError is true")
	}
}

func Test_I15_BaseIsContinueOnError_IsExitOnError_ContinueFalse(t *testing.T) {
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: false}
	if !b.IsExitOnError() {
		t.Fatal("expected true when IsContinueOnError is false")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsSecure — NewSecure, NewPlain, IsPlainText, IsIncludePayload (lines 7-25)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_BaseIsSecure_NewSecure(t *testing.T) {
	s := coreinstruction.NewSecure()
	if !s.IsSecure {
		t.Fatal("expected secure")
	}
	if s.IsPlainText() {
		t.Fatal("expected not plain text")
	}
	if s.IsIncludePayload() {
		t.Fatal("expected not include payload")
	}
}

func Test_I15_BaseIsSecure_NewPlain(t *testing.T) {
	s := coreinstruction.NewPlain()
	if s.IsSecure {
		t.Fatal("expected not secure")
	}
	if !s.IsPlainText() {
		t.Fatal("expected plain text")
	}
	if !s.IsIncludePayload() {
		t.Fatal("expected include payload")
	}
}

func Test_I15_BaseTags_TagsHashset_Cached(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"x", "y"})
	h1 := bt.TagsHashset()
	h2 := bt.TagsHashset()
	if h1 != h2 {
		t.Fatal("expected same pointer (cached)")
	}
}

func Test_I15_BaseTags_HasAllTags(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"a", "b", "c"})
	if !bt.HasAllTags("a", "b") {
		t.Fatal("expected true")
	}
	if bt.HasAllTags("a", "z") {
		t.Fatal("expected false")
	}
	// Empty tags => true
	if !bt.HasAllTags() {
		t.Fatal("expected true for empty")
	}
}

func Test_I15_BaseTags_HasAnyTags(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"a", "b"})
	if !bt.HasAnyTags("a", "z") {
		t.Fatal("expected true")
	}
	if bt.HasAnyTags("x", "z") {
		t.Fatal("expected false")
	}
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_Empty(t *testing.T) {
	bt := coreinstruction.NewTags([]string{})
	r := regexp.MustCompile(`.*`)
	if bt.IsAnyTagMatchesRegex(r) {
		t.Fatal("expected false for empty tags")
	}
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_Match(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"hello-world", "foo"})
	r := regexp.MustCompile(`^hello`)
	if !bt.IsAnyTagMatchesRegex(r) {
		t.Fatal("expected match")
	}
}

func Test_I15_BaseTags_IsAnyTagMatchesRegex_NoMatch(t *testing.T) {
	bt := coreinstruction.NewTags([]string{"foo", "bar"})
	r := regexp.MustCompile(`^hello`)
	if bt.IsAnyTagMatchesRegex(r) {
		t.Fatal("expected no match")
	}
}
