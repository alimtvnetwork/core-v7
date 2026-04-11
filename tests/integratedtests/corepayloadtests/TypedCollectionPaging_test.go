package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// createNumberedUsers is defined in shared_typed_helpers.go

// =============================================================================
// Tests: GetPagesSize
// =============================================================================

func Test_TypedPayloadCollection_GetPagesSize_FromTypedCollectionPagin(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pageSize, _ := input.GetAsInt("pageSize")

		count := 3 // default from createTestCollection (Alice, Bob, Carol)
		if c, ok := input.GetAsInt("count"); ok {
			count = c
		}

		collection := createNumberedUsers(count)

		// Act
		pagesSize := collection.GetPagesSize(pageSize)

		// Assert
		actual := args.Map{
			"pagesSize": pagesSize,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetSinglePageCollection
// =============================================================================

func Test_TypedPayloadCollection_GetSinglePageCollection(t *testing.T) {
	for caseIndex, testCase := range typedCollectionSinglePageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")
		pageIndex, _ := input.GetAsInt("pageIndex")

		collection := createNumberedUsers(count)

		// Act
		page := collection.GetSinglePageCollection(pageSize, pageIndex)

		actual := args.Map{
			"pageItemCount": page.Length(),
		}
		for i, item := range page.Items() {
			actual[fmt.Sprintf("item%d", i)] = item.Identifier()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetPagedCollection
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollection_FromTypedCollectionPagin(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollection(pageSize)

		actual := args.Map{
			"pageCount": len(pages),
		}
		for i, page := range pages {
			actual[fmt.Sprintf("page%dItems", i+1)] = page.Length()
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: GetPagedCollectionWithInfo
// =============================================================================

func Test_TypedPayloadCollection_GetPagedCollectionWithInfo_FromTypedCollectionPagin(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagedWithInfoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		pageSize, _ := input.GetAsInt("pageSize")

		collection := createNumberedUsers(count)

		// Act
		pages := collection.GetPagedCollectionWithInfo(pageSize)

		actual := args.Map{
			"pageCount": len(pages),
		}

		// Emit first 2 pages' PagingInfo fields
		for i := 0; i < 2 && i < len(pages); i++ {
			page := pages[i]
			prefix := fmt.Sprintf("p%d", i+1)
			actual[prefix+"CurrentPageIndex"] = page.Paging.CurrentPageIndex
			actual[prefix+"TotalPages"] = page.Paging.TotalPages
			actual[prefix+"PerPageItems"] = page.Paging.PerPageItems
			actual[prefix+"TotalItems"] = page.Paging.TotalItems
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Tests: Paging on empty collection
// =============================================================================

func Test_TypedPayloadCollection_Paging_Empty(t *testing.T) {
	for caseIndex, testCase := range typedCollectionPagingEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pageSize, _ := input.GetAsInt("pageSize")

		collection := corepayload.EmptyTypedPayloadCollection[testUser]()

		// Act
		pages := collection.GetPagedCollection(pageSize)

		actual := args.Map{
			"pageCount":  len(pages),
			"page1Items": pages[0].Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
