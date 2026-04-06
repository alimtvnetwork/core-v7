package coreapitests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

// =============================================================================
// RequestAttribute — uncovered branches
// =============================================================================

func Test_Cov_RequestAttribute_HasSearchRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{SearchRequest: &coreapi.SearchRequest{}}
	if !attr.HasSearchRequest() {
		t.Error("should have search request")
	}
}

func Test_Cov_RequestAttribute_HasSearchRequest_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if attr.HasSearchRequest() {
		t.Error("nil should not have search request")
	}
}

func Test_Cov_RequestAttribute_HasPageRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{PageRequest: &coreapi.PageRequest{PageSize: 10}}
	if !attr.HasPageRequest() {
		t.Error("should have page request")
	}
}

func Test_Cov_RequestAttribute_HasPageRequest_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if attr.HasPageRequest() {
		t.Error("nil should not have page request")
	}
}

func Test_Cov_RequestAttribute_IsEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if !attr.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Cov_RequestAttribute_IsAnyNull_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if !attr.IsAnyNull() {
		t.Error("nil should be null")
	}
}

func Test_Cov_RequestAttribute_IsPageRequestEmpty(t *testing.T) {
	attr := &coreapi.RequestAttribute{}
	if !attr.IsPageRequestEmpty() {
		t.Error("should be empty without page request")
	}
}

func Test_Cov_RequestAttribute_IsPageRequestEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if !attr.IsPageRequestEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Cov_RequestAttribute_IsSearchRequestEmpty(t *testing.T) {
	attr := &coreapi.RequestAttribute{}
	if !attr.IsSearchRequestEmpty() {
		t.Error("should be empty without search request")
	}
}

func Test_Cov_RequestAttribute_IsSearchRequestEmpty_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if !attr.IsSearchRequestEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Cov_RequestAttribute_Clone(t *testing.T) {
	attr := &coreapi.RequestAttribute{
		Url:          "http://test",
		IsValid:      true,
		SearchRequest: &coreapi.SearchRequest{SearchTerm: "test"},
		PageRequest:   &coreapi.PageRequest{PageSize: 10},
	}
	c := attr.Clone()
	if c == nil || c.Url != "http://test" {
		t.Error("should clone")
	}
}

func Test_Cov_RequestAttribute_Clone_Nil(t *testing.T) {
	var attr *coreapi.RequestAttribute
	if attr.Clone() != nil {
		t.Error("nil clone should return nil")
	}
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
	if c == nil || len(c.StepsPerformed) != 1 || len(c.DebugInfos) != 1 {
		t.Error("should clone with slices")
	}
}

func Test_Cov_ResponseAttribute_Clone_Nil(t *testing.T) {
	var attr *coreapi.ResponseAttribute
	if attr.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_ResponseAttribute_Clone_EmptySlices(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	c := attr.Clone()
	if c == nil || c.StepsPerformed != nil {
		t.Error("should clone without slices")
	}
}

// =============================================================================
// InvalidRequestAttribute / InvalidResponseAttribute
// =============================================================================

func Test_Cov_InvalidRequestAttribute(t *testing.T) {
	attr := coreapi.InvalidRequestAttribute()
	if attr.IsValid {
		t.Error("should be invalid")
	}
}

func Test_Cov_InvalidResponseAttribute(t *testing.T) {
	attr := coreapi.InvalidResponseAttribute("test error")
	if attr.IsValid || attr.Message != "test error" {
		t.Error("should be invalid with message")
	}
}

// =============================================================================
// SearchRequest — uncovered branches
// =============================================================================

func Test_Cov_SearchRequest_Clone(t *testing.T) {
	sr := &coreapi.SearchRequest{SearchTerm: "test", IsContains: true}
	c := sr.Clone()
	if c == nil || c.SearchTerm != "test" || !c.IsContains {
		t.Error("should clone")
	}
}

func Test_Cov_SearchRequest_Clone_Nil(t *testing.T) {
	var sr *coreapi.SearchRequest
	if sr.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

// =============================================================================
// PageRequest — uncovered branches
// =============================================================================

func Test_Cov_PageRequest_IsPageSizeEmpty_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	if !pr.IsPageSizeEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Cov_PageRequest_IsPageIndexEmpty_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	if !pr.IsPageIndexEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_Cov_PageRequest_HasPageSize(t *testing.T) {
	pr := &coreapi.PageRequest{PageSize: 10}
	if !pr.HasPageSize() {
		t.Error("should have page size")
	}
}

func Test_Cov_PageRequest_HasPageSize_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	if pr.HasPageSize() {
		t.Error("nil should not have page size")
	}
}

func Test_Cov_PageRequest_HasPageIndex(t *testing.T) {
	pr := &coreapi.PageRequest{PageIndex: 5}
	if !pr.HasPageIndex() {
		t.Error("should have page index")
	}
}

func Test_Cov_PageRequest_HasPageIndex_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	if pr.HasPageIndex() {
		t.Error("nil should not have page index")
	}
}

