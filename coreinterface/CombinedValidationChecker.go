package coreinterface

type CombinedValidationChecker interface {
	HasErrorChecker
	HasAnyItemChecker
	HasAnyIssueChecker
	ValidationErrorGetter
	ErrorGetter
}
