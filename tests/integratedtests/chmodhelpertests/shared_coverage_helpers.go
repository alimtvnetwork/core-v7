package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
)

// covTempDir creates a temporary directory for test use.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covTempDir(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

// covWriteFile writes a file in a directory and returns the full path.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covWriteFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	err := os.WriteFile(p, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

// newTestRW creates a SimpleFileReaderWriter for testing.
// Moved here from Coverage12_SimpleFileRW_test.go so split-recovery subfolders can see it.
func newTestRW(dir, file string) chmodhelper.SimpleFileReaderWriter {
	return chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               filepath.Join(dir, file),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}
