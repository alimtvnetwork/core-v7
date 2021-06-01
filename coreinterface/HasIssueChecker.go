package coreinterface

type HasIssueChecker interface {
	HasAnyItemChecker
	HasIssues() bool
	IsEmptyOrIssues() bool
	// HasValidItems Has items and there is no issues
	HasValidItems() bool
}
