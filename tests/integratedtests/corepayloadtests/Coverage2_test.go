package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── PagingInfo ──

func Test_Cov2_PagingInfo(t *testing.T) {
	p := corepayload.PagingInfo{
		CurrentPageIndex: 1,
		PerPageItems:     10,
		TotalItems:       25,
		TotalPages:       3,
	}

	actual := args.Map{
		"totalPages": p.TotalPages,
		"hasTotalPages": p.HasTotalPages(),
		"hasCurrentPage": p.HasCurrentPageIndex(),
	}
	expected := args.Map{
		"totalPages": 3,
		"hasTotalPages": true,
		"hasCurrentPage": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- with args", actual)
}

func Test_Cov2_PagingInfo_Empty(t *testing.T) {
	p := corepayload.PagingInfo{}

	actual := args.Map{
		"isEmpty":    p.IsEmpty(),
		"hasTotalPages": p.HasTotalPages(),
	}
	expected := args.Map{
		"isEmpty":    true,
		"hasTotalPages": false,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns empty -- empty", actual)
}

func Test_Cov2_PagingInfo_Clone(t *testing.T) {
	p := corepayload.PagingInfo{
		CurrentPageIndex: 2,
		PerPageItems:     10,
		TotalItems:       50,
		TotalPages:       5,
	}
	cloned := p.Clone()

	actual := args.Map{
		"isEqual": p.IsEqual(&cloned),
		"totalPages": cloned.TotalPages,
	}
	expected := args.Map{
		"isEqual": true,
		"totalPages": 5,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- clone", actual)
}

func Test_Cov2_PagingInfo_ClonePtr(t *testing.T) {
	p := &corepayload.PagingInfo{TotalPages: 3, TotalItems: 30}
	cloned := p.ClonePtr()

	actual := args.Map{
		"notNil":  cloned != nil,
		"isEqual": p.IsEqual(cloned),
	}
	expected := args.Map{
		"notNil":  true,
		"isEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- clonePtr", actual)
}

func Test_Cov2_PagingInfo_ClonePtr_Nil(t *testing.T) {
	var p *corepayload.PagingInfo
	cloned := p.ClonePtr()

	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- clonePtr nil", actual)
}

func Test_Cov2_PagingInfo_InvalidChecks(t *testing.T) {
	p := corepayload.PagingInfo{}

	actual := args.Map{
		"invalidTotalPages":       p.IsInvalidTotalPages(),
		"invalidCurrentPageIndex": p.IsInvalidCurrentPageIndex(),
		"invalidPerPageItems":     p.IsInvalidPerPageItems(),
		"invalidTotalItems":       p.IsInvalidTotalItems(),
	}
	expected := args.Map{
		"invalidTotalPages":       true,
		"invalidCurrentPageIndex": true,
		"invalidPerPageItems":     true,
		"invalidTotalItems":       true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns error -- invalid checks", actual)
}

// ── SessionInfo ──

func Test_Cov2_SessionInfo(t *testing.T) {
	s := corepayload.SessionInfo{Id: "abc123"}

	actual := args.Map{
		"isValid": s.IsValid(),
		"id":      s.Id,
	}
	expected := args.Map{
		"isValid": true,
		"id":      "abc123",
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- with args", actual)
}

func Test_Cov2_SessionInfo_Empty(t *testing.T) {
	s := corepayload.SessionInfo{}

	actual := args.Map{"isEmpty": s.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty -- all zero fields", actual)
}

// ── AuthInfo ──

func Test_Cov2_AuthInfo(t *testing.T) {
	a := corepayload.AuthInfo{Identifier: "id1", ActionType: "login", ResourceName: "/api"}

	actual := args.Map{
		"hasAction":   a.HasActionType(),
		"hasResource": a.HasResourceName(),
		"identifier":  a.Identifier,
	}
	expected := args.Map{
		"hasAction":   true,
		"hasResource": true,
		"identifier":  "id1",
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- with args", actual)
}

// ── PayloadWrapper ──

func Test_Cov2_PayloadWrapper_Basic(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"test", "id1", "task", "cat",
		map[string]string{"k": "v"},
	)

	actual := args.Map{
		"notNil":  pw != nil,
		"hasAny":  pw.HasAnyItem(),
	}
	expected := args.Map{
		"notNil":  true,
		"hasAny":  true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_Cov2_PayloadWrapper_Empty(t *testing.T) {
	pw := corepayload.New.PayloadWrapper.Empty()

	actual := args.Map{
		"notNil":  pw != nil,
		"hasAny":  pw.HasAnyItem(),
	}
	expected := args.Map{
		"notNil":  true,
		"hasAny":  false,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns empty -- empty", actual)
}
