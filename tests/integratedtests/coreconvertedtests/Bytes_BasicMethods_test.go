package coreconvertedtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/converters/coreconverted"
)

func Test_Bytes_BasicMethods(t *testing.T) {
	// Arrange
	b := &coreconverted.Bytes{Values: []byte{1, 2, 3}, CombinedError: nil}
	bErr := &coreconverted.Bytes{Values: nil, CombinedError: errors.New("err")}
	var bNil *coreconverted.Bytes

	// Act
	actual := args.Map{
		"hasError":      b.HasError(),
		"length":        b.Length(),
		"hasAny":        b.HasAnyItem(),
		"isEmpty":       b.IsEmpty(),
		"hasIssues":     b.HasIssuesOrEmpty(),
		"errHasError":   bErr.HasError(),
		"errIsEmpty":    bErr.IsEmpty(),
		"errHasIssues":  bErr.HasIssuesOrEmpty(),
		"nilLength":     bNil.Length(),
	}
	expected := args.Map{
		"hasError":      false,
		"length":        3,
		"hasAny":        true,
		"isEmpty":       false,
		"hasIssues":     false,
		"errHasError":   true,
		"errIsEmpty":    true,
		"errHasIssues":  true,
		"nilLength":     0,
	}
	expected.ShouldBeEqual(t, 0, "Bytes_BasicMethods returns correct value -- with args", actual)
}

func Test_Integers_BasicMethods(t *testing.T) {
	// Arrange
	i := &coreconverted.Integers{Values: []int{1, 2, 3}, CombinedError: nil}
	iErr := &coreconverted.Integers{Values: nil, CombinedError: errors.New("err")}
	var iNil *coreconverted.Integers

	// Act
	actual := args.Map{
		"hasError":     i.HasError(),
		"length":       i.Length(),
		"hasAny":       i.HasAnyItem(),
		"isEmpty":      i.IsEmpty(),
		"hasIssues":    i.HasIssuesOrEmpty(),
		"errHasError":  iErr.HasError(),
		"errHasIssues": iErr.HasIssuesOrEmpty(),
		"nilLength":    iNil.Length(),
	}
	expected := args.Map{
		"hasError":     false,
		"length":       3,
		"hasAny":       true,
		"isEmpty":      false,
		"hasIssues":    false,
		"errHasError":  true,
		"errHasIssues": true,
		"nilLength":    0,
	}
	expected.ShouldBeEqual(t, 0, "Integers_BasicMethods returns correct value -- with args", actual)
}
