package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// =============================================================================
// Helpers
// =============================================================================

func createTestUsers() []*corepayload.TypedPayloadWrapper[testUser] {
	users := []testUser{
		{Name: "Alice", Email: "alice@test.com", Age: 30},
		{Name: "Bob", Email: "bob@test.com", Age: 25},
		{Name: "Carol", Email: "carol@test.com", Age: 35},
	}

	wrappers := make([]*corepayload.TypedPayloadWrapper[testUser], 0, len(users))

	for i, user := range users {
		category := "senior"
		if user.Age < 30 {
			category = "junior"
		}

		typed, err := corepayload.TypedPayloadWrapperNameIdCategory[testUser](
			user.Name,
			fmt.Sprintf("usr-%d", i),
			category,
			user,
		)
		errcore.HandleErr(err)

		wrappers = append(wrappers, typed)
	}

	return wrappers
}

func createTestCollection() *corepayload.TypedPayloadCollection[testUser] {
	return corepayload.TypedPayloadCollectionFrom[testUser](createTestUsers())
}

// =============================================================================
// Tests
// =============================================================================

func Test_TypedPayloadCollection_Creation(t *testing.T) {
	for caseIndex, testCase := range typedCollectionCreationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetDirectLower("capacity").(int)

		// Act
		var collection *corepayload.TypedPayloadCollection[testUser]
		if capacity == 0 {
			collection = corepayload.EmptyTypedPayloadCollection[testUser]()
		} else {
			collection = corepayload.NewTypedPayloadCollection[testUser](capacity)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"length":  collection.Length(),
			"isEmpty": collection.IsEmpty(),
		})
	}
}

func Test_TypedPayloadCollection_Add(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		email, _ := input.GetAsString("email")
		age := input.GetDirectLower("age").(int)
		collection := corepayload.EmptyTypedPayloadCollection[testUser]()

		firstUser := testUser{Name: name, Email: email, Age: age}
		firstTyped, firstErr := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
			name, "id-1", firstUser,
		)
		errcore.HandleErr(firstErr)

		// Act
		collection.Add(firstTyped)

		actual := args.Map{
			"length":    collection.Length(),
			"isEmpty":   collection.IsEmpty(),
			"firstName": collection.First().Data().Name,
		}

		name2, hasSecond := input.GetAsString("name2")
		if hasSecond {
			email2, _ := input.GetAsString("email2")
			age2 := input.GetDirectLower("age2").(int)
			secondUser := testUser{Name: name2, Email: email2, Age: age2}
			secondTyped, secondErr := corepayload.TypedPayloadWrapperNameIdRecord[testUser](
				name2, "id-2", secondUser,
			)
			errcore.HandleErr(secondErr)
			collection.Add(secondTyped)

			actual = args.Map{
				"length":     collection.Length(),
				"isEmpty":    collection.IsEmpty(),
				"firstName":  collection.First().Data().Name,
				"secondName": collection.Last().Data().Name,
			}
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedPayloadCollection_FilterByData_FromTypedCollection(t *testing.T) {
	for caseIndex, testCase := range typedCollectionFilterTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		filtered := collection.FilterByData(func(user testUser) bool {
			return user.Age >= 30
		})

		actual := args.Map{
			"filteredCount": filtered.Length(),
		}

		filtered.ForEachData(func(index int, data testUser) {
			actual[fmt.Sprintf("match%d", index+1)] = data.Name
		})

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedPayloadCollection_MapData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionMapTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		names := corepayload.MapTypedPayloadData[testUser, string](
			collection,
			func(user testUser) string { return user.Name },
		)

		actual := args.Map{
			"count": len(names),
		}

		for i, name := range names {
			actual[fmt.Sprintf("name%d", i)] = name
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedPayloadCollection_ReduceData(t *testing.T) {
	for caseIndex, testCase := range typedCollectionReduceTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		totalAge := corepayload.ReduceTypedPayloadData[testUser, int](
			collection,
			0,
			func(acc int, user testUser) int { return acc + user.Age },
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"totalAge": totalAge,
		})
	}
}

func Test_TypedPayloadCollection_GroupByCategory(t *testing.T) {
	for caseIndex, testCase := range typedCollectionGroupTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		groups := corepayload.GroupTypedPayloads[testUser, string](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) string {
				return item.CategoryName()
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"groupCount":      len(groups),
			"juniorGroupSize": groups["junior"].Length(),
			"seniorGroupSize": groups["senior"].Length(),
		})
	}
}

