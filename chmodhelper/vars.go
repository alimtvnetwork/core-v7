package chmodhelper

import (
	"sync"

	"gitlab.com/auk-go/core/internal/pathinternal"
)

var (
	SimpleFileWriter = simpleFileWriter{}
	New              = newCreator{}
	ChmodApply       = chmodApplier{}
	ChmodVerify      = chmodVerifier{}
	globalMutex      = sync.Mutex{}
	TempDirDefault   = pathinternal.GetTemp() // eg. unix : /tmp, windows: %temp%
	TempDirGetter    = tempDirGetter{}
	newError         = errorCreator{}
)
