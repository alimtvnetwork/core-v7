package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RwxWrapper.LinuxApplyRecursive valid dir (line 328-345) ──

func Test_RwxWrapper_LinuxApplyRecursive_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "sub")
	_ = os.Mkdir(subDir, 0o755)
	_ = os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.LinuxApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive valid dir", expected)
}

// ── RwxWrapper.LinuxApplyRecursive invalid path, skip=false ──

func Test_RwxWrapper_LinuxApplyRecursive_InvalidNotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.LinuxApplyRecursive(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive invalid not skip", expected)
}

// ── RwxWrapper.LinuxApplyRecursive invalid path, skip=true ──

func Test_RwxWrapper_LinuxApplyRecursive_InvalidSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.LinuxApplyRecursive(true, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive invalid skip", expected)
}

// ── RwxWrapper.ApplyRecursive valid dir ──

func Test_RwxWrapper_ApplyRecursive_ValidDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.ApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive valid dir", expected)
}

// ── RwxWrapper.ApplyRecursive invalid path not skip ──

func Test_RwxWrapper_ApplyRecursive_InvalidNotSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	wrapper := mustRwxWrapper("rwxr-xr-x")

	// Act
	err := wrapper.ApplyRecursive(false, "/no/such/path/ever")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive invalid not skip", expected)
}

// ── RwxWrapper.ApplyRecursive on a single file ──

func Test_RwxWrapper_ApplyRecursive_SingleFile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}

	// Arrange
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "single.txt")
	_ = os.WriteFile(testFile, []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.ApplyRecursive(false, testFile)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "ApplyRecursive single file", expected)
}

// ── RwxWrapper.ApplyRecursive_Dir_CmdPath — linux cmd-based recursive chmod ──

func Test_RwxWrapper_ApplyRecursive_Dir_CmdPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}
	if runtime.GOOS != "linux" {
		t.Skip("linux only for cmd-based recursive chmod")
	}

	// Arrange
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "inner")
	_ = os.Mkdir(subDir, 0o755)
	_ = os.WriteFile(filepath.Join(subDir, "f.txt"), []byte("x"), 0o644)
	wrapper := mustRwxWrapper("rwxrwxrwx")

	// Act
	err := wrapper.LinuxApplyRecursive(false, tmpDir)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "LinuxApplyRecursive cmd path", expected)
}