func Test_TypedPayloadCollection_Partition(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPartitionTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		senior, junior := corepayload.PartitionTypedPayloads[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Age >= 30
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"seniorCount": senior.Length(),
			"juniorCount": junior.Length(),
		})
	}
}

func Test_TypedPayloadCollection_AllData_FromTypedCollection(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAllDataTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		allData := collection.AllData()

		actual := args.Map{
			"count": len(allData),
		}

		for i, user := range allData {
			actual[fmt.Sprintf("data%d", i)] = user.Name
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedPayloadCollection_ElementAccess(t *testing.T) {
	for caseIndex, testCase := range typedCollectionElementAccessTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		firstName := collection.First().Data().Name
		lastName := collection.Last().Data().Name

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"firstName": firstName,
			"lastName":  lastName,
		})
	}
}

func Test_TypedPayloadCollection_AnyAll(t *testing.T) {
	for caseIndex, testCase := range typedCollectionAnyAllTestCases {
		// Arrange
		collection := createTestCollection()

		// Act
		hasBob := corepayload.AnyTypedPayload[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Name == "Bob"
			},
		)

		hasNonExistent := corepayload.AnyTypedPayload[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.Data().Name == "Nonexistent"
			},
		)

		allParsed := corepayload.AllTypedPayloads[testUser](
			collection,
			func(item *corepayload.TypedPayloadWrapper[testUser]) bool {
				return item.IsParsed()
			},
		)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"anyBob":         hasBob,
			"anyNonexistent": hasNonExistent,
			"allAreParsed":   allParsed,
		})
	}
}

func Test_TypedPayloadCollection_EmptyBehavior(t *testing.T) {
	// Arrange
	tc := typedCollectionEmptyOpsTestCase
	collection := corepayload.EmptyTypedPayloadCollection[testUser]()

	allData := collection.AllData()
	names := corepayload.MapTypedPayloadData[testUser, string](
		collection, func(u testUser) string { return u.Name },
	)
	filtered := collection.FilterByData(func(u testUser) bool { return true })
	totalAge := corepayload.ReduceTypedPayloadData[testUser, int](
		collection, 0, func(acc int, u testUser) int { return acc + u.Age },
	)

	// Act
	actual := args.Map{
		"allDataLen":  len(allData),
		"namesLen":    len(names),
		"filteredLen": filtered.Length(),
		"totalAge":    totalAge,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TypedPayloadCollection_FirstByName_FromTypedCollection(t *testing.T) {
	// Arrange
	tc := typedCollectionFirstByNameTestCase
	collection := createTestCollection()

	found := collection.FirstByName("Bob")
	notFound := collection.FirstByName("Nonexistent")

	// Act
	actual := args.Map{
		"foundName":   found.Data().Name,
		"notFoundNil": notFound == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TypedPayloadCollection_RemoveAt_FromTypedCollection(t *testing.T) {
	// Arrange
	tc := typedCollectionRemoveAtTestCase
	collection := createTestCollection()

	removed := collection.RemoveAt(1)
	invalidRemove := collection.RemoveAt(99)

	// Act
	actual := args.Map{
		"removed":       removed,
		"lengthAfter":   collection.Length(),
		"firstName":     collection.First().Data().Name,
		"lastName":      collection.Last().Data().Name,
		"invalidRemove": invalidRemove,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TypedPayloadCollection_ToPayloadsCollection_FromTypedCollection(t *testing.T) {
	// Arrange
	tc := typedCollectionToPayloadsTestCase
	collection := createTestCollection()
	payloads := collection.ToPayloadsCollection()

	// Act
	actual := args.Map{
		"length":    payloads.Length(),
		"firstName": payloads.First().Name,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
