package coreinstructiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── FlatSpecification ──

func Test_FlatSpec_InvalidFlatSpecification(t *testing.T) {
	// Arrange
	flat := coreinstruction.InvalidFlatSpecification()

	// Act
	actual := args.Map{
		"valid": flat.IsValid,
		"id": flat.Id,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"id": "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidFlatSpecification returns error -- with args", actual)
}

func Test_FlatSpec_BaseAccessors(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecification("id1", "disp", "tp", []string{"t1"}, true)
	flat := spec.FlatSpecification()

	// Act
	actual := args.Map{
		"baseId":       flat.BaseIdentifier().Id,
		"baseDisplay":  flat.BaseDisplay().Display,
		"baseType":     flat.BaseType().Type,
		"baseIsGlobal": flat.BaseIsGlobal().IsGlobal,
		"tagsLen":      len(flat.BaseTags().Tags),
	}

	// Assert
	expected := args.Map{
		"baseId":       "id1",
		"baseDisplay":  "disp",
		"baseType":     "tp",
		"baseIsGlobal": true,
		"tagsLen":      1,
	}
	expected.ShouldBeEqual(t, 0, "FlatSpec_BaseAccessors returns correct value -- with args", actual)
}

func Test_FlatSpec_SpecCaching(t *testing.T) {
	// Arrange
	flat := &coreinstruction.FlatSpecification{Id: "x", Display: "d", Type: "t"}
	s1 := flat.Spec()
	s2 := flat.Spec()

	// Act
	actual := args.Map{
		"same": s1 == s2,
		"id": s1.Id,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"id": "x",
	}
	expected.ShouldBeEqual(t, 0, "FlatSpec_SpecCaching returns correct value -- with args", actual)
}

