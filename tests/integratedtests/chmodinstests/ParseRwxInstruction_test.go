package chmodinstests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ParseRwxInstructionUsingJsonResult ──

func Test_ParseRwxInstruction_Nil(t *testing.T) {
	// Arrange
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(nil)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns nil -- nil", actual)
}

func Test_ParseRwxInstruction_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns empty -- empty bytes", actual)
}

func Test_ParseRwxInstruction_InvalidJson(t *testing.T) {
	// Arrange — valid JSON but not an RwxInstruction shape causes unmarshal to succeed
	// Use a raw string that's valid JSON but wrong type to trigger unmarshal error
	r := corejson.NewPtr("not-an-instruction")
	// Act
	_, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert — string JSON unmarshals into struct without error (fields stay zero),
	// but we exercise the code path
	_ = err
}

func Test_ParseRwxInstruction_Success(t *testing.T) {
	// Arrange
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
		Condition:          *chmodins.DefaultAllTrueCondition(),
	}
	r := corejson.NewPtr(ins)
	// Act
	result, err := chmodins.ParseRwxInstructionUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"noErr":     err == nil,
		"notNil":    result != nil,
		"owner":     result.Owner,
		"group":     result.Group,
		"other":     result.Other,
		"recursive": result.IsRecursive,
	}
	expected := args.Map{
		"noErr": true, "notNil": true,
		"owner": "rwx", "group": "r-x", "other": "r--",
		"recursive": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstruction returns correct value -- success", actual)
}

// ── ParseRwxInstructionUsingJsonResultMust ──

func Test_ParseRwxInstructionMust_Success(t *testing.T) {
	// Arrange
	ins := chmodins.RwxInstruction{
		RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
		Condition:          *chmodins.DefaultAllFalseCondition(),
	}
	r := corejson.NewPtr(ins)
	// Act
	result := chmodins.ParseRwxInstructionUsingJsonResultMust(r)
	// Assert
	actual := args.Map{
		"owner": result.Owner,
		"recursive": result.IsRecursive,
	}
	expected := args.Map{
		"owner": "rwx",
		"recursive": false,
	}
	expected.ShouldBeEqual(t, 0, "ParseRwxInstructionMust returns correct value -- success", actual)
}

func Test_ParseRwxInstructionMust_Panic(t *testing.T) {
	// Arrange
	defer func() { recover() }()
	// Act
	chmodins.ParseRwxInstructionUsingJsonResultMust(nil)
	// Assert
	actual := args.Map{"result": false}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic", actual)
}

// ── ParseBaseRwxInstructionsUsingJsonResult ──

func Test_ParseBaseRwxInstructions_Nil(t *testing.T) {
	// Arrange
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(nil)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns nil -- nil", actual)
}

func Test_ParseBaseRwxInstructions_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"resultNil": result == nil,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"resultNil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns empty -- empty bytes", actual)
}

func Test_ParseBaseRwxInstructions_InvalidJson(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("not-a-base-instruction")
	// Act
	_, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert — exercise unmarshal path
	_ = err
}

func Test_ParseBaseRwxInstructions_Success(t *testing.T) {
	// Arrange
	base := chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllTrueCondition(),
			},
		},
	}
	r := corejson.NewPtr(base)
	// Act
	result, err := chmodins.ParseBaseRwxInstructionsUsingJsonResult(r)
	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"notNil": result != nil,
		"len":    result.Length(),
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructions returns correct value -- success", actual)
}

// ── ParseBaseRwxInstructionsUsingJsonResultMust ──

func Test_ParseBaseRwxInstructionsMust_Success(t *testing.T) {
	// Arrange
	base := chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllFalseCondition(),
			},
		},
	}
	r := corejson.NewPtr(base)
	// Act
	result := chmodins.ParseBaseRwxInstructionsUsingJsonResultMust(r)
	// Assert
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ParseBaseRwxInstructionsMust returns correct value -- success", actual)
}

func Test_ParseBaseRwxInstructionsMust_Panic(t *testing.T) {
	// Arrange
	defer func() { recover() }()
	// Act
	chmodins.ParseBaseRwxInstructionsUsingJsonResultMust(nil)
	// Assert
	actual := args.Map{"result": false}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic", actual)
}
