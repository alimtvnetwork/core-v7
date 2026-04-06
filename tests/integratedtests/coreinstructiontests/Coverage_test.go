package coreinstructiontests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// BaseIdentifier
// ==========================================

func Test_BaseIdentifier_IdString(t *testing.T) {
	id := coreinstruction.NewIdentifier("test-id")
	actual := args.Map{"result": id.IdString() != "test-id"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test-id', got ''", actual)
}

func Test_BaseIdentifier_IsIdEmpty(t *testing.T) {
	id := coreinstruction.NewIdentifier("")
	actual := args.Map{"result": id.IsIdEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty id should be empty", actual)
	id2 := coreinstruction.NewIdentifier("x")
	actual := args.Map{"result": id2.IsIdEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty should not be empty", actual)
}

func Test_BaseIdentifier_IsIdWhitespace(t *testing.T) {
	id := coreinstruction.NewIdentifier("   ")
	actual := args.Map{"result": id.IsIdWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "whitespace id should be whitespace", actual)
}

func Test_BaseIdentifier_IsId(t *testing.T) {
	id := coreinstruction.NewIdentifier("test")
	actual := args.Map{"result": id.IsId("test")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match", actual)
	actual := args.Map{"result": id.IsId("other")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}

func Test_BaseIdentifier_IsIdCaseInsensitive(t *testing.T) {
	id := coreinstruction.NewIdentifier("Test")
	actual := args.Map{"result": id.IsIdCaseInsensitive("test")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match case insensitive", actual)
}

func Test_BaseIdentifier_IsIdContains(t *testing.T) {
	id := coreinstruction.NewIdentifier("hello-world")
	actual := args.Map{"result": id.IsIdContains("world")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain 'world'", actual)
}

func Test_BaseIdentifier_IsIdRegexMatches(t *testing.T) {
	id := coreinstruction.NewIdentifier("test-123")
	re := regexp.MustCompile(`\d+`)
	actual := args.Map{"result": id.IsIdRegexMatches(re)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match regex", actual)
}

func Test_BaseIdentifier_Clone(t *testing.T) {
	id := coreinstruction.NewIdentifier("orig")
	cloned := id.Clone()
	actual := args.Map{"result": cloned.IdString() != "orig"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should have same id", actual)
}

// ==========================================
// BaseDisplay
// ==========================================

func Test_BaseDisplay_IsDisplay(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	actual := args.Map{"result": spec.IsDisplay("MyDisplay")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match display", actual)
}

func Test_BaseDisplay_IsDisplayCaseInsensitive(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	actual := args.Map{"result": spec.IsDisplayCaseInsensitive("mydisplay")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match case insensitive", actual)
}

func Test_BaseDisplay_IsDisplayContains(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")
	actual := args.Map{"result": spec.IsDisplayContains("Disp")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain 'Disp'", actual)
}

func Test_BaseDisplay_IsDisplayRegexMatches(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id", "display-123", "type")
	re := regexp.MustCompile(`\d+`)
	actual := args.Map{"result": spec.IsDisplayRegexMatches(re)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match regex", actual)
}

// ==========================================
// BaseEnabler
// ==========================================

func Test_BaseEnabler_SetEnable(t *testing.T) {
	e := &coreinstruction.BaseEnabler{}
	e.SetEnable()
	actual := args.Map{"result": e.IsEnabled}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be enabled", actual)
}

func Test_BaseEnabler_SetDisable(t *testing.T) {
	e := &coreinstruction.BaseEnabler{IsEnabled: true}
	e.SetDisable()
	actual := args.Map{"result": e.IsEnabled}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be disabled", actual)
}

func Test_BaseEnabler_SetEnableVal(t *testing.T) {
	e := &coreinstruction.BaseEnabler{}
	e.SetEnableVal(true)
	actual := args.Map{"result": e.IsEnabled}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be enabled", actual)
	e.SetEnableVal(false)
	actual := args.Map{"result": e.IsEnabled}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be disabled", actual)
}

// ==========================================
// BaseFromTo
// ==========================================

func Test_BaseFromTo_Create(t *testing.T) {
	ft := coreinstruction.NewBaseFromTo("src", "dst")
	actual := args.Map{"result": ft.From != "src" || ft.To != "dst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from/to not set correctly", actual)
}

// ==========================================
// Specification
// ==========================================

func Test_Specification_Simple(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")
	actual := args.Map{"result": spec.Id != "id1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'id1', got ''", actual)
	actual := args.Map{"result": spec.Display != "Display1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Display1'", actual)
	actual := args.Map{"result": spec.Type != "Type1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Type1'", actual)
}

func Test_Specification_SimpleGlobal(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimpleGlobal("id1", "Display1", "Type1")
	actual := args.Map{"result": spec.IsGlobal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be global", actual)
}

func Test_Specification_Full(t *testing.T) {
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1", "tag2"}, true)
	actual := args.Map{"result": len(spec.Tags) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 tags", actual)
	actual := args.Map{"result": spec.IsGlobal}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be global", actual)
}

func Test_Specification_Clone(t *testing.T) {
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1"}, true)
	cloned := spec.Clone()
	actual := args.Map{"result": cloned.Id != "id1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone id mismatch", actual)
	cloned.Tags[0] = "modified"
	actual := args.Map{"result": spec.Tags[0] == "modified"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should be independent", actual)
}

func Test_Specification_Clone_Nil(t *testing.T) {
	var spec *coreinstruction.Specification
	cloned := spec.Clone()
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Specification_FlatSpecification(t *testing.T) {
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")
	flat := spec.FlatSpecification()
	actual := args.Map{"result": flat == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil flat spec", actual)
	// Second call should return cached
	flat2 := spec.FlatSpecification()
	actual := args.Map{"result": flat != flat2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return same cached instance", actual)
}

func Test_Specification_FlatSpecification_Nil(t *testing.T) {
	var spec *coreinstruction.Specification
	flat := spec.FlatSpecification()
	actual := args.Map{"result": flat != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil spec should return nil flat", actual)
}

// ==========================================
// Rename
// ==========================================

func Test_Rename_Properties(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	actual := args.Map{"result": r.FromName() != "old" || r.ToName() != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "properties mismatch", actual)
	actual := args.Map{"result": r.ExistingName() != "old" || r.NewName() != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "alias properties mismatch", actual)
}

func Test_Rename_IsNull(t *testing.T) {
	var r *coreinstruction.Rename
	actual := args.Map{"result": r.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_Rename_IsExistingEmpty(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "", New: "new"}
	actual := args.Map{"result": r.IsExistingEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty existing should be empty", actual)
}

func Test_Rename_IsNewEmpty(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: ""}
	actual := args.Map{"result": r.IsNewEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty new should be empty", actual)
}

func Test_Rename_String(t *testing.T) {
	r := coreinstruction.Rename{Existing: "old", New: "new"}
	s := r.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_Rename_SourceDestination(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	sd := r.SourceDestination()
	actual := args.Map{"result": sd == nil || sd.Source != "old" || sd.Destination != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "source destination conversion failed", actual)
}

func Test_Rename_SourceDestination_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	sd := r.SourceDestination()
	actual := args.Map{"result": sd != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Rename_FromTo(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	ft := r.FromTo()
	actual := args.Map{"result": ft == nil || ft.From != "old" || ft.To != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from-to conversion failed", actual)
}

func Test_Rename_FromTo_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	ft := r.FromTo()
	actual := args.Map{"result": ft != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Rename_SetFromToName(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	r.SetFromName("newFrom")
	r.SetToName("newTo")
	actual := args.Map{"result": r.Existing != "newFrom" || r.New != "newTo"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "set methods failed", actual)
}

func Test_Rename_SetFromToName_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	r.SetFromName("x") // should not panic
	r.SetToName("y")   // should not panic
}

func Test_Rename_Clone(t *testing.T) {
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	cloned := r.Clone()
	actual := args.Map{"result": cloned.Existing != "old" || cloned.New != "new"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
}

func Test_Rename_Clone_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	cloned := r.Clone()
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// ==========================================
// SourceDestination
// ==========================================

func Test_SourceDestination_Properties(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	actual := args.Map{"result": sd.FromName() != "src" || sd.ToName() != "dst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "properties mismatch", actual)
}

func Test_SourceDestination_IsNull(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	actual := args.Map{"result": sd.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_SourceDestination_IsSourceEmpty(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: ""}
	actual := args.Map{"result": sd.IsSourceEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty source should be empty", actual)
}

func Test_SourceDestination_IsDestinationEmpty(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Destination: ""}
	actual := args.Map{"result": sd.IsDestinationEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty destination should be empty", actual)
}

func Test_SourceDestination_String(t *testing.T) {
	sd := coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	actual := args.Map{"result": sd.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_SourceDestination_SetFromToName(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	sd.SetFromName("newSrc")
	sd.SetToName("newDst")
	actual := args.Map{"result": sd.Source != "newSrc" || sd.Destination != "newDst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "set methods failed", actual)
}

func Test_SourceDestination_SetFromToName_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	sd.SetFromName("x")
	sd.SetToName("y")
}

func Test_SourceDestination_FromTo(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	ft := sd.FromTo()
	actual := args.Map{"result": ft == nil || ft.From != "src" || ft.To != "dst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from-to conversion failed", actual)
}

func Test_SourceDestination_FromTo_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	actual := args.Map{"result": sd.FromTo() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SourceDestination_Rename(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	r := sd.Rename()
	actual := args.Map{"result": r == nil || r.Existing != "src" || r.New != "dst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "rename conversion failed", actual)
}

func Test_SourceDestination_Rename_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	actual := args.Map{"result": sd.Rename() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SourceDestination_Clone(t *testing.T) {
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	cloned := sd.Clone()
	actual := args.Map{"result": cloned.Source != "src" || cloned.Destination != "dst"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
}

func Test_SourceDestination_Clone_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	actual := args.Map{"result": sd.Clone() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// ==========================================
// NameList
// ==========================================

func Test_NameList_IsNull(t *testing.T) {
	var nl *coreinstruction.NameList
	actual := args.Map{"result": nl.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_NameList_IsAnyNull_Nil(t *testing.T) {
	var nl *coreinstruction.NameList
	actual := args.Map{"result": nl.IsAnyNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be any null", actual)
}

func Test_NameList_IsAnyNull_NilList(t *testing.T) {
	nl := &coreinstruction.NameList{Name: "test"}
	actual := args.Map{"result": nl.IsAnyNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil list should be any null", actual)
}

func Test_NameList_IsNameEmpty(t *testing.T) {
	nl := &coreinstruction.NameList{Name: ""}
	actual := args.Map{"result": nl.IsNameEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty name should be empty", actual)
}

func Test_NameList_HasName(t *testing.T) {
	nl := &coreinstruction.NameList{Name: "test"}
	actual := args.Map{"result": nl.HasName()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have name", actual)
}

func Test_NameList_Clone_Nil(t *testing.T) {
	var nl *coreinstruction.NameList
	actual := args.Map{"result": nl.Clone(true) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}
