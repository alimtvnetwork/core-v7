package baseactioninf

import (
	"fmt"
)

type SimpleExecutor interface {
	Namer

	CategoryTypeNamer

	Executor
	IsApply() (isSuccess bool)

	fmt.Stringer
}

type StandardExecutor interface {
	SimpleExecutor

	Initializer
	DefaultsInjector
	ValidationErrorGetter
	Strings() []string
}

type SimpleResulter interface {
	Namer

	CategoryTypeNamer

	Exec() ([]byte, error)
	ExecAny() (anyItem interface{}, err error)

	IsApply() (isSuccess bool)
}

type StandardResulter interface {
	SimpleResulter

	Initializer
	DefaultsInjector
	ValidationErrorGetter

	Strings() []string
}

type ConditionStandardResulter interface {
	StandardResulter

	IsExecutableChecker
	IsExecutableUsingMapChecker
}

type ConditionSimpleResulter interface {
	SimpleResulter

	IsExecutableChecker
	IsExecutableUsingMapChecker
}

type ConditionSimpleExecutor interface {
	SimpleExecutor

	IsExecutableChecker
	IsExecutableUsingMapChecker
}

type ConditionStandardExecutor interface {
	StandardExecutor

	IsExecutableChecker
	IsExecutableUsingMapChecker
}
