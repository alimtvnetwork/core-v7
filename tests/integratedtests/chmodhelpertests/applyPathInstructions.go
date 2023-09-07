package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func applyPathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		testCase.RwxInstructions)

	errcore.SimpleHandleErr(err, "applyPathInstructions")

	for _, createPath := range testCase.CreatePaths {
		err2 := executors.ApplyOnPathsPtr(createPath.GetPathsPtr())

		if err2 != nil {
			return err2
		}
	}

	return nil
}
