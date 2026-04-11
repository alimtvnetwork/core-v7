package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// All tests below use mustRwxWrapper() from Coverage19_helpers.go
// Source API: chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen(string) (RwxWrapper, error)
// Input: 9-char rwx string WITHOUT leading hyphen (e.g., "rwxr-xr-x")

// ── RwxWrapper.ApplyChmod on valid file (lines 227-255) ──

func Test_RwxWrapper_ApplyChmod_ValidFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(testFile, []byte("data"), 0o644)
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(false, testFile)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmod valid file", expected)
}

// ── RwxWrapper.ApplyChmod on invalid path (lines 237-239) ──

func Test_RwxWrapper_ApplyChmod_InvalidPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyChmod invalid path", expected)
}

// ── RwxWrapper.ApplyChmod skipOnInvalid with invalid path ──

func Test_RwxWrapper_ApplyChmod_SkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmod(true, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmod skip invalid", expected)
}

// ── RwxWrapper.ApplyChmodSkipInvalid (line 304-307) ──

func Test_RwxWrapper_ApplyChmodSkipInvalid(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyChmodSkipInvalid("/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyChmodSkipInvalid", expected)
}

// ── RwxWrapper: invalidPathErr (line 86-93) ──

func Test_RwxWrapper_ApplyChmod_InvalidPath_NotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	applyErr := wrapper.ApplyChmod(false, "/nonexistent/path/file.txt")

	// Assert
	actual := args.Map{"hasError": applyErr != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyChmod invalidPathErr", expected)
}

// ── RwxWrapper.IsEqualVarWrapper (line 579-588) ──

func Test_RwxWrapper_IsEqualVarWrapper_Nil_FromRwxWrapperApplyChmod(t *testing.T) {
	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	result := wrapper.IsEqualVarWrapper(nil)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	actual.ShouldBeEqual(t, 1, "IsEqualVarWrapper nil", expected)
}
