package chmodhelpertests

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func linuxApplyRecursivePathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		testCase.RwxInstructions)

	errcore.SimpleHandleErr(err, "linuxApplyRecursivePathInstructions")

	for _, createPath := range testCase.CreatePaths {
		err2 := executors.ApplyOnPath(createPath.Dir)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
