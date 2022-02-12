package loggerinf

import "gitlab.com/evatix-go/core/internal/internalinterface"

type PathExtender interface {
	internalinterface.PathExtender
	ParentDirExtender() PathExtender
	RootExtender() PathExtender
}
