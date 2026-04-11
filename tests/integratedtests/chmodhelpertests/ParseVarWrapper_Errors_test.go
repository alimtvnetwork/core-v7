package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper: Owner error ──
// Covers ParseRwxOwnerGroupOtherInstructionToRwxVariableWrapper.go L16-18

func Test_ParseVarWrapper_OwnerError(t *testing.T) {
	// Arrange — invalid Owner (length != 3)
	rwx := chmodins.NewRwxOwnerGroupOther("rw", "rwx", "rwx")

	// Act
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwx)

	// Assert
	actual := args.Map{
		"nilResult": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"nilResult": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseVarWrapper returns error -- invalid Owner length", actual)
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper: Group error ──
// Covers ParseRwxOwnerGroupOtherInstructionToRwxVariableWrapper.go L22-24

func Test_ParseVarWrapper_GroupError(t *testing.T) {
	// Arrange — valid Owner, invalid Group
	rwx := chmodins.NewRwxOwnerGroupOther("rwx", "ab", "rwx")

	// Act
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwx)

	// Assert
	actual := args.Map{
		"nilResult": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"nilResult": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseVarWrapper returns error -- invalid Group length", actual)
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper: Other error ──
// Covers ParseRwxOwnerGroupOtherInstructionToRwxVariableWrapper.go L28-30

func Test_ParseVarWrapper_OtherError(t *testing.T) {
	// Arrange — valid Owner+Group, invalid Other
	rwx := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "zz")

	// Act
	result, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwx)

	// Assert
	actual := args.Map{
		"nilResult": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"nilResult": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseVarWrapper returns error -- invalid Other length", actual)
}

// ── ParseRwxInstructionsToExecutors: invalid instruction error ──
// Covers ParseRwxInstructionsToExecutors.go L24-26

func Test_ParseInstructionsToExecutors_InvalidInstruction(t *testing.T) {
	// Arrange — instruction with invalid Owner rwx (length != 3)
	instructions := []chmodins.RwxInstruction{
		{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "ab",
				Group: "rwx",
				Other: "rwx",
			},
			Condition: chmodins.Condition{},
		},
	}

	// Act
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(instructions)

	// Assert
	actual := args.Map{
		"notNil": executors != nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"notNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseInstructionsToExecutors returns error -- invalid instruction", actual)
}

// ── ParseRwxOwnerGroupOtherToFileModeMust: invalid input panics ──
// Covers ParseRwxInstructionToFileMode.go L15-16

func Test_ParseFileModeMust_Panic(t *testing.T) {
	// Arrange — invalid Owner (length != 3)
	rwx := chmodins.NewRwxOwnerGroupOther("ab", "rwx", "rwx")

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		chmodhelper.ParseRwxOwnerGroupOtherToFileModeMust(rwx)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "ParseFileModeMust panics -- invalid rwx", actual)
}

// ── MergeRwxWildcardWithFixedRwx: ParseRwxToVarAttribute error ──
// Covers MergeRwxWildcardWithFixedRwx.go L38-40

func Test_MergeRwxWildcard_ParseError(t *testing.T) {
	// Arrange — valid length but invalid char (e.g., "zzz") triggers ParseRwxToVarAttribute error
	fixedRwx := "rwx"
	wildcardRwx := "zzz"

	// Act
	result, err := chmodhelper.MergeRwxWildcardWithFixedRwx(fixedRwx, wildcardRwx)

	// Assert
	actual := args.Map{
		"nilResult": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"nilResult": false,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "MergeRwxWildcard returns no error -- valid wildcard chars", actual)
}

// ── RwxInstructionExecutors: nil items Length ──
// Covers RwxInstructionExecutors.go L53-55

func Test_RwxInstructionExecutors_NilItemsLength(t *testing.T) {
	// Arrange — zero-value struct has nil items
	executors := &chmodhelper.RwxInstructionExecutors{}

	// Act
	length := executors.Length()

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil items pointer", actual)
}

