package coreinstructiontests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/reqtype"
)

func regexpCompile(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

// ── BaseLineIdentifier ──

func Test_BaseLineIdentifier_New(t *testing.T) {
	// Arrange
	bli := coreinstruction.NewBaseLineIdentifier(5, reqtype.Create)

	// Act
	actual := args.Map{
		"line": bli.LineNumber,
		"isCreate": bli.IsNewLineRequest(),
	}

	// Assert
	expected := args.Map{
		"line": 5,
		"isCreate": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLineIdentifier_New returns correct value -- with args", actual)
}

func Test_BaseLineIdentifier_ToNewLineIdentifier(t *testing.T) {
	// Arrange
	bli := coreinstruction.NewBaseLineIdentifier(3, reqtype.Update)
	li := bli.ToNewLineIdentifier()

	// Act
	actual := args.Map{
		"line": li.LineNumber,
		"isModify": li.IsModifyLineRequest(),
	}

	// Assert
	expected := args.Map{
		"line": 3,
		"isModify": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLineIdentifier_ToNew returns correct value -- with args", actual)
}

func Test_BaseLineIdentifier_ToNewLineIdentifier_Nil(t *testing.T) {
	// Arrange
	var bli *coreinstruction.BaseLineIdentifier

	// Act
	actual := args.Map{"isNil": bli.ToNewLineIdentifier() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseLineIdentifier_ToNew_Nil returns nil -- with args", actual)
}

func Test_BaseLineIdentifier_Clone(t *testing.T) {
	// Arrange
	bli := coreinstruction.NewBaseLineIdentifier(7, reqtype.Delete)
	cloned := bli.Clone()

	// Act
	actual := args.Map{
		"line": cloned.LineNumber,
		"isDel": cloned.IsDeleteLineRequest(),
	}

	// Assert
	expected := args.Map{
		"line": 7,
		"isDel": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseLineIdentifier_Clone returns correct value -- with args", actual)
}

func Test_BaseLineIdentifier_Clone_Nil(t *testing.T) {
	// Arrange
	var bli *coreinstruction.BaseLineIdentifier

	// Act
	actual := args.Map{"isNil": bli.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseLineIdentifier_Clone_Nil returns nil -- with args", actual)
}

// ── LineIdentifier non-nil paths ──

func Test_LineIdentifier_NonNil(t *testing.T) {
	// Arrange
	create := &coreinstruction.LineIdentifier{LineNumber: 1, LineModifyAs: reqtype.Create}
	update := &coreinstruction.LineIdentifier{LineNumber: 2, LineModifyAs: reqtype.Update}
	del := &coreinstruction.LineIdentifier{LineNumber: 3, LineModifyAs: reqtype.Delete}

	// Act
	actual := args.Map{
		"createIsNew":   create.IsNewLineRequest(),
		"createAddMod":  create.IsAddNewOrModifyLineRequest(),
		"updateModify":  update.IsModifyLineRequest(),
		"updateAddMod":  update.IsAddNewOrModifyLineRequest(),
		"delDelete":     del.IsDeleteLineRequest(),
		"delHasLine":    del.HasLineNumber(),
		"delInvalid":    del.IsInvalidLineNumber(),
		"toBase":        create.ToBaseLineIdentifier() != nil,
		"clone":         create.Clone() != nil,
		"invalidUsing":  create.IsInvalidLineNumberUsingLastLineNumber(0),
		"validUsing":    create.IsInvalidLineNumberUsingLastLineNumber(10),
	}

	// Assert
	expected := args.Map{
		"createIsNew":   true,
		"createAddMod":  true,
		"updateModify":  true,
		"updateAddMod":  true,
		"delDelete":     true,
		"delHasLine":    true,
		"delInvalid":    false,
		"toBase":        true,
		"clone":         true,
		"invalidUsing":  true,
		"validUsing":    false,
	}
	expected.ShouldBeEqual(t, 0, "LineIdentifier_NonNil returns nil -- with args", actual)
}

// ── BaseModifyAs ──

func Test_BaseModifyAs(t *testing.T) {
	// Arrange
	bm := coreinstruction.NewModifyAs(reqtype.Create)

	// Act
	actual := args.Map{"modifyAs": string(bm.ModifyAs)}

	// Assert
	expected := args.Map{"modifyAs": string(reqtype.Create)}
	expected.ShouldBeEqual(t, 0, "BaseModifyAs_New returns correct value -- with args", actual)

	bm.SetModifyAs(reqtype.Update)
	actual2 := args.Map{"modifyAs": string(bm.ModifyAs)}
	expected2 := args.Map{"modifyAs": string(reqtype.Update)}
	expected2.ShouldBeEqual(t, 0, "BaseModifyAs_Set returns correct value -- with args", actual2)
}

// ── BaseSpecification ──

func Test_BaseSpecification(t *testing.T) {
	// Arrange
	bs := coreinstruction.NewBaseSpecification("id1", "disp", "tp", []string{"t1"}, true)

	// Act
	actual := args.Map{
		"id":      bs.Identifier().Id,
		"display": bs.Display().Display,
		"type":    bs.Type().Type,
		"hasSpec": bs.HasSpec(),
		"empty":   bs.IsEmptySpec(),
	}

	// Assert
	expected := args.Map{
		"id":      "id1",
		"display": "disp",
		"type":    "tp",
		"hasSpec": true,
		"empty":   false,
	}
	expected.ShouldBeEqual(t, 0, "BaseSpecification returns correct value -- with args", actual)
}

func Test_BaseSpecification_Clone(t *testing.T) {
	// Arrange
	bs := coreinstruction.NewBaseSpecification("id1", "disp", "tp", nil, false)
	cloned := bs.Clone()

	// Act
	actual := args.Map{"id": cloned.Specification.Id}

	// Assert
	expected := args.Map{"id": "id1"}
	expected.ShouldBeEqual(t, 0, "BaseSpecification_Clone returns correct value -- with args", actual)
}

func Test_BaseSpecification_Clone_Nil(t *testing.T) {
	// Arrange
	var bs *coreinstruction.BaseSpecification

	// Act
	actual := args.Map{"isNil": bs.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseSpecification_Clone_Nil returns nil -- with args", actual)
}

// ── BaseSpecPlusRequestIds ──

func Test_BaseSpecPlusRequestIds_SpecOnly(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id1", "d", "t")
	bspr := coreinstruction.NewBaseSpecPlusRequestIdsUsingSpecOnly(spec)

	// Act
	actual := args.Map{
		"specId": bspr.Specification.Id,
		"reqLen": len(bspr.RequestIds),
	}

	// Assert
	expected := args.Map{
		"specId": "id1",
		"reqLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "BaseSpecPlusReqIds_SpecOnly returns correct value -- with args", actual)
}

func Test_BaseSpecPlusRequestIds_Full(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id1", "d", "t")
	reqIds := coreinstruction.NewRequestIds(true, "r1", "r2")
	bspr := coreinstruction.NewBaseSpecPlusRequestIds(spec, reqIds)

	// Act
	actual := args.Map{"reqLen": len(bspr.RequestIds)}

	// Assert
	expected := args.Map{"reqLen": 2}
	expected.ShouldBeEqual(t, 0, "BaseSpecPlusReqIds_Full returns correct value -- with args", actual)
}

func Test_BaseSpecPlusRequestIds_Clone(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id1", "d", "t")
	bspr := coreinstruction.NewBaseSpecPlusRequestIdsUsingSpecOnly(spec)
	cloned := bspr.Clone()

	// Act
	actual := args.Map{"specId": cloned.Specification.Id}

	// Assert
	expected := args.Map{"specId": "id1"}
	expected.ShouldBeEqual(t, 0, "BaseSpecPlusReqIds_Clone returns correct value -- with args", actual)
}

func Test_BaseSpecPlusRequestIds_Clone_Nil(t *testing.T) {
	// Arrange
	var bspr *coreinstruction.BaseSpecPlusRequestIds

	// Act
	actual := args.Map{"isNil": bspr.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseSpecPlusReqIds_Clone_Nil returns nil -- with args", actual)
}

// ── BaseRequestIds ──

func Test_BaseRequestIds(t *testing.T) {
	// Arrange
	bri := coreinstruction.NewBaseRequestIds(true, "a", "b")

	// Act
	actual := args.Map{
		"len":    bri.RequestIdsLength(),
		"hasIds": bri.HasRequestIds(),
		"empty":  bri.IsEmptyRequestIds(),
	}

	// Assert
	expected := args.Map{
		"len":    2,
		"hasIds": true,
		"empty":  false,
	}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds returns correct value -- with args", actual)
}

func Test_BaseRequestIds_AddReqId(t *testing.T) {
	// Arrange
	bri := coreinstruction.NewBaseRequestIds(false, "a")
	rid := coreinstruction.IdentifierWithIsGlobal{
		BaseIdentifier: coreinstruction.BaseIdentifier{Id: "b"},
		IsGlobal:       true,
	}
	bri.AddReqId(rid)

	// Act
	actual := args.Map{"len": bri.RequestIdsLength()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_AddReqId returns correct value -- with args", actual)
}

func Test_BaseRequestIds_AddIds(t *testing.T) {
	// Arrange
	bri := coreinstruction.NewBaseRequestIds(false, "a")
	bri.AddIds(true, "b", "c")

	// Act
	actual := args.Map{"len": bri.RequestIdsLength()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_AddIds returns correct value -- with args", actual)
}

func Test_BaseRequestIds_AddIds_Empty(t *testing.T) {
	// Arrange
	bri := coreinstruction.NewBaseRequestIds(false, "a")
	bri.AddIds(true)

	// Act
	actual := args.Map{"len": bri.RequestIdsLength()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_AddIds_Empty returns empty -- with args", actual)
}

func Test_BaseRequestIds_Clone(t *testing.T) {
	// Arrange
	bri := coreinstruction.NewBaseRequestIds(true, "a", "b")
	cloned := bri.Clone()

	// Act
	actual := args.Map{"len": cloned.RequestIdsLength()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_Clone returns correct value -- with args", actual)
}

func Test_BaseRequestIds_Clone_Nil(t *testing.T) {
	// Arrange
	var bri *coreinstruction.BaseRequestIds

	// Act
	actual := args.Map{"isNil": bri.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_Clone_Nil returns nil -- with args", actual)
}

func Test_BaseRequestIds_NilLen(t *testing.T) {
	// Arrange
	var bri *coreinstruction.BaseRequestIds

	// Act
	actual := args.Map{
		"len": bri.RequestIdsLength(),
		"empty": bri.IsEmptyRequestIds(),
		"has": bri.HasRequestIds(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseRequestIds_Nil returns nil -- with args", actual)
}

// ── NewRequestIds / NewRequestId ──

func Test_NewRequestIds_Empty(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewRequestIds(true)

	// Act
	actual := args.Map{"len": len(ids)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewRequestIds_Empty returns empty -- with args", actual)
}

func Test_NewRequestId(t *testing.T) {
	// Arrange
	rid := coreinstruction.NewRequestId(true, "x")

	// Act
	actual := args.Map{
		"id": rid.Id,
		"global": rid.IsGlobal,
	}

	// Assert
	expected := args.Map{
		"id": "x",
		"global": true,
	}
	expected.ShouldBeEqual(t, 0, "NewRequestId returns correct value -- with args", actual)
}

// ── BaseSourceDestination ──

func Test_BaseSourceDestination(t *testing.T) {
	// Arrange
	bsd := coreinstruction.NewBaseSourceDestination("src", "dst")

	// Act
	actual := args.Map{
		"src": bsd.Source,
		"dst": bsd.Destination,
	}

	// Assert
	expected := args.Map{
		"src": "src",
		"dst": "dst",
	}
	expected.ShouldBeEqual(t, 0, "BaseSourceDestination returns correct value -- with args", actual)
}

// ── BaseIsRename ──

func Test_BaseIsRename(t *testing.T) {
	// Arrange
	r := coreinstruction.NewRename(true)

	// Act
	actual := args.Map{"isRename": r.IsRename}

	// Assert
	expected := args.Map{"isRename": true}
	expected.ShouldBeEqual(t, 0, "BaseIsRename returns correct value -- with args", actual)
}

// ── BaseUsername ──

func Test_BaseUsername(t *testing.T) {
	// Arrange
	u := coreinstruction.NewUsername("admin")

	// Act
	actual := args.Map{
		"str":              u.UsernameString(),
		"isEmpty":          u.IsUsernameEmpty(),
		"isWhitespace":     u.IsUsernameWhitespace(),
		"isAdmin":          u.IsUsername("admin"),
		"isNotAdmin":       u.IsUsername("other"),
		"caseInsensitive":  u.IsUsernameCaseInsensitive("ADMIN"),
		"contains":         u.IsUsernameContains("dmi"),
	}

	// Assert
	expected := args.Map{
		"str":              "admin",
		"isEmpty":          false,
		"isWhitespace":     false,
		"isAdmin":          true,
		"isNotAdmin":       false,
		"caseInsensitive":  true,
		"contains":         true,
	}
	expected.ShouldBeEqual(t, 0, "BaseUsername returns correct value -- with args", actual)
}

func Test_BaseUsername_Nil(t *testing.T) {
	// Arrange
	var u *coreinstruction.BaseUsername

	// Act
	actual := args.Map{
		"isEmpty": u.IsUsernameEmpty(),
		"isWs": u.IsUsernameWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"isWs": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseUsername_Nil returns nil -- with args", actual)
}

func Test_BaseUsername_IsEqual(t *testing.T) {
	// Arrange
	u1 := coreinstruction.NewUsername("a")
	u2 := coreinstruction.NewUsername("a")
	u3 := coreinstruction.NewUsername("b")
	var nilU *coreinstruction.BaseUsername

	// Act
	actual := args.Map{
		"equal":       u1.IsEqual(u2),
		"notEqual":    u1.IsNotEqual(u3),
		"nilBoth":     nilU.IsEqual(nil),
		"nilLeft":     nilU.IsEqual(u1),
		"nilRight":    u1.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"equal":       true,
		"notEqual":    true,
		"nilBoth":     true,
		"nilLeft":     false,
		"nilRight":    false,
	}
	expected.ShouldBeEqual(t, 0, "BaseUsername_IsEqual returns correct value -- with args", actual)
}

func Test_BaseUsername_Clone(t *testing.T) {
	// Arrange
	u := coreinstruction.NewUsername("test")
	cloned := u.ClonePtr()
	val := u.Clone()

	// Act
	actual := args.Map{
		"ptrName": cloned.Username,
		"valName": val.Username,
	}

	// Assert
	expected := args.Map{
		"ptrName": "test",
		"valName": "test",
	}
	expected.ShouldBeEqual(t, 0, "BaseUsername_Clone returns correct value -- with args", actual)
}

func Test_BaseUsername_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var u *coreinstruction.BaseUsername

	// Act
	actual := args.Map{"isNil": u.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseUsername_ClonePtr_Nil returns nil -- with args", actual)
}

func Test_BaseUsername_Regex(t *testing.T) {
	// Arrange
	u := coreinstruction.NewUsername("user123")
	re := regexpCompile(`\d+`)

	// Act
	actual := args.Map{"match": u.IsUsernameRegexMatches(re)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "BaseUsername_Regex returns correct value -- with args", actual)
}

// ── IdentifierWithIsGlobal ──

func Test_IdentifierWithIsGlobal_Clone(t *testing.T) {
	// Arrange
	iwg := coreinstruction.NewIdentifierWithIsGlobal("x", true)
	cloned := iwg.Clone()

	// Act
	actual := args.Map{
		"id": cloned.Id,
		"global": cloned.IsGlobal,
	}

	// Assert
	expected := args.Map{
		"id": "x",
		"global": true,
	}
	expected.ShouldBeEqual(t, 0, "IdentifierWithIsGlobal_Clone returns non-empty -- with args", actual)
}

func Test_IdentifierWithIsGlobal_Clone_Nil(t *testing.T) {
	// Arrange
	var iwg *coreinstruction.IdentifierWithIsGlobal

	// Act
	actual := args.Map{"isNil": iwg.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "IdentifierWithIsGlobal_Clone_Nil returns nil -- with args", actual)
}

// ── IdentifiersWithGlobals additional ──

func Test_IdentifiersWithGlobals_Full(t *testing.T) {
	// Arrange
	iwgs := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b")
	iwgs.Add(false, "c")
	iwgs.Adds(true, "d", "e")

	// Act
	actual := args.Map{
		"len":     iwgs.Length(),
		"hasAny":  iwgs.HasAnyItem(),
		"indexOf": iwgs.IndexOf("c"),
		"getById": iwgs.GetById("a") != nil,
		"getNil":  iwgs.GetById("z") == nil,
	}

	// Assert
	expected := args.Map{
		"len":     5,
		"hasAny":  true,
		"indexOf": 2,
		"getById": true,
		"getNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "IdentifiersWithGlobals_Full returns non-empty -- with args", actual)
}

func Test_IdentifiersWithGlobals_Empty(t *testing.T) {
	// Arrange
	iwgs := coreinstruction.EmptyIdentifiersWithGlobals()
	iwgs.Add(false, "") // skip empty
	iwgs.Adds(true)     // skip empty args

	// Act
	actual := args.Map{
		"isEmpty": iwgs.IsEmpty(),
		"indexOf": iwgs.IndexOf(""),
		"getById": iwgs.GetById("") == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"indexOf": -1,
		"getById": true,
	}
	expected.ShouldBeEqual(t, 0, "IdentifiersWithGlobals_Empty returns empty -- with args", actual)
}

func Test_IdentifiersWithGlobals_Clone(t *testing.T) {
	// Arrange
	iwgs := coreinstruction.NewIdentifiersWithGlobals(true, "a")
	cloned := iwgs.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "IdentifiersWithGlobals_Clone returns non-empty -- with args", actual)
}

func Test_IdentifiersWithGlobals_EmptyClone(t *testing.T) {
	// Arrange
	iwgs := coreinstruction.NewIdentifiersWithGlobals(true)
	cloned := iwgs.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IdentifiersWithGlobals_EmptyClone returns empty -- with args", actual)
}

// ── Identifiers additional ──

func Test_Identifiers_NewEmpty(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiers()

	// Act
	actual := args.Map{
		"len": ids.Length(),
		"empty": ids.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers_NewEmpty returns empty -- with args", actual)
}

func Test_Identifiers_EmptyClone(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiers()
	cloned := ids.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Identifiers_EmptyClone returns empty -- with args", actual)
}

func Test_Identifiers_AddsEmpty(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiers()
	ids.Adds()

	// Act
	actual := args.Map{"len": ids.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Identifiers_AddsEmpty returns empty -- with args", actual)
}

// ── StringSearch non-nil ──

func Test_StringSearch_NonNil(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "hello",
	}

	// Act
	actual := args.Map{
		"isEmpty":    ss.IsEmpty(),
		"isExist":    ss.IsExist(),
		"has":        ss.Has(),
		"match":      ss.IsMatch("say hello world"),
		"matchFail":  ss.IsMatchFailed("goodbye"),
		"allMatch":   ss.IsAllMatch("hello there", "say hello"),
		"anyFail":    ss.IsAnyMatchFailed("goodbye", "hello"),
		"verifyErr":  ss.VerifyError("hello world") == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty":    false,
		"isExist":    true,
		"has":        true,
		"match":      true,
		"matchFail":  true,
		"allMatch":   true,
		"anyFail":    true,
		"verifyErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_NonNil returns nil -- with args", actual)
}

func Test_StringSearch_Regex(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Regex,
		Search:        `\d+`,
	}

	// Act
	actual := args.Map{
		"match":     ss.IsMatch("abc123"),
		"noMatch":   ss.IsMatch("abcdef"),
		"verifyErr": ss.VerifyError("abc123") == nil,
	}

	// Assert
	expected := args.Map{
		"match":     true,
		"noMatch":   false,
		"verifyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_Regex returns correct value -- with args", actual)
}

// ── StringCompare non-nil paths ──

func Test_StringCompare_MatchFailed(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEqual("abc", "xyz")

	// Act
	actual := args.Map{
		"isDefined":  sc.IsDefined(),
		"isInvalid":  sc.IsInvalid(),
		"matchFail":  sc.IsMatchFailed(),
		"isMatch":    sc.IsMatch(),
	}

	// Assert
	expected := args.Map{
		"isDefined":  true,
		"isInvalid":  false,
		"matchFail":  true,
		"isMatch":    false,
	}
	expected.ShouldBeEqual(t, 0, "StringCompare_MatchFailed returns correct value -- with args", actual)
}

func Test_StringCompare_VerifyError_Fail(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEqual("abc", "xyz")

	// Act
	actual := args.Map{"hasErr": sc.VerifyError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringCompare_VerifyError_Fail returns error -- with args", actual)
}

func Test_StringCompare_Regex_Fail(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareRegex(`\d+`, "nodigits")

	// Act
	actual := args.Map{"hasErr": sc.VerifyError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringCompare_Regex_Fail returns correct value -- with args", actual)
}

// ── NameList clone with list ──

func Test_NameList_CloneWithList(t *testing.T) {
	// Arrange
	nl := &coreinstruction.NameList{Name: "n"}
	cloned := nl.Clone(false)

	// Act
	actual := args.Map{"name": cloned.Name}

	// Assert
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NameList_Clone returns correct value -- with args", actual)
}

// ── NameRequests / NameRequestsCollection ──

func Test_NameRequests(t *testing.T) {
	// Arrange
	nr := coreinstruction.NameRequests{Name: "test"}

	// Act
	actual := args.Map{"name": nr.Name}

	// Assert
	expected := args.Map{"name": "test"}
	expected.ShouldBeEqual(t, 0, "NameRequests returns correct value -- with args", actual)
}

func Test_NameRequestsCollection(t *testing.T) {
	// Arrange
	nrc := coreinstruction.NameRequestsCollection{
		NameRequestsList: []coreinstruction.NameRequests{{Name: "a"}},
	}

	// Act
	actual := args.Map{"len": len(nrc.NameRequestsList)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NameRequestsCollection returns correct value -- with args", actual)
}

// ── ParentIdentifier / ById / DependsOn / DependencyName / SpecificVersion ──

func Test_ParentIdentifier(t *testing.T) {
	// Arrange
	pi := coreinstruction.ParentIdentifier{ParentId: "p1", ParentName: "pn", ParentVersion: "v1"}

	// Act
	actual := args.Map{
		"id": pi.ParentId,
		"name": pi.ParentName,
		"ver": pi.ParentVersion,
	}

	// Assert
	expected := args.Map{
		"id": "p1",
		"name": "pn",
		"ver": "v1",
	}
	expected.ShouldBeEqual(t, 0, "ParentIdentifier returns correct value -- with args", actual)
}

func Test_ById(t *testing.T) {
	// Arrange
	b := coreinstruction.ById{Id: "x"}

	// Act
	actual := args.Map{"id": b.Id}

	// Assert
	expected := args.Map{"id": "x"}
	expected.ShouldBeEqual(t, 0, "ById returns correct value -- with args", actual)
}

func Test_DependsOn(t *testing.T) {
	// Arrange
	d := coreinstruction.DependsOn{
		SpecificVersion: coreinstruction.SpecificVersion{Version: "1.0", IsSpecific: true},
		DependencyName:  coreinstruction.DependencyName{Name: "dep"},
	}

	// Act
	actual := args.Map{
		"name": d.Name,
		"ver": d.Version,
		"isSpec": d.IsSpecific,
	}

	// Assert
	expected := args.Map{
		"name": "dep",
		"ver": "1.0",
		"isSpec": true,
	}
	expected.ShouldBeEqual(t, 0, "DependsOn returns correct value -- with args", actual)
}

// ── BaseByIds ──

func Test_BaseByIds(t *testing.T) {
	// Arrange
	bb := coreinstruction.BaseByIds{ByIds: []coreinstruction.ById{{Id: "a"}, {Id: "b"}}}

	// Act
	actual := args.Map{"len": len(bb.ByIds)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BaseByIds returns correct value -- with args", actual)
}

// ── SourceDestination nil paths for IsSourceEmpty/IsDestinationEmpty ──

func Test_SourceDestination_NilEmpty(t *testing.T) {
	// Arrange
	var sd *coreinstruction.SourceDestination

	// Act
	actual := args.Map{
		"srcEmpty": sd.IsSourceEmpty(),
		"dstEmpty": sd.IsDestinationEmpty(),
	}

	// Assert
	expected := args.Map{
		"srcEmpty": true,
		"dstEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SourceDestination_NilEmpty returns nil -- with args", actual)
}

// ── Rename nil paths for IsExistingEmpty/IsNewEmpty ──

func Test_Rename_NilEmpty(t *testing.T) {
	// Arrange
	var r *coreinstruction.Rename

	// Act
	actual := args.Map{
		"existEmpty": r.IsExistingEmpty(),
		"newEmpty":   r.IsNewEmpty(),
	}

	// Assert
	expected := args.Map{
		"existEmpty": true,
		"newEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "Rename_NilEmpty returns nil -- with args", actual)
}

// ── NameList nil paths ──

func Test_NameList_NilPaths(t *testing.T) {
	// Arrange
	var nl *coreinstruction.NameList

	// Act
	actual := args.Map{
		"isNameEmpty": nl.IsNameEmpty(),
		"hasName":     nl.HasName(),
	}

	// Assert
	expected := args.Map{
		"isNameEmpty": true,
		"hasName":     false,
	}
	expected.ShouldBeEqual(t, 0, "NameList_NilPaths returns nil -- with args", actual)
}

// ── BaseSpecification HasSpec nil ──

func Test_BaseSpec_HasSpec_Nil(t *testing.T) {
	// Arrange
	var bs *coreinstruction.BaseSpecification

	// Act
	actual := args.Map{"has": bs.HasSpec()}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "BaseSpec_HasSpec_Nil returns nil -- with args", actual)
}

// ── IdentifiersWithGlobals nil ──

func Test_IdentifiersWithGlobals_Nil(t *testing.T) {
	// Arrange
	var iwgs *coreinstruction.IdentifiersWithGlobals

	// Act
	actual := args.Map{
		"len":   iwgs.Length(),
		"empty": iwgs.IsEmpty(),
		"has":   iwgs.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len":   0,
		"empty": true,
		"has":   false,
	}
	expected.ShouldBeEqual(t, 0, "IdentifiersWithGlobals_Nil returns nil -- with args", actual)
}

// ── Identifiers nil ──

func Test_Identifiers_Nil(t *testing.T) {
	// Arrange
	var ids *coreinstruction.Identifiers

	// Act
	actual := args.Map{
		"len": ids.Length(),
		"empty": ids.IsEmpty(),
		"has": ids.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"empty": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers_Nil returns nil -- with args", actual)
}
