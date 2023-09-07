package pathextendinf

import "gitlab.com/auk-go/core/coreinterface/errcoreinf"

type ActionExecutor interface {
	HasAnyAction() bool
	IsEmptyActions() bool
	Exec() errcoreinf.BasicErrWrapper
	ExecAll() errcoreinf.BaseErrorWrapperCollectionDefiner
}
