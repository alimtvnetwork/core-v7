package chmodhelpertests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// covTempDir and covWriteFile are defined in shared_coverage_helpers.go

// ── tempDirGetter ──

func Test_TempDirGetter_TempDefault(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.TempDirGetter.TempDefault()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "TempDirGetter.TempDefault returns non-empty path",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TempDirGetter_TempPermanent_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.TempDirGetter.TempPermanent()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "TempDirGetter.TempPermanent returns non-empty path",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TempDirGetter_TempOption_Permanent(t *testing.T) {
	// Arrange / Act
	perm := chmodhelper.TempDirGetter.TempOption(true)
	nonPerm := chmodhelper.TempDirGetter.TempOption(false)

	// Assert
	actual := args.Map{
		"permNotEmpty":    fmt.Sprintf("%v", perm != ""),
		"nonPermNotEmpty": fmt.Sprintf("%v", nonPerm != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "TempDirGetter.TempOption returns paths for both options",
		ExpectedInput: args.Map{
			"permNotEmpty":    "true",
			"nonPermNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── IsPathExists / IsPathInvalid / IsDirectory ──

func Test_IsPathExists_Valid(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	result := chmodhelper.IsPathExists(dir)

	// Assert
	actual := args.Map{
		"exists": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsPathExists returns true for existing dir",
		ExpectedInput: args.Map{
			"exists": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IsPathExists_Invalid(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.IsPathExists("/nonexistent_path_xyz_12345")

	// Assert
	actual := args.Map{
		"exists": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsPathExists returns false for non-existing path",
		ExpectedInput: args.Map{
			"exists": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IsPathInvalid_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.IsPathInvalid("/nonexistent_path_xyz_12345")

	// Assert
	actual := args.Map{
		"invalid": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsPathInvalid returns true for non-existing path",
		ExpectedInput: args.Map{
			"invalid": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IsDirectory_Dir(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	result := chmodhelper.IsDirectory(dir)

	// Assert
	actual := args.Map{
		"isDir": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsDirectory returns true for directory",
		ExpectedInput: args.Map{
			"isDir": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IsDirectory_File(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "test.txt", "content")

	// Act
	result := chmodhelper.IsDirectory(filePath)

	// Assert
	actual := args.Map{
		"isDir": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsDirectory returns false for file",
		ExpectedInput: args.Map{
			"isDir": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IsDirectory_NonExist(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.IsDirectory("/nonexistent_12345")

	// Assert
	actual := args.Map{
		"isDir": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "IsDirectory returns false for non-existing path",
		ExpectedInput: args.Map{
			"isDir": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── ExpandCharRwx ──

func Test_ExpandCharRwx_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	r, w, x := chmodhelper.ExpandCharRwx("rwx")

	// Assert
	actual := args.Map{
		"r": string(r),
		"w": string(w),
		"x": string(x),
	}

	tc := coretestcases.CaseV1{
		Title:         "ExpandCharRwx returns r w x bytes -- rwx input",
		ExpectedInput: args.Map{
			"r": "r",
			"w": "w",
			"x": "x",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ExpandCharRwx_Hyphens(t *testing.T) {
	// Arrange / Act
	r, w, x := chmodhelper.ExpandCharRwx("r-x")

	// Assert
	actual := args.Map{
		"r": string(r),
		"w": string(w),
		"x": string(x),
	}

	tc := coretestcases.CaseV1{
		Title:         "ExpandCharRwx returns correct chars -- r-x input",
		ExpectedInput: args.Map{
			"r": "r",
			"w": "-",
			"x": "x",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── FileModeFriendlyString ──

func Test_FileModeFriendlyString_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.FileModeFriendlyString(0755)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "FileModeFriendlyString returns non-empty for 0755",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── GetPathExistStat ──

func Test_GetPathExistStat_ValidDir(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	stat := chmodhelper.GetPathExistStat(dir)

	// Assert
	actual := args.Map{
		"isExist":    fmt.Sprintf("%v", stat.IsExist),
		"isDir":      fmt.Sprintf("%v", stat.IsDir()),
		"isFile":     fmt.Sprintf("%v", stat.IsFile()),
		"isInvalid":  fmt.Sprintf("%v", stat.IsInvalid()),
		"hasIssues":  fmt.Sprintf("%v", stat.HasAnyIssues()),
		"hasError":   fmt.Sprintf("%v", stat.HasError()),
		"emptyError": fmt.Sprintf("%v", stat.IsEmptyError()),
		"hasInfo":    fmt.Sprintf("%v", stat.HasFileInfo()),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetPathExistStat returns valid stat for existing dir",
		ExpectedInput: args.Map{
			"isExist":    "true",
			"isDir":      "true",
			"isFile":     "false",
			"isInvalid":  "false",
			"hasIssues":  "false",
			"hasError":   "false",
			"emptyError": "true",
			"hasInfo":    "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_GetPathExistStat_ValidFile(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "stat.txt", "data")

	// Act
	stat := chmodhelper.GetPathExistStat(filePath)

	// Assert
	actual := args.Map{
		"isExist":     fmt.Sprintf("%v", stat.IsExist),
		"isFile":      fmt.Sprintf("%v", stat.IsFile()),
		"isDir":       fmt.Sprintf("%v", stat.IsDir()),
		"hasMode":     fmt.Sprintf("%v", stat.FileMode() != nil),
		"hasSize":     fmt.Sprintf("%v", stat.Size() != nil),
		"hasModified": fmt.Sprintf("%v", stat.LastModifiedDate() != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetPathExistStat returns valid stat for file",
		ExpectedInput: args.Map{
			"isExist":     "true",
			"isFile":      "true",
			"isDir":       "false",
			"hasMode":     "true",
			"hasSize":     "true",
			"hasModified": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_GetPathExistStat_NonExist(t *testing.T) {
	// Arrange / Act
	stat := chmodhelper.GetPathExistStat("/nonexistent_xyz_99")

	// Assert
	actual := args.Map{
		"isExist":   fmt.Sprintf("%v", stat.IsExist),
		"isInvalid": fmt.Sprintf("%v", stat.IsInvalid()),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetPathExistStat returns invalid stat for non-existing path",
		ExpectedInput: args.Map{
			"isExist":   "false",
			"isInvalid": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_Nil_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	var stat *chmodhelper.PathExistStat

	// Act / Assert
	actual := args.Map{
		"hasError":    fmt.Sprintf("%v", stat.HasError()),
		"emptyError":  fmt.Sprintf("%v", stat.IsEmptyError()),
		"hasInfo":     fmt.Sprintf("%v", stat.HasFileInfo()),
		"invalidInfo": fmt.Sprintf("%v", stat.IsInvalidFileInfo()),
		"isInvalid":   fmt.Sprintf("%v", stat.IsInvalid()),
		"hasIssues":   fmt.Sprintf("%v", stat.HasAnyIssues()),
		"isFile":      fmt.Sprintf("%v", stat.IsFile()),
		"isDir":       fmt.Sprintf("%v", stat.IsDir()),
		"nilModified": fmt.Sprintf("%v", stat.LastModifiedDate() == nil),
		"nilMode":     fmt.Sprintf("%v", stat.FileMode() == nil),
		"nilSize":     fmt.Sprintf("%v", stat.Size() == nil),
		"string":      stat.String(),
		"notExist":    fmt.Sprintf("%v", stat.NotExistError() == nil),
		"notFile":     fmt.Sprintf("%v", stat.NotAFileError() == nil),
		"notDir":      fmt.Sprintf("%v", stat.NotADirError() == nil),
		"meaningFul":  fmt.Sprintf("%v", stat.MeaningFullError() == nil),
		"msgWrapped":  stat.MessageWithPathWrapped("test"),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat nil receiver returns safe defaults",
		ExpectedInput: args.Map{
			"hasError":    "false",
			"emptyError":  "true",
			"hasInfo":     "false",
			"invalidInfo": "true",
			"isInvalid":   "true",
			"hasIssues":   "true",
			"isFile":      "false",
			"isDir":       "false",
			"nilModified": "true",
			"nilMode":     "true",
			"nilSize":     "true",
			"string":      "",
			"notExist":    "true",
			"notFile":     "true",
			"notDir":      "true",
			"meaningFul":  "true",
			"msgWrapped":  "",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_Split(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "split.txt", "x")
	stat := chmodhelper.GetPathExistStat(filePath)

	// Act
	_, fileName := stat.Split()
	parentDir := stat.ParentDir()

	// Assert
	actual := args.Map{
		"fileName":       fileName,
		"parentNotEmpty": fmt.Sprintf("%v", parentDir != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.Split returns dir and filename -- valid file",
		ExpectedInput: args.Map{
			"fileName":       "split.txt",
			"parentNotEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_Split_Dir(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)

	// Act
	dirResult, fileName := stat.Split()

	// Assert
	actual := args.Map{
		"dir":      dirResult,
		"fileName": fileName,
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.Split returns empty for directory",
		ExpectedInput: args.Map{
			"dir":      "",
			"fileName": "",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_DotExt(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "data.json", "{}")
	stat := chmodhelper.GetPathExistStat(filePath)

	// Act
	ext := stat.DotExt()

	// Assert
	actual := args.Map{
		"ext": ext,
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.DotExt returns extension -- .json file",
		ExpectedInput: args.Map{
			"ext": ".json",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_CombineWith(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)

	// Act
	combined := stat.CombineWithNewPath("sub", "file.txt")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", combined != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.CombineWithNewPath joins paths -- sub/file.txt",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_Dispose_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)

	// Act
	stat.Dispose()

	// Assert
	actual := args.Map{
		"location": stat.Location,
		"isExist":  fmt.Sprintf("%v", stat.IsExist),
		"nilInfo":  fmt.Sprintf("%v", stat.FileInfo == nil),
		"nilError": fmt.Sprintf("%v", stat.Error == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.Dispose clears all fields",
		ExpectedInput: args.Map{
			"location": "",
			"isExist":  "false",
			"nilInfo":  "true",
			"nilError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_NotExistError(t *testing.T) {
	// Arrange
	stat := chmodhelper.GetPathExistStat("/nonexistent_abc_777")

	// Act
	err := stat.NotExistError()

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.NotExistError returns error for non-existing -- path missing",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_NotAFileError_Dir_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)

	// Act
	err := stat.NotAFileError()

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.NotAFileError returns error for directory -- expecting file",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_NotADirError_File_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "f.txt", "x")
	stat := chmodhelper.GetPathExistStat(filePath)

	// Act
	err := stat.NotADirError()

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.NotADirError returns error for file -- expecting dir",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PathExistStat_String_Valid(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	stat := chmodhelper.GetPathExistStat(dir)

	// Act
	result := stat.String()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "PathExistStat.String returns formatted output -- valid dir",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── GetPathExistStatExpand ──

func Test_GetPathExistStatExpand_Valid(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	fileInfo, isExist, err := chmodhelper.GetPathExistStatExpand(dir)

	// Assert
	actual := args.Map{
		"isExist":     fmt.Sprintf("%v", isExist),
		"noError":     fmt.Sprintf("%v", err == nil),
		"hasFileInfo": fmt.Sprintf("%v", fileInfo != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "GetPathExistStatExpand returns valid info -- existing dir",
		ExpectedInput: args.Map{
			"isExist":     "true",
			"noError":     "true",
			"hasFileInfo": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── chmodVerifier ──

func Test_ChmodVerifier_GetRwxFull_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.GetRwxFull(0755)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
		"length":   fmt.Sprintf("%d", len(result)),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.GetRwxFull returns 10 char rwx string -- 0755",
		ExpectedInput: args.Map{
			"notEmpty": "true",
			"length":   "10",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_GetRwx9_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.GetRwx9(0755)

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
		"length":   fmt.Sprintf("%d", len(result)),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.GetRwx9 returns 9 char rwx string -- 0755",
		ExpectedInput: args.Map{
			"notEmpty": "true",
			"length":   "9",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_GetRwx9_Zero(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.GetRwx9(0)

	// Assert
	actual := args.Map{
		"result": result,
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.GetRwx9 returns empty for zero file mode",
		ExpectedInput: args.Map{
			"result": "---------",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_IsEqual_InvalidPath(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.IsEqual("/nonexistent_9999", 0755)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.IsEqual returns false for non-existing path",
		ExpectedInput: args.Map{
			"isEqual": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_IsEqualSkipInvalid_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.IsEqualSkipInvalid("/nonexistent_9999", 0755)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.IsEqualSkipInvalid returns true for invalid path",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_IsMismatch_InvalidPath(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.IsMismatch("/nonexistent_9999", 0755)

	// Assert
	actual := args.Map{
		"isMismatch": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.IsMismatch returns true for non-existing path",
		ExpectedInput: args.Map{
			"isMismatch": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_MismatchError_InvalidPath(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodVerify.MismatchError("/nonexistent_9999", 0755)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.MismatchError returns error for non-existing path",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_PathIf_SkipVerify(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodVerify.PathIf(false, "/nonexistent", 0755)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.PathIf returns nil when isVerify is false",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_RwxFull_InvalidLength_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodVerify.RwxFull("/tmp", "rwx")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.RwxFull returns error for invalid rwx length",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_IsEqualRwxFullSkipInvalid_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.ChmodVerify.IsEqualRwxFullSkipInvalid("/nonexistent", "-rwxrwxrwx")

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.IsEqualRwxFullSkipInvalid returns true on invalid path",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_MismatchErrorUsingRwxFull(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodVerify.MismatchErrorUsingRwxFull("/nonexistent", "-rwxrwxrwx")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.MismatchErrorUsingRwxFull returns error for non-existing",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_GetExisting_Invalid(t *testing.T) {
	// Arrange / Act
	_, err := chmodhelper.ChmodVerify.GetExisting("/nonexistent_chmod_test")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.GetExisting returns error for non-existing path",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_GetExistingRwxWrapper_Invalid(t *testing.T) {
	// Arrange / Act
	_, err := chmodhelper.ChmodVerify.GetExistingRwxWrapper("/nonexistent_wrapper_test")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.GetExistingRwxWrapper returns error for non-existing",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodVerifier_UsingRwxOwnerGroupOther_Nil_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodVerify.UsingRwxOwnerGroupOther(nil, "/tmp")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodVerify.UsingRwxOwnerGroupOther returns error for nil rwx",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── chmodApplier ──

func Test_ChmodApplier_ApplyIf_SkipApply(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodApply.ApplyIf(false, 0755, "/nonexistent")

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodApply.ApplyIf returns nil when isApply is false",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodApplier_OnMismatchOption_SkipApply(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodApply.OnMismatchOption(false, true, 0755, "/nonexistent")

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodApply.OnMismatchOption returns nil when isApply is false",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_EmptyLocations(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, &chmodins.Condition{})

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodApply.PathsUsingFileModeConditions returns nil for empty locations",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodApplier_PathsUsingFileModeConditions_NilCondition(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(0755, nil, "/tmp")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodApply.PathsUsingFileModeConditions returns error for nil condition",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ChmodApplier_RwxPartial_EmptyLocations(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.ChmodApply.RwxPartial("-rwx", &chmodins.Condition{})

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "ChmodApply.RwxPartial returns nil for empty locations",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxStringApplyChmod_EmptyLocations(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxStringApplyChmod("-rwxrwxrwx", &chmodins.Condition{})

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxStringApplyChmod returns nil for empty locations",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_EmptyLocations(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
		&chmodins.Condition{},
	)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxOwnerGroupOtherApplyChmod returns nil for empty locations",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilRwx_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(nil, &chmodins.Condition{}, "/tmp")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxOwnerGroupOtherApplyChmod returns error for nil rwx",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxOwnerGroupOtherApplyChmod_NilCondition_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxOwnerGroupOtherApplyChmod(
		&chmodins.RwxOwnerGroupOther{Owner: "rwx", Group: "r-x", Other: "r-x"},
		nil,
		"/tmp",
	)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxOwnerGroupOtherApplyChmod returns error for nil condition",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxStringApplyChmod_InvalidLength_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxStringApplyChmod("rwx", &chmodins.Condition{}, "/tmp")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxStringApplyChmod returns error for invalid rwx length",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_RwxStringApplyChmod_NilCondition_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.RwxStringApplyChmod("-rwxrwxrwx", nil, "/tmp")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "RwxStringApplyChmod returns error for nil condition",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── fileWriter (via SimpleFileWriter.FileWriter) ──

func Test_FileWriter_ParentDir_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	result := chmodhelper.SimpleFileWriter.FileWriter.ParentDir("/some/path/file.txt")

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileWriter.ParentDir returns parent directory",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileWriter_RemoveIf_False_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(false, "/some/path")

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileWriter.RemoveIf returns nil when isRemove is false",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileWriter_RemoveIf_NonExist_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(true, "/nonexistent_remove_test_12345")

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileWriter.RemoveIf returns nil for non-existing path",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileWriter_RemoveIf_Exists(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "remove_test.txt", "data")

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.RemoveIf(true, filePath)
	existsAfter := chmodhelper.IsPathExists(filePath)

	// Assert
	actual := args.Map{
		"noError":     fmt.Sprintf("%v", err == nil),
		"existsAfter": fmt.Sprintf("%v", existsAfter),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileWriter.RemoveIf removes existing file",
		ExpectedInput: args.Map{
			"noError":     "true",
			"existsAfter": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileWriter_WriteAndRead_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "write_test.txt")

	// Act
	err := chmodhelper.SimpleFileWriter.FileWriter.All(
		0755, 0644,
		false, false, false, true,
		dir, filePath, []byte("hello write"),
	)
	content, readErr := os.ReadFile(filePath)

	// Assert
	actual := args.Map{
		"writeNoErr": fmt.Sprintf("%v", err == nil),
		"readNoErr":  fmt.Sprintf("%v", readErr == nil),
		"content":    string(content),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileWriter.All writes file successfully -- basic write",
		ExpectedInput: args.Map{
			"writeNoErr": "true",
			"readNoErr":  "true",
			"content":    "hello write",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── fileReader (via SimpleFileWriter.FileReader) ──

func Test_FileReader_Read_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "read_test.txt", "read me")

	// Act
	content, err := chmodhelper.SimpleFileWriter.FileReader.Read(filePath)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"content": content,
	}

	tc := coretestcases.CaseV1{
		Title:         "fileReader.Read returns file content as string",
		ExpectedInput: args.Map{
			"noError": "true",
			"content": "read me",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileReader_ReadBytes_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "readbytes.txt", "bytes")

	// Act
	b, err := chmodhelper.SimpleFileWriter.FileReader.ReadBytes(filePath)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"content": string(b),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileReader.ReadBytes returns file content as bytes",
		ExpectedInput: args.Map{
			"noError": "true",
			"content": "bytes",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FileReader_Read_NonExist(t *testing.T) {
	// Arrange / Act
	_, err := chmodhelper.SimpleFileWriter.FileReader.Read("/nonexistent_reader_test")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "fileReader.Read returns error for non-existing file",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── simpleFileWriter Lock/Unlock ──

func Test_SimpleFileWriter_LockUnlock_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act - should not deadlock
	chmodhelper.SimpleFileWriter.Lock()
	chmodhelper.SimpleFileWriter.Unlock()

	// Assert
	actual := args.Map{
		"noPanic": "true",
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileWriter.Lock and Unlock do not deadlock",
		ExpectedInput: args.Map{
			"noPanic": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── dirCreator (via SimpleFileWriter.CreateDir) ──

func Test_DirCreator_If_False_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	err := chmodhelper.SimpleFileWriter.CreateDir.If(false, 0755, "/some/path")

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "dirCreator.If returns nil when isCreate is false",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DirCreator_IfMissing_Existing(t *testing.T) {
	// Arrange
	dir := covTempDir(t)

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.IfMissing(0755, dir)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "dirCreator.IfMissing returns nil for existing dir",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DirCreator_Default_NewDir(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	newDir := filepath.Join(dir, "newsubdir")

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.Default(0755, newDir)
	exists := chmodhelper.IsPathExists(newDir)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"exists":  fmt.Sprintf("%v", exists),
	}

	tc := coretestcases.CaseV1{
		Title:         "dirCreator.Default creates new directory",
		ExpectedInput: args.Map{
			"noError": "true",
			"exists":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DirCreator_Direct_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	newDir := filepath.Join(dir, "direct_sub")

	// Act
	err := chmodhelper.SimpleFileWriter.CreateDir.Direct(newDir)
	exists := chmodhelper.IsPathExists(newDir)

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"exists":  fmt.Sprintf("%v", exists),
	}

	tc := coretestcases.CaseV1{
		Title:         "dirCreator.Direct creates directory with default chmod",
		ExpectedInput: args.Map{
			"noError": "true",
			"exists":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SimpleFileReaderWriter ──

func Test_SimpleFileReaderWriter_Create(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "rw_test.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Assert
	actual := args.Map{
		"notNil":    fmt.Sprintf("%v", rw != nil),
		"parentDir": fmt.Sprintf("%v", rw.ParentDir == dir),
		"filePath":  fmt.Sprintf("%v", rw.FilePath == filePath),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.Create initializes correctly",
		ExpectedInput: args.Map{
			"notNil":    "true",
			"parentDir": "true",
			"filePath":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_Default(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "default_test.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Default(true, filePath)

	// Assert
	actual := args.Map{
		"notNil":       fmt.Sprintf("%v", rw != nil),
		"hasParentDir": fmt.Sprintf("%v", rw.ParentDir != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.Default extracts parent from path",
		ExpectedInput: args.Map{
			"notNil":       "true",
			"hasParentDir": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_WriteAndReadString(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "writeread.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	writeErr := rw.WriteString("hello rw")
	content, readErr := rw.ReadString()

	// Assert
	actual := args.Map{
		"writeNoErr": fmt.Sprintf("%v", writeErr == nil),
		"readNoErr":  fmt.Sprintf("%v", readErr == nil),
		"content":    content,
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.WriteString and ReadString work correctly",
		ExpectedInput: args.Map{
			"writeNoErr": "true",
			"readNoErr":  "true",
			"content":    "hello rw",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_WriteAndRead_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "bytes.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	writeErr := rw.Write([]byte("byte content"))
	content, readErr := rw.Read()

	// Assert
	actual := args.Map{
		"writeNoErr": fmt.Sprintf("%v", writeErr == nil),
		"readNoErr":  fmt.Sprintf("%v", readErr == nil),
		"content":    string(content),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.Write and Read work for bytes",
		ExpectedInput: args.Map{
			"writeNoErr": "true",
			"readNoErr":  "true",
			"content":    "byte content",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_IsExist_IsPathInvalid(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "exist_check.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	existBefore := rw.IsExist()
	invalidBefore := rw.IsPathInvalid()
	parentExists := rw.IsParentExist()

	// Assert
	actual := args.Map{
		"existBefore":   fmt.Sprintf("%v", existBefore),
		"invalidBefore": fmt.Sprintf("%v", invalidBefore),
		"parentExists":  fmt.Sprintf("%v", parentExists),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter path checks before write",
		ExpectedInput: args.Map{
			"existBefore":   "false",
			"invalidBefore": "true",
			"parentExists":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_Clone_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "clone.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	cloned := rw.Clone()

	// Assert
	actual := args.Map{
		"sameDir":  fmt.Sprintf("%v", cloned.ParentDir == rw.ParentDir),
		"samePath": fmt.Sprintf("%v", cloned.FilePath == rw.FilePath),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.Clone creates deep copy",
		ExpectedInput: args.Map{
			"sameDir":  "true",
			"samePath": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_ClonePtr_Nil_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	var rw *chmodhelper.SimpleFileReaderWriter

	// Act
	result := rw.ClonePtr()

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.ClonePtr returns nil for nil receiver",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_String_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "str.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	result := rw.String()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", result != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.String returns formatted output",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_JoinRelPath_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filepath.Join(dir, "f.txt"),
	)

	// Act
	result := rw.JoinRelPath("sub/file.txt")
	emptyResult := rw.JoinRelPath("")

	// Assert
	actual := args.Map{
		"notEmpty":     fmt.Sprintf("%v", result != ""),
		"emptyIsClean": fmt.Sprintf("%v", emptyResult != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.JoinRelPath joins paths correctly",
		ExpectedInput: args.Map{
			"notEmpty":     "true",
			"emptyIsClean": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_Expire_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := covWriteFile(t, dir, "expire.txt", "data")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	err := rw.Expire()
	existsAfter := rw.IsExist()

	// Assert
	actual := args.Map{
		"noError":     fmt.Sprintf("%v", err == nil),
		"existsAfter": fmt.Sprintf("%v", existsAfter),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.Expire removes existing file",
		ExpectedInput: args.Map{
			"noError":     "true",
			"existsAfter": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_Expire_NonExist(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filepath.Join(dir, "nonexist.txt"),
	)

	// Act
	err := rw.Expire()

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.Expire returns nil for non-existing file",
		ExpectedInput: args.Map{
			"noError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_ReadOnExist_NonExist(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filepath.Join(dir, "missing.txt"),
	)

	// Act
	bytes, err := rw.ReadOnExist()

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"isNil":   fmt.Sprintf("%v", bytes == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.ReadOnExist returns nil for non-existing",
		ExpectedInput: args.Map{
			"noError": "true",
			"isNil":   "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_ReadStringOnExist_NonExist(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filepath.Join(dir, "missing.txt"),
	)

	// Act
	content, err := rw.ReadStringOnExist()

	// Assert
	actual := args.Map{
		"noError": fmt.Sprintf("%v", err == nil),
		"content": content,
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.ReadStringOnExist returns empty for non-existing",
		ExpectedInput: args.Map{
			"noError": "true",
			"content": "",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_HasPathIssues(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filepath.Join(dir, "missing.txt"),
	)

	// Act
	result := rw.HasPathIssues()
	hasAny := rw.HasAnyIssues()

	// Assert
	actual := args.Map{
		"hasPathIssues": fmt.Sprintf("%v", result),
		"hasAnyIssues":  fmt.Sprintf("%v", hasAny),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter path issue detection -- missing file",
		ExpectedInput: args.Map{
			"hasPathIssues": "true",
			"hasAnyIssues":  "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_InitializeDefault_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "init.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  filePath,
	}

	// Act
	initialized := rw.InitializeDefault(true)

	// Assert
	actual := args.Map{
		"notNil":       fmt.Sprintf("%v", initialized != nil),
		"hasParentDir": fmt.Sprintf("%v", initialized.ParentDir != ""),
		"mustChmod":    fmt.Sprintf("%v", initialized.IsMustChmodApplyOnFile),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.InitializeDefault sets parentDir and flags",
		ExpectedInput: args.Map{
			"notNil":       "true",
			"hasParentDir": "true",
			"mustChmod":    "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_InitializeDefaultApplyChmod_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "init2.txt")
	rw := chmodhelper.SimpleFileReaderWriter{
		ChmodDir:  0755,
		ChmodFile: 0644,
		FilePath:  filePath,
	}

	// Act
	initialized := rw.InitializeDefaultApplyChmod()

	// Assert
	actual := args.Map{
		"notNil":    fmt.Sprintf("%v", initialized != nil),
		"mustChmod": fmt.Sprintf("%v", initialized.IsMustChmodApplyOnFile),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.InitializeDefaultApplyChmod sets mustChmod true",
		ExpectedInput: args.Map{
			"notNil":    "true",
			"mustChmod": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SimpleFileReaderWriter_Json_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "json.txt")
	rw := chmodhelper.New.SimpleFileReaderWriter.Create(
		true, 0755, 0644, dir, filePath,
	)

	// Act
	jsonResult := rw.Json()

	// Assert
	actual := args.Map{
		"notEmpty": fmt.Sprintf("%v", len(jsonResult.Bytes) > 0),
	}

	tc := coretestcases.CaseV1{
		Title:         "SimpleFileReaderWriter.Json returns non-empty result",
		ExpectedInput: args.Map{
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── newSimpleFileReaderWriterCreator variants ──

func Test_NewSimpleFileReaderWriter_All(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "all.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.All(
		0755, 0644, true, true, true, dir, filePath,
	)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", rw != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.All creates with all params",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_Options(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "options.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Options(true, true, true, filePath)

	// Assert
	actual := args.Map{
		"notNil":       fmt.Sprintf("%v", rw != nil),
		"hasParentDir": fmt.Sprintf("%v", rw.ParentDir != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.Options creates with defaults",
		ExpectedInput: args.Map{
			"notNil":       "true",
			"hasParentDir": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_Path(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "path.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.Path(true, 0755, 0644, filePath)

	// Assert
	actual := args.Map{
		"notNil":       fmt.Sprintf("%v", rw != nil),
		"hasParentDir": fmt.Sprintf("%v", rw.ParentDir != ""),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.Path extracts parent from path",
		ExpectedInput: args.Map{
			"notNil":       "true",
			"hasParentDir": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_DefaultCleanPath(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "clean.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.DefaultCleanPath(true, filePath)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", rw != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.DefaultCleanPath cleans and creates",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_PathCondition(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "condition.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathCondition(true, true, 0755, 0644, filePath)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", rw != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.PathCondition creates with clean option",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_PathDirDefaultChmod(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "dirdefault.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.PathDirDefaultChmod(true, 0644, filePath)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", rw != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.PathDirDefaultChmod creates with dir default",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewSimpleFileReaderWriter_CreateClean(t *testing.T) {
	// Arrange
	dir := covTempDir(t)
	filePath := filepath.Join(dir, "createclean.txt")

	// Act
	rw := chmodhelper.New.SimpleFileReaderWriter.CreateClean(true, 0755, 0644, dir, filePath)

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", rw != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.SimpleFileReaderWriter.CreateClean cleans paths and creates",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── Variant ──

func Test_Variant_String_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	str := v.String()
	r, w, x := v.ExpandOctalByte()

	// Assert
	actual := args.Map{
		"string": str,
		"r":      string(r),
		"w":      string(w),
		"x":      string(x),
	}

	tc := coretestcases.CaseV1{
		Title:         "Variant.String and ExpandOctalByte work for X755",
		ExpectedInput: args.Map{
			"string": "755",
			"r":      "7",
			"w":      "5",
			"x":      "5",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Variant_ToWrapper_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	v := chmodhelper.X777

	// Act
	wrapper, err := v.ToWrapper()

	// Assert
	actual := args.Map{
		"noError":  fmt.Sprintf("%v", err == nil),
		"notEmpty": fmt.Sprintf("%v", wrapper.Owner.IsRead),
	}

	tc := coretestcases.CaseV1{
		Title:         "Variant.ToWrapper creates valid RwxWrapper -- X777",
		ExpectedInput: args.Map{
			"noError":  "true",
			"notEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── newRwxWrapperCreator ──

func Test_NewRwxWrapper_Invalid_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	wrapper := chmodhelper.New.RwxWrapper.Invalid()

	// Assert
	actual := args.Map{
		"ownerRead": fmt.Sprintf("%v", wrapper.Owner.IsRead),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.Invalid returns empty wrapper",
		ExpectedInput: args.Map{
			"ownerRead": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_InvalidPtr_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	wrapper := chmodhelper.New.RwxWrapper.InvalidPtr()

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", wrapper != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.InvalidPtr returns non-nil empty pointer",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_Empty_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	wrapper := chmodhelper.New.RwxWrapper.Empty()

	// Assert
	actual := args.Map{
		"notNil": fmt.Sprintf("%v", wrapper != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.Empty returns non-nil empty pointer",
		ExpectedInput: args.Map{
			"notNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_UsingFileMode_Zero_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(0)

	// Assert
	actual := args.Map{
		"ownerRead": fmt.Sprintf("%v", wrapper.Owner.IsRead),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.UsingFileMode returns empty for zero mode",
		ExpectedInput: args.Map{
			"ownerRead": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_UsingFileModePtr_Zero_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange / Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileModePtr(0)

	// Assert
	actual := args.Map{
		"notNil":    fmt.Sprintf("%v", wrapper != nil),
		"ownerRead": fmt.Sprintf("%v", wrapper.Owner.IsRead),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.UsingFileModePtr returns empty for zero mode",
		ExpectedInput: args.Map{
			"notNil":    "true",
			"ownerRead": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_RwxFullString_InvalidLength(t *testing.T) {
	// Arrange / Act
	_, err := chmodhelper.New.RwxWrapper.RwxFullString("rwx")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.RwxFullString returns error for invalid length",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_NewRwxWrapper_RwxFullStringWtHyphen_InvalidLength(t *testing.T) {
	// Arrange / Act
	_, err := chmodhelper.New.RwxWrapper.RwxFullStringWtHyphen("rwx")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "New.RwxWrapper.RwxFullStringWtHyphen returns error for invalid length",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── VarAttribute ──

func Test_VarAttribute_Nil_Clone(t *testing.T) {
	// Arrange
	var va *chmodhelper.VarAttribute

	// Act
	result := va.Clone()

	// Assert
	actual := args.Map{
		"isNil": fmt.Sprintf("%v", result == nil),
	}

	tc := coretestcases.CaseV1{
		Title:         "VarAttribute.Clone returns nil for nil receiver",
		ExpectedInput: args.Map{
			"isNil": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_VarAttribute_IsEqualPtr_BothNil_FromTempDirGetterFileWri(t *testing.T) {
	// Arrange
	var a, b *chmodhelper.VarAttribute

	// Act
	result := a.IsEqualPtr(b)

	// Assert
	actual := args.Map{
		"isEqual": fmt.Sprintf("%v", result),
	}

	tc := coretestcases.CaseV1{
		Title:         "VarAttribute.IsEqualPtr returns true for both nil",
		ExpectedInput: args.Map{
			"isEqual": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}
