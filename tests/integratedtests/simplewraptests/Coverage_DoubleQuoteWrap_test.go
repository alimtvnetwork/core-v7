package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_DoubleQuoteWrapElements_SkipOnExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, "hello", "\"already\"")
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_DoubleQuoteWrapElements_NoSkip(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, "hello")
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_DoubleQuoteWrapElements_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Cov_DoubleQuoteWrapElementsWithIndexes(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}
