package coreappendtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PrependAppendAnyItemsToStringsUsingFunc_SkipEmpty(t *testing.T) {
	// Arrange
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true,
		func(item any) string { return fmt.Sprintf("%v", item) },
		"pre",
		"post",
		"a", nil, "",
	)
	// nil items are skipped, empty string items are skipped (isSkipEmptyString=true)
	// "a" -> "a", nil -> skipped, "" -> "" which is empty so skipped

	// Act
	actual := args.Map{"result": len(result) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2 items", actual)
}

func Test_PrependAppendAnyItemsToStringsUsingFunc_SkipEmptyString_InLoop(t *testing.T) {
	// Arrange
	// This specifically targets the branch: isSkipEmptyString && currentStr == ""
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true,
		func(item any) string {
			if item == nil {
				return ""
			}
			return fmt.Sprintf("%v", item)
		},
		"pre",
		"post",
		"hello", "world",
	)

	// Act
	actual := args.Map{"result": len(result) < 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3 items", actual)
}
