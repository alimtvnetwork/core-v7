package coreinterface

type BasicSlicer interface {
	LengthGetter
	CountGetter
	EmptyChecker
	LastIndexGetter
	HasIndexChecker
	HasAnyItemChecker
	Stringer
}
