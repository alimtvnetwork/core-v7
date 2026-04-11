package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Attribute_ToByte_ToRwxString(t *testing.T) {
	for caseIndex, testCase := range attributeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		read, _ := input.GetAsBool("read")
		write, _ := input.GetAsBool("write")
		execute, _ := input.GetAsBool("execute")

		// Act
		attr := chmodhelper.New.Attribute.Create(read, write, execute)

		actual := args.Map{
			"toByte":      int(attr.ToByte()),
			"toRwxString": attr.ToRwxString(),
			"isEmpty":     attr.IsEmpty(),
			"isDefined":   attr.IsDefined(),
			"isZero":      attr.IsZero(),
			"isInvalid":   attr.IsInvalid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_IsEqual_FromAttribute(t *testing.T) {
	for caseIndex, testCase := range attributeEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftRwx, _ := input.GetAsString("leftRwx")
		rightRwx, _ := input.GetAsString("rightRwx")

		// Act
		left := chmodhelper.New.Attribute.UsingRwxString(leftRwx)
		right := chmodhelper.New.Attribute.UsingRwxString(rightRwx)

		actual := args.Map{
			"isEqual": left.IsEqual(right),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_Clone(t *testing.T) {
	for caseIndex, testCase := range attributeCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		read, _ := input.GetAsBool("read")
		write, _ := input.GetAsBool("write")
		execute, _ := input.GetAsBool("execute")

		// Act
		attr := chmodhelper.New.Attribute.Create(read, write, execute)
		cloned := attr.Clone()

		actual := args.Map{
			"cloneRead":    cloned.IsRead,
			"cloneWrite":   cloned.IsWrite,
			"cloneExecute": cloned.IsExecute,
			"isEqual":      attr.IsEqualPtr(cloned),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_UsingByte(t *testing.T) {
	for caseIndex, testCase := range usingByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsInt("input")

		// Act
		attr, err := chmodhelper.New.Attribute.UsingByte(byte(inputVal))
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "UsingByte() returned error:", actual)

		actual = args.Map{
			"read":    attr.IsRead,
			"write":   attr.IsWrite,
			"execute": attr.IsExecute,
			"toByte":  int(attr.ToByte()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_UsingRwxString(t *testing.T) {
	for caseIndex, testCase := range usingRwxStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		attr := chmodhelper.New.Attribute.UsingRwxString(inputStr)

		actual := args.Map{
			"read":    attr.IsRead,
			"write":   attr.IsWrite,
			"execute": attr.IsExecute,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Attribute_ToAttributeValue_FromAttribute(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, false)

	// Act
	attrVal := attr.ToAttributeValue()

	// Assert
	actual := args.Map{"result": attrVal.Read != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Read=4", actual)
	actual = args.Map{"result": attrVal.Write != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Write=2", actual)
	actual = args.Map{"result": attrVal.Execute != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Execute=0", actual)
	actual = args.Map{"result": attrVal.Sum != 6}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Sum=6", actual)
}

func Test_Attribute_ToVariant_FromAttribute(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	v := attr.ToVariant()

	// Assert
	actual := args.Map{"result": v.Value() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected variant value 5", actual)
}

func Test_ExpandCharRwx_FromAttribute(t *testing.T) {
	for caseIndex, testCase := range expandCharRwxTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		r, w, x := chmodhelper.ExpandCharRwx(inputStr)

		actual := args.Map{
			"r": fmt.Sprintf("%d", r),
			"w": fmt.Sprintf("%d", w),
			"x": fmt.Sprintf("%d", x),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsPathExists(t *testing.T) {
	for caseIndex, testCase := range isPathExistsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		path, _ := input.GetAsString("input")

		// Act
		actual := args.Map{
			"exists":  chmodhelper.IsPathExists(path),
			"invalid": chmodhelper.IsPathInvalid(path),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsDirectory(t *testing.T) {
	for caseIndex, testCase := range isDirectoryTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		path, _ := input.GetAsString("input")

		// Act
		actual := args.Map{
			"isDir": chmodhelper.IsDirectory(path),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_GetRwxLengthError(t *testing.T) {
	for caseIndex, testCase := range getRwxLengthErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		err := chmodhelper.GetRwxLengthError(inputStr)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsPathExistsPlusFileInfo(t *testing.T) {
	// Arrange - existing path
	exists, fileInfo := chmodhelper.IsPathExistsPlusFileInfo(".")

	// Assert
	actual := args.Map{"result": exists}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected '.' to exist", actual)
	actual = args.Map{"result": fileInfo == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected fileInfo to not be nil", actual)

	// Arrange - non-existing path
	exists2, _ := chmodhelper.IsPathExistsPlusFileInfo("/non/existing/path/xyz")

	// Assert
	actual = args.Map{"result": exists2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-existing path to return false", actual)
}

func Test_Attribute_NilIsEmpty(t *testing.T) {
	// Arrange
	var attr *chmodhelper.Attribute

	// Act & Assert
	actual := args.Map{"result": attr.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil attribute IsNull to be true", actual)
	actual = args.Map{"result": attr.IsAnyNull()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil attribute IsAnyNull to be true", actual)
	actual = args.Map{"result": attr.IsEmpty()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected nil attribute IsEmpty to be true", actual)
}

func Test_Attribute_NilClone(t *testing.T) {
	// Arrange
	var attr *chmodhelper.Attribute

	// Act
	cloned := attr.Clone()

	// Assert
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil clone to be nil", actual)
}

func Test_Attribute_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var left *chmodhelper.Attribute
	var right *chmodhelper.Attribute

	// Act & Assert
	actual := args.Map{"result": left.IsEqualPtr(right)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected both nil IsEqualPtr to be true", actual)
}

func Test_Attribute_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, true)
	var nilAttr *chmodhelper.Attribute

	// Act & Assert
	actual := args.Map{"result": attr.IsEqualPtr(nilAttr)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected one nil IsEqualPtr to be false", actual)
}