func Test_Cov_PageRequest_Clone_Nil(t *testing.T) {
	var pr *coreapi.PageRequest
	if pr.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_PageRequest_Clone(t *testing.T) {
	pr := &coreapi.PageRequest{PageSize: 10, PageIndex: 2}
	c := pr.Clone()
	if c.PageSize != 10 || c.PageIndex != 2 {
		t.Error("should clone")
	}
}

// =============================================================================
// TypedResponse — uncovered branches
// =============================================================================

func Test_Cov_TypedResponse_Clone_Nil(t *testing.T) {
	var resp *coreapi.TypedResponse[string]
	if resp.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_TypedResponse_TypedResponseResult_Nil(t *testing.T) {
	var resp *coreapi.TypedResponse[string]
	if resp.TypedResponseResult() != nil {
		t.Error("nil should return nil")
	}
}

func Test_Cov_TypedResponse_TypedResponseResult(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	resp := coreapi.NewTypedResponse(attr, "hello")
	result := resp.TypedResponseResult()
	if result == nil || result.Response != "hello" {
		t.Error("should convert")
	}
}

func Test_Cov_InvalidTypedResponse_NilAttribute(t *testing.T) {
	resp := coreapi.InvalidTypedResponse[string](nil)
	if resp.Attribute == nil {
		t.Error("should have default invalid attribute")
	}
}

// =============================================================================
// TypedResponseResult — uncovered branches
// =============================================================================

func Test_Cov_TypedResponseResult_IsValid(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	if !rr.IsValid() {
		t.Error("should be valid")
	}
}

func Test_Cov_TypedResponseResult_IsValid_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	if rr.IsValid() {
		t.Error("nil should be invalid")
	}
}

func Test_Cov_TypedResponseResult_IsInvalid(t *testing.T) {
	rr := coreapi.InvalidTypedResponseResult[string](nil)
	if !rr.IsInvalid() {
		t.Error("should be invalid")
	}
}

func Test_Cov_TypedResponseResult_Message_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	if rr.Message() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Cov_TypedResponseResult_Message(t *testing.T) {
	attr := &coreapi.ResponseAttribute{Message: "ok"}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	if rr.Message() != "ok" {
		t.Error("should return message")
	}
}

func Test_Cov_TypedResponseResult_ClonePtr_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	if rr.ClonePtr() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_TypedResponseResult_ClonePtr(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	c := rr.ClonePtr()
	if c == nil || c.Response != "data" {
		t.Error("should clone")
	}
}

func Test_Cov_TypedResponseResult_ToTypedResponse_Nil(t *testing.T) {
	var rr *coreapi.TypedResponseResult[string]
	if rr.ToTypedResponse() != nil {
		t.Error("nil should return nil")
	}
}

func Test_Cov_TypedResponseResult_ToTypedResponse(t *testing.T) {
	attr := &coreapi.ResponseAttribute{IsValid: true}
	rr := coreapi.NewTypedResponseResult(attr, "data")
	resp := rr.ToTypedResponse()
	if resp == nil || resp.Response != "data" {
		t.Error("should convert")
	}
}

func Test_Cov_InvalidTypedResponseResult_NilAttribute(t *testing.T) {
	rr := coreapi.InvalidTypedResponseResult[string](nil)
	if rr.Attribute == nil {
		t.Error("should have default invalid attribute")
	}
}

// =============================================================================
// TypedRequest — uncovered branches
// =============================================================================

func Test_Cov_TypedRequest_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedRequest[string]
	if req.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_TypedRequest_ToTypedSimpleGenericRequest_Nil(t *testing.T) {
	var req *coreapi.TypedRequest[string]
	if req.ToTypedSimpleGenericRequest(true, "") != nil {
		t.Error("nil should return nil")
	}
}

func Test_Cov_TypedRequest_ToTypedSimpleGenericRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequest(attr, "payload")
	sgr := req.ToTypedSimpleGenericRequest(true, "")
	if sgr == nil {
		t.Error("should convert")
	}
}

