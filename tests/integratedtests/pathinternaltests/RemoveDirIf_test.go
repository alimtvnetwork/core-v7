package pathinternaltests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

// ── RemoveDirIf — true condition with existing dir ──

func Test_RemoveDirIf_ExistingDir(t *testing.T) {
	// Arrange
	dir := filepath.Join(os.TempDir(), "pathinternal_test_cov3_removedir")
	_ = os.MkdirAll(dir, 0755)

	// Act
	err := pathinternal.RemoveDirIf(true, dir, "test")

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"removed": !pathinternal.IsPathExists(dir),
	}
	expected := args.Map{
		"noErr": true,
		"removed": true,
	}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf returns correct value -- existing dir", actual)
}

// ── RemoveDirIfMust — no panic on non-existent ──

func Test_RemoveDirIfMust_NonExistent(t *testing.T) {
	// Act — should not panic
	pathinternal.RemoveDirIfMust(true, "/tmp/nonexistent_pathcov3", "test")

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIfMust returns non-empty -- non-existent", actual)
}

func Test_RemoveDirIfMust_FalseCondition(t *testing.T) {
	// Act
	pathinternal.RemoveDirIfMust(false, "/tmp/whatever", "test")

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIfMust returns non-empty -- false condition", actual)
}

// ── RemoveDirMust ──

func Test_RemoveDirMust(t *testing.T) {
	// Arrange
	dir := filepath.Join(os.TempDir(), "pathinternal_test_cov3_removemust")
	_ = os.MkdirAll(dir, 0755)

	// Act
	pathinternal.RemoveDirMust(dir, "test")

	// Assert
	actual := args.Map{"removed": !pathinternal.IsPathExists(dir)}
	expected := args.Map{"removed": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirMust returns correct value -- with args", actual)
}

// ── RemoveDirMustSimple ──

func Test_RemoveDirMustSimple(t *testing.T) {
	// Arrange
	dir := filepath.Join(os.TempDir(), "pathinternal_test_cov3_removesimple")
	_ = os.MkdirAll(dir, 0755)

	// Act
	pathinternal.RemoveDirMustSimple(dir)

	// Assert
	actual := args.Map{"removed": !pathinternal.IsPathExists(dir)}
	expected := args.Map{"removed": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirMustSimple returns correct value -- with args", actual)
}

// ── Clean with triple/quadruple slashes ──

func Test_Clean_MultiSlash(t *testing.T) {
	// Act
	actual := args.Map{
		"triple": pathinternal.Clean("/a///b"),
		"quad":   pathinternal.Clean("/a////b"),
	}

	// Assert
	expected := args.Map{
		"triple": "/a/b",
		"quad":   "/a/b",
	}
	expected.ShouldBeEqual(t, 0, "Clean returns correct value -- multi-slash", actual)
}
