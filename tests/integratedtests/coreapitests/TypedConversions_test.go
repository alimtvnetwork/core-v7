package coreapitests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: NewTypedSimpleGenericRequest
// ==========================================

func Test_NewTypedSimpleGenericRequest(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)

		// Act
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		actual := args.Map{
			"isValid":      fmt.Sprintf("%v", req.IsValid()),
			"hasAttribute": fmt.Sprintf("%v", req.Attribute.IsValid),
			"payload":      req.Data(),
			"hasRequest":   fmt.Sprintf("%v", req.Request != nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: InvalidTypedSimpleGenericRequest
// ==========================================

func Test_InvalidTypedSimpleGenericRequest(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestInvalidTestCases {
		// Act
		req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
		actual := args.Map{
			"isValid":      fmt.Sprintf("%v", req.IsValid()),
			"isInvalid":    fmt.Sprintf("%v", req.Attribute != nil),
			"isNilRequest": fmt.Sprintf("%v", req.Request == nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest IsValid / IsInvalid
// ==========================================

func Test_TypedSimpleGenericRequest_Validity(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestValidityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequestVal, _ := input.Get("nilRequest")
		nilRequest, _ := nilRequestVal.(bool)
		invalidAttributeVal, _ := input.Get("invalidAttribute")
		invalidAttribute, _ := invalidAttributeVal.(bool)
		nilAttributeVal, _ := input.Get("nilAttribute")
		nilAttribute, _ := nilAttributeVal.(bool)

		var req *coreapi.TypedSimpleGenericRequest[string]

		var attr *coreapi.RequestAttribute

		if nilAttribute {
			attr = nil
		} else if invalidAttribute {
			attr = &coreapi.RequestAttribute{IsValid: false}
		} else {
			attr = &coreapi.RequestAttribute{IsValid: true}
		}

		if nilRequest {
			req = &coreapi.TypedSimpleGenericRequest[string]{
				Attribute: attr,
				Request:   nil,
			}
		} else {
			payload, _ := input.GetAsString("payload")
			simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
			req = coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		}

		// Act
		actual := args.Map{
			"isValid":   fmt.Sprintf("%v", req.IsValid()),
			"isInvalid": fmt.Sprintf("%v", req.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Message / InvalidError
// ==========================================

func Test_TypedSimpleGenericRequest_Message(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequestVal, _ := input.Get("nilRequest")
		nilRequest, _ := nilRequestVal.(bool)

		var req *coreapi.TypedSimpleGenericRequest[string]

		if nilRequest {
			req = coreapi.InvalidTypedSimpleGenericRequest[string](nil)
		} else {
			payload, _ := input.GetAsString("payload")
			message, _ := input.GetAsString("message")
			attr := &coreapi.RequestAttribute{IsValid: true}
			simpleReq := coredynamic.NewTypedSimpleRequest[string](payload, false, message)
			req = coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		}

		// Act
		actual := args.Map{
			"message":    req.Message(),
			"isNilError": fmt.Sprintf("%v", req.InvalidError() == nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver (CaseNilSafe pattern)
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestNilSafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestInvalidUnderlyingTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")
		message, _ := input.GetAsString("message")
		check, _ := input.GetAsString("check")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequest[string](payload, false, message)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		switch check {
		case "validity":
			actual := args.Map{
				"isValid":   fmt.Sprintf("%v", req.IsValid()),
				"isInvalid": fmt.Sprintf("%v", req.IsInvalid()),
			}
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		case "message":
			tc.ShouldBeEqual(t, caseIndex, req.Message())
		case "invalidError":
			actual = args.Map{
				"isNilError":   fmt.Sprintf("%v", req.InvalidError() == nil),
				"errorMessage": req.InvalidError().Error(),
			}
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		case "invalidErrorNil":
			actual = args.Map{
				"isNilError":   fmt.Sprintf("%v", req.InvalidError() == nil),
				"errorMessage": req.Message(),
			}
			tc.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Clone
// ==========================================

func Test_TypedSimpleGenericRequest_Clone_FromTypedConversions(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		cloned := req.Clone()
		actual := args.Map{
			"payload":       cloned.Data(),
			"isValid":       fmt.Sprintf("%v", cloned.IsValid()),
			"isIndependent": fmt.Sprintf("%v", cloned != req),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedSimpleGenericRequest_Clone_Nil_FromTypedConversions(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedSimpleGenericRequest[string]

		// Act
		cloned := req.Clone()
		actLines := []string{fmt.Sprintf("%v", cloned == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedRequestIn.TypedSimpleGenericRequest conversion
// ==========================================

func Test_TypedRequestIn_TypedSimpleGenericRequest_FromTypedConversions(t *testing.T) {
	for caseIndex, testCase := range typedRequestInToTypedSimpleGenericTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")
		isValid := true
		message := ""

		if vRaw, ok := input.Get("isValid"); ok {
			if v, isBool := vRaw.(bool); isBool {
				isValid = v
			}
		}

		if v, ok := input.GetAsString("message"); ok {
			message = v
		}

		attr := &coreapi.RequestAttribute{IsValid: true}
		reqIn := coreapi.NewTypedRequestIn[string](attr, payload)

		// Act
		tsgr := reqIn.TypedSimpleGenericRequest(isValid, message)

		// Build actual map — key depends on test case
		actual := args.Map{
			"isValid": fmt.Sprintf("%v", tsgr.Request.IsValid()),
			"payload": tsgr.Data(),
			"message": tsgr.Message(),
		}

		// Use "hasRequest" or "hasValidRequest" depending on expected keys
		expected := testCase.ExpectedInput.(args.Map)
		if _, ok := expected["hasRequest"]; ok {
			actual["hasRequest"] = fmt.Sprintf("%v", tsgr.Attribute.IsValid)
		}
		if _, ok := expected["hasValidRequest"]; ok {
			actual["hasValidRequest"] = fmt.Sprintf("%v", tsgr.Request.IsValid())
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedRequestIn_TypedSimpleGenericRequest_Nil_FromTypedConversions(t *testing.T) {
	for caseIndex, tc := range typedRequestInToTypedSimpleGenericNilTestCases {
		// Arrange
		var reqIn *coreapi.TypedRequestIn[string]

		// Act
		result := reqIn.TypedSimpleGenericRequest(true, "")
		actLines := []string{fmt.Sprintf("%v", result == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponse.TypedResponseResult conversion
// ==========================================

func Test_TypedResponse_TypedResponseResult_FromTypedConversions(t *testing.T) {
	for caseIndex, testCase := range typedResponseToTypedResponseResultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")
		message, _ := input.GetAsString("message")
		isValid := true

		if vRaw, ok := input.Get("isValid"); ok {
			if v, isBool := vRaw.(bool); isBool {
				isValid = v
			}
		}

		attr := &coreapi.ResponseAttribute{IsValid: isValid, Message: message}
		resp := coreapi.NewTypedResponse[string](attr, response)

		// Act
		result := resp.TypedResponseResult()
		actual := args.Map{
			"response": result.Response,
			"isValid":  fmt.Sprintf("%v", result.IsValid()),
			"message":  result.Message(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TypedResponse_TypedResponseResult_Nil_FromTypedConversions(t *testing.T) {
	for caseIndex, tc := range typedResponseToTypedResponseResultNilTestCases {
		// Arrange
		var resp *coreapi.TypedResponse[string]

		// Act
		result := resp.TypedResponseResult()
		actLines := []string{fmt.Sprintf("%v", result == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
