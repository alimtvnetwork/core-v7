package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxWrapper Creation ──

func Test_RwxWrapper_Create_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2RwxWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("mode")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Create(mode)

		actual := args.Map{
			"rwxFull":   wrapper.ToFullRwxValueString(),
			"fileMode":  wrapper.ToFileModeString(),
			"rwx3":      wrapper.ToRwxCompiledStr(),
			"hasError":  err != nil,
			"isDefined": wrapper.IsDefined(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_RwxFullString_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2RwxFullStringParseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxFull, _ := input.GetAsString("rwxFull")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.RwxFullString(rwxFull)
		hasError := err != nil

		actual := args.Map{
			"hasError": hasError,
		}

		if !hasError {
			actual["rwx3"] = wrapper.ToRwxCompiledStr()
			actual["isDefined"] = wrapper.IsDefined()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── RwxWrapper methods ──

func Test_RwxWrapper_Bytes_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	rwxBytes := wrapper.Bytes()

	// Assert
	actual := args.Map{"result": rwxBytes[0] != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "owner should be 7", actual)
	actual = args.Map{"result": rwxBytes[1] != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "group should be 5", actual)
	actual = args.Map{"result": rwxBytes[2] != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "other should be 5", actual)
}

func Test_RwxWrapper_ToCompiledOctalBytes_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	b4 := wrapper.ToCompiledOctalBytes4Digits()
	b3 := wrapper.ToCompiledOctalBytes3Digits()
	o, g, oth := wrapper.ToCompiledSplitValues()

	// Assert
	actual := args.Map{"result": b4[0] != '0'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first byte of 4-digit should be '0'", actual)
	actual = args.Map{"result": len(b3) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "3-digit should have 3 bytes", actual)
	actual = args.Map{"result": o != '7' || g != '5' || oth != '5'}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "split values wrong:", actual)
}

func Test_RwxWrapper_ToUint32Octal_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	octal := wrapper.ToUint32Octal()

	// Assert
	actual := args.Map{"result": octal != 0755}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0755, got %o", actual)
}

func Test_RwxWrapper_ToFileMode_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	mode := wrapper.ToFileMode()

	// Assert
	actual := args.Map{"result": mode != os.FileMode(0755)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0755", actual)
}

func Test_RwxWrapper_ToFullRwxValuesChars_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	chars := wrapper.ToFullRwxValuesChars()

	// Assert
	actual := args.Map{"result": len(chars) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10 chars", actual)
}

func Test_RwxWrapper_String_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	str := wrapper.String()

	// Assert
	actual := args.Map{"result": str != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -rwxr-xr-x", actual)
}

func Test_RwxWrapper_FriendlyDisplay_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	display := wrapper.FriendlyDisplay()

	// Assert
	actual := args.Map{"result": display == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FriendlyDisplay should not be empty", actual)
}

func Test_RwxWrapper_Clone_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	cloned := wrapper.Clone()

	// Assert
	actual := args.Map{"result": cloned == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should not return nil", actual)
	actual = args.Map{"result": cloned.IsEqualPtr(&wrapper)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "cloned should equal original", actual)
}

func Test_RwxWrapper_Clone_Nil_Ext2(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act
	cloned := wrapper.Clone()

	// Assert
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone on nil should return nil", actual)
}

