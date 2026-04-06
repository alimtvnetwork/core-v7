package coreinstruction

import (
	"regexp"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

type BaseTags struct {
	tagsHashset *corestr.Hashset
	Tags        []string `json:"Tags,omitempty"`
}

func NewTagsPtr(tags []string) *BaseTags {
	if len(tags) == 0 {
		return NewTags(nil)
	}

	return NewTags(tags)
}

func NewTags(tags []string) *BaseTags {
	if len(tags) == 0 {
		return &BaseTags{
			Tags: []string{},
		}
	}

	return &BaseTags{
		Tags: tags,
	}
}

func (it BaseTags) TagsLength() int {
	if it.Tags == nil {
		return constants.Zero
	}

	return len(it.Tags)
}

func (it BaseTags) IsTagsEmpty() bool {
	return it.TagsLength() == 0
}

func (it *BaseTags) TagsHashset() *corestr.Hashset {
	if it.tagsHashset != nil {
		return it.tagsHashset
	}

	it.tagsHashset = corestr.New.Hashset.Strings(
		it.Tags)

	return it.tagsHashset
}

func (it BaseTags) HasAllTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := it.TagsHashset()

	return hashset.HasAll(tags...)
}

func (it BaseTags) HasAnyTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := it.TagsHashset()

	return hashset.HasAny(tags...)
}

func (it BaseTags) IsAnyTagMatchesRegex(regexp2 *regexp.Regexp) bool {
	if it.IsTagsEmpty() {
		return false
	}

	for _, s := range it.Tags {
		if regexp2.MatchString(s) {
			return true
		}
	}

	return false
}
