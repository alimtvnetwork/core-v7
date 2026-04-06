package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Equal_BothNonNilSameMessage(t *testing.T) {
	// Cover the Error() comparison branch
	e1 := errors.New("same")
	e2 := errors.New("same")
	actual := args.Map{"result": iserror.Equal(e1, e2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for same message", actual)
}

func Test_QW_Equal_BothNonNilDiffMessage(t *testing.T) {
	e1 := errors.New("a")
	e2 := errors.New("b")
	actual := args.Map{"result": iserror.Equal(e1, e2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for different message", actual)
}

func Test_QW_Equal_LeftNilRightNot(t *testing.T) {
	actual := args.Map{"result": iserror.Equal(nil, errors.New("a"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
