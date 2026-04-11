package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Result_Unmarshal_Valid_FromResultUnmarshal(t *testing.T) {
	tc := resultUnmarshalValidTestCase

	// Arrange
	src := exampleStruct{Name: "Alice", Age: 30}
	jsonResult := corejson.NewPtr(src)
	target := &exampleStruct{}

	// Act
	err := jsonResult.Unmarshal(target)

	actual := args.Map{
		"error":            fmt.Sprintf("%v", err),
		"deserializedName": target.Name,
		"deserializedAge":  fmt.Sprintf("%v", target.Age),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_NilReceiver(t *testing.T) {
	tc := resultUnmarshalNilTestCase

	// Arrange
	var nilResult *corejson.Result
	target := &exampleStruct{}

	// Act
	err := nilResult.Unmarshal(target)

	actual := args.Map{
		"hasError":          err != nil,
		"errorContainsNull": strings.Contains(err.Error(), "null"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_InvalidBytes(t *testing.T) {
	tc := resultUnmarshalInvalidTestCase

	// Arrange
	result := corejson.NewResult.UsingBytesTypePtr([]byte(`{invalid-json`), "TestType")
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actual := args.Map{
		"hasError":               err != nil,
		"errorContainsUnmarshal": strings.Contains(err.Error(), "unmarshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_ExistingError(t *testing.T) {
	tc := resultUnmarshalExistingErrorTestCase

	// Arrange
	ch := make(chan int)
	result := corejson.NewPtr(ch)
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actual := args.Map{
		"hasError":               err != nil,
		"errorContainsUnmarshal": strings.Contains(err.Error(), "unmarshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
