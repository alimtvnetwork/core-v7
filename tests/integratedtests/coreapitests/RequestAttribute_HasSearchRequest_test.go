package coreapitests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// RequestAttribute — uncovered branches
// =============================================================================

func Test_RequestAttribute_HasSearchRequest(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{SearchRequest: &coreapi.SearchRequest{}}

	// Act
	actual := args.Map{"result": attr.HasSearchRequest()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have search request", actual)
}

func Test_RequestAttribute_HasSearchRequest_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.HasSearchRequest()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have search request", actual)
}

func Test_RequestAttribute_HasPageRequest(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{PageRequest: &coreapi.PageRequest{PageSize: 10}}

	// Act
	actual := args.Map{"result": attr.HasPageRequest()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page request", actual)
}

func Test_RequestAttribute_HasPageRequest_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.HasPageRequest()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page request", actual)
}

func Test_RequestAttribute_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_RequestAttribute_IsAnyNull_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.IsAnyNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_RequestAttribute_IsPageRequestEmpty(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{}

	// Act
	actual := args.Map{"result": attr.IsPageRequestEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty without page request", actual)
}

func Test_RequestAttribute_IsPageRequestEmpty_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.IsPageRequestEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_RequestAttribute_IsSearchRequestEmpty(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{}

	// Act
	actual := args.Map{"result": attr.IsSearchRequestEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty without search request", actual)
}

func Test_RequestAttribute_IsSearchRequestEmpty_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.IsSearchRequestEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_RequestAttribute_Clone(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{
		Url:          "http://test",
		IsValid:      true,
		SearchRequest: &coreapi.SearchRequest{SearchTerm: "test"},
		PageRequest:   &coreapi.PageRequest{PageSize: 10},
	}
	c := attr.Clone()

	// Act
	actual := args.Map{"result": c == nil || c.Url != "http://test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_RequestAttribute_Clone_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.RequestAttribute

	// Act
	actual := args.Map{"result": attr.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// =============================================================================
// ResponseAttribute — uncovered branches
// =============================================================================

func Test_ResponseAttribute_Clone_WithSlices(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{
		IsValid:        true,
		Message:        "ok",
		StepsPerformed: []string{"step1"},
		DebugInfos:     []string{"debug1"},
	}
	c := attr.Clone()

	// Act
	actual := args.Map{"result": c == nil || len(c.StepsPerformed) != 1 || len(c.DebugInfos) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone with slices", actual)
}

func Test_ResponseAttribute_Clone_Nil(t *testing.T) {
	// Arrange
	var attr *coreapi.ResponseAttribute

	// Act
	actual := args.Map{"result": attr.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_ResponseAttribute_Clone_EmptySlices(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{IsValid: true}
	c := attr.Clone()

	// Act
	actual := args.Map{"result": c == nil || c.StepsPerformed != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone without slices", actual)
}

// =============================================================================
// InvalidRequestAttribute / InvalidResponseAttribute
// =============================================================================

func Test_InvalidRequestAttribute(t *testing.T) {
	// Arrange
	attr := coreapi.InvalidRequestAttribute()

	// Act
	actual := args.Map{"result": attr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_InvalidResponseAttribute(t *testing.T) {
	// Arrange
	attr := coreapi.InvalidResponseAttribute("test error")

	// Act
	actual := args.Map{"result": attr.IsValid || attr.Message != "test error"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid with message", actual)
}

// =============================================================================
// SearchRequest — uncovered branches
// =============================================================================

func Test_SearchRequest_Clone(t *testing.T) {
	// Arrange
	sr := &coreapi.SearchRequest{SearchTerm: "test", IsContains: true}
	c := sr.Clone()

	// Act
	actual := args.Map{"result": c == nil || c.SearchTerm != "test" || !c.IsContains}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_SearchRequest_Clone_Nil(t *testing.T) {
	// Arrange
	var sr *coreapi.SearchRequest

	// Act
	actual := args.Map{"result": sr.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// =============================================================================
// PageRequest — uncovered branches
// =============================================================================

func Test_PageRequest_IsPageSizeEmpty_Nil(t *testing.T) {
	// Arrange
	var pr *coreapi.PageRequest

	// Act
	actual := args.Map{"result": pr.IsPageSizeEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_PageRequest_IsPageIndexEmpty_Nil(t *testing.T) {
	// Arrange
	var pr *coreapi.PageRequest

	// Act
	actual := args.Map{"result": pr.IsPageIndexEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_PageRequest_HasPageSize_FromRequestAttributeHasS(t *testing.T) {
	// Arrange
	pr := &coreapi.PageRequest{PageSize: 10}

	// Act
	actual := args.Map{"result": pr.HasPageSize()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page size", actual)
}

func Test_PageRequest_HasPageSize_Nil(t *testing.T) {
	// Arrange
	var pr *coreapi.PageRequest

	// Act
	actual := args.Map{"result": pr.HasPageSize()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page size", actual)
}

func Test_PageRequest_HasPageIndex_FromRequestAttributeHasS(t *testing.T) {
	// Arrange
	pr := &coreapi.PageRequest{PageIndex: 5}

	// Act
	actual := args.Map{"result": pr.HasPageIndex()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page index", actual)
}

func Test_PageRequest_HasPageIndex_Nil(t *testing.T) {
	// Arrange
	var pr *coreapi.PageRequest

	// Act
	actual := args.Map{"result": pr.HasPageIndex()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page index", actual)
}

func Test_PageRequest_Clone_Nil_FromRequestAttributeHasS(t *testing.T) {
	// Arrange
	var pr *coreapi.PageRequest

	// Act
	actual := args.Map{"result": pr.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_PageRequest_Clone(t *testing.T) {
	// Arrange
	pr := &coreapi.PageRequest{PageSize: 10, PageIndex: 2}
	c := pr.Clone()

	// Act
	actual := args.Map{"result": c.PageSize != 10 || c.PageIndex != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

// =============================================================================
// TypedResponse — uncovered branches
// =============================================================================

func Test_TypedResponse_Clone_Nil(t *testing.T) {
	// Arrange
	var resp *coreapi.TypedResponse[string]

	// Act
	actual := args.Map{"result": resp.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_TypedResponse_TypedResponseResult_Nil(t *testing.T) {
	// Arrange
	var resp *coreapi.TypedResponse[string]

	// Act
	actual := args.Map{"result": resp.TypedResponseResult() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TypedResponse_TypedResponseResult(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{IsValid: true}
	resp := coreapi.NewTypedResponse(attr, "hello")
	result := resp.TypedResponseResult()

	// Act
	actual := args.Map{"result": result == nil || result.Response != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_InvalidTypedResponse_NilAttribute(t *testing.T) {
	// Arrange
	resp := coreapi.InvalidTypedResponse[string](nil)

	// Act
	actual := args.Map{"result": resp.Attribute == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedResponseResult — uncovered branches
// =============================================================================

func Test_TypedResponseResult_IsValid(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")

	// Act
	actual := args.Map{"result": rr.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_TypedResponseResult_IsValid_Nil(t *testing.T) {
	// Arrange
	var rr *coreapi.TypedResponseResult[string]

	// Act
	actual := args.Map{"result": rr.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_TypedResponseResult_IsInvalid(t *testing.T) {
	// Arrange
	rr := coreapi.InvalidTypedResponseResult[string](nil)

	// Act
	actual := args.Map{"result": rr.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_TypedResponseResult_Message_Nil(t *testing.T) {
	// Arrange
	var rr *coreapi.TypedResponseResult[string]

	// Act
	actual := args.Map{"result": rr.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_TypedResponseResult_Message(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{Message: "ok"}
	rr := coreapi.NewTypedResponseResult(attr, "data")

	// Act
	actual := args.Map{"result": rr.Message() != "ok"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return message", actual)
}

func Test_TypedResponseResult_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var rr *coreapi.TypedResponseResult[string]

	// Act
	actual := args.Map{"result": rr.ClonePtr() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_TypedResponseResult_ClonePtr(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	c := rr.ClonePtr()

	// Act
	actual := args.Map{"result": c == nil || c.Response != "data"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_TypedResponseResult_ToTypedResponse_Nil(t *testing.T) {
	// Arrange
	var rr *coreapi.TypedResponseResult[string]

	// Act
	actual := args.Map{"result": rr.ToTypedResponse() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TypedResponseResult_ToTypedResponse(t *testing.T) {
	// Arrange
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	resp := rr.ToTypedResponse()

	// Act
	actual := args.Map{"result": resp == nil || resp.Response != "data"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_InvalidTypedResponseResult_NilAttribute(t *testing.T) {
	// Arrange
	rr := coreapi.InvalidTypedResponseResult[string](nil)

	// Act
	actual := args.Map{"result": rr.Attribute == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedRequest — uncovered branches
// =============================================================================

func Test_TypedRequest_Clone_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedRequest[string]

	// Act
	actual := args.Map{"result": req.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_TypedRequest_ToTypedSimpleGenericRequest_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedRequest[string]

	// Act
	actual := args.Map{"result": req.ToTypedSimpleGenericRequest(true, "") != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TypedRequest_ToTypedSimpleGenericRequest(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequest(attr, "payload")
	sgr := req.ToTypedSimpleGenericRequest(true, "")

	// Act
	actual := args.Map{"result": sgr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_InvalidTypedRequest_NilAttribute(t *testing.T) {
	// Arrange
	req := coreapi.InvalidTypedRequest[string](nil)

	// Act
	actual := args.Map{"result": req.Attribute == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedRequestIn — uncovered branches
// =============================================================================

func Test_TypedRequestIn_Clone_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedRequestIn[string]

	// Act
	actual := args.Map{"result": req.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_TypedRequestIn_TypedSimpleGenericRequest_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedRequestIn[string]

	// Act
	actual := args.Map{"result": req.TypedSimpleGenericRequest(true, "") != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TypedRequestIn_TypedSimpleGenericRequest(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequestIn(attr, "payload")
	sgr := req.TypedSimpleGenericRequest(true, "")

	// Act
	actual := args.Map{"result": sgr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_InvalidTypedRequestIn_NilAttribute(t *testing.T) {
	// Arrange
	req := coreapi.InvalidTypedRequestIn[string](nil)

	// Act
	actual := args.Map{"result": req.Attribute == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedSimpleGenericRequest — uncovered branches
// =============================================================================

func Test_TypedSimpleGenericRequest_IsValid_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actual := args.Map{"result": req.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_TypedSimpleGenericRequest_IsValid_NilRequest(t *testing.T) {
	// Arrange
	req := &coreapi.TypedSimpleGenericRequest[string]{
		Attribute: &coreapi.RequestAttribute{IsValid: true},
	}

	// Act
	actual := args.Map{"result": req.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should be invalid", actual)
}

func Test_TypedSimpleGenericRequest_IsValid_True(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)

	// Act
	actual := args.Map{"result": req.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_TypedSimpleGenericRequest_IsInvalid(t *testing.T) {
	// Arrange
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)

	// Act
	actual := args.Map{"result": req.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_TypedSimpleGenericRequest_Message_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actual := args.Map{"result": req.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_TypedSimpleGenericRequest_Message_NilRequest(t *testing.T) {
	// Arrange
	req := &coreapi.TypedSimpleGenericRequest[string]{}

	// Act
	actual := args.Map{"result": req.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should return empty", actual)
}

func Test_TypedSimpleGenericRequest_InvalidError_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actual := args.Map{"result": req.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_TypedSimpleGenericRequest_InvalidError_NilRequest(t *testing.T) {
	// Arrange
	req := &coreapi.TypedSimpleGenericRequest[string]{}

	// Act
	actual := args.Map{"result": req.InvalidError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should return nil", actual)
}

func Test_TypedSimpleGenericRequest_Clone_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actual := args.Map{"result": req.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_TypedSimpleGenericRequest_Clone(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	c := req.Clone()

	// Act
	actual := args.Map{"result": c == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_InvalidTypedSimpleGenericRequest_NilAttribute(t *testing.T) {
	// Arrange
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)

	// Act
	actual := args.Map{"result": req.Attribute == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

func Test_TypedSimpleGenericRequest_Data(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("hello", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)

	// Act
	actual := args.Map{"result": req.Data() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return data", actual)
}
