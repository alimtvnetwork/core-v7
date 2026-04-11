package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// =============================================================================
// Helpers
// =============================================================================

func createTypedProduct(input args.Map) *corepayload.TypedPayloadWrapper[testProduct] {
	name, _ := input.GetAsString("name")
	id, _ := input.GetAsString("id")
	sku, _ := input.GetAsString("sku")
	title, _ := input.GetAsString("title")
	price := input.GetDirectLower("price").(float64)

	product := testProduct{SKU: sku, Title: title, Price: price}
	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](name, id, product)
	errcore.HandleErr(err)

	return typed
}

// =============================================================================
// Tests
// =============================================================================

func Test_TypedPayloadWrapper_Deserialization(t *testing.T) {
	for caseIndex, testCase := range typedWrapperDeserializationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typed := createTypedProduct(input)

		// Act
		serialized, serializeErr := typed.Serialize()
		errcore.HandleErr(serializeErr)

		deserialized, deserializeErr := corepayload.TypedPayloadWrapperDeserialize[testProduct](serialized)
		errcore.HandleErr(deserializeErr)

		data := deserialized.Data()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"name":  deserialized.Name(),
			"id":    deserialized.Identifier(),
			"sku":   data.SKU,
			"title": data.Title,
			"price": fmt.Sprintf("%.2f", data.Price),
		})
	}
}

func Test_TypedPayloadWrapper_RoundTrip(t *testing.T) {
	for caseIndex, testCase := range typedWrapperRoundTripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original := createTypedProduct(input)

		// Act
		jsonResult := original.JsonPtr()
		restored, restoreErr := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[testProduct](jsonResult)
		errcore.HandleErr(restoreErr)

		restoredData := restored.Data()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"name":  restored.Name(),
			"id":    restored.Identifier(),
			"sku":   restoredData.SKU,
			"title": restoredData.Title,
			"price": fmt.Sprintf("%.2f", restoredData.Price),
		})
	}
}

func Test_TypedPayloadWrapper_DeepClone(t *testing.T) {
	for caseIndex, testCase := range typedWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original := createTypedProduct(input)

		// Act
		cloned, cloneErr := original.ClonePtr(true)
		errcore.HandleErr(cloneErr)

		mutatedProduct := testProduct{
			SKU:   cloned.Data().SKU,
			Title: "Modified",
			Price: cloned.Data().Price,
		}
		setErr := cloned.SetTypedData(mutatedProduct)
		errcore.HandleErr(setErr)

		originalData := original.Data()
		clonedData := cloned.Data()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"originalName":  original.Name(),
			"originalId":    original.Identifier(),
			"originalSku":   originalData.SKU,
			"originalTitle": originalData.Title,
			"originalPrice": fmt.Sprintf("%.2f", originalData.Price),
			"clonedTitle":   clonedData.Title,
		})
	}
}

func Test_TypedPayloadWrapper_SetTypedData_FromTypedWrapper(t *testing.T) {
	for caseIndex, testCase := range typedWrapperSetDataTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		typed := createTypedProduct(input)

		newTitle, _ := input.GetAsString("new_title")
		newPrice := input.GetDirectLower("new_price").(float64)

		// Act
		updatedProduct := testProduct{
			SKU:   typed.Data().SKU,
			Title: newTitle,
			Price: newPrice,
		}
		setErr := typed.SetTypedData(updatedProduct)
		errcore.HandleErr(setErr)

		reparsed, reparseErr := corepayload.NewTypedPayloadWrapper[testProduct](typed.ToPayloadWrapper())
		errcore.HandleErr(reparseErr)

		directData := typed.Data()
		reparsedData := reparsed.Data()

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"directTitle":   directData.Title,
			"directPrice":   fmt.Sprintf("%.2f", directData.Price),
			"reparsedTitle": reparsedData.Title,
			"reparsedPrice": fmt.Sprintf("%.2f", reparsedData.Price),
		})
	}
}