func Test_FlatSpec_NilSpec(t *testing.T) {
	// Arrange
	var flat *coreinstruction.FlatSpecification

	// Act
	actual := args.Map{"isNil": flat.Spec() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FlatSpec_NilSpec returns nil -- with args", actual)
}

func Test_FlatSpec_Clone(t *testing.T) {
	// Arrange
	flat := &coreinstruction.FlatSpecification{Id: "x", Tags: []string{"a"}}
	cloned := flat.Clone()
	var nilFlat *coreinstruction.FlatSpecification

	// Act
	actual := args.Map{
		"id":       cloned.Id,
		"tagsLen":  len(cloned.Tags),
		"nilClone": nilFlat.Clone() == nil,
	}

	// Assert
	expected := args.Map{
		"id":       "x",
		"tagsLen":  1,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "FlatSpec_Clone returns correct value -- with args", actual)
}

// ── StringCompare ──

func Test_StringCompare_Constructors(t *testing.T) {
	// Arrange
	eq := coreinstruction.NewStringCompareEqual("abc", "abc")
	contains := coreinstruction.NewStringCompareContains(false, "bc", "abc")
	starts := coreinstruction.NewStringCompareStartsWith(false, "ab", "abc")
	ends := coreinstruction.NewStringCompareEndsWith(false, "bc", "abc")
	regex := coreinstruction.NewStringCompareRegex(`\w+`, "abc")

	// Act
	actual := args.Map{
		"eqMatch":       eq.IsMatch(),
		"containsMatch": contains.IsMatch(),
		"startsMatch":   starts.IsMatch(),
		"endsMatch":     ends.IsMatch(),
		"regexMatch":    regex.IsMatch(),
	}

	// Assert
	expected := args.Map{
		"eqMatch":       true,
		"containsMatch": true,
		"startsMatch":   true,
		"endsMatch":     true,
		"regexMatch":    true,
	}
	expected.ShouldBeEqual(t, 0, "StringCompare_Constructors returns correct value -- with args", actual)
}

func Test_StringCompare_NilMethods(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actual := args.Map{
		"isInvalid":  sc.IsInvalid(),
		"isDefined":  sc.IsDefined(),
		"isMatch":    sc.IsMatch(),
		"matchFail":  sc.IsMatchFailed(),
		"verifyErr":  sc.VerifyError() == nil,
	}

	// Assert
	expected := args.Map{
		"isInvalid":  true,
		"isDefined":  false,
		"isMatch":    true,
		"matchFail":  false,
		"verifyErr":  true,
	}
	expected.ShouldBeEqual(t, 0, "StringCompare_Nil returns nil -- with args", actual)
}

func Test_StringCompare_VerifyError_Regex(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareRegex(`\w+`, "abc")

	// Act
	actual := args.Map{"noErr": sc.VerifyError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_Regex returns error -- with args", actual)
}

func Test_StringCompare_VerifyError_NonRegex(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompare(stringcompareas.Equal, false, "abc", "abc")

	// Act
	actual := args.Map{"noErr": sc.VerifyError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_NonRegex returns error -- with args", actual)
}

// ── StringSearch ──

func Test_StringSearch_Methods(t *testing.T) {
	// Arrange
	var nilSS *coreinstruction.StringSearch

	// Act
	actual := args.Map{
		"nilIsEmpty": nilSS.IsEmpty(),
		"nilIsExist": nilSS.IsExist(),
		"nilHas":     nilSS.Has(),
		"nilMatch":   nilSS.IsMatch("anything"),
		"nilVerify":  nilSS.VerifyError("anything") == nil,
	}

	// Assert
	expected := args.Map{
		"nilIsEmpty": true,
		"nilIsExist": false,
		"nilHas":     false,
		"nilMatch":   true,
		"nilVerify":  true,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_Nil returns nil -- with args", actual)
}

func Test_StringSearch_IsAllMatch(t *testing.T) {
	// Arrange
	var nilSS *coreinstruction.StringSearch

	// Act
	actual := args.Map{
		"emptyMatch":   nilSS.IsAllMatch(),
		"anyMatchFail": nilSS.IsAnyMatchFailed("a"),
	}

	// Assert
	expected := args.Map{
		"emptyMatch":   true,
		"anyMatchFail": false,
	}
	expected.ShouldBeEqual(t, 0, "StringSearch_AllMatch returns correct value -- with args", actual)
}

// ── NameList / NameListCollection ──

func Test_NameList_DeepClone(t *testing.T) {
	// Arrange
	nl := &coreinstruction.NameList{Name: "test", List: corestr.New.SimpleSlice.Lines("a", "b")}
	cloned := nl.DeepClone()

	// Act
	actual := args.Map{
		"name": cloned.Name,
		"notNil": cloned.List != nil,
	}

	// Assert
	expected := args.Map{
		"name": "test",
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NameList_DeepClone returns correct value -- with args", actual)
}

func Test_NameList_String(t *testing.T) {
	// Arrange
	nl := coreinstruction.NameList{Name: "test"}

	// Act
	actual := args.Map{"notEmpty": nl.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameList_String returns correct value -- with args", actual)
}

func Test_NameListCollection(t *testing.T) {
	// Arrange
	var nilNlc *coreinstruction.NameListCollection
	nlc := &coreinstruction.NameListCollection{NameLists: []coreinstruction.NameList{{Name: "a"}}}

	// Act
	actual := args.Map{
		"nilIsNull":  nilNlc.IsNull(),
		"nilIsAny":   nilNlc.IsAnyNull(),
		"nilIsEmpty": nilNlc.IsEmpty(),
		"nilLength":  nilNlc.Length(),
		"length":     nlc.Length(),
		"hasAny":     nlc.HasAnyItem(),
		"notEmpty":   nlc.String() != "",
	}

	// Assert
	expected := args.Map{
		"nilIsNull":  true,
		"nilIsAny":   true,
		"nilIsEmpty": true,
		"nilLength":  0,
		"length":     1,
		"hasAny":     true,
		"notEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "NameListCollection returns correct value -- with args", actual)
}

// ── Identifiers ──

func Test_Identifiers_AddsAndLookup(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiers()
	ids.Add("a").Adds("b", "c")

	// Act
	actual := args.Map{
		"length":   ids.Length(),
		"hasAny":   ids.HasAnyItem(),
		"indexOf":  ids.IndexOf("b"),
		"notFound": ids.IndexOf("z"),
		"getById":  ids.GetById("a") != nil,
		"getNil":   ids.GetById("z") == nil,
	}

	// Assert
	expected := args.Map{
		"length":   3,
		"hasAny":   true,
		"indexOf":  1,
		"notFound": -1,
		"getById":  true,
		"getNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers returns correct value -- with args", actual)
}

func Test_Identifiers_EmptyOps(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiers()
	ids.Add("") // should skip

	// Act
	actual := args.Map{
		"isEmpty":  ids.IsEmpty(),
		"indexOf":  ids.IndexOf(""),
		"getEmpty": ids.GetById("") == nil,
	}

	// Assert
	expected := args.Map{
		"isEmpty":  true,
		"indexOf":  -1,
		"getEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Identifiers_Empty returns empty -- with args", actual)
}

func Test_Identifiers_Clone(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiers("a", "b")
	cloned := ids.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Identifiers_Clone returns correct value -- with args", actual)
}

func Test_Identifiers_NewUsingCap(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersUsingCap(5)

	// Act
	actual := args.Map{"len": ids.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewIdentifiersUsingCap returns correct value -- with args", actual)
}

// ── LineIdentifier ──

func Test_LineIdentifier_Methods(t *testing.T) {
	// Arrange
	var nilLi *coreinstruction.LineIdentifier

	// Act
	actual := args.Map{
		"nilInvalid":   nilLi.IsInvalidLineNumber(),
		"nilHasLine":   nilLi.HasLineNumber(),
		"nilNewLine":   nilLi.IsNewLineRequest(),
		"nilDelete":    nilLi.IsDeleteLineRequest(),
		"nilModify":    nilLi.IsModifyLineRequest(),
		"nilAddOrMod":  nilLi.IsAddNewOrModifyLineRequest(),
		"nilToBase":    nilLi.ToBaseLineIdentifier() == nil,
		"nilClone":     nilLi.Clone() == nil,
		"nilInvalidLn": nilLi.IsInvalidLineNumberUsingLastLineNumber(10),
	}

	// Assert
	expected := args.Map{
		"nilInvalid":   true,
		"nilHasLine":   false,
		"nilNewLine":   false,
		"nilDelete":    false,
		"nilModify":    false,
		"nilAddOrMod":  false,
		"nilToBase":    true,
		"nilClone":     true,
		"nilInvalidLn": true,
	}
	expected.ShouldBeEqual(t, 0, "LineIdentifier_Nil returns nil -- with args", actual)
}

// ── BaseTypeDotFilter ──

func Test_BaseTypeDotFilter(t *testing.T) {
	// Arrange
	f := &coreinstruction.BaseTypeDotFilter{}
	// Uses zero value — just trigger the function
	splits := f.GetDotSplitTypes()
	splits2 := f.GetDotSplitTypes() // cached

	// Act
	actual := args.Map{"same": len(splits) == len(splits2)}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "BaseTypeDotFilter returns correct value -- with args", actual)
}

// ── RequestSpecification ──

func Test_RequestSpecification_Clone(t *testing.T) {
	// Arrange
	rs := coreinstruction.RequestSpecification{}
	rs.Id = "x"
	rs.Tags = []string{"t1"}
	cloned := rs.Clone()

	// Act
	actual := args.Map{
		"id": cloned.Id,
		"tagsLen": len(cloned.Tags),
	}

	// Assert
	expected := args.Map{
		"id": "x",
		"tagsLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "RequestSpecification_Clone returns correct value -- with args", actual)
}
