package coreinterface

type StringHasAnyChecker interface {
	HasAny(searchTerms ...string) bool
}
