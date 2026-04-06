package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// getBool extracts a boolean flag from the input map, defaulting to false.
func getBool(input args.Map, key string) bool {
	return input.GetAsBoolDefault(key, false)
}

// buildMapFromInput constructs a *MapAnyItems from the test case input.
func buildMapFromInput(input args.Map) *coredynamic.MapAnyItems {
	if getBool(input, "leftNil") {
		return nil
	}

	leftMap := input["leftMap"].(map[string]any)

	return coredynamic.NewMapAnyItemsUsingItems(leftMap)
}

// ==========================================
// IsEqual — named tests
// ==========================================

func Test_MapAnyItems_IsEqual_BothNil(t *testing.T) {
	tc := mapAnyItemsIsEqualBothNilTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	var left, right *coredynamic.MapAnyItems
	if !getBool(input, "leftNil") {
		left = coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	}
	if !getBool(input, "rightNil") {
		right = coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))
	}

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_LeftNil(t *testing.T) {
	tc := mapAnyItemsIsEqualLeftNilTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	var left *coredynamic.MapAnyItems
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_RightNil(t *testing.T) {
	tc := mapAnyItemsIsEqualRightNilTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(nil)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_SameContent(t *testing.T) {
	tc := mapAnyItemsIsEqualSameContentTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_DiffValues(t *testing.T) {
	tc := mapAnyItemsIsEqualDiffValuesTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_DiffKeys(t *testing.T) {
	tc := mapAnyItemsIsEqualDiffKeysTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_DiffLengths(t *testing.T) {
	tc := mapAnyItemsIsEqualDiffLengthsTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_IsEqual_BothEmpty(t *testing.T) {
	tc := mapAnyItemsIsEqualBothEmptyTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	right := coredynamic.NewMapAnyItemsUsingItems(input["rightMap"].(map[string]any))

	// Act
	actual := args.Map{"isEqual": left.IsEqual(right)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// IsEqualRaw — named tests
// ==========================================

// Note: IsEqualRaw nil receiver tests migrated to NilReceiver_test.go using CaseNilSafe pattern.

func Test_MapAnyItems_IsEqualRaw_Matching(t *testing.T) {
	tc := mapAnyItemsIsEqualRawMatchingTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	rawMap := input["rightMap"].(map[string]any)

	// Act
	actual := args.Map{"isEqualRaw": m.IsEqualRaw(rawMap)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// ClonePtr — named tests
// ==========================================

// Note: ClonePtr nil receiver test migrated to NilReceiver_test.go using CaseNilSafe pattern.

func Test_MapAnyItems_ClonePtr_ValidData(t *testing.T) {
	tc := mapAnyItemsClonePtrValidTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))

	// Act
	clone, err := m.ClonePtr()

	actual := args.Map{
		"hasError":    err != nil,
		"cloneIsNil":  clone == nil,
		"cloneLength": clone.Length(),
		"hasName":     clone.HasKey("name"),
		"hasAge":      clone.HasKey("age"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_ClonePtr_Empty(t *testing.T) {
	tc := mapAnyItemsClonePtrEmptyTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))

	// Act
	clone, err := m.ClonePtr()

	actual := args.Map{
		"hasError":    err != nil,
		"cloneIsNil":  clone == nil,
		"cloneLength": clone.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_ClonePtr_Independence(t *testing.T) {
	tc := mapAnyItemsClonePtrIndependenceTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))

	// Act
	clone, err := m.ClonePtr()
	clone.Add("new_key", "new_val")

	actual := args.Map{
		"hasError":          err != nil,
		"cloneIsNil":        clone == nil,
		"originalHasNewKey": m.HasKey("new_key"),
		"cloneHasNewKey":    clone.HasKey("new_key"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Edge cases — named tests
// ==========================================

// Note: Length/IsEmpty/HasAnyItem/HasKey nil receiver tests migrated to NilReceiver_test.go using CaseNilSafe pattern.

func Test_MapAnyItems_HasKey_Exists(t *testing.T) {
	tc := mapAnyItemsHasKeyExistsTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	key, _ := input.GetAsString("key")

	// Act
	actual := args.Map{"hasKey": m.HasKey(key)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_HasKey_Missing(t *testing.T) {
	tc := mapAnyItemsHasKeyMissingTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	key, _ := input.GetAsString("key")

	// Act
	actual := args.Map{"hasKey": m.HasKey(key)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_Add_NewKey(t *testing.T) {
	tc := mapAnyItemsAddNewKeyTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	addKey := input["addKey"].(string)
	addValue := input["addValue"]

	// Act
	isNew := m.Add(addKey, addValue)

	actual := args.Map{
		"isNew":       isNew,
		"lengthAfter": m.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_Add_ExistingKey(t *testing.T) {
	tc := mapAnyItemsAddExistingKeyTestCase
	input := tc.ArrangeInput.(args.Map)

	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(input["leftMap"].(map[string]any))
	addKey := input["addKey"].(string)
	addValue := input["addValue"]

	// Act
	isNew := m.Add(addKey, addValue)

	actual := args.Map{
		"isNew":        isNew,
		"updatedValue": m.GetValue(addKey),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
