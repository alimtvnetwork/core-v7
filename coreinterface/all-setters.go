package coreinterface

type SetEnabler interface {
	SetEnable()
}

type SetEnablerReturnsError interface {
	SetEnable() error
}

type SetFlager interface {
	SetFlag(flagName string) error
}
