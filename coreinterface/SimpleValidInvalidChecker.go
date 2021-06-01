package coreinterface

type SimpleValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
	InvalidMessageGetter
}
