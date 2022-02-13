package internalpathextender

// PathInfoer
//
// FullPath:
//  - Refers to the FullPath (dir or filename through absolute path).
// FileName:
//  - Refers to the file name at the end
// DirName:
//  - Refers to the file name at the end
// Name:
//  - Refers to the end of name could be file or dir.
// Extension:
//  - Refers to dot extension (.db, .back etc)
// RootDir:
//  - Refers to root dir where things started from.
// Relative:
//  - Refers to relative from RootDir
// ParentDir:
//  - Refers to parent dir of the FullPath and different from RootDir
type PathInfoer interface {
	FullPath() string
	FileName() string
	DirName() string
	Name() string
	// Extension
	//
	//  Refers to dot extension always
	Extension() string
	// RootDir
	//
	//  Refers to start of the dir
	//  For example a repo start point
	RootDir() string
	// Relative
	//
	//  Refers to relative path from root
	//  Joining Root + Relative should give absolute or FullPath()
	Relative() string
	// ParentDir
	//
	//  Refers to current full-path's parent dir.
	//  ParentDir is different from RootDir.
	ParentDir() string
	// String
	//
	//  returns full summary info
	String() string

	Size() uint64
}