func Test_Cov_InvalidTypedRequest_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedRequest[string](nil)
	if req.Attribute == nil {
		t.Error("should have default invalid attribute")
	}
}

// =============================================================================
// TypedRequestIn — uncovered branches
// =============================================================================

func Test_Cov_TypedRequestIn_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedRequestIn[string]
	if req.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_TypedRequestIn_TypedSimpleGenericRequest_Nil(t *testing.T) {
	var req *coreapi.TypedRequestIn[string]
	if req.TypedSimpleGenericRequest(true, "") != nil {
		t.Error("nil should return nil")
	}
}

func Test_Cov_TypedRequestIn_TypedSimpleGenericRequest(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	req := coreapi.NewTypedRequestIn(attr, "payload")
	sgr := req.TypedSimpleGenericRequest(true, "")
	if sgr == nil {
		t.Error("should convert")
	}
}

func Test_Cov_InvalidTypedRequestIn_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedRequestIn[string](nil)
	if req.Attribute == nil {
		t.Error("should have default invalid attribute")
	}
}

// =============================================================================
// TypedSimpleGenericRequest — uncovered branches
// =============================================================================

func Test_Cov_TypedSimpleGenericRequest_IsValid_NilReceiver(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	if req.IsValid() {
		t.Error("nil should be invalid")
	}
}

func Test_Cov_TypedSimpleGenericRequest_IsValid_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{
		Attribute: &coreapi.RequestAttribute{IsValid: true},
	}
	if req.IsValid() {
		t.Error("nil request should be invalid")
	}
}

func Test_Cov_TypedSimpleGenericRequest_IsValid_True(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	if !req.IsValid() {
		t.Error("should be valid")
	}
}

func Test_Cov_TypedSimpleGenericRequest_IsInvalid(t *testing.T) {
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
	if !req.IsInvalid() {
		t.Error("should be invalid")
	}
}

func Test_Cov_TypedSimpleGenericRequest_Message_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	if req.Message() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Cov_TypedSimpleGenericRequest_Message_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{}
	if req.Message() != "" {
		t.Error("nil request should return empty")
	}
}

func Test_Cov_TypedSimpleGenericRequest_InvalidError_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	if req.InvalidError() != nil {
		t.Error("nil should return nil")
	}
}

func Test_Cov_TypedSimpleGenericRequest_InvalidError_NilRequest(t *testing.T) {
	req := &coreapi.TypedSimpleGenericRequest[string]{}
	if req.InvalidError() != nil {
		t.Error("nil request should return nil")
	}
}

func Test_Cov_TypedSimpleGenericRequest_Clone_Nil(t *testing.T) {
	var req *coreapi.TypedSimpleGenericRequest[string]
	if req.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Cov_TypedSimpleGenericRequest_Clone(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("data", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	c := req.Clone()
	if c == nil {
		t.Error("should clone")
	}
}

func Test_Cov_InvalidTypedSimpleGenericRequest_NilAttribute(t *testing.T) {
	req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
	if req.Attribute == nil {
		t.Error("should have default invalid attribute")
	}
}

func Test_Cov_TypedSimpleGenericRequest_Data(t *testing.T) {
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest("hello", true, "")
	req := coreapi.NewTypedSimpleGenericRequest(attr, simpleReq)
	if req.Data() != "hello" {
		t.Error("should return data")
	}
}
