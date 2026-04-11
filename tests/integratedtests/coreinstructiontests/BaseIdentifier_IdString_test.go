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
	// Arrange
	id := coreinstruction.NewIdentifier("test-id")

	// Act
	actual := args.Map{"result": id.IdString() != "test-id"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test-id', got ''", actual)
}

func Test_BaseIdentifier_IsIdEmpty(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("")

	// Act
	actual := args.Map{"result": id.IsIdEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty id should be empty", actual)
	id2 := coreinstruction.NewIdentifier("x")
	actual = args.Map{"result": id2.IsIdEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty should not be empty", actual)
}

func Test_BaseIdentifier_IsIdWhitespace(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("   ")

	// Act
	actual := args.Map{"result": id.IsIdWhitespace()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "whitespace id should be whitespace", actual)
}

func Test_BaseIdentifier_IsId(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("test")

	// Act
	actual := args.Map{"result": id.IsId("test")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match", actual)
	actual = args.Map{"result": id.IsId("other")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}

func Test_BaseIdentifier_IsIdCaseInsensitive(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("Test")

	// Act
	actual := args.Map{"result": id.IsIdCaseInsensitive("test")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match case insensitive", actual)
}

func Test_BaseIdentifier_IsIdContains(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("hello-world")

	// Act
	actual := args.Map{"result": id.IsIdContains("world")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain 'world'", actual)
}

func Test_BaseIdentifier_IsIdRegexMatches(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("test-123")
	re := regexp.MustCompile(`\d+`)

	// Act
	actual := args.Map{"result": id.IsIdRegexMatches(re)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match regex", actual)
}

func Test_BaseIdentifier_Clone(t *testing.T) {
	// Arrange
	id := coreinstruction.NewIdentifier("orig")
	cloned := id.Clone()

	// Act
	actual := args.Map{"result": cloned.IdString() != "orig"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should have same id", actual)
}

// ==========================================
// BaseDisplay
// ==========================================

func Test_BaseDisplay_IsDisplay(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")

	// Act
	actual := args.Map{"result": spec.IsDisplay("MyDisplay")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match display", actual)
}

func Test_BaseDisplay_IsDisplayCaseInsensitive(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")

	// Act
	actual := args.Map{"result": spec.IsDisplayCaseInsensitive("mydisplay")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match case insensitive", actual)
}

func Test_BaseDisplay_IsDisplayContains(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id", "MyDisplay", "type")

	// Act
	actual := args.Map{"result": spec.IsDisplayContains("Disp")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain 'Disp'", actual)
}

func Test_BaseDisplay_IsDisplayRegexMatches(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id", "display-123", "type")
	re := regexp.MustCompile(`\d+`)

	// Act
	actual := args.Map{"result": spec.IsDisplayRegexMatches(re)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match regex", actual)
}

// ==========================================
// BaseEnabler
// ==========================================

func Test_BaseEnabler_SetEnable(t *testing.T) {
	// Arrange
	e := &coreinstruction.BaseEnabler{}
	e.SetEnable()

	// Act
	actual := args.Map{"result": e.IsEnabled}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be enabled", actual)
}

func Test_BaseEnabler_SetDisable(t *testing.T) {
	// Arrange
	e := &coreinstruction.BaseEnabler{IsEnabled: true}
	e.SetDisable()

	// Act
	actual := args.Map{"result": e.IsEnabled}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be disabled", actual)
}

func Test_BaseEnabler_SetEnableVal(t *testing.T) {
	// Arrange
	e := &coreinstruction.BaseEnabler{}
	e.SetEnableVal(true)

	// Act
	actual := args.Map{"result": e.IsEnabled}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be enabled", actual)
	e.SetEnableVal(false)
	actual = args.Map{"result": e.IsEnabled}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be disabled", actual)
}

// ==========================================
// BaseFromTo
// ==========================================

func Test_BaseFromTo_Create(t *testing.T) {
	// Arrange
	ft := coreinstruction.NewBaseFromTo("src", "dst")

	// Act
	actual := args.Map{"result": ft.From != "src" || ft.To != "dst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from/to not set correctly", actual)
}

// ==========================================
// Specification
// ==========================================

func Test_Specification_Simple(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")

	// Act
	actual := args.Map{"result": spec.Id != "id1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'id1', got ''", actual)
	actual = args.Map{"result": spec.Display != "Display1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Display1'", actual)
	actual = args.Map{"result": spec.Type != "Type1"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Type1'", actual)
}

func Test_Specification_SimpleGlobal(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimpleGlobal("id1", "Display1", "Type1")

	// Act
	actual := args.Map{"result": spec.IsGlobal}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be global", actual)
}

func Test_Specification_Full(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1", "tag2"}, true)

	// Act
	actual := args.Map{"result": len(spec.Tags) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 tags", actual)
	actual = args.Map{"result": spec.IsGlobal}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be global", actual)
}

func Test_Specification_Clone(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecification("id1", "Display1", "Type1", []string{"tag1"}, true)
	cloned := spec.Clone()

	// Act
	actual := args.Map{"result": cloned.Id != "id1"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone id mismatch", actual)
	cloned.Tags[0] = "modified"
	actual = args.Map{"result": spec.Tags[0] == "modified"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should be independent", actual)
}

func Test_Specification_Clone_Nil(t *testing.T) {
	// Arrange
	var spec *coreinstruction.Specification
	cloned := spec.Clone()

	// Act
	actual := args.Map{"result": cloned != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Specification_FlatSpecification(t *testing.T) {
	// Arrange
	spec := coreinstruction.NewSpecificationSimple("id1", "Display1", "Type1")
	flat := spec.FlatSpecification()

	// Act
	actual := args.Map{"result": flat == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-nil flat spec", actual)
	// Second call should return cached
	flat2 := spec.FlatSpecification()
	actual = args.Map{"result": flat != flat2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return same cached instance", actual)
}

func Test_Specification_FlatSpecification_Nil(t *testing.T) {
	// Arrange
	var spec *coreinstruction.Specification
	flat := spec.FlatSpecification()

	// Act
	actual := args.Map{"result": flat != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil spec should return nil flat", actual)
}

// ==========================================
// Rename
// ==========================================

func Test_Rename_Properties(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: "new"}

	// Act
	actual := args.Map{"result": r.FromName() != "old" || r.ToName() != "new"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "properties mismatch", actual)
	actual = args.Map{"result": r.ExistingName() != "old" || r.NewName() != "new"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "alias properties mismatch", actual)
}

func Test_Rename_IsNull(t *testing.T) {
	// Arrange
	var r *coreinstruction.Rename

	// Act
	actual := args.Map{"result": r.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_Rename_IsExistingEmpty(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "", New: "new"}

	// Act
	actual := args.Map{"result": r.IsExistingEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty existing should be empty", actual)
}

func Test_Rename_IsNewEmpty(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: ""}

	// Act
	actual := args.Map{"result": r.IsNewEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty new should be empty", actual)
}

func Test_Rename_String(t *testing.T) {
	// Arrange
	r := coreinstruction.Rename{Existing: "old", New: "new"}
	s := r.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_Rename_SourceDestination(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	sd := r.SourceDestination()

	// Act
	actual := args.Map{"result": sd == nil || sd.Source != "old" || sd.Destination != "new"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "source destination conversion failed", actual)
}

func Test_Rename_SourceDestination_Nil(t *testing.T) {
	// Arrange
	var r *coreinstruction.Rename
	sd := r.SourceDestination()

	// Act
	actual := args.Map{"result": sd != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Rename_FromTo(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	ft := r.FromTo()

	// Act
	actual := args.Map{"result": ft == nil || ft.From != "old" || ft.To != "new"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from-to conversion failed", actual)
}

func Test_Rename_FromTo_Nil(t *testing.T) {
	// Arrange
	var r *coreinstruction.Rename
	ft := r.FromTo()

	// Act
	actual := args.Map{"result": ft != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_Rename_SetFromToName(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	r.SetFromName("newFrom")
	r.SetToName("newTo")

	// Act
	actual := args.Map{"result": r.Existing != "newFrom" || r.New != "newTo"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "set methods failed", actual)
}

func Test_Rename_SetFromToName_Nil(t *testing.T) {
	var r *coreinstruction.Rename
	r.SetFromName("x") // should not panic
	r.SetToName("y")   // should not panic
}

func Test_Rename_Clone(t *testing.T) {
	// Arrange
	r := &coreinstruction.Rename{Existing: "old", New: "new"}
	cloned := r.Clone()

	// Act
	actual := args.Map{"result": cloned.Existing != "old" || cloned.New != "new"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
}

func Test_Rename_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coreinstruction.Rename
	cloned := r.Clone()

	// Act
	actual := args.Map{"result": cloned != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// ==========================================
// SourceDestination
// ==========================================

func Test_SourceDestination_Properties(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}

	// Act
	actual := args.Map{"result": sd.FromName() != "src" || sd.ToName() != "dst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "properties mismatch", actual)
}

func Test_SourceDestination_IsNull(t *testing.T) {
	// Arrange
	var sd *coreinstruction.SourceDestination

	// Act
	actual := args.Map{"result": sd.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_SourceDestination_IsSourceEmpty(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: ""}

	// Act
	actual := args.Map{"result": sd.IsSourceEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty source should be empty", actual)
}

func Test_SourceDestination_IsDestinationEmpty(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Destination: ""}

	// Act
	actual := args.Map{"result": sd.IsDestinationEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty destination should be empty", actual)
}

func Test_SourceDestination_String(t *testing.T) {
	// Arrange
	sd := coreinstruction.SourceDestination{Source: "src", Destination: "dst"}

	// Act
	actual := args.Map{"result": sd.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty string", actual)
}

func Test_SourceDestination_SetFromToName(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	sd.SetFromName("newSrc")
	sd.SetToName("newDst")

	// Act
	actual := args.Map{"result": sd.Source != "newSrc" || sd.Destination != "newDst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "set methods failed", actual)
}

func Test_SourceDestination_SetFromToName_Nil(t *testing.T) {
	var sd *coreinstruction.SourceDestination
	sd.SetFromName("x")
	sd.SetToName("y")
}

func Test_SourceDestination_FromTo(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	ft := sd.FromTo()

	// Act
	actual := args.Map{"result": ft == nil || ft.From != "src" || ft.To != "dst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "from-to conversion failed", actual)
}

func Test_SourceDestination_FromTo_Nil(t *testing.T) {
	// Arrange
	var sd *coreinstruction.SourceDestination

	// Act
	actual := args.Map{"result": sd.FromTo() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SourceDestination_Rename(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	r := sd.Rename()

	// Act
	actual := args.Map{"result": r == nil || r.Existing != "src" || r.New != "dst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "rename conversion failed", actual)
}

func Test_SourceDestination_Rename_Nil(t *testing.T) {
	// Arrange
	var sd *coreinstruction.SourceDestination

	// Act
	actual := args.Map{"result": sd.Rename() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SourceDestination_Clone(t *testing.T) {
	// Arrange
	sd := &coreinstruction.SourceDestination{Source: "src", Destination: "dst"}
	cloned := sd.Clone()

	// Act
	actual := args.Map{"result": cloned.Source != "src" || cloned.Destination != "dst"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
}

func Test_SourceDestination_Clone_Nil(t *testing.T) {
	// Arrange
	var sd *coreinstruction.SourceDestination

	// Act
	actual := args.Map{"result": sd.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

// ==========================================
// NameList
// ==========================================

func Test_NameList_IsNull(t *testing.T) {
	// Arrange
	var nl *coreinstruction.NameList

	// Act
	actual := args.Map{"result": nl.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be null", actual)
}

func Test_NameList_IsAnyNull_Nil(t *testing.T) {
	// Arrange
	var nl *coreinstruction.NameList

	// Act
	actual := args.Map{"result": nl.IsAnyNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be any null", actual)
}

func Test_NameList_IsAnyNull_NilList(t *testing.T) {
	// Arrange
	nl := &coreinstruction.NameList{Name: "test"}

	// Act
	actual := args.Map{"result": nl.IsAnyNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil list should be any null", actual)
}

func Test_NameList_IsNameEmpty(t *testing.T) {
	// Arrange
	nl := &coreinstruction.NameList{Name: ""}

	// Act
	actual := args.Map{"result": nl.IsNameEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty name should be empty", actual)
}

func Test_NameList_HasName(t *testing.T) {
	// Arrange
	nl := &coreinstruction.NameList{Name: "test"}

	// Act
	actual := args.Map{"result": nl.HasName()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have name", actual)
}

func Test_NameList_Clone_Nil(t *testing.T) {
	// Arrange
	var nl *coreinstruction.NameList

	// Act
	actual := args.Map{"result": nl.Clone(true) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}
