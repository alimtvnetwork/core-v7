package chmodhelper

import (
	"os"
	"sync"
)

var (
	SimpleFileWriter = simpleFileWriter{}
	New              = newCreator{}
	ChmodApply       = chmodApplier{}
	ChmodVerify      = chmodVerifier{}
	globalMutex      = sync.Mutex{}
	TempDirDefault   = os.TempDir()
	TempDirGetter    = tempDirGetter{}
)
