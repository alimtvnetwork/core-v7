package chmodinstests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Condition ──

func Test_Condition_Defaults(t *testing.T) {
	// Arrange
	allTrue := chmodins.DefaultAllTrueCondition()
	allFalse := chmodins.DefaultAllFalseCondition()
	exceptRecurse := chmodins.DefaultAllFalseExceptRecurse()

	// Act
	actual := args.Map{
		"trueSkip": allTrue.IsSkipOnInvalid, "trueContin": allTrue.IsContinueOnError, "trueRecurse": allTrue.IsRecursive,
		"falseSkip": allFalse.IsSkipOnInvalid, "falseContin": allFalse.IsContinueOnError, "falseRecurse": allFalse.IsRecursive,
		"exceptRecurse": exceptRecurse.IsRecursive, "exceptSkip": exceptRecurse.IsSkipOnInvalid,
	}

	// Assert
	expected := args.Map{
		"trueSkip": true, "trueContin": true, "trueRecurse": true,
		"falseSkip": false, "falseContin": false, "falseRecurse": false,
		"exceptRecurse": true, "exceptSkip": false,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- Defaults", actual)
}

func Test_Condition_IsExitOnInvalid(t *testing.T) {
	// Arrange
	c := &chmodins.Condition{IsSkipOnInvalid: true}
	var nilC *chmodins.Condition

	// Act
	actual := args.Map{
		"skipTrue":  c.IsExitOnInvalid(),
		"nil":       nilC.IsExitOnInvalid(),
		"collectNil": nilC.IsCollectErrorOnInvalid(),
	}

	// Assert
	expected := args.Map{
		"skipTrue": false,
		"nil": true,
		"collectNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns error -- IsExitOnInvalid", actual)
}

func Test_Condition_Clone(t *testing.T) {
	// Arrange
	c := &chmodins.Condition{IsSkipOnInvalid: true, IsContinueOnError: true, IsRecursive: true}
	cloned := c.Clone()
	nonPtr := c.CloneNonPtr()
	var nilC *chmodins.Condition

	// Act
	actual := args.Map{
		"skip": cloned.IsSkipOnInvalid, "continue": cloned.IsContinueOnError,
		"recursive": cloned.IsRecursive, "nonPtrSkip": nonPtr.IsSkipOnInvalid,
		"nilClone": nilC.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"skip": true, "continue": true, "recursive": true,
		"nonPtrSkip": true, "nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- Clone", actual)
}

// ── RwxOwnerGroupOther ──

func Test_RwxOGO_Methods(t *testing.T) {
	// Arrange
	ogo := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "-w-")

	// Act
	actual := args.Map{
		"isOwner":  ogo.IsOwner("rwx"),
		"isGroup":  ogo.IsGroup("r-x"),
		"isOther":  ogo.IsOther("-w-"),
		"is":       ogo.Is("rwx", "r-x", "-w-"),
		"toString": ogo.ToString(true),
		"noHyphen": ogo.ToString(false),
		"string":   ogo.String(),
	}

	// Assert
	expected := args.Map{
		"isOwner": true, "isGroup": true, "isOther": true,
		"is": true, "toString": "-rwxr-x-w-", "noHyphen": "rwxr-x-w-",
		"string": "-rwxr-x-w-",
	}
	expected.ShouldBeEqual(t, 0, "RwxOGO returns correct value -- Methods", actual)
}

func Test_RwxOGO_ExpandChars(t *testing.T) {
	// Arrange
	ogo := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "-w-")
	or, ow, ox := ogo.ExpandCharOwner()
	gr, gw, gx := ogo.ExpandCharGroup()
	otr, otw, otx := ogo.ExpandCharOther()

	// Act
	actual := args.Map{
		"or": string(or), "ow": string(ow), "ox": string(ox),
		"gr": string(gr), "gw": string(gw), "gx": string(gx),
		"otr": string(otr), "otw": string(otw), "otx": string(otx),
	}

	// Assert
	expected := args.Map{
		"or": "r", "ow": "w", "ox": "x",
		"gr": "r", "gw": "-", "gx": "x",
		"otr": "-", "otw": "w", "otx": "-",
	}
	expected.ShouldBeEqual(t, 0, "RwxOGO returns correct value -- ExpandChars", actual)
}

func Test_RwxOGO_IsEqual(t *testing.T) {
	// Arrange
	a := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "-w-")
	b := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "-w-")
	c := chmodins.NewRwxOwnerGroupOther("r-x", "r-x", "-w-")
	var nilR *chmodins.RwxOwnerGroupOther

	// Act
	actual := args.Map{
		"equal":    a.IsEqual(b),
		"notEqual": a.IsEqual(c),
		"nilNil":   nilR.IsEqual(nilR),
		"nilA":     nilR.IsEqual(a),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"nilNil": true,
		"nilA": false,
	}
	expected.ShouldBeEqual(t, 0, "RwxOGO returns correct value -- IsEqual", actual)
}

