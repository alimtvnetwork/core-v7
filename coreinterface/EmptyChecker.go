package coreinterface

type EmptyChecker interface {
	IsEmpty() bool
	HasAnyItemChecker
}
