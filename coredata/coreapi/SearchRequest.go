package coreapi

type SearchRequest struct {
	SearchTerm                                                     string
	IsNaturalSearch, IsContains, IsStartsWith, IsEndsWith, IsRegex bool
}