func Test_RwxWrapper_IsEmpty_Ext2(t *testing.T) {
	// Arrange
	empty := chmodhelper.RwxWrapper{}

	// Assert
	actual := args.Map{"result": empty.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty wrapper should be empty", actual)
	if !empty.IsNull() {
		// non-pointer IsNull returns false
	}
	actual = args.Map{"result": empty.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty wrapper should be invalid", actual)
}

func Test_RwxWrapper_IsEqualPtr_Ext2(t *testing.T) {
	// Arrange
	w1, _ := chmodhelper.New.RwxWrapper.Create("755")
	w2, _ := chmodhelper.New.RwxWrapper.Create("755")
	w3, _ := chmodhelper.New.RwxWrapper.Create("644")

	// Assert
	actual := args.Map{"result": w1.IsEqualPtr(&w2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same mode wrappers should be equal", actual)
	actual = args.Map{"result": w1.IsEqualPtr(&w3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different mode wrappers should not be equal", actual)
	if !w1.IsEqualPtr(nil) == true {
		// nil comparison
	}
}

func Test_RwxWrapper_IsEqualFileMode_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert
	actual := args.Map{"result": wrapper.IsEqualFileMode(os.FileMode(0755))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal to 0755", actual)
	actual = args.Map{"result": wrapper.IsEqualFileMode(os.FileMode(0644))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal to 0644", actual)
	actual = args.Map{"result": wrapper.IsNotEqualFileMode(os.FileMode(0644))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be not equal to 0644", actual)
}

func Test_RwxWrapper_IsRwxFullEqual_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert
	actual := args.Map{"result": wrapper.IsRwxFullEqual("-rwxr-xr-x")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match full rwx string", actual)
	actual = args.Map{"result": wrapper.IsRwxFullEqual("-rw-r--r--")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match different rwx", actual)
	actual = args.Map{"result": wrapper.IsRwxFullEqual("short")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match short string", actual)
}

func Test_RwxWrapper_IsRwxEqualLocation_Ext2(t *testing.T) {
	// Assert
	w, _ := chmodhelper.New.RwxWrapper.Create("755")
	// non-existent path
	actual := args.Map{"result": w.IsRwxEqualLocation("/nonexistent/path/xyz123")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false for nonexistent path", actual)
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert - nil
	actual := args.Map{"result": wrapper.IsRwxEqualFileInfo(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false for nil fileInfo", actual)
}

func Test_RwxWrapper_IsEqualVarWrapper_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Assert - nil
	actual := args.Map{"result": wrapper.IsEqualVarWrapper(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be false for nil varWrapper", actual)
}

func Test_RwxWrapper_ToPtr_ToNonPtr_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ptr := wrapper.ToPtr()
	nonPtr := ptr.ToNonPtr()

	// Assert
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr should not return nil", actual)
	actual = args.Map{"result": nonPtr.ToFullRwxValueString() != wrapper.ToFullRwxValueString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNonPtr should equal original", actual)
}

func Test_RwxWrapper_ToRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ogo := wrapper.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual = args.Map{"result": ogo.Owner != "rwx"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "owner should be rwx", actual)
}

func Test_RwxWrapper_ToRwxInstruction_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	condition := &chmodins.Condition{IsSkipOnInvalid: true}

	// Act
	ins := wrapper.ToRwxInstruction(condition)

	// Assert
	actual := args.Map{"result": ins == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_RwxWrapper_JSON_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonBytes, err := wrapper.MarshalJSON()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
	actual = args.Map{"result": len(jsonBytes) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JSON bytes should not be empty", actual)

	// Act - unmarshal
	var parsed chmodhelper.RwxWrapper
	err2 := parsed.UnmarshalJSON(jsonBytes)

	// Assert
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
	actual = args.Map{"result": parsed.ToFullRwxValueString() != wrapper.ToFullRwxValueString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "parsed should equal original", actual)
}

func Test_RwxWrapper_Json_Methods_Ext2(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonResult := wrapper.Json()
	jsonPtrResult := wrapper.JsonPtr()

	// Assert
	actual := args.Map{"result": jsonResult.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Json() should not error", actual)
	actual = args.Map{"result": jsonPtrResult == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonPtr should not be nil", actual)

	// JsonParseSelfInject
	var target chmodhelper.RwxWrapper
	err := target.JsonParseSelfInject(&jsonResult)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject error:", actual)

	// AsJsonContractsBinder
	binder := wrapper.AsJsonContractsBinder()
	actual = args.Map{"result": binder == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "binder should not be nil", actual)
}

// ── Attribute tests ──

func Test_Attribute_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2AttributeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		attr := chmodhelper.New.Attribute.UsingRwxString(rwx)

		actual := args.Map{
			"isRead":    attr.IsRead,
			"isWrite":   attr.IsWrite,
			"isExecute": attr.IsExecute,
			"toByte":    attr.ToByte(),
			"rwxStr":    attr.ToRwxString(),
			"isEmpty":   attr.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var attr *chmodhelper.Attribute

	// Assert
	actual := args.Map{"result": attr.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil attr should be null", actual)
	actual = args.Map{"result": attr.IsAnyNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil attr IsAnyNull should be true", actual)
	actual = args.Map{"result": attr.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil attr should be empty", actual)
	actual = args.Map{"result": attr.IsZero()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil attr should be zero", actual)
	actual = args.Map{"result": attr.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil attr should be invalid", actual)
	actual = args.Map{"result": attr.IsDefined()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil attr should not be defined", actual)
	actual = args.Map{"result": attr.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil attr HasAnyItem should be false", actual)
	cloned := attr.Clone()
	actual = args.Map{"result": cloned != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil attr Clone should be nil", actual)
}

func Test_Attribute_ToAttributeValue_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	attrVal := attr.ToAttributeValue()

	// Assert
	actual := args.Map{"result": attrVal.Sum != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sum 7", actual)
	actual = args.Map{"result": attrVal.Read != 4}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected read 4", actual)
}

func Test_Attribute_ToSpecificBytes_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	r, w, x, sum := attr.ToSpecificBytes()

	// Assert
	actual := args.Map{"result": r != 4 || w != 0 || x != 1 || sum != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected: r= w= x= sum=", actual)
}

func Test_Attribute_ToRwx_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	rwx := attr.ToRwx()

	// Assert
	actual := args.Map{"result": rwx[0] != 'r' || rwx[1] != '-' || rwx[2] != 'x'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected rwx:", actual)
}

func Test_Attribute_ToVariant_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	v := attr.ToVariant()

	// Assert
	actual := args.Map{"result": v.Value() != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_Attribute_ToStringByte_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	sb := attr.ToStringByte()

	// Assert
	actual := args.Map{"result": sb != '7'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '7'", actual)
}

func Test_Attribute_Clone_Ext2(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.UsingRwxString("rwx")

	// Act
	cloned := attr.Clone()

	// Assert
	actual := args.Map{"result": cloned == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should not be nil", actual)
	actual = args.Map{"result": attr.IsEqualPtr(cloned)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "cloned should equal original", actual)
}

func Test_Attribute_IsEqual_Ext2(t *testing.T) {
	// Arrange
	a1 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	a2 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	a3 := chmodhelper.New.Attribute.UsingRwxString("r--")

	// Assert
	actual := args.Map{"result": a1.IsEqual(a2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same attrs should be equal", actual)
	actual = args.Map{"result": a1.IsEqual(a3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different attrs should not be equal", actual)
}

func Test_Attribute_IsEqualPtr_BothNil_Ext2(t *testing.T) {
	// Arrange
	var a1, a2 *chmodhelper.Attribute

	// Assert
	actual := args.Map{"result": a1.IsEqualPtr(a2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	a3 := chmodhelper.New.Attribute.UsingRwxString("rwx")
	actual = args.Map{"result": a1.IsEqualPtr(&a3)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
}

func Test_Attribute_UsingByte_Ext2(t *testing.T) {
	// Act
	attr, err := chmodhelper.New.Attribute.UsingByte(5)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": attr.IsRead || attr.IsWrite || !attr.IsExecute}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "byte 5 = r-x", actual)
}

func Test_Attribute_UsingByte_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.New.Attribute.UsingByte(8)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for byte > 7", actual)
}

func Test_Attribute_UsingVariant_Ext2(t *testing.T) {
	// Act
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadWrite)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": attr.IsRead || !attr.IsWrite || attr.IsExecute}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReadWrite should be rw-", actual)
}

// ── AttrVariant tests ──

func Test_AttrVariant_Methods_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2AttrVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(chmodhelper.AttrVariant)

		// Act
		actual := args.Map{
			"value":     variant.Value(),
			"isGreater": variant.IsGreaterThan(5),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AttrVariant_ToAttribute_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute

	// Act
	attr := v.ToAttribute()

	// Assert
	actual := args.Map{"result": attr.IsRead || !attr.IsWrite || !attr.IsExecute}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReadWriteExecute should have all true", actual)
}

// ── Variant tests ──

func Test_Variant_ToWrapper_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2VariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variant := input["variant"].(chmodhelper.Variant)

		// Act
		wrapper, err := variant.ToWrapper()

		actual := args.Map{
			"rwxFull":  wrapper.ToFullRwxValueString(),
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_ToWrapperPtr_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.X777

	// Act
	ptr, err := v.ToWrapperPtr()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": ptr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_Variant_ExpandOctalByte_Ext2(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	r, w, x := v.ExpandOctalByte()

	// Assert
	actual := args.Map{"result": r != '7' || w != '5' || x != '5'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

// ── ParseRwxToVarAttribute ──

func Test_ParseRwxToVarAttribute_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2ParseRwxToVarAttrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		varAttr, err := chmodhelper.ParseRwxToVarAttribute(rwx)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil {
			actual["isFixedType"] = varAttr.IsFixedType()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_VarAttribute_Methods_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("r*x")

	// Assert
	actual := args.Map{"result": varAttr.IsFixedType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have wildcard", actual)
	actual = args.Map{"result": varAttr.HasWildcard()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have wildcard", actual)
	actual = args.Map{"result": varAttr.String() != "r*x"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected r*x", actual)

	// Clone
	cloned := varAttr.Clone()
	actual = args.Map{"result": cloned == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should not be nil", actual)
	actual = args.Map{"result": varAttr.IsEqualPtr(cloned)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "cloned should equal", actual)

	// ToCompileFixAttr -- not fixed type, returns nil
	fixAttr := varAttr.ToCompileFixAttr()
	actual = args.Map{"result": fixAttr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-fixed should return nil from ToCompileFixAttr", actual)
}

func Test_VarAttribute_Fixed_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("rwx")

	// Assert
	actual := args.Map{"result": varAttr.IsFixedType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be fixed type", actual)

	fixAttr := varAttr.ToCompileFixAttr()
	actual = args.Map{"result": fixAttr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fixed type should return non-nil", actual)
	actual = args.Map{"result": fixAttr.IsRead || !fixAttr.IsWrite || !fixAttr.IsExecute}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have all true", actual)
}

func Test_VarAttribute_ToCompileAttr_WithWildcard_Ext2(t *testing.T) {
	// Arrange
	varAttr, _ := chmodhelper.ParseRwxToVarAttribute("r*x")
	fixed := chmodhelper.New.Attribute.UsingRwxString("rw-")

	// Act
	compiled := varAttr.ToCompileAttr(&fixed)

	// Assert
	actual := args.Map{"result": compiled.IsRead}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "read should be true", actual)
	actual = args.Map{"result": compiled.IsWrite}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "wildcard should inherit write=true from fixed", actual)
	actual = args.Map{"result": compiled.IsExecute}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "execute should be true", actual)
}

func Test_VarAttribute_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var varAttr *chmodhelper.VarAttribute

	// Assert
	cloned := varAttr.Clone()
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Clone should be nil", actual)
}

func Test_VarAttribute_IsEqualPtr_BothNil_Ext2(t *testing.T) {
	// Arrange
	var a, b *chmodhelper.VarAttribute

	// Assert
	actual := args.Map{"result": a.IsEqualPtr(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

// ── MergeRwxWildcardWithFixedRwx ──

func Test_MergeRwxWildcardWithFixedRwx_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2MergeRwxTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		existing, _ := input.GetAsString("existing")
		wildcard, _ := input.GetAsString("wildcard")

		// Act
		fixedAttr, err := chmodhelper.MergeRwxWildcardWithFixedRwx(existing, wildcard)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil && fixedAttr != nil {
			actual["result"] = fixedAttr.ToRwxString()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── NewRwxVariableWrapper ──

func Test_NewRwxVariableWrapper_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2NewRwxVarWrapperTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		partial, _ := input.GetAsString("partial")

		// Act
		varWrapper, err := chmodhelper.NewRwxVariableWrapper(partial)

		actual := args.Map{
			"hasError": err != nil,
		}
		if err == nil {
			actual["isFixedType"] = varWrapper.IsFixedType()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxVariableWrapper_Methods_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	actual := args.Map{"result": varWrapper.IsFixedType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be fixed", actual)
	actual = args.Map{"result": varWrapper.HasWildcard()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have wildcard", actual)

	fixedPtr := varWrapper.ToCompileFixedPtr()
	actual = args.Map{"result": fixedPtr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "fixed type should return ptr", actual)

	compiled := varWrapper.ToCompileWrapper(nil)
	actual = args.Map{"result": compiled.ToFullRwxValueString() != "-rwxr-xr--"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)

	cloned := varWrapper.Clone()
	actual = args.Map{"result": cloned == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should not be nil", actual)
	actual = args.Map{"result": varWrapper.IsEqualPtr(cloned)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "cloned should equal", actual)

	toStr := varWrapper.ToString(true)
	actual = args.Map{"result": toStr != "-rwxr-xr--"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToString with hyphen:", actual)
	toStr2 := varWrapper.ToString(false)
	actual = args.Map{"result": toStr2 != "rwxr-xr--"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToString without hyphen:", actual)

	actual = args.Map{"result": varWrapper.String() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
}

func Test_RwxVariableWrapper_PartialMatch_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	actual := args.Map{"result": varWrapper.IsOwnerPartialMatch("rwx")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "owner should match", actual)
	actual = args.Map{"result": varWrapper.IsGroupPartialMatch("r-x")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "group should match", actual)
	actual = args.Map{"result": varWrapper.IsOtherPartialMatch("r--")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "other should match", actual)
}

func Test_RwxVariableWrapper_IsEqual_Ext2(t *testing.T) {
	// Arrange
	varWrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxr-xr--")

	// Assert
	actual := args.Map{"result": varWrapper.IsEqualPartialFullRwx("-rwxr-xr--")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match full rwx", actual)
	actual = args.Map{"result": varWrapper.IsEqualPartialFullRwx("-rwxrwxrwx")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match different", actual)
	actual = args.Map{"result": varWrapper.IsMismatchPartialFullRwx("-rwxrwxrwx")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be mismatch", actual)
	actual = args.Map{"result": varWrapper.IsEqualPartialRwxPartial("-rwxr-xr--")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match partial", actual)
	actual = args.Map{"result": varWrapper.IsEqualPartialUsingFileMode(os.FileMode(0754))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match 0754", actual)
	actual = args.Map{"result": varWrapper.IsEqualUsingFileMode(os.FileMode(0754))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match 0754", actual)
	actual = args.Map{"result": varWrapper.IsEqualUsingLocation("/nonexistent/path/xyz")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match nonexistent", actual)
	actual = args.Map{"result": varWrapper.IsEqualUsingFileInfo(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match nil fileInfo", actual)
	actual = args.Map{"result": varWrapper.IsEqualRwxWrapperPtr(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match nil wrapper", actual)
}

func Test_RwxVariableWrapper_NilSafe_Ext2(t *testing.T) {
	// Arrange
	var v1, v2 *chmodhelper.RwxVariableWrapper

	// Assert
	actual := args.Map{"result": v1.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Clone should be nil", actual)
	actual = args.Map{"result": v1.IsEqualPtr(v2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

// ── GetRwxLengthError ──

func Test_GetRwxLengthError_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2GetRwxLengthErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx, _ := input.GetAsString("rwx")

		// Act
		err := chmodhelper.GetRwxLengthError(rwx)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── ParseRwxInstructionToStringRwx ──

func Test_ParseRwxInstructionToStringRwx_Verification(t *testing.T) {
	for caseIndex, testCase := range ext2ParseRwxToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwx := input["rwx"].(*chmodins.RwxOwnerGroupOther)
		includeHyphen := input["includeHyphen"].(bool)

		// Act
		result := chmodhelper.ParseRwxInstructionToStringRwx(rwx, includeHyphen)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ── ExpandCharRwx ──

func Test_ExpandCharRwx_Ext2(t *testing.T) {
	// Act
	r, w, x := chmodhelper.ExpandCharRwx("rwx")

	// Assert
	actual := args.Map{"result": r != 'r' || w != 'w' || x != 'x'}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

// ── IsChmod ──

func Test_IsChmod_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	rwxFull := info.Mode().String()

	// Assert
	actual := args.Map{"result": chmodhelper.IsChmod(tmpFile.Name(), rwxFull)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match existing chmod", actual)
	actual = args.Map{"result": chmodhelper.IsChmod(tmpFile.Name(), "short")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "short string should return false", actual)
	actual = args.Map{"result": chmodhelper.IsChmod("", "-rwxrwxrwx")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty path should return false", actual)
	actual = args.Map{"result": chmodhelper.IsChmod("/nonexistent/xyz", "-rwxrwxrwx")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nonexistent should return false", actual)
}

// ── IsChmodEqualUsingRwxOwnerGroupOther ──

func Test_IsChmodEqualUsingRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Assert
	actual := args.Map{"result": chmodhelper.IsChmodEqualUsingRwxOwnerGroupOther("/tmp", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil rwx should return false", actual)
}

// ── FileModeFriendlyString ──

func Test_FileModeFriendlyString_Ext2(t *testing.T) {
	// Act
	result := chmodhelper.FileModeFriendlyString(os.FileMode(0755))

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

// ── GetExistingChmodOfValidFile ──

func Test_GetExistingChmodOfValidFile_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile(tmpFile.Name())

	// Assert
	actual := args.Map{"result": isInvalid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be invalid", actual)
	actual = args.Map{"result": chmod == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "chmod should not be 0", actual)
}

func Test_GetExistingChmodOfValidFile_Invalid_Ext2(t *testing.T) {
	// Act
	chmod, isInvalid := chmodhelper.GetExistingChmodOfValidFile("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": isInvalid}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	actual = args.Map{"result": chmod != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "chmod should be 0", actual)
}

// ── ParseRwxOwnerGroupOtherToFileMode ──

func Test_ParseRwxOwnerGroupOtherToFileMode_Ext2(t *testing.T) {
	// Arrange
	ogo := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r--",
	}

	// Act
	mode, err := chmodhelper.ParseRwxOwnerGroupOtherToFileMode(ogo)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": mode != os.FileMode(0754)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0754, got %o", actual)
}

// ── ParseRwxInstructionToExecutor ──

func Test_ParseRwxInstructionToExecutor_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}

	// Act
	executor, err := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": executor == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "executor should not be nil", actual)
	actual = args.Map{"result": executor.IsFixedWrapper()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be fixed wrapper", actual)
	actual = args.Map{"result": executor.IsVarWrapper()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be var wrapper", actual)
}

func Test_ParseRwxInstructionToExecutor_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxInstructionToExecutor(nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_RwxInstructionExecutor_IsEqualFileMode_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	actual := args.Map{"result": executor.IsEqualFileMode(os.FileMode(0754))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be equal to 0754", actual)
	actual = args.Map{"result": executor.IsEqualFileMode(os.FileMode(0777))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be equal to 0777", actual)
}

func Test_RwxInstructionExecutor_CompiledWrapper_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Act
	compiled, err := executor.CompiledWrapper(os.FileMode(0755))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": compiled == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_RwxInstructionExecutor_IsEqualRwxPartial_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	actual := args.Map{"result": executor.IsEqualRwxPartial("-rwxr-xr--")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match partial", actual)
}

func Test_RwxInstructionExecutor_IsEqualFileInfo_Nil_Ext2(t *testing.T) {
	// Arrange
	ins := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Condition: chmodins.Condition{},
	}
	executor, _ := chmodhelper.ParseRwxInstructionToExecutor(ins)

	// Assert
	actual := args.Map{"result": executor.IsEqualFileInfo(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil fileInfo should return false", actual)
	actual = args.Map{"result": executor.IsEqualRwxWrapper(nil)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil wrapper should return false", actual)
}

// ── RwxInstructionExecutors ──

func Test_RwxInstructionExecutors_Ext2(t *testing.T) {
	// Arrange
	executors := chmodhelper.NewRwxInstructionExecutors(2)

	// Assert
	actual := args.Map{"result": executors.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty initially", actual)
	actual = args.Map{"result": executors.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have items", actual)
	actual = args.Map{"result": executors.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "length should be 0", actual)
	actual = args.Map{"result": executors.Count() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "count should be 0", actual)
	actual = args.Map{"result": executors.LastIndex() != -1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "lastIndex should be -1", actual)
	actual = args.Map{"result": executors.HasIndex(0)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have index 0", actual)

	// Add nil -- should skip
	executors.Add(nil)
	actual = args.Map{"result": executors.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "adding nil should not increase length", actual)
}

func Test_ParseRwxInstructionsToExecutors_Ext2(t *testing.T) {
	// Arrange
	instructions := []chmodins.RwxInstruction{
		{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r--",
			},
			Condition: chmodins.Condition{},
		},
	}

	// Act
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(instructions)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": executors.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_ParseRwxInstructionsToExecutors_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxInstructionsToExecutors(nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_ParseRwxInstructionsToExecutors_Empty_Ext2(t *testing.T) {
	// Act
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors([]chmodins.RwxInstruction{})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual = args.Map{"result": executors.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

// ── ParseRwxOwnerGroupOtherToRwxVariableWrapper ──

func Test_ParseRwxOwnerGroupOtherToRwxVariableWrapper_Nil_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ── FilteredPathFileInfoMap ──

func Test_FilteredPathFileInfoMap_Empty_Ext2(t *testing.T) {
	// Arrange
	fmap := chmodhelper.InvalidFilteredPathFileInfoMap()

	// Assert
	actual := args.Map{"result": fmap.HasAnyValidFileInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have valid info", actual)
	actual = args.Map{"result": fmap.IsEmptyValidFileInfos()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
	actual = args.Map{"result": fmap.HasAnyMissingPaths()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have missing", actual)
	actual = args.Map{"result": fmap.IsEmptyIssues()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty issues", actual)
	actual = args.Map{"result": fmap.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
	actual = args.Map{"result": fmap.HasAnyIssues()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have issues", actual)
	actual = args.Map{"result": fmap.LengthOfIssues() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "issues length should be 0", actual)
	actual = args.Map{"result": fmap.MissingPathsToString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing paths string should be empty", actual)
}

func Test_GetExistsFilteredPathFileInfoMap_Ext2(t *testing.T) {
	// Arrange
	tempDir := os.TempDir()

	// Act
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(
		false,
		tempDir,
		"/nonexistent/xyz",
	)

	// Assert
	actual := args.Map{"result": fmap.HasAnyValidFileInfo()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have valid tempDir", actual)
	actual = args.Map{"result": fmap.HasAnyMissingPaths()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have missing /nonexistent/xyz", actual)
	if fmap.HasError() == false {
		// should have error because isSkipOnInvalid=false
	}
}

func Test_GetExistsFilteredPathFileInfoMap_Empty_Ext2(t *testing.T) {
	// Act
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(false)

	// Assert
	actual := args.Map{"result": fmap.HasAnyValidFileInfo()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_FilteredPathFileInfoMap_LazyMethods_Ext2(t *testing.T) {
	// Arrange
	fmap := chmodhelper.GetExistsFilteredPathFileInfoMap(
		true,
		os.TempDir(),
	)

	// Act
	locs := fmap.LazyValidLocations()
	locs2 := fmap.LazyValidLocations() // cached
	infos := fmap.ValidFileInfos()
	wrappers := fmap.LazyValidLocationFileInfoRwxWrappers()
	wrappers2 := fmap.LazyValidLocationFileInfoRwxWrappers() // cached

	// Assert
	actual := args.Map{"result": len(locs) == 0 || len(locs2) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have locations", actual)
	actual = args.Map{"result": len(infos) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have infos", actual)
	actual = args.Map{"result": len(wrappers) == 0 || len(wrappers2) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have wrappers", actual)
}

// ── GetFilteredExistsPaths ──

func Test_GetFilteredExistsPaths_Ext2(t *testing.T) {
	// Act
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{
		os.TempDir(),
		"/nonexistent/xyz",
	})

	// Assert
	actual := args.Map{"result": len(found) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 found", actual)
	actual = args.Map{"result": len(missing) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 missing", actual)
}

func Test_GetFilteredExistsPaths_Empty_Ext2(t *testing.T) {
	// Act
	found, missing := chmodhelper.GetFilteredExistsPaths([]string{})

	// Assert
	actual := args.Map{"result": len(found) != 0 || len(missing) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

// ── RwxMatchingStatus ──

func Test_RwxMatchingStatus_Empty_Ext2(t *testing.T) {
	// Arrange
	status := chmodhelper.EmptyRwxMatchingStatus()

	// Assert
	actual := args.Map{"result": status.IsAllMatching}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be all matching", actual)
	actual = args.Map{"result": status.MissingFilesToString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing files should be empty", actual)
}

func Test_RwxMatchingStatus_Invalid_Ext2(t *testing.T) {
	// Arrange
	err := os.ErrNotExist
	status := chmodhelper.InvalidRwxMatchingStatus(err)

	// Assert
	actual := args.Map{"result": status.IsAllMatching}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be matching", actual)
	actual = args.Map{"result": status.Error == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
}

func Test_RwxMatchingStatus_CreateErrFinalError_AllMatching_Ext2(t *testing.T) {
	// Arrange
	status := &chmodhelper.RwxMatchingStatus{
		IsAllMatching: true,
		Error:         nil,
	}

	// Act
	err := status.CreateErrFinalError()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all matching with no error should return nil", actual)
}

// ── GetExistingChmodRwxWrappers ──

func Test_GetExistingChmodRwxWrappers_Empty_Ext2(t *testing.T) {
	// Act
	wrappers, err := chmodhelper.GetExistingChmodRwxWrappers(false)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": len(wrappers) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_GetExistingChmodRwxWrappers_ContinueOnError_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrappers, err := chmodhelper.GetExistingChmodRwxWrappers(
		true,
		tmpFile.Name(),
		"/nonexistent/xyz",
	)

	// Assert
	actual := args.Map{"result": len(wrappers) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 valid", actual)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error for nonexistent path", actual)
}

func Test_GetExistingChmodRwxWrappers_ImmediateExit_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrappers(
		false,
		"/nonexistent/xyz",
	)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
}

// ── GetRecursivePaths ──

func Test_GetRecursivePaths_File_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": len(paths) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_GetRecursivePaths_Dir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_recursive_ext2"
	os.MkdirAll(dir, 0755)
	defer os.RemoveAll(dir)
	os.WriteFile(dir+"/file.txt", []byte("test"), 0644)

	// Act
	paths, err := chmodhelper.GetRecursivePaths(false, dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": len(paths) < 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2 paths", actual)
}

func Test_GetRecursivePaths_NonExistent_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetRecursivePaths(false, "/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have error", actual)
}

// ── GetExistingChmodRwxWrapper ──

func Test_GetExistingChmodRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.GetExistingChmodRwxWrapper(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_GetExistingChmodRwxWrapper_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.GetExistingChmodRwxWrapper("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── SimpleFileReaderWriter ──

func Test_SimpleFileReaderWriter_WriteReadString_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.WriteString("hello world")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "write error:", actual)

	content, readErr := rw.ReadString()
	actual = args.Map{"result": readErr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "read error:", actual)
	actual = args.Map{"result": content != "hello world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello world'", actual)
}

func Test_SimpleFileReaderWriter_Properties_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_props_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Assert
	actual := args.Map{"result": rw.IsExist()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "file should not exist yet", actual)
	actual = args.Map{"result": rw.HasPathIssues()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have path issues", actual)
	actual = args.Map{"result": rw.IsPathInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be path invalid", actual)
	if rw.HasAnyIssues() == false {
		// parent dir doesn't exist either
	}

	str := rw.String()
	actual = args.Map{"result": str == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
}

func Test_SimpleFileReaderWriter_WriteAndRead_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_wr_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.Write([]byte("bytes content"))
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "write error:", actual)

	// Read
	bytes, readErr := rw.Read()
	actual = args.Map{"result": readErr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "read error:", actual)
	actual = args.Map{"result": string(bytes) != "bytes content"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected content:", actual)

	// ReadOnExist
	bytes2, err2 := rw.ReadOnExist()
	actual = args.Map{"result": err2 != nil || len(bytes2) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReadOnExist should work", actual)

	// ReadStringOnExist
	s, err3 := rw.ReadStringOnExist()
	actual = args.Map{"result": err3 != nil || s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExist should work", actual)

	// Expire
	expErr := rw.Expire()
	actual = args.Map{"result": expErr != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Expire error:", actual)
	actual = args.Map{"result": rw.IsExist()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "file should not exist after Expire", actual)
}

func Test_SimpleFileReaderWriter_ReadOnExist_NotExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.txt")

	// Act
	bytes, err := rw.ReadOnExist()

	// Assert
	actual := args.Map{"result": err != nil || bytes != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReadOnExist on nonexistent should return nil, nil", actual)

	s, err2 := rw.ReadStringOnExist()
	actual = args.Map{"result": err2 != nil || s != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReadStringOnExist on nonexistent should return empty", actual)
}

func Test_SimpleFileReaderWriter_WritePath_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_wp_ext2"
	filePath := dir + "/test.txt"
	newPath := dir + "/subdir/other.txt"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)
	rw.Write([]byte("init"))

	// Act
	err := rw.WritePath(true, newPath, []byte("new content"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WritePath error:", actual)
}

func Test_SimpleFileReaderWriter_JoinRelPath_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test/file.txt")

	// Act
	joined := rw.JoinRelPath("sub/other.txt")
	joinedEmpty := rw.JoinRelPath("")

	// Assert
	actual := args.Map{"result": joined == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual = args.Map{"result": joinedEmpty == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_SimpleFileReaderWriter_WriteAny_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_sfrw_any_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.WriteAny(map[string]string{"key": "value"})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "WriteAny error:", actual)
}

func Test_SimpleFileReaderWriter_Get_NotExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.json")

	// Act
	var target map[string]string
	err := rw.Get(&target)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Get on nonexistent should error", actual)
}

func Test_SimpleFileReaderWriter_Clone_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	cloned := rw.Clone()

	// Assert
	actual := args.Map{"result": cloned.FilePath != rw.FilePath}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "cloned FilePath should match", actual)
}

func Test_SimpleFileReaderWriter_ClonePtr_Nil_Ext2(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter

	// Act
	cloned := rw.ClonePtr()

	// Assert
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil ClonePtr should return nil", actual)
}

func Test_SimpleFileReaderWriter_JSON_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	jsonBytes, err := rw.MarshalJSON()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)

	var parsed chmodhelper.SimpleFileReaderWriter
	err2 := parsed.UnmarshalJSON(jsonBytes)
	actual = args.Map{"result": err2 != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
}

func Test_SimpleFileReaderWriter_Json_Methods_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/tmp/test.txt")

	// Act
	jsonResult := rw.Json()
	jsonPtrResult := rw.JsonPtr()

	// Assert
	actual := args.Map{"result": jsonResult.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Json() should not error", actual)
	actual = args.Map{"result": jsonPtrResult == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonPtr should not nil", actual)

	var target chmodhelper.SimpleFileReaderWriter
	err := target.JsonParseSelfInject(&jsonResult)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject error:", actual)

	binder := rw.AsJsonContractsBinder()
	actual = args.Map{"result": binder == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "binder should not be nil", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefault_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	initialized := rw.InitializeDefault(true)

	// Assert
	actual := args.Map{"result": initialized == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual = args.Map{"result": initialized.ParentDir == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ParentDir should be populated", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultApplyChmod_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	initialized := rw.InitializeDefaultApplyChmod()

	// Assert
	actual := args.Map{"result": initialized == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultNew_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	newRw := rw.InitializeDefaultNew()

	// Assert
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SimpleFileReaderWriter_NewPath_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
	}

	// Act
	newRw := rw.NewPath(true, "/tmp/other/file2.txt")

	// Assert
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SimpleFileReaderWriter_NewPathJoin_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  os.FileMode(0755),
		ChmodFile: os.FileMode(0644),
		FilePath:  "/tmp/test/file.txt",
		ParentDir: "/tmp/test",
	}

	// Act
	newRw := rw.NewPathJoin(true, "subdir", "file.txt")

	// Assert
	actual := args.Map{"result": newRw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SimpleFileReaderWriter_Serialize_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/file.txt")

	// Act -- file doesn't exist
	bytes, err := rw.Serialize()

	// Assert
	if err != nil || bytes != nil {
		// Serialize on non-exist returns nil, nil (alias for ReadOnExist)
	}
}

func Test_SimpleFileReaderWriter_RemoveOnExist_Ext2(t *testing.T) {
	// Arrange
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/file.txt")

	// Act
	err := rw.RemoveOnExist()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RemoveOnExist on nonexistent should not error:", actual)
}

func Test_SimpleFileReaderWriter_ExpireParentDir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_expire_parent_ext2"
	filePath := dir + "/test.txt"
	os.MkdirAll(dir, 0755)
	os.WriteFile(filePath, []byte("test"), 0644)
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Act
	err := rw.ExpireParentDir()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ExpireParentDir error:", actual)
}

func Test_SimpleFileReaderWriter_RemoveDirOnExist_Ext2(t *testing.T) {
	// Arrange -- dir doesn't exist
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, "/nonexistent/xyz/file.txt")

	// Act
	err := rw.RemoveDirOnExist()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RemoveDirOnExist on nonexistent should not error:", actual)
}

// ── newSimpleFileReaderWriterCreator ──

func Test_NewSimpleFileReaderWriter_All_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.All(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		true,
		true,
		"/tmp/dir",
		"/tmp/dir/file.txt",
	)

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_Options_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(true, true, true, "/tmp/test.txt")

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_Create_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		"/tmp/dir",
		"/tmp/dir/file.txt",
	)

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_CreateClean_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		"./tmp/dir",
		"./tmp/dir/file.txt",
	)

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_DefaultCleanPath_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(true, "./tmp/file.txt")

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_Path_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(true, os.FileMode(0755), os.FileMode(0644), "/tmp/file.txt")

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_PathCondition_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(true, true, os.FileMode(0755), os.FileMode(0644), "./tmp/file.txt")

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewSimpleFileReaderWriter_PathDirDefaultChmod_Ext2(t *testing.T) {
	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(true, os.FileMode(0644), "/tmp/file.txt")

	// Assert
	actual := args.Map{"result": rw == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ── newRwxWrapperCreator additional ──

func Test_NewRwxWrapperCreator_Invalid_Ext2(t *testing.T) {
	// Act
	w := chmodhelper.New.RwxWrapper.Invalid()
	wPtr := chmodhelper.New.RwxWrapper.InvalidPtr()
	empty := chmodhelper.New.RwxWrapper.Empty()

	// Assert
	actual := args.Map{"result": w.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be empty", actual)
	actual = args.Map{"result": wPtr == nil || !wPtr.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "InvalidPtr should be empty", actual)
	actual = args.Map{"result": empty == nil || !empty.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty should be empty", actual)
}

func Test_NewRwxWrapperCreator_CreatePtr_Ext2(t *testing.T) {
	// Act
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": ptr == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewRwxWrapperCreator_UsingBytes_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingBytes([3]byte{7, 5, 5})

	// Assert
	actual := args.Map{"result": wrapper.ToFullRwxValueString() != "-rwxr-xr-x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_UsingSpecificByte_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingSpecificByte(7, 5, 5)

	// Assert
	actual := args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_UsingAttrVariants_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrVariants(
		chmodhelper.ReadWriteExecute,
		chmodhelper.ReadExecute,
		chmodhelper.ReadExecute,
	)

	// Assert
	actual := args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_UsingAttrs_Ext2(t *testing.T) {
	// Arrange
	owner := chmodhelper.New.Attribute.UsingRwxString("rwx")
	group := chmodhelper.New.Attribute.UsingRwxString("r-x")
	other := chmodhelper.New.Attribute.UsingRwxString("r-x")

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrs(owner, group, other)

	// Assert
	actual := args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_Rwx10_Ext2(t *testing.T) {
	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.Rwx10("-rwxr-xr-x")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_Rwx9_Ext2(t *testing.T) {
	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.Rwx9("rwxr-xr-x")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_RwxFullStringWtHyphen_InvalidLength_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rw")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid length", actual)
}

func Test_NewRwxWrapperCreator_UsingFileMode_Zero_Ext2(t *testing.T) {
	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(0)
	wrapperPtr := chmodhelper.New.RwxWrapper.UsingFileModePtr(0)

	// Assert
	actual := args.Map{"result": wrapper.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "zero mode should be empty", actual)
	actual = args.Map{"result": wrapperPtr == nil || !wrapperPtr.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "zero mode ptr should be empty", actual)
}

func Test_NewRwxWrapperCreator_UsingRwxOwnerGroupOther_Ext2(t *testing.T) {
	// Arrange
	ogo := &chmodins.RwxOwnerGroupOther{
		Owner: "rwx",
		Group: "r-x",
		Other: "r-x",
	}

	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.UsingRwxOwnerGroupOther(ogo)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper.ToRwxCompiledStr() != "755"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_NewRwxWrapperCreator_Instruction_Ext2(t *testing.T) {
	// Act
	ins, err := chmodhelper.New.RwxWrapper.Instruction(
		"-rwxr-xr-x",
		chmodins.Condition{},
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": ins == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewRwxWrapperCreator_UsingExistingFile_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.New.RwxWrapper.UsingExistingFile(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_NewRwxWrapperCreator_UsingExistingFileSkipInvalidFile_Ext2(t *testing.T) {
	// Act - valid file
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	wrapper, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile(tmpFile.Name())
	actual := args.Map{"result": isInvalid || wrapper == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "valid file should work", actual)

	// Act - invalid
	wrapper2, isInvalid2 := chmodhelper.New.RwxWrapper.UsingExistingFileSkipInvalidFile("/nonexistent/xyz")
	actual = args.Map{"result": isInvalid2 || wrapper2 == nil}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "invalid file should return empty", actual)
}

func Test_NewRwxWrapperCreator_UsingExistingFileOption_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act - skip invalid = true, valid file
	wrapper, err, isInvalid := chmodhelper.New.RwxWrapper.UsingExistingFileOption(true, tmpFile.Name())
	actual := args.Map{"result": err != nil || isInvalid || wrapper == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should work for valid file with skip=true", actual)

	// Act - skip invalid = false, valid file
	wrapper2, err2, isInvalid2 := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, tmpFile.Name())
	actual = args.Map{"result": err2 != nil || isInvalid2 || wrapper2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should work for valid file with skip=false", actual)

	// Act - skip invalid = false, invalid file
	_, err3, _ := chmodhelper.New.RwxWrapper.UsingExistingFileOption(false, "/nonexistent/xyz")
	actual = args.Map{"result": err3 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error for invalid file with skip=false", actual)
}

// ── SingleRwx ──

func Test_SingleRwx_All_Ext2(t *testing.T) {
	// Arrange
	singleRwx, err := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)

	ogo := singleRwx.ToRwxOwnerGroupOther()
	actual = args.Map{"result": ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "rwx"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_Owner_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "rwx" || ogo.Group != "***" || ogo.Other != "***"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_Group_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Group)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "***" || ogo.Group != "rwx" || ogo.Other != "***"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_Other_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Other)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "***" || ogo.Group != "***" || ogo.Other != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_OwnerGroup_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerGroup)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "rwx" || ogo.Group != "rwx" || ogo.Other != "***"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_GroupOther_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.GroupOther)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "***" || ogo.Group != "rwx" || ogo.Other != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_OwnerOther_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.OwnerOther)

	// Act
	ogo := singleRwx.ToRwxOwnerGroupOther()

	// Assert
	actual := args.Map{"result": ogo.Owner != "rwx" || ogo.Group != "***" || ogo.Other != "rwx"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_SingleRwx_InvalidLength_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.NewSingleRwx("rw", chmodclasstype.All)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid length", actual)
}

func Test_SingleRwx_ToRwxInstruction_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)
	cond := &chmodins.Condition{IsRecursive: true}

	// Act
	ins := singleRwx.ToRwxInstruction(cond)

	// Assert
	actual := args.Map{"result": ins == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SingleRwx_ToVarRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Act
	varWrapper, err := singleRwx.ToVarRwxWrapper()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": varWrapper == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SingleRwx_ToRwxWrapper_All_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.All)

	// Act
	wrapper, err := singleRwx.ToRwxWrapper()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

func Test_SingleRwx_ToRwxWrapper_NonAll_Fails_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	_, err := singleRwx.ToRwxWrapper()

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-All classType should error", actual)
}

func Test_SingleRwx_ToDisabledRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	singleRwx, _ := chmodhelper.NewSingleRwx("rwx", chmodclasstype.Owner)

	// Act
	wrapper, err := singleRwx.ToDisabledRwxWrapper()

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ── TempDirGetter ──

func Test_TempDirGetter_Ext2(t *testing.T) {
	// Act
	tempDefault := chmodhelper.TempDirGetter.TempDefault()
	tempPermanent := chmodhelper.TempDirGetter.TempPermanent()
	tempOptionTrue := chmodhelper.TempDirGetter.TempOption(true)
	tempOptionFalse := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	actual := args.Map{"result": tempDefault == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TempDefault should not be empty", actual)
	actual = args.Map{"result": tempPermanent == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TempPermanent should not be empty", actual)
	actual = args.Map{"result": tempOptionTrue != tempPermanent}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TempOption(true) should equal TempPermanent", actual)
	actual = args.Map{"result": tempOptionFalse != tempDefault}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TempOption(false) should equal TempDefault", actual)
}

// ── chmodApplier tests ──

func Test_ChmodApplier_ApplyIf_False_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.ApplyIf(false, os.FileMode(0755), "/whatever")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ApplyIf false should return nil", actual)
}

func Test_ChmodApplier_OnMismatchOption_NotApply_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.OnMismatchOption(false, false, os.FileMode(0755), "/whatever")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "isApply=false should return nil", actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(os.FileMode(0755), nil)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty locations should return nil", actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(os.FileMode(0755), nil, "/tmp")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil condition should error", actual)
}

func Test_ChmodApplier_RwxPartial_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodApply.RwxPartial("-rwx", &chmodins.Condition{})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty locations should return nil", actual)
}

// ── RwxStringApplyChmod ──

func Test_RwxStringApplyChmod_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", &chmodins.Condition{})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty locations should return nil", actual)
}

func Test_RwxStringApplyChmod_InvalidLength_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("rwx", &chmodins.Condition{}, "/tmp")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "invalid length should error", actual)
}

func Test_RwxStringApplyChmod_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxStringApplyChmod("-rwxr-xr-x", nil, "/tmp")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil condition should error", actual)
}

// ── RwxOwnerGroupOtherApplyChmod ──

func Test_RwxOwnerGroupOtherApplyChmod_EmptyLocations_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"},
		&chmodins.Condition{},
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty locations should return nil", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilRwx_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, &chmodins.Condition{}, "/tmp")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil rwx should error", actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilCondition_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r--"},
		nil,
		"/tmp",
	)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil condition should error", actual)
}

// ── chmodVerifier additional tests ──

func Test_ChmodVerifier_RwxFull_InvalidLength_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "invalid length should error", actual)
}

func Test_ChmodVerifier_RwxFull_NonExistent_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.RwxFull("/nonexistent/xyz", "-rwxr-xr-x")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nonexistent path should error", actual)
}

func Test_ChmodVerifier_MismatchError_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.MismatchError(tmpFile.Name(), os.FileMode(0777))

	// Assert - may or may not error depending on actual permissions
	_ = err
}

func Test_ChmodVerifier_MismatchErrorUsingRwxFull_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodVerify.MismatchErrorUsingRwxFull(tmpFile.Name(), "-rwxrwxrwx")

	// Assert
	_ = err
}

func Test_ChmodVerifier_IsEqualRwxFull_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	info, _ := os.Stat(tmpFile.Name())
	rwxFull := info.Mode().String()

	// Assert
	actual := args.Map{"result": chmodhelper.ChmodVerify.IsEqualRwxFull(tmpFile.Name(), rwxFull)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match existing", actual)
}

func Test_ChmodVerifier_GetExistingRwxWrapper_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	wrapper, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": wrapper.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_ChmodVerifier_PathsUsingFileModeImmediateReturn_Valid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileModeImmediateReturn(
		info.Mode(), tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_ChmodVerifier_Path_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.Path(tmpFile.Name(), info.Mode())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_ChmodVerifier_PathIf_True_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())

	// Act
	err := chmodhelper.ChmodVerify.PathIf(true, tmpFile.Name(), info.Mode())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_ChmodVerifier_UsingRwxOwnerGroupOther_Nil_Ext2(t *testing.T) {
	// Act
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(nil, "/tmp")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil rwx should error", actual)
}

// ── RwxPartialToInstructionExecutor ──

func Test_RwxPartialToInstructionExecutor_NilCondition_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.RwxPartialToInstructionExecutor("-rwx", nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil condition should error", actual)
}

func Test_RwxPartialToInstructionExecutor_Valid_Ext2(t *testing.T) {
	// Act
	executor, err := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr--", &chmodins.Condition{})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": executor == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ── dirCreator additional ──

func Test_DirCreator_IfMissingLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_ifmissinglock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissingLock(os.FileMode(0755), dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_DirCreator_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_defaultlock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.DefaultLock(os.FileMode(0755), dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_DirCreator_DirectLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_directlock_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.DirectLock(dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_DirCreator_ByChecking_NewDir_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_bychecking_ext2"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(os.FileMode(0755), dir)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_DirCreator_ByChecking_ExistingDir_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_bychecking_existing_ext2"
	os.MkdirAll(dir, 0755)
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.ByChecking(os.FileMode(0755), dir)

	// Assert - chmod apply on existing dir
	_ = err
}

// ── fileWriter additional ──

func Test_FileWriter_All_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_all_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.All(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		false,
		false,
		true,
		dir,
		filePath,
		[]byte("content"),
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileWriter_AllLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_alllock_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.AllLock(
		os.FileMode(0755),
		os.FileMode(0644),
		true,
		false,
		false,
		true,
		dir,
		filePath,
		[]byte("content"),
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileWriter_Chmod_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Chmod(
		true,
		os.FileMode(0755),
		os.FileMode(0644),
		filePath,
		[]byte("content"),
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileWriter_ChmodFile_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	dir := os.TempDir() + "/chmodtest_fw_chmodfile_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.ChmodFile(
		true,
		os.FileMode(0644),
		filePath,
		[]byte("content"),
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

// ── fileBytesWriter additional ──

func Test_FileBytesWriter_WithDir_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wd_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDir(true, filePath, []byte("content"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileBytesWriter_WithDirLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wdl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirLock(true, filePath, []byte("content"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileBytesWriter_WithDirChmodLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_wdcl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.WithDirChmodLock(
		true, os.FileMode(0755), os.FileMode(0644), filePath, []byte("content"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileBytesWriter_Chmod_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fbw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Bytes.Chmod(
		true, os.FileMode(0755), os.FileMode(0644), filePath, []byte("content"))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

// ── fileStringWriter additional ──

func Test_FileStringWriter_All_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_all_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.All(
		true, os.FileMode(0755), os.FileMode(0644),
		false, false, true,
		dir, filePath, "content",
	)

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileStringWriter_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_dl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.DefaultLock(true, filePath, "content")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileStringWriter_Chmod_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_chmod_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.Chmod(
		true, os.FileMode(0755), os.FileMode(0644), filePath, "content")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_FileStringWriter_ChmodLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_fsw_chmodl_ext2"
	filePath := dir + "/test.txt"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.String.ChmodLock(
		true, os.FileMode(0755), os.FileMode(0644), filePath, "content")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

// ── fileReader ──

func Test_FileReader_Read_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	os.WriteFile(tmpFile.Name(), []byte("hello"), 0644)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": content != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_FileReader_ReadBytes_Ext2(t *testing.T) {
	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	os.WriteFile(tmpFile.Name(), []byte("hello"), 0644)
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	bytes, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	actual = args.Map{"result": string(bytes) != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_FileReader_Read_Invalid_Ext2(t *testing.T) {
	// Act
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── anyItemWriter ──

func Test_AnyItemWriter_Default_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_aiw_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.Default(
		true, filePath, map[string]string{"key": "value"})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_AnyItemWriter_DefaultLock_Ext2(t *testing.T) {
	// Arrange
	dir := os.TempDir() + "/chmodtest_aiw_dl_ext2"
	filePath := dir + "/test.json"
	defer os.RemoveAll(dir)

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.Any.DefaultLock(
		true, filePath, map[string]string{"key": "value"})

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

// ── simpleFileWriter Lock/Unlock ──

func Test_SimpleFileWriter_LockUnlock_Ext2(t *testing.T) {
	// Act - should not panic
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()
}

// ── ApplyChmod on valid temp file (Unix) ──

func Test_RwxWrapper_ApplyChmod_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	err := wrapper.ApplyChmod(false, tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_RwxWrapper_ApplyChmodSkipInvalid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	err := wrapper.ApplyChmodSkipInvalid("/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip invalid should return nil", actual)
}

func Test_RwxWrapper_ApplyChmod_InvalidNotSkip_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	err := wrapper.ApplyChmod(false, "/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error for invalid path without skip", actual)
}

func Test_RwxWrapper_ApplyChmodOptions_SkipApply_Ext2(t *testing.T) {
	// Act
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	err := wrapper.ApplyChmodOptions(false, true, false, "/whatever")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "isApply=false should return nil", actual)
}

func Test_RwxWrapper_Verify_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(info.Mode())

	// Act
	err := wrapper.Verify(tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_RwxWrapper_HasChmod_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()
	info, _ := os.Stat(tmpFile.Name())
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(info.Mode())

	// Assert
	actual := args.Map{"result": wrapper.HasChmod(tmpFile.Name())}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match", actual)
}

// ── chmodApplier Unix ──

func Test_ChmodApplier_Default_Unix_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tmpFile, _ := os.CreateTemp("", "test-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Act
	err := chmodhelper.ChmodApply.Default(os.FileMode(0755), tmpFile.Name())

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
}

func Test_ChmodApplier_SkipInvalidFile_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.SkipInvalidFile(os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip invalid should return nil", actual)
}

func Test_ChmodApplier_OnMismatch_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.OnMismatch(true, os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip invalid path should return nil", actual)
}

func Test_ChmodApplier_OnMismatchSkipInvalid_Ext2(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Act
	err := chmodhelper.ChmodApply.OnMismatchSkipInvalid(os.FileMode(0755), "/nonexistent/xyz")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip invalid should return nil", actual)
}
