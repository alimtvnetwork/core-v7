package typesconvtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/typesconv"
)

func Test_IntPtrToSimple_NonNil(t *testing.T) {
	// Arrange
	v := 42

	// Act
	actual := args.Map{"result": typesconv.IntPtrToSimple(&v)}

	// Assert
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "IntPtrToSimple_NonNil returns nil -- with args", actual)
}

func Test_IntPtrToSimpleDef_NonNil(t *testing.T) {
	// Arrange
	v := 10

	// Act
	actual := args.Map{"result": typesconv.IntPtrToSimpleDef(&v, 99)}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrToSimpleDef_NonNil returns nil -- with args", actual)
}

func Test_IntPtrToDefPtr_NonNil(t *testing.T) {
	// Arrange
	v := 10
	r := typesconv.IntPtrToDefPtr(&v, 99)

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrToDefPtr_NonNil returns nil -- with args", actual)
}

func Test_IntPtrDefValFunc_NonNil(t *testing.T) {
	// Arrange
	v := 10
	r := typesconv.IntPtrDefValFunc(&v, func() int { return 99 })

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IntPtrDefValFunc_NonNil returns nil -- with args", actual)
}

func Test_BytePtrToSimple_NonNil(t *testing.T) {
	// Arrange
	v := byte(5)

	// Act
	actual := args.Map{"result": int(typesconv.BytePtrToSimple(&v))}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToSimple_NonNil returns nil -- with args", actual)
}

func Test_BytePtrToSimpleDef_NonNil(t *testing.T) {
	// Arrange
	v := byte(5)

	// Act
	actual := args.Map{"result": int(typesconv.BytePtrToSimpleDef(&v, 9))}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToSimpleDef_NonNil returns nil -- with args", actual)
}

func Test_BytePtrToDefPtr_NonNil(t *testing.T) {
	// Arrange
	v := byte(5)
	r := typesconv.BytePtrToDefPtr(&v, 9)

	// Act
	actual := args.Map{"result": int(*r)}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrToDefPtr_NonNil returns nil -- with args", actual)
}

func Test_BytePtrDefValFunc_NonNil(t *testing.T) {
	// Arrange
	v := byte(5)
	r := typesconv.BytePtrDefValFunc(&v, func() byte { return 9 })

	// Act
	actual := args.Map{"result": int(*r)}

	// Assert
	expected := args.Map{"result": 5}
	expected.ShouldBeEqual(t, 0, "BytePtrDefValFunc_NonNil returns nil -- with args", actual)
}

func Test_FloatPtrToSimple_NonNil(t *testing.T) {
	// Arrange
	v := float32(3.14)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimple(&v))}

	// Assert
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToSimple_NonNil returns nil -- with args", actual)
}

func Test_FloatPtrToSimpleDef_NonNil(t *testing.T) {
	// Arrange
	v := float32(3.14)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%.2f", typesconv.FloatPtrToSimpleDef(&v, 9.9))}

	// Assert
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToSimpleDef_NonNil returns nil -- with args", actual)
}

func Test_FloatPtrToDefPtr_NonNil(t *testing.T) {
	// Arrange
	v := float32(3.14)
	r := typesconv.FloatPtrToDefPtr(&v, 9.9)

	// Act
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}

	// Assert
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrToDefPtr_NonNil returns nil -- with args", actual)
}

func Test_FloatPtrDefValFunc_NonNil(t *testing.T) {
	// Arrange
	v := float32(3.14)
	r := typesconv.FloatPtrDefValFunc(&v, func() float32 { return 9.9 })

	// Act
	actual := args.Map{"result": fmt.Sprintf("%.2f", *r)}

	// Assert
	expected := args.Map{"result": "3.14"}
	expected.ShouldBeEqual(t, 0, "FloatPtrDefValFunc_NonNil returns nil -- with args", actual)
}

func Test_StringPtrToSimpleDef_NonNil(t *testing.T) {
	// Arrange
	v := "hello"

	// Act
	actual := args.Map{"result": typesconv.StringPtrToSimpleDef(&v, "fb")}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrToSimpleDef_NonNil returns nil -- with args", actual)
}

func Test_StringPtrToDefPtr_NonNil(t *testing.T) {
	// Arrange
	v := "hello"
	r := typesconv.StringPtrToDefPtr(&v, "fb")

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrToDefPtr_NonNil returns nil -- with args", actual)
}

func Test_StringPtrDefValFunc_NonNil(t *testing.T) {
	// Arrange
	v := "hello"
	r := typesconv.StringPtrDefValFunc(&v, func() string { return "fb" })

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringPtrDefValFunc_NonNil returns nil -- with args", actual)
}

func Test_StringToBool_YES(t *testing.T) {
	// Act
	actual := args.Map{"result": typesconv.StringToBool("YES")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringToBool_YES returns correct value -- with args", actual)
}

func Test_StringToBool_No(t *testing.T) {
	// Act
	actual := args.Map{"result": typesconv.StringToBool("No")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_No returns correct value -- with args", actual)
}

func Test_StringToBool_NO(t *testing.T) {
	// Act
	actual := args.Map{"result": typesconv.StringToBool("NO")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_NO returns correct value -- with args", actual)
}

func Test_StringToBool_false(t *testing.T) {
	// Act
	actual := args.Map{"result": typesconv.StringToBool("false")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringToBool_false returns non-empty -- with args", actual)
}

func Test_StringPointerToBool_Empty(t *testing.T) {
	// Arrange
	s := ""

	// Act
	actual := args.Map{"result": typesconv.StringPointerToBool(&s)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringPointerToBool_Empty returns empty -- with args", actual)
}

func Test_StringPointerToBoolPtr_Empty(t *testing.T) {
	// Arrange
	s := ""
	r := typesconv.StringPointerToBoolPtr(&s)

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringPointerToBoolPtr_Empty returns empty -- with args", actual)
}

func Test_StringToBoolPtr_NonEmpty(t *testing.T) {
	// Arrange
	r := typesconv.StringToBoolPtr("yes")

	// Act
	actual := args.Map{"result": *r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringToBoolPtr_NonEmpty returns empty -- with args", actual)
}
