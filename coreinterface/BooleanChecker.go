package coreinterface

type BooleanChecker interface {
	IsAnyByOrder(booleans ...bool) bool
	HasAll(searchTerms ...string) bool
	HasAny(searchTerms ...string) bool
	HasItemsWithoutIssues() bool
}
