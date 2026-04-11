package fsinternaltests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/fsinternal"
)

// ── IsPathExists ──

func Test_IsPathExists_Valid(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()

	// Act
	actual := args.Map{"result": fsinternal.IsPathExists(tmpDir)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns true -- existing dir", actual)
}

func Test_IsPathExists_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"result": fsinternal.IsPathExists("/nonexistent/path/xyz123")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns false -- nonexistent path", actual)
}

func Test_IsPathExists_File(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("hello"), 0644)

	// Act
	actual := args.Map{"result": fsinternal.IsPathExists(tmpFile)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns true -- existing file", actual)
}

// ── IsPathInvalid ──

func Test_IsPathInvalid_Valid(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()

	// Act
	actual := args.Map{"result": fsinternal.IsPathInvalid(tmpDir)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsPathInvalid returns false -- existing dir", actual)
}

func Test_IsPathInvalid_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"result": fsinternal.IsPathInvalid("/nonexistent/path/xyz123")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsPathInvalid returns true -- nonexistent path", actual)
}

// ── IsDirectory ──

func Test_IsDirectory_Dir(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()

	// Act
	actual := args.Map{"result": fsinternal.IsDirectory(tmpDir)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDirectory returns true -- existing dir", actual)
}

func Test_IsDirectory_File(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("hello"), 0644)

	// Act
	actual := args.Map{"result": fsinternal.IsDirectory(tmpFile)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDirectory returns false -- file not dir", actual)
}

func Test_IsDirectory_Nonexistent(t *testing.T) {
	// Act
	actual := args.Map{"result": fsinternal.IsDirectory("/nonexistent/path/xyz123")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDirectory returns false -- nonexistent path", actual)
}

// ── GetPathExistStat ──

func Test_GetPathExistStat_Exists(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	info, isExist, err := fsinternal.GetPathExistStat(tmpDir)

	// Act
	actual := args.Map{
		"isExist":  isExist,
		"infoNil":  info == nil,
		"errIsNil": err == nil,
		"isDir":    info != nil && info.IsDir(),
	}

	// Assert
	expected := args.Map{
		"isExist":  true,
		"infoNil":  false,
		"errIsNil": true,
		"isDir":    true,
	}
	expected.ShouldBeEqual(t, 0, "GetPathExistStat returns valid info -- existing dir", actual)
}

func Test_GetPathExistStat_NotExists(t *testing.T) {
	// Arrange
	_, isExist, _ := fsinternal.GetPathExistStat("/nonexistent/path/xyz123")

	// Act
	actual := args.Map{"isExist": isExist}

	// Assert
	expected := args.Map{"isExist": false}
	expected.ShouldBeEqual(t, 0, "GetPathExistStat returns false -- nonexistent path", actual)
}

func Test_GetPathExistStat_File(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(tmpFile, []byte("hello"), 0644)
	info, isExist, err := fsinternal.GetPathExistStat(tmpFile)

	// Act
	actual := args.Map{
		"isExist":  isExist,
		"errIsNil": err == nil,
		"isDir":    info.IsDir(),
		"size":     info.Size() > 0,
	}

	// Assert
	expected := args.Map{
		"isExist":  true,
		"errIsNil": true,
		"isDir":    false,
		"size":     true,
	}
	expected.ShouldBeEqual(t, 0, "GetPathExistStat returns file info -- existing file", actual)
}
