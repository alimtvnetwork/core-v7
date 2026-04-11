package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Condition ──

func Test_Chmodins_Condition_Defaults(t *testing.T) {
	// Arrange
	allTrue := chmodins.DefaultAllTrueCondition()
	allFalse := chmodins.DefaultAllFalseCondition()
	recurse := chmodins.DefaultAllFalseExceptRecurse()

	// Act
	actual := args.Map{
		"trueSkip":    allTrue.IsSkipOnInvalid,
		"trueCont":    allTrue.IsContinueOnError,
		"trueRecur":   allTrue.IsRecursive,
		"falseSkip":   allFalse.IsSkipOnInvalid,
		"falseCont":   allFalse.IsContinueOnError,
		"falseRecur":  allFalse.IsRecursive,
		"recurSkip":   recurse.IsSkipOnInvalid,
		"recurRecur":  recurse.IsRecursive,
	}

	// Assert
	expected := args.Map{
		"trueSkip":    true,
		"trueCont":    true,
		"trueRecur":   true,
		"falseSkip":   false,
		"falseCont":   false,
		"falseRecur":  false,
		"recurSkip":   false,
		"recurRecur":  true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- defaults", actual)
}

func Test_Chmodins_Condition_Methods(t *testing.T) {
	// Arrange
	c := chmodins.DefaultAllTrueCondition()
	var nilC *chmodins.Condition

	// Act
	actual := args.Map{
		"exitOnInvalid":    c.IsExitOnInvalid(),
		"collectOnInvalid": c.IsCollectErrorOnInvalid(),
		"nilExit":          nilC.IsExitOnInvalid(),
		"nilCollect":       nilC.IsCollectErrorOnInvalid(),
	}

	// Assert
	expected := args.Map{
		"exitOnInvalid":    false,
		"collectOnInvalid": false,
		"nilExit":          true,
		"nilCollect":       true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- methods", actual)
}

func Test_Chmodins_Condition_Clone(t *testing.T) {
	// Arrange
	c := chmodins.DefaultAllTrueCondition()
	clone := c.Clone()
	var nilC *chmodins.Condition
	nilClone := nilC.Clone()
	nonPtrClone := c.CloneNonPtr()

	// Act
	actual := args.Map{
		"cloneNotNil":    clone != nil,
		"cloneSkip":      clone.IsSkipOnInvalid,
		"nilCloneIsNil":  nilClone == nil,
		"nonPtrRecursive": nonPtrClone.IsRecursive,
	}

	// Assert
	expected := args.Map{
		"cloneNotNil":    true,
		"cloneSkip":      true,
		"nilCloneIsNil":  true,
		"nonPtrRecursive": true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- clone", actual)
}

// ── RwxOwnerGroupOther ──

func Test_Chmodins_RwxOwnerGroupOther_New(t *testing.T) {
	// Arrange
	r := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")

	// Act
	actual := args.Map{
		"owner":    r.Owner,
		"group":    r.Group,
		"other":    r.Other,
		"isOwner":  r.IsOwner("rwx"),
		"isGroup":  r.IsGroup("r-x"),
		"isOther":  r.IsOther("r--"),
		"is":       r.Is("rwx", "r-x", "r--"),
		"isNot":    r.Is("rwx", "rwx", "rwx"),
	}

	// Assert
	expected := args.Map{
		"owner":    "rwx",
		"group":    "r-x",
		"other":    "r--",
		"isOwner":  true,
		"isGroup":  true,
		"isOther":  true,
		"is":       true,
		"isNot":    false,
	}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOther returns correct value -- new", actual)
}

func Test_Chmodins_RwxOwnerGroupOther_ExpandChars(t *testing.T) {
	// Arrange
	r := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")
	or, ow, ox := r.ExpandCharOwner()
	gr, gw, gx := r.ExpandCharGroup()
	otr, otw, otx := r.ExpandCharOther()

	// Act
	actual := args.Map{
		"or": or, "ow": ow, "ox": ox,
		"gr": gr, "gw": gw, "gx": gx,
		"otr": otr, "otw": otw, "otx": otx,
	}

	// Assert
	expected := args.Map{
		"or": byte('r'), "ow": byte('w'), "ox": byte('x'),
		"gr": byte('r'), "gw": byte('-'), "gx": byte('x'),
		"otr": byte('r'), "otw": byte('-'), "otx": byte('-'),
	}
	expected.ShouldBeEqual(t, 0, "ExpandChars returns correct value -- with args", actual)
}

func Test_Chmodins_RwxOwnerGroupOther_IsEqual(t *testing.T) {
	// Arrange
	a := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")
	b := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")
	c := chmodins.NewRwxOwnerGroupOther("rwx", "rwx", "rwx")
	var nilR *chmodins.RwxOwnerGroupOther

	// Act
	actual := args.Map{
		"equal":     a.IsEqual(b),
		"notEqual":  a.IsEqual(c),
		"nilBoth":   nilR.IsEqual(nil),
		"nilLeft":   nilR.IsEqual(a),
		"nilRight":  a.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"equal":     true,
		"notEqual":  false,
		"nilBoth":   true,
		"nilLeft":   false,
		"nilRight":  false,
	}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOther returns correct value -- IsEqual", actual)
}

func Test_Chmodins_RwxOwnerGroupOther_ToString(t *testing.T) {
	// Arrange
	r := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")

	// Act
	actual := args.Map{
		"withHyphen":    r.ToString(true),
		"withoutHyphen": r.ToString(false),
		"string":        r.String(),
	}

	// Assert
	expected := args.Map{
		"withHyphen":    "-rwxr-xr--",
		"withoutHyphen": "rwxr-xr--",
		"string":        "-rwxr-xr--",
	}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOther returns correct value -- ToString", actual)
}

func Test_Chmodins_RwxOwnerGroupOther_Clone(t *testing.T) {
	// Arrange
	r := chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--")
	clone := r.Clone()
	var nilR *chmodins.RwxOwnerGroupOther
	nilClone := nilR.Clone()

	// Act
	actual := args.Map{
		"cloneNotNil":  clone != nil,
		"cloneOwner":   clone.Owner,
		"nilCloneNil":  nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"cloneNotNil":  true,
		"cloneOwner":   "rwx",
		"nilCloneNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "RwxOwnerGroupOther returns correct value -- Clone", actual)
}

// ── RwxInstruction ──

func Test_Chmodins_RwxInstruction_Clone(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
		Condition:          *chmodins.DefaultAllTrueCondition(),
	}
	clone := ins.Clone()
	var nilIns *chmodins.RwxInstruction
	nilClone := nilIns.Clone()

	// Act
	actual := args.Map{
		"cloneNotNil": clone != nil,
		"cloneOwner":  clone.Owner,
		"cloneRecur":  clone.IsRecursive,
		"nilCloneNil": nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"cloneNotNil": true,
		"cloneOwner":  "rwx",
		"cloneRecur":  true,
		"nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "RwxInstruction returns correct value -- Clone", actual)
}

// ── BaseRwxInstructions ──

func Test_Chmodins_BaseRwxInstructions(t *testing.T) {
	// Arrange
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllFalseCondition(),
			},
		},
	}
	var nilBase *chmodins.BaseRwxInstructions

	// Act
	actual := args.Map{
		"len":        base.Length(),
		"isEmpty":    base.IsEmpty(),
		"hasAny":     base.HasAnyItem(),
		"nilLen":     nilBase.Length(),
		"nilEmpty":   nilBase.IsEmpty(),
		"nilHasAny":  nilBase.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len":        1,
		"isEmpty":    false,
		"hasAny":     true,
		"nilLen":     0,
		"nilEmpty":   true,
		"nilHasAny":  false,
	}
	expected.ShouldBeEqual(t, 0, "BaseRwxInstructions returns correct value -- with args", actual)
}

func Test_Chmodins_BaseRwxInstructions_Clone(t *testing.T) {
	// Arrange
	base := &chmodins.BaseRwxInstructions{
		RwxInstructions: []chmodins.RwxInstruction{
			{
				RwxOwnerGroupOther: *chmodins.NewRwxOwnerGroupOther("rwx", "r-x", "r--"),
				Condition:          *chmodins.DefaultAllFalseCondition(),
			},
		},
	}
	clone := base.Clone()
	nonPtrClone := base.CloneNonPtr()
	var nilBase *chmodins.BaseRwxInstructions
	nilClone := nilBase.Clone()

	// Act
	actual := args.Map{
		"cloneNotNil":    clone != nil,
		"cloneLen":       clone.Length(),
		"nonPtrLen":      nonPtrClone.Length(),
		"nilCloneIsNil":  nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"cloneNotNil":    true,
		"cloneLen":       1,
		"nonPtrLen":      1,
		"nilCloneIsNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "BaseRwxInstructions returns correct value -- Clone", actual)
}

// ── ExpandRwxFullStringToOwnerGroupOther ──

func Test_Chmodins_ExpandRwx_Valid(t *testing.T) {
	// Arrange
	r, err := chmodins.ExpandRwxFullStringToOwnerGroupOther("-rwxr-xr--")

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"owner":  r.Owner,
		"group":  r.Group,
		"other":  r.Other,
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"owner":  "rwx",
		"group":  "r-x",
		"other":  "r--",
	}
	expected.ShouldBeEqual(t, 0, "ExpandRwx returns non-empty -- valid", actual)
}

func Test_Chmodins_ExpandRwx_Invalid(t *testing.T) {
	// Arrange
	_, err := chmodins.ExpandRwxFullStringToOwnerGroupOther("short")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpandRwx returns error -- invalid", actual)
}

// ── FixRwxFullStringWithWildcards ──

func Test_Chmodins_FixRwx_Full(t *testing.T) {
	// Act
	actual := args.Map{"result": chmodins.FixRwxFullStringWithWildcards("-rwxrwxrwx")}

	// Assert
	expected := args.Map{"result": "-rwxrwxrwx"}
	expected.ShouldBeEqual(t, 0, "FixRwx returns correct value -- full length", actual)
}

func Test_Chmodins_FixRwx_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": chmodins.FixRwxFullStringWithWildcards("")}

	// Assert
	expected := args.Map{"result": "-*********"}
	expected.ShouldBeEqual(t, 0, "FixRwx returns empty -- empty", actual)
}

