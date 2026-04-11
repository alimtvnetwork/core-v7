package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Equal_BothSameMessage(t *testing.T) {
	// Arrange
	a := errors.New("err")
	b := errors.New("err")

	// Act
	actual := args.Map{"result": iserror.Equal(a, b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal by message", actual)
}

func Test_Equal_DiffMessage(t *testing.T) {
	// Arrange
	a := errors.New("a")
	b := errors.New("b")

	// Act
	actual := args.Map{"result": iserror.Equal(a, b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}
