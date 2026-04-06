package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Conclusive_BothReflectNull(t *testing.T) {
	var p1 *string
	var p2 *string
	isEqual, isConclusive := isany.Conclusive(p1, p2)
	actual := args.Map{"result": isEqual || !isConclusive}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil ptr should be equal conclusive", actual)
}

func Test_Cov_Conclusive_OneReflectNull(t *testing.T) {
	var p1 *string
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(p1, &s)
	actual := args.Map{"result": isEqual || !isConclusive}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal but conclusive", actual)
}

func Test_Cov_Conclusive_DiffTypes(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(42, "hello")
	actual := args.Map{"result": isEqual || !isConclusive}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "diff types should be not equal but conclusive", actual)
}

func Test_Cov_Conclusive_Inconclusive(t *testing.T) {
	a := "hello"
	b := "world"
	isEqual, isConclusive := isany.Conclusive(&a, &b)
	actual := args.Map{"result": isEqual || isConclusive}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same type diff values should be inconclusive", actual)
}

func Test_Cov_JsonEqual_IntEqual(t *testing.T) {
	actual := args.Map{"result": isany.JsonEqual(42, 42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual := args.Map{"result": isany.JsonEqual(42, 43)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Cov_JsonEqual_JsonMarshal(t *testing.T) {
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 1}
	actual := args.Map{"result": isany.JsonEqual(a, b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}