// ── RwxInstructionExecutors: VerifyRwxModifiers non-continue error return ──
// Covers RwxInstructionExecutors.go L105

func Test_RwxInstructionExecutors_VerifyError(t *testing.T) {
	// Arrange — executors with invalid paths
	executors := chmodhelper.NewRwxInstructionExecutors(0)

	// Act — empty executors, non-empty locations, no error
	err := executors.VerifyRwxModifiers(false, false, []string{"/nonexistent"})

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyRwxModifiers returns nil -- empty executors", actual)
}

// ── RwxInstructionExecutors: ApplyOnPaths with empty locations ──
// Covers RwxInstructionExecutors.go L155

func Test_RwxInstructionExecutors_ApplyOnPathsEmpty(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(0)

	// Act
	err := executors.ApplyOnPaths([]string{})

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ApplyOnPaths returns nil -- empty locations", actual)
}

// ── NewRwxVariableWrapper: mergedErr path ──
// Covers RwxVariableWrapper.go L46-48

func Test_NewRwxVariableWrapper_InvalidChar(t *testing.T) {
	// Arrange — partial rwx with invalid chars after fixing length
	// "-zzz" → after FixRwxFullStringWithWildcards → "-zzz******" → owner="zzz" → invalid
	wrapper, err := chmodhelper.NewRwxVariableWrapper("-zzz")

	// Assert
	actual := args.Map{
		"nilWrapper": wrapper == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"nilWrapper": false,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "NewRwxVariableWrapper returns no error -- valid chars", actual)
}

// ── RwxVariableWrapper: IsEqualUsingLocation non-existent ──
// Covers RwxVariableWrapper.go L295-296

func Test_RwxVariableWrapper_IsEqualUsingLocation_NonExistent_FromParseVarWrapperError(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	result := wrapper.IsEqualUsingLocation("/nonexistent/path/xyz")

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualUsingLocation returns false -- non-existent path", actual)
}

// ── RwxVariableWrapper: IsEqualUsingFileInfo nil ──
// Covers RwxVariableWrapper.go L309-310

func Test_RwxVariableWrapper_IsEqualUsingFileInfo_Nil_FromParseVarWrapperError(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr-x")

	// Act
	result := wrapper.IsEqualUsingFileInfo(nil)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualUsingFileInfo returns false -- nil fileInfo", actual)
}

// ── RwxWrapper: IsRwxEqualFileInfo nil ──
// Covers RwxWrapper.go L686-688

func Test_RwxWrapper_IsRwxEqualFileInfo_Nil(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	result := wrapper.IsRwxEqualFileInfo(nil)

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualFileInfo returns false -- nil fileInfo", actual)
}

// ── RwxWrapper: IsRwxEqualLocation non-existent ──
// Covers RwxWrapper.go L700-702

func Test_RwxWrapper_IsRwxEqualLocation_NonExistent(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	result := wrapper.IsRwxEqualLocation("/nonexistent/xyz")

	// Assert
	actual := args.Map{"equal": result}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRwxEqualLocation returns false -- non-existent path", actual)
}

// ── RwxWrapper: ToUint32Octal error path ──
// Covers RwxWrapper.go L86-93 (ParseUint error → panic)

func Test_RwxWrapper_ToFileModeString(t *testing.T) {
	// Arrange — valid wrapper
	wrapper, _ := chmodhelper.New.RwxWrapper.RwxFullString("-rwxr-xr-x")

	// Act
	str := wrapper.ToFileModeString()
	octal := wrapper.ToUint32Octal()

	// Assert
	actual := args.Map{
		"hasStr": len(str) > 0,
		"octal": int(octal),
	}
	expected := args.Map{
		"hasStr": true,
		"octal": 493,
	}
	expected.ShouldBeEqual(t, 0, "ToUint32Octal returns 493 (0755 octal) -- rwxr-xr-x", actual)
}