func Test_Chmodins_FixRwx_Partial(t *testing.T) {
	// Arrange
	result := chmodins.FixRwxFullStringWithWildcards("-rwx")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 10}
	expected.ShouldBeEqual(t, 0, "FixRwx returns correct value -- partial", actual)
}

// ── ExpandRwxFullStringToOwnerGroupOtherByFixingFirst ──

func Test_Chmodins_ExpandByFixing(t *testing.T) {
	// Arrange
	r, err := chmodins.ExpandRwxFullStringToOwnerGroupOtherByFixingFirst("-rwx")

	// Act
	actual := args.Map{
		"noErr":   err == nil,
		"notNil":  r != nil,
	}

	// Assert
	expected := args.Map{
		"noErr":   true,
		"notNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "ExpandByFixing returns correct value -- with args", actual)
}

// ── GetRwxFullLengthError ──

func Test_Chmodins_GetRwxFullLengthError_Valid(t *testing.T) {
	// Act
	actual := args.Map{"isNil": chmodins.GetRwxFullLengthError("-rwxrwxrwx") == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "GetRwxFullLengthError returns error -- valid", actual)
}

func Test_Chmodins_GetRwxFullLengthError_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"hasErr": chmodins.GetRwxFullLengthError("short") != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetRwxFullLengthError returns error -- invalid", actual)
}
