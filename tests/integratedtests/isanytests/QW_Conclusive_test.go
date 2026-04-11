package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Conclusive_LeftNilRightDefined(t *testing.T) {
	// Arrange
	// Cover the branch: left==nil || right==nil (with left nil, right defined)
	isEqual, isConclusive := isany.Conclusive(nil, "hello")

	// Act
	actual := args.Map{"result": isEqual}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	actual = args.Map{"result": isConclusive}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected conclusive", actual)
}

func Test_QW_JsonEqual_BothMarshalError(t *testing.T) {
	// Arrange
	// Cover the branch where both marshal errors exist and are different
	ch1 := make(chan int)
	ch2 := make(chan string)
	result := isany.JsonEqual(ch1, ch2)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for different marshal errors", actual)
}
