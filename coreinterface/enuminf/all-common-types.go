package enuminf

import (
	"gitlab.com/auk-go/core/internal/internalinterface"
	"gitlab.com/auk-go/core/internal/internalinterface/internalenuminf"
)

type EnvironmentFlagTyper interface {
	BasicEnumer
	internalenuminf.EnvironmentFlagTyper
}

type EnvironmentTyper interface {
	BasicEnumer
	internalenuminf.EnvironmentTyper
}

type EnvironmentOptioner interface {
	EnvTyper() EnvironmentTyper
	FlagTyper() EnvironmentFlagTyper
}

type CrudTyper interface {
	BasicEnumer
	internalenuminf.CrudTyper
}

type LinuxTyper interface {
	BasicEnumer
	internalenuminf.LinuxTyper
}

type LoggerTyper interface {
	BasicEnumer
	internalinterface.LoggerTyper
}

type LogLevelTyper interface {
	BasicEnumer
	internalenuminf.LogLevelTyper
}

type CompareMethodsTyper interface {
	BasicEnumer
	internalenuminf.CompareMethodsTyper
}

type OnOffLater interface {
	BasicEnumer
	internalenuminf.OnOffLater
}

type PrivilegeTyper interface {
	BasicEnumer
	internalenuminf.PrivilegeTyper
}

type OverwriteOrRideOrEnforcer interface {
	BasicEnumer
	internalenuminf.OverwriteOrRideOrEnforcer
}

type HttpMethodTyper interface {
	BasicEnumer
	internalenuminf.HttpMethodTyper
}

type ActionTyper interface {
	BasicEnumer
	internalenuminf.ActionTyper
}

type CompletionStateTyper interface {
	BasicEnumer
	internalenuminf.CompletionStateTyper
}

type EventTyper interface {
	BasicEnumer
	internalenuminf.EventTyper
}

type StringCompareTyper interface {
	BasicEnumer
	internalenuminf.StringCompareTyper
}
