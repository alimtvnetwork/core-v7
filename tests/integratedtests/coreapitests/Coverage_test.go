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

func Test_Cov_RequestAttribute_HasSearchRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{SearchRequest: &coreapi.SearchRequest{}}
	actual := args.Map{"result": attr.HasSearchRequest()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have search request", actual)
}

func Test_Cov_RequestAttribute_HasSearchRequest_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.HasSearchRequest()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have search request", actual)
}

func Test_Cov_RequestAttribute_HasPageRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{PageRequest: &coreapi.PageRequest{PageSize: 10}}
	actual := args.Map{"result": attr.HasPageRequest()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page request", actual)
}

func Test_Cov_RequestAttribute_HasPageRequest_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.HasPageRequest()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page request", actual)
}

func Test_Cov_RequestAttribute_IsEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov_RequestAttribute_IsAnyNull_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.IsAnyNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_Cov_RequestAttribute_IsPageRequestEmpty(t *testing.T) {
	attr := &coreapi.RequestAttribute{}
	actual := args.Map{"result": attr.IsPageRequestEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty without page request", actual)
}

func Test_Cov_RequestAttribute_IsPageRequestEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.IsPageRequestEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov_RequestAttribute_IsSearchRequestEmpty(t *testing.T) {
	attr := &coreapi.RequestAttribute{}
	actual := args.Map{"result": attr.IsSearchRequestEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty without search request", actual)
}

func Test_Cov_RequestAttribute_IsSearchRequestEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.IsSearchRequestEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov_RequestAttribute_Clone(t *testing.T) {
	attr := &coreapi.RequestAttribute{
		Url:          "http://test",
		IsValid:      true,
		SearchRequest: &coreapi.SearchRequest{SearchTerm: "test"},
		PageRequest:   &coreapi.PageRequest{PageSize: 10},
	}
	c := attr.Clone()
	actual := args.Map{"result": c == nil || c.Url != "http://test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_Cov_RequestAttribute_Clone_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	actual := args.Map{"result": attr.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// =============================================================================
// ResponseAttribute — uncovered branches
// =============================================================================

func Test_Cov_ResponseAttribute_Clone_WithSlices(t *testing.T) {
	attr := &coreapi.ResponseAttribute{
		IsValid:        true,
		Message:        "ok",
		StepsPerformed: []string{"step1"},
		DebugInfos:     []string{"debug1"},
	}
	c := attr.Clone()
	actual := args.Map{"result": c == nil || len(c.StepsPerformed) != 1 || len(c.DebugInfos) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone with slices", actual)
}

func Test_Cov_ResponseAttribute_Clone_Nil(t *testing.T) {
	var attr *coreapi.ResponseAttribute
	actual := args.Map{"result": attr.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_ResponseAttribute_Clone_EmptySlices(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	c := attr.Clone()
	actual := args.Map{"result": c == nil || c.StepsPerformed != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone without slices", actual)
}

// =============================================================================
// InvalidRequestAttribute / InvalidResponseAttribute
// =============================================================================

func Test_Cov_InvalidRequestAttribute(t *testing.T) {
	attr := coreapi.InvalidRequestAttribute()
	actual := args.Map{"result": attr.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Cov_InvalidResponseAttribute(t *testing.T) {
	attr := coreapi.InvalidResponseAttribute("test error")
	actual := args.Map{"result": attr.IsValid || attr.Message != "test error"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid with message", actual)
}

// =============================================================================
// SearchRequest — uncovered branches
// =============================================================================

func Test_Cov_SearchRequest_Clone(t *testing.T) {
	sr := &coreapi.SearchRequest{SearchTerm: "test", IsContains: true}
	c := sr.Clone()
	actual := args.Map{"result": c == nil || c.SearchTerm != "test" || !c.IsContains}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_Cov_SearchRequest_Clone_Nil(t *testing.T) {
	var sr *coreapi.SearchRequest
	actual := args.Map{"result": sr.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// =============================================================================
// PageRequest — uncovered branches
// =============================================================================

func Test_Cov_PageRequest_IsPageSizeEmpty_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	actual := args.Map{"result": pr.IsPageSizeEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov_PageRequest_IsPageIndexEmpty_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	actual := args.Map{"result": pr.IsPageIndexEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_Cov_PageRequest_HasPageSize(t *testing.T) {
	pr := &coreapi.PageRequest{PageSize: 10}
	actual := args.Map{"result": pr.HasPageSize()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page size", actual)
}

func Test_Cov_PageRequest_HasPageSize_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	actual := args.Map{"result": pr.HasPageSize()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page size", actual)
}

func Test_Cov_PageRequest_HasPageIndex(t *testing.T) {
	pr := &coreapi.PageRequest{PageIndex: 5}
	actual := args.Map{"result": pr.HasPageIndex()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have page index", actual)
}

func Test_Cov_PageRequest_HasPageIndex_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	actual := args.Map{"result": pr.HasPageIndex()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have page index", actual)
}

func Test_Cov_PageRequest_Clone_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	actual := args.Map{"result": pr.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_PageRequest_Clone(t *testing.T) {
	pr := &coreapi.PageRequest{PageSize: 10, PageIndex: 2}
	c := pr.Clone()
	actual := args.Map{"result": c.PageSize != 10 || c.PageIndex != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

// =============================================================================
// TypedResponse — uncovered branches
// =============================================================================

func Test_Cov_TypedResponse_Clone_Nil(t *testing.T) {
	var resp *coreapi.TypedResponse[string]
	actual := args.Map{"result": resp.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_TypedResponse_TypedResponseResult_Nil(t *testing.T) {
	var resp *coreapi.TypedResponse[string]
	actual := args.Map{"result": resp.TypedResponseResult() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Cov_TypedResponse_TypedResponseResult(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	resp := coreapi.NewTypedResponse(attr, "hello")
	result := resp.TypedResponseResult()
	actual := args.Map{"result": result == nil || result.Response != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_Cov_InvalidTypedResponse_NilAttribute(t *testing.T) {
	resp := coreapi.InvalidTypedResponse[string](nil)
	actual := args.Map{"result": resp.Attribute == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedResponseResult — uncovered branches
// =============================================================================

func Test_Cov_TypedResponseResult_IsValid(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	actual := args.Map{"result": rr.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_Cov_TypedResponseResult_IsValid_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	actual := args.Map{"result": rr.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov_TypedResponseResult_IsInvalid(t *testing.T) {
	rr := coreapi.InvalidTypedResponseResult[string](nil)
	actual := args.Map{"result": rr.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Cov_TypedResponseResult_Message_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	actual := args.Map{"result": rr.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov_TypedResponseResult_Message(t *testing.T) {
	attr := &coreapi.ResponseAttribute{Message: "ok"}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	actual := args.Map{"result": rr.Message() != "ok"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return message", actual)
}

func Test_Cov_TypedResponseResult_ClonePtr_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	actual := args.Map{"result": rr.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_TypedResponseResult_ClonePtr(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	c := rr.ClonePtr()
	actual := args.Map{"result": c == nil || c.Response != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_Cov_TypedResponseResult_ToTypedResponse_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	actual := args.Map{"result": rr.ToTypedResponse() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Cov_TypedResponseResult_ToTypedResponse(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	resp := rr.ToTypedResponse()
	actual := args.Map{"result": resp == nil || resp.Response != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_Cov_InvalidTypedResponseResult_NilAttribute(t *testing.T) {
	rr := coreapi.InvalidTypedResponseResult[string](nil)
	actual := args.Map{"result": rr.Attribute == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedRequest — uncovered branches
// =============================================================================

func Test_Cov_TypedRequest_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedRequest[string]
	actual := args.Map{"result": req.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_TypedRequest_ToTypedSimpleGenericRequest_Nil(t *testing.T) {
	var req *coreapi.TypedRequest[string]
	actual := args.Map{"result": req.ToTypedSimpleGenericRequest(true, "") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Cov_TypedRequest_ToTypedSimpleGenericRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequest(attr, "payload")
	sgr := req.ToTypedSimpleGenericRequest(true, "")
	actual := args.Map{"result": sgr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_Cov_InvalidTypedRequest_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedRequest[string](nil)
	actual := args.Map{"result": req.Attribute == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedRequestIn — uncovered branches
// =============================================================================

func Test_Cov_TypedRequestIn_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedRequestIn[string]
	actual := args.Map{"result": req.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_TypedRequestIn_TypedSimpleGenericRequest_Nil(t *testing.T) {
	var req *coreapi.TypedRequestIn[string]
	actual := args.Map{"result": req.TypedSimpleGenericRequest(true, "") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Cov_TypedRequestIn_TypedSimpleGenericRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequestIn(attr, "payload")
	sgr := req.TypedSimpleGenericRequest(true, "")
	actual := args.Map{"result": sgr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}

func Test_Cov_InvalidTypedRequestIn_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedRequestIn[string](nil)
	actual := args.Map{"result": req.Attribute == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

// =============================================================================
// TypedSimpleGenericRequest — uncovered branches
// =============================================================================

func Test_Cov_TypedSimpleGenericRequest_IsValid_NilReceiver(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	actual := args.Map{"result": req.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Cov_TypedSimpleGenericRequest_IsValid_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{
		Attribute: &coreapi.RequestAttribute{IsValid: true},
	}
	actual := args.Map{"result": req.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should be invalid", actual)
}

func Test_Cov_TypedSimpleGenericRequest_IsValid_True(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	actual := args.Map{"result": req.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_Cov_TypedSimpleGenericRequest_IsInvalid(t *testing.T) {
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
	actual := args.Map{"result": req.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Cov_TypedSimpleGenericRequest_Message_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	actual := args.Map{"result": req.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Cov_TypedSimpleGenericRequest_Message_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{}
	actual := args.Map{"result": req.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should return empty", actual)
}

func Test_Cov_TypedSimpleGenericRequest_InvalidError_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	actual := args.Map{"result": req.InvalidError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Cov_TypedSimpleGenericRequest_InvalidError_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{}
	actual := args.Map{"result": req.InvalidError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil request should return nil", actual)
}

func Test_Cov_TypedSimpleGenericRequest_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	actual := args.Map{"result": req.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Cov_TypedSimpleGenericRequest_Clone(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	c := req.Clone()
	actual := args.Map{"result": c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clone", actual)
}

func Test_Cov_InvalidTypedSimpleGenericRequest_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
	actual := args.Map{"result": req.Attribute == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have default invalid attribute", actual)
}

func Test_Cov_TypedSimpleGenericRequest_Data(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("hello", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	actual := args.Map{"result": req.Data() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return data", actual)
}
