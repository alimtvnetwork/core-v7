package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Conclusive_BothReflectNull(t *testing.T) {
	// Arrange
	var p1 *string
	var p2 *string
	isEqual, isConclusive := isany.Conclusive(p1, p2)

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil ptr should be equal conclusive", actual)
}

func Test_Conclusive_OneReflectNull(t *testing.T) {
	// Arrange
	var p1 *string
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(p1, &s)

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal but conclusive", actual)
}

func Test_Conclusive_DiffTypes_FromConclusiveBothReflec(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(42, "hello")

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "diff types should be not equal but conclusive", actual)
}

func Test_Conclusive_Inconclusive_FromConclusiveBothReflec(t *testing.T) {
	// Arrange
	a := "hello"
	b := "world"
	isEqual, isConclusive := isany.Conclusive(&a, &b)

	// Act
	actual := args.Map{"result": isEqual || isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same type diff values should be inconclusive", actual)
}

func Test_JsonEqual_IntEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual(42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": isany.JsonEqual(42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_JsonEqual_JsonMarshal(t *testing.T) {
	// Arrange
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 1}

	// Act
	actual := args.Map{"result": isany.JsonEqual(a, b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}
