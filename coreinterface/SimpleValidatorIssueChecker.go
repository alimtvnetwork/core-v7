package coreinterface

type SimpleValidatorIssueChecker interface {
	SimpleValidInvalidChecker
	HasAnyItemChecker
	InvalidDirectErrorGetter
}
