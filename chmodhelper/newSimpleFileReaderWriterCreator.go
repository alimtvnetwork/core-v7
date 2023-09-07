package chmodhelper

import (
	"os"
	"path"
	"path/filepath"
)

type newSimpleFileReaderWriterCreator struct{}

// Create
//
// Arguments:
// 	- chmodDir     : applying on parentDir
// 	- chmodFile    : applying on file
// 	- absParentDir : absolute parentDir ( it can be two level before ),
// 	    it doesn't have to be relative to absFilePath but can be relative.
// 	- absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) Create(
	chmodDir,
	chmodFile os.FileMode,
	absParentDir,
	absFilePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              absParentDir,
		FilePath:               absFilePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// All
//
// Arguments:
// 	- chmodDir     : applying on parentDir
// 	- chmodFile    : applying on file
// 	- absParentDir : absolute parentDir ( it can be two level before ),
// 	    it doesn't have to be relative to absFilePath but can be relative.
// 	- absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) All(
	chmodDir,
	chmodFile os.FileMode,
	isApplyChmodMust bool,
	isApplyOnMismatch bool,
	absParentDir,
	absFilePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              absParentDir,
		FilePath:               absFilePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// CreateClean
//
//  Applies path.Clean() to relative actual path from relative(cmd/..) or (.) path
//  then create the reader, writer.
//
// Arguments:
// 	- chmodDir  : applying on parentDir
// 	- chmodFile : applying on file
// 	- parentDir : absolute parentDir ( it can be two level before ) after clean,
// 	    it doesn't have to be relative to absFilePath but can be relative.
// 	- filePath  : absolute file path after clean or else will not work
func (it newSimpleFileReaderWriterCreator) CreateClean(
	chmodDir,
	chmodFile os.FileMode,
	parentDir,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir = path.Clean(parentDir)
	filePath = path.Clean(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// Default
//
//  applies default chmod dir - 0755
//  (filemode.DirDefault), file - 0644  (filemode.FileDefault)
//
// Arguments:
// 	- chmodDir     : applying on parentDir
// 	- chmodFile    : applying on file
// 	- absParentDir : absolute parentDir will be from parent of absFilePath
// 	- absFilePath  : absolute file path
func (it newSimpleFileReaderWriterCreator) Default(
	absFilePath string,
) *SimpleFileReaderWriter {
	parentDir := filepath.Dir(absFilePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              fileDefaultChmod,
		ParentDir:              parentDir,
		FilePath:               absFilePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// DefaultCleanPath
//
//  Applies path.Clean() to relative actual path from relative(cmd/..) or (.) path
//  then create the reader, writer.
//
//  applies default chmod dir - 0755
//  (filemode.DirDefault), file - 0644  (filemode.FileDefault)
//
// Arguments:
// 	- chmodDir     : applying on parentDir
// 	- chmodFile    : applying on file
// 	- absFilePath  : absolute file path after clean.
// 	- absParentDir : absolute parentDir will be from parent of absFilePath after clean
func (it newSimpleFileReaderWriterCreator) DefaultCleanPath(
	filePath string,
) *SimpleFileReaderWriter {
	filePath = path.Clean(filePath)
	parentDir := filepath.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              fileDefaultChmod,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// Path
//
// Arguments:
// 	- chmodDir     : applying on parentDir
// 	- chmodFile    : applying on file
// 	- absFilePath  : absolute file path.
// 	- parentDir    : will be extracted from absFilePath.
func (it newSimpleFileReaderWriterCreator) Path(
	chmodDir,
	chmodFile os.FileMode,
	absFilePath string,
) *SimpleFileReaderWriter {
	parentDir := filepath.Dir(absFilePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               absFilePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

func (it newSimpleFileReaderWriterCreator) PathCondition(
	isApplyClean bool,
	chmodDir,
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	if isApplyClean {
		filePath = path.Clean(filePath)
	}

	parentDir := filepath.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               chmodDir,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

// PathDirDefaultChmod
//
//  dir will be applied with default chmod - 0755
func (it newSimpleFileReaderWriterCreator) PathDirDefaultChmod(
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := filepath.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:               dirDefaultChmod,
		ChmodFile:              chmodFile,
		ParentDir:              parentDir,
		FilePath:               filePath,
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}
