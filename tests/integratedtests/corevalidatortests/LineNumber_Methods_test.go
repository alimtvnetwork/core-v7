package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_LineNumber_Methods(t *testing.T) {
	// Arrange
	ln := corevalidator.LineNumber{LineNumber: 5}

	// Act
	actual := args.Map{
		"has": ln.HasLineNumber(),
		"isMatch": ln.IsMatch(5),
		"mismatchErr": ln.VerifyError(6) != nil,
	}

	// Assert
	expected := args.Map{
		"has": true,
		"isMatch": true,
		"mismatchErr": true,
	}
	expected.ShouldBeEqual(t, 0, "LineNumber returns correct value -- methods", actual)
}

func Test_Condition_Methods(t *testing.T) {
	// Arrange
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"splitByWhitespace": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"splitByWhitespace": true}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- methods", actual)
}

func Test_Parameter_Methods(t *testing.T) {
	// Arrange
	p := corevalidator.Parameter{CaseIndex: 1, Header: "h", IsCaseSensitive: false}

	// Act
	actual := args.Map{
		"caseIndex": p.CaseIndex,
		"header": p.Header,
		"ignoreCase": p.IsIgnoreCase(),
	}

	// Assert
	expected := args.Map{
		"caseIndex": 1,
		"header": "h",
		"ignoreCase": true,
	}
	expected.ShouldBeEqual(t, 0, "Parameter returns correct value -- methods", actual)
}

func Test_RangesSegment(t *testing.T) {
	// Arrange
	rs := corevalidator.RangesSegment{RangeInt: corerange.RangeInt{Start: 0, End: 10}, ExpectedLines: []string{"a"}, CompareAs: stringcompareas.Equal}

	// Act
	actual := args.Map{
		"start": rs.Start,
		"end": rs.End,
		"expLen": len(rs.ExpectedLines),
		"compareAs": rs.CompareAs.Name(),
	}

	// Assert
	expected := args.Map{
		"start": 0,
		"end": 10,
		"expLen": 1,
		"compareAs": "Equal",
	}
	expected.ShouldBeEqual(t, 0, "RangesSegment returns correct value -- struct", actual)
}

func Test_BaseValidatorCoreCondition_Default(t *testing.T) {
	// Arrange
	bv := corevalidator.BaseValidatorCoreCondition{}
	cond := bv.ValidatorCoreConditionDefault()

	// Act
	actual := args.Map{
		"notNil": bv.ValidatorCoreCondition != nil,
		"split": cond.IsSplitByWhitespace(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"split": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseValidatorCoreCondition returns non-empty -- default", actual)
}
