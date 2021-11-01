package coreinterface

type IsSuccessValidator interface {
	IsValidChecker
	IsSuccessChecker
	IsFailedChecker
}