func Test_RwxOGO_Clone(t *testing.T) {
	// Arrange
	a := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "-w-")
	cloned := a.Clone()
	var nilR *chmodins.RwxOwnerGroupOther

	// Act
	actual := args.Map{
		"owner": cloned.Owner, "nilClone": nilR.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"owner": "rwx",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxOGO returns correct value -- Clone", actual)
}

// ── RwxInstruction ──

func Test_RwxInstruction_Clone(t *testing.T) {
	// Arrange
	ri := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "-w-"},
		Condition:          chmodins.Condition{IsRecursive: true},
	}
	cloned := ri.Clone()
	var nilRI *chmodins.RwxInstruction

	// Act
	actual := args.Map{
		"owner":     cloned.Owner,
		"recursive": cloned.IsRecursive,
		"nilClone":  nilRI.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"owner": "rwx",
		"recursive": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxInstruction returns correct value -- Clone", actual)
}

// ── BaseRwxInstructions ──

func Test_BaseRwxInstructions_Methods(t *testing.T) {
	// Arrange
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "-w-"}},
		},
	}

	// Act
	actual := args.Map{
		"length":  base.Length(),
		"isEmpty": base.IsEmpty(),
		"hasAny":  base.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"length": 1,
		"isEmpty": false,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseRwxInstructions returns correct value -- Methods", actual)
}

func Test_BaseRwxInstructions_Nil(t *testing.T) {
	// Arrange
	var base *chmodins.BaseRwxInstructions

	// Act
	actual := args.Map{
		"length": base.Length(), "isEmpty": base.IsEmpty(),
		"hasAny": base.HasAnyItem(), "cloneNil": base.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"length": 0,
		"isEmpty": true,
		"hasAny": false,
		"cloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseRwxInstructions returns nil -- Nil", actual)
}

func Test_BaseRwxInstructions_Clone(t *testing.T) {
	// Arrange
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "-w-"}},
		},
	}
	cloned := base.Clone()
	nonPtr := base.CloneNonPtr()

	// Act
	actual := args.Map{
		"len":       cloned.Length(),
		"nonPtrLen": nonPtr.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"nonPtrLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "BaseRwxInstructions returns correct value -- Clone", actual)
}

// ── ExpandRwxFullString ──

func Test_ExpandRwxFullString(t *testing.T) {
	// Arrange
	result, err := chmodins.ExpandRwxFullStringToOwnerGroupOther("-rwxr-x-w-")
	_, errBad := chmodins.ExpandRwxFullStringToOwnerGroupOther("short")

	// Act
	actual := args.Map{
		"owner": result.Owner, "group": result.Group, "other": result.Other,
		"noErr": err == nil, "hasErrBad": errBad != nil,
	}

	// Assert
	expected := args.Map{
		"owner": "rwx", "group": "r-x", "other": "-w-",
		"noErr": true, "hasErrBad": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpandRwxFullString returns correct value -- with args", actual)
}

func Test_ExpandRwxFullStringByFixingFirst(t *testing.T) {
	// Arrange
	result, err := chmodins.ExpandRwxFullStringToOwnerGroupOtherByFixingFirst("-rwx")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"owner": result.Owner,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"owner": "rwx",
	}
	expected.ShouldBeEqual(t, 0, "ExpandByFixingFirst returns correct value -- with args", actual)
}

// ── FixRwxFullStringWithWildcards ──

func Test_FixRwxFullString(t *testing.T) {
	// Act
	actual := args.Map{
		"full":   chmodins.FixRwxFullStringWithWildcards("-rwxr-xr-x"),
		"short":  len(chmodins.FixRwxFullStringWithWildcards("-rwx")),
		"empty":  chmodins.FixRwxFullStringWithWildcards(""),
	}

	// Assert
	expected := args.Map{
		"full": "-rwxr-xr-x", "short": 10,
		"empty": chmodins.AllWildCardsRwxFullString,
	}
	expected.ShouldBeEqual(t, 0, "FixRwxFullString returns correct value -- with args", actual)
}

// ── GetRwxFullLengthError ──

func Test_GetRwxFullLengthError(t *testing.T) {
	// Arrange
	noErr := chmodins.GetRwxFullLengthError("-rwxr-xr-x")
	hasErr := chmodins.GetRwxFullLengthError("short")

	// Act
	actual := args.Map{
		"noErr": noErr == nil,
		"hasErr": hasErr != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetRwxFullLengthError returns error -- with args", actual)
}
