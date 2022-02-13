package pathextendinf

import "gitlab.com/evatix-go/core/internal/internalinterface/internalpathextender"

type PathInfoer interface {
	IsPathChecker
	internalpathextender.PathInfoer
}
