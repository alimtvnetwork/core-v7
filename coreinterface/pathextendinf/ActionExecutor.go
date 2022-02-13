package pathextendinf

import "gitlab.com/evatix-go/core/coreinterface/errcoreinf"

type ActionExecutor interface {
	HasAnyAction() bool
	IsEmptyActions() bool
	Exec() errcoreinf.BasicErrWrapper
	ExecAll() errcoreinf.BaseErrorWrapperCollectionDefiner
}
