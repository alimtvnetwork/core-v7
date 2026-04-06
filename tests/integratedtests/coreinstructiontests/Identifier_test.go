package coreinstructiontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_BaseIdentifier_Verification(t *testing.T) {
	for caseIndex, testCase := range identifierTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		id, _ := input.GetAsString("id")

		// Act
		identifier := coreinstruction.NewIdentifier(id)
		actual := args.Map{
			"id":                  identifier.IdString(),
			"isEmpty":             identifier.IsIdEmpty(),
			"isEmptyOrWhitespace": identifier.IsIdWhitespace(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Identifiers_Length_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, ok := input.GetAsStrings("ids")
		if !ok {
			errcore.HandleErrMessage("ids required")
		}

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		actual := args.Map{
			"length":     identifiers.Length(),
			"isEmpty":    identifiers.IsEmpty(),
			"hasAnyItem": identifiers.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Identifiers_GetById_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersGetByIdTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, ok := input.GetAsStrings("ids")
		if !ok {
			errcore.HandleErrMessage("ids required")
		}
		searchId, _ := input.GetAsString("searchId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		result := identifiers.GetById(searchId)

		found := result != nil
		foundId := ""
		if found {
			foundId = result.Id
		}

		actual := args.Map{
			"found": found,
			"id":    foundId,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Identifiers_IndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersIndexOfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, ok := input.GetAsStrings("ids")
		if !ok {
			errcore.HandleErrMessage("ids required")
		}
		searchId, _ := input.GetAsString("searchId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		index := identifiers.IndexOf(searchId)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", index))
	}
}

func Test_Identifiers_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, ok := input.GetAsStrings("ids")
		if !ok {
			errcore.HandleErrMessage("ids required")
		}

		// Act
		original := coreinstruction.NewIdentifiers(ids...)
		cloned := original.Clone()

		actual := args.Map{
			"length": cloned.Length(),
		}
		for i, baseId := range cloned.Ids {
			actual[fmt.Sprintf("id%d", i)] = baseId.Id
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Identifiers_Add_Verification(t *testing.T) {
	for caseIndex, testCase := range identifiersAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ids, ok := input.GetAsStrings("ids")
		if !ok {
			errcore.HandleErrMessage("ids required")
		}
		addId, _ := input.GetAsString("addId")

		// Act
		identifiers := coreinstruction.NewIdentifiers(ids...)
		identifiers.Add(addId)

		actual := args.Map{
			"length": identifiers.Length(),
		}
		for i, baseId := range identifiers.Ids {
			actual[fmt.Sprintf("id%d", i)] = baseId.Id
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Specification_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range specificationCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		id, _ := input.GetAsString("id")
		display, _ := input.GetAsString("display")
		typeName, _ := input.GetAsString("typeName")
		tags, ok := input.GetAsStrings("tags")
		if !ok {
			errcore.HandleErrMessage("tags required")
		}
		isGlobalRaw, _ := input.Get("isGlobal")
		isGlobal, _ := isGlobalRaw.(bool)

		// Act
		spec := coreinstruction.NewSpecification(id, display, typeName, tags, isGlobal)
		cloned := spec.Clone()

		actual := args.Map{
			"id":        cloned.Id,
			"display":   cloned.Display,
			"typeName":  cloned.Type,
			"tagsCount": cloned.TagsLength(),
			"isGlobal":  cloned.IsGlobal,
		}
		for i, tag := range cloned.Tags {
			actual[fmt.Sprintf("tag%d", i)] = tag
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BaseTags_Verification(t *testing.T) {
	for caseIndex, testCase := range baseTagsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		tags, ok := input.GetAsStrings("tags")
		if !ok {
			errcore.HandleErrMessage("tags required")
		}
		searchTags, ok := input.GetAsStrings("searchTags")
		if !ok {
			errcore.HandleErrMessage("searchTags required")
		}

		// Act
		baseTags := coreinstruction.NewTags(tags)
		actual := args.Map{
			"tagsCount":  baseTags.TagsLength(),
			"isEmpty":    baseTags.IsTagsEmpty(),
			"hasAllTags": baseTags.HasAllTags(searchTags...),
			"hasAnyTag":  baseTags.HasAnyTags(searchTags...),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Specification_Clone_NilSafety(t *testing.T) {
	// Arrange
	tc := specificationCloneNilTestCase
	var spec *coreinstruction.Specification
	result := spec.Clone()

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Specification_Clone_DeepCopy_Tags(t *testing.T) {
	// Arrange
	tc := specificationCloneDeepCopyTestCase
	original := coreinstruction.NewSpecification("id", "display", "type", []string{"a", "b"}, false)
	cloned := original.Clone()

	cloned.Tags[0] = "MUTATED"

	// Act
	actual := args.Map{
		"originalTag0": original.Tags[0],
		"cloneTag0":    cloned.Tags[0],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
