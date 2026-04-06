package coreinstructiontests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_FromTo_ClonePtr(t *testing.T) {
	// Case 0: positive
	{
		tc := fromToClonePtrCopiesTestCase
		orig := &coreinstruction.FromTo{From: "source", To: "destination"}
		cloned := orig.ClonePtr()

		actual := args.Map{
			"isNotNil": cloned != nil,
			"from":     cloned.From,
			"to":       cloned.To,
		}

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
	tc := fromToCloneCopiesTestCase
	orig := coreinstruction.FromTo{From: "a", To: "b"}
	c := orig.Clone()

	actual := args.Map{
		"from": c.From,
		"to":   c.To,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_IsNull(t *testing.T) {
	{
		tc := fromToIsNullNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"result": nilFT.IsNull()}

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
	{
		tc := fromToIsFromEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "", To: "dest"}

		actual := args.Map{"result": ft.IsFromEmpty()}

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
	{
		tc := fromToIsToEmptyEmptyTestCase
		ft := &coreinstruction.FromTo{From: "src", To: ""}

		actual := args.Map{"result": ft.IsToEmpty()}

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
	tc := fromToStringContainsTestCase
	ft := coreinstruction.FromTo{From: "alpha", To: "beta"}
	s := ft.String()

	actual := args.Map{
		"containsFrom": len(s) > 0 && strings.Contains(s, "alpha"),
		"containsTo":   strings.Contains(s, "beta"),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_Names(t *testing.T) {
	tc := fromToNamesTestCase
	ft := coreinstruction.FromTo{From: "src", To: "dst"}

	actual := args.Map{
		"fromName": ft.FromName(),
		"toName":   ft.ToName(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_SetFromName(t *testing.T) {
	{
		tc := fromToSetFromNameUpdatesTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "t"}
		ft.SetFromName("new")

		actual := args.Map{"from": ft.From}

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
	tc := fromToSetToNameUpdatesTestCase
	ft := &coreinstruction.FromTo{From: "f", To: "old"}
	ft.SetToName("new")

	actual := args.Map{"to": ft.To}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FromTo_SourceDestination(t *testing.T) {
	{
		tc := fromToSourceDestMapsTestCase
		ft := &coreinstruction.FromTo{From: "src", To: "dst"}
		sd := ft.SourceDestination()

		actual := args.Map{
			"isNotNil":    sd != nil,
			"source":      sd.Source,
			"destination": sd.Destination,
		}

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
	{
		tc := fromToRenameMapsTestCase
		ft := &coreinstruction.FromTo{From: "old", To: "new"}
		rn := ft.Rename()

		actual := args.Map{
			"isNotNil": rn != nil,
			"existing": rn.Existing,
			"newName":  rn.New,
		}

		tc.ShouldBeEqualMapFirst(t, actual)
	}

	{
		tc := fromToRenameNilTestCase
		var nilFT *coreinstruction.FromTo

		actual := args.Map{"isNil": nilFT.Rename() == nil}

		tc.ShouldBeEqualMapFirst(t, actual)
	}
}
