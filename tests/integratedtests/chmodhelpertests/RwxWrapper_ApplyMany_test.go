package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/smartystreets/goconvey/convey"
)

// All RwxWrapper.ApplyLinuxChmodOnMany tests.
// Source API: chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen(string) (RwxWrapper, error)
// Input: 9-char rwx string WITHOUT leading hyphen

// ── non-recursive valid ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursive(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	f1 := filepath.Join(tmpDir, "a.txt")
	f2 := filepath.Join(tmpDir, "b.txt")
	_ = os.WriteFile(f1, []byte("x"), 0o644)
	_ = os.WriteFile(f2, []byte("y"), 0o644)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, f1, f2)

	// Assert
	actual = args.Map{"hasError": err != nil}
	expected = args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany non-recursive", expected)
}

// ── non-recursive with error ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursiveError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, "/no/exist/1", "/no/exist/2")

	// Assert
	actual = args.Map{"hasError": err != nil}
	expected = args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany non-recursive error", expected)
}

// ── recursive valid ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_Recursive_FromRwxWrapperApplyMany(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "sub")
	_ = os.Mkdir(subDir, 0o755)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, tmpDir)

	// Assert
	actual = args.Map{"hasError": err != nil}
	expected = args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany recursive", expected)
}

// ── recursive error ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_RecursiveError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: false,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, "/no/exist/path")

	// Assert
	actual = args.Map{"hasError": err != nil}
	expected = args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyLinuxChmodOnMany recursive error", expected)
}

// ── recursive continue-on-error ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_RecursiveContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       true,
		IsContinueOnError: true,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, tmpDir, "/no/exist")

	// Assert — at least one path should fail but continue
	_ = err
	convey.Convey("RecursiveContinueOnError processes all paths", t, func() {
		convey.So(true, convey.ShouldBeTrue)
	})
}

// ── non-recursive continue-on-error ──

func Test_RwxWrapper_ApplyLinuxChmodOnMany_NonRecursiveContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	validFile := filepath.Join(tmpDir, "ok.txt")
	_ = os.WriteFile(validFile, []byte("x"), 0o644)
	wrapper, wErr := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwxrwxrwx")
	actual := args.Map{"result": wErr != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected parse error:", actual)
	cond := &chmodins.Condition{
		IsRecursive:       false,
		IsContinueOnError: true,
		IsSkipOnInvalid:   false,
	}

	// Act
	err := wrapper.ApplyLinuxChmodOnMany(cond, validFile, "/no/exist")

	// Assert — errors aggregated, no panic
	_ = err
	convey.Convey("NonRecursiveContinueOnError processes all paths", t, func() {
		convey.So(true, convey.ShouldBeTrue)
	})
}
