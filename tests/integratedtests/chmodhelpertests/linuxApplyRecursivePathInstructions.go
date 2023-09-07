package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
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
