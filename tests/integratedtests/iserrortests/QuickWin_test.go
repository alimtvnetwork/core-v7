package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Equal_BothNonNilSameMessage(t *testing.T) {
	// Arrange
	// Cover the Error() comparison branch
	e1 := errors.New("same")
	e2 := errors.New("same")

	// Act
	actual := args.Map{"result": iserror.Equal(e1, e2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for same message", actual)
}

func Test_QW_Equal_BothNonNilDiffMessage(t *testing.T) {
	// Arrange
	e1 := errors.New("a")
	e2 := errors.New("b")

	// Act
	actual := args.Map{"result": iserror.Equal(e1, e2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for different message", actual)
}

func Test_QW_Equal_LeftNilRightNot(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Equal(nil, errors.New("a"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
