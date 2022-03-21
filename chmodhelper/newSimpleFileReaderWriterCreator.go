package chmodhelper

import (
	"os"
	"path"
	"path/filepath"
)

type newSimpleFileReaderWriterCreator struct{}

func (it newSimpleFileReaderWriterCreator) Create(
	chmodDir,
	chmodFile os.FileMode,
	parentDir,
	filePath string,
) *SimpleFileReaderWriter {
	return &SimpleFileReaderWriter{
		ChmodDir:  chmodDir,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

func (it newSimpleFileReaderWriterCreator) CreateClean(
	chmodDir,
	chmodFile os.FileMode,
	parentDir,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir = path.Clean(parentDir)
	filePath = filepath.Clean(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  chmodDir,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

// Default
//
//  Default chmod dir - 0755, file - 0644
func (it newSimpleFileReaderWriterCreator) Default(
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := path.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  dirDefaultChmod,
		ChmodFile: fileDefaultChmod,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

func (it newSimpleFileReaderWriterCreator) DefaultCleanPath(
	filePath string,
) *SimpleFileReaderWriter {
	filePath = path.Clean(filePath)
	parentDir := path.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  dirDefaultChmod,
		ChmodFile: fileDefaultChmod,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

// Path
//
//  Default chmod dir - 0755, file - 0644
func (it newSimpleFileReaderWriterCreator) Path(
	chmodDir,
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := path.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  chmodDir,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

// PathCondition
//
//  Default chmod dir - 0755, file - 0644
func (it newSimpleFileReaderWriterCreator) PathCondition(
	isApplyClean bool,
	chmodDir,
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	if isApplyClean {
		filePath = path.Clean(filePath)
	}

	parentDir := path.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  chmodDir,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}

// PathDirDefaultChmod
//
//  dir will be applied with default chmod - 0755
func (it newSimpleFileReaderWriterCreator) PathDirDefaultChmod(
	chmodFile os.FileMode,
	filePath string,
) *SimpleFileReaderWriter {
	parentDir := path.Dir(filePath)

	return &SimpleFileReaderWriter{
		ChmodDir:  dirDefaultChmod,
		ChmodFile: chmodFile,
		ParentDir: parentDir,
		FilePath:  filePath,
	}
}
