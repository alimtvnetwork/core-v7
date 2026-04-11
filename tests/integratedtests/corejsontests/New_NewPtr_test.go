package corejsontests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: New - valid
// ==========================================================================

func Test_New_Valid_FromNewNewPtr(t *testing.T) {
	tc := newValidTestCase
	result := corejson.New(struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 30})

	actual := args.Map{
		"hasError":    result.HasError(),
		"isEmpty":     result.IsEmpty(),
		"hasBytes":    len(result.Bytes) > 0,
		"hasTypeName": result.TypeName != "",
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: New - nil
// ==========================================================================

func Test_New_Nil(t *testing.T) {
	tc := newNilTestCase
	result := corejson.New(nil)

	actual := args.Map{
		"hasError":     result.HasError(),
		"bytesContent": string(result.Bytes),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: New - channel
// ==========================================================================

func Test_New_Channel(t *testing.T) {
	tc := newChannelTestCase
	result := corejson.New(make(chan int))

	actual := args.Map{
		"hasError":             result.HasError(),
		"errorContainsMarshal": strings.Contains(result.Error.Error(), "marshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: NewPtr - valid
// ==========================================================================

func Test_NewPtr_Valid_FromNewNewPtr(t *testing.T) {
	tc := newPtrValidTestCase
	result := corejson.NewPtr(struct {
		Name string
		Age  int
	}{Name: "Bob", Age: 25})

	actual := args.Map{
		"isNonNil": result != nil,
		"hasError": result.HasError(),
		"isEmpty":  result.IsEmpty(),
		"hasBytes": len(result.Bytes) > 0,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: NewPtr - nil
// ==========================================================================

func Test_NewPtr_Nil(t *testing.T) {
	tc := newPtrNilTestCase
	result := corejson.NewPtr(nil)

	actual := args.Map{
		"isNonNil":     result != nil,
		"hasError":     result.HasError(),
		"bytesContent": string(result.Bytes),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: NewPtr - channel
// ==========================================================================

func Test_NewPtr_Channel(t *testing.T) {
	tc := newPtrChannelTestCase
	result := corejson.NewPtr(make(chan string))

	actual := args.Map{
		"isNonNil":             result != nil,
		"hasError":             result.HasError(),
		"errorContainsMarshal": strings.Contains(result.Error.Error(), "marshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
