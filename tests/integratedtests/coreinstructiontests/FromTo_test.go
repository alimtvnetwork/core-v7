package coreinstructiontests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_FromTo_ClonePtr(t *testing.T) {
	// Arrange
	// Case 0: positive
	{
		tc := fromToClonePtrCopiesTestCase
		orig := &coreinstruction.FromTo{From: "source", To: "destination"}
		cloned := orig.ClonePtr()

	// Act
		actual := args.Map{
			"isNotNil": cloned != nil,
			"from":     cloned.From,
			"to":       cloned.To,
		}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	// Case 1: nil receiver
	{
		tc := fromToClonePtrNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"isNil": nilFT.ClonePtr() == nil}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_Clone(t *testing.T) {
	// Arrange
	tc := fromToCloneCopiesTestCase
	orig := coreinstruction.FromTo{From: "a", To: "b"}
	c := orig.Clone()

	// Act
	actual := args.Map{
		"from": c.From,
		"to":   c.To,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_IsNull(t *testing.T) {
	// Arrange
	{
		tc := fromToIsNullNilTestCase
		var nilFT *coreinstruction.FromTo

	// Act
		actual := args.Map{"result": nilFT.IsNull()}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToIsNullNonNilTestCase
		ft := &coreinstruction.FromTo{From: "x", To: "y"}

		actual := args.Map{"result": ft.IsNull()}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_IsFromEmpty(t *testing.T) {
	// Arrange
	{
		tc := fromToIsFromEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "", To: "dest"}

	// Act
		actual := args.Map{"result": ft.IsFromEmpty()}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToIsFromEmptyNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"result": nilFT.IsFromEmpty()}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_IsToEmpty(t *testing.T) {
	// Arrange
	{
		tc := fromToIsToEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "src", To: ""}

	// Act
		actual := args.Map{"result": ft.IsToEmpty()}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToIsToEmptyNonEmptyTestCase
		ft := &coreinstruction.FromTo{From: "src", To: "dest"}

		actual := args.Map{"result": ft.IsToEmpty()}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_String(t *testing.T) {
	// Arrange
	tc := fromToStringContainsTestCase
	ft := coreinstruction.FromTo{From: "alpha", To: "beta"}
	s := ft.String()

	// Act
	actual := args.Map{
		"containsFrom": len(s) > 0 && strings.Contains(s, "alpha"),
		"containsTo":   strings.Contains(s, "beta"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_Names(t *testing.T) {
	// Arrange
	tc := fromToNamesTestCase
	ft := coreinstruction.FromTo{From: "src", To: "dst"}

	// Act
	actual := args.Map{
		"fromName": ft.FromName(),
		"toName":   ft.ToName(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_SetFromName(t *testing.T) {
	// Arrange
	{
		tc := fromToSetFromNameUpdatesTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "t"}
		ft.SetFromName("new")

	// Act
		actual := args.Map{"from": ft.From}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToSetFromNameNilTestCase
		var nilFT *coreinstruction.FromTo
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()
			nilFT.SetFromName("x")
		}()

		actual := args.Map{"noPanic": !didPanic}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_SetToName(t *testing.T) {
	// Arrange
	tc := fromToSetToNameUpdatesTestCase
	ft := &coreinstruction.FromTo{From: "f", To: "old"}
	ft.SetToName("new")

	// Act
	actual := args.Map{"to": ft.To}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_SourceDestination(t *testing.T) {
	// Arrange
	{
		tc := fromToSourceDestMapsTestCase
		ft := &coreinstruction.FromTo{From: "src", To: "dst"}
		sd := ft.SourceDestination()

	// Act
		actual := args.Map{
			"isNotNil":    sd != nil,
			"source":      sd.Source,
			"destination": sd.Destination,
		}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToSourceDestNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"isNil": nilFT.SourceDestination() == nil}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}

func Test_FromTo_Rename(t *testing.T) {
	// Arrange
	{
		tc := fromToRenameMapsTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "new"}
		rn := ft.Rename()

	// Act
		actual := args.Map{
			"isNotNil": rn != nil,
			"existing": rn.Existing,
			"newName":  rn.New,
		}

	// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToRenameNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"isNil": nilFT.Rename() == nil}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}
