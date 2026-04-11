package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_RwxWrapper_Create(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Create(mode)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Create() returned error:", actual)

		actual := args.Map{
			"ownerRwx":    wrapper.Owner.ToRwxString(),
			"groupRwx":    wrapper.Group.ToRwxString(),
			"otherRwx":    wrapper.Other.ToRwxString(),
			"fullRwx":     wrapper.ToFullRwxValueString(),
			"rwx9":        wrapper.ToFullRwxValueStringExceptHyphen(),
			"fileMode":    wrapper.ToFileModeString(),
			"rwxCompiled": wrapper.ToRwxCompiledStr(),
			"isEmpty":     wrapper.IsEmpty(),
			"isDefined":   wrapper.IsDefined(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_RwxFullString(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperRwxFullStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxStr, _ := input.GetAsString("input")
		expected := testCase.ExpectedInput.(args.Map)
		hasError, _ := expected.GetAsBool("hasError")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.RwxFullString(rwxStr)

		if hasError {
			// Assert
			actual := args.Map{
				"hasError": err != nil,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			// Assert
			actual := args.Map{
				"ownerRwx":  wrapper.Owner.ToRwxString(),
				"groupRwx":  wrapper.Group.ToRwxString(),
				"otherRwx":  wrapper.Other.ToRwxString(),
				"hasError":  err != nil,
				"isDefined": wrapper.IsDefined(),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_RwxWrapper_Rwx9(t *testing.T) {
	for caseIndex, testCase := range rwxWrapper9StringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxStr, _ := input.GetAsString("input")
		expected := testCase.ExpectedInput.(args.Map)
		hasError, _ := expected.GetAsBool("hasError")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Rwx9(rwxStr)

		if hasError {
			// Assert
			actual := args.Map{
				"hasError": err != nil,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			// Assert
			actual := args.Map{
				"ownerRwx":  wrapper.Owner.ToRwxString(),
				"groupRwx":  wrapper.Group.ToRwxString(),
				"otherRwx":  wrapper.Other.ToRwxString(),
				"hasError":  err != nil,
				"isDefined": wrapper.IsDefined(),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_RwxWrapper_Bytes_FromRwxWrapper(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperBytesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		bytes := wrapper.Bytes()

		actual := args.Map{
			"byte0": int(bytes[0]),
			"byte1": int(bytes[1]),
			"byte2": int(bytes[2]),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_Clone(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		cloned := wrapper.Clone()

		actual := args.Map{
			"isEqual":    wrapper.IsEqualPtr(cloned),
			"ownerRwx":   cloned.Owner.ToRwxString(),
			"groupRwx":   cloned.Group.ToRwxString(),
			"otherRwx":   cloned.Other.ToRwxString(),
			"clonedNull": cloned.IsNull(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_IsEqualPtr_FromRwxWrapper(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftMode, _ := input.GetAsString("left")
		rightMode, _ := input.GetAsString("right")

		// Act
		left, _ := chmodhelper.New.RwxWrapper.Create(leftMode)
		right, _ := chmodhelper.New.RwxWrapper.Create(rightMode)

		actual := args.Map{
			"isEqual": left.IsEqualPtr(&right),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_NilClone(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act
	cloned := wrapper.Clone()

	// Assert
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil clone to be nil", actual)
}

func Test_RwxWrapper_NilIsEmpty(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act & Assert
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil wrapper IsEmpty to be true", actual)
	actual := args.Map{"result": wrapper.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil wrapper IsNull to be true", actual)
	actual := args.Map{"result": wrapper.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil wrapper IsInvalid to be true", actual)
}

func Test_RwxWrapper_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var left *chmodhelper.RwxWrapper
	var right *chmodhelper.RwxWrapper

	// Act & Assert
	actual := args.Map{"result": left.IsEqualPtr(right)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected both nil IsEqualPtr to be true", actual)
}

func Test_RwxWrapper_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	var nilWrapper *chmodhelper.RwxWrapper

	// Act & Assert
	actual := args.Map{"result": wrapper.IsEqualPtr(nilWrapper)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected one nil IsEqualPtr to be false", actual)
}

func Test_RwxWrapper_ToFileMode_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	mode := wrapper.ToFileMode()

	// Assert
	expectedMode := os.FileMode(0755)
	actual := args.Map{"result": mode != expectedMode}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FileMode", actual)
}

func Test_RwxWrapper_ToUint32Octal_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	octal := wrapper.ToUint32Octal()

	// Assert
	actual := args.Map{"result": octal != 0755}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected octal 0755 ()", actual)
}

func Test_RwxWrapper_String(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	str := wrapper.String()

	// Assert
	actual := args.Map{"result": str != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_UsingBytes(t *testing.T) {
	// Arrange
	bytes := [3]byte{7, 5, 5}

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingBytes(bytes)

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_UsingSpecificByte(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.UsingSpecificByte(7, 5, 5)

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_Invalid(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.Invalid()

	// Assert
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Invalid wrapper to be empty", actual)
}

func Test_RwxWrapper_InvalidPtr(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.InvalidPtr()

	// Assert
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected InvalidPtr wrapper to be empty", actual)
}

func Test_RwxWrapper_Empty(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.Empty()

	// Assert
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Empty wrapper to be empty", actual)
	if !wrapper.IsNull() {
		// Note: Empty returns *RwxWrapper{}, not nil
		// IsNull checks for nil
	}
}

func Test_RwxWrapper_ToCompiledOctalBytes_FromRwxWrapper(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperOctalTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		octal4 := wrapper.ToCompiledOctalBytes4Digits()

		actual := args.Map{
			"octal4": string(octal4[:]),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_ToCompiledSplitValues(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	owner, group, other := wrapper.ToCompiledSplitValues()

	// Assert
	expectedOwner := byte('7')
	expectedGroup := byte('5')
	expectedOther := byte('5')
	actual := args.Map{"result": owner != expectedOwner}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected owner '', got ''", actual)
	actual := args.Map{"result": group != expectedGroup}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected group '', got ''", actual)
	actual := args.Map{"result": other != expectedOther}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected other '', got ''", actual)
}

func Test_RwxWrapper_FriendlyDisplay_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	display := wrapper.FriendlyDisplay()

	// Assert
	actual := args.Map{"result": display == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FriendlyDisplay to not be empty", actual)
}

func Test_RwxWrapper_ToRwxOwnerGroupOther_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ogo := wrapper.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToRwxOwnerGroupOther to not be nil", actual)
	actual := args.Map{"result": ogo.Owner != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Owner 'rwx', got ''", actual)
	actual := args.Map{"result": ogo.Group != "r-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Group 'r-x', got ''", actual)
	actual := args.Map{"result": ogo.Other != "r-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Other 'r-x', got ''", actual)
}

func Test_RwxWrapper_IsRwxFullEqual_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	actual := args.Map{"result": wrapper.IsRwxFullEqual("-rwxr-xr-x")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsRwxFullEqual('-rwxr-xr-x') true", actual)
	actual := args.Map{"result": wrapper.IsRwxFullEqual("-rw-r--r--")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsRwxFullEqual('-rw-r--r--') false", actual)
	actual := args.Map{"result": wrapper.IsRwxFullEqual("short")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsRwxFullEqual('short') false for short string", actual)
}

func Test_RwxWrapper_IsEqualFileMode_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	actual := args.Map{"result": wrapper.IsEqualFileMode(os.FileMode(0755))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEqualFileMode(0755) true", actual)
	actual := args.Map{"result": wrapper.IsEqualFileMode(os.FileMode(0644))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEqualFileMode(0644) false", actual)
	actual := args.Map{"result": wrapper.IsNotEqualFileMode(os.FileMode(0644))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsNotEqualFileMode(0644) true", actual)
}

func Test_RwxWrapper_HasAnyItem(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	emptyWrapper := chmodhelper.New.RwxWrapper.Invalid()

	// Act & Assert
	actual := args.Map{"result": wrapper.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true for 755", actual)
	actual := args.Map{"result": emptyWrapper.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem false for empty wrapper", actual)
}

func Test_RwxWrapper_ToPtr_ToNonPtr_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ptr := wrapper.ToPtr()
	nonPtr := ptr.ToNonPtr()

	// Assert
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToPtr to not be nil", actual)
	actual := args.Map{"result": nonPtr.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToNonPtr to preserve value, got ''", actual)
}

func Test_RwxWrapper_Json_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonResult := wrapper.Json()
	jsonStr := jsonResult.JsonString()

	// Assert
	actual := args.Map{"result": jsonStr == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Json string to not be empty", actual)
}

func Test_RwxWrapper_UsingVariant(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variantStr, _ := input.GetAsString("input")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.Variant(variantStr))

		actual := args.Map{
			"fullRwx":  wrapper.ToFullRwxValueString(),
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_UsingAttrVariants(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrVariants(
		chmodhelper.ReadWriteExecute,
		chmodhelper.ReadExecute,
		chmodhelper.ReadExecute,
	)

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_UsingAttrs(t *testing.T) {
	// Arrange
	owner := chmodhelper.New.Attribute.Create(true, true, true)
	group := chmodhelper.New.Attribute.Create(true, false, true)
	other := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrs(owner, group, other)

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_CreatePtr(t *testing.T) {
	// Arrange & Act
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CreatePtr returned error:", actual)
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected CreatePtr to not be nil", actual)
	actual := args.Map{"result": ptr.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_UsingFileModePtr(t *testing.T) {
	// Arrange
	mode := os.FileMode(0755)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileModePtr(mode)

	// Assert
	actual := args.Map{"result": wrapper == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected UsingFileModePtr to not be nil", actual)
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_UsingFileModePtr_Zero(t *testing.T) {
	// Arrange
	mode := os.FileMode(0)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileModePtr(mode)

	// Assert
	actual := args.Map{"result": wrapper == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected UsingFileModePtr to not be nil even for zero", actual)
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected zero FileMode to create empty wrapper", actual)
}

func Test_RwxWrapper_UsingFileMode(t *testing.T) {
	// Arrange
	mode := os.FileMode(0644)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(mode)

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rw-r--r--"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rw-r--r--', got ''", actual)
}

func Test_RwxWrapper_UsingVariantPtr(t *testing.T) {
	// Arrange & Act
	ptr, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("755"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UsingVariantPtr returned error:", actual)
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected UsingVariantPtr to not be nil", actual)
	actual := args.Map{"result": ptr.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_ToFullRwxValuesChars(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	chars := wrapper.ToFullRwxValuesChars()

	// Assert
	actual := args.Map{"result": len(chars) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10 chars", actual)
	actual := args.Map{"result": string(chars) != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_RwxWrapper_IsEqualVarWrapper_Nil_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	actual := args.Map{"result": wrapper.IsEqualVarWrapper(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEqualVarWrapper(nil) false", actual)
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Nil_FromRwxWrapper(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	actual := args.Map{"result": wrapper.IsRwxEqualFileInfo(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsRwxEqualFileInfo(nil) false", actual)
}

func Test_FileModeFriendlyString_FromRwxWrapper(t *testing.T) {
	// Arrange
	mode := os.FileMode(0755)

	// Act
	result := chmodhelper.FileModeFriendlyString(mode)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected FileModeFriendlyString to not be empty", actual)
}

func Test_AttrVariant(t *testing.T) {
	for caseIndex, testCase := range attrVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("input")

		// Act
		variant := chmodhelper.AttrVariant(val)
		attr := variant.ToAttribute()

		actual := args.Map{
			"value":       int(variant.Value()),
			"attrRead":    attr.IsRead,
			"attrWrite":   attr.IsWrite,
			"attrExecute": attr.IsExecute,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AttrVariant_IsGreaterThan(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute // 7

	// Act & Assert
	actual := args.Map{"result": v.IsGreaterThan(8)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsGreaterThan(8) true (8 > 7)", actual)
	actual := args.Map{"result": v.IsGreaterThan(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsGreaterThan(5) false (5 < 7)", actual)
}

func Test_Variant_String_FromRwxWrapper(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	str := v.String()

	// Assert
	actual := args.Map{"result": str != "755"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '755', got ''", actual)
}

func Test_Variant_ExpandOctalByte_FromRwxWrapper(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	r, w, x := v.ExpandOctalByte()

	// Assert
	actual := args.Map{"result": r != '7'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected r='7' ()", actual)
	actual := args.Map{"result": w != '5'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected w='5' ()", actual)
	actual := args.Map{"result": x != '5'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x='5' ()", actual)
}

func Test_Variant_ToWrapper_FromRwxWrapper(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	wrapper, err := v.ToWrapper()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToWrapper returned error:", actual)
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rwxr-xr-x', got ''", actual)
}

func Test_Variant_ToWrapperPtr_FromRwxWrapper(t *testing.T) {
	// Arrange
	v := chmodhelper.X644

	// Act
	wrapper, err := v.ToWrapperPtr()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToWrapperPtr returned error:", actual)
	actual := args.Map{"result": wrapper == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToWrapperPtr to not be nil", actual)
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rw-r--r--"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '-rw-r--r--', got ''", actual)
}

func Test_Attribute_HasAnyItem(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, false)
	emptyAttr := chmodhelper.New.Attribute.Create(false, false, false)

	// Act & Assert
	actual := args.Map{"result": attr.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem true for read-only attribute", actual)
	actual := args.Map{"result": emptyAttr.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasAnyItem false for empty attribute", actual)
}

func Test_Attribute_ToSum(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, true)

	// Act
	sum := attr.ToSum()

	// Assert
	actual := args.Map{"result": sum != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ToSum 7", actual)
}

func Test_Attribute_ToRwx_FromRwxWrapper(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	rwx := attr.ToRwx()

	// Assert
	actual := args.Map{"result": rwx != [3]byte{'r', '-', 'x'}}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [r,-,x]", actual)
}

func Test_Attribute_ToStringByte(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, true)

	// Act
	sb := attr.ToStringByte()

	// Assert - 7 + '0' = '7'
	actual := args.Map{"result": sb != '7'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '7' ()", actual)
}

func Test_Attribute_ToSpecificBytes_FromRwxWrapper(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, false)

	// Act
	read, write, exe, sum := attr.ToSpecificBytes()

	// Assert
	actual := args.Map{"result": read != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected read=4", actual)
	actual := args.Map{"result": write != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected write=2", actual)
	actual := args.Map{"result": exe != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected exe=0", actual)
	actual := args.Map{"result": sum != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sum=6", actual)
}

func Test_Attribute_UsingByteMust(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.UsingByteMust(5)

	// Assert
	actual := args.Map{"result": attr.IsRead}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsRead true for byte 5", actual)
	actual := args.Map{"result": attr.IsWrite}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsWrite false for byte 5", actual)
	actual := args.Map{"result": attr.IsExecute}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsExecute true for byte 5", actual)
}

func Test_Attribute_UsingVariantMust(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.UsingVariantMust(chmodhelper.ReadWriteExecute)

	// Assert
	actual := args.Map{"result": attr.IsRead || !attr.IsWrite || !attr.IsExecute}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected all permissions true for ReadWriteExecute variant", actual)
}

func Test_Attribute_UsingVariant(t *testing.T) {
	// Arrange & Act
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadExecute)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UsingVariant returned error:", actual)
	actual := args.Map{"result": attr.IsRead || attr.IsWrite || !attr.IsExecute}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected read+execute only for ReadExecute variant", actual)
}

func Test_Attribute_Default(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.Default(true, false, true)

	// Assert
	actual := args.Map{"result": attr.IsRead || attr.IsWrite || !attr.IsExecute}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Default to create attribute with given values", actual)
}

func Test_IsChmod_EmptyLocation(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{"result": chmodhelper.IsChmod("", "-rwxr-xr-x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsChmod empty location to be false", actual)
}

func Test_IsChmod_InvalidLength(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{"result": chmodhelper.IsChmod(".", "rwx")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsChmod invalid rwx length to be false", actual)
}
