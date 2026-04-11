package chmodhelpertests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage17 — exported API gaps for chmodhelper
// ══════════════════════════════════════════════════════════════════════════════

// ── RwxWrapper ───────────────────────────────────────────────────────────────

func Test_RwxWrapper_VerifyPaths_InvalidPaths(t *testing.T) {
	// Arrange
	wrapper, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.X755)
	actual := args.Map{"result": err}
	expected := args.Map{"result": nil}
	expected.ShouldBeEqual(t, 0, "err", actual)

	// Act
	verifyErr := wrapper.VerifyPaths(false, "/nonexistent/path1")

	// Assert
	convey.Convey("VerifyPaths returns error for non-existent paths", t, func() {
		convey.So(verifyErr, convey.ShouldNotBeNil)
	})
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Nil_FromRwxWrapperVerifyAndA(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.X755)

	// Act
	result := wrapper.IsRwxEqualFileInfo(nil)

	// Assert
	convey.Convey("IsRwxEqualFileInfo returns false for nil fileInfo", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_RwxWrapper_IsRwxEqualLocation_NonExistent_FromRwxWrapperVerifyAndA(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.X755)

	// Act
	result := wrapper.IsRwxEqualLocation("/nonexistent/xyz")

	// Assert
	convey.Convey("IsRwxEqualLocation returns false for non-existent path", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

// ── chmodVerifier ────────────────────────────────────────────────────────────

func Test_ChmodVerify_GetRwx9_Valid(t *testing.T) {
	// Arrange
	mode := os.FileMode(0755)

	// Act
	result := chmodhelper.ChmodVerify.GetRwx9(mode)

	// Assert
	convey.Convey("GetRwx9 returns 9-char string for valid mode", t, func() {
		convey.So(len(result), convey.ShouldEqual, 9)
	})
}

func Test_ChmodVerify_PathIf_VerifyTrue(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support Unix file permissions")
	}

	// Arrange — valid file
	dir := t.TempDir()
	fp := filepath.Join(dir, "test.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := chmodhelper.ChmodVerify.PathIf(true, fp, 0644)

	// Assert
	convey.Convey("PathIf verifies when isVerify is true", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_ChmodVerify_PathsUsingFileModeImmediateReturn(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support Unix file permissions")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "a.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileModeImmediateReturn(0644, fp)

	// Assert
	convey.Convey("PathsUsingFileModeImmediateReturn works on valid file", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_ChmodVerify_PathsUsingFileModeContinueOnError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support Unix file permissions")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "b.txt")
	os.WriteFile(fp, []byte("y"), 0644)

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileModeContinueOnError(0644, fp)

	// Assert
	convey.Convey("PathsUsingFileModeContinueOnError works on valid file", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_ChmodVerify_PathsUsingFileMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support Unix file permissions")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "c.txt")
	os.WriteFile(fp, []byte("z"), 0644)

	// Act
	err := chmodhelper.ChmodVerify.PathsUsingFileMode(false, 0644, fp)

	// Assert
	convey.Convey("PathsUsingFileMode works on valid file", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_ChmodVerify_PathsUsingPartialRwxOptions_InvalidPath(t *testing.T) {
	// Arrange & Act
	err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
		false, true, "-rwxr-xr-x", "/nonexistent/xyz")

	// Assert
	convey.Convey("PathsUsingPartialRwxOptions skips invalid when flag set", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func Test_ChmodVerify_PathsUsingRwxFull_ImmediateReturn_Error(t *testing.T) {
	// Arrange & Act — non-existent path, isContinueOnError=false
	err := chmodhelper.ChmodVerify.PathsUsingRwxFull(false, "-rwxr-xr-x", "/nonexistent/xyz")

	// Assert
	convey.Convey("PathsUsingRwxFull returns error for non-existent path", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── RwxVariableWrapper ───────────────────────────────────────────────────────

func Test_RwxVariableWrapper_IsEqualPartialRwxPartial(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwx******")

	// Act
	result := wrapper.IsEqualPartialRwxPartial("-rwxr-xr-x")

	// Assert
	convey.Convey("IsEqualPartialRwxPartial matches wildcard owner", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

func Test_RwxVariableWrapper_IsEqualUsingLocation_NonExistent_FromRwxWrapperVerifyAndA(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxrwxrwx")

	// Act
	result := wrapper.IsEqualUsingLocation("/nonexistent/xyz")

	// Assert
	convey.Convey("IsEqualUsingLocation returns false for non-existent", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_RwxVariableWrapper_IsEqualUsingFileInfo_Nil_FromRwxWrapperVerifyAndA(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.NewRwxVariableWrapper("-rwxrwxrwx")

	// Act
	result := wrapper.IsEqualUsingFileInfo(nil)

	// Assert
	convey.Convey("IsEqualUsingFileInfo returns false for nil", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

// ── SingleRwx ────────────────────────────────────────────────────────────────

func Test_SingleRwx_ToDisabledRwxWrapper_Valid(t *testing.T) {
	// Arrange
	single := &chmodhelper.SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.All,
	}

	// Act
	wrapper, err := single.ToDisabledRwxWrapper()

	// Assert
	convey.Convey("ToDisabledRwxWrapper returns wrapper for All class", t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(wrapper, convey.ShouldNotBeNil)
	})
}

func Test_SingleRwx_ToRwxWrapper_NonAll_FromRwxWrapperVerifyAndA(t *testing.T) {
	// Arrange
	single := &chmodhelper.SingleRwx{
		Rwx:       "rwx",
		ClassType: chmodclasstype.Owner,
	}

	// Act
	_, err := single.ToRwxWrapper()

	// Assert
	convey.Convey("ToRwxWrapper returns error for non-All class", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── CreateDirFilesWithRwxPermissions ─────────────────────────────────────────

func Test_CreateDirFilesWithRwxPermissions_Error(t *testing.T) {
	// Arrange — invalid dir path
	perms := []chmodhelper.DirFilesWithRwxPermission{
		{
			DirWithFiles: chmodhelper.DirWithFiles{
				Dir:   string([]byte{0}) + "/impossible",
				Files: []string{},
			},
			ApplyRwx: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "r-x",
			},
		},
	}

	// Act
	err := chmodhelper.CreateDirFilesWithRwxPermissions(false, perms)

	// Assert
	convey.Convey("CreateDirFilesWithRwxPermissions returns error for bad path", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── MergeRwxWildcardWithFixedRwx ─────────────────────────────────────────────

func Test_MergeRwxWildcardWithFixedRwx_ParseError(t *testing.T) {
	// Arrange — wildcard input with wrong length
	_, err := chmodhelper.MergeRwxWildcardWithFixedRwx("x", "rwx")

	// Assert
	convey.Convey("MergeRwxWildcardWithFixedRwx returns error for bad input", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── SimpleFileReaderWriter ───────────────────────────────────────────────────

func Test_SimpleFileReaderWriter_WriteRelativePath_Error(t *testing.T) {
	// Arrange
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		ParentDir: string([]byte{0}),
		FilePath:  string([]byte{0}) + "/file.txt",
	}

	// Act
	err := rw.WriteRelativePath(false, "sub/file.txt", []byte("data"))

	// Assert
	convey.Convey("WriteRelativePath returns error for invalid path", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// NOTE: getOnExist is unexported — tested in-package via Coverage6_UnexportedGaps_test.go

// ── fwChmodVerifier ──────────────────────────────────────────────────────────

func Test_FwChmodVerifier_IsEqualFile_FromRwxWrapperVerifyAndA(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not support Unix file permissions")
	}

	// Arrange
	dir := t.TempDir()
	fp := filepath.Join(dir, "test.txt")
	os.WriteFile(fp, []byte("x"), 0644)

	rw := &chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  fp,
		ParentDir: dir,
	}

	// Act
	verifier := rw.ChmodVerifier()
	result := verifier.IsEqualFile()

	// Assert
	convey.Convey("IsEqualFile returns true for matching chmod", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// ── PathExistStat ────────────────────────────────────────────────────────────

func Test_PathExistStat_MeaningfulError_HasError(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent/xyz")

	// Act
	err := stat.MeaningFullError()

	// Assert
	convey.Convey("MeaningFullError returns error for non-existent path", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── RwxPartialToInstructionExecutor ──────────────────────────────────────────

func Test_RwxPartialToInstructionExecutor_ParseError(t *testing.T) {
	// Arrange & Act — nil condition triggers CannotBeNilOrEmpty error
	_, err := chmodhelper.RwxPartialToInstructionExecutor(
		"-rwxr-xr-x",
		nil,
	)

	// Assert
	convey.Convey("RwxPartialToInstructionExecutor returns error for nil condition", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── CreateDirWithFiles ───────────────────────────────────────────────────────

func Test_CreateDirWithFiles_FileCloseError_ChmodError(t *testing.T) {
	// Arrange — valid dir, create files, verify chmod path
	dir := t.TempDir()
	subDir := filepath.Join(dir, "sub")
	dw := &chmodhelper.DirWithFiles{
		Dir:   subDir,
		Files: []string{"a.txt"},
	}

	// Act
	err := chmodhelper.CreateDirWithFiles(false, 0644, dw)

	// Assert
	convey.Convey("CreateDirWithFiles succeeds with valid files", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// ── DirFilesWithContent ──────────────────────────────────────────────────────

func Test_DirFilesWithContent_Create_RemoveDirError(t *testing.T) {
	// Arrange — invalid dir path
	dfc := &chmodhelper.DirFilesWithContent{
		Dir:         string([]byte{0}) + "/impossible",
		DirFileMode: 0755,
		Files:       []chmodhelper.FileWithContent{},
	}

	// Act
	err := dfc.Create(true)

	// Assert
	convey.Convey("Create returns error for invalid dir path", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── newRwxWrapperCreator ─────────────────────────────────────────────────────

func Test_NewRwxWrapper_UsingVariantPtr_InvalidDigit(t *testing.T) {
	// Arrange & Act — digit > 7 returns error
	_, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("899"))

	// Assert
	convey.Convey("UsingVariantPtr returns error for invalid octal digit", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// ── tempDirGetter ────────────────────────────────────────────────────────────

func Test_TempDirGetter_TempOption_NonPermanent(t *testing.T) {
	// Arrange & Act
	result := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	convey.Convey("TempOption(false) returns non-empty temp dir", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
	})
}
