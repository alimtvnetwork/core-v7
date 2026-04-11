package msgcreatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/msgcreator"
)

// Cover Assert.SortedArray with isPrint=true
func Test_Assert_SortedArray_Print(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedArray(true, true, "c b a")

	// Act
	actual := args.Map{
		"first": result[0],
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedArray returns non-empty -- isPrint=true", actual)
}

func Test_Assert_SortedArray_NoSort(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedArray(false, false, "c b a")

	// Act
	actual := args.Map{
		"first": result[0],
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedArray returns non-empty -- isSort=false", actual)
}

func Test_Assert_SortedMessage_Print(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedMessage(true, "c b a", ",")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SortedMessage returns non-empty -- isPrint=true", actual)
}

// Cover ToStrings with various types
func Test_Assert_ToStrings_Slice(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- slice", actual)
}

func Test_Assert_ToStrings_Int(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(42)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- int", actual)
}

func Test_Assert_ToStrings_Bool(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- bool", actual)
}

func Test_Assert_ToStrings_MapStringAny(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- map[string]any", actual)
}

func Test_Assert_ToStringsWithSpace_Empty(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStringsWithSpace(4, []string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsWithSpace returns empty -- empty", actual)
}
