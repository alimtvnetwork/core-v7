package pathextendinf

type PathLister interface {
	// ParentDir
	//
	//  Actions related to ParentDir
	ParentDir() DirFileLister
	// RootDir
	//
	//  Actions related to RootDir
	RootDir() DirFileLister
	// FilePath
	//
	//  Actions related to FilePath
	FilePath() DirFileLister
}
