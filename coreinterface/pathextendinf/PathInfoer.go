package pathextendinf

import "gitlab.com/auk-go/core/internal/internalinterface/internalpathextender"

type PathInfoer interface {
	IsPathChecker
	internalpathextender.PathInfoer
}
