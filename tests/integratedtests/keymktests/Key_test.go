package keymktests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

func Test_Key_Compile_Verification(t *testing.T) {
	for caseIndex, testCase := range keyCompileTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		compiled := key.Compile()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, compiled)
	}
}

func Test_Key_AppendChain_Verification(t *testing.T) {
	for caseIndex, testCase := range keyAppendChainTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		initialRaw, _ := input.Get("initial")
		initial := initialRaw.([]string)
		appendRaw, _ := input.Get("append")
		appendItems := appendRaw.([]string)

		// Act
		key := keymk.NewKey.DefaultStrings(main, initial...)
		key.AppendChainStrings(appendItems...)
		compiled := key.Compile()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, compiled)
	}
}

func Test_Key_Finalized_Verification(t *testing.T) {
	for caseIndex, testCase := range keyFinalizedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		key.Finalized()
		compiled := key.CompiledChain()
		isComplete := fmt.Sprintf("%v", key.IsComplete())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			compiled,
			isComplete,
		)
	}
}

func Test_Key_HasInChains_Verification(t *testing.T) {
	for caseIndex, testCase := range keyHasInChainsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")
		search, _ := input.GetAsString("search")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		result := key.HasInChains(search)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Key_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range keyClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		cloned := key.ClonePtr()
		originalCompiled := key.Compile()
		clonedCompiled := cloned.Compile()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			originalCompiled,
			clonedCompiled,
		)
	}
}

func Test_Key_Length_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		length := key.Length()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", length))
	}
}

func Test_Key_IsEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range keyIsEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		isEmpty := key.IsEmpty()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", isEmpty))
	}
}

func Test_Key_CompileWithAdditional_Verification(t *testing.T) {
	for caseIndex, testCase := range keyCompileWithAdditionalTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		main, _ := input.GetAsString("main")
		chains, _ := input.GetAsStrings("chains")
		additional, _ := input.GetAsString("additional")

		// Act
		key := keymk.NewKey.DefaultStrings(main, chains...)
		compiled := key.Compile(additional)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, compiled)
	}
}