func Test_TypedPayloadWrapper_NilWrapper(t *testing.T) {
	tc := typedWrapperNilTestCase

	// Act
	_, err := corepayload.NewTypedPayloadWrapper[testProduct](nil)

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"hasError": err != nil,
	})
}

func Test_TypedPayloadWrapper_InvalidJson(t *testing.T) {
	tc := typedWrapperInvalidJsonTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	invalidBytes, _ := input.GetAsString("bytes")

	// Act
	_, err := corepayload.TypedPayloadWrapperDeserialize[testProduct]([]byte(invalidBytes))

	// Assert
	tc.ShouldBeEqualMap(t, 0, args.Map{
		"hasError": err != nil,
	})
}

func Test_TypedPayloadWrapper_DeserializeToMany(t *testing.T) {
	for caseIndex, testCase := range typedWrapperDeserializeToManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetDirectLower("count").(int)

		wrappers := make([]*corepayload.TypedPayloadWrapper[testProduct], 0, count)

		for i := 0; i < count; i++ {
			product := testProduct{
				SKU:   fmt.Sprintf("SKU-%d", i),
				Title: fmt.Sprintf("item-%d", i),
				Price: float64(i) * 10.0,
			}

			typed, createErr := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](
				fmt.Sprintf("item-%d", i),
				fmt.Sprintf("id-%d", i),
				product,
			)
			errcore.HandleErr(createErr)
			wrappers = append(wrappers, typed)
		}

		payloadWrappers := make([]*corepayload.PayloadWrapper, len(wrappers))

		for i, w := range wrappers {
			payloadWrappers[i] = w.ToPayloadWrapper()
		}

		jsonSlice := corejson.Serialize.Apply(payloadWrappers)
		jsonSlice.HandleError()

		// Act
		deserialized, deserializeErr := corepayload.TypedPayloadWrapperDeserializeToMany[testProduct](jsonSlice.Bytes)
		errcore.HandleErr(deserializeErr)

		actual := args.Map{
			"count": len(deserialized),
		}

		for i, item := range deserialized {
			actual[fmt.Sprintf("title%d", i)] = item.Data().Title
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedPayloadWrapper_MetadataAccessors(t *testing.T) {
	// Arrange
	tc := typedWrapperMetadataAccessorsTestCase
	product := testProduct{SKU: "META-1", Title: "Meta Test", Price: 42.0}
	typed, err := corepayload.TypedPayloadWrapperNameIdCategory[testProduct](
		"meta-name", "meta-id", "meta-category", product,
	)
	errcore.HandleErr(err)

	// Act
	actual := args.Map{
		"name":             typed.Name(),
		"identifier":       typed.Identifier(),
		"categoryName":     typed.CategoryName(),
		"isParsed":         typed.IsParsed(),
		"isEmpty":          typed.IsEmpty(),
		"hasError":         typed.HasError(),
		"hasSingleRecord":  typed.HasSingleRecord(),
		"payloadsNonEmpty": typed.PayloadsString() != "",
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TypedPayloadWrapper_TypedDataJson(t *testing.T) {
	// Arrange
	tc := typedWrapperTypedDataJsonTestCase
	product := testProduct{SKU: "JSON-1", Title: "Json Test", Price: 99.99}
	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[testProduct](
		"json-test", "jt-1", product,
	)
	errcore.HandleErr(err)

	dataJson := typed.TypedDataJson()
	dataJsonPtr := typed.TypedDataJsonPtr()
	jsonBytes, jsonErr := typed.TypedDataJsonBytes()

	// Act
	actual := args.Map{
		"dataJsonNonEmpty":  !dataJson.IsEmpty(),
		"dataJsonPtrNonNil": dataJsonPtr != nil && !dataJsonPtr.IsEmpty(),
		"jsonBytesNonEmpty": len(jsonBytes) > 0,
		"jsonBytesNoError":  jsonErr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
